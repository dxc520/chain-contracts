// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ArtExhibition

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

// ArtExhibitionMetaData contains all meta data concerning the ArtExhibition contract.
var ArtExhibitionMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"}],\"name\":\"BurnBatch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"BurnSingle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Gift\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"amount\",\"type\":\"uint256[]\"}],\"name\":\"GiftBatch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"contratAdress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"enumAuthorityUpgradeable.AuthorityRole\",\"name\":\"role\",\"type\":\"uint8\"}],\"name\":\"GrantRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"string[]\",\"name\":\"tokenUris\",\"type\":\"string[]\"}],\"name\":\"MintBatchItems\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"tokenURI\",\"type\":\"string\"}],\"name\":\"MintSingleItem\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"contratAdress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"enumAuthorityUpgradeable.AuthorityRole\",\"name\":\"role\",\"type\":\"uint8\"}],\"name\":\"RevokeRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"}],\"name\":\"TransferBatch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"TransferSingle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"URI\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"}],\"name\":\"balanceOfBatch\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"}],\"name\":\"burnBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantAdminRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantGiftRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantManagerRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantMintRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasAdminRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasGiftRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasManagerRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasMintRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeAdminRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeGiftRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeManagerRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeMintRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeBatchTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"uri\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"walletProxy\",\"type\":\"address\"}],\"name\":\"setWalletProxyInit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"tokenURI\",\"type\":\"string\"}],\"name\":\"mintSingleItem\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"string[]\",\"name\":\"tokenUris\",\"type\":\"string[]\"}],\"name\":\"mintBatchItems\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"gift\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"giftBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ArtExhibitionABI is the input ABI used to generate the binding from.
// Deprecated: Use ArtExhibitionMetaData.ABI instead.
var ArtExhibitionABI = ArtExhibitionMetaData.ABI

// ArtExhibition is an auto generated Go binding around an Ethereum contract.
type ArtExhibition struct {
	ArtExhibitionCaller     // Read-only binding to the contract
	ArtExhibitionTransactor // Write-only binding to the contract
	ArtExhibitionFilterer   // Log filterer for contract events
}

// ArtExhibitionCaller is an auto generated read-only Go binding around an Ethereum contract.
type ArtExhibitionCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArtExhibitionTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ArtExhibitionTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArtExhibitionFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ArtExhibitionFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArtExhibitionSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ArtExhibitionSession struct {
	Contract     *ArtExhibition    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ArtExhibitionCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ArtExhibitionCallerSession struct {
	Contract *ArtExhibitionCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// ArtExhibitionTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ArtExhibitionTransactorSession struct {
	Contract     *ArtExhibitionTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// ArtExhibitionRaw is an auto generated low-level Go binding around an Ethereum contract.
type ArtExhibitionRaw struct {
	Contract *ArtExhibition // Generic contract binding to access the raw methods on
}

// ArtExhibitionCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ArtExhibitionCallerRaw struct {
	Contract *ArtExhibitionCaller // Generic read-only contract binding to access the raw methods on
}

// ArtExhibitionTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ArtExhibitionTransactorRaw struct {
	Contract *ArtExhibitionTransactor // Generic write-only contract binding to access the raw methods on
}

// NewArtExhibition creates a new instance of ArtExhibition, bound to a specific deployed contract.
func NewArtExhibition(address common.Address, backend bind.ContractBackend) (*ArtExhibition, error) {
	contract, err := bindArtExhibition(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ArtExhibition{ArtExhibitionCaller: ArtExhibitionCaller{contract: contract}, ArtExhibitionTransactor: ArtExhibitionTransactor{contract: contract}, ArtExhibitionFilterer: ArtExhibitionFilterer{contract: contract}}, nil
}

// NewArtExhibitionCaller creates a new read-only instance of ArtExhibition, bound to a specific deployed contract.
func NewArtExhibitionCaller(address common.Address, caller bind.ContractCaller) (*ArtExhibitionCaller, error) {
	contract, err := bindArtExhibition(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ArtExhibitionCaller{contract: contract}, nil
}

// NewArtExhibitionTransactor creates a new write-only instance of ArtExhibition, bound to a specific deployed contract.
func NewArtExhibitionTransactor(address common.Address, transactor bind.ContractTransactor) (*ArtExhibitionTransactor, error) {
	contract, err := bindArtExhibition(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ArtExhibitionTransactor{contract: contract}, nil
}

// NewArtExhibitionFilterer creates a new log filterer instance of ArtExhibition, bound to a specific deployed contract.
func NewArtExhibitionFilterer(address common.Address, filterer bind.ContractFilterer) (*ArtExhibitionFilterer, error) {
	contract, err := bindArtExhibition(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ArtExhibitionFilterer{contract: contract}, nil
}

// bindArtExhibition binds a generic wrapper to an already deployed contract.
func bindArtExhibition(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ArtExhibitionABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ArtExhibition *ArtExhibitionRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ArtExhibition.Contract.ArtExhibitionCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ArtExhibition *ArtExhibitionRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ArtExhibition.Contract.ArtExhibitionTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ArtExhibition *ArtExhibitionRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ArtExhibition.Contract.ArtExhibitionTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ArtExhibition *ArtExhibitionCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ArtExhibition.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ArtExhibition *ArtExhibitionTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ArtExhibition.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ArtExhibition *ArtExhibitionTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ArtExhibition.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ArtExhibition *ArtExhibitionCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ArtExhibition.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ArtExhibition *ArtExhibitionSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _ArtExhibition.Contract.DEFAULTADMINROLE(&_ArtExhibition.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ArtExhibition *ArtExhibitionCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _ArtExhibition.Contract.DEFAULTADMINROLE(&_ArtExhibition.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_ArtExhibition *ArtExhibitionCaller) BalanceOf(opts *bind.CallOpts, account common.Address, id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ArtExhibition.contract.Call(opts, &out, "balanceOf", account, id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_ArtExhibition *ArtExhibitionSession) BalanceOf(account common.Address, id *big.Int) (*big.Int, error) {
	return _ArtExhibition.Contract.BalanceOf(&_ArtExhibition.CallOpts, account, id)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_ArtExhibition *ArtExhibitionCallerSession) BalanceOf(account common.Address, id *big.Int) (*big.Int, error) {
	return _ArtExhibition.Contract.BalanceOf(&_ArtExhibition.CallOpts, account, id)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_ArtExhibition *ArtExhibitionCaller) BalanceOfBatch(opts *bind.CallOpts, accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _ArtExhibition.contract.Call(opts, &out, "balanceOfBatch", accounts, ids)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_ArtExhibition *ArtExhibitionSession) BalanceOfBatch(accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	return _ArtExhibition.Contract.BalanceOfBatch(&_ArtExhibition.CallOpts, accounts, ids)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_ArtExhibition *ArtExhibitionCallerSession) BalanceOfBatch(accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	return _ArtExhibition.Contract.BalanceOfBatch(&_ArtExhibition.CallOpts, accounts, ids)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ArtExhibition *ArtExhibitionCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _ArtExhibition.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ArtExhibition *ArtExhibitionSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _ArtExhibition.Contract.GetRoleAdmin(&_ArtExhibition.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ArtExhibition *ArtExhibitionCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _ArtExhibition.Contract.GetRoleAdmin(&_ArtExhibition.CallOpts, role)
}

// HasAdminRole is a free data retrieval call binding the contract method 0xc395fcb3.
//
// Solidity: function hasAdminRole(address account) view returns(bool)
func (_ArtExhibition *ArtExhibitionCaller) HasAdminRole(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _ArtExhibition.contract.Call(opts, &out, "hasAdminRole", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasAdminRole is a free data retrieval call binding the contract method 0xc395fcb3.
//
// Solidity: function hasAdminRole(address account) view returns(bool)
func (_ArtExhibition *ArtExhibitionSession) HasAdminRole(account common.Address) (bool, error) {
	return _ArtExhibition.Contract.HasAdminRole(&_ArtExhibition.CallOpts, account)
}

// HasAdminRole is a free data retrieval call binding the contract method 0xc395fcb3.
//
// Solidity: function hasAdminRole(address account) view returns(bool)
func (_ArtExhibition *ArtExhibitionCallerSession) HasAdminRole(account common.Address) (bool, error) {
	return _ArtExhibition.Contract.HasAdminRole(&_ArtExhibition.CallOpts, account)
}

// HasGiftRole is a free data retrieval call binding the contract method 0x629962dc.
//
// Solidity: function hasGiftRole(address account) view returns(bool)
func (_ArtExhibition *ArtExhibitionCaller) HasGiftRole(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _ArtExhibition.contract.Call(opts, &out, "hasGiftRole", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasGiftRole is a free data retrieval call binding the contract method 0x629962dc.
//
// Solidity: function hasGiftRole(address account) view returns(bool)
func (_ArtExhibition *ArtExhibitionSession) HasGiftRole(account common.Address) (bool, error) {
	return _ArtExhibition.Contract.HasGiftRole(&_ArtExhibition.CallOpts, account)
}

// HasGiftRole is a free data retrieval call binding the contract method 0x629962dc.
//
// Solidity: function hasGiftRole(address account) view returns(bool)
func (_ArtExhibition *ArtExhibitionCallerSession) HasGiftRole(account common.Address) (bool, error) {
	return _ArtExhibition.Contract.HasGiftRole(&_ArtExhibition.CallOpts, account)
}

// HasManagerRole is a free data retrieval call binding the contract method 0x5026c826.
//
// Solidity: function hasManagerRole(address account) view returns(bool)
func (_ArtExhibition *ArtExhibitionCaller) HasManagerRole(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _ArtExhibition.contract.Call(opts, &out, "hasManagerRole", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasManagerRole is a free data retrieval call binding the contract method 0x5026c826.
//
// Solidity: function hasManagerRole(address account) view returns(bool)
func (_ArtExhibition *ArtExhibitionSession) HasManagerRole(account common.Address) (bool, error) {
	return _ArtExhibition.Contract.HasManagerRole(&_ArtExhibition.CallOpts, account)
}

// HasManagerRole is a free data retrieval call binding the contract method 0x5026c826.
//
// Solidity: function hasManagerRole(address account) view returns(bool)
func (_ArtExhibition *ArtExhibitionCallerSession) HasManagerRole(account common.Address) (bool, error) {
	return _ArtExhibition.Contract.HasManagerRole(&_ArtExhibition.CallOpts, account)
}

// HasMintRole is a free data retrieval call binding the contract method 0xc4e58e54.
//
// Solidity: function hasMintRole(address account) view returns(bool)
func (_ArtExhibition *ArtExhibitionCaller) HasMintRole(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _ArtExhibition.contract.Call(opts, &out, "hasMintRole", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasMintRole is a free data retrieval call binding the contract method 0xc4e58e54.
//
// Solidity: function hasMintRole(address account) view returns(bool)
func (_ArtExhibition *ArtExhibitionSession) HasMintRole(account common.Address) (bool, error) {
	return _ArtExhibition.Contract.HasMintRole(&_ArtExhibition.CallOpts, account)
}

// HasMintRole is a free data retrieval call binding the contract method 0xc4e58e54.
//
// Solidity: function hasMintRole(address account) view returns(bool)
func (_ArtExhibition *ArtExhibitionCallerSession) HasMintRole(account common.Address) (bool, error) {
	return _ArtExhibition.Contract.HasMintRole(&_ArtExhibition.CallOpts, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ArtExhibition *ArtExhibitionCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _ArtExhibition.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ArtExhibition *ArtExhibitionSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _ArtExhibition.Contract.HasRole(&_ArtExhibition.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ArtExhibition *ArtExhibitionCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _ArtExhibition.Contract.HasRole(&_ArtExhibition.CallOpts, role, account)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_ArtExhibition *ArtExhibitionCaller) IsApprovedForAll(opts *bind.CallOpts, account common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _ArtExhibition.contract.Call(opts, &out, "isApprovedForAll", account, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_ArtExhibition *ArtExhibitionSession) IsApprovedForAll(account common.Address, operator common.Address) (bool, error) {
	return _ArtExhibition.Contract.IsApprovedForAll(&_ArtExhibition.CallOpts, account, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_ArtExhibition *ArtExhibitionCallerSession) IsApprovedForAll(account common.Address, operator common.Address) (bool, error) {
	return _ArtExhibition.Contract.IsApprovedForAll(&_ArtExhibition.CallOpts, account, operator)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ArtExhibition *ArtExhibitionCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ArtExhibition.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ArtExhibition *ArtExhibitionSession) Name() (string, error) {
	return _ArtExhibition.Contract.Name(&_ArtExhibition.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ArtExhibition *ArtExhibitionCallerSession) Name() (string, error) {
	return _ArtExhibition.Contract.Name(&_ArtExhibition.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ArtExhibition *ArtExhibitionCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ArtExhibition.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ArtExhibition *ArtExhibitionSession) Owner() (common.Address, error) {
	return _ArtExhibition.Contract.Owner(&_ArtExhibition.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ArtExhibition *ArtExhibitionCallerSession) Owner() (common.Address, error) {
	return _ArtExhibition.Contract.Owner(&_ArtExhibition.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ArtExhibition *ArtExhibitionCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _ArtExhibition.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ArtExhibition *ArtExhibitionSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ArtExhibition.Contract.SupportsInterface(&_ArtExhibition.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ArtExhibition *ArtExhibitionCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ArtExhibition.Contract.SupportsInterface(&_ArtExhibition.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ArtExhibition *ArtExhibitionCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ArtExhibition.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ArtExhibition *ArtExhibitionSession) Symbol() (string, error) {
	return _ArtExhibition.Contract.Symbol(&_ArtExhibition.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ArtExhibition *ArtExhibitionCallerSession) Symbol() (string, error) {
	return _ArtExhibition.Contract.Symbol(&_ArtExhibition.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_ArtExhibition *ArtExhibitionCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _ArtExhibition.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_ArtExhibition *ArtExhibitionSession) TokenURI(tokenId *big.Int) (string, error) {
	return _ArtExhibition.Contract.TokenURI(&_ArtExhibition.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_ArtExhibition *ArtExhibitionCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _ArtExhibition.Contract.TokenURI(&_ArtExhibition.CallOpts, tokenId)
}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 ) view returns(string)
func (_ArtExhibition *ArtExhibitionCaller) Uri(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _ArtExhibition.contract.Call(opts, &out, "uri", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 ) view returns(string)
func (_ArtExhibition *ArtExhibitionSession) Uri(arg0 *big.Int) (string, error) {
	return _ArtExhibition.Contract.Uri(&_ArtExhibition.CallOpts, arg0)
}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 ) view returns(string)
func (_ArtExhibition *ArtExhibitionCallerSession) Uri(arg0 *big.Int) (string, error) {
	return _ArtExhibition.Contract.Uri(&_ArtExhibition.CallOpts, arg0)
}

// Burn is a paid mutator transaction binding the contract method 0xf5298aca.
//
// Solidity: function burn(address account, uint256 id, uint256 value) returns()
func (_ArtExhibition *ArtExhibitionTransactor) Burn(opts *bind.TransactOpts, account common.Address, id *big.Int, value *big.Int) (*types.Transaction, error) {
	return _ArtExhibition.contract.Transact(opts, "burn", account, id, value)
}

// Burn is a paid mutator transaction binding the contract method 0xf5298aca.
//
// Solidity: function burn(address account, uint256 id, uint256 value) returns()
func (_ArtExhibition *ArtExhibitionSession) Burn(account common.Address, id *big.Int, value *big.Int) (*types.Transaction, error) {
	return _ArtExhibition.Contract.Burn(&_ArtExhibition.TransactOpts, account, id, value)
}

// Burn is a paid mutator transaction binding the contract method 0xf5298aca.
//
// Solidity: function burn(address account, uint256 id, uint256 value) returns()
func (_ArtExhibition *ArtExhibitionTransactorSession) Burn(account common.Address, id *big.Int, value *big.Int) (*types.Transaction, error) {
	return _ArtExhibition.Contract.Burn(&_ArtExhibition.TransactOpts, account, id, value)
}

// BurnBatch is a paid mutator transaction binding the contract method 0x6b20c454.
//
// Solidity: function burnBatch(address account, uint256[] ids, uint256[] values) returns()
func (_ArtExhibition *ArtExhibitionTransactor) BurnBatch(opts *bind.TransactOpts, account common.Address, ids []*big.Int, values []*big.Int) (*types.Transaction, error) {
	return _ArtExhibition.contract.Transact(opts, "burnBatch", account, ids, values)
}

// BurnBatch is a paid mutator transaction binding the contract method 0x6b20c454.
//
// Solidity: function burnBatch(address account, uint256[] ids, uint256[] values) returns()
func (_ArtExhibition *ArtExhibitionSession) BurnBatch(account common.Address, ids []*big.Int, values []*big.Int) (*types.Transaction, error) {
	return _ArtExhibition.Contract.BurnBatch(&_ArtExhibition.TransactOpts, account, ids, values)
}

// BurnBatch is a paid mutator transaction binding the contract method 0x6b20c454.
//
// Solidity: function burnBatch(address account, uint256[] ids, uint256[] values) returns()
func (_ArtExhibition *ArtExhibitionTransactorSession) BurnBatch(account common.Address, ids []*big.Int, values []*big.Int) (*types.Transaction, error) {
	return _ArtExhibition.Contract.BurnBatch(&_ArtExhibition.TransactOpts, account, ids, values)
}

// Gift is a paid mutator transaction binding the contract method 0x388a1cf2.
//
// Solidity: function gift(address from, address to, uint256 tokenId, uint256 amount) returns()
func (_ArtExhibition *ArtExhibitionTransactor) Gift(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _ArtExhibition.contract.Transact(opts, "gift", from, to, tokenId, amount)
}

// Gift is a paid mutator transaction binding the contract method 0x388a1cf2.
//
// Solidity: function gift(address from, address to, uint256 tokenId, uint256 amount) returns()
func (_ArtExhibition *ArtExhibitionSession) Gift(from common.Address, to common.Address, tokenId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _ArtExhibition.Contract.Gift(&_ArtExhibition.TransactOpts, from, to, tokenId, amount)
}

// Gift is a paid mutator transaction binding the contract method 0x388a1cf2.
//
// Solidity: function gift(address from, address to, uint256 tokenId, uint256 amount) returns()
func (_ArtExhibition *ArtExhibitionTransactorSession) Gift(from common.Address, to common.Address, tokenId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _ArtExhibition.Contract.Gift(&_ArtExhibition.TransactOpts, from, to, tokenId, amount)
}

// GiftBatch is a paid mutator transaction binding the contract method 0x21bb0b08.
//
// Solidity: function giftBatch(address from, address to, uint256[] tokenIds, uint256[] amounts) returns()
func (_ArtExhibition *ArtExhibitionTransactor) GiftBatch(opts *bind.TransactOpts, from common.Address, to common.Address, tokenIds []*big.Int, amounts []*big.Int) (*types.Transaction, error) {
	return _ArtExhibition.contract.Transact(opts, "giftBatch", from, to, tokenIds, amounts)
}

// GiftBatch is a paid mutator transaction binding the contract method 0x21bb0b08.
//
// Solidity: function giftBatch(address from, address to, uint256[] tokenIds, uint256[] amounts) returns()
func (_ArtExhibition *ArtExhibitionSession) GiftBatch(from common.Address, to common.Address, tokenIds []*big.Int, amounts []*big.Int) (*types.Transaction, error) {
	return _ArtExhibition.Contract.GiftBatch(&_ArtExhibition.TransactOpts, from, to, tokenIds, amounts)
}

// GiftBatch is a paid mutator transaction binding the contract method 0x21bb0b08.
//
// Solidity: function giftBatch(address from, address to, uint256[] tokenIds, uint256[] amounts) returns()
func (_ArtExhibition *ArtExhibitionTransactorSession) GiftBatch(from common.Address, to common.Address, tokenIds []*big.Int, amounts []*big.Int) (*types.Transaction, error) {
	return _ArtExhibition.Contract.GiftBatch(&_ArtExhibition.TransactOpts, from, to, tokenIds, amounts)
}

// GrantAdminRole is a paid mutator transaction binding the contract method 0xc634b78e.
//
// Solidity: function grantAdminRole(address account) returns()
func (_ArtExhibition *ArtExhibitionTransactor) GrantAdminRole(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _ArtExhibition.contract.Transact(opts, "grantAdminRole", account)
}

// GrantAdminRole is a paid mutator transaction binding the contract method 0xc634b78e.
//
// Solidity: function grantAdminRole(address account) returns()
func (_ArtExhibition *ArtExhibitionSession) GrantAdminRole(account common.Address) (*types.Transaction, error) {
	return _ArtExhibition.Contract.GrantAdminRole(&_ArtExhibition.TransactOpts, account)
}

// GrantAdminRole is a paid mutator transaction binding the contract method 0xc634b78e.
//
// Solidity: function grantAdminRole(address account) returns()
func (_ArtExhibition *ArtExhibitionTransactorSession) GrantAdminRole(account common.Address) (*types.Transaction, error) {
	return _ArtExhibition.Contract.GrantAdminRole(&_ArtExhibition.TransactOpts, account)
}

// GrantGiftRole is a paid mutator transaction binding the contract method 0x0bd9c905.
//
// Solidity: function grantGiftRole(address account) returns()
func (_ArtExhibition *ArtExhibitionTransactor) GrantGiftRole(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _ArtExhibition.contract.Transact(opts, "grantGiftRole", account)
}

// GrantGiftRole is a paid mutator transaction binding the contract method 0x0bd9c905.
//
// Solidity: function grantGiftRole(address account) returns()
func (_ArtExhibition *ArtExhibitionSession) GrantGiftRole(account common.Address) (*types.Transaction, error) {
	return _ArtExhibition.Contract.GrantGiftRole(&_ArtExhibition.TransactOpts, account)
}

// GrantGiftRole is a paid mutator transaction binding the contract method 0x0bd9c905.
//
// Solidity: function grantGiftRole(address account) returns()
func (_ArtExhibition *ArtExhibitionTransactorSession) GrantGiftRole(account common.Address) (*types.Transaction, error) {
	return _ArtExhibition.Contract.GrantGiftRole(&_ArtExhibition.TransactOpts, account)
}

// GrantManagerRole is a paid mutator transaction binding the contract method 0x26e885e3.
//
// Solidity: function grantManagerRole(address account) returns()
func (_ArtExhibition *ArtExhibitionTransactor) GrantManagerRole(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _ArtExhibition.contract.Transact(opts, "grantManagerRole", account)
}

// GrantManagerRole is a paid mutator transaction binding the contract method 0x26e885e3.
//
// Solidity: function grantManagerRole(address account) returns()
func (_ArtExhibition *ArtExhibitionSession) GrantManagerRole(account common.Address) (*types.Transaction, error) {
	return _ArtExhibition.Contract.GrantManagerRole(&_ArtExhibition.TransactOpts, account)
}

// GrantManagerRole is a paid mutator transaction binding the contract method 0x26e885e3.
//
// Solidity: function grantManagerRole(address account) returns()
func (_ArtExhibition *ArtExhibitionTransactorSession) GrantManagerRole(account common.Address) (*types.Transaction, error) {
	return _ArtExhibition.Contract.GrantManagerRole(&_ArtExhibition.TransactOpts, account)
}

// GrantMintRole is a paid mutator transaction binding the contract method 0xc2e3273d.
//
// Solidity: function grantMintRole(address account) returns()
func (_ArtExhibition *ArtExhibitionTransactor) GrantMintRole(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _ArtExhibition.contract.Transact(opts, "grantMintRole", account)
}

// GrantMintRole is a paid mutator transaction binding the contract method 0xc2e3273d.
//
// Solidity: function grantMintRole(address account) returns()
func (_ArtExhibition *ArtExhibitionSession) GrantMintRole(account common.Address) (*types.Transaction, error) {
	return _ArtExhibition.Contract.GrantMintRole(&_ArtExhibition.TransactOpts, account)
}

// GrantMintRole is a paid mutator transaction binding the contract method 0xc2e3273d.
//
// Solidity: function grantMintRole(address account) returns()
func (_ArtExhibition *ArtExhibitionTransactorSession) GrantMintRole(account common.Address) (*types.Transaction, error) {
	return _ArtExhibition.Contract.GrantMintRole(&_ArtExhibition.TransactOpts, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ArtExhibition *ArtExhibitionTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ArtExhibition.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ArtExhibition *ArtExhibitionSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ArtExhibition.Contract.GrantRole(&_ArtExhibition.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ArtExhibition *ArtExhibitionTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ArtExhibition.Contract.GrantRole(&_ArtExhibition.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0xa6487c53.
//
// Solidity: function initialize(string name, string symbol, string uri) returns()
func (_ArtExhibition *ArtExhibitionTransactor) Initialize(opts *bind.TransactOpts, name string, symbol string, uri string) (*types.Transaction, error) {
	return _ArtExhibition.contract.Transact(opts, "initialize", name, symbol, uri)
}

// Initialize is a paid mutator transaction binding the contract method 0xa6487c53.
//
// Solidity: function initialize(string name, string symbol, string uri) returns()
func (_ArtExhibition *ArtExhibitionSession) Initialize(name string, symbol string, uri string) (*types.Transaction, error) {
	return _ArtExhibition.Contract.Initialize(&_ArtExhibition.TransactOpts, name, symbol, uri)
}

// Initialize is a paid mutator transaction binding the contract method 0xa6487c53.
//
// Solidity: function initialize(string name, string symbol, string uri) returns()
func (_ArtExhibition *ArtExhibitionTransactorSession) Initialize(name string, symbol string, uri string) (*types.Transaction, error) {
	return _ArtExhibition.Contract.Initialize(&_ArtExhibition.TransactOpts, name, symbol, uri)
}

// MintBatchItems is a paid mutator transaction binding the contract method 0x239079bb.
//
// Solidity: function mintBatchItems(address to, uint256[] tokenIds, uint256[] amounts, string[] tokenUris) returns()
func (_ArtExhibition *ArtExhibitionTransactor) MintBatchItems(opts *bind.TransactOpts, to common.Address, tokenIds []*big.Int, amounts []*big.Int, tokenUris []string) (*types.Transaction, error) {
	return _ArtExhibition.contract.Transact(opts, "mintBatchItems", to, tokenIds, amounts, tokenUris)
}

// MintBatchItems is a paid mutator transaction binding the contract method 0x239079bb.
//
// Solidity: function mintBatchItems(address to, uint256[] tokenIds, uint256[] amounts, string[] tokenUris) returns()
func (_ArtExhibition *ArtExhibitionSession) MintBatchItems(to common.Address, tokenIds []*big.Int, amounts []*big.Int, tokenUris []string) (*types.Transaction, error) {
	return _ArtExhibition.Contract.MintBatchItems(&_ArtExhibition.TransactOpts, to, tokenIds, amounts, tokenUris)
}

// MintBatchItems is a paid mutator transaction binding the contract method 0x239079bb.
//
// Solidity: function mintBatchItems(address to, uint256[] tokenIds, uint256[] amounts, string[] tokenUris) returns()
func (_ArtExhibition *ArtExhibitionTransactorSession) MintBatchItems(to common.Address, tokenIds []*big.Int, amounts []*big.Int, tokenUris []string) (*types.Transaction, error) {
	return _ArtExhibition.Contract.MintBatchItems(&_ArtExhibition.TransactOpts, to, tokenIds, amounts, tokenUris)
}

// MintSingleItem is a paid mutator transaction binding the contract method 0x6f492d51.
//
// Solidity: function mintSingleItem(address to, uint256 tokenId, uint256 amount, string tokenURI) returns()
func (_ArtExhibition *ArtExhibitionTransactor) MintSingleItem(opts *bind.TransactOpts, to common.Address, tokenId *big.Int, amount *big.Int, tokenURI string) (*types.Transaction, error) {
	return _ArtExhibition.contract.Transact(opts, "mintSingleItem", to, tokenId, amount, tokenURI)
}

// MintSingleItem is a paid mutator transaction binding the contract method 0x6f492d51.
//
// Solidity: function mintSingleItem(address to, uint256 tokenId, uint256 amount, string tokenURI) returns()
func (_ArtExhibition *ArtExhibitionSession) MintSingleItem(to common.Address, tokenId *big.Int, amount *big.Int, tokenURI string) (*types.Transaction, error) {
	return _ArtExhibition.Contract.MintSingleItem(&_ArtExhibition.TransactOpts, to, tokenId, amount, tokenURI)
}

// MintSingleItem is a paid mutator transaction binding the contract method 0x6f492d51.
//
// Solidity: function mintSingleItem(address to, uint256 tokenId, uint256 amount, string tokenURI) returns()
func (_ArtExhibition *ArtExhibitionTransactorSession) MintSingleItem(to common.Address, tokenId *big.Int, amount *big.Int, tokenURI string) (*types.Transaction, error) {
	return _ArtExhibition.Contract.MintSingleItem(&_ArtExhibition.TransactOpts, to, tokenId, amount, tokenURI)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ArtExhibition *ArtExhibitionTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ArtExhibition.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ArtExhibition *ArtExhibitionSession) RenounceOwnership() (*types.Transaction, error) {
	return _ArtExhibition.Contract.RenounceOwnership(&_ArtExhibition.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ArtExhibition *ArtExhibitionTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ArtExhibition.Contract.RenounceOwnership(&_ArtExhibition.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_ArtExhibition *ArtExhibitionTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ArtExhibition.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_ArtExhibition *ArtExhibitionSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ArtExhibition.Contract.RenounceRole(&_ArtExhibition.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_ArtExhibition *ArtExhibitionTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ArtExhibition.Contract.RenounceRole(&_ArtExhibition.TransactOpts, role, account)
}

// RevokeAdminRole is a paid mutator transaction binding the contract method 0x9a19c7b0.
//
// Solidity: function revokeAdminRole(address account) returns()
func (_ArtExhibition *ArtExhibitionTransactor) RevokeAdminRole(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _ArtExhibition.contract.Transact(opts, "revokeAdminRole", account)
}

// RevokeAdminRole is a paid mutator transaction binding the contract method 0x9a19c7b0.
//
// Solidity: function revokeAdminRole(address account) returns()
func (_ArtExhibition *ArtExhibitionSession) RevokeAdminRole(account common.Address) (*types.Transaction, error) {
	return _ArtExhibition.Contract.RevokeAdminRole(&_ArtExhibition.TransactOpts, account)
}

// RevokeAdminRole is a paid mutator transaction binding the contract method 0x9a19c7b0.
//
// Solidity: function revokeAdminRole(address account) returns()
func (_ArtExhibition *ArtExhibitionTransactorSession) RevokeAdminRole(account common.Address) (*types.Transaction, error) {
	return _ArtExhibition.Contract.RevokeAdminRole(&_ArtExhibition.TransactOpts, account)
}

// RevokeGiftRole is a paid mutator transaction binding the contract method 0x543309a0.
//
// Solidity: function revokeGiftRole(address account) returns()
func (_ArtExhibition *ArtExhibitionTransactor) RevokeGiftRole(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _ArtExhibition.contract.Transact(opts, "revokeGiftRole", account)
}

// RevokeGiftRole is a paid mutator transaction binding the contract method 0x543309a0.
//
// Solidity: function revokeGiftRole(address account) returns()
func (_ArtExhibition *ArtExhibitionSession) RevokeGiftRole(account common.Address) (*types.Transaction, error) {
	return _ArtExhibition.Contract.RevokeGiftRole(&_ArtExhibition.TransactOpts, account)
}

// RevokeGiftRole is a paid mutator transaction binding the contract method 0x543309a0.
//
// Solidity: function revokeGiftRole(address account) returns()
func (_ArtExhibition *ArtExhibitionTransactorSession) RevokeGiftRole(account common.Address) (*types.Transaction, error) {
	return _ArtExhibition.Contract.RevokeGiftRole(&_ArtExhibition.TransactOpts, account)
}

// RevokeManagerRole is a paid mutator transaction binding the contract method 0xbe4dc94f.
//
// Solidity: function revokeManagerRole(address account) returns()
func (_ArtExhibition *ArtExhibitionTransactor) RevokeManagerRole(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _ArtExhibition.contract.Transact(opts, "revokeManagerRole", account)
}

// RevokeManagerRole is a paid mutator transaction binding the contract method 0xbe4dc94f.
//
// Solidity: function revokeManagerRole(address account) returns()
func (_ArtExhibition *ArtExhibitionSession) RevokeManagerRole(account common.Address) (*types.Transaction, error) {
	return _ArtExhibition.Contract.RevokeManagerRole(&_ArtExhibition.TransactOpts, account)
}

// RevokeManagerRole is a paid mutator transaction binding the contract method 0xbe4dc94f.
//
// Solidity: function revokeManagerRole(address account) returns()
func (_ArtExhibition *ArtExhibitionTransactorSession) RevokeManagerRole(account common.Address) (*types.Transaction, error) {
	return _ArtExhibition.Contract.RevokeManagerRole(&_ArtExhibition.TransactOpts, account)
}

// RevokeMintRole is a paid mutator transaction binding the contract method 0xf81094f3.
//
// Solidity: function revokeMintRole(address account) returns()
func (_ArtExhibition *ArtExhibitionTransactor) RevokeMintRole(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _ArtExhibition.contract.Transact(opts, "revokeMintRole", account)
}

// RevokeMintRole is a paid mutator transaction binding the contract method 0xf81094f3.
//
// Solidity: function revokeMintRole(address account) returns()
func (_ArtExhibition *ArtExhibitionSession) RevokeMintRole(account common.Address) (*types.Transaction, error) {
	return _ArtExhibition.Contract.RevokeMintRole(&_ArtExhibition.TransactOpts, account)
}

// RevokeMintRole is a paid mutator transaction binding the contract method 0xf81094f3.
//
// Solidity: function revokeMintRole(address account) returns()
func (_ArtExhibition *ArtExhibitionTransactorSession) RevokeMintRole(account common.Address) (*types.Transaction, error) {
	return _ArtExhibition.Contract.RevokeMintRole(&_ArtExhibition.TransactOpts, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ArtExhibition *ArtExhibitionTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ArtExhibition.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ArtExhibition *ArtExhibitionSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ArtExhibition.Contract.RevokeRole(&_ArtExhibition.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ArtExhibition *ArtExhibitionTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ArtExhibition.Contract.RevokeRole(&_ArtExhibition.TransactOpts, role, account)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_ArtExhibition *ArtExhibitionTransactor) SafeBatchTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _ArtExhibition.contract.Transact(opts, "safeBatchTransferFrom", from, to, ids, amounts, data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_ArtExhibition *ArtExhibitionSession) SafeBatchTransferFrom(from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _ArtExhibition.Contract.SafeBatchTransferFrom(&_ArtExhibition.TransactOpts, from, to, ids, amounts, data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_ArtExhibition *ArtExhibitionTransactorSession) SafeBatchTransferFrom(from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _ArtExhibition.Contract.SafeBatchTransferFrom(&_ArtExhibition.TransactOpts, from, to, ids, amounts, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_ArtExhibition *ArtExhibitionTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _ArtExhibition.contract.Transact(opts, "safeTransferFrom", from, to, id, amount, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_ArtExhibition *ArtExhibitionSession) SafeTransferFrom(from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _ArtExhibition.Contract.SafeTransferFrom(&_ArtExhibition.TransactOpts, from, to, id, amount, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_ArtExhibition *ArtExhibitionTransactorSession) SafeTransferFrom(from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _ArtExhibition.Contract.SafeTransferFrom(&_ArtExhibition.TransactOpts, from, to, id, amount, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_ArtExhibition *ArtExhibitionTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _ArtExhibition.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_ArtExhibition *ArtExhibitionSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _ArtExhibition.Contract.SetApprovalForAll(&_ArtExhibition.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_ArtExhibition *ArtExhibitionTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _ArtExhibition.Contract.SetApprovalForAll(&_ArtExhibition.TransactOpts, operator, approved)
}

// SetWalletProxyInit is a paid mutator transaction binding the contract method 0x0b8c9321.
//
// Solidity: function setWalletProxyInit(address walletProxy) returns()
func (_ArtExhibition *ArtExhibitionTransactor) SetWalletProxyInit(opts *bind.TransactOpts, walletProxy common.Address) (*types.Transaction, error) {
	return _ArtExhibition.contract.Transact(opts, "setWalletProxyInit", walletProxy)
}

// SetWalletProxyInit is a paid mutator transaction binding the contract method 0x0b8c9321.
//
// Solidity: function setWalletProxyInit(address walletProxy) returns()
func (_ArtExhibition *ArtExhibitionSession) SetWalletProxyInit(walletProxy common.Address) (*types.Transaction, error) {
	return _ArtExhibition.Contract.SetWalletProxyInit(&_ArtExhibition.TransactOpts, walletProxy)
}

// SetWalletProxyInit is a paid mutator transaction binding the contract method 0x0b8c9321.
//
// Solidity: function setWalletProxyInit(address walletProxy) returns()
func (_ArtExhibition *ArtExhibitionTransactorSession) SetWalletProxyInit(walletProxy common.Address) (*types.Transaction, error) {
	return _ArtExhibition.Contract.SetWalletProxyInit(&_ArtExhibition.TransactOpts, walletProxy)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ArtExhibition *ArtExhibitionTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ArtExhibition.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ArtExhibition *ArtExhibitionSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ArtExhibition.Contract.TransferOwnership(&_ArtExhibition.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ArtExhibition *ArtExhibitionTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ArtExhibition.Contract.TransferOwnership(&_ArtExhibition.TransactOpts, newOwner)
}

// ArtExhibitionApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the ArtExhibition contract.
type ArtExhibitionApprovalForAllIterator struct {
	Event *ArtExhibitionApprovalForAll // Event containing the contract specifics and raw log

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
func (it *ArtExhibitionApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArtExhibitionApprovalForAll)
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
		it.Event = new(ArtExhibitionApprovalForAll)
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
func (it *ArtExhibitionApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ArtExhibitionApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ArtExhibitionApprovalForAll represents a ApprovalForAll event raised by the ArtExhibition contract.
type ArtExhibitionApprovalForAll struct {
	Account  common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_ArtExhibition *ArtExhibitionFilterer) FilterApprovalForAll(opts *bind.FilterOpts, account []common.Address, operator []common.Address) (*ArtExhibitionApprovalForAllIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ArtExhibition.contract.FilterLogs(opts, "ApprovalForAll", accountRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &ArtExhibitionApprovalForAllIterator{contract: _ArtExhibition.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_ArtExhibition *ArtExhibitionFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *ArtExhibitionApprovalForAll, account []common.Address, operator []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ArtExhibition.contract.WatchLogs(opts, "ApprovalForAll", accountRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ArtExhibitionApprovalForAll)
				if err := _ArtExhibition.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_ArtExhibition *ArtExhibitionFilterer) ParseApprovalForAll(log types.Log) (*ArtExhibitionApprovalForAll, error) {
	event := new(ArtExhibitionApprovalForAll)
	if err := _ArtExhibition.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ArtExhibitionBurnBatchIterator is returned from FilterBurnBatch and is used to iterate over the raw logs and unpacked data for BurnBatch events raised by the ArtExhibition contract.
type ArtExhibitionBurnBatchIterator struct {
	Event *ArtExhibitionBurnBatch // Event containing the contract specifics and raw log

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
func (it *ArtExhibitionBurnBatchIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArtExhibitionBurnBatch)
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
		it.Event = new(ArtExhibitionBurnBatch)
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
func (it *ArtExhibitionBurnBatchIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ArtExhibitionBurnBatchIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ArtExhibitionBurnBatch represents a BurnBatch event raised by the ArtExhibition contract.
type ArtExhibitionBurnBatch struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	TokenIds []*big.Int
	Values   []*big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterBurnBatch is a free log retrieval operation binding the contract event 0xf65bab0bc14d082dd634b49742360c60a6d1a172715c7ae677c3c74dfa8cd2f9.
//
// Solidity: event BurnBatch(address indexed operator, address indexed from, address indexed to, uint256[] tokenIds, uint256[] values)
func (_ArtExhibition *ArtExhibitionFilterer) FilterBurnBatch(opts *bind.FilterOpts, operator []common.Address, from []common.Address, to []common.Address) (*ArtExhibitionBurnBatchIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ArtExhibition.contract.FilterLogs(opts, "BurnBatch", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ArtExhibitionBurnBatchIterator{contract: _ArtExhibition.contract, event: "BurnBatch", logs: logs, sub: sub}, nil
}

// WatchBurnBatch is a free log subscription operation binding the contract event 0xf65bab0bc14d082dd634b49742360c60a6d1a172715c7ae677c3c74dfa8cd2f9.
//
// Solidity: event BurnBatch(address indexed operator, address indexed from, address indexed to, uint256[] tokenIds, uint256[] values)
func (_ArtExhibition *ArtExhibitionFilterer) WatchBurnBatch(opts *bind.WatchOpts, sink chan<- *ArtExhibitionBurnBatch, operator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ArtExhibition.contract.WatchLogs(opts, "BurnBatch", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ArtExhibitionBurnBatch)
				if err := _ArtExhibition.contract.UnpackLog(event, "BurnBatch", log); err != nil {
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

// ParseBurnBatch is a log parse operation binding the contract event 0xf65bab0bc14d082dd634b49742360c60a6d1a172715c7ae677c3c74dfa8cd2f9.
//
// Solidity: event BurnBatch(address indexed operator, address indexed from, address indexed to, uint256[] tokenIds, uint256[] values)
func (_ArtExhibition *ArtExhibitionFilterer) ParseBurnBatch(log types.Log) (*ArtExhibitionBurnBatch, error) {
	event := new(ArtExhibitionBurnBatch)
	if err := _ArtExhibition.contract.UnpackLog(event, "BurnBatch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ArtExhibitionBurnSingleIterator is returned from FilterBurnSingle and is used to iterate over the raw logs and unpacked data for BurnSingle events raised by the ArtExhibition contract.
type ArtExhibitionBurnSingleIterator struct {
	Event *ArtExhibitionBurnSingle // Event containing the contract specifics and raw log

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
func (it *ArtExhibitionBurnSingleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArtExhibitionBurnSingle)
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
		it.Event = new(ArtExhibitionBurnSingle)
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
func (it *ArtExhibitionBurnSingleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ArtExhibitionBurnSingleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ArtExhibitionBurnSingle represents a BurnSingle event raised by the ArtExhibition contract.
type ArtExhibitionBurnSingle struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	TokenId  *big.Int
	Value    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterBurnSingle is a free log retrieval operation binding the contract event 0x8d3d421b8b55218c70b2351439776839bded6f4eae983cde6e4e4d44474a8805.
//
// Solidity: event BurnSingle(address indexed operator, address indexed from, address indexed to, uint256 tokenId, uint256 value)
func (_ArtExhibition *ArtExhibitionFilterer) FilterBurnSingle(opts *bind.FilterOpts, operator []common.Address, from []common.Address, to []common.Address) (*ArtExhibitionBurnSingleIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ArtExhibition.contract.FilterLogs(opts, "BurnSingle", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ArtExhibitionBurnSingleIterator{contract: _ArtExhibition.contract, event: "BurnSingle", logs: logs, sub: sub}, nil
}

// WatchBurnSingle is a free log subscription operation binding the contract event 0x8d3d421b8b55218c70b2351439776839bded6f4eae983cde6e4e4d44474a8805.
//
// Solidity: event BurnSingle(address indexed operator, address indexed from, address indexed to, uint256 tokenId, uint256 value)
func (_ArtExhibition *ArtExhibitionFilterer) WatchBurnSingle(opts *bind.WatchOpts, sink chan<- *ArtExhibitionBurnSingle, operator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ArtExhibition.contract.WatchLogs(opts, "BurnSingle", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ArtExhibitionBurnSingle)
				if err := _ArtExhibition.contract.UnpackLog(event, "BurnSingle", log); err != nil {
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

// ParseBurnSingle is a log parse operation binding the contract event 0x8d3d421b8b55218c70b2351439776839bded6f4eae983cde6e4e4d44474a8805.
//
// Solidity: event BurnSingle(address indexed operator, address indexed from, address indexed to, uint256 tokenId, uint256 value)
func (_ArtExhibition *ArtExhibitionFilterer) ParseBurnSingle(log types.Log) (*ArtExhibitionBurnSingle, error) {
	event := new(ArtExhibitionBurnSingle)
	if err := _ArtExhibition.contract.UnpackLog(event, "BurnSingle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ArtExhibitionGiftIterator is returned from FilterGift and is used to iterate over the raw logs and unpacked data for Gift events raised by the ArtExhibition contract.
type ArtExhibitionGiftIterator struct {
	Event *ArtExhibitionGift // Event containing the contract specifics and raw log

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
func (it *ArtExhibitionGiftIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArtExhibitionGift)
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
		it.Event = new(ArtExhibitionGift)
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
func (it *ArtExhibitionGiftIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ArtExhibitionGiftIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ArtExhibitionGift represents a Gift event raised by the ArtExhibition contract.
type ArtExhibitionGift struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	TokenId  *big.Int
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterGift is a free log retrieval operation binding the contract event 0xb092eacc598875ec236f41a54517e10c908d3e9cab1622a983fe32ecf2f7fe69.
//
// Solidity: event Gift(address operator, address indexed from, address indexed to, uint256 indexed tokenId, uint256 amount)
func (_ArtExhibition *ArtExhibitionFilterer) FilterGift(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*ArtExhibitionGiftIterator, error) {

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

	logs, sub, err := _ArtExhibition.contract.FilterLogs(opts, "Gift", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ArtExhibitionGiftIterator{contract: _ArtExhibition.contract, event: "Gift", logs: logs, sub: sub}, nil
}

// WatchGift is a free log subscription operation binding the contract event 0xb092eacc598875ec236f41a54517e10c908d3e9cab1622a983fe32ecf2f7fe69.
//
// Solidity: event Gift(address operator, address indexed from, address indexed to, uint256 indexed tokenId, uint256 amount)
func (_ArtExhibition *ArtExhibitionFilterer) WatchGift(opts *bind.WatchOpts, sink chan<- *ArtExhibitionGift, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _ArtExhibition.contract.WatchLogs(opts, "Gift", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ArtExhibitionGift)
				if err := _ArtExhibition.contract.UnpackLog(event, "Gift", log); err != nil {
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

// ParseGift is a log parse operation binding the contract event 0xb092eacc598875ec236f41a54517e10c908d3e9cab1622a983fe32ecf2f7fe69.
//
// Solidity: event Gift(address operator, address indexed from, address indexed to, uint256 indexed tokenId, uint256 amount)
func (_ArtExhibition *ArtExhibitionFilterer) ParseGift(log types.Log) (*ArtExhibitionGift, error) {
	event := new(ArtExhibitionGift)
	if err := _ArtExhibition.contract.UnpackLog(event, "Gift", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ArtExhibitionGiftBatchIterator is returned from FilterGiftBatch and is used to iterate over the raw logs and unpacked data for GiftBatch events raised by the ArtExhibition contract.
type ArtExhibitionGiftBatchIterator struct {
	Event *ArtExhibitionGiftBatch // Event containing the contract specifics and raw log

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
func (it *ArtExhibitionGiftBatchIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArtExhibitionGiftBatch)
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
		it.Event = new(ArtExhibitionGiftBatch)
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
func (it *ArtExhibitionGiftBatchIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ArtExhibitionGiftBatchIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ArtExhibitionGiftBatch represents a GiftBatch event raised by the ArtExhibition contract.
type ArtExhibitionGiftBatch struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	TokenIds []*big.Int
	Amount   []*big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterGiftBatch is a free log retrieval operation binding the contract event 0x3aa1c9897b88302c4e6d8e9e4095cc87c6065f7028a29f26c713c9ab1e725ee7.
//
// Solidity: event GiftBatch(address operator, address indexed from, address indexed to, uint256[] tokenIds, uint256[] amount)
func (_ArtExhibition *ArtExhibitionFilterer) FilterGiftBatch(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ArtExhibitionGiftBatchIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ArtExhibition.contract.FilterLogs(opts, "GiftBatch", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ArtExhibitionGiftBatchIterator{contract: _ArtExhibition.contract, event: "GiftBatch", logs: logs, sub: sub}, nil
}

// WatchGiftBatch is a free log subscription operation binding the contract event 0x3aa1c9897b88302c4e6d8e9e4095cc87c6065f7028a29f26c713c9ab1e725ee7.
//
// Solidity: event GiftBatch(address operator, address indexed from, address indexed to, uint256[] tokenIds, uint256[] amount)
func (_ArtExhibition *ArtExhibitionFilterer) WatchGiftBatch(opts *bind.WatchOpts, sink chan<- *ArtExhibitionGiftBatch, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ArtExhibition.contract.WatchLogs(opts, "GiftBatch", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ArtExhibitionGiftBatch)
				if err := _ArtExhibition.contract.UnpackLog(event, "GiftBatch", log); err != nil {
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

// ParseGiftBatch is a log parse operation binding the contract event 0x3aa1c9897b88302c4e6d8e9e4095cc87c6065f7028a29f26c713c9ab1e725ee7.
//
// Solidity: event GiftBatch(address operator, address indexed from, address indexed to, uint256[] tokenIds, uint256[] amount)
func (_ArtExhibition *ArtExhibitionFilterer) ParseGiftBatch(log types.Log) (*ArtExhibitionGiftBatch, error) {
	event := new(ArtExhibitionGiftBatch)
	if err := _ArtExhibition.contract.UnpackLog(event, "GiftBatch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ArtExhibitionGrantRoleIterator is returned from FilterGrantRole and is used to iterate over the raw logs and unpacked data for GrantRole events raised by the ArtExhibition contract.
type ArtExhibitionGrantRoleIterator struct {
	Event *ArtExhibitionGrantRole // Event containing the contract specifics and raw log

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
func (it *ArtExhibitionGrantRoleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArtExhibitionGrantRole)
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
		it.Event = new(ArtExhibitionGrantRole)
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
func (it *ArtExhibitionGrantRoleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ArtExhibitionGrantRoleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ArtExhibitionGrantRole represents a GrantRole event raised by the ArtExhibition contract.
type ArtExhibitionGrantRole struct {
	ContratAdress common.Address
	Operator      common.Address
	Account       common.Address
	Role          uint8
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterGrantRole is a free log retrieval operation binding the contract event 0x7ee499468b11c5fd78e2af0a6821cac160d8b71e0d88ed9398fe819db39ee58d.
//
// Solidity: event GrantRole(address contratAdress, address indexed operator, address indexed account, uint8 indexed role)
func (_ArtExhibition *ArtExhibitionFilterer) FilterGrantRole(opts *bind.FilterOpts, operator []common.Address, account []common.Address, role []uint8) (*ArtExhibitionGrantRoleIterator, error) {

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

	logs, sub, err := _ArtExhibition.contract.FilterLogs(opts, "GrantRole", operatorRule, accountRule, roleRule)
	if err != nil {
		return nil, err
	}
	return &ArtExhibitionGrantRoleIterator{contract: _ArtExhibition.contract, event: "GrantRole", logs: logs, sub: sub}, nil
}

// WatchGrantRole is a free log subscription operation binding the contract event 0x7ee499468b11c5fd78e2af0a6821cac160d8b71e0d88ed9398fe819db39ee58d.
//
// Solidity: event GrantRole(address contratAdress, address indexed operator, address indexed account, uint8 indexed role)
func (_ArtExhibition *ArtExhibitionFilterer) WatchGrantRole(opts *bind.WatchOpts, sink chan<- *ArtExhibitionGrantRole, operator []common.Address, account []common.Address, role []uint8) (event.Subscription, error) {

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

	logs, sub, err := _ArtExhibition.contract.WatchLogs(opts, "GrantRole", operatorRule, accountRule, roleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ArtExhibitionGrantRole)
				if err := _ArtExhibition.contract.UnpackLog(event, "GrantRole", log); err != nil {
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
func (_ArtExhibition *ArtExhibitionFilterer) ParseGrantRole(log types.Log) (*ArtExhibitionGrantRole, error) {
	event := new(ArtExhibitionGrantRole)
	if err := _ArtExhibition.contract.UnpackLog(event, "GrantRole", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ArtExhibitionMintBatchItemsIterator is returned from FilterMintBatchItems and is used to iterate over the raw logs and unpacked data for MintBatchItems events raised by the ArtExhibition contract.
type ArtExhibitionMintBatchItemsIterator struct {
	Event *ArtExhibitionMintBatchItems // Event containing the contract specifics and raw log

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
func (it *ArtExhibitionMintBatchItemsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArtExhibitionMintBatchItems)
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
		it.Event = new(ArtExhibitionMintBatchItems)
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
func (it *ArtExhibitionMintBatchItemsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ArtExhibitionMintBatchItemsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ArtExhibitionMintBatchItems represents a MintBatchItems event raised by the ArtExhibition contract.
type ArtExhibitionMintBatchItems struct {
	Operator  common.Address
	To        common.Address
	TokenIds  []*big.Int
	Amounts   []*big.Int
	TokenUris []string
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterMintBatchItems is a free log retrieval operation binding the contract event 0x67cbb50eda25c4f94b31205b994dd650951f4fa187dccd35c59bfdbc465b2998.
//
// Solidity: event MintBatchItems(address indexed operator, address indexed to, uint256[] tokenIds, uint256[] amounts, string[] tokenUris)
func (_ArtExhibition *ArtExhibitionFilterer) FilterMintBatchItems(opts *bind.FilterOpts, operator []common.Address, to []common.Address) (*ArtExhibitionMintBatchItemsIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ArtExhibition.contract.FilterLogs(opts, "MintBatchItems", operatorRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ArtExhibitionMintBatchItemsIterator{contract: _ArtExhibition.contract, event: "MintBatchItems", logs: logs, sub: sub}, nil
}

// WatchMintBatchItems is a free log subscription operation binding the contract event 0x67cbb50eda25c4f94b31205b994dd650951f4fa187dccd35c59bfdbc465b2998.
//
// Solidity: event MintBatchItems(address indexed operator, address indexed to, uint256[] tokenIds, uint256[] amounts, string[] tokenUris)
func (_ArtExhibition *ArtExhibitionFilterer) WatchMintBatchItems(opts *bind.WatchOpts, sink chan<- *ArtExhibitionMintBatchItems, operator []common.Address, to []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ArtExhibition.contract.WatchLogs(opts, "MintBatchItems", operatorRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ArtExhibitionMintBatchItems)
				if err := _ArtExhibition.contract.UnpackLog(event, "MintBatchItems", log); err != nil {
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

// ParseMintBatchItems is a log parse operation binding the contract event 0x67cbb50eda25c4f94b31205b994dd650951f4fa187dccd35c59bfdbc465b2998.
//
// Solidity: event MintBatchItems(address indexed operator, address indexed to, uint256[] tokenIds, uint256[] amounts, string[] tokenUris)
func (_ArtExhibition *ArtExhibitionFilterer) ParseMintBatchItems(log types.Log) (*ArtExhibitionMintBatchItems, error) {
	event := new(ArtExhibitionMintBatchItems)
	if err := _ArtExhibition.contract.UnpackLog(event, "MintBatchItems", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ArtExhibitionMintSingleItemIterator is returned from FilterMintSingleItem and is used to iterate over the raw logs and unpacked data for MintSingleItem events raised by the ArtExhibition contract.
type ArtExhibitionMintSingleItemIterator struct {
	Event *ArtExhibitionMintSingleItem // Event containing the contract specifics and raw log

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
func (it *ArtExhibitionMintSingleItemIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArtExhibitionMintSingleItem)
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
		it.Event = new(ArtExhibitionMintSingleItem)
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
func (it *ArtExhibitionMintSingleItemIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ArtExhibitionMintSingleItemIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ArtExhibitionMintSingleItem represents a MintSingleItem event raised by the ArtExhibition contract.
type ArtExhibitionMintSingleItem struct {
	Operator common.Address
	To       common.Address
	TokenId  *big.Int
	Amount   *big.Int
	TokenURI string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterMintSingleItem is a free log retrieval operation binding the contract event 0xa100de2e3bb58a491db343c5680e49f0a98376bd96d1071edd8b45e4cbadb321.
//
// Solidity: event MintSingleItem(address operator, address indexed to, uint256 indexed tokenId, uint256 indexed amount, string tokenURI)
func (_ArtExhibition *ArtExhibitionFilterer) FilterMintSingleItem(opts *bind.FilterOpts, to []common.Address, tokenId []*big.Int, amount []*big.Int) (*ArtExhibitionMintSingleItemIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _ArtExhibition.contract.FilterLogs(opts, "MintSingleItem", toRule, tokenIdRule, amountRule)
	if err != nil {
		return nil, err
	}
	return &ArtExhibitionMintSingleItemIterator{contract: _ArtExhibition.contract, event: "MintSingleItem", logs: logs, sub: sub}, nil
}

// WatchMintSingleItem is a free log subscription operation binding the contract event 0xa100de2e3bb58a491db343c5680e49f0a98376bd96d1071edd8b45e4cbadb321.
//
// Solidity: event MintSingleItem(address operator, address indexed to, uint256 indexed tokenId, uint256 indexed amount, string tokenURI)
func (_ArtExhibition *ArtExhibitionFilterer) WatchMintSingleItem(opts *bind.WatchOpts, sink chan<- *ArtExhibitionMintSingleItem, to []common.Address, tokenId []*big.Int, amount []*big.Int) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _ArtExhibition.contract.WatchLogs(opts, "MintSingleItem", toRule, tokenIdRule, amountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ArtExhibitionMintSingleItem)
				if err := _ArtExhibition.contract.UnpackLog(event, "MintSingleItem", log); err != nil {
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

// ParseMintSingleItem is a log parse operation binding the contract event 0xa100de2e3bb58a491db343c5680e49f0a98376bd96d1071edd8b45e4cbadb321.
//
// Solidity: event MintSingleItem(address operator, address indexed to, uint256 indexed tokenId, uint256 indexed amount, string tokenURI)
func (_ArtExhibition *ArtExhibitionFilterer) ParseMintSingleItem(log types.Log) (*ArtExhibitionMintSingleItem, error) {
	event := new(ArtExhibitionMintSingleItem)
	if err := _ArtExhibition.contract.UnpackLog(event, "MintSingleItem", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ArtExhibitionOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ArtExhibition contract.
type ArtExhibitionOwnershipTransferredIterator struct {
	Event *ArtExhibitionOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ArtExhibitionOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArtExhibitionOwnershipTransferred)
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
		it.Event = new(ArtExhibitionOwnershipTransferred)
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
func (it *ArtExhibitionOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ArtExhibitionOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ArtExhibitionOwnershipTransferred represents a OwnershipTransferred event raised by the ArtExhibition contract.
type ArtExhibitionOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ArtExhibition *ArtExhibitionFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ArtExhibitionOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ArtExhibition.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ArtExhibitionOwnershipTransferredIterator{contract: _ArtExhibition.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ArtExhibition *ArtExhibitionFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ArtExhibitionOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ArtExhibition.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ArtExhibitionOwnershipTransferred)
				if err := _ArtExhibition.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_ArtExhibition *ArtExhibitionFilterer) ParseOwnershipTransferred(log types.Log) (*ArtExhibitionOwnershipTransferred, error) {
	event := new(ArtExhibitionOwnershipTransferred)
	if err := _ArtExhibition.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ArtExhibitionRevokeRoleIterator is returned from FilterRevokeRole and is used to iterate over the raw logs and unpacked data for RevokeRole events raised by the ArtExhibition contract.
type ArtExhibitionRevokeRoleIterator struct {
	Event *ArtExhibitionRevokeRole // Event containing the contract specifics and raw log

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
func (it *ArtExhibitionRevokeRoleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArtExhibitionRevokeRole)
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
		it.Event = new(ArtExhibitionRevokeRole)
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
func (it *ArtExhibitionRevokeRoleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ArtExhibitionRevokeRoleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ArtExhibitionRevokeRole represents a RevokeRole event raised by the ArtExhibition contract.
type ArtExhibitionRevokeRole struct {
	ContratAdress common.Address
	Operator      common.Address
	Account       common.Address
	Role          uint8
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterRevokeRole is a free log retrieval operation binding the contract event 0xa3847ba69b1901955bb65a1e552d84a42c25411b4a60509c2ebf380c3a0be301.
//
// Solidity: event RevokeRole(address contratAdress, address indexed operator, address indexed account, uint8 indexed role)
func (_ArtExhibition *ArtExhibitionFilterer) FilterRevokeRole(opts *bind.FilterOpts, operator []common.Address, account []common.Address, role []uint8) (*ArtExhibitionRevokeRoleIterator, error) {

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

	logs, sub, err := _ArtExhibition.contract.FilterLogs(opts, "RevokeRole", operatorRule, accountRule, roleRule)
	if err != nil {
		return nil, err
	}
	return &ArtExhibitionRevokeRoleIterator{contract: _ArtExhibition.contract, event: "RevokeRole", logs: logs, sub: sub}, nil
}

// WatchRevokeRole is a free log subscription operation binding the contract event 0xa3847ba69b1901955bb65a1e552d84a42c25411b4a60509c2ebf380c3a0be301.
//
// Solidity: event RevokeRole(address contratAdress, address indexed operator, address indexed account, uint8 indexed role)
func (_ArtExhibition *ArtExhibitionFilterer) WatchRevokeRole(opts *bind.WatchOpts, sink chan<- *ArtExhibitionRevokeRole, operator []common.Address, account []common.Address, role []uint8) (event.Subscription, error) {

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

	logs, sub, err := _ArtExhibition.contract.WatchLogs(opts, "RevokeRole", operatorRule, accountRule, roleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ArtExhibitionRevokeRole)
				if err := _ArtExhibition.contract.UnpackLog(event, "RevokeRole", log); err != nil {
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
func (_ArtExhibition *ArtExhibitionFilterer) ParseRevokeRole(log types.Log) (*ArtExhibitionRevokeRole, error) {
	event := new(ArtExhibitionRevokeRole)
	if err := _ArtExhibition.contract.UnpackLog(event, "RevokeRole", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ArtExhibitionRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the ArtExhibition contract.
type ArtExhibitionRoleAdminChangedIterator struct {
	Event *ArtExhibitionRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *ArtExhibitionRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArtExhibitionRoleAdminChanged)
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
		it.Event = new(ArtExhibitionRoleAdminChanged)
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
func (it *ArtExhibitionRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ArtExhibitionRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ArtExhibitionRoleAdminChanged represents a RoleAdminChanged event raised by the ArtExhibition contract.
type ArtExhibitionRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_ArtExhibition *ArtExhibitionFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*ArtExhibitionRoleAdminChangedIterator, error) {

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

	logs, sub, err := _ArtExhibition.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &ArtExhibitionRoleAdminChangedIterator{contract: _ArtExhibition.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_ArtExhibition *ArtExhibitionFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *ArtExhibitionRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _ArtExhibition.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ArtExhibitionRoleAdminChanged)
				if err := _ArtExhibition.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_ArtExhibition *ArtExhibitionFilterer) ParseRoleAdminChanged(log types.Log) (*ArtExhibitionRoleAdminChanged, error) {
	event := new(ArtExhibitionRoleAdminChanged)
	if err := _ArtExhibition.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ArtExhibitionRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the ArtExhibition contract.
type ArtExhibitionRoleGrantedIterator struct {
	Event *ArtExhibitionRoleGranted // Event containing the contract specifics and raw log

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
func (it *ArtExhibitionRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArtExhibitionRoleGranted)
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
		it.Event = new(ArtExhibitionRoleGranted)
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
func (it *ArtExhibitionRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ArtExhibitionRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ArtExhibitionRoleGranted represents a RoleGranted event raised by the ArtExhibition contract.
type ArtExhibitionRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_ArtExhibition *ArtExhibitionFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ArtExhibitionRoleGrantedIterator, error) {

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

	logs, sub, err := _ArtExhibition.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ArtExhibitionRoleGrantedIterator{contract: _ArtExhibition.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_ArtExhibition *ArtExhibitionFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *ArtExhibitionRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _ArtExhibition.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ArtExhibitionRoleGranted)
				if err := _ArtExhibition.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_ArtExhibition *ArtExhibitionFilterer) ParseRoleGranted(log types.Log) (*ArtExhibitionRoleGranted, error) {
	event := new(ArtExhibitionRoleGranted)
	if err := _ArtExhibition.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ArtExhibitionRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the ArtExhibition contract.
type ArtExhibitionRoleRevokedIterator struct {
	Event *ArtExhibitionRoleRevoked // Event containing the contract specifics and raw log

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
func (it *ArtExhibitionRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArtExhibitionRoleRevoked)
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
		it.Event = new(ArtExhibitionRoleRevoked)
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
func (it *ArtExhibitionRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ArtExhibitionRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ArtExhibitionRoleRevoked represents a RoleRevoked event raised by the ArtExhibition contract.
type ArtExhibitionRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_ArtExhibition *ArtExhibitionFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ArtExhibitionRoleRevokedIterator, error) {

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

	logs, sub, err := _ArtExhibition.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ArtExhibitionRoleRevokedIterator{contract: _ArtExhibition.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_ArtExhibition *ArtExhibitionFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *ArtExhibitionRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _ArtExhibition.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ArtExhibitionRoleRevoked)
				if err := _ArtExhibition.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_ArtExhibition *ArtExhibitionFilterer) ParseRoleRevoked(log types.Log) (*ArtExhibitionRoleRevoked, error) {
	event := new(ArtExhibitionRoleRevoked)
	if err := _ArtExhibition.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ArtExhibitionTransferBatchIterator is returned from FilterTransferBatch and is used to iterate over the raw logs and unpacked data for TransferBatch events raised by the ArtExhibition contract.
type ArtExhibitionTransferBatchIterator struct {
	Event *ArtExhibitionTransferBatch // Event containing the contract specifics and raw log

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
func (it *ArtExhibitionTransferBatchIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArtExhibitionTransferBatch)
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
		it.Event = new(ArtExhibitionTransferBatch)
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
func (it *ArtExhibitionTransferBatchIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ArtExhibitionTransferBatchIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ArtExhibitionTransferBatch represents a TransferBatch event raised by the ArtExhibition contract.
type ArtExhibitionTransferBatch struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	Ids      []*big.Int
	Values   []*big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransferBatch is a free log retrieval operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_ArtExhibition *ArtExhibitionFilterer) FilterTransferBatch(opts *bind.FilterOpts, operator []common.Address, from []common.Address, to []common.Address) (*ArtExhibitionTransferBatchIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ArtExhibition.contract.FilterLogs(opts, "TransferBatch", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ArtExhibitionTransferBatchIterator{contract: _ArtExhibition.contract, event: "TransferBatch", logs: logs, sub: sub}, nil
}

// WatchTransferBatch is a free log subscription operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_ArtExhibition *ArtExhibitionFilterer) WatchTransferBatch(opts *bind.WatchOpts, sink chan<- *ArtExhibitionTransferBatch, operator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ArtExhibition.contract.WatchLogs(opts, "TransferBatch", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ArtExhibitionTransferBatch)
				if err := _ArtExhibition.contract.UnpackLog(event, "TransferBatch", log); err != nil {
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

// ParseTransferBatch is a log parse operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_ArtExhibition *ArtExhibitionFilterer) ParseTransferBatch(log types.Log) (*ArtExhibitionTransferBatch, error) {
	event := new(ArtExhibitionTransferBatch)
	if err := _ArtExhibition.contract.UnpackLog(event, "TransferBatch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ArtExhibitionTransferSingleIterator is returned from FilterTransferSingle and is used to iterate over the raw logs and unpacked data for TransferSingle events raised by the ArtExhibition contract.
type ArtExhibitionTransferSingleIterator struct {
	Event *ArtExhibitionTransferSingle // Event containing the contract specifics and raw log

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
func (it *ArtExhibitionTransferSingleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArtExhibitionTransferSingle)
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
		it.Event = new(ArtExhibitionTransferSingle)
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
func (it *ArtExhibitionTransferSingleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ArtExhibitionTransferSingleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ArtExhibitionTransferSingle represents a TransferSingle event raised by the ArtExhibition contract.
type ArtExhibitionTransferSingle struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	Id       *big.Int
	Value    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransferSingle is a free log retrieval operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_ArtExhibition *ArtExhibitionFilterer) FilterTransferSingle(opts *bind.FilterOpts, operator []common.Address, from []common.Address, to []common.Address) (*ArtExhibitionTransferSingleIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ArtExhibition.contract.FilterLogs(opts, "TransferSingle", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ArtExhibitionTransferSingleIterator{contract: _ArtExhibition.contract, event: "TransferSingle", logs: logs, sub: sub}, nil
}

// WatchTransferSingle is a free log subscription operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_ArtExhibition *ArtExhibitionFilterer) WatchTransferSingle(opts *bind.WatchOpts, sink chan<- *ArtExhibitionTransferSingle, operator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ArtExhibition.contract.WatchLogs(opts, "TransferSingle", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ArtExhibitionTransferSingle)
				if err := _ArtExhibition.contract.UnpackLog(event, "TransferSingle", log); err != nil {
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

// ParseTransferSingle is a log parse operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_ArtExhibition *ArtExhibitionFilterer) ParseTransferSingle(log types.Log) (*ArtExhibitionTransferSingle, error) {
	event := new(ArtExhibitionTransferSingle)
	if err := _ArtExhibition.contract.UnpackLog(event, "TransferSingle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ArtExhibitionURIIterator is returned from FilterURI and is used to iterate over the raw logs and unpacked data for URI events raised by the ArtExhibition contract.
type ArtExhibitionURIIterator struct {
	Event *ArtExhibitionURI // Event containing the contract specifics and raw log

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
func (it *ArtExhibitionURIIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArtExhibitionURI)
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
		it.Event = new(ArtExhibitionURI)
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
func (it *ArtExhibitionURIIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ArtExhibitionURIIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ArtExhibitionURI represents a URI event raised by the ArtExhibition contract.
type ArtExhibitionURI struct {
	Value string
	Id    *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterURI is a free log retrieval operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_ArtExhibition *ArtExhibitionFilterer) FilterURI(opts *bind.FilterOpts, id []*big.Int) (*ArtExhibitionURIIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _ArtExhibition.contract.FilterLogs(opts, "URI", idRule)
	if err != nil {
		return nil, err
	}
	return &ArtExhibitionURIIterator{contract: _ArtExhibition.contract, event: "URI", logs: logs, sub: sub}, nil
}

// WatchURI is a free log subscription operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_ArtExhibition *ArtExhibitionFilterer) WatchURI(opts *bind.WatchOpts, sink chan<- *ArtExhibitionURI, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _ArtExhibition.contract.WatchLogs(opts, "URI", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ArtExhibitionURI)
				if err := _ArtExhibition.contract.UnpackLog(event, "URI", log); err != nil {
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

// ParseURI is a log parse operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_ArtExhibition *ArtExhibitionFilterer) ParseURI(log types.Log) (*ArtExhibitionURI, error) {
	event := new(ArtExhibitionURI)
	if err := _ArtExhibition.contract.UnpackLog(event, "URI", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
