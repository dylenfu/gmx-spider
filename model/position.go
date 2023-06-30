package model

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

/*
function getPosition(address _account, address _collateralToken, address _indexToken, bool _isLong) public view
    returns (uint256 size, uint256 collateral, uint256 averagePrice, uint256 entryFundingRate, uint256 reserveAmount, int256 realisedPnl, uint256 lastIncreasedTime, int256 floatingPnl, uint256 uncollectedFundingFee) {
        return abi.decode(_getPosition(_account, _collateralToken, _indexToken, _isLong), (uint256, uint256, uint256, uint256, uint256, int256, uint256, int256, uint256));
    }
*/
type Position struct {
	Size                  *big.Int
	Collateral            *big.Int
	AveragePrice          *big.Int
	EntryFundingRate      *big.Int
	ReserveAmount         *big.Int
	RealisedPnl           *big.Int
	LastIncreasedTime     *big.Int
	FloatingPnl           *big.Int
	UncollectedFundingFee *big.Int

	BlockNumber     uint64
	Timestamp       uint64
	Account         common.Address
	CollateralToken common.Address
	IndexToken      common.Address
	IsLong          bool
}
