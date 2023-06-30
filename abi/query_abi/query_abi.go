// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package query_abi

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

var (
	MethodCumulativeFundingRates = "cumulativeFundingRates"

	MethodGetDelta = "getDelta"

	MethodGetPosition = "getPosition"

	MethodGetPositionKey = "getPositionKey"

	MethodPositions = "positions"
)

// IVaultABI is the input ABI used to generate the binding from.
const IVaultABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"collateralToken\",\"type\":\"address\"}],\"name\":\"cumulativeFundingRates\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_indexToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_size\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_averagePrice\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_isLong\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"_lastIncreasedTime\",\"type\":\"uint256\"}],\"name\":\"getDelta\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_collateralToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_indexToken\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_isLong\",\"type\":\"bool\"}],\"name\":\"getPosition\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_collateralToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_indexToken\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_isLong\",\"type\":\"bool\"}],\"name\":\"getPositionKey\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"positionKey\",\"type\":\"bytes32\"}],\"name\":\"positions\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// IVaultFuncSigs maps the 4-byte function signature to its string representation.
var IVaultFuncSigs = map[string]string{
	"c65bc7b1": "cumulativeFundingRates(address)",
	"5c07eaab": "getDelta(address,uint256,uint256,bool,uint256)",
	"4a3f088d": "getPosition(address,address,address,bool)",
	"2d4b0576": "getPositionKey(address,address,address,bool)",
	"514ea4bf": "positions(bytes32)",
}

// IVault is an auto generated Go binding around an Ethereum contract.
type IVault struct {
	IVaultCaller     // Read-only binding to the contract
	IVaultTransactor // Write-only binding to the contract
	IVaultFilterer   // Log filterer for contract events
}

// IVaultCaller is an auto generated read-only Go binding around an Ethereum contract.
type IVaultCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IVaultTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IVaultTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IVaultFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IVaultFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IVaultSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IVaultSession struct {
	Contract     *IVault           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IVaultCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IVaultCallerSession struct {
	Contract *IVaultCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IVaultTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IVaultTransactorSession struct {
	Contract     *IVaultTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IVaultRaw is an auto generated low-level Go binding around an Ethereum contract.
type IVaultRaw struct {
	Contract *IVault // Generic contract binding to access the raw methods on
}

// IVaultCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IVaultCallerRaw struct {
	Contract *IVaultCaller // Generic read-only contract binding to access the raw methods on
}

// IVaultTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IVaultTransactorRaw struct {
	Contract *IVaultTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIVault creates a new instance of IVault, bound to a specific deployed contract.
func NewIVault(address common.Address, backend bind.ContractBackend) (*IVault, error) {
	contract, err := bindIVault(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IVault{IVaultCaller: IVaultCaller{contract: contract}, IVaultTransactor: IVaultTransactor{contract: contract}, IVaultFilterer: IVaultFilterer{contract: contract}}, nil
}

// NewIVaultCaller creates a new read-only instance of IVault, bound to a specific deployed contract.
func NewIVaultCaller(address common.Address, caller bind.ContractCaller) (*IVaultCaller, error) {
	contract, err := bindIVault(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IVaultCaller{contract: contract}, nil
}

// NewIVaultTransactor creates a new write-only instance of IVault, bound to a specific deployed contract.
func NewIVaultTransactor(address common.Address, transactor bind.ContractTransactor) (*IVaultTransactor, error) {
	contract, err := bindIVault(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IVaultTransactor{contract: contract}, nil
}

// NewIVaultFilterer creates a new log filterer instance of IVault, bound to a specific deployed contract.
func NewIVaultFilterer(address common.Address, filterer bind.ContractFilterer) (*IVaultFilterer, error) {
	contract, err := bindIVault(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IVaultFilterer{contract: contract}, nil
}

// bindIVault binds a generic wrapper to an already deployed contract.
func bindIVault(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IVaultABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IVault *IVaultRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IVault.Contract.IVaultCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IVault *IVaultRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IVault.Contract.IVaultTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IVault *IVaultRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IVault.Contract.IVaultTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IVault *IVaultCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IVault.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IVault *IVaultTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IVault.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IVault *IVaultTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IVault.Contract.contract.Transact(opts, method, params...)
}

// CumulativeFundingRates is a free data retrieval call binding the contract method 0xc65bc7b1.
//
// Solidity: function cumulativeFundingRates(address collateralToken) view returns(uint256)
func (_IVault *IVaultCaller) CumulativeFundingRates(opts *bind.CallOpts, collateralToken common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IVault.contract.Call(opts, &out, "cumulativeFundingRates", collateralToken)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CumulativeFundingRates is a free data retrieval call binding the contract method 0xc65bc7b1.
//
// Solidity: function cumulativeFundingRates(address collateralToken) view returns(uint256)
func (_IVault *IVaultSession) CumulativeFundingRates(collateralToken common.Address) (*big.Int, error) {
	return _IVault.Contract.CumulativeFundingRates(&_IVault.CallOpts, collateralToken)
}

// CumulativeFundingRates is a free data retrieval call binding the contract method 0xc65bc7b1.
//
// Solidity: function cumulativeFundingRates(address collateralToken) view returns(uint256)
func (_IVault *IVaultCallerSession) CumulativeFundingRates(collateralToken common.Address) (*big.Int, error) {
	return _IVault.Contract.CumulativeFundingRates(&_IVault.CallOpts, collateralToken)
}

// GetDelta is a free data retrieval call binding the contract method 0x5c07eaab.
//
// Solidity: function getDelta(address _indexToken, uint256 _size, uint256 _averagePrice, bool _isLong, uint256 _lastIncreasedTime) view returns(bool, uint256)
func (_IVault *IVaultCaller) GetDelta(opts *bind.CallOpts, _indexToken common.Address, _size *big.Int, _averagePrice *big.Int, _isLong bool, _lastIncreasedTime *big.Int) (bool, *big.Int, error) {
	var out []interface{}
	err := _IVault.contract.Call(opts, &out, "getDelta", _indexToken, _size, _averagePrice, _isLong, _lastIncreasedTime)

	if err != nil {
		return *new(bool), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetDelta is a free data retrieval call binding the contract method 0x5c07eaab.
//
// Solidity: function getDelta(address _indexToken, uint256 _size, uint256 _averagePrice, bool _isLong, uint256 _lastIncreasedTime) view returns(bool, uint256)
func (_IVault *IVaultSession) GetDelta(_indexToken common.Address, _size *big.Int, _averagePrice *big.Int, _isLong bool, _lastIncreasedTime *big.Int) (bool, *big.Int, error) {
	return _IVault.Contract.GetDelta(&_IVault.CallOpts, _indexToken, _size, _averagePrice, _isLong, _lastIncreasedTime)
}

// GetDelta is a free data retrieval call binding the contract method 0x5c07eaab.
//
// Solidity: function getDelta(address _indexToken, uint256 _size, uint256 _averagePrice, bool _isLong, uint256 _lastIncreasedTime) view returns(bool, uint256)
func (_IVault *IVaultCallerSession) GetDelta(_indexToken common.Address, _size *big.Int, _averagePrice *big.Int, _isLong bool, _lastIncreasedTime *big.Int) (bool, *big.Int, error) {
	return _IVault.Contract.GetDelta(&_IVault.CallOpts, _indexToken, _size, _averagePrice, _isLong, _lastIncreasedTime)
}

// GetPosition is a free data retrieval call binding the contract method 0x4a3f088d.
//
// Solidity: function getPosition(address _account, address _collateralToken, address _indexToken, bool _isLong) view returns(uint256, uint256, uint256, uint256, uint256, uint256, bool, uint256)
func (_IVault *IVaultCaller) GetPosition(opts *bind.CallOpts, _account common.Address, _collateralToken common.Address, _indexToken common.Address, _isLong bool) (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, bool, *big.Int, error) {
	var out []interface{}
	err := _IVault.contract.Call(opts, &out, "getPosition", _account, _collateralToken, _indexToken, _isLong)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(bool), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	out3 := *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	out4 := *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	out5 := *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	out6 := *abi.ConvertType(out[6], new(bool)).(*bool)
	out7 := *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)

	return out0, out1, out2, out3, out4, out5, out6, out7, err

}

// GetPosition is a free data retrieval call binding the contract method 0x4a3f088d.
//
// Solidity: function getPosition(address _account, address _collateralToken, address _indexToken, bool _isLong) view returns(uint256, uint256, uint256, uint256, uint256, uint256, bool, uint256)
func (_IVault *IVaultSession) GetPosition(_account common.Address, _collateralToken common.Address, _indexToken common.Address, _isLong bool) (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, bool, *big.Int, error) {
	return _IVault.Contract.GetPosition(&_IVault.CallOpts, _account, _collateralToken, _indexToken, _isLong)
}

// GetPosition is a free data retrieval call binding the contract method 0x4a3f088d.
//
// Solidity: function getPosition(address _account, address _collateralToken, address _indexToken, bool _isLong) view returns(uint256, uint256, uint256, uint256, uint256, uint256, bool, uint256)
func (_IVault *IVaultCallerSession) GetPosition(_account common.Address, _collateralToken common.Address, _indexToken common.Address, _isLong bool) (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, bool, *big.Int, error) {
	return _IVault.Contract.GetPosition(&_IVault.CallOpts, _account, _collateralToken, _indexToken, _isLong)
}

// GetPositionKey is a free data retrieval call binding the contract method 0x2d4b0576.
//
// Solidity: function getPositionKey(address _account, address _collateralToken, address _indexToken, bool _isLong) pure returns(bytes32)
func (_IVault *IVaultCaller) GetPositionKey(opts *bind.CallOpts, _account common.Address, _collateralToken common.Address, _indexToken common.Address, _isLong bool) ([32]byte, error) {
	var out []interface{}
	err := _IVault.contract.Call(opts, &out, "getPositionKey", _account, _collateralToken, _indexToken, _isLong)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetPositionKey is a free data retrieval call binding the contract method 0x2d4b0576.
//
// Solidity: function getPositionKey(address _account, address _collateralToken, address _indexToken, bool _isLong) pure returns(bytes32)
func (_IVault *IVaultSession) GetPositionKey(_account common.Address, _collateralToken common.Address, _indexToken common.Address, _isLong bool) ([32]byte, error) {
	return _IVault.Contract.GetPositionKey(&_IVault.CallOpts, _account, _collateralToken, _indexToken, _isLong)
}

// GetPositionKey is a free data retrieval call binding the contract method 0x2d4b0576.
//
// Solidity: function getPositionKey(address _account, address _collateralToken, address _indexToken, bool _isLong) pure returns(bytes32)
func (_IVault *IVaultCallerSession) GetPositionKey(_account common.Address, _collateralToken common.Address, _indexToken common.Address, _isLong bool) ([32]byte, error) {
	return _IVault.Contract.GetPositionKey(&_IVault.CallOpts, _account, _collateralToken, _indexToken, _isLong)
}

// Positions is a free data retrieval call binding the contract method 0x514ea4bf.
//
// Solidity: function positions(bytes32 positionKey) view returns(uint256, uint256, uint256, uint256, uint256, int256, uint256)
func (_IVault *IVaultCaller) Positions(opts *bind.CallOpts, positionKey [32]byte) (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _IVault.contract.Call(opts, &out, "positions", positionKey)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	out3 := *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	out4 := *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	out5 := *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	out6 := *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)

	return out0, out1, out2, out3, out4, out5, out6, err

}

// Positions is a free data retrieval call binding the contract method 0x514ea4bf.
//
// Solidity: function positions(bytes32 positionKey) view returns(uint256, uint256, uint256, uint256, uint256, int256, uint256)
func (_IVault *IVaultSession) Positions(positionKey [32]byte) (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _IVault.Contract.Positions(&_IVault.CallOpts, positionKey)
}

// Positions is a free data retrieval call binding the contract method 0x514ea4bf.
//
// Solidity: function positions(bytes32 positionKey) view returns(uint256, uint256, uint256, uint256, uint256, int256, uint256)
func (_IVault *IVaultCallerSession) Positions(positionKey [32]byte) (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _IVault.Contract.Positions(&_IVault.CallOpts, positionKey)
}

var (
	MethodFUNDINGRATEPRECISION = "FUNDING_RATE_PRECISION"

	MethodVault = "vault"
)

// QueryABI is the input ABI used to generate the binding from.
const QueryABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"vaultAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"FUNDING_RATE_PRECISION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_collateralToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_indexToken\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_isLong\",\"type\":\"bool\"}],\"name\":\"getPosition\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"size\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"averagePrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"entryFundingRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveAmount\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"realisedPnl\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"lastIncreasedTime\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"floatingPnl\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"uncollectedFundingFee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vault\",\"outputs\":[{\"internalType\":\"contractIVault\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// QueryFuncSigs maps the 4-byte function signature to its string representation.
var QueryFuncSigs = map[string]string{
	"6be6026b": "FUNDING_RATE_PRECISION()",
	"4a3f088d": "getPosition(address,address,address,bool)",
	"fbfa77cf": "vault()",
}

// QueryBin is the compiled bytecode used for deploying new contracts.
var QueryBin = "0x608060405234801561001057600080fd5b506040516107b13803806107b18339818101604052602081101561003357600080fd5b5051600080546001600160a01b039092166001600160a01b031990921691909117905561074c806100656000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c80634a3f088d146100465780636be6026b146100cb578063fbfa77cf146100e5575b600080fd5b6100846004803603608081101561005c57600080fd5b506001600160a01b038135811691602081013582169160408201351690606001351515610109565b60408051998a5260208a0198909852888801969096526060880194909452608087019290925260a086015260c085015260e084015261010083015251908190036101200190f35b6100d36101bb565b60408051918252519081900360200190f35b6100ed6101c2565b604080516001600160a01b039092168252519081900360200190f35b60008060008060008060008060006101238d8d8d8d6101d1565b80602001905161012081101561013857600080fd5b810190808051906020019092919080519060200190929190805190602001909291908051906020019092919080519060200190929190805190602001909291908051906020019092919080519060200190929190805190602001909291905050509850985098509850985098509850985098509499509499509499509499909550565b620f424081565b6000546001600160a01b031681565b60008054604080516316a582bb60e11b81526001600160a01b0388811660048301528781166024830152868116604483015285151560648301529151606094859490931691632d4b0576916084808301926020929190829003018186803b15801561023b57600080fd5b505afa15801561024f573d6000803e3d6000fd5b505050506040513d602081101561026557600080fd5b50516000805460408051602480820186905282518083039091018152604490910182526020810180516001600160e01b031663514ea4bf60e01b1781529151815195965093946060946001600160a01b0390941693919282918083835b602083106102e15780518252601f1990920191602091820191016102c2565b6001836020036101000a038019825116818451168082178552505050505050905001915050600060405180830381855afa9150503d8060008114610341576040519150601f19603f3d011682016040523d82523d6000602084013e610346565b606091505b509150915081610394576040805162461bcd60e51b81526020600482015260146024820152733330b4b6103a379033b2ba103837b9b4ba34b7b760611b604482015290519081900360640190fd5b8093505050506103a58185856103bb565b90506103b18186610583565b9695505050505050565b606060008060008680602001905160e08110156103d757600080fd5b50805160408083015160c090930151600080548351635c07eaab60e01b81526001600160a01b038d8116600483015260248201879052604482018890528c15156064830152608482018590528551969a5096985092965090948594911692635c07eaab9260a48082019391829003018186803b15801561045657600080fd5b505afa15801561046a573d6000803e3d6000fd5b505050506040513d604081101561048057600080fd5b508051602090910151909250905060008261049d576000196104a0565b60015b820290508981604051602001808281526020019150506040516020818303038152906040526040516020018083805190602001908083835b602083106104f75780518252601f1990920191602091820191016104d8565b51815160209384036101000a600019018019909216911617905285519190930192850191508083835b6020831061053f5780518252601f199092019160209182019101610520565b6001836020036101000a0380198251168184511680821785525050505050509050019250505060405160208183030381529060405296505050505050509392505050565b60606000808480602001905161010081101561059e57600080fd5b508051606090910151600080546040805163c65bc7b160e01b81526001600160a01b038a81166004830152915195975093955091938593929091169163c65bc7b1916024808301926020929190829003018186803b1580156105ff57600080fd5b505afa158015610613573d6000803e3d6000fd5b505050506040513d602081101561062957600080fd5b50510390506000620f42408483020490508681604051602001808281526020019150506040516020818303038152906040526040516020018083805190602001908083835b6020831061068d5780518252601f19909201916020918201910161066e565b51815160209384036101000a600019018019909216911617905285519190930192850191508083835b602083106106d55780518252601f1990920191602091820191016106b6565b6001836020036101000a038019825116818451168082178552505050505050905001925050506040516020818303038152906040529450505050509291505056fea26469706673582212203d3350fd8cc2baf7e301d1c4518f1182fa642915ee8c65e0cbbc02e5a96688cb64736f6c634300060c0033"

// DeployQuery deploys a new Ethereum contract, binding an instance of Query to it.
func DeployQuery(auth *bind.TransactOpts, backend bind.ContractBackend, vaultAddress common.Address) (common.Address, *types.Transaction, *Query, error) {
	parsed, err := abi.JSON(strings.NewReader(QueryABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(QueryBin), backend, vaultAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Query{QueryCaller: QueryCaller{contract: contract}, QueryTransactor: QueryTransactor{contract: contract}, QueryFilterer: QueryFilterer{contract: contract}}, nil
}

// Query is an auto generated Go binding around an Ethereum contract.
type Query struct {
	QueryCaller     // Read-only binding to the contract
	QueryTransactor // Write-only binding to the contract
	QueryFilterer   // Log filterer for contract events
}

// QueryCaller is an auto generated read-only Go binding around an Ethereum contract.
type QueryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// QueryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type QueryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// QueryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type QueryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// QuerySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type QuerySession struct {
	Contract     *Query            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// QueryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type QueryCallerSession struct {
	Contract *QueryCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// QueryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type QueryTransactorSession struct {
	Contract     *QueryTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// QueryRaw is an auto generated low-level Go binding around an Ethereum contract.
type QueryRaw struct {
	Contract *Query // Generic contract binding to access the raw methods on
}

// QueryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type QueryCallerRaw struct {
	Contract *QueryCaller // Generic read-only contract binding to access the raw methods on
}

// QueryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type QueryTransactorRaw struct {
	Contract *QueryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewQuery creates a new instance of Query, bound to a specific deployed contract.
func NewQuery(address common.Address, backend bind.ContractBackend) (*Query, error) {
	contract, err := bindQuery(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Query{QueryCaller: QueryCaller{contract: contract}, QueryTransactor: QueryTransactor{contract: contract}, QueryFilterer: QueryFilterer{contract: contract}}, nil
}

// NewQueryCaller creates a new read-only instance of Query, bound to a specific deployed contract.
func NewQueryCaller(address common.Address, caller bind.ContractCaller) (*QueryCaller, error) {
	contract, err := bindQuery(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &QueryCaller{contract: contract}, nil
}

// NewQueryTransactor creates a new write-only instance of Query, bound to a specific deployed contract.
func NewQueryTransactor(address common.Address, transactor bind.ContractTransactor) (*QueryTransactor, error) {
	contract, err := bindQuery(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &QueryTransactor{contract: contract}, nil
}

// NewQueryFilterer creates a new log filterer instance of Query, bound to a specific deployed contract.
func NewQueryFilterer(address common.Address, filterer bind.ContractFilterer) (*QueryFilterer, error) {
	contract, err := bindQuery(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &QueryFilterer{contract: contract}, nil
}

// bindQuery binds a generic wrapper to an already deployed contract.
func bindQuery(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(QueryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Query *QueryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Query.Contract.QueryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Query *QueryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Query.Contract.QueryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Query *QueryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Query.Contract.QueryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Query *QueryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Query.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Query *QueryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Query.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Query *QueryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Query.Contract.contract.Transact(opts, method, params...)
}

// FUNDINGRATEPRECISION is a free data retrieval call binding the contract method 0x6be6026b.
//
// Solidity: function FUNDING_RATE_PRECISION() view returns(uint256)
func (_Query *QueryCaller) FUNDINGRATEPRECISION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Query.contract.Call(opts, &out, "FUNDING_RATE_PRECISION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FUNDINGRATEPRECISION is a free data retrieval call binding the contract method 0x6be6026b.
//
// Solidity: function FUNDING_RATE_PRECISION() view returns(uint256)
func (_Query *QuerySession) FUNDINGRATEPRECISION() (*big.Int, error) {
	return _Query.Contract.FUNDINGRATEPRECISION(&_Query.CallOpts)
}

// FUNDINGRATEPRECISION is a free data retrieval call binding the contract method 0x6be6026b.
//
// Solidity: function FUNDING_RATE_PRECISION() view returns(uint256)
func (_Query *QueryCallerSession) FUNDINGRATEPRECISION() (*big.Int, error) {
	return _Query.Contract.FUNDINGRATEPRECISION(&_Query.CallOpts)
}

// GetPosition is a free data retrieval call binding the contract method 0x4a3f088d.
//
// Solidity: function getPosition(address _account, address _collateralToken, address _indexToken, bool _isLong) view returns(uint256 size, uint256 collateral, uint256 averagePrice, uint256 entryFundingRate, uint256 reserveAmount, int256 realisedPnl, uint256 lastIncreasedTime, int256 floatingPnl, uint256 uncollectedFundingFee)
func (_Query *QueryCaller) GetPosition(opts *bind.CallOpts, _account common.Address, _collateralToken common.Address, _indexToken common.Address, _isLong bool) (struct {
	Size                  *big.Int
	Collateral            *big.Int
	AveragePrice          *big.Int
	EntryFundingRate      *big.Int
	ReserveAmount         *big.Int
	RealisedPnl           *big.Int
	LastIncreasedTime     *big.Int
	FloatingPnl           *big.Int
	UncollectedFundingFee *big.Int
}, error) {
	var out []interface{}
	err := _Query.contract.Call(opts, &out, "getPosition", _account, _collateralToken, _indexToken, _isLong)

	outstruct := new(struct {
		Size                  *big.Int
		Collateral            *big.Int
		AveragePrice          *big.Int
		EntryFundingRate      *big.Int
		ReserveAmount         *big.Int
		RealisedPnl           *big.Int
		LastIncreasedTime     *big.Int
		FloatingPnl           *big.Int
		UncollectedFundingFee *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Size = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Collateral = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.AveragePrice = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.EntryFundingRate = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.ReserveAmount = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.RealisedPnl = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.LastIncreasedTime = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.FloatingPnl = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)
	outstruct.UncollectedFundingFee = *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetPosition is a free data retrieval call binding the contract method 0x4a3f088d.
//
// Solidity: function getPosition(address _account, address _collateralToken, address _indexToken, bool _isLong) view returns(uint256 size, uint256 collateral, uint256 averagePrice, uint256 entryFundingRate, uint256 reserveAmount, int256 realisedPnl, uint256 lastIncreasedTime, int256 floatingPnl, uint256 uncollectedFundingFee)
func (_Query *QuerySession) GetPosition(_account common.Address, _collateralToken common.Address, _indexToken common.Address, _isLong bool) (struct {
	Size                  *big.Int
	Collateral            *big.Int
	AveragePrice          *big.Int
	EntryFundingRate      *big.Int
	ReserveAmount         *big.Int
	RealisedPnl           *big.Int
	LastIncreasedTime     *big.Int
	FloatingPnl           *big.Int
	UncollectedFundingFee *big.Int
}, error) {
	return _Query.Contract.GetPosition(&_Query.CallOpts, _account, _collateralToken, _indexToken, _isLong)
}

// GetPosition is a free data retrieval call binding the contract method 0x4a3f088d.
//
// Solidity: function getPosition(address _account, address _collateralToken, address _indexToken, bool _isLong) view returns(uint256 size, uint256 collateral, uint256 averagePrice, uint256 entryFundingRate, uint256 reserveAmount, int256 realisedPnl, uint256 lastIncreasedTime, int256 floatingPnl, uint256 uncollectedFundingFee)
func (_Query *QueryCallerSession) GetPosition(_account common.Address, _collateralToken common.Address, _indexToken common.Address, _isLong bool) (struct {
	Size                  *big.Int
	Collateral            *big.Int
	AveragePrice          *big.Int
	EntryFundingRate      *big.Int
	ReserveAmount         *big.Int
	RealisedPnl           *big.Int
	LastIncreasedTime     *big.Int
	FloatingPnl           *big.Int
	UncollectedFundingFee *big.Int
}, error) {
	return _Query.Contract.GetPosition(&_Query.CallOpts, _account, _collateralToken, _indexToken, _isLong)
}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_Query *QueryCaller) Vault(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Query.contract.Call(opts, &out, "vault")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_Query *QuerySession) Vault() (common.Address, error) {
	return _Query.Contract.Vault(&_Query.CallOpts)
}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_Query *QueryCallerSession) Vault() (common.Address, error) {
	return _Query.Contract.Vault(&_Query.CallOpts)
}

