package nats

import "time"

type Config struct {
	NatsEnabled bool          		`mapstructure:"nats-enabled"`
	NatsAsyncAtx bool         		`mapstructure:"nats-async-atx"`
	NatsAsyncRewards bool     		`mapstructure:"nats-async-rewards"`
	NatsAsyncLayers bool      		`mapstructure:"nats-async-layers"`
	NatsAsyncTransactions bool      `mapstructure:"nats-async-transactions"`
	NatsAsyncTransactionsResult bool`mapstructure:"nats-async-transactions-result"`
	NatsAsyncMalfeasance bool		`mapstructure:"nats-async-malfeasance"`
	NatsUrl     string        		`mapstructure:"nats-url"`
	NatsMaxAge  time.Duration 		`mapstructure:"nats-max-age"`
}

func DefaultConfig() Config {
	return Config{
		NatsEnabled: false,
		NatsUrl:     "nats://0.0.0.0:4222",
		NatsMaxAge:  365 * 24 * time.Hour,
		NatsAsyncAtx: true,
		NatsAsyncLayers: false,
		NatsAsyncRewards: false,
		NatsAsyncTransactions: false,
		NatsAsyncTransactionsResult: false,
		NatsAsyncMalfeasance: false,
	}
}
