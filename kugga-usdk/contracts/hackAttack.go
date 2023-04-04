// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package hackAttack

import (
	"errors"
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
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// HackAttackMetaData contains all meta data concerning the HackAttack contract.
var HackAttackMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"idx\",\"type\":\"uint256\"}],\"name\":\"AddressEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bool\",\"name\":\"result\",\"type\":\"bool\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"mark\",\"type\":\"uint256\"}],\"name\":\"BoolEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"typeHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"idx\",\"type\":\"uint256\"}],\"name\":\"Bytes32Event\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"idx\",\"type\":\"uint256\"}],\"name\":\"BytesEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"v\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"idx\",\"type\":\"uint256\"}],\"name\":\"Uint256Event\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"usdk_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"bountyPldge_\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferSolarChain\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"cid_\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"reward_\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"rewardUnit_\",\"type\":\"string\"}],\"name\":\"pledge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"bountyCid\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"attackTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"attactWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receipt_\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// HackAttackABI is the input ABI used to generate the binding from.
// Deprecated: Use HackAttackMetaData.ABI instead.
var HackAttackABI = HackAttackMetaData.ABI

// HackAttack is an auto generated Go binding around an Ethereum contract.
type HackAttack struct {
	HackAttackCaller     // Read-only binding to the contract
	HackAttackTransactor // Write-only binding to the contract
	HackAttackFilterer   // Log filterer for contract events
}

// HackAttackCaller is an auto generated read-only Go binding around an Ethereum contract.
type HackAttackCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HackAttackTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HackAttackTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HackAttackFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HackAttackFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HackAttackSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HackAttackSession struct {
	Contract     *HackAttack       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HackAttackCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HackAttackCallerSession struct {
	Contract *HackAttackCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// HackAttackTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HackAttackTransactorSession struct {
	Contract     *HackAttackTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// HackAttackRaw is an auto generated low-level Go binding around an Ethereum contract.
type HackAttackRaw struct {
	Contract *HackAttack // Generic contract binding to access the raw methods on
}

// HackAttackCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HackAttackCallerRaw struct {
	Contract *HackAttackCaller // Generic read-only contract binding to access the raw methods on
}

// HackAttackTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HackAttackTransactorRaw struct {
	Contract *HackAttackTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHackAttack creates a new instance of HackAttack, bound to a specific deployed contract.
func NewHackAttack(address common.Address, backend bind.ContractBackend) (*HackAttack, error) {
	contract, err := bindHackAttack(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &HackAttack{HackAttackCaller: HackAttackCaller{contract: contract}, HackAttackTransactor: HackAttackTransactor{contract: contract}, HackAttackFilterer: HackAttackFilterer{contract: contract}}, nil
}

// NewHackAttackCaller creates a new read-only instance of HackAttack, bound to a specific deployed contract.
func NewHackAttackCaller(address common.Address, caller bind.ContractCaller) (*HackAttackCaller, error) {
	contract, err := bindHackAttack(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HackAttackCaller{contract: contract}, nil
}

// NewHackAttackTransactor creates a new write-only instance of HackAttack, bound to a specific deployed contract.
func NewHackAttackTransactor(address common.Address, transactor bind.ContractTransactor) (*HackAttackTransactor, error) {
	contract, err := bindHackAttack(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HackAttackTransactor{contract: contract}, nil
}

// NewHackAttackFilterer creates a new log filterer instance of HackAttack, bound to a specific deployed contract.
func NewHackAttackFilterer(address common.Address, filterer bind.ContractFilterer) (*HackAttackFilterer, error) {
	contract, err := bindHackAttack(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HackAttackFilterer{contract: contract}, nil
}

// bindHackAttack binds a generic wrapper to an already deployed contract.
func bindHackAttack(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(HackAttackABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HackAttack *HackAttackRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HackAttack.Contract.HackAttackCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HackAttack *HackAttackRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HackAttack.Contract.HackAttackTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HackAttack *HackAttackRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HackAttack.Contract.HackAttackTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HackAttack *HackAttackCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HackAttack.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HackAttack *HackAttackTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HackAttack.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HackAttack *HackAttackTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HackAttack.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_HackAttack *HackAttackCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _HackAttack.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_HackAttack *HackAttackSession) Owner() (common.Address, error) {
	return _HackAttack.Contract.Owner(&_HackAttack.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_HackAttack *HackAttackCallerSession) Owner() (common.Address, error) {
	return _HackAttack.Contract.Owner(&_HackAttack.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0xb759f954.
//
// Solidity: function approve(uint256 amount) payable returns()
func (_HackAttack *HackAttackTransactor) Approve(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _HackAttack.contract.Transact(opts, "approve", amount)
}

// Approve is a paid mutator transaction binding the contract method 0xb759f954.
//
// Solidity: function approve(uint256 amount) payable returns()
func (_HackAttack *HackAttackSession) Approve(amount *big.Int) (*types.Transaction, error) {
	return _HackAttack.Contract.Approve(&_HackAttack.TransactOpts, amount)
}

// Approve is a paid mutator transaction binding the contract method 0xb759f954.
//
// Solidity: function approve(uint256 amount) payable returns()
func (_HackAttack *HackAttackTransactorSession) Approve(amount *big.Int) (*types.Transaction, error) {
	return _HackAttack.Contract.Approve(&_HackAttack.TransactOpts, amount)
}

// AttackTransfer is a paid mutator transaction binding the contract method 0x95ecf3c6.
//
// Solidity: function attackTransfer(string bountyCid, uint256 amount) returns()
func (_HackAttack *HackAttackTransactor) AttackTransfer(opts *bind.TransactOpts, bountyCid string, amount *big.Int) (*types.Transaction, error) {
	return _HackAttack.contract.Transact(opts, "attackTransfer", bountyCid, amount)
}

// AttackTransfer is a paid mutator transaction binding the contract method 0x95ecf3c6.
//
// Solidity: function attackTransfer(string bountyCid, uint256 amount) returns()
func (_HackAttack *HackAttackSession) AttackTransfer(bountyCid string, amount *big.Int) (*types.Transaction, error) {
	return _HackAttack.Contract.AttackTransfer(&_HackAttack.TransactOpts, bountyCid, amount)
}

// AttackTransfer is a paid mutator transaction binding the contract method 0x95ecf3c6.
//
// Solidity: function attackTransfer(string bountyCid, uint256 amount) returns()
func (_HackAttack *HackAttackTransactorSession) AttackTransfer(bountyCid string, amount *big.Int) (*types.Transaction, error) {
	return _HackAttack.Contract.AttackTransfer(&_HackAttack.TransactOpts, bountyCid, amount)
}

// AttactWithdraw is a paid mutator transaction binding the contract method 0x979b7285.
//
// Solidity: function attactWithdraw() returns()
func (_HackAttack *HackAttackTransactor) AttactWithdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HackAttack.contract.Transact(opts, "attactWithdraw")
}

// AttactWithdraw is a paid mutator transaction binding the contract method 0x979b7285.
//
// Solidity: function attactWithdraw() returns()
func (_HackAttack *HackAttackSession) AttactWithdraw() (*types.Transaction, error) {
	return _HackAttack.Contract.AttactWithdraw(&_HackAttack.TransactOpts)
}

// AttactWithdraw is a paid mutator transaction binding the contract method 0x979b7285.
//
// Solidity: function attactWithdraw() returns()
func (_HackAttack *HackAttackTransactorSession) AttactWithdraw() (*types.Transaction, error) {
	return _HackAttack.Contract.AttactWithdraw(&_HackAttack.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address usdk_, address bountyPldge_) payable returns()
func (_HackAttack *HackAttackTransactor) Initialize(opts *bind.TransactOpts, usdk_ common.Address, bountyPldge_ common.Address) (*types.Transaction, error) {
	return _HackAttack.contract.Transact(opts, "initialize", usdk_, bountyPldge_)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address usdk_, address bountyPldge_) payable returns()
func (_HackAttack *HackAttackSession) Initialize(usdk_ common.Address, bountyPldge_ common.Address) (*types.Transaction, error) {
	return _HackAttack.Contract.Initialize(&_HackAttack.TransactOpts, usdk_, bountyPldge_)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address usdk_, address bountyPldge_) payable returns()
func (_HackAttack *HackAttackTransactorSession) Initialize(usdk_ common.Address, bountyPldge_ common.Address) (*types.Transaction, error) {
	return _HackAttack.Contract.Initialize(&_HackAttack.TransactOpts, usdk_, bountyPldge_)
}

// Pledge is a paid mutator transaction binding the contract method 0xac35b881.
//
// Solidity: function pledge(string cid_, uint256 reward_, string rewardUnit_) returns()
func (_HackAttack *HackAttackTransactor) Pledge(opts *bind.TransactOpts, cid_ string, reward_ *big.Int, rewardUnit_ string) (*types.Transaction, error) {
	return _HackAttack.contract.Transact(opts, "pledge", cid_, reward_, rewardUnit_)
}

// Pledge is a paid mutator transaction binding the contract method 0xac35b881.
//
// Solidity: function pledge(string cid_, uint256 reward_, string rewardUnit_) returns()
func (_HackAttack *HackAttackSession) Pledge(cid_ string, reward_ *big.Int, rewardUnit_ string) (*types.Transaction, error) {
	return _HackAttack.Contract.Pledge(&_HackAttack.TransactOpts, cid_, reward_, rewardUnit_)
}

// Pledge is a paid mutator transaction binding the contract method 0xac35b881.
//
// Solidity: function pledge(string cid_, uint256 reward_, string rewardUnit_) returns()
func (_HackAttack *HackAttackTransactorSession) Pledge(cid_ string, reward_ *big.Int, rewardUnit_ string) (*types.Transaction, error) {
	return _HackAttack.Contract.Pledge(&_HackAttack.TransactOpts, cid_, reward_, rewardUnit_)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_HackAttack *HackAttackTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HackAttack.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_HackAttack *HackAttackSession) RenounceOwnership() (*types.Transaction, error) {
	return _HackAttack.Contract.RenounceOwnership(&_HackAttack.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_HackAttack *HackAttackTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _HackAttack.Contract.RenounceOwnership(&_HackAttack.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_HackAttack *HackAttackTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _HackAttack.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_HackAttack *HackAttackSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _HackAttack.Contract.TransferOwnership(&_HackAttack.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_HackAttack *HackAttackTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _HackAttack.Contract.TransferOwnership(&_HackAttack.TransactOpts, newOwner)
}

// TransferSolarChain is a paid mutator transaction binding the contract method 0x73fd24fc.
//
// Solidity: function transferSolarChain(uint256 amount) payable returns()
func (_HackAttack *HackAttackTransactor) TransferSolarChain(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _HackAttack.contract.Transact(opts, "transferSolarChain", amount)
}

// TransferSolarChain is a paid mutator transaction binding the contract method 0x73fd24fc.
//
// Solidity: function transferSolarChain(uint256 amount) payable returns()
func (_HackAttack *HackAttackSession) TransferSolarChain(amount *big.Int) (*types.Transaction, error) {
	return _HackAttack.Contract.TransferSolarChain(&_HackAttack.TransactOpts, amount)
}

// TransferSolarChain is a paid mutator transaction binding the contract method 0x73fd24fc.
//
// Solidity: function transferSolarChain(uint256 amount) payable returns()
func (_HackAttack *HackAttackTransactorSession) TransferSolarChain(amount *big.Int) (*types.Transaction, error) {
	return _HackAttack.Contract.TransferSolarChain(&_HackAttack.TransactOpts, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address receipt_) returns()
func (_HackAttack *HackAttackTransactor) Withdraw(opts *bind.TransactOpts, receipt_ common.Address) (*types.Transaction, error) {
	return _HackAttack.contract.Transact(opts, "withdraw", receipt_)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address receipt_) returns()
func (_HackAttack *HackAttackSession) Withdraw(receipt_ common.Address) (*types.Transaction, error) {
	return _HackAttack.Contract.Withdraw(&_HackAttack.TransactOpts, receipt_)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address receipt_) returns()
func (_HackAttack *HackAttackTransactorSession) Withdraw(receipt_ common.Address) (*types.Transaction, error) {
	return _HackAttack.Contract.Withdraw(&_HackAttack.TransactOpts, receipt_)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_HackAttack *HackAttackTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _HackAttack.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_HackAttack *HackAttackSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _HackAttack.Contract.Fallback(&_HackAttack.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_HackAttack *HackAttackTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _HackAttack.Contract.Fallback(&_HackAttack.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_HackAttack *HackAttackTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HackAttack.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_HackAttack *HackAttackSession) Receive() (*types.Transaction, error) {
	return _HackAttack.Contract.Receive(&_HackAttack.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_HackAttack *HackAttackTransactorSession) Receive() (*types.Transaction, error) {
	return _HackAttack.Contract.Receive(&_HackAttack.TransactOpts)
}

// HackAttackAddressEventIterator is returned from FilterAddressEvent and is used to iterate over the raw logs and unpacked data for AddressEvent events raised by the HackAttack contract.
type HackAttackAddressEventIterator struct {
	Event *HackAttackAddressEvent // Event containing the contract specifics and raw log

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
func (it *HackAttackAddressEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HackAttackAddressEvent)
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
		it.Event = new(HackAttackAddressEvent)
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
func (it *HackAttackAddressEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HackAttackAddressEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HackAttackAddressEvent represents a AddressEvent event raised by the HackAttack contract.
type HackAttackAddressEvent struct {
	Sender common.Address
	Value  *big.Int
	Idx    *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterAddressEvent is a free log retrieval operation binding the contract event 0x73e26ef1bfac6b32eb0475f358f0de686fec3fd1096940b43f0de915905f6f79.
//
// Solidity: event AddressEvent(address indexed sender, uint256 indexed value, uint256 indexed idx)
func (_HackAttack *HackAttackFilterer) FilterAddressEvent(opts *bind.FilterOpts, sender []common.Address, value []*big.Int, idx []*big.Int) (*HackAttackAddressEventIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var valueRule []interface{}
	for _, valueItem := range value {
		valueRule = append(valueRule, valueItem)
	}
	var idxRule []interface{}
	for _, idxItem := range idx {
		idxRule = append(idxRule, idxItem)
	}

	logs, sub, err := _HackAttack.contract.FilterLogs(opts, "AddressEvent", senderRule, valueRule, idxRule)
	if err != nil {
		return nil, err
	}
	return &HackAttackAddressEventIterator{contract: _HackAttack.contract, event: "AddressEvent", logs: logs, sub: sub}, nil
}

// WatchAddressEvent is a free log subscription operation binding the contract event 0x73e26ef1bfac6b32eb0475f358f0de686fec3fd1096940b43f0de915905f6f79.
//
// Solidity: event AddressEvent(address indexed sender, uint256 indexed value, uint256 indexed idx)
func (_HackAttack *HackAttackFilterer) WatchAddressEvent(opts *bind.WatchOpts, sink chan<- *HackAttackAddressEvent, sender []common.Address, value []*big.Int, idx []*big.Int) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var valueRule []interface{}
	for _, valueItem := range value {
		valueRule = append(valueRule, valueItem)
	}
	var idxRule []interface{}
	for _, idxItem := range idx {
		idxRule = append(idxRule, idxItem)
	}

	logs, sub, err := _HackAttack.contract.WatchLogs(opts, "AddressEvent", senderRule, valueRule, idxRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HackAttackAddressEvent)
				if err := _HackAttack.contract.UnpackLog(event, "AddressEvent", log); err != nil {
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

// ParseAddressEvent is a log parse operation binding the contract event 0x73e26ef1bfac6b32eb0475f358f0de686fec3fd1096940b43f0de915905f6f79.
//
// Solidity: event AddressEvent(address indexed sender, uint256 indexed value, uint256 indexed idx)
func (_HackAttack *HackAttackFilterer) ParseAddressEvent(log types.Log) (*HackAttackAddressEvent, error) {
	event := new(HackAttackAddressEvent)
	if err := _HackAttack.contract.UnpackLog(event, "AddressEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HackAttackBoolEventIterator is returned from FilterBoolEvent and is used to iterate over the raw logs and unpacked data for BoolEvent events raised by the HackAttack contract.
type HackAttackBoolEventIterator struct {
	Event *HackAttackBoolEvent // Event containing the contract specifics and raw log

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
func (it *HackAttackBoolEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HackAttackBoolEvent)
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
		it.Event = new(HackAttackBoolEvent)
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
func (it *HackAttackBoolEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HackAttackBoolEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HackAttackBoolEvent represents a BoolEvent event raised by the HackAttack contract.
type HackAttackBoolEvent struct {
	Result bool
	Mark   *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBoolEvent is a free log retrieval operation binding the contract event 0x4c048f1faa88972c3b95075da7792557ce1e4e582059693989a2bd95abe1bdc3.
//
// Solidity: event BoolEvent(bool indexed result, uint256 indexed mark)
func (_HackAttack *HackAttackFilterer) FilterBoolEvent(opts *bind.FilterOpts, result []bool, mark []*big.Int) (*HackAttackBoolEventIterator, error) {

	var resultRule []interface{}
	for _, resultItem := range result {
		resultRule = append(resultRule, resultItem)
	}
	var markRule []interface{}
	for _, markItem := range mark {
		markRule = append(markRule, markItem)
	}

	logs, sub, err := _HackAttack.contract.FilterLogs(opts, "BoolEvent", resultRule, markRule)
	if err != nil {
		return nil, err
	}
	return &HackAttackBoolEventIterator{contract: _HackAttack.contract, event: "BoolEvent", logs: logs, sub: sub}, nil
}

// WatchBoolEvent is a free log subscription operation binding the contract event 0x4c048f1faa88972c3b95075da7792557ce1e4e582059693989a2bd95abe1bdc3.
//
// Solidity: event BoolEvent(bool indexed result, uint256 indexed mark)
func (_HackAttack *HackAttackFilterer) WatchBoolEvent(opts *bind.WatchOpts, sink chan<- *HackAttackBoolEvent, result []bool, mark []*big.Int) (event.Subscription, error) {

	var resultRule []interface{}
	for _, resultItem := range result {
		resultRule = append(resultRule, resultItem)
	}
	var markRule []interface{}
	for _, markItem := range mark {
		markRule = append(markRule, markItem)
	}

	logs, sub, err := _HackAttack.contract.WatchLogs(opts, "BoolEvent", resultRule, markRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HackAttackBoolEvent)
				if err := _HackAttack.contract.UnpackLog(event, "BoolEvent", log); err != nil {
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

// ParseBoolEvent is a log parse operation binding the contract event 0x4c048f1faa88972c3b95075da7792557ce1e4e582059693989a2bd95abe1bdc3.
//
// Solidity: event BoolEvent(bool indexed result, uint256 indexed mark)
func (_HackAttack *HackAttackFilterer) ParseBoolEvent(log types.Log) (*HackAttackBoolEvent, error) {
	event := new(HackAttackBoolEvent)
	if err := _HackAttack.contract.UnpackLog(event, "BoolEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HackAttackBytes32EventIterator is returned from FilterBytes32Event and is used to iterate over the raw logs and unpacked data for Bytes32Event events raised by the HackAttack contract.
type HackAttackBytes32EventIterator struct {
	Event *HackAttackBytes32Event // Event containing the contract specifics and raw log

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
func (it *HackAttackBytes32EventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HackAttackBytes32Event)
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
		it.Event = new(HackAttackBytes32Event)
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
func (it *HackAttackBytes32EventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HackAttackBytes32EventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HackAttackBytes32Event represents a Bytes32Event event raised by the HackAttack contract.
type HackAttackBytes32Event struct {
	TypeHash [32]byte
	Idx      *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterBytes32Event is a free log retrieval operation binding the contract event 0x7a2e7f2872791f9e379e4000877dd090a80304cc08f547e570985984ec10dbc0.
//
// Solidity: event Bytes32Event(bytes32 indexed typeHash, uint256 indexed idx)
func (_HackAttack *HackAttackFilterer) FilterBytes32Event(opts *bind.FilterOpts, typeHash [][32]byte, idx []*big.Int) (*HackAttackBytes32EventIterator, error) {

	var typeHashRule []interface{}
	for _, typeHashItem := range typeHash {
		typeHashRule = append(typeHashRule, typeHashItem)
	}
	var idxRule []interface{}
	for _, idxItem := range idx {
		idxRule = append(idxRule, idxItem)
	}

	logs, sub, err := _HackAttack.contract.FilterLogs(opts, "Bytes32Event", typeHashRule, idxRule)
	if err != nil {
		return nil, err
	}
	return &HackAttackBytes32EventIterator{contract: _HackAttack.contract, event: "Bytes32Event", logs: logs, sub: sub}, nil
}

// WatchBytes32Event is a free log subscription operation binding the contract event 0x7a2e7f2872791f9e379e4000877dd090a80304cc08f547e570985984ec10dbc0.
//
// Solidity: event Bytes32Event(bytes32 indexed typeHash, uint256 indexed idx)
func (_HackAttack *HackAttackFilterer) WatchBytes32Event(opts *bind.WatchOpts, sink chan<- *HackAttackBytes32Event, typeHash [][32]byte, idx []*big.Int) (event.Subscription, error) {

	var typeHashRule []interface{}
	for _, typeHashItem := range typeHash {
		typeHashRule = append(typeHashRule, typeHashItem)
	}
	var idxRule []interface{}
	for _, idxItem := range idx {
		idxRule = append(idxRule, idxItem)
	}

	logs, sub, err := _HackAttack.contract.WatchLogs(opts, "Bytes32Event", typeHashRule, idxRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HackAttackBytes32Event)
				if err := _HackAttack.contract.UnpackLog(event, "Bytes32Event", log); err != nil {
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

// ParseBytes32Event is a log parse operation binding the contract event 0x7a2e7f2872791f9e379e4000877dd090a80304cc08f547e570985984ec10dbc0.
//
// Solidity: event Bytes32Event(bytes32 indexed typeHash, uint256 indexed idx)
func (_HackAttack *HackAttackFilterer) ParseBytes32Event(log types.Log) (*HackAttackBytes32Event, error) {
	event := new(HackAttackBytes32Event)
	if err := _HackAttack.contract.UnpackLog(event, "Bytes32Event", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HackAttackBytesEventIterator is returned from FilterBytesEvent and is used to iterate over the raw logs and unpacked data for BytesEvent events raised by the HackAttack contract.
type HackAttackBytesEventIterator struct {
	Event *HackAttackBytesEvent // Event containing the contract specifics and raw log

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
func (it *HackAttackBytesEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HackAttackBytesEvent)
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
		it.Event = new(HackAttackBytesEvent)
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
func (it *HackAttackBytesEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HackAttackBytesEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HackAttackBytesEvent represents a BytesEvent event raised by the HackAttack contract.
type HackAttackBytesEvent struct {
	Data []byte
	Idx  *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterBytesEvent is a free log retrieval operation binding the contract event 0xad18c571ec138d35e20a584c2f3f9ecc206da98839156635a9b62a9e5989832f.
//
// Solidity: event BytesEvent(bytes data, uint256 indexed idx)
func (_HackAttack *HackAttackFilterer) FilterBytesEvent(opts *bind.FilterOpts, idx []*big.Int) (*HackAttackBytesEventIterator, error) {

	var idxRule []interface{}
	for _, idxItem := range idx {
		idxRule = append(idxRule, idxItem)
	}

	logs, sub, err := _HackAttack.contract.FilterLogs(opts, "BytesEvent", idxRule)
	if err != nil {
		return nil, err
	}
	return &HackAttackBytesEventIterator{contract: _HackAttack.contract, event: "BytesEvent", logs: logs, sub: sub}, nil
}

// WatchBytesEvent is a free log subscription operation binding the contract event 0xad18c571ec138d35e20a584c2f3f9ecc206da98839156635a9b62a9e5989832f.
//
// Solidity: event BytesEvent(bytes data, uint256 indexed idx)
func (_HackAttack *HackAttackFilterer) WatchBytesEvent(opts *bind.WatchOpts, sink chan<- *HackAttackBytesEvent, idx []*big.Int) (event.Subscription, error) {

	var idxRule []interface{}
	for _, idxItem := range idx {
		idxRule = append(idxRule, idxItem)
	}

	logs, sub, err := _HackAttack.contract.WatchLogs(opts, "BytesEvent", idxRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HackAttackBytesEvent)
				if err := _HackAttack.contract.UnpackLog(event, "BytesEvent", log); err != nil {
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

// ParseBytesEvent is a log parse operation binding the contract event 0xad18c571ec138d35e20a584c2f3f9ecc206da98839156635a9b62a9e5989832f.
//
// Solidity: event BytesEvent(bytes data, uint256 indexed idx)
func (_HackAttack *HackAttackFilterer) ParseBytesEvent(log types.Log) (*HackAttackBytesEvent, error) {
	event := new(HackAttackBytesEvent)
	if err := _HackAttack.contract.UnpackLog(event, "BytesEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HackAttackInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the HackAttack contract.
type HackAttackInitializedIterator struct {
	Event *HackAttackInitialized // Event containing the contract specifics and raw log

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
func (it *HackAttackInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HackAttackInitialized)
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
		it.Event = new(HackAttackInitialized)
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
func (it *HackAttackInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HackAttackInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HackAttackInitialized represents a Initialized event raised by the HackAttack contract.
type HackAttackInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_HackAttack *HackAttackFilterer) FilterInitialized(opts *bind.FilterOpts) (*HackAttackInitializedIterator, error) {

	logs, sub, err := _HackAttack.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &HackAttackInitializedIterator{contract: _HackAttack.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_HackAttack *HackAttackFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *HackAttackInitialized) (event.Subscription, error) {

	logs, sub, err := _HackAttack.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HackAttackInitialized)
				if err := _HackAttack.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_HackAttack *HackAttackFilterer) ParseInitialized(log types.Log) (*HackAttackInitialized, error) {
	event := new(HackAttackInitialized)
	if err := _HackAttack.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HackAttackOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the HackAttack contract.
type HackAttackOwnershipTransferredIterator struct {
	Event *HackAttackOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *HackAttackOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HackAttackOwnershipTransferred)
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
		it.Event = new(HackAttackOwnershipTransferred)
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
func (it *HackAttackOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HackAttackOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HackAttackOwnershipTransferred represents a OwnershipTransferred event raised by the HackAttack contract.
type HackAttackOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_HackAttack *HackAttackFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*HackAttackOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _HackAttack.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &HackAttackOwnershipTransferredIterator{contract: _HackAttack.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_HackAttack *HackAttackFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *HackAttackOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _HackAttack.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HackAttackOwnershipTransferred)
				if err := _HackAttack.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_HackAttack *HackAttackFilterer) ParseOwnershipTransferred(log types.Log) (*HackAttackOwnershipTransferred, error) {
	event := new(HackAttackOwnershipTransferred)
	if err := _HackAttack.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HackAttackUint256EventIterator is returned from FilterUint256Event and is used to iterate over the raw logs and unpacked data for Uint256Event events raised by the HackAttack contract.
type HackAttackUint256EventIterator struct {
	Event *HackAttackUint256Event // Event containing the contract specifics and raw log

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
func (it *HackAttackUint256EventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HackAttackUint256Event)
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
		it.Event = new(HackAttackUint256Event)
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
func (it *HackAttackUint256EventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HackAttackUint256EventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HackAttackUint256Event represents a Uint256Event event raised by the HackAttack contract.
type HackAttackUint256Event struct {
	V   *big.Int
	Idx *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterUint256Event is a free log retrieval operation binding the contract event 0x04ab70fac4c593e7954b088c03560c671d7f64602e32f846aeb0f9ad5a3f6227.
//
// Solidity: event Uint256Event(uint256 indexed v, uint256 indexed idx)
func (_HackAttack *HackAttackFilterer) FilterUint256Event(opts *bind.FilterOpts, v []*big.Int, idx []*big.Int) (*HackAttackUint256EventIterator, error) {

	var vRule []interface{}
	for _, vItem := range v {
		vRule = append(vRule, vItem)
	}
	var idxRule []interface{}
	for _, idxItem := range idx {
		idxRule = append(idxRule, idxItem)
	}

	logs, sub, err := _HackAttack.contract.FilterLogs(opts, "Uint256Event", vRule, idxRule)
	if err != nil {
		return nil, err
	}
	return &HackAttackUint256EventIterator{contract: _HackAttack.contract, event: "Uint256Event", logs: logs, sub: sub}, nil
}

// WatchUint256Event is a free log subscription operation binding the contract event 0x04ab70fac4c593e7954b088c03560c671d7f64602e32f846aeb0f9ad5a3f6227.
//
// Solidity: event Uint256Event(uint256 indexed v, uint256 indexed idx)
func (_HackAttack *HackAttackFilterer) WatchUint256Event(opts *bind.WatchOpts, sink chan<- *HackAttackUint256Event, v []*big.Int, idx []*big.Int) (event.Subscription, error) {

	var vRule []interface{}
	for _, vItem := range v {
		vRule = append(vRule, vItem)
	}
	var idxRule []interface{}
	for _, idxItem := range idx {
		idxRule = append(idxRule, idxItem)
	}

	logs, sub, err := _HackAttack.contract.WatchLogs(opts, "Uint256Event", vRule, idxRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HackAttackUint256Event)
				if err := _HackAttack.contract.UnpackLog(event, "Uint256Event", log); err != nil {
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

// ParseUint256Event is a log parse operation binding the contract event 0x04ab70fac4c593e7954b088c03560c671d7f64602e32f846aeb0f9ad5a3f6227.
//
// Solidity: event Uint256Event(uint256 indexed v, uint256 indexed idx)
func (_HackAttack *HackAttackFilterer) ParseUint256Event(log types.Log) (*HackAttackUint256Event, error) {
	event := new(HackAttackUint256Event)
	if err := _HackAttack.contract.UnpackLog(event, "Uint256Event", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
