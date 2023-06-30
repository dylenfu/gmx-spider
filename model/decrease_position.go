package model

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type DecreasePosition struct {
	Key             common.Hash
	Account         common.Address
	CollateralToken common.Address
	IndexToken      common.Address
	CollateralDelta *big.Int
	SizeDelta       *big.Int
	IsLong          bool
	Price           *big.Int
	Fee             *big.Int

	// raw data
	BlockNumber uint64
	Timestamp   uint64
	TxHash      common.Hash
}
