// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package solarTokenCoin

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

// SolarTokenCoinMetaData contains all meta data concerning the SolarTokenCoin contract.
var SolarTokenCoinMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_buyer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"MintTokens\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_buyer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_seller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_saleOrderId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_purchasePrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_topic\",\"type\":\"string\"}],\"name\":\"PurchasePromise\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tradeId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"feePercent\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sellerIncome\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"organiserIncome\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"topic\",\"type\":\"string\"}],\"name\":\"Trade\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"TransferDirectSolar\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"TransferTokens\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[],\"name\":\"MINTER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"grantMintRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasMintRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"mintTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceMintRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"decimal\",\"type\":\"uint8\"}],\"name\":\"setDecimal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transferTokens\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractICellarNftSaleInterface\",\"name\":\"solarNftSale\",\"type\":\"address\"},{\"internalType\":\"contractICellarNft1155SaleInterface\",\"name\":\"cellarNft1155Sale\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimal\",\"type\":\"uint8\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spenderERC721\",\"type\":\"address\"}],\"name\":\"setSpenderERC721\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spenderERC1155\",\"type\":\"address\"}],\"name\":\"setSpenderERC1155\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"organiser\",\"type\":\"address\"}],\"name\":\"setOrganiser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSpenderERC721\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[],\"name\":\"getSpenderERC1155\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[],\"name\":\"getOrganiser\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"purchasePrice\",\"type\":\"uint256\"}],\"name\":\"increase721Allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"}],\"name\":\"removeAllowanceOrderItem\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"purchasePrice\",\"type\":\"uint256\"}],\"name\":\"increase1155Allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// SolarTokenCoinABI is the input ABI used to generate the binding from.
// Deprecated: Use SolarTokenCoinMetaData.ABI instead.
var SolarTokenCoinABI = SolarTokenCoinMetaData.ABI

// SolarTokenCoin is an auto generated Go binding around an Ethereum contract.
type SolarTokenCoin struct {
	SolarTokenCoinCaller     // Read-only binding to the contract
	SolarTokenCoinTransactor // Write-only binding to the contract
	SolarTokenCoinFilterer   // Log filterer for contract events
}

// SolarTokenCoinCaller is an auto generated read-only Go binding around an Ethereum contract.
type SolarTokenCoinCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SolarTokenCoinTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SolarTokenCoinTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SolarTokenCoinFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SolarTokenCoinFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SolarTokenCoinSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SolarTokenCoinSession struct {
	Contract     *SolarTokenCoin   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SolarTokenCoinCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SolarTokenCoinCallerSession struct {
	Contract *SolarTokenCoinCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// SolarTokenCoinTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SolarTokenCoinTransactorSession struct {
	Contract     *SolarTokenCoinTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// SolarTokenCoinRaw is an auto generated low-level Go binding around an Ethereum contract.
type SolarTokenCoinRaw struct {
	Contract *SolarTokenCoin // Generic contract binding to access the raw methods on
}

// SolarTokenCoinCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SolarTokenCoinCallerRaw struct {
	Contract *SolarTokenCoinCaller // Generic read-only contract binding to access the raw methods on
}

// SolarTokenCoinTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SolarTokenCoinTransactorRaw struct {
	Contract *SolarTokenCoinTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSolarTokenCoin creates a new instance of SolarTokenCoin, bound to a specific deployed contract.
func NewSolarTokenCoin(address common.Address, backend bind.ContractBackend) (*SolarTokenCoin, error) {
	contract, err := bindSolarTokenCoin(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SolarTokenCoin{SolarTokenCoinCaller: SolarTokenCoinCaller{contract: contract}, SolarTokenCoinTransactor: SolarTokenCoinTransactor{contract: contract}, SolarTokenCoinFilterer: SolarTokenCoinFilterer{contract: contract}}, nil
}

// NewSolarTokenCoinCaller creates a new read-only instance of SolarTokenCoin, bound to a specific deployed contract.
func NewSolarTokenCoinCaller(address common.Address, caller bind.ContractCaller) (*SolarTokenCoinCaller, error) {
	contract, err := bindSolarTokenCoin(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SolarTokenCoinCaller{contract: contract}, nil
}

// NewSolarTokenCoinTransactor creates a new write-only instance of SolarTokenCoin, bound to a specific deployed contract.
func NewSolarTokenCoinTransactor(address common.Address, transactor bind.ContractTransactor) (*SolarTokenCoinTransactor, error) {
	contract, err := bindSolarTokenCoin(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SolarTokenCoinTransactor{contract: contract}, nil
}

// NewSolarTokenCoinFilterer creates a new log filterer instance of SolarTokenCoin, bound to a specific deployed contract.
func NewSolarTokenCoinFilterer(address common.Address, filterer bind.ContractFilterer) (*SolarTokenCoinFilterer, error) {
	contract, err := bindSolarTokenCoin(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SolarTokenCoinFilterer{contract: contract}, nil
}

// bindSolarTokenCoin binds a generic wrapper to an already deployed contract.
func bindSolarTokenCoin(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SolarTokenCoinABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SolarTokenCoin *SolarTokenCoinRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SolarTokenCoin.Contract.SolarTokenCoinCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SolarTokenCoin *SolarTokenCoinRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SolarTokenCoin.Contract.SolarTokenCoinTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SolarTokenCoin *SolarTokenCoinRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SolarTokenCoin.Contract.SolarTokenCoinTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SolarTokenCoin *SolarTokenCoinCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SolarTokenCoin.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SolarTokenCoin *SolarTokenCoinTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SolarTokenCoin.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SolarTokenCoin *SolarTokenCoinTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SolarTokenCoin.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_SolarTokenCoin *SolarTokenCoinCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _SolarTokenCoin.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_SolarTokenCoin *SolarTokenCoinSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _SolarTokenCoin.Contract.DEFAULTADMINROLE(&_SolarTokenCoin.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_SolarTokenCoin *SolarTokenCoinCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _SolarTokenCoin.Contract.DEFAULTADMINROLE(&_SolarTokenCoin.CallOpts)
}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_SolarTokenCoin *SolarTokenCoinCaller) MINTERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _SolarTokenCoin.contract.Call(opts, &out, "MINTER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_SolarTokenCoin *SolarTokenCoinSession) MINTERROLE() ([32]byte, error) {
	return _SolarTokenCoin.Contract.MINTERROLE(&_SolarTokenCoin.CallOpts)
}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_SolarTokenCoin *SolarTokenCoinCallerSession) MINTERROLE() ([32]byte, error) {
	return _SolarTokenCoin.Contract.MINTERROLE(&_SolarTokenCoin.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_SolarTokenCoin *SolarTokenCoinCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SolarTokenCoin.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_SolarTokenCoin *SolarTokenCoinSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _SolarTokenCoin.Contract.Allowance(&_SolarTokenCoin.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_SolarTokenCoin *SolarTokenCoinCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _SolarTokenCoin.Contract.Allowance(&_SolarTokenCoin.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_SolarTokenCoin *SolarTokenCoinCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SolarTokenCoin.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_SolarTokenCoin *SolarTokenCoinSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _SolarTokenCoin.Contract.BalanceOf(&_SolarTokenCoin.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_SolarTokenCoin *SolarTokenCoinCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _SolarTokenCoin.Contract.BalanceOf(&_SolarTokenCoin.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_SolarTokenCoin *SolarTokenCoinCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _SolarTokenCoin.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_SolarTokenCoin *SolarTokenCoinSession) Decimals() (uint8, error) {
	return _SolarTokenCoin.Contract.Decimals(&_SolarTokenCoin.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_SolarTokenCoin *SolarTokenCoinCallerSession) Decimals() (uint8, error) {
	return _SolarTokenCoin.Contract.Decimals(&_SolarTokenCoin.CallOpts)
}

// GetOrganiser is a free data retrieval call binding the contract method 0x0b705578.
//
// Solidity: function getOrganiser() view returns(address)
func (_SolarTokenCoin *SolarTokenCoinCaller) GetOrganiser(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SolarTokenCoin.contract.Call(opts, &out, "getOrganiser")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOrganiser is a free data retrieval call binding the contract method 0x0b705578.
//
// Solidity: function getOrganiser() view returns(address)
func (_SolarTokenCoin *SolarTokenCoinSession) GetOrganiser() (common.Address, error) {
	return _SolarTokenCoin.Contract.GetOrganiser(&_SolarTokenCoin.CallOpts)
}

// GetOrganiser is a free data retrieval call binding the contract method 0x0b705578.
//
// Solidity: function getOrganiser() view returns(address)
func (_SolarTokenCoin *SolarTokenCoinCallerSession) GetOrganiser() (common.Address, error) {
	return _SolarTokenCoin.Contract.GetOrganiser(&_SolarTokenCoin.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_SolarTokenCoin *SolarTokenCoinCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _SolarTokenCoin.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_SolarTokenCoin *SolarTokenCoinSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _SolarTokenCoin.Contract.GetRoleAdmin(&_SolarTokenCoin.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_SolarTokenCoin *SolarTokenCoinCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _SolarTokenCoin.Contract.GetRoleAdmin(&_SolarTokenCoin.CallOpts, role)
}

// GetSpenderERC1155 is a free data retrieval call binding the contract method 0x0d61c96e.
//
// Solidity: function getSpenderERC1155() view returns(address)
func (_SolarTokenCoin *SolarTokenCoinCaller) GetSpenderERC1155(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SolarTokenCoin.contract.Call(opts, &out, "getSpenderERC1155")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetSpenderERC1155 is a free data retrieval call binding the contract method 0x0d61c96e.
//
// Solidity: function getSpenderERC1155() view returns(address)
func (_SolarTokenCoin *SolarTokenCoinSession) GetSpenderERC1155() (common.Address, error) {
	return _SolarTokenCoin.Contract.GetSpenderERC1155(&_SolarTokenCoin.CallOpts)
}

// GetSpenderERC1155 is a free data retrieval call binding the contract method 0x0d61c96e.
//
// Solidity: function getSpenderERC1155() view returns(address)
func (_SolarTokenCoin *SolarTokenCoinCallerSession) GetSpenderERC1155() (common.Address, error) {
	return _SolarTokenCoin.Contract.GetSpenderERC1155(&_SolarTokenCoin.CallOpts)
}

// GetSpenderERC721 is a free data retrieval call binding the contract method 0x3bd8cb38.
//
// Solidity: function getSpenderERC721() view returns(address)
func (_SolarTokenCoin *SolarTokenCoinCaller) GetSpenderERC721(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SolarTokenCoin.contract.Call(opts, &out, "getSpenderERC721")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetSpenderERC721 is a free data retrieval call binding the contract method 0x3bd8cb38.
//
// Solidity: function getSpenderERC721() view returns(address)
func (_SolarTokenCoin *SolarTokenCoinSession) GetSpenderERC721() (common.Address, error) {
	return _SolarTokenCoin.Contract.GetSpenderERC721(&_SolarTokenCoin.CallOpts)
}

// GetSpenderERC721 is a free data retrieval call binding the contract method 0x3bd8cb38.
//
// Solidity: function getSpenderERC721() view returns(address)
func (_SolarTokenCoin *SolarTokenCoinCallerSession) GetSpenderERC721() (common.Address, error) {
	return _SolarTokenCoin.Contract.GetSpenderERC721(&_SolarTokenCoin.CallOpts)
}

// HasMintRole is a free data retrieval call binding the contract method 0xc4e58e54.
//
// Solidity: function hasMintRole(address account) view returns(bool)
func (_SolarTokenCoin *SolarTokenCoinCaller) HasMintRole(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _SolarTokenCoin.contract.Call(opts, &out, "hasMintRole", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasMintRole is a free data retrieval call binding the contract method 0xc4e58e54.
//
// Solidity: function hasMintRole(address account) view returns(bool)
func (_SolarTokenCoin *SolarTokenCoinSession) HasMintRole(account common.Address) (bool, error) {
	return _SolarTokenCoin.Contract.HasMintRole(&_SolarTokenCoin.CallOpts, account)
}

// HasMintRole is a free data retrieval call binding the contract method 0xc4e58e54.
//
// Solidity: function hasMintRole(address account) view returns(bool)
func (_SolarTokenCoin *SolarTokenCoinCallerSession) HasMintRole(account common.Address) (bool, error) {
	return _SolarTokenCoin.Contract.HasMintRole(&_SolarTokenCoin.CallOpts, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_SolarTokenCoin *SolarTokenCoinCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _SolarTokenCoin.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_SolarTokenCoin *SolarTokenCoinSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _SolarTokenCoin.Contract.HasRole(&_SolarTokenCoin.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_SolarTokenCoin *SolarTokenCoinCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _SolarTokenCoin.Contract.HasRole(&_SolarTokenCoin.CallOpts, role, account)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_SolarTokenCoin *SolarTokenCoinCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _SolarTokenCoin.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_SolarTokenCoin *SolarTokenCoinSession) Name() (string, error) {
	return _SolarTokenCoin.Contract.Name(&_SolarTokenCoin.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_SolarTokenCoin *SolarTokenCoinCallerSession) Name() (string, error) {
	return _SolarTokenCoin.Contract.Name(&_SolarTokenCoin.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SolarTokenCoin *SolarTokenCoinCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SolarTokenCoin.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SolarTokenCoin *SolarTokenCoinSession) Owner() (common.Address, error) {
	return _SolarTokenCoin.Contract.Owner(&_SolarTokenCoin.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SolarTokenCoin *SolarTokenCoinCallerSession) Owner() (common.Address, error) {
	return _SolarTokenCoin.Contract.Owner(&_SolarTokenCoin.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_SolarTokenCoin *SolarTokenCoinCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _SolarTokenCoin.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_SolarTokenCoin *SolarTokenCoinSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _SolarTokenCoin.Contract.SupportsInterface(&_SolarTokenCoin.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_SolarTokenCoin *SolarTokenCoinCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _SolarTokenCoin.Contract.SupportsInterface(&_SolarTokenCoin.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_SolarTokenCoin *SolarTokenCoinCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _SolarTokenCoin.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_SolarTokenCoin *SolarTokenCoinSession) Symbol() (string, error) {
	return _SolarTokenCoin.Contract.Symbol(&_SolarTokenCoin.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_SolarTokenCoin *SolarTokenCoinCallerSession) Symbol() (string, error) {
	return _SolarTokenCoin.Contract.Symbol(&_SolarTokenCoin.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_SolarTokenCoin *SolarTokenCoinCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SolarTokenCoin.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_SolarTokenCoin *SolarTokenCoinSession) TotalSupply() (*big.Int, error) {
	return _SolarTokenCoin.Contract.TotalSupply(&_SolarTokenCoin.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_SolarTokenCoin *SolarTokenCoinCallerSession) TotalSupply() (*big.Int, error) {
	return _SolarTokenCoin.Contract.TotalSupply(&_SolarTokenCoin.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_SolarTokenCoin *SolarTokenCoinTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SolarTokenCoin.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_SolarTokenCoin *SolarTokenCoinSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SolarTokenCoin.Contract.Approve(&_SolarTokenCoin.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_SolarTokenCoin *SolarTokenCoinTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SolarTokenCoin.Contract.Approve(&_SolarTokenCoin.TransactOpts, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_SolarTokenCoin *SolarTokenCoinTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _SolarTokenCoin.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_SolarTokenCoin *SolarTokenCoinSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _SolarTokenCoin.Contract.DecreaseAllowance(&_SolarTokenCoin.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_SolarTokenCoin *SolarTokenCoinTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _SolarTokenCoin.Contract.DecreaseAllowance(&_SolarTokenCoin.TransactOpts, spender, subtractedValue)
}

// GrantMintRole is a paid mutator transaction binding the contract method 0xc2e3273d.
//
// Solidity: function grantMintRole(address operator) returns()
func (_SolarTokenCoin *SolarTokenCoinTransactor) GrantMintRole(opts *bind.TransactOpts, operator common.Address) (*types.Transaction, error) {
	return _SolarTokenCoin.contract.Transact(opts, "grantMintRole", operator)
}

// GrantMintRole is a paid mutator transaction binding the contract method 0xc2e3273d.
//
// Solidity: function grantMintRole(address operator) returns()
func (_SolarTokenCoin *SolarTokenCoinSession) GrantMintRole(operator common.Address) (*types.Transaction, error) {
	return _SolarTokenCoin.Contract.GrantMintRole(&_SolarTokenCoin.TransactOpts, operator)
}

// GrantMintRole is a paid mutator transaction binding the contract method 0xc2e3273d.
//
// Solidity: function grantMintRole(address operator) returns()
func (_SolarTokenCoin *SolarTokenCoinTransactorSession) GrantMintRole(operator common.Address) (*types.Transaction, error) {
	return _SolarTokenCoin.Contract.GrantMintRole(&_SolarTokenCoin.TransactOpts, operator)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_SolarTokenCoin *SolarTokenCoinTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _SolarTokenCoin.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_SolarTokenCoin *SolarTokenCoinSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _SolarTokenCoin.Contract.GrantRole(&_SolarTokenCoin.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_SolarTokenCoin *SolarTokenCoinTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _SolarTokenCoin.Contract.GrantRole(&_SolarTokenCoin.TransactOpts, role, account)
}

// Increase1155Allowance is a paid mutator transaction binding the contract method 0xf5d57a29.
//
// Solidity: function increase1155Allowance(uint256 orderId, uint256 purchasePrice) returns(uint256)
func (_SolarTokenCoin *SolarTokenCoinTransactor) Increase1155Allowance(opts *bind.TransactOpts, orderId *big.Int, purchasePrice *big.Int) (*types.Transaction, error) {
	return _SolarTokenCoin.contract.Transact(opts, "increase1155Allowance", orderId, purchasePrice)
}

// Increase1155Allowance is a paid mutator transaction binding the contract method 0xf5d57a29.
//
// Solidity: function increase1155Allowance(uint256 orderId, uint256 purchasePrice) returns(uint256)
func (_SolarTokenCoin *SolarTokenCoinSession) Increase1155Allowance(orderId *big.Int, purchasePrice *big.Int) (*types.Transaction, error) {
	return _SolarTokenCoin.Contract.Increase1155Allowance(&_SolarTokenCoin.TransactOpts, orderId, purchasePrice)
}

// Increase1155Allowance is a paid mutator transaction binding the contract method 0xf5d57a29.
//
// Solidity: function increase1155Allowance(uint256 orderId, uint256 purchasePrice) returns(uint256)
func (_SolarTokenCoin *SolarTokenCoinTransactorSession) Increase1155Allowance(orderId *big.Int, purchasePrice *big.Int) (*types.Transaction, error) {
	return _SolarTokenCoin.Contract.Increase1155Allowance(&_SolarTokenCoin.TransactOpts, orderId, purchasePrice)
}

// Increase721Allowance is a paid mutator transaction binding the contract method 0xccc3a52f.
//
// Solidity: function increase721Allowance(uint256 orderId, uint256 purchasePrice) returns(uint256)
func (_SolarTokenCoin *SolarTokenCoinTransactor) Increase721Allowance(opts *bind.TransactOpts, orderId *big.Int, purchasePrice *big.Int) (*types.Transaction, error) {
	return _SolarTokenCoin.contract.Transact(opts, "increase721Allowance", orderId, purchasePrice)
}

// Increase721Allowance is a paid mutator transaction binding the contract method 0xccc3a52f.
//
// Solidity: function increase721Allowance(uint256 orderId, uint256 purchasePrice) returns(uint256)
func (_SolarTokenCoin *SolarTokenCoinSession) Increase721Allowance(orderId *big.Int, purchasePrice *big.Int) (*types.Transaction, error) {
	return _SolarTokenCoin.Contract.Increase721Allowance(&_SolarTokenCoin.TransactOpts, orderId, purchasePrice)
}

// Increase721Allowance is a paid mutator transaction binding the contract method 0xccc3a52f.
//
// Solidity: function increase721Allowance(uint256 orderId, uint256 purchasePrice) returns(uint256)
func (_SolarTokenCoin *SolarTokenCoinTransactorSession) Increase721Allowance(orderId *big.Int, purchasePrice *big.Int) (*types.Transaction, error) {
	return _SolarTokenCoin.Contract.Increase721Allowance(&_SolarTokenCoin.TransactOpts, orderId, purchasePrice)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_SolarTokenCoin *SolarTokenCoinTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _SolarTokenCoin.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_SolarTokenCoin *SolarTokenCoinSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _SolarTokenCoin.Contract.IncreaseAllowance(&_SolarTokenCoin.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_SolarTokenCoin *SolarTokenCoinTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _SolarTokenCoin.Contract.IncreaseAllowance(&_SolarTokenCoin.TransactOpts, spender, addedValue)
}

// Initialize is a paid mutator transaction binding the contract method 0x8420ce99.
//
// Solidity: function initialize(address solarNftSale, address cellarNft1155Sale, string name, string symbol, uint8 decimal) returns()
func (_SolarTokenCoin *SolarTokenCoinTransactor) Initialize(opts *bind.TransactOpts, solarNftSale common.Address, cellarNft1155Sale common.Address, name string, symbol string, decimal uint8) (*types.Transaction, error) {
	return _SolarTokenCoin.contract.Transact(opts, "initialize", solarNftSale, cellarNft1155Sale, name, symbol, decimal)
}

// Initialize is a paid mutator transaction binding the contract method 0x8420ce99.
//
// Solidity: function initialize(address solarNftSale, address cellarNft1155Sale, string name, string symbol, uint8 decimal) returns()
func (_SolarTokenCoin *SolarTokenCoinSession) Initialize(solarNftSale common.Address, cellarNft1155Sale common.Address, name string, symbol string, decimal uint8) (*types.Transaction, error) {
	return _SolarTokenCoin.Contract.Initialize(&_SolarTokenCoin.TransactOpts, solarNftSale, cellarNft1155Sale, name, symbol, decimal)
}

// Initialize is a paid mutator transaction binding the contract method 0x8420ce99.
//
// Solidity: function initialize(address solarNftSale, address cellarNft1155Sale, string name, string symbol, uint8 decimal) returns()
func (_SolarTokenCoin *SolarTokenCoinTransactorSession) Initialize(solarNftSale common.Address, cellarNft1155Sale common.Address, name string, symbol string, decimal uint8) (*types.Transaction, error) {
	return _SolarTokenCoin.Contract.Initialize(&_SolarTokenCoin.TransactOpts, solarNftSale, cellarNft1155Sale, name, symbol, decimal)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address account, uint256 amount) returns(bool)
func (_SolarTokenCoin *SolarTokenCoinTransactor) Mint(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SolarTokenCoin.contract.Transact(opts, "mint", account, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address account, uint256 amount) returns(bool)
func (_SolarTokenCoin *SolarTokenCoinSession) Mint(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SolarTokenCoin.Contract.Mint(&_SolarTokenCoin.TransactOpts, account, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address account, uint256 amount) returns(bool)
func (_SolarTokenCoin *SolarTokenCoinTransactorSession) Mint(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SolarTokenCoin.Contract.Mint(&_SolarTokenCoin.TransactOpts, account, amount)
}

// MintTokens is a paid mutator transaction binding the contract method 0xf0dda65c.
//
// Solidity: function mintTokens(address _recipient, uint256 _amount) returns()
func (_SolarTokenCoin *SolarTokenCoinTransactor) MintTokens(opts *bind.TransactOpts, _recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _SolarTokenCoin.contract.Transact(opts, "mintTokens", _recipient, _amount)
}

// MintTokens is a paid mutator transaction binding the contract method 0xf0dda65c.
//
// Solidity: function mintTokens(address _recipient, uint256 _amount) returns()
func (_SolarTokenCoin *SolarTokenCoinSession) MintTokens(_recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _SolarTokenCoin.Contract.MintTokens(&_SolarTokenCoin.TransactOpts, _recipient, _amount)
}

// MintTokens is a paid mutator transaction binding the contract method 0xf0dda65c.
//
// Solidity: function mintTokens(address _recipient, uint256 _amount) returns()
func (_SolarTokenCoin *SolarTokenCoinTransactorSession) MintTokens(_recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _SolarTokenCoin.Contract.MintTokens(&_SolarTokenCoin.TransactOpts, _recipient, _amount)
}

// RemoveAllowanceOrderItem is a paid mutator transaction binding the contract method 0x02e2fcdd.
//
// Solidity: function removeAllowanceOrderItem(uint256 orderId) returns()
func (_SolarTokenCoin *SolarTokenCoinTransactor) RemoveAllowanceOrderItem(opts *bind.TransactOpts, orderId *big.Int) (*types.Transaction, error) {
	return _SolarTokenCoin.contract.Transact(opts, "removeAllowanceOrderItem", orderId)
}

// RemoveAllowanceOrderItem is a paid mutator transaction binding the contract method 0x02e2fcdd.
//
// Solidity: function removeAllowanceOrderItem(uint256 orderId) returns()
func (_SolarTokenCoin *SolarTokenCoinSession) RemoveAllowanceOrderItem(orderId *big.Int) (*types.Transaction, error) {
	return _SolarTokenCoin.Contract.RemoveAllowanceOrderItem(&_SolarTokenCoin.TransactOpts, orderId)
}

// RemoveAllowanceOrderItem is a paid mutator transaction binding the contract method 0x02e2fcdd.
//
// Solidity: function removeAllowanceOrderItem(uint256 orderId) returns()
func (_SolarTokenCoin *SolarTokenCoinTransactorSession) RemoveAllowanceOrderItem(orderId *big.Int) (*types.Transaction, error) {
	return _SolarTokenCoin.Contract.RemoveAllowanceOrderItem(&_SolarTokenCoin.TransactOpts, orderId)
}

// RenounceMintRole is a paid mutator transaction binding the contract method 0xa49a177a.
//
// Solidity: function renounceMintRole(address account) returns()
func (_SolarTokenCoin *SolarTokenCoinTransactor) RenounceMintRole(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _SolarTokenCoin.contract.Transact(opts, "renounceMintRole", account)
}

// RenounceMintRole is a paid mutator transaction binding the contract method 0xa49a177a.
//
// Solidity: function renounceMintRole(address account) returns()
func (_SolarTokenCoin *SolarTokenCoinSession) RenounceMintRole(account common.Address) (*types.Transaction, error) {
	return _SolarTokenCoin.Contract.RenounceMintRole(&_SolarTokenCoin.TransactOpts, account)
}

// RenounceMintRole is a paid mutator transaction binding the contract method 0xa49a177a.
//
// Solidity: function renounceMintRole(address account) returns()
func (_SolarTokenCoin *SolarTokenCoinTransactorSession) RenounceMintRole(account common.Address) (*types.Transaction, error) {
	return _SolarTokenCoin.Contract.RenounceMintRole(&_SolarTokenCoin.TransactOpts, account)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SolarTokenCoin *SolarTokenCoinTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SolarTokenCoin.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SolarTokenCoin *SolarTokenCoinSession) RenounceOwnership() (*types.Transaction, error) {
	return _SolarTokenCoin.Contract.RenounceOwnership(&_SolarTokenCoin.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SolarTokenCoin *SolarTokenCoinTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _SolarTokenCoin.Contract.RenounceOwnership(&_SolarTokenCoin.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_SolarTokenCoin *SolarTokenCoinTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _SolarTokenCoin.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_SolarTokenCoin *SolarTokenCoinSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _SolarTokenCoin.Contract.RenounceRole(&_SolarTokenCoin.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_SolarTokenCoin *SolarTokenCoinTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _SolarTokenCoin.Contract.RenounceRole(&_SolarTokenCoin.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_SolarTokenCoin *SolarTokenCoinTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _SolarTokenCoin.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_SolarTokenCoin *SolarTokenCoinSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _SolarTokenCoin.Contract.RevokeRole(&_SolarTokenCoin.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_SolarTokenCoin *SolarTokenCoinTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _SolarTokenCoin.Contract.RevokeRole(&_SolarTokenCoin.TransactOpts, role, account)
}

// SetDecimal is a paid mutator transaction binding the contract method 0x1a9a0426.
//
// Solidity: function setDecimal(uint8 decimal) returns()
func (_SolarTokenCoin *SolarTokenCoinTransactor) SetDecimal(opts *bind.TransactOpts, decimal uint8) (*types.Transaction, error) {
	return _SolarTokenCoin.contract.Transact(opts, "setDecimal", decimal)
}

// SetDecimal is a paid mutator transaction binding the contract method 0x1a9a0426.
//
// Solidity: function setDecimal(uint8 decimal) returns()
func (_SolarTokenCoin *SolarTokenCoinSession) SetDecimal(decimal uint8) (*types.Transaction, error) {
	return _SolarTokenCoin.Contract.SetDecimal(&_SolarTokenCoin.TransactOpts, decimal)
}

// SetDecimal is a paid mutator transaction binding the contract method 0x1a9a0426.
//
// Solidity: function setDecimal(uint8 decimal) returns()
func (_SolarTokenCoin *SolarTokenCoinTransactorSession) SetDecimal(decimal uint8) (*types.Transaction, error) {
	return _SolarTokenCoin.Contract.SetDecimal(&_SolarTokenCoin.TransactOpts, decimal)
}

// SetOrganiser is a paid mutator transaction binding the contract method 0x8c4b0e7e.
//
// Solidity: function setOrganiser(address organiser) returns()
func (_SolarTokenCoin *SolarTokenCoinTransactor) SetOrganiser(opts *bind.TransactOpts, organiser common.Address) (*types.Transaction, error) {
	return _SolarTokenCoin.contract.Transact(opts, "setOrganiser", organiser)
}

// SetOrganiser is a paid mutator transaction binding the contract method 0x8c4b0e7e.
//
// Solidity: function setOrganiser(address organiser) returns()
func (_SolarTokenCoin *SolarTokenCoinSession) SetOrganiser(organiser common.Address) (*types.Transaction, error) {
	return _SolarTokenCoin.Contract.SetOrganiser(&_SolarTokenCoin.TransactOpts, organiser)
}

// SetOrganiser is a paid mutator transaction binding the contract method 0x8c4b0e7e.
//
// Solidity: function setOrganiser(address organiser) returns()
func (_SolarTokenCoin *SolarTokenCoinTransactorSession) SetOrganiser(organiser common.Address) (*types.Transaction, error) {
	return _SolarTokenCoin.Contract.SetOrganiser(&_SolarTokenCoin.TransactOpts, organiser)
}

// SetSpenderERC1155 is a paid mutator transaction binding the contract method 0xfe71d675.
//
// Solidity: function setSpenderERC1155(address spenderERC1155) returns()
func (_SolarTokenCoin *SolarTokenCoinTransactor) SetSpenderERC1155(opts *bind.TransactOpts, spenderERC1155 common.Address) (*types.Transaction, error) {
	return _SolarTokenCoin.contract.Transact(opts, "setSpenderERC1155", spenderERC1155)
}

// SetSpenderERC1155 is a paid mutator transaction binding the contract method 0xfe71d675.
//
// Solidity: function setSpenderERC1155(address spenderERC1155) returns()
func (_SolarTokenCoin *SolarTokenCoinSession) SetSpenderERC1155(spenderERC1155 common.Address) (*types.Transaction, error) {
	return _SolarTokenCoin.Contract.SetSpenderERC1155(&_SolarTokenCoin.TransactOpts, spenderERC1155)
}

// SetSpenderERC1155 is a paid mutator transaction binding the contract method 0xfe71d675.
//
// Solidity: function setSpenderERC1155(address spenderERC1155) returns()
func (_SolarTokenCoin *SolarTokenCoinTransactorSession) SetSpenderERC1155(spenderERC1155 common.Address) (*types.Transaction, error) {
	return _SolarTokenCoin.Contract.SetSpenderERC1155(&_SolarTokenCoin.TransactOpts, spenderERC1155)
}

// SetSpenderERC721 is a paid mutator transaction binding the contract method 0x2aa7bab0.
//
// Solidity: function setSpenderERC721(address spenderERC721) returns()
func (_SolarTokenCoin *SolarTokenCoinTransactor) SetSpenderERC721(opts *bind.TransactOpts, spenderERC721 common.Address) (*types.Transaction, error) {
	return _SolarTokenCoin.contract.Transact(opts, "setSpenderERC721", spenderERC721)
}

// SetSpenderERC721 is a paid mutator transaction binding the contract method 0x2aa7bab0.
//
// Solidity: function setSpenderERC721(address spenderERC721) returns()
func (_SolarTokenCoin *SolarTokenCoinSession) SetSpenderERC721(spenderERC721 common.Address) (*types.Transaction, error) {
	return _SolarTokenCoin.Contract.SetSpenderERC721(&_SolarTokenCoin.TransactOpts, spenderERC721)
}

// SetSpenderERC721 is a paid mutator transaction binding the contract method 0x2aa7bab0.
//
// Solidity: function setSpenderERC721(address spenderERC721) returns()
func (_SolarTokenCoin *SolarTokenCoinTransactorSession) SetSpenderERC721(spenderERC721 common.Address) (*types.Transaction, error) {
	return _SolarTokenCoin.Contract.SetSpenderERC721(&_SolarTokenCoin.TransactOpts, spenderERC721)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_SolarTokenCoin *SolarTokenCoinTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SolarTokenCoin.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_SolarTokenCoin *SolarTokenCoinSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SolarTokenCoin.Contract.Transfer(&_SolarTokenCoin.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_SolarTokenCoin *SolarTokenCoinTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SolarTokenCoin.Contract.Transfer(&_SolarTokenCoin.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_SolarTokenCoin *SolarTokenCoinTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SolarTokenCoin.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_SolarTokenCoin *SolarTokenCoinSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SolarTokenCoin.Contract.TransferFrom(&_SolarTokenCoin.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_SolarTokenCoin *SolarTokenCoinTransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SolarTokenCoin.Contract.TransferFrom(&_SolarTokenCoin.TransactOpts, from, to, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SolarTokenCoin *SolarTokenCoinTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _SolarTokenCoin.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SolarTokenCoin *SolarTokenCoinSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SolarTokenCoin.Contract.TransferOwnership(&_SolarTokenCoin.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SolarTokenCoin *SolarTokenCoinTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SolarTokenCoin.Contract.TransferOwnership(&_SolarTokenCoin.TransactOpts, newOwner)
}

// TransferTokens is a paid mutator transaction binding the contract method 0xa64b6e5f.
//
// Solidity: function transferTokens(address _from, address _recipient, uint256 _amount) returns(bool)
func (_SolarTokenCoin *SolarTokenCoinTransactor) TransferTokens(opts *bind.TransactOpts, _from common.Address, _recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _SolarTokenCoin.contract.Transact(opts, "transferTokens", _from, _recipient, _amount)
}

// TransferTokens is a paid mutator transaction binding the contract method 0xa64b6e5f.
//
// Solidity: function transferTokens(address _from, address _recipient, uint256 _amount) returns(bool)
func (_SolarTokenCoin *SolarTokenCoinSession) TransferTokens(_from common.Address, _recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _SolarTokenCoin.Contract.TransferTokens(&_SolarTokenCoin.TransactOpts, _from, _recipient, _amount)
}

// TransferTokens is a paid mutator transaction binding the contract method 0xa64b6e5f.
//
// Solidity: function transferTokens(address _from, address _recipient, uint256 _amount) returns(bool)
func (_SolarTokenCoin *SolarTokenCoinTransactorSession) TransferTokens(_from common.Address, _recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _SolarTokenCoin.Contract.TransferTokens(&_SolarTokenCoin.TransactOpts, _from, _recipient, _amount)
}

// SolarTokenCoinApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the SolarTokenCoin contract.
type SolarTokenCoinApprovalIterator struct {
	Event *SolarTokenCoinApproval // Event containing the contract specifics and raw log

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
func (it *SolarTokenCoinApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SolarTokenCoinApproval)
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
		it.Event = new(SolarTokenCoinApproval)
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
func (it *SolarTokenCoinApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SolarTokenCoinApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SolarTokenCoinApproval represents a Approval event raised by the SolarTokenCoin contract.
type SolarTokenCoinApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_SolarTokenCoin *SolarTokenCoinFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*SolarTokenCoinApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _SolarTokenCoin.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &SolarTokenCoinApprovalIterator{contract: _SolarTokenCoin.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_SolarTokenCoin *SolarTokenCoinFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *SolarTokenCoinApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _SolarTokenCoin.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SolarTokenCoinApproval)
				if err := _SolarTokenCoin.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_SolarTokenCoin *SolarTokenCoinFilterer) ParseApproval(log types.Log) (*SolarTokenCoinApproval, error) {
	event := new(SolarTokenCoinApproval)
	if err := _SolarTokenCoin.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SolarTokenCoinMintTokensIterator is returned from FilterMintTokens and is used to iterate over the raw logs and unpacked data for MintTokens events raised by the SolarTokenCoin contract.
type SolarTokenCoinMintTokensIterator struct {
	Event *SolarTokenCoinMintTokens // Event containing the contract specifics and raw log

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
func (it *SolarTokenCoinMintTokensIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SolarTokenCoinMintTokens)
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
		it.Event = new(SolarTokenCoinMintTokens)
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
func (it *SolarTokenCoinMintTokensIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SolarTokenCoinMintTokensIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SolarTokenCoinMintTokens represents a MintTokens event raised by the SolarTokenCoin contract.
type SolarTokenCoinMintTokens struct {
	From   common.Address
	Buyer  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMintTokens is a free log retrieval operation binding the contract event 0x21f9c9a1a1f9a311a50f15fec5c1faa9e21fc9edf964f0fdecba5bd490484c5e.
//
// Solidity: event MintTokens(address indexed _from, address indexed _buyer, uint256 indexed _amount)
func (_SolarTokenCoin *SolarTokenCoinFilterer) FilterMintTokens(opts *bind.FilterOpts, _from []common.Address, _buyer []common.Address, _amount []*big.Int) (*SolarTokenCoinMintTokensIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _buyerRule []interface{}
	for _, _buyerItem := range _buyer {
		_buyerRule = append(_buyerRule, _buyerItem)
	}
	var _amountRule []interface{}
	for _, _amountItem := range _amount {
		_amountRule = append(_amountRule, _amountItem)
	}

	logs, sub, err := _SolarTokenCoin.contract.FilterLogs(opts, "MintTokens", _fromRule, _buyerRule, _amountRule)
	if err != nil {
		return nil, err
	}
	return &SolarTokenCoinMintTokensIterator{contract: _SolarTokenCoin.contract, event: "MintTokens", logs: logs, sub: sub}, nil
}

// WatchMintTokens is a free log subscription operation binding the contract event 0x21f9c9a1a1f9a311a50f15fec5c1faa9e21fc9edf964f0fdecba5bd490484c5e.
//
// Solidity: event MintTokens(address indexed _from, address indexed _buyer, uint256 indexed _amount)
func (_SolarTokenCoin *SolarTokenCoinFilterer) WatchMintTokens(opts *bind.WatchOpts, sink chan<- *SolarTokenCoinMintTokens, _from []common.Address, _buyer []common.Address, _amount []*big.Int) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _buyerRule []interface{}
	for _, _buyerItem := range _buyer {
		_buyerRule = append(_buyerRule, _buyerItem)
	}
	var _amountRule []interface{}
	for _, _amountItem := range _amount {
		_amountRule = append(_amountRule, _amountItem)
	}

	logs, sub, err := _SolarTokenCoin.contract.WatchLogs(opts, "MintTokens", _fromRule, _buyerRule, _amountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SolarTokenCoinMintTokens)
				if err := _SolarTokenCoin.contract.UnpackLog(event, "MintTokens", log); err != nil {
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

// ParseMintTokens is a log parse operation binding the contract event 0x21f9c9a1a1f9a311a50f15fec5c1faa9e21fc9edf964f0fdecba5bd490484c5e.
//
// Solidity: event MintTokens(address indexed _from, address indexed _buyer, uint256 indexed _amount)
func (_SolarTokenCoin *SolarTokenCoinFilterer) ParseMintTokens(log types.Log) (*SolarTokenCoinMintTokens, error) {
	event := new(SolarTokenCoinMintTokens)
	if err := _SolarTokenCoin.contract.UnpackLog(event, "MintTokens", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SolarTokenCoinOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the SolarTokenCoin contract.
type SolarTokenCoinOwnershipTransferredIterator struct {
	Event *SolarTokenCoinOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *SolarTokenCoinOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SolarTokenCoinOwnershipTransferred)
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
		it.Event = new(SolarTokenCoinOwnershipTransferred)
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
func (it *SolarTokenCoinOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SolarTokenCoinOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SolarTokenCoinOwnershipTransferred represents a OwnershipTransferred event raised by the SolarTokenCoin contract.
type SolarTokenCoinOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SolarTokenCoin *SolarTokenCoinFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*SolarTokenCoinOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SolarTokenCoin.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &SolarTokenCoinOwnershipTransferredIterator{contract: _SolarTokenCoin.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SolarTokenCoin *SolarTokenCoinFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *SolarTokenCoinOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SolarTokenCoin.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SolarTokenCoinOwnershipTransferred)
				if err := _SolarTokenCoin.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_SolarTokenCoin *SolarTokenCoinFilterer) ParseOwnershipTransferred(log types.Log) (*SolarTokenCoinOwnershipTransferred, error) {
	event := new(SolarTokenCoinOwnershipTransferred)
	if err := _SolarTokenCoin.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SolarTokenCoinPurchasePromiseIterator is returned from FilterPurchasePromise and is used to iterate over the raw logs and unpacked data for PurchasePromise events raised by the SolarTokenCoin contract.
type SolarTokenCoinPurchasePromiseIterator struct {
	Event *SolarTokenCoinPurchasePromise // Event containing the contract specifics and raw log

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
func (it *SolarTokenCoinPurchasePromiseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SolarTokenCoinPurchasePromise)
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
		it.Event = new(SolarTokenCoinPurchasePromise)
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
func (it *SolarTokenCoinPurchasePromiseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SolarTokenCoinPurchasePromiseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SolarTokenCoinPurchasePromise represents a PurchasePromise event raised by the SolarTokenCoin contract.
type SolarTokenCoinPurchasePromise struct {
	Buyer         common.Address
	Seller        common.Address
	TokenId       *big.Int
	SaleOrderId   *big.Int
	Price         *big.Int
	PurchasePrice *big.Int
	Amount        *big.Int
	Topic         string
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterPurchasePromise is a free log retrieval operation binding the contract event 0x722cb00b101d30c882a148444d3d516b21be84ecb29aa6e077d2a535a1da4882.
//
// Solidity: event PurchasePromise(address indexed _buyer, address _seller, uint256 indexed _tokenId, uint256 _saleOrderId, uint256 _price, uint256 indexed _purchasePrice, uint256 _amount, string _topic)
func (_SolarTokenCoin *SolarTokenCoinFilterer) FilterPurchasePromise(opts *bind.FilterOpts, _buyer []common.Address, _tokenId []*big.Int, _purchasePrice []*big.Int) (*SolarTokenCoinPurchasePromiseIterator, error) {

	var _buyerRule []interface{}
	for _, _buyerItem := range _buyer {
		_buyerRule = append(_buyerRule, _buyerItem)
	}

	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}

	var _purchasePriceRule []interface{}
	for _, _purchasePriceItem := range _purchasePrice {
		_purchasePriceRule = append(_purchasePriceRule, _purchasePriceItem)
	}

	logs, sub, err := _SolarTokenCoin.contract.FilterLogs(opts, "PurchasePromise", _buyerRule, _tokenIdRule, _purchasePriceRule)
	if err != nil {
		return nil, err
	}
	return &SolarTokenCoinPurchasePromiseIterator{contract: _SolarTokenCoin.contract, event: "PurchasePromise", logs: logs, sub: sub}, nil
}

// WatchPurchasePromise is a free log subscription operation binding the contract event 0x722cb00b101d30c882a148444d3d516b21be84ecb29aa6e077d2a535a1da4882.
//
// Solidity: event PurchasePromise(address indexed _buyer, address _seller, uint256 indexed _tokenId, uint256 _saleOrderId, uint256 _price, uint256 indexed _purchasePrice, uint256 _amount, string _topic)
func (_SolarTokenCoin *SolarTokenCoinFilterer) WatchPurchasePromise(opts *bind.WatchOpts, sink chan<- *SolarTokenCoinPurchasePromise, _buyer []common.Address, _tokenId []*big.Int, _purchasePrice []*big.Int) (event.Subscription, error) {

	var _buyerRule []interface{}
	for _, _buyerItem := range _buyer {
		_buyerRule = append(_buyerRule, _buyerItem)
	}

	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}

	var _purchasePriceRule []interface{}
	for _, _purchasePriceItem := range _purchasePrice {
		_purchasePriceRule = append(_purchasePriceRule, _purchasePriceItem)
	}

	logs, sub, err := _SolarTokenCoin.contract.WatchLogs(opts, "PurchasePromise", _buyerRule, _tokenIdRule, _purchasePriceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SolarTokenCoinPurchasePromise)
				if err := _SolarTokenCoin.contract.UnpackLog(event, "PurchasePromise", log); err != nil {
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

// ParsePurchasePromise is a log parse operation binding the contract event 0x722cb00b101d30c882a148444d3d516b21be84ecb29aa6e077d2a535a1da4882.
//
// Solidity: event PurchasePromise(address indexed _buyer, address _seller, uint256 indexed _tokenId, uint256 _saleOrderId, uint256 _price, uint256 indexed _purchasePrice, uint256 _amount, string _topic)
func (_SolarTokenCoin *SolarTokenCoinFilterer) ParsePurchasePromise(log types.Log) (*SolarTokenCoinPurchasePromise, error) {
	event := new(SolarTokenCoinPurchasePromise)
	if err := _SolarTokenCoin.contract.UnpackLog(event, "PurchasePromise", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SolarTokenCoinRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the SolarTokenCoin contract.
type SolarTokenCoinRoleAdminChangedIterator struct {
	Event *SolarTokenCoinRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *SolarTokenCoinRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SolarTokenCoinRoleAdminChanged)
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
		it.Event = new(SolarTokenCoinRoleAdminChanged)
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
func (it *SolarTokenCoinRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SolarTokenCoinRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SolarTokenCoinRoleAdminChanged represents a RoleAdminChanged event raised by the SolarTokenCoin contract.
type SolarTokenCoinRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_SolarTokenCoin *SolarTokenCoinFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*SolarTokenCoinRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _SolarTokenCoin.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &SolarTokenCoinRoleAdminChangedIterator{contract: _SolarTokenCoin.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_SolarTokenCoin *SolarTokenCoinFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *SolarTokenCoinRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _SolarTokenCoin.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SolarTokenCoinRoleAdminChanged)
				if err := _SolarTokenCoin.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_SolarTokenCoin *SolarTokenCoinFilterer) ParseRoleAdminChanged(log types.Log) (*SolarTokenCoinRoleAdminChanged, error) {
	event := new(SolarTokenCoinRoleAdminChanged)
	if err := _SolarTokenCoin.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SolarTokenCoinRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the SolarTokenCoin contract.
type SolarTokenCoinRoleGrantedIterator struct {
	Event *SolarTokenCoinRoleGranted // Event containing the contract specifics and raw log

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
func (it *SolarTokenCoinRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SolarTokenCoinRoleGranted)
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
		it.Event = new(SolarTokenCoinRoleGranted)
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
func (it *SolarTokenCoinRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SolarTokenCoinRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SolarTokenCoinRoleGranted represents a RoleGranted event raised by the SolarTokenCoin contract.
type SolarTokenCoinRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_SolarTokenCoin *SolarTokenCoinFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*SolarTokenCoinRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _SolarTokenCoin.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &SolarTokenCoinRoleGrantedIterator{contract: _SolarTokenCoin.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_SolarTokenCoin *SolarTokenCoinFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *SolarTokenCoinRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _SolarTokenCoin.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SolarTokenCoinRoleGranted)
				if err := _SolarTokenCoin.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_SolarTokenCoin *SolarTokenCoinFilterer) ParseRoleGranted(log types.Log) (*SolarTokenCoinRoleGranted, error) {
	event := new(SolarTokenCoinRoleGranted)
	if err := _SolarTokenCoin.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SolarTokenCoinRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the SolarTokenCoin contract.
type SolarTokenCoinRoleRevokedIterator struct {
	Event *SolarTokenCoinRoleRevoked // Event containing the contract specifics and raw log

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
func (it *SolarTokenCoinRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SolarTokenCoinRoleRevoked)
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
		it.Event = new(SolarTokenCoinRoleRevoked)
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
func (it *SolarTokenCoinRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SolarTokenCoinRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SolarTokenCoinRoleRevoked represents a RoleRevoked event raised by the SolarTokenCoin contract.
type SolarTokenCoinRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_SolarTokenCoin *SolarTokenCoinFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*SolarTokenCoinRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _SolarTokenCoin.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &SolarTokenCoinRoleRevokedIterator{contract: _SolarTokenCoin.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_SolarTokenCoin *SolarTokenCoinFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *SolarTokenCoinRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _SolarTokenCoin.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SolarTokenCoinRoleRevoked)
				if err := _SolarTokenCoin.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_SolarTokenCoin *SolarTokenCoinFilterer) ParseRoleRevoked(log types.Log) (*SolarTokenCoinRoleRevoked, error) {
	event := new(SolarTokenCoinRoleRevoked)
	if err := _SolarTokenCoin.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SolarTokenCoinTradeIterator is returned from FilterTrade and is used to iterate over the raw logs and unpacked data for Trade events raised by the SolarTokenCoin contract.
type SolarTokenCoinTradeIterator struct {
	Event *SolarTokenCoinTrade // Event containing the contract specifics and raw log

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
func (it *SolarTokenCoinTradeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SolarTokenCoinTrade)
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
		it.Event = new(SolarTokenCoinTrade)
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
func (it *SolarTokenCoinTradeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SolarTokenCoinTradeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SolarTokenCoinTrade represents a Trade event raised by the SolarTokenCoin contract.
type SolarTokenCoinTrade struct {
	Buyer           common.Address
	Seller          common.Address
	TokenId         *big.Int
	TradeId         *big.Int
	OrderId         *big.Int
	Price           *big.Int
	FeePercent      uint16
	SellerIncome    *big.Int
	OrganiserIncome *big.Int
	Amount          *big.Int
	Topic           string
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterTrade is a free log retrieval operation binding the contract event 0x51673e66d6b42dde09f3f4eaf694071ca1c153777bb4a7837a2d2af828ea55ab.
//
// Solidity: event Trade(address indexed buyer, address indexed seller, uint256 indexed tokenId, uint256 tradeId, uint256 orderId, uint256 price, uint16 feePercent, uint256 sellerIncome, uint256 organiserIncome, uint256 amount, string topic)
func (_SolarTokenCoin *SolarTokenCoinFilterer) FilterTrade(opts *bind.FilterOpts, buyer []common.Address, seller []common.Address, tokenId []*big.Int) (*SolarTokenCoinTradeIterator, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}
	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _SolarTokenCoin.contract.FilterLogs(opts, "Trade", buyerRule, sellerRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &SolarTokenCoinTradeIterator{contract: _SolarTokenCoin.contract, event: "Trade", logs: logs, sub: sub}, nil
}

// WatchTrade is a free log subscription operation binding the contract event 0x51673e66d6b42dde09f3f4eaf694071ca1c153777bb4a7837a2d2af828ea55ab.
//
// Solidity: event Trade(address indexed buyer, address indexed seller, uint256 indexed tokenId, uint256 tradeId, uint256 orderId, uint256 price, uint16 feePercent, uint256 sellerIncome, uint256 organiserIncome, uint256 amount, string topic)
func (_SolarTokenCoin *SolarTokenCoinFilterer) WatchTrade(opts *bind.WatchOpts, sink chan<- *SolarTokenCoinTrade, buyer []common.Address, seller []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}
	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _SolarTokenCoin.contract.WatchLogs(opts, "Trade", buyerRule, sellerRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SolarTokenCoinTrade)
				if err := _SolarTokenCoin.contract.UnpackLog(event, "Trade", log); err != nil {
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

// ParseTrade is a log parse operation binding the contract event 0x51673e66d6b42dde09f3f4eaf694071ca1c153777bb4a7837a2d2af828ea55ab.
//
// Solidity: event Trade(address indexed buyer, address indexed seller, uint256 indexed tokenId, uint256 tradeId, uint256 orderId, uint256 price, uint16 feePercent, uint256 sellerIncome, uint256 organiserIncome, uint256 amount, string topic)
func (_SolarTokenCoin *SolarTokenCoinFilterer) ParseTrade(log types.Log) (*SolarTokenCoinTrade, error) {
	event := new(SolarTokenCoinTrade)
	if err := _SolarTokenCoin.contract.UnpackLog(event, "Trade", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SolarTokenCoinTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the SolarTokenCoin contract.
type SolarTokenCoinTransferIterator struct {
	Event *SolarTokenCoinTransfer // Event containing the contract specifics and raw log

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
func (it *SolarTokenCoinTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SolarTokenCoinTransfer)
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
		it.Event = new(SolarTokenCoinTransfer)
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
func (it *SolarTokenCoinTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SolarTokenCoinTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SolarTokenCoinTransfer represents a Transfer event raised by the SolarTokenCoin contract.
type SolarTokenCoinTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_SolarTokenCoin *SolarTokenCoinFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*SolarTokenCoinTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _SolarTokenCoin.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &SolarTokenCoinTransferIterator{contract: _SolarTokenCoin.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_SolarTokenCoin *SolarTokenCoinFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *SolarTokenCoinTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _SolarTokenCoin.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SolarTokenCoinTransfer)
				if err := _SolarTokenCoin.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_SolarTokenCoin *SolarTokenCoinFilterer) ParseTransfer(log types.Log) (*SolarTokenCoinTransfer, error) {
	event := new(SolarTokenCoinTransfer)
	if err := _SolarTokenCoin.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SolarTokenCoinTransferDirectSolarIterator is returned from FilterTransferDirectSolar and is used to iterate over the raw logs and unpacked data for TransferDirectSolar events raised by the SolarTokenCoin contract.
type SolarTokenCoinTransferDirectSolarIterator struct {
	Event *SolarTokenCoinTransferDirectSolar // Event containing the contract specifics and raw log

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
func (it *SolarTokenCoinTransferDirectSolarIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SolarTokenCoinTransferDirectSolar)
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
		it.Event = new(SolarTokenCoinTransferDirectSolar)
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
func (it *SolarTokenCoinTransferDirectSolarIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SolarTokenCoinTransferDirectSolarIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SolarTokenCoinTransferDirectSolar represents a TransferDirectSolar event raised by the SolarTokenCoin contract.
type SolarTokenCoinTransferDirectSolar struct {
	Owner     common.Address
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTransferDirectSolar is a free log retrieval operation binding the contract event 0x2a26d8faa722f3f9b6c2e53a3875378dd062f9d54bfda18357d24bcd431c0375.
//
// Solidity: event TransferDirectSolar(address indexed _owner, address indexed _recipient, uint256 indexed _amount)
func (_SolarTokenCoin *SolarTokenCoinFilterer) FilterTransferDirectSolar(opts *bind.FilterOpts, _owner []common.Address, _recipient []common.Address, _amount []*big.Int) (*SolarTokenCoinTransferDirectSolarIterator, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _recipientRule []interface{}
	for _, _recipientItem := range _recipient {
		_recipientRule = append(_recipientRule, _recipientItem)
	}
	var _amountRule []interface{}
	for _, _amountItem := range _amount {
		_amountRule = append(_amountRule, _amountItem)
	}

	logs, sub, err := _SolarTokenCoin.contract.FilterLogs(opts, "TransferDirectSolar", _ownerRule, _recipientRule, _amountRule)
	if err != nil {
		return nil, err
	}
	return &SolarTokenCoinTransferDirectSolarIterator{contract: _SolarTokenCoin.contract, event: "TransferDirectSolar", logs: logs, sub: sub}, nil
}

// WatchTransferDirectSolar is a free log subscription operation binding the contract event 0x2a26d8faa722f3f9b6c2e53a3875378dd062f9d54bfda18357d24bcd431c0375.
//
// Solidity: event TransferDirectSolar(address indexed _owner, address indexed _recipient, uint256 indexed _amount)
func (_SolarTokenCoin *SolarTokenCoinFilterer) WatchTransferDirectSolar(opts *bind.WatchOpts, sink chan<- *SolarTokenCoinTransferDirectSolar, _owner []common.Address, _recipient []common.Address, _amount []*big.Int) (event.Subscription, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _recipientRule []interface{}
	for _, _recipientItem := range _recipient {
		_recipientRule = append(_recipientRule, _recipientItem)
	}
	var _amountRule []interface{}
	for _, _amountItem := range _amount {
		_amountRule = append(_amountRule, _amountItem)
	}

	logs, sub, err := _SolarTokenCoin.contract.WatchLogs(opts, "TransferDirectSolar", _ownerRule, _recipientRule, _amountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SolarTokenCoinTransferDirectSolar)
				if err := _SolarTokenCoin.contract.UnpackLog(event, "TransferDirectSolar", log); err != nil {
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

// ParseTransferDirectSolar is a log parse operation binding the contract event 0x2a26d8faa722f3f9b6c2e53a3875378dd062f9d54bfda18357d24bcd431c0375.
//
// Solidity: event TransferDirectSolar(address indexed _owner, address indexed _recipient, uint256 indexed _amount)
func (_SolarTokenCoin *SolarTokenCoinFilterer) ParseTransferDirectSolar(log types.Log) (*SolarTokenCoinTransferDirectSolar, error) {
	event := new(SolarTokenCoinTransferDirectSolar)
	if err := _SolarTokenCoin.contract.UnpackLog(event, "TransferDirectSolar", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SolarTokenCoinTransferTokensIterator is returned from FilterTransferTokens and is used to iterate over the raw logs and unpacked data for TransferTokens events raised by the SolarTokenCoin contract.
type SolarTokenCoinTransferTokensIterator struct {
	Event *SolarTokenCoinTransferTokens // Event containing the contract specifics and raw log

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
func (it *SolarTokenCoinTransferTokensIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SolarTokenCoinTransferTokens)
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
		it.Event = new(SolarTokenCoinTransferTokens)
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
func (it *SolarTokenCoinTransferTokensIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SolarTokenCoinTransferTokensIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SolarTokenCoinTransferTokens represents a TransferTokens event raised by the SolarTokenCoin contract.
type SolarTokenCoinTransferTokens struct {
	From      common.Address
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTransferTokens is a free log retrieval operation binding the contract event 0x58908a0fd75f7db2ca358a37b3076327d374ee1403d013a2efbc255535501edf.
//
// Solidity: event TransferTokens(address indexed _from, address indexed _recipient, uint256 indexed _amount)
func (_SolarTokenCoin *SolarTokenCoinFilterer) FilterTransferTokens(opts *bind.FilterOpts, _from []common.Address, _recipient []common.Address, _amount []*big.Int) (*SolarTokenCoinTransferTokensIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _recipientRule []interface{}
	for _, _recipientItem := range _recipient {
		_recipientRule = append(_recipientRule, _recipientItem)
	}
	var _amountRule []interface{}
	for _, _amountItem := range _amount {
		_amountRule = append(_amountRule, _amountItem)
	}

	logs, sub, err := _SolarTokenCoin.contract.FilterLogs(opts, "TransferTokens", _fromRule, _recipientRule, _amountRule)
	if err != nil {
		return nil, err
	}
	return &SolarTokenCoinTransferTokensIterator{contract: _SolarTokenCoin.contract, event: "TransferTokens", logs: logs, sub: sub}, nil
}

// WatchTransferTokens is a free log subscription operation binding the contract event 0x58908a0fd75f7db2ca358a37b3076327d374ee1403d013a2efbc255535501edf.
//
// Solidity: event TransferTokens(address indexed _from, address indexed _recipient, uint256 indexed _amount)
func (_SolarTokenCoin *SolarTokenCoinFilterer) WatchTransferTokens(opts *bind.WatchOpts, sink chan<- *SolarTokenCoinTransferTokens, _from []common.Address, _recipient []common.Address, _amount []*big.Int) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _recipientRule []interface{}
	for _, _recipientItem := range _recipient {
		_recipientRule = append(_recipientRule, _recipientItem)
	}
	var _amountRule []interface{}
	for _, _amountItem := range _amount {
		_amountRule = append(_amountRule, _amountItem)
	}

	logs, sub, err := _SolarTokenCoin.contract.WatchLogs(opts, "TransferTokens", _fromRule, _recipientRule, _amountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SolarTokenCoinTransferTokens)
				if err := _SolarTokenCoin.contract.UnpackLog(event, "TransferTokens", log); err != nil {
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

// ParseTransferTokens is a log parse operation binding the contract event 0x58908a0fd75f7db2ca358a37b3076327d374ee1403d013a2efbc255535501edf.
//
// Solidity: event TransferTokens(address indexed _from, address indexed _recipient, uint256 indexed _amount)
func (_SolarTokenCoin *SolarTokenCoinFilterer) ParseTransferTokens(log types.Log) (*SolarTokenCoinTransferTokens, error) {
	event := new(SolarTokenCoinTransferTokens)
	if err := _SolarTokenCoin.contract.UnpackLog(event, "TransferTokens", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
