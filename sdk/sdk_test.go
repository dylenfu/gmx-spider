package sdk

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v github.com/dylenfu/gmx-spider/sdk -run TestQueryEvent
func TestQueryEvent(t *testing.T) {
	sdk, err := NewEthereumSdk(TestNode)
	assert.Nil(t, err)

	chainID, err := sdk.GetClient().ChainID(context.Background())
	assert.Nil(t, err)

	t.Logf("chain id: %d", chainID)
}

// go test -v github.com/dylenfu/gmx-spider/sdk -run TestFilterIncreasePositionEvent
func TestFilterIncreasePositionEvent(t *testing.T) {
	sdk, err := NewEthereumSdk(TestNode)
	assert.Nil(t, err)

	height, err := sdk.GetCurrentBlockHeight()
	assert.Nil(t, err)

	length := 1000
	start := height - uint64(length)
	end := height

	events, err := sdk.GetIncreasePositionEventByBlockNumber(start, end)
	if err != nil {
		t.Fatalf("get increasePostion failed, err: %v", err)
	}

	for _, data := range events {
		t.Log(data.BlockNumber, data.TxHash, data)
	}
}

// go test -v -count=1 github.com/dylenfu/gmx-spider/sdk -run TestFilterDecreasePositionEvent
func TestFilterDecreasePositionEvent(t *testing.T) {
	sdk, err := NewEthereumSdk(TestNode)
	assert.Nil(t, err)

	height, err := sdk.GetCurrentBlockHeight()
	assert.Nil(t, err)

	length := 2000
	start := height - uint64(length)
	end := height

	events, err := sdk.GetDecreasePositionEventByBlockNumber(start, end)
	if err != nil {
		t.Fatalf("get decreasePosition failed, err: %v", err)
	}

	for _, data := range events {
		t.Log(data.BlockNumber, data.TxHash, data)
	}
}

// go test -v -count=1 github.com/dylenfu/gmx-spider/sdk -run TestGetPosition
func TestGetPosition(t *testing.T) {
	sdk, err := NewEthereumSdk(TestNode)
	assert.Nil(t, err)

	// get event in an scope
	height, err := sdk.GetCurrentBlockHeight()
	assert.Nil(t, err)
	length := 1000
	start := height - uint64(length)
	end := height

	events, err := sdk.GetIncreasePositionEventByBlockNumber(start, end)
	if err != nil {
		t.Fatalf("get decreasePosition failed, err: %v", err)
	}
	if len(events) == 0 {
		t.Fatalf("no increasePosition events")
	}
	event := events[0]

	// get position
	blockNumber := event.BlockNumber
	position, err := sdk.GetPosition(blockNumber, event.Account, event.CollateralToken, event.IndexToken, event.IsLong)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(position)
}
