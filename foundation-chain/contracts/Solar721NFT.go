// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package solar721NFT

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

// Solar721NFTMetaData contains all meta data concerning the Solar721NFT contract.
var Solar721NFTMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_orderId\",\"type\":\"uint256\"}],\"name\":\"CancelOrder\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Gift\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sellingPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"feePercent\",\"type\":\"uint16\"}],\"name\":\"MakeOrder\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_tokenURI\",\"type\":\"string\"}],\"name\":\"Mint721Item\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_buyer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_seller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_tradeId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_orderId\",\"type\":\"uint256\"}],\"name\":\"TransferNFT\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_orderId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_sellingPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"_feePercent\",\"type\":\"uint16\"}],\"name\":\"UpdateOrder\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"gift\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint16\",\"name\":\"platformFeePercent\",\"type\":\"uint16\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"hasMintRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"grantMintRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceMintRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"}],\"name\":\"grantMarket\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOrganiser\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"platformFeePercent\",\"type\":\"uint16\"}],\"name\":\"setPlatformFeePercent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPlatformFeePercent\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"tokenURI\",\"type\":\"string\"}],\"name\":\"mintItem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sellingPrice\",\"type\":\"uint256\"}],\"name\":\"makeOrder\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"cancleOrder\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"cancleOrder\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sellingPrice\",\"type\":\"uint256\"}],\"name\":\"updateOrder\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sellingPrice\",\"type\":\"uint256\"}],\"name\":\"updateOrder\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferNFT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getOrderSellingPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getOrderFeePercent\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getNFTOwnerByTokenId\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getSaleOrderIdByToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentMaxSaleOrderId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// Solar721NFTABI is the input ABI used to generate the binding from.
// Deprecated: Use Solar721NFTMetaData.ABI instead.
var Solar721NFTABI = Solar721NFTMetaData.ABI

// Solar721NFT is an auto generated Go binding around an Ethereum contract.
type Solar721NFT struct {
	Solar721NFTCaller     // Read-only binding to the contract
	Solar721NFTTransactor // Write-only binding to the contract
	Solar721NFTFilterer   // Log filterer for contract events
}

// Solar721NFTCaller is an auto generated read-only Go binding around an Ethereum contract.
type Solar721NFTCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Solar721NFTTransactor is an auto generated write-only Go binding around an Ethereum contract.
type Solar721NFTTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Solar721NFTFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Solar721NFTFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Solar721NFTSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Solar721NFTSession struct {
	Contract     *Solar721NFT      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Solar721NFTCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Solar721NFTCallerSession struct {
	Contract *Solar721NFTCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// Solar721NFTTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Solar721NFTTransactorSession struct {
	Contract     *Solar721NFTTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// Solar721NFTRaw is an auto generated low-level Go binding around an Ethereum contract.
type Solar721NFTRaw struct {
	Contract *Solar721NFT // Generic contract binding to access the raw methods on
}

// Solar721NFTCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Solar721NFTCallerRaw struct {
	Contract *Solar721NFTCaller // Generic read-only contract binding to access the raw methods on
}

// Solar721NFTTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Solar721NFTTransactorRaw struct {
	Contract *Solar721NFTTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSolar721NFT creates a new instance of Solar721NFT, bound to a specific deployed contract.
func NewSolar721NFT(address common.Address, backend bind.ContractBackend) (*Solar721NFT, error) {
	contract, err := bindSolar721NFT(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Solar721NFT{Solar721NFTCaller: Solar721NFTCaller{contract: contract}, Solar721NFTTransactor: Solar721NFTTransactor{contract: contract}, Solar721NFTFilterer: Solar721NFTFilterer{contract: contract}}, nil
}

// NewSolar721NFTCaller creates a new read-only instance of Solar721NFT, bound to a specific deployed contract.
func NewSolar721NFTCaller(address common.Address, caller bind.ContractCaller) (*Solar721NFTCaller, error) {
	contract, err := bindSolar721NFT(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Solar721NFTCaller{contract: contract}, nil
}

// NewSolar721NFTTransactor creates a new write-only instance of Solar721NFT, bound to a specific deployed contract.
func NewSolar721NFTTransactor(address common.Address, transactor bind.ContractTransactor) (*Solar721NFTTransactor, error) {
	contract, err := bindSolar721NFT(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Solar721NFTTransactor{contract: contract}, nil
}

// NewSolar721NFTFilterer creates a new log filterer instance of Solar721NFT, bound to a specific deployed contract.
func NewSolar721NFTFilterer(address common.Address, filterer bind.ContractFilterer) (*Solar721NFTFilterer, error) {
	contract, err := bindSolar721NFT(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Solar721NFTFilterer{contract: contract}, nil
}

// bindSolar721NFT binds a generic wrapper to an already deployed contract.
func bindSolar721NFT(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(Solar721NFTABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Solar721NFT *Solar721NFTRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Solar721NFT.Contract.Solar721NFTCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Solar721NFT *Solar721NFTRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Solar721NFT.Contract.Solar721NFTTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Solar721NFT *Solar721NFTRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Solar721NFT.Contract.Solar721NFTTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Solar721NFT *Solar721NFTCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Solar721NFT.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Solar721NFT *Solar721NFTTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Solar721NFT.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Solar721NFT *Solar721NFTTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Solar721NFT.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Solar721NFT *Solar721NFTCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Solar721NFT.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Solar721NFT *Solar721NFTSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Solar721NFT.Contract.DEFAULTADMINROLE(&_Solar721NFT.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Solar721NFT *Solar721NFTCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Solar721NFT.Contract.DEFAULTADMINROLE(&_Solar721NFT.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Solar721NFT *Solar721NFTCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Solar721NFT.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Solar721NFT *Solar721NFTSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Solar721NFT.Contract.BalanceOf(&_Solar721NFT.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Solar721NFT *Solar721NFTCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Solar721NFT.Contract.BalanceOf(&_Solar721NFT.CallOpts, owner)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Solar721NFT *Solar721NFTCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Solar721NFT.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Solar721NFT *Solar721NFTSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Solar721NFT.Contract.GetApproved(&_Solar721NFT.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Solar721NFT *Solar721NFTCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Solar721NFT.Contract.GetApproved(&_Solar721NFT.CallOpts, tokenId)
}

// GetCurrentMaxSaleOrderId is a free data retrieval call binding the contract method 0x69c47870.
//
// Solidity: function getCurrentMaxSaleOrderId() view returns(uint256)
func (_Solar721NFT *Solar721NFTCaller) GetCurrentMaxSaleOrderId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Solar721NFT.contract.Call(opts, &out, "getCurrentMaxSaleOrderId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentMaxSaleOrderId is a free data retrieval call binding the contract method 0x69c47870.
//
// Solidity: function getCurrentMaxSaleOrderId() view returns(uint256)
func (_Solar721NFT *Solar721NFTSession) GetCurrentMaxSaleOrderId() (*big.Int, error) {
	return _Solar721NFT.Contract.GetCurrentMaxSaleOrderId(&_Solar721NFT.CallOpts)
}

// GetCurrentMaxSaleOrderId is a free data retrieval call binding the contract method 0x69c47870.
//
// Solidity: function getCurrentMaxSaleOrderId() view returns(uint256)
func (_Solar721NFT *Solar721NFTCallerSession) GetCurrentMaxSaleOrderId() (*big.Int, error) {
	return _Solar721NFT.Contract.GetCurrentMaxSaleOrderId(&_Solar721NFT.CallOpts)
}

// GetNFTOwnerByTokenId is a free data retrieval call binding the contract method 0xcff68615.
//
// Solidity: function getNFTOwnerByTokenId(uint256 tokenId) view returns(address)
func (_Solar721NFT *Solar721NFTCaller) GetNFTOwnerByTokenId(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Solar721NFT.contract.Call(opts, &out, "getNFTOwnerByTokenId", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetNFTOwnerByTokenId is a free data retrieval call binding the contract method 0xcff68615.
//
// Solidity: function getNFTOwnerByTokenId(uint256 tokenId) view returns(address)
func (_Solar721NFT *Solar721NFTSession) GetNFTOwnerByTokenId(tokenId *big.Int) (common.Address, error) {
	return _Solar721NFT.Contract.GetNFTOwnerByTokenId(&_Solar721NFT.CallOpts, tokenId)
}

// GetNFTOwnerByTokenId is a free data retrieval call binding the contract method 0xcff68615.
//
// Solidity: function getNFTOwnerByTokenId(uint256 tokenId) view returns(address)
func (_Solar721NFT *Solar721NFTCallerSession) GetNFTOwnerByTokenId(tokenId *big.Int) (common.Address, error) {
	return _Solar721NFT.Contract.GetNFTOwnerByTokenId(&_Solar721NFT.CallOpts, tokenId)
}

// GetOrderFeePercent is a free data retrieval call binding the contract method 0xedc59055.
//
// Solidity: function getOrderFeePercent(uint256 tokenId) view returns(uint16)
func (_Solar721NFT *Solar721NFTCaller) GetOrderFeePercent(opts *bind.CallOpts, tokenId *big.Int) (uint16, error) {
	var out []interface{}
	err := _Solar721NFT.contract.Call(opts, &out, "getOrderFeePercent", tokenId)

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// GetOrderFeePercent is a free data retrieval call binding the contract method 0xedc59055.
//
// Solidity: function getOrderFeePercent(uint256 tokenId) view returns(uint16)
func (_Solar721NFT *Solar721NFTSession) GetOrderFeePercent(tokenId *big.Int) (uint16, error) {
	return _Solar721NFT.Contract.GetOrderFeePercent(&_Solar721NFT.CallOpts, tokenId)
}

// GetOrderFeePercent is a free data retrieval call binding the contract method 0xedc59055.
//
// Solidity: function getOrderFeePercent(uint256 tokenId) view returns(uint16)
func (_Solar721NFT *Solar721NFTCallerSession) GetOrderFeePercent(tokenId *big.Int) (uint16, error) {
	return _Solar721NFT.Contract.GetOrderFeePercent(&_Solar721NFT.CallOpts, tokenId)
}

// GetOrderSellingPrice is a free data retrieval call binding the contract method 0xd7e5a39d.
//
// Solidity: function getOrderSellingPrice(uint256 tokenId) view returns(uint256)
func (_Solar721NFT *Solar721NFTCaller) GetOrderSellingPrice(opts *bind.CallOpts, tokenId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Solar721NFT.contract.Call(opts, &out, "getOrderSellingPrice", tokenId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetOrderSellingPrice is a free data retrieval call binding the contract method 0xd7e5a39d.
//
// Solidity: function getOrderSellingPrice(uint256 tokenId) view returns(uint256)
func (_Solar721NFT *Solar721NFTSession) GetOrderSellingPrice(tokenId *big.Int) (*big.Int, error) {
	return _Solar721NFT.Contract.GetOrderSellingPrice(&_Solar721NFT.CallOpts, tokenId)
}

// GetOrderSellingPrice is a free data retrieval call binding the contract method 0xd7e5a39d.
//
// Solidity: function getOrderSellingPrice(uint256 tokenId) view returns(uint256)
func (_Solar721NFT *Solar721NFTCallerSession) GetOrderSellingPrice(tokenId *big.Int) (*big.Int, error) {
	return _Solar721NFT.Contract.GetOrderSellingPrice(&_Solar721NFT.CallOpts, tokenId)
}

// GetOrganiser is a free data retrieval call binding the contract method 0x0b705578.
//
// Solidity: function getOrganiser() view returns(address)
func (_Solar721NFT *Solar721NFTCaller) GetOrganiser(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Solar721NFT.contract.Call(opts, &out, "getOrganiser")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOrganiser is a free data retrieval call binding the contract method 0x0b705578.
//
// Solidity: function getOrganiser() view returns(address)
func (_Solar721NFT *Solar721NFTSession) GetOrganiser() (common.Address, error) {
	return _Solar721NFT.Contract.GetOrganiser(&_Solar721NFT.CallOpts)
}

// GetOrganiser is a free data retrieval call binding the contract method 0x0b705578.
//
// Solidity: function getOrganiser() view returns(address)
func (_Solar721NFT *Solar721NFTCallerSession) GetOrganiser() (common.Address, error) {
	return _Solar721NFT.Contract.GetOrganiser(&_Solar721NFT.CallOpts)
}

// GetPlatformFeePercent is a free data retrieval call binding the contract method 0x4c9c3f85.
//
// Solidity: function getPlatformFeePercent() view returns(uint16)
func (_Solar721NFT *Solar721NFTCaller) GetPlatformFeePercent(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _Solar721NFT.contract.Call(opts, &out, "getPlatformFeePercent")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// GetPlatformFeePercent is a free data retrieval call binding the contract method 0x4c9c3f85.
//
// Solidity: function getPlatformFeePercent() view returns(uint16)
func (_Solar721NFT *Solar721NFTSession) GetPlatformFeePercent() (uint16, error) {
	return _Solar721NFT.Contract.GetPlatformFeePercent(&_Solar721NFT.CallOpts)
}

// GetPlatformFeePercent is a free data retrieval call binding the contract method 0x4c9c3f85.
//
// Solidity: function getPlatformFeePercent() view returns(uint16)
func (_Solar721NFT *Solar721NFTCallerSession) GetPlatformFeePercent() (uint16, error) {
	return _Solar721NFT.Contract.GetPlatformFeePercent(&_Solar721NFT.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Solar721NFT *Solar721NFTCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Solar721NFT.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Solar721NFT *Solar721NFTSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Solar721NFT.Contract.GetRoleAdmin(&_Solar721NFT.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Solar721NFT *Solar721NFTCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Solar721NFT.Contract.GetRoleAdmin(&_Solar721NFT.CallOpts, role)
}

// GetSaleOrderIdByToken is a free data retrieval call binding the contract method 0xa0fbc878.
//
// Solidity: function getSaleOrderIdByToken(uint256 tokenId) view returns(uint256)
func (_Solar721NFT *Solar721NFTCaller) GetSaleOrderIdByToken(opts *bind.CallOpts, tokenId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Solar721NFT.contract.Call(opts, &out, "getSaleOrderIdByToken", tokenId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSaleOrderIdByToken is a free data retrieval call binding the contract method 0xa0fbc878.
//
// Solidity: function getSaleOrderIdByToken(uint256 tokenId) view returns(uint256)
func (_Solar721NFT *Solar721NFTSession) GetSaleOrderIdByToken(tokenId *big.Int) (*big.Int, error) {
	return _Solar721NFT.Contract.GetSaleOrderIdByToken(&_Solar721NFT.CallOpts, tokenId)
}

// GetSaleOrderIdByToken is a free data retrieval call binding the contract method 0xa0fbc878.
//
// Solidity: function getSaleOrderIdByToken(uint256 tokenId) view returns(uint256)
func (_Solar721NFT *Solar721NFTCallerSession) GetSaleOrderIdByToken(tokenId *big.Int) (*big.Int, error) {
	return _Solar721NFT.Contract.GetSaleOrderIdByToken(&_Solar721NFT.CallOpts, tokenId)
}

// HasMintRole is a free data retrieval call binding the contract method 0xc4e58e54.
//
// Solidity: function hasMintRole(address _account) view returns(bool)
func (_Solar721NFT *Solar721NFTCaller) HasMintRole(opts *bind.CallOpts, _account common.Address) (bool, error) {
	var out []interface{}
	err := _Solar721NFT.contract.Call(opts, &out, "hasMintRole", _account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasMintRole is a free data retrieval call binding the contract method 0xc4e58e54.
//
// Solidity: function hasMintRole(address _account) view returns(bool)
func (_Solar721NFT *Solar721NFTSession) HasMintRole(_account common.Address) (bool, error) {
	return _Solar721NFT.Contract.HasMintRole(&_Solar721NFT.CallOpts, _account)
}

// HasMintRole is a free data retrieval call binding the contract method 0xc4e58e54.
//
// Solidity: function hasMintRole(address _account) view returns(bool)
func (_Solar721NFT *Solar721NFTCallerSession) HasMintRole(_account common.Address) (bool, error) {
	return _Solar721NFT.Contract.HasMintRole(&_Solar721NFT.CallOpts, _account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Solar721NFT *Solar721NFTCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Solar721NFT.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Solar721NFT *Solar721NFTSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Solar721NFT.Contract.HasRole(&_Solar721NFT.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Solar721NFT *Solar721NFTCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Solar721NFT.Contract.HasRole(&_Solar721NFT.CallOpts, role, account)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Solar721NFT *Solar721NFTCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Solar721NFT.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Solar721NFT *Solar721NFTSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Solar721NFT.Contract.IsApprovedForAll(&_Solar721NFT.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Solar721NFT *Solar721NFTCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Solar721NFT.Contract.IsApprovedForAll(&_Solar721NFT.CallOpts, owner, operator)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Solar721NFT *Solar721NFTCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Solar721NFT.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Solar721NFT *Solar721NFTSession) Name() (string, error) {
	return _Solar721NFT.Contract.Name(&_Solar721NFT.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Solar721NFT *Solar721NFTCallerSession) Name() (string, error) {
	return _Solar721NFT.Contract.Name(&_Solar721NFT.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Solar721NFT *Solar721NFTCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Solar721NFT.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Solar721NFT *Solar721NFTSession) Owner() (common.Address, error) {
	return _Solar721NFT.Contract.Owner(&_Solar721NFT.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Solar721NFT *Solar721NFTCallerSession) Owner() (common.Address, error) {
	return _Solar721NFT.Contract.Owner(&_Solar721NFT.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Solar721NFT *Solar721NFTCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Solar721NFT.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Solar721NFT *Solar721NFTSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Solar721NFT.Contract.OwnerOf(&_Solar721NFT.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Solar721NFT *Solar721NFTCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Solar721NFT.Contract.OwnerOf(&_Solar721NFT.CallOpts, tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Solar721NFT *Solar721NFTCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Solar721NFT.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Solar721NFT *Solar721NFTSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Solar721NFT.Contract.SupportsInterface(&_Solar721NFT.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Solar721NFT *Solar721NFTCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Solar721NFT.Contract.SupportsInterface(&_Solar721NFT.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Solar721NFT *Solar721NFTCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Solar721NFT.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Solar721NFT *Solar721NFTSession) Symbol() (string, error) {
	return _Solar721NFT.Contract.Symbol(&_Solar721NFT.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Solar721NFT *Solar721NFTCallerSession) Symbol() (string, error) {
	return _Solar721NFT.Contract.Symbol(&_Solar721NFT.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Solar721NFT *Solar721NFTCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Solar721NFT.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Solar721NFT *Solar721NFTSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Solar721NFT.Contract.TokenURI(&_Solar721NFT.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Solar721NFT *Solar721NFTCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Solar721NFT.Contract.TokenURI(&_Solar721NFT.CallOpts, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Solar721NFT *Solar721NFTTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Solar721NFT.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Solar721NFT *Solar721NFTSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Solar721NFT.Contract.Approve(&_Solar721NFT.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Solar721NFT *Solar721NFTTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Solar721NFT.Contract.Approve(&_Solar721NFT.TransactOpts, to, tokenId)
}

// CancleOrder is a paid mutator transaction binding the contract method 0x9c589601.
//
// Solidity: function cancleOrder(address operator, uint256 tokenId) returns(bool)
func (_Solar721NFT *Solar721NFTTransactor) CancleOrder(opts *bind.TransactOpts, operator common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Solar721NFT.contract.Transact(opts, "cancleOrder", operator, tokenId)
}

// CancleOrder is a paid mutator transaction binding the contract method 0x9c589601.
//
// Solidity: function cancleOrder(address operator, uint256 tokenId) returns(bool)
func (_Solar721NFT *Solar721NFTSession) CancleOrder(operator common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Solar721NFT.Contract.CancleOrder(&_Solar721NFT.TransactOpts, operator, tokenId)
}

// CancleOrder is a paid mutator transaction binding the contract method 0x9c589601.
//
// Solidity: function cancleOrder(address operator, uint256 tokenId) returns(bool)
func (_Solar721NFT *Solar721NFTTransactorSession) CancleOrder(operator common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Solar721NFT.Contract.CancleOrder(&_Solar721NFT.TransactOpts, operator, tokenId)
}

// CancleOrder0 is a paid mutator transaction binding the contract method 0xf406cb8a.
//
// Solidity: function cancleOrder(uint256 tokenId) returns(bool)
func (_Solar721NFT *Solar721NFTTransactor) CancleOrder0(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _Solar721NFT.contract.Transact(opts, "cancleOrder0", tokenId)
}

// CancleOrder0 is a paid mutator transaction binding the contract method 0xf406cb8a.
//
// Solidity: function cancleOrder(uint256 tokenId) returns(bool)
func (_Solar721NFT *Solar721NFTSession) CancleOrder0(tokenId *big.Int) (*types.Transaction, error) {
	return _Solar721NFT.Contract.CancleOrder0(&_Solar721NFT.TransactOpts, tokenId)
}

// CancleOrder0 is a paid mutator transaction binding the contract method 0xf406cb8a.
//
// Solidity: function cancleOrder(uint256 tokenId) returns(bool)
func (_Solar721NFT *Solar721NFTTransactorSession) CancleOrder0(tokenId *big.Int) (*types.Transaction, error) {
	return _Solar721NFT.Contract.CancleOrder0(&_Solar721NFT.TransactOpts, tokenId)
}

// Gift is a paid mutator transaction binding the contract method 0x95dbee2d.
//
// Solidity: function gift(address from, address to, uint256 tokenId) returns()
func (_Solar721NFT *Solar721NFTTransactor) Gift(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Solar721NFT.contract.Transact(opts, "gift", from, to, tokenId)
}

// Gift is a paid mutator transaction binding the contract method 0x95dbee2d.
//
// Solidity: function gift(address from, address to, uint256 tokenId) returns()
func (_Solar721NFT *Solar721NFTSession) Gift(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Solar721NFT.Contract.Gift(&_Solar721NFT.TransactOpts, from, to, tokenId)
}

// Gift is a paid mutator transaction binding the contract method 0x95dbee2d.
//
// Solidity: function gift(address from, address to, uint256 tokenId) returns()
func (_Solar721NFT *Solar721NFTTransactorSession) Gift(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Solar721NFT.Contract.Gift(&_Solar721NFT.TransactOpts, from, to, tokenId)
}

// GrantMarket is a paid mutator transaction binding the contract method 0x040add28.
//
// Solidity: function grantMarket(address market) returns(bool)
func (_Solar721NFT *Solar721NFTTransactor) GrantMarket(opts *bind.TransactOpts, market common.Address) (*types.Transaction, error) {
	return _Solar721NFT.contract.Transact(opts, "grantMarket", market)
}

// GrantMarket is a paid mutator transaction binding the contract method 0x040add28.
//
// Solidity: function grantMarket(address market) returns(bool)
func (_Solar721NFT *Solar721NFTSession) GrantMarket(market common.Address) (*types.Transaction, error) {
	return _Solar721NFT.Contract.GrantMarket(&_Solar721NFT.TransactOpts, market)
}

// GrantMarket is a paid mutator transaction binding the contract method 0x040add28.
//
// Solidity: function grantMarket(address market) returns(bool)
func (_Solar721NFT *Solar721NFTTransactorSession) GrantMarket(market common.Address) (*types.Transaction, error) {
	return _Solar721NFT.Contract.GrantMarket(&_Solar721NFT.TransactOpts, market)
}

// GrantMintRole is a paid mutator transaction binding the contract method 0xc2e3273d.
//
// Solidity: function grantMintRole(address operator) returns()
func (_Solar721NFT *Solar721NFTTransactor) GrantMintRole(opts *bind.TransactOpts, operator common.Address) (*types.Transaction, error) {
	return _Solar721NFT.contract.Transact(opts, "grantMintRole", operator)
}

// GrantMintRole is a paid mutator transaction binding the contract method 0xc2e3273d.
//
// Solidity: function grantMintRole(address operator) returns()
func (_Solar721NFT *Solar721NFTSession) GrantMintRole(operator common.Address) (*types.Transaction, error) {
	return _Solar721NFT.Contract.GrantMintRole(&_Solar721NFT.TransactOpts, operator)
}

// GrantMintRole is a paid mutator transaction binding the contract method 0xc2e3273d.
//
// Solidity: function grantMintRole(address operator) returns()
func (_Solar721NFT *Solar721NFTTransactorSession) GrantMintRole(operator common.Address) (*types.Transaction, error) {
	return _Solar721NFT.Contract.GrantMintRole(&_Solar721NFT.TransactOpts, operator)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Solar721NFT *Solar721NFTTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Solar721NFT.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Solar721NFT *Solar721NFTSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Solar721NFT.Contract.GrantRole(&_Solar721NFT.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Solar721NFT *Solar721NFTTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Solar721NFT.Contract.GrantRole(&_Solar721NFT.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0x2549edd0.
//
// Solidity: function initialize(string name, string symbol, uint16 platformFeePercent) returns()
func (_Solar721NFT *Solar721NFTTransactor) Initialize(opts *bind.TransactOpts, name string, symbol string, platformFeePercent uint16) (*types.Transaction, error) {
	return _Solar721NFT.contract.Transact(opts, "initialize", name, symbol, platformFeePercent)
}

// Initialize is a paid mutator transaction binding the contract method 0x2549edd0.
//
// Solidity: function initialize(string name, string symbol, uint16 platformFeePercent) returns()
func (_Solar721NFT *Solar721NFTSession) Initialize(name string, symbol string, platformFeePercent uint16) (*types.Transaction, error) {
	return _Solar721NFT.Contract.Initialize(&_Solar721NFT.TransactOpts, name, symbol, platformFeePercent)
}

// Initialize is a paid mutator transaction binding the contract method 0x2549edd0.
//
// Solidity: function initialize(string name, string symbol, uint16 platformFeePercent) returns()
func (_Solar721NFT *Solar721NFTTransactorSession) Initialize(name string, symbol string, platformFeePercent uint16) (*types.Transaction, error) {
	return _Solar721NFT.Contract.Initialize(&_Solar721NFT.TransactOpts, name, symbol, platformFeePercent)
}

// MakeOrder is a paid mutator transaction binding the contract method 0x9bc80102.
//
// Solidity: function makeOrder(uint256 tokenId, uint256 sellingPrice) returns(uint256)
func (_Solar721NFT *Solar721NFTTransactor) MakeOrder(opts *bind.TransactOpts, tokenId *big.Int, sellingPrice *big.Int) (*types.Transaction, error) {
	return _Solar721NFT.contract.Transact(opts, "makeOrder", tokenId, sellingPrice)
}

// MakeOrder is a paid mutator transaction binding the contract method 0x9bc80102.
//
// Solidity: function makeOrder(uint256 tokenId, uint256 sellingPrice) returns(uint256)
func (_Solar721NFT *Solar721NFTSession) MakeOrder(tokenId *big.Int, sellingPrice *big.Int) (*types.Transaction, error) {
	return _Solar721NFT.Contract.MakeOrder(&_Solar721NFT.TransactOpts, tokenId, sellingPrice)
}

// MakeOrder is a paid mutator transaction binding the contract method 0x9bc80102.
//
// Solidity: function makeOrder(uint256 tokenId, uint256 sellingPrice) returns(uint256)
func (_Solar721NFT *Solar721NFTTransactorSession) MakeOrder(tokenId *big.Int, sellingPrice *big.Int) (*types.Transaction, error) {
	return _Solar721NFT.Contract.MakeOrder(&_Solar721NFT.TransactOpts, tokenId, sellingPrice)
}

// MintItem is a paid mutator transaction binding the contract method 0x8337afdc.
//
// Solidity: function mintItem(address to, uint256 tokenId, string tokenURI) returns(uint256)
func (_Solar721NFT *Solar721NFTTransactor) MintItem(opts *bind.TransactOpts, to common.Address, tokenId *big.Int, tokenURI string) (*types.Transaction, error) {
	return _Solar721NFT.contract.Transact(opts, "mintItem", to, tokenId, tokenURI)
}

// MintItem is a paid mutator transaction binding the contract method 0x8337afdc.
//
// Solidity: function mintItem(address to, uint256 tokenId, string tokenURI) returns(uint256)
func (_Solar721NFT *Solar721NFTSession) MintItem(to common.Address, tokenId *big.Int, tokenURI string) (*types.Transaction, error) {
	return _Solar721NFT.Contract.MintItem(&_Solar721NFT.TransactOpts, to, tokenId, tokenURI)
}

// MintItem is a paid mutator transaction binding the contract method 0x8337afdc.
//
// Solidity: function mintItem(address to, uint256 tokenId, string tokenURI) returns(uint256)
func (_Solar721NFT *Solar721NFTTransactorSession) MintItem(to common.Address, tokenId *big.Int, tokenURI string) (*types.Transaction, error) {
	return _Solar721NFT.Contract.MintItem(&_Solar721NFT.TransactOpts, to, tokenId, tokenURI)
}

// RenounceMintRole is a paid mutator transaction binding the contract method 0xa49a177a.
//
// Solidity: function renounceMintRole(address account) returns()
func (_Solar721NFT *Solar721NFTTransactor) RenounceMintRole(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _Solar721NFT.contract.Transact(opts, "renounceMintRole", account)
}

// RenounceMintRole is a paid mutator transaction binding the contract method 0xa49a177a.
//
// Solidity: function renounceMintRole(address account) returns()
func (_Solar721NFT *Solar721NFTSession) RenounceMintRole(account common.Address) (*types.Transaction, error) {
	return _Solar721NFT.Contract.RenounceMintRole(&_Solar721NFT.TransactOpts, account)
}

// RenounceMintRole is a paid mutator transaction binding the contract method 0xa49a177a.
//
// Solidity: function renounceMintRole(address account) returns()
func (_Solar721NFT *Solar721NFTTransactorSession) RenounceMintRole(account common.Address) (*types.Transaction, error) {
	return _Solar721NFT.Contract.RenounceMintRole(&_Solar721NFT.TransactOpts, account)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Solar721NFT *Solar721NFTTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Solar721NFT.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Solar721NFT *Solar721NFTSession) RenounceOwnership() (*types.Transaction, error) {
	return _Solar721NFT.Contract.RenounceOwnership(&_Solar721NFT.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Solar721NFT *Solar721NFTTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Solar721NFT.Contract.RenounceOwnership(&_Solar721NFT.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Solar721NFT *Solar721NFTTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Solar721NFT.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Solar721NFT *Solar721NFTSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Solar721NFT.Contract.RenounceRole(&_Solar721NFT.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Solar721NFT *Solar721NFTTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Solar721NFT.Contract.RenounceRole(&_Solar721NFT.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Solar721NFT *Solar721NFTTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Solar721NFT.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Solar721NFT *Solar721NFTSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Solar721NFT.Contract.RevokeRole(&_Solar721NFT.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Solar721NFT *Solar721NFTTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Solar721NFT.Contract.RevokeRole(&_Solar721NFT.TransactOpts, role, account)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Solar721NFT *Solar721NFTTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Solar721NFT.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Solar721NFT *Solar721NFTSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Solar721NFT.Contract.SafeTransferFrom(&_Solar721NFT.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Solar721NFT *Solar721NFTTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Solar721NFT.Contract.SafeTransferFrom(&_Solar721NFT.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Solar721NFT *Solar721NFTTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Solar721NFT.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Solar721NFT *Solar721NFTSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Solar721NFT.Contract.SafeTransferFrom0(&_Solar721NFT.TransactOpts, from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Solar721NFT *Solar721NFTTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Solar721NFT.Contract.SafeTransferFrom0(&_Solar721NFT.TransactOpts, from, to, tokenId, _data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Solar721NFT *Solar721NFTTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _Solar721NFT.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Solar721NFT *Solar721NFTSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Solar721NFT.Contract.SetApprovalForAll(&_Solar721NFT.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Solar721NFT *Solar721NFTTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Solar721NFT.Contract.SetApprovalForAll(&_Solar721NFT.TransactOpts, operator, approved)
}

// SetPlatformFeePercent is a paid mutator transaction binding the contract method 0x685538df.
//
// Solidity: function setPlatformFeePercent(uint16 platformFeePercent) returns()
func (_Solar721NFT *Solar721NFTTransactor) SetPlatformFeePercent(opts *bind.TransactOpts, platformFeePercent uint16) (*types.Transaction, error) {
	return _Solar721NFT.contract.Transact(opts, "setPlatformFeePercent", platformFeePercent)
}

// SetPlatformFeePercent is a paid mutator transaction binding the contract method 0x685538df.
//
// Solidity: function setPlatformFeePercent(uint16 platformFeePercent) returns()
func (_Solar721NFT *Solar721NFTSession) SetPlatformFeePercent(platformFeePercent uint16) (*types.Transaction, error) {
	return _Solar721NFT.Contract.SetPlatformFeePercent(&_Solar721NFT.TransactOpts, platformFeePercent)
}

// SetPlatformFeePercent is a paid mutator transaction binding the contract method 0x685538df.
//
// Solidity: function setPlatformFeePercent(uint16 platformFeePercent) returns()
func (_Solar721NFT *Solar721NFTTransactorSession) SetPlatformFeePercent(platformFeePercent uint16) (*types.Transaction, error) {
	return _Solar721NFT.Contract.SetPlatformFeePercent(&_Solar721NFT.TransactOpts, platformFeePercent)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Solar721NFT *Solar721NFTTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Solar721NFT.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Solar721NFT *Solar721NFTSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Solar721NFT.Contract.TransferFrom(&_Solar721NFT.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Solar721NFT *Solar721NFTTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Solar721NFT.Contract.TransferFrom(&_Solar721NFT.TransactOpts, from, to, tokenId)
}

// TransferNFT is a paid mutator transaction binding the contract method 0x94ab67fe.
//
// Solidity: function transferNFT(address buyer, uint256 tokenId) returns(uint256, uint256)
func (_Solar721NFT *Solar721NFTTransactor) TransferNFT(opts *bind.TransactOpts, buyer common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Solar721NFT.contract.Transact(opts, "transferNFT", buyer, tokenId)
}

// TransferNFT is a paid mutator transaction binding the contract method 0x94ab67fe.
//
// Solidity: function transferNFT(address buyer, uint256 tokenId) returns(uint256, uint256)
func (_Solar721NFT *Solar721NFTSession) TransferNFT(buyer common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Solar721NFT.Contract.TransferNFT(&_Solar721NFT.TransactOpts, buyer, tokenId)
}

// TransferNFT is a paid mutator transaction binding the contract method 0x94ab67fe.
//
// Solidity: function transferNFT(address buyer, uint256 tokenId) returns(uint256, uint256)
func (_Solar721NFT *Solar721NFTTransactorSession) TransferNFT(buyer common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Solar721NFT.Contract.TransferNFT(&_Solar721NFT.TransactOpts, buyer, tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Solar721NFT *Solar721NFTTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Solar721NFT.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Solar721NFT *Solar721NFTSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Solar721NFT.Contract.TransferOwnership(&_Solar721NFT.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Solar721NFT *Solar721NFTTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Solar721NFT.Contract.TransferOwnership(&_Solar721NFT.TransactOpts, newOwner)
}

// UpdateOrder is a paid mutator transaction binding the contract method 0x587b71a6.
//
// Solidity: function updateOrder(uint256 tokenId, uint256 sellingPrice) returns(bool)
func (_Solar721NFT *Solar721NFTTransactor) UpdateOrder(opts *bind.TransactOpts, tokenId *big.Int, sellingPrice *big.Int) (*types.Transaction, error) {
	return _Solar721NFT.contract.Transact(opts, "updateOrder", tokenId, sellingPrice)
}

// UpdateOrder is a paid mutator transaction binding the contract method 0x587b71a6.
//
// Solidity: function updateOrder(uint256 tokenId, uint256 sellingPrice) returns(bool)
func (_Solar721NFT *Solar721NFTSession) UpdateOrder(tokenId *big.Int, sellingPrice *big.Int) (*types.Transaction, error) {
	return _Solar721NFT.Contract.UpdateOrder(&_Solar721NFT.TransactOpts, tokenId, sellingPrice)
}

// UpdateOrder is a paid mutator transaction binding the contract method 0x587b71a6.
//
// Solidity: function updateOrder(uint256 tokenId, uint256 sellingPrice) returns(bool)
func (_Solar721NFT *Solar721NFTTransactorSession) UpdateOrder(tokenId *big.Int, sellingPrice *big.Int) (*types.Transaction, error) {
	return _Solar721NFT.Contract.UpdateOrder(&_Solar721NFT.TransactOpts, tokenId, sellingPrice)
}

// UpdateOrder0 is a paid mutator transaction binding the contract method 0xf7f33eda.
//
// Solidity: function updateOrder(address owner, uint256 tokenId, uint256 sellingPrice) returns(bool)
func (_Solar721NFT *Solar721NFTTransactor) UpdateOrder0(opts *bind.TransactOpts, owner common.Address, tokenId *big.Int, sellingPrice *big.Int) (*types.Transaction, error) {
	return _Solar721NFT.contract.Transact(opts, "updateOrder0", owner, tokenId, sellingPrice)
}

// UpdateOrder0 is a paid mutator transaction binding the contract method 0xf7f33eda.
//
// Solidity: function updateOrder(address owner, uint256 tokenId, uint256 sellingPrice) returns(bool)
func (_Solar721NFT *Solar721NFTSession) UpdateOrder0(owner common.Address, tokenId *big.Int, sellingPrice *big.Int) (*types.Transaction, error) {
	return _Solar721NFT.Contract.UpdateOrder0(&_Solar721NFT.TransactOpts, owner, tokenId, sellingPrice)
}

// UpdateOrder0 is a paid mutator transaction binding the contract method 0xf7f33eda.
//
// Solidity: function updateOrder(address owner, uint256 tokenId, uint256 sellingPrice) returns(bool)
func (_Solar721NFT *Solar721NFTTransactorSession) UpdateOrder0(owner common.Address, tokenId *big.Int, sellingPrice *big.Int) (*types.Transaction, error) {
	return _Solar721NFT.Contract.UpdateOrder0(&_Solar721NFT.TransactOpts, owner, tokenId, sellingPrice)
}

// Solar721NFTApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Solar721NFT contract.
type Solar721NFTApprovalIterator struct {
	Event *Solar721NFTApproval // Event containing the contract specifics and raw log

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
func (it *Solar721NFTApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Solar721NFTApproval)
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
		it.Event = new(Solar721NFTApproval)
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
func (it *Solar721NFTApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Solar721NFTApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Solar721NFTApproval represents a Approval event raised by the Solar721NFT contract.
type Solar721NFTApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Solar721NFT *Solar721NFTFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*Solar721NFTApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Solar721NFT.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &Solar721NFTApprovalIterator{contract: _Solar721NFT.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Solar721NFT *Solar721NFTFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *Solar721NFTApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Solar721NFT.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Solar721NFTApproval)
				if err := _Solar721NFT.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Solar721NFT *Solar721NFTFilterer) ParseApproval(log types.Log) (*Solar721NFTApproval, error) {
	event := new(Solar721NFTApproval)
	if err := _Solar721NFT.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Solar721NFTApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Solar721NFT contract.
type Solar721NFTApprovalForAllIterator struct {
	Event *Solar721NFTApprovalForAll // Event containing the contract specifics and raw log

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
func (it *Solar721NFTApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Solar721NFTApprovalForAll)
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
		it.Event = new(Solar721NFTApprovalForAll)
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
func (it *Solar721NFTApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Solar721NFTApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Solar721NFTApprovalForAll represents a ApprovalForAll event raised by the Solar721NFT contract.
type Solar721NFTApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Solar721NFT *Solar721NFTFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*Solar721NFTApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Solar721NFT.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &Solar721NFTApprovalForAllIterator{contract: _Solar721NFT.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Solar721NFT *Solar721NFTFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *Solar721NFTApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Solar721NFT.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Solar721NFTApprovalForAll)
				if err := _Solar721NFT.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Solar721NFT *Solar721NFTFilterer) ParseApprovalForAll(log types.Log) (*Solar721NFTApprovalForAll, error) {
	event := new(Solar721NFTApprovalForAll)
	if err := _Solar721NFT.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Solar721NFTCancelOrderIterator is returned from FilterCancelOrder and is used to iterate over the raw logs and unpacked data for CancelOrder events raised by the Solar721NFT contract.
type Solar721NFTCancelOrderIterator struct {
	Event *Solar721NFTCancelOrder // Event containing the contract specifics and raw log

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
func (it *Solar721NFTCancelOrderIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Solar721NFTCancelOrder)
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
		it.Event = new(Solar721NFTCancelOrder)
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
func (it *Solar721NFTCancelOrderIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Solar721NFTCancelOrderIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Solar721NFTCancelOrder represents a CancelOrder event raised by the Solar721NFT contract.
type Solar721NFTCancelOrder struct {
	Operator common.Address
	TokenId  *big.Int
	OrderId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterCancelOrder is a free log retrieval operation binding the contract event 0x1e0e8a7901ee871b939f18e3937f188af21bf9885762ec9c84917abc9fbaf1fa.
//
// Solidity: event CancelOrder(address indexed _operator, uint256 indexed _tokenId, uint256 indexed _orderId)
func (_Solar721NFT *Solar721NFTFilterer) FilterCancelOrder(opts *bind.FilterOpts, _operator []common.Address, _tokenId []*big.Int, _orderId []*big.Int) (*Solar721NFTCancelOrderIterator, error) {

	var _operatorRule []interface{}
	for _, _operatorItem := range _operator {
		_operatorRule = append(_operatorRule, _operatorItem)
	}
	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}
	var _orderIdRule []interface{}
	for _, _orderIdItem := range _orderId {
		_orderIdRule = append(_orderIdRule, _orderIdItem)
	}

	logs, sub, err := _Solar721NFT.contract.FilterLogs(opts, "CancelOrder", _operatorRule, _tokenIdRule, _orderIdRule)
	if err != nil {
		return nil, err
	}
	return &Solar721NFTCancelOrderIterator{contract: _Solar721NFT.contract, event: "CancelOrder", logs: logs, sub: sub}, nil
}

// WatchCancelOrder is a free log subscription operation binding the contract event 0x1e0e8a7901ee871b939f18e3937f188af21bf9885762ec9c84917abc9fbaf1fa.
//
// Solidity: event CancelOrder(address indexed _operator, uint256 indexed _tokenId, uint256 indexed _orderId)
func (_Solar721NFT *Solar721NFTFilterer) WatchCancelOrder(opts *bind.WatchOpts, sink chan<- *Solar721NFTCancelOrder, _operator []common.Address, _tokenId []*big.Int, _orderId []*big.Int) (event.Subscription, error) {

	var _operatorRule []interface{}
	for _, _operatorItem := range _operator {
		_operatorRule = append(_operatorRule, _operatorItem)
	}
	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}
	var _orderIdRule []interface{}
	for _, _orderIdItem := range _orderId {
		_orderIdRule = append(_orderIdRule, _orderIdItem)
	}

	logs, sub, err := _Solar721NFT.contract.WatchLogs(opts, "CancelOrder", _operatorRule, _tokenIdRule, _orderIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Solar721NFTCancelOrder)
				if err := _Solar721NFT.contract.UnpackLog(event, "CancelOrder", log); err != nil {
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

// ParseCancelOrder is a log parse operation binding the contract event 0x1e0e8a7901ee871b939f18e3937f188af21bf9885762ec9c84917abc9fbaf1fa.
//
// Solidity: event CancelOrder(address indexed _operator, uint256 indexed _tokenId, uint256 indexed _orderId)
func (_Solar721NFT *Solar721NFTFilterer) ParseCancelOrder(log types.Log) (*Solar721NFTCancelOrder, error) {
	event := new(Solar721NFTCancelOrder)
	if err := _Solar721NFT.contract.UnpackLog(event, "CancelOrder", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Solar721NFTGiftIterator is returned from FilterGift and is used to iterate over the raw logs and unpacked data for Gift events raised by the Solar721NFT contract.
type Solar721NFTGiftIterator struct {
	Event *Solar721NFTGift // Event containing the contract specifics and raw log

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
func (it *Solar721NFTGiftIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Solar721NFTGift)
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
		it.Event = new(Solar721NFTGift)
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
func (it *Solar721NFTGiftIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Solar721NFTGiftIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Solar721NFTGift represents a Gift event raised by the Solar721NFT contract.
type Solar721NFTGift struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterGift is a free log retrieval operation binding the contract event 0xa64e36941fd6ccca91d7915b34e77bde711b5ec2338168949868a32e3e5dd1fc.
//
// Solidity: event Gift(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Solar721NFT *Solar721NFTFilterer) FilterGift(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*Solar721NFTGiftIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Solar721NFT.contract.FilterLogs(opts, "Gift", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &Solar721NFTGiftIterator{contract: _Solar721NFT.contract, event: "Gift", logs: logs, sub: sub}, nil
}

// WatchGift is a free log subscription operation binding the contract event 0xa64e36941fd6ccca91d7915b34e77bde711b5ec2338168949868a32e3e5dd1fc.
//
// Solidity: event Gift(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Solar721NFT *Solar721NFTFilterer) WatchGift(opts *bind.WatchOpts, sink chan<- *Solar721NFTGift, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Solar721NFT.contract.WatchLogs(opts, "Gift", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Solar721NFTGift)
				if err := _Solar721NFT.contract.UnpackLog(event, "Gift", log); err != nil {
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

// ParseGift is a log parse operation binding the contract event 0xa64e36941fd6ccca91d7915b34e77bde711b5ec2338168949868a32e3e5dd1fc.
//
// Solidity: event Gift(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Solar721NFT *Solar721NFTFilterer) ParseGift(log types.Log) (*Solar721NFTGift, error) {
	event := new(Solar721NFTGift)
	if err := _Solar721NFT.contract.UnpackLog(event, "Gift", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Solar721NFTMakeOrderIterator is returned from FilterMakeOrder and is used to iterate over the raw logs and unpacked data for MakeOrder events raised by the Solar721NFT contract.
type Solar721NFTMakeOrderIterator struct {
	Event *Solar721NFTMakeOrder // Event containing the contract specifics and raw log

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
func (it *Solar721NFTMakeOrderIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Solar721NFTMakeOrder)
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
		it.Event = new(Solar721NFTMakeOrder)
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
func (it *Solar721NFTMakeOrderIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Solar721NFTMakeOrderIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Solar721NFTMakeOrder represents a MakeOrder event raised by the Solar721NFT contract.
type Solar721NFTMakeOrder struct {
	Operator     common.Address
	TokenId      *big.Int
	OrderId      *big.Int
	SellingPrice *big.Int
	FeePercent   uint16
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterMakeOrder is a free log retrieval operation binding the contract event 0x68c4798c8d7b0fcc2897f7ea2655def6035ac4d0e363dfae63e68b1de3d19d10.
//
// Solidity: event MakeOrder(address indexed operator, uint256 indexed tokenId, uint256 indexed orderId, uint256 sellingPrice, uint16 feePercent)
func (_Solar721NFT *Solar721NFTFilterer) FilterMakeOrder(opts *bind.FilterOpts, operator []common.Address, tokenId []*big.Int, orderId []*big.Int) (*Solar721NFTMakeOrderIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}

	logs, sub, err := _Solar721NFT.contract.FilterLogs(opts, "MakeOrder", operatorRule, tokenIdRule, orderIdRule)
	if err != nil {
		return nil, err
	}
	return &Solar721NFTMakeOrderIterator{contract: _Solar721NFT.contract, event: "MakeOrder", logs: logs, sub: sub}, nil
}

// WatchMakeOrder is a free log subscription operation binding the contract event 0x68c4798c8d7b0fcc2897f7ea2655def6035ac4d0e363dfae63e68b1de3d19d10.
//
// Solidity: event MakeOrder(address indexed operator, uint256 indexed tokenId, uint256 indexed orderId, uint256 sellingPrice, uint16 feePercent)
func (_Solar721NFT *Solar721NFTFilterer) WatchMakeOrder(opts *bind.WatchOpts, sink chan<- *Solar721NFTMakeOrder, operator []common.Address, tokenId []*big.Int, orderId []*big.Int) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}

	logs, sub, err := _Solar721NFT.contract.WatchLogs(opts, "MakeOrder", operatorRule, tokenIdRule, orderIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Solar721NFTMakeOrder)
				if err := _Solar721NFT.contract.UnpackLog(event, "MakeOrder", log); err != nil {
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

// ParseMakeOrder is a log parse operation binding the contract event 0x68c4798c8d7b0fcc2897f7ea2655def6035ac4d0e363dfae63e68b1de3d19d10.
//
// Solidity: event MakeOrder(address indexed operator, uint256 indexed tokenId, uint256 indexed orderId, uint256 sellingPrice, uint16 feePercent)
func (_Solar721NFT *Solar721NFTFilterer) ParseMakeOrder(log types.Log) (*Solar721NFTMakeOrder, error) {
	event := new(Solar721NFTMakeOrder)
	if err := _Solar721NFT.contract.UnpackLog(event, "MakeOrder", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Solar721NFTMint721ItemIterator is returned from FilterMint721Item and is used to iterate over the raw logs and unpacked data for Mint721Item events raised by the Solar721NFT contract.
type Solar721NFTMint721ItemIterator struct {
	Event *Solar721NFTMint721Item // Event containing the contract specifics and raw log

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
func (it *Solar721NFTMint721ItemIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Solar721NFTMint721Item)
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
		it.Event = new(Solar721NFTMint721Item)
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
func (it *Solar721NFTMint721ItemIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Solar721NFTMint721ItemIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Solar721NFTMint721Item represents a Mint721Item event raised by the Solar721NFT contract.
type Solar721NFTMint721Item struct {
	Operator common.Address
	To       common.Address
	TokenId  *big.Int
	Amount   *big.Int
	TokenURI string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterMint721Item is a free log retrieval operation binding the contract event 0x3b20d8fb008c95b0a7cc81ea049761553e3c8fa49f6d8903222bea81964de562.
//
// Solidity: event Mint721Item(address _operator, address indexed _to, uint256 indexed _tokenId, uint256 indexed _amount, string _tokenURI)
func (_Solar721NFT *Solar721NFTFilterer) FilterMint721Item(opts *bind.FilterOpts, _to []common.Address, _tokenId []*big.Int, _amount []*big.Int) (*Solar721NFTMint721ItemIterator, error) {

	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}
	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}
	var _amountRule []interface{}
	for _, _amountItem := range _amount {
		_amountRule = append(_amountRule, _amountItem)
	}

	logs, sub, err := _Solar721NFT.contract.FilterLogs(opts, "Mint721Item", _toRule, _tokenIdRule, _amountRule)
	if err != nil {
		return nil, err
	}
	return &Solar721NFTMint721ItemIterator{contract: _Solar721NFT.contract, event: "Mint721Item", logs: logs, sub: sub}, nil
}

// WatchMint721Item is a free log subscription operation binding the contract event 0x3b20d8fb008c95b0a7cc81ea049761553e3c8fa49f6d8903222bea81964de562.
//
// Solidity: event Mint721Item(address _operator, address indexed _to, uint256 indexed _tokenId, uint256 indexed _amount, string _tokenURI)
func (_Solar721NFT *Solar721NFTFilterer) WatchMint721Item(opts *bind.WatchOpts, sink chan<- *Solar721NFTMint721Item, _to []common.Address, _tokenId []*big.Int, _amount []*big.Int) (event.Subscription, error) {

	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}
	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}
	var _amountRule []interface{}
	for _, _amountItem := range _amount {
		_amountRule = append(_amountRule, _amountItem)
	}

	logs, sub, err := _Solar721NFT.contract.WatchLogs(opts, "Mint721Item", _toRule, _tokenIdRule, _amountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Solar721NFTMint721Item)
				if err := _Solar721NFT.contract.UnpackLog(event, "Mint721Item", log); err != nil {
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

// ParseMint721Item is a log parse operation binding the contract event 0x3b20d8fb008c95b0a7cc81ea049761553e3c8fa49f6d8903222bea81964de562.
//
// Solidity: event Mint721Item(address _operator, address indexed _to, uint256 indexed _tokenId, uint256 indexed _amount, string _tokenURI)
func (_Solar721NFT *Solar721NFTFilterer) ParseMint721Item(log types.Log) (*Solar721NFTMint721Item, error) {
	event := new(Solar721NFTMint721Item)
	if err := _Solar721NFT.contract.UnpackLog(event, "Mint721Item", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Solar721NFTOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Solar721NFT contract.
type Solar721NFTOwnershipTransferredIterator struct {
	Event *Solar721NFTOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *Solar721NFTOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Solar721NFTOwnershipTransferred)
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
		it.Event = new(Solar721NFTOwnershipTransferred)
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
func (it *Solar721NFTOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Solar721NFTOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Solar721NFTOwnershipTransferred represents a OwnershipTransferred event raised by the Solar721NFT contract.
type Solar721NFTOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Solar721NFT *Solar721NFTFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*Solar721NFTOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Solar721NFT.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &Solar721NFTOwnershipTransferredIterator{contract: _Solar721NFT.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Solar721NFT *Solar721NFTFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *Solar721NFTOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Solar721NFT.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Solar721NFTOwnershipTransferred)
				if err := _Solar721NFT.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Solar721NFT *Solar721NFTFilterer) ParseOwnershipTransferred(log types.Log) (*Solar721NFTOwnershipTransferred, error) {
	event := new(Solar721NFTOwnershipTransferred)
	if err := _Solar721NFT.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Solar721NFTRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Solar721NFT contract.
type Solar721NFTRoleAdminChangedIterator struct {
	Event *Solar721NFTRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *Solar721NFTRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Solar721NFTRoleAdminChanged)
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
		it.Event = new(Solar721NFTRoleAdminChanged)
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
func (it *Solar721NFTRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Solar721NFTRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Solar721NFTRoleAdminChanged represents a RoleAdminChanged event raised by the Solar721NFT contract.
type Solar721NFTRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Solar721NFT *Solar721NFTFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*Solar721NFTRoleAdminChangedIterator, error) {

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

	logs, sub, err := _Solar721NFT.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &Solar721NFTRoleAdminChangedIterator{contract: _Solar721NFT.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Solar721NFT *Solar721NFTFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *Solar721NFTRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _Solar721NFT.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Solar721NFTRoleAdminChanged)
				if err := _Solar721NFT.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_Solar721NFT *Solar721NFTFilterer) ParseRoleAdminChanged(log types.Log) (*Solar721NFTRoleAdminChanged, error) {
	event := new(Solar721NFTRoleAdminChanged)
	if err := _Solar721NFT.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Solar721NFTRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Solar721NFT contract.
type Solar721NFTRoleGrantedIterator struct {
	Event *Solar721NFTRoleGranted // Event containing the contract specifics and raw log

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
func (it *Solar721NFTRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Solar721NFTRoleGranted)
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
		it.Event = new(Solar721NFTRoleGranted)
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
func (it *Solar721NFTRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Solar721NFTRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Solar721NFTRoleGranted represents a RoleGranted event raised by the Solar721NFT contract.
type Solar721NFTRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Solar721NFT *Solar721NFTFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*Solar721NFTRoleGrantedIterator, error) {

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

	logs, sub, err := _Solar721NFT.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &Solar721NFTRoleGrantedIterator{contract: _Solar721NFT.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Solar721NFT *Solar721NFTFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *Solar721NFTRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Solar721NFT.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Solar721NFTRoleGranted)
				if err := _Solar721NFT.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_Solar721NFT *Solar721NFTFilterer) ParseRoleGranted(log types.Log) (*Solar721NFTRoleGranted, error) {
	event := new(Solar721NFTRoleGranted)
	if err := _Solar721NFT.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Solar721NFTRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Solar721NFT contract.
type Solar721NFTRoleRevokedIterator struct {
	Event *Solar721NFTRoleRevoked // Event containing the contract specifics and raw log

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
func (it *Solar721NFTRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Solar721NFTRoleRevoked)
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
		it.Event = new(Solar721NFTRoleRevoked)
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
func (it *Solar721NFTRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Solar721NFTRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Solar721NFTRoleRevoked represents a RoleRevoked event raised by the Solar721NFT contract.
type Solar721NFTRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Solar721NFT *Solar721NFTFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*Solar721NFTRoleRevokedIterator, error) {

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

	logs, sub, err := _Solar721NFT.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &Solar721NFTRoleRevokedIterator{contract: _Solar721NFT.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Solar721NFT *Solar721NFTFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *Solar721NFTRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Solar721NFT.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Solar721NFTRoleRevoked)
				if err := _Solar721NFT.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_Solar721NFT *Solar721NFTFilterer) ParseRoleRevoked(log types.Log) (*Solar721NFTRoleRevoked, error) {
	event := new(Solar721NFTRoleRevoked)
	if err := _Solar721NFT.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Solar721NFTTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Solar721NFT contract.
type Solar721NFTTransferIterator struct {
	Event *Solar721NFTTransfer // Event containing the contract specifics and raw log

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
func (it *Solar721NFTTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Solar721NFTTransfer)
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
		it.Event = new(Solar721NFTTransfer)
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
func (it *Solar721NFTTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Solar721NFTTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Solar721NFTTransfer represents a Transfer event raised by the Solar721NFT contract.
type Solar721NFTTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Solar721NFT *Solar721NFTFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*Solar721NFTTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Solar721NFT.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &Solar721NFTTransferIterator{contract: _Solar721NFT.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Solar721NFT *Solar721NFTFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *Solar721NFTTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Solar721NFT.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Solar721NFTTransfer)
				if err := _Solar721NFT.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Solar721NFT *Solar721NFTFilterer) ParseTransfer(log types.Log) (*Solar721NFTTransfer, error) {
	event := new(Solar721NFTTransfer)
	if err := _Solar721NFT.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Solar721NFTTransferNFTIterator is returned from FilterTransferNFT and is used to iterate over the raw logs and unpacked data for TransferNFT events raised by the Solar721NFT contract.
type Solar721NFTTransferNFTIterator struct {
	Event *Solar721NFTTransferNFT // Event containing the contract specifics and raw log

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
func (it *Solar721NFTTransferNFTIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Solar721NFTTransferNFT)
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
		it.Event = new(Solar721NFTTransferNFT)
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
func (it *Solar721NFTTransferNFTIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Solar721NFTTransferNFTIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Solar721NFTTransferNFT represents a TransferNFT event raised by the Solar721NFT contract.
type Solar721NFTTransferNFT struct {
	Buyer   common.Address
	Seller  common.Address
	TokenId *big.Int
	Amount  *big.Int
	TradeId *big.Int
	OrderId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransferNFT is a free log retrieval operation binding the contract event 0x5f3324640c72e908deeb7ed6936cb3d6b3fe6fff62c68d6dbe7f61abb716e7f1.
//
// Solidity: event TransferNFT(address indexed _buyer, address indexed _seller, uint256 indexed _tokenId, uint256 _amount, uint256 _tradeId, uint256 _orderId)
func (_Solar721NFT *Solar721NFTFilterer) FilterTransferNFT(opts *bind.FilterOpts, _buyer []common.Address, _seller []common.Address, _tokenId []*big.Int) (*Solar721NFTTransferNFTIterator, error) {

	var _buyerRule []interface{}
	for _, _buyerItem := range _buyer {
		_buyerRule = append(_buyerRule, _buyerItem)
	}
	var _sellerRule []interface{}
	for _, _sellerItem := range _seller {
		_sellerRule = append(_sellerRule, _sellerItem)
	}
	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}

	logs, sub, err := _Solar721NFT.contract.FilterLogs(opts, "TransferNFT", _buyerRule, _sellerRule, _tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &Solar721NFTTransferNFTIterator{contract: _Solar721NFT.contract, event: "TransferNFT", logs: logs, sub: sub}, nil
}

// WatchTransferNFT is a free log subscription operation binding the contract event 0x5f3324640c72e908deeb7ed6936cb3d6b3fe6fff62c68d6dbe7f61abb716e7f1.
//
// Solidity: event TransferNFT(address indexed _buyer, address indexed _seller, uint256 indexed _tokenId, uint256 _amount, uint256 _tradeId, uint256 _orderId)
func (_Solar721NFT *Solar721NFTFilterer) WatchTransferNFT(opts *bind.WatchOpts, sink chan<- *Solar721NFTTransferNFT, _buyer []common.Address, _seller []common.Address, _tokenId []*big.Int) (event.Subscription, error) {

	var _buyerRule []interface{}
	for _, _buyerItem := range _buyer {
		_buyerRule = append(_buyerRule, _buyerItem)
	}
	var _sellerRule []interface{}
	for _, _sellerItem := range _seller {
		_sellerRule = append(_sellerRule, _sellerItem)
	}
	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}

	logs, sub, err := _Solar721NFT.contract.WatchLogs(opts, "TransferNFT", _buyerRule, _sellerRule, _tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Solar721NFTTransferNFT)
				if err := _Solar721NFT.contract.UnpackLog(event, "TransferNFT", log); err != nil {
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

// ParseTransferNFT is a log parse operation binding the contract event 0x5f3324640c72e908deeb7ed6936cb3d6b3fe6fff62c68d6dbe7f61abb716e7f1.
//
// Solidity: event TransferNFT(address indexed _buyer, address indexed _seller, uint256 indexed _tokenId, uint256 _amount, uint256 _tradeId, uint256 _orderId)
func (_Solar721NFT *Solar721NFTFilterer) ParseTransferNFT(log types.Log) (*Solar721NFTTransferNFT, error) {
	event := new(Solar721NFTTransferNFT)
	if err := _Solar721NFT.contract.UnpackLog(event, "TransferNFT", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Solar721NFTUpdateOrderIterator is returned from FilterUpdateOrder and is used to iterate over the raw logs and unpacked data for UpdateOrder events raised by the Solar721NFT contract.
type Solar721NFTUpdateOrderIterator struct {
	Event *Solar721NFTUpdateOrder // Event containing the contract specifics and raw log

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
func (it *Solar721NFTUpdateOrderIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Solar721NFTUpdateOrder)
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
		it.Event = new(Solar721NFTUpdateOrder)
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
func (it *Solar721NFTUpdateOrderIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Solar721NFTUpdateOrderIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Solar721NFTUpdateOrder represents a UpdateOrder event raised by the Solar721NFT contract.
type Solar721NFTUpdateOrder struct {
	Operator     common.Address
	TokenId      *big.Int
	OrderId      *big.Int
	SellingPrice *big.Int
	FeePercent   uint16
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterUpdateOrder is a free log retrieval operation binding the contract event 0x3919119cd19163ce3bba7a2b60dcc5add194c3217075e6e9a19640a8bfa512a9.
//
// Solidity: event UpdateOrder(address indexed _operator, uint256 indexed _tokenId, uint256 indexed _orderId, uint256 _sellingPrice, uint16 _feePercent)
func (_Solar721NFT *Solar721NFTFilterer) FilterUpdateOrder(opts *bind.FilterOpts, _operator []common.Address, _tokenId []*big.Int, _orderId []*big.Int) (*Solar721NFTUpdateOrderIterator, error) {

	var _operatorRule []interface{}
	for _, _operatorItem := range _operator {
		_operatorRule = append(_operatorRule, _operatorItem)
	}
	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}
	var _orderIdRule []interface{}
	for _, _orderIdItem := range _orderId {
		_orderIdRule = append(_orderIdRule, _orderIdItem)
	}

	logs, sub, err := _Solar721NFT.contract.FilterLogs(opts, "UpdateOrder", _operatorRule, _tokenIdRule, _orderIdRule)
	if err != nil {
		return nil, err
	}
	return &Solar721NFTUpdateOrderIterator{contract: _Solar721NFT.contract, event: "UpdateOrder", logs: logs, sub: sub}, nil
}

// WatchUpdateOrder is a free log subscription operation binding the contract event 0x3919119cd19163ce3bba7a2b60dcc5add194c3217075e6e9a19640a8bfa512a9.
//
// Solidity: event UpdateOrder(address indexed _operator, uint256 indexed _tokenId, uint256 indexed _orderId, uint256 _sellingPrice, uint16 _feePercent)
func (_Solar721NFT *Solar721NFTFilterer) WatchUpdateOrder(opts *bind.WatchOpts, sink chan<- *Solar721NFTUpdateOrder, _operator []common.Address, _tokenId []*big.Int, _orderId []*big.Int) (event.Subscription, error) {

	var _operatorRule []interface{}
	for _, _operatorItem := range _operator {
		_operatorRule = append(_operatorRule, _operatorItem)
	}
	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}
	var _orderIdRule []interface{}
	for _, _orderIdItem := range _orderId {
		_orderIdRule = append(_orderIdRule, _orderIdItem)
	}

	logs, sub, err := _Solar721NFT.contract.WatchLogs(opts, "UpdateOrder", _operatorRule, _tokenIdRule, _orderIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Solar721NFTUpdateOrder)
				if err := _Solar721NFT.contract.UnpackLog(event, "UpdateOrder", log); err != nil {
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

// ParseUpdateOrder is a log parse operation binding the contract event 0x3919119cd19163ce3bba7a2b60dcc5add194c3217075e6e9a19640a8bfa512a9.
//
// Solidity: event UpdateOrder(address indexed _operator, uint256 indexed _tokenId, uint256 indexed _orderId, uint256 _sellingPrice, uint16 _feePercent)
func (_Solar721NFT *Solar721NFTFilterer) ParseUpdateOrder(log types.Log) (*Solar721NFTUpdateOrder, error) {
	event := new(Solar721NFTUpdateOrder)
	if err := _Solar721NFT.contract.UnpackLog(event, "UpdateOrder", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
