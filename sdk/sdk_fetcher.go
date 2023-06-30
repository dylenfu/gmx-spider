package sdk

// import (
// 	"context"
// 	"fmt"
// 	"math/big"

// 	"github.com/dylenfu/gmx-spider/model"
// 	"github.com/ethereum/go-ethereum/accounts/abi/bind"
// 	"github.com/ethereum/go-ethereum/common"
// )

// var (
// 	TestNode                         = "https://arb1.arbitrum.io/rpc"                                    // mainnet
// 	TestChainID       uint64         = 42161                                                             // mainnet
// 	TestVaultContract common.Address = common.HexToAddress("0x489ee077994B6658eAfA855C308275EAd8097C4A") // mainnet
// 	TestQueryContract common.Address = common.HexToAddress("0xeaE393358a1cdbCf80f1285c44B402CF58239Fc9") // mainnet
// )

// type Fetcher struct {
// 	chainId       uint64
// 	urls          []string
// 	term          uint64
// 	sdk           *EthereumSdkPro
// 	vaultContract common.Address
// }

// func NewFetcher(chainID uint64, urls []string, nodeListenTerm uint64, vaultContract common.Address) *Fetcher {
// 	fetcher := &Fetcher{
// 		chainId:       chainID,
// 		urls:          urls,
// 		term:          nodeListenTerm,
// 		vaultContract: vaultContract,
// 	}

// 	fetcher.sdk = NewEthereumSdkPro(fetcher.urls, fetcher.term, fetcher.chainId)

// 	return fetcher
// }

// func (f *Fetcher) GetLatestHeight() (uint64, error) {
// 	return f.sdk.GetLatestHeight()
// }

// func (f *Fetcher) GetChainListenTerm() uint64 {
// 	return f.term
// }

// func (f *Fetcher) GetChainId() uint64 {
// 	return f.chainId
// }

// func (f *Fetcher) GetPosition(blockNumber uint64, account, collateralToken, indexToken common.Address, isLong bool) (*model.Position, error) {
// 	out, err := f.sdk.GetPosition(blockNumber, TestQueryContract, account, collateralToken, indexToken, isLong)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &model.Position{
// 		Size              out.Size,
// 	// Collateral            *big.Int
// 	// AveragePrice          *big.Int
// 	// EntryFundingRate      *big.Int
// 	// ReserveAmount         *big.Int
// 	// RealisedPnl           *big.Int
// 	// LastIncreasedTime     *big.Int
// 	// FloatingPnl           *big.Int
// 	// UncollectedFundingFee *big.Int

// 	// BlockNumber     uint64
// 	// Timestamp       uint64
// 	// Account         common.Address
// 	// CollateralToken common.Address
// 	// IndexToken      common.Address
// 	// IsLong          bool
// 	}
// }

// func (f *Fetcher) HandleNewBlock(start, end uint64) ([]*model.IncreasePosition, []*model.DecreasePosition, error) {
// 	if start < end {
// 		return nil, nil, fmt.Errorf("invalid scan range, start height %d < end height %d", start, end)
// 	}

// 	// map block number to block.timestamp
// 	blockTimer := make(map[uint64]uint64)
// 	for i := start; i <= end; i++ {
// 		timestamp, err := f.sdk.GetBlockTimeByNumber(i)
// 		if err != nil {
// 			return nil, nil, fmt.Errorf("getBlockTimeByNumber failed, err: %v", err)
// 		}
// 		blockTimer[i] = timestamp
// 	}

// 	// fetch abi events
// 	eIncreasePostions, err := f.sdk.infos.(start, end)
// 	if err != nil {
// 		return nil, nil, err
// 	}
// 	eDecreasePositions, err := f.getDecreasePositionEventByBlockNumber(start, end)
// 	if err != nil {
// 		return nil, nil, err
// 	}

// 	// convert to model
// 	mIncreasePositions := make([]*model.IncreasePosition, 0)
// 	mDecreasePositions := make([]*model.DecreasePosition, 0)
// 	for _, event := range eIncreasePostions {
// 		data := ConvertIncreasePosition(event)
// 		data.Timestamp = blockTimer[data.BlockNumber]
// 		mIncreasePositions = append(mIncreasePositions, data)
// 	}
// 	for _, event := range eDecreasePositions {
// 		data := ConvertDecreasePosition(event)
// 		data.Timestamp = blockTimer[data.BlockNumber]
// 		mDecreasePositions = append(mDecreasePositions, data)
// 	}

// 	return mIncreasePositions, mDecreasePositions, nil
// }

// func (f *Fetcher) backend() bind.ContractBackend {
// 	return f.sdk.GetLatest().sdk.backend()
// }
