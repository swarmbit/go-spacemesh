package nats

import (
	"encoding/json"

	"github.com/nats-io/nats.go"
	"github.com/spacemeshos/go-spacemesh/log"
)

type NatsConnector struct {
	nc *nats.Conn
	js nats.JetStreamContext
	config Config
}

func NewNatsConnector(config Config) (*NatsConnector, error) {
	natsURL := config.NatsUrl
	nc, err := nats.Connect(natsURL)
	if err != nil {
		log.With().Warning("failed to connect to nats")
		return &NatsConnector{}, err
	}
	js, err := nc.JetStream()
	if err != nil {
		log.With().Warning("failed to create jetstream")
		return &NatsConnector{}, err
	}
	js.AddStream(&nats.StreamConfig{
		Name:     "layers",
		Subjects: []string{"layers"},
		MaxAge:   config.NatsMaxAge,
	})
	js.AddStream(&nats.StreamConfig{
		Name:     "rewards",
		Subjects: []string{"rewards"},
		MaxAge:   config.NatsMaxAge,
	})
	js.AddStream(&nats.StreamConfig{
		Name:     "transactions",
		Subjects: []string{"transactions.created", "transactions.result"},
		MaxAge:   config.NatsMaxAge,
	})
	js.AddStream(&nats.StreamConfig{
		Name:     "atx",
		Subjects: []string{"atx"},
		MaxAge:   config.NatsMaxAge,
	})
	js.AddStream(&nats.StreamConfig{
		Name:     "malfeasance",
		Subjects: []string{"malfeasance"},
		MaxAge:   config.NatsMaxAge,
	})
	return &NatsConnector{
		nc: nc,
		js: js,
		config: config,
	}, err

}

type Malfeasance struct {
	LayerID  uint32 `json:"layer"`
	NodeID   string `json:"node_id"`
	Received int64  `json:"received"`
}

type LayerUpdate struct {
	LayerID uint32 `json:"layer"`
	Status  int    `json:"status"`
}

type Reward struct {
	ID          string `json:"id"`
	Layer       uint32 `json:"layer"`
	Total       uint64 `json:"totalReward"`
	LayerReward uint64 `json:"layerReward"`
	Coinbase    string `json:"coinbase"`
	AtxID       string `json:"atxID"`
	NodeID      string `json:"nodeID"`
}

type Atx struct {
	Received          int64  `json:"received"`
	BaseTick          uint64 `json:"baseTick"`
	TickCount         uint64 `json:"tickCount"`
	EffectiveNumUnits uint32 `json:"EffectiveNumUnits"`
	AtxID             string `json:"atxID"`
	NodeID            string `json:"nodeID"`
	Sequence          uint64 `json:"sequence"`
	PublishEpoch      uint32 `json:"publishEpoch"`
	Coinbase          string `json:"coinbase"`
}

type Transaction struct {
	ID     string             `json:"id"`
	Header *TransactionHeader `json:"header"`
	Raw    []byte             `json:"raw"`
}

type TransactionHeader struct {
	Message         string   `json:"message"`
	Status          uint8    `json:"status"`
	BlockID         string   `json:"block_id"`
	LayerID         uint32   `json:"layer_id"`
	Principal       string   `json:"principal"`
	TemplateAddress string   `json:"template_address"`
	Method          uint8    `json:"method"`
	Nonce           uint64   `json:"nonce"`
	Gas             uint64   `json:"gas"`
	Fee             uint64   `json:"fee"`
	Addresses       []string `json:"addresses"`
}

func (n *NatsConnector) PublishLayer(layerUpdate *LayerUpdate) {
	jsonData, err := json.Marshal(layerUpdate)
	if err != nil {
		log.With().Warning("failed to serialize event")
		panic("Failed to serialize layer")
	}
	if n.config.NatsAsyncLayers {
		_, err = n.js.PublishAsync("layers", jsonData)
		if err != nil {
			panic("failed to publish layers: " + err.Error())
		}
	} else {
		_, err = n.js.Publish("layers", jsonData)
		if err != nil {
			panic("failed to publish layers: " + err.Error())
		}
	}
}

func (n *NatsConnector) PublishRewards(reward *Reward) {
	jsonData, err := json.Marshal(reward)
	if err != nil {
		log.Warning("failed to serialize event")
		panic("Failed to serialize rewards")
	}
	if n.config.NatsAsyncRewards {
		_, err = n.js.PublishAsync("rewards", jsonData)
		if err != nil {
			panic("failed to publish rewards: " + err.Error())
		}
	} else {
		_, err = n.js.Publish("rewards", jsonData)
		if err != nil {
			panic("failed to publish rewards: " + err.Error())
		}
	}
}

func (n *NatsConnector) PublishATX(atx *Atx) {
	jsonData, err := json.Marshal(atx)
	if err != nil {
		log.Warning("failed to serialize event")
		panic("Failed to serialize transaction atx")
	}
	if n.config.NatsAsyncAtx {
		_, err = n.js.PublishAsync("atx", jsonData)
		if err != nil {
			panic("failed to publish atx: " + err.Error())
		}
	} else {
		_, err = n.js.Publish("atx", jsonData)
		if err != nil {
			panic("failed to publish atx: " + err.Error())
		}
	}
}

func (n *NatsConnector) PublishNewTransaction(transaction *Transaction) {
	jsonData, err := json.Marshal(transaction)
	if err != nil {
		log.Warning("failed to serialize event")
		panic("Failed to serialize new transaction")
	}
	if n.config.NatsAsyncTransactions {
		_, err = n.js.PublishAsync("transactions.created", jsonData)
		if err != nil {
			panic("failed to publish transactions.created: " + err.Error())
		}
	} else {
		_, err = n.js.Publish("transactions.created", jsonData)
		if err != nil {
			panic("failed to publish transactions.created: " + err.Error())
		}
	}

}

func (n *NatsConnector) PublishTransactionResult(transaction *Transaction) {
	jsonData, err := json.Marshal(transaction)
	if err != nil {
		log.Warning("failed to serialize event")
		panic("Failed to serialize transaction result")
	}
	if n.config.NatsAsyncTransactionsResult {
		_, err = n.js.PublishAsync("transactions.result", jsonData)
		if err != nil {
			panic("failed to publish transactions.result: " + err.Error())
		}
	} else  {
		_, err = n.js.Publish("transactions.result", jsonData)
		if err != nil {
			panic("failed to publish transactions.result: " + err.Error())
		}
	}
}

func (n *NatsConnector) PublishMalfeasance(malfeasance *Malfeasance) {
	jsonData, err := json.Marshal(malfeasance)
	if err != nil {
		log.Warning("failed to serialize event")
		panic("Failed to serialize malfeasance")
	}

	if n.config.NatsAsyncAtx {
		_, err = n.js.PublishAsync("malfeasance", jsonData)
		if err != nil {
			panic("failed to publish malfeasance: " + err.Error())
		}
	} else {
		_, err = n.js.Publish("malfeasance", jsonData)
		if err != nil {
			panic("failed to publish malfeasance: " + err.Error())
		}
	}

}
