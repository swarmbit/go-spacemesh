package nats

import (
	"encoding/json"

	"github.com/nats-io/nats.go"
	"github.com/spacemeshos/go-spacemesh/log"
)

type NatsConnector struct {
	nc *nats.Conn
	js nats.JetStreamContext
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
	return &NatsConnector{
		nc: nc,
		js: js,
	}, err

}

type LayerUpdate struct {
	LayerID int `json:"layer"`
	Status  int `json:"status"`
}

type Reward struct {
	Layer       int    `json:"layer"`
	Total       uint64 `json:"totalReward"`
	LayerReward uint64 `json:"layerReward"`
	Coinbase    string `json:"coinbase"`
	AtxID       string `json:"atxID"`
	NodeID      string `json:"nodeID"`
}

type Atx struct {
	Received          int64
	BaseTick          uint64 `json:"baseTick"`
	TickCount         uint64 `json:"tickCount"`
	EffectiveNumUnits uint32 `json:"EffectiveNumUnits"`
	AtxID             string `json:"atxID"`
	NodeID            string `json:"nodeID"`
	Sequence          uint64 `json:"sequence"`
	PublishEpoch      uint32 `json:"publishEpoch"`
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
	n.js.Publish("layers", jsonData)
}

func (n *NatsConnector) PublishRewards(reward *Reward) {
	jsonData, err := json.Marshal(reward)
	if err != nil {
		log.With().Warning("failed to serialize event")
		panic("Failed to serialize rewards")
	}
	n.js.Publish("rewards", jsonData)
}

func (n *NatsConnector) PublishATX(atx *Atx) {
	jsonData, err := json.Marshal(atx)
	if err != nil {
		log.With().Warning("failed to serialize event")
		panic("Failed to serialize transaction atx")
	}
	n.js.Publish("atx", jsonData)
}

func (n *NatsConnector) PublishNewTransaction(transaction *Transaction) {
	jsonData, err := json.Marshal(transaction)
	if err != nil {
		log.With().Warning("failed to serialize event")
		panic("Failed to serialize new transaction")
	}
	n.js.Publish("transactions.created", jsonData)
}

func (n *NatsConnector) PublishTransactionResult(transaction *Transaction) {
	jsonData, err := json.Marshal(transaction)
	if err != nil {
		log.With().Warning("failed to serialize event")
		panic("Failed to serialize transaction result")
	}
	n.js.Publish("transactions.result", jsonData)
}
