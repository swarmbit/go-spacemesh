package types

import (
	"bytes"
	"fmt"
	"math/big"
	"sort"

	"github.com/google/go-cmp/cmp"
	"github.com/spacemeshos/go-scale"
	"go.uber.org/zap/zapcore"

	"github.com/spacemeshos/go-spacemesh/codec"
	"github.com/spacemeshos/go-spacemesh/common/util"
)

const (
	// BlockIDSize in bytes.
	// FIXME(dshulyak) why do we cast to hash32 when returning bytes?
	// probably required for fetching by hash between peers.
	BlockIDSize = Hash32Length
)

//go:generate scalegen -types Block,InnerBlock,RatNum,AnyReward,Certificate,CertifyMessage,CertifyContent

// BlockID is a 20-byte blake3 sum of the serialized block used to identify a Block.
type BlockID Hash20

// EmptyBlockID is a canonical empty BlockID.
var EmptyBlockID = BlockID{}

// NewExistingBlock creates a block from existing data.
func NewExistingBlock(id BlockID, inner InnerBlock) *Block {
	return &Block{blockID: id, InnerBlock: inner}
}

// EncodeScale implements scale codec interface.
func (id *BlockID) EncodeScale(e *scale.Encoder) (int, error) {
	return scale.EncodeByteArray(e, id[:])
}

// DecodeScale implements scale codec interface.
func (id *BlockID) DecodeScale(d *scale.Decoder) (int, error) {
	return scale.DecodeByteArray(d, id[:])
}

func (id *BlockID) IsEmpty() bool {
	return *id == EmptyBlockID
}

func (id *BlockID) MarshalText() ([]byte, error) {
	return util.Base64Encode(id[:]), nil
}

func (id *BlockID) UnmarshalText(buf []byte) error {
	return util.Base64Decode(id[:], buf)
}

// Block contains the content of a layer on the mesh history.
type Block struct {
	InnerBlock
	// the following fields are kept private and from being serialized
	blockID BlockID
}

func (b Block) Equal(other Block) bool {
	return cmp.Equal(b.InnerBlock, other.InnerBlock)
}

// InnerBlock contains the transactions and rewards of a block.
type InnerBlock struct {
	LayerIndex LayerID
	TickHeight uint64
	// Rewards are the rewards for the block.
	//
	// Worst case scenario is that a single smesher identity has > 99.97% of the total weight of the network.
	// In this case they will get all 50 available slots in all 4032 layers of the epoch.
	// Additionally every other identity on the network that successfully published an ATX will get 1 slot.
	//
	// If we expect 8.0 Mio ATXs that would be a total of 8.0 Mio + 50 * 4032 = 8 201 600 slots.
	// Since these are randomly distributed across the epoch, we can expect an average of n * p =
	// 8 201 600 / 4032 = 2034.1 rewards in a block with a standard deviation of sqrt(n * p * (1 - p)) =
	// sqrt(8 201 600 * 1/4032 * 4031/4032) = 45.1
	//
	// This means that we can expect a maximum of 2034.1 + 6*45.1 = 2304.7 rewards per block with
	// > 99.9997% probability.
	Rewards []AnyReward     `scale:"max=2350"`
	TxIDs   []TransactionID `scale:"max=100000"`
}

// RatNum represents a rational number with the numerator and denominator.
// note: RatNum aims to be a generic representation of a rational number and parse-able by
// different programming languages.
// for doing math around weight inside go-spacemesh codebase, use util.Weight.
type RatNum struct {
	Num, Denom uint64
}

// String implements fmt.Stringer interface for RatNum.
func (r *RatNum) String() string {
	return fmt.Sprintf("%d/%d", r.Num, r.Denom)
}

// ToBigRat creates big.Rat instance.
func (r *RatNum) ToBigRat() *big.Rat {
	return new(big.Rat).SetFrac(
		new(big.Int).SetUint64(r.Num),
		new(big.Int).SetUint64(r.Denom),
	)
}

func RatNumFromBigRat(r *big.Rat) RatNum {
	return RatNum{Num: r.Num().Uint64(), Denom: r.Denom().Uint64()}
}

// AnyReward contains the reward information by ATXID.
type AnyReward struct {
	AtxID  ATXID
	Weight RatNum
}

// CoinbaseReward contains the reward information by coinbase and smesher, used as an interface to VM.
type CoinbaseReward struct {
	SmesherID NodeID
	Coinbase  Address
	Weight    RatNum
}

// Initialize calculates and sets the Block's cached blockID.
func (b *Block) Initialize() {
	b.blockID = BlockID(CalcHash20(b.Bytes()))
}

// Bytes returns the serialization of the InnerBlock.
func (b *Block) Bytes() []byte {
	return codec.MustEncode(&b.InnerBlock)
}

// ID returns the BlockID.
func (b *Block) ID() BlockID {
	return b.blockID
}

// ToVote creates Vote struct from block.
func (b *Block) ToVote() Vote {
	return Vote{ID: b.ID(), LayerID: b.LayerIndex, Height: b.TickHeight}
}

// MarshalLogObject implements logging encoder for Block.
func (b *Block) MarshalLogObject(encoder zapcore.ObjectEncoder) error {
	encoder.AddString("block_id", b.ID().String())
	encoder.AddUint32("layer_id", b.LayerIndex.Uint32())
	encoder.AddUint64("tick_height", b.TickHeight)
	encoder.AddInt("num_tx", len(b.TxIDs))
	encoder.AddInt("num_rewards", len(b.Rewards))
	return nil
}

// Bytes returns the BlockID as a byte slice.
func (id BlockID) Bytes() []byte {
	return id.AsHash32().Bytes()
}

// AsHash32 returns a Hash32 whose first 20 bytes are the bytes of this BlockID, it is right-padded with zeros.
func (id BlockID) AsHash32() Hash32 {
	return Hash20(id).ToHash32()
}

// String implements the Stringer interface.
func (id BlockID) String() string {
	return Hash20(id).ShortString()
}

// Compare returns true if other (the given BlockID) is less than this BlockID, by lexicographic comparison.
func (id BlockID) Compare(other BlockID) bool {
	return bytes.Compare(id.Bytes(), other.Bytes()) < 0
}

// BlockIDsToHashes turns a list of BlockID into their Hash32 representation.
func BlockIDsToHashes(ids []BlockID) []Hash32 {
	hashes := make([]Hash32, 0, len(ids))
	for _, id := range ids {
		hashes = append(hashes, id.AsHash32())
	}
	return hashes
}

type blockIDs []BlockID

func (ids blockIDs) MarshalLogArray(encoder zapcore.ArrayEncoder) error {
	for i := range ids {
		encoder.AppendString(ids[i].String())
	}
	return nil
}

// SortBlockIDs sorts a list of BlockID in lexicographic order, in-place.
func SortBlockIDs(ids blockIDs) []BlockID {
	sort.Slice(ids, func(i, j int) bool { return ids[i].Compare(ids[j]) })
	return ids
}

// ToBlockIDs returns a slice of BlockID corresponding to the given list of Block.
func ToBlockIDs(blocks []*Block) []BlockID {
	ids := make([]BlockID, 0, len(blocks))
	for _, b := range blocks {
		ids = append(ids, b.ID())
	}
	return ids
}

// BlockContextualValidity represents the contextual validity of a block.
type BlockContextualValidity struct {
	ID       BlockID
	Validity bool
}

// Certificate represents a certified block.
type Certificate struct {
	BlockID    BlockID
	Signatures []CertifyMessage `scale:"max=1000"` // the max. size depends on HARE.N config parameter + safety buffer
}

// CertifyMessage is generated by a node that's eligible to certify the hare output and is gossiped to the network.
type CertifyMessage struct {
	CertifyContent

	Signature EdSignature
	SmesherID NodeID
}

// CertifyContent is actual content the node would sign to certify a hare output.
type CertifyContent struct {
	LayerID LayerID
	BlockID BlockID
	// EligibilityCnt is the number of eligibility of a NodeID on the given Layer
	// as a hare output certifier.
	EligibilityCnt uint16
	// Proof is the role proof for being a hare output certifier on the given Layer.
	Proof VrfSignature
}

// Bytes returns the actual data being signed in a CertifyMessage.
func (cm *CertifyMessage) Bytes() []byte {
	return codec.MustEncode(&cm.CertifyContent)
}
