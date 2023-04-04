// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package solar1155NFT

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

// Solar1155NFTMetaData contains all meta data concerning the Solar1155NFT contract.
var Solar1155NFTMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"}],\"name\":\"BurnBatch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"BurnSingle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_seller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_orderId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_unitPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"_feePercent\",\"type\":\"uint16\"}],\"name\":\"CancelOrder\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Gift\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"amount\",\"type\":\"uint256[]\"}],\"name\":\"GiftBatch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"unitPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"feePercent\",\"type\":\"uint16\"}],\"name\":\"MakeOrder\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_tokenURI\",\"type\":\"string\"}],\"name\":\"Mint1155Item\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"_tokenIds\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"_amounts\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"string[]\",\"name\":\"_tokenUris\",\"type\":\"string[]\"}],\"name\":\"MintBatchItems\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_buyer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_seller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_tradeId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_orderId\",\"type\":\"uint256\"}],\"name\":\"TradeNft\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"}],\"name\":\"TransferBatch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"TransferSingle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"URI\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_seller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_orderId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_unitPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"_feePercent\",\"type\":\"uint16\"}],\"name\":\"UpdateOrder\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"}],\"name\":\"balanceOfBatch\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"}],\"name\":\"burnBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"gift\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"giftBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"grantMintRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"hasMintRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceMintRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeBatchTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"uri\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"},{\"internalType\":\"uint16\",\"name\":\"platformFeePercent\",\"type\":\"uint16\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"platformFeePercent\",\"type\":\"uint16\"}],\"name\":\"setPlatformFeePercent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPlatformFeePercent\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"}],\"name\":\"grantMarket\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMarket\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"tokenURI\",\"type\":\"string\"}],\"name\":\"mintItem\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"string[]\",\"name\":\"tokenUris\",\"type\":\"string[]\"}],\"name\":\"mintBatchItems\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unitPrice\",\"type\":\"uint256\"}],\"name\":\"makeOrder\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unitPrice\",\"type\":\"uint256\"}],\"name\":\"makeOrderProxy\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"}],\"name\":\"cancleOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stockAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sellingUnitPrice\",\"type\":\"uint256\"}],\"name\":\"updateOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"orderId_\",\"type\":\"uint256\"}],\"name\":\"getOrder\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"fee\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"sellingPrice\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"buyer_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"orderId_\",\"type\":\"uint256\"}],\"name\":\"tradeNft\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// Solar1155NFTABI is the input ABI used to generate the binding from.
// Deprecated: Use Solar1155NFTMetaData.ABI instead.
var Solar1155NFTABI = Solar1155NFTMetaData.ABI

// Solar1155NFT is an auto generated Go binding around an Ethereum contract.
type Solar1155NFT struct {
	Solar1155NFTCaller     // Read-only binding to the contract
	Solar1155NFTTransactor // Write-only binding to the contract
	Solar1155NFTFilterer   // Log filterer for contract events
}

// Solar1155NFTCaller is an auto generated read-only Go binding around an Ethereum contract.
type Solar1155NFTCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Solar1155NFTTransactor is an auto generated write-only Go binding around an Ethereum contract.
type Solar1155NFTTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Solar1155NFTFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Solar1155NFTFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Solar1155NFTSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Solar1155NFTSession struct {
	Contract     *Solar1155NFT     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Solar1155NFTCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Solar1155NFTCallerSession struct {
	Contract *Solar1155NFTCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// Solar1155NFTTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Solar1155NFTTransactorSession struct {
	Contract     *Solar1155NFTTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// Solar1155NFTRaw is an auto generated low-level Go binding around an Ethereum contract.
type Solar1155NFTRaw struct {
	Contract *Solar1155NFT // Generic contract binding to access the raw methods on
}

// Solar1155NFTCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Solar1155NFTCallerRaw struct {
	Contract *Solar1155NFTCaller // Generic read-only contract binding to access the raw methods on
}

// Solar1155NFTTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Solar1155NFTTransactorRaw struct {
	Contract *Solar1155NFTTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSolar1155NFT creates a new instance of Solar1155NFT, bound to a specific deployed contract.
func NewSolar1155NFT(address common.Address, backend bind.ContractBackend) (*Solar1155NFT, error) {
	contract, err := bindSolar1155NFT(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Solar1155NFT{Solar1155NFTCaller: Solar1155NFTCaller{contract: contract}, Solar1155NFTTransactor: Solar1155NFTTransactor{contract: contract}, Solar1155NFTFilterer: Solar1155NFTFilterer{contract: contract}}, nil
}

// NewSolar1155NFTCaller creates a new read-only instance of Solar1155NFT, bound to a specific deployed contract.
func NewSolar1155NFTCaller(address common.Address, caller bind.ContractCaller) (*Solar1155NFTCaller, error) {
	contract, err := bindSolar1155NFT(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Solar1155NFTCaller{contract: contract}, nil
}

// NewSolar1155NFTTransactor creates a new write-only instance of Solar1155NFT, bound to a specific deployed contract.
func NewSolar1155NFTTransactor(address common.Address, transactor bind.ContractTransactor) (*Solar1155NFTTransactor, error) {
	contract, err := bindSolar1155NFT(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Solar1155NFTTransactor{contract: contract}, nil
}

// NewSolar1155NFTFilterer creates a new log filterer instance of Solar1155NFT, bound to a specific deployed contract.
func NewSolar1155NFTFilterer(address common.Address, filterer bind.ContractFilterer) (*Solar1155NFTFilterer, error) {
	contract, err := bindSolar1155NFT(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Solar1155NFTFilterer{contract: contract}, nil
}

// bindSolar1155NFT binds a generic wrapper to an already deployed contract.
func bindSolar1155NFT(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(Solar1155NFTABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Solar1155NFT *Solar1155NFTRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Solar1155NFT.Contract.Solar1155NFTCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Solar1155NFT *Solar1155NFTRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Solar1155NFT.Contract.Solar1155NFTTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Solar1155NFT *Solar1155NFTRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Solar1155NFT.Contract.Solar1155NFTTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Solar1155NFT *Solar1155NFTCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Solar1155NFT.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Solar1155NFT *Solar1155NFTTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Solar1155NFT.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Solar1155NFT *Solar1155NFTTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Solar1155NFT.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Solar1155NFT *Solar1155NFTCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Solar1155NFT.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Solar1155NFT *Solar1155NFTSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Solar1155NFT.Contract.DEFAULTADMINROLE(&_Solar1155NFT.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Solar1155NFT *Solar1155NFTCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Solar1155NFT.Contract.DEFAULTADMINROLE(&_Solar1155NFT.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_Solar1155NFT *Solar1155NFTCaller) BalanceOf(opts *bind.CallOpts, account common.Address, id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Solar1155NFT.contract.Call(opts, &out, "balanceOf", account, id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_Solar1155NFT *Solar1155NFTSession) BalanceOf(account common.Address, id *big.Int) (*big.Int, error) {
	return _Solar1155NFT.Contract.BalanceOf(&_Solar1155NFT.CallOpts, account, id)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_Solar1155NFT *Solar1155NFTCallerSession) BalanceOf(account common.Address, id *big.Int) (*big.Int, error) {
	return _Solar1155NFT.Contract.BalanceOf(&_Solar1155NFT.CallOpts, account, id)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_Solar1155NFT *Solar1155NFTCaller) BalanceOfBatch(opts *bind.CallOpts, accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _Solar1155NFT.contract.Call(opts, &out, "balanceOfBatch", accounts, ids)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_Solar1155NFT *Solar1155NFTSession) BalanceOfBatch(accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	return _Solar1155NFT.Contract.BalanceOfBatch(&_Solar1155NFT.CallOpts, accounts, ids)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_Solar1155NFT *Solar1155NFTCallerSession) BalanceOfBatch(accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	return _Solar1155NFT.Contract.BalanceOfBatch(&_Solar1155NFT.CallOpts, accounts, ids)
}

// GetMarket is a free data retrieval call binding the contract method 0xf1be1679.
//
// Solidity: function getMarket() view returns(address)
func (_Solar1155NFT *Solar1155NFTCaller) GetMarket(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Solar1155NFT.contract.Call(opts, &out, "getMarket")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetMarket is a free data retrieval call binding the contract method 0xf1be1679.
//
// Solidity: function getMarket() view returns(address)
func (_Solar1155NFT *Solar1155NFTSession) GetMarket() (common.Address, error) {
	return _Solar1155NFT.Contract.GetMarket(&_Solar1155NFT.CallOpts)
}

// GetMarket is a free data retrieval call binding the contract method 0xf1be1679.
//
// Solidity: function getMarket() view returns(address)
func (_Solar1155NFT *Solar1155NFTCallerSession) GetMarket() (common.Address, error) {
	return _Solar1155NFT.Contract.GetMarket(&_Solar1155NFT.CallOpts)
}

// GetOrder is a free data retrieval call binding the contract method 0xd09ef241.
//
// Solidity: function getOrder(uint256 orderId_) view returns(address seller, uint256 tokenId, uint256 amount, uint16 fee, uint256 sellingPrice)
func (_Solar1155NFT *Solar1155NFTCaller) GetOrder(opts *bind.CallOpts, orderId_ *big.Int) (struct {
	Seller       common.Address
	TokenId      *big.Int
	Amount       *big.Int
	Fee          uint16
	SellingPrice *big.Int
}, error) {
	var out []interface{}
	err := _Solar1155NFT.contract.Call(opts, &out, "getOrder", orderId_)

	outstruct := new(struct {
		Seller       common.Address
		TokenId      *big.Int
		Amount       *big.Int
		Fee          uint16
		SellingPrice *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Seller = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.TokenId = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Amount = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Fee = *abi.ConvertType(out[3], new(uint16)).(*uint16)
	outstruct.SellingPrice = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetOrder is a free data retrieval call binding the contract method 0xd09ef241.
//
// Solidity: function getOrder(uint256 orderId_) view returns(address seller, uint256 tokenId, uint256 amount, uint16 fee, uint256 sellingPrice)
func (_Solar1155NFT *Solar1155NFTSession) GetOrder(orderId_ *big.Int) (struct {
	Seller       common.Address
	TokenId      *big.Int
	Amount       *big.Int
	Fee          uint16
	SellingPrice *big.Int
}, error) {
	return _Solar1155NFT.Contract.GetOrder(&_Solar1155NFT.CallOpts, orderId_)
}

// GetOrder is a free data retrieval call binding the contract method 0xd09ef241.
//
// Solidity: function getOrder(uint256 orderId_) view returns(address seller, uint256 tokenId, uint256 amount, uint16 fee, uint256 sellingPrice)
func (_Solar1155NFT *Solar1155NFTCallerSession) GetOrder(orderId_ *big.Int) (struct {
	Seller       common.Address
	TokenId      *big.Int
	Amount       *big.Int
	Fee          uint16
	SellingPrice *big.Int
}, error) {
	return _Solar1155NFT.Contract.GetOrder(&_Solar1155NFT.CallOpts, orderId_)
}

// GetPlatformFeePercent is a free data retrieval call binding the contract method 0x4c9c3f85.
//
// Solidity: function getPlatformFeePercent() view returns(uint16)
func (_Solar1155NFT *Solar1155NFTCaller) GetPlatformFeePercent(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _Solar1155NFT.contract.Call(opts, &out, "getPlatformFeePercent")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// GetPlatformFeePercent is a free data retrieval call binding the contract method 0x4c9c3f85.
//
// Solidity: function getPlatformFeePercent() view returns(uint16)
func (_Solar1155NFT *Solar1155NFTSession) GetPlatformFeePercent() (uint16, error) {
	return _Solar1155NFT.Contract.GetPlatformFeePercent(&_Solar1155NFT.CallOpts)
}

// GetPlatformFeePercent is a free data retrieval call binding the contract method 0x4c9c3f85.
//
// Solidity: function getPlatformFeePercent() view returns(uint16)
func (_Solar1155NFT *Solar1155NFTCallerSession) GetPlatformFeePercent() (uint16, error) {
	return _Solar1155NFT.Contract.GetPlatformFeePercent(&_Solar1155NFT.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Solar1155NFT *Solar1155NFTCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Solar1155NFT.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Solar1155NFT *Solar1155NFTSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Solar1155NFT.Contract.GetRoleAdmin(&_Solar1155NFT.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Solar1155NFT *Solar1155NFTCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Solar1155NFT.Contract.GetRoleAdmin(&_Solar1155NFT.CallOpts, role)
}

// HasMintRole is a free data retrieval call binding the contract method 0xc4e58e54.
//
// Solidity: function hasMintRole(address _account) view returns(bool)
func (_Solar1155NFT *Solar1155NFTCaller) HasMintRole(opts *bind.CallOpts, _account common.Address) (bool, error) {
	var out []interface{}
	err := _Solar1155NFT.contract.Call(opts, &out, "hasMintRole", _account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasMintRole is a free data retrieval call binding the contract method 0xc4e58e54.
//
// Solidity: function hasMintRole(address _account) view returns(bool)
func (_Solar1155NFT *Solar1155NFTSession) HasMintRole(_account common.Address) (bool, error) {
	return _Solar1155NFT.Contract.HasMintRole(&_Solar1155NFT.CallOpts, _account)
}

// HasMintRole is a free data retrieval call binding the contract method 0xc4e58e54.
//
// Solidity: function hasMintRole(address _account) view returns(bool)
func (_Solar1155NFT *Solar1155NFTCallerSession) HasMintRole(_account common.Address) (bool, error) {
	return _Solar1155NFT.Contract.HasMintRole(&_Solar1155NFT.CallOpts, _account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Solar1155NFT *Solar1155NFTCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Solar1155NFT.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Solar1155NFT *Solar1155NFTSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Solar1155NFT.Contract.HasRole(&_Solar1155NFT.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Solar1155NFT *Solar1155NFTCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Solar1155NFT.Contract.HasRole(&_Solar1155NFT.CallOpts, role, account)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_Solar1155NFT *Solar1155NFTCaller) IsApprovedForAll(opts *bind.CallOpts, account common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Solar1155NFT.contract.Call(opts, &out, "isApprovedForAll", account, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_Solar1155NFT *Solar1155NFTSession) IsApprovedForAll(account common.Address, operator common.Address) (bool, error) {
	return _Solar1155NFT.Contract.IsApprovedForAll(&_Solar1155NFT.CallOpts, account, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_Solar1155NFT *Solar1155NFTCallerSession) IsApprovedForAll(account common.Address, operator common.Address) (bool, error) {
	return _Solar1155NFT.Contract.IsApprovedForAll(&_Solar1155NFT.CallOpts, account, operator)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Solar1155NFT *Solar1155NFTCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Solar1155NFT.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Solar1155NFT *Solar1155NFTSession) Name() (string, error) {
	return _Solar1155NFT.Contract.Name(&_Solar1155NFT.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Solar1155NFT *Solar1155NFTCallerSession) Name() (string, error) {
	return _Solar1155NFT.Contract.Name(&_Solar1155NFT.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Solar1155NFT *Solar1155NFTCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Solar1155NFT.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Solar1155NFT *Solar1155NFTSession) Owner() (common.Address, error) {
	return _Solar1155NFT.Contract.Owner(&_Solar1155NFT.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Solar1155NFT *Solar1155NFTCallerSession) Owner() (common.Address, error) {
	return _Solar1155NFT.Contract.Owner(&_Solar1155NFT.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Solar1155NFT *Solar1155NFTCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Solar1155NFT.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Solar1155NFT *Solar1155NFTSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Solar1155NFT.Contract.SupportsInterface(&_Solar1155NFT.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Solar1155NFT *Solar1155NFTCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Solar1155NFT.Contract.SupportsInterface(&_Solar1155NFT.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Solar1155NFT *Solar1155NFTCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Solar1155NFT.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Solar1155NFT *Solar1155NFTSession) Symbol() (string, error) {
	return _Solar1155NFT.Contract.Symbol(&_Solar1155NFT.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Solar1155NFT *Solar1155NFTCallerSession) Symbol() (string, error) {
	return _Solar1155NFT.Contract.Symbol(&_Solar1155NFT.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Solar1155NFT *Solar1155NFTCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Solar1155NFT.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Solar1155NFT *Solar1155NFTSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Solar1155NFT.Contract.TokenURI(&_Solar1155NFT.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Solar1155NFT *Solar1155NFTCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Solar1155NFT.Contract.TokenURI(&_Solar1155NFT.CallOpts, tokenId)
}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 tokenId) view returns(string)
func (_Solar1155NFT *Solar1155NFTCaller) Uri(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Solar1155NFT.contract.Call(opts, &out, "uri", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 tokenId) view returns(string)
func (_Solar1155NFT *Solar1155NFTSession) Uri(tokenId *big.Int) (string, error) {
	return _Solar1155NFT.Contract.Uri(&_Solar1155NFT.CallOpts, tokenId)
}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 tokenId) view returns(string)
func (_Solar1155NFT *Solar1155NFTCallerSession) Uri(tokenId *big.Int) (string, error) {
	return _Solar1155NFT.Contract.Uri(&_Solar1155NFT.CallOpts, tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0xf5298aca.
//
// Solidity: function burn(address account, uint256 id, uint256 value) returns()
func (_Solar1155NFT *Solar1155NFTTransactor) Burn(opts *bind.TransactOpts, account common.Address, id *big.Int, value *big.Int) (*types.Transaction, error) {
	return _Solar1155NFT.contract.Transact(opts, "burn", account, id, value)
}

// Burn is a paid mutator transaction binding the contract method 0xf5298aca.
//
// Solidity: function burn(address account, uint256 id, uint256 value) returns()
func (_Solar1155NFT *Solar1155NFTSession) Burn(account common.Address, id *big.Int, value *big.Int) (*types.Transaction, error) {
	return _Solar1155NFT.Contract.Burn(&_Solar1155NFT.TransactOpts, account, id, value)
}

// Burn is a paid mutator transaction binding the contract method 0xf5298aca.
//
// Solidity: function burn(address account, uint256 id, uint256 value) returns()
func (_Solar1155NFT *Solar1155NFTTransactorSession) Burn(account common.Address, id *big.Int, value *big.Int) (*types.Transaction, error) {
	return _Solar1155NFT.Contract.Burn(&_Solar1155NFT.TransactOpts, account, id, value)
}

// BurnBatch is a paid mutator transaction binding the contract method 0x6b20c454.
//
// Solidity: function burnBatch(address account, uint256[] ids, uint256[] values) returns()
func (_Solar1155NFT *Solar1155NFTTransactor) BurnBatch(opts *bind.TransactOpts, account common.Address, ids []*big.Int, values []*big.Int) (*types.Transaction, error) {
	return _Solar1155NFT.contract.Transact(opts, "burnBatch", account, ids, values)
}

// BurnBatch is a paid mutator transaction binding the contract method 0x6b20c454.
//
// Solidity: function burnBatch(address account, uint256[] ids, uint256[] values) returns()
func (_Solar1155NFT *Solar1155NFTSession) BurnBatch(account common.Address, ids []*big.Int, values []*big.Int) (*types.Transaction, error) {
	return _Solar1155NFT.Contract.BurnBatch(&_Solar1155NFT.TransactOpts, account, ids, values)
}

// BurnBatch is a paid mutator transaction binding the contract method 0x6b20c454.
//
// Solidity: function burnBatch(address account, uint256[] ids, uint256[] values) returns()
func (_Solar1155NFT *Solar1155NFTTransactorSession) BurnBatch(account common.Address, ids []*big.Int, values []*big.Int) (*types.Transaction, error) {
	return _Solar1155NFT.Contract.BurnBatch(&_Solar1155NFT.TransactOpts, account, ids, values)
}

// CancleOrder is a paid mutator transaction binding the contract method 0xf406cb8a.
//
// Solidity: function cancleOrder(uint256 orderId) returns()
func (_Solar1155NFT *Solar1155NFTTransactor) CancleOrder(opts *bind.TransactOpts, orderId *big.Int) (*types.Transaction, error) {
	return _Solar1155NFT.contract.Transact(opts, "cancleOrder", orderId)
}

// CancleOrder is a paid mutator transaction binding the contract method 0xf406cb8a.
//
// Solidity: function cancleOrder(uint256 orderId) returns()
func (_Solar1155NFT *Solar1155NFTSession) CancleOrder(orderId *big.Int) (*types.Transaction, error) {
	return _Solar1155NFT.Contract.CancleOrder(&_Solar1155NFT.TransactOpts, orderId)
}

// CancleOrder is a paid mutator transaction binding the contract method 0xf406cb8a.
//
// Solidity: function cancleOrder(uint256 orderId) returns()
func (_Solar1155NFT *Solar1155NFTTransactorSession) CancleOrder(orderId *big.Int) (*types.Transaction, error) {
	return _Solar1155NFT.Contract.CancleOrder(&_Solar1155NFT.TransactOpts, orderId)
}

// Gift is a paid mutator transaction binding the contract method 0x388a1cf2.
//
// Solidity: function gift(address from, address to, uint256 id, uint256 amount) returns()
func (_Solar1155NFT *Solar1155NFTTransactor) Gift(opts *bind.TransactOpts, from common.Address, to common.Address, id *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Solar1155NFT.contract.Transact(opts, "gift", from, to, id, amount)
}

// Gift is a paid mutator transaction binding the contract method 0x388a1cf2.
//
// Solidity: function gift(address from, address to, uint256 id, uint256 amount) returns()
func (_Solar1155NFT *Solar1155NFTSession) Gift(from common.Address, to common.Address, id *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Solar1155NFT.Contract.Gift(&_Solar1155NFT.TransactOpts, from, to, id, amount)
}

// Gift is a paid mutator transaction binding the contract method 0x388a1cf2.
//
// Solidity: function gift(address from, address to, uint256 id, uint256 amount) returns()
func (_Solar1155NFT *Solar1155NFTTransactorSession) Gift(from common.Address, to common.Address, id *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Solar1155NFT.Contract.Gift(&_Solar1155NFT.TransactOpts, from, to, id, amount)
}

// GiftBatch is a paid mutator transaction binding the contract method 0x21bb0b08.
//
// Solidity: function giftBatch(address from, address to, uint256[] ids, uint256[] amounts) returns()
func (_Solar1155NFT *Solar1155NFTTransactor) GiftBatch(opts *bind.TransactOpts, from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int) (*types.Transaction, error) {
	return _Solar1155NFT.contract.Transact(opts, "giftBatch", from, to, ids, amounts)
}

// GiftBatch is a paid mutator transaction binding the contract method 0x21bb0b08.
//
// Solidity: function giftBatch(address from, address to, uint256[] ids, uint256[] amounts) returns()
func (_Solar1155NFT *Solar1155NFTSession) GiftBatch(from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int) (*types.Transaction, error) {
	return _Solar1155NFT.Contract.GiftBatch(&_Solar1155NFT.TransactOpts, from, to, ids, amounts)
}

// GiftBatch is a paid mutator transaction binding the contract method 0x21bb0b08.
//
// Solidity: function giftBatch(address from, address to, uint256[] ids, uint256[] amounts) returns()
func (_Solar1155NFT *Solar1155NFTTransactorSession) GiftBatch(from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int) (*types.Transaction, error) {
	return _Solar1155NFT.Contract.GiftBatch(&_Solar1155NFT.TransactOpts, from, to, ids, amounts)
}

// GrantMarket is a paid mutator transaction binding the contract method 0x040add28.
//
// Solidity: function grantMarket(address market) returns(bool)
func (_Solar1155NFT *Solar1155NFTTransactor) GrantMarket(opts *bind.TransactOpts, market common.Address) (*types.Transaction, error) {
	return _Solar1155NFT.contract.Transact(opts, "grantMarket", market)
}

// GrantMarket is a paid mutator transaction binding the contract method 0x040add28.
//
// Solidity: function grantMarket(address market) returns(bool)
func (_Solar1155NFT *Solar1155NFTSession) GrantMarket(market common.Address) (*types.Transaction, error) {
	return _Solar1155NFT.Contract.GrantMarket(&_Solar1155NFT.TransactOpts, market)
}

// GrantMarket is a paid mutator transaction binding the contract method 0x040add28.
//
// Solidity: function grantMarket(address market) returns(bool)
func (_Solar1155NFT *Solar1155NFTTransactorSession) GrantMarket(market common.Address) (*types.Transaction, error) {
	return _Solar1155NFT.Contract.GrantMarket(&_Solar1155NFT.TransactOpts, market)
}

// GrantMintRole is a paid mutator transaction binding the contract method 0xc2e3273d.
//
// Solidity: function grantMintRole(address operator) returns()
func (_Solar1155NFT *Solar1155NFTTransactor) GrantMintRole(opts *bind.TransactOpts, operator common.Address) (*types.Transaction, error) {
	return _Solar1155NFT.contract.Transact(opts, "grantMintRole", operator)
}

// GrantMintRole is a paid mutator transaction binding the contract method 0xc2e3273d.
//
// Solidity: function grantMintRole(address operator) returns()
func (_Solar1155NFT *Solar1155NFTSession) GrantMintRole(operator common.Address) (*types.Transaction, error) {
	return _Solar1155NFT.Contract.GrantMintRole(&_Solar1155NFT.TransactOpts, operator)
}

// GrantMintRole is a paid mutator transaction binding the contract method 0xc2e3273d.
//
// Solidity: function grantMintRole(address operator) returns()
func (_Solar1155NFT *Solar1155NFTTransactorSession) GrantMintRole(operator common.Address) (*types.Transaction, error) {
	return _Solar1155NFT.Contract.GrantMintRole(&_Solar1155NFT.TransactOpts, operator)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Solar1155NFT *Solar1155NFTTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Solar1155NFT.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Solar1155NFT *Solar1155NFTSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Solar1155NFT.Contract.GrantRole(&_Solar1155NFT.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Solar1155NFT *Solar1155NFTTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Solar1155NFT.Contract.GrantRole(&_Solar1155NFT.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0x335a3e6c.
//
// Solidity: function initialize(string name, string symbol, string uri, uint16 platformFeePercent) returns()
func (_Solar1155NFT *Solar1155NFTTransactor) Initialize(opts *bind.TransactOpts, name string, symbol string, uri string, platformFeePercent uint16) (*types.Transaction, error) {
	return _Solar1155NFT.contract.Transact(opts, "initialize", name, symbol, uri, platformFeePercent)
}

// Initialize is a paid mutator transaction binding the contract method 0x335a3e6c.
//
// Solidity: function initialize(string name, string symbol, string uri, uint16 platformFeePercent) returns()
func (_Solar1155NFT *Solar1155NFTSession) Initialize(name string, symbol string, uri string, platformFeePercent uint16) (*types.Transaction, error) {
	return _Solar1155NFT.Contract.Initialize(&_Solar1155NFT.TransactOpts, name, symbol, uri, platformFeePercent)
}

// Initialize is a paid mutator transaction binding the contract method 0x335a3e6c.
//
// Solidity: function initialize(string name, string symbol, string uri, uint16 platformFeePercent) returns()
func (_Solar1155NFT *Solar1155NFTTransactorSession) Initialize(name string, symbol string, uri string, platformFeePercent uint16) (*types.Transaction, error) {
	return _Solar1155NFT.Contract.Initialize(&_Solar1155NFT.TransactOpts, name, symbol, uri, platformFeePercent)
}

// MakeOrder is a paid mutator transaction binding the contract method 0x19d616bc.
//
// Solidity: function makeOrder(uint256 tokenId, uint256 amount, uint256 unitPrice) returns(uint256)
func (_Solar1155NFT *Solar1155NFTTransactor) MakeOrder(opts *bind.TransactOpts, tokenId *big.Int, amount *big.Int, unitPrice *big.Int) (*types.Transaction, error) {
	return _Solar1155NFT.contract.Transact(opts, "makeOrder", tokenId, amount, unitPrice)
}

// MakeOrder is a paid mutator transaction binding the contract method 0x19d616bc.
//
// Solidity: function makeOrder(uint256 tokenId, uint256 amount, uint256 unitPrice) returns(uint256)
func (_Solar1155NFT *Solar1155NFTSession) MakeOrder(tokenId *big.Int, amount *big.Int, unitPrice *big.Int) (*types.Transaction, error) {
	return _Solar1155NFT.Contract.MakeOrder(&_Solar1155NFT.TransactOpts, tokenId, amount, unitPrice)
}

// MakeOrder is a paid mutator transaction binding the contract method 0x19d616bc.
//
// Solidity: function makeOrder(uint256 tokenId, uint256 amount, uint256 unitPrice) returns(uint256)
func (_Solar1155NFT *Solar1155NFTTransactorSession) MakeOrder(tokenId *big.Int, amount *big.Int, unitPrice *big.Int) (*types.Transaction, error) {
	return _Solar1155NFT.Contract.MakeOrder(&_Solar1155NFT.TransactOpts, tokenId, amount, unitPrice)
}

// MakeOrderProxy is a paid mutator transaction binding the contract method 0xec3599e3.
//
// Solidity: function makeOrderProxy(address from, uint256 tokenId, uint256 amount, uint256 unitPrice) returns(uint256)
func (_Solar1155NFT *Solar1155NFTTransactor) MakeOrderProxy(opts *bind.TransactOpts, from common.Address, tokenId *big.Int, amount *big.Int, unitPrice *big.Int) (*types.Transaction, error) {
	return _Solar1155NFT.contract.Transact(opts, "makeOrderProxy", from, tokenId, amount, unitPrice)
}

// MakeOrderProxy is a paid mutator transaction binding the contract method 0xec3599e3.
//
// Solidity: function makeOrderProxy(address from, uint256 tokenId, uint256 amount, uint256 unitPrice) returns(uint256)
func (_Solar1155NFT *Solar1155NFTSession) MakeOrderProxy(from common.Address, tokenId *big.Int, amount *big.Int, unitPrice *big.Int) (*types.Transaction, error) {
	return _Solar1155NFT.Contract.MakeOrderProxy(&_Solar1155NFT.TransactOpts, from, tokenId, amount, unitPrice)
}

// MakeOrderProxy is a paid mutator transaction binding the contract method 0xec3599e3.
//
// Solidity: function makeOrderProxy(address from, uint256 tokenId, uint256 amount, uint256 unitPrice) returns(uint256)
func (_Solar1155NFT *Solar1155NFTTransactorSession) MakeOrderProxy(from common.Address, tokenId *big.Int, amount *big.Int, unitPrice *big.Int) (*types.Transaction, error) {
	return _Solar1155NFT.Contract.MakeOrderProxy(&_Solar1155NFT.TransactOpts, from, tokenId, amount, unitPrice)
}

// MintBatchItems is a paid mutator transaction binding the contract method 0x239079bb.
//
// Solidity: function mintBatchItems(address to, uint256[] ids, uint256[] amounts, string[] tokenUris) returns()
func (_Solar1155NFT *Solar1155NFTTransactor) MintBatchItems(opts *bind.TransactOpts, to common.Address, ids []*big.Int, amounts []*big.Int, tokenUris []string) (*types.Transaction, error) {
	return _Solar1155NFT.contract.Transact(opts, "mintBatchItems", to, ids, amounts, tokenUris)
}

// MintBatchItems is a paid mutator transaction binding the contract method 0x239079bb.
//
// Solidity: function mintBatchItems(address to, uint256[] ids, uint256[] amounts, string[] tokenUris) returns()
func (_Solar1155NFT *Solar1155NFTSession) MintBatchItems(to common.Address, ids []*big.Int, amounts []*big.Int, tokenUris []string) (*types.Transaction, error) {
	return _Solar1155NFT.Contract.MintBatchItems(&_Solar1155NFT.TransactOpts, to, ids, amounts, tokenUris)
}

// MintBatchItems is a paid mutator transaction binding the contract method 0x239079bb.
//
// Solidity: function mintBatchItems(address to, uint256[] ids, uint256[] amounts, string[] tokenUris) returns()
func (_Solar1155NFT *Solar1155NFTTransactorSession) MintBatchItems(to common.Address, ids []*big.Int, amounts []*big.Int, tokenUris []string) (*types.Transaction, error) {
	return _Solar1155NFT.Contract.MintBatchItems(&_Solar1155NFT.TransactOpts, to, ids, amounts, tokenUris)
}

// MintItem is a paid mutator transaction binding the contract method 0x21fc9040.
//
// Solidity: function mintItem(address to, uint256 id, uint256 amount, string tokenURI) returns()
func (_Solar1155NFT *Solar1155NFTTransactor) MintItem(opts *bind.TransactOpts, to common.Address, id *big.Int, amount *big.Int, tokenURI string) (*types.Transaction, error) {
	return _Solar1155NFT.contract.Transact(opts, "mintItem", to, id, amount, tokenURI)
}

// MintItem is a paid mutator transaction binding the contract method 0x21fc9040.
//
// Solidity: function mintItem(address to, uint256 id, uint256 amount, string tokenURI) returns()
func (_Solar1155NFT *Solar1155NFTSession) MintItem(to common.Address, id *big.Int, amount *big.Int, tokenURI string) (*types.Transaction, error) {
	return _Solar1155NFT.Contract.MintItem(&_Solar1155NFT.TransactOpts, to, id, amount, tokenURI)
}

// MintItem is a paid mutator transaction binding the contract method 0x21fc9040.
//
// Solidity: function mintItem(address to, uint256 id, uint256 amount, string tokenURI) returns()
func (_Solar1155NFT *Solar1155NFTTransactorSession) MintItem(to common.Address, id *big.Int, amount *big.Int, tokenURI string) (*types.Transaction, error) {
	return _Solar1155NFT.Contract.MintItem(&_Solar1155NFT.TransactOpts, to, id, amount, tokenURI)
}

// RenounceMintRole is a paid mutator transaction binding the contract method 0xa49a177a.
//
// Solidity: function renounceMintRole(address account) returns()
func (_Solar1155NFT *Solar1155NFTTransactor) RenounceMintRole(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _Solar1155NFT.contract.Transact(opts, "renounceMintRole", account)
}

// RenounceMintRole is a paid mutator transaction binding the contract method 0xa49a177a.
//
// Solidity: function renounceMintRole(address account) returns()
func (_Solar1155NFT *Solar1155NFTSession) RenounceMintRole(account common.Address) (*types.Transaction, error) {
	return _Solar1155NFT.Contract.RenounceMintRole(&_Solar1155NFT.TransactOpts, account)
}

// RenounceMintRole is a paid mutator transaction binding the contract method 0xa49a177a.
//
// Solidity: function renounceMintRole(address account) returns()
func (_Solar1155NFT *Solar1155NFTTransactorSession) RenounceMintRole(account common.Address) (*types.Transaction, error) {
	return _Solar1155NFT.Contract.RenounceMintRole(&_Solar1155NFT.TransactOpts, account)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Solar1155NFT *Solar1155NFTTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Solar1155NFT.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Solar1155NFT *Solar1155NFTSession) RenounceOwnership() (*types.Transaction, error) {
	return _Solar1155NFT.Contract.RenounceOwnership(&_Solar1155NFT.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Solar1155NFT *Solar1155NFTTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Solar1155NFT.Contract.RenounceOwnership(&_Solar1155NFT.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Solar1155NFT *Solar1155NFTTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Solar1155NFT.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Solar1155NFT *Solar1155NFTSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Solar1155NFT.Contract.RenounceRole(&_Solar1155NFT.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Solar1155NFT *Solar1155NFTTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Solar1155NFT.Contract.RenounceRole(&_Solar1155NFT.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Solar1155NFT *Solar1155NFTTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Solar1155NFT.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Solar1155NFT *Solar1155NFTSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Solar1155NFT.Contract.RevokeRole(&_Solar1155NFT.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Solar1155NFT *Solar1155NFTTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Solar1155NFT.Contract.RevokeRole(&_Solar1155NFT.TransactOpts, role, account)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_Solar1155NFT *Solar1155NFTTransactor) SafeBatchTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _Solar1155NFT.contract.Transact(opts, "safeBatchTransferFrom", from, to, ids, amounts, data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_Solar1155NFT *Solar1155NFTSession) SafeBatchTransferFrom(from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _Solar1155NFT.Contract.SafeBatchTransferFrom(&_Solar1155NFT.TransactOpts, from, to, ids, amounts, data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_Solar1155NFT *Solar1155NFTTransactorSession) SafeBatchTransferFrom(from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _Solar1155NFT.Contract.SafeBatchTransferFrom(&_Solar1155NFT.TransactOpts, from, to, ids, amounts, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_Solar1155NFT *Solar1155NFTTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _Solar1155NFT.contract.Transact(opts, "safeTransferFrom", from, to, id, amount, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_Solar1155NFT *Solar1155NFTSession) SafeTransferFrom(from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _Solar1155NFT.Contract.SafeTransferFrom(&_Solar1155NFT.TransactOpts, from, to, id, amount, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_Solar1155NFT *Solar1155NFTTransactorSession) SafeTransferFrom(from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _Solar1155NFT.Contract.SafeTransferFrom(&_Solar1155NFT.TransactOpts, from, to, id, amount, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Solar1155NFT *Solar1155NFTTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _Solar1155NFT.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Solar1155NFT *Solar1155NFTSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Solar1155NFT.Contract.SetApprovalForAll(&_Solar1155NFT.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Solar1155NFT *Solar1155NFTTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Solar1155NFT.Contract.SetApprovalForAll(&_Solar1155NFT.TransactOpts, operator, approved)
}

// SetPlatformFeePercent is a paid mutator transaction binding the contract method 0x685538df.
//
// Solidity: function setPlatformFeePercent(uint16 platformFeePercent) returns()
func (_Solar1155NFT *Solar1155NFTTransactor) SetPlatformFeePercent(opts *bind.TransactOpts, platformFeePercent uint16) (*types.Transaction, error) {
	return _Solar1155NFT.contract.Transact(opts, "setPlatformFeePercent", platformFeePercent)
}

// SetPlatformFeePercent is a paid mutator transaction binding the contract method 0x685538df.
//
// Solidity: function setPlatformFeePercent(uint16 platformFeePercent) returns()
func (_Solar1155NFT *Solar1155NFTSession) SetPlatformFeePercent(platformFeePercent uint16) (*types.Transaction, error) {
	return _Solar1155NFT.Contract.SetPlatformFeePercent(&_Solar1155NFT.TransactOpts, platformFeePercent)
}

// SetPlatformFeePercent is a paid mutator transaction binding the contract method 0x685538df.
//
// Solidity: function setPlatformFeePercent(uint16 platformFeePercent) returns()
func (_Solar1155NFT *Solar1155NFTTransactorSession) SetPlatformFeePercent(platformFeePercent uint16) (*types.Transaction, error) {
	return _Solar1155NFT.Contract.SetPlatformFeePercent(&_Solar1155NFT.TransactOpts, platformFeePercent)
}

// TradeNft is a paid mutator transaction binding the contract method 0x7d3405ac.
//
// Solidity: function tradeNft(address buyer_, uint256 orderId_) returns(uint256)
func (_Solar1155NFT *Solar1155NFTTransactor) TradeNft(opts *bind.TransactOpts, buyer_ common.Address, orderId_ *big.Int) (*types.Transaction, error) {
	return _Solar1155NFT.contract.Transact(opts, "tradeNft", buyer_, orderId_)
}

// TradeNft is a paid mutator transaction binding the contract method 0x7d3405ac.
//
// Solidity: function tradeNft(address buyer_, uint256 orderId_) returns(uint256)
func (_Solar1155NFT *Solar1155NFTSession) TradeNft(buyer_ common.Address, orderId_ *big.Int) (*types.Transaction, error) {
	return _Solar1155NFT.Contract.TradeNft(&_Solar1155NFT.TransactOpts, buyer_, orderId_)
}

// TradeNft is a paid mutator transaction binding the contract method 0x7d3405ac.
//
// Solidity: function tradeNft(address buyer_, uint256 orderId_) returns(uint256)
func (_Solar1155NFT *Solar1155NFTTransactorSession) TradeNft(buyer_ common.Address, orderId_ *big.Int) (*types.Transaction, error) {
	return _Solar1155NFT.Contract.TradeNft(&_Solar1155NFT.TransactOpts, buyer_, orderId_)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Solar1155NFT *Solar1155NFTTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Solar1155NFT.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Solar1155NFT *Solar1155NFTSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Solar1155NFT.Contract.TransferOwnership(&_Solar1155NFT.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Solar1155NFT *Solar1155NFTTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Solar1155NFT.Contract.TransferOwnership(&_Solar1155NFT.TransactOpts, newOwner)
}

// UpdateOrder is a paid mutator transaction binding the contract method 0x2ede65df.
//
// Solidity: function updateOrder(uint256 orderId, uint256 stockAmount, uint256 sellingUnitPrice) returns()
func (_Solar1155NFT *Solar1155NFTTransactor) UpdateOrder(opts *bind.TransactOpts, orderId *big.Int, stockAmount *big.Int, sellingUnitPrice *big.Int) (*types.Transaction, error) {
	return _Solar1155NFT.contract.Transact(opts, "updateOrder", orderId, stockAmount, sellingUnitPrice)
}

// UpdateOrder is a paid mutator transaction binding the contract method 0x2ede65df.
//
// Solidity: function updateOrder(uint256 orderId, uint256 stockAmount, uint256 sellingUnitPrice) returns()
func (_Solar1155NFT *Solar1155NFTSession) UpdateOrder(orderId *big.Int, stockAmount *big.Int, sellingUnitPrice *big.Int) (*types.Transaction, error) {
	return _Solar1155NFT.Contract.UpdateOrder(&_Solar1155NFT.TransactOpts, orderId, stockAmount, sellingUnitPrice)
}

// UpdateOrder is a paid mutator transaction binding the contract method 0x2ede65df.
//
// Solidity: function updateOrder(uint256 orderId, uint256 stockAmount, uint256 sellingUnitPrice) returns()
func (_Solar1155NFT *Solar1155NFTTransactorSession) UpdateOrder(orderId *big.Int, stockAmount *big.Int, sellingUnitPrice *big.Int) (*types.Transaction, error) {
	return _Solar1155NFT.Contract.UpdateOrder(&_Solar1155NFT.TransactOpts, orderId, stockAmount, sellingUnitPrice)
}

// Solar1155NFTApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Solar1155NFT contract.
type Solar1155NFTApprovalForAllIterator struct {
	Event *Solar1155NFTApprovalForAll // Event containing the contract specifics and raw log

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
func (it *Solar1155NFTApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Solar1155NFTApprovalForAll)
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
		it.Event = new(Solar1155NFTApprovalForAll)
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
func (it *Solar1155NFTApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Solar1155NFTApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Solar1155NFTApprovalForAll represents a ApprovalForAll event raised by the Solar1155NFT contract.
type Solar1155NFTApprovalForAll struct {
	Account  common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_Solar1155NFT *Solar1155NFTFilterer) FilterApprovalForAll(opts *bind.FilterOpts, account []common.Address, operator []common.Address) (*Solar1155NFTApprovalForAllIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Solar1155NFT.contract.FilterLogs(opts, "ApprovalForAll", accountRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &Solar1155NFTApprovalForAllIterator{contract: _Solar1155NFT.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_Solar1155NFT *Solar1155NFTFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *Solar1155NFTApprovalForAll, account []common.Address, operator []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Solar1155NFT.contract.WatchLogs(opts, "ApprovalForAll", accountRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Solar1155NFTApprovalForAll)
				if err := _Solar1155NFT.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
func (_Solar1155NFT *Solar1155NFTFilterer) ParseApprovalForAll(log types.Log) (*Solar1155NFTApprovalForAll, error) {
	event := new(Solar1155NFTApprovalForAll)
	if err := _Solar1155NFT.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Solar1155NFTBurnBatchIterator is returned from FilterBurnBatch and is used to iterate over the raw logs and unpacked data for BurnBatch events raised by the Solar1155NFT contract.
type Solar1155NFTBurnBatchIterator struct {
	Event *Solar1155NFTBurnBatch // Event containing the contract specifics and raw log

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
func (it *Solar1155NFTBurnBatchIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Solar1155NFTBurnBatch)
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
		it.Event = new(Solar1155NFTBurnBatch)
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
func (it *Solar1155NFTBurnBatchIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Solar1155NFTBurnBatchIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Solar1155NFTBurnBatch represents a BurnBatch event raised by the Solar1155NFT contract.
type Solar1155NFTBurnBatch struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	Ids      []*big.Int
	Values   []*big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterBurnBatch is a free log retrieval operation binding the contract event 0xf65bab0bc14d082dd634b49742360c60a6d1a172715c7ae677c3c74dfa8cd2f9.
//
// Solidity: event BurnBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_Solar1155NFT *Solar1155NFTFilterer) FilterBurnBatch(opts *bind.FilterOpts, operator []common.Address, from []common.Address, to []common.Address) (*Solar1155NFTBurnBatchIterator, error) {

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

	logs, sub, err := _Solar1155NFT.contract.FilterLogs(opts, "BurnBatch", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &Solar1155NFTBurnBatchIterator{contract: _Solar1155NFT.contract, event: "BurnBatch", logs: logs, sub: sub}, nil
}

// WatchBurnBatch is a free log subscription operation binding the contract event 0xf65bab0bc14d082dd634b49742360c60a6d1a172715c7ae677c3c74dfa8cd2f9.
//
// Solidity: event BurnBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_Solar1155NFT *Solar1155NFTFilterer) WatchBurnBatch(opts *bind.WatchOpts, sink chan<- *Solar1155NFTBurnBatch, operator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Solar1155NFT.contract.WatchLogs(opts, "BurnBatch", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Solar1155NFTBurnBatch)
				if err := _Solar1155NFT.contract.UnpackLog(event, "BurnBatch", log); err != nil {
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
// Solidity: event BurnBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_Solar1155NFT *Solar1155NFTFilterer) ParseBurnBatch(log types.Log) (*Solar1155NFTBurnBatch, error) {
	event := new(Solar1155NFTBurnBatch)
	if err := _Solar1155NFT.contract.UnpackLog(event, "BurnBatch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Solar1155NFTBurnSingleIterator is returned from FilterBurnSingle and is used to iterate over the raw logs and unpacked data for BurnSingle events raised by the Solar1155NFT contract.
type Solar1155NFTBurnSingleIterator struct {
	Event *Solar1155NFTBurnSingle // Event containing the contract specifics and raw log

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
func (it *Solar1155NFTBurnSingleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Solar1155NFTBurnSingle)
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
		it.Event = new(Solar1155NFTBurnSingle)
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
func (it *Solar1155NFTBurnSingleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Solar1155NFTBurnSingleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Solar1155NFTBurnSingle represents a BurnSingle event raised by the Solar1155NFT contract.
type Solar1155NFTBurnSingle struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	Id       *big.Int
	Value    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterBurnSingle is a free log retrieval operation binding the contract event 0x8d3d421b8b55218c70b2351439776839bded6f4eae983cde6e4e4d44474a8805.
//
// Solidity: event BurnSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_Solar1155NFT *Solar1155NFTFilterer) FilterBurnSingle(opts *bind.FilterOpts, operator []common.Address, from []common.Address, to []common.Address) (*Solar1155NFTBurnSingleIterator, error) {

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

	logs, sub, err := _Solar1155NFT.contract.FilterLogs(opts, "BurnSingle", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &Solar1155NFTBurnSingleIterator{contract: _Solar1155NFT.contract, event: "BurnSingle", logs: logs, sub: sub}, nil
}

// WatchBurnSingle is a free log subscription operation binding the contract event 0x8d3d421b8b55218c70b2351439776839bded6f4eae983cde6e4e4d44474a8805.
//
// Solidity: event BurnSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_Solar1155NFT *Solar1155NFTFilterer) WatchBurnSingle(opts *bind.WatchOpts, sink chan<- *Solar1155NFTBurnSingle, operator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Solar1155NFT.contract.WatchLogs(opts, "BurnSingle", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Solar1155NFTBurnSingle)
				if err := _Solar1155NFT.contract.UnpackLog(event, "BurnSingle", log); err != nil {
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
// Solidity: event BurnSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_Solar1155NFT *Solar1155NFTFilterer) ParseBurnSingle(log types.Log) (*Solar1155NFTBurnSingle, error) {
	event := new(Solar1155NFTBurnSingle)
	if err := _Solar1155NFT.contract.UnpackLog(event, "BurnSingle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Solar1155NFTCancelOrderIterator is returned from FilterCancelOrder and is used to iterate over the raw logs and unpacked data for CancelOrder events raised by the Solar1155NFT contract.
type Solar1155NFTCancelOrderIterator struct {
	Event *Solar1155NFTCancelOrder // Event containing the contract specifics and raw log

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
func (it *Solar1155NFTCancelOrderIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Solar1155NFTCancelOrder)
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
		it.Event = new(Solar1155NFTCancelOrder)
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
func (it *Solar1155NFTCancelOrderIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Solar1155NFTCancelOrderIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Solar1155NFTCancelOrder represents a CancelOrder event raised by the Solar1155NFT contract.
type Solar1155NFTCancelOrder struct {
	Operator   common.Address
	Seller     common.Address
	TokenId    *big.Int
	Amount     *big.Int
	OrderId    *big.Int
	UnitPrice  *big.Int
	FeePercent uint16
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterCancelOrder is a free log retrieval operation binding the contract event 0x4bcba0ad617ccaa9ebb90efacca1a243ad72e0b94523d17199422eefda213843.
//
// Solidity: event CancelOrder(address _operator, address indexed _seller, uint256 indexed _tokenId, uint256 _amount, uint256 indexed _orderId, uint256 _unitPrice, uint16 _feePercent)
func (_Solar1155NFT *Solar1155NFTFilterer) FilterCancelOrder(opts *bind.FilterOpts, _seller []common.Address, _tokenId []*big.Int, _orderId []*big.Int) (*Solar1155NFTCancelOrderIterator, error) {

	var _sellerRule []interface{}
	for _, _sellerItem := range _seller {
		_sellerRule = append(_sellerRule, _sellerItem)
	}
	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}

	var _orderIdRule []interface{}
	for _, _orderIdItem := range _orderId {
		_orderIdRule = append(_orderIdRule, _orderIdItem)
	}

	logs, sub, err := _Solar1155NFT.contract.FilterLogs(opts, "CancelOrder", _sellerRule, _tokenIdRule, _orderIdRule)
	if err != nil {
		return nil, err
	}
	return &Solar1155NFTCancelOrderIterator{contract: _Solar1155NFT.contract, event: "CancelOrder", logs: logs, sub: sub}, nil
}

// WatchCancelOrder is a free log subscription operation binding the contract event 0x4bcba0ad617ccaa9ebb90efacca1a243ad72e0b94523d17199422eefda213843.
//
// Solidity: event CancelOrder(address _operator, address indexed _seller, uint256 indexed _tokenId, uint256 _amount, uint256 indexed _orderId, uint256 _unitPrice, uint16 _feePercent)
func (_Solar1155NFT *Solar1155NFTFilterer) WatchCancelOrder(opts *bind.WatchOpts, sink chan<- *Solar1155NFTCancelOrder, _seller []common.Address, _tokenId []*big.Int, _orderId []*big.Int) (event.Subscription, error) {

	var _sellerRule []interface{}
	for _, _sellerItem := range _seller {
		_sellerRule = append(_sellerRule, _sellerItem)
	}
	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}

	var _orderIdRule []interface{}
	for _, _orderIdItem := range _orderId {
		_orderIdRule = append(_orderIdRule, _orderIdItem)
	}

	logs, sub, err := _Solar1155NFT.contract.WatchLogs(opts, "CancelOrder", _sellerRule, _tokenIdRule, _orderIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Solar1155NFTCancelOrder)
				if err := _Solar1155NFT.contract.UnpackLog(event, "CancelOrder", log); err != nil {
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

// ParseCancelOrder is a log parse operation binding the contract event 0x4bcba0ad617ccaa9ebb90efacca1a243ad72e0b94523d17199422eefda213843.
//
// Solidity: event CancelOrder(address _operator, address indexed _seller, uint256 indexed _tokenId, uint256 _amount, uint256 indexed _orderId, uint256 _unitPrice, uint16 _feePercent)
func (_Solar1155NFT *Solar1155NFTFilterer) ParseCancelOrder(log types.Log) (*Solar1155NFTCancelOrder, error) {
	event := new(Solar1155NFTCancelOrder)
	if err := _Solar1155NFT.contract.UnpackLog(event, "CancelOrder", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Solar1155NFTGiftIterator is returned from FilterGift and is used to iterate over the raw logs and unpacked data for Gift events raised by the Solar1155NFT contract.
type Solar1155NFTGiftIterator struct {
	Event *Solar1155NFTGift // Event containing the contract specifics and raw log

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
func (it *Solar1155NFTGiftIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Solar1155NFTGift)
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
		it.Event = new(Solar1155NFTGift)
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
func (it *Solar1155NFTGiftIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Solar1155NFTGiftIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Solar1155NFTGift represents a Gift event raised by the Solar1155NFT contract.
type Solar1155NFTGift struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterGift is a free log retrieval operation binding the contract event 0x695f2cbe6dd7aa6bf5917c368e632d9acededcfd29fb98b009e9eefc9da5ecad.
//
// Solidity: event Gift(address indexed from, address indexed to, uint256 indexed tokenId, uint256 amount)
func (_Solar1155NFT *Solar1155NFTFilterer) FilterGift(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*Solar1155NFTGiftIterator, error) {

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

	logs, sub, err := _Solar1155NFT.contract.FilterLogs(opts, "Gift", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &Solar1155NFTGiftIterator{contract: _Solar1155NFT.contract, event: "Gift", logs: logs, sub: sub}, nil
}

// WatchGift is a free log subscription operation binding the contract event 0x695f2cbe6dd7aa6bf5917c368e632d9acededcfd29fb98b009e9eefc9da5ecad.
//
// Solidity: event Gift(address indexed from, address indexed to, uint256 indexed tokenId, uint256 amount)
func (_Solar1155NFT *Solar1155NFTFilterer) WatchGift(opts *bind.WatchOpts, sink chan<- *Solar1155NFTGift, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _Solar1155NFT.contract.WatchLogs(opts, "Gift", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Solar1155NFTGift)
				if err := _Solar1155NFT.contract.UnpackLog(event, "Gift", log); err != nil {
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

// ParseGift is a log parse operation binding the contract event 0x695f2cbe6dd7aa6bf5917c368e632d9acededcfd29fb98b009e9eefc9da5ecad.
//
// Solidity: event Gift(address indexed from, address indexed to, uint256 indexed tokenId, uint256 amount)
func (_Solar1155NFT *Solar1155NFTFilterer) ParseGift(log types.Log) (*Solar1155NFTGift, error) {
	event := new(Solar1155NFTGift)
	if err := _Solar1155NFT.contract.UnpackLog(event, "Gift", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Solar1155NFTGiftBatchIterator is returned from FilterGiftBatch and is used to iterate over the raw logs and unpacked data for GiftBatch events raised by the Solar1155NFT contract.
type Solar1155NFTGiftBatchIterator struct {
	Event *Solar1155NFTGiftBatch // Event containing the contract specifics and raw log

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
func (it *Solar1155NFTGiftBatchIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Solar1155NFTGiftBatch)
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
		it.Event = new(Solar1155NFTGiftBatch)
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
func (it *Solar1155NFTGiftBatchIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Solar1155NFTGiftBatchIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Solar1155NFTGiftBatch represents a GiftBatch event raised by the Solar1155NFT contract.
type Solar1155NFTGiftBatch struct {
	From     common.Address
	To       common.Address
	TokenIds []*big.Int
	Amount   []*big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterGiftBatch is a free log retrieval operation binding the contract event 0x8a5776ef1fec4a5dd9cf679b51062e85dd2f172c194a033cd5cbef388320da96.
//
// Solidity: event GiftBatch(address indexed from, address indexed to, uint256[] tokenIds, uint256[] amount)
func (_Solar1155NFT *Solar1155NFTFilterer) FilterGiftBatch(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*Solar1155NFTGiftBatchIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Solar1155NFT.contract.FilterLogs(opts, "GiftBatch", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &Solar1155NFTGiftBatchIterator{contract: _Solar1155NFT.contract, event: "GiftBatch", logs: logs, sub: sub}, nil
}

// WatchGiftBatch is a free log subscription operation binding the contract event 0x8a5776ef1fec4a5dd9cf679b51062e85dd2f172c194a033cd5cbef388320da96.
//
// Solidity: event GiftBatch(address indexed from, address indexed to, uint256[] tokenIds, uint256[] amount)
func (_Solar1155NFT *Solar1155NFTFilterer) WatchGiftBatch(opts *bind.WatchOpts, sink chan<- *Solar1155NFTGiftBatch, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Solar1155NFT.contract.WatchLogs(opts, "GiftBatch", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Solar1155NFTGiftBatch)
				if err := _Solar1155NFT.contract.UnpackLog(event, "GiftBatch", log); err != nil {
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

// ParseGiftBatch is a log parse operation binding the contract event 0x8a5776ef1fec4a5dd9cf679b51062e85dd2f172c194a033cd5cbef388320da96.
//
// Solidity: event GiftBatch(address indexed from, address indexed to, uint256[] tokenIds, uint256[] amount)
func (_Solar1155NFT *Solar1155NFTFilterer) ParseGiftBatch(log types.Log) (*Solar1155NFTGiftBatch, error) {
	event := new(Solar1155NFTGiftBatch)
	if err := _Solar1155NFT.contract.UnpackLog(event, "GiftBatch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Solar1155NFTMakeOrderIterator is returned from FilterMakeOrder and is used to iterate over the raw logs and unpacked data for MakeOrder events raised by the Solar1155NFT contract.
type Solar1155NFTMakeOrderIterator struct {
	Event *Solar1155NFTMakeOrder // Event containing the contract specifics and raw log

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
func (it *Solar1155NFTMakeOrderIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Solar1155NFTMakeOrder)
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
		it.Event = new(Solar1155NFTMakeOrder)
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
func (it *Solar1155NFTMakeOrderIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Solar1155NFTMakeOrderIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Solar1155NFTMakeOrder represents a MakeOrder event raised by the Solar1155NFT contract.
type Solar1155NFTMakeOrder struct {
	Operator   common.Address
	Seller     common.Address
	TokenId    *big.Int
	Amount     *big.Int
	OrderId    *big.Int
	UnitPrice  *big.Int
	FeePercent uint16
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterMakeOrder is a free log retrieval operation binding the contract event 0xa2d5e68b643c0e0ce64df9740849284545362cd9cf83fbd6264ecf59eada0f4f.
//
// Solidity: event MakeOrder(address operator, address indexed seller, uint256 indexed tokenId, uint256 amount, uint256 indexed orderId, uint256 unitPrice, uint16 feePercent)
func (_Solar1155NFT *Solar1155NFTFilterer) FilterMakeOrder(opts *bind.FilterOpts, seller []common.Address, tokenId []*big.Int, orderId []*big.Int) (*Solar1155NFTMakeOrderIterator, error) {

	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}

	logs, sub, err := _Solar1155NFT.contract.FilterLogs(opts, "MakeOrder", sellerRule, tokenIdRule, orderIdRule)
	if err != nil {
		return nil, err
	}
	return &Solar1155NFTMakeOrderIterator{contract: _Solar1155NFT.contract, event: "MakeOrder", logs: logs, sub: sub}, nil
}

// WatchMakeOrder is a free log subscription operation binding the contract event 0xa2d5e68b643c0e0ce64df9740849284545362cd9cf83fbd6264ecf59eada0f4f.
//
// Solidity: event MakeOrder(address operator, address indexed seller, uint256 indexed tokenId, uint256 amount, uint256 indexed orderId, uint256 unitPrice, uint16 feePercent)
func (_Solar1155NFT *Solar1155NFTFilterer) WatchMakeOrder(opts *bind.WatchOpts, sink chan<- *Solar1155NFTMakeOrder, seller []common.Address, tokenId []*big.Int, orderId []*big.Int) (event.Subscription, error) {

	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}

	logs, sub, err := _Solar1155NFT.contract.WatchLogs(opts, "MakeOrder", sellerRule, tokenIdRule, orderIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Solar1155NFTMakeOrder)
				if err := _Solar1155NFT.contract.UnpackLog(event, "MakeOrder", log); err != nil {
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

// ParseMakeOrder is a log parse operation binding the contract event 0xa2d5e68b643c0e0ce64df9740849284545362cd9cf83fbd6264ecf59eada0f4f.
//
// Solidity: event MakeOrder(address operator, address indexed seller, uint256 indexed tokenId, uint256 amount, uint256 indexed orderId, uint256 unitPrice, uint16 feePercent)
func (_Solar1155NFT *Solar1155NFTFilterer) ParseMakeOrder(log types.Log) (*Solar1155NFTMakeOrder, error) {
	event := new(Solar1155NFTMakeOrder)
	if err := _Solar1155NFT.contract.UnpackLog(event, "MakeOrder", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Solar1155NFTMint1155ItemIterator is returned from FilterMint1155Item and is used to iterate over the raw logs and unpacked data for Mint1155Item events raised by the Solar1155NFT contract.
type Solar1155NFTMint1155ItemIterator struct {
	Event *Solar1155NFTMint1155Item // Event containing the contract specifics and raw log

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
func (it *Solar1155NFTMint1155ItemIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Solar1155NFTMint1155Item)
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
		it.Event = new(Solar1155NFTMint1155Item)
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
func (it *Solar1155NFTMint1155ItemIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Solar1155NFTMint1155ItemIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Solar1155NFTMint1155Item represents a Mint1155Item event raised by the Solar1155NFT contract.
type Solar1155NFTMint1155Item struct {
	Operator common.Address
	To       common.Address
	TokenId  *big.Int
	Amount   *big.Int
	TokenURI string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterMint1155Item is a free log retrieval operation binding the contract event 0x19bbbbdd0f802c36b55f41676205367ff3171811a6e27b41a4cf5d355d7ce64b.
//
// Solidity: event Mint1155Item(address operator, address indexed _to, uint256 indexed _tokenId, uint256 indexed _amount, string _tokenURI)
func (_Solar1155NFT *Solar1155NFTFilterer) FilterMint1155Item(opts *bind.FilterOpts, _to []common.Address, _tokenId []*big.Int, _amount []*big.Int) (*Solar1155NFTMint1155ItemIterator, error) {

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

	logs, sub, err := _Solar1155NFT.contract.FilterLogs(opts, "Mint1155Item", _toRule, _tokenIdRule, _amountRule)
	if err != nil {
		return nil, err
	}
	return &Solar1155NFTMint1155ItemIterator{contract: _Solar1155NFT.contract, event: "Mint1155Item", logs: logs, sub: sub}, nil
}

// WatchMint1155Item is a free log subscription operation binding the contract event 0x19bbbbdd0f802c36b55f41676205367ff3171811a6e27b41a4cf5d355d7ce64b.
//
// Solidity: event Mint1155Item(address operator, address indexed _to, uint256 indexed _tokenId, uint256 indexed _amount, string _tokenURI)
func (_Solar1155NFT *Solar1155NFTFilterer) WatchMint1155Item(opts *bind.WatchOpts, sink chan<- *Solar1155NFTMint1155Item, _to []common.Address, _tokenId []*big.Int, _amount []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _Solar1155NFT.contract.WatchLogs(opts, "Mint1155Item", _toRule, _tokenIdRule, _amountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Solar1155NFTMint1155Item)
				if err := _Solar1155NFT.contract.UnpackLog(event, "Mint1155Item", log); err != nil {
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

// ParseMint1155Item is a log parse operation binding the contract event 0x19bbbbdd0f802c36b55f41676205367ff3171811a6e27b41a4cf5d355d7ce64b.
//
// Solidity: event Mint1155Item(address operator, address indexed _to, uint256 indexed _tokenId, uint256 indexed _amount, string _tokenURI)
func (_Solar1155NFT *Solar1155NFTFilterer) ParseMint1155Item(log types.Log) (*Solar1155NFTMint1155Item, error) {
	event := new(Solar1155NFTMint1155Item)
	if err := _Solar1155NFT.contract.UnpackLog(event, "Mint1155Item", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Solar1155NFTMintBatchItemsIterator is returned from FilterMintBatchItems and is used to iterate over the raw logs and unpacked data for MintBatchItems events raised by the Solar1155NFT contract.
type Solar1155NFTMintBatchItemsIterator struct {
	Event *Solar1155NFTMintBatchItems // Event containing the contract specifics and raw log

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
func (it *Solar1155NFTMintBatchItemsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Solar1155NFTMintBatchItems)
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
		it.Event = new(Solar1155NFTMintBatchItems)
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
func (it *Solar1155NFTMintBatchItemsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Solar1155NFTMintBatchItemsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Solar1155NFTMintBatchItems represents a MintBatchItems event raised by the Solar1155NFT contract.
type Solar1155NFTMintBatchItems struct {
	Operator  common.Address
	To        common.Address
	TokenIds  []*big.Int
	Amounts   []*big.Int
	TokenUris []string
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterMintBatchItems is a free log retrieval operation binding the contract event 0x67cbb50eda25c4f94b31205b994dd650951f4fa187dccd35c59bfdbc465b2998.
//
// Solidity: event MintBatchItems(address indexed operator, address indexed _to, uint256[] _tokenIds, uint256[] _amounts, string[] _tokenUris)
func (_Solar1155NFT *Solar1155NFTFilterer) FilterMintBatchItems(opts *bind.FilterOpts, operator []common.Address, _to []common.Address) (*Solar1155NFTMintBatchItemsIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _Solar1155NFT.contract.FilterLogs(opts, "MintBatchItems", operatorRule, _toRule)
	if err != nil {
		return nil, err
	}
	return &Solar1155NFTMintBatchItemsIterator{contract: _Solar1155NFT.contract, event: "MintBatchItems", logs: logs, sub: sub}, nil
}

// WatchMintBatchItems is a free log subscription operation binding the contract event 0x67cbb50eda25c4f94b31205b994dd650951f4fa187dccd35c59bfdbc465b2998.
//
// Solidity: event MintBatchItems(address indexed operator, address indexed _to, uint256[] _tokenIds, uint256[] _amounts, string[] _tokenUris)
func (_Solar1155NFT *Solar1155NFTFilterer) WatchMintBatchItems(opts *bind.WatchOpts, sink chan<- *Solar1155NFTMintBatchItems, operator []common.Address, _to []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _Solar1155NFT.contract.WatchLogs(opts, "MintBatchItems", operatorRule, _toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Solar1155NFTMintBatchItems)
				if err := _Solar1155NFT.contract.UnpackLog(event, "MintBatchItems", log); err != nil {
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
// Solidity: event MintBatchItems(address indexed operator, address indexed _to, uint256[] _tokenIds, uint256[] _amounts, string[] _tokenUris)
func (_Solar1155NFT *Solar1155NFTFilterer) ParseMintBatchItems(log types.Log) (*Solar1155NFTMintBatchItems, error) {
	event := new(Solar1155NFTMintBatchItems)
	if err := _Solar1155NFT.contract.UnpackLog(event, "MintBatchItems", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Solar1155NFTOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Solar1155NFT contract.
type Solar1155NFTOwnershipTransferredIterator struct {
	Event *Solar1155NFTOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *Solar1155NFTOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Solar1155NFTOwnershipTransferred)
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
		it.Event = new(Solar1155NFTOwnershipTransferred)
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
func (it *Solar1155NFTOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Solar1155NFTOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Solar1155NFTOwnershipTransferred represents a OwnershipTransferred event raised by the Solar1155NFT contract.
type Solar1155NFTOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Solar1155NFT *Solar1155NFTFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*Solar1155NFTOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Solar1155NFT.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &Solar1155NFTOwnershipTransferredIterator{contract: _Solar1155NFT.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Solar1155NFT *Solar1155NFTFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *Solar1155NFTOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Solar1155NFT.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Solar1155NFTOwnershipTransferred)
				if err := _Solar1155NFT.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Solar1155NFT *Solar1155NFTFilterer) ParseOwnershipTransferred(log types.Log) (*Solar1155NFTOwnershipTransferred, error) {
	event := new(Solar1155NFTOwnershipTransferred)
	if err := _Solar1155NFT.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Solar1155NFTRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Solar1155NFT contract.
type Solar1155NFTRoleAdminChangedIterator struct {
	Event *Solar1155NFTRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *Solar1155NFTRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Solar1155NFTRoleAdminChanged)
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
		it.Event = new(Solar1155NFTRoleAdminChanged)
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
func (it *Solar1155NFTRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Solar1155NFTRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Solar1155NFTRoleAdminChanged represents a RoleAdminChanged event raised by the Solar1155NFT contract.
type Solar1155NFTRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Solar1155NFT *Solar1155NFTFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*Solar1155NFTRoleAdminChangedIterator, error) {

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

	logs, sub, err := _Solar1155NFT.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &Solar1155NFTRoleAdminChangedIterator{contract: _Solar1155NFT.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Solar1155NFT *Solar1155NFTFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *Solar1155NFTRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _Solar1155NFT.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Solar1155NFTRoleAdminChanged)
				if err := _Solar1155NFT.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_Solar1155NFT *Solar1155NFTFilterer) ParseRoleAdminChanged(log types.Log) (*Solar1155NFTRoleAdminChanged, error) {
	event := new(Solar1155NFTRoleAdminChanged)
	if err := _Solar1155NFT.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Solar1155NFTRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Solar1155NFT contract.
type Solar1155NFTRoleGrantedIterator struct {
	Event *Solar1155NFTRoleGranted // Event containing the contract specifics and raw log

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
func (it *Solar1155NFTRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Solar1155NFTRoleGranted)
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
		it.Event = new(Solar1155NFTRoleGranted)
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
func (it *Solar1155NFTRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Solar1155NFTRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Solar1155NFTRoleGranted represents a RoleGranted event raised by the Solar1155NFT contract.
type Solar1155NFTRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Solar1155NFT *Solar1155NFTFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*Solar1155NFTRoleGrantedIterator, error) {

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

	logs, sub, err := _Solar1155NFT.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &Solar1155NFTRoleGrantedIterator{contract: _Solar1155NFT.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Solar1155NFT *Solar1155NFTFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *Solar1155NFTRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Solar1155NFT.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Solar1155NFTRoleGranted)
				if err := _Solar1155NFT.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_Solar1155NFT *Solar1155NFTFilterer) ParseRoleGranted(log types.Log) (*Solar1155NFTRoleGranted, error) {
	event := new(Solar1155NFTRoleGranted)
	if err := _Solar1155NFT.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Solar1155NFTRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Solar1155NFT contract.
type Solar1155NFTRoleRevokedIterator struct {
	Event *Solar1155NFTRoleRevoked // Event containing the contract specifics and raw log

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
func (it *Solar1155NFTRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Solar1155NFTRoleRevoked)
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
		it.Event = new(Solar1155NFTRoleRevoked)
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
func (it *Solar1155NFTRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Solar1155NFTRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Solar1155NFTRoleRevoked represents a RoleRevoked event raised by the Solar1155NFT contract.
type Solar1155NFTRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Solar1155NFT *Solar1155NFTFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*Solar1155NFTRoleRevokedIterator, error) {

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

	logs, sub, err := _Solar1155NFT.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &Solar1155NFTRoleRevokedIterator{contract: _Solar1155NFT.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Solar1155NFT *Solar1155NFTFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *Solar1155NFTRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Solar1155NFT.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Solar1155NFTRoleRevoked)
				if err := _Solar1155NFT.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_Solar1155NFT *Solar1155NFTFilterer) ParseRoleRevoked(log types.Log) (*Solar1155NFTRoleRevoked, error) {
	event := new(Solar1155NFTRoleRevoked)
	if err := _Solar1155NFT.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Solar1155NFTTradeNftIterator is returned from FilterTradeNft and is used to iterate over the raw logs and unpacked data for TradeNft events raised by the Solar1155NFT contract.
type Solar1155NFTTradeNftIterator struct {
	Event *Solar1155NFTTradeNft // Event containing the contract specifics and raw log

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
func (it *Solar1155NFTTradeNftIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Solar1155NFTTradeNft)
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
		it.Event = new(Solar1155NFTTradeNft)
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
func (it *Solar1155NFTTradeNftIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Solar1155NFTTradeNftIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Solar1155NFTTradeNft represents a TradeNft event raised by the Solar1155NFT contract.
type Solar1155NFTTradeNft struct {
	Buyer   common.Address
	Seller  common.Address
	TokenId *big.Int
	Amount  *big.Int
	TradeId *big.Int
	OrderId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTradeNft is a free log retrieval operation binding the contract event 0x2b5a7ee0b1c788b20b39eed13b1509332c783cf171e537256cea53ed52b29dd1.
//
// Solidity: event TradeNft(address indexed _buyer, address indexed _seller, uint256 indexed _tokenId, uint256 amount, uint256 _tradeId, uint256 _orderId)
func (_Solar1155NFT *Solar1155NFTFilterer) FilterTradeNft(opts *bind.FilterOpts, _buyer []common.Address, _seller []common.Address, _tokenId []*big.Int) (*Solar1155NFTTradeNftIterator, error) {

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

	logs, sub, err := _Solar1155NFT.contract.FilterLogs(opts, "TradeNft", _buyerRule, _sellerRule, _tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &Solar1155NFTTradeNftIterator{contract: _Solar1155NFT.contract, event: "TradeNft", logs: logs, sub: sub}, nil
}

// WatchTradeNft is a free log subscription operation binding the contract event 0x2b5a7ee0b1c788b20b39eed13b1509332c783cf171e537256cea53ed52b29dd1.
//
// Solidity: event TradeNft(address indexed _buyer, address indexed _seller, uint256 indexed _tokenId, uint256 amount, uint256 _tradeId, uint256 _orderId)
func (_Solar1155NFT *Solar1155NFTFilterer) WatchTradeNft(opts *bind.WatchOpts, sink chan<- *Solar1155NFTTradeNft, _buyer []common.Address, _seller []common.Address, _tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _Solar1155NFT.contract.WatchLogs(opts, "TradeNft", _buyerRule, _sellerRule, _tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Solar1155NFTTradeNft)
				if err := _Solar1155NFT.contract.UnpackLog(event, "TradeNft", log); err != nil {
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

// ParseTradeNft is a log parse operation binding the contract event 0x2b5a7ee0b1c788b20b39eed13b1509332c783cf171e537256cea53ed52b29dd1.
//
// Solidity: event TradeNft(address indexed _buyer, address indexed _seller, uint256 indexed _tokenId, uint256 amount, uint256 _tradeId, uint256 _orderId)
func (_Solar1155NFT *Solar1155NFTFilterer) ParseTradeNft(log types.Log) (*Solar1155NFTTradeNft, error) {
	event := new(Solar1155NFTTradeNft)
	if err := _Solar1155NFT.contract.UnpackLog(event, "TradeNft", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Solar1155NFTTransferBatchIterator is returned from FilterTransferBatch and is used to iterate over the raw logs and unpacked data for TransferBatch events raised by the Solar1155NFT contract.
type Solar1155NFTTransferBatchIterator struct {
	Event *Solar1155NFTTransferBatch // Event containing the contract specifics and raw log

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
func (it *Solar1155NFTTransferBatchIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Solar1155NFTTransferBatch)
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
		it.Event = new(Solar1155NFTTransferBatch)
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
func (it *Solar1155NFTTransferBatchIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Solar1155NFTTransferBatchIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Solar1155NFTTransferBatch represents a TransferBatch event raised by the Solar1155NFT contract.
type Solar1155NFTTransferBatch struct {
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
func (_Solar1155NFT *Solar1155NFTFilterer) FilterTransferBatch(opts *bind.FilterOpts, operator []common.Address, from []common.Address, to []common.Address) (*Solar1155NFTTransferBatchIterator, error) {

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

	logs, sub, err := _Solar1155NFT.contract.FilterLogs(opts, "TransferBatch", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &Solar1155NFTTransferBatchIterator{contract: _Solar1155NFT.contract, event: "TransferBatch", logs: logs, sub: sub}, nil
}

// WatchTransferBatch is a free log subscription operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_Solar1155NFT *Solar1155NFTFilterer) WatchTransferBatch(opts *bind.WatchOpts, sink chan<- *Solar1155NFTTransferBatch, operator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Solar1155NFT.contract.WatchLogs(opts, "TransferBatch", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Solar1155NFTTransferBatch)
				if err := _Solar1155NFT.contract.UnpackLog(event, "TransferBatch", log); err != nil {
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
func (_Solar1155NFT *Solar1155NFTFilterer) ParseTransferBatch(log types.Log) (*Solar1155NFTTransferBatch, error) {
	event := new(Solar1155NFTTransferBatch)
	if err := _Solar1155NFT.contract.UnpackLog(event, "TransferBatch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Solar1155NFTTransferSingleIterator is returned from FilterTransferSingle and is used to iterate over the raw logs and unpacked data for TransferSingle events raised by the Solar1155NFT contract.
type Solar1155NFTTransferSingleIterator struct {
	Event *Solar1155NFTTransferSingle // Event containing the contract specifics and raw log

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
func (it *Solar1155NFTTransferSingleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Solar1155NFTTransferSingle)
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
		it.Event = new(Solar1155NFTTransferSingle)
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
func (it *Solar1155NFTTransferSingleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Solar1155NFTTransferSingleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Solar1155NFTTransferSingle represents a TransferSingle event raised by the Solar1155NFT contract.
type Solar1155NFTTransferSingle struct {
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
func (_Solar1155NFT *Solar1155NFTFilterer) FilterTransferSingle(opts *bind.FilterOpts, operator []common.Address, from []common.Address, to []common.Address) (*Solar1155NFTTransferSingleIterator, error) {

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

	logs, sub, err := _Solar1155NFT.contract.FilterLogs(opts, "TransferSingle", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &Solar1155NFTTransferSingleIterator{contract: _Solar1155NFT.contract, event: "TransferSingle", logs: logs, sub: sub}, nil
}

// WatchTransferSingle is a free log subscription operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_Solar1155NFT *Solar1155NFTFilterer) WatchTransferSingle(opts *bind.WatchOpts, sink chan<- *Solar1155NFTTransferSingle, operator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Solar1155NFT.contract.WatchLogs(opts, "TransferSingle", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Solar1155NFTTransferSingle)
				if err := _Solar1155NFT.contract.UnpackLog(event, "TransferSingle", log); err != nil {
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
func (_Solar1155NFT *Solar1155NFTFilterer) ParseTransferSingle(log types.Log) (*Solar1155NFTTransferSingle, error) {
	event := new(Solar1155NFTTransferSingle)
	if err := _Solar1155NFT.contract.UnpackLog(event, "TransferSingle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Solar1155NFTURIIterator is returned from FilterURI and is used to iterate over the raw logs and unpacked data for URI events raised by the Solar1155NFT contract.
type Solar1155NFTURIIterator struct {
	Event *Solar1155NFTURI // Event containing the contract specifics and raw log

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
func (it *Solar1155NFTURIIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Solar1155NFTURI)
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
		it.Event = new(Solar1155NFTURI)
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
func (it *Solar1155NFTURIIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Solar1155NFTURIIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Solar1155NFTURI represents a URI event raised by the Solar1155NFT contract.
type Solar1155NFTURI struct {
	Value string
	Id    *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterURI is a free log retrieval operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_Solar1155NFT *Solar1155NFTFilterer) FilterURI(opts *bind.FilterOpts, id []*big.Int) (*Solar1155NFTURIIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Solar1155NFT.contract.FilterLogs(opts, "URI", idRule)
	if err != nil {
		return nil, err
	}
	return &Solar1155NFTURIIterator{contract: _Solar1155NFT.contract, event: "URI", logs: logs, sub: sub}, nil
}

// WatchURI is a free log subscription operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_Solar1155NFT *Solar1155NFTFilterer) WatchURI(opts *bind.WatchOpts, sink chan<- *Solar1155NFTURI, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Solar1155NFT.contract.WatchLogs(opts, "URI", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Solar1155NFTURI)
				if err := _Solar1155NFT.contract.UnpackLog(event, "URI", log); err != nil {
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
func (_Solar1155NFT *Solar1155NFTFilterer) ParseURI(log types.Log) (*Solar1155NFTURI, error) {
	event := new(Solar1155NFTURI)
	if err := _Solar1155NFT.contract.UnpackLog(event, "URI", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Solar1155NFTUpdateOrderIterator is returned from FilterUpdateOrder and is used to iterate over the raw logs and unpacked data for UpdateOrder events raised by the Solar1155NFT contract.
type Solar1155NFTUpdateOrderIterator struct {
	Event *Solar1155NFTUpdateOrder // Event containing the contract specifics and raw log

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
func (it *Solar1155NFTUpdateOrderIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Solar1155NFTUpdateOrder)
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
		it.Event = new(Solar1155NFTUpdateOrder)
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
func (it *Solar1155NFTUpdateOrderIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Solar1155NFTUpdateOrderIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Solar1155NFTUpdateOrder represents a UpdateOrder event raised by the Solar1155NFT contract.
type Solar1155NFTUpdateOrder struct {
	Operator   common.Address
	Seller     common.Address
	TokenId    *big.Int
	Amount     *big.Int
	OrderId    *big.Int
	UnitPrice  *big.Int
	FeePercent uint16
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterUpdateOrder is a free log retrieval operation binding the contract event 0x57f3ae470c83e1b112c2537ac7f869af4dc7bc7341b8acfa03b969efd3c576a2.
//
// Solidity: event UpdateOrder(address _operator, address indexed _seller, uint256 indexed _tokenId, uint256 _amount, uint256 indexed _orderId, uint256 _unitPrice, uint16 _feePercent)
func (_Solar1155NFT *Solar1155NFTFilterer) FilterUpdateOrder(opts *bind.FilterOpts, _seller []common.Address, _tokenId []*big.Int, _orderId []*big.Int) (*Solar1155NFTUpdateOrderIterator, error) {

	var _sellerRule []interface{}
	for _, _sellerItem := range _seller {
		_sellerRule = append(_sellerRule, _sellerItem)
	}
	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}

	var _orderIdRule []interface{}
	for _, _orderIdItem := range _orderId {
		_orderIdRule = append(_orderIdRule, _orderIdItem)
	}

	logs, sub, err := _Solar1155NFT.contract.FilterLogs(opts, "UpdateOrder", _sellerRule, _tokenIdRule, _orderIdRule)
	if err != nil {
		return nil, err
	}
	return &Solar1155NFTUpdateOrderIterator{contract: _Solar1155NFT.contract, event: "UpdateOrder", logs: logs, sub: sub}, nil
}

// WatchUpdateOrder is a free log subscription operation binding the contract event 0x57f3ae470c83e1b112c2537ac7f869af4dc7bc7341b8acfa03b969efd3c576a2.
//
// Solidity: event UpdateOrder(address _operator, address indexed _seller, uint256 indexed _tokenId, uint256 _amount, uint256 indexed _orderId, uint256 _unitPrice, uint16 _feePercent)
func (_Solar1155NFT *Solar1155NFTFilterer) WatchUpdateOrder(opts *bind.WatchOpts, sink chan<- *Solar1155NFTUpdateOrder, _seller []common.Address, _tokenId []*big.Int, _orderId []*big.Int) (event.Subscription, error) {

	var _sellerRule []interface{}
	for _, _sellerItem := range _seller {
		_sellerRule = append(_sellerRule, _sellerItem)
	}
	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}

	var _orderIdRule []interface{}
	for _, _orderIdItem := range _orderId {
		_orderIdRule = append(_orderIdRule, _orderIdItem)
	}

	logs, sub, err := _Solar1155NFT.contract.WatchLogs(opts, "UpdateOrder", _sellerRule, _tokenIdRule, _orderIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Solar1155NFTUpdateOrder)
				if err := _Solar1155NFT.contract.UnpackLog(event, "UpdateOrder", log); err != nil {
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

// ParseUpdateOrder is a log parse operation binding the contract event 0x57f3ae470c83e1b112c2537ac7f869af4dc7bc7341b8acfa03b969efd3c576a2.
//
// Solidity: event UpdateOrder(address _operator, address indexed _seller, uint256 indexed _tokenId, uint256 _amount, uint256 indexed _orderId, uint256 _unitPrice, uint16 _feePercent)
func (_Solar1155NFT *Solar1155NFTFilterer) ParseUpdateOrder(log types.Log) (*Solar1155NFTUpdateOrder, error) {
	event := new(Solar1155NFTUpdateOrder)
	if err := _Solar1155NFT.contract.UnpackLog(event, "UpdateOrder", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
