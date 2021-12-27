package blocks

import (
	"context"
	"errors"
	"math/rand"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/spacemeshos/go-spacemesh/blocks/mocks"
	"github.com/spacemeshos/go-spacemesh/codec"
	"github.com/spacemeshos/go-spacemesh/common/types"
	"github.com/spacemeshos/go-spacemesh/log/logtest"
	smocks "github.com/spacemeshos/go-spacemesh/system/mocks"
)

type testHandler struct {
	*Handler
	ctrl        *gomock.Controller
	mockFetcher *smocks.MockFetcher
	mockMesh    *mocks.MockmeshProvider
}

func createTestHandler(t *testing.T) *testHandler {
	ctrl := gomock.NewController(t)
	th := &testHandler{
		ctrl:        ctrl,
		mockFetcher: smocks.NewMockFetcher(ctrl),
		mockMesh:    mocks.NewMockmeshProvider(ctrl),
	}
	th.Handler = NewHandler(th.mockFetcher, th.mockMesh, WithLogger(logtest.New(t)))
	return th
}

func createBlockData(t *testing.T, layerID types.LayerID, txIDs []types.TransactionID) (*types.UCBlock, []byte) {
	t.Helper()
	block := &types.UCBlock{
		InnerBlock: types.InnerBlock{
			LayerIndex: layerID,
			TxIDs:      txIDs,
		},
	}
	block.Initialize()
	data, err := codec.Encode(block)
	require.NoError(t, err)
	return block, data
}

func Test_HandleBlockData_MalformedData(t *testing.T) {
	th := createTestHandler(t)
	layerID := types.NewLayerID(99)
	_, txIDs, _ := createTransactions(t, rand.Intn(100))

	_, data := createBlockData(t, layerID, txIDs)
	assert.ErrorIs(t, th.HandleBlockData(context.TODO(), data[1:]), errMalformedData)
}

func Test_HandleBlockData_AlreadyHasBlock(t *testing.T) {
	th := createTestHandler(t)
	layerID := types.NewLayerID(99)
	_, txIDs, _ := createTransactions(t, rand.Intn(100))

	block, data := createBlockData(t, layerID, txIDs)
	th.mockMesh.EXPECT().HasBlock(block.ID()).Return(true).Times(1)
	assert.NoError(t, th.HandleBlockData(context.TODO(), data))
}

func Test_HandleBlockData_FailedToFetchTXs(t *testing.T) {
	th := createTestHandler(t)
	layerID := types.NewLayerID(99)
	_, txIDs, _ := createTransactions(t, rand.Intn(100))

	block, data := createBlockData(t, layerID, txIDs)
	th.mockMesh.EXPECT().HasBlock(block.ID()).Return(false).Times(1)
	errUnknown := errors.New("unknown")
	th.mockFetcher.EXPECT().GetTxs(gomock.Any(), txIDs).Return(errUnknown).Times(1)
	assert.ErrorIs(t, th.HandleBlockData(context.TODO(), data), errUnknown)
}

func Test_HandleBlockData_FailedToAddBlock(t *testing.T) {
	th := createTestHandler(t)
	layerID := types.NewLayerID(99)
	_, txIDs, _ := createTransactions(t, rand.Intn(100))

	block, data := createBlockData(t, layerID, txIDs)
	th.mockMesh.EXPECT().HasBlock(block.ID()).Return(false).Times(1)
	th.mockFetcher.EXPECT().GetTxs(gomock.Any(), txIDs).Return(nil).Times(1)
	errUnknown := errors.New("unknown")
	th.mockMesh.EXPECT().AddBlock(block).Return(errUnknown).Times(1)
	assert.ErrorIs(t, th.HandleBlockData(context.TODO(), data), errUnknown)
}

func Test_HandleBlockData(t *testing.T) {
	th := createTestHandler(t)
	layerID := types.NewLayerID(99)
	_, txIDs, _ := createTransactions(t, rand.Intn(100))

	block, data := createBlockData(t, layerID, txIDs)
	th.mockMesh.EXPECT().HasBlock(block.ID()).Return(false).Times(1)
	th.mockFetcher.EXPECT().GetTxs(gomock.Any(), txIDs).Return(nil).Times(1)
	th.mockMesh.EXPECT().AddBlock(block).Return(nil).Times(1)
	assert.NoError(t, th.HandleBlockData(context.TODO(), data))
}