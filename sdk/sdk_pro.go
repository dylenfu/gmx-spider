package sdk

// import (
// 	"context"
// 	"errors"
// 	"fmt"
// 	"math"
// 	"math/big"
// 	"runtime/debug"
// 	"sync"
// 	"time"

// 	"github.com/beego/beego/v2/core/logs"
// 	qabi "github.com/dylenfu/gmx-spider/abi/query_abi"
// 	abi "github.com/dylenfu/gmx-spider/abi/vault_abi"
// 	"github.com/dylenfu/gmx-spider/model"
// 	"github.com/ethereum/go-ethereum"
// 	"github.com/ethereum/go-ethereum/accounts/abi/bind"
// 	"github.com/ethereum/go-ethereum/common"
// 	"github.com/ethereum/go-ethereum/core/types"
// 	"github.com/ethereum/go-ethereum/ethclient"
// )

// var ErrNodesInvalid = errors.New("all node is not working")

// const maxNodeSelection = 3

// type EthereumInfo struct {
// 	sdk          *EthereumSdk
// 	latestHeight uint64
// }

// func (info *EthereumInfo) backend() bind.ContractBackend {
// 	return info.sdk.backend()
// }

// func NewEthereumInfo(url string) *EthereumInfo {
// 	sdk, err := NewEthereumSdk(url)
// 	if err != nil || sdk == nil {
// 		panic(err)
// 	}
// 	return &EthereumInfo{
// 		sdk:          sdk,
// 		latestHeight: 0,
// 	}
// }

// type EthereumSdkPro struct {
// 	infos         map[string]*EthereumInfo
// 	selectionSlot uint64
// 	id            uint64
// 	mutex         sync.Mutex
// }

// func NewEthereumSdkPro(urls []string, slot uint64, id uint64) *EthereumSdkPro {
// 	infos := make(map[string]*EthereumInfo, len(urls))
// 	for _, url := range urls {
// 		infos[url] = NewEthereumInfo(url)
// 	}
// 	pro := &EthereumSdkPro{infos: infos, selectionSlot: slot, id: id}
// 	pro.selection()
// 	go pro.NodeSelection()
// 	return pro
// }

// func (pro *EthereumSdkPro) NodeSelection() {
// 	for {
// 		pro.nodeSelection()
// 	}
// }

// func (pro *EthereumSdkPro) nodeSelection() {
// 	defer func() {
// 		if r := recover(); r != nil {
// 			logs.Error("node selection, recover info: %s", string(debug.Stack()))
// 		}
// 	}()
// 	logs.Debug("node selection of chain : %d......", pro.id)
// 	ticker := time.NewTicker(time.Second * time.Duration(pro.selectionSlot))
// 	for {
// 		select {
// 		case <-ticker.C:
// 			pro.selection()
// 		}
// 	}
// }

// func (pro *EthereumSdkPro) selection() {
// 	for url, info := range pro.infos {
// 		height, err := info.sdk.GetCurrentBlockHeight()
// 		if err != nil || height == math.MaxUint64 || height == 0 {
// 			logs.Error("nodeselection get current block height err, chain %v, url: %s", pro.id, url)
// 			height = 1
// 		}
// 		/*
// 			 block, err := info.sdk.GetBlockByNumber(height)
// 			 if err != nil || block == nil {
// 				 logs.Error("get current block err: %v, url: %s", err, url)
// 				 info.latestHeight = height - 1
// 				 continue
// 			 }
// 			 transactions := block.Transactions()
// 			 if len(transactions) > 0 {
// 				 transaction := transactions[0]
// 				 receipt, err := info.sdk.GetTransactionReceipt(transaction.Hash())
// 				 if err != nil || receipt == nil {
// 					 logs.Error("get transaction receipt err: %v, url: %s", err, url)
// 					 info.latestHeight = height - 1
// 					 continue
// 				 }
// 			 }
// 		*/
// 		pro.mutex.Lock()
// 		info.latestHeight = height - 1
// 		pro.mutex.Unlock()
// 	}
// }

// func (pro *EthereumSdkPro) GetLatest() *EthereumInfo {
// 	pro.mutex.Lock()
// 	defer func() {
// 		pro.mutex.Unlock()
// 	}()
// 	height := uint64(0)
// 	var latestInfo *EthereumInfo = nil
// 	for _, info := range pro.infos {
// 		if info != nil && info.latestHeight > height {
// 			height = info.latestHeight
// 			latestInfo = info
// 		}
// 	}
// 	return latestInfo
// }

// func (pro *EthereumSdkPro) GetClient() *ethclient.Client {
// 	info := pro.GetLatest()
// 	if info == nil {
// 		return nil
// 	}
// 	return info.sdk.GetClient()
// }

// func (pro *EthereumSdkPro) SetClientHeightZero(cli *ethclient.Client) {
// 	for node, info := range pro.infos {
// 		if info.sdk.GetClient() == cli {
// 			logs.Error("SetClientHeightZero node:%v is err", node)
// 			info.latestHeight = 0
// 			break
// 		}
// 	}
// }

// func (pro *EthereumSdkPro) GetLatestHeight() (uint64, error) {
// 	info := pro.GetLatest()
// 	if info == nil {
// 		return 0, ErrNodesInvalid
// 	}
// 	return info.latestHeight, nil
// }

// func (pro *EthereumSdkPro) GetHeaderByNumber(number uint64) (*types.Header, error) {
// 	info := pro.GetLatest()
// 	if info == nil {
// 		return nil, ErrNodesInvalid
// 	}
// 	flag := 0
// 	for info != nil {
// 		header, err := info.sdk.GetHeaderByNumber(number)
// 		if err != nil {
// 			flag++
// 			if flag > maxNodeSelection {
// 				logs.Error("GetHeaderByNumber_chain:%v,node:%v,GetHeaderByNumber err", pro.id, info.sdk.url)
// 				flag = 0
// 				time.Sleep(time.Second)
// 			}
// 			info.latestHeight = 0
// 			info = pro.GetLatest()
// 		} else {
// 			flag = 0
// 			return header, nil
// 		}
// 	}
// 	return nil, ErrNodesInvalid
// }

// func (pro *EthereumSdkPro) GetBlockTimeByNumber(number uint64) (uint64, error) {
// 	info := pro.GetLatest()
// 	if info == nil {
// 		return 0, ErrNodesInvalid
// 	}
// 	flag := 0
// 	for info != nil {
// 		timestamp, err := info.sdk.GetBlockTimeByNumber(pro.id, number)
// 		if err != nil {
// 			flag++
// 			if flag > maxNodeSelection {
// 				logs.Error("GetBlockTimeByNumber_chain:%v,node:%v,eth_getBlockByNumber err %v", pro.id, info.sdk.url, err)
// 				flag = 0
// 				time.Sleep(time.Second)
// 			}
// 			info.latestHeight = 0
// 			info = pro.GetLatest()
// 		} else {
// 			flag = 0
// 			return timestamp, nil
// 		}
// 	}
// 	return 0, ErrNodesInvalid
// }

// func (pro *EthereumSdkPro) GetTransactionByHash(hash common.Hash) (*types.Transaction, error) {
// 	info := pro.GetLatest()
// 	if info == nil {
// 		return nil, ErrNodesInvalid
// 	}
// 	flag := 0
// 	for info != nil {
// 		tx, err := info.sdk.GetTransactionByHash(hash)
// 		if err != nil {
// 			flag++
// 			if flag > maxNodeSelection {
// 				logs.Error("chain:%v,node:%v,GetTransactionByHash err", pro.id, info.sdk.url)
// 				flag = 0
// 				time.Sleep(time.Second)
// 			}
// 			info.latestHeight = 0
// 			info = pro.GetLatest()
// 		} else {
// 			flag = 0
// 			return tx, nil
// 		}
// 	}
// 	return nil, ErrNodesInvalid
// }

// func (pro *EthereumSdkPro) GetTransactionReceipt(hash common.Hash) (*types.Receipt, error) {
// 	info := pro.GetLatest()
// 	if info == nil {
// 		return nil, ErrNodesInvalid
// 	}
// 	flag := 0
// 	for info != nil {
// 		receipt, err := info.sdk.GetTransactionReceipt(hash)
// 		if err != nil {
// 			flag++
// 			if flag > maxNodeSelection {
// 				logs.Error("chain:%v,node:%v,GetTransactionReceipt err", pro.id, info.sdk.url)
// 				flag = 0
// 				time.Sleep(time.Second)
// 			}
// 			info.latestHeight = 0
// 			info = pro.GetLatest()
// 		} else {
// 			flag = 0
// 			return receipt, nil
// 		}
// 	}
// 	return nil, ErrNodesInvalid
// }

// func (pro *EthereumSdkPro) NonceAt(addr common.Address) (uint64, error) {
// 	info := pro.GetLatest()
// 	if info == nil {
// 		return 0, ErrNodesInvalid
// 	}

// 	for info != nil {
// 		nonce, err := info.sdk.NonceAt(addr)
// 		if err != nil {
// 			info.latestHeight = 0
// 			info = pro.GetLatest()
// 		} else {
// 			return nonce, nil
// 		}
// 	}
// 	return 0, ErrNodesInvalid
// }

// func (pro *EthereumSdkPro) SendRawTransaction(tx *types.Transaction) error {
// 	info := pro.GetLatest()
// 	if info == nil {
// 		return ErrNodesInvalid
// 	}

// 	for info != nil {
// 		err := info.sdk.SendRawTransaction(tx)
// 		if err != nil {
// 			info.latestHeight = 0
// 			info = pro.GetLatest()
// 		} else {
// 			return nil
// 		}
// 	}
// 	return ErrNodesInvalid
// }

// func (pro *EthereumSdkPro) TransactionByHash(hash common.Hash) (*types.Transaction, bool, error) {
// 	info := pro.GetLatest()
// 	if info == nil {
// 		return nil, false, ErrNodesInvalid
// 	}

// 	for info != nil {
// 		tx, isPending, err := info.sdk.TransactionByHash(hash)
// 		if err != nil {
// 			info.latestHeight = 0
// 			info = pro.GetLatest()
// 		} else {
// 			return tx, isPending, nil
// 		}
// 	}
// 	return nil, false, ErrNodesInvalid
// }

// func (pro *EthereumSdkPro) SuggestGasPrice() (*big.Int, error) {
// 	info := pro.GetLatest()
// 	if info == nil {
// 		return nil, ErrNodesInvalid
// 	}

// 	for info != nil {
// 		gas, err := info.sdk.SuggestGasPrice()
// 		if err != nil {
// 			info.latestHeight = 0
// 			info = pro.GetLatest()
// 		} else {
// 			return gas, nil
// 		}
// 	}
// 	return nil, ErrNodesInvalid
// }

// func (pro *EthereumSdkPro) EstimateGas(msg ethereum.CallMsg) (uint64, error) {
// 	info := pro.GetLatest()
// 	if info == nil {
// 		return 0, ErrNodesInvalid
// 	}

// 	for info != nil {
// 		gas, err := info.sdk.EstimateGas(msg)
// 		if err != nil {
// 			info.latestHeight = 0
// 			info = pro.GetLatest()
// 		} else {
// 			return gas, nil
// 		}
// 	}
// 	return 0, ErrNodesInvalid
// }

// //func (pro *EthereumSdkPro) Erc20Info(hash string) (string, string, int64, string, error) {
// //	info := pro.GetLatest()
// //	if info == nil {
// //		return "", "", 0, "", fmt.Errorf("all node is not working")
// //	}
// //	for info != nil {
// //		hash, name, decimal, symbol, err := info.sdk.Erc20Info(hash)
// //		if err != nil {
// //			info.latestHeight = 0
// //			info = pro.GetLatest()
// //		} else {
// //			return hash, name, decimal, symbol, nil
// //		}
// //	}
// //	return "", "", 0, "", fmt.Errorf("all node is not working")
// //}

// func (pro *EthereumSdkPro) IsEthAddress(addr string) bool {
// 	if addr == "0000000000000000000000000000000000000000" {
// 		return true
// 	} else {
// 		return false
// 	}
// }

// func (pro *EthereumSdkPro) Erc20Balance(erc20 string, addr string) (*big.Int, error) {
// 	info := pro.GetLatest()
// 	if info == nil {
// 		return new(big.Int).SetUint64(0), ErrNodesInvalid
// 	}
// 	erc20Address := common.HexToAddress(erc20)
// 	userAddress := common.HexToAddress(addr)
// 	for info != nil {
// 		balance := new(big.Int).SetUint64(0)
// 		var err error
// 		if pro.IsEthAddress(erc20) {
// 			balance, err = info.sdk.EthBalance(addr)
// 			//logs.Error("eth, addr: %s, balance: %s", addr, balance.String())
// 		} else {
// 			balance, err = info.sdk.GetERC20Balance(erc20Address, userAddress)
// 			//logs.Error("erc20, addr: %s, balance: %s", addr, balance.String())
// 		}
// 		if err != nil {
// 			info.latestHeight = 0
// 			info = pro.GetLatest()
// 		} else {
// 			return balance, nil
// 		}
// 	}
// 	return new(big.Int).SetUint64(0), ErrNodesInvalid
// }

// func (pro *EthereumSdkPro) Erc20TotalSupply(erc20 string) (*big.Int, error) {
// 	info := pro.GetLatest()
// 	if info == nil {
// 		return new(big.Int).SetUint64(0), ErrNodesInvalid
// 	}
// 	erc20Address := common.HexToAddress(erc20)
// 	for info != nil {
// 		totalSupply := new(big.Int).SetUint64(0)
// 		if pro.IsEthAddress(erc20) {
// 			return totalSupply, nil
// 		} else {
// 			totalSupply, err := info.sdk.GetERC20TotalSupply(erc20Address)
// 			if err != nil {
// 				info.latestHeight = 0
// 				info = pro.GetLatest()
// 			} else {
// 				return totalSupply, nil
// 			}
// 		}
// 	}
// 	return new(big.Int).SetUint64(0), ErrNodesInvalid
// }

// func (pro *EthereumSdkPro) WaitTransactionConfirm(hash common.Hash) bool {
// 	num := 0
// 	for num < 300 {
// 		time.Sleep(time.Second * 2)
// 		_, ispending, err := pro.TransactionByHash(hash)
// 		if err != nil {
// 			num++
// 			continue
// 		}
// 		if ispending {
// 			num++
// 			continue
// 		} else {
// 			return true
// 		}
// 	}
// 	return false
// }

// func (pro *EthereumSdkPro) reset(info *EthereumInfo) *EthereumInfo {
// 	info.latestHeight = 0
// 	return pro.GetLatest()
// }

// func (pro *EthereumSdkPro) FilterLog(FromBlock *big.Int, ToBlock *big.Int, Addresses []common.Address) ([]types.Log, error) {
// 	info := pro.GetLatest()
// 	if info == nil {
// 		return nil, fmt.Errorf("chain:%v FilterLog all node is not working", pro.id)
// 	}
// 	return info.sdk.FilterLog(FromBlock, ToBlock, Addresses)
// }
