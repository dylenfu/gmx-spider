// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package vault_abi

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

var ()

// AddressABI is the input ABI used to generate the binding from.
const AddressABI = "[]"

// AddressBin is the compiled bytecode used for deploying new contracts.
var AddressBin = "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212206e6d0bb3e4585a29e1dddcf6db3a68e21cf7861bde835a417547858c6d108d1c64736f6c634300060c0033"

// DeployAddress deploys a new Ethereum contract, binding an instance of Address to it.
func DeployAddress(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Address, error) {
	parsed, err := abi.JSON(strings.NewReader(AddressABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(AddressBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Address{AddressCaller: AddressCaller{contract: contract}, AddressTransactor: AddressTransactor{contract: contract}, AddressFilterer: AddressFilterer{contract: contract}}, nil
}

// Address is an auto generated Go binding around an Ethereum contract.
type Address struct {
	AddressCaller     // Read-only binding to the contract
	AddressTransactor // Write-only binding to the contract
	AddressFilterer   // Log filterer for contract events
}

// AddressCaller is an auto generated read-only Go binding around an Ethereum contract.
type AddressCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AddressTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AddressFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AddressSession struct {
	Contract     *Address          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AddressCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AddressCallerSession struct {
	Contract *AddressCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// AddressTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AddressTransactorSession struct {
	Contract     *AddressTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// AddressRaw is an auto generated low-level Go binding around an Ethereum contract.
type AddressRaw struct {
	Contract *Address // Generic contract binding to access the raw methods on
}

// AddressCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AddressCallerRaw struct {
	Contract *AddressCaller // Generic read-only contract binding to access the raw methods on
}

// AddressTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AddressTransactorRaw struct {
	Contract *AddressTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAddress creates a new instance of Address, bound to a specific deployed contract.
func NewAddress(address common.Address, backend bind.ContractBackend) (*Address, error) {
	contract, err := bindAddress(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Address{AddressCaller: AddressCaller{contract: contract}, AddressTransactor: AddressTransactor{contract: contract}, AddressFilterer: AddressFilterer{contract: contract}}, nil
}

// NewAddressCaller creates a new read-only instance of Address, bound to a specific deployed contract.
func NewAddressCaller(address common.Address, caller bind.ContractCaller) (*AddressCaller, error) {
	contract, err := bindAddress(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AddressCaller{contract: contract}, nil
}

// NewAddressTransactor creates a new write-only instance of Address, bound to a specific deployed contract.
func NewAddressTransactor(address common.Address, transactor bind.ContractTransactor) (*AddressTransactor, error) {
	contract, err := bindAddress(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AddressTransactor{contract: contract}, nil
}

// NewAddressFilterer creates a new log filterer instance of Address, bound to a specific deployed contract.
func NewAddressFilterer(address common.Address, filterer bind.ContractFilterer) (*AddressFilterer, error) {
	contract, err := bindAddress(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AddressFilterer{contract: contract}, nil
}

// bindAddress binds a generic wrapper to an already deployed contract.
func bindAddress(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AddressABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Address *AddressRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Address.Contract.AddressCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Address *AddressRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Address.Contract.AddressTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Address *AddressRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Address.Contract.AddressTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Address *AddressCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Address.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Address *AddressTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Address.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Address *AddressTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Address.Contract.contract.Transact(opts, method, params...)
}

var (
	MethodApprove = "approve"

	MethodTransfer = "transfer"

	MethodTransferFrom = "transferFrom"

	MethodAllowance = "allowance"

	MethodBalanceOf = "balanceOf"

	MethodTotalSupply = "totalSupply"

	EventApproval = "Approval"

	EventTransfer = "Transfer"
)

// IERC20ABI is the input ABI used to generate the binding from.
const IERC20ABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IERC20FuncSigs maps the 4-byte function signature to its string representation.
var IERC20FuncSigs = map[string]string{
	"dd62ed3e": "allowance(address,address)",
	"095ea7b3": "approve(address,uint256)",
	"70a08231": "balanceOf(address)",
	"18160ddd": "totalSupply()",
	"a9059cbb": "transfer(address,uint256)",
	"23b872dd": "transferFrom(address,address,uint256)",
}

// IERC20 is an auto generated Go binding around an Ethereum contract.
type IERC20 struct {
	IERC20Caller     // Read-only binding to the contract
	IERC20Transactor // Write-only binding to the contract
	IERC20Filterer   // Log filterer for contract events
}

// IERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type IERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC20Session struct {
	Contract     *IERC20           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC20CallerSession struct {
	Contract *IERC20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC20TransactorSession struct {
	Contract     *IERC20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type IERC20Raw struct {
	Contract *IERC20 // Generic contract binding to access the raw methods on
}

// IERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC20CallerRaw struct {
	Contract *IERC20Caller // Generic read-only contract binding to access the raw methods on
}

// IERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC20TransactorRaw struct {
	Contract *IERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC20 creates a new instance of IERC20, bound to a specific deployed contract.
func NewIERC20(address common.Address, backend bind.ContractBackend) (*IERC20, error) {
	contract, err := bindIERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC20{IERC20Caller: IERC20Caller{contract: contract}, IERC20Transactor: IERC20Transactor{contract: contract}, IERC20Filterer: IERC20Filterer{contract: contract}}, nil
}

// NewIERC20Caller creates a new read-only instance of IERC20, bound to a specific deployed contract.
func NewIERC20Caller(address common.Address, caller bind.ContractCaller) (*IERC20Caller, error) {
	contract, err := bindIERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20Caller{contract: contract}, nil
}

// NewIERC20Transactor creates a new write-only instance of IERC20, bound to a specific deployed contract.
func NewIERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*IERC20Transactor, error) {
	contract, err := bindIERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20Transactor{contract: contract}, nil
}

// NewIERC20Filterer creates a new log filterer instance of IERC20, bound to a specific deployed contract.
func NewIERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*IERC20Filterer, error) {
	contract, err := bindIERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC20Filterer{contract: contract}, nil
}

// bindIERC20 binds a generic wrapper to an already deployed contract.
func bindIERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IERC20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20 *IERC20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC20.Contract.IERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20 *IERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20.Contract.IERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20 *IERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20.Contract.IERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20 *IERC20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20 *IERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20 *IERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20 *IERC20Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20 *IERC20Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC20.Contract.Allowance(&_IERC20.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20 *IERC20CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC20.Contract.Allowance(&_IERC20.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20 *IERC20Caller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20 *IERC20Session) BalanceOf(account common.Address) (*big.Int, error) {
	return _IERC20.Contract.BalanceOf(&_IERC20.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20 *IERC20CallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _IERC20.Contract.BalanceOf(&_IERC20.CallOpts, account)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20 *IERC20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20 *IERC20Session) TotalSupply() (*big.Int, error) {
	return _IERC20.Contract.TotalSupply(&_IERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20 *IERC20CallerSession) TotalSupply() (*big.Int, error) {
	return _IERC20.Contract.TotalSupply(&_IERC20.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20 *IERC20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20 *IERC20Session) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Approve(&_IERC20.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20 *IERC20TransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Approve(&_IERC20.TransactOpts, spender, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Transactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Session) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Transfer(&_IERC20.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20TransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Transfer(&_IERC20.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Transactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Session) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.TransferFrom(&_IERC20.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20TransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.TransferFrom(&_IERC20.TransactOpts, sender, recipient, amount)
}

// IERC20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the IERC20 contract.
type IERC20ApprovalIterator struct {
	Event *IERC20Approval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IERC20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20Approval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IERC20Approval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IERC20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20Approval represents a Approval event raised by the IERC20 contract.
type IERC20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*IERC20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &IERC20ApprovalIterator{contract: _IERC20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *IERC20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20Approval)
				if err := _IERC20.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) ParseApproval(log types.Log) (*IERC20Approval, error) {
	event := new(IERC20Approval)
	if err := _IERC20.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the IERC20 contract.
type IERC20TransferIterator struct {
	Event *IERC20Transfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IERC20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20Transfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IERC20Transfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IERC20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20Transfer represents a Transfer event raised by the IERC20 contract.
type IERC20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*IERC20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IERC20TransferIterator{contract: _IERC20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *IERC20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20Transfer)
				if err := _IERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) ParseTransfer(log types.Log) (*IERC20Transfer, error) {
	event := new(IERC20Transfer)
	if err := _IERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

var (
	MethodAddVault = "addVault"

	MethodBurn = "burn"

	MethodMint = "mint"

	MethodRemoveVault = "removeVault"
)

// IUSDGABI is the input ABI used to generate the binding from.
const IUSDGABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_vault\",\"type\":\"address\"}],\"name\":\"addVault\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_vault\",\"type\":\"address\"}],\"name\":\"removeVault\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IUSDGFuncSigs maps the 4-byte function signature to its string representation.
var IUSDGFuncSigs = map[string]string{
	"256b5a02": "addVault(address)",
	"9dc29fac": "burn(address,uint256)",
	"40c10f19": "mint(address,uint256)",
	"ceb68c23": "removeVault(address)",
}

// IUSDG is an auto generated Go binding around an Ethereum contract.
type IUSDG struct {
	IUSDGCaller     // Read-only binding to the contract
	IUSDGTransactor // Write-only binding to the contract
	IUSDGFilterer   // Log filterer for contract events
}

// IUSDGCaller is an auto generated read-only Go binding around an Ethereum contract.
type IUSDGCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IUSDGTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IUSDGTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IUSDGFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IUSDGFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IUSDGSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IUSDGSession struct {
	Contract     *IUSDG            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IUSDGCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IUSDGCallerSession struct {
	Contract *IUSDGCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IUSDGTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IUSDGTransactorSession struct {
	Contract     *IUSDGTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IUSDGRaw is an auto generated low-level Go binding around an Ethereum contract.
type IUSDGRaw struct {
	Contract *IUSDG // Generic contract binding to access the raw methods on
}

// IUSDGCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IUSDGCallerRaw struct {
	Contract *IUSDGCaller // Generic read-only contract binding to access the raw methods on
}

// IUSDGTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IUSDGTransactorRaw struct {
	Contract *IUSDGTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIUSDG creates a new instance of IUSDG, bound to a specific deployed contract.
func NewIUSDG(address common.Address, backend bind.ContractBackend) (*IUSDG, error) {
	contract, err := bindIUSDG(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IUSDG{IUSDGCaller: IUSDGCaller{contract: contract}, IUSDGTransactor: IUSDGTransactor{contract: contract}, IUSDGFilterer: IUSDGFilterer{contract: contract}}, nil
}

// NewIUSDGCaller creates a new read-only instance of IUSDG, bound to a specific deployed contract.
func NewIUSDGCaller(address common.Address, caller bind.ContractCaller) (*IUSDGCaller, error) {
	contract, err := bindIUSDG(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IUSDGCaller{contract: contract}, nil
}

// NewIUSDGTransactor creates a new write-only instance of IUSDG, bound to a specific deployed contract.
func NewIUSDGTransactor(address common.Address, transactor bind.ContractTransactor) (*IUSDGTransactor, error) {
	contract, err := bindIUSDG(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IUSDGTransactor{contract: contract}, nil
}

// NewIUSDGFilterer creates a new log filterer instance of IUSDG, bound to a specific deployed contract.
func NewIUSDGFilterer(address common.Address, filterer bind.ContractFilterer) (*IUSDGFilterer, error) {
	contract, err := bindIUSDG(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IUSDGFilterer{contract: contract}, nil
}

// bindIUSDG binds a generic wrapper to an already deployed contract.
func bindIUSDG(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IUSDGABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IUSDG *IUSDGRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IUSDG.Contract.IUSDGCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IUSDG *IUSDGRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IUSDG.Contract.IUSDGTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IUSDG *IUSDGRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IUSDG.Contract.IUSDGTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IUSDG *IUSDGCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IUSDG.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IUSDG *IUSDGTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IUSDG.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IUSDG *IUSDGTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IUSDG.Contract.contract.Transact(opts, method, params...)
}

// AddVault is a paid mutator transaction binding the contract method 0x256b5a02.
//
// Solidity: function addVault(address _vault) returns()
func (_IUSDG *IUSDGTransactor) AddVault(opts *bind.TransactOpts, _vault common.Address) (*types.Transaction, error) {
	return _IUSDG.contract.Transact(opts, "addVault", _vault)
}

// AddVault is a paid mutator transaction binding the contract method 0x256b5a02.
//
// Solidity: function addVault(address _vault) returns()
func (_IUSDG *IUSDGSession) AddVault(_vault common.Address) (*types.Transaction, error) {
	return _IUSDG.Contract.AddVault(&_IUSDG.TransactOpts, _vault)
}

// AddVault is a paid mutator transaction binding the contract method 0x256b5a02.
//
// Solidity: function addVault(address _vault) returns()
func (_IUSDG *IUSDGTransactorSession) AddVault(_vault common.Address) (*types.Transaction, error) {
	return _IUSDG.Contract.AddVault(&_IUSDG.TransactOpts, _vault)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address _account, uint256 _amount) returns()
func (_IUSDG *IUSDGTransactor) Burn(opts *bind.TransactOpts, _account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IUSDG.contract.Transact(opts, "burn", _account, _amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address _account, uint256 _amount) returns()
func (_IUSDG *IUSDGSession) Burn(_account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IUSDG.Contract.Burn(&_IUSDG.TransactOpts, _account, _amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address _account, uint256 _amount) returns()
func (_IUSDG *IUSDGTransactorSession) Burn(_account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IUSDG.Contract.Burn(&_IUSDG.TransactOpts, _account, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _account, uint256 _amount) returns()
func (_IUSDG *IUSDGTransactor) Mint(opts *bind.TransactOpts, _account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IUSDG.contract.Transact(opts, "mint", _account, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _account, uint256 _amount) returns()
func (_IUSDG *IUSDGSession) Mint(_account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IUSDG.Contract.Mint(&_IUSDG.TransactOpts, _account, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _account, uint256 _amount) returns()
func (_IUSDG *IUSDGTransactorSession) Mint(_account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IUSDG.Contract.Mint(&_IUSDG.TransactOpts, _account, _amount)
}

// RemoveVault is a paid mutator transaction binding the contract method 0xceb68c23.
//
// Solidity: function removeVault(address _vault) returns()
func (_IUSDG *IUSDGTransactor) RemoveVault(opts *bind.TransactOpts, _vault common.Address) (*types.Transaction, error) {
	return _IUSDG.contract.Transact(opts, "removeVault", _vault)
}

// RemoveVault is a paid mutator transaction binding the contract method 0xceb68c23.
//
// Solidity: function removeVault(address _vault) returns()
func (_IUSDG *IUSDGSession) RemoveVault(_vault common.Address) (*types.Transaction, error) {
	return _IUSDG.Contract.RemoveVault(&_IUSDG.TransactOpts, _vault)
}

// RemoveVault is a paid mutator transaction binding the contract method 0xceb68c23.
//
// Solidity: function removeVault(address _vault) returns()
func (_IUSDG *IUSDGTransactorSession) RemoveVault(_vault common.Address) (*types.Transaction, error) {
	return _IUSDG.Contract.RemoveVault(&_IUSDG.TransactOpts, _vault)
}

var (
	MethodBuyUSDG = "buyUSDG"

	MethodDecreasePosition = "decreasePosition"

	MethodDirectPoolDeposit = "directPoolDeposit"

	MethodIncreasePosition = "increasePosition"

	MethodLiquidatePosition = "liquidatePosition"

	MethodSellUSDG = "sellUSDG"

	MethodSetBufferAmount = "setBufferAmount"

	MethodSetError = "setError"

	MethodSetFees = "setFees"

	MethodSetFundingRate = "setFundingRate"

	MethodSetInManagerMode = "setInManagerMode"

	MethodSetInPrivateLiquidationMode = "setInPrivateLiquidationMode"

	MethodSetIsLeverageEnabled = "setIsLeverageEnabled"

	MethodSetIsSwapEnabled = "setIsSwapEnabled"

	MethodSetLiquidator = "setLiquidator"

	MethodSetManager = "setManager"

	MethodSetMaxGasPrice = "setMaxGasPrice"

	MethodSetMaxGlobalShortSize = "setMaxGlobalShortSize"

	MethodSetMaxLeverage = "setMaxLeverage"

	MethodSetPriceFeed = "setPriceFeed"

	MethodSetTokenConfig = "setTokenConfig"

	MethodSetUsdgAmount = "setUsdgAmount"

	MethodSetVaultUtils = "setVaultUtils"

	MethodSwap = "swap"

	MethodWithdrawFees = "withdrawFees"

	MethodAllWhitelistedTokens = "allWhitelistedTokens"

	MethodAllWhitelistedTokensLength = "allWhitelistedTokensLength"

	MethodApprovedRouters = "approvedRouters"

	MethodBufferAmounts = "bufferAmounts"

	MethodCumulativeFundingRates = "cumulativeFundingRates"

	MethodFeeReserves = "feeReserves"

	MethodFundingInterval = "fundingInterval"

	MethodFundingRateFactor = "fundingRateFactor"

	MethodGetDelta = "getDelta"

	MethodGetFeeBasisPoints = "getFeeBasisPoints"

	MethodGetMaxPrice = "getMaxPrice"

	MethodGetMinPrice = "getMinPrice"

	MethodGetNextFundingRate = "getNextFundingRate"

	MethodGetPosition = "getPosition"

	MethodGetRedemptionAmount = "getRedemptionAmount"

	MethodGetTargetUsdgAmount = "getTargetUsdgAmount"

	MethodGlobalShortAveragePrices = "globalShortAveragePrices"

	MethodGlobalShortSizes = "globalShortSizes"

	MethodGov = "gov"

	MethodGuaranteedUsd = "guaranteedUsd"

	MethodHasDynamicFees = "hasDynamicFees"

	MethodInManagerMode = "inManagerMode"

	MethodInPrivateLiquidationMode = "inPrivateLiquidationMode"

	MethodIsInitialized = "isInitialized"

	MethodIsLeverageEnabled = "isLeverageEnabled"

	MethodIsLiquidator = "isLiquidator"

	MethodIsManager = "isManager"

	MethodIsSwapEnabled = "isSwapEnabled"

	MethodLastFundingTimes = "lastFundingTimes"

	MethodLiquidationFeeUsd = "liquidationFeeUsd"

	MethodMarginFeeBasisPoints = "marginFeeBasisPoints"

	MethodMaxGasPrice = "maxGasPrice"

	MethodMaxGlobalShortSizes = "maxGlobalShortSizes"

	MethodMaxLeverage = "maxLeverage"

	MethodMaxUsdgAmounts = "maxUsdgAmounts"

	MethodMinProfitBasisPoints = "minProfitBasisPoints"

	MethodMinProfitTime = "minProfitTime"

	MethodMintBurnFeeBasisPoints = "mintBurnFeeBasisPoints"

	MethodPoolAmounts = "poolAmounts"

	MethodPriceFeed = "priceFeed"

	MethodReservedAmounts = "reservedAmounts"

	MethodRouter = "router"

	MethodShortableTokens = "shortableTokens"

	MethodStableFundingRateFactor = "stableFundingRateFactor"

	MethodStableSwapFeeBasisPoints = "stableSwapFeeBasisPoints"

	MethodStableTaxBasisPoints = "stableTaxBasisPoints"

	MethodStableTokens = "stableTokens"

	MethodSwapFeeBasisPoints = "swapFeeBasisPoints"

	MethodTaxBasisPoints = "taxBasisPoints"

	MethodTokenBalances = "tokenBalances"

	MethodTokenDecimals = "tokenDecimals"

	MethodTokenToUsdMin = "tokenToUsdMin"

	MethodTokenWeights = "tokenWeights"

	MethodTotalTokenWeights = "totalTokenWeights"

	MethodUsdg = "usdg"

	MethodUsdgAmounts = "usdgAmounts"

	MethodValidateLiquidation = "validateLiquidation"

	MethodWhitelistedTokenCount = "whitelistedTokenCount"

	MethodWhitelistedTokens = "whitelistedTokens"
)

// IVaultABI is the input ABI used to generate the binding from.
const IVaultABI = "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"allWhitelistedTokens\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allWhitelistedTokensLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_router\",\"type\":\"address\"}],\"name\":\"approvedRouters\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"bufferAmounts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"buyUSDG\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"cumulativeFundingRates\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_collateralToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_indexToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_collateralDelta\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_sizeDelta\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_isLong\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"decreasePosition\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"directPoolDeposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"feeReserves\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fundingInterval\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fundingRateFactor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_indexToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_size\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_averagePrice\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_isLong\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"_lastIncreasedTime\",\"type\":\"uint256\"}],\"name\":\"getDelta\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_usdgDelta\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_feeBasisPoints\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_taxBasisPoints\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_increment\",\"type\":\"bool\"}],\"name\":\"getFeeBasisPoints\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"getMaxPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"getMinPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"getNextFundingRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_collateralToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_indexToken\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_isLong\",\"type\":\"bool\"}],\"name\":\"getPosition\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_usdgAmount\",\"type\":\"uint256\"}],\"name\":\"getRedemptionAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"getTargetUsdgAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"globalShortAveragePrices\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"globalShortSizes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gov\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"guaranteedUsd\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"hasDynamicFees\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"inManagerMode\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"inPrivateLiquidationMode\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_collateralToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_indexToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_sizeDelta\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_isLong\",\"type\":\"bool\"}],\"name\":\"increasePosition\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isInitialized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isLeverageEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"isLiquidator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"isManager\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isSwapEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"lastFundingTimes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_collateralToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_indexToken\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_isLong\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"_feeReceiver\",\"type\":\"address\"}],\"name\":\"liquidatePosition\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"liquidationFeeUsd\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"marginFeeBasisPoints\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxGasPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"maxGlobalShortSizes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxLeverage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"maxUsdgAmounts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"minProfitBasisPoints\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minProfitTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mintBurnFeeBasisPoints\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"poolAmounts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"priceFeed\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"reservedAmounts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"router\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"sellUSDG\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"setBufferAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_errorCode\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_error\",\"type\":\"string\"}],\"name\":\"setError\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_taxBasisPoints\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_stableTaxBasisPoints\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_mintBurnFeeBasisPoints\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_swapFeeBasisPoints\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_stableSwapFeeBasisPoints\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_marginFeeBasisPoints\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_liquidationFeeUsd\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minProfitTime\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_hasDynamicFees\",\"type\":\"bool\"}],\"name\":\"setFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_fundingInterval\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_fundingRateFactor\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_stableFundingRateFactor\",\"type\":\"uint256\"}],\"name\":\"setFundingRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_inManagerMode\",\"type\":\"bool\"}],\"name\":\"setInManagerMode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_inPrivateLiquidationMode\",\"type\":\"bool\"}],\"name\":\"setInPrivateLiquidationMode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_isLeverageEnabled\",\"type\":\"bool\"}],\"name\":\"setIsLeverageEnabled\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_isSwapEnabled\",\"type\":\"bool\"}],\"name\":\"setIsSwapEnabled\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_liquidator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_isActive\",\"type\":\"bool\"}],\"name\":\"setLiquidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_manager\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_isManager\",\"type\":\"bool\"}],\"name\":\"setManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_maxGasPrice\",\"type\":\"uint256\"}],\"name\":\"setMaxGasPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"setMaxGlobalShortSize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_maxLeverage\",\"type\":\"uint256\"}],\"name\":\"setMaxLeverage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_priceFeed\",\"type\":\"address\"}],\"name\":\"setPriceFeed\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenDecimals\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_redemptionBps\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minProfitBps\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxUsdgAmount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_isStable\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"_isShortable\",\"type\":\"bool\"}],\"name\":\"setTokenConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"setUsdgAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIVaultUtils\",\"name\":\"_vaultUtils\",\"type\":\"address\"}],\"name\":\"setVaultUtils\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"shortableTokens\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stableFundingRateFactor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stableSwapFeeBasisPoints\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stableTaxBasisPoints\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"stableTokens\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_tokenOut\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"swap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"swapFeeBasisPoints\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"taxBasisPoints\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"tokenBalances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"tokenDecimals\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenAmount\",\"type\":\"uint256\"}],\"name\":\"tokenToUsdMin\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"tokenWeights\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalTokenWeights\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"usdg\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"usdgAmounts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_collateralToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_indexToken\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_isLong\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"_raise\",\"type\":\"bool\"}],\"name\":\"validateLiquidation\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"whitelistedTokenCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"whitelistedTokens\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"withdrawFees\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IVaultFuncSigs maps the 4-byte function signature to its string representation.
var IVaultFuncSigs = map[string]string{
	"e468baf0": "allWhitelistedTokens(uint256)",
	"0842b076": "allWhitelistedTokensLength()",
	"60922199": "approvedRouters(address,address)",
	"4a993ee9": "bufferAmounts(address)",
	"817bb857": "buyUSDG(address,address)",
	"c65bc7b1": "cumulativeFundingRates(address)",
	"82a08490": "decreasePosition(address,address,address,uint256,uint256,bool,address)",
	"5f7bc119": "directPoolDeposit(address)",
	"1ce9cb8f": "feeReserves(address)",
	"9849e412": "fundingInterval()",
	"c4f718bf": "fundingRateFactor()",
	"5c07eaab": "getDelta(address,uint256,uint256,bool,uint256)",
	"c7e074c3": "getFeeBasisPoints(address,uint256,uint256,uint256,bool)",
	"e124e6d2": "getMaxPrice(address)",
	"81a612d6": "getMinPrice(address)",
	"a93acac2": "getNextFundingRate(address)",
	"4a3f088d": "getPosition(address,address,address,bool)",
	"2c668ec1": "getRedemptionAmount(address,uint256)",
	"3a05dcc1": "getTargetUsdgAmount(address)",
	"62749803": "globalShortAveragePrices(address)",
	"8a78daa8": "globalShortSizes(address)",
	"12d43a51": "gov()",
	"f07456ce": "guaranteedUsd(address)",
	"9f392eb3": "hasDynamicFees()",
	"9060b1ca": "inManagerMode()",
	"181e210e": "inPrivateLiquidationMode()",
	"48d91abf": "increasePosition(address,address,address,uint256,bool)",
	"392e53cd": "isInitialized()",
	"3e72a262": "isLeverageEnabled()",
	"529a356f": "isLiquidator(address)",
	"f3ae2415": "isManager(address)",
	"351a964d": "isSwapEnabled()",
	"d8f897c3": "lastFundingTimes(address)",
	"de2ea948": "liquidatePosition(address,address,address,bool,address)",
	"174d2694": "liquidationFeeUsd()",
	"318bc689": "marginFeeBasisPoints()",
	"3de39c11": "maxGasPrice()",
	"9698d25a": "maxGlobalShortSizes(address)",
	"ae3302c2": "maxLeverage()",
	"ad1e4f8d": "maxUsdgAmounts(address)",
	"88b1fbdf": "minProfitBasisPoints(address)",
	"d9ac4225": "minProfitTime()",
	"4d47b304": "mintBurnFeeBasisPoints()",
	"52f55eed": "poolAmounts(address)",
	"741bef1a": "priceFeed()",
	"c3c7b9e9": "reservedAmounts(address)",
	"f887ea40": "router()",
	"711e6190": "sellUSDG(address,address)",
	"8585f4d2": "setBufferAmount(address,uint256)",
	"28e67be5": "setError(uint256,string)",
	"40eb3802": "setFees(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,bool)",
	"8a27d468": "setFundingRate(uint256,uint256,uint256)",
	"24b0c04d": "setInManagerMode(bool)",
	"f07bbf77": "setInPrivateLiquidationMode(bool)",
	"7c2eb9f7": "setIsLeverageEnabled(bool)",
	"30455ede": "setIsSwapEnabled(bool)",
	"4453a374": "setLiquidator(address,bool)",
	"a5e90eee": "setManager(address,bool)",
	"d2fa635e": "setMaxGasPrice(uint256)",
	"efa10a6e": "setMaxGlobalShortSize(address,uint256)",
	"d3127e63": "setMaxLeverage(uint256)",
	"724e78da": "setPriceFeed(address)",
	"3c5a6e35": "setTokenConfig(address,uint256,uint256,uint256,uint256,bool,bool)",
	"d66b000d": "setUsdgAmount(address,uint256)",
	"71089f4d": "setVaultUtils(address)",
	"db3555fb": "shortableTokens(address)",
	"134ca63b": "stableFundingRateFactor()",
	"df73a267": "stableSwapFeeBasisPoints()",
	"10eb56c2": "stableTaxBasisPoints()",
	"42b60b03": "stableTokens(address)",
	"93316212": "swap(address,address,address)",
	"a22f2392": "swapFeeBasisPoints()",
	"7a210a2b": "taxBasisPoints()",
	"523fba7f": "tokenBalances(address)",
	"8ee573ac": "tokenDecimals(address)",
	"0a48d5a9": "tokenToUsdMin(address,uint256)",
	"ab2f3ad4": "tokenWeights(address)",
	"dc8f5fac": "totalTokenWeights()",
	"f5b91b7b": "usdg()",
	"1aa4ace5": "usdgAmounts(address)",
	"d54d5a9f": "validateLiquidation(address,address,address,bool,bool)",
	"62287a32": "whitelistedTokenCount()",
	"daf9c210": "whitelistedTokens(address)",
	"f2555278": "withdrawFees(address,address)",
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

// AllWhitelistedTokens is a free data retrieval call binding the contract method 0xe468baf0.
//
// Solidity: function allWhitelistedTokens(uint256 ) view returns(address)
func (_IVault *IVaultCaller) AllWhitelistedTokens(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _IVault.contract.Call(opts, &out, "allWhitelistedTokens", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AllWhitelistedTokens is a free data retrieval call binding the contract method 0xe468baf0.
//
// Solidity: function allWhitelistedTokens(uint256 ) view returns(address)
func (_IVault *IVaultSession) AllWhitelistedTokens(arg0 *big.Int) (common.Address, error) {
	return _IVault.Contract.AllWhitelistedTokens(&_IVault.CallOpts, arg0)
}

// AllWhitelistedTokens is a free data retrieval call binding the contract method 0xe468baf0.
//
// Solidity: function allWhitelistedTokens(uint256 ) view returns(address)
func (_IVault *IVaultCallerSession) AllWhitelistedTokens(arg0 *big.Int) (common.Address, error) {
	return _IVault.Contract.AllWhitelistedTokens(&_IVault.CallOpts, arg0)
}

// AllWhitelistedTokensLength is a free data retrieval call binding the contract method 0x0842b076.
//
// Solidity: function allWhitelistedTokensLength() view returns(uint256)
func (_IVault *IVaultCaller) AllWhitelistedTokensLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IVault.contract.Call(opts, &out, "allWhitelistedTokensLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AllWhitelistedTokensLength is a free data retrieval call binding the contract method 0x0842b076.
//
// Solidity: function allWhitelistedTokensLength() view returns(uint256)
func (_IVault *IVaultSession) AllWhitelistedTokensLength() (*big.Int, error) {
	return _IVault.Contract.AllWhitelistedTokensLength(&_IVault.CallOpts)
}

// AllWhitelistedTokensLength is a free data retrieval call binding the contract method 0x0842b076.
//
// Solidity: function allWhitelistedTokensLength() view returns(uint256)
func (_IVault *IVaultCallerSession) AllWhitelistedTokensLength() (*big.Int, error) {
	return _IVault.Contract.AllWhitelistedTokensLength(&_IVault.CallOpts)
}

// ApprovedRouters is a free data retrieval call binding the contract method 0x60922199.
//
// Solidity: function approvedRouters(address _account, address _router) view returns(bool)
func (_IVault *IVaultCaller) ApprovedRouters(opts *bind.CallOpts, _account common.Address, _router common.Address) (bool, error) {
	var out []interface{}
	err := _IVault.contract.Call(opts, &out, "approvedRouters", _account, _router)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ApprovedRouters is a free data retrieval call binding the contract method 0x60922199.
//
// Solidity: function approvedRouters(address _account, address _router) view returns(bool)
func (_IVault *IVaultSession) ApprovedRouters(_account common.Address, _router common.Address) (bool, error) {
	return _IVault.Contract.ApprovedRouters(&_IVault.CallOpts, _account, _router)
}

// ApprovedRouters is a free data retrieval call binding the contract method 0x60922199.
//
// Solidity: function approvedRouters(address _account, address _router) view returns(bool)
func (_IVault *IVaultCallerSession) ApprovedRouters(_account common.Address, _router common.Address) (bool, error) {
	return _IVault.Contract.ApprovedRouters(&_IVault.CallOpts, _account, _router)
}

// BufferAmounts is a free data retrieval call binding the contract method 0x4a993ee9.
//
// Solidity: function bufferAmounts(address _token) view returns(uint256)
func (_IVault *IVaultCaller) BufferAmounts(opts *bind.CallOpts, _token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IVault.contract.Call(opts, &out, "bufferAmounts", _token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BufferAmounts is a free data retrieval call binding the contract method 0x4a993ee9.
//
// Solidity: function bufferAmounts(address _token) view returns(uint256)
func (_IVault *IVaultSession) BufferAmounts(_token common.Address) (*big.Int, error) {
	return _IVault.Contract.BufferAmounts(&_IVault.CallOpts, _token)
}

// BufferAmounts is a free data retrieval call binding the contract method 0x4a993ee9.
//
// Solidity: function bufferAmounts(address _token) view returns(uint256)
func (_IVault *IVaultCallerSession) BufferAmounts(_token common.Address) (*big.Int, error) {
	return _IVault.Contract.BufferAmounts(&_IVault.CallOpts, _token)
}

// CumulativeFundingRates is a free data retrieval call binding the contract method 0xc65bc7b1.
//
// Solidity: function cumulativeFundingRates(address _token) view returns(uint256)
func (_IVault *IVaultCaller) CumulativeFundingRates(opts *bind.CallOpts, _token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IVault.contract.Call(opts, &out, "cumulativeFundingRates", _token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CumulativeFundingRates is a free data retrieval call binding the contract method 0xc65bc7b1.
//
// Solidity: function cumulativeFundingRates(address _token) view returns(uint256)
func (_IVault *IVaultSession) CumulativeFundingRates(_token common.Address) (*big.Int, error) {
	return _IVault.Contract.CumulativeFundingRates(&_IVault.CallOpts, _token)
}

// CumulativeFundingRates is a free data retrieval call binding the contract method 0xc65bc7b1.
//
// Solidity: function cumulativeFundingRates(address _token) view returns(uint256)
func (_IVault *IVaultCallerSession) CumulativeFundingRates(_token common.Address) (*big.Int, error) {
	return _IVault.Contract.CumulativeFundingRates(&_IVault.CallOpts, _token)
}

// FeeReserves is a free data retrieval call binding the contract method 0x1ce9cb8f.
//
// Solidity: function feeReserves(address _token) view returns(uint256)
func (_IVault *IVaultCaller) FeeReserves(opts *bind.CallOpts, _token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IVault.contract.Call(opts, &out, "feeReserves", _token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FeeReserves is a free data retrieval call binding the contract method 0x1ce9cb8f.
//
// Solidity: function feeReserves(address _token) view returns(uint256)
func (_IVault *IVaultSession) FeeReserves(_token common.Address) (*big.Int, error) {
	return _IVault.Contract.FeeReserves(&_IVault.CallOpts, _token)
}

// FeeReserves is a free data retrieval call binding the contract method 0x1ce9cb8f.
//
// Solidity: function feeReserves(address _token) view returns(uint256)
func (_IVault *IVaultCallerSession) FeeReserves(_token common.Address) (*big.Int, error) {
	return _IVault.Contract.FeeReserves(&_IVault.CallOpts, _token)
}

// FundingInterval is a free data retrieval call binding the contract method 0x9849e412.
//
// Solidity: function fundingInterval() view returns(uint256)
func (_IVault *IVaultCaller) FundingInterval(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IVault.contract.Call(opts, &out, "fundingInterval")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FundingInterval is a free data retrieval call binding the contract method 0x9849e412.
//
// Solidity: function fundingInterval() view returns(uint256)
func (_IVault *IVaultSession) FundingInterval() (*big.Int, error) {
	return _IVault.Contract.FundingInterval(&_IVault.CallOpts)
}

// FundingInterval is a free data retrieval call binding the contract method 0x9849e412.
//
// Solidity: function fundingInterval() view returns(uint256)
func (_IVault *IVaultCallerSession) FundingInterval() (*big.Int, error) {
	return _IVault.Contract.FundingInterval(&_IVault.CallOpts)
}

// FundingRateFactor is a free data retrieval call binding the contract method 0xc4f718bf.
//
// Solidity: function fundingRateFactor() view returns(uint256)
func (_IVault *IVaultCaller) FundingRateFactor(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IVault.contract.Call(opts, &out, "fundingRateFactor")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FundingRateFactor is a free data retrieval call binding the contract method 0xc4f718bf.
//
// Solidity: function fundingRateFactor() view returns(uint256)
func (_IVault *IVaultSession) FundingRateFactor() (*big.Int, error) {
	return _IVault.Contract.FundingRateFactor(&_IVault.CallOpts)
}

// FundingRateFactor is a free data retrieval call binding the contract method 0xc4f718bf.
//
// Solidity: function fundingRateFactor() view returns(uint256)
func (_IVault *IVaultCallerSession) FundingRateFactor() (*big.Int, error) {
	return _IVault.Contract.FundingRateFactor(&_IVault.CallOpts)
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

// GetFeeBasisPoints is a free data retrieval call binding the contract method 0xc7e074c3.
//
// Solidity: function getFeeBasisPoints(address _token, uint256 _usdgDelta, uint256 _feeBasisPoints, uint256 _taxBasisPoints, bool _increment) view returns(uint256)
func (_IVault *IVaultCaller) GetFeeBasisPoints(opts *bind.CallOpts, _token common.Address, _usdgDelta *big.Int, _feeBasisPoints *big.Int, _taxBasisPoints *big.Int, _increment bool) (*big.Int, error) {
	var out []interface{}
	err := _IVault.contract.Call(opts, &out, "getFeeBasisPoints", _token, _usdgDelta, _feeBasisPoints, _taxBasisPoints, _increment)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetFeeBasisPoints is a free data retrieval call binding the contract method 0xc7e074c3.
//
// Solidity: function getFeeBasisPoints(address _token, uint256 _usdgDelta, uint256 _feeBasisPoints, uint256 _taxBasisPoints, bool _increment) view returns(uint256)
func (_IVault *IVaultSession) GetFeeBasisPoints(_token common.Address, _usdgDelta *big.Int, _feeBasisPoints *big.Int, _taxBasisPoints *big.Int, _increment bool) (*big.Int, error) {
	return _IVault.Contract.GetFeeBasisPoints(&_IVault.CallOpts, _token, _usdgDelta, _feeBasisPoints, _taxBasisPoints, _increment)
}

// GetFeeBasisPoints is a free data retrieval call binding the contract method 0xc7e074c3.
//
// Solidity: function getFeeBasisPoints(address _token, uint256 _usdgDelta, uint256 _feeBasisPoints, uint256 _taxBasisPoints, bool _increment) view returns(uint256)
func (_IVault *IVaultCallerSession) GetFeeBasisPoints(_token common.Address, _usdgDelta *big.Int, _feeBasisPoints *big.Int, _taxBasisPoints *big.Int, _increment bool) (*big.Int, error) {
	return _IVault.Contract.GetFeeBasisPoints(&_IVault.CallOpts, _token, _usdgDelta, _feeBasisPoints, _taxBasisPoints, _increment)
}

// GetMaxPrice is a free data retrieval call binding the contract method 0xe124e6d2.
//
// Solidity: function getMaxPrice(address _token) view returns(uint256)
func (_IVault *IVaultCaller) GetMaxPrice(opts *bind.CallOpts, _token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IVault.contract.Call(opts, &out, "getMaxPrice", _token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMaxPrice is a free data retrieval call binding the contract method 0xe124e6d2.
//
// Solidity: function getMaxPrice(address _token) view returns(uint256)
func (_IVault *IVaultSession) GetMaxPrice(_token common.Address) (*big.Int, error) {
	return _IVault.Contract.GetMaxPrice(&_IVault.CallOpts, _token)
}

// GetMaxPrice is a free data retrieval call binding the contract method 0xe124e6d2.
//
// Solidity: function getMaxPrice(address _token) view returns(uint256)
func (_IVault *IVaultCallerSession) GetMaxPrice(_token common.Address) (*big.Int, error) {
	return _IVault.Contract.GetMaxPrice(&_IVault.CallOpts, _token)
}

// GetMinPrice is a free data retrieval call binding the contract method 0x81a612d6.
//
// Solidity: function getMinPrice(address _token) view returns(uint256)
func (_IVault *IVaultCaller) GetMinPrice(opts *bind.CallOpts, _token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IVault.contract.Call(opts, &out, "getMinPrice", _token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMinPrice is a free data retrieval call binding the contract method 0x81a612d6.
//
// Solidity: function getMinPrice(address _token) view returns(uint256)
func (_IVault *IVaultSession) GetMinPrice(_token common.Address) (*big.Int, error) {
	return _IVault.Contract.GetMinPrice(&_IVault.CallOpts, _token)
}

// GetMinPrice is a free data retrieval call binding the contract method 0x81a612d6.
//
// Solidity: function getMinPrice(address _token) view returns(uint256)
func (_IVault *IVaultCallerSession) GetMinPrice(_token common.Address) (*big.Int, error) {
	return _IVault.Contract.GetMinPrice(&_IVault.CallOpts, _token)
}

// GetNextFundingRate is a free data retrieval call binding the contract method 0xa93acac2.
//
// Solidity: function getNextFundingRate(address _token) view returns(uint256)
func (_IVault *IVaultCaller) GetNextFundingRate(opts *bind.CallOpts, _token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IVault.contract.Call(opts, &out, "getNextFundingRate", _token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNextFundingRate is a free data retrieval call binding the contract method 0xa93acac2.
//
// Solidity: function getNextFundingRate(address _token) view returns(uint256)
func (_IVault *IVaultSession) GetNextFundingRate(_token common.Address) (*big.Int, error) {
	return _IVault.Contract.GetNextFundingRate(&_IVault.CallOpts, _token)
}

// GetNextFundingRate is a free data retrieval call binding the contract method 0xa93acac2.
//
// Solidity: function getNextFundingRate(address _token) view returns(uint256)
func (_IVault *IVaultCallerSession) GetNextFundingRate(_token common.Address) (*big.Int, error) {
	return _IVault.Contract.GetNextFundingRate(&_IVault.CallOpts, _token)
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

// GetRedemptionAmount is a free data retrieval call binding the contract method 0x2c668ec1.
//
// Solidity: function getRedemptionAmount(address _token, uint256 _usdgAmount) view returns(uint256)
func (_IVault *IVaultCaller) GetRedemptionAmount(opts *bind.CallOpts, _token common.Address, _usdgAmount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IVault.contract.Call(opts, &out, "getRedemptionAmount", _token, _usdgAmount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRedemptionAmount is a free data retrieval call binding the contract method 0x2c668ec1.
//
// Solidity: function getRedemptionAmount(address _token, uint256 _usdgAmount) view returns(uint256)
func (_IVault *IVaultSession) GetRedemptionAmount(_token common.Address, _usdgAmount *big.Int) (*big.Int, error) {
	return _IVault.Contract.GetRedemptionAmount(&_IVault.CallOpts, _token, _usdgAmount)
}

// GetRedemptionAmount is a free data retrieval call binding the contract method 0x2c668ec1.
//
// Solidity: function getRedemptionAmount(address _token, uint256 _usdgAmount) view returns(uint256)
func (_IVault *IVaultCallerSession) GetRedemptionAmount(_token common.Address, _usdgAmount *big.Int) (*big.Int, error) {
	return _IVault.Contract.GetRedemptionAmount(&_IVault.CallOpts, _token, _usdgAmount)
}

// GetTargetUsdgAmount is a free data retrieval call binding the contract method 0x3a05dcc1.
//
// Solidity: function getTargetUsdgAmount(address _token) view returns(uint256)
func (_IVault *IVaultCaller) GetTargetUsdgAmount(opts *bind.CallOpts, _token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IVault.contract.Call(opts, &out, "getTargetUsdgAmount", _token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTargetUsdgAmount is a free data retrieval call binding the contract method 0x3a05dcc1.
//
// Solidity: function getTargetUsdgAmount(address _token) view returns(uint256)
func (_IVault *IVaultSession) GetTargetUsdgAmount(_token common.Address) (*big.Int, error) {
	return _IVault.Contract.GetTargetUsdgAmount(&_IVault.CallOpts, _token)
}

// GetTargetUsdgAmount is a free data retrieval call binding the contract method 0x3a05dcc1.
//
// Solidity: function getTargetUsdgAmount(address _token) view returns(uint256)
func (_IVault *IVaultCallerSession) GetTargetUsdgAmount(_token common.Address) (*big.Int, error) {
	return _IVault.Contract.GetTargetUsdgAmount(&_IVault.CallOpts, _token)
}

// GlobalShortAveragePrices is a free data retrieval call binding the contract method 0x62749803.
//
// Solidity: function globalShortAveragePrices(address _token) view returns(uint256)
func (_IVault *IVaultCaller) GlobalShortAveragePrices(opts *bind.CallOpts, _token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IVault.contract.Call(opts, &out, "globalShortAveragePrices", _token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GlobalShortAveragePrices is a free data retrieval call binding the contract method 0x62749803.
//
// Solidity: function globalShortAveragePrices(address _token) view returns(uint256)
func (_IVault *IVaultSession) GlobalShortAveragePrices(_token common.Address) (*big.Int, error) {
	return _IVault.Contract.GlobalShortAveragePrices(&_IVault.CallOpts, _token)
}

// GlobalShortAveragePrices is a free data retrieval call binding the contract method 0x62749803.
//
// Solidity: function globalShortAveragePrices(address _token) view returns(uint256)
func (_IVault *IVaultCallerSession) GlobalShortAveragePrices(_token common.Address) (*big.Int, error) {
	return _IVault.Contract.GlobalShortAveragePrices(&_IVault.CallOpts, _token)
}

// GlobalShortSizes is a free data retrieval call binding the contract method 0x8a78daa8.
//
// Solidity: function globalShortSizes(address _token) view returns(uint256)
func (_IVault *IVaultCaller) GlobalShortSizes(opts *bind.CallOpts, _token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IVault.contract.Call(opts, &out, "globalShortSizes", _token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GlobalShortSizes is a free data retrieval call binding the contract method 0x8a78daa8.
//
// Solidity: function globalShortSizes(address _token) view returns(uint256)
func (_IVault *IVaultSession) GlobalShortSizes(_token common.Address) (*big.Int, error) {
	return _IVault.Contract.GlobalShortSizes(&_IVault.CallOpts, _token)
}

// GlobalShortSizes is a free data retrieval call binding the contract method 0x8a78daa8.
//
// Solidity: function globalShortSizes(address _token) view returns(uint256)
func (_IVault *IVaultCallerSession) GlobalShortSizes(_token common.Address) (*big.Int, error) {
	return _IVault.Contract.GlobalShortSizes(&_IVault.CallOpts, _token)
}

// Gov is a free data retrieval call binding the contract method 0x12d43a51.
//
// Solidity: function gov() view returns(address)
func (_IVault *IVaultCaller) Gov(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IVault.contract.Call(opts, &out, "gov")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Gov is a free data retrieval call binding the contract method 0x12d43a51.
//
// Solidity: function gov() view returns(address)
func (_IVault *IVaultSession) Gov() (common.Address, error) {
	return _IVault.Contract.Gov(&_IVault.CallOpts)
}

// Gov is a free data retrieval call binding the contract method 0x12d43a51.
//
// Solidity: function gov() view returns(address)
func (_IVault *IVaultCallerSession) Gov() (common.Address, error) {
	return _IVault.Contract.Gov(&_IVault.CallOpts)
}

// GuaranteedUsd is a free data retrieval call binding the contract method 0xf07456ce.
//
// Solidity: function guaranteedUsd(address _token) view returns(uint256)
func (_IVault *IVaultCaller) GuaranteedUsd(opts *bind.CallOpts, _token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IVault.contract.Call(opts, &out, "guaranteedUsd", _token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GuaranteedUsd is a free data retrieval call binding the contract method 0xf07456ce.
//
// Solidity: function guaranteedUsd(address _token) view returns(uint256)
func (_IVault *IVaultSession) GuaranteedUsd(_token common.Address) (*big.Int, error) {
	return _IVault.Contract.GuaranteedUsd(&_IVault.CallOpts, _token)
}

// GuaranteedUsd is a free data retrieval call binding the contract method 0xf07456ce.
//
// Solidity: function guaranteedUsd(address _token) view returns(uint256)
func (_IVault *IVaultCallerSession) GuaranteedUsd(_token common.Address) (*big.Int, error) {
	return _IVault.Contract.GuaranteedUsd(&_IVault.CallOpts, _token)
}

// HasDynamicFees is a free data retrieval call binding the contract method 0x9f392eb3.
//
// Solidity: function hasDynamicFees() view returns(bool)
func (_IVault *IVaultCaller) HasDynamicFees(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _IVault.contract.Call(opts, &out, "hasDynamicFees")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasDynamicFees is a free data retrieval call binding the contract method 0x9f392eb3.
//
// Solidity: function hasDynamicFees() view returns(bool)
func (_IVault *IVaultSession) HasDynamicFees() (bool, error) {
	return _IVault.Contract.HasDynamicFees(&_IVault.CallOpts)
}

// HasDynamicFees is a free data retrieval call binding the contract method 0x9f392eb3.
//
// Solidity: function hasDynamicFees() view returns(bool)
func (_IVault *IVaultCallerSession) HasDynamicFees() (bool, error) {
	return _IVault.Contract.HasDynamicFees(&_IVault.CallOpts)
}

// InManagerMode is a free data retrieval call binding the contract method 0x9060b1ca.
//
// Solidity: function inManagerMode() view returns(bool)
func (_IVault *IVaultCaller) InManagerMode(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _IVault.contract.Call(opts, &out, "inManagerMode")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// InManagerMode is a free data retrieval call binding the contract method 0x9060b1ca.
//
// Solidity: function inManagerMode() view returns(bool)
func (_IVault *IVaultSession) InManagerMode() (bool, error) {
	return _IVault.Contract.InManagerMode(&_IVault.CallOpts)
}

// InManagerMode is a free data retrieval call binding the contract method 0x9060b1ca.
//
// Solidity: function inManagerMode() view returns(bool)
func (_IVault *IVaultCallerSession) InManagerMode() (bool, error) {
	return _IVault.Contract.InManagerMode(&_IVault.CallOpts)
}

// InPrivateLiquidationMode is a free data retrieval call binding the contract method 0x181e210e.
//
// Solidity: function inPrivateLiquidationMode() view returns(bool)
func (_IVault *IVaultCaller) InPrivateLiquidationMode(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _IVault.contract.Call(opts, &out, "inPrivateLiquidationMode")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// InPrivateLiquidationMode is a free data retrieval call binding the contract method 0x181e210e.
//
// Solidity: function inPrivateLiquidationMode() view returns(bool)
func (_IVault *IVaultSession) InPrivateLiquidationMode() (bool, error) {
	return _IVault.Contract.InPrivateLiquidationMode(&_IVault.CallOpts)
}

// InPrivateLiquidationMode is a free data retrieval call binding the contract method 0x181e210e.
//
// Solidity: function inPrivateLiquidationMode() view returns(bool)
func (_IVault *IVaultCallerSession) InPrivateLiquidationMode() (bool, error) {
	return _IVault.Contract.InPrivateLiquidationMode(&_IVault.CallOpts)
}

// IsInitialized is a free data retrieval call binding the contract method 0x392e53cd.
//
// Solidity: function isInitialized() view returns(bool)
func (_IVault *IVaultCaller) IsInitialized(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _IVault.contract.Call(opts, &out, "isInitialized")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsInitialized is a free data retrieval call binding the contract method 0x392e53cd.
//
// Solidity: function isInitialized() view returns(bool)
func (_IVault *IVaultSession) IsInitialized() (bool, error) {
	return _IVault.Contract.IsInitialized(&_IVault.CallOpts)
}

// IsInitialized is a free data retrieval call binding the contract method 0x392e53cd.
//
// Solidity: function isInitialized() view returns(bool)
func (_IVault *IVaultCallerSession) IsInitialized() (bool, error) {
	return _IVault.Contract.IsInitialized(&_IVault.CallOpts)
}

// IsLeverageEnabled is a free data retrieval call binding the contract method 0x3e72a262.
//
// Solidity: function isLeverageEnabled() view returns(bool)
func (_IVault *IVaultCaller) IsLeverageEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _IVault.contract.Call(opts, &out, "isLeverageEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsLeverageEnabled is a free data retrieval call binding the contract method 0x3e72a262.
//
// Solidity: function isLeverageEnabled() view returns(bool)
func (_IVault *IVaultSession) IsLeverageEnabled() (bool, error) {
	return _IVault.Contract.IsLeverageEnabled(&_IVault.CallOpts)
}

// IsLeverageEnabled is a free data retrieval call binding the contract method 0x3e72a262.
//
// Solidity: function isLeverageEnabled() view returns(bool)
func (_IVault *IVaultCallerSession) IsLeverageEnabled() (bool, error) {
	return _IVault.Contract.IsLeverageEnabled(&_IVault.CallOpts)
}

// IsLiquidator is a free data retrieval call binding the contract method 0x529a356f.
//
// Solidity: function isLiquidator(address _account) view returns(bool)
func (_IVault *IVaultCaller) IsLiquidator(opts *bind.CallOpts, _account common.Address) (bool, error) {
	var out []interface{}
	err := _IVault.contract.Call(opts, &out, "isLiquidator", _account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsLiquidator is a free data retrieval call binding the contract method 0x529a356f.
//
// Solidity: function isLiquidator(address _account) view returns(bool)
func (_IVault *IVaultSession) IsLiquidator(_account common.Address) (bool, error) {
	return _IVault.Contract.IsLiquidator(&_IVault.CallOpts, _account)
}

// IsLiquidator is a free data retrieval call binding the contract method 0x529a356f.
//
// Solidity: function isLiquidator(address _account) view returns(bool)
func (_IVault *IVaultCallerSession) IsLiquidator(_account common.Address) (bool, error) {
	return _IVault.Contract.IsLiquidator(&_IVault.CallOpts, _account)
}

// IsManager is a free data retrieval call binding the contract method 0xf3ae2415.
//
// Solidity: function isManager(address _account) view returns(bool)
func (_IVault *IVaultCaller) IsManager(opts *bind.CallOpts, _account common.Address) (bool, error) {
	var out []interface{}
	err := _IVault.contract.Call(opts, &out, "isManager", _account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsManager is a free data retrieval call binding the contract method 0xf3ae2415.
//
// Solidity: function isManager(address _account) view returns(bool)
func (_IVault *IVaultSession) IsManager(_account common.Address) (bool, error) {
	return _IVault.Contract.IsManager(&_IVault.CallOpts, _account)
}

// IsManager is a free data retrieval call binding the contract method 0xf3ae2415.
//
// Solidity: function isManager(address _account) view returns(bool)
func (_IVault *IVaultCallerSession) IsManager(_account common.Address) (bool, error) {
	return _IVault.Contract.IsManager(&_IVault.CallOpts, _account)
}

// IsSwapEnabled is a free data retrieval call binding the contract method 0x351a964d.
//
// Solidity: function isSwapEnabled() view returns(bool)
func (_IVault *IVaultCaller) IsSwapEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _IVault.contract.Call(opts, &out, "isSwapEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSwapEnabled is a free data retrieval call binding the contract method 0x351a964d.
//
// Solidity: function isSwapEnabled() view returns(bool)
func (_IVault *IVaultSession) IsSwapEnabled() (bool, error) {
	return _IVault.Contract.IsSwapEnabled(&_IVault.CallOpts)
}

// IsSwapEnabled is a free data retrieval call binding the contract method 0x351a964d.
//
// Solidity: function isSwapEnabled() view returns(bool)
func (_IVault *IVaultCallerSession) IsSwapEnabled() (bool, error) {
	return _IVault.Contract.IsSwapEnabled(&_IVault.CallOpts)
}

// LastFundingTimes is a free data retrieval call binding the contract method 0xd8f897c3.
//
// Solidity: function lastFundingTimes(address _token) view returns(uint256)
func (_IVault *IVaultCaller) LastFundingTimes(opts *bind.CallOpts, _token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IVault.contract.Call(opts, &out, "lastFundingTimes", _token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastFundingTimes is a free data retrieval call binding the contract method 0xd8f897c3.
//
// Solidity: function lastFundingTimes(address _token) view returns(uint256)
func (_IVault *IVaultSession) LastFundingTimes(_token common.Address) (*big.Int, error) {
	return _IVault.Contract.LastFundingTimes(&_IVault.CallOpts, _token)
}

// LastFundingTimes is a free data retrieval call binding the contract method 0xd8f897c3.
//
// Solidity: function lastFundingTimes(address _token) view returns(uint256)
func (_IVault *IVaultCallerSession) LastFundingTimes(_token common.Address) (*big.Int, error) {
	return _IVault.Contract.LastFundingTimes(&_IVault.CallOpts, _token)
}

// LiquidationFeeUsd is a free data retrieval call binding the contract method 0x174d2694.
//
// Solidity: function liquidationFeeUsd() view returns(uint256)
func (_IVault *IVaultCaller) LiquidationFeeUsd(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IVault.contract.Call(opts, &out, "liquidationFeeUsd")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LiquidationFeeUsd is a free data retrieval call binding the contract method 0x174d2694.
//
// Solidity: function liquidationFeeUsd() view returns(uint256)
func (_IVault *IVaultSession) LiquidationFeeUsd() (*big.Int, error) {
	return _IVault.Contract.LiquidationFeeUsd(&_IVault.CallOpts)
}

// LiquidationFeeUsd is a free data retrieval call binding the contract method 0x174d2694.
//
// Solidity: function liquidationFeeUsd() view returns(uint256)
func (_IVault *IVaultCallerSession) LiquidationFeeUsd() (*big.Int, error) {
	return _IVault.Contract.LiquidationFeeUsd(&_IVault.CallOpts)
}

// MarginFeeBasisPoints is a free data retrieval call binding the contract method 0x318bc689.
//
// Solidity: function marginFeeBasisPoints() view returns(uint256)
func (_IVault *IVaultCaller) MarginFeeBasisPoints(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IVault.contract.Call(opts, &out, "marginFeeBasisPoints")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MarginFeeBasisPoints is a free data retrieval call binding the contract method 0x318bc689.
//
// Solidity: function marginFeeBasisPoints() view returns(uint256)
func (_IVault *IVaultSession) MarginFeeBasisPoints() (*big.Int, error) {
	return _IVault.Contract.MarginFeeBasisPoints(&_IVault.CallOpts)
}

// MarginFeeBasisPoints is a free data retrieval call binding the contract method 0x318bc689.
//
// Solidity: function marginFeeBasisPoints() view returns(uint256)
func (_IVault *IVaultCallerSession) MarginFeeBasisPoints() (*big.Int, error) {
	return _IVault.Contract.MarginFeeBasisPoints(&_IVault.CallOpts)
}

// MaxGasPrice is a free data retrieval call binding the contract method 0x3de39c11.
//
// Solidity: function maxGasPrice() view returns(uint256)
func (_IVault *IVaultCaller) MaxGasPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IVault.contract.Call(opts, &out, "maxGasPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxGasPrice is a free data retrieval call binding the contract method 0x3de39c11.
//
// Solidity: function maxGasPrice() view returns(uint256)
func (_IVault *IVaultSession) MaxGasPrice() (*big.Int, error) {
	return _IVault.Contract.MaxGasPrice(&_IVault.CallOpts)
}

// MaxGasPrice is a free data retrieval call binding the contract method 0x3de39c11.
//
// Solidity: function maxGasPrice() view returns(uint256)
func (_IVault *IVaultCallerSession) MaxGasPrice() (*big.Int, error) {
	return _IVault.Contract.MaxGasPrice(&_IVault.CallOpts)
}

// MaxGlobalShortSizes is a free data retrieval call binding the contract method 0x9698d25a.
//
// Solidity: function maxGlobalShortSizes(address _token) view returns(uint256)
func (_IVault *IVaultCaller) MaxGlobalShortSizes(opts *bind.CallOpts, _token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IVault.contract.Call(opts, &out, "maxGlobalShortSizes", _token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxGlobalShortSizes is a free data retrieval call binding the contract method 0x9698d25a.
//
// Solidity: function maxGlobalShortSizes(address _token) view returns(uint256)
func (_IVault *IVaultSession) MaxGlobalShortSizes(_token common.Address) (*big.Int, error) {
	return _IVault.Contract.MaxGlobalShortSizes(&_IVault.CallOpts, _token)
}

// MaxGlobalShortSizes is a free data retrieval call binding the contract method 0x9698d25a.
//
// Solidity: function maxGlobalShortSizes(address _token) view returns(uint256)
func (_IVault *IVaultCallerSession) MaxGlobalShortSizes(_token common.Address) (*big.Int, error) {
	return _IVault.Contract.MaxGlobalShortSizes(&_IVault.CallOpts, _token)
}

// MaxLeverage is a free data retrieval call binding the contract method 0xae3302c2.
//
// Solidity: function maxLeverage() view returns(uint256)
func (_IVault *IVaultCaller) MaxLeverage(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IVault.contract.Call(opts, &out, "maxLeverage")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxLeverage is a free data retrieval call binding the contract method 0xae3302c2.
//
// Solidity: function maxLeverage() view returns(uint256)
func (_IVault *IVaultSession) MaxLeverage() (*big.Int, error) {
	return _IVault.Contract.MaxLeverage(&_IVault.CallOpts)
}

// MaxLeverage is a free data retrieval call binding the contract method 0xae3302c2.
//
// Solidity: function maxLeverage() view returns(uint256)
func (_IVault *IVaultCallerSession) MaxLeverage() (*big.Int, error) {
	return _IVault.Contract.MaxLeverage(&_IVault.CallOpts)
}

// MaxUsdgAmounts is a free data retrieval call binding the contract method 0xad1e4f8d.
//
// Solidity: function maxUsdgAmounts(address _token) view returns(uint256)
func (_IVault *IVaultCaller) MaxUsdgAmounts(opts *bind.CallOpts, _token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IVault.contract.Call(opts, &out, "maxUsdgAmounts", _token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxUsdgAmounts is a free data retrieval call binding the contract method 0xad1e4f8d.
//
// Solidity: function maxUsdgAmounts(address _token) view returns(uint256)
func (_IVault *IVaultSession) MaxUsdgAmounts(_token common.Address) (*big.Int, error) {
	return _IVault.Contract.MaxUsdgAmounts(&_IVault.CallOpts, _token)
}

// MaxUsdgAmounts is a free data retrieval call binding the contract method 0xad1e4f8d.
//
// Solidity: function maxUsdgAmounts(address _token) view returns(uint256)
func (_IVault *IVaultCallerSession) MaxUsdgAmounts(_token common.Address) (*big.Int, error) {
	return _IVault.Contract.MaxUsdgAmounts(&_IVault.CallOpts, _token)
}

// MinProfitBasisPoints is a free data retrieval call binding the contract method 0x88b1fbdf.
//
// Solidity: function minProfitBasisPoints(address _token) view returns(uint256)
func (_IVault *IVaultCaller) MinProfitBasisPoints(opts *bind.CallOpts, _token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IVault.contract.Call(opts, &out, "minProfitBasisPoints", _token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinProfitBasisPoints is a free data retrieval call binding the contract method 0x88b1fbdf.
//
// Solidity: function minProfitBasisPoints(address _token) view returns(uint256)
func (_IVault *IVaultSession) MinProfitBasisPoints(_token common.Address) (*big.Int, error) {
	return _IVault.Contract.MinProfitBasisPoints(&_IVault.CallOpts, _token)
}

// MinProfitBasisPoints is a free data retrieval call binding the contract method 0x88b1fbdf.
//
// Solidity: function minProfitBasisPoints(address _token) view returns(uint256)
func (_IVault *IVaultCallerSession) MinProfitBasisPoints(_token common.Address) (*big.Int, error) {
	return _IVault.Contract.MinProfitBasisPoints(&_IVault.CallOpts, _token)
}

// MinProfitTime is a free data retrieval call binding the contract method 0xd9ac4225.
//
// Solidity: function minProfitTime() view returns(uint256)
func (_IVault *IVaultCaller) MinProfitTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IVault.contract.Call(opts, &out, "minProfitTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinProfitTime is a free data retrieval call binding the contract method 0xd9ac4225.
//
// Solidity: function minProfitTime() view returns(uint256)
func (_IVault *IVaultSession) MinProfitTime() (*big.Int, error) {
	return _IVault.Contract.MinProfitTime(&_IVault.CallOpts)
}

// MinProfitTime is a free data retrieval call binding the contract method 0xd9ac4225.
//
// Solidity: function minProfitTime() view returns(uint256)
func (_IVault *IVaultCallerSession) MinProfitTime() (*big.Int, error) {
	return _IVault.Contract.MinProfitTime(&_IVault.CallOpts)
}

// MintBurnFeeBasisPoints is a free data retrieval call binding the contract method 0x4d47b304.
//
// Solidity: function mintBurnFeeBasisPoints() view returns(uint256)
func (_IVault *IVaultCaller) MintBurnFeeBasisPoints(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IVault.contract.Call(opts, &out, "mintBurnFeeBasisPoints")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MintBurnFeeBasisPoints is a free data retrieval call binding the contract method 0x4d47b304.
//
// Solidity: function mintBurnFeeBasisPoints() view returns(uint256)
func (_IVault *IVaultSession) MintBurnFeeBasisPoints() (*big.Int, error) {
	return _IVault.Contract.MintBurnFeeBasisPoints(&_IVault.CallOpts)
}

// MintBurnFeeBasisPoints is a free data retrieval call binding the contract method 0x4d47b304.
//
// Solidity: function mintBurnFeeBasisPoints() view returns(uint256)
func (_IVault *IVaultCallerSession) MintBurnFeeBasisPoints() (*big.Int, error) {
	return _IVault.Contract.MintBurnFeeBasisPoints(&_IVault.CallOpts)
}

// PoolAmounts is a free data retrieval call binding the contract method 0x52f55eed.
//
// Solidity: function poolAmounts(address _token) view returns(uint256)
func (_IVault *IVaultCaller) PoolAmounts(opts *bind.CallOpts, _token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IVault.contract.Call(opts, &out, "poolAmounts", _token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PoolAmounts is a free data retrieval call binding the contract method 0x52f55eed.
//
// Solidity: function poolAmounts(address _token) view returns(uint256)
func (_IVault *IVaultSession) PoolAmounts(_token common.Address) (*big.Int, error) {
	return _IVault.Contract.PoolAmounts(&_IVault.CallOpts, _token)
}

// PoolAmounts is a free data retrieval call binding the contract method 0x52f55eed.
//
// Solidity: function poolAmounts(address _token) view returns(uint256)
func (_IVault *IVaultCallerSession) PoolAmounts(_token common.Address) (*big.Int, error) {
	return _IVault.Contract.PoolAmounts(&_IVault.CallOpts, _token)
}

// PriceFeed is a free data retrieval call binding the contract method 0x741bef1a.
//
// Solidity: function priceFeed() view returns(address)
func (_IVault *IVaultCaller) PriceFeed(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IVault.contract.Call(opts, &out, "priceFeed")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PriceFeed is a free data retrieval call binding the contract method 0x741bef1a.
//
// Solidity: function priceFeed() view returns(address)
func (_IVault *IVaultSession) PriceFeed() (common.Address, error) {
	return _IVault.Contract.PriceFeed(&_IVault.CallOpts)
}

// PriceFeed is a free data retrieval call binding the contract method 0x741bef1a.
//
// Solidity: function priceFeed() view returns(address)
func (_IVault *IVaultCallerSession) PriceFeed() (common.Address, error) {
	return _IVault.Contract.PriceFeed(&_IVault.CallOpts)
}

// ReservedAmounts is a free data retrieval call binding the contract method 0xc3c7b9e9.
//
// Solidity: function reservedAmounts(address _token) view returns(uint256)
func (_IVault *IVaultCaller) ReservedAmounts(opts *bind.CallOpts, _token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IVault.contract.Call(opts, &out, "reservedAmounts", _token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ReservedAmounts is a free data retrieval call binding the contract method 0xc3c7b9e9.
//
// Solidity: function reservedAmounts(address _token) view returns(uint256)
func (_IVault *IVaultSession) ReservedAmounts(_token common.Address) (*big.Int, error) {
	return _IVault.Contract.ReservedAmounts(&_IVault.CallOpts, _token)
}

// ReservedAmounts is a free data retrieval call binding the contract method 0xc3c7b9e9.
//
// Solidity: function reservedAmounts(address _token) view returns(uint256)
func (_IVault *IVaultCallerSession) ReservedAmounts(_token common.Address) (*big.Int, error) {
	return _IVault.Contract.ReservedAmounts(&_IVault.CallOpts, _token)
}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_IVault *IVaultCaller) Router(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IVault.contract.Call(opts, &out, "router")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_IVault *IVaultSession) Router() (common.Address, error) {
	return _IVault.Contract.Router(&_IVault.CallOpts)
}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_IVault *IVaultCallerSession) Router() (common.Address, error) {
	return _IVault.Contract.Router(&_IVault.CallOpts)
}

// ShortableTokens is a free data retrieval call binding the contract method 0xdb3555fb.
//
// Solidity: function shortableTokens(address _token) view returns(bool)
func (_IVault *IVaultCaller) ShortableTokens(opts *bind.CallOpts, _token common.Address) (bool, error) {
	var out []interface{}
	err := _IVault.contract.Call(opts, &out, "shortableTokens", _token)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ShortableTokens is a free data retrieval call binding the contract method 0xdb3555fb.
//
// Solidity: function shortableTokens(address _token) view returns(bool)
func (_IVault *IVaultSession) ShortableTokens(_token common.Address) (bool, error) {
	return _IVault.Contract.ShortableTokens(&_IVault.CallOpts, _token)
}

// ShortableTokens is a free data retrieval call binding the contract method 0xdb3555fb.
//
// Solidity: function shortableTokens(address _token) view returns(bool)
func (_IVault *IVaultCallerSession) ShortableTokens(_token common.Address) (bool, error) {
	return _IVault.Contract.ShortableTokens(&_IVault.CallOpts, _token)
}

// StableFundingRateFactor is a free data retrieval call binding the contract method 0x134ca63b.
//
// Solidity: function stableFundingRateFactor() view returns(uint256)
func (_IVault *IVaultCaller) StableFundingRateFactor(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IVault.contract.Call(opts, &out, "stableFundingRateFactor")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StableFundingRateFactor is a free data retrieval call binding the contract method 0x134ca63b.
//
// Solidity: function stableFundingRateFactor() view returns(uint256)
func (_IVault *IVaultSession) StableFundingRateFactor() (*big.Int, error) {
	return _IVault.Contract.StableFundingRateFactor(&_IVault.CallOpts)
}

// StableFundingRateFactor is a free data retrieval call binding the contract method 0x134ca63b.
//
// Solidity: function stableFundingRateFactor() view returns(uint256)
func (_IVault *IVaultCallerSession) StableFundingRateFactor() (*big.Int, error) {
	return _IVault.Contract.StableFundingRateFactor(&_IVault.CallOpts)
}

// StableSwapFeeBasisPoints is a free data retrieval call binding the contract method 0xdf73a267.
//
// Solidity: function stableSwapFeeBasisPoints() view returns(uint256)
func (_IVault *IVaultCaller) StableSwapFeeBasisPoints(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IVault.contract.Call(opts, &out, "stableSwapFeeBasisPoints")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StableSwapFeeBasisPoints is a free data retrieval call binding the contract method 0xdf73a267.
//
// Solidity: function stableSwapFeeBasisPoints() view returns(uint256)
func (_IVault *IVaultSession) StableSwapFeeBasisPoints() (*big.Int, error) {
	return _IVault.Contract.StableSwapFeeBasisPoints(&_IVault.CallOpts)
}

// StableSwapFeeBasisPoints is a free data retrieval call binding the contract method 0xdf73a267.
//
// Solidity: function stableSwapFeeBasisPoints() view returns(uint256)
func (_IVault *IVaultCallerSession) StableSwapFeeBasisPoints() (*big.Int, error) {
	return _IVault.Contract.StableSwapFeeBasisPoints(&_IVault.CallOpts)
}

// StableTaxBasisPoints is a free data retrieval call binding the contract method 0x10eb56c2.
//
// Solidity: function stableTaxBasisPoints() view returns(uint256)
func (_IVault *IVaultCaller) StableTaxBasisPoints(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IVault.contract.Call(opts, &out, "stableTaxBasisPoints")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StableTaxBasisPoints is a free data retrieval call binding the contract method 0x10eb56c2.
//
// Solidity: function stableTaxBasisPoints() view returns(uint256)
func (_IVault *IVaultSession) StableTaxBasisPoints() (*big.Int, error) {
	return _IVault.Contract.StableTaxBasisPoints(&_IVault.CallOpts)
}

// StableTaxBasisPoints is a free data retrieval call binding the contract method 0x10eb56c2.
//
// Solidity: function stableTaxBasisPoints() view returns(uint256)
func (_IVault *IVaultCallerSession) StableTaxBasisPoints() (*big.Int, error) {
	return _IVault.Contract.StableTaxBasisPoints(&_IVault.CallOpts)
}

// StableTokens is a free data retrieval call binding the contract method 0x42b60b03.
//
// Solidity: function stableTokens(address _token) view returns(bool)
func (_IVault *IVaultCaller) StableTokens(opts *bind.CallOpts, _token common.Address) (bool, error) {
	var out []interface{}
	err := _IVault.contract.Call(opts, &out, "stableTokens", _token)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// StableTokens is a free data retrieval call binding the contract method 0x42b60b03.
//
// Solidity: function stableTokens(address _token) view returns(bool)
func (_IVault *IVaultSession) StableTokens(_token common.Address) (bool, error) {
	return _IVault.Contract.StableTokens(&_IVault.CallOpts, _token)
}

// StableTokens is a free data retrieval call binding the contract method 0x42b60b03.
//
// Solidity: function stableTokens(address _token) view returns(bool)
func (_IVault *IVaultCallerSession) StableTokens(_token common.Address) (bool, error) {
	return _IVault.Contract.StableTokens(&_IVault.CallOpts, _token)
}

// SwapFeeBasisPoints is a free data retrieval call binding the contract method 0xa22f2392.
//
// Solidity: function swapFeeBasisPoints() view returns(uint256)
func (_IVault *IVaultCaller) SwapFeeBasisPoints(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IVault.contract.Call(opts, &out, "swapFeeBasisPoints")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SwapFeeBasisPoints is a free data retrieval call binding the contract method 0xa22f2392.
//
// Solidity: function swapFeeBasisPoints() view returns(uint256)
func (_IVault *IVaultSession) SwapFeeBasisPoints() (*big.Int, error) {
	return _IVault.Contract.SwapFeeBasisPoints(&_IVault.CallOpts)
}

// SwapFeeBasisPoints is a free data retrieval call binding the contract method 0xa22f2392.
//
// Solidity: function swapFeeBasisPoints() view returns(uint256)
func (_IVault *IVaultCallerSession) SwapFeeBasisPoints() (*big.Int, error) {
	return _IVault.Contract.SwapFeeBasisPoints(&_IVault.CallOpts)
}

// TaxBasisPoints is a free data retrieval call binding the contract method 0x7a210a2b.
//
// Solidity: function taxBasisPoints() view returns(uint256)
func (_IVault *IVaultCaller) TaxBasisPoints(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IVault.contract.Call(opts, &out, "taxBasisPoints")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TaxBasisPoints is a free data retrieval call binding the contract method 0x7a210a2b.
//
// Solidity: function taxBasisPoints() view returns(uint256)
func (_IVault *IVaultSession) TaxBasisPoints() (*big.Int, error) {
	return _IVault.Contract.TaxBasisPoints(&_IVault.CallOpts)
}

// TaxBasisPoints is a free data retrieval call binding the contract method 0x7a210a2b.
//
// Solidity: function taxBasisPoints() view returns(uint256)
func (_IVault *IVaultCallerSession) TaxBasisPoints() (*big.Int, error) {
	return _IVault.Contract.TaxBasisPoints(&_IVault.CallOpts)
}

// TokenBalances is a free data retrieval call binding the contract method 0x523fba7f.
//
// Solidity: function tokenBalances(address _token) view returns(uint256)
func (_IVault *IVaultCaller) TokenBalances(opts *bind.CallOpts, _token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IVault.contract.Call(opts, &out, "tokenBalances", _token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenBalances is a free data retrieval call binding the contract method 0x523fba7f.
//
// Solidity: function tokenBalances(address _token) view returns(uint256)
func (_IVault *IVaultSession) TokenBalances(_token common.Address) (*big.Int, error) {
	return _IVault.Contract.TokenBalances(&_IVault.CallOpts, _token)
}

// TokenBalances is a free data retrieval call binding the contract method 0x523fba7f.
//
// Solidity: function tokenBalances(address _token) view returns(uint256)
func (_IVault *IVaultCallerSession) TokenBalances(_token common.Address) (*big.Int, error) {
	return _IVault.Contract.TokenBalances(&_IVault.CallOpts, _token)
}

// TokenDecimals is a free data retrieval call binding the contract method 0x8ee573ac.
//
// Solidity: function tokenDecimals(address _token) view returns(uint256)
func (_IVault *IVaultCaller) TokenDecimals(opts *bind.CallOpts, _token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IVault.contract.Call(opts, &out, "tokenDecimals", _token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenDecimals is a free data retrieval call binding the contract method 0x8ee573ac.
//
// Solidity: function tokenDecimals(address _token) view returns(uint256)
func (_IVault *IVaultSession) TokenDecimals(_token common.Address) (*big.Int, error) {
	return _IVault.Contract.TokenDecimals(&_IVault.CallOpts, _token)
}

// TokenDecimals is a free data retrieval call binding the contract method 0x8ee573ac.
//
// Solidity: function tokenDecimals(address _token) view returns(uint256)
func (_IVault *IVaultCallerSession) TokenDecimals(_token common.Address) (*big.Int, error) {
	return _IVault.Contract.TokenDecimals(&_IVault.CallOpts, _token)
}

// TokenToUsdMin is a free data retrieval call binding the contract method 0x0a48d5a9.
//
// Solidity: function tokenToUsdMin(address _token, uint256 _tokenAmount) view returns(uint256)
func (_IVault *IVaultCaller) TokenToUsdMin(opts *bind.CallOpts, _token common.Address, _tokenAmount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IVault.contract.Call(opts, &out, "tokenToUsdMin", _token, _tokenAmount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenToUsdMin is a free data retrieval call binding the contract method 0x0a48d5a9.
//
// Solidity: function tokenToUsdMin(address _token, uint256 _tokenAmount) view returns(uint256)
func (_IVault *IVaultSession) TokenToUsdMin(_token common.Address, _tokenAmount *big.Int) (*big.Int, error) {
	return _IVault.Contract.TokenToUsdMin(&_IVault.CallOpts, _token, _tokenAmount)
}

// TokenToUsdMin is a free data retrieval call binding the contract method 0x0a48d5a9.
//
// Solidity: function tokenToUsdMin(address _token, uint256 _tokenAmount) view returns(uint256)
func (_IVault *IVaultCallerSession) TokenToUsdMin(_token common.Address, _tokenAmount *big.Int) (*big.Int, error) {
	return _IVault.Contract.TokenToUsdMin(&_IVault.CallOpts, _token, _tokenAmount)
}

// TokenWeights is a free data retrieval call binding the contract method 0xab2f3ad4.
//
// Solidity: function tokenWeights(address _token) view returns(uint256)
func (_IVault *IVaultCaller) TokenWeights(opts *bind.CallOpts, _token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IVault.contract.Call(opts, &out, "tokenWeights", _token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenWeights is a free data retrieval call binding the contract method 0xab2f3ad4.
//
// Solidity: function tokenWeights(address _token) view returns(uint256)
func (_IVault *IVaultSession) TokenWeights(_token common.Address) (*big.Int, error) {
	return _IVault.Contract.TokenWeights(&_IVault.CallOpts, _token)
}

// TokenWeights is a free data retrieval call binding the contract method 0xab2f3ad4.
//
// Solidity: function tokenWeights(address _token) view returns(uint256)
func (_IVault *IVaultCallerSession) TokenWeights(_token common.Address) (*big.Int, error) {
	return _IVault.Contract.TokenWeights(&_IVault.CallOpts, _token)
}

// TotalTokenWeights is a free data retrieval call binding the contract method 0xdc8f5fac.
//
// Solidity: function totalTokenWeights() view returns(uint256)
func (_IVault *IVaultCaller) TotalTokenWeights(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IVault.contract.Call(opts, &out, "totalTokenWeights")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalTokenWeights is a free data retrieval call binding the contract method 0xdc8f5fac.
//
// Solidity: function totalTokenWeights() view returns(uint256)
func (_IVault *IVaultSession) TotalTokenWeights() (*big.Int, error) {
	return _IVault.Contract.TotalTokenWeights(&_IVault.CallOpts)
}

// TotalTokenWeights is a free data retrieval call binding the contract method 0xdc8f5fac.
//
// Solidity: function totalTokenWeights() view returns(uint256)
func (_IVault *IVaultCallerSession) TotalTokenWeights() (*big.Int, error) {
	return _IVault.Contract.TotalTokenWeights(&_IVault.CallOpts)
}

// Usdg is a free data retrieval call binding the contract method 0xf5b91b7b.
//
// Solidity: function usdg() view returns(address)
func (_IVault *IVaultCaller) Usdg(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IVault.contract.Call(opts, &out, "usdg")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Usdg is a free data retrieval call binding the contract method 0xf5b91b7b.
//
// Solidity: function usdg() view returns(address)
func (_IVault *IVaultSession) Usdg() (common.Address, error) {
	return _IVault.Contract.Usdg(&_IVault.CallOpts)
}

// Usdg is a free data retrieval call binding the contract method 0xf5b91b7b.
//
// Solidity: function usdg() view returns(address)
func (_IVault *IVaultCallerSession) Usdg() (common.Address, error) {
	return _IVault.Contract.Usdg(&_IVault.CallOpts)
}

// UsdgAmounts is a free data retrieval call binding the contract method 0x1aa4ace5.
//
// Solidity: function usdgAmounts(address _token) view returns(uint256)
func (_IVault *IVaultCaller) UsdgAmounts(opts *bind.CallOpts, _token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IVault.contract.Call(opts, &out, "usdgAmounts", _token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UsdgAmounts is a free data retrieval call binding the contract method 0x1aa4ace5.
//
// Solidity: function usdgAmounts(address _token) view returns(uint256)
func (_IVault *IVaultSession) UsdgAmounts(_token common.Address) (*big.Int, error) {
	return _IVault.Contract.UsdgAmounts(&_IVault.CallOpts, _token)
}

// UsdgAmounts is a free data retrieval call binding the contract method 0x1aa4ace5.
//
// Solidity: function usdgAmounts(address _token) view returns(uint256)
func (_IVault *IVaultCallerSession) UsdgAmounts(_token common.Address) (*big.Int, error) {
	return _IVault.Contract.UsdgAmounts(&_IVault.CallOpts, _token)
}

// ValidateLiquidation is a free data retrieval call binding the contract method 0xd54d5a9f.
//
// Solidity: function validateLiquidation(address _account, address _collateralToken, address _indexToken, bool _isLong, bool _raise) view returns(uint256, uint256)
func (_IVault *IVaultCaller) ValidateLiquidation(opts *bind.CallOpts, _account common.Address, _collateralToken common.Address, _indexToken common.Address, _isLong bool, _raise bool) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _IVault.contract.Call(opts, &out, "validateLiquidation", _account, _collateralToken, _indexToken, _isLong, _raise)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// ValidateLiquidation is a free data retrieval call binding the contract method 0xd54d5a9f.
//
// Solidity: function validateLiquidation(address _account, address _collateralToken, address _indexToken, bool _isLong, bool _raise) view returns(uint256, uint256)
func (_IVault *IVaultSession) ValidateLiquidation(_account common.Address, _collateralToken common.Address, _indexToken common.Address, _isLong bool, _raise bool) (*big.Int, *big.Int, error) {
	return _IVault.Contract.ValidateLiquidation(&_IVault.CallOpts, _account, _collateralToken, _indexToken, _isLong, _raise)
}

// ValidateLiquidation is a free data retrieval call binding the contract method 0xd54d5a9f.
//
// Solidity: function validateLiquidation(address _account, address _collateralToken, address _indexToken, bool _isLong, bool _raise) view returns(uint256, uint256)
func (_IVault *IVaultCallerSession) ValidateLiquidation(_account common.Address, _collateralToken common.Address, _indexToken common.Address, _isLong bool, _raise bool) (*big.Int, *big.Int, error) {
	return _IVault.Contract.ValidateLiquidation(&_IVault.CallOpts, _account, _collateralToken, _indexToken, _isLong, _raise)
}

// WhitelistedTokenCount is a free data retrieval call binding the contract method 0x62287a32.
//
// Solidity: function whitelistedTokenCount() view returns(uint256)
func (_IVault *IVaultCaller) WhitelistedTokenCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IVault.contract.Call(opts, &out, "whitelistedTokenCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WhitelistedTokenCount is a free data retrieval call binding the contract method 0x62287a32.
//
// Solidity: function whitelistedTokenCount() view returns(uint256)
func (_IVault *IVaultSession) WhitelistedTokenCount() (*big.Int, error) {
	return _IVault.Contract.WhitelistedTokenCount(&_IVault.CallOpts)
}

// WhitelistedTokenCount is a free data retrieval call binding the contract method 0x62287a32.
//
// Solidity: function whitelistedTokenCount() view returns(uint256)
func (_IVault *IVaultCallerSession) WhitelistedTokenCount() (*big.Int, error) {
	return _IVault.Contract.WhitelistedTokenCount(&_IVault.CallOpts)
}

// WhitelistedTokens is a free data retrieval call binding the contract method 0xdaf9c210.
//
// Solidity: function whitelistedTokens(address _token) view returns(bool)
func (_IVault *IVaultCaller) WhitelistedTokens(opts *bind.CallOpts, _token common.Address) (bool, error) {
	var out []interface{}
	err := _IVault.contract.Call(opts, &out, "whitelistedTokens", _token)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// WhitelistedTokens is a free data retrieval call binding the contract method 0xdaf9c210.
//
// Solidity: function whitelistedTokens(address _token) view returns(bool)
func (_IVault *IVaultSession) WhitelistedTokens(_token common.Address) (bool, error) {
	return _IVault.Contract.WhitelistedTokens(&_IVault.CallOpts, _token)
}

// WhitelistedTokens is a free data retrieval call binding the contract method 0xdaf9c210.
//
// Solidity: function whitelistedTokens(address _token) view returns(bool)
func (_IVault *IVaultCallerSession) WhitelistedTokens(_token common.Address) (bool, error) {
	return _IVault.Contract.WhitelistedTokens(&_IVault.CallOpts, _token)
}

// BuyUSDG is a paid mutator transaction binding the contract method 0x817bb857.
//
// Solidity: function buyUSDG(address _token, address _receiver) returns(uint256)
func (_IVault *IVaultTransactor) BuyUSDG(opts *bind.TransactOpts, _token common.Address, _receiver common.Address) (*types.Transaction, error) {
	return _IVault.contract.Transact(opts, "buyUSDG", _token, _receiver)
}

// BuyUSDG is a paid mutator transaction binding the contract method 0x817bb857.
//
// Solidity: function buyUSDG(address _token, address _receiver) returns(uint256)
func (_IVault *IVaultSession) BuyUSDG(_token common.Address, _receiver common.Address) (*types.Transaction, error) {
	return _IVault.Contract.BuyUSDG(&_IVault.TransactOpts, _token, _receiver)
}

// BuyUSDG is a paid mutator transaction binding the contract method 0x817bb857.
//
// Solidity: function buyUSDG(address _token, address _receiver) returns(uint256)
func (_IVault *IVaultTransactorSession) BuyUSDG(_token common.Address, _receiver common.Address) (*types.Transaction, error) {
	return _IVault.Contract.BuyUSDG(&_IVault.TransactOpts, _token, _receiver)
}

// DecreasePosition is a paid mutator transaction binding the contract method 0x82a08490.
//
// Solidity: function decreasePosition(address _account, address _collateralToken, address _indexToken, uint256 _collateralDelta, uint256 _sizeDelta, bool _isLong, address _receiver) returns(uint256)
func (_IVault *IVaultTransactor) DecreasePosition(opts *bind.TransactOpts, _account common.Address, _collateralToken common.Address, _indexToken common.Address, _collateralDelta *big.Int, _sizeDelta *big.Int, _isLong bool, _receiver common.Address) (*types.Transaction, error) {
	return _IVault.contract.Transact(opts, "decreasePosition", _account, _collateralToken, _indexToken, _collateralDelta, _sizeDelta, _isLong, _receiver)
}

// DecreasePosition is a paid mutator transaction binding the contract method 0x82a08490.
//
// Solidity: function decreasePosition(address _account, address _collateralToken, address _indexToken, uint256 _collateralDelta, uint256 _sizeDelta, bool _isLong, address _receiver) returns(uint256)
func (_IVault *IVaultSession) DecreasePosition(_account common.Address, _collateralToken common.Address, _indexToken common.Address, _collateralDelta *big.Int, _sizeDelta *big.Int, _isLong bool, _receiver common.Address) (*types.Transaction, error) {
	return _IVault.Contract.DecreasePosition(&_IVault.TransactOpts, _account, _collateralToken, _indexToken, _collateralDelta, _sizeDelta, _isLong, _receiver)
}

// DecreasePosition is a paid mutator transaction binding the contract method 0x82a08490.
//
// Solidity: function decreasePosition(address _account, address _collateralToken, address _indexToken, uint256 _collateralDelta, uint256 _sizeDelta, bool _isLong, address _receiver) returns(uint256)
func (_IVault *IVaultTransactorSession) DecreasePosition(_account common.Address, _collateralToken common.Address, _indexToken common.Address, _collateralDelta *big.Int, _sizeDelta *big.Int, _isLong bool, _receiver common.Address) (*types.Transaction, error) {
	return _IVault.Contract.DecreasePosition(&_IVault.TransactOpts, _account, _collateralToken, _indexToken, _collateralDelta, _sizeDelta, _isLong, _receiver)
}

// DirectPoolDeposit is a paid mutator transaction binding the contract method 0x5f7bc119.
//
// Solidity: function directPoolDeposit(address _token) returns()
func (_IVault *IVaultTransactor) DirectPoolDeposit(opts *bind.TransactOpts, _token common.Address) (*types.Transaction, error) {
	return _IVault.contract.Transact(opts, "directPoolDeposit", _token)
}

// DirectPoolDeposit is a paid mutator transaction binding the contract method 0x5f7bc119.
//
// Solidity: function directPoolDeposit(address _token) returns()
func (_IVault *IVaultSession) DirectPoolDeposit(_token common.Address) (*types.Transaction, error) {
	return _IVault.Contract.DirectPoolDeposit(&_IVault.TransactOpts, _token)
}

// DirectPoolDeposit is a paid mutator transaction binding the contract method 0x5f7bc119.
//
// Solidity: function directPoolDeposit(address _token) returns()
func (_IVault *IVaultTransactorSession) DirectPoolDeposit(_token common.Address) (*types.Transaction, error) {
	return _IVault.Contract.DirectPoolDeposit(&_IVault.TransactOpts, _token)
}

// IncreasePosition is a paid mutator transaction binding the contract method 0x48d91abf.
//
// Solidity: function increasePosition(address _account, address _collateralToken, address _indexToken, uint256 _sizeDelta, bool _isLong) returns()
func (_IVault *IVaultTransactor) IncreasePosition(opts *bind.TransactOpts, _account common.Address, _collateralToken common.Address, _indexToken common.Address, _sizeDelta *big.Int, _isLong bool) (*types.Transaction, error) {
	return _IVault.contract.Transact(opts, "increasePosition", _account, _collateralToken, _indexToken, _sizeDelta, _isLong)
}

// IncreasePosition is a paid mutator transaction binding the contract method 0x48d91abf.
//
// Solidity: function increasePosition(address _account, address _collateralToken, address _indexToken, uint256 _sizeDelta, bool _isLong) returns()
func (_IVault *IVaultSession) IncreasePosition(_account common.Address, _collateralToken common.Address, _indexToken common.Address, _sizeDelta *big.Int, _isLong bool) (*types.Transaction, error) {
	return _IVault.Contract.IncreasePosition(&_IVault.TransactOpts, _account, _collateralToken, _indexToken, _sizeDelta, _isLong)
}

// IncreasePosition is a paid mutator transaction binding the contract method 0x48d91abf.
//
// Solidity: function increasePosition(address _account, address _collateralToken, address _indexToken, uint256 _sizeDelta, bool _isLong) returns()
func (_IVault *IVaultTransactorSession) IncreasePosition(_account common.Address, _collateralToken common.Address, _indexToken common.Address, _sizeDelta *big.Int, _isLong bool) (*types.Transaction, error) {
	return _IVault.Contract.IncreasePosition(&_IVault.TransactOpts, _account, _collateralToken, _indexToken, _sizeDelta, _isLong)
}

// LiquidatePosition is a paid mutator transaction binding the contract method 0xde2ea948.
//
// Solidity: function liquidatePosition(address _account, address _collateralToken, address _indexToken, bool _isLong, address _feeReceiver) returns()
func (_IVault *IVaultTransactor) LiquidatePosition(opts *bind.TransactOpts, _account common.Address, _collateralToken common.Address, _indexToken common.Address, _isLong bool, _feeReceiver common.Address) (*types.Transaction, error) {
	return _IVault.contract.Transact(opts, "liquidatePosition", _account, _collateralToken, _indexToken, _isLong, _feeReceiver)
}

// LiquidatePosition is a paid mutator transaction binding the contract method 0xde2ea948.
//
// Solidity: function liquidatePosition(address _account, address _collateralToken, address _indexToken, bool _isLong, address _feeReceiver) returns()
func (_IVault *IVaultSession) LiquidatePosition(_account common.Address, _collateralToken common.Address, _indexToken common.Address, _isLong bool, _feeReceiver common.Address) (*types.Transaction, error) {
	return _IVault.Contract.LiquidatePosition(&_IVault.TransactOpts, _account, _collateralToken, _indexToken, _isLong, _feeReceiver)
}

// LiquidatePosition is a paid mutator transaction binding the contract method 0xde2ea948.
//
// Solidity: function liquidatePosition(address _account, address _collateralToken, address _indexToken, bool _isLong, address _feeReceiver) returns()
func (_IVault *IVaultTransactorSession) LiquidatePosition(_account common.Address, _collateralToken common.Address, _indexToken common.Address, _isLong bool, _feeReceiver common.Address) (*types.Transaction, error) {
	return _IVault.Contract.LiquidatePosition(&_IVault.TransactOpts, _account, _collateralToken, _indexToken, _isLong, _feeReceiver)
}

// SellUSDG is a paid mutator transaction binding the contract method 0x711e6190.
//
// Solidity: function sellUSDG(address _token, address _receiver) returns(uint256)
func (_IVault *IVaultTransactor) SellUSDG(opts *bind.TransactOpts, _token common.Address, _receiver common.Address) (*types.Transaction, error) {
	return _IVault.contract.Transact(opts, "sellUSDG", _token, _receiver)
}

// SellUSDG is a paid mutator transaction binding the contract method 0x711e6190.
//
// Solidity: function sellUSDG(address _token, address _receiver) returns(uint256)
func (_IVault *IVaultSession) SellUSDG(_token common.Address, _receiver common.Address) (*types.Transaction, error) {
	return _IVault.Contract.SellUSDG(&_IVault.TransactOpts, _token, _receiver)
}

// SellUSDG is a paid mutator transaction binding the contract method 0x711e6190.
//
// Solidity: function sellUSDG(address _token, address _receiver) returns(uint256)
func (_IVault *IVaultTransactorSession) SellUSDG(_token common.Address, _receiver common.Address) (*types.Transaction, error) {
	return _IVault.Contract.SellUSDG(&_IVault.TransactOpts, _token, _receiver)
}

// SetBufferAmount is a paid mutator transaction binding the contract method 0x8585f4d2.
//
// Solidity: function setBufferAmount(address _token, uint256 _amount) returns()
func (_IVault *IVaultTransactor) SetBufferAmount(opts *bind.TransactOpts, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IVault.contract.Transact(opts, "setBufferAmount", _token, _amount)
}

// SetBufferAmount is a paid mutator transaction binding the contract method 0x8585f4d2.
//
// Solidity: function setBufferAmount(address _token, uint256 _amount) returns()
func (_IVault *IVaultSession) SetBufferAmount(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IVault.Contract.SetBufferAmount(&_IVault.TransactOpts, _token, _amount)
}

// SetBufferAmount is a paid mutator transaction binding the contract method 0x8585f4d2.
//
// Solidity: function setBufferAmount(address _token, uint256 _amount) returns()
func (_IVault *IVaultTransactorSession) SetBufferAmount(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IVault.Contract.SetBufferAmount(&_IVault.TransactOpts, _token, _amount)
}

// SetError is a paid mutator transaction binding the contract method 0x28e67be5.
//
// Solidity: function setError(uint256 _errorCode, string _error) returns()
func (_IVault *IVaultTransactor) SetError(opts *bind.TransactOpts, _errorCode *big.Int, _error string) (*types.Transaction, error) {
	return _IVault.contract.Transact(opts, "setError", _errorCode, _error)
}

// SetError is a paid mutator transaction binding the contract method 0x28e67be5.
//
// Solidity: function setError(uint256 _errorCode, string _error) returns()
func (_IVault *IVaultSession) SetError(_errorCode *big.Int, _error string) (*types.Transaction, error) {
	return _IVault.Contract.SetError(&_IVault.TransactOpts, _errorCode, _error)
}

// SetError is a paid mutator transaction binding the contract method 0x28e67be5.
//
// Solidity: function setError(uint256 _errorCode, string _error) returns()
func (_IVault *IVaultTransactorSession) SetError(_errorCode *big.Int, _error string) (*types.Transaction, error) {
	return _IVault.Contract.SetError(&_IVault.TransactOpts, _errorCode, _error)
}

// SetFees is a paid mutator transaction binding the contract method 0x40eb3802.
//
// Solidity: function setFees(uint256 _taxBasisPoints, uint256 _stableTaxBasisPoints, uint256 _mintBurnFeeBasisPoints, uint256 _swapFeeBasisPoints, uint256 _stableSwapFeeBasisPoints, uint256 _marginFeeBasisPoints, uint256 _liquidationFeeUsd, uint256 _minProfitTime, bool _hasDynamicFees) returns()
func (_IVault *IVaultTransactor) SetFees(opts *bind.TransactOpts, _taxBasisPoints *big.Int, _stableTaxBasisPoints *big.Int, _mintBurnFeeBasisPoints *big.Int, _swapFeeBasisPoints *big.Int, _stableSwapFeeBasisPoints *big.Int, _marginFeeBasisPoints *big.Int, _liquidationFeeUsd *big.Int, _minProfitTime *big.Int, _hasDynamicFees bool) (*types.Transaction, error) {
	return _IVault.contract.Transact(opts, "setFees", _taxBasisPoints, _stableTaxBasisPoints, _mintBurnFeeBasisPoints, _swapFeeBasisPoints, _stableSwapFeeBasisPoints, _marginFeeBasisPoints, _liquidationFeeUsd, _minProfitTime, _hasDynamicFees)
}

// SetFees is a paid mutator transaction binding the contract method 0x40eb3802.
//
// Solidity: function setFees(uint256 _taxBasisPoints, uint256 _stableTaxBasisPoints, uint256 _mintBurnFeeBasisPoints, uint256 _swapFeeBasisPoints, uint256 _stableSwapFeeBasisPoints, uint256 _marginFeeBasisPoints, uint256 _liquidationFeeUsd, uint256 _minProfitTime, bool _hasDynamicFees) returns()
func (_IVault *IVaultSession) SetFees(_taxBasisPoints *big.Int, _stableTaxBasisPoints *big.Int, _mintBurnFeeBasisPoints *big.Int, _swapFeeBasisPoints *big.Int, _stableSwapFeeBasisPoints *big.Int, _marginFeeBasisPoints *big.Int, _liquidationFeeUsd *big.Int, _minProfitTime *big.Int, _hasDynamicFees bool) (*types.Transaction, error) {
	return _IVault.Contract.SetFees(&_IVault.TransactOpts, _taxBasisPoints, _stableTaxBasisPoints, _mintBurnFeeBasisPoints, _swapFeeBasisPoints, _stableSwapFeeBasisPoints, _marginFeeBasisPoints, _liquidationFeeUsd, _minProfitTime, _hasDynamicFees)
}

// SetFees is a paid mutator transaction binding the contract method 0x40eb3802.
//
// Solidity: function setFees(uint256 _taxBasisPoints, uint256 _stableTaxBasisPoints, uint256 _mintBurnFeeBasisPoints, uint256 _swapFeeBasisPoints, uint256 _stableSwapFeeBasisPoints, uint256 _marginFeeBasisPoints, uint256 _liquidationFeeUsd, uint256 _minProfitTime, bool _hasDynamicFees) returns()
func (_IVault *IVaultTransactorSession) SetFees(_taxBasisPoints *big.Int, _stableTaxBasisPoints *big.Int, _mintBurnFeeBasisPoints *big.Int, _swapFeeBasisPoints *big.Int, _stableSwapFeeBasisPoints *big.Int, _marginFeeBasisPoints *big.Int, _liquidationFeeUsd *big.Int, _minProfitTime *big.Int, _hasDynamicFees bool) (*types.Transaction, error) {
	return _IVault.Contract.SetFees(&_IVault.TransactOpts, _taxBasisPoints, _stableTaxBasisPoints, _mintBurnFeeBasisPoints, _swapFeeBasisPoints, _stableSwapFeeBasisPoints, _marginFeeBasisPoints, _liquidationFeeUsd, _minProfitTime, _hasDynamicFees)
}

// SetFundingRate is a paid mutator transaction binding the contract method 0x8a27d468.
//
// Solidity: function setFundingRate(uint256 _fundingInterval, uint256 _fundingRateFactor, uint256 _stableFundingRateFactor) returns()
func (_IVault *IVaultTransactor) SetFundingRate(opts *bind.TransactOpts, _fundingInterval *big.Int, _fundingRateFactor *big.Int, _stableFundingRateFactor *big.Int) (*types.Transaction, error) {
	return _IVault.contract.Transact(opts, "setFundingRate", _fundingInterval, _fundingRateFactor, _stableFundingRateFactor)
}

// SetFundingRate is a paid mutator transaction binding the contract method 0x8a27d468.
//
// Solidity: function setFundingRate(uint256 _fundingInterval, uint256 _fundingRateFactor, uint256 _stableFundingRateFactor) returns()
func (_IVault *IVaultSession) SetFundingRate(_fundingInterval *big.Int, _fundingRateFactor *big.Int, _stableFundingRateFactor *big.Int) (*types.Transaction, error) {
	return _IVault.Contract.SetFundingRate(&_IVault.TransactOpts, _fundingInterval, _fundingRateFactor, _stableFundingRateFactor)
}

// SetFundingRate is a paid mutator transaction binding the contract method 0x8a27d468.
//
// Solidity: function setFundingRate(uint256 _fundingInterval, uint256 _fundingRateFactor, uint256 _stableFundingRateFactor) returns()
func (_IVault *IVaultTransactorSession) SetFundingRate(_fundingInterval *big.Int, _fundingRateFactor *big.Int, _stableFundingRateFactor *big.Int) (*types.Transaction, error) {
	return _IVault.Contract.SetFundingRate(&_IVault.TransactOpts, _fundingInterval, _fundingRateFactor, _stableFundingRateFactor)
}

// SetInManagerMode is a paid mutator transaction binding the contract method 0x24b0c04d.
//
// Solidity: function setInManagerMode(bool _inManagerMode) returns()
func (_IVault *IVaultTransactor) SetInManagerMode(opts *bind.TransactOpts, _inManagerMode bool) (*types.Transaction, error) {
	return _IVault.contract.Transact(opts, "setInManagerMode", _inManagerMode)
}

// SetInManagerMode is a paid mutator transaction binding the contract method 0x24b0c04d.
//
// Solidity: function setInManagerMode(bool _inManagerMode) returns()
func (_IVault *IVaultSession) SetInManagerMode(_inManagerMode bool) (*types.Transaction, error) {
	return _IVault.Contract.SetInManagerMode(&_IVault.TransactOpts, _inManagerMode)
}

// SetInManagerMode is a paid mutator transaction binding the contract method 0x24b0c04d.
//
// Solidity: function setInManagerMode(bool _inManagerMode) returns()
func (_IVault *IVaultTransactorSession) SetInManagerMode(_inManagerMode bool) (*types.Transaction, error) {
	return _IVault.Contract.SetInManagerMode(&_IVault.TransactOpts, _inManagerMode)
}

// SetInPrivateLiquidationMode is a paid mutator transaction binding the contract method 0xf07bbf77.
//
// Solidity: function setInPrivateLiquidationMode(bool _inPrivateLiquidationMode) returns()
func (_IVault *IVaultTransactor) SetInPrivateLiquidationMode(opts *bind.TransactOpts, _inPrivateLiquidationMode bool) (*types.Transaction, error) {
	return _IVault.contract.Transact(opts, "setInPrivateLiquidationMode", _inPrivateLiquidationMode)
}

// SetInPrivateLiquidationMode is a paid mutator transaction binding the contract method 0xf07bbf77.
//
// Solidity: function setInPrivateLiquidationMode(bool _inPrivateLiquidationMode) returns()
func (_IVault *IVaultSession) SetInPrivateLiquidationMode(_inPrivateLiquidationMode bool) (*types.Transaction, error) {
	return _IVault.Contract.SetInPrivateLiquidationMode(&_IVault.TransactOpts, _inPrivateLiquidationMode)
}

// SetInPrivateLiquidationMode is a paid mutator transaction binding the contract method 0xf07bbf77.
//
// Solidity: function setInPrivateLiquidationMode(bool _inPrivateLiquidationMode) returns()
func (_IVault *IVaultTransactorSession) SetInPrivateLiquidationMode(_inPrivateLiquidationMode bool) (*types.Transaction, error) {
	return _IVault.Contract.SetInPrivateLiquidationMode(&_IVault.TransactOpts, _inPrivateLiquidationMode)
}

// SetIsLeverageEnabled is a paid mutator transaction binding the contract method 0x7c2eb9f7.
//
// Solidity: function setIsLeverageEnabled(bool _isLeverageEnabled) returns()
func (_IVault *IVaultTransactor) SetIsLeverageEnabled(opts *bind.TransactOpts, _isLeverageEnabled bool) (*types.Transaction, error) {
	return _IVault.contract.Transact(opts, "setIsLeverageEnabled", _isLeverageEnabled)
}

// SetIsLeverageEnabled is a paid mutator transaction binding the contract method 0x7c2eb9f7.
//
// Solidity: function setIsLeverageEnabled(bool _isLeverageEnabled) returns()
func (_IVault *IVaultSession) SetIsLeverageEnabled(_isLeverageEnabled bool) (*types.Transaction, error) {
	return _IVault.Contract.SetIsLeverageEnabled(&_IVault.TransactOpts, _isLeverageEnabled)
}

// SetIsLeverageEnabled is a paid mutator transaction binding the contract method 0x7c2eb9f7.
//
// Solidity: function setIsLeverageEnabled(bool _isLeverageEnabled) returns()
func (_IVault *IVaultTransactorSession) SetIsLeverageEnabled(_isLeverageEnabled bool) (*types.Transaction, error) {
	return _IVault.Contract.SetIsLeverageEnabled(&_IVault.TransactOpts, _isLeverageEnabled)
}

// SetIsSwapEnabled is a paid mutator transaction binding the contract method 0x30455ede.
//
// Solidity: function setIsSwapEnabled(bool _isSwapEnabled) returns()
func (_IVault *IVaultTransactor) SetIsSwapEnabled(opts *bind.TransactOpts, _isSwapEnabled bool) (*types.Transaction, error) {
	return _IVault.contract.Transact(opts, "setIsSwapEnabled", _isSwapEnabled)
}

// SetIsSwapEnabled is a paid mutator transaction binding the contract method 0x30455ede.
//
// Solidity: function setIsSwapEnabled(bool _isSwapEnabled) returns()
func (_IVault *IVaultSession) SetIsSwapEnabled(_isSwapEnabled bool) (*types.Transaction, error) {
	return _IVault.Contract.SetIsSwapEnabled(&_IVault.TransactOpts, _isSwapEnabled)
}

// SetIsSwapEnabled is a paid mutator transaction binding the contract method 0x30455ede.
//
// Solidity: function setIsSwapEnabled(bool _isSwapEnabled) returns()
func (_IVault *IVaultTransactorSession) SetIsSwapEnabled(_isSwapEnabled bool) (*types.Transaction, error) {
	return _IVault.Contract.SetIsSwapEnabled(&_IVault.TransactOpts, _isSwapEnabled)
}

// SetLiquidator is a paid mutator transaction binding the contract method 0x4453a374.
//
// Solidity: function setLiquidator(address _liquidator, bool _isActive) returns()
func (_IVault *IVaultTransactor) SetLiquidator(opts *bind.TransactOpts, _liquidator common.Address, _isActive bool) (*types.Transaction, error) {
	return _IVault.contract.Transact(opts, "setLiquidator", _liquidator, _isActive)
}

// SetLiquidator is a paid mutator transaction binding the contract method 0x4453a374.
//
// Solidity: function setLiquidator(address _liquidator, bool _isActive) returns()
func (_IVault *IVaultSession) SetLiquidator(_liquidator common.Address, _isActive bool) (*types.Transaction, error) {
	return _IVault.Contract.SetLiquidator(&_IVault.TransactOpts, _liquidator, _isActive)
}

// SetLiquidator is a paid mutator transaction binding the contract method 0x4453a374.
//
// Solidity: function setLiquidator(address _liquidator, bool _isActive) returns()
func (_IVault *IVaultTransactorSession) SetLiquidator(_liquidator common.Address, _isActive bool) (*types.Transaction, error) {
	return _IVault.Contract.SetLiquidator(&_IVault.TransactOpts, _liquidator, _isActive)
}

// SetManager is a paid mutator transaction binding the contract method 0xa5e90eee.
//
// Solidity: function setManager(address _manager, bool _isManager) returns()
func (_IVault *IVaultTransactor) SetManager(opts *bind.TransactOpts, _manager common.Address, _isManager bool) (*types.Transaction, error) {
	return _IVault.contract.Transact(opts, "setManager", _manager, _isManager)
}

// SetManager is a paid mutator transaction binding the contract method 0xa5e90eee.
//
// Solidity: function setManager(address _manager, bool _isManager) returns()
func (_IVault *IVaultSession) SetManager(_manager common.Address, _isManager bool) (*types.Transaction, error) {
	return _IVault.Contract.SetManager(&_IVault.TransactOpts, _manager, _isManager)
}

// SetManager is a paid mutator transaction binding the contract method 0xa5e90eee.
//
// Solidity: function setManager(address _manager, bool _isManager) returns()
func (_IVault *IVaultTransactorSession) SetManager(_manager common.Address, _isManager bool) (*types.Transaction, error) {
	return _IVault.Contract.SetManager(&_IVault.TransactOpts, _manager, _isManager)
}

// SetMaxGasPrice is a paid mutator transaction binding the contract method 0xd2fa635e.
//
// Solidity: function setMaxGasPrice(uint256 _maxGasPrice) returns()
func (_IVault *IVaultTransactor) SetMaxGasPrice(opts *bind.TransactOpts, _maxGasPrice *big.Int) (*types.Transaction, error) {
	return _IVault.contract.Transact(opts, "setMaxGasPrice", _maxGasPrice)
}

// SetMaxGasPrice is a paid mutator transaction binding the contract method 0xd2fa635e.
//
// Solidity: function setMaxGasPrice(uint256 _maxGasPrice) returns()
func (_IVault *IVaultSession) SetMaxGasPrice(_maxGasPrice *big.Int) (*types.Transaction, error) {
	return _IVault.Contract.SetMaxGasPrice(&_IVault.TransactOpts, _maxGasPrice)
}

// SetMaxGasPrice is a paid mutator transaction binding the contract method 0xd2fa635e.
//
// Solidity: function setMaxGasPrice(uint256 _maxGasPrice) returns()
func (_IVault *IVaultTransactorSession) SetMaxGasPrice(_maxGasPrice *big.Int) (*types.Transaction, error) {
	return _IVault.Contract.SetMaxGasPrice(&_IVault.TransactOpts, _maxGasPrice)
}

// SetMaxGlobalShortSize is a paid mutator transaction binding the contract method 0xefa10a6e.
//
// Solidity: function setMaxGlobalShortSize(address _token, uint256 _amount) returns()
func (_IVault *IVaultTransactor) SetMaxGlobalShortSize(opts *bind.TransactOpts, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IVault.contract.Transact(opts, "setMaxGlobalShortSize", _token, _amount)
}

// SetMaxGlobalShortSize is a paid mutator transaction binding the contract method 0xefa10a6e.
//
// Solidity: function setMaxGlobalShortSize(address _token, uint256 _amount) returns()
func (_IVault *IVaultSession) SetMaxGlobalShortSize(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IVault.Contract.SetMaxGlobalShortSize(&_IVault.TransactOpts, _token, _amount)
}

// SetMaxGlobalShortSize is a paid mutator transaction binding the contract method 0xefa10a6e.
//
// Solidity: function setMaxGlobalShortSize(address _token, uint256 _amount) returns()
func (_IVault *IVaultTransactorSession) SetMaxGlobalShortSize(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IVault.Contract.SetMaxGlobalShortSize(&_IVault.TransactOpts, _token, _amount)
}

// SetMaxLeverage is a paid mutator transaction binding the contract method 0xd3127e63.
//
// Solidity: function setMaxLeverage(uint256 _maxLeverage) returns()
func (_IVault *IVaultTransactor) SetMaxLeverage(opts *bind.TransactOpts, _maxLeverage *big.Int) (*types.Transaction, error) {
	return _IVault.contract.Transact(opts, "setMaxLeverage", _maxLeverage)
}

// SetMaxLeverage is a paid mutator transaction binding the contract method 0xd3127e63.
//
// Solidity: function setMaxLeverage(uint256 _maxLeverage) returns()
func (_IVault *IVaultSession) SetMaxLeverage(_maxLeverage *big.Int) (*types.Transaction, error) {
	return _IVault.Contract.SetMaxLeverage(&_IVault.TransactOpts, _maxLeverage)
}

// SetMaxLeverage is a paid mutator transaction binding the contract method 0xd3127e63.
//
// Solidity: function setMaxLeverage(uint256 _maxLeverage) returns()
func (_IVault *IVaultTransactorSession) SetMaxLeverage(_maxLeverage *big.Int) (*types.Transaction, error) {
	return _IVault.Contract.SetMaxLeverage(&_IVault.TransactOpts, _maxLeverage)
}

// SetPriceFeed is a paid mutator transaction binding the contract method 0x724e78da.
//
// Solidity: function setPriceFeed(address _priceFeed) returns()
func (_IVault *IVaultTransactor) SetPriceFeed(opts *bind.TransactOpts, _priceFeed common.Address) (*types.Transaction, error) {
	return _IVault.contract.Transact(opts, "setPriceFeed", _priceFeed)
}

// SetPriceFeed is a paid mutator transaction binding the contract method 0x724e78da.
//
// Solidity: function setPriceFeed(address _priceFeed) returns()
func (_IVault *IVaultSession) SetPriceFeed(_priceFeed common.Address) (*types.Transaction, error) {
	return _IVault.Contract.SetPriceFeed(&_IVault.TransactOpts, _priceFeed)
}

// SetPriceFeed is a paid mutator transaction binding the contract method 0x724e78da.
//
// Solidity: function setPriceFeed(address _priceFeed) returns()
func (_IVault *IVaultTransactorSession) SetPriceFeed(_priceFeed common.Address) (*types.Transaction, error) {
	return _IVault.Contract.SetPriceFeed(&_IVault.TransactOpts, _priceFeed)
}

// SetTokenConfig is a paid mutator transaction binding the contract method 0x3c5a6e35.
//
// Solidity: function setTokenConfig(address _token, uint256 _tokenDecimals, uint256 _redemptionBps, uint256 _minProfitBps, uint256 _maxUsdgAmount, bool _isStable, bool _isShortable) returns()
func (_IVault *IVaultTransactor) SetTokenConfig(opts *bind.TransactOpts, _token common.Address, _tokenDecimals *big.Int, _redemptionBps *big.Int, _minProfitBps *big.Int, _maxUsdgAmount *big.Int, _isStable bool, _isShortable bool) (*types.Transaction, error) {
	return _IVault.contract.Transact(opts, "setTokenConfig", _token, _tokenDecimals, _redemptionBps, _minProfitBps, _maxUsdgAmount, _isStable, _isShortable)
}

// SetTokenConfig is a paid mutator transaction binding the contract method 0x3c5a6e35.
//
// Solidity: function setTokenConfig(address _token, uint256 _tokenDecimals, uint256 _redemptionBps, uint256 _minProfitBps, uint256 _maxUsdgAmount, bool _isStable, bool _isShortable) returns()
func (_IVault *IVaultSession) SetTokenConfig(_token common.Address, _tokenDecimals *big.Int, _redemptionBps *big.Int, _minProfitBps *big.Int, _maxUsdgAmount *big.Int, _isStable bool, _isShortable bool) (*types.Transaction, error) {
	return _IVault.Contract.SetTokenConfig(&_IVault.TransactOpts, _token, _tokenDecimals, _redemptionBps, _minProfitBps, _maxUsdgAmount, _isStable, _isShortable)
}

// SetTokenConfig is a paid mutator transaction binding the contract method 0x3c5a6e35.
//
// Solidity: function setTokenConfig(address _token, uint256 _tokenDecimals, uint256 _redemptionBps, uint256 _minProfitBps, uint256 _maxUsdgAmount, bool _isStable, bool _isShortable) returns()
func (_IVault *IVaultTransactorSession) SetTokenConfig(_token common.Address, _tokenDecimals *big.Int, _redemptionBps *big.Int, _minProfitBps *big.Int, _maxUsdgAmount *big.Int, _isStable bool, _isShortable bool) (*types.Transaction, error) {
	return _IVault.Contract.SetTokenConfig(&_IVault.TransactOpts, _token, _tokenDecimals, _redemptionBps, _minProfitBps, _maxUsdgAmount, _isStable, _isShortable)
}

// SetUsdgAmount is a paid mutator transaction binding the contract method 0xd66b000d.
//
// Solidity: function setUsdgAmount(address _token, uint256 _amount) returns()
func (_IVault *IVaultTransactor) SetUsdgAmount(opts *bind.TransactOpts, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IVault.contract.Transact(opts, "setUsdgAmount", _token, _amount)
}

// SetUsdgAmount is a paid mutator transaction binding the contract method 0xd66b000d.
//
// Solidity: function setUsdgAmount(address _token, uint256 _amount) returns()
func (_IVault *IVaultSession) SetUsdgAmount(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IVault.Contract.SetUsdgAmount(&_IVault.TransactOpts, _token, _amount)
}

// SetUsdgAmount is a paid mutator transaction binding the contract method 0xd66b000d.
//
// Solidity: function setUsdgAmount(address _token, uint256 _amount) returns()
func (_IVault *IVaultTransactorSession) SetUsdgAmount(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IVault.Contract.SetUsdgAmount(&_IVault.TransactOpts, _token, _amount)
}

// SetVaultUtils is a paid mutator transaction binding the contract method 0x71089f4d.
//
// Solidity: function setVaultUtils(address _vaultUtils) returns()
func (_IVault *IVaultTransactor) SetVaultUtils(opts *bind.TransactOpts, _vaultUtils common.Address) (*types.Transaction, error) {
	return _IVault.contract.Transact(opts, "setVaultUtils", _vaultUtils)
}

// SetVaultUtils is a paid mutator transaction binding the contract method 0x71089f4d.
//
// Solidity: function setVaultUtils(address _vaultUtils) returns()
func (_IVault *IVaultSession) SetVaultUtils(_vaultUtils common.Address) (*types.Transaction, error) {
	return _IVault.Contract.SetVaultUtils(&_IVault.TransactOpts, _vaultUtils)
}

// SetVaultUtils is a paid mutator transaction binding the contract method 0x71089f4d.
//
// Solidity: function setVaultUtils(address _vaultUtils) returns()
func (_IVault *IVaultTransactorSession) SetVaultUtils(_vaultUtils common.Address) (*types.Transaction, error) {
	return _IVault.Contract.SetVaultUtils(&_IVault.TransactOpts, _vaultUtils)
}

// Swap is a paid mutator transaction binding the contract method 0x93316212.
//
// Solidity: function swap(address _tokenIn, address _tokenOut, address _receiver) returns(uint256)
func (_IVault *IVaultTransactor) Swap(opts *bind.TransactOpts, _tokenIn common.Address, _tokenOut common.Address, _receiver common.Address) (*types.Transaction, error) {
	return _IVault.contract.Transact(opts, "swap", _tokenIn, _tokenOut, _receiver)
}

// Swap is a paid mutator transaction binding the contract method 0x93316212.
//
// Solidity: function swap(address _tokenIn, address _tokenOut, address _receiver) returns(uint256)
func (_IVault *IVaultSession) Swap(_tokenIn common.Address, _tokenOut common.Address, _receiver common.Address) (*types.Transaction, error) {
	return _IVault.Contract.Swap(&_IVault.TransactOpts, _tokenIn, _tokenOut, _receiver)
}

// Swap is a paid mutator transaction binding the contract method 0x93316212.
//
// Solidity: function swap(address _tokenIn, address _tokenOut, address _receiver) returns(uint256)
func (_IVault *IVaultTransactorSession) Swap(_tokenIn common.Address, _tokenOut common.Address, _receiver common.Address) (*types.Transaction, error) {
	return _IVault.Contract.Swap(&_IVault.TransactOpts, _tokenIn, _tokenOut, _receiver)
}

// WithdrawFees is a paid mutator transaction binding the contract method 0xf2555278.
//
// Solidity: function withdrawFees(address _token, address _receiver) returns(uint256)
func (_IVault *IVaultTransactor) WithdrawFees(opts *bind.TransactOpts, _token common.Address, _receiver common.Address) (*types.Transaction, error) {
	return _IVault.contract.Transact(opts, "withdrawFees", _token, _receiver)
}

// WithdrawFees is a paid mutator transaction binding the contract method 0xf2555278.
//
// Solidity: function withdrawFees(address _token, address _receiver) returns(uint256)
func (_IVault *IVaultSession) WithdrawFees(_token common.Address, _receiver common.Address) (*types.Transaction, error) {
	return _IVault.Contract.WithdrawFees(&_IVault.TransactOpts, _token, _receiver)
}

// WithdrawFees is a paid mutator transaction binding the contract method 0xf2555278.
//
// Solidity: function withdrawFees(address _token, address _receiver) returns(uint256)
func (_IVault *IVaultTransactorSession) WithdrawFees(_token common.Address, _receiver common.Address) (*types.Transaction, error) {
	return _IVault.Contract.WithdrawFees(&_IVault.TransactOpts, _token, _receiver)
}

var (
	MethodSetAdjustment = "setAdjustment"

	MethodSetFavorPrimaryPrice = "setFavorPrimaryPrice"

	MethodSetIsAmmEnabled = "setIsAmmEnabled"

	MethodSetIsSecondaryPriceEnabled = "setIsSecondaryPriceEnabled"

	MethodSetMaxStrictPriceDeviation = "setMaxStrictPriceDeviation"

	MethodSetPriceSampleSpace = "setPriceSampleSpace"

	MethodSetSpreadBasisPoints = "setSpreadBasisPoints"

	MethodSetSpreadThresholdBasisPoints = "setSpreadThresholdBasisPoints"

	MethodSetUseV2Pricing = "setUseV2Pricing"

	MethodAdjustmentBasisPoints = "adjustmentBasisPoints"

	MethodGetAmmPrice = "getAmmPrice"

	MethodGetLatestPrimaryPrice = "getLatestPrimaryPrice"

	MethodGetPrice = "getPrice"

	MethodGetPrimaryPrice = "getPrimaryPrice"

	MethodIsAdjustmentAdditive = "isAdjustmentAdditive"
)

// IVaultPriceFeedABI is the input ABI used to generate the binding from.
const IVaultPriceFeedABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"adjustmentBasisPoints\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"getAmmPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"getLatestPrimaryPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_maximise\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"_includeAmmPrice\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"_useSwapPricing\",\"type\":\"bool\"}],\"name\":\"getPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_maximise\",\"type\":\"bool\"}],\"name\":\"getPrimaryPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"isAdjustmentAdditive\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_isAdditive\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"_adjustmentBps\",\"type\":\"uint256\"}],\"name\":\"setAdjustment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_favorPrimaryPrice\",\"type\":\"bool\"}],\"name\":\"setFavorPrimaryPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_isEnabled\",\"type\":\"bool\"}],\"name\":\"setIsAmmEnabled\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_isEnabled\",\"type\":\"bool\"}],\"name\":\"setIsSecondaryPriceEnabled\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_maxStrictPriceDeviation\",\"type\":\"uint256\"}],\"name\":\"setMaxStrictPriceDeviation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_priceSampleSpace\",\"type\":\"uint256\"}],\"name\":\"setPriceSampleSpace\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_spreadBasisPoints\",\"type\":\"uint256\"}],\"name\":\"setSpreadBasisPoints\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_spreadThresholdBasisPoints\",\"type\":\"uint256\"}],\"name\":\"setSpreadThresholdBasisPoints\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_priceFeed\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_priceDecimals\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_isStrictStable\",\"type\":\"bool\"}],\"name\":\"setTokenConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_useV2Pricing\",\"type\":\"bool\"}],\"name\":\"setUseV2Pricing\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IVaultPriceFeedFuncSigs maps the 4-byte function signature to its string representation.
var IVaultPriceFeedFuncSigs = map[string]string{
	"48cac277": "adjustmentBasisPoints(address)",
	"c2138d8c": "getAmmPrice(address)",
	"56bf9de4": "getLatestPrimaryPrice(address)",
	"2fc3a70a": "getPrice(address,bool,bool,bool)",
	"56c8c2c1": "getPrimaryPrice(address,bool)",
	"6ce8a44b": "isAdjustmentAdditive(address)",
	"d694376c": "setAdjustment(address,bool,uint256)",
	"604f37e9": "setFavorPrimaryPrice(bool)",
	"9917dc74": "setIsAmmEnabled(bool)",
	"eb1c92a9": "setIsSecondaryPriceEnabled(bool)",
	"2fbfe3d3": "setMaxStrictPriceDeviation(uint256)",
	"2fa03b8f": "setPriceSampleSpace(uint256)",
	"9b889380": "setSpreadBasisPoints(address,uint256)",
	"b731dd87": "setSpreadThresholdBasisPoints(uint256)",
	"4b9ade47": "setTokenConfig(address,address,uint256,bool)",
	"fd34ec40": "setUseV2Pricing(bool)",
}

// IVaultPriceFeed is an auto generated Go binding around an Ethereum contract.
type IVaultPriceFeed struct {
	IVaultPriceFeedCaller     // Read-only binding to the contract
	IVaultPriceFeedTransactor // Write-only binding to the contract
	IVaultPriceFeedFilterer   // Log filterer for contract events
}

// IVaultPriceFeedCaller is an auto generated read-only Go binding around an Ethereum contract.
type IVaultPriceFeedCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IVaultPriceFeedTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IVaultPriceFeedTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IVaultPriceFeedFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IVaultPriceFeedFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IVaultPriceFeedSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IVaultPriceFeedSession struct {
	Contract     *IVaultPriceFeed  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IVaultPriceFeedCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IVaultPriceFeedCallerSession struct {
	Contract *IVaultPriceFeedCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// IVaultPriceFeedTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IVaultPriceFeedTransactorSession struct {
	Contract     *IVaultPriceFeedTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// IVaultPriceFeedRaw is an auto generated low-level Go binding around an Ethereum contract.
type IVaultPriceFeedRaw struct {
	Contract *IVaultPriceFeed // Generic contract binding to access the raw methods on
}

// IVaultPriceFeedCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IVaultPriceFeedCallerRaw struct {
	Contract *IVaultPriceFeedCaller // Generic read-only contract binding to access the raw methods on
}

// IVaultPriceFeedTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IVaultPriceFeedTransactorRaw struct {
	Contract *IVaultPriceFeedTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIVaultPriceFeed creates a new instance of IVaultPriceFeed, bound to a specific deployed contract.
func NewIVaultPriceFeed(address common.Address, backend bind.ContractBackend) (*IVaultPriceFeed, error) {
	contract, err := bindIVaultPriceFeed(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IVaultPriceFeed{IVaultPriceFeedCaller: IVaultPriceFeedCaller{contract: contract}, IVaultPriceFeedTransactor: IVaultPriceFeedTransactor{contract: contract}, IVaultPriceFeedFilterer: IVaultPriceFeedFilterer{contract: contract}}, nil
}

// NewIVaultPriceFeedCaller creates a new read-only instance of IVaultPriceFeed, bound to a specific deployed contract.
func NewIVaultPriceFeedCaller(address common.Address, caller bind.ContractCaller) (*IVaultPriceFeedCaller, error) {
	contract, err := bindIVaultPriceFeed(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IVaultPriceFeedCaller{contract: contract}, nil
}

// NewIVaultPriceFeedTransactor creates a new write-only instance of IVaultPriceFeed, bound to a specific deployed contract.
func NewIVaultPriceFeedTransactor(address common.Address, transactor bind.ContractTransactor) (*IVaultPriceFeedTransactor, error) {
	contract, err := bindIVaultPriceFeed(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IVaultPriceFeedTransactor{contract: contract}, nil
}

// NewIVaultPriceFeedFilterer creates a new log filterer instance of IVaultPriceFeed, bound to a specific deployed contract.
func NewIVaultPriceFeedFilterer(address common.Address, filterer bind.ContractFilterer) (*IVaultPriceFeedFilterer, error) {
	contract, err := bindIVaultPriceFeed(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IVaultPriceFeedFilterer{contract: contract}, nil
}

// bindIVaultPriceFeed binds a generic wrapper to an already deployed contract.
func bindIVaultPriceFeed(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IVaultPriceFeedABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IVaultPriceFeed *IVaultPriceFeedRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IVaultPriceFeed.Contract.IVaultPriceFeedCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IVaultPriceFeed *IVaultPriceFeedRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IVaultPriceFeed.Contract.IVaultPriceFeedTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IVaultPriceFeed *IVaultPriceFeedRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IVaultPriceFeed.Contract.IVaultPriceFeedTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IVaultPriceFeed *IVaultPriceFeedCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IVaultPriceFeed.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IVaultPriceFeed *IVaultPriceFeedTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IVaultPriceFeed.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IVaultPriceFeed *IVaultPriceFeedTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IVaultPriceFeed.Contract.contract.Transact(opts, method, params...)
}

// AdjustmentBasisPoints is a free data retrieval call binding the contract method 0x48cac277.
//
// Solidity: function adjustmentBasisPoints(address _token) view returns(uint256)
func (_IVaultPriceFeed *IVaultPriceFeedCaller) AdjustmentBasisPoints(opts *bind.CallOpts, _token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IVaultPriceFeed.contract.Call(opts, &out, "adjustmentBasisPoints", _token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AdjustmentBasisPoints is a free data retrieval call binding the contract method 0x48cac277.
//
// Solidity: function adjustmentBasisPoints(address _token) view returns(uint256)
func (_IVaultPriceFeed *IVaultPriceFeedSession) AdjustmentBasisPoints(_token common.Address) (*big.Int, error) {
	return _IVaultPriceFeed.Contract.AdjustmentBasisPoints(&_IVaultPriceFeed.CallOpts, _token)
}

// AdjustmentBasisPoints is a free data retrieval call binding the contract method 0x48cac277.
//
// Solidity: function adjustmentBasisPoints(address _token) view returns(uint256)
func (_IVaultPriceFeed *IVaultPriceFeedCallerSession) AdjustmentBasisPoints(_token common.Address) (*big.Int, error) {
	return _IVaultPriceFeed.Contract.AdjustmentBasisPoints(&_IVaultPriceFeed.CallOpts, _token)
}

// GetAmmPrice is a free data retrieval call binding the contract method 0xc2138d8c.
//
// Solidity: function getAmmPrice(address _token) view returns(uint256)
func (_IVaultPriceFeed *IVaultPriceFeedCaller) GetAmmPrice(opts *bind.CallOpts, _token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IVaultPriceFeed.contract.Call(opts, &out, "getAmmPrice", _token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAmmPrice is a free data retrieval call binding the contract method 0xc2138d8c.
//
// Solidity: function getAmmPrice(address _token) view returns(uint256)
func (_IVaultPriceFeed *IVaultPriceFeedSession) GetAmmPrice(_token common.Address) (*big.Int, error) {
	return _IVaultPriceFeed.Contract.GetAmmPrice(&_IVaultPriceFeed.CallOpts, _token)
}

// GetAmmPrice is a free data retrieval call binding the contract method 0xc2138d8c.
//
// Solidity: function getAmmPrice(address _token) view returns(uint256)
func (_IVaultPriceFeed *IVaultPriceFeedCallerSession) GetAmmPrice(_token common.Address) (*big.Int, error) {
	return _IVaultPriceFeed.Contract.GetAmmPrice(&_IVaultPriceFeed.CallOpts, _token)
}

// GetLatestPrimaryPrice is a free data retrieval call binding the contract method 0x56bf9de4.
//
// Solidity: function getLatestPrimaryPrice(address _token) view returns(uint256)
func (_IVaultPriceFeed *IVaultPriceFeedCaller) GetLatestPrimaryPrice(opts *bind.CallOpts, _token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IVaultPriceFeed.contract.Call(opts, &out, "getLatestPrimaryPrice", _token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLatestPrimaryPrice is a free data retrieval call binding the contract method 0x56bf9de4.
//
// Solidity: function getLatestPrimaryPrice(address _token) view returns(uint256)
func (_IVaultPriceFeed *IVaultPriceFeedSession) GetLatestPrimaryPrice(_token common.Address) (*big.Int, error) {
	return _IVaultPriceFeed.Contract.GetLatestPrimaryPrice(&_IVaultPriceFeed.CallOpts, _token)
}

// GetLatestPrimaryPrice is a free data retrieval call binding the contract method 0x56bf9de4.
//
// Solidity: function getLatestPrimaryPrice(address _token) view returns(uint256)
func (_IVaultPriceFeed *IVaultPriceFeedCallerSession) GetLatestPrimaryPrice(_token common.Address) (*big.Int, error) {
	return _IVaultPriceFeed.Contract.GetLatestPrimaryPrice(&_IVaultPriceFeed.CallOpts, _token)
}

// GetPrice is a free data retrieval call binding the contract method 0x2fc3a70a.
//
// Solidity: function getPrice(address _token, bool _maximise, bool _includeAmmPrice, bool _useSwapPricing) view returns(uint256)
func (_IVaultPriceFeed *IVaultPriceFeedCaller) GetPrice(opts *bind.CallOpts, _token common.Address, _maximise bool, _includeAmmPrice bool, _useSwapPricing bool) (*big.Int, error) {
	var out []interface{}
	err := _IVaultPriceFeed.contract.Call(opts, &out, "getPrice", _token, _maximise, _includeAmmPrice, _useSwapPricing)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPrice is a free data retrieval call binding the contract method 0x2fc3a70a.
//
// Solidity: function getPrice(address _token, bool _maximise, bool _includeAmmPrice, bool _useSwapPricing) view returns(uint256)
func (_IVaultPriceFeed *IVaultPriceFeedSession) GetPrice(_token common.Address, _maximise bool, _includeAmmPrice bool, _useSwapPricing bool) (*big.Int, error) {
	return _IVaultPriceFeed.Contract.GetPrice(&_IVaultPriceFeed.CallOpts, _token, _maximise, _includeAmmPrice, _useSwapPricing)
}

// GetPrice is a free data retrieval call binding the contract method 0x2fc3a70a.
//
// Solidity: function getPrice(address _token, bool _maximise, bool _includeAmmPrice, bool _useSwapPricing) view returns(uint256)
func (_IVaultPriceFeed *IVaultPriceFeedCallerSession) GetPrice(_token common.Address, _maximise bool, _includeAmmPrice bool, _useSwapPricing bool) (*big.Int, error) {
	return _IVaultPriceFeed.Contract.GetPrice(&_IVaultPriceFeed.CallOpts, _token, _maximise, _includeAmmPrice, _useSwapPricing)
}

// GetPrimaryPrice is a free data retrieval call binding the contract method 0x56c8c2c1.
//
// Solidity: function getPrimaryPrice(address _token, bool _maximise) view returns(uint256)
func (_IVaultPriceFeed *IVaultPriceFeedCaller) GetPrimaryPrice(opts *bind.CallOpts, _token common.Address, _maximise bool) (*big.Int, error) {
	var out []interface{}
	err := _IVaultPriceFeed.contract.Call(opts, &out, "getPrimaryPrice", _token, _maximise)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPrimaryPrice is a free data retrieval call binding the contract method 0x56c8c2c1.
//
// Solidity: function getPrimaryPrice(address _token, bool _maximise) view returns(uint256)
func (_IVaultPriceFeed *IVaultPriceFeedSession) GetPrimaryPrice(_token common.Address, _maximise bool) (*big.Int, error) {
	return _IVaultPriceFeed.Contract.GetPrimaryPrice(&_IVaultPriceFeed.CallOpts, _token, _maximise)
}

// GetPrimaryPrice is a free data retrieval call binding the contract method 0x56c8c2c1.
//
// Solidity: function getPrimaryPrice(address _token, bool _maximise) view returns(uint256)
func (_IVaultPriceFeed *IVaultPriceFeedCallerSession) GetPrimaryPrice(_token common.Address, _maximise bool) (*big.Int, error) {
	return _IVaultPriceFeed.Contract.GetPrimaryPrice(&_IVaultPriceFeed.CallOpts, _token, _maximise)
}

// IsAdjustmentAdditive is a free data retrieval call binding the contract method 0x6ce8a44b.
//
// Solidity: function isAdjustmentAdditive(address _token) view returns(bool)
func (_IVaultPriceFeed *IVaultPriceFeedCaller) IsAdjustmentAdditive(opts *bind.CallOpts, _token common.Address) (bool, error) {
	var out []interface{}
	err := _IVaultPriceFeed.contract.Call(opts, &out, "isAdjustmentAdditive", _token)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAdjustmentAdditive is a free data retrieval call binding the contract method 0x6ce8a44b.
//
// Solidity: function isAdjustmentAdditive(address _token) view returns(bool)
func (_IVaultPriceFeed *IVaultPriceFeedSession) IsAdjustmentAdditive(_token common.Address) (bool, error) {
	return _IVaultPriceFeed.Contract.IsAdjustmentAdditive(&_IVaultPriceFeed.CallOpts, _token)
}

// IsAdjustmentAdditive is a free data retrieval call binding the contract method 0x6ce8a44b.
//
// Solidity: function isAdjustmentAdditive(address _token) view returns(bool)
func (_IVaultPriceFeed *IVaultPriceFeedCallerSession) IsAdjustmentAdditive(_token common.Address) (bool, error) {
	return _IVaultPriceFeed.Contract.IsAdjustmentAdditive(&_IVaultPriceFeed.CallOpts, _token)
}

// SetAdjustment is a paid mutator transaction binding the contract method 0xd694376c.
//
// Solidity: function setAdjustment(address _token, bool _isAdditive, uint256 _adjustmentBps) returns()
func (_IVaultPriceFeed *IVaultPriceFeedTransactor) SetAdjustment(opts *bind.TransactOpts, _token common.Address, _isAdditive bool, _adjustmentBps *big.Int) (*types.Transaction, error) {
	return _IVaultPriceFeed.contract.Transact(opts, "setAdjustment", _token, _isAdditive, _adjustmentBps)
}

// SetAdjustment is a paid mutator transaction binding the contract method 0xd694376c.
//
// Solidity: function setAdjustment(address _token, bool _isAdditive, uint256 _adjustmentBps) returns()
func (_IVaultPriceFeed *IVaultPriceFeedSession) SetAdjustment(_token common.Address, _isAdditive bool, _adjustmentBps *big.Int) (*types.Transaction, error) {
	return _IVaultPriceFeed.Contract.SetAdjustment(&_IVaultPriceFeed.TransactOpts, _token, _isAdditive, _adjustmentBps)
}

// SetAdjustment is a paid mutator transaction binding the contract method 0xd694376c.
//
// Solidity: function setAdjustment(address _token, bool _isAdditive, uint256 _adjustmentBps) returns()
func (_IVaultPriceFeed *IVaultPriceFeedTransactorSession) SetAdjustment(_token common.Address, _isAdditive bool, _adjustmentBps *big.Int) (*types.Transaction, error) {
	return _IVaultPriceFeed.Contract.SetAdjustment(&_IVaultPriceFeed.TransactOpts, _token, _isAdditive, _adjustmentBps)
}

// SetFavorPrimaryPrice is a paid mutator transaction binding the contract method 0x604f37e9.
//
// Solidity: function setFavorPrimaryPrice(bool _favorPrimaryPrice) returns()
func (_IVaultPriceFeed *IVaultPriceFeedTransactor) SetFavorPrimaryPrice(opts *bind.TransactOpts, _favorPrimaryPrice bool) (*types.Transaction, error) {
	return _IVaultPriceFeed.contract.Transact(opts, "setFavorPrimaryPrice", _favorPrimaryPrice)
}

// SetFavorPrimaryPrice is a paid mutator transaction binding the contract method 0x604f37e9.
//
// Solidity: function setFavorPrimaryPrice(bool _favorPrimaryPrice) returns()
func (_IVaultPriceFeed *IVaultPriceFeedSession) SetFavorPrimaryPrice(_favorPrimaryPrice bool) (*types.Transaction, error) {
	return _IVaultPriceFeed.Contract.SetFavorPrimaryPrice(&_IVaultPriceFeed.TransactOpts, _favorPrimaryPrice)
}

// SetFavorPrimaryPrice is a paid mutator transaction binding the contract method 0x604f37e9.
//
// Solidity: function setFavorPrimaryPrice(bool _favorPrimaryPrice) returns()
func (_IVaultPriceFeed *IVaultPriceFeedTransactorSession) SetFavorPrimaryPrice(_favorPrimaryPrice bool) (*types.Transaction, error) {
	return _IVaultPriceFeed.Contract.SetFavorPrimaryPrice(&_IVaultPriceFeed.TransactOpts, _favorPrimaryPrice)
}

// SetIsAmmEnabled is a paid mutator transaction binding the contract method 0x9917dc74.
//
// Solidity: function setIsAmmEnabled(bool _isEnabled) returns()
func (_IVaultPriceFeed *IVaultPriceFeedTransactor) SetIsAmmEnabled(opts *bind.TransactOpts, _isEnabled bool) (*types.Transaction, error) {
	return _IVaultPriceFeed.contract.Transact(opts, "setIsAmmEnabled", _isEnabled)
}

// SetIsAmmEnabled is a paid mutator transaction binding the contract method 0x9917dc74.
//
// Solidity: function setIsAmmEnabled(bool _isEnabled) returns()
func (_IVaultPriceFeed *IVaultPriceFeedSession) SetIsAmmEnabled(_isEnabled bool) (*types.Transaction, error) {
	return _IVaultPriceFeed.Contract.SetIsAmmEnabled(&_IVaultPriceFeed.TransactOpts, _isEnabled)
}

// SetIsAmmEnabled is a paid mutator transaction binding the contract method 0x9917dc74.
//
// Solidity: function setIsAmmEnabled(bool _isEnabled) returns()
func (_IVaultPriceFeed *IVaultPriceFeedTransactorSession) SetIsAmmEnabled(_isEnabled bool) (*types.Transaction, error) {
	return _IVaultPriceFeed.Contract.SetIsAmmEnabled(&_IVaultPriceFeed.TransactOpts, _isEnabled)
}

// SetIsSecondaryPriceEnabled is a paid mutator transaction binding the contract method 0xeb1c92a9.
//
// Solidity: function setIsSecondaryPriceEnabled(bool _isEnabled) returns()
func (_IVaultPriceFeed *IVaultPriceFeedTransactor) SetIsSecondaryPriceEnabled(opts *bind.TransactOpts, _isEnabled bool) (*types.Transaction, error) {
	return _IVaultPriceFeed.contract.Transact(opts, "setIsSecondaryPriceEnabled", _isEnabled)
}

// SetIsSecondaryPriceEnabled is a paid mutator transaction binding the contract method 0xeb1c92a9.
//
// Solidity: function setIsSecondaryPriceEnabled(bool _isEnabled) returns()
func (_IVaultPriceFeed *IVaultPriceFeedSession) SetIsSecondaryPriceEnabled(_isEnabled bool) (*types.Transaction, error) {
	return _IVaultPriceFeed.Contract.SetIsSecondaryPriceEnabled(&_IVaultPriceFeed.TransactOpts, _isEnabled)
}

// SetIsSecondaryPriceEnabled is a paid mutator transaction binding the contract method 0xeb1c92a9.
//
// Solidity: function setIsSecondaryPriceEnabled(bool _isEnabled) returns()
func (_IVaultPriceFeed *IVaultPriceFeedTransactorSession) SetIsSecondaryPriceEnabled(_isEnabled bool) (*types.Transaction, error) {
	return _IVaultPriceFeed.Contract.SetIsSecondaryPriceEnabled(&_IVaultPriceFeed.TransactOpts, _isEnabled)
}

// SetMaxStrictPriceDeviation is a paid mutator transaction binding the contract method 0x2fbfe3d3.
//
// Solidity: function setMaxStrictPriceDeviation(uint256 _maxStrictPriceDeviation) returns()
func (_IVaultPriceFeed *IVaultPriceFeedTransactor) SetMaxStrictPriceDeviation(opts *bind.TransactOpts, _maxStrictPriceDeviation *big.Int) (*types.Transaction, error) {
	return _IVaultPriceFeed.contract.Transact(opts, "setMaxStrictPriceDeviation", _maxStrictPriceDeviation)
}

// SetMaxStrictPriceDeviation is a paid mutator transaction binding the contract method 0x2fbfe3d3.
//
// Solidity: function setMaxStrictPriceDeviation(uint256 _maxStrictPriceDeviation) returns()
func (_IVaultPriceFeed *IVaultPriceFeedSession) SetMaxStrictPriceDeviation(_maxStrictPriceDeviation *big.Int) (*types.Transaction, error) {
	return _IVaultPriceFeed.Contract.SetMaxStrictPriceDeviation(&_IVaultPriceFeed.TransactOpts, _maxStrictPriceDeviation)
}

// SetMaxStrictPriceDeviation is a paid mutator transaction binding the contract method 0x2fbfe3d3.
//
// Solidity: function setMaxStrictPriceDeviation(uint256 _maxStrictPriceDeviation) returns()
func (_IVaultPriceFeed *IVaultPriceFeedTransactorSession) SetMaxStrictPriceDeviation(_maxStrictPriceDeviation *big.Int) (*types.Transaction, error) {
	return _IVaultPriceFeed.Contract.SetMaxStrictPriceDeviation(&_IVaultPriceFeed.TransactOpts, _maxStrictPriceDeviation)
}

// SetPriceSampleSpace is a paid mutator transaction binding the contract method 0x2fa03b8f.
//
// Solidity: function setPriceSampleSpace(uint256 _priceSampleSpace) returns()
func (_IVaultPriceFeed *IVaultPriceFeedTransactor) SetPriceSampleSpace(opts *bind.TransactOpts, _priceSampleSpace *big.Int) (*types.Transaction, error) {
	return _IVaultPriceFeed.contract.Transact(opts, "setPriceSampleSpace", _priceSampleSpace)
}

// SetPriceSampleSpace is a paid mutator transaction binding the contract method 0x2fa03b8f.
//
// Solidity: function setPriceSampleSpace(uint256 _priceSampleSpace) returns()
func (_IVaultPriceFeed *IVaultPriceFeedSession) SetPriceSampleSpace(_priceSampleSpace *big.Int) (*types.Transaction, error) {
	return _IVaultPriceFeed.Contract.SetPriceSampleSpace(&_IVaultPriceFeed.TransactOpts, _priceSampleSpace)
}

// SetPriceSampleSpace is a paid mutator transaction binding the contract method 0x2fa03b8f.
//
// Solidity: function setPriceSampleSpace(uint256 _priceSampleSpace) returns()
func (_IVaultPriceFeed *IVaultPriceFeedTransactorSession) SetPriceSampleSpace(_priceSampleSpace *big.Int) (*types.Transaction, error) {
	return _IVaultPriceFeed.Contract.SetPriceSampleSpace(&_IVaultPriceFeed.TransactOpts, _priceSampleSpace)
}

// SetSpreadBasisPoints is a paid mutator transaction binding the contract method 0x9b889380.
//
// Solidity: function setSpreadBasisPoints(address _token, uint256 _spreadBasisPoints) returns()
func (_IVaultPriceFeed *IVaultPriceFeedTransactor) SetSpreadBasisPoints(opts *bind.TransactOpts, _token common.Address, _spreadBasisPoints *big.Int) (*types.Transaction, error) {
	return _IVaultPriceFeed.contract.Transact(opts, "setSpreadBasisPoints", _token, _spreadBasisPoints)
}

// SetSpreadBasisPoints is a paid mutator transaction binding the contract method 0x9b889380.
//
// Solidity: function setSpreadBasisPoints(address _token, uint256 _spreadBasisPoints) returns()
func (_IVaultPriceFeed *IVaultPriceFeedSession) SetSpreadBasisPoints(_token common.Address, _spreadBasisPoints *big.Int) (*types.Transaction, error) {
	return _IVaultPriceFeed.Contract.SetSpreadBasisPoints(&_IVaultPriceFeed.TransactOpts, _token, _spreadBasisPoints)
}

// SetSpreadBasisPoints is a paid mutator transaction binding the contract method 0x9b889380.
//
// Solidity: function setSpreadBasisPoints(address _token, uint256 _spreadBasisPoints) returns()
func (_IVaultPriceFeed *IVaultPriceFeedTransactorSession) SetSpreadBasisPoints(_token common.Address, _spreadBasisPoints *big.Int) (*types.Transaction, error) {
	return _IVaultPriceFeed.Contract.SetSpreadBasisPoints(&_IVaultPriceFeed.TransactOpts, _token, _spreadBasisPoints)
}

// SetSpreadThresholdBasisPoints is a paid mutator transaction binding the contract method 0xb731dd87.
//
// Solidity: function setSpreadThresholdBasisPoints(uint256 _spreadThresholdBasisPoints) returns()
func (_IVaultPriceFeed *IVaultPriceFeedTransactor) SetSpreadThresholdBasisPoints(opts *bind.TransactOpts, _spreadThresholdBasisPoints *big.Int) (*types.Transaction, error) {
	return _IVaultPriceFeed.contract.Transact(opts, "setSpreadThresholdBasisPoints", _spreadThresholdBasisPoints)
}

// SetSpreadThresholdBasisPoints is a paid mutator transaction binding the contract method 0xb731dd87.
//
// Solidity: function setSpreadThresholdBasisPoints(uint256 _spreadThresholdBasisPoints) returns()
func (_IVaultPriceFeed *IVaultPriceFeedSession) SetSpreadThresholdBasisPoints(_spreadThresholdBasisPoints *big.Int) (*types.Transaction, error) {
	return _IVaultPriceFeed.Contract.SetSpreadThresholdBasisPoints(&_IVaultPriceFeed.TransactOpts, _spreadThresholdBasisPoints)
}

// SetSpreadThresholdBasisPoints is a paid mutator transaction binding the contract method 0xb731dd87.
//
// Solidity: function setSpreadThresholdBasisPoints(uint256 _spreadThresholdBasisPoints) returns()
func (_IVaultPriceFeed *IVaultPriceFeedTransactorSession) SetSpreadThresholdBasisPoints(_spreadThresholdBasisPoints *big.Int) (*types.Transaction, error) {
	return _IVaultPriceFeed.Contract.SetSpreadThresholdBasisPoints(&_IVaultPriceFeed.TransactOpts, _spreadThresholdBasisPoints)
}

// SetTokenConfig is a paid mutator transaction binding the contract method 0x4b9ade47.
//
// Solidity: function setTokenConfig(address _token, address _priceFeed, uint256 _priceDecimals, bool _isStrictStable) returns()
func (_IVaultPriceFeed *IVaultPriceFeedTransactor) SetTokenConfig(opts *bind.TransactOpts, _token common.Address, _priceFeed common.Address, _priceDecimals *big.Int, _isStrictStable bool) (*types.Transaction, error) {
	return _IVaultPriceFeed.contract.Transact(opts, "setTokenConfig", _token, _priceFeed, _priceDecimals, _isStrictStable)
}

// SetTokenConfig is a paid mutator transaction binding the contract method 0x4b9ade47.
//
// Solidity: function setTokenConfig(address _token, address _priceFeed, uint256 _priceDecimals, bool _isStrictStable) returns()
func (_IVaultPriceFeed *IVaultPriceFeedSession) SetTokenConfig(_token common.Address, _priceFeed common.Address, _priceDecimals *big.Int, _isStrictStable bool) (*types.Transaction, error) {
	return _IVaultPriceFeed.Contract.SetTokenConfig(&_IVaultPriceFeed.TransactOpts, _token, _priceFeed, _priceDecimals, _isStrictStable)
}

// SetTokenConfig is a paid mutator transaction binding the contract method 0x4b9ade47.
//
// Solidity: function setTokenConfig(address _token, address _priceFeed, uint256 _priceDecimals, bool _isStrictStable) returns()
func (_IVaultPriceFeed *IVaultPriceFeedTransactorSession) SetTokenConfig(_token common.Address, _priceFeed common.Address, _priceDecimals *big.Int, _isStrictStable bool) (*types.Transaction, error) {
	return _IVaultPriceFeed.Contract.SetTokenConfig(&_IVaultPriceFeed.TransactOpts, _token, _priceFeed, _priceDecimals, _isStrictStable)
}

// SetUseV2Pricing is a paid mutator transaction binding the contract method 0xfd34ec40.
//
// Solidity: function setUseV2Pricing(bool _useV2Pricing) returns()
func (_IVaultPriceFeed *IVaultPriceFeedTransactor) SetUseV2Pricing(opts *bind.TransactOpts, _useV2Pricing bool) (*types.Transaction, error) {
	return _IVaultPriceFeed.contract.Transact(opts, "setUseV2Pricing", _useV2Pricing)
}

// SetUseV2Pricing is a paid mutator transaction binding the contract method 0xfd34ec40.
//
// Solidity: function setUseV2Pricing(bool _useV2Pricing) returns()
func (_IVaultPriceFeed *IVaultPriceFeedSession) SetUseV2Pricing(_useV2Pricing bool) (*types.Transaction, error) {
	return _IVaultPriceFeed.Contract.SetUseV2Pricing(&_IVaultPriceFeed.TransactOpts, _useV2Pricing)
}

// SetUseV2Pricing is a paid mutator transaction binding the contract method 0xfd34ec40.
//
// Solidity: function setUseV2Pricing(bool _useV2Pricing) returns()
func (_IVaultPriceFeed *IVaultPriceFeedTransactorSession) SetUseV2Pricing(_useV2Pricing bool) (*types.Transaction, error) {
	return _IVaultPriceFeed.Contract.SetUseV2Pricing(&_IVaultPriceFeed.TransactOpts, _useV2Pricing)
}

var (
	MethodUpdateCumulativeFundingRate = "updateCumulativeFundingRate"

	MethodGetBuyUsdgFeeBasisPoints = "getBuyUsdgFeeBasisPoints"

	MethodGetEntryFundingRate = "getEntryFundingRate"

	MethodGetFundingFee = "getFundingFee"

	MethodGetPositionFee = "getPositionFee"

	MethodGetSellUsdgFeeBasisPoints = "getSellUsdgFeeBasisPoints"

	MethodGetSwapFeeBasisPoints = "getSwapFeeBasisPoints"

	MethodValidateDecreasePosition = "validateDecreasePosition"

	MethodValidateIncreasePosition = "validateIncreasePosition"
)

// IVaultUtilsABI is the input ABI used to generate the binding from.
const IVaultUtilsABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_usdgAmount\",\"type\":\"uint256\"}],\"name\":\"getBuyUsdgFeeBasisPoints\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_collateralToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_indexToken\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_isLong\",\"type\":\"bool\"}],\"name\":\"getEntryFundingRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_usdgDelta\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_feeBasisPoints\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_taxBasisPoints\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_increment\",\"type\":\"bool\"}],\"name\":\"getFeeBasisPoints\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_collateralToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_indexToken\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_isLong\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"_size\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_entryFundingRate\",\"type\":\"uint256\"}],\"name\":\"getFundingFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_collateralToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_indexToken\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_isLong\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"_sizeDelta\",\"type\":\"uint256\"}],\"name\":\"getPositionFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_usdgAmount\",\"type\":\"uint256\"}],\"name\":\"getSellUsdgFeeBasisPoints\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_usdgAmount\",\"type\":\"uint256\"}],\"name\":\"getSwapFeeBasisPoints\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_collateralToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_indexToken\",\"type\":\"address\"}],\"name\":\"updateCumulativeFundingRate\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_collateralToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_indexToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_collateralDelta\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_sizeDelta\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_isLong\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"validateDecreasePosition\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_collateralToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_indexToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_sizeDelta\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_isLong\",\"type\":\"bool\"}],\"name\":\"validateIncreasePosition\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_collateralToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_indexToken\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_isLong\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"_raise\",\"type\":\"bool\"}],\"name\":\"validateLiquidation\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// IVaultUtilsFuncSigs maps the 4-byte function signature to its string representation.
var IVaultUtilsFuncSigs = map[string]string{
	"4adeddc6": "getBuyUsdgFeeBasisPoints(address,uint256)",
	"b1cc53ab": "getEntryFundingRate(address,address,bool)",
	"c7e074c3": "getFeeBasisPoints(address,uint256,uint256,uint256,bool)",
	"da76524c": "getFundingFee(address,address,address,bool,uint256,uint256)",
	"fdaf6ac3": "getPositionFee(address,address,address,bool,uint256)",
	"eb0835bf": "getSellUsdgFeeBasisPoints(address,uint256)",
	"da133816": "getSwapFeeBasisPoints(address,address,uint256)",
	"fbfded6d": "updateCumulativeFundingRate(address,address)",
	"81d11a23": "validateDecreasePosition(address,address,address,uint256,uint256,bool,address)",
	"9d5c28fa": "validateIncreasePosition(address,address,address,uint256,bool)",
	"d54d5a9f": "validateLiquidation(address,address,address,bool,bool)",
}

// IVaultUtils is an auto generated Go binding around an Ethereum contract.
type IVaultUtils struct {
	IVaultUtilsCaller     // Read-only binding to the contract
	IVaultUtilsTransactor // Write-only binding to the contract
	IVaultUtilsFilterer   // Log filterer for contract events
}

// IVaultUtilsCaller is an auto generated read-only Go binding around an Ethereum contract.
type IVaultUtilsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IVaultUtilsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IVaultUtilsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IVaultUtilsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IVaultUtilsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IVaultUtilsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IVaultUtilsSession struct {
	Contract     *IVaultUtils      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IVaultUtilsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IVaultUtilsCallerSession struct {
	Contract *IVaultUtilsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// IVaultUtilsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IVaultUtilsTransactorSession struct {
	Contract     *IVaultUtilsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// IVaultUtilsRaw is an auto generated low-level Go binding around an Ethereum contract.
type IVaultUtilsRaw struct {
	Contract *IVaultUtils // Generic contract binding to access the raw methods on
}

// IVaultUtilsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IVaultUtilsCallerRaw struct {
	Contract *IVaultUtilsCaller // Generic read-only contract binding to access the raw methods on
}

// IVaultUtilsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IVaultUtilsTransactorRaw struct {
	Contract *IVaultUtilsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIVaultUtils creates a new instance of IVaultUtils, bound to a specific deployed contract.
func NewIVaultUtils(address common.Address, backend bind.ContractBackend) (*IVaultUtils, error) {
	contract, err := bindIVaultUtils(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IVaultUtils{IVaultUtilsCaller: IVaultUtilsCaller{contract: contract}, IVaultUtilsTransactor: IVaultUtilsTransactor{contract: contract}, IVaultUtilsFilterer: IVaultUtilsFilterer{contract: contract}}, nil
}

// NewIVaultUtilsCaller creates a new read-only instance of IVaultUtils, bound to a specific deployed contract.
func NewIVaultUtilsCaller(address common.Address, caller bind.ContractCaller) (*IVaultUtilsCaller, error) {
	contract, err := bindIVaultUtils(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IVaultUtilsCaller{contract: contract}, nil
}

// NewIVaultUtilsTransactor creates a new write-only instance of IVaultUtils, bound to a specific deployed contract.
func NewIVaultUtilsTransactor(address common.Address, transactor bind.ContractTransactor) (*IVaultUtilsTransactor, error) {
	contract, err := bindIVaultUtils(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IVaultUtilsTransactor{contract: contract}, nil
}

// NewIVaultUtilsFilterer creates a new log filterer instance of IVaultUtils, bound to a specific deployed contract.
func NewIVaultUtilsFilterer(address common.Address, filterer bind.ContractFilterer) (*IVaultUtilsFilterer, error) {
	contract, err := bindIVaultUtils(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IVaultUtilsFilterer{contract: contract}, nil
}

// bindIVaultUtils binds a generic wrapper to an already deployed contract.
func bindIVaultUtils(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IVaultUtilsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IVaultUtils *IVaultUtilsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IVaultUtils.Contract.IVaultUtilsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IVaultUtils *IVaultUtilsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IVaultUtils.Contract.IVaultUtilsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IVaultUtils *IVaultUtilsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IVaultUtils.Contract.IVaultUtilsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IVaultUtils *IVaultUtilsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IVaultUtils.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IVaultUtils *IVaultUtilsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IVaultUtils.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IVaultUtils *IVaultUtilsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IVaultUtils.Contract.contract.Transact(opts, method, params...)
}

// GetBuyUsdgFeeBasisPoints is a free data retrieval call binding the contract method 0x4adeddc6.
//
// Solidity: function getBuyUsdgFeeBasisPoints(address _token, uint256 _usdgAmount) view returns(uint256)
func (_IVaultUtils *IVaultUtilsCaller) GetBuyUsdgFeeBasisPoints(opts *bind.CallOpts, _token common.Address, _usdgAmount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IVaultUtils.contract.Call(opts, &out, "getBuyUsdgFeeBasisPoints", _token, _usdgAmount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBuyUsdgFeeBasisPoints is a free data retrieval call binding the contract method 0x4adeddc6.
//
// Solidity: function getBuyUsdgFeeBasisPoints(address _token, uint256 _usdgAmount) view returns(uint256)
func (_IVaultUtils *IVaultUtilsSession) GetBuyUsdgFeeBasisPoints(_token common.Address, _usdgAmount *big.Int) (*big.Int, error) {
	return _IVaultUtils.Contract.GetBuyUsdgFeeBasisPoints(&_IVaultUtils.CallOpts, _token, _usdgAmount)
}

// GetBuyUsdgFeeBasisPoints is a free data retrieval call binding the contract method 0x4adeddc6.
//
// Solidity: function getBuyUsdgFeeBasisPoints(address _token, uint256 _usdgAmount) view returns(uint256)
func (_IVaultUtils *IVaultUtilsCallerSession) GetBuyUsdgFeeBasisPoints(_token common.Address, _usdgAmount *big.Int) (*big.Int, error) {
	return _IVaultUtils.Contract.GetBuyUsdgFeeBasisPoints(&_IVaultUtils.CallOpts, _token, _usdgAmount)
}

// GetEntryFundingRate is a free data retrieval call binding the contract method 0xb1cc53ab.
//
// Solidity: function getEntryFundingRate(address _collateralToken, address _indexToken, bool _isLong) view returns(uint256)
func (_IVaultUtils *IVaultUtilsCaller) GetEntryFundingRate(opts *bind.CallOpts, _collateralToken common.Address, _indexToken common.Address, _isLong bool) (*big.Int, error) {
	var out []interface{}
	err := _IVaultUtils.contract.Call(opts, &out, "getEntryFundingRate", _collateralToken, _indexToken, _isLong)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetEntryFundingRate is a free data retrieval call binding the contract method 0xb1cc53ab.
//
// Solidity: function getEntryFundingRate(address _collateralToken, address _indexToken, bool _isLong) view returns(uint256)
func (_IVaultUtils *IVaultUtilsSession) GetEntryFundingRate(_collateralToken common.Address, _indexToken common.Address, _isLong bool) (*big.Int, error) {
	return _IVaultUtils.Contract.GetEntryFundingRate(&_IVaultUtils.CallOpts, _collateralToken, _indexToken, _isLong)
}

// GetEntryFundingRate is a free data retrieval call binding the contract method 0xb1cc53ab.
//
// Solidity: function getEntryFundingRate(address _collateralToken, address _indexToken, bool _isLong) view returns(uint256)
func (_IVaultUtils *IVaultUtilsCallerSession) GetEntryFundingRate(_collateralToken common.Address, _indexToken common.Address, _isLong bool) (*big.Int, error) {
	return _IVaultUtils.Contract.GetEntryFundingRate(&_IVaultUtils.CallOpts, _collateralToken, _indexToken, _isLong)
}

// GetFeeBasisPoints is a free data retrieval call binding the contract method 0xc7e074c3.
//
// Solidity: function getFeeBasisPoints(address _token, uint256 _usdgDelta, uint256 _feeBasisPoints, uint256 _taxBasisPoints, bool _increment) view returns(uint256)
func (_IVaultUtils *IVaultUtilsCaller) GetFeeBasisPoints(opts *bind.CallOpts, _token common.Address, _usdgDelta *big.Int, _feeBasisPoints *big.Int, _taxBasisPoints *big.Int, _increment bool) (*big.Int, error) {
	var out []interface{}
	err := _IVaultUtils.contract.Call(opts, &out, "getFeeBasisPoints", _token, _usdgDelta, _feeBasisPoints, _taxBasisPoints, _increment)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetFeeBasisPoints is a free data retrieval call binding the contract method 0xc7e074c3.
//
// Solidity: function getFeeBasisPoints(address _token, uint256 _usdgDelta, uint256 _feeBasisPoints, uint256 _taxBasisPoints, bool _increment) view returns(uint256)
func (_IVaultUtils *IVaultUtilsSession) GetFeeBasisPoints(_token common.Address, _usdgDelta *big.Int, _feeBasisPoints *big.Int, _taxBasisPoints *big.Int, _increment bool) (*big.Int, error) {
	return _IVaultUtils.Contract.GetFeeBasisPoints(&_IVaultUtils.CallOpts, _token, _usdgDelta, _feeBasisPoints, _taxBasisPoints, _increment)
}

// GetFeeBasisPoints is a free data retrieval call binding the contract method 0xc7e074c3.
//
// Solidity: function getFeeBasisPoints(address _token, uint256 _usdgDelta, uint256 _feeBasisPoints, uint256 _taxBasisPoints, bool _increment) view returns(uint256)
func (_IVaultUtils *IVaultUtilsCallerSession) GetFeeBasisPoints(_token common.Address, _usdgDelta *big.Int, _feeBasisPoints *big.Int, _taxBasisPoints *big.Int, _increment bool) (*big.Int, error) {
	return _IVaultUtils.Contract.GetFeeBasisPoints(&_IVaultUtils.CallOpts, _token, _usdgDelta, _feeBasisPoints, _taxBasisPoints, _increment)
}

// GetFundingFee is a free data retrieval call binding the contract method 0xda76524c.
//
// Solidity: function getFundingFee(address _account, address _collateralToken, address _indexToken, bool _isLong, uint256 _size, uint256 _entryFundingRate) view returns(uint256)
func (_IVaultUtils *IVaultUtilsCaller) GetFundingFee(opts *bind.CallOpts, _account common.Address, _collateralToken common.Address, _indexToken common.Address, _isLong bool, _size *big.Int, _entryFundingRate *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IVaultUtils.contract.Call(opts, &out, "getFundingFee", _account, _collateralToken, _indexToken, _isLong, _size, _entryFundingRate)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetFundingFee is a free data retrieval call binding the contract method 0xda76524c.
//
// Solidity: function getFundingFee(address _account, address _collateralToken, address _indexToken, bool _isLong, uint256 _size, uint256 _entryFundingRate) view returns(uint256)
func (_IVaultUtils *IVaultUtilsSession) GetFundingFee(_account common.Address, _collateralToken common.Address, _indexToken common.Address, _isLong bool, _size *big.Int, _entryFundingRate *big.Int) (*big.Int, error) {
	return _IVaultUtils.Contract.GetFundingFee(&_IVaultUtils.CallOpts, _account, _collateralToken, _indexToken, _isLong, _size, _entryFundingRate)
}

// GetFundingFee is a free data retrieval call binding the contract method 0xda76524c.
//
// Solidity: function getFundingFee(address _account, address _collateralToken, address _indexToken, bool _isLong, uint256 _size, uint256 _entryFundingRate) view returns(uint256)
func (_IVaultUtils *IVaultUtilsCallerSession) GetFundingFee(_account common.Address, _collateralToken common.Address, _indexToken common.Address, _isLong bool, _size *big.Int, _entryFundingRate *big.Int) (*big.Int, error) {
	return _IVaultUtils.Contract.GetFundingFee(&_IVaultUtils.CallOpts, _account, _collateralToken, _indexToken, _isLong, _size, _entryFundingRate)
}

// GetPositionFee is a free data retrieval call binding the contract method 0xfdaf6ac3.
//
// Solidity: function getPositionFee(address _account, address _collateralToken, address _indexToken, bool _isLong, uint256 _sizeDelta) view returns(uint256)
func (_IVaultUtils *IVaultUtilsCaller) GetPositionFee(opts *bind.CallOpts, _account common.Address, _collateralToken common.Address, _indexToken common.Address, _isLong bool, _sizeDelta *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IVaultUtils.contract.Call(opts, &out, "getPositionFee", _account, _collateralToken, _indexToken, _isLong, _sizeDelta)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPositionFee is a free data retrieval call binding the contract method 0xfdaf6ac3.
//
// Solidity: function getPositionFee(address _account, address _collateralToken, address _indexToken, bool _isLong, uint256 _sizeDelta) view returns(uint256)
func (_IVaultUtils *IVaultUtilsSession) GetPositionFee(_account common.Address, _collateralToken common.Address, _indexToken common.Address, _isLong bool, _sizeDelta *big.Int) (*big.Int, error) {
	return _IVaultUtils.Contract.GetPositionFee(&_IVaultUtils.CallOpts, _account, _collateralToken, _indexToken, _isLong, _sizeDelta)
}

// GetPositionFee is a free data retrieval call binding the contract method 0xfdaf6ac3.
//
// Solidity: function getPositionFee(address _account, address _collateralToken, address _indexToken, bool _isLong, uint256 _sizeDelta) view returns(uint256)
func (_IVaultUtils *IVaultUtilsCallerSession) GetPositionFee(_account common.Address, _collateralToken common.Address, _indexToken common.Address, _isLong bool, _sizeDelta *big.Int) (*big.Int, error) {
	return _IVaultUtils.Contract.GetPositionFee(&_IVaultUtils.CallOpts, _account, _collateralToken, _indexToken, _isLong, _sizeDelta)
}

// GetSellUsdgFeeBasisPoints is a free data retrieval call binding the contract method 0xeb0835bf.
//
// Solidity: function getSellUsdgFeeBasisPoints(address _token, uint256 _usdgAmount) view returns(uint256)
func (_IVaultUtils *IVaultUtilsCaller) GetSellUsdgFeeBasisPoints(opts *bind.CallOpts, _token common.Address, _usdgAmount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IVaultUtils.contract.Call(opts, &out, "getSellUsdgFeeBasisPoints", _token, _usdgAmount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSellUsdgFeeBasisPoints is a free data retrieval call binding the contract method 0xeb0835bf.
//
// Solidity: function getSellUsdgFeeBasisPoints(address _token, uint256 _usdgAmount) view returns(uint256)
func (_IVaultUtils *IVaultUtilsSession) GetSellUsdgFeeBasisPoints(_token common.Address, _usdgAmount *big.Int) (*big.Int, error) {
	return _IVaultUtils.Contract.GetSellUsdgFeeBasisPoints(&_IVaultUtils.CallOpts, _token, _usdgAmount)
}

// GetSellUsdgFeeBasisPoints is a free data retrieval call binding the contract method 0xeb0835bf.
//
// Solidity: function getSellUsdgFeeBasisPoints(address _token, uint256 _usdgAmount) view returns(uint256)
func (_IVaultUtils *IVaultUtilsCallerSession) GetSellUsdgFeeBasisPoints(_token common.Address, _usdgAmount *big.Int) (*big.Int, error) {
	return _IVaultUtils.Contract.GetSellUsdgFeeBasisPoints(&_IVaultUtils.CallOpts, _token, _usdgAmount)
}

// GetSwapFeeBasisPoints is a free data retrieval call binding the contract method 0xda133816.
//
// Solidity: function getSwapFeeBasisPoints(address _tokenIn, address _tokenOut, uint256 _usdgAmount) view returns(uint256)
func (_IVaultUtils *IVaultUtilsCaller) GetSwapFeeBasisPoints(opts *bind.CallOpts, _tokenIn common.Address, _tokenOut common.Address, _usdgAmount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IVaultUtils.contract.Call(opts, &out, "getSwapFeeBasisPoints", _tokenIn, _tokenOut, _usdgAmount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSwapFeeBasisPoints is a free data retrieval call binding the contract method 0xda133816.
//
// Solidity: function getSwapFeeBasisPoints(address _tokenIn, address _tokenOut, uint256 _usdgAmount) view returns(uint256)
func (_IVaultUtils *IVaultUtilsSession) GetSwapFeeBasisPoints(_tokenIn common.Address, _tokenOut common.Address, _usdgAmount *big.Int) (*big.Int, error) {
	return _IVaultUtils.Contract.GetSwapFeeBasisPoints(&_IVaultUtils.CallOpts, _tokenIn, _tokenOut, _usdgAmount)
}

// GetSwapFeeBasisPoints is a free data retrieval call binding the contract method 0xda133816.
//
// Solidity: function getSwapFeeBasisPoints(address _tokenIn, address _tokenOut, uint256 _usdgAmount) view returns(uint256)
func (_IVaultUtils *IVaultUtilsCallerSession) GetSwapFeeBasisPoints(_tokenIn common.Address, _tokenOut common.Address, _usdgAmount *big.Int) (*big.Int, error) {
	return _IVaultUtils.Contract.GetSwapFeeBasisPoints(&_IVaultUtils.CallOpts, _tokenIn, _tokenOut, _usdgAmount)
}

// ValidateDecreasePosition is a free data retrieval call binding the contract method 0x81d11a23.
//
// Solidity: function validateDecreasePosition(address _account, address _collateralToken, address _indexToken, uint256 _collateralDelta, uint256 _sizeDelta, bool _isLong, address _receiver) view returns()
func (_IVaultUtils *IVaultUtilsCaller) ValidateDecreasePosition(opts *bind.CallOpts, _account common.Address, _collateralToken common.Address, _indexToken common.Address, _collateralDelta *big.Int, _sizeDelta *big.Int, _isLong bool, _receiver common.Address) error {
	var out []interface{}
	err := _IVaultUtils.contract.Call(opts, &out, "validateDecreasePosition", _account, _collateralToken, _indexToken, _collateralDelta, _sizeDelta, _isLong, _receiver)

	if err != nil {
		return err
	}

	return err

}

// ValidateDecreasePosition is a free data retrieval call binding the contract method 0x81d11a23.
//
// Solidity: function validateDecreasePosition(address _account, address _collateralToken, address _indexToken, uint256 _collateralDelta, uint256 _sizeDelta, bool _isLong, address _receiver) view returns()
func (_IVaultUtils *IVaultUtilsSession) ValidateDecreasePosition(_account common.Address, _collateralToken common.Address, _indexToken common.Address, _collateralDelta *big.Int, _sizeDelta *big.Int, _isLong bool, _receiver common.Address) error {
	return _IVaultUtils.Contract.ValidateDecreasePosition(&_IVaultUtils.CallOpts, _account, _collateralToken, _indexToken, _collateralDelta, _sizeDelta, _isLong, _receiver)
}

// ValidateDecreasePosition is a free data retrieval call binding the contract method 0x81d11a23.
//
// Solidity: function validateDecreasePosition(address _account, address _collateralToken, address _indexToken, uint256 _collateralDelta, uint256 _sizeDelta, bool _isLong, address _receiver) view returns()
func (_IVaultUtils *IVaultUtilsCallerSession) ValidateDecreasePosition(_account common.Address, _collateralToken common.Address, _indexToken common.Address, _collateralDelta *big.Int, _sizeDelta *big.Int, _isLong bool, _receiver common.Address) error {
	return _IVaultUtils.Contract.ValidateDecreasePosition(&_IVaultUtils.CallOpts, _account, _collateralToken, _indexToken, _collateralDelta, _sizeDelta, _isLong, _receiver)
}

// ValidateIncreasePosition is a free data retrieval call binding the contract method 0x9d5c28fa.
//
// Solidity: function validateIncreasePosition(address _account, address _collateralToken, address _indexToken, uint256 _sizeDelta, bool _isLong) view returns()
func (_IVaultUtils *IVaultUtilsCaller) ValidateIncreasePosition(opts *bind.CallOpts, _account common.Address, _collateralToken common.Address, _indexToken common.Address, _sizeDelta *big.Int, _isLong bool) error {
	var out []interface{}
	err := _IVaultUtils.contract.Call(opts, &out, "validateIncreasePosition", _account, _collateralToken, _indexToken, _sizeDelta, _isLong)

	if err != nil {
		return err
	}

	return err

}

// ValidateIncreasePosition is a free data retrieval call binding the contract method 0x9d5c28fa.
//
// Solidity: function validateIncreasePosition(address _account, address _collateralToken, address _indexToken, uint256 _sizeDelta, bool _isLong) view returns()
func (_IVaultUtils *IVaultUtilsSession) ValidateIncreasePosition(_account common.Address, _collateralToken common.Address, _indexToken common.Address, _sizeDelta *big.Int, _isLong bool) error {
	return _IVaultUtils.Contract.ValidateIncreasePosition(&_IVaultUtils.CallOpts, _account, _collateralToken, _indexToken, _sizeDelta, _isLong)
}

// ValidateIncreasePosition is a free data retrieval call binding the contract method 0x9d5c28fa.
//
// Solidity: function validateIncreasePosition(address _account, address _collateralToken, address _indexToken, uint256 _sizeDelta, bool _isLong) view returns()
func (_IVaultUtils *IVaultUtilsCallerSession) ValidateIncreasePosition(_account common.Address, _collateralToken common.Address, _indexToken common.Address, _sizeDelta *big.Int, _isLong bool) error {
	return _IVaultUtils.Contract.ValidateIncreasePosition(&_IVaultUtils.CallOpts, _account, _collateralToken, _indexToken, _sizeDelta, _isLong)
}

// ValidateLiquidation is a free data retrieval call binding the contract method 0xd54d5a9f.
//
// Solidity: function validateLiquidation(address _account, address _collateralToken, address _indexToken, bool _isLong, bool _raise) view returns(uint256, uint256)
func (_IVaultUtils *IVaultUtilsCaller) ValidateLiquidation(opts *bind.CallOpts, _account common.Address, _collateralToken common.Address, _indexToken common.Address, _isLong bool, _raise bool) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _IVaultUtils.contract.Call(opts, &out, "validateLiquidation", _account, _collateralToken, _indexToken, _isLong, _raise)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// ValidateLiquidation is a free data retrieval call binding the contract method 0xd54d5a9f.
//
// Solidity: function validateLiquidation(address _account, address _collateralToken, address _indexToken, bool _isLong, bool _raise) view returns(uint256, uint256)
func (_IVaultUtils *IVaultUtilsSession) ValidateLiquidation(_account common.Address, _collateralToken common.Address, _indexToken common.Address, _isLong bool, _raise bool) (*big.Int, *big.Int, error) {
	return _IVaultUtils.Contract.ValidateLiquidation(&_IVaultUtils.CallOpts, _account, _collateralToken, _indexToken, _isLong, _raise)
}

// ValidateLiquidation is a free data retrieval call binding the contract method 0xd54d5a9f.
//
// Solidity: function validateLiquidation(address _account, address _collateralToken, address _indexToken, bool _isLong, bool _raise) view returns(uint256, uint256)
func (_IVaultUtils *IVaultUtilsCallerSession) ValidateLiquidation(_account common.Address, _collateralToken common.Address, _indexToken common.Address, _isLong bool, _raise bool) (*big.Int, *big.Int, error) {
	return _IVaultUtils.Contract.ValidateLiquidation(&_IVaultUtils.CallOpts, _account, _collateralToken, _indexToken, _isLong, _raise)
}

// UpdateCumulativeFundingRate is a paid mutator transaction binding the contract method 0xfbfded6d.
//
// Solidity: function updateCumulativeFundingRate(address _collateralToken, address _indexToken) returns(bool)
func (_IVaultUtils *IVaultUtilsTransactor) UpdateCumulativeFundingRate(opts *bind.TransactOpts, _collateralToken common.Address, _indexToken common.Address) (*types.Transaction, error) {
	return _IVaultUtils.contract.Transact(opts, "updateCumulativeFundingRate", _collateralToken, _indexToken)
}

// UpdateCumulativeFundingRate is a paid mutator transaction binding the contract method 0xfbfded6d.
//
// Solidity: function updateCumulativeFundingRate(address _collateralToken, address _indexToken) returns(bool)
func (_IVaultUtils *IVaultUtilsSession) UpdateCumulativeFundingRate(_collateralToken common.Address, _indexToken common.Address) (*types.Transaction, error) {
	return _IVaultUtils.Contract.UpdateCumulativeFundingRate(&_IVaultUtils.TransactOpts, _collateralToken, _indexToken)
}

// UpdateCumulativeFundingRate is a paid mutator transaction binding the contract method 0xfbfded6d.
//
// Solidity: function updateCumulativeFundingRate(address _collateralToken, address _indexToken) returns(bool)
func (_IVaultUtils *IVaultUtilsTransactorSession) UpdateCumulativeFundingRate(_collateralToken common.Address, _indexToken common.Address) (*types.Transaction, error) {
	return _IVaultUtils.Contract.UpdateCumulativeFundingRate(&_IVaultUtils.TransactOpts, _collateralToken, _indexToken)
}

var ()

// ReentrancyGuardABI is the input ABI used to generate the binding from.
const ReentrancyGuardABI = "[]"

// ReentrancyGuard is an auto generated Go binding around an Ethereum contract.
type ReentrancyGuard struct {
	ReentrancyGuardCaller     // Read-only binding to the contract
	ReentrancyGuardTransactor // Write-only binding to the contract
	ReentrancyGuardFilterer   // Log filterer for contract events
}

// ReentrancyGuardCaller is an auto generated read-only Go binding around an Ethereum contract.
type ReentrancyGuardCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReentrancyGuardTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ReentrancyGuardTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReentrancyGuardFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ReentrancyGuardFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReentrancyGuardSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ReentrancyGuardSession struct {
	Contract     *ReentrancyGuard  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ReentrancyGuardCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ReentrancyGuardCallerSession struct {
	Contract *ReentrancyGuardCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// ReentrancyGuardTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ReentrancyGuardTransactorSession struct {
	Contract     *ReentrancyGuardTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// ReentrancyGuardRaw is an auto generated low-level Go binding around an Ethereum contract.
type ReentrancyGuardRaw struct {
	Contract *ReentrancyGuard // Generic contract binding to access the raw methods on
}

// ReentrancyGuardCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ReentrancyGuardCallerRaw struct {
	Contract *ReentrancyGuardCaller // Generic read-only contract binding to access the raw methods on
}

// ReentrancyGuardTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ReentrancyGuardTransactorRaw struct {
	Contract *ReentrancyGuardTransactor // Generic write-only contract binding to access the raw methods on
}

// NewReentrancyGuard creates a new instance of ReentrancyGuard, bound to a specific deployed contract.
func NewReentrancyGuard(address common.Address, backend bind.ContractBackend) (*ReentrancyGuard, error) {
	contract, err := bindReentrancyGuard(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ReentrancyGuard{ReentrancyGuardCaller: ReentrancyGuardCaller{contract: contract}, ReentrancyGuardTransactor: ReentrancyGuardTransactor{contract: contract}, ReentrancyGuardFilterer: ReentrancyGuardFilterer{contract: contract}}, nil
}

// NewReentrancyGuardCaller creates a new read-only instance of ReentrancyGuard, bound to a specific deployed contract.
func NewReentrancyGuardCaller(address common.Address, caller bind.ContractCaller) (*ReentrancyGuardCaller, error) {
	contract, err := bindReentrancyGuard(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ReentrancyGuardCaller{contract: contract}, nil
}

// NewReentrancyGuardTransactor creates a new write-only instance of ReentrancyGuard, bound to a specific deployed contract.
func NewReentrancyGuardTransactor(address common.Address, transactor bind.ContractTransactor) (*ReentrancyGuardTransactor, error) {
	contract, err := bindReentrancyGuard(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ReentrancyGuardTransactor{contract: contract}, nil
}

// NewReentrancyGuardFilterer creates a new log filterer instance of ReentrancyGuard, bound to a specific deployed contract.
func NewReentrancyGuardFilterer(address common.Address, filterer bind.ContractFilterer) (*ReentrancyGuardFilterer, error) {
	contract, err := bindReentrancyGuard(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ReentrancyGuardFilterer{contract: contract}, nil
}

// bindReentrancyGuard binds a generic wrapper to an already deployed contract.
func bindReentrancyGuard(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ReentrancyGuardABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ReentrancyGuard *ReentrancyGuardRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ReentrancyGuard.Contract.ReentrancyGuardCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ReentrancyGuard *ReentrancyGuardRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReentrancyGuard.Contract.ReentrancyGuardTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ReentrancyGuard *ReentrancyGuardRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ReentrancyGuard.Contract.ReentrancyGuardTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ReentrancyGuard *ReentrancyGuardCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ReentrancyGuard.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ReentrancyGuard *ReentrancyGuardTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReentrancyGuard.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ReentrancyGuard *ReentrancyGuardTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ReentrancyGuard.Contract.contract.Transact(opts, method, params...)
}

var ()

// SafeERC20ABI is the input ABI used to generate the binding from.
const SafeERC20ABI = "[]"

// SafeERC20Bin is the compiled bytecode used for deploying new contracts.
var SafeERC20Bin = "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220aa6b8b5dc6bd700359779ba35030d78146d4314951b1c9344a63aa4e7f86fc0164736f6c634300060c0033"

// DeploySafeERC20 deploys a new Ethereum contract, binding an instance of SafeERC20 to it.
func DeploySafeERC20(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeERC20, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeERC20ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SafeERC20Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SafeERC20{SafeERC20Caller: SafeERC20Caller{contract: contract}, SafeERC20Transactor: SafeERC20Transactor{contract: contract}, SafeERC20Filterer: SafeERC20Filterer{contract: contract}}, nil
}

// SafeERC20 is an auto generated Go binding around an Ethereum contract.
type SafeERC20 struct {
	SafeERC20Caller     // Read-only binding to the contract
	SafeERC20Transactor // Write-only binding to the contract
	SafeERC20Filterer   // Log filterer for contract events
}

// SafeERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type SafeERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeERC20Session struct {
	Contract     *SafeERC20        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeERC20CallerSession struct {
	Contract *SafeERC20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// SafeERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeERC20TransactorSession struct {
	Contract     *SafeERC20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// SafeERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type SafeERC20Raw struct {
	Contract *SafeERC20 // Generic contract binding to access the raw methods on
}

// SafeERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeERC20CallerRaw struct {
	Contract *SafeERC20Caller // Generic read-only contract binding to access the raw methods on
}

// SafeERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeERC20TransactorRaw struct {
	Contract *SafeERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeERC20 creates a new instance of SafeERC20, bound to a specific deployed contract.
func NewSafeERC20(address common.Address, backend bind.ContractBackend) (*SafeERC20, error) {
	contract, err := bindSafeERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeERC20{SafeERC20Caller: SafeERC20Caller{contract: contract}, SafeERC20Transactor: SafeERC20Transactor{contract: contract}, SafeERC20Filterer: SafeERC20Filterer{contract: contract}}, nil
}

// NewSafeERC20Caller creates a new read-only instance of SafeERC20, bound to a specific deployed contract.
func NewSafeERC20Caller(address common.Address, caller bind.ContractCaller) (*SafeERC20Caller, error) {
	contract, err := bindSafeERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeERC20Caller{contract: contract}, nil
}

// NewSafeERC20Transactor creates a new write-only instance of SafeERC20, bound to a specific deployed contract.
func NewSafeERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*SafeERC20Transactor, error) {
	contract, err := bindSafeERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeERC20Transactor{contract: contract}, nil
}

// NewSafeERC20Filterer creates a new log filterer instance of SafeERC20, bound to a specific deployed contract.
func NewSafeERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*SafeERC20Filterer, error) {
	contract, err := bindSafeERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeERC20Filterer{contract: contract}, nil
}

// bindSafeERC20 binds a generic wrapper to an already deployed contract.
func bindSafeERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeERC20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeERC20 *SafeERC20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeERC20.Contract.SafeERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeERC20 *SafeERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeERC20.Contract.SafeERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeERC20 *SafeERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeERC20.Contract.SafeERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeERC20 *SafeERC20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeERC20 *SafeERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeERC20 *SafeERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeERC20.Contract.contract.Transact(opts, method, params...)
}

var ()

// SafeMathABI is the input ABI used to generate the binding from.
const SafeMathABI = "[]"

// SafeMathBin is the compiled bytecode used for deploying new contracts.
var SafeMathBin = "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220ed64cb4fa4bf062de13fd2c58dd55da6653c3d2ed14f82df59edf24268a8e2a064736f6c634300060c0033"

// DeploySafeMath deploys a new Ethereum contract, binding an instance of SafeMath to it.
func DeploySafeMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeMath, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SafeMathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// SafeMath is an auto generated Go binding around an Ethereum contract.
type SafeMath struct {
	SafeMathCaller     // Read-only binding to the contract
	SafeMathTransactor // Write-only binding to the contract
	SafeMathFilterer   // Log filterer for contract events
}

// SafeMathCaller is an auto generated read-only Go binding around an Ethereum contract.
type SafeMathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeMathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeMathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeMathSession struct {
	Contract     *SafeMath         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeMathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeMathCallerSession struct {
	Contract *SafeMathCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// SafeMathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeMathTransactorSession struct {
	Contract     *SafeMathTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SafeMathRaw is an auto generated low-level Go binding around an Ethereum contract.
type SafeMathRaw struct {
	Contract *SafeMath // Generic contract binding to access the raw methods on
}

// SafeMathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeMathCallerRaw struct {
	Contract *SafeMathCaller // Generic read-only contract binding to access the raw methods on
}

// SafeMathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeMathTransactorRaw struct {
	Contract *SafeMathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeMath creates a new instance of SafeMath, bound to a specific deployed contract.
func NewSafeMath(address common.Address, backend bind.ContractBackend) (*SafeMath, error) {
	contract, err := bindSafeMath(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// NewSafeMathCaller creates a new read-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathCaller(address common.Address, caller bind.ContractCaller) (*SafeMathCaller, error) {
	contract, err := bindSafeMath(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathCaller{contract: contract}, nil
}

// NewSafeMathTransactor creates a new write-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathTransactor(address common.Address, transactor bind.ContractTransactor) (*SafeMathTransactor, error) {
	contract, err := bindSafeMath(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathTransactor{contract: contract}, nil
}

// NewSafeMathFilterer creates a new log filterer instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathFilterer(address common.Address, filterer bind.ContractFilterer) (*SafeMathFilterer, error) {
	contract, err := bindSafeMath(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeMathFilterer{contract: contract}, nil
}

// bindSafeMath binds a generic wrapper to an already deployed contract.
func bindSafeMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.SafeMathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transact(opts, method, params...)
}

var (
	MethodAddRouter = "addRouter"

	MethodClearTokenConfig = "clearTokenConfig"

	MethodInitialize = "initialize"

	MethodRemoveRouter = "removeRouter"

	MethodSetErrorController = "setErrorController"

	MethodSetGov = "setGov"

	MethodUpgradeVault = "upgradeVault"

	MethodBASISPOINTSDIVISOR = "BASIS_POINTS_DIVISOR"

	MethodFUNDINGRATEPRECISION = "FUNDING_RATE_PRECISION"

	MethodMAXFEEBASISPOINTS = "MAX_FEE_BASIS_POINTS"

	MethodMAXFUNDINGRATEFACTOR = "MAX_FUNDING_RATE_FACTOR"

	MethodMAXLIQUIDATIONFEEUSD = "MAX_LIQUIDATION_FEE_USD"

	MethodMINFUNDINGRATEINTERVAL = "MIN_FUNDING_RATE_INTERVAL"

	MethodMINLEVERAGE = "MIN_LEVERAGE"

	MethodPRICEPRECISION = "PRICE_PRECISION"

	MethodUSDGDECIMALS = "USDG_DECIMALS"

	MethodAdjustForDecimals = "adjustForDecimals"

	MethodErrorController = "errorController"

	MethodErrors = "errors"

	MethodGetGlobalShortDelta = "getGlobalShortDelta"

	MethodGetNextAveragePrice = "getNextAveragePrice"

	MethodGetNextGlobalShortAveragePrice = "getNextGlobalShortAveragePrice"

	MethodGetPositionDelta = "getPositionDelta"

	MethodGetPositionKey = "getPositionKey"

	MethodGetPositionLeverage = "getPositionLeverage"

	MethodGetRedemptionCollateral = "getRedemptionCollateral"

	MethodGetRedemptionCollateralUsd = "getRedemptionCollateralUsd"

	MethodGetUtilisation = "getUtilisation"

	MethodIncludeAmmPrice = "includeAmmPrice"

	MethodPositions = "positions"

	MethodUsdToToken = "usdToToken"

	MethodUsdToTokenMax = "usdToTokenMax"

	MethodUsdToTokenMin = "usdToTokenMin"

	MethodUseSwapPricing = "useSwapPricing"

	MethodVaultUtils = "vaultUtils"

	EventBuyUSDG = "BuyUSDG"

	EventClosePosition = "ClosePosition"

	EventCollectMarginFees = "CollectMarginFees"

	EventCollectSwapFees = "CollectSwapFees"

	EventDecreaseGuaranteedUsd = "DecreaseGuaranteedUsd"

	EventDecreasePoolAmount = "DecreasePoolAmount"

	EventDecreasePosition = "DecreasePosition"

	EventDecreaseReservedAmount = "DecreaseReservedAmount"

	EventDecreaseUsdgAmount = "DecreaseUsdgAmount"

	EventDirectPoolDeposit = "DirectPoolDeposit"

	EventIncreaseGuaranteedUsd = "IncreaseGuaranteedUsd"

	EventIncreasePoolAmount = "IncreasePoolAmount"

	EventIncreasePosition = "IncreasePosition"

	EventIncreaseReservedAmount = "IncreaseReservedAmount"

	EventIncreaseUsdgAmount = "IncreaseUsdgAmount"

	EventLiquidatePosition = "LiquidatePosition"

	EventSellUSDG = "SellUSDG"

	EventSwap = "Swap"

	EventUpdateFundingRate = "UpdateFundingRate"

	EventUpdatePnl = "UpdatePnl"

	EventUpdatePosition = "UpdatePosition"
)

// VaultABI is the input ABI used to generate the binding from.
const VaultABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"usdgAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeBasisPoints\",\"type\":\"uint256\"}],\"name\":\"BuyUSDG\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"size\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"collateral\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"averagePrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"entryFundingRate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reserveAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"realisedPnl\",\"type\":\"int256\"}],\"name\":\"ClosePosition\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeUsd\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeTokens\",\"type\":\"uint256\"}],\"name\":\"CollectMarginFees\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeUsd\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeTokens\",\"type\":\"uint256\"}],\"name\":\"CollectSwapFees\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"DecreaseGuaranteedUsd\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"DecreasePoolAmount\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"collateralToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"indexToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"collateralDelta\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sizeDelta\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isLong\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"DecreasePosition\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"DecreaseReservedAmount\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"DecreaseUsdgAmount\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"DirectPoolDeposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"IncreaseGuaranteedUsd\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"IncreasePoolAmount\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"collateralToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"indexToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"collateralDelta\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sizeDelta\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isLong\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"IncreasePosition\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"IncreaseReservedAmount\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"IncreaseUsdgAmount\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"collateralToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"indexToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isLong\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"size\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"collateral\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reserveAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"realisedPnl\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"markPrice\",\"type\":\"uint256\"}],\"name\":\"LiquidatePosition\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"usdgAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeBasisPoints\",\"type\":\"uint256\"}],\"name\":\"SellUSDG\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountOutAfterFees\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeBasisPoints\",\"type\":\"uint256\"}],\"name\":\"Swap\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fundingRate\",\"type\":\"uint256\"}],\"name\":\"UpdateFundingRate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"hasProfit\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"delta\",\"type\":\"uint256\"}],\"name\":\"UpdatePnl\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"size\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"collateral\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"averagePrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"entryFundingRate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reserveAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"realisedPnl\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"markPrice\",\"type\":\"uint256\"}],\"name\":\"UpdatePosition\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BASIS_POINTS_DIVISOR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"FUNDING_RATE_PRECISION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_FEE_BASIS_POINTS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_FUNDING_RATE_FACTOR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_LIQUIDATION_FEE_USD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MIN_FUNDING_RATE_INTERVAL\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MIN_LEVERAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PRICE_PRECISION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"USDG_DECIMALS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_router\",\"type\":\"address\"}],\"name\":\"addRouter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_tokenDiv\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_tokenMul\",\"type\":\"address\"}],\"name\":\"adjustForDecimals\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"allWhitelistedTokens\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allWhitelistedTokensLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"approvedRouters\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"bufferAmounts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"buyUSDG\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"clearTokenConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"cumulativeFundingRates\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_collateralToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_indexToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_collateralDelta\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_sizeDelta\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_isLong\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"decreasePosition\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"directPoolDeposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"errorController\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"errors\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"feeReserves\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fundingInterval\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fundingRateFactor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_indexToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_size\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_averagePrice\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_isLong\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"_lastIncreasedTime\",\"type\":\"uint256\"}],\"name\":\"getDelta\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_collateralToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_indexToken\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_isLong\",\"type\":\"bool\"}],\"name\":\"getEntryFundingRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_usdgDelta\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_feeBasisPoints\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_taxBasisPoints\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_increment\",\"type\":\"bool\"}],\"name\":\"getFeeBasisPoints\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_collateralToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_indexToken\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_isLong\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"_size\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_entryFundingRate\",\"type\":\"uint256\"}],\"name\":\"getFundingFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"getGlobalShortDelta\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"getMaxPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"getMinPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_indexToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_size\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_averagePrice\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_isLong\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"_nextPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_sizeDelta\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_lastIncreasedTime\",\"type\":\"uint256\"}],\"name\":\"getNextAveragePrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"getNextFundingRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_indexToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_nextPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_sizeDelta\",\"type\":\"uint256\"}],\"name\":\"getNextGlobalShortAveragePrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_collateralToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_indexToken\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_isLong\",\"type\":\"bool\"}],\"name\":\"getPosition\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_collateralToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_indexToken\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_isLong\",\"type\":\"bool\"}],\"name\":\"getPositionDelta\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_collateralToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_indexToken\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_isLong\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"_sizeDelta\",\"type\":\"uint256\"}],\"name\":\"getPositionFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_collateralToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_indexToken\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_isLong\",\"type\":\"bool\"}],\"name\":\"getPositionKey\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_collateralToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_indexToken\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_isLong\",\"type\":\"bool\"}],\"name\":\"getPositionLeverage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_usdgAmount\",\"type\":\"uint256\"}],\"name\":\"getRedemptionAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"getRedemptionCollateral\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"getRedemptionCollateralUsd\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"getTargetUsdgAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"getUtilisation\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"globalShortAveragePrices\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"globalShortSizes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gov\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"guaranteedUsd\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"hasDynamicFees\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"inManagerMode\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"inPrivateLiquidationMode\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"includeAmmPrice\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_collateralToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_indexToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_sizeDelta\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_isLong\",\"type\":\"bool\"}],\"name\":\"increasePosition\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_router\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_usdg\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_priceFeed\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_liquidationFeeUsd\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_fundingRateFactor\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_stableFundingRateFactor\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isInitialized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isLeverageEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isLiquidator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isManager\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isSwapEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"lastFundingTimes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_collateralToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_indexToken\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_isLong\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"_feeReceiver\",\"type\":\"address\"}],\"name\":\"liquidatePosition\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"liquidationFeeUsd\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"marginFeeBasisPoints\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxGasPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"maxGlobalShortSizes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxLeverage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"maxUsdgAmounts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"minProfitBasisPoints\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minProfitTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mintBurnFeeBasisPoints\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"poolAmounts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"positions\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"size\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"averagePrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"entryFundingRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveAmount\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"realisedPnl\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"lastIncreasedTime\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"priceFeed\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_router\",\"type\":\"address\"}],\"name\":\"removeRouter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"reservedAmounts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"router\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"sellUSDG\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"setBufferAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_errorCode\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_error\",\"type\":\"string\"}],\"name\":\"setError\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_errorController\",\"type\":\"address\"}],\"name\":\"setErrorController\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_taxBasisPoints\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_stableTaxBasisPoints\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_mintBurnFeeBasisPoints\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_swapFeeBasisPoints\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_stableSwapFeeBasisPoints\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_marginFeeBasisPoints\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_liquidationFeeUsd\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minProfitTime\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_hasDynamicFees\",\"type\":\"bool\"}],\"name\":\"setFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_fundingInterval\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_fundingRateFactor\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_stableFundingRateFactor\",\"type\":\"uint256\"}],\"name\":\"setFundingRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_gov\",\"type\":\"address\"}],\"name\":\"setGov\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_inManagerMode\",\"type\":\"bool\"}],\"name\":\"setInManagerMode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_inPrivateLiquidationMode\",\"type\":\"bool\"}],\"name\":\"setInPrivateLiquidationMode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_isLeverageEnabled\",\"type\":\"bool\"}],\"name\":\"setIsLeverageEnabled\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_isSwapEnabled\",\"type\":\"bool\"}],\"name\":\"setIsSwapEnabled\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_liquidator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_isActive\",\"type\":\"bool\"}],\"name\":\"setLiquidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_manager\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_isManager\",\"type\":\"bool\"}],\"name\":\"setManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_maxGasPrice\",\"type\":\"uint256\"}],\"name\":\"setMaxGasPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"setMaxGlobalShortSize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_maxLeverage\",\"type\":\"uint256\"}],\"name\":\"setMaxLeverage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_priceFeed\",\"type\":\"address\"}],\"name\":\"setPriceFeed\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenDecimals\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_tokenWeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minProfitBps\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxUsdgAmount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_isStable\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"_isShortable\",\"type\":\"bool\"}],\"name\":\"setTokenConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"setUsdgAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIVaultUtils\",\"name\":\"_vaultUtils\",\"type\":\"address\"}],\"name\":\"setVaultUtils\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"shortableTokens\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stableFundingRateFactor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stableSwapFeeBasisPoints\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stableTaxBasisPoints\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"stableTokens\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_tokenOut\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"swap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"swapFeeBasisPoints\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"taxBasisPoints\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"tokenBalances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"tokenDecimals\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenAmount\",\"type\":\"uint256\"}],\"name\":\"tokenToUsdMin\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"tokenWeights\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalTokenWeights\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_collateralToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_indexToken\",\"type\":\"address\"}],\"name\":\"updateCumulativeFundingRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newVault\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"upgradeVault\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_usdAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"}],\"name\":\"usdToToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_usdAmount\",\"type\":\"uint256\"}],\"name\":\"usdToTokenMax\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_usdAmount\",\"type\":\"uint256\"}],\"name\":\"usdToTokenMin\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"usdg\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"usdgAmounts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"useSwapPricing\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_collateralToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_indexToken\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_isLong\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"_raise\",\"type\":\"bool\"}],\"name\":\"validateLiquidation\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vaultUtils\",\"outputs\":[{\"internalType\":\"contractIVaultUtils\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"whitelistedTokenCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"whitelistedTokens\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"withdrawFees\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// VaultFuncSigs maps the 4-byte function signature to its string representation.
var VaultFuncSigs = map[string]string{
	"126082cf": "BASIS_POINTS_DIVISOR()",
	"6be6026b": "FUNDING_RATE_PRECISION()",
	"4befe2ca": "MAX_FEE_BASIS_POINTS()",
	"8a39735a": "MAX_FUNDING_RATE_FACTOR()",
	"07c58752": "MAX_LIQUIDATION_FEE_USD()",
	"fce28c10": "MIN_FUNDING_RATE_INTERVAL()",
	"34c1557d": "MIN_LEVERAGE()",
	"95082d25": "PRICE_PRECISION()",
	"870d917c": "USDG_DECIMALS()",
	"24ca984e": "addRouter(address)",
	"42152873": "adjustForDecimals(uint256,address,address)",
	"e468baf0": "allWhitelistedTokens(uint256)",
	"0842b076": "allWhitelistedTokensLength()",
	"60922199": "approvedRouters(address,address)",
	"4a993ee9": "bufferAmounts(address)",
	"817bb857": "buyUSDG(address,address)",
	"e67f59a7": "clearTokenConfig(address)",
	"c65bc7b1": "cumulativeFundingRates(address)",
	"82a08490": "decreasePosition(address,address,address,uint256,uint256,bool,address)",
	"5f7bc119": "directPoolDeposit(address)",
	"48f35cbb": "errorController()",
	"fed1a606": "errors(uint256)",
	"1ce9cb8f": "feeReserves(address)",
	"9849e412": "fundingInterval()",
	"c4f718bf": "fundingRateFactor()",
	"5c07eaab": "getDelta(address,uint256,uint256,bool,uint256)",
	"b1cc53ab": "getEntryFundingRate(address,address,bool)",
	"c7e074c3": "getFeeBasisPoints(address,uint256,uint256,uint256,bool)",
	"da76524c": "getFundingFee(address,address,address,bool,uint256,uint256)",
	"b364accb": "getGlobalShortDelta(address)",
	"e124e6d2": "getMaxPrice(address)",
	"81a612d6": "getMinPrice(address)",
	"db97495f": "getNextAveragePrice(address,uint256,uint256,bool,uint256,uint256,uint256)",
	"a93acac2": "getNextFundingRate(address)",
	"9d7432ca": "getNextGlobalShortAveragePrice(address,uint256,uint256)",
	"4a3f088d": "getPosition(address,address,address,bool)",
	"45a6f370": "getPositionDelta(address,address,address,bool)",
	"fdaf6ac3": "getPositionFee(address,address,address,bool,uint256)",
	"2d4b0576": "getPositionKey(address,address,address,bool)",
	"51723e82": "getPositionLeverage(address,address,address,bool)",
	"2c668ec1": "getRedemptionAmount(address,uint256)",
	"b136ca49": "getRedemptionCollateral(address)",
	"29ff9615": "getRedemptionCollateralUsd(address)",
	"3a05dcc1": "getTargetUsdgAmount(address)",
	"04fef1db": "getUtilisation(address)",
	"62749803": "globalShortAveragePrices(address)",
	"8a78daa8": "globalShortSizes(address)",
	"12d43a51": "gov()",
	"f07456ce": "guaranteedUsd(address)",
	"9f392eb3": "hasDynamicFees()",
	"9060b1ca": "inManagerMode()",
	"181e210e": "inPrivateLiquidationMode()",
	"ab08c1c6": "includeAmmPrice()",
	"48d91abf": "increasePosition(address,address,address,uint256,bool)",
	"728cdbca": "initialize(address,address,address,uint256,uint256,uint256)",
	"392e53cd": "isInitialized()",
	"3e72a262": "isLeverageEnabled()",
	"529a356f": "isLiquidator(address)",
	"f3ae2415": "isManager(address)",
	"351a964d": "isSwapEnabled()",
	"d8f897c3": "lastFundingTimes(address)",
	"de2ea948": "liquidatePosition(address,address,address,bool,address)",
	"174d2694": "liquidationFeeUsd()",
	"318bc689": "marginFeeBasisPoints()",
	"3de39c11": "maxGasPrice()",
	"9698d25a": "maxGlobalShortSizes(address)",
	"ae3302c2": "maxLeverage()",
	"ad1e4f8d": "maxUsdgAmounts(address)",
	"88b1fbdf": "minProfitBasisPoints(address)",
	"d9ac4225": "minProfitTime()",
	"4d47b304": "mintBurnFeeBasisPoints()",
	"52f55eed": "poolAmounts(address)",
	"514ea4bf": "positions(bytes32)",
	"741bef1a": "priceFeed()",
	"6ae0b154": "removeRouter(address)",
	"c3c7b9e9": "reservedAmounts(address)",
	"f887ea40": "router()",
	"711e6190": "sellUSDG(address,address)",
	"8585f4d2": "setBufferAmount(address,uint256)",
	"28e67be5": "setError(uint256,string)",
	"8f7b8404": "setErrorController(address)",
	"40eb3802": "setFees(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,bool)",
	"8a27d468": "setFundingRate(uint256,uint256,uint256)",
	"cfad57a2": "setGov(address)",
	"24b0c04d": "setInManagerMode(bool)",
	"f07bbf77": "setInPrivateLiquidationMode(bool)",
	"7c2eb9f7": "setIsLeverageEnabled(bool)",
	"30455ede": "setIsSwapEnabled(bool)",
	"4453a374": "setLiquidator(address,bool)",
	"a5e90eee": "setManager(address,bool)",
	"d2fa635e": "setMaxGasPrice(uint256)",
	"efa10a6e": "setMaxGlobalShortSize(address,uint256)",
	"d3127e63": "setMaxLeverage(uint256)",
	"724e78da": "setPriceFeed(address)",
	"3c5a6e35": "setTokenConfig(address,uint256,uint256,uint256,uint256,bool,bool)",
	"d66b000d": "setUsdgAmount(address,uint256)",
	"71089f4d": "setVaultUtils(address)",
	"db3555fb": "shortableTokens(address)",
	"134ca63b": "stableFundingRateFactor()",
	"df73a267": "stableSwapFeeBasisPoints()",
	"10eb56c2": "stableTaxBasisPoints()",
	"42b60b03": "stableTokens(address)",
	"93316212": "swap(address,address,address)",
	"a22f2392": "swapFeeBasisPoints()",
	"7a210a2b": "taxBasisPoints()",
	"523fba7f": "tokenBalances(address)",
	"8ee573ac": "tokenDecimals(address)",
	"0a48d5a9": "tokenToUsdMin(address,uint256)",
	"ab2f3ad4": "tokenWeights(address)",
	"dc8f5fac": "totalTokenWeights()",
	"fbfded6d": "updateCumulativeFundingRate(address,address)",
	"cea0c328": "upgradeVault(address,address,uint256)",
	"fa12dbc0": "usdToToken(address,uint256,uint256)",
	"a42ab3d2": "usdToTokenMax(address,uint256)",
	"9899cd02": "usdToTokenMin(address,uint256)",
	"f5b91b7b": "usdg()",
	"1aa4ace5": "usdgAmounts(address)",
	"b06423f3": "useSwapPricing()",
	"d54d5a9f": "validateLiquidation(address,address,address,bool,bool)",
	"6abbe0c8": "vaultUtils()",
	"62287a32": "whitelistedTokenCount()",
	"daf9c210": "whitelistedTokens(address)",
	"f2555278": "withdrawFees(address,address)",
}

// VaultBin is the compiled bytecode used for deploying new contracts.
var VaultBin = "0x60806040526001805462ff00001961ff00199091166101001716620100001781556207a1206008556032600a9081556014600b55601e600c819055600d556004600e55600f556011805460ff199081169091556170806012556016805463ffffff0019921690921716905534801561007657600080fd5b506001600055600680546001600160a01b031916331790556160fb806200009e6000396000f3fe608060405234801561001057600080fd5b50600436106107595760003560e01c80638585f4d2116103c9578063c7e074c3116101ff578063df73a26711610125578063f3ae2415116100b8578063fbfded6d11610087578063fbfded6d14611820578063fce28c101461184e578063fdaf6ac314611856578063fed1a6061461189a57610759565b8063f3ae2415146117b8578063f5b91b7b146117de578063f887ea40146117e6578063fa12dbc0146117ee57610759565b8063efa10a6e116100f4578063efa10a6e14611719578063f07456ce14611745578063f07bbf771461176b578063f25552781461178a57610759565b8063df73a267146116a8578063e124e6d2146116b0578063e468baf0146116d6578063e67f59a7146116f357610759565b8063d8f897c31161019d578063db3555fb1161016c578063db3555fb146115e6578063db97495f1461160c578063dc8f5fac14611658578063de2ea9481461166057610759565b8063d8f897c314611548578063d9ac42251461156e578063da76524c14611576578063daf9c210146115c057610759565b8063d2fa635e116101d9578063d2fa635e14611483578063d3127e63146114a0578063d54d5a9f146114bd578063d66b000d1461151c57610759565b8063c7e074c3146113e7578063cea0c32814611427578063cfad57a21461145d57610759565b80639f392eb3116102ef578063ae3302c211610282578063b364accb11610251578063b364accb1461136d578063c3c7b9e914611393578063c4f718bf146113b9578063c65bc7b1146113c157610759565b8063ae3302c2146112ff578063b06423f314611307578063b136ca491461130f578063b1cc53ab1461133557610759565b8063a93acac2116102be578063a93acac214611285578063ab08c1c6146112ab578063ab2f3ad4146112b3578063ad1e4f8d146112d957610759565b80639f392eb31461121b578063a22f239214611223578063a42ab3d21461122b578063a5e90eee1461125757610759565b80638f7b8404116103675780639698d25a116103365780639698d25a1461118f5780639849e412146111b55780639899cd02146111bd5780639d7432ca146111e957610759565b80638f7b8404146111215780639060b1ca14611147578063933162121461114f57806395082d251461118757610759565b80638a27d468116103a35780638a27d468146110ac5780638a39735a146107da5780638a78daa8146110d55780638ee573ac146110fb57610759565b80638585f4d214611052578063870d917c1461107e57806388b1fbdf1461108657610759565b80634453a3741161059e57806360922199116104c4578063724e78da116104575780637c2eb9f7116104265780637c2eb9f714610f8b578063817bb85714610faa57806381a612d614610fd857806382a0849014610ffe57610759565b8063724e78da14610f0d578063728cdbca14610f33578063741bef1a14610f7b5780637a210a2b14610f8357610759565b80636ae0b154116104935780636ae0b15414610e8b5780636be6026b14610eb157806371089f4d14610eb9578063711e619014610edf57610759565b80636092219914610e2757806362287a3214610e555780636274980314610e5d5780636abbe0c814610e8357610759565b80634d47b3041161053c578063529a356f1161050b578063529a356f14610d7557806352f55eed14610d9b5780635c07eaab14610dc15780635f7bc11914610e0157610759565b80634d47b30414610cb4578063514ea4bf14610cbc57806351723e8214610d11578063523fba7f14610d4f57610759565b806348f35cbb1161057857806348f35cbb14610bfd5780634a3f088d14610c055780634a993ee914610c865780634befe2ca14610cac57610759565b80634453a37414610b3257806345a6f37014610b6057806348d91abf14610bb957610759565b806329ff961511610683578063392e53cd116106215780633e72a262116105f05780633e72a26214610a7f57806340eb380214610a875780634215287314610ad857806342b60b0314610b0c57610759565b8063392e53cd146109fb5780633a05dcc114610a035780633c5a6e3514610a295780633de39c1114610a7757610759565b806330455ede1161065d57806330455ede146109cc578063318bc689146109eb57806334c1557d146107da578063351a964d146109f357610759565b806329ff96151461093c5780632c668ec1146109625780632d4b05761461098e57610759565b8063134ca63b116106fb5780631ce9cb8f116106ca5780631ce9cb8f1461085857806324b0c04d1461087e57806324ca984e1461089f57806328e67be5146108c557610759565b8063134ca63b14610806578063174d26941461080e578063181e210e146108165780631aa4ace51461083257610759565b80630a48d5a9116107375780630a48d5a9146107a657806310eb56c2146107d2578063126082cf146107da57806312d43a51146107e257610759565b806304fef1db1461075e57806307c58752146107965780630842b0761461079e575b600080fd5b6107846004803603602081101561077457600080fd5b50356001600160a01b031661192c565b60408051918252519081900360200190f35b61078461198f565b6107846119a1565b610784600480360360408110156107bc57600080fd5b506001600160a01b0381351690602001356119a7565b6107846119fa565b610784611a00565b6107ea611a06565b604080516001600160a01b039092168252519081900360200190f35b610784611a15565b610784611a1b565b61081e611a21565b604080519115158252519081900360200190f35b6107846004803603602081101561084857600080fd5b50356001600160a01b0316611a31565b6107846004803603602081101561086e57600080fd5b50356001600160a01b0316611a43565b61089d6004803603602081101561089457600080fd5b50351515611a55565b005b61089d600480360360208110156108b557600080fd5b50356001600160a01b0316611a79565b61089d600480360360408110156108db57600080fd5b813591908101906040810160208201356401000000008111156108fd57600080fd5b82018360208201111561090f57600080fd5b8035906020019184600183028401116401000000008311171561093157600080fd5b509092509050611aaa565b6107846004803603602081101561095257600080fd5b50356001600160a01b0316611b28565b6107846004803603604081101561097857600080fd5b506001600160a01b038135169060200135611b3c565b610784600480360360808110156109a457600080fd5b506001600160a01b038135811691602081013582169160408201351690606001351515611b82565b61089d600480360360208110156109e257600080fd5b50351515611bdd565b610784611bff565b61081e611c05565b61081e611c13565b61078460048036036020811015610a1957600080fd5b50356001600160a01b0316611c1c565b61089d600480360360e0811015610a3f57600080fd5b506001600160a01b038135169060208101359060408101359060608101359060808101359060a081013515159060c001351515611cdc565b610784611e1e565b61081e611e24565b61089d6004803603610120811015610a9e57600080fd5b5080359060208101359060408101359060608101359060808101359060a08101359060c08101359060e08101359061010001351515611e33565b61078460048036036060811015610aee57600080fd5b508035906001600160a01b0360208201358116916040013516611eee565b61081e60048036036020811015610b2257600080fd5b50356001600160a01b0316611f90565b61089d60048036036040811015610b4857600080fd5b506001600160a01b0381351690602001351515611fa5565b610b9e60048036036080811015610b7657600080fd5b506001600160a01b038135811691602081013582169160408201351690606001351515611fd8565b60408051921515835260208301919091528051918290030190f35b61089d600480360360a0811015610bcf57600080fd5b506001600160a01b038135811691602081013582169160408201351690606081013590608001351515612076565b6107ea6124b5565b610c4360048036036080811015610c1b57600080fd5b506001600160a01b0381358116916020810135821691604082013516906060013515156124c4565b604080519889526020890197909752878701959095526060870193909352608086019190915260a0850152151560c084015260e083015251908190036101000190f35b61078460048036036020811015610c9c57600080fd5b50356001600160a01b03166125ba565b6107846125cc565b6107846125d2565b610cd960048036036020811015610cd257600080fd5b50356125d8565b604080519788526020880196909652868601949094526060860192909252608085015260a084015260c0830152519081900360e00190f35b61078460048036036080811015610d2757600080fd5b506001600160a01b038135811691602081013582169160408201351690606001351515612615565b61078460048036036020811015610d6557600080fd5b50356001600160a01b03166126bb565b61081e60048036036020811015610d8b57600080fd5b50356001600160a01b03166126cd565b61078460048036036020811015610db157600080fd5b50356001600160a01b03166126e2565b610b9e600480360360a0811015610dd757600080fd5b506001600160a01b03813516906020810135906040810135906060810135151590608001356126f4565b61089d60048036036020811015610e1757600080fd5b50356001600160a01b03166127ec565b61081e60048036036040811015610e3d57600080fd5b506001600160a01b03813581169160200135166128cf565b6107846128ef565b61078460048036036020811015610e7357600080fd5b50356001600160a01b03166128f5565b6107ea612907565b61089d60048036036020811015610ea157600080fd5b50356001600160a01b031661291d565b61078461294b565b61089d60048036036020811015610ecf57600080fd5b50356001600160a01b0316612952565b61078460048036036040811015610ef557600080fd5b506001600160a01b0381358116916020013516612986565b61089d60048036036020811015610f2357600080fd5b50356001600160a01b0316612c1a565b61089d600480360360c0811015610f4957600080fd5b506001600160a01b03813581169160208101358216916040820135169060608101359060808101359060a00135612c44565b6107ea612cbb565b610784612cca565b61089d60048036036020811015610fa157600080fd5b50351515612cd0565b61078460048036036040811015610fc057600080fd5b506001600160a01b0381358116916020013516612cf4565b61078460048036036020811015610fee57600080fd5b50356001600160a01b0316612fc6565b610784600480360360e081101561101457600080fd5b506001600160a01b0381358116916020810135821691604082013581169160608101359160808201359160a081013515159160c09091013516613071565b61089d6004803603604081101561106857600080fd5b506001600160a01b0381351690602001356130ef565b610784613113565b6107846004803603602081101561109c57600080fd5b50356001600160a01b0316613118565b61089d600480360360608110156110c257600080fd5b508035906020810135906040013561312a565b610784600480360360208110156110eb57600080fd5b50356001600160a01b0316613170565b6107846004803603602081101561111157600080fd5b50356001600160a01b0316613182565b61089d6004803603602081101561113757600080fd5b50356001600160a01b0316613194565b61081e6131be565b6107846004803603606081101561116557600080fd5b506001600160a01b0381358116916020810135821691604090910135166131cd565b6107846134b6565b610784600480360360208110156111a557600080fd5b50356001600160a01b03166134c7565b6107846134d9565b610784600480360360408110156111d357600080fd5b506001600160a01b0381351690602001356134df565b610784600480360360608110156111ff57600080fd5b506001600160a01b038135169060208101359060400135613501565b61081e6135aa565b6107846135b3565b6107846004803603604081101561124157600080fd5b506001600160a01b0381351690602001356135b9565b61089d6004803603604081101561126d57600080fd5b506001600160a01b03813516906020013515156135d6565b6107846004803603602081101561129b57600080fd5b50356001600160a01b0316613609565b61081e613707565b610784600480360360208110156112c957600080fd5b50356001600160a01b0316613710565b610784600480360360208110156112ef57600080fd5b50356001600160a01b0316613722565b610784613734565b61081e61373a565b6107846004803603602081101561132557600080fd5b50356001600160a01b0316613748565b6107846004803603606081101561134b57600080fd5b506001600160a01b0381358116916020810135909116906040013515156137e9565b610b9e6004803603602081101561138357600080fd5b50356001600160a01b0316613884565b610784600480360360208110156113a957600080fd5b50356001600160a01b031661391c565b61078461392e565b610784600480360360208110156113d757600080fd5b50356001600160a01b0316613934565b610784600480360360a08110156113fd57600080fd5b506001600160a01b0381351690602081013590604081013590606081013590608001351515613946565b61089d6004803603606081101561143d57600080fd5b506001600160a01b038135811691602081013590911690604001356139f0565b61089d6004803603602081101561147357600080fd5b50356001600160a01b0316613a11565b61089d6004803603602081101561149957600080fd5b5035613a3b565b61089d600480360360208110156114b657600080fd5b5035613a48565b611503600480360360a08110156114d357600080fd5b506001600160a01b0381358116916020810135821691604082013516906060810135151590608001351515613a64565b6040805192835260208301919091528051918290030190f35b61089d6004803603604081101561153257600080fd5b506001600160a01b038135169060200135613b1b565b6107846004803603602081101561155e57600080fd5b50356001600160a01b0316613b74565b610784613b86565b610784600480360360c081101561158c57600080fd5b506001600160a01b038135811691602081013582169160408201351690606081013515159060808101359060a00135613b8c565b61081e600480360360208110156115d657600080fd5b50356001600160a01b0316613c40565b61081e600480360360208110156115fc57600080fd5b50356001600160a01b0316613c55565b610784600480360360e081101561162257600080fd5b506001600160a01b0381351690602081013590604081013590606081013515159060808101359060a08101359060c00135613c69565b610784613cfa565b61089d600480360360a081101561167657600080fd5b506001600160a01b03813581169160208101358216916040820135811691606081013515159160809091013516613d00565b6107846140d8565b610784600480360360208110156116c657600080fd5b50356001600160a01b03166140de565b6107ea600480360360208110156116ec57600080fd5b5035614156565b61089d6004803603602081101561170957600080fd5b50356001600160a01b031661417d565b61089d6004803603604081101561172f57600080fd5b506001600160a01b038135169060200135614250565b6107846004803603602081101561175b57600080fd5b50356001600160a01b0316614274565b61089d6004803603602081101561178157600080fd5b50351515614286565b610784600480360360408110156117a057600080fd5b506001600160a01b03813581169160200135166142ac565b61081e600480360360208110156117ce57600080fd5b50356001600160a01b0316614302565b6107ea614317565b6107ea614326565b6107846004803603606081101561180457600080fd5b506001600160a01b038135169060208101359060400135614335565b61089d6004803603604081101561183657600080fd5b506001600160a01b038135811691602001351661436f565b610784614548565b610784600480360360a081101561186c57600080fd5b506001600160a01b03813581169160208101358216916040820135169060608101351515906080013561454e565b6118b7600480360360208110156118b057600080fd5b50356145c4565b6040805160208082528351818301528351919283929083019185019080838360005b838110156118f15781810151838201526020016118d9565b50505050905090810190601f16801561191e5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6001600160a01b0381166000908152602560205260408120548061195457600091505061198a565b6001600160a01b03831660009081526026602052604090205461198690829061198090620f424061465f565b906146b8565b9150505b919050565b6d04ee2d6d415b85acef810000000081565b601b5490565b6000816119b6575060006119f4565b60006119c184612fc6565b6001600160a01b0385166000908152601d60205260409020549091506119ef600a82900a611980868561465f565b925050505b92915050565b600b5481565b61271081565b6006546001600160a01b031681565b60145481565b60095481565b6016546301000000900460ff1681565b60236020526000908152604090205481565b602c6020526000908152604090205481565b611a5d6146fa565b60168054911515620100000262ff000019909216919091179055565b3360009081526018602090815260408083206001600160a01b0394909416835292905220805460ff19166001179055565b6002546001600160a01b03163314611b09576040805162461bcd60e51b815260206004820152601e60248201527f5661756c743a20696e76616c6964206572726f72436f6e74726f6c6c65720000604482015290519081900360640190fd5b6000838152603060205260409020611b22908383615f8a565b50505050565b60006119f482611b3784613748565b6119a7565b600080611b48846140de565b90506000611b6782611980866c0c9f2c9cd04674edea4000000061465f565b6005549091506119ef9082906001600160a01b031687611eee565b604080516bffffffffffffffffffffffff19606096871b811660208084019190915295871b811660348301529390951b9092166048850152151560f81b605c8401528051603d818503018152605d9093019052815191012090565b611be56146fa565b600180549115156101000261ff0019909216919091179055565b600f5481565b600154610100900460ff1681565b60015460ff1681565b600080600560009054906101000a90046001600160a01b03166001600160a01b03166318160ddd6040518163ffffffff1660e01b815260040160206040518083038186803b158015611c6d57600080fd5b505afa158015611c81573d6000803e3d6000fd5b505050506040513d6020811015611c9757600080fd5b5051905080611caa57600091505061198a565b6001600160a01b038316600090815260226020526040902054601554611cd490611980838561465f565b949350505050565b611ce46146fa565b6001600160a01b0387166000908152601c602052604090205460ff16611d6157600754611d12906001614715565b600755601b80546001810182556000919091527f3ad8aa4f87544323a9d1e5dd902f40c356527a7955687113db5f9a85ad579dc10180546001600160a01b0319166001600160a01b0389161790555b6015546001600160a01b038816600090815260226020526040902054611d8890829061476f565b6001600160a01b0389166000908152601c602090815260408083208054600160ff1991821617909155601d83528184208c9055602283528184208b9055601e83528184208a905560248352818420899055601f83528184208054821689151517905591805290912080549091168415151790559050611e078187614715565b601555611e13886140de565b505050505050505050565b60175481565b60015462010000900460ff1681565b611e3b6146fa565b611e4b6101f48a111560036147b1565b611e5b6101f489111560046147b1565b611e6b6101f488111560056147b1565b611e7b6101f487111560066147b1565b611e8b6101f486111560076147b1565b611e9b6101f485111560086147b1565b611eb76d04ee2d6d415b85acef810000000084111560096147b1565b600a98909855600b96909655600c94909455600d92909255600e55600f556009556010556011805460ff1916911515919091179055565b60055460009081906001600160a01b03858116911614611f26576001600160a01b0384166000908152601d6020526040902054611f29565b60125b6005549091506000906001600160a01b03858116911614611f62576001600160a01b0384166000908152601d6020526040902054611f65565b60125b9050611f8482600a0a61198083600a0a8961465f90919063ffffffff16565b925050505b9392505050565b601f6020526000908152604090205460ff1681565b611fad6146fa565b6001600160a01b03919091166000908152601960205260409020805460ff1916911515919091179055565b6000806000611fe987878787611b82565b9050611ff3616008565b506000818152602b6020908152604091829020825160e081018452815480825260018301549382019390935260028201549381018490526003820154606082015260048201546080820152600582015460a082015260069091015460c08201819052909261206792899290919089906126f4565b93509350505094509492505050565b600260005414156120bc576040805162461bcd60e51b815260206004820152601f602482015260008051602061605b833981519152604482015290519081900360640190fd5b60026000556001546120d89062010000900460ff16601c6147b1565b6120e061485d565b6120e985614879565b6120f48484836148dc565b60015460408051634eae147d60e11b81526001600160a01b038881166004830152878116602483015286811660448301526064820186905284151560848301529151630100000090930490911691639d5c28fa9160a480820192600092909190829003018186803b15801561216857600080fd5b505afa15801561217c573d6000803e3d6000fd5b5050505061218a848461436f565b600061219886868685611b82565b6000818152602b60205260408120919250836121bc576121b786612fc6565b6121c5565b6121c5866140de565b82549091506121d657600282018190555b8154158015906121e65750600085115b1561220c57612206868360000154846002015487858a8860060154613c69565b60028301555b6000612225898989888a886000015489600301546149f1565b9050600061223289614abf565b905060006122408a836119a7565b60018601549091506122529082614715565b6001860181905561226790841115601d6147b1565b6001850154612276908461476f565b60018601556122868a8a896137e9565b600386015584546122979089614715565b8086554260068701556122ad901515601e6147b1565b6122bf85600001548660010154614b6c565b6122cd8b8b8b8a6001613a64565b505060006122db8b8a6135b9565b60048701549091506122ed9082614715565b60048701556122fc8b82614b90565b8715612341576123158b6123108b87614715565b614c30565b61231f8b83614caf565b6123298b84614d2e565b61233c8b6123378d876134df565b614e46565b6123ad565b6001600160a01b038a166000908152602d602052604090205461237e576001600160a01b038a166000908152602e602052604090208590556123a3565b6123898a868b613501565b6001600160a01b038b166000908152602e60205260409020555b6123ad8a8a614f1c565b604080518881526001600160a01b03808f166020830152808e16828401528c1660608201526080810184905260a081018b905289151560c082015260e08101879052610100810186905290517f2fe68525253654c21998f35787a8d0f361905ef647c854092430ab65f2f15022918190036101200190a1855460018701546002880154600389015460048a015460058b0154604080518e81526020810197909752868101959095526060860193909352608085019190915260a084015260c083015260e08201879052517f20853733b590dce729d9f4628682ebd9a34d2354e72679e66f43a008fc03b773918190036101000190a15050600160005550505050505050505050565b6002546001600160a01b031681565b60008060008060008060008060006124de8d8d8d8d611b82565b90506124e8616008565b602b60008381526020019081526020016000206040518060e001604052908160008201548152602001600182015481526020016002820154815260200160038201548152602001600482015481526020016005820154815260200160068201548152505090506000808260a0015113612568578160a0015160000361256e565b8160a001515b9050816000015182602001518360400151846060015185608001518560008860a0015112158860c001519a509a509a509a509a509a509a509a5050505094995094995094999196509450565b60276020526000908152604090205481565b6101f481565b600c5481565b602b602052600090815260409020805460018201546002830154600384015460048501546005860154600690960154949593949293919290919087565b60008061262486868686611b82565b905061262e616008565b506000818152602b6020908152604091829020825160e0810184528154815260018201549281018390526002820154938101939093526003810154606084015260048101546080840152600581015460a08401526006015460c083015261269890151560256147b1565b602081015181516126b091906119809061271061465f565b979650505050505050565b60216020526000908152604090205481565b60196020526000908152604090205460ff1681565b60256020526000908152604090205481565b6000806127056000861160266147b1565b60008461271a57612715886140de565b612723565b61272388612fc6565b9050600081871161273d57612738828861476f565b612747565b612747878361476f565b90506000612759886119808b8561465f565b90506000871561276c5750878311612771565b508288115b60006127886010548961471590919063ffffffff16565b42116127ac576001600160a01b038c166000908152601e60205260409020546127af565b60005b90508180156127d257506127c38b8261465f565b6127cf8461271061465f565b11155b156127dc57600092505b509a909950975050505050505050565b60026000541415612832576040805162461bcd60e51b815260206004820152601f602482015260008051602061605b833981519152604482015290519081900360640190fd5b600260009081556001600160a01b0382168152601c602052604090205461285d9060ff16600e6147b1565b600061286882614abf565b905061287860008211600f6147b1565b6128828282614d2e565b604080516001600160a01b03841681526020810183905281517fa5a389190ebf6170a133bda5c769b77f4d6715b8aa172ec0ddf8473d0b4944bd929181900390910190a150506001600055565b601860209081526000928352604080842090915290825290205460ff1681565b60075481565b602e6020526000908152604090205481565b600154630100000090046001600160a01b031681565b3360009081526018602090815260408083206001600160a01b0394909416835292905220805460ff19169055565b620f424081565b61295a6146fa565b600180546001600160a01b039092166301000000026301000000600160b81b0319909216919091179055565b6000600260005414156129ce576040805162461bcd60e51b815260206004820152601f602482015260008051602061605b833981519152604482015290519081900360640190fd5b60026000556129db614fd7565b6001600160a01b0383166000908152601c6020526040902054612a029060ff1660136147b1565b6016805461ff001916610100179055600554600090612a29906001600160a01b0316614abf565b9050612a396000821160146147b1565b612a43848561436f565b6000612a4f8583611b3c565b9050612a5f6000821160156147b1565b612a698583615006565b612a738582614e46565b60055460408051632770a7eb60e21b81523060048201526024810185905290516001600160a01b0390921691639dc29fac9160448082019260009290919082900301818387803b158015612ac657600080fd5b505af1158015612ada573d6000803e3d6000fd5b5050600554612af492506001600160a01b031690506150e8565b6001546040805163eb0835bf60e01b81526001600160a01b038881166004830152602482018690529151600093630100000090049092169163eb0835bf91604480820192602092909190829003018186803b158015612b5257600080fd5b505afa158015612b66573d6000803e3d6000fd5b505050506040513d6020811015612b7c57600080fd5b505190506000612b8d878484615183565b9050612b9d6000821160166147b1565b612ba887828861524b565b604080516001600160a01b03808916825289166020820152808201869052606081018390526080810184905290517fd732b7828fa6cee72c285eac756fc66a7477e3dc22e22e7c432f1c265d40b4839181900360a00190a16016805461ff001916905560016000559695505050505050565b612c226146fa565b600480546001600160a01b0319166001600160a01b0392909216919091179055565b612c4c6146fa565b60018054612c609160ff90911615906147b1565b6001805460ff191681179055600380546001600160a01b03199081166001600160a01b039889161790915560058054821696881696909617909555600480549095169390951692909217909255600991909155601355601455565b6004546001600160a01b031681565b600a5481565b612cd86146fa565b60018054911515620100000262ff000019909216919091179055565b600060026000541415612d3c576040805162461bcd60e51b815260206004820152601f602482015260008051602061605b833981519152604482015290519081900360640190fd5b6002600055612d49614fd7565b6001600160a01b0383166000908152601c6020526040902054612d709060ff1660106147b1565b6016805461ff0019166101001790556000612d8a84614abf565b9050612d9a6000821160116147b1565b612da4848561436f565b6000612daf85612fc6565b90506000612dce6c0c9f2c9cd04674edea40000000611980858561465f565b600554909150612dea90829088906001600160a01b0316611eee565b9050612dfa6000821160126147b1565b6001546040805163256f6ee360e11b81526001600160a01b0389811660048301526024820185905291516000936301000000900490921691634adeddc691604480820192602092909190829003018186803b158015612e5857600080fd5b505afa158015612e6c573d6000803e3d6000fd5b505050506040513d6020811015612e8257600080fd5b505190506000612e93888684615183565b90506000612eb26c0c9f2c9cd04674edea40000000611980848861465f565b600554909150612ece9082908b906001600160a01b0316611eee565b9050612eda89826152f2565b612ee48983614d2e565b600554604080516340c10f1960e01b81526001600160a01b038b8116600483015260248201859052915191909216916340c10f1991604480830192600092919082900301818387803b158015612f3957600080fd5b505af1158015612f4d573d6000803e3d6000fd5b5050604080516001600160a01b03808d1682528d1660208201528082018a9052606081018590526080810187905290517fab4c77c74cd32c85f35416cf03e7ce9e2d4387f7b7f2c1f4bf53daaecf8ea72d93509081900360a0019150a16016805461ff0019169055600160005598975050505050505050565b60048054601654604080516317e1d38560e11b81526001600160a01b038681169582019590955260006024820181905260ff80851615156044840152610100909404909316151560648201529051919390921691632fc3a70a916084808301926020929190829003018186803b15801561303f57600080fd5b505afa158015613053573d6000803e3d6000fd5b505050506040513d602081101561306957600080fd5b505192915050565b6000600260005414156130b9576040805162461bcd60e51b815260206004820152601f602482015260008051602061605b833981519152604482015290519081900360640190fd5b60026000556130c661485d565b6130cf88614879565b6130de888888888888886153b0565b600160005598975050505050505050565b6130f76146fa565b6001600160a01b03909116600090815260276020526040902055565b601281565b601e6020526000908152604090205481565b6131326146fa565b613142610e10841015600a6147b1565b613152612710831115600b6147b1565b613162612710821115600c6147b1565b601292909255601355601455565b602d6020526000908152604090205481565b601d6020526000908152604090205481565b61319c6146fa565b600280546001600160a01b0319166001600160a01b0392909216919091179055565b60165462010000900460ff1681565b600060026000541415613215576040805162461bcd60e51b815260206004820152601f602482015260008051602061605b833981519152604482015290519081900360640190fd5b600260005560015461323090610100900460ff1660176147b1565b6001600160a01b0384166000908152601c60205260409020546132579060ff1660186147b1565b6001600160a01b0383166000908152601c602052604090205461327e9060ff1660196147b1565b61329e836001600160a01b0316856001600160a01b03161415601a6147b1565b6016805461ff0019166101001790556132b7848061436f565b6132c1838461436f565b60006132cc85614abf565b90506132dc60008211601b6147b1565b60006132e786612fc6565b905060006132f4866140de565b9050600061330682611980868661465f565b9050613313818989611eee565b905060006133326c0c9f2c9cd04674edea40000000611980878761465f565b60055490915061334e9082908b906001600160a01b0316611eee565b60015460408051636d099c0b60e11b81526001600160a01b038d811660048301528c8116602483015260448201859052915193945060009363010000009093049091169163da13381691606480820192602092909190829003018186803b1580156133b857600080fd5b505afa1580156133cc573d6000803e3d6000fd5b505050506040513d60208110156133e257600080fd5b5051905060006133f38a8584615183565b90506133ff8b846152f2565b6134098a84615006565b6134138b88614d2e565b61341d8a85614e46565b6134268a615873565b6134318a828b61524b565b604080516001600160a01b03808c168252808e1660208301528c1681830152606081018990526080810186905260a0810183905260c0810184905290517f0874b2d545cb271cdbda4e093020c452328b24af12382ed62c4d00f5c26709db9181900360e00190a16016805461ff001916905560016000559a9950505050505050505050565b6c0c9f2c9cd04674edea4000000081565b602f6020526000908152604090205481565b60125481565b6000816134ee575060006119f4565b611f8983836134fc866140de565b614335565b6001600160a01b0383166000908152602d6020908152604080832054602e9092528220548285821161353c57613537868361476f565b613546565b613546828761476f565b9050600061355883611980868561465f565b905086831160006135698689614715565b90506000826135815761357c8285614715565b61358b565b61358b828561476f565b905061359b816119808c8561465f565b9b9a5050505050505050505050565b60115460ff1681565b600d5481565b6000816135c8575060006119f4565b611f8983836134fc86612fc6565b6135de6146fa565b6001600160a01b03919091166000908152601a60205260409020805460ff1916911515919091179055565b6012546001600160a01b0382166000908152602a60205260408120549091429161363291614715565b11156136405750600061198a565b6012546001600160a01b0383166000908152602a6020526040812054909161366d9161198090429061476f565b6001600160a01b038416600090815260256020526040902054909150806136995760009250505061198a565b6001600160a01b0384166000908152601f602052604081205460ff166136c1576013546136c5565b6014545b6001600160a01b0386166000908152602660205260409020549091506136fe9083906119809086906136f890869061465f565b9061465f565b95945050505050565b60165460ff1681565b60226020526000908152604090205481565b60246020526000908152604090205481565b60085481565b601654610100900460ff1681565b6001600160a01b0381166000908152601f602052604081205460ff161561378857506001600160a01b03811660009081526025602052604090205461198a565b6001600160a01b0382166000908152602860205260408120546137ac9084906134df565b6001600160a01b038416600090815260266020908152604080832054602590925290912054919250611986916137e3908490614715565b9061476f565b6001546040805163b1cc53ab60e01b81526001600160a01b038681166004830152858116602483015284151560448301529151600093630100000090049092169163b1cc53ab91606480820192602092909190829003018186803b15801561385057600080fd5b505afa158015613864573d6000803e3d6000fd5b505050506040513d602081101561387a57600080fd5b5051949350505050565b6001600160a01b0381166000908152602d60205260408120548190806138b1576000809250925050613917565b60006138bc856140de565b6001600160a01b0386166000908152602e60205260408120549192508282116138ee576138e9838361476f565b6138f8565b6138f8828461476f565b9050600061390a83611980878561465f565b9390921195509193505050505b915091565b60266020526000908152604090205481565b60135481565b60296020526000908152604090205481565b6001546040805163c7e074c360e01b81526001600160a01b03888116600483015260248201889052604482018790526064820186905284151560848301529151600093630100000090049092169163c7e074c39160a480820192602092909190829003018186803b1580156139ba57600080fd5b505afa1580156139ce573d6000803e3d6000fd5b505050506040513d60208110156139e457600080fd5b50519695505050505050565b6139f86146fa565b613a0c6001600160a01b03831684836158eb565b505050565b613a196146fa565b600680546001600160a01b0319166001600160a01b0392909216919091179055565b613a436146fa565b601755565b613a506146fa565b613a5f612710821160026147b1565b600855565b6001546040805163d54d5a9f60e01b81526001600160a01b03888116600483015287811660248301528681166044830152851515606483015284151560848301528251600094859463010000009091049092169263d54d5a9f9260a4808301939192829003018186803b158015613ada57600080fd5b505afa158015613aee573d6000803e3d6000fd5b505050506040513d6040811015613b0457600080fd5b508051602090910151909890975095505050505050565b613b236146fa565b6001600160a01b03821660009081526023602052604090205480821115613b5d57613b5783613b52848461476f565b6152f2565b50613b70565b613a0c83613b6b838561476f565b615006565b5050565b602a6020526000908152604090205481565b60105481565b6001546040805163369d949360e21b81526001600160a01b0389811660048301528881166024830152878116604483015286151560648301526084820186905260a482018590529151600093630100000090049092169163da76524c9160c480820192602092909190829003018186803b158015613c0957600080fd5b505afa158015613c1d573d6000803e3d6000fd5b505050506040513d6020811015613c3357600080fd5b5051979650505050505050565b601c6020526000908152604090205460ff1681565b602080526000908152604090205460ff1681565b6000806000613c7b8a8a8a8a886126f4565b90925090506000613c8c8a87614715565b905060008815613cbb5783613caa57613ca5828461476f565b613cb4565b613cb48284614715565b9050613cdc565b83613ccf57613cca8284614715565b613cd9565b613cd9828461476f565b90505b613cea816119808a8561465f565b9c9b505050505050505050505050565b60155481565b60026000541415613d46576040805162461bcd60e51b815260206004820152601f602482015260008051602061605b833981519152604482015290519081900360640190fd5b60026000556016546301000000900460ff1615613d7b5733600090815260196020526040902054613d7b9060ff1660226147b1565b6016805460ff19169055613d8f848461436f565b6000613d9d86868686611b82565b9050613da7616008565b506000818152602b6020908152604091829020825160e08101845281548082526001830154938201939093526002820154938101939093526003810154606084015260048101546080840152600581015460a08401526006015460c0830152613e1390151560236147b1565b600080613e24898989896000613a64565b91509150613e37826000141560246147b1565b8160021415613e6c57613e54898989600087600001518b8f6153b0565b50506016805460ff19166001179055506140cc915050565b6000613e7889836134df565b6001600160a01b038a166000908152602c6020526040902054909150613e9e9082614715565b6001600160a01b038a166000818152602c602090815260409182902093909355805191825291810184905280820183905290517f5d0c0019d3d45fadeb74eff9d2c9924d146d000ac6bcf3c28bf0ac3c9baa011a9181900360600190a1613f0989856080015161593d565b8615613f385760208401518451613f2a918b91613f259161476f565b614caf565b613f38896123378b856134df565b600087613f4d57613f48896140de565b613f56565b613f5689612fc6565b90507f2e1f85a64a2f22cf2f0c42584e7c919ed4abe8d53675cff0f62bf1e95a1c676f868c8c8c8c8a600001518b602001518c608001518d60a001518a604051808b81526020018a6001600160a01b03168152602001896001600160a01b03168152602001886001600160a01b0316815260200187151581526020018681526020018581526020018481526020018381526020018281526020019a505050505050505050505060405180910390a1871580156140155750846020015183105b1561404357602085015160009061402c908561476f565b90506140418b61403c8d846134df565b614d2e565b505b87614056576140568986600001516159f4565b6000868152602b602052604081208181556001810182905560028101829055600381018290556004810182905560058101829055600601556009546140a2908b906123379082906134df565b6140b88a6140b28c6009546134df565b8961524b565b50506016805460ff19166001179055505050505b50506001600055505050565b600e5481565b60048054601654604080516317e1d38560e11b81526001600160a01b03868116958201959095526001602482015260ff80841615156044830152610100909304909216151560648301525160009390921691632fc3a70a91608480820192602092909190829003018186803b15801561303f57600080fd5b601b818154811061416357fe5b6000918252602090912001546001600160a01b0316905081565b6141856146fa565b6001600160a01b0381166000908152601c60205260409020546141ac9060ff16600d6147b1565b6001600160a01b0381166000908152602260205260409020546015546141d19161476f565b6015556001600160a01b0381166000908152601c60209081526040808320805460ff19908116909155601d835281842084905560228352818420849055601e835281842084905560248352818420849055601f83528184208054821690559180529091208054909116905560075461424a90600161476f565b60075550565b6142586146fa565b6001600160a01b039091166000908152602f6020526040902055565b60286020526000908152604090205481565b61428e6146fa565b6016805491151563010000000263ff00000019909216919091179055565b60006142b66146fa565b6001600160a01b0383166000908152602c6020526040902054806142de5760009150506119f4565b6001600160a01b0384166000908152602c6020526040812055611f8984828561524b565b601a6020526000908152604090205460ff1681565b6005546001600160a01b031681565b6003546001600160a01b031681565b60008261434457506000611f89565b6001600160a01b0384166000908152601d60205260409020546136fe8361198086600a85900a61465f565b6001546040805163fbfded6d60e01b81526001600160a01b03858116600483015284811660248301529151600093630100000090049092169163fbfded6d9160448082019260209290919082900301818787803b1580156143cf57600080fd5b505af11580156143e3573d6000803e3d6000fd5b505050506040513d60208110156143f957600080fd5b50519050806144085750613b70565b6001600160a01b0383166000908152602a602052604090205461445557601254614436906136f842826146b8565b6001600160a01b0384166000908152602a602052604090205550613b70565b6012546001600160a01b0384166000908152602a6020526040902054429161447d9190614715565b11156144895750613b70565b600061449484613609565b6001600160a01b0385166000908152602960205260409020549091506144ba9082614715565b6001600160a01b0385166000908152602960205260409020556012546144e4906136f842826146b8565b6001600160a01b0385166000818152602a602090815260408083209490945560298152908390205483519283529082015281517fa146fc154e1913322e9817d49f0d5c37466c24326e15de10e739a948be815eab929181900390910190a150505050565b610e1081565b6001546040805163fdaf6ac360e01b81526001600160a01b038881166004830152878116602483015286811660448301528515156064830152608482018590529151600093630100000090049092169163fdaf6ac39160a480820192602092909190829003018186803b1580156139ba57600080fd5b60306020908152600091825260409182902080548351601f6002600019610100600186161502019093169290920491820184900484028101840190945280845290918301828280156146575780601f1061462c57610100808354040283529160200191614657565b820191906000526020600020905b81548152906001019060200180831161463a57829003601f168201915b505050505081565b60008261466e575060006119f4565b8282028284828161467b57fe5b0414611f895760405162461bcd60e51b815260040180806020018281038252602181526020018061607b6021913960400191505060405180910390fd5b6000611f8983836040518060400160405280601a81526020017f536166654d6174683a206469766973696f6e206279207a65726f000000000000815250615a5c565b600654614713906001600160a01b0316331460356147b1565b565b600082820183811015611f89576040805162461bcd60e51b815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b6000611f8983836040518060400160405280601e81526020017f536166654d6174683a207375627472616374696f6e206f766572666c6f770000815250615afe565b600081815260306020526040902082613a0c5760405162461bcd60e51b815260206004820190815282546002600019610100600184161502019091160460248301819052909182916044909101908490801561484e5780601f106148235761010080835404028352916020019161484e565b820191906000526020600020905b81548152906001019060200180831161483157829003601f168201915b50509250505060405180910390fd5b60175461486957614713565b6147136017543a111560376147b1565b336001600160a01b038216141561488f576148d9565b6003546001600160a01b03163314156148a7576148d9565b6001600160a01b03811660009081526018602090815260408083203384529091529020546148d99060ff1660296147b1565b50565b801561495557614901826001600160a01b0316846001600160a01b031614602a6147b1565b6001600160a01b0383166000908152601c60205260409020546149289060ff16602b6147b1565b6001600160a01b0383166000908152601f60205260409020546149509060ff1615602c6147b1565b613a0c565b6001600160a01b0383166000908152601c602052604090205461497c9060ff16602d6147b1565b6001600160a01b0383166000908152601f60205260409020546149a39060ff16602e6147b1565b6001600160a01b0382166000908152601f60205260409020546149cb9060ff1615602f6147b1565b6001600160a01b0382166000908152602080526040902054613a0c9060ff1660306147b1565b600080614a01898989898961454e565b90506000614a138a8a8a8a8989613b8c565b9050614a1f8282614715565b91506000614a2d8a846134df565b6001600160a01b038b166000908152602c6020526040902054909150614a539082614715565b6001600160a01b038b166000818152602c602090815260409182902093909355805191825291810185905280820183905290517f5d0c0019d3d45fadeb74eff9d2c9924d146d000ac6bcf3c28bf0ac3c9baa011a9181900360600190a150909998505050505050505050565b6001600160a01b03811660008181526021602090815260408083205481516370a0823160e01b8152306004820152915193949093859391926370a08231926024808301939192829003018186803b158015614b1957600080fd5b505afa158015614b2d573d6000803e3d6000fd5b505050506040513d6020811015614b4357600080fd5b50516001600160a01b03851660009081526021602052604090208190559050611cd4818361476f565b81614b8257614b7d811560276147b1565b613b70565b613b708183101560286147b1565b6001600160a01b038216600090815260266020526040902054614bb39082614715565b6001600160a01b038316600090815260266020818152604080842085905560258252909220549152614be891111560346147b1565b604080516001600160a01b03841681526020810183905281517faa5649d82f5462be9d19b0f2b31a59b2259950a6076550bac9f3a1c07db9f66d929181900390910190a15050565b6001600160a01b038216600090815260286020526040902054614c539082614715565b6001600160a01b03831660008181526028602090815260409182902093909355805191825291810183905281517fd9d4761f75e0d0103b5cbeab941eeb443d7a56a35b5baf2a0787c03f03f4e474929181900390910190a15050565b6001600160a01b038216600090815260286020526040902054614cd2908261476f565b6001600160a01b03831660008181526028602090815260409182902093909355805191825291810183905281517f34e07158b9db50df5613e591c44ea2ebc82834eff4a4dc3a46e000e608261d68929181900390910190a15050565b6001600160a01b038216600090815260256020526040902054614d519082614715565b6001600160a01b03831660008181526025602090815260408083209490945583516370a0823160e01b8152306004820152935191936370a082319260248083019392829003018186803b158015614da757600080fd5b505afa158015614dbb573d6000803e3d6000fd5b505050506040513d6020811015614dd157600080fd5b50516001600160a01b038416600090815260256020526040902054909150614dfd9082101560316147b1565b604080516001600160a01b03851681526020810184905281517f976177fbe09a15e5e43f848844963a42b41ef919ef17ff21a17a5421de8f4737929181900390910190a1505050565b604080518082018252601a81527f5661756c743a20706f6f6c416d6f756e742065786365656465640000000000006020808301919091526001600160a01b038516600090815260259091529190912054614ea1918390615afe565b6001600160a01b03831660009081526025602090815260408083208490556026909152902054614ed491101560326147b1565b604080516001600160a01b03841681526020810183905281517f112726233fbeaeed0f5b1dba5cb0b2b81883dee49fb35ff99fd98ed9f6d31eb0929181900390910190a15050565b6001600160a01b0382166000908152602d6020526040902054614f3f9082614715565b6001600160a01b0383166000908152602d6020908152604080832093909355602f905220548015613a0c576001600160a01b0383166000908152602d6020526040902054811015613a0c576040805162461bcd60e51b815260206004820152601a60248201527f5661756c743a206d61782073686f727473206578636565646564000000000000604482015290519081900360640190fd5b60165462010000900460ff161561471357336000908152601a60205260409020546147139060ff1660366147b1565b6001600160a01b038216600090815260236020526040902054818111615081576001600160a01b0383166000818152602360209081526040808320929092558151928352820183905280517fe1e812596aac93a06ecc4ca627014d18e30f5c33b825160cc9d5c0ba61e452279281900390910190a150613b70565b61508b818361476f565b6001600160a01b03841660008181526023602090815260409182902093909355805191825291810184905281517fe1e812596aac93a06ecc4ca627014d18e30f5c33b825160cc9d5c0ba61e45227929181900390910190a1505050565b6000816001600160a01b03166370a08231306040518263ffffffff1660e01b815260040180826001600160a01b0316815260200191505060206040518083038186803b15801561513757600080fd5b505afa15801561514b573d6000803e3d6000fd5b505050506040513d602081101561516157600080fd5b50516001600160a01b0390921660009081526021602052604090209190915550565b6000806151a0612710611980615199828761476f565b879061465f565b905060006151ae858361476f565b6001600160a01b0387166000908152602c60205260409020549091506151d49082614715565b6001600160a01b0387166000908152602c60205260409020557f47cd9dda0e50ce30bcaaacd0488452b596221c07ac402a581cfae4d3933cac2b8661521981846119a7565b604080516001600160a01b0390931683526020830191909152818101849052519081900360600190a150949350505050565b61525f6001600160a01b03841682846158eb565b604080516370a0823160e01b815230600482015290516001600160a01b038516916370a08231916024808301926020929190829003018186803b1580156152a557600080fd5b505afa1580156152b9573d6000803e3d6000fd5b505050506040513d60208110156152cf57600080fd5b50516001600160a01b039093166000908152602160205260409020929092555050565b6001600160a01b0382166000908152602360205260409020546153159082614715565b6001600160a01b0383166000908152602360209081526040808320939093556024905220548015615367576001600160a01b0383166000908152602360205260409020546153679082101560336147b1565b604080516001600160a01b03851681526020810184905281517f64243679a443432e2293343b77d411ff6144370404618f00ca0d2025d9ca9882929181900390910190a1505050565b600154604080516381d11a2360e01b81526001600160a01b038a8116600483015289811660248301528881166044830152606482018890526084820187905285151560a483015284811660c4830152915160009363010000009004909216916381d11a239160e4808201928692909190829003018186803b15801561543457600080fd5b505afa158015615448573d6000803e3d6000fd5b50505050615456878761436f565b600061546489898987611b82565b6000818152602b60205260409020805491925090615485901515601f6147b1565b615497868260000154101560206147b1565b6154a9878260010154101560216147b1565b6001810154815460048301546000916154c691611980908b61465f565b60048401549091506154d8908261476f565b60048401556154e78b8261593d565b506000806154f98d8d8d8d8d8d615b58565b85549193509150891461569d576155118c8c8a6137e9565b60038501558354615522908a61476f565b80855560018501546155349190614b6c565b6155428d8d8d8b6001613a64565b5050871561556f576155658c61231086600101548661476f90919063ffffffff16565b61556f8c8a614caf565b6000886155845761557f8c6140de565b61558d565b61558d8c612fc6565b90507f93d75d64d1f84fc6f430a64fc578bdd4c1e090e90ea2d51773e626d19de56d30868f8f8f8f8f8f886155c28c8c61476f565b60408051998a526001600160a01b0398891660208b015296881689880152949096166060880152608087019290925260a0860152151560c085015260e084019290925261010083019190915251908190036101200190a18454600186015460028701546003880154600489015460058a0154604080518d81526020810197909752868101959095526060860193909352608085019190915260a084015260c083015260e08201839052517f20853733b590dce729d9f4628682ebd9a34d2354e72679e66f43a008fc03b773918190036101000190a150615811565b87156156b7576156ad8c84614c30565b6156b78c8a614caf565b6000886156cc576156c78c6140de565b6156d5565b6156d58c612fc6565b90507f93d75d64d1f84fc6f430a64fc578bdd4c1e090e90ea2d51773e626d19de56d30868f8f8f8f8f8f8861570a8c8c61476f565b60408051998a526001600160a01b0398891660208b015296881689880152949096166060880152608087019290925260a0860152151560c085015260e084019290925261010083019190915251908190036101200190a18454600186015460028701546003880154600489015460058a0154604080518d81526020810197909752868101959095526060860193909352608085019190915260a084015260c0830152517f73af1d417d82c240fdb6d319b34ad884487c6bf2845d98980cc52ad9171cb4559181900360e00190a1506000858152602b602052604081208181556001810182905560028101829055600381018290556004810182905560058101829055600601555b87615820576158208b8a6159f4565b811561586057871561583a5761583a8c6123378e856134df565b60006158468d836134df565b90506158538d828a61524b565b95506126b0945050505050565b5060009c9b505050505050505050505050565b6001600160a01b03811660009081526027602090815260408083205460259092529091205410156148d9576040805162461bcd60e51b815260206004820152601a60248201527f5661756c743a20706f6f6c416d6f756e74203c20627566666572000000000000604482015290519081900360640190fd5b604080516001600160a01b038416602482015260448082018490528251808303909101815260649091019091526020810180516001600160e01b031663a9059cbb60e01b179052613a0c908490615d52565b604080518082018252601b81527f5661756c743a20696e73756666696369656e74207265736572766500000000006020808301919091526001600160a01b038516600090815260269091529190912054615998918390615afe565b6001600160a01b03831660008181526026602090815260409182902093909355805191825291810183905281517f533cb5ed32be6a90284e96b5747a1bfc2d38fdb5768a6b5f67ff7d62144ed67b929181900390910190a15050565b6001600160a01b0382166000908152602d602052604090205480821115615a3457506001600160a01b0382166000908152602d6020526040812055613b70565b615a3e818361476f565b6001600160a01b0384166000908152602d6020526040902055505050565b60008183615ae85760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b83811015615aad578181015183820152602001615a95565b50505050905090810190601f168015615ada5780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b506000838581615af457fe5b0495945050505050565b60008184841115615b505760405162461bcd60e51b8152602060048201818152835160248401528351909283926044909101919085019080838360008315615aad578181015183820152602001615a95565b505050900390565b6000806000615b6989898987611b82565b6000818152602b60205260408120805460038201549394509092615b96918d918d918d918b918d916149f1565b9050600080600080615bb78d876000015488600201548d8a600601546126f4565b87549195508593509150615bcf906119808d8461465f565b925050506000828015615be25750600082115b15615c125750600584018054820190558088615c12576000615c048e846134df565b9050615c108e82614e46565b505b82158015615c205750600082115b15615c64576001850154615c34908361476f565b600186015588615c58576000615c4a8e846134df565b9050615c568e82614d2e565b505b60058501805483900390555b8a15615c8c57615c74818c614715565b6001860154909150615c86908c61476f565b60018601555b84548a1415615caf576001850154615ca5908290614715565b6000600187015590505b8084811115615cc957615cc2828661476f565b9050615cfd565b6001860154615cd8908661476f565b60018701558915615cfd576000615cef8f876134df565b9050615cfb8f82614e46565b505b60408051888152851515602082015280820185905290517f3ff41bdde87755b687ae83d0221a232b6be51a803330ed9661c1b5d0105e0d8a9181900360600190a1909e909d509b505050505050505050505050565b6060615da7826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b0316615e039092919063ffffffff16565b805190915015613a0c57808060200190516020811015615dc657600080fd5b5051613a0c5760405162461bcd60e51b815260040180806020018281038252602a81526020018061609c602a913960400191505060405180910390fd5b6060611cd4848460008585615e1785615f1e565b615e68576040805162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000604482015290519081900360640190fd5b60006060866001600160a01b031685876040518082805190602001908083835b60208310615ea75780518252601f199092019160209182019101615e88565b6001836020036101000a03801982511681845116808217855250505050505090500191505060006040518083038185875af1925050503d8060008114615f09576040519150601f19603f3d011682016040523d82523d6000602084013e615f0e565b606091505b50915091506126b0828286615f24565b3b151590565b60608315615f33575081611f89565b825115615f435782518084602001fd5b60405162461bcd60e51b8152602060048201818152845160248401528451859391928392604401919085019080838360008315615aad578181015183820152602001615a95565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10615fcb5782800160ff19823516178555615ff8565b82800160010185558215615ff8579182015b82811115615ff8578235825591602001919060010190615fdd565b50616004929150616045565b5090565b6040518060e00160405280600081526020016000815260200160008152602001600081526020016000815260200160008152602001600081525090565b5b80821115616004576000815560010161604656fe5265656e7472616e637947756172643a207265656e7472616e742063616c6c00536166654d6174683a206d756c7469706c69636174696f6e206f766572666c6f775361666545524332303a204552433230206f7065726174696f6e20646964206e6f742073756363656564a2646970667358221220fc57b97e9fd8f0730a17f8e9d8e9f96c2681ba5cbaeb0ec9884f37e49213546c64736f6c634300060c0033"

// DeployVault deploys a new Ethereum contract, binding an instance of Vault to it.
func DeployVault(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Vault, error) {
	parsed, err := abi.JSON(strings.NewReader(VaultABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(VaultBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Vault{VaultCaller: VaultCaller{contract: contract}, VaultTransactor: VaultTransactor{contract: contract}, VaultFilterer: VaultFilterer{contract: contract}}, nil
}

// Vault is an auto generated Go binding around an Ethereum contract.
type Vault struct {
	VaultCaller     // Read-only binding to the contract
	VaultTransactor // Write-only binding to the contract
	VaultFilterer   // Log filterer for contract events
}

// VaultCaller is an auto generated read-only Go binding around an Ethereum contract.
type VaultCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VaultTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VaultTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VaultFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VaultFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VaultSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VaultSession struct {
	Contract     *Vault            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VaultCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VaultCallerSession struct {
	Contract *VaultCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// VaultTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VaultTransactorSession struct {
	Contract     *VaultTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VaultRaw is an auto generated low-level Go binding around an Ethereum contract.
type VaultRaw struct {
	Contract *Vault // Generic contract binding to access the raw methods on
}

// VaultCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VaultCallerRaw struct {
	Contract *VaultCaller // Generic read-only contract binding to access the raw methods on
}

// VaultTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VaultTransactorRaw struct {
	Contract *VaultTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVault creates a new instance of Vault, bound to a specific deployed contract.
func NewVault(address common.Address, backend bind.ContractBackend) (*Vault, error) {
	contract, err := bindVault(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Vault{VaultCaller: VaultCaller{contract: contract}, VaultTransactor: VaultTransactor{contract: contract}, VaultFilterer: VaultFilterer{contract: contract}}, nil
}

// NewVaultCaller creates a new read-only instance of Vault, bound to a specific deployed contract.
func NewVaultCaller(address common.Address, caller bind.ContractCaller) (*VaultCaller, error) {
	contract, err := bindVault(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VaultCaller{contract: contract}, nil
}

// NewVaultTransactor creates a new write-only instance of Vault, bound to a specific deployed contract.
func NewVaultTransactor(address common.Address, transactor bind.ContractTransactor) (*VaultTransactor, error) {
	contract, err := bindVault(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VaultTransactor{contract: contract}, nil
}

// NewVaultFilterer creates a new log filterer instance of Vault, bound to a specific deployed contract.
func NewVaultFilterer(address common.Address, filterer bind.ContractFilterer) (*VaultFilterer, error) {
	contract, err := bindVault(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VaultFilterer{contract: contract}, nil
}

// bindVault binds a generic wrapper to an already deployed contract.
func bindVault(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VaultABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Vault *VaultRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Vault.Contract.VaultCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Vault *VaultRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vault.Contract.VaultTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Vault *VaultRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Vault.Contract.VaultTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Vault *VaultCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Vault.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Vault *VaultTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vault.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Vault *VaultTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Vault.Contract.contract.Transact(opts, method, params...)
}

// BASISPOINTSDIVISOR is a free data retrieval call binding the contract method 0x126082cf.
//
// Solidity: function BASIS_POINTS_DIVISOR() view returns(uint256)
func (_Vault *VaultCaller) BASISPOINTSDIVISOR(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "BASIS_POINTS_DIVISOR")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BASISPOINTSDIVISOR is a free data retrieval call binding the contract method 0x126082cf.
//
// Solidity: function BASIS_POINTS_DIVISOR() view returns(uint256)
func (_Vault *VaultSession) BASISPOINTSDIVISOR() (*big.Int, error) {
	return _Vault.Contract.BASISPOINTSDIVISOR(&_Vault.CallOpts)
}

// BASISPOINTSDIVISOR is a free data retrieval call binding the contract method 0x126082cf.
//
// Solidity: function BASIS_POINTS_DIVISOR() view returns(uint256)
func (_Vault *VaultCallerSession) BASISPOINTSDIVISOR() (*big.Int, error) {
	return _Vault.Contract.BASISPOINTSDIVISOR(&_Vault.CallOpts)
}

// FUNDINGRATEPRECISION is a free data retrieval call binding the contract method 0x6be6026b.
//
// Solidity: function FUNDING_RATE_PRECISION() view returns(uint256)
func (_Vault *VaultCaller) FUNDINGRATEPRECISION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "FUNDING_RATE_PRECISION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FUNDINGRATEPRECISION is a free data retrieval call binding the contract method 0x6be6026b.
//
// Solidity: function FUNDING_RATE_PRECISION() view returns(uint256)
func (_Vault *VaultSession) FUNDINGRATEPRECISION() (*big.Int, error) {
	return _Vault.Contract.FUNDINGRATEPRECISION(&_Vault.CallOpts)
}

// FUNDINGRATEPRECISION is a free data retrieval call binding the contract method 0x6be6026b.
//
// Solidity: function FUNDING_RATE_PRECISION() view returns(uint256)
func (_Vault *VaultCallerSession) FUNDINGRATEPRECISION() (*big.Int, error) {
	return _Vault.Contract.FUNDINGRATEPRECISION(&_Vault.CallOpts)
}

// MAXFEEBASISPOINTS is a free data retrieval call binding the contract method 0x4befe2ca.
//
// Solidity: function MAX_FEE_BASIS_POINTS() view returns(uint256)
func (_Vault *VaultCaller) MAXFEEBASISPOINTS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "MAX_FEE_BASIS_POINTS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXFEEBASISPOINTS is a free data retrieval call binding the contract method 0x4befe2ca.
//
// Solidity: function MAX_FEE_BASIS_POINTS() view returns(uint256)
func (_Vault *VaultSession) MAXFEEBASISPOINTS() (*big.Int, error) {
	return _Vault.Contract.MAXFEEBASISPOINTS(&_Vault.CallOpts)
}

// MAXFEEBASISPOINTS is a free data retrieval call binding the contract method 0x4befe2ca.
//
// Solidity: function MAX_FEE_BASIS_POINTS() view returns(uint256)
func (_Vault *VaultCallerSession) MAXFEEBASISPOINTS() (*big.Int, error) {
	return _Vault.Contract.MAXFEEBASISPOINTS(&_Vault.CallOpts)
}

// MAXFUNDINGRATEFACTOR is a free data retrieval call binding the contract method 0x8a39735a.
//
// Solidity: function MAX_FUNDING_RATE_FACTOR() view returns(uint256)
func (_Vault *VaultCaller) MAXFUNDINGRATEFACTOR(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "MAX_FUNDING_RATE_FACTOR")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXFUNDINGRATEFACTOR is a free data retrieval call binding the contract method 0x8a39735a.
//
// Solidity: function MAX_FUNDING_RATE_FACTOR() view returns(uint256)
func (_Vault *VaultSession) MAXFUNDINGRATEFACTOR() (*big.Int, error) {
	return _Vault.Contract.MAXFUNDINGRATEFACTOR(&_Vault.CallOpts)
}

// MAXFUNDINGRATEFACTOR is a free data retrieval call binding the contract method 0x8a39735a.
//
// Solidity: function MAX_FUNDING_RATE_FACTOR() view returns(uint256)
func (_Vault *VaultCallerSession) MAXFUNDINGRATEFACTOR() (*big.Int, error) {
	return _Vault.Contract.MAXFUNDINGRATEFACTOR(&_Vault.CallOpts)
}

// MAXLIQUIDATIONFEEUSD is a free data retrieval call binding the contract method 0x07c58752.
//
// Solidity: function MAX_LIQUIDATION_FEE_USD() view returns(uint256)
func (_Vault *VaultCaller) MAXLIQUIDATIONFEEUSD(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "MAX_LIQUIDATION_FEE_USD")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXLIQUIDATIONFEEUSD is a free data retrieval call binding the contract method 0x07c58752.
//
// Solidity: function MAX_LIQUIDATION_FEE_USD() view returns(uint256)
func (_Vault *VaultSession) MAXLIQUIDATIONFEEUSD() (*big.Int, error) {
	return _Vault.Contract.MAXLIQUIDATIONFEEUSD(&_Vault.CallOpts)
}

// MAXLIQUIDATIONFEEUSD is a free data retrieval call binding the contract method 0x07c58752.
//
// Solidity: function MAX_LIQUIDATION_FEE_USD() view returns(uint256)
func (_Vault *VaultCallerSession) MAXLIQUIDATIONFEEUSD() (*big.Int, error) {
	return _Vault.Contract.MAXLIQUIDATIONFEEUSD(&_Vault.CallOpts)
}

// MINFUNDINGRATEINTERVAL is a free data retrieval call binding the contract method 0xfce28c10.
//
// Solidity: function MIN_FUNDING_RATE_INTERVAL() view returns(uint256)
func (_Vault *VaultCaller) MINFUNDINGRATEINTERVAL(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "MIN_FUNDING_RATE_INTERVAL")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINFUNDINGRATEINTERVAL is a free data retrieval call binding the contract method 0xfce28c10.
//
// Solidity: function MIN_FUNDING_RATE_INTERVAL() view returns(uint256)
func (_Vault *VaultSession) MINFUNDINGRATEINTERVAL() (*big.Int, error) {
	return _Vault.Contract.MINFUNDINGRATEINTERVAL(&_Vault.CallOpts)
}

// MINFUNDINGRATEINTERVAL is a free data retrieval call binding the contract method 0xfce28c10.
//
// Solidity: function MIN_FUNDING_RATE_INTERVAL() view returns(uint256)
func (_Vault *VaultCallerSession) MINFUNDINGRATEINTERVAL() (*big.Int, error) {
	return _Vault.Contract.MINFUNDINGRATEINTERVAL(&_Vault.CallOpts)
}

// MINLEVERAGE is a free data retrieval call binding the contract method 0x34c1557d.
//
// Solidity: function MIN_LEVERAGE() view returns(uint256)
func (_Vault *VaultCaller) MINLEVERAGE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "MIN_LEVERAGE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINLEVERAGE is a free data retrieval call binding the contract method 0x34c1557d.
//
// Solidity: function MIN_LEVERAGE() view returns(uint256)
func (_Vault *VaultSession) MINLEVERAGE() (*big.Int, error) {
	return _Vault.Contract.MINLEVERAGE(&_Vault.CallOpts)
}

// MINLEVERAGE is a free data retrieval call binding the contract method 0x34c1557d.
//
// Solidity: function MIN_LEVERAGE() view returns(uint256)
func (_Vault *VaultCallerSession) MINLEVERAGE() (*big.Int, error) {
	return _Vault.Contract.MINLEVERAGE(&_Vault.CallOpts)
}

// PRICEPRECISION is a free data retrieval call binding the contract method 0x95082d25.
//
// Solidity: function PRICE_PRECISION() view returns(uint256)
func (_Vault *VaultCaller) PRICEPRECISION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "PRICE_PRECISION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PRICEPRECISION is a free data retrieval call binding the contract method 0x95082d25.
//
// Solidity: function PRICE_PRECISION() view returns(uint256)
func (_Vault *VaultSession) PRICEPRECISION() (*big.Int, error) {
	return _Vault.Contract.PRICEPRECISION(&_Vault.CallOpts)
}

// PRICEPRECISION is a free data retrieval call binding the contract method 0x95082d25.
//
// Solidity: function PRICE_PRECISION() view returns(uint256)
func (_Vault *VaultCallerSession) PRICEPRECISION() (*big.Int, error) {
	return _Vault.Contract.PRICEPRECISION(&_Vault.CallOpts)
}

// USDGDECIMALS is a free data retrieval call binding the contract method 0x870d917c.
//
// Solidity: function USDG_DECIMALS() view returns(uint256)
func (_Vault *VaultCaller) USDGDECIMALS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "USDG_DECIMALS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// USDGDECIMALS is a free data retrieval call binding the contract method 0x870d917c.
//
// Solidity: function USDG_DECIMALS() view returns(uint256)
func (_Vault *VaultSession) USDGDECIMALS() (*big.Int, error) {
	return _Vault.Contract.USDGDECIMALS(&_Vault.CallOpts)
}

// USDGDECIMALS is a free data retrieval call binding the contract method 0x870d917c.
//
// Solidity: function USDG_DECIMALS() view returns(uint256)
func (_Vault *VaultCallerSession) USDGDECIMALS() (*big.Int, error) {
	return _Vault.Contract.USDGDECIMALS(&_Vault.CallOpts)
}

// AdjustForDecimals is a free data retrieval call binding the contract method 0x42152873.
//
// Solidity: function adjustForDecimals(uint256 _amount, address _tokenDiv, address _tokenMul) view returns(uint256)
func (_Vault *VaultCaller) AdjustForDecimals(opts *bind.CallOpts, _amount *big.Int, _tokenDiv common.Address, _tokenMul common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "adjustForDecimals", _amount, _tokenDiv, _tokenMul)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AdjustForDecimals is a free data retrieval call binding the contract method 0x42152873.
//
// Solidity: function adjustForDecimals(uint256 _amount, address _tokenDiv, address _tokenMul) view returns(uint256)
func (_Vault *VaultSession) AdjustForDecimals(_amount *big.Int, _tokenDiv common.Address, _tokenMul common.Address) (*big.Int, error) {
	return _Vault.Contract.AdjustForDecimals(&_Vault.CallOpts, _amount, _tokenDiv, _tokenMul)
}

// AdjustForDecimals is a free data retrieval call binding the contract method 0x42152873.
//
// Solidity: function adjustForDecimals(uint256 _amount, address _tokenDiv, address _tokenMul) view returns(uint256)
func (_Vault *VaultCallerSession) AdjustForDecimals(_amount *big.Int, _tokenDiv common.Address, _tokenMul common.Address) (*big.Int, error) {
	return _Vault.Contract.AdjustForDecimals(&_Vault.CallOpts, _amount, _tokenDiv, _tokenMul)
}

// AllWhitelistedTokens is a free data retrieval call binding the contract method 0xe468baf0.
//
// Solidity: function allWhitelistedTokens(uint256 ) view returns(address)
func (_Vault *VaultCaller) AllWhitelistedTokens(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "allWhitelistedTokens", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AllWhitelistedTokens is a free data retrieval call binding the contract method 0xe468baf0.
//
// Solidity: function allWhitelistedTokens(uint256 ) view returns(address)
func (_Vault *VaultSession) AllWhitelistedTokens(arg0 *big.Int) (common.Address, error) {
	return _Vault.Contract.AllWhitelistedTokens(&_Vault.CallOpts, arg0)
}

// AllWhitelistedTokens is a free data retrieval call binding the contract method 0xe468baf0.
//
// Solidity: function allWhitelistedTokens(uint256 ) view returns(address)
func (_Vault *VaultCallerSession) AllWhitelistedTokens(arg0 *big.Int) (common.Address, error) {
	return _Vault.Contract.AllWhitelistedTokens(&_Vault.CallOpts, arg0)
}

// AllWhitelistedTokensLength is a free data retrieval call binding the contract method 0x0842b076.
//
// Solidity: function allWhitelistedTokensLength() view returns(uint256)
func (_Vault *VaultCaller) AllWhitelistedTokensLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "allWhitelistedTokensLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AllWhitelistedTokensLength is a free data retrieval call binding the contract method 0x0842b076.
//
// Solidity: function allWhitelistedTokensLength() view returns(uint256)
func (_Vault *VaultSession) AllWhitelistedTokensLength() (*big.Int, error) {
	return _Vault.Contract.AllWhitelistedTokensLength(&_Vault.CallOpts)
}

// AllWhitelistedTokensLength is a free data retrieval call binding the contract method 0x0842b076.
//
// Solidity: function allWhitelistedTokensLength() view returns(uint256)
func (_Vault *VaultCallerSession) AllWhitelistedTokensLength() (*big.Int, error) {
	return _Vault.Contract.AllWhitelistedTokensLength(&_Vault.CallOpts)
}

// ApprovedRouters is a free data retrieval call binding the contract method 0x60922199.
//
// Solidity: function approvedRouters(address , address ) view returns(bool)
func (_Vault *VaultCaller) ApprovedRouters(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (bool, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "approvedRouters", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ApprovedRouters is a free data retrieval call binding the contract method 0x60922199.
//
// Solidity: function approvedRouters(address , address ) view returns(bool)
func (_Vault *VaultSession) ApprovedRouters(arg0 common.Address, arg1 common.Address) (bool, error) {
	return _Vault.Contract.ApprovedRouters(&_Vault.CallOpts, arg0, arg1)
}

// ApprovedRouters is a free data retrieval call binding the contract method 0x60922199.
//
// Solidity: function approvedRouters(address , address ) view returns(bool)
func (_Vault *VaultCallerSession) ApprovedRouters(arg0 common.Address, arg1 common.Address) (bool, error) {
	return _Vault.Contract.ApprovedRouters(&_Vault.CallOpts, arg0, arg1)
}

// BufferAmounts is a free data retrieval call binding the contract method 0x4a993ee9.
//
// Solidity: function bufferAmounts(address ) view returns(uint256)
func (_Vault *VaultCaller) BufferAmounts(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "bufferAmounts", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BufferAmounts is a free data retrieval call binding the contract method 0x4a993ee9.
//
// Solidity: function bufferAmounts(address ) view returns(uint256)
func (_Vault *VaultSession) BufferAmounts(arg0 common.Address) (*big.Int, error) {
	return _Vault.Contract.BufferAmounts(&_Vault.CallOpts, arg0)
}

// BufferAmounts is a free data retrieval call binding the contract method 0x4a993ee9.
//
// Solidity: function bufferAmounts(address ) view returns(uint256)
func (_Vault *VaultCallerSession) BufferAmounts(arg0 common.Address) (*big.Int, error) {
	return _Vault.Contract.BufferAmounts(&_Vault.CallOpts, arg0)
}

// CumulativeFundingRates is a free data retrieval call binding the contract method 0xc65bc7b1.
//
// Solidity: function cumulativeFundingRates(address ) view returns(uint256)
func (_Vault *VaultCaller) CumulativeFundingRates(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "cumulativeFundingRates", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CumulativeFundingRates is a free data retrieval call binding the contract method 0xc65bc7b1.
//
// Solidity: function cumulativeFundingRates(address ) view returns(uint256)
func (_Vault *VaultSession) CumulativeFundingRates(arg0 common.Address) (*big.Int, error) {
	return _Vault.Contract.CumulativeFundingRates(&_Vault.CallOpts, arg0)
}

// CumulativeFundingRates is a free data retrieval call binding the contract method 0xc65bc7b1.
//
// Solidity: function cumulativeFundingRates(address ) view returns(uint256)
func (_Vault *VaultCallerSession) CumulativeFundingRates(arg0 common.Address) (*big.Int, error) {
	return _Vault.Contract.CumulativeFundingRates(&_Vault.CallOpts, arg0)
}

// ErrorController is a free data retrieval call binding the contract method 0x48f35cbb.
//
// Solidity: function errorController() view returns(address)
func (_Vault *VaultCaller) ErrorController(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "errorController")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ErrorController is a free data retrieval call binding the contract method 0x48f35cbb.
//
// Solidity: function errorController() view returns(address)
func (_Vault *VaultSession) ErrorController() (common.Address, error) {
	return _Vault.Contract.ErrorController(&_Vault.CallOpts)
}

// ErrorController is a free data retrieval call binding the contract method 0x48f35cbb.
//
// Solidity: function errorController() view returns(address)
func (_Vault *VaultCallerSession) ErrorController() (common.Address, error) {
	return _Vault.Contract.ErrorController(&_Vault.CallOpts)
}

// Errors is a free data retrieval call binding the contract method 0xfed1a606.
//
// Solidity: function errors(uint256 ) view returns(string)
func (_Vault *VaultCaller) Errors(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "errors", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Errors is a free data retrieval call binding the contract method 0xfed1a606.
//
// Solidity: function errors(uint256 ) view returns(string)
func (_Vault *VaultSession) Errors(arg0 *big.Int) (string, error) {
	return _Vault.Contract.Errors(&_Vault.CallOpts, arg0)
}

// Errors is a free data retrieval call binding the contract method 0xfed1a606.
//
// Solidity: function errors(uint256 ) view returns(string)
func (_Vault *VaultCallerSession) Errors(arg0 *big.Int) (string, error) {
	return _Vault.Contract.Errors(&_Vault.CallOpts, arg0)
}

// FeeReserves is a free data retrieval call binding the contract method 0x1ce9cb8f.
//
// Solidity: function feeReserves(address ) view returns(uint256)
func (_Vault *VaultCaller) FeeReserves(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "feeReserves", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FeeReserves is a free data retrieval call binding the contract method 0x1ce9cb8f.
//
// Solidity: function feeReserves(address ) view returns(uint256)
func (_Vault *VaultSession) FeeReserves(arg0 common.Address) (*big.Int, error) {
	return _Vault.Contract.FeeReserves(&_Vault.CallOpts, arg0)
}

// FeeReserves is a free data retrieval call binding the contract method 0x1ce9cb8f.
//
// Solidity: function feeReserves(address ) view returns(uint256)
func (_Vault *VaultCallerSession) FeeReserves(arg0 common.Address) (*big.Int, error) {
	return _Vault.Contract.FeeReserves(&_Vault.CallOpts, arg0)
}

// FundingInterval is a free data retrieval call binding the contract method 0x9849e412.
//
// Solidity: function fundingInterval() view returns(uint256)
func (_Vault *VaultCaller) FundingInterval(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "fundingInterval")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FundingInterval is a free data retrieval call binding the contract method 0x9849e412.
//
// Solidity: function fundingInterval() view returns(uint256)
func (_Vault *VaultSession) FundingInterval() (*big.Int, error) {
	return _Vault.Contract.FundingInterval(&_Vault.CallOpts)
}

// FundingInterval is a free data retrieval call binding the contract method 0x9849e412.
//
// Solidity: function fundingInterval() view returns(uint256)
func (_Vault *VaultCallerSession) FundingInterval() (*big.Int, error) {
	return _Vault.Contract.FundingInterval(&_Vault.CallOpts)
}

// FundingRateFactor is a free data retrieval call binding the contract method 0xc4f718bf.
//
// Solidity: function fundingRateFactor() view returns(uint256)
func (_Vault *VaultCaller) FundingRateFactor(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "fundingRateFactor")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FundingRateFactor is a free data retrieval call binding the contract method 0xc4f718bf.
//
// Solidity: function fundingRateFactor() view returns(uint256)
func (_Vault *VaultSession) FundingRateFactor() (*big.Int, error) {
	return _Vault.Contract.FundingRateFactor(&_Vault.CallOpts)
}

// FundingRateFactor is a free data retrieval call binding the contract method 0xc4f718bf.
//
// Solidity: function fundingRateFactor() view returns(uint256)
func (_Vault *VaultCallerSession) FundingRateFactor() (*big.Int, error) {
	return _Vault.Contract.FundingRateFactor(&_Vault.CallOpts)
}

// GetDelta is a free data retrieval call binding the contract method 0x5c07eaab.
//
// Solidity: function getDelta(address _indexToken, uint256 _size, uint256 _averagePrice, bool _isLong, uint256 _lastIncreasedTime) view returns(bool, uint256)
func (_Vault *VaultCaller) GetDelta(opts *bind.CallOpts, _indexToken common.Address, _size *big.Int, _averagePrice *big.Int, _isLong bool, _lastIncreasedTime *big.Int) (bool, *big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "getDelta", _indexToken, _size, _averagePrice, _isLong, _lastIncreasedTime)

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
func (_Vault *VaultSession) GetDelta(_indexToken common.Address, _size *big.Int, _averagePrice *big.Int, _isLong bool, _lastIncreasedTime *big.Int) (bool, *big.Int, error) {
	return _Vault.Contract.GetDelta(&_Vault.CallOpts, _indexToken, _size, _averagePrice, _isLong, _lastIncreasedTime)
}

// GetDelta is a free data retrieval call binding the contract method 0x5c07eaab.
//
// Solidity: function getDelta(address _indexToken, uint256 _size, uint256 _averagePrice, bool _isLong, uint256 _lastIncreasedTime) view returns(bool, uint256)
func (_Vault *VaultCallerSession) GetDelta(_indexToken common.Address, _size *big.Int, _averagePrice *big.Int, _isLong bool, _lastIncreasedTime *big.Int) (bool, *big.Int, error) {
	return _Vault.Contract.GetDelta(&_Vault.CallOpts, _indexToken, _size, _averagePrice, _isLong, _lastIncreasedTime)
}

// GetEntryFundingRate is a free data retrieval call binding the contract method 0xb1cc53ab.
//
// Solidity: function getEntryFundingRate(address _collateralToken, address _indexToken, bool _isLong) view returns(uint256)
func (_Vault *VaultCaller) GetEntryFundingRate(opts *bind.CallOpts, _collateralToken common.Address, _indexToken common.Address, _isLong bool) (*big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "getEntryFundingRate", _collateralToken, _indexToken, _isLong)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetEntryFundingRate is a free data retrieval call binding the contract method 0xb1cc53ab.
//
// Solidity: function getEntryFundingRate(address _collateralToken, address _indexToken, bool _isLong) view returns(uint256)
func (_Vault *VaultSession) GetEntryFundingRate(_collateralToken common.Address, _indexToken common.Address, _isLong bool) (*big.Int, error) {
	return _Vault.Contract.GetEntryFundingRate(&_Vault.CallOpts, _collateralToken, _indexToken, _isLong)
}

// GetEntryFundingRate is a free data retrieval call binding the contract method 0xb1cc53ab.
//
// Solidity: function getEntryFundingRate(address _collateralToken, address _indexToken, bool _isLong) view returns(uint256)
func (_Vault *VaultCallerSession) GetEntryFundingRate(_collateralToken common.Address, _indexToken common.Address, _isLong bool) (*big.Int, error) {
	return _Vault.Contract.GetEntryFundingRate(&_Vault.CallOpts, _collateralToken, _indexToken, _isLong)
}

// GetFeeBasisPoints is a free data retrieval call binding the contract method 0xc7e074c3.
//
// Solidity: function getFeeBasisPoints(address _token, uint256 _usdgDelta, uint256 _feeBasisPoints, uint256 _taxBasisPoints, bool _increment) view returns(uint256)
func (_Vault *VaultCaller) GetFeeBasisPoints(opts *bind.CallOpts, _token common.Address, _usdgDelta *big.Int, _feeBasisPoints *big.Int, _taxBasisPoints *big.Int, _increment bool) (*big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "getFeeBasisPoints", _token, _usdgDelta, _feeBasisPoints, _taxBasisPoints, _increment)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetFeeBasisPoints is a free data retrieval call binding the contract method 0xc7e074c3.
//
// Solidity: function getFeeBasisPoints(address _token, uint256 _usdgDelta, uint256 _feeBasisPoints, uint256 _taxBasisPoints, bool _increment) view returns(uint256)
func (_Vault *VaultSession) GetFeeBasisPoints(_token common.Address, _usdgDelta *big.Int, _feeBasisPoints *big.Int, _taxBasisPoints *big.Int, _increment bool) (*big.Int, error) {
	return _Vault.Contract.GetFeeBasisPoints(&_Vault.CallOpts, _token, _usdgDelta, _feeBasisPoints, _taxBasisPoints, _increment)
}

// GetFeeBasisPoints is a free data retrieval call binding the contract method 0xc7e074c3.
//
// Solidity: function getFeeBasisPoints(address _token, uint256 _usdgDelta, uint256 _feeBasisPoints, uint256 _taxBasisPoints, bool _increment) view returns(uint256)
func (_Vault *VaultCallerSession) GetFeeBasisPoints(_token common.Address, _usdgDelta *big.Int, _feeBasisPoints *big.Int, _taxBasisPoints *big.Int, _increment bool) (*big.Int, error) {
	return _Vault.Contract.GetFeeBasisPoints(&_Vault.CallOpts, _token, _usdgDelta, _feeBasisPoints, _taxBasisPoints, _increment)
}

// GetFundingFee is a free data retrieval call binding the contract method 0xda76524c.
//
// Solidity: function getFundingFee(address _account, address _collateralToken, address _indexToken, bool _isLong, uint256 _size, uint256 _entryFundingRate) view returns(uint256)
func (_Vault *VaultCaller) GetFundingFee(opts *bind.CallOpts, _account common.Address, _collateralToken common.Address, _indexToken common.Address, _isLong bool, _size *big.Int, _entryFundingRate *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "getFundingFee", _account, _collateralToken, _indexToken, _isLong, _size, _entryFundingRate)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetFundingFee is a free data retrieval call binding the contract method 0xda76524c.
//
// Solidity: function getFundingFee(address _account, address _collateralToken, address _indexToken, bool _isLong, uint256 _size, uint256 _entryFundingRate) view returns(uint256)
func (_Vault *VaultSession) GetFundingFee(_account common.Address, _collateralToken common.Address, _indexToken common.Address, _isLong bool, _size *big.Int, _entryFundingRate *big.Int) (*big.Int, error) {
	return _Vault.Contract.GetFundingFee(&_Vault.CallOpts, _account, _collateralToken, _indexToken, _isLong, _size, _entryFundingRate)
}

// GetFundingFee is a free data retrieval call binding the contract method 0xda76524c.
//
// Solidity: function getFundingFee(address _account, address _collateralToken, address _indexToken, bool _isLong, uint256 _size, uint256 _entryFundingRate) view returns(uint256)
func (_Vault *VaultCallerSession) GetFundingFee(_account common.Address, _collateralToken common.Address, _indexToken common.Address, _isLong bool, _size *big.Int, _entryFundingRate *big.Int) (*big.Int, error) {
	return _Vault.Contract.GetFundingFee(&_Vault.CallOpts, _account, _collateralToken, _indexToken, _isLong, _size, _entryFundingRate)
}

// GetGlobalShortDelta is a free data retrieval call binding the contract method 0xb364accb.
//
// Solidity: function getGlobalShortDelta(address _token) view returns(bool, uint256)
func (_Vault *VaultCaller) GetGlobalShortDelta(opts *bind.CallOpts, _token common.Address) (bool, *big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "getGlobalShortDelta", _token)

	if err != nil {
		return *new(bool), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetGlobalShortDelta is a free data retrieval call binding the contract method 0xb364accb.
//
// Solidity: function getGlobalShortDelta(address _token) view returns(bool, uint256)
func (_Vault *VaultSession) GetGlobalShortDelta(_token common.Address) (bool, *big.Int, error) {
	return _Vault.Contract.GetGlobalShortDelta(&_Vault.CallOpts, _token)
}

// GetGlobalShortDelta is a free data retrieval call binding the contract method 0xb364accb.
//
// Solidity: function getGlobalShortDelta(address _token) view returns(bool, uint256)
func (_Vault *VaultCallerSession) GetGlobalShortDelta(_token common.Address) (bool, *big.Int, error) {
	return _Vault.Contract.GetGlobalShortDelta(&_Vault.CallOpts, _token)
}

// GetMaxPrice is a free data retrieval call binding the contract method 0xe124e6d2.
//
// Solidity: function getMaxPrice(address _token) view returns(uint256)
func (_Vault *VaultCaller) GetMaxPrice(opts *bind.CallOpts, _token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "getMaxPrice", _token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMaxPrice is a free data retrieval call binding the contract method 0xe124e6d2.
//
// Solidity: function getMaxPrice(address _token) view returns(uint256)
func (_Vault *VaultSession) GetMaxPrice(_token common.Address) (*big.Int, error) {
	return _Vault.Contract.GetMaxPrice(&_Vault.CallOpts, _token)
}

// GetMaxPrice is a free data retrieval call binding the contract method 0xe124e6d2.
//
// Solidity: function getMaxPrice(address _token) view returns(uint256)
func (_Vault *VaultCallerSession) GetMaxPrice(_token common.Address) (*big.Int, error) {
	return _Vault.Contract.GetMaxPrice(&_Vault.CallOpts, _token)
}

// GetMinPrice is a free data retrieval call binding the contract method 0x81a612d6.
//
// Solidity: function getMinPrice(address _token) view returns(uint256)
func (_Vault *VaultCaller) GetMinPrice(opts *bind.CallOpts, _token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "getMinPrice", _token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMinPrice is a free data retrieval call binding the contract method 0x81a612d6.
//
// Solidity: function getMinPrice(address _token) view returns(uint256)
func (_Vault *VaultSession) GetMinPrice(_token common.Address) (*big.Int, error) {
	return _Vault.Contract.GetMinPrice(&_Vault.CallOpts, _token)
}

// GetMinPrice is a free data retrieval call binding the contract method 0x81a612d6.
//
// Solidity: function getMinPrice(address _token) view returns(uint256)
func (_Vault *VaultCallerSession) GetMinPrice(_token common.Address) (*big.Int, error) {
	return _Vault.Contract.GetMinPrice(&_Vault.CallOpts, _token)
}

// GetNextAveragePrice is a free data retrieval call binding the contract method 0xdb97495f.
//
// Solidity: function getNextAveragePrice(address _indexToken, uint256 _size, uint256 _averagePrice, bool _isLong, uint256 _nextPrice, uint256 _sizeDelta, uint256 _lastIncreasedTime) view returns(uint256)
func (_Vault *VaultCaller) GetNextAveragePrice(opts *bind.CallOpts, _indexToken common.Address, _size *big.Int, _averagePrice *big.Int, _isLong bool, _nextPrice *big.Int, _sizeDelta *big.Int, _lastIncreasedTime *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "getNextAveragePrice", _indexToken, _size, _averagePrice, _isLong, _nextPrice, _sizeDelta, _lastIncreasedTime)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNextAveragePrice is a free data retrieval call binding the contract method 0xdb97495f.
//
// Solidity: function getNextAveragePrice(address _indexToken, uint256 _size, uint256 _averagePrice, bool _isLong, uint256 _nextPrice, uint256 _sizeDelta, uint256 _lastIncreasedTime) view returns(uint256)
func (_Vault *VaultSession) GetNextAveragePrice(_indexToken common.Address, _size *big.Int, _averagePrice *big.Int, _isLong bool, _nextPrice *big.Int, _sizeDelta *big.Int, _lastIncreasedTime *big.Int) (*big.Int, error) {
	return _Vault.Contract.GetNextAveragePrice(&_Vault.CallOpts, _indexToken, _size, _averagePrice, _isLong, _nextPrice, _sizeDelta, _lastIncreasedTime)
}

// GetNextAveragePrice is a free data retrieval call binding the contract method 0xdb97495f.
//
// Solidity: function getNextAveragePrice(address _indexToken, uint256 _size, uint256 _averagePrice, bool _isLong, uint256 _nextPrice, uint256 _sizeDelta, uint256 _lastIncreasedTime) view returns(uint256)
func (_Vault *VaultCallerSession) GetNextAveragePrice(_indexToken common.Address, _size *big.Int, _averagePrice *big.Int, _isLong bool, _nextPrice *big.Int, _sizeDelta *big.Int, _lastIncreasedTime *big.Int) (*big.Int, error) {
	return _Vault.Contract.GetNextAveragePrice(&_Vault.CallOpts, _indexToken, _size, _averagePrice, _isLong, _nextPrice, _sizeDelta, _lastIncreasedTime)
}

// GetNextFundingRate is a free data retrieval call binding the contract method 0xa93acac2.
//
// Solidity: function getNextFundingRate(address _token) view returns(uint256)
func (_Vault *VaultCaller) GetNextFundingRate(opts *bind.CallOpts, _token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "getNextFundingRate", _token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNextFundingRate is a free data retrieval call binding the contract method 0xa93acac2.
//
// Solidity: function getNextFundingRate(address _token) view returns(uint256)
func (_Vault *VaultSession) GetNextFundingRate(_token common.Address) (*big.Int, error) {
	return _Vault.Contract.GetNextFundingRate(&_Vault.CallOpts, _token)
}

// GetNextFundingRate is a free data retrieval call binding the contract method 0xa93acac2.
//
// Solidity: function getNextFundingRate(address _token) view returns(uint256)
func (_Vault *VaultCallerSession) GetNextFundingRate(_token common.Address) (*big.Int, error) {
	return _Vault.Contract.GetNextFundingRate(&_Vault.CallOpts, _token)
}

// GetNextGlobalShortAveragePrice is a free data retrieval call binding the contract method 0x9d7432ca.
//
// Solidity: function getNextGlobalShortAveragePrice(address _indexToken, uint256 _nextPrice, uint256 _sizeDelta) view returns(uint256)
func (_Vault *VaultCaller) GetNextGlobalShortAveragePrice(opts *bind.CallOpts, _indexToken common.Address, _nextPrice *big.Int, _sizeDelta *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "getNextGlobalShortAveragePrice", _indexToken, _nextPrice, _sizeDelta)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNextGlobalShortAveragePrice is a free data retrieval call binding the contract method 0x9d7432ca.
//
// Solidity: function getNextGlobalShortAveragePrice(address _indexToken, uint256 _nextPrice, uint256 _sizeDelta) view returns(uint256)
func (_Vault *VaultSession) GetNextGlobalShortAveragePrice(_indexToken common.Address, _nextPrice *big.Int, _sizeDelta *big.Int) (*big.Int, error) {
	return _Vault.Contract.GetNextGlobalShortAveragePrice(&_Vault.CallOpts, _indexToken, _nextPrice, _sizeDelta)
}

// GetNextGlobalShortAveragePrice is a free data retrieval call binding the contract method 0x9d7432ca.
//
// Solidity: function getNextGlobalShortAveragePrice(address _indexToken, uint256 _nextPrice, uint256 _sizeDelta) view returns(uint256)
func (_Vault *VaultCallerSession) GetNextGlobalShortAveragePrice(_indexToken common.Address, _nextPrice *big.Int, _sizeDelta *big.Int) (*big.Int, error) {
	return _Vault.Contract.GetNextGlobalShortAveragePrice(&_Vault.CallOpts, _indexToken, _nextPrice, _sizeDelta)
}

// GetPosition is a free data retrieval call binding the contract method 0x4a3f088d.
//
// Solidity: function getPosition(address _account, address _collateralToken, address _indexToken, bool _isLong) view returns(uint256, uint256, uint256, uint256, uint256, uint256, bool, uint256)
func (_Vault *VaultCaller) GetPosition(opts *bind.CallOpts, _account common.Address, _collateralToken common.Address, _indexToken common.Address, _isLong bool) (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, bool, *big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "getPosition", _account, _collateralToken, _indexToken, _isLong)

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
func (_Vault *VaultSession) GetPosition(_account common.Address, _collateralToken common.Address, _indexToken common.Address, _isLong bool) (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, bool, *big.Int, error) {
	return _Vault.Contract.GetPosition(&_Vault.CallOpts, _account, _collateralToken, _indexToken, _isLong)
}

// GetPosition is a free data retrieval call binding the contract method 0x4a3f088d.
//
// Solidity: function getPosition(address _account, address _collateralToken, address _indexToken, bool _isLong) view returns(uint256, uint256, uint256, uint256, uint256, uint256, bool, uint256)
func (_Vault *VaultCallerSession) GetPosition(_account common.Address, _collateralToken common.Address, _indexToken common.Address, _isLong bool) (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, bool, *big.Int, error) {
	return _Vault.Contract.GetPosition(&_Vault.CallOpts, _account, _collateralToken, _indexToken, _isLong)
}

// GetPositionDelta is a free data retrieval call binding the contract method 0x45a6f370.
//
// Solidity: function getPositionDelta(address _account, address _collateralToken, address _indexToken, bool _isLong) view returns(bool, uint256)
func (_Vault *VaultCaller) GetPositionDelta(opts *bind.CallOpts, _account common.Address, _collateralToken common.Address, _indexToken common.Address, _isLong bool) (bool, *big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "getPositionDelta", _account, _collateralToken, _indexToken, _isLong)

	if err != nil {
		return *new(bool), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetPositionDelta is a free data retrieval call binding the contract method 0x45a6f370.
//
// Solidity: function getPositionDelta(address _account, address _collateralToken, address _indexToken, bool _isLong) view returns(bool, uint256)
func (_Vault *VaultSession) GetPositionDelta(_account common.Address, _collateralToken common.Address, _indexToken common.Address, _isLong bool) (bool, *big.Int, error) {
	return _Vault.Contract.GetPositionDelta(&_Vault.CallOpts, _account, _collateralToken, _indexToken, _isLong)
}

// GetPositionDelta is a free data retrieval call binding the contract method 0x45a6f370.
//
// Solidity: function getPositionDelta(address _account, address _collateralToken, address _indexToken, bool _isLong) view returns(bool, uint256)
func (_Vault *VaultCallerSession) GetPositionDelta(_account common.Address, _collateralToken common.Address, _indexToken common.Address, _isLong bool) (bool, *big.Int, error) {
	return _Vault.Contract.GetPositionDelta(&_Vault.CallOpts, _account, _collateralToken, _indexToken, _isLong)
}

// GetPositionFee is a free data retrieval call binding the contract method 0xfdaf6ac3.
//
// Solidity: function getPositionFee(address _account, address _collateralToken, address _indexToken, bool _isLong, uint256 _sizeDelta) view returns(uint256)
func (_Vault *VaultCaller) GetPositionFee(opts *bind.CallOpts, _account common.Address, _collateralToken common.Address, _indexToken common.Address, _isLong bool, _sizeDelta *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "getPositionFee", _account, _collateralToken, _indexToken, _isLong, _sizeDelta)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPositionFee is a free data retrieval call binding the contract method 0xfdaf6ac3.
//
// Solidity: function getPositionFee(address _account, address _collateralToken, address _indexToken, bool _isLong, uint256 _sizeDelta) view returns(uint256)
func (_Vault *VaultSession) GetPositionFee(_account common.Address, _collateralToken common.Address, _indexToken common.Address, _isLong bool, _sizeDelta *big.Int) (*big.Int, error) {
	return _Vault.Contract.GetPositionFee(&_Vault.CallOpts, _account, _collateralToken, _indexToken, _isLong, _sizeDelta)
}

// GetPositionFee is a free data retrieval call binding the contract method 0xfdaf6ac3.
//
// Solidity: function getPositionFee(address _account, address _collateralToken, address _indexToken, bool _isLong, uint256 _sizeDelta) view returns(uint256)
func (_Vault *VaultCallerSession) GetPositionFee(_account common.Address, _collateralToken common.Address, _indexToken common.Address, _isLong bool, _sizeDelta *big.Int) (*big.Int, error) {
	return _Vault.Contract.GetPositionFee(&_Vault.CallOpts, _account, _collateralToken, _indexToken, _isLong, _sizeDelta)
}

// GetPositionKey is a free data retrieval call binding the contract method 0x2d4b0576.
//
// Solidity: function getPositionKey(address _account, address _collateralToken, address _indexToken, bool _isLong) pure returns(bytes32)
func (_Vault *VaultCaller) GetPositionKey(opts *bind.CallOpts, _account common.Address, _collateralToken common.Address, _indexToken common.Address, _isLong bool) ([32]byte, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "getPositionKey", _account, _collateralToken, _indexToken, _isLong)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetPositionKey is a free data retrieval call binding the contract method 0x2d4b0576.
//
// Solidity: function getPositionKey(address _account, address _collateralToken, address _indexToken, bool _isLong) pure returns(bytes32)
func (_Vault *VaultSession) GetPositionKey(_account common.Address, _collateralToken common.Address, _indexToken common.Address, _isLong bool) ([32]byte, error) {
	return _Vault.Contract.GetPositionKey(&_Vault.CallOpts, _account, _collateralToken, _indexToken, _isLong)
}

// GetPositionKey is a free data retrieval call binding the contract method 0x2d4b0576.
//
// Solidity: function getPositionKey(address _account, address _collateralToken, address _indexToken, bool _isLong) pure returns(bytes32)
func (_Vault *VaultCallerSession) GetPositionKey(_account common.Address, _collateralToken common.Address, _indexToken common.Address, _isLong bool) ([32]byte, error) {
	return _Vault.Contract.GetPositionKey(&_Vault.CallOpts, _account, _collateralToken, _indexToken, _isLong)
}

// GetPositionLeverage is a free data retrieval call binding the contract method 0x51723e82.
//
// Solidity: function getPositionLeverage(address _account, address _collateralToken, address _indexToken, bool _isLong) view returns(uint256)
func (_Vault *VaultCaller) GetPositionLeverage(opts *bind.CallOpts, _account common.Address, _collateralToken common.Address, _indexToken common.Address, _isLong bool) (*big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "getPositionLeverage", _account, _collateralToken, _indexToken, _isLong)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPositionLeverage is a free data retrieval call binding the contract method 0x51723e82.
//
// Solidity: function getPositionLeverage(address _account, address _collateralToken, address _indexToken, bool _isLong) view returns(uint256)
func (_Vault *VaultSession) GetPositionLeverage(_account common.Address, _collateralToken common.Address, _indexToken common.Address, _isLong bool) (*big.Int, error) {
	return _Vault.Contract.GetPositionLeverage(&_Vault.CallOpts, _account, _collateralToken, _indexToken, _isLong)
}

// GetPositionLeverage is a free data retrieval call binding the contract method 0x51723e82.
//
// Solidity: function getPositionLeverage(address _account, address _collateralToken, address _indexToken, bool _isLong) view returns(uint256)
func (_Vault *VaultCallerSession) GetPositionLeverage(_account common.Address, _collateralToken common.Address, _indexToken common.Address, _isLong bool) (*big.Int, error) {
	return _Vault.Contract.GetPositionLeverage(&_Vault.CallOpts, _account, _collateralToken, _indexToken, _isLong)
}

// GetRedemptionAmount is a free data retrieval call binding the contract method 0x2c668ec1.
//
// Solidity: function getRedemptionAmount(address _token, uint256 _usdgAmount) view returns(uint256)
func (_Vault *VaultCaller) GetRedemptionAmount(opts *bind.CallOpts, _token common.Address, _usdgAmount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "getRedemptionAmount", _token, _usdgAmount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRedemptionAmount is a free data retrieval call binding the contract method 0x2c668ec1.
//
// Solidity: function getRedemptionAmount(address _token, uint256 _usdgAmount) view returns(uint256)
func (_Vault *VaultSession) GetRedemptionAmount(_token common.Address, _usdgAmount *big.Int) (*big.Int, error) {
	return _Vault.Contract.GetRedemptionAmount(&_Vault.CallOpts, _token, _usdgAmount)
}

// GetRedemptionAmount is a free data retrieval call binding the contract method 0x2c668ec1.
//
// Solidity: function getRedemptionAmount(address _token, uint256 _usdgAmount) view returns(uint256)
func (_Vault *VaultCallerSession) GetRedemptionAmount(_token common.Address, _usdgAmount *big.Int) (*big.Int, error) {
	return _Vault.Contract.GetRedemptionAmount(&_Vault.CallOpts, _token, _usdgAmount)
}

// GetRedemptionCollateral is a free data retrieval call binding the contract method 0xb136ca49.
//
// Solidity: function getRedemptionCollateral(address _token) view returns(uint256)
func (_Vault *VaultCaller) GetRedemptionCollateral(opts *bind.CallOpts, _token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "getRedemptionCollateral", _token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRedemptionCollateral is a free data retrieval call binding the contract method 0xb136ca49.
//
// Solidity: function getRedemptionCollateral(address _token) view returns(uint256)
func (_Vault *VaultSession) GetRedemptionCollateral(_token common.Address) (*big.Int, error) {
	return _Vault.Contract.GetRedemptionCollateral(&_Vault.CallOpts, _token)
}

// GetRedemptionCollateral is a free data retrieval call binding the contract method 0xb136ca49.
//
// Solidity: function getRedemptionCollateral(address _token) view returns(uint256)
func (_Vault *VaultCallerSession) GetRedemptionCollateral(_token common.Address) (*big.Int, error) {
	return _Vault.Contract.GetRedemptionCollateral(&_Vault.CallOpts, _token)
}

// GetRedemptionCollateralUsd is a free data retrieval call binding the contract method 0x29ff9615.
//
// Solidity: function getRedemptionCollateralUsd(address _token) view returns(uint256)
func (_Vault *VaultCaller) GetRedemptionCollateralUsd(opts *bind.CallOpts, _token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "getRedemptionCollateralUsd", _token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRedemptionCollateralUsd is a free data retrieval call binding the contract method 0x29ff9615.
//
// Solidity: function getRedemptionCollateralUsd(address _token) view returns(uint256)
func (_Vault *VaultSession) GetRedemptionCollateralUsd(_token common.Address) (*big.Int, error) {
	return _Vault.Contract.GetRedemptionCollateralUsd(&_Vault.CallOpts, _token)
}

// GetRedemptionCollateralUsd is a free data retrieval call binding the contract method 0x29ff9615.
//
// Solidity: function getRedemptionCollateralUsd(address _token) view returns(uint256)
func (_Vault *VaultCallerSession) GetRedemptionCollateralUsd(_token common.Address) (*big.Int, error) {
	return _Vault.Contract.GetRedemptionCollateralUsd(&_Vault.CallOpts, _token)
}

// GetTargetUsdgAmount is a free data retrieval call binding the contract method 0x3a05dcc1.
//
// Solidity: function getTargetUsdgAmount(address _token) view returns(uint256)
func (_Vault *VaultCaller) GetTargetUsdgAmount(opts *bind.CallOpts, _token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "getTargetUsdgAmount", _token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTargetUsdgAmount is a free data retrieval call binding the contract method 0x3a05dcc1.
//
// Solidity: function getTargetUsdgAmount(address _token) view returns(uint256)
func (_Vault *VaultSession) GetTargetUsdgAmount(_token common.Address) (*big.Int, error) {
	return _Vault.Contract.GetTargetUsdgAmount(&_Vault.CallOpts, _token)
}

// GetTargetUsdgAmount is a free data retrieval call binding the contract method 0x3a05dcc1.
//
// Solidity: function getTargetUsdgAmount(address _token) view returns(uint256)
func (_Vault *VaultCallerSession) GetTargetUsdgAmount(_token common.Address) (*big.Int, error) {
	return _Vault.Contract.GetTargetUsdgAmount(&_Vault.CallOpts, _token)
}

// GetUtilisation is a free data retrieval call binding the contract method 0x04fef1db.
//
// Solidity: function getUtilisation(address _token) view returns(uint256)
func (_Vault *VaultCaller) GetUtilisation(opts *bind.CallOpts, _token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "getUtilisation", _token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUtilisation is a free data retrieval call binding the contract method 0x04fef1db.
//
// Solidity: function getUtilisation(address _token) view returns(uint256)
func (_Vault *VaultSession) GetUtilisation(_token common.Address) (*big.Int, error) {
	return _Vault.Contract.GetUtilisation(&_Vault.CallOpts, _token)
}

// GetUtilisation is a free data retrieval call binding the contract method 0x04fef1db.
//
// Solidity: function getUtilisation(address _token) view returns(uint256)
func (_Vault *VaultCallerSession) GetUtilisation(_token common.Address) (*big.Int, error) {
	return _Vault.Contract.GetUtilisation(&_Vault.CallOpts, _token)
}

// GlobalShortAveragePrices is a free data retrieval call binding the contract method 0x62749803.
//
// Solidity: function globalShortAveragePrices(address ) view returns(uint256)
func (_Vault *VaultCaller) GlobalShortAveragePrices(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "globalShortAveragePrices", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GlobalShortAveragePrices is a free data retrieval call binding the contract method 0x62749803.
//
// Solidity: function globalShortAveragePrices(address ) view returns(uint256)
func (_Vault *VaultSession) GlobalShortAveragePrices(arg0 common.Address) (*big.Int, error) {
	return _Vault.Contract.GlobalShortAveragePrices(&_Vault.CallOpts, arg0)
}

// GlobalShortAveragePrices is a free data retrieval call binding the contract method 0x62749803.
//
// Solidity: function globalShortAveragePrices(address ) view returns(uint256)
func (_Vault *VaultCallerSession) GlobalShortAveragePrices(arg0 common.Address) (*big.Int, error) {
	return _Vault.Contract.GlobalShortAveragePrices(&_Vault.CallOpts, arg0)
}

// GlobalShortSizes is a free data retrieval call binding the contract method 0x8a78daa8.
//
// Solidity: function globalShortSizes(address ) view returns(uint256)
func (_Vault *VaultCaller) GlobalShortSizes(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "globalShortSizes", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GlobalShortSizes is a free data retrieval call binding the contract method 0x8a78daa8.
//
// Solidity: function globalShortSizes(address ) view returns(uint256)
func (_Vault *VaultSession) GlobalShortSizes(arg0 common.Address) (*big.Int, error) {
	return _Vault.Contract.GlobalShortSizes(&_Vault.CallOpts, arg0)
}

// GlobalShortSizes is a free data retrieval call binding the contract method 0x8a78daa8.
//
// Solidity: function globalShortSizes(address ) view returns(uint256)
func (_Vault *VaultCallerSession) GlobalShortSizes(arg0 common.Address) (*big.Int, error) {
	return _Vault.Contract.GlobalShortSizes(&_Vault.CallOpts, arg0)
}

// Gov is a free data retrieval call binding the contract method 0x12d43a51.
//
// Solidity: function gov() view returns(address)
func (_Vault *VaultCaller) Gov(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "gov")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Gov is a free data retrieval call binding the contract method 0x12d43a51.
//
// Solidity: function gov() view returns(address)
func (_Vault *VaultSession) Gov() (common.Address, error) {
	return _Vault.Contract.Gov(&_Vault.CallOpts)
}

// Gov is a free data retrieval call binding the contract method 0x12d43a51.
//
// Solidity: function gov() view returns(address)
func (_Vault *VaultCallerSession) Gov() (common.Address, error) {
	return _Vault.Contract.Gov(&_Vault.CallOpts)
}

// GuaranteedUsd is a free data retrieval call binding the contract method 0xf07456ce.
//
// Solidity: function guaranteedUsd(address ) view returns(uint256)
func (_Vault *VaultCaller) GuaranteedUsd(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "guaranteedUsd", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GuaranteedUsd is a free data retrieval call binding the contract method 0xf07456ce.
//
// Solidity: function guaranteedUsd(address ) view returns(uint256)
func (_Vault *VaultSession) GuaranteedUsd(arg0 common.Address) (*big.Int, error) {
	return _Vault.Contract.GuaranteedUsd(&_Vault.CallOpts, arg0)
}

// GuaranteedUsd is a free data retrieval call binding the contract method 0xf07456ce.
//
// Solidity: function guaranteedUsd(address ) view returns(uint256)
func (_Vault *VaultCallerSession) GuaranteedUsd(arg0 common.Address) (*big.Int, error) {
	return _Vault.Contract.GuaranteedUsd(&_Vault.CallOpts, arg0)
}

// HasDynamicFees is a free data retrieval call binding the contract method 0x9f392eb3.
//
// Solidity: function hasDynamicFees() view returns(bool)
func (_Vault *VaultCaller) HasDynamicFees(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "hasDynamicFees")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasDynamicFees is a free data retrieval call binding the contract method 0x9f392eb3.
//
// Solidity: function hasDynamicFees() view returns(bool)
func (_Vault *VaultSession) HasDynamicFees() (bool, error) {
	return _Vault.Contract.HasDynamicFees(&_Vault.CallOpts)
}

// HasDynamicFees is a free data retrieval call binding the contract method 0x9f392eb3.
//
// Solidity: function hasDynamicFees() view returns(bool)
func (_Vault *VaultCallerSession) HasDynamicFees() (bool, error) {
	return _Vault.Contract.HasDynamicFees(&_Vault.CallOpts)
}

// InManagerMode is a free data retrieval call binding the contract method 0x9060b1ca.
//
// Solidity: function inManagerMode() view returns(bool)
func (_Vault *VaultCaller) InManagerMode(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "inManagerMode")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// InManagerMode is a free data retrieval call binding the contract method 0x9060b1ca.
//
// Solidity: function inManagerMode() view returns(bool)
func (_Vault *VaultSession) InManagerMode() (bool, error) {
	return _Vault.Contract.InManagerMode(&_Vault.CallOpts)
}

// InManagerMode is a free data retrieval call binding the contract method 0x9060b1ca.
//
// Solidity: function inManagerMode() view returns(bool)
func (_Vault *VaultCallerSession) InManagerMode() (bool, error) {
	return _Vault.Contract.InManagerMode(&_Vault.CallOpts)
}

// InPrivateLiquidationMode is a free data retrieval call binding the contract method 0x181e210e.
//
// Solidity: function inPrivateLiquidationMode() view returns(bool)
func (_Vault *VaultCaller) InPrivateLiquidationMode(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "inPrivateLiquidationMode")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// InPrivateLiquidationMode is a free data retrieval call binding the contract method 0x181e210e.
//
// Solidity: function inPrivateLiquidationMode() view returns(bool)
func (_Vault *VaultSession) InPrivateLiquidationMode() (bool, error) {
	return _Vault.Contract.InPrivateLiquidationMode(&_Vault.CallOpts)
}

// InPrivateLiquidationMode is a free data retrieval call binding the contract method 0x181e210e.
//
// Solidity: function inPrivateLiquidationMode() view returns(bool)
func (_Vault *VaultCallerSession) InPrivateLiquidationMode() (bool, error) {
	return _Vault.Contract.InPrivateLiquidationMode(&_Vault.CallOpts)
}

// IncludeAmmPrice is a free data retrieval call binding the contract method 0xab08c1c6.
//
// Solidity: function includeAmmPrice() view returns(bool)
func (_Vault *VaultCaller) IncludeAmmPrice(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "includeAmmPrice")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IncludeAmmPrice is a free data retrieval call binding the contract method 0xab08c1c6.
//
// Solidity: function includeAmmPrice() view returns(bool)
func (_Vault *VaultSession) IncludeAmmPrice() (bool, error) {
	return _Vault.Contract.IncludeAmmPrice(&_Vault.CallOpts)
}

// IncludeAmmPrice is a free data retrieval call binding the contract method 0xab08c1c6.
//
// Solidity: function includeAmmPrice() view returns(bool)
func (_Vault *VaultCallerSession) IncludeAmmPrice() (bool, error) {
	return _Vault.Contract.IncludeAmmPrice(&_Vault.CallOpts)
}

// IsInitialized is a free data retrieval call binding the contract method 0x392e53cd.
//
// Solidity: function isInitialized() view returns(bool)
func (_Vault *VaultCaller) IsInitialized(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "isInitialized")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsInitialized is a free data retrieval call binding the contract method 0x392e53cd.
//
// Solidity: function isInitialized() view returns(bool)
func (_Vault *VaultSession) IsInitialized() (bool, error) {
	return _Vault.Contract.IsInitialized(&_Vault.CallOpts)
}

// IsInitialized is a free data retrieval call binding the contract method 0x392e53cd.
//
// Solidity: function isInitialized() view returns(bool)
func (_Vault *VaultCallerSession) IsInitialized() (bool, error) {
	return _Vault.Contract.IsInitialized(&_Vault.CallOpts)
}

// IsLeverageEnabled is a free data retrieval call binding the contract method 0x3e72a262.
//
// Solidity: function isLeverageEnabled() view returns(bool)
func (_Vault *VaultCaller) IsLeverageEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "isLeverageEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsLeverageEnabled is a free data retrieval call binding the contract method 0x3e72a262.
//
// Solidity: function isLeverageEnabled() view returns(bool)
func (_Vault *VaultSession) IsLeverageEnabled() (bool, error) {
	return _Vault.Contract.IsLeverageEnabled(&_Vault.CallOpts)
}

// IsLeverageEnabled is a free data retrieval call binding the contract method 0x3e72a262.
//
// Solidity: function isLeverageEnabled() view returns(bool)
func (_Vault *VaultCallerSession) IsLeverageEnabled() (bool, error) {
	return _Vault.Contract.IsLeverageEnabled(&_Vault.CallOpts)
}

// IsLiquidator is a free data retrieval call binding the contract method 0x529a356f.
//
// Solidity: function isLiquidator(address ) view returns(bool)
func (_Vault *VaultCaller) IsLiquidator(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "isLiquidator", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsLiquidator is a free data retrieval call binding the contract method 0x529a356f.
//
// Solidity: function isLiquidator(address ) view returns(bool)
func (_Vault *VaultSession) IsLiquidator(arg0 common.Address) (bool, error) {
	return _Vault.Contract.IsLiquidator(&_Vault.CallOpts, arg0)
}

// IsLiquidator is a free data retrieval call binding the contract method 0x529a356f.
//
// Solidity: function isLiquidator(address ) view returns(bool)
func (_Vault *VaultCallerSession) IsLiquidator(arg0 common.Address) (bool, error) {
	return _Vault.Contract.IsLiquidator(&_Vault.CallOpts, arg0)
}

// IsManager is a free data retrieval call binding the contract method 0xf3ae2415.
//
// Solidity: function isManager(address ) view returns(bool)
func (_Vault *VaultCaller) IsManager(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "isManager", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsManager is a free data retrieval call binding the contract method 0xf3ae2415.
//
// Solidity: function isManager(address ) view returns(bool)
func (_Vault *VaultSession) IsManager(arg0 common.Address) (bool, error) {
	return _Vault.Contract.IsManager(&_Vault.CallOpts, arg0)
}

// IsManager is a free data retrieval call binding the contract method 0xf3ae2415.
//
// Solidity: function isManager(address ) view returns(bool)
func (_Vault *VaultCallerSession) IsManager(arg0 common.Address) (bool, error) {
	return _Vault.Contract.IsManager(&_Vault.CallOpts, arg0)
}

// IsSwapEnabled is a free data retrieval call binding the contract method 0x351a964d.
//
// Solidity: function isSwapEnabled() view returns(bool)
func (_Vault *VaultCaller) IsSwapEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "isSwapEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSwapEnabled is a free data retrieval call binding the contract method 0x351a964d.
//
// Solidity: function isSwapEnabled() view returns(bool)
func (_Vault *VaultSession) IsSwapEnabled() (bool, error) {
	return _Vault.Contract.IsSwapEnabled(&_Vault.CallOpts)
}

// IsSwapEnabled is a free data retrieval call binding the contract method 0x351a964d.
//
// Solidity: function isSwapEnabled() view returns(bool)
func (_Vault *VaultCallerSession) IsSwapEnabled() (bool, error) {
	return _Vault.Contract.IsSwapEnabled(&_Vault.CallOpts)
}

// LastFundingTimes is a free data retrieval call binding the contract method 0xd8f897c3.
//
// Solidity: function lastFundingTimes(address ) view returns(uint256)
func (_Vault *VaultCaller) LastFundingTimes(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "lastFundingTimes", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastFundingTimes is a free data retrieval call binding the contract method 0xd8f897c3.
//
// Solidity: function lastFundingTimes(address ) view returns(uint256)
func (_Vault *VaultSession) LastFundingTimes(arg0 common.Address) (*big.Int, error) {
	return _Vault.Contract.LastFundingTimes(&_Vault.CallOpts, arg0)
}

// LastFundingTimes is a free data retrieval call binding the contract method 0xd8f897c3.
//
// Solidity: function lastFundingTimes(address ) view returns(uint256)
func (_Vault *VaultCallerSession) LastFundingTimes(arg0 common.Address) (*big.Int, error) {
	return _Vault.Contract.LastFundingTimes(&_Vault.CallOpts, arg0)
}

// LiquidationFeeUsd is a free data retrieval call binding the contract method 0x174d2694.
//
// Solidity: function liquidationFeeUsd() view returns(uint256)
func (_Vault *VaultCaller) LiquidationFeeUsd(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "liquidationFeeUsd")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LiquidationFeeUsd is a free data retrieval call binding the contract method 0x174d2694.
//
// Solidity: function liquidationFeeUsd() view returns(uint256)
func (_Vault *VaultSession) LiquidationFeeUsd() (*big.Int, error) {
	return _Vault.Contract.LiquidationFeeUsd(&_Vault.CallOpts)
}

// LiquidationFeeUsd is a free data retrieval call binding the contract method 0x174d2694.
//
// Solidity: function liquidationFeeUsd() view returns(uint256)
func (_Vault *VaultCallerSession) LiquidationFeeUsd() (*big.Int, error) {
	return _Vault.Contract.LiquidationFeeUsd(&_Vault.CallOpts)
}

// MarginFeeBasisPoints is a free data retrieval call binding the contract method 0x318bc689.
//
// Solidity: function marginFeeBasisPoints() view returns(uint256)
func (_Vault *VaultCaller) MarginFeeBasisPoints(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "marginFeeBasisPoints")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MarginFeeBasisPoints is a free data retrieval call binding the contract method 0x318bc689.
//
// Solidity: function marginFeeBasisPoints() view returns(uint256)
func (_Vault *VaultSession) MarginFeeBasisPoints() (*big.Int, error) {
	return _Vault.Contract.MarginFeeBasisPoints(&_Vault.CallOpts)
}

// MarginFeeBasisPoints is a free data retrieval call binding the contract method 0x318bc689.
//
// Solidity: function marginFeeBasisPoints() view returns(uint256)
func (_Vault *VaultCallerSession) MarginFeeBasisPoints() (*big.Int, error) {
	return _Vault.Contract.MarginFeeBasisPoints(&_Vault.CallOpts)
}

// MaxGasPrice is a free data retrieval call binding the contract method 0x3de39c11.
//
// Solidity: function maxGasPrice() view returns(uint256)
func (_Vault *VaultCaller) MaxGasPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "maxGasPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxGasPrice is a free data retrieval call binding the contract method 0x3de39c11.
//
// Solidity: function maxGasPrice() view returns(uint256)
func (_Vault *VaultSession) MaxGasPrice() (*big.Int, error) {
	return _Vault.Contract.MaxGasPrice(&_Vault.CallOpts)
}

// MaxGasPrice is a free data retrieval call binding the contract method 0x3de39c11.
//
// Solidity: function maxGasPrice() view returns(uint256)
func (_Vault *VaultCallerSession) MaxGasPrice() (*big.Int, error) {
	return _Vault.Contract.MaxGasPrice(&_Vault.CallOpts)
}

// MaxGlobalShortSizes is a free data retrieval call binding the contract method 0x9698d25a.
//
// Solidity: function maxGlobalShortSizes(address ) view returns(uint256)
func (_Vault *VaultCaller) MaxGlobalShortSizes(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "maxGlobalShortSizes", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxGlobalShortSizes is a free data retrieval call binding the contract method 0x9698d25a.
//
// Solidity: function maxGlobalShortSizes(address ) view returns(uint256)
func (_Vault *VaultSession) MaxGlobalShortSizes(arg0 common.Address) (*big.Int, error) {
	return _Vault.Contract.MaxGlobalShortSizes(&_Vault.CallOpts, arg0)
}

// MaxGlobalShortSizes is a free data retrieval call binding the contract method 0x9698d25a.
//
// Solidity: function maxGlobalShortSizes(address ) view returns(uint256)
func (_Vault *VaultCallerSession) MaxGlobalShortSizes(arg0 common.Address) (*big.Int, error) {
	return _Vault.Contract.MaxGlobalShortSizes(&_Vault.CallOpts, arg0)
}

// MaxLeverage is a free data retrieval call binding the contract method 0xae3302c2.
//
// Solidity: function maxLeverage() view returns(uint256)
func (_Vault *VaultCaller) MaxLeverage(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "maxLeverage")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxLeverage is a free data retrieval call binding the contract method 0xae3302c2.
//
// Solidity: function maxLeverage() view returns(uint256)
func (_Vault *VaultSession) MaxLeverage() (*big.Int, error) {
	return _Vault.Contract.MaxLeverage(&_Vault.CallOpts)
}

// MaxLeverage is a free data retrieval call binding the contract method 0xae3302c2.
//
// Solidity: function maxLeverage() view returns(uint256)
func (_Vault *VaultCallerSession) MaxLeverage() (*big.Int, error) {
	return _Vault.Contract.MaxLeverage(&_Vault.CallOpts)
}

// MaxUsdgAmounts is a free data retrieval call binding the contract method 0xad1e4f8d.
//
// Solidity: function maxUsdgAmounts(address ) view returns(uint256)
func (_Vault *VaultCaller) MaxUsdgAmounts(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "maxUsdgAmounts", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxUsdgAmounts is a free data retrieval call binding the contract method 0xad1e4f8d.
//
// Solidity: function maxUsdgAmounts(address ) view returns(uint256)
func (_Vault *VaultSession) MaxUsdgAmounts(arg0 common.Address) (*big.Int, error) {
	return _Vault.Contract.MaxUsdgAmounts(&_Vault.CallOpts, arg0)
}

// MaxUsdgAmounts is a free data retrieval call binding the contract method 0xad1e4f8d.
//
// Solidity: function maxUsdgAmounts(address ) view returns(uint256)
func (_Vault *VaultCallerSession) MaxUsdgAmounts(arg0 common.Address) (*big.Int, error) {
	return _Vault.Contract.MaxUsdgAmounts(&_Vault.CallOpts, arg0)
}

// MinProfitBasisPoints is a free data retrieval call binding the contract method 0x88b1fbdf.
//
// Solidity: function minProfitBasisPoints(address ) view returns(uint256)
func (_Vault *VaultCaller) MinProfitBasisPoints(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "minProfitBasisPoints", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinProfitBasisPoints is a free data retrieval call binding the contract method 0x88b1fbdf.
//
// Solidity: function minProfitBasisPoints(address ) view returns(uint256)
func (_Vault *VaultSession) MinProfitBasisPoints(arg0 common.Address) (*big.Int, error) {
	return _Vault.Contract.MinProfitBasisPoints(&_Vault.CallOpts, arg0)
}

// MinProfitBasisPoints is a free data retrieval call binding the contract method 0x88b1fbdf.
//
// Solidity: function minProfitBasisPoints(address ) view returns(uint256)
func (_Vault *VaultCallerSession) MinProfitBasisPoints(arg0 common.Address) (*big.Int, error) {
	return _Vault.Contract.MinProfitBasisPoints(&_Vault.CallOpts, arg0)
}

// MinProfitTime is a free data retrieval call binding the contract method 0xd9ac4225.
//
// Solidity: function minProfitTime() view returns(uint256)
func (_Vault *VaultCaller) MinProfitTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "minProfitTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinProfitTime is a free data retrieval call binding the contract method 0xd9ac4225.
//
// Solidity: function minProfitTime() view returns(uint256)
func (_Vault *VaultSession) MinProfitTime() (*big.Int, error) {
	return _Vault.Contract.MinProfitTime(&_Vault.CallOpts)
}

// MinProfitTime is a free data retrieval call binding the contract method 0xd9ac4225.
//
// Solidity: function minProfitTime() view returns(uint256)
func (_Vault *VaultCallerSession) MinProfitTime() (*big.Int, error) {
	return _Vault.Contract.MinProfitTime(&_Vault.CallOpts)
}

// MintBurnFeeBasisPoints is a free data retrieval call binding the contract method 0x4d47b304.
//
// Solidity: function mintBurnFeeBasisPoints() view returns(uint256)
func (_Vault *VaultCaller) MintBurnFeeBasisPoints(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "mintBurnFeeBasisPoints")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MintBurnFeeBasisPoints is a free data retrieval call binding the contract method 0x4d47b304.
//
// Solidity: function mintBurnFeeBasisPoints() view returns(uint256)
func (_Vault *VaultSession) MintBurnFeeBasisPoints() (*big.Int, error) {
	return _Vault.Contract.MintBurnFeeBasisPoints(&_Vault.CallOpts)
}

// MintBurnFeeBasisPoints is a free data retrieval call binding the contract method 0x4d47b304.
//
// Solidity: function mintBurnFeeBasisPoints() view returns(uint256)
func (_Vault *VaultCallerSession) MintBurnFeeBasisPoints() (*big.Int, error) {
	return _Vault.Contract.MintBurnFeeBasisPoints(&_Vault.CallOpts)
}

// PoolAmounts is a free data retrieval call binding the contract method 0x52f55eed.
//
// Solidity: function poolAmounts(address ) view returns(uint256)
func (_Vault *VaultCaller) PoolAmounts(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "poolAmounts", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PoolAmounts is a free data retrieval call binding the contract method 0x52f55eed.
//
// Solidity: function poolAmounts(address ) view returns(uint256)
func (_Vault *VaultSession) PoolAmounts(arg0 common.Address) (*big.Int, error) {
	return _Vault.Contract.PoolAmounts(&_Vault.CallOpts, arg0)
}

// PoolAmounts is a free data retrieval call binding the contract method 0x52f55eed.
//
// Solidity: function poolAmounts(address ) view returns(uint256)
func (_Vault *VaultCallerSession) PoolAmounts(arg0 common.Address) (*big.Int, error) {
	return _Vault.Contract.PoolAmounts(&_Vault.CallOpts, arg0)
}

// Positions is a free data retrieval call binding the contract method 0x514ea4bf.
//
// Solidity: function positions(bytes32 ) view returns(uint256 size, uint256 collateral, uint256 averagePrice, uint256 entryFundingRate, uint256 reserveAmount, int256 realisedPnl, uint256 lastIncreasedTime)
func (_Vault *VaultCaller) Positions(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Size              *big.Int
	Collateral        *big.Int
	AveragePrice      *big.Int
	EntryFundingRate  *big.Int
	ReserveAmount     *big.Int
	RealisedPnl       *big.Int
	LastIncreasedTime *big.Int
}, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "positions", arg0)

	outstruct := new(struct {
		Size              *big.Int
		Collateral        *big.Int
		AveragePrice      *big.Int
		EntryFundingRate  *big.Int
		ReserveAmount     *big.Int
		RealisedPnl       *big.Int
		LastIncreasedTime *big.Int
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

	return *outstruct, err

}

// Positions is a free data retrieval call binding the contract method 0x514ea4bf.
//
// Solidity: function positions(bytes32 ) view returns(uint256 size, uint256 collateral, uint256 averagePrice, uint256 entryFundingRate, uint256 reserveAmount, int256 realisedPnl, uint256 lastIncreasedTime)
func (_Vault *VaultSession) Positions(arg0 [32]byte) (struct {
	Size              *big.Int
	Collateral        *big.Int
	AveragePrice      *big.Int
	EntryFundingRate  *big.Int
	ReserveAmount     *big.Int
	RealisedPnl       *big.Int
	LastIncreasedTime *big.Int
}, error) {
	return _Vault.Contract.Positions(&_Vault.CallOpts, arg0)
}

// Positions is a free data retrieval call binding the contract method 0x514ea4bf.
//
// Solidity: function positions(bytes32 ) view returns(uint256 size, uint256 collateral, uint256 averagePrice, uint256 entryFundingRate, uint256 reserveAmount, int256 realisedPnl, uint256 lastIncreasedTime)
func (_Vault *VaultCallerSession) Positions(arg0 [32]byte) (struct {
	Size              *big.Int
	Collateral        *big.Int
	AveragePrice      *big.Int
	EntryFundingRate  *big.Int
	ReserveAmount     *big.Int
	RealisedPnl       *big.Int
	LastIncreasedTime *big.Int
}, error) {
	return _Vault.Contract.Positions(&_Vault.CallOpts, arg0)
}

// PriceFeed is a free data retrieval call binding the contract method 0x741bef1a.
//
// Solidity: function priceFeed() view returns(address)
func (_Vault *VaultCaller) PriceFeed(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "priceFeed")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PriceFeed is a free data retrieval call binding the contract method 0x741bef1a.
//
// Solidity: function priceFeed() view returns(address)
func (_Vault *VaultSession) PriceFeed() (common.Address, error) {
	return _Vault.Contract.PriceFeed(&_Vault.CallOpts)
}

// PriceFeed is a free data retrieval call binding the contract method 0x741bef1a.
//
// Solidity: function priceFeed() view returns(address)
func (_Vault *VaultCallerSession) PriceFeed() (common.Address, error) {
	return _Vault.Contract.PriceFeed(&_Vault.CallOpts)
}

// ReservedAmounts is a free data retrieval call binding the contract method 0xc3c7b9e9.
//
// Solidity: function reservedAmounts(address ) view returns(uint256)
func (_Vault *VaultCaller) ReservedAmounts(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "reservedAmounts", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ReservedAmounts is a free data retrieval call binding the contract method 0xc3c7b9e9.
//
// Solidity: function reservedAmounts(address ) view returns(uint256)
func (_Vault *VaultSession) ReservedAmounts(arg0 common.Address) (*big.Int, error) {
	return _Vault.Contract.ReservedAmounts(&_Vault.CallOpts, arg0)
}

// ReservedAmounts is a free data retrieval call binding the contract method 0xc3c7b9e9.
//
// Solidity: function reservedAmounts(address ) view returns(uint256)
func (_Vault *VaultCallerSession) ReservedAmounts(arg0 common.Address) (*big.Int, error) {
	return _Vault.Contract.ReservedAmounts(&_Vault.CallOpts, arg0)
}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_Vault *VaultCaller) Router(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "router")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_Vault *VaultSession) Router() (common.Address, error) {
	return _Vault.Contract.Router(&_Vault.CallOpts)
}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_Vault *VaultCallerSession) Router() (common.Address, error) {
	return _Vault.Contract.Router(&_Vault.CallOpts)
}

// ShortableTokens is a free data retrieval call binding the contract method 0xdb3555fb.
//
// Solidity: function shortableTokens(address ) view returns(bool)
func (_Vault *VaultCaller) ShortableTokens(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "shortableTokens", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ShortableTokens is a free data retrieval call binding the contract method 0xdb3555fb.
//
// Solidity: function shortableTokens(address ) view returns(bool)
func (_Vault *VaultSession) ShortableTokens(arg0 common.Address) (bool, error) {
	return _Vault.Contract.ShortableTokens(&_Vault.CallOpts, arg0)
}

// ShortableTokens is a free data retrieval call binding the contract method 0xdb3555fb.
//
// Solidity: function shortableTokens(address ) view returns(bool)
func (_Vault *VaultCallerSession) ShortableTokens(arg0 common.Address) (bool, error) {
	return _Vault.Contract.ShortableTokens(&_Vault.CallOpts, arg0)
}

// StableFundingRateFactor is a free data retrieval call binding the contract method 0x134ca63b.
//
// Solidity: function stableFundingRateFactor() view returns(uint256)
func (_Vault *VaultCaller) StableFundingRateFactor(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "stableFundingRateFactor")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StableFundingRateFactor is a free data retrieval call binding the contract method 0x134ca63b.
//
// Solidity: function stableFundingRateFactor() view returns(uint256)
func (_Vault *VaultSession) StableFundingRateFactor() (*big.Int, error) {
	return _Vault.Contract.StableFundingRateFactor(&_Vault.CallOpts)
}

// StableFundingRateFactor is a free data retrieval call binding the contract method 0x134ca63b.
//
// Solidity: function stableFundingRateFactor() view returns(uint256)
func (_Vault *VaultCallerSession) StableFundingRateFactor() (*big.Int, error) {
	return _Vault.Contract.StableFundingRateFactor(&_Vault.CallOpts)
}

// StableSwapFeeBasisPoints is a free data retrieval call binding the contract method 0xdf73a267.
//
// Solidity: function stableSwapFeeBasisPoints() view returns(uint256)
func (_Vault *VaultCaller) StableSwapFeeBasisPoints(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "stableSwapFeeBasisPoints")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StableSwapFeeBasisPoints is a free data retrieval call binding the contract method 0xdf73a267.
//
// Solidity: function stableSwapFeeBasisPoints() view returns(uint256)
func (_Vault *VaultSession) StableSwapFeeBasisPoints() (*big.Int, error) {
	return _Vault.Contract.StableSwapFeeBasisPoints(&_Vault.CallOpts)
}

// StableSwapFeeBasisPoints is a free data retrieval call binding the contract method 0xdf73a267.
//
// Solidity: function stableSwapFeeBasisPoints() view returns(uint256)
func (_Vault *VaultCallerSession) StableSwapFeeBasisPoints() (*big.Int, error) {
	return _Vault.Contract.StableSwapFeeBasisPoints(&_Vault.CallOpts)
}

// StableTaxBasisPoints is a free data retrieval call binding the contract method 0x10eb56c2.
//
// Solidity: function stableTaxBasisPoints() view returns(uint256)
func (_Vault *VaultCaller) StableTaxBasisPoints(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "stableTaxBasisPoints")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StableTaxBasisPoints is a free data retrieval call binding the contract method 0x10eb56c2.
//
// Solidity: function stableTaxBasisPoints() view returns(uint256)
func (_Vault *VaultSession) StableTaxBasisPoints() (*big.Int, error) {
	return _Vault.Contract.StableTaxBasisPoints(&_Vault.CallOpts)
}

// StableTaxBasisPoints is a free data retrieval call binding the contract method 0x10eb56c2.
//
// Solidity: function stableTaxBasisPoints() view returns(uint256)
func (_Vault *VaultCallerSession) StableTaxBasisPoints() (*big.Int, error) {
	return _Vault.Contract.StableTaxBasisPoints(&_Vault.CallOpts)
}

// StableTokens is a free data retrieval call binding the contract method 0x42b60b03.
//
// Solidity: function stableTokens(address ) view returns(bool)
func (_Vault *VaultCaller) StableTokens(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "stableTokens", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// StableTokens is a free data retrieval call binding the contract method 0x42b60b03.
//
// Solidity: function stableTokens(address ) view returns(bool)
func (_Vault *VaultSession) StableTokens(arg0 common.Address) (bool, error) {
	return _Vault.Contract.StableTokens(&_Vault.CallOpts, arg0)
}

// StableTokens is a free data retrieval call binding the contract method 0x42b60b03.
//
// Solidity: function stableTokens(address ) view returns(bool)
func (_Vault *VaultCallerSession) StableTokens(arg0 common.Address) (bool, error) {
	return _Vault.Contract.StableTokens(&_Vault.CallOpts, arg0)
}

// SwapFeeBasisPoints is a free data retrieval call binding the contract method 0xa22f2392.
//
// Solidity: function swapFeeBasisPoints() view returns(uint256)
func (_Vault *VaultCaller) SwapFeeBasisPoints(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "swapFeeBasisPoints")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SwapFeeBasisPoints is a free data retrieval call binding the contract method 0xa22f2392.
//
// Solidity: function swapFeeBasisPoints() view returns(uint256)
func (_Vault *VaultSession) SwapFeeBasisPoints() (*big.Int, error) {
	return _Vault.Contract.SwapFeeBasisPoints(&_Vault.CallOpts)
}

// SwapFeeBasisPoints is a free data retrieval call binding the contract method 0xa22f2392.
//
// Solidity: function swapFeeBasisPoints() view returns(uint256)
func (_Vault *VaultCallerSession) SwapFeeBasisPoints() (*big.Int, error) {
	return _Vault.Contract.SwapFeeBasisPoints(&_Vault.CallOpts)
}

// TaxBasisPoints is a free data retrieval call binding the contract method 0x7a210a2b.
//
// Solidity: function taxBasisPoints() view returns(uint256)
func (_Vault *VaultCaller) TaxBasisPoints(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "taxBasisPoints")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TaxBasisPoints is a free data retrieval call binding the contract method 0x7a210a2b.
//
// Solidity: function taxBasisPoints() view returns(uint256)
func (_Vault *VaultSession) TaxBasisPoints() (*big.Int, error) {
	return _Vault.Contract.TaxBasisPoints(&_Vault.CallOpts)
}

// TaxBasisPoints is a free data retrieval call binding the contract method 0x7a210a2b.
//
// Solidity: function taxBasisPoints() view returns(uint256)
func (_Vault *VaultCallerSession) TaxBasisPoints() (*big.Int, error) {
	return _Vault.Contract.TaxBasisPoints(&_Vault.CallOpts)
}

// TokenBalances is a free data retrieval call binding the contract method 0x523fba7f.
//
// Solidity: function tokenBalances(address ) view returns(uint256)
func (_Vault *VaultCaller) TokenBalances(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "tokenBalances", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenBalances is a free data retrieval call binding the contract method 0x523fba7f.
//
// Solidity: function tokenBalances(address ) view returns(uint256)
func (_Vault *VaultSession) TokenBalances(arg0 common.Address) (*big.Int, error) {
	return _Vault.Contract.TokenBalances(&_Vault.CallOpts, arg0)
}

// TokenBalances is a free data retrieval call binding the contract method 0x523fba7f.
//
// Solidity: function tokenBalances(address ) view returns(uint256)
func (_Vault *VaultCallerSession) TokenBalances(arg0 common.Address) (*big.Int, error) {
	return _Vault.Contract.TokenBalances(&_Vault.CallOpts, arg0)
}

// TokenDecimals is a free data retrieval call binding the contract method 0x8ee573ac.
//
// Solidity: function tokenDecimals(address ) view returns(uint256)
func (_Vault *VaultCaller) TokenDecimals(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "tokenDecimals", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenDecimals is a free data retrieval call binding the contract method 0x8ee573ac.
//
// Solidity: function tokenDecimals(address ) view returns(uint256)
func (_Vault *VaultSession) TokenDecimals(arg0 common.Address) (*big.Int, error) {
	return _Vault.Contract.TokenDecimals(&_Vault.CallOpts, arg0)
}

// TokenDecimals is a free data retrieval call binding the contract method 0x8ee573ac.
//
// Solidity: function tokenDecimals(address ) view returns(uint256)
func (_Vault *VaultCallerSession) TokenDecimals(arg0 common.Address) (*big.Int, error) {
	return _Vault.Contract.TokenDecimals(&_Vault.CallOpts, arg0)
}

// TokenToUsdMin is a free data retrieval call binding the contract method 0x0a48d5a9.
//
// Solidity: function tokenToUsdMin(address _token, uint256 _tokenAmount) view returns(uint256)
func (_Vault *VaultCaller) TokenToUsdMin(opts *bind.CallOpts, _token common.Address, _tokenAmount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "tokenToUsdMin", _token, _tokenAmount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenToUsdMin is a free data retrieval call binding the contract method 0x0a48d5a9.
//
// Solidity: function tokenToUsdMin(address _token, uint256 _tokenAmount) view returns(uint256)
func (_Vault *VaultSession) TokenToUsdMin(_token common.Address, _tokenAmount *big.Int) (*big.Int, error) {
	return _Vault.Contract.TokenToUsdMin(&_Vault.CallOpts, _token, _tokenAmount)
}

// TokenToUsdMin is a free data retrieval call binding the contract method 0x0a48d5a9.
//
// Solidity: function tokenToUsdMin(address _token, uint256 _tokenAmount) view returns(uint256)
func (_Vault *VaultCallerSession) TokenToUsdMin(_token common.Address, _tokenAmount *big.Int) (*big.Int, error) {
	return _Vault.Contract.TokenToUsdMin(&_Vault.CallOpts, _token, _tokenAmount)
}

// TokenWeights is a free data retrieval call binding the contract method 0xab2f3ad4.
//
// Solidity: function tokenWeights(address ) view returns(uint256)
func (_Vault *VaultCaller) TokenWeights(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "tokenWeights", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenWeights is a free data retrieval call binding the contract method 0xab2f3ad4.
//
// Solidity: function tokenWeights(address ) view returns(uint256)
func (_Vault *VaultSession) TokenWeights(arg0 common.Address) (*big.Int, error) {
	return _Vault.Contract.TokenWeights(&_Vault.CallOpts, arg0)
}

// TokenWeights is a free data retrieval call binding the contract method 0xab2f3ad4.
//
// Solidity: function tokenWeights(address ) view returns(uint256)
func (_Vault *VaultCallerSession) TokenWeights(arg0 common.Address) (*big.Int, error) {
	return _Vault.Contract.TokenWeights(&_Vault.CallOpts, arg0)
}

// TotalTokenWeights is a free data retrieval call binding the contract method 0xdc8f5fac.
//
// Solidity: function totalTokenWeights() view returns(uint256)
func (_Vault *VaultCaller) TotalTokenWeights(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "totalTokenWeights")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalTokenWeights is a free data retrieval call binding the contract method 0xdc8f5fac.
//
// Solidity: function totalTokenWeights() view returns(uint256)
func (_Vault *VaultSession) TotalTokenWeights() (*big.Int, error) {
	return _Vault.Contract.TotalTokenWeights(&_Vault.CallOpts)
}

// TotalTokenWeights is a free data retrieval call binding the contract method 0xdc8f5fac.
//
// Solidity: function totalTokenWeights() view returns(uint256)
func (_Vault *VaultCallerSession) TotalTokenWeights() (*big.Int, error) {
	return _Vault.Contract.TotalTokenWeights(&_Vault.CallOpts)
}

// UsdToToken is a free data retrieval call binding the contract method 0xfa12dbc0.
//
// Solidity: function usdToToken(address _token, uint256 _usdAmount, uint256 _price) view returns(uint256)
func (_Vault *VaultCaller) UsdToToken(opts *bind.CallOpts, _token common.Address, _usdAmount *big.Int, _price *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "usdToToken", _token, _usdAmount, _price)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UsdToToken is a free data retrieval call binding the contract method 0xfa12dbc0.
//
// Solidity: function usdToToken(address _token, uint256 _usdAmount, uint256 _price) view returns(uint256)
func (_Vault *VaultSession) UsdToToken(_token common.Address, _usdAmount *big.Int, _price *big.Int) (*big.Int, error) {
	return _Vault.Contract.UsdToToken(&_Vault.CallOpts, _token, _usdAmount, _price)
}

// UsdToToken is a free data retrieval call binding the contract method 0xfa12dbc0.
//
// Solidity: function usdToToken(address _token, uint256 _usdAmount, uint256 _price) view returns(uint256)
func (_Vault *VaultCallerSession) UsdToToken(_token common.Address, _usdAmount *big.Int, _price *big.Int) (*big.Int, error) {
	return _Vault.Contract.UsdToToken(&_Vault.CallOpts, _token, _usdAmount, _price)
}

// UsdToTokenMax is a free data retrieval call binding the contract method 0xa42ab3d2.
//
// Solidity: function usdToTokenMax(address _token, uint256 _usdAmount) view returns(uint256)
func (_Vault *VaultCaller) UsdToTokenMax(opts *bind.CallOpts, _token common.Address, _usdAmount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "usdToTokenMax", _token, _usdAmount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UsdToTokenMax is a free data retrieval call binding the contract method 0xa42ab3d2.
//
// Solidity: function usdToTokenMax(address _token, uint256 _usdAmount) view returns(uint256)
func (_Vault *VaultSession) UsdToTokenMax(_token common.Address, _usdAmount *big.Int) (*big.Int, error) {
	return _Vault.Contract.UsdToTokenMax(&_Vault.CallOpts, _token, _usdAmount)
}

// UsdToTokenMax is a free data retrieval call binding the contract method 0xa42ab3d2.
//
// Solidity: function usdToTokenMax(address _token, uint256 _usdAmount) view returns(uint256)
func (_Vault *VaultCallerSession) UsdToTokenMax(_token common.Address, _usdAmount *big.Int) (*big.Int, error) {
	return _Vault.Contract.UsdToTokenMax(&_Vault.CallOpts, _token, _usdAmount)
}

// UsdToTokenMin is a free data retrieval call binding the contract method 0x9899cd02.
//
// Solidity: function usdToTokenMin(address _token, uint256 _usdAmount) view returns(uint256)
func (_Vault *VaultCaller) UsdToTokenMin(opts *bind.CallOpts, _token common.Address, _usdAmount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "usdToTokenMin", _token, _usdAmount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UsdToTokenMin is a free data retrieval call binding the contract method 0x9899cd02.
//
// Solidity: function usdToTokenMin(address _token, uint256 _usdAmount) view returns(uint256)
func (_Vault *VaultSession) UsdToTokenMin(_token common.Address, _usdAmount *big.Int) (*big.Int, error) {
	return _Vault.Contract.UsdToTokenMin(&_Vault.CallOpts, _token, _usdAmount)
}

// UsdToTokenMin is a free data retrieval call binding the contract method 0x9899cd02.
//
// Solidity: function usdToTokenMin(address _token, uint256 _usdAmount) view returns(uint256)
func (_Vault *VaultCallerSession) UsdToTokenMin(_token common.Address, _usdAmount *big.Int) (*big.Int, error) {
	return _Vault.Contract.UsdToTokenMin(&_Vault.CallOpts, _token, _usdAmount)
}

// Usdg is a free data retrieval call binding the contract method 0xf5b91b7b.
//
// Solidity: function usdg() view returns(address)
func (_Vault *VaultCaller) Usdg(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "usdg")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Usdg is a free data retrieval call binding the contract method 0xf5b91b7b.
//
// Solidity: function usdg() view returns(address)
func (_Vault *VaultSession) Usdg() (common.Address, error) {
	return _Vault.Contract.Usdg(&_Vault.CallOpts)
}

// Usdg is a free data retrieval call binding the contract method 0xf5b91b7b.
//
// Solidity: function usdg() view returns(address)
func (_Vault *VaultCallerSession) Usdg() (common.Address, error) {
	return _Vault.Contract.Usdg(&_Vault.CallOpts)
}

// UsdgAmounts is a free data retrieval call binding the contract method 0x1aa4ace5.
//
// Solidity: function usdgAmounts(address ) view returns(uint256)
func (_Vault *VaultCaller) UsdgAmounts(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "usdgAmounts", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UsdgAmounts is a free data retrieval call binding the contract method 0x1aa4ace5.
//
// Solidity: function usdgAmounts(address ) view returns(uint256)
func (_Vault *VaultSession) UsdgAmounts(arg0 common.Address) (*big.Int, error) {
	return _Vault.Contract.UsdgAmounts(&_Vault.CallOpts, arg0)
}

// UsdgAmounts is a free data retrieval call binding the contract method 0x1aa4ace5.
//
// Solidity: function usdgAmounts(address ) view returns(uint256)
func (_Vault *VaultCallerSession) UsdgAmounts(arg0 common.Address) (*big.Int, error) {
	return _Vault.Contract.UsdgAmounts(&_Vault.CallOpts, arg0)
}

// UseSwapPricing is a free data retrieval call binding the contract method 0xb06423f3.
//
// Solidity: function useSwapPricing() view returns(bool)
func (_Vault *VaultCaller) UseSwapPricing(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "useSwapPricing")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// UseSwapPricing is a free data retrieval call binding the contract method 0xb06423f3.
//
// Solidity: function useSwapPricing() view returns(bool)
func (_Vault *VaultSession) UseSwapPricing() (bool, error) {
	return _Vault.Contract.UseSwapPricing(&_Vault.CallOpts)
}

// UseSwapPricing is a free data retrieval call binding the contract method 0xb06423f3.
//
// Solidity: function useSwapPricing() view returns(bool)
func (_Vault *VaultCallerSession) UseSwapPricing() (bool, error) {
	return _Vault.Contract.UseSwapPricing(&_Vault.CallOpts)
}

// ValidateLiquidation is a free data retrieval call binding the contract method 0xd54d5a9f.
//
// Solidity: function validateLiquidation(address _account, address _collateralToken, address _indexToken, bool _isLong, bool _raise) view returns(uint256, uint256)
func (_Vault *VaultCaller) ValidateLiquidation(opts *bind.CallOpts, _account common.Address, _collateralToken common.Address, _indexToken common.Address, _isLong bool, _raise bool) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "validateLiquidation", _account, _collateralToken, _indexToken, _isLong, _raise)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// ValidateLiquidation is a free data retrieval call binding the contract method 0xd54d5a9f.
//
// Solidity: function validateLiquidation(address _account, address _collateralToken, address _indexToken, bool _isLong, bool _raise) view returns(uint256, uint256)
func (_Vault *VaultSession) ValidateLiquidation(_account common.Address, _collateralToken common.Address, _indexToken common.Address, _isLong bool, _raise bool) (*big.Int, *big.Int, error) {
	return _Vault.Contract.ValidateLiquidation(&_Vault.CallOpts, _account, _collateralToken, _indexToken, _isLong, _raise)
}

// ValidateLiquidation is a free data retrieval call binding the contract method 0xd54d5a9f.
//
// Solidity: function validateLiquidation(address _account, address _collateralToken, address _indexToken, bool _isLong, bool _raise) view returns(uint256, uint256)
func (_Vault *VaultCallerSession) ValidateLiquidation(_account common.Address, _collateralToken common.Address, _indexToken common.Address, _isLong bool, _raise bool) (*big.Int, *big.Int, error) {
	return _Vault.Contract.ValidateLiquidation(&_Vault.CallOpts, _account, _collateralToken, _indexToken, _isLong, _raise)
}

// VaultUtils is a free data retrieval call binding the contract method 0x6abbe0c8.
//
// Solidity: function vaultUtils() view returns(address)
func (_Vault *VaultCaller) VaultUtils(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "vaultUtils")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VaultUtils is a free data retrieval call binding the contract method 0x6abbe0c8.
//
// Solidity: function vaultUtils() view returns(address)
func (_Vault *VaultSession) VaultUtils() (common.Address, error) {
	return _Vault.Contract.VaultUtils(&_Vault.CallOpts)
}

// VaultUtils is a free data retrieval call binding the contract method 0x6abbe0c8.
//
// Solidity: function vaultUtils() view returns(address)
func (_Vault *VaultCallerSession) VaultUtils() (common.Address, error) {
	return _Vault.Contract.VaultUtils(&_Vault.CallOpts)
}

// WhitelistedTokenCount is a free data retrieval call binding the contract method 0x62287a32.
//
// Solidity: function whitelistedTokenCount() view returns(uint256)
func (_Vault *VaultCaller) WhitelistedTokenCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "whitelistedTokenCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WhitelistedTokenCount is a free data retrieval call binding the contract method 0x62287a32.
//
// Solidity: function whitelistedTokenCount() view returns(uint256)
func (_Vault *VaultSession) WhitelistedTokenCount() (*big.Int, error) {
	return _Vault.Contract.WhitelistedTokenCount(&_Vault.CallOpts)
}

// WhitelistedTokenCount is a free data retrieval call binding the contract method 0x62287a32.
//
// Solidity: function whitelistedTokenCount() view returns(uint256)
func (_Vault *VaultCallerSession) WhitelistedTokenCount() (*big.Int, error) {
	return _Vault.Contract.WhitelistedTokenCount(&_Vault.CallOpts)
}

// WhitelistedTokens is a free data retrieval call binding the contract method 0xdaf9c210.
//
// Solidity: function whitelistedTokens(address ) view returns(bool)
func (_Vault *VaultCaller) WhitelistedTokens(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "whitelistedTokens", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// WhitelistedTokens is a free data retrieval call binding the contract method 0xdaf9c210.
//
// Solidity: function whitelistedTokens(address ) view returns(bool)
func (_Vault *VaultSession) WhitelistedTokens(arg0 common.Address) (bool, error) {
	return _Vault.Contract.WhitelistedTokens(&_Vault.CallOpts, arg0)
}

// WhitelistedTokens is a free data retrieval call binding the contract method 0xdaf9c210.
//
// Solidity: function whitelistedTokens(address ) view returns(bool)
func (_Vault *VaultCallerSession) WhitelistedTokens(arg0 common.Address) (bool, error) {
	return _Vault.Contract.WhitelistedTokens(&_Vault.CallOpts, arg0)
}

// AddRouter is a paid mutator transaction binding the contract method 0x24ca984e.
//
// Solidity: function addRouter(address _router) returns()
func (_Vault *VaultTransactor) AddRouter(opts *bind.TransactOpts, _router common.Address) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "addRouter", _router)
}

// AddRouter is a paid mutator transaction binding the contract method 0x24ca984e.
//
// Solidity: function addRouter(address _router) returns()
func (_Vault *VaultSession) AddRouter(_router common.Address) (*types.Transaction, error) {
	return _Vault.Contract.AddRouter(&_Vault.TransactOpts, _router)
}

// AddRouter is a paid mutator transaction binding the contract method 0x24ca984e.
//
// Solidity: function addRouter(address _router) returns()
func (_Vault *VaultTransactorSession) AddRouter(_router common.Address) (*types.Transaction, error) {
	return _Vault.Contract.AddRouter(&_Vault.TransactOpts, _router)
}

// BuyUSDG is a paid mutator transaction binding the contract method 0x817bb857.
//
// Solidity: function buyUSDG(address _token, address _receiver) returns(uint256)
func (_Vault *VaultTransactor) BuyUSDG(opts *bind.TransactOpts, _token common.Address, _receiver common.Address) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "buyUSDG", _token, _receiver)
}

// BuyUSDG is a paid mutator transaction binding the contract method 0x817bb857.
//
// Solidity: function buyUSDG(address _token, address _receiver) returns(uint256)
func (_Vault *VaultSession) BuyUSDG(_token common.Address, _receiver common.Address) (*types.Transaction, error) {
	return _Vault.Contract.BuyUSDG(&_Vault.TransactOpts, _token, _receiver)
}

// BuyUSDG is a paid mutator transaction binding the contract method 0x817bb857.
//
// Solidity: function buyUSDG(address _token, address _receiver) returns(uint256)
func (_Vault *VaultTransactorSession) BuyUSDG(_token common.Address, _receiver common.Address) (*types.Transaction, error) {
	return _Vault.Contract.BuyUSDG(&_Vault.TransactOpts, _token, _receiver)
}

// ClearTokenConfig is a paid mutator transaction binding the contract method 0xe67f59a7.
//
// Solidity: function clearTokenConfig(address _token) returns()
func (_Vault *VaultTransactor) ClearTokenConfig(opts *bind.TransactOpts, _token common.Address) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "clearTokenConfig", _token)
}

// ClearTokenConfig is a paid mutator transaction binding the contract method 0xe67f59a7.
//
// Solidity: function clearTokenConfig(address _token) returns()
func (_Vault *VaultSession) ClearTokenConfig(_token common.Address) (*types.Transaction, error) {
	return _Vault.Contract.ClearTokenConfig(&_Vault.TransactOpts, _token)
}

// ClearTokenConfig is a paid mutator transaction binding the contract method 0xe67f59a7.
//
// Solidity: function clearTokenConfig(address _token) returns()
func (_Vault *VaultTransactorSession) ClearTokenConfig(_token common.Address) (*types.Transaction, error) {
	return _Vault.Contract.ClearTokenConfig(&_Vault.TransactOpts, _token)
}

// DecreasePosition is a paid mutator transaction binding the contract method 0x82a08490.
//
// Solidity: function decreasePosition(address _account, address _collateralToken, address _indexToken, uint256 _collateralDelta, uint256 _sizeDelta, bool _isLong, address _receiver) returns(uint256)
func (_Vault *VaultTransactor) DecreasePosition(opts *bind.TransactOpts, _account common.Address, _collateralToken common.Address, _indexToken common.Address, _collateralDelta *big.Int, _sizeDelta *big.Int, _isLong bool, _receiver common.Address) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "decreasePosition", _account, _collateralToken, _indexToken, _collateralDelta, _sizeDelta, _isLong, _receiver)
}

// DecreasePosition is a paid mutator transaction binding the contract method 0x82a08490.
//
// Solidity: function decreasePosition(address _account, address _collateralToken, address _indexToken, uint256 _collateralDelta, uint256 _sizeDelta, bool _isLong, address _receiver) returns(uint256)
func (_Vault *VaultSession) DecreasePosition(_account common.Address, _collateralToken common.Address, _indexToken common.Address, _collateralDelta *big.Int, _sizeDelta *big.Int, _isLong bool, _receiver common.Address) (*types.Transaction, error) {
	return _Vault.Contract.DecreasePosition(&_Vault.TransactOpts, _account, _collateralToken, _indexToken, _collateralDelta, _sizeDelta, _isLong, _receiver)
}

// DecreasePosition is a paid mutator transaction binding the contract method 0x82a08490.
//
// Solidity: function decreasePosition(address _account, address _collateralToken, address _indexToken, uint256 _collateralDelta, uint256 _sizeDelta, bool _isLong, address _receiver) returns(uint256)
func (_Vault *VaultTransactorSession) DecreasePosition(_account common.Address, _collateralToken common.Address, _indexToken common.Address, _collateralDelta *big.Int, _sizeDelta *big.Int, _isLong bool, _receiver common.Address) (*types.Transaction, error) {
	return _Vault.Contract.DecreasePosition(&_Vault.TransactOpts, _account, _collateralToken, _indexToken, _collateralDelta, _sizeDelta, _isLong, _receiver)
}

// DirectPoolDeposit is a paid mutator transaction binding the contract method 0x5f7bc119.
//
// Solidity: function directPoolDeposit(address _token) returns()
func (_Vault *VaultTransactor) DirectPoolDeposit(opts *bind.TransactOpts, _token common.Address) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "directPoolDeposit", _token)
}

// DirectPoolDeposit is a paid mutator transaction binding the contract method 0x5f7bc119.
//
// Solidity: function directPoolDeposit(address _token) returns()
func (_Vault *VaultSession) DirectPoolDeposit(_token common.Address) (*types.Transaction, error) {
	return _Vault.Contract.DirectPoolDeposit(&_Vault.TransactOpts, _token)
}

// DirectPoolDeposit is a paid mutator transaction binding the contract method 0x5f7bc119.
//
// Solidity: function directPoolDeposit(address _token) returns()
func (_Vault *VaultTransactorSession) DirectPoolDeposit(_token common.Address) (*types.Transaction, error) {
	return _Vault.Contract.DirectPoolDeposit(&_Vault.TransactOpts, _token)
}

// IncreasePosition is a paid mutator transaction binding the contract method 0x48d91abf.
//
// Solidity: function increasePosition(address _account, address _collateralToken, address _indexToken, uint256 _sizeDelta, bool _isLong) returns()
func (_Vault *VaultTransactor) IncreasePosition(opts *bind.TransactOpts, _account common.Address, _collateralToken common.Address, _indexToken common.Address, _sizeDelta *big.Int, _isLong bool) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "increasePosition", _account, _collateralToken, _indexToken, _sizeDelta, _isLong)
}

// IncreasePosition is a paid mutator transaction binding the contract method 0x48d91abf.
//
// Solidity: function increasePosition(address _account, address _collateralToken, address _indexToken, uint256 _sizeDelta, bool _isLong) returns()
func (_Vault *VaultSession) IncreasePosition(_account common.Address, _collateralToken common.Address, _indexToken common.Address, _sizeDelta *big.Int, _isLong bool) (*types.Transaction, error) {
	return _Vault.Contract.IncreasePosition(&_Vault.TransactOpts, _account, _collateralToken, _indexToken, _sizeDelta, _isLong)
}

// IncreasePosition is a paid mutator transaction binding the contract method 0x48d91abf.
//
// Solidity: function increasePosition(address _account, address _collateralToken, address _indexToken, uint256 _sizeDelta, bool _isLong) returns()
func (_Vault *VaultTransactorSession) IncreasePosition(_account common.Address, _collateralToken common.Address, _indexToken common.Address, _sizeDelta *big.Int, _isLong bool) (*types.Transaction, error) {
	return _Vault.Contract.IncreasePosition(&_Vault.TransactOpts, _account, _collateralToken, _indexToken, _sizeDelta, _isLong)
}

// Initialize is a paid mutator transaction binding the contract method 0x728cdbca.
//
// Solidity: function initialize(address _router, address _usdg, address _priceFeed, uint256 _liquidationFeeUsd, uint256 _fundingRateFactor, uint256 _stableFundingRateFactor) returns()
func (_Vault *VaultTransactor) Initialize(opts *bind.TransactOpts, _router common.Address, _usdg common.Address, _priceFeed common.Address, _liquidationFeeUsd *big.Int, _fundingRateFactor *big.Int, _stableFundingRateFactor *big.Int) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "initialize", _router, _usdg, _priceFeed, _liquidationFeeUsd, _fundingRateFactor, _stableFundingRateFactor)
}

// Initialize is a paid mutator transaction binding the contract method 0x728cdbca.
//
// Solidity: function initialize(address _router, address _usdg, address _priceFeed, uint256 _liquidationFeeUsd, uint256 _fundingRateFactor, uint256 _stableFundingRateFactor) returns()
func (_Vault *VaultSession) Initialize(_router common.Address, _usdg common.Address, _priceFeed common.Address, _liquidationFeeUsd *big.Int, _fundingRateFactor *big.Int, _stableFundingRateFactor *big.Int) (*types.Transaction, error) {
	return _Vault.Contract.Initialize(&_Vault.TransactOpts, _router, _usdg, _priceFeed, _liquidationFeeUsd, _fundingRateFactor, _stableFundingRateFactor)
}

// Initialize is a paid mutator transaction binding the contract method 0x728cdbca.
//
// Solidity: function initialize(address _router, address _usdg, address _priceFeed, uint256 _liquidationFeeUsd, uint256 _fundingRateFactor, uint256 _stableFundingRateFactor) returns()
func (_Vault *VaultTransactorSession) Initialize(_router common.Address, _usdg common.Address, _priceFeed common.Address, _liquidationFeeUsd *big.Int, _fundingRateFactor *big.Int, _stableFundingRateFactor *big.Int) (*types.Transaction, error) {
	return _Vault.Contract.Initialize(&_Vault.TransactOpts, _router, _usdg, _priceFeed, _liquidationFeeUsd, _fundingRateFactor, _stableFundingRateFactor)
}

// LiquidatePosition is a paid mutator transaction binding the contract method 0xde2ea948.
//
// Solidity: function liquidatePosition(address _account, address _collateralToken, address _indexToken, bool _isLong, address _feeReceiver) returns()
func (_Vault *VaultTransactor) LiquidatePosition(opts *bind.TransactOpts, _account common.Address, _collateralToken common.Address, _indexToken common.Address, _isLong bool, _feeReceiver common.Address) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "liquidatePosition", _account, _collateralToken, _indexToken, _isLong, _feeReceiver)
}

// LiquidatePosition is a paid mutator transaction binding the contract method 0xde2ea948.
//
// Solidity: function liquidatePosition(address _account, address _collateralToken, address _indexToken, bool _isLong, address _feeReceiver) returns()
func (_Vault *VaultSession) LiquidatePosition(_account common.Address, _collateralToken common.Address, _indexToken common.Address, _isLong bool, _feeReceiver common.Address) (*types.Transaction, error) {
	return _Vault.Contract.LiquidatePosition(&_Vault.TransactOpts, _account, _collateralToken, _indexToken, _isLong, _feeReceiver)
}

// LiquidatePosition is a paid mutator transaction binding the contract method 0xde2ea948.
//
// Solidity: function liquidatePosition(address _account, address _collateralToken, address _indexToken, bool _isLong, address _feeReceiver) returns()
func (_Vault *VaultTransactorSession) LiquidatePosition(_account common.Address, _collateralToken common.Address, _indexToken common.Address, _isLong bool, _feeReceiver common.Address) (*types.Transaction, error) {
	return _Vault.Contract.LiquidatePosition(&_Vault.TransactOpts, _account, _collateralToken, _indexToken, _isLong, _feeReceiver)
}

// RemoveRouter is a paid mutator transaction binding the contract method 0x6ae0b154.
//
// Solidity: function removeRouter(address _router) returns()
func (_Vault *VaultTransactor) RemoveRouter(opts *bind.TransactOpts, _router common.Address) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "removeRouter", _router)
}

// RemoveRouter is a paid mutator transaction binding the contract method 0x6ae0b154.
//
// Solidity: function removeRouter(address _router) returns()
func (_Vault *VaultSession) RemoveRouter(_router common.Address) (*types.Transaction, error) {
	return _Vault.Contract.RemoveRouter(&_Vault.TransactOpts, _router)
}

// RemoveRouter is a paid mutator transaction binding the contract method 0x6ae0b154.
//
// Solidity: function removeRouter(address _router) returns()
func (_Vault *VaultTransactorSession) RemoveRouter(_router common.Address) (*types.Transaction, error) {
	return _Vault.Contract.RemoveRouter(&_Vault.TransactOpts, _router)
}

// SellUSDG is a paid mutator transaction binding the contract method 0x711e6190.
//
// Solidity: function sellUSDG(address _token, address _receiver) returns(uint256)
func (_Vault *VaultTransactor) SellUSDG(opts *bind.TransactOpts, _token common.Address, _receiver common.Address) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "sellUSDG", _token, _receiver)
}

// SellUSDG is a paid mutator transaction binding the contract method 0x711e6190.
//
// Solidity: function sellUSDG(address _token, address _receiver) returns(uint256)
func (_Vault *VaultSession) SellUSDG(_token common.Address, _receiver common.Address) (*types.Transaction, error) {
	return _Vault.Contract.SellUSDG(&_Vault.TransactOpts, _token, _receiver)
}

// SellUSDG is a paid mutator transaction binding the contract method 0x711e6190.
//
// Solidity: function sellUSDG(address _token, address _receiver) returns(uint256)
func (_Vault *VaultTransactorSession) SellUSDG(_token common.Address, _receiver common.Address) (*types.Transaction, error) {
	return _Vault.Contract.SellUSDG(&_Vault.TransactOpts, _token, _receiver)
}

// SetBufferAmount is a paid mutator transaction binding the contract method 0x8585f4d2.
//
// Solidity: function setBufferAmount(address _token, uint256 _amount) returns()
func (_Vault *VaultTransactor) SetBufferAmount(opts *bind.TransactOpts, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "setBufferAmount", _token, _amount)
}

// SetBufferAmount is a paid mutator transaction binding the contract method 0x8585f4d2.
//
// Solidity: function setBufferAmount(address _token, uint256 _amount) returns()
func (_Vault *VaultSession) SetBufferAmount(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Vault.Contract.SetBufferAmount(&_Vault.TransactOpts, _token, _amount)
}

// SetBufferAmount is a paid mutator transaction binding the contract method 0x8585f4d2.
//
// Solidity: function setBufferAmount(address _token, uint256 _amount) returns()
func (_Vault *VaultTransactorSession) SetBufferAmount(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Vault.Contract.SetBufferAmount(&_Vault.TransactOpts, _token, _amount)
}

// SetError is a paid mutator transaction binding the contract method 0x28e67be5.
//
// Solidity: function setError(uint256 _errorCode, string _error) returns()
func (_Vault *VaultTransactor) SetError(opts *bind.TransactOpts, _errorCode *big.Int, _error string) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "setError", _errorCode, _error)
}

// SetError is a paid mutator transaction binding the contract method 0x28e67be5.
//
// Solidity: function setError(uint256 _errorCode, string _error) returns()
func (_Vault *VaultSession) SetError(_errorCode *big.Int, _error string) (*types.Transaction, error) {
	return _Vault.Contract.SetError(&_Vault.TransactOpts, _errorCode, _error)
}

// SetError is a paid mutator transaction binding the contract method 0x28e67be5.
//
// Solidity: function setError(uint256 _errorCode, string _error) returns()
func (_Vault *VaultTransactorSession) SetError(_errorCode *big.Int, _error string) (*types.Transaction, error) {
	return _Vault.Contract.SetError(&_Vault.TransactOpts, _errorCode, _error)
}

// SetErrorController is a paid mutator transaction binding the contract method 0x8f7b8404.
//
// Solidity: function setErrorController(address _errorController) returns()
func (_Vault *VaultTransactor) SetErrorController(opts *bind.TransactOpts, _errorController common.Address) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "setErrorController", _errorController)
}

// SetErrorController is a paid mutator transaction binding the contract method 0x8f7b8404.
//
// Solidity: function setErrorController(address _errorController) returns()
func (_Vault *VaultSession) SetErrorController(_errorController common.Address) (*types.Transaction, error) {
	return _Vault.Contract.SetErrorController(&_Vault.TransactOpts, _errorController)
}

// SetErrorController is a paid mutator transaction binding the contract method 0x8f7b8404.
//
// Solidity: function setErrorController(address _errorController) returns()
func (_Vault *VaultTransactorSession) SetErrorController(_errorController common.Address) (*types.Transaction, error) {
	return _Vault.Contract.SetErrorController(&_Vault.TransactOpts, _errorController)
}

// SetFees is a paid mutator transaction binding the contract method 0x40eb3802.
//
// Solidity: function setFees(uint256 _taxBasisPoints, uint256 _stableTaxBasisPoints, uint256 _mintBurnFeeBasisPoints, uint256 _swapFeeBasisPoints, uint256 _stableSwapFeeBasisPoints, uint256 _marginFeeBasisPoints, uint256 _liquidationFeeUsd, uint256 _minProfitTime, bool _hasDynamicFees) returns()
func (_Vault *VaultTransactor) SetFees(opts *bind.TransactOpts, _taxBasisPoints *big.Int, _stableTaxBasisPoints *big.Int, _mintBurnFeeBasisPoints *big.Int, _swapFeeBasisPoints *big.Int, _stableSwapFeeBasisPoints *big.Int, _marginFeeBasisPoints *big.Int, _liquidationFeeUsd *big.Int, _minProfitTime *big.Int, _hasDynamicFees bool) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "setFees", _taxBasisPoints, _stableTaxBasisPoints, _mintBurnFeeBasisPoints, _swapFeeBasisPoints, _stableSwapFeeBasisPoints, _marginFeeBasisPoints, _liquidationFeeUsd, _minProfitTime, _hasDynamicFees)
}

// SetFees is a paid mutator transaction binding the contract method 0x40eb3802.
//
// Solidity: function setFees(uint256 _taxBasisPoints, uint256 _stableTaxBasisPoints, uint256 _mintBurnFeeBasisPoints, uint256 _swapFeeBasisPoints, uint256 _stableSwapFeeBasisPoints, uint256 _marginFeeBasisPoints, uint256 _liquidationFeeUsd, uint256 _minProfitTime, bool _hasDynamicFees) returns()
func (_Vault *VaultSession) SetFees(_taxBasisPoints *big.Int, _stableTaxBasisPoints *big.Int, _mintBurnFeeBasisPoints *big.Int, _swapFeeBasisPoints *big.Int, _stableSwapFeeBasisPoints *big.Int, _marginFeeBasisPoints *big.Int, _liquidationFeeUsd *big.Int, _minProfitTime *big.Int, _hasDynamicFees bool) (*types.Transaction, error) {
	return _Vault.Contract.SetFees(&_Vault.TransactOpts, _taxBasisPoints, _stableTaxBasisPoints, _mintBurnFeeBasisPoints, _swapFeeBasisPoints, _stableSwapFeeBasisPoints, _marginFeeBasisPoints, _liquidationFeeUsd, _minProfitTime, _hasDynamicFees)
}

// SetFees is a paid mutator transaction binding the contract method 0x40eb3802.
//
// Solidity: function setFees(uint256 _taxBasisPoints, uint256 _stableTaxBasisPoints, uint256 _mintBurnFeeBasisPoints, uint256 _swapFeeBasisPoints, uint256 _stableSwapFeeBasisPoints, uint256 _marginFeeBasisPoints, uint256 _liquidationFeeUsd, uint256 _minProfitTime, bool _hasDynamicFees) returns()
func (_Vault *VaultTransactorSession) SetFees(_taxBasisPoints *big.Int, _stableTaxBasisPoints *big.Int, _mintBurnFeeBasisPoints *big.Int, _swapFeeBasisPoints *big.Int, _stableSwapFeeBasisPoints *big.Int, _marginFeeBasisPoints *big.Int, _liquidationFeeUsd *big.Int, _minProfitTime *big.Int, _hasDynamicFees bool) (*types.Transaction, error) {
	return _Vault.Contract.SetFees(&_Vault.TransactOpts, _taxBasisPoints, _stableTaxBasisPoints, _mintBurnFeeBasisPoints, _swapFeeBasisPoints, _stableSwapFeeBasisPoints, _marginFeeBasisPoints, _liquidationFeeUsd, _minProfitTime, _hasDynamicFees)
}

// SetFundingRate is a paid mutator transaction binding the contract method 0x8a27d468.
//
// Solidity: function setFundingRate(uint256 _fundingInterval, uint256 _fundingRateFactor, uint256 _stableFundingRateFactor) returns()
func (_Vault *VaultTransactor) SetFundingRate(opts *bind.TransactOpts, _fundingInterval *big.Int, _fundingRateFactor *big.Int, _stableFundingRateFactor *big.Int) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "setFundingRate", _fundingInterval, _fundingRateFactor, _stableFundingRateFactor)
}

// SetFundingRate is a paid mutator transaction binding the contract method 0x8a27d468.
//
// Solidity: function setFundingRate(uint256 _fundingInterval, uint256 _fundingRateFactor, uint256 _stableFundingRateFactor) returns()
func (_Vault *VaultSession) SetFundingRate(_fundingInterval *big.Int, _fundingRateFactor *big.Int, _stableFundingRateFactor *big.Int) (*types.Transaction, error) {
	return _Vault.Contract.SetFundingRate(&_Vault.TransactOpts, _fundingInterval, _fundingRateFactor, _stableFundingRateFactor)
}

// SetFundingRate is a paid mutator transaction binding the contract method 0x8a27d468.
//
// Solidity: function setFundingRate(uint256 _fundingInterval, uint256 _fundingRateFactor, uint256 _stableFundingRateFactor) returns()
func (_Vault *VaultTransactorSession) SetFundingRate(_fundingInterval *big.Int, _fundingRateFactor *big.Int, _stableFundingRateFactor *big.Int) (*types.Transaction, error) {
	return _Vault.Contract.SetFundingRate(&_Vault.TransactOpts, _fundingInterval, _fundingRateFactor, _stableFundingRateFactor)
}

// SetGov is a paid mutator transaction binding the contract method 0xcfad57a2.
//
// Solidity: function setGov(address _gov) returns()
func (_Vault *VaultTransactor) SetGov(opts *bind.TransactOpts, _gov common.Address) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "setGov", _gov)
}

// SetGov is a paid mutator transaction binding the contract method 0xcfad57a2.
//
// Solidity: function setGov(address _gov) returns()
func (_Vault *VaultSession) SetGov(_gov common.Address) (*types.Transaction, error) {
	return _Vault.Contract.SetGov(&_Vault.TransactOpts, _gov)
}

// SetGov is a paid mutator transaction binding the contract method 0xcfad57a2.
//
// Solidity: function setGov(address _gov) returns()
func (_Vault *VaultTransactorSession) SetGov(_gov common.Address) (*types.Transaction, error) {
	return _Vault.Contract.SetGov(&_Vault.TransactOpts, _gov)
}

// SetInManagerMode is a paid mutator transaction binding the contract method 0x24b0c04d.
//
// Solidity: function setInManagerMode(bool _inManagerMode) returns()
func (_Vault *VaultTransactor) SetInManagerMode(opts *bind.TransactOpts, _inManagerMode bool) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "setInManagerMode", _inManagerMode)
}

// SetInManagerMode is a paid mutator transaction binding the contract method 0x24b0c04d.
//
// Solidity: function setInManagerMode(bool _inManagerMode) returns()
func (_Vault *VaultSession) SetInManagerMode(_inManagerMode bool) (*types.Transaction, error) {
	return _Vault.Contract.SetInManagerMode(&_Vault.TransactOpts, _inManagerMode)
}

// SetInManagerMode is a paid mutator transaction binding the contract method 0x24b0c04d.
//
// Solidity: function setInManagerMode(bool _inManagerMode) returns()
func (_Vault *VaultTransactorSession) SetInManagerMode(_inManagerMode bool) (*types.Transaction, error) {
	return _Vault.Contract.SetInManagerMode(&_Vault.TransactOpts, _inManagerMode)
}

// SetInPrivateLiquidationMode is a paid mutator transaction binding the contract method 0xf07bbf77.
//
// Solidity: function setInPrivateLiquidationMode(bool _inPrivateLiquidationMode) returns()
func (_Vault *VaultTransactor) SetInPrivateLiquidationMode(opts *bind.TransactOpts, _inPrivateLiquidationMode bool) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "setInPrivateLiquidationMode", _inPrivateLiquidationMode)
}

// SetInPrivateLiquidationMode is a paid mutator transaction binding the contract method 0xf07bbf77.
//
// Solidity: function setInPrivateLiquidationMode(bool _inPrivateLiquidationMode) returns()
func (_Vault *VaultSession) SetInPrivateLiquidationMode(_inPrivateLiquidationMode bool) (*types.Transaction, error) {
	return _Vault.Contract.SetInPrivateLiquidationMode(&_Vault.TransactOpts, _inPrivateLiquidationMode)
}

// SetInPrivateLiquidationMode is a paid mutator transaction binding the contract method 0xf07bbf77.
//
// Solidity: function setInPrivateLiquidationMode(bool _inPrivateLiquidationMode) returns()
func (_Vault *VaultTransactorSession) SetInPrivateLiquidationMode(_inPrivateLiquidationMode bool) (*types.Transaction, error) {
	return _Vault.Contract.SetInPrivateLiquidationMode(&_Vault.TransactOpts, _inPrivateLiquidationMode)
}

// SetIsLeverageEnabled is a paid mutator transaction binding the contract method 0x7c2eb9f7.
//
// Solidity: function setIsLeverageEnabled(bool _isLeverageEnabled) returns()
func (_Vault *VaultTransactor) SetIsLeverageEnabled(opts *bind.TransactOpts, _isLeverageEnabled bool) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "setIsLeverageEnabled", _isLeverageEnabled)
}

// SetIsLeverageEnabled is a paid mutator transaction binding the contract method 0x7c2eb9f7.
//
// Solidity: function setIsLeverageEnabled(bool _isLeverageEnabled) returns()
func (_Vault *VaultSession) SetIsLeverageEnabled(_isLeverageEnabled bool) (*types.Transaction, error) {
	return _Vault.Contract.SetIsLeverageEnabled(&_Vault.TransactOpts, _isLeverageEnabled)
}

// SetIsLeverageEnabled is a paid mutator transaction binding the contract method 0x7c2eb9f7.
//
// Solidity: function setIsLeverageEnabled(bool _isLeverageEnabled) returns()
func (_Vault *VaultTransactorSession) SetIsLeverageEnabled(_isLeverageEnabled bool) (*types.Transaction, error) {
	return _Vault.Contract.SetIsLeverageEnabled(&_Vault.TransactOpts, _isLeverageEnabled)
}

// SetIsSwapEnabled is a paid mutator transaction binding the contract method 0x30455ede.
//
// Solidity: function setIsSwapEnabled(bool _isSwapEnabled) returns()
func (_Vault *VaultTransactor) SetIsSwapEnabled(opts *bind.TransactOpts, _isSwapEnabled bool) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "setIsSwapEnabled", _isSwapEnabled)
}

// SetIsSwapEnabled is a paid mutator transaction binding the contract method 0x30455ede.
//
// Solidity: function setIsSwapEnabled(bool _isSwapEnabled) returns()
func (_Vault *VaultSession) SetIsSwapEnabled(_isSwapEnabled bool) (*types.Transaction, error) {
	return _Vault.Contract.SetIsSwapEnabled(&_Vault.TransactOpts, _isSwapEnabled)
}

// SetIsSwapEnabled is a paid mutator transaction binding the contract method 0x30455ede.
//
// Solidity: function setIsSwapEnabled(bool _isSwapEnabled) returns()
func (_Vault *VaultTransactorSession) SetIsSwapEnabled(_isSwapEnabled bool) (*types.Transaction, error) {
	return _Vault.Contract.SetIsSwapEnabled(&_Vault.TransactOpts, _isSwapEnabled)
}

// SetLiquidator is a paid mutator transaction binding the contract method 0x4453a374.
//
// Solidity: function setLiquidator(address _liquidator, bool _isActive) returns()
func (_Vault *VaultTransactor) SetLiquidator(opts *bind.TransactOpts, _liquidator common.Address, _isActive bool) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "setLiquidator", _liquidator, _isActive)
}

// SetLiquidator is a paid mutator transaction binding the contract method 0x4453a374.
//
// Solidity: function setLiquidator(address _liquidator, bool _isActive) returns()
func (_Vault *VaultSession) SetLiquidator(_liquidator common.Address, _isActive bool) (*types.Transaction, error) {
	return _Vault.Contract.SetLiquidator(&_Vault.TransactOpts, _liquidator, _isActive)
}

// SetLiquidator is a paid mutator transaction binding the contract method 0x4453a374.
//
// Solidity: function setLiquidator(address _liquidator, bool _isActive) returns()
func (_Vault *VaultTransactorSession) SetLiquidator(_liquidator common.Address, _isActive bool) (*types.Transaction, error) {
	return _Vault.Contract.SetLiquidator(&_Vault.TransactOpts, _liquidator, _isActive)
}

// SetManager is a paid mutator transaction binding the contract method 0xa5e90eee.
//
// Solidity: function setManager(address _manager, bool _isManager) returns()
func (_Vault *VaultTransactor) SetManager(opts *bind.TransactOpts, _manager common.Address, _isManager bool) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "setManager", _manager, _isManager)
}

// SetManager is a paid mutator transaction binding the contract method 0xa5e90eee.
//
// Solidity: function setManager(address _manager, bool _isManager) returns()
func (_Vault *VaultSession) SetManager(_manager common.Address, _isManager bool) (*types.Transaction, error) {
	return _Vault.Contract.SetManager(&_Vault.TransactOpts, _manager, _isManager)
}

// SetManager is a paid mutator transaction binding the contract method 0xa5e90eee.
//
// Solidity: function setManager(address _manager, bool _isManager) returns()
func (_Vault *VaultTransactorSession) SetManager(_manager common.Address, _isManager bool) (*types.Transaction, error) {
	return _Vault.Contract.SetManager(&_Vault.TransactOpts, _manager, _isManager)
}

// SetMaxGasPrice is a paid mutator transaction binding the contract method 0xd2fa635e.
//
// Solidity: function setMaxGasPrice(uint256 _maxGasPrice) returns()
func (_Vault *VaultTransactor) SetMaxGasPrice(opts *bind.TransactOpts, _maxGasPrice *big.Int) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "setMaxGasPrice", _maxGasPrice)
}

// SetMaxGasPrice is a paid mutator transaction binding the contract method 0xd2fa635e.
//
// Solidity: function setMaxGasPrice(uint256 _maxGasPrice) returns()
func (_Vault *VaultSession) SetMaxGasPrice(_maxGasPrice *big.Int) (*types.Transaction, error) {
	return _Vault.Contract.SetMaxGasPrice(&_Vault.TransactOpts, _maxGasPrice)
}

// SetMaxGasPrice is a paid mutator transaction binding the contract method 0xd2fa635e.
//
// Solidity: function setMaxGasPrice(uint256 _maxGasPrice) returns()
func (_Vault *VaultTransactorSession) SetMaxGasPrice(_maxGasPrice *big.Int) (*types.Transaction, error) {
	return _Vault.Contract.SetMaxGasPrice(&_Vault.TransactOpts, _maxGasPrice)
}

// SetMaxGlobalShortSize is a paid mutator transaction binding the contract method 0xefa10a6e.
//
// Solidity: function setMaxGlobalShortSize(address _token, uint256 _amount) returns()
func (_Vault *VaultTransactor) SetMaxGlobalShortSize(opts *bind.TransactOpts, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "setMaxGlobalShortSize", _token, _amount)
}

// SetMaxGlobalShortSize is a paid mutator transaction binding the contract method 0xefa10a6e.
//
// Solidity: function setMaxGlobalShortSize(address _token, uint256 _amount) returns()
func (_Vault *VaultSession) SetMaxGlobalShortSize(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Vault.Contract.SetMaxGlobalShortSize(&_Vault.TransactOpts, _token, _amount)
}

// SetMaxGlobalShortSize is a paid mutator transaction binding the contract method 0xefa10a6e.
//
// Solidity: function setMaxGlobalShortSize(address _token, uint256 _amount) returns()
func (_Vault *VaultTransactorSession) SetMaxGlobalShortSize(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Vault.Contract.SetMaxGlobalShortSize(&_Vault.TransactOpts, _token, _amount)
}

// SetMaxLeverage is a paid mutator transaction binding the contract method 0xd3127e63.
//
// Solidity: function setMaxLeverage(uint256 _maxLeverage) returns()
func (_Vault *VaultTransactor) SetMaxLeverage(opts *bind.TransactOpts, _maxLeverage *big.Int) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "setMaxLeverage", _maxLeverage)
}

// SetMaxLeverage is a paid mutator transaction binding the contract method 0xd3127e63.
//
// Solidity: function setMaxLeverage(uint256 _maxLeverage) returns()
func (_Vault *VaultSession) SetMaxLeverage(_maxLeverage *big.Int) (*types.Transaction, error) {
	return _Vault.Contract.SetMaxLeverage(&_Vault.TransactOpts, _maxLeverage)
}

// SetMaxLeverage is a paid mutator transaction binding the contract method 0xd3127e63.
//
// Solidity: function setMaxLeverage(uint256 _maxLeverage) returns()
func (_Vault *VaultTransactorSession) SetMaxLeverage(_maxLeverage *big.Int) (*types.Transaction, error) {
	return _Vault.Contract.SetMaxLeverage(&_Vault.TransactOpts, _maxLeverage)
}

// SetPriceFeed is a paid mutator transaction binding the contract method 0x724e78da.
//
// Solidity: function setPriceFeed(address _priceFeed) returns()
func (_Vault *VaultTransactor) SetPriceFeed(opts *bind.TransactOpts, _priceFeed common.Address) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "setPriceFeed", _priceFeed)
}

// SetPriceFeed is a paid mutator transaction binding the contract method 0x724e78da.
//
// Solidity: function setPriceFeed(address _priceFeed) returns()
func (_Vault *VaultSession) SetPriceFeed(_priceFeed common.Address) (*types.Transaction, error) {
	return _Vault.Contract.SetPriceFeed(&_Vault.TransactOpts, _priceFeed)
}

// SetPriceFeed is a paid mutator transaction binding the contract method 0x724e78da.
//
// Solidity: function setPriceFeed(address _priceFeed) returns()
func (_Vault *VaultTransactorSession) SetPriceFeed(_priceFeed common.Address) (*types.Transaction, error) {
	return _Vault.Contract.SetPriceFeed(&_Vault.TransactOpts, _priceFeed)
}

// SetTokenConfig is a paid mutator transaction binding the contract method 0x3c5a6e35.
//
// Solidity: function setTokenConfig(address _token, uint256 _tokenDecimals, uint256 _tokenWeight, uint256 _minProfitBps, uint256 _maxUsdgAmount, bool _isStable, bool _isShortable) returns()
func (_Vault *VaultTransactor) SetTokenConfig(opts *bind.TransactOpts, _token common.Address, _tokenDecimals *big.Int, _tokenWeight *big.Int, _minProfitBps *big.Int, _maxUsdgAmount *big.Int, _isStable bool, _isShortable bool) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "setTokenConfig", _token, _tokenDecimals, _tokenWeight, _minProfitBps, _maxUsdgAmount, _isStable, _isShortable)
}

// SetTokenConfig is a paid mutator transaction binding the contract method 0x3c5a6e35.
//
// Solidity: function setTokenConfig(address _token, uint256 _tokenDecimals, uint256 _tokenWeight, uint256 _minProfitBps, uint256 _maxUsdgAmount, bool _isStable, bool _isShortable) returns()
func (_Vault *VaultSession) SetTokenConfig(_token common.Address, _tokenDecimals *big.Int, _tokenWeight *big.Int, _minProfitBps *big.Int, _maxUsdgAmount *big.Int, _isStable bool, _isShortable bool) (*types.Transaction, error) {
	return _Vault.Contract.SetTokenConfig(&_Vault.TransactOpts, _token, _tokenDecimals, _tokenWeight, _minProfitBps, _maxUsdgAmount, _isStable, _isShortable)
}

// SetTokenConfig is a paid mutator transaction binding the contract method 0x3c5a6e35.
//
// Solidity: function setTokenConfig(address _token, uint256 _tokenDecimals, uint256 _tokenWeight, uint256 _minProfitBps, uint256 _maxUsdgAmount, bool _isStable, bool _isShortable) returns()
func (_Vault *VaultTransactorSession) SetTokenConfig(_token common.Address, _tokenDecimals *big.Int, _tokenWeight *big.Int, _minProfitBps *big.Int, _maxUsdgAmount *big.Int, _isStable bool, _isShortable bool) (*types.Transaction, error) {
	return _Vault.Contract.SetTokenConfig(&_Vault.TransactOpts, _token, _tokenDecimals, _tokenWeight, _minProfitBps, _maxUsdgAmount, _isStable, _isShortable)
}

// SetUsdgAmount is a paid mutator transaction binding the contract method 0xd66b000d.
//
// Solidity: function setUsdgAmount(address _token, uint256 _amount) returns()
func (_Vault *VaultTransactor) SetUsdgAmount(opts *bind.TransactOpts, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "setUsdgAmount", _token, _amount)
}

// SetUsdgAmount is a paid mutator transaction binding the contract method 0xd66b000d.
//
// Solidity: function setUsdgAmount(address _token, uint256 _amount) returns()
func (_Vault *VaultSession) SetUsdgAmount(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Vault.Contract.SetUsdgAmount(&_Vault.TransactOpts, _token, _amount)
}

// SetUsdgAmount is a paid mutator transaction binding the contract method 0xd66b000d.
//
// Solidity: function setUsdgAmount(address _token, uint256 _amount) returns()
func (_Vault *VaultTransactorSession) SetUsdgAmount(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Vault.Contract.SetUsdgAmount(&_Vault.TransactOpts, _token, _amount)
}

// SetVaultUtils is a paid mutator transaction binding the contract method 0x71089f4d.
//
// Solidity: function setVaultUtils(address _vaultUtils) returns()
func (_Vault *VaultTransactor) SetVaultUtils(opts *bind.TransactOpts, _vaultUtils common.Address) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "setVaultUtils", _vaultUtils)
}

// SetVaultUtils is a paid mutator transaction binding the contract method 0x71089f4d.
//
// Solidity: function setVaultUtils(address _vaultUtils) returns()
func (_Vault *VaultSession) SetVaultUtils(_vaultUtils common.Address) (*types.Transaction, error) {
	return _Vault.Contract.SetVaultUtils(&_Vault.TransactOpts, _vaultUtils)
}

// SetVaultUtils is a paid mutator transaction binding the contract method 0x71089f4d.
//
// Solidity: function setVaultUtils(address _vaultUtils) returns()
func (_Vault *VaultTransactorSession) SetVaultUtils(_vaultUtils common.Address) (*types.Transaction, error) {
	return _Vault.Contract.SetVaultUtils(&_Vault.TransactOpts, _vaultUtils)
}

// Swap is a paid mutator transaction binding the contract method 0x93316212.
//
// Solidity: function swap(address _tokenIn, address _tokenOut, address _receiver) returns(uint256)
func (_Vault *VaultTransactor) Swap(opts *bind.TransactOpts, _tokenIn common.Address, _tokenOut common.Address, _receiver common.Address) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "swap", _tokenIn, _tokenOut, _receiver)
}

// Swap is a paid mutator transaction binding the contract method 0x93316212.
//
// Solidity: function swap(address _tokenIn, address _tokenOut, address _receiver) returns(uint256)
func (_Vault *VaultSession) Swap(_tokenIn common.Address, _tokenOut common.Address, _receiver common.Address) (*types.Transaction, error) {
	return _Vault.Contract.Swap(&_Vault.TransactOpts, _tokenIn, _tokenOut, _receiver)
}

// Swap is a paid mutator transaction binding the contract method 0x93316212.
//
// Solidity: function swap(address _tokenIn, address _tokenOut, address _receiver) returns(uint256)
func (_Vault *VaultTransactorSession) Swap(_tokenIn common.Address, _tokenOut common.Address, _receiver common.Address) (*types.Transaction, error) {
	return _Vault.Contract.Swap(&_Vault.TransactOpts, _tokenIn, _tokenOut, _receiver)
}

// UpdateCumulativeFundingRate is a paid mutator transaction binding the contract method 0xfbfded6d.
//
// Solidity: function updateCumulativeFundingRate(address _collateralToken, address _indexToken) returns()
func (_Vault *VaultTransactor) UpdateCumulativeFundingRate(opts *bind.TransactOpts, _collateralToken common.Address, _indexToken common.Address) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "updateCumulativeFundingRate", _collateralToken, _indexToken)
}

// UpdateCumulativeFundingRate is a paid mutator transaction binding the contract method 0xfbfded6d.
//
// Solidity: function updateCumulativeFundingRate(address _collateralToken, address _indexToken) returns()
func (_Vault *VaultSession) UpdateCumulativeFundingRate(_collateralToken common.Address, _indexToken common.Address) (*types.Transaction, error) {
	return _Vault.Contract.UpdateCumulativeFundingRate(&_Vault.TransactOpts, _collateralToken, _indexToken)
}

// UpdateCumulativeFundingRate is a paid mutator transaction binding the contract method 0xfbfded6d.
//
// Solidity: function updateCumulativeFundingRate(address _collateralToken, address _indexToken) returns()
func (_Vault *VaultTransactorSession) UpdateCumulativeFundingRate(_collateralToken common.Address, _indexToken common.Address) (*types.Transaction, error) {
	return _Vault.Contract.UpdateCumulativeFundingRate(&_Vault.TransactOpts, _collateralToken, _indexToken)
}

// UpgradeVault is a paid mutator transaction binding the contract method 0xcea0c328.
//
// Solidity: function upgradeVault(address _newVault, address _token, uint256 _amount) returns()
func (_Vault *VaultTransactor) UpgradeVault(opts *bind.TransactOpts, _newVault common.Address, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "upgradeVault", _newVault, _token, _amount)
}

// UpgradeVault is a paid mutator transaction binding the contract method 0xcea0c328.
//
// Solidity: function upgradeVault(address _newVault, address _token, uint256 _amount) returns()
func (_Vault *VaultSession) UpgradeVault(_newVault common.Address, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Vault.Contract.UpgradeVault(&_Vault.TransactOpts, _newVault, _token, _amount)
}

// UpgradeVault is a paid mutator transaction binding the contract method 0xcea0c328.
//
// Solidity: function upgradeVault(address _newVault, address _token, uint256 _amount) returns()
func (_Vault *VaultTransactorSession) UpgradeVault(_newVault common.Address, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Vault.Contract.UpgradeVault(&_Vault.TransactOpts, _newVault, _token, _amount)
}

// WithdrawFees is a paid mutator transaction binding the contract method 0xf2555278.
//
// Solidity: function withdrawFees(address _token, address _receiver) returns(uint256)
func (_Vault *VaultTransactor) WithdrawFees(opts *bind.TransactOpts, _token common.Address, _receiver common.Address) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "withdrawFees", _token, _receiver)
}

// WithdrawFees is a paid mutator transaction binding the contract method 0xf2555278.
//
// Solidity: function withdrawFees(address _token, address _receiver) returns(uint256)
func (_Vault *VaultSession) WithdrawFees(_token common.Address, _receiver common.Address) (*types.Transaction, error) {
	return _Vault.Contract.WithdrawFees(&_Vault.TransactOpts, _token, _receiver)
}

// WithdrawFees is a paid mutator transaction binding the contract method 0xf2555278.
//
// Solidity: function withdrawFees(address _token, address _receiver) returns(uint256)
func (_Vault *VaultTransactorSession) WithdrawFees(_token common.Address, _receiver common.Address) (*types.Transaction, error) {
	return _Vault.Contract.WithdrawFees(&_Vault.TransactOpts, _token, _receiver)
}

// VaultBuyUSDGIterator is returned from FilterBuyUSDG and is used to iterate over the raw logs and unpacked data for BuyUSDG events raised by the Vault contract.
type VaultBuyUSDGIterator struct {
	Event *VaultBuyUSDG // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VaultBuyUSDGIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultBuyUSDG)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VaultBuyUSDG)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VaultBuyUSDGIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultBuyUSDGIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultBuyUSDG represents a BuyUSDG event raised by the Vault contract.
type VaultBuyUSDG struct {
	Account        common.Address
	Token          common.Address
	TokenAmount    *big.Int
	UsdgAmount     *big.Int
	FeeBasisPoints *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterBuyUSDG is a free log retrieval operation binding the contract event 0xab4c77c74cd32c85f35416cf03e7ce9e2d4387f7b7f2c1f4bf53daaecf8ea72d.
//
// Solidity: event BuyUSDG(address account, address token, uint256 tokenAmount, uint256 usdgAmount, uint256 feeBasisPoints)
func (_Vault *VaultFilterer) FilterBuyUSDG(opts *bind.FilterOpts) (*VaultBuyUSDGIterator, error) {

	logs, sub, err := _Vault.contract.FilterLogs(opts, "BuyUSDG")
	if err != nil {
		return nil, err
	}
	return &VaultBuyUSDGIterator{contract: _Vault.contract, event: "BuyUSDG", logs: logs, sub: sub}, nil
}

// WatchBuyUSDG is a free log subscription operation binding the contract event 0xab4c77c74cd32c85f35416cf03e7ce9e2d4387f7b7f2c1f4bf53daaecf8ea72d.
//
// Solidity: event BuyUSDG(address account, address token, uint256 tokenAmount, uint256 usdgAmount, uint256 feeBasisPoints)
func (_Vault *VaultFilterer) WatchBuyUSDG(opts *bind.WatchOpts, sink chan<- *VaultBuyUSDG) (event.Subscription, error) {

	logs, sub, err := _Vault.contract.WatchLogs(opts, "BuyUSDG")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultBuyUSDG)
				if err := _Vault.contract.UnpackLog(event, "BuyUSDG", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBuyUSDG is a log parse operation binding the contract event 0xab4c77c74cd32c85f35416cf03e7ce9e2d4387f7b7f2c1f4bf53daaecf8ea72d.
//
// Solidity: event BuyUSDG(address account, address token, uint256 tokenAmount, uint256 usdgAmount, uint256 feeBasisPoints)
func (_Vault *VaultFilterer) ParseBuyUSDG(log types.Log) (*VaultBuyUSDG, error) {
	event := new(VaultBuyUSDG)
	if err := _Vault.contract.UnpackLog(event, "BuyUSDG", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultClosePositionIterator is returned from FilterClosePosition and is used to iterate over the raw logs and unpacked data for ClosePosition events raised by the Vault contract.
type VaultClosePositionIterator struct {
	Event *VaultClosePosition // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VaultClosePositionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultClosePosition)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VaultClosePosition)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VaultClosePositionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultClosePositionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultClosePosition represents a ClosePosition event raised by the Vault contract.
type VaultClosePosition struct {
	Key              [32]byte
	Size             *big.Int
	Collateral       *big.Int
	AveragePrice     *big.Int
	EntryFundingRate *big.Int
	ReserveAmount    *big.Int
	RealisedPnl      *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterClosePosition is a free log retrieval operation binding the contract event 0x73af1d417d82c240fdb6d319b34ad884487c6bf2845d98980cc52ad9171cb455.
//
// Solidity: event ClosePosition(bytes32 key, uint256 size, uint256 collateral, uint256 averagePrice, uint256 entryFundingRate, uint256 reserveAmount, int256 realisedPnl)
func (_Vault *VaultFilterer) FilterClosePosition(opts *bind.FilterOpts) (*VaultClosePositionIterator, error) {

	logs, sub, err := _Vault.contract.FilterLogs(opts, "ClosePosition")
	if err != nil {
		return nil, err
	}
	return &VaultClosePositionIterator{contract: _Vault.contract, event: "ClosePosition", logs: logs, sub: sub}, nil
}

// WatchClosePosition is a free log subscription operation binding the contract event 0x73af1d417d82c240fdb6d319b34ad884487c6bf2845d98980cc52ad9171cb455.
//
// Solidity: event ClosePosition(bytes32 key, uint256 size, uint256 collateral, uint256 averagePrice, uint256 entryFundingRate, uint256 reserveAmount, int256 realisedPnl)
func (_Vault *VaultFilterer) WatchClosePosition(opts *bind.WatchOpts, sink chan<- *VaultClosePosition) (event.Subscription, error) {

	logs, sub, err := _Vault.contract.WatchLogs(opts, "ClosePosition")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultClosePosition)
				if err := _Vault.contract.UnpackLog(event, "ClosePosition", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseClosePosition is a log parse operation binding the contract event 0x73af1d417d82c240fdb6d319b34ad884487c6bf2845d98980cc52ad9171cb455.
//
// Solidity: event ClosePosition(bytes32 key, uint256 size, uint256 collateral, uint256 averagePrice, uint256 entryFundingRate, uint256 reserveAmount, int256 realisedPnl)
func (_Vault *VaultFilterer) ParseClosePosition(log types.Log) (*VaultClosePosition, error) {
	event := new(VaultClosePosition)
	if err := _Vault.contract.UnpackLog(event, "ClosePosition", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultCollectMarginFeesIterator is returned from FilterCollectMarginFees and is used to iterate over the raw logs and unpacked data for CollectMarginFees events raised by the Vault contract.
type VaultCollectMarginFeesIterator struct {
	Event *VaultCollectMarginFees // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VaultCollectMarginFeesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultCollectMarginFees)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VaultCollectMarginFees)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VaultCollectMarginFeesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultCollectMarginFeesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultCollectMarginFees represents a CollectMarginFees event raised by the Vault contract.
type VaultCollectMarginFees struct {
	Token     common.Address
	FeeUsd    *big.Int
	FeeTokens *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterCollectMarginFees is a free log retrieval operation binding the contract event 0x5d0c0019d3d45fadeb74eff9d2c9924d146d000ac6bcf3c28bf0ac3c9baa011a.
//
// Solidity: event CollectMarginFees(address token, uint256 feeUsd, uint256 feeTokens)
func (_Vault *VaultFilterer) FilterCollectMarginFees(opts *bind.FilterOpts) (*VaultCollectMarginFeesIterator, error) {

	logs, sub, err := _Vault.contract.FilterLogs(opts, "CollectMarginFees")
	if err != nil {
		return nil, err
	}
	return &VaultCollectMarginFeesIterator{contract: _Vault.contract, event: "CollectMarginFees", logs: logs, sub: sub}, nil
}

// WatchCollectMarginFees is a free log subscription operation binding the contract event 0x5d0c0019d3d45fadeb74eff9d2c9924d146d000ac6bcf3c28bf0ac3c9baa011a.
//
// Solidity: event CollectMarginFees(address token, uint256 feeUsd, uint256 feeTokens)
func (_Vault *VaultFilterer) WatchCollectMarginFees(opts *bind.WatchOpts, sink chan<- *VaultCollectMarginFees) (event.Subscription, error) {

	logs, sub, err := _Vault.contract.WatchLogs(opts, "CollectMarginFees")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultCollectMarginFees)
				if err := _Vault.contract.UnpackLog(event, "CollectMarginFees", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCollectMarginFees is a log parse operation binding the contract event 0x5d0c0019d3d45fadeb74eff9d2c9924d146d000ac6bcf3c28bf0ac3c9baa011a.
//
// Solidity: event CollectMarginFees(address token, uint256 feeUsd, uint256 feeTokens)
func (_Vault *VaultFilterer) ParseCollectMarginFees(log types.Log) (*VaultCollectMarginFees, error) {
	event := new(VaultCollectMarginFees)
	if err := _Vault.contract.UnpackLog(event, "CollectMarginFees", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultCollectSwapFeesIterator is returned from FilterCollectSwapFees and is used to iterate over the raw logs and unpacked data for CollectSwapFees events raised by the Vault contract.
type VaultCollectSwapFeesIterator struct {
	Event *VaultCollectSwapFees // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VaultCollectSwapFeesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultCollectSwapFees)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VaultCollectSwapFees)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VaultCollectSwapFeesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultCollectSwapFeesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultCollectSwapFees represents a CollectSwapFees event raised by the Vault contract.
type VaultCollectSwapFees struct {
	Token     common.Address
	FeeUsd    *big.Int
	FeeTokens *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterCollectSwapFees is a free log retrieval operation binding the contract event 0x47cd9dda0e50ce30bcaaacd0488452b596221c07ac402a581cfae4d3933cac2b.
//
// Solidity: event CollectSwapFees(address token, uint256 feeUsd, uint256 feeTokens)
func (_Vault *VaultFilterer) FilterCollectSwapFees(opts *bind.FilterOpts) (*VaultCollectSwapFeesIterator, error) {

	logs, sub, err := _Vault.contract.FilterLogs(opts, "CollectSwapFees")
	if err != nil {
		return nil, err
	}
	return &VaultCollectSwapFeesIterator{contract: _Vault.contract, event: "CollectSwapFees", logs: logs, sub: sub}, nil
}

// WatchCollectSwapFees is a free log subscription operation binding the contract event 0x47cd9dda0e50ce30bcaaacd0488452b596221c07ac402a581cfae4d3933cac2b.
//
// Solidity: event CollectSwapFees(address token, uint256 feeUsd, uint256 feeTokens)
func (_Vault *VaultFilterer) WatchCollectSwapFees(opts *bind.WatchOpts, sink chan<- *VaultCollectSwapFees) (event.Subscription, error) {

	logs, sub, err := _Vault.contract.WatchLogs(opts, "CollectSwapFees")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultCollectSwapFees)
				if err := _Vault.contract.UnpackLog(event, "CollectSwapFees", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCollectSwapFees is a log parse operation binding the contract event 0x47cd9dda0e50ce30bcaaacd0488452b596221c07ac402a581cfae4d3933cac2b.
//
// Solidity: event CollectSwapFees(address token, uint256 feeUsd, uint256 feeTokens)
func (_Vault *VaultFilterer) ParseCollectSwapFees(log types.Log) (*VaultCollectSwapFees, error) {
	event := new(VaultCollectSwapFees)
	if err := _Vault.contract.UnpackLog(event, "CollectSwapFees", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultDecreaseGuaranteedUsdIterator is returned from FilterDecreaseGuaranteedUsd and is used to iterate over the raw logs and unpacked data for DecreaseGuaranteedUsd events raised by the Vault contract.
type VaultDecreaseGuaranteedUsdIterator struct {
	Event *VaultDecreaseGuaranteedUsd // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VaultDecreaseGuaranteedUsdIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultDecreaseGuaranteedUsd)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VaultDecreaseGuaranteedUsd)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VaultDecreaseGuaranteedUsdIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultDecreaseGuaranteedUsdIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultDecreaseGuaranteedUsd represents a DecreaseGuaranteedUsd event raised by the Vault contract.
type VaultDecreaseGuaranteedUsd struct {
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDecreaseGuaranteedUsd is a free log retrieval operation binding the contract event 0x34e07158b9db50df5613e591c44ea2ebc82834eff4a4dc3a46e000e608261d68.
//
// Solidity: event DecreaseGuaranteedUsd(address token, uint256 amount)
func (_Vault *VaultFilterer) FilterDecreaseGuaranteedUsd(opts *bind.FilterOpts) (*VaultDecreaseGuaranteedUsdIterator, error) {

	logs, sub, err := _Vault.contract.FilterLogs(opts, "DecreaseGuaranteedUsd")
	if err != nil {
		return nil, err
	}
	return &VaultDecreaseGuaranteedUsdIterator{contract: _Vault.contract, event: "DecreaseGuaranteedUsd", logs: logs, sub: sub}, nil
}

// WatchDecreaseGuaranteedUsd is a free log subscription operation binding the contract event 0x34e07158b9db50df5613e591c44ea2ebc82834eff4a4dc3a46e000e608261d68.
//
// Solidity: event DecreaseGuaranteedUsd(address token, uint256 amount)
func (_Vault *VaultFilterer) WatchDecreaseGuaranteedUsd(opts *bind.WatchOpts, sink chan<- *VaultDecreaseGuaranteedUsd) (event.Subscription, error) {

	logs, sub, err := _Vault.contract.WatchLogs(opts, "DecreaseGuaranteedUsd")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultDecreaseGuaranteedUsd)
				if err := _Vault.contract.UnpackLog(event, "DecreaseGuaranteedUsd", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDecreaseGuaranteedUsd is a log parse operation binding the contract event 0x34e07158b9db50df5613e591c44ea2ebc82834eff4a4dc3a46e000e608261d68.
//
// Solidity: event DecreaseGuaranteedUsd(address token, uint256 amount)
func (_Vault *VaultFilterer) ParseDecreaseGuaranteedUsd(log types.Log) (*VaultDecreaseGuaranteedUsd, error) {
	event := new(VaultDecreaseGuaranteedUsd)
	if err := _Vault.contract.UnpackLog(event, "DecreaseGuaranteedUsd", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultDecreasePoolAmountIterator is returned from FilterDecreasePoolAmount and is used to iterate over the raw logs and unpacked data for DecreasePoolAmount events raised by the Vault contract.
type VaultDecreasePoolAmountIterator struct {
	Event *VaultDecreasePoolAmount // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VaultDecreasePoolAmountIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultDecreasePoolAmount)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VaultDecreasePoolAmount)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VaultDecreasePoolAmountIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultDecreasePoolAmountIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultDecreasePoolAmount represents a DecreasePoolAmount event raised by the Vault contract.
type VaultDecreasePoolAmount struct {
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDecreasePoolAmount is a free log retrieval operation binding the contract event 0x112726233fbeaeed0f5b1dba5cb0b2b81883dee49fb35ff99fd98ed9f6d31eb0.
//
// Solidity: event DecreasePoolAmount(address token, uint256 amount)
func (_Vault *VaultFilterer) FilterDecreasePoolAmount(opts *bind.FilterOpts) (*VaultDecreasePoolAmountIterator, error) {

	logs, sub, err := _Vault.contract.FilterLogs(opts, "DecreasePoolAmount")
	if err != nil {
		return nil, err
	}
	return &VaultDecreasePoolAmountIterator{contract: _Vault.contract, event: "DecreasePoolAmount", logs: logs, sub: sub}, nil
}

// WatchDecreasePoolAmount is a free log subscription operation binding the contract event 0x112726233fbeaeed0f5b1dba5cb0b2b81883dee49fb35ff99fd98ed9f6d31eb0.
//
// Solidity: event DecreasePoolAmount(address token, uint256 amount)
func (_Vault *VaultFilterer) WatchDecreasePoolAmount(opts *bind.WatchOpts, sink chan<- *VaultDecreasePoolAmount) (event.Subscription, error) {

	logs, sub, err := _Vault.contract.WatchLogs(opts, "DecreasePoolAmount")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultDecreasePoolAmount)
				if err := _Vault.contract.UnpackLog(event, "DecreasePoolAmount", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDecreasePoolAmount is a log parse operation binding the contract event 0x112726233fbeaeed0f5b1dba5cb0b2b81883dee49fb35ff99fd98ed9f6d31eb0.
//
// Solidity: event DecreasePoolAmount(address token, uint256 amount)
func (_Vault *VaultFilterer) ParseDecreasePoolAmount(log types.Log) (*VaultDecreasePoolAmount, error) {
	event := new(VaultDecreasePoolAmount)
	if err := _Vault.contract.UnpackLog(event, "DecreasePoolAmount", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultDecreasePositionIterator is returned from FilterDecreasePosition and is used to iterate over the raw logs and unpacked data for DecreasePosition events raised by the Vault contract.
type VaultDecreasePositionIterator struct {
	Event *VaultDecreasePosition // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VaultDecreasePositionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultDecreasePosition)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VaultDecreasePosition)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VaultDecreasePositionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultDecreasePositionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultDecreasePosition represents a DecreasePosition event raised by the Vault contract.
type VaultDecreasePosition struct {
	Key             [32]byte
	Account         common.Address
	CollateralToken common.Address
	IndexToken      common.Address
	CollateralDelta *big.Int
	SizeDelta       *big.Int
	IsLong          bool
	Price           *big.Int
	Fee             *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterDecreasePosition is a free log retrieval operation binding the contract event 0x93d75d64d1f84fc6f430a64fc578bdd4c1e090e90ea2d51773e626d19de56d30.
//
// Solidity: event DecreasePosition(bytes32 key, address account, address collateralToken, address indexToken, uint256 collateralDelta, uint256 sizeDelta, bool isLong, uint256 price, uint256 fee)
func (_Vault *VaultFilterer) FilterDecreasePosition(opts *bind.FilterOpts) (*VaultDecreasePositionIterator, error) {

	logs, sub, err := _Vault.contract.FilterLogs(opts, "DecreasePosition")
	if err != nil {
		return nil, err
	}
	return &VaultDecreasePositionIterator{contract: _Vault.contract, event: "DecreasePosition", logs: logs, sub: sub}, nil
}

// WatchDecreasePosition is a free log subscription operation binding the contract event 0x93d75d64d1f84fc6f430a64fc578bdd4c1e090e90ea2d51773e626d19de56d30.
//
// Solidity: event DecreasePosition(bytes32 key, address account, address collateralToken, address indexToken, uint256 collateralDelta, uint256 sizeDelta, bool isLong, uint256 price, uint256 fee)
func (_Vault *VaultFilterer) WatchDecreasePosition(opts *bind.WatchOpts, sink chan<- *VaultDecreasePosition) (event.Subscription, error) {

	logs, sub, err := _Vault.contract.WatchLogs(opts, "DecreasePosition")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultDecreasePosition)
				if err := _Vault.contract.UnpackLog(event, "DecreasePosition", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDecreasePosition is a log parse operation binding the contract event 0x93d75d64d1f84fc6f430a64fc578bdd4c1e090e90ea2d51773e626d19de56d30.
//
// Solidity: event DecreasePosition(bytes32 key, address account, address collateralToken, address indexToken, uint256 collateralDelta, uint256 sizeDelta, bool isLong, uint256 price, uint256 fee)
func (_Vault *VaultFilterer) ParseDecreasePosition(log types.Log) (*VaultDecreasePosition, error) {
	event := new(VaultDecreasePosition)
	if err := _Vault.contract.UnpackLog(event, "DecreasePosition", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultDecreaseReservedAmountIterator is returned from FilterDecreaseReservedAmount and is used to iterate over the raw logs and unpacked data for DecreaseReservedAmount events raised by the Vault contract.
type VaultDecreaseReservedAmountIterator struct {
	Event *VaultDecreaseReservedAmount // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VaultDecreaseReservedAmountIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultDecreaseReservedAmount)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VaultDecreaseReservedAmount)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VaultDecreaseReservedAmountIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultDecreaseReservedAmountIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultDecreaseReservedAmount represents a DecreaseReservedAmount event raised by the Vault contract.
type VaultDecreaseReservedAmount struct {
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDecreaseReservedAmount is a free log retrieval operation binding the contract event 0x533cb5ed32be6a90284e96b5747a1bfc2d38fdb5768a6b5f67ff7d62144ed67b.
//
// Solidity: event DecreaseReservedAmount(address token, uint256 amount)
func (_Vault *VaultFilterer) FilterDecreaseReservedAmount(opts *bind.FilterOpts) (*VaultDecreaseReservedAmountIterator, error) {

	logs, sub, err := _Vault.contract.FilterLogs(opts, "DecreaseReservedAmount")
	if err != nil {
		return nil, err
	}
	return &VaultDecreaseReservedAmountIterator{contract: _Vault.contract, event: "DecreaseReservedAmount", logs: logs, sub: sub}, nil
}

// WatchDecreaseReservedAmount is a free log subscription operation binding the contract event 0x533cb5ed32be6a90284e96b5747a1bfc2d38fdb5768a6b5f67ff7d62144ed67b.
//
// Solidity: event DecreaseReservedAmount(address token, uint256 amount)
func (_Vault *VaultFilterer) WatchDecreaseReservedAmount(opts *bind.WatchOpts, sink chan<- *VaultDecreaseReservedAmount) (event.Subscription, error) {

	logs, sub, err := _Vault.contract.WatchLogs(opts, "DecreaseReservedAmount")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultDecreaseReservedAmount)
				if err := _Vault.contract.UnpackLog(event, "DecreaseReservedAmount", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDecreaseReservedAmount is a log parse operation binding the contract event 0x533cb5ed32be6a90284e96b5747a1bfc2d38fdb5768a6b5f67ff7d62144ed67b.
//
// Solidity: event DecreaseReservedAmount(address token, uint256 amount)
func (_Vault *VaultFilterer) ParseDecreaseReservedAmount(log types.Log) (*VaultDecreaseReservedAmount, error) {
	event := new(VaultDecreaseReservedAmount)
	if err := _Vault.contract.UnpackLog(event, "DecreaseReservedAmount", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultDecreaseUsdgAmountIterator is returned from FilterDecreaseUsdgAmount and is used to iterate over the raw logs and unpacked data for DecreaseUsdgAmount events raised by the Vault contract.
type VaultDecreaseUsdgAmountIterator struct {
	Event *VaultDecreaseUsdgAmount // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VaultDecreaseUsdgAmountIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultDecreaseUsdgAmount)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VaultDecreaseUsdgAmount)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VaultDecreaseUsdgAmountIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultDecreaseUsdgAmountIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultDecreaseUsdgAmount represents a DecreaseUsdgAmount event raised by the Vault contract.
type VaultDecreaseUsdgAmount struct {
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDecreaseUsdgAmount is a free log retrieval operation binding the contract event 0xe1e812596aac93a06ecc4ca627014d18e30f5c33b825160cc9d5c0ba61e45227.
//
// Solidity: event DecreaseUsdgAmount(address token, uint256 amount)
func (_Vault *VaultFilterer) FilterDecreaseUsdgAmount(opts *bind.FilterOpts) (*VaultDecreaseUsdgAmountIterator, error) {

	logs, sub, err := _Vault.contract.FilterLogs(opts, "DecreaseUsdgAmount")
	if err != nil {
		return nil, err
	}
	return &VaultDecreaseUsdgAmountIterator{contract: _Vault.contract, event: "DecreaseUsdgAmount", logs: logs, sub: sub}, nil
}

// WatchDecreaseUsdgAmount is a free log subscription operation binding the contract event 0xe1e812596aac93a06ecc4ca627014d18e30f5c33b825160cc9d5c0ba61e45227.
//
// Solidity: event DecreaseUsdgAmount(address token, uint256 amount)
func (_Vault *VaultFilterer) WatchDecreaseUsdgAmount(opts *bind.WatchOpts, sink chan<- *VaultDecreaseUsdgAmount) (event.Subscription, error) {

	logs, sub, err := _Vault.contract.WatchLogs(opts, "DecreaseUsdgAmount")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultDecreaseUsdgAmount)
				if err := _Vault.contract.UnpackLog(event, "DecreaseUsdgAmount", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDecreaseUsdgAmount is a log parse operation binding the contract event 0xe1e812596aac93a06ecc4ca627014d18e30f5c33b825160cc9d5c0ba61e45227.
//
// Solidity: event DecreaseUsdgAmount(address token, uint256 amount)
func (_Vault *VaultFilterer) ParseDecreaseUsdgAmount(log types.Log) (*VaultDecreaseUsdgAmount, error) {
	event := new(VaultDecreaseUsdgAmount)
	if err := _Vault.contract.UnpackLog(event, "DecreaseUsdgAmount", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultDirectPoolDepositIterator is returned from FilterDirectPoolDeposit and is used to iterate over the raw logs and unpacked data for DirectPoolDeposit events raised by the Vault contract.
type VaultDirectPoolDepositIterator struct {
	Event *VaultDirectPoolDeposit // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VaultDirectPoolDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultDirectPoolDeposit)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VaultDirectPoolDeposit)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VaultDirectPoolDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultDirectPoolDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultDirectPoolDeposit represents a DirectPoolDeposit event raised by the Vault contract.
type VaultDirectPoolDeposit struct {
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDirectPoolDeposit is a free log retrieval operation binding the contract event 0xa5a389190ebf6170a133bda5c769b77f4d6715b8aa172ec0ddf8473d0b4944bd.
//
// Solidity: event DirectPoolDeposit(address token, uint256 amount)
func (_Vault *VaultFilterer) FilterDirectPoolDeposit(opts *bind.FilterOpts) (*VaultDirectPoolDepositIterator, error) {

	logs, sub, err := _Vault.contract.FilterLogs(opts, "DirectPoolDeposit")
	if err != nil {
		return nil, err
	}
	return &VaultDirectPoolDepositIterator{contract: _Vault.contract, event: "DirectPoolDeposit", logs: logs, sub: sub}, nil
}

// WatchDirectPoolDeposit is a free log subscription operation binding the contract event 0xa5a389190ebf6170a133bda5c769b77f4d6715b8aa172ec0ddf8473d0b4944bd.
//
// Solidity: event DirectPoolDeposit(address token, uint256 amount)
func (_Vault *VaultFilterer) WatchDirectPoolDeposit(opts *bind.WatchOpts, sink chan<- *VaultDirectPoolDeposit) (event.Subscription, error) {

	logs, sub, err := _Vault.contract.WatchLogs(opts, "DirectPoolDeposit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultDirectPoolDeposit)
				if err := _Vault.contract.UnpackLog(event, "DirectPoolDeposit", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDirectPoolDeposit is a log parse operation binding the contract event 0xa5a389190ebf6170a133bda5c769b77f4d6715b8aa172ec0ddf8473d0b4944bd.
//
// Solidity: event DirectPoolDeposit(address token, uint256 amount)
func (_Vault *VaultFilterer) ParseDirectPoolDeposit(log types.Log) (*VaultDirectPoolDeposit, error) {
	event := new(VaultDirectPoolDeposit)
	if err := _Vault.contract.UnpackLog(event, "DirectPoolDeposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultIncreaseGuaranteedUsdIterator is returned from FilterIncreaseGuaranteedUsd and is used to iterate over the raw logs and unpacked data for IncreaseGuaranteedUsd events raised by the Vault contract.
type VaultIncreaseGuaranteedUsdIterator struct {
	Event *VaultIncreaseGuaranteedUsd // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VaultIncreaseGuaranteedUsdIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultIncreaseGuaranteedUsd)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VaultIncreaseGuaranteedUsd)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VaultIncreaseGuaranteedUsdIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultIncreaseGuaranteedUsdIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultIncreaseGuaranteedUsd represents a IncreaseGuaranteedUsd event raised by the Vault contract.
type VaultIncreaseGuaranteedUsd struct {
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterIncreaseGuaranteedUsd is a free log retrieval operation binding the contract event 0xd9d4761f75e0d0103b5cbeab941eeb443d7a56a35b5baf2a0787c03f03f4e474.
//
// Solidity: event IncreaseGuaranteedUsd(address token, uint256 amount)
func (_Vault *VaultFilterer) FilterIncreaseGuaranteedUsd(opts *bind.FilterOpts) (*VaultIncreaseGuaranteedUsdIterator, error) {

	logs, sub, err := _Vault.contract.FilterLogs(opts, "IncreaseGuaranteedUsd")
	if err != nil {
		return nil, err
	}
	return &VaultIncreaseGuaranteedUsdIterator{contract: _Vault.contract, event: "IncreaseGuaranteedUsd", logs: logs, sub: sub}, nil
}

// WatchIncreaseGuaranteedUsd is a free log subscription operation binding the contract event 0xd9d4761f75e0d0103b5cbeab941eeb443d7a56a35b5baf2a0787c03f03f4e474.
//
// Solidity: event IncreaseGuaranteedUsd(address token, uint256 amount)
func (_Vault *VaultFilterer) WatchIncreaseGuaranteedUsd(opts *bind.WatchOpts, sink chan<- *VaultIncreaseGuaranteedUsd) (event.Subscription, error) {

	logs, sub, err := _Vault.contract.WatchLogs(opts, "IncreaseGuaranteedUsd")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultIncreaseGuaranteedUsd)
				if err := _Vault.contract.UnpackLog(event, "IncreaseGuaranteedUsd", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseIncreaseGuaranteedUsd is a log parse operation binding the contract event 0xd9d4761f75e0d0103b5cbeab941eeb443d7a56a35b5baf2a0787c03f03f4e474.
//
// Solidity: event IncreaseGuaranteedUsd(address token, uint256 amount)
func (_Vault *VaultFilterer) ParseIncreaseGuaranteedUsd(log types.Log) (*VaultIncreaseGuaranteedUsd, error) {
	event := new(VaultIncreaseGuaranteedUsd)
	if err := _Vault.contract.UnpackLog(event, "IncreaseGuaranteedUsd", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultIncreasePoolAmountIterator is returned from FilterIncreasePoolAmount and is used to iterate over the raw logs and unpacked data for IncreasePoolAmount events raised by the Vault contract.
type VaultIncreasePoolAmountIterator struct {
	Event *VaultIncreasePoolAmount // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VaultIncreasePoolAmountIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultIncreasePoolAmount)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VaultIncreasePoolAmount)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VaultIncreasePoolAmountIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultIncreasePoolAmountIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultIncreasePoolAmount represents a IncreasePoolAmount event raised by the Vault contract.
type VaultIncreasePoolAmount struct {
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterIncreasePoolAmount is a free log retrieval operation binding the contract event 0x976177fbe09a15e5e43f848844963a42b41ef919ef17ff21a17a5421de8f4737.
//
// Solidity: event IncreasePoolAmount(address token, uint256 amount)
func (_Vault *VaultFilterer) FilterIncreasePoolAmount(opts *bind.FilterOpts) (*VaultIncreasePoolAmountIterator, error) {

	logs, sub, err := _Vault.contract.FilterLogs(opts, "IncreasePoolAmount")
	if err != nil {
		return nil, err
	}
	return &VaultIncreasePoolAmountIterator{contract: _Vault.contract, event: "IncreasePoolAmount", logs: logs, sub: sub}, nil
}

// WatchIncreasePoolAmount is a free log subscription operation binding the contract event 0x976177fbe09a15e5e43f848844963a42b41ef919ef17ff21a17a5421de8f4737.
//
// Solidity: event IncreasePoolAmount(address token, uint256 amount)
func (_Vault *VaultFilterer) WatchIncreasePoolAmount(opts *bind.WatchOpts, sink chan<- *VaultIncreasePoolAmount) (event.Subscription, error) {

	logs, sub, err := _Vault.contract.WatchLogs(opts, "IncreasePoolAmount")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultIncreasePoolAmount)
				if err := _Vault.contract.UnpackLog(event, "IncreasePoolAmount", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseIncreasePoolAmount is a log parse operation binding the contract event 0x976177fbe09a15e5e43f848844963a42b41ef919ef17ff21a17a5421de8f4737.
//
// Solidity: event IncreasePoolAmount(address token, uint256 amount)
func (_Vault *VaultFilterer) ParseIncreasePoolAmount(log types.Log) (*VaultIncreasePoolAmount, error) {
	event := new(VaultIncreasePoolAmount)
	if err := _Vault.contract.UnpackLog(event, "IncreasePoolAmount", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultIncreasePositionIterator is returned from FilterIncreasePosition and is used to iterate over the raw logs and unpacked data for IncreasePosition events raised by the Vault contract.
type VaultIncreasePositionIterator struct {
	Event *VaultIncreasePosition // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VaultIncreasePositionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultIncreasePosition)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VaultIncreasePosition)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VaultIncreasePositionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultIncreasePositionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultIncreasePosition represents a IncreasePosition event raised by the Vault contract.
type VaultIncreasePosition struct {
	Key             [32]byte
	Account         common.Address
	CollateralToken common.Address
	IndexToken      common.Address
	CollateralDelta *big.Int
	SizeDelta       *big.Int
	IsLong          bool
	Price           *big.Int
	Fee             *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterIncreasePosition is a free log retrieval operation binding the contract event 0x2fe68525253654c21998f35787a8d0f361905ef647c854092430ab65f2f15022.
//
// Solidity: event IncreasePosition(bytes32 key, address account, address collateralToken, address indexToken, uint256 collateralDelta, uint256 sizeDelta, bool isLong, uint256 price, uint256 fee)
func (_Vault *VaultFilterer) FilterIncreasePosition(opts *bind.FilterOpts) (*VaultIncreasePositionIterator, error) {

	logs, sub, err := _Vault.contract.FilterLogs(opts, "IncreasePosition")
	if err != nil {
		return nil, err
	}
	return &VaultIncreasePositionIterator{contract: _Vault.contract, event: "IncreasePosition", logs: logs, sub: sub}, nil
}

// WatchIncreasePosition is a free log subscription operation binding the contract event 0x2fe68525253654c21998f35787a8d0f361905ef647c854092430ab65f2f15022.
//
// Solidity: event IncreasePosition(bytes32 key, address account, address collateralToken, address indexToken, uint256 collateralDelta, uint256 sizeDelta, bool isLong, uint256 price, uint256 fee)
func (_Vault *VaultFilterer) WatchIncreasePosition(opts *bind.WatchOpts, sink chan<- *VaultIncreasePosition) (event.Subscription, error) {

	logs, sub, err := _Vault.contract.WatchLogs(opts, "IncreasePosition")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultIncreasePosition)
				if err := _Vault.contract.UnpackLog(event, "IncreasePosition", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseIncreasePosition is a log parse operation binding the contract event 0x2fe68525253654c21998f35787a8d0f361905ef647c854092430ab65f2f15022.
//
// Solidity: event IncreasePosition(bytes32 key, address account, address collateralToken, address indexToken, uint256 collateralDelta, uint256 sizeDelta, bool isLong, uint256 price, uint256 fee)
func (_Vault *VaultFilterer) ParseIncreasePosition(log types.Log) (*VaultIncreasePosition, error) {
	event := new(VaultIncreasePosition)
	if err := _Vault.contract.UnpackLog(event, "IncreasePosition", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultIncreaseReservedAmountIterator is returned from FilterIncreaseReservedAmount and is used to iterate over the raw logs and unpacked data for IncreaseReservedAmount events raised by the Vault contract.
type VaultIncreaseReservedAmountIterator struct {
	Event *VaultIncreaseReservedAmount // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VaultIncreaseReservedAmountIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultIncreaseReservedAmount)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VaultIncreaseReservedAmount)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VaultIncreaseReservedAmountIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultIncreaseReservedAmountIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultIncreaseReservedAmount represents a IncreaseReservedAmount event raised by the Vault contract.
type VaultIncreaseReservedAmount struct {
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterIncreaseReservedAmount is a free log retrieval operation binding the contract event 0xaa5649d82f5462be9d19b0f2b31a59b2259950a6076550bac9f3a1c07db9f66d.
//
// Solidity: event IncreaseReservedAmount(address token, uint256 amount)
func (_Vault *VaultFilterer) FilterIncreaseReservedAmount(opts *bind.FilterOpts) (*VaultIncreaseReservedAmountIterator, error) {

	logs, sub, err := _Vault.contract.FilterLogs(opts, "IncreaseReservedAmount")
	if err != nil {
		return nil, err
	}
	return &VaultIncreaseReservedAmountIterator{contract: _Vault.contract, event: "IncreaseReservedAmount", logs: logs, sub: sub}, nil
}

// WatchIncreaseReservedAmount is a free log subscription operation binding the contract event 0xaa5649d82f5462be9d19b0f2b31a59b2259950a6076550bac9f3a1c07db9f66d.
//
// Solidity: event IncreaseReservedAmount(address token, uint256 amount)
func (_Vault *VaultFilterer) WatchIncreaseReservedAmount(opts *bind.WatchOpts, sink chan<- *VaultIncreaseReservedAmount) (event.Subscription, error) {

	logs, sub, err := _Vault.contract.WatchLogs(opts, "IncreaseReservedAmount")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultIncreaseReservedAmount)
				if err := _Vault.contract.UnpackLog(event, "IncreaseReservedAmount", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseIncreaseReservedAmount is a log parse operation binding the contract event 0xaa5649d82f5462be9d19b0f2b31a59b2259950a6076550bac9f3a1c07db9f66d.
//
// Solidity: event IncreaseReservedAmount(address token, uint256 amount)
func (_Vault *VaultFilterer) ParseIncreaseReservedAmount(log types.Log) (*VaultIncreaseReservedAmount, error) {
	event := new(VaultIncreaseReservedAmount)
	if err := _Vault.contract.UnpackLog(event, "IncreaseReservedAmount", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultIncreaseUsdgAmountIterator is returned from FilterIncreaseUsdgAmount and is used to iterate over the raw logs and unpacked data for IncreaseUsdgAmount events raised by the Vault contract.
type VaultIncreaseUsdgAmountIterator struct {
	Event *VaultIncreaseUsdgAmount // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VaultIncreaseUsdgAmountIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultIncreaseUsdgAmount)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VaultIncreaseUsdgAmount)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VaultIncreaseUsdgAmountIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultIncreaseUsdgAmountIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultIncreaseUsdgAmount represents a IncreaseUsdgAmount event raised by the Vault contract.
type VaultIncreaseUsdgAmount struct {
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterIncreaseUsdgAmount is a free log retrieval operation binding the contract event 0x64243679a443432e2293343b77d411ff6144370404618f00ca0d2025d9ca9882.
//
// Solidity: event IncreaseUsdgAmount(address token, uint256 amount)
func (_Vault *VaultFilterer) FilterIncreaseUsdgAmount(opts *bind.FilterOpts) (*VaultIncreaseUsdgAmountIterator, error) {

	logs, sub, err := _Vault.contract.FilterLogs(opts, "IncreaseUsdgAmount")
	if err != nil {
		return nil, err
	}
	return &VaultIncreaseUsdgAmountIterator{contract: _Vault.contract, event: "IncreaseUsdgAmount", logs: logs, sub: sub}, nil
}

// WatchIncreaseUsdgAmount is a free log subscription operation binding the contract event 0x64243679a443432e2293343b77d411ff6144370404618f00ca0d2025d9ca9882.
//
// Solidity: event IncreaseUsdgAmount(address token, uint256 amount)
func (_Vault *VaultFilterer) WatchIncreaseUsdgAmount(opts *bind.WatchOpts, sink chan<- *VaultIncreaseUsdgAmount) (event.Subscription, error) {

	logs, sub, err := _Vault.contract.WatchLogs(opts, "IncreaseUsdgAmount")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultIncreaseUsdgAmount)
				if err := _Vault.contract.UnpackLog(event, "IncreaseUsdgAmount", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseIncreaseUsdgAmount is a log parse operation binding the contract event 0x64243679a443432e2293343b77d411ff6144370404618f00ca0d2025d9ca9882.
//
// Solidity: event IncreaseUsdgAmount(address token, uint256 amount)
func (_Vault *VaultFilterer) ParseIncreaseUsdgAmount(log types.Log) (*VaultIncreaseUsdgAmount, error) {
	event := new(VaultIncreaseUsdgAmount)
	if err := _Vault.contract.UnpackLog(event, "IncreaseUsdgAmount", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultLiquidatePositionIterator is returned from FilterLiquidatePosition and is used to iterate over the raw logs and unpacked data for LiquidatePosition events raised by the Vault contract.
type VaultLiquidatePositionIterator struct {
	Event *VaultLiquidatePosition // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VaultLiquidatePositionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultLiquidatePosition)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VaultLiquidatePosition)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VaultLiquidatePositionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultLiquidatePositionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultLiquidatePosition represents a LiquidatePosition event raised by the Vault contract.
type VaultLiquidatePosition struct {
	Key             [32]byte
	Account         common.Address
	CollateralToken common.Address
	IndexToken      common.Address
	IsLong          bool
	Size            *big.Int
	Collateral      *big.Int
	ReserveAmount   *big.Int
	RealisedPnl     *big.Int
	MarkPrice       *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterLiquidatePosition is a free log retrieval operation binding the contract event 0x2e1f85a64a2f22cf2f0c42584e7c919ed4abe8d53675cff0f62bf1e95a1c676f.
//
// Solidity: event LiquidatePosition(bytes32 key, address account, address collateralToken, address indexToken, bool isLong, uint256 size, uint256 collateral, uint256 reserveAmount, int256 realisedPnl, uint256 markPrice)
func (_Vault *VaultFilterer) FilterLiquidatePosition(opts *bind.FilterOpts) (*VaultLiquidatePositionIterator, error) {

	logs, sub, err := _Vault.contract.FilterLogs(opts, "LiquidatePosition")
	if err != nil {
		return nil, err
	}
	return &VaultLiquidatePositionIterator{contract: _Vault.contract, event: "LiquidatePosition", logs: logs, sub: sub}, nil
}

// WatchLiquidatePosition is a free log subscription operation binding the contract event 0x2e1f85a64a2f22cf2f0c42584e7c919ed4abe8d53675cff0f62bf1e95a1c676f.
//
// Solidity: event LiquidatePosition(bytes32 key, address account, address collateralToken, address indexToken, bool isLong, uint256 size, uint256 collateral, uint256 reserveAmount, int256 realisedPnl, uint256 markPrice)
func (_Vault *VaultFilterer) WatchLiquidatePosition(opts *bind.WatchOpts, sink chan<- *VaultLiquidatePosition) (event.Subscription, error) {

	logs, sub, err := _Vault.contract.WatchLogs(opts, "LiquidatePosition")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultLiquidatePosition)
				if err := _Vault.contract.UnpackLog(event, "LiquidatePosition", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLiquidatePosition is a log parse operation binding the contract event 0x2e1f85a64a2f22cf2f0c42584e7c919ed4abe8d53675cff0f62bf1e95a1c676f.
//
// Solidity: event LiquidatePosition(bytes32 key, address account, address collateralToken, address indexToken, bool isLong, uint256 size, uint256 collateral, uint256 reserveAmount, int256 realisedPnl, uint256 markPrice)
func (_Vault *VaultFilterer) ParseLiquidatePosition(log types.Log) (*VaultLiquidatePosition, error) {
	event := new(VaultLiquidatePosition)
	if err := _Vault.contract.UnpackLog(event, "LiquidatePosition", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultSellUSDGIterator is returned from FilterSellUSDG and is used to iterate over the raw logs and unpacked data for SellUSDG events raised by the Vault contract.
type VaultSellUSDGIterator struct {
	Event *VaultSellUSDG // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VaultSellUSDGIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultSellUSDG)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VaultSellUSDG)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VaultSellUSDGIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultSellUSDGIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultSellUSDG represents a SellUSDG event raised by the Vault contract.
type VaultSellUSDG struct {
	Account        common.Address
	Token          common.Address
	UsdgAmount     *big.Int
	TokenAmount    *big.Int
	FeeBasisPoints *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterSellUSDG is a free log retrieval operation binding the contract event 0xd732b7828fa6cee72c285eac756fc66a7477e3dc22e22e7c432f1c265d40b483.
//
// Solidity: event SellUSDG(address account, address token, uint256 usdgAmount, uint256 tokenAmount, uint256 feeBasisPoints)
func (_Vault *VaultFilterer) FilterSellUSDG(opts *bind.FilterOpts) (*VaultSellUSDGIterator, error) {

	logs, sub, err := _Vault.contract.FilterLogs(opts, "SellUSDG")
	if err != nil {
		return nil, err
	}
	return &VaultSellUSDGIterator{contract: _Vault.contract, event: "SellUSDG", logs: logs, sub: sub}, nil
}

// WatchSellUSDG is a free log subscription operation binding the contract event 0xd732b7828fa6cee72c285eac756fc66a7477e3dc22e22e7c432f1c265d40b483.
//
// Solidity: event SellUSDG(address account, address token, uint256 usdgAmount, uint256 tokenAmount, uint256 feeBasisPoints)
func (_Vault *VaultFilterer) WatchSellUSDG(opts *bind.WatchOpts, sink chan<- *VaultSellUSDG) (event.Subscription, error) {

	logs, sub, err := _Vault.contract.WatchLogs(opts, "SellUSDG")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultSellUSDG)
				if err := _Vault.contract.UnpackLog(event, "SellUSDG", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSellUSDG is a log parse operation binding the contract event 0xd732b7828fa6cee72c285eac756fc66a7477e3dc22e22e7c432f1c265d40b483.
//
// Solidity: event SellUSDG(address account, address token, uint256 usdgAmount, uint256 tokenAmount, uint256 feeBasisPoints)
func (_Vault *VaultFilterer) ParseSellUSDG(log types.Log) (*VaultSellUSDG, error) {
	event := new(VaultSellUSDG)
	if err := _Vault.contract.UnpackLog(event, "SellUSDG", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultSwapIterator is returned from FilterSwap and is used to iterate over the raw logs and unpacked data for Swap events raised by the Vault contract.
type VaultSwapIterator struct {
	Event *VaultSwap // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VaultSwapIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultSwap)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VaultSwap)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VaultSwapIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultSwapIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultSwap represents a Swap event raised by the Vault contract.
type VaultSwap struct {
	Account            common.Address
	TokenIn            common.Address
	TokenOut           common.Address
	AmountIn           *big.Int
	AmountOut          *big.Int
	AmountOutAfterFees *big.Int
	FeeBasisPoints     *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterSwap is a free log retrieval operation binding the contract event 0x0874b2d545cb271cdbda4e093020c452328b24af12382ed62c4d00f5c26709db.
//
// Solidity: event Swap(address account, address tokenIn, address tokenOut, uint256 amountIn, uint256 amountOut, uint256 amountOutAfterFees, uint256 feeBasisPoints)
func (_Vault *VaultFilterer) FilterSwap(opts *bind.FilterOpts) (*VaultSwapIterator, error) {

	logs, sub, err := _Vault.contract.FilterLogs(opts, "Swap")
	if err != nil {
		return nil, err
	}
	return &VaultSwapIterator{contract: _Vault.contract, event: "Swap", logs: logs, sub: sub}, nil
}

// WatchSwap is a free log subscription operation binding the contract event 0x0874b2d545cb271cdbda4e093020c452328b24af12382ed62c4d00f5c26709db.
//
// Solidity: event Swap(address account, address tokenIn, address tokenOut, uint256 amountIn, uint256 amountOut, uint256 amountOutAfterFees, uint256 feeBasisPoints)
func (_Vault *VaultFilterer) WatchSwap(opts *bind.WatchOpts, sink chan<- *VaultSwap) (event.Subscription, error) {

	logs, sub, err := _Vault.contract.WatchLogs(opts, "Swap")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultSwap)
				if err := _Vault.contract.UnpackLog(event, "Swap", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSwap is a log parse operation binding the contract event 0x0874b2d545cb271cdbda4e093020c452328b24af12382ed62c4d00f5c26709db.
//
// Solidity: event Swap(address account, address tokenIn, address tokenOut, uint256 amountIn, uint256 amountOut, uint256 amountOutAfterFees, uint256 feeBasisPoints)
func (_Vault *VaultFilterer) ParseSwap(log types.Log) (*VaultSwap, error) {
	event := new(VaultSwap)
	if err := _Vault.contract.UnpackLog(event, "Swap", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultUpdateFundingRateIterator is returned from FilterUpdateFundingRate and is used to iterate over the raw logs and unpacked data for UpdateFundingRate events raised by the Vault contract.
type VaultUpdateFundingRateIterator struct {
	Event *VaultUpdateFundingRate // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VaultUpdateFundingRateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultUpdateFundingRate)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VaultUpdateFundingRate)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VaultUpdateFundingRateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultUpdateFundingRateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultUpdateFundingRate represents a UpdateFundingRate event raised by the Vault contract.
type VaultUpdateFundingRate struct {
	Token       common.Address
	FundingRate *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterUpdateFundingRate is a free log retrieval operation binding the contract event 0xa146fc154e1913322e9817d49f0d5c37466c24326e15de10e739a948be815eab.
//
// Solidity: event UpdateFundingRate(address token, uint256 fundingRate)
func (_Vault *VaultFilterer) FilterUpdateFundingRate(opts *bind.FilterOpts) (*VaultUpdateFundingRateIterator, error) {

	logs, sub, err := _Vault.contract.FilterLogs(opts, "UpdateFundingRate")
	if err != nil {
		return nil, err
	}
	return &VaultUpdateFundingRateIterator{contract: _Vault.contract, event: "UpdateFundingRate", logs: logs, sub: sub}, nil
}

// WatchUpdateFundingRate is a free log subscription operation binding the contract event 0xa146fc154e1913322e9817d49f0d5c37466c24326e15de10e739a948be815eab.
//
// Solidity: event UpdateFundingRate(address token, uint256 fundingRate)
func (_Vault *VaultFilterer) WatchUpdateFundingRate(opts *bind.WatchOpts, sink chan<- *VaultUpdateFundingRate) (event.Subscription, error) {

	logs, sub, err := _Vault.contract.WatchLogs(opts, "UpdateFundingRate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultUpdateFundingRate)
				if err := _Vault.contract.UnpackLog(event, "UpdateFundingRate", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpdateFundingRate is a log parse operation binding the contract event 0xa146fc154e1913322e9817d49f0d5c37466c24326e15de10e739a948be815eab.
//
// Solidity: event UpdateFundingRate(address token, uint256 fundingRate)
func (_Vault *VaultFilterer) ParseUpdateFundingRate(log types.Log) (*VaultUpdateFundingRate, error) {
	event := new(VaultUpdateFundingRate)
	if err := _Vault.contract.UnpackLog(event, "UpdateFundingRate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultUpdatePnlIterator is returned from FilterUpdatePnl and is used to iterate over the raw logs and unpacked data for UpdatePnl events raised by the Vault contract.
type VaultUpdatePnlIterator struct {
	Event *VaultUpdatePnl // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VaultUpdatePnlIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultUpdatePnl)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VaultUpdatePnl)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VaultUpdatePnlIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultUpdatePnlIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultUpdatePnl represents a UpdatePnl event raised by the Vault contract.
type VaultUpdatePnl struct {
	Key       [32]byte
	HasProfit bool
	Delta     *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUpdatePnl is a free log retrieval operation binding the contract event 0x3ff41bdde87755b687ae83d0221a232b6be51a803330ed9661c1b5d0105e0d8a.
//
// Solidity: event UpdatePnl(bytes32 key, bool hasProfit, uint256 delta)
func (_Vault *VaultFilterer) FilterUpdatePnl(opts *bind.FilterOpts) (*VaultUpdatePnlIterator, error) {

	logs, sub, err := _Vault.contract.FilterLogs(opts, "UpdatePnl")
	if err != nil {
		return nil, err
	}
	return &VaultUpdatePnlIterator{contract: _Vault.contract, event: "UpdatePnl", logs: logs, sub: sub}, nil
}

// WatchUpdatePnl is a free log subscription operation binding the contract event 0x3ff41bdde87755b687ae83d0221a232b6be51a803330ed9661c1b5d0105e0d8a.
//
// Solidity: event UpdatePnl(bytes32 key, bool hasProfit, uint256 delta)
func (_Vault *VaultFilterer) WatchUpdatePnl(opts *bind.WatchOpts, sink chan<- *VaultUpdatePnl) (event.Subscription, error) {

	logs, sub, err := _Vault.contract.WatchLogs(opts, "UpdatePnl")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultUpdatePnl)
				if err := _Vault.contract.UnpackLog(event, "UpdatePnl", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpdatePnl is a log parse operation binding the contract event 0x3ff41bdde87755b687ae83d0221a232b6be51a803330ed9661c1b5d0105e0d8a.
//
// Solidity: event UpdatePnl(bytes32 key, bool hasProfit, uint256 delta)
func (_Vault *VaultFilterer) ParseUpdatePnl(log types.Log) (*VaultUpdatePnl, error) {
	event := new(VaultUpdatePnl)
	if err := _Vault.contract.UnpackLog(event, "UpdatePnl", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultUpdatePositionIterator is returned from FilterUpdatePosition and is used to iterate over the raw logs and unpacked data for UpdatePosition events raised by the Vault contract.
type VaultUpdatePositionIterator struct {
	Event *VaultUpdatePosition // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VaultUpdatePositionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultUpdatePosition)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VaultUpdatePosition)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VaultUpdatePositionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultUpdatePositionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultUpdatePosition represents a UpdatePosition event raised by the Vault contract.
type VaultUpdatePosition struct {
	Key              [32]byte
	Size             *big.Int
	Collateral       *big.Int
	AveragePrice     *big.Int
	EntryFundingRate *big.Int
	ReserveAmount    *big.Int
	RealisedPnl      *big.Int
	MarkPrice        *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterUpdatePosition is a free log retrieval operation binding the contract event 0x20853733b590dce729d9f4628682ebd9a34d2354e72679e66f43a008fc03b773.
//
// Solidity: event UpdatePosition(bytes32 key, uint256 size, uint256 collateral, uint256 averagePrice, uint256 entryFundingRate, uint256 reserveAmount, int256 realisedPnl, uint256 markPrice)
func (_Vault *VaultFilterer) FilterUpdatePosition(opts *bind.FilterOpts) (*VaultUpdatePositionIterator, error) {

	logs, sub, err := _Vault.contract.FilterLogs(opts, "UpdatePosition")
	if err != nil {
		return nil, err
	}
	return &VaultUpdatePositionIterator{contract: _Vault.contract, event: "UpdatePosition", logs: logs, sub: sub}, nil
}

// WatchUpdatePosition is a free log subscription operation binding the contract event 0x20853733b590dce729d9f4628682ebd9a34d2354e72679e66f43a008fc03b773.
//
// Solidity: event UpdatePosition(bytes32 key, uint256 size, uint256 collateral, uint256 averagePrice, uint256 entryFundingRate, uint256 reserveAmount, int256 realisedPnl, uint256 markPrice)
func (_Vault *VaultFilterer) WatchUpdatePosition(opts *bind.WatchOpts, sink chan<- *VaultUpdatePosition) (event.Subscription, error) {

	logs, sub, err := _Vault.contract.WatchLogs(opts, "UpdatePosition")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultUpdatePosition)
				if err := _Vault.contract.UnpackLog(event, "UpdatePosition", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpdatePosition is a log parse operation binding the contract event 0x20853733b590dce729d9f4628682ebd9a34d2354e72679e66f43a008fc03b773.
//
// Solidity: event UpdatePosition(bytes32 key, uint256 size, uint256 collateral, uint256 averagePrice, uint256 entryFundingRate, uint256 reserveAmount, int256 realisedPnl, uint256 markPrice)
func (_Vault *VaultFilterer) ParseUpdatePosition(log types.Log) (*VaultUpdatePosition, error) {
	event := new(VaultUpdatePosition)
	if err := _Vault.contract.UnpackLog(event, "UpdatePosition", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

