// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package WalletProxy

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

// WalletProxyMetaData contains all meta data concerning the WalletProxy contract.
var WalletProxyMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"}],\"name\":\"BatchReceived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"contratAdress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"enumAuthorityUpgradeable.AuthorityRole\",\"name\":\"role\",\"type\":\"uint8\"}],\"name\":\"GrantRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"}],\"name\":\"Received\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"contratAdress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"enumAuthorityUpgradeable.AuthorityRole\",\"name\":\"role\",\"type\":\"uint8\"}],\"name\":\"RevokeRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantAdminRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantGiftRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantManagerRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantMintRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasAdminRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasGiftRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasManagerRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasMintRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeAdminRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeGiftRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeManagerRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeMintRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC1155ExtUpgradeable\",\"name\":\"erc1155\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"tokenURI\",\"type\":\"string\"}],\"name\":\"mintSingleItem\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"string[]\",\"name\":\"tokenUris\",\"type\":\"string[]\"}],\"name\":\"mintBatchItems\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"gift\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"giftBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"onERC1155Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"onERC1155BatchReceived\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// WalletProxyABI is the input ABI used to generate the binding from.
// Deprecated: Use WalletProxyMetaData.ABI instead.
var WalletProxyABI = WalletProxyMetaData.ABI

// WalletProxy is an auto generated Go binding around an Ethereum contract.
type WalletProxy struct {
	WalletProxyCaller     // Read-only binding to the contract
	WalletProxyTransactor // Write-only binding to the contract
	WalletProxyFilterer   // Log filterer for contract events
}

// WalletProxyCaller is an auto generated read-only Go binding around an Ethereum contract.
type WalletProxyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WalletProxyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WalletProxyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WalletProxyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WalletProxyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WalletProxySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WalletProxySession struct {
	Contract     *WalletProxy      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WalletProxyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WalletProxyCallerSession struct {
	Contract *WalletProxyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// WalletProxyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WalletProxyTransactorSession struct {
	Contract     *WalletProxyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// WalletProxyRaw is an auto generated low-level Go binding around an Ethereum contract.
type WalletProxyRaw struct {
	Contract *WalletProxy // Generic contract binding to access the raw methods on
}

// WalletProxyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WalletProxyCallerRaw struct {
	Contract *WalletProxyCaller // Generic read-only contract binding to access the raw methods on
}

// WalletProxyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WalletProxyTransactorRaw struct {
	Contract *WalletProxyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWalletProxy creates a new instance of WalletProxy, bound to a specific deployed contract.
func NewWalletProxy(address common.Address, backend bind.ContractBackend) (*WalletProxy, error) {
	contract, err := bindWalletProxy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &WalletProxy{WalletProxyCaller: WalletProxyCaller{contract: contract}, WalletProxyTransactor: WalletProxyTransactor{contract: contract}, WalletProxyFilterer: WalletProxyFilterer{contract: contract}}, nil
}

// NewWalletProxyCaller creates a new read-only instance of WalletProxy, bound to a specific deployed contract.
func NewWalletProxyCaller(address common.Address, caller bind.ContractCaller) (*WalletProxyCaller, error) {
	contract, err := bindWalletProxy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WalletProxyCaller{contract: contract}, nil
}

// NewWalletProxyTransactor creates a new write-only instance of WalletProxy, bound to a specific deployed contract.
func NewWalletProxyTransactor(address common.Address, transactor bind.ContractTransactor) (*WalletProxyTransactor, error) {
	contract, err := bindWalletProxy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WalletProxyTransactor{contract: contract}, nil
}

// NewWalletProxyFilterer creates a new log filterer instance of WalletProxy, bound to a specific deployed contract.
func NewWalletProxyFilterer(address common.Address, filterer bind.ContractFilterer) (*WalletProxyFilterer, error) {
	contract, err := bindWalletProxy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WalletProxyFilterer{contract: contract}, nil
}

// bindWalletProxy binds a generic wrapper to an already deployed contract.
func bindWalletProxy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(WalletProxyABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WalletProxy *WalletProxyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WalletProxy.Contract.WalletProxyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WalletProxy *WalletProxyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WalletProxy.Contract.WalletProxyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WalletProxy *WalletProxyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WalletProxy.Contract.WalletProxyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WalletProxy *WalletProxyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WalletProxy.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WalletProxy *WalletProxyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WalletProxy.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WalletProxy *WalletProxyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WalletProxy.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_WalletProxy *WalletProxyCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _WalletProxy.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_WalletProxy *WalletProxySession) DEFAULTADMINROLE() ([32]byte, error) {
	return _WalletProxy.Contract.DEFAULTADMINROLE(&_WalletProxy.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_WalletProxy *WalletProxyCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _WalletProxy.Contract.DEFAULTADMINROLE(&_WalletProxy.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_WalletProxy *WalletProxyCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _WalletProxy.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_WalletProxy *WalletProxySession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _WalletProxy.Contract.GetRoleAdmin(&_WalletProxy.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_WalletProxy *WalletProxyCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _WalletProxy.Contract.GetRoleAdmin(&_WalletProxy.CallOpts, role)
}

// HasAdminRole is a free data retrieval call binding the contract method 0xc395fcb3.
//
// Solidity: function hasAdminRole(address account) view returns(bool)
func (_WalletProxy *WalletProxyCaller) HasAdminRole(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _WalletProxy.contract.Call(opts, &out, "hasAdminRole", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasAdminRole is a free data retrieval call binding the contract method 0xc395fcb3.
//
// Solidity: function hasAdminRole(address account) view returns(bool)
func (_WalletProxy *WalletProxySession) HasAdminRole(account common.Address) (bool, error) {
	return _WalletProxy.Contract.HasAdminRole(&_WalletProxy.CallOpts, account)
}

// HasAdminRole is a free data retrieval call binding the contract method 0xc395fcb3.
//
// Solidity: function hasAdminRole(address account) view returns(bool)
func (_WalletProxy *WalletProxyCallerSession) HasAdminRole(account common.Address) (bool, error) {
	return _WalletProxy.Contract.HasAdminRole(&_WalletProxy.CallOpts, account)
}

// HasGiftRole is a free data retrieval call binding the contract method 0x629962dc.
//
// Solidity: function hasGiftRole(address account) view returns(bool)
func (_WalletProxy *WalletProxyCaller) HasGiftRole(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _WalletProxy.contract.Call(opts, &out, "hasGiftRole", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasGiftRole is a free data retrieval call binding the contract method 0x629962dc.
//
// Solidity: function hasGiftRole(address account) view returns(bool)
func (_WalletProxy *WalletProxySession) HasGiftRole(account common.Address) (bool, error) {
	return _WalletProxy.Contract.HasGiftRole(&_WalletProxy.CallOpts, account)
}

// HasGiftRole is a free data retrieval call binding the contract method 0x629962dc.
//
// Solidity: function hasGiftRole(address account) view returns(bool)
func (_WalletProxy *WalletProxyCallerSession) HasGiftRole(account common.Address) (bool, error) {
	return _WalletProxy.Contract.HasGiftRole(&_WalletProxy.CallOpts, account)
}

// HasManagerRole is a free data retrieval call binding the contract method 0x5026c826.
//
// Solidity: function hasManagerRole(address account) view returns(bool)
func (_WalletProxy *WalletProxyCaller) HasManagerRole(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _WalletProxy.contract.Call(opts, &out, "hasManagerRole", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasManagerRole is a free data retrieval call binding the contract method 0x5026c826.
//
// Solidity: function hasManagerRole(address account) view returns(bool)
func (_WalletProxy *WalletProxySession) HasManagerRole(account common.Address) (bool, error) {
	return _WalletProxy.Contract.HasManagerRole(&_WalletProxy.CallOpts, account)
}

// HasManagerRole is a free data retrieval call binding the contract method 0x5026c826.
//
// Solidity: function hasManagerRole(address account) view returns(bool)
func (_WalletProxy *WalletProxyCallerSession) HasManagerRole(account common.Address) (bool, error) {
	return _WalletProxy.Contract.HasManagerRole(&_WalletProxy.CallOpts, account)
}

// HasMintRole is a free data retrieval call binding the contract method 0xc4e58e54.
//
// Solidity: function hasMintRole(address account) view returns(bool)
func (_WalletProxy *WalletProxyCaller) HasMintRole(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _WalletProxy.contract.Call(opts, &out, "hasMintRole", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasMintRole is a free data retrieval call binding the contract method 0xc4e58e54.
//
// Solidity: function hasMintRole(address account) view returns(bool)
func (_WalletProxy *WalletProxySession) HasMintRole(account common.Address) (bool, error) {
	return _WalletProxy.Contract.HasMintRole(&_WalletProxy.CallOpts, account)
}

// HasMintRole is a free data retrieval call binding the contract method 0xc4e58e54.
//
// Solidity: function hasMintRole(address account) view returns(bool)
func (_WalletProxy *WalletProxyCallerSession) HasMintRole(account common.Address) (bool, error) {
	return _WalletProxy.Contract.HasMintRole(&_WalletProxy.CallOpts, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_WalletProxy *WalletProxyCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _WalletProxy.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_WalletProxy *WalletProxySession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _WalletProxy.Contract.HasRole(&_WalletProxy.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_WalletProxy *WalletProxyCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _WalletProxy.Contract.HasRole(&_WalletProxy.CallOpts, role, account)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_WalletProxy *WalletProxyCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WalletProxy.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_WalletProxy *WalletProxySession) Owner() (common.Address, error) {
	return _WalletProxy.Contract.Owner(&_WalletProxy.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_WalletProxy *WalletProxyCallerSession) Owner() (common.Address, error) {
	return _WalletProxy.Contract.Owner(&_WalletProxy.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_WalletProxy *WalletProxyCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _WalletProxy.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_WalletProxy *WalletProxySession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _WalletProxy.Contract.SupportsInterface(&_WalletProxy.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_WalletProxy *WalletProxyCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _WalletProxy.Contract.SupportsInterface(&_WalletProxy.CallOpts, interfaceId)
}

// Gift is a paid mutator transaction binding the contract method 0x598d00a9.
//
// Solidity: function gift(address to, uint256 tokenId, uint256 amount) returns()
func (_WalletProxy *WalletProxyTransactor) Gift(opts *bind.TransactOpts, to common.Address, tokenId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _WalletProxy.contract.Transact(opts, "gift", to, tokenId, amount)
}

// Gift is a paid mutator transaction binding the contract method 0x598d00a9.
//
// Solidity: function gift(address to, uint256 tokenId, uint256 amount) returns()
func (_WalletProxy *WalletProxySession) Gift(to common.Address, tokenId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _WalletProxy.Contract.Gift(&_WalletProxy.TransactOpts, to, tokenId, amount)
}

// Gift is a paid mutator transaction binding the contract method 0x598d00a9.
//
// Solidity: function gift(address to, uint256 tokenId, uint256 amount) returns()
func (_WalletProxy *WalletProxyTransactorSession) Gift(to common.Address, tokenId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _WalletProxy.Contract.Gift(&_WalletProxy.TransactOpts, to, tokenId, amount)
}

// GiftBatch is a paid mutator transaction binding the contract method 0xd741fb24.
//
// Solidity: function giftBatch(address to, uint256[] tokenIds, uint256[] amounts) returns()
func (_WalletProxy *WalletProxyTransactor) GiftBatch(opts *bind.TransactOpts, to common.Address, tokenIds []*big.Int, amounts []*big.Int) (*types.Transaction, error) {
	return _WalletProxy.contract.Transact(opts, "giftBatch", to, tokenIds, amounts)
}

// GiftBatch is a paid mutator transaction binding the contract method 0xd741fb24.
//
// Solidity: function giftBatch(address to, uint256[] tokenIds, uint256[] amounts) returns()
func (_WalletProxy *WalletProxySession) GiftBatch(to common.Address, tokenIds []*big.Int, amounts []*big.Int) (*types.Transaction, error) {
	return _WalletProxy.Contract.GiftBatch(&_WalletProxy.TransactOpts, to, tokenIds, amounts)
}

// GiftBatch is a paid mutator transaction binding the contract method 0xd741fb24.
//
// Solidity: function giftBatch(address to, uint256[] tokenIds, uint256[] amounts) returns()
func (_WalletProxy *WalletProxyTransactorSession) GiftBatch(to common.Address, tokenIds []*big.Int, amounts []*big.Int) (*types.Transaction, error) {
	return _WalletProxy.Contract.GiftBatch(&_WalletProxy.TransactOpts, to, tokenIds, amounts)
}

// GrantAdminRole is a paid mutator transaction binding the contract method 0xc634b78e.
//
// Solidity: function grantAdminRole(address account) returns()
func (_WalletProxy *WalletProxyTransactor) GrantAdminRole(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _WalletProxy.contract.Transact(opts, "grantAdminRole", account)
}

// GrantAdminRole is a paid mutator transaction binding the contract method 0xc634b78e.
//
// Solidity: function grantAdminRole(address account) returns()
func (_WalletProxy *WalletProxySession) GrantAdminRole(account common.Address) (*types.Transaction, error) {
	return _WalletProxy.Contract.GrantAdminRole(&_WalletProxy.TransactOpts, account)
}

// GrantAdminRole is a paid mutator transaction binding the contract method 0xc634b78e.
//
// Solidity: function grantAdminRole(address account) returns()
func (_WalletProxy *WalletProxyTransactorSession) GrantAdminRole(account common.Address) (*types.Transaction, error) {
	return _WalletProxy.Contract.GrantAdminRole(&_WalletProxy.TransactOpts, account)
}

// GrantGiftRole is a paid mutator transaction binding the contract method 0x0bd9c905.
//
// Solidity: function grantGiftRole(address account) returns()
func (_WalletProxy *WalletProxyTransactor) GrantGiftRole(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _WalletProxy.contract.Transact(opts, "grantGiftRole", account)
}

// GrantGiftRole is a paid mutator transaction binding the contract method 0x0bd9c905.
//
// Solidity: function grantGiftRole(address account) returns()
func (_WalletProxy *WalletProxySession) GrantGiftRole(account common.Address) (*types.Transaction, error) {
	return _WalletProxy.Contract.GrantGiftRole(&_WalletProxy.TransactOpts, account)
}

// GrantGiftRole is a paid mutator transaction binding the contract method 0x0bd9c905.
//
// Solidity: function grantGiftRole(address account) returns()
func (_WalletProxy *WalletProxyTransactorSession) GrantGiftRole(account common.Address) (*types.Transaction, error) {
	return _WalletProxy.Contract.GrantGiftRole(&_WalletProxy.TransactOpts, account)
}

// GrantManagerRole is a paid mutator transaction binding the contract method 0x26e885e3.
//
// Solidity: function grantManagerRole(address account) returns()
func (_WalletProxy *WalletProxyTransactor) GrantManagerRole(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _WalletProxy.contract.Transact(opts, "grantManagerRole", account)
}

// GrantManagerRole is a paid mutator transaction binding the contract method 0x26e885e3.
//
// Solidity: function grantManagerRole(address account) returns()
func (_WalletProxy *WalletProxySession) GrantManagerRole(account common.Address) (*types.Transaction, error) {
	return _WalletProxy.Contract.GrantManagerRole(&_WalletProxy.TransactOpts, account)
}

// GrantManagerRole is a paid mutator transaction binding the contract method 0x26e885e3.
//
// Solidity: function grantManagerRole(address account) returns()
func (_WalletProxy *WalletProxyTransactorSession) GrantManagerRole(account common.Address) (*types.Transaction, error) {
	return _WalletProxy.Contract.GrantManagerRole(&_WalletProxy.TransactOpts, account)
}

// GrantMintRole is a paid mutator transaction binding the contract method 0xc2e3273d.
//
// Solidity: function grantMintRole(address account) returns()
func (_WalletProxy *WalletProxyTransactor) GrantMintRole(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _WalletProxy.contract.Transact(opts, "grantMintRole", account)
}

// GrantMintRole is a paid mutator transaction binding the contract method 0xc2e3273d.
//
// Solidity: function grantMintRole(address account) returns()
func (_WalletProxy *WalletProxySession) GrantMintRole(account common.Address) (*types.Transaction, error) {
	return _WalletProxy.Contract.GrantMintRole(&_WalletProxy.TransactOpts, account)
}

// GrantMintRole is a paid mutator transaction binding the contract method 0xc2e3273d.
//
// Solidity: function grantMintRole(address account) returns()
func (_WalletProxy *WalletProxyTransactorSession) GrantMintRole(account common.Address) (*types.Transaction, error) {
	return _WalletProxy.Contract.GrantMintRole(&_WalletProxy.TransactOpts, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_WalletProxy *WalletProxyTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _WalletProxy.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_WalletProxy *WalletProxySession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _WalletProxy.Contract.GrantRole(&_WalletProxy.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_WalletProxy *WalletProxyTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _WalletProxy.Contract.GrantRole(&_WalletProxy.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0x90657147.
//
// Solidity: function initialize(address erc1155, string name, string symbol) returns()
func (_WalletProxy *WalletProxyTransactor) Initialize(opts *bind.TransactOpts, erc1155 common.Address, name string, symbol string) (*types.Transaction, error) {
	return _WalletProxy.contract.Transact(opts, "initialize", erc1155, name, symbol)
}

// Initialize is a paid mutator transaction binding the contract method 0x90657147.
//
// Solidity: function initialize(address erc1155, string name, string symbol) returns()
func (_WalletProxy *WalletProxySession) Initialize(erc1155 common.Address, name string, symbol string) (*types.Transaction, error) {
	return _WalletProxy.Contract.Initialize(&_WalletProxy.TransactOpts, erc1155, name, symbol)
}

// Initialize is a paid mutator transaction binding the contract method 0x90657147.
//
// Solidity: function initialize(address erc1155, string name, string symbol) returns()
func (_WalletProxy *WalletProxyTransactorSession) Initialize(erc1155 common.Address, name string, symbol string) (*types.Transaction, error) {
	return _WalletProxy.Contract.Initialize(&_WalletProxy.TransactOpts, erc1155, name, symbol)
}

// MintBatchItems is a paid mutator transaction binding the contract method 0xd353676a.
//
// Solidity: function mintBatchItems(uint256[] tokenIds, uint256[] amounts, string[] tokenUris) returns()
func (_WalletProxy *WalletProxyTransactor) MintBatchItems(opts *bind.TransactOpts, tokenIds []*big.Int, amounts []*big.Int, tokenUris []string) (*types.Transaction, error) {
	return _WalletProxy.contract.Transact(opts, "mintBatchItems", tokenIds, amounts, tokenUris)
}

// MintBatchItems is a paid mutator transaction binding the contract method 0xd353676a.
//
// Solidity: function mintBatchItems(uint256[] tokenIds, uint256[] amounts, string[] tokenUris) returns()
func (_WalletProxy *WalletProxySession) MintBatchItems(tokenIds []*big.Int, amounts []*big.Int, tokenUris []string) (*types.Transaction, error) {
	return _WalletProxy.Contract.MintBatchItems(&_WalletProxy.TransactOpts, tokenIds, amounts, tokenUris)
}

// MintBatchItems is a paid mutator transaction binding the contract method 0xd353676a.
//
// Solidity: function mintBatchItems(uint256[] tokenIds, uint256[] amounts, string[] tokenUris) returns()
func (_WalletProxy *WalletProxyTransactorSession) MintBatchItems(tokenIds []*big.Int, amounts []*big.Int, tokenUris []string) (*types.Transaction, error) {
	return _WalletProxy.Contract.MintBatchItems(&_WalletProxy.TransactOpts, tokenIds, amounts, tokenUris)
}

// MintSingleItem is a paid mutator transaction binding the contract method 0xea2e0c1d.
//
// Solidity: function mintSingleItem(uint256 tokenId, uint256 amount, string tokenURI) returns()
func (_WalletProxy *WalletProxyTransactor) MintSingleItem(opts *bind.TransactOpts, tokenId *big.Int, amount *big.Int, tokenURI string) (*types.Transaction, error) {
	return _WalletProxy.contract.Transact(opts, "mintSingleItem", tokenId, amount, tokenURI)
}

// MintSingleItem is a paid mutator transaction binding the contract method 0xea2e0c1d.
//
// Solidity: function mintSingleItem(uint256 tokenId, uint256 amount, string tokenURI) returns()
func (_WalletProxy *WalletProxySession) MintSingleItem(tokenId *big.Int, amount *big.Int, tokenURI string) (*types.Transaction, error) {
	return _WalletProxy.Contract.MintSingleItem(&_WalletProxy.TransactOpts, tokenId, amount, tokenURI)
}

// MintSingleItem is a paid mutator transaction binding the contract method 0xea2e0c1d.
//
// Solidity: function mintSingleItem(uint256 tokenId, uint256 amount, string tokenURI) returns()
func (_WalletProxy *WalletProxyTransactorSession) MintSingleItem(tokenId *big.Int, amount *big.Int, tokenURI string) (*types.Transaction, error) {
	return _WalletProxy.Contract.MintSingleItem(&_WalletProxy.TransactOpts, tokenId, amount, tokenURI)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address operator, address from, uint256[] tokenIds, uint256[] values, bytes data) returns(bytes4)
func (_WalletProxy *WalletProxyTransactor) OnERC1155BatchReceived(opts *bind.TransactOpts, operator common.Address, from common.Address, tokenIds []*big.Int, values []*big.Int, data []byte) (*types.Transaction, error) {
	return _WalletProxy.contract.Transact(opts, "onERC1155BatchReceived", operator, from, tokenIds, values, data)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address operator, address from, uint256[] tokenIds, uint256[] values, bytes data) returns(bytes4)
func (_WalletProxy *WalletProxySession) OnERC1155BatchReceived(operator common.Address, from common.Address, tokenIds []*big.Int, values []*big.Int, data []byte) (*types.Transaction, error) {
	return _WalletProxy.Contract.OnERC1155BatchReceived(&_WalletProxy.TransactOpts, operator, from, tokenIds, values, data)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address operator, address from, uint256[] tokenIds, uint256[] values, bytes data) returns(bytes4)
func (_WalletProxy *WalletProxyTransactorSession) OnERC1155BatchReceived(operator common.Address, from common.Address, tokenIds []*big.Int, values []*big.Int, data []byte) (*types.Transaction, error) {
	return _WalletProxy.Contract.OnERC1155BatchReceived(&_WalletProxy.TransactOpts, operator, from, tokenIds, values, data)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address operator, address from, uint256 tokenId, uint256 value, bytes data) returns(bytes4)
func (_WalletProxy *WalletProxyTransactor) OnERC1155Received(opts *bind.TransactOpts, operator common.Address, from common.Address, tokenId *big.Int, value *big.Int, data []byte) (*types.Transaction, error) {
	return _WalletProxy.contract.Transact(opts, "onERC1155Received", operator, from, tokenId, value, data)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address operator, address from, uint256 tokenId, uint256 value, bytes data) returns(bytes4)
func (_WalletProxy *WalletProxySession) OnERC1155Received(operator common.Address, from common.Address, tokenId *big.Int, value *big.Int, data []byte) (*types.Transaction, error) {
	return _WalletProxy.Contract.OnERC1155Received(&_WalletProxy.TransactOpts, operator, from, tokenId, value, data)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address operator, address from, uint256 tokenId, uint256 value, bytes data) returns(bytes4)
func (_WalletProxy *WalletProxyTransactorSession) OnERC1155Received(operator common.Address, from common.Address, tokenId *big.Int, value *big.Int, data []byte) (*types.Transaction, error) {
	return _WalletProxy.Contract.OnERC1155Received(&_WalletProxy.TransactOpts, operator, from, tokenId, value, data)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_WalletProxy *WalletProxyTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WalletProxy.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_WalletProxy *WalletProxySession) RenounceOwnership() (*types.Transaction, error) {
	return _WalletProxy.Contract.RenounceOwnership(&_WalletProxy.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_WalletProxy *WalletProxyTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _WalletProxy.Contract.RenounceOwnership(&_WalletProxy.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_WalletProxy *WalletProxyTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _WalletProxy.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_WalletProxy *WalletProxySession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _WalletProxy.Contract.RenounceRole(&_WalletProxy.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_WalletProxy *WalletProxyTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _WalletProxy.Contract.RenounceRole(&_WalletProxy.TransactOpts, role, account)
}

// RevokeAdminRole is a paid mutator transaction binding the contract method 0x9a19c7b0.
//
// Solidity: function revokeAdminRole(address account) returns()
func (_WalletProxy *WalletProxyTransactor) RevokeAdminRole(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _WalletProxy.contract.Transact(opts, "revokeAdminRole", account)
}

// RevokeAdminRole is a paid mutator transaction binding the contract method 0x9a19c7b0.
//
// Solidity: function revokeAdminRole(address account) returns()
func (_WalletProxy *WalletProxySession) RevokeAdminRole(account common.Address) (*types.Transaction, error) {
	return _WalletProxy.Contract.RevokeAdminRole(&_WalletProxy.TransactOpts, account)
}

// RevokeAdminRole is a paid mutator transaction binding the contract method 0x9a19c7b0.
//
// Solidity: function revokeAdminRole(address account) returns()
func (_WalletProxy *WalletProxyTransactorSession) RevokeAdminRole(account common.Address) (*types.Transaction, error) {
	return _WalletProxy.Contract.RevokeAdminRole(&_WalletProxy.TransactOpts, account)
}

// RevokeGiftRole is a paid mutator transaction binding the contract method 0x543309a0.
//
// Solidity: function revokeGiftRole(address account) returns()
func (_WalletProxy *WalletProxyTransactor) RevokeGiftRole(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _WalletProxy.contract.Transact(opts, "revokeGiftRole", account)
}

// RevokeGiftRole is a paid mutator transaction binding the contract method 0x543309a0.
//
// Solidity: function revokeGiftRole(address account) returns()
func (_WalletProxy *WalletProxySession) RevokeGiftRole(account common.Address) (*types.Transaction, error) {
	return _WalletProxy.Contract.RevokeGiftRole(&_WalletProxy.TransactOpts, account)
}

// RevokeGiftRole is a paid mutator transaction binding the contract method 0x543309a0.
//
// Solidity: function revokeGiftRole(address account) returns()
func (_WalletProxy *WalletProxyTransactorSession) RevokeGiftRole(account common.Address) (*types.Transaction, error) {
	return _WalletProxy.Contract.RevokeGiftRole(&_WalletProxy.TransactOpts, account)
}

// RevokeManagerRole is a paid mutator transaction binding the contract method 0xbe4dc94f.
//
// Solidity: function revokeManagerRole(address account) returns()
func (_WalletProxy *WalletProxyTransactor) RevokeManagerRole(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _WalletProxy.contract.Transact(opts, "revokeManagerRole", account)
}

// RevokeManagerRole is a paid mutator transaction binding the contract method 0xbe4dc94f.
//
// Solidity: function revokeManagerRole(address account) returns()
func (_WalletProxy *WalletProxySession) RevokeManagerRole(account common.Address) (*types.Transaction, error) {
	return _WalletProxy.Contract.RevokeManagerRole(&_WalletProxy.TransactOpts, account)
}

// RevokeManagerRole is a paid mutator transaction binding the contract method 0xbe4dc94f.
//
// Solidity: function revokeManagerRole(address account) returns()
func (_WalletProxy *WalletProxyTransactorSession) RevokeManagerRole(account common.Address) (*types.Transaction, error) {
	return _WalletProxy.Contract.RevokeManagerRole(&_WalletProxy.TransactOpts, account)
}

// RevokeMintRole is a paid mutator transaction binding the contract method 0xf81094f3.
//
// Solidity: function revokeMintRole(address account) returns()
func (_WalletProxy *WalletProxyTransactor) RevokeMintRole(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _WalletProxy.contract.Transact(opts, "revokeMintRole", account)
}

// RevokeMintRole is a paid mutator transaction binding the contract method 0xf81094f3.
//
// Solidity: function revokeMintRole(address account) returns()
func (_WalletProxy *WalletProxySession) RevokeMintRole(account common.Address) (*types.Transaction, error) {
	return _WalletProxy.Contract.RevokeMintRole(&_WalletProxy.TransactOpts, account)
}

// RevokeMintRole is a paid mutator transaction binding the contract method 0xf81094f3.
//
// Solidity: function revokeMintRole(address account) returns()
func (_WalletProxy *WalletProxyTransactorSession) RevokeMintRole(account common.Address) (*types.Transaction, error) {
	return _WalletProxy.Contract.RevokeMintRole(&_WalletProxy.TransactOpts, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_WalletProxy *WalletProxyTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _WalletProxy.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_WalletProxy *WalletProxySession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _WalletProxy.Contract.RevokeRole(&_WalletProxy.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_WalletProxy *WalletProxyTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _WalletProxy.Contract.RevokeRole(&_WalletProxy.TransactOpts, role, account)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_WalletProxy *WalletProxyTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _WalletProxy.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_WalletProxy *WalletProxySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _WalletProxy.Contract.TransferOwnership(&_WalletProxy.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_WalletProxy *WalletProxyTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _WalletProxy.Contract.TransferOwnership(&_WalletProxy.TransactOpts, newOwner)
}

// WalletProxyBatchReceivedIterator is returned from FilterBatchReceived and is used to iterate over the raw logs and unpacked data for BatchReceived events raised by the WalletProxy contract.
type WalletProxyBatchReceivedIterator struct {
	Event *WalletProxyBatchReceived // Event containing the contract specifics and raw log

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
func (it *WalletProxyBatchReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletProxyBatchReceived)
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
		it.Event = new(WalletProxyBatchReceived)
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
func (it *WalletProxyBatchReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletProxyBatchReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletProxyBatchReceived represents a BatchReceived event raised by the WalletProxy contract.
type WalletProxyBatchReceived struct {
	Operator common.Address
	From     common.Address
	TokenIds []*big.Int
	Values   []*big.Int
	Data     []byte
	Gas      *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterBatchReceived is a free log retrieval operation binding the contract event 0x0bcad9224ba33b574e9c85298de2f44b4c80015a21aa5df474896444909863d8.
//
// Solidity: event BatchReceived(address operator, address from, uint256[] tokenIds, uint256[] values, bytes data, uint256 gas)
func (_WalletProxy *WalletProxyFilterer) FilterBatchReceived(opts *bind.FilterOpts) (*WalletProxyBatchReceivedIterator, error) {

	logs, sub, err := _WalletProxy.contract.FilterLogs(opts, "BatchReceived")
	if err != nil {
		return nil, err
	}
	return &WalletProxyBatchReceivedIterator{contract: _WalletProxy.contract, event: "BatchReceived", logs: logs, sub: sub}, nil
}

// WatchBatchReceived is a free log subscription operation binding the contract event 0x0bcad9224ba33b574e9c85298de2f44b4c80015a21aa5df474896444909863d8.
//
// Solidity: event BatchReceived(address operator, address from, uint256[] tokenIds, uint256[] values, bytes data, uint256 gas)
func (_WalletProxy *WalletProxyFilterer) WatchBatchReceived(opts *bind.WatchOpts, sink chan<- *WalletProxyBatchReceived) (event.Subscription, error) {

	logs, sub, err := _WalletProxy.contract.WatchLogs(opts, "BatchReceived")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletProxyBatchReceived)
				if err := _WalletProxy.contract.UnpackLog(event, "BatchReceived", log); err != nil {
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

// ParseBatchReceived is a log parse operation binding the contract event 0x0bcad9224ba33b574e9c85298de2f44b4c80015a21aa5df474896444909863d8.
//
// Solidity: event BatchReceived(address operator, address from, uint256[] tokenIds, uint256[] values, bytes data, uint256 gas)
func (_WalletProxy *WalletProxyFilterer) ParseBatchReceived(log types.Log) (*WalletProxyBatchReceived, error) {
	event := new(WalletProxyBatchReceived)
	if err := _WalletProxy.contract.UnpackLog(event, "BatchReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WalletProxyGrantRoleIterator is returned from FilterGrantRole and is used to iterate over the raw logs and unpacked data for GrantRole events raised by the WalletProxy contract.
type WalletProxyGrantRoleIterator struct {
	Event *WalletProxyGrantRole // Event containing the contract specifics and raw log

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
func (it *WalletProxyGrantRoleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletProxyGrantRole)
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
		it.Event = new(WalletProxyGrantRole)
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
func (it *WalletProxyGrantRoleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletProxyGrantRoleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletProxyGrantRole represents a GrantRole event raised by the WalletProxy contract.
type WalletProxyGrantRole struct {
	ContratAdress common.Address
	Operator      common.Address
	Account       common.Address
	Role          uint8
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterGrantRole is a free log retrieval operation binding the contract event 0x7ee499468b11c5fd78e2af0a6821cac160d8b71e0d88ed9398fe819db39ee58d.
//
// Solidity: event GrantRole(address contratAdress, address indexed operator, address indexed account, uint8 indexed role)
func (_WalletProxy *WalletProxyFilterer) FilterGrantRole(opts *bind.FilterOpts, operator []common.Address, account []common.Address, role []uint8) (*WalletProxyGrantRoleIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}

	logs, sub, err := _WalletProxy.contract.FilterLogs(opts, "GrantRole", operatorRule, accountRule, roleRule)
	if err != nil {
		return nil, err
	}
	return &WalletProxyGrantRoleIterator{contract: _WalletProxy.contract, event: "GrantRole", logs: logs, sub: sub}, nil
}

// WatchGrantRole is a free log subscription operation binding the contract event 0x7ee499468b11c5fd78e2af0a6821cac160d8b71e0d88ed9398fe819db39ee58d.
//
// Solidity: event GrantRole(address contratAdress, address indexed operator, address indexed account, uint8 indexed role)
func (_WalletProxy *WalletProxyFilterer) WatchGrantRole(opts *bind.WatchOpts, sink chan<- *WalletProxyGrantRole, operator []common.Address, account []common.Address, role []uint8) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}

	logs, sub, err := _WalletProxy.contract.WatchLogs(opts, "GrantRole", operatorRule, accountRule, roleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletProxyGrantRole)
				if err := _WalletProxy.contract.UnpackLog(event, "GrantRole", log); err != nil {
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

// ParseGrantRole is a log parse operation binding the contract event 0x7ee499468b11c5fd78e2af0a6821cac160d8b71e0d88ed9398fe819db39ee58d.
//
// Solidity: event GrantRole(address contratAdress, address indexed operator, address indexed account, uint8 indexed role)
func (_WalletProxy *WalletProxyFilterer) ParseGrantRole(log types.Log) (*WalletProxyGrantRole, error) {
	event := new(WalletProxyGrantRole)
	if err := _WalletProxy.contract.UnpackLog(event, "GrantRole", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WalletProxyOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the WalletProxy contract.
type WalletProxyOwnershipTransferredIterator struct {
	Event *WalletProxyOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *WalletProxyOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletProxyOwnershipTransferred)
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
		it.Event = new(WalletProxyOwnershipTransferred)
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
func (it *WalletProxyOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletProxyOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletProxyOwnershipTransferred represents a OwnershipTransferred event raised by the WalletProxy contract.
type WalletProxyOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_WalletProxy *WalletProxyFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*WalletProxyOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _WalletProxy.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &WalletProxyOwnershipTransferredIterator{contract: _WalletProxy.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_WalletProxy *WalletProxyFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *WalletProxyOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _WalletProxy.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletProxyOwnershipTransferred)
				if err := _WalletProxy.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_WalletProxy *WalletProxyFilterer) ParseOwnershipTransferred(log types.Log) (*WalletProxyOwnershipTransferred, error) {
	event := new(WalletProxyOwnershipTransferred)
	if err := _WalletProxy.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WalletProxyReceivedIterator is returned from FilterReceived and is used to iterate over the raw logs and unpacked data for Received events raised by the WalletProxy contract.
type WalletProxyReceivedIterator struct {
	Event *WalletProxyReceived // Event containing the contract specifics and raw log

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
func (it *WalletProxyReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletProxyReceived)
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
		it.Event = new(WalletProxyReceived)
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
func (it *WalletProxyReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletProxyReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletProxyReceived represents a Received event raised by the WalletProxy contract.
type WalletProxyReceived struct {
	Operator common.Address
	From     common.Address
	TokenId  *big.Int
	Value    *big.Int
	Data     []byte
	Gas      *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterReceived is a free log retrieval operation binding the contract event 0x20dfa9f79060c8c4d7fe892c97c71bcf6e3b63d1dcf66fea7aefc0211628cf29.
//
// Solidity: event Received(address operator, address from, uint256 tokenId, uint256 value, bytes data, uint256 gas)
func (_WalletProxy *WalletProxyFilterer) FilterReceived(opts *bind.FilterOpts) (*WalletProxyReceivedIterator, error) {

	logs, sub, err := _WalletProxy.contract.FilterLogs(opts, "Received")
	if err != nil {
		return nil, err
	}
	return &WalletProxyReceivedIterator{contract: _WalletProxy.contract, event: "Received", logs: logs, sub: sub}, nil
}

// WatchReceived is a free log subscription operation binding the contract event 0x20dfa9f79060c8c4d7fe892c97c71bcf6e3b63d1dcf66fea7aefc0211628cf29.
//
// Solidity: event Received(address operator, address from, uint256 tokenId, uint256 value, bytes data, uint256 gas)
func (_WalletProxy *WalletProxyFilterer) WatchReceived(opts *bind.WatchOpts, sink chan<- *WalletProxyReceived) (event.Subscription, error) {

	logs, sub, err := _WalletProxy.contract.WatchLogs(opts, "Received")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletProxyReceived)
				if err := _WalletProxy.contract.UnpackLog(event, "Received", log); err != nil {
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

// ParseReceived is a log parse operation binding the contract event 0x20dfa9f79060c8c4d7fe892c97c71bcf6e3b63d1dcf66fea7aefc0211628cf29.
//
// Solidity: event Received(address operator, address from, uint256 tokenId, uint256 value, bytes data, uint256 gas)
func (_WalletProxy *WalletProxyFilterer) ParseReceived(log types.Log) (*WalletProxyReceived, error) {
	event := new(WalletProxyReceived)
	if err := _WalletProxy.contract.UnpackLog(event, "Received", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WalletProxyRevokeRoleIterator is returned from FilterRevokeRole and is used to iterate over the raw logs and unpacked data for RevokeRole events raised by the WalletProxy contract.
type WalletProxyRevokeRoleIterator struct {
	Event *WalletProxyRevokeRole // Event containing the contract specifics and raw log

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
func (it *WalletProxyRevokeRoleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletProxyRevokeRole)
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
		it.Event = new(WalletProxyRevokeRole)
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
func (it *WalletProxyRevokeRoleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletProxyRevokeRoleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletProxyRevokeRole represents a RevokeRole event raised by the WalletProxy contract.
type WalletProxyRevokeRole struct {
	ContratAdress common.Address
	Operator      common.Address
	Account       common.Address
	Role          uint8
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterRevokeRole is a free log retrieval operation binding the contract event 0xa3847ba69b1901955bb65a1e552d84a42c25411b4a60509c2ebf380c3a0be301.
//
// Solidity: event RevokeRole(address contratAdress, address indexed operator, address indexed account, uint8 indexed role)
func (_WalletProxy *WalletProxyFilterer) FilterRevokeRole(opts *bind.FilterOpts, operator []common.Address, account []common.Address, role []uint8) (*WalletProxyRevokeRoleIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}

	logs, sub, err := _WalletProxy.contract.FilterLogs(opts, "RevokeRole", operatorRule, accountRule, roleRule)
	if err != nil {
		return nil, err
	}
	return &WalletProxyRevokeRoleIterator{contract: _WalletProxy.contract, event: "RevokeRole", logs: logs, sub: sub}, nil
}

// WatchRevokeRole is a free log subscription operation binding the contract event 0xa3847ba69b1901955bb65a1e552d84a42c25411b4a60509c2ebf380c3a0be301.
//
// Solidity: event RevokeRole(address contratAdress, address indexed operator, address indexed account, uint8 indexed role)
func (_WalletProxy *WalletProxyFilterer) WatchRevokeRole(opts *bind.WatchOpts, sink chan<- *WalletProxyRevokeRole, operator []common.Address, account []common.Address, role []uint8) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}

	logs, sub, err := _WalletProxy.contract.WatchLogs(opts, "RevokeRole", operatorRule, accountRule, roleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletProxyRevokeRole)
				if err := _WalletProxy.contract.UnpackLog(event, "RevokeRole", log); err != nil {
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

// ParseRevokeRole is a log parse operation binding the contract event 0xa3847ba69b1901955bb65a1e552d84a42c25411b4a60509c2ebf380c3a0be301.
//
// Solidity: event RevokeRole(address contratAdress, address indexed operator, address indexed account, uint8 indexed role)
func (_WalletProxy *WalletProxyFilterer) ParseRevokeRole(log types.Log) (*WalletProxyRevokeRole, error) {
	event := new(WalletProxyRevokeRole)
	if err := _WalletProxy.contract.UnpackLog(event, "RevokeRole", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WalletProxyRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the WalletProxy contract.
type WalletProxyRoleAdminChangedIterator struct {
	Event *WalletProxyRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *WalletProxyRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletProxyRoleAdminChanged)
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
		it.Event = new(WalletProxyRoleAdminChanged)
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
func (it *WalletProxyRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletProxyRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletProxyRoleAdminChanged represents a RoleAdminChanged event raised by the WalletProxy contract.
type WalletProxyRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_WalletProxy *WalletProxyFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*WalletProxyRoleAdminChangedIterator, error) {

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

	logs, sub, err := _WalletProxy.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &WalletProxyRoleAdminChangedIterator{contract: _WalletProxy.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_WalletProxy *WalletProxyFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *WalletProxyRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _WalletProxy.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletProxyRoleAdminChanged)
				if err := _WalletProxy.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_WalletProxy *WalletProxyFilterer) ParseRoleAdminChanged(log types.Log) (*WalletProxyRoleAdminChanged, error) {
	event := new(WalletProxyRoleAdminChanged)
	if err := _WalletProxy.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WalletProxyRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the WalletProxy contract.
type WalletProxyRoleGrantedIterator struct {
	Event *WalletProxyRoleGranted // Event containing the contract specifics and raw log

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
func (it *WalletProxyRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletProxyRoleGranted)
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
		it.Event = new(WalletProxyRoleGranted)
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
func (it *WalletProxyRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletProxyRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletProxyRoleGranted represents a RoleGranted event raised by the WalletProxy contract.
type WalletProxyRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_WalletProxy *WalletProxyFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*WalletProxyRoleGrantedIterator, error) {

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

	logs, sub, err := _WalletProxy.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &WalletProxyRoleGrantedIterator{contract: _WalletProxy.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_WalletProxy *WalletProxyFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *WalletProxyRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _WalletProxy.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletProxyRoleGranted)
				if err := _WalletProxy.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_WalletProxy *WalletProxyFilterer) ParseRoleGranted(log types.Log) (*WalletProxyRoleGranted, error) {
	event := new(WalletProxyRoleGranted)
	if err := _WalletProxy.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WalletProxyRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the WalletProxy contract.
type WalletProxyRoleRevokedIterator struct {
	Event *WalletProxyRoleRevoked // Event containing the contract specifics and raw log

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
func (it *WalletProxyRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletProxyRoleRevoked)
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
		it.Event = new(WalletProxyRoleRevoked)
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
func (it *WalletProxyRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletProxyRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletProxyRoleRevoked represents a RoleRevoked event raised by the WalletProxy contract.
type WalletProxyRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_WalletProxy *WalletProxyFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*WalletProxyRoleRevokedIterator, error) {

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

	logs, sub, err := _WalletProxy.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &WalletProxyRoleRevokedIterator{contract: _WalletProxy.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_WalletProxy *WalletProxyFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *WalletProxyRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _WalletProxy.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletProxyRoleRevoked)
				if err := _WalletProxy.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_WalletProxy *WalletProxyFilterer) ParseRoleRevoked(log types.Log) (*WalletProxyRoleRevoked, error) {
	event := new(WalletProxyRoleRevoked)
	if err := _WalletProxy.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
