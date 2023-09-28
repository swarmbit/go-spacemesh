package config

import (
	"math"
	"math/big"
	"os"
	"path/filepath"
	"runtime"
	"time"

	"go.uber.org/zap/zapcore"

	"github.com/spacemeshos/go-spacemesh/activation"
	"github.com/spacemeshos/go-spacemesh/api/grpcserver"
	"github.com/spacemeshos/go-spacemesh/beacon"
	"github.com/spacemeshos/go-spacemesh/bootstrap"
	"github.com/spacemeshos/go-spacemesh/checkpoint"
	"github.com/spacemeshos/go-spacemesh/datastore"
	"github.com/spacemeshos/go-spacemesh/fetch"
	hareConfig "github.com/spacemeshos/go-spacemesh/hare/config"
	eligConfig "github.com/spacemeshos/go-spacemesh/hare/eligibility/config"
	"github.com/spacemeshos/go-spacemesh/nats"
	"github.com/spacemeshos/go-spacemesh/p2p"
	"github.com/spacemeshos/go-spacemesh/syncer"
	timeConfig "github.com/spacemeshos/go-spacemesh/timesync/config"
	"github.com/spacemeshos/go-spacemesh/tortoise"
)

func MainnetConfig() Config {
	var postPowDifficulty activation.PowDifficulty
	if err := postPowDifficulty.UnmarshalText([]byte("000dfb23b0979b4b000000000000000000000000000000000000000000000000")); err != nil {
		panic(err)
	}
	p2pconfig := p2p.DefaultConfig()

	p2pconfig.Bootnodes = []string{
		"/dns4/mainnet-bootnode-0.spacemesh.network/tcp/5000/p2p/12D3KooWPStnitMbLyWAGr32gHmPr538mT658Thp6zTUujZt3LRf",
		"/dns4/mainnet-bootnode-2.spacemesh.network/tcp/5000/p2p/12D3KooWAsMgXLpyGdsRNjHBF3FaXwnXhyMEqWQYBXUpvCHNzFNK",
		"/dns4/mainnet-bootnode-4.spacemesh.network/tcp/5000/p2p/12D3KooWRcTWDHzptnhJn5h6CtwnokzzMaDLcXv6oM9CxQEXd5FL",
		"/dns4/mainnet-bootnode-6.spacemesh.network/tcp/5000/p2p/12D3KooWRS47KAs3ZLkBtE2AqjJCwxRYqZKmyLkvombJJdrca8Hz",
		"/dns4/mainnet-bootnode-8.spacemesh.network/tcp/5000/p2p/12D3KooWFYv99aGbtXnZQy6UZxyf72NpkWJp3K4HS8Py35WhKtzE",
		"/dns4/mainnet-bootnode-10.spacemesh.network/tcp/5000/p2p/12D3KooWHK5m83sNj2eNMJMGAngcS9gBja27ho83t79Q2CD4iRjQ",
		"/dns4/mainnet-bootnode-12.spacemesh.network/tcp/5000/p2p/12D3KooWG4gk8GtMsAjYxHtbNC7oEoBTMRLbLDpKgSQMQkYBFRsw",
		"/dns4/mainnet-bootnode-14.spacemesh.network/tcp/5000/p2p/12D3KooWRkZMjGNrQfRyeKQC9U58cUwAfyQMtjNsupixkBFag8AY",
		"/dns4/mainnet-bootnode-16.spacemesh.network/tcp/5000/p2p/12D3KooWDAFRuFrMNgVQMDy8cgD71GLtPyYyfQzFxMZr2yUBgjHK",
		"/dns4/mainnet-bootnode-18.spacemesh.network/tcp/5000/p2p/12D3KooWMJmdfwxDctuGGoTYJD8Wj9jubQBbPfrgrzzXaQ1RTKE6",
	}

	smeshing := DefaultSmeshingConfig()
	smeshing.ProvingOpts.Nonces = 288
	smeshing.ProvingOpts.Threads = uint(runtime.NumCPU() * 3 / 4)
	if smeshing.ProvingOpts.Threads < 1 {
		smeshing.ProvingOpts.Threads = 1
	}
	logging := DefaultLoggingConfig()
	logging.TrtlLoggerLevel = zapcore.WarnLevel.String()
	return Config{
		BaseConfig: BaseConfig{
			DataDirParent:         defaultDataDir,
			FileLock:              filepath.Join(os.TempDir(), "spacemesh.lock"),
			MetricsPort:           1010,
			DatabaseConnections:   16,
			DatabasePruneInterval: 30 * time.Minute,
			NetworkHRP:            "sm",

			LayerDuration:  5 * time.Minute,
			LayerAvgSize:   50,
			LayersPerEpoch: 4032,

			TxsPerProposal: 700,       // https://github.com/spacemeshos/go-spacemesh/issues/4559
			BlockGasLimit:  100107000, // 3000 of spends

			OptFilterThreshold: 90,

			TickSize: 9331200,
			PoETServers: []string{
				"https://mainnet-poet-0.spacemesh.network",
				"https://mainnet-poet-1.spacemesh.network",
				"https://mainnet-poet-2.spacemesh.network",
				"https://poet-110.spacemesh.network",
				"https://poet-111.spacemesh.network",
				"https://poet-112.spacemesh.network",
			},
			RegossipAtxInterval: 2 * time.Hour,
		},
		Genesis: &GenesisConfig{
			GenesisTime: "2023-07-14T08:00:00Z",
			ExtraData:   "00000000000000000001a6bc150307b5c1998045752b3c87eccf3c013036f3cc",
			Accounts:    MainnetAccounts(),
		},
		Tortoise: tortoise.Config{
			Hdist:                    10,
			Zdist:                    2,
			WindowSize:               10000,
			MaxExceptions:            1000,
			BadBeaconVoteDelayLayers: 4032,
			// 1000 - is assumed minimal number of units
			// 5000 - half of the expected poet ticks
			MinimalActiveSetWeight: 1000 * 5000,
		},
		HARE: hareConfig.Config{
			N:               200,
			ExpectedLeaders: 5,
			RoundDuration:   25 * time.Second,
			WakeupDelta:     25 * time.Second,
			LimitConcurrent: 2,
			LimitIterations: 4,
		},
		HareEligibility: eligConfig.Config{
			ConfidenceParam: 200,
		},
		Beacon: beacon.Config{
			Kappa:                    40,
			Q:                        big.NewRat(1, 3),
			Theta:                    big.NewRat(1, 4),
			GracePeriodDuration:      10 * time.Minute,
			ProposalDuration:         4 * time.Minute,
			FirstVotingRoundDuration: 30 * time.Minute,
			RoundsNumber:             0,
			VotingRoundDuration:      4 * time.Minute,
			WeakCoinRoundDuration:    4 * time.Minute,
			VotesLimit:               100,
			BeaconSyncWeightUnits:    800,
		},
		POET: activation.PoetConfig{
			PhaseShift:        240 * time.Hour,
			CycleGap:          12 * time.Hour,
			GracePeriod:       1 * time.Hour,
			RequestTimeout:    1100 * time.Second, // RequestRetryDelay * 2 * MaxRequestRetries*(MaxRequestRetries+1)/2
			RequestRetryDelay: 10 * time.Second,
			MaxRequestRetries: 10,
		},
		POST: activation.PostConfig{
			MinNumUnits:   4,
			MaxNumUnits:   math.MaxUint32,
			LabelsPerUnit: 4294967296,
			K1:            26,
			K2:            37,
			K3:            37,
			PowDifficulty: postPowDifficulty,
		},
		Bootstrap: bootstrap.Config{
			URL:      "https://bootstrap.spacemesh.network/mainnet",
			Version:  "https://spacemesh.io/bootstrap.schema.json.1.0",
			DataDir:  os.TempDir(),
			Interval: 30 * time.Second,
		},
		P2P:      p2pconfig,
		API:      grpcserver.DefaultConfig(),
		NATS:     nats.DefaultConfig(),
		TIME:     timeConfig.DefaultConfig(),
		SMESHING: smeshing,
		FETCH:    fetch.DefaultConfig(),
		LOGGING:  logging,
		Sync: syncer.Config{
			Interval:                 time.Minute,
			EpochEndFraction:         0.8,
			MaxStaleDuration:         time.Hour,
			Standalone:               false,
			GossipDuration:           50 * time.Second,
			OutOfSyncThresholdLayers: 36, // 3h
		},
		Recovery: checkpoint.DefaultConfig(),
		Cache:    datastore.DefaultConfig(),
	}
}
