package sdk

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"time"

	"github.com/beego/beego/v2/core/logs"
	erc20 "github.com/dylenfu/gmx-spider/abi/mintable_erc20_abi"
	qabi "github.com/dylenfu/gmx-spider/abi/query_abi"
	vabi "github.com/dylenfu/gmx-spider/abi/vault_abi"
	"github.com/dylenfu/gmx-spider/model"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/ethereum/go-ethereum/rpc"
)

var (
	DefaultGasLimit uint64         = 50000
	EmptyAddress                   = common.Address{}
	EmptyHash                      = common.Hash{}
	TestNode                       = "https://arb1.arbitrum.io/rpc"                                    // mainnet
	ChainID         uint64         = 42161                                                             // mainnet
	VaultContract   common.Address = common.HexToAddress("0x489ee077994B6658eAfA855C308275EAd8097C4A") // mainnet
	QueryContract   common.Address = common.HexToAddress("0xeaE393358a1cdbCf80f1285c44B402CF58239Fc9") // mainnet
)

type EthereumSdk struct {
	vaultContract common.Address
	queryContract common.Address
	chainId       uint64
	rpcClient     *rpc.Client
	rawClient     *ethclient.Client
	url           string
}

func NewEthereumSdk(url string) (*EthereumSdk, error) {
	rpcClient, err1 := rpc.Dial(url)
	rawClient, err2 := ethclient.Dial(url)
	if rpcClient == nil || err1 != nil || rawClient == nil || err2 != nil {
		return nil, fmt.Errorf("ethereum node is not working!, err1: %v, err2: %v", err1, err2)
	}
	return &EthereumSdk{
		rpcClient:     rpcClient,
		rawClient:     rawClient,
		url:           url,
		vaultContract: VaultContract,
		queryContract: QueryContract,
		chainId:       ChainID,
	}, nil
}

func (sdk *EthereumSdk) HandleNewBlock(start, end uint64) ([]*model.IncreasePosition, []*model.DecreasePosition, error) {
	if start < end {
		return nil, nil, fmt.Errorf("invalid scan range, start height %d < end height %d", start, end)
	}

	// map block number to block.timestamp
	blockTimer := make(map[uint64]uint64)
	for i := start; i <= end; i++ {
		timestamp, err := sdk.GetBlockTimeByNumber(i)
		if err != nil {
			return nil, nil, fmt.Errorf("getBlockTimeByNumber failed, err: %v", err)
		}
		blockTimer[i] = timestamp
	}

	increasePositionEvents, err := sdk.GetIncreasePositionEventByBlockNumber(start, end)
	if err != nil {
		return nil, nil, err
	}
	for _, data := range increasePositionEvents {
		data.Timestamp = blockTimer[data.BlockNumber]
	}
	decreasePositionEvents, err := sdk.GetDecreasePositionEventByBlockNumber(start, end)
	if err != nil {
		return nil, nil, err
	}
	for _, data := range decreasePositionEvents {
		data.Timestamp = blockTimer[data.BlockNumber]
	}
	return increasePositionEvents, decreasePositionEvents, nil
}

func (sdk *EthereumSdk) GetIncreasePositionEventByBlockNumber(start, end uint64) ([]*model.IncreasePosition, error) {
	filterOpts := newFilterOption(start, end)
	vaultABI, err := vabi.NewVault(sdk.vaultContract, sdk.backend())
	if err != nil {
		return nil, fmt.Errorf("NewVaultABI failed, err: %v", err)
	}
	iter, err := vaultABI.FilterIncreasePosition(filterOpts)
	if err != nil {
		return nil, fmt.Errorf("vaultABI.FilterIncreasePosition failed, err: %v", err)
	}

	list := make([]*model.IncreasePosition, 0)
	for iter.Next() {
		event := ConvertIncreasePosition(iter.Event)
		list = append(list, event)
	}
	return list, nil
}

func (sdk *EthereumSdk) GetDecreasePositionEventByBlockNumber(start, end uint64) ([]*model.DecreasePosition, error) {
	filterOpts := newFilterOption(start, end)
	vaultABI, err := vabi.NewVault(sdk.vaultContract, sdk.backend())
	if err != nil {
		return nil, fmt.Errorf("NewVaultABI failed, err: %v", err)
	}
	iter, err := vaultABI.FilterDecreasePosition(filterOpts)
	if err != nil {
		return nil, fmt.Errorf("vaultABI.FilterDecreasePosition failed, err: %v", err)
	}
	list := make([]*model.DecreasePosition, 0)
	for iter.Next() {
		event := ConvertDecreasePosition(iter.Event)
		list = append(list, event)
	}
	return list, nil
}

func (sdk *EthereumSdk) GetPosition(blockNumber uint64, account, collateralToken, indexToken common.Address, isLong bool) (*model.Position, error) {
	timestamp, err := sdk.GetBlockTimeByNumber(blockNumber)
	if err != nil {
		return nil, fmt.Errorf("getBlockTimeByNumber failed, err: %v", err)
	}

	queryABI, err := qabi.NewQuery(sdk.queryContract, sdk.backend())
	if err != nil {
		return nil, fmt.Errorf("NewQueryABI failed, err: %v", err)
	}
	opt := &bind.CallOpts{BlockNumber: new(big.Int).SetUint64(blockNumber)}
	out, err := queryABI.GetPosition(opt, account, collateralToken, indexToken, isLong)
	if err != nil {
		return nil, fmt.Errorf("queryABI.GetPosition failed, err: %v", err)
	}
	position := &model.Position{}
	position.Size = out.Size
	position.Collateral = out.Collateral
	position.AveragePrice = out.AveragePrice
	position.EntryFundingRate = out.EntryFundingRate
	position.ReserveAmount = out.ReserveAmount
	position.RealisedPnl = out.RealisedPnl
	position.LastIncreasedTime = out.LastIncreasedTime
	position.FloatingPnl = out.FloatingPnl
	position.UncollectedFundingFee = out.UncollectedFundingFee

	position.BlockNumber = blockNumber
	position.Timestamp = timestamp
	position.Account = account
	position.CollateralToken = collateralToken
	position.IndexToken = indexToken
	position.IsLong = isLong
	return position, nil
}

func (s *EthereumSdk) GetClient() *ethclient.Client {
	return s.rawClient
}

func (s *EthereumSdk) GetUrl() string {
	return s.url
}

func (s *EthereumSdk) GetCurrentBlockHeight() (uint64, error) {
	var result hexutil.Big
	err := s.rpcClient.CallContext(context.Background(), &result, "eth_blockNumber")
	for err != nil {
		return 0, err
	}
	return (*big.Int)(&result).Uint64(), err
}

func toBlockNumArg(number *big.Int) string {
	if number == nil {
		return "latest"
	}
	return hexutil.EncodeBig(number)
}

// GetHeaderByNumber returns the given header
func (s *EthereumSdk) GetHeaderByNumber(number uint64) (*types.Header, error) {
	var header *types.Header
	var newNumber *big.Int
	if number < 0 {
		newNumber = nil
	} else {
		newNumber = big.NewInt(int64(number))
	}
	err := s.rpcClient.CallContext(context.Background(), &header, "eth_getBlockByNumber", toBlockNumArg(newNumber), false)
	for err != nil {
		return nil, err
	}
	return header, err
}

// GetBlockTimeByNumber returns the timestamp of given block number
func (s *EthereumSdk) GetBlockTimeByNumber(number uint64) (timestamp uint64, err error) {
	var header *types.Header
	newNumber := big.NewInt(int64(number))

	err = s.rpcClient.CallContext(context.Background(), &header, "eth_getBlockByNumber", toBlockNumArg(newNumber), false)
	for err != nil {
		return 0, err
	}
	timestamp = header.Time
	return
}

func (s *EthereumSdk) GetBlockByNumber(number uint64) (*types.Block, error) {
	return s.rawClient.BlockByNumber(context.Background(), new(big.Int).SetUint64(number))
}

func (s *EthereumSdk) GetTransactionByHash(hash common.Hash) (*types.Transaction, error) {
	tx, _, err := s.rawClient.TransactionByHash(context.Background(), hash)
	for err != nil {
		return nil, err
	}
	return tx, err
}

func (s *EthereumSdk) GetTransactionReceipt(hash common.Hash) (*types.Receipt, error) {
	receipt, err := s.rawClient.TransactionReceipt(context.Background(), hash)
	for err != nil {
		return nil, err
	}
	return receipt, nil
}

func (s *EthereumSdk) NonceAt(addr common.Address) (uint64, error) {
	nonce, err := s.rawClient.PendingNonceAt(context.Background(), addr)
	for err != nil {
		return 0, err
	}
	return nonce, nil
}

func (s *EthereumSdk) SendRawTransaction(tx *types.Transaction) error {
	data, err := rlp.EncodeToBytes(tx)
	if err != nil {
		return err
	}
	err = s.rpcClient.CallContext(context.Background(), nil, "eth_sendRawTransaction", hexutil.Encode(data))
	for err != nil {
		return err
	}
	return nil
}

func (s *EthereumSdk) TransactionByHash(hash common.Hash) (*types.Transaction, bool, error) {
	tx, isPending, err := s.rawClient.TransactionByHash(context.Background(), hash)
	for err != nil {
		return nil, false, err
	}
	return tx, isPending, err
}

func (s *EthereumSdk) SuggestGasPrice() (*big.Int, error) {
	gasPrice, err := s.rawClient.SuggestGasPrice(context.Background())
	for err != nil {
		return nil, err
	}
	return gasPrice, err
}

func (s *EthereumSdk) EstimateGas(msg ethereum.CallMsg) (uint64, error) {
	gasLimit, err := s.rawClient.EstimateGas(context.Background(), msg)
	for err != nil {
		return 0, err
	}
	return gasLimit, err
}

func (s *EthereumSdk) GetERC20Balance(asset, owner common.Address) (*big.Int, error) {
	contract, err := erc20.NewERC20Mintable(asset, s.backend())
	if err != nil {
		return nil, err
	}
	return contract.BalanceOf(nil, owner)
}

func (s *EthereumSdk) GetERC20TotalSupply(asset common.Address) (*big.Int, error) {
	contract, err := erc20.NewERC20Mintable(asset, s.backend())
	if err != nil {
		return nil, err
	}
	return contract.TotalSupply(nil)
}

func (s *EthereumSdk) ApproveERC20Token(
	key *ecdsa.PrivateKey,
	asset, spender common.Address,
	amount *big.Int,
) (common.Hash, error) {

	contract, err := erc20.NewERC20Mintable(asset, s.backend())
	if err != nil {
		return EmptyHash, err
	}

	auth, err := s.makeAuth(key, DefaultGasLimit)
	if err != nil {
		return EmptyHash, err
	}
	tx, err := contract.Approve(auth, spender, amount)
	if err != nil {
		return EmptyHash, err
	}

	if err := s.waitTxConfirm(tx.Hash()); err != nil {
		return EmptyHash, err
	}

	return tx.Hash(), nil
}

func (s *EthereumSdk) GetERC20Allowance(asset, owner, spender common.Address) (*big.Int, error) {
	contract, err := erc20.NewERC20Mintable(asset, s.backend())
	if err != nil {
		return nil, err
	}
	return contract.Allowance(nil, owner, spender)
}

//func (ec *EthereumSdk) Erc20Info(hash string) (string, string, int64, string, error) {
//	erc20Address := common.HexToAddress(hash)
//	erc20Contract, err := usdt_abi.NewTetherToken(erc20Address, ec.rawClient)
//	if err != nil {
//		return "", "", 0, "", err
//	}
//	name, err := erc20Contract.Name(&bind.CallOpts{})
//	if err != nil {
//		return "", "", 0, "", err
//	}
//	/*
//		totolSupply, err := erc20Contract.TotalSupply(&bind.CallOpts{})
//		if err != nil {
//			return "", "", 0, "", err
//		}
//	*/
//	decimal, err := erc20Contract.Decimals(&bind.CallOpts{})
//	if err != nil {
//		return "", "", 0, "", err
//	}
//	symbol, err := erc20Contract.Symbol(&bind.CallOpts{})
//	if err != nil {
//		return "", "", 0, "", err
//	}
//	return hash, name, decimal.Int64(), symbol, nil
//}

//func (ec *EthereumSdk) Erc20Balance(erc20 string, addr string) (uint64, error) {
//	erc20Address := common.HexToAddress(erc20)
//	erc20Contract, err := erc20_abi.NewERC20(erc20Address, ec.rawClient)
//	if err != nil {
//		return 0, fmt.Errorf("erc20 address is not right")
//	}
//	userAddress := common.HexToAddress(addr)
//	balance, err := erc20Contract.BalanceOf(&bind.CallOpts{}, userAddress)
//	if err != nil {
//		return 0, err
//	}
//	return balance.Uint64(), nil
//}

func (s *EthereumSdk) EthBalance(addr string) (*big.Int, error) {
	var result hexutil.Big
	ctx := context.Background()
	err := s.rpcClient.CallContext(ctx, &result, "eth_getBalance", "0x"+addr, "latest")
	return (*big.Int)(&result), err
}

func (s *EthereumSdk) FilterLog(FromBlock *big.Int, ToBlock *big.Int, Addresses []common.Address) ([]types.Log, error) {
	ctx := context.Background()
	var filterQuery ethereum.FilterQuery
	filterQuery.FromBlock = FromBlock
	filterQuery.ToBlock = ToBlock
	filterQuery.Addresses = Addresses
	return s.rawClient.FilterLogs(ctx, filterQuery)
}

func (s *EthereumSdk) DeployERC20(key *ecdsa.PrivateKey) (common.Address, error) {
	auth, err := s.makeAuth(key, DefaultGasLimit)
	if err != nil {
		return EmptyAddress, err
	}

	addr, tx, _, err := erc20.DeployERC20(auth, s.backend())
	if err != nil {
		return EmptyAddress, err
	}

	if err := s.waitTxConfirm(tx.Hash()); err != nil {
		return EmptyAddress, err
	}

	return addr, nil
}

func (s *EthereumSdk) dumpTx(hash common.Hash) error {
	tx, err := s.GetTransactionReceipt(hash)
	if err != nil {
		return fmt.Errorf("faild to get receipt %s", hash.Hex())
	}

	if tx.Status == 0 {
		return fmt.Errorf("receipt failed %s", hash.Hex())
	}

	logs.Info("txhash %s, block height %d", hash.Hex(), tx.BlockNumber.Uint64())
	for _, event := range tx.Logs {
		logs.Info("eventlog address %s", event.Address.Hex())
		logs.Info("eventlog data %s", new(big.Int).SetBytes(event.Data).String())
		for i, topic := range event.Topics {
			logs.Info("eventlog topic[%d] %s", i, topic.String())
		}
	}
	return nil
}

func (s *EthereumSdk) waitTxConfirm(hash common.Hash) error {
	for {
		time.Sleep(time.Second * 1)
		_, ispending, err := s.TransactionByHash(hash)
		if err != nil {
			logs.Error("failed to call TransactionByHash: %v", err)
			continue
		}
		if ispending == true {
			continue
		} else {
			break
		}
	}
	logs.Info("tx %s confirmed", hash.Hex())
	if err := s.dumpTx(hash); err != nil {
		return err
	}
	return nil
}

func (s *EthereumSdk) makeAuth(key *ecdsa.PrivateKey, gasLimit uint64) (*bind.TransactOpts, error) {
	authAddress := Key2address(key)
	nonce, err := s.NonceAt(authAddress)
	if err != nil {
		return nil, fmt.Errorf("makeAuth, addr %s, err %v", authAddress.Hex(), err)
	}

	gasPrice, err := s.SuggestGasPrice()
	if err != nil {
		return nil, fmt.Errorf("makeAuth, get suggest gas price err: %v", err)
	}

	auth := bind.NewKeyedTransactor(key)
	auth.From = authAddress
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(int64(0)) // in wei
	auth.GasLimit = gasLimit
	auth.GasPrice = gasPrice
	//if DefaultAddGasPrice.Cmp(gasPrice) > 0 {
	//	auth.GasPrice = DefaultAddGasPrice
	//} else {
	//	auth.GasPrice = gasPrice
	//}

	return auth, nil
}

func (s *EthereumSdk) backend() bind.ContractBackend {
	return s.rawClient
}

func Key2address(key *ecdsa.PrivateKey) common.Address {
	publicKey := key.Public()
	publicKeyECDSA := publicKey.(*ecdsa.PublicKey)
	return crypto.PubkeyToAddress(*publicKeyECDSA)
}

func newFilterOption(start, end uint64) *bind.FilterOpts {
	return &bind.FilterOpts{
		Start:   start,
		End:     &end,
		Context: context.Background(),
	}
}

func ConvertIncreasePosition(event *vabi.VaultIncreasePosition) *model.IncreasePosition {
	data := &model.IncreasePosition{
		Key:             common.BytesToHash(event.Key[:]),
		Account:         event.Account,
		CollateralToken: event.CollateralToken,
		IndexToken:      event.IndexToken,
		CollateralDelta: event.CollateralDelta,
		SizeDelta:       event.SizeDelta,
		IsLong:          event.IsLong,
		Price:           event.Price,
		Fee:             event.Fee,
		BlockNumber:     event.Raw.BlockNumber,
		TxHash:          event.Raw.TxHash,
	}
	return data
}

func ConvertDecreasePosition(event *vabi.VaultDecreasePosition) *model.DecreasePosition {
	data := &model.DecreasePosition{
		Key:             common.BytesToHash(event.Key[:]),
		Account:         event.Account,
		CollateralToken: event.CollateralToken,
		IndexToken:      event.IndexToken,
		CollateralDelta: event.CollateralDelta,
		SizeDelta:       event.SizeDelta,
		IsLong:          event.IsLong,
		Price:           event.Price,
		Fee:             event.Fee,
		BlockNumber:     event.Raw.BlockNumber,
		TxHash:          event.Raw.TxHash,
	}
	return data
}
