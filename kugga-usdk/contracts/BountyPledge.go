// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bountyPledge

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

// BountyPledgeDataUpgradeableAssetCoin is an auto generated low-level Go binding around an user-defined struct.
type BountyPledgeDataUpgradeableAssetCoin struct {
	Asset   common.Address
	Symbol  string
	Id      uint8
	Enabled bool
}

// BountyPledgeDataUpgradeableBounty is an auto generated low-level Go binding around an user-defined struct.
type BountyPledgeDataUpgradeableBounty struct {
	StartTime  *big.Int
	EndTime    *big.Int
	Owner      common.Address
	BountyType uint8
	Cid        string
	Reward     BountyPledgeDataUpgradeableReward
	Status     uint8
}

// BountyPledgeDataUpgradeablePledgeRequest is an auto generated low-level Go binding around an user-defined struct.
type BountyPledgeDataUpgradeablePledgeRequest struct {
	StartTime  *big.Int
	EndTime    *big.Int
	BountyType uint8
	Cid        string
	Reward     *big.Int
	RewardUnit string
}

// BountyPledgeDataUpgradeableReward is an auto generated low-level Go binding around an user-defined struct.
type BountyPledgeDataUpgradeableReward struct {
	Amount  *big.Int
	Balance *big.Int
	Unit    string
}

// BountyPledgeMetaData contains all meta data concerning the BountyPledge contract.
var BountyPledgeMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"Id\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"model\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"AssetActivity\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"keccak256CID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"bountyType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"status\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"rewardUnit\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"cid\",\"type\":\"string\"}],\"name\":\"BountyActivity\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"keccak256CID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"rewardUnit\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"model\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"status\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"cid\",\"type\":\"string\"}],\"name\":\"RewardActivity\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"keccak256BountyCID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"status\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"rewardUnit\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"bountyCid\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"contributeCid\",\"type\":\"string\"}],\"name\":\"TransferActivity\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"keccak256BountyCID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"status\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"rewardUnit\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"bountyCid\",\"type\":\"string\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"CHILD_CHAIN_ID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CHILD_CHAIN_ID_BYTES\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ERC712_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ROOT_CHAIN_ID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ROOT_CHAIN_ID_BYTES\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset_\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"symbol_\",\"type\":\"string\"}],\"name\":\"addSupporteAsset\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset_\",\"type\":\"address\"}],\"name\":\"disableAsset\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"disableWorldStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset_\",\"type\":\"address\"}],\"name\":\"enableAsset\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"enableWorldStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"cid\",\"type\":\"string\"}],\"name\":\"getBounty\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"bountyType\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"cid\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"unit\",\"type\":\"string\"}],\"internalType\":\"structBountyPledgeDataUpgradeable.Reward\",\"name\":\"reward\",\"type\":\"tuple\"},{\"internalType\":\"uint8\",\"name\":\"status\",\"type\":\"uint8\"}],\"internalType\":\"structBountyPledgeDataUpgradeable.Bounty\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset_\",\"type\":\"address\"}],\"name\":\"getSupportedAsset\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"id\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"internalType\":\"structBountyPledgeDataUpgradeable.AssetCoin\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"symbol_\",\"type\":\"string\"}],\"name\":\"getSupportedAssetBySymbol\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"id\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"internalType\":\"structBountyPledgeDataUpgradeable.AssetCoin\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"cid\",\"type\":\"string\"}],\"name\":\"getWinners\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getWorldStatus\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"cidUnit\",\"type\":\"string\"}],\"name\":\"getkeccak256\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol_\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"usdt_\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"coinUnit_\",\"type\":\"string\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"bountyType\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"cid\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"rewardUnit\",\"type\":\"string\"}],\"internalType\":\"structBountyPledgeDataUpgradeable.PledgeRequest\",\"name\":\"req\",\"type\":\"tuple\"}],\"name\":\"pledge\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"startTime_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime_\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"bountyType_\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"cid_\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"reward_\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"rewardUnit_\",\"type\":\"string\"}],\"name\":\"pledge2\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"cid\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"deadLine\",\"type\":\"uint256\"}],\"name\":\"setBountyExpired\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"}],\"name\":\"setName\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"symbol_\",\"type\":\"string\"}],\"name\":\"setSymbol\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"bountyCid\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"contributeCid\",\"type\":\"string\"}],\"name\":\"transfer\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"bountyCid\",\"type\":\"string\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"cid\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"increaseReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"cid\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approveReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// BountyPledgeABI is the input ABI used to generate the binding from.
// Deprecated: Use BountyPledgeMetaData.ABI instead.
var BountyPledgeABI = BountyPledgeMetaData.ABI

// BountyPledge is an auto generated Go binding around an Ethereum contract.
type BountyPledge struct {
	BountyPledgeCaller     // Read-only binding to the contract
	BountyPledgeTransactor // Write-only binding to the contract
	BountyPledgeFilterer   // Log filterer for contract events
}

// BountyPledgeCaller is an auto generated read-only Go binding around an Ethereum contract.
type BountyPledgeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BountyPledgeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BountyPledgeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BountyPledgeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BountyPledgeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BountyPledgeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BountyPledgeSession struct {
	Contract     *BountyPledge     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BountyPledgeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BountyPledgeCallerSession struct {
	Contract *BountyPledgeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// BountyPledgeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BountyPledgeTransactorSession struct {
	Contract     *BountyPledgeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// BountyPledgeRaw is an auto generated low-level Go binding around an Ethereum contract.
type BountyPledgeRaw struct {
	Contract *BountyPledge // Generic contract binding to access the raw methods on
}

// BountyPledgeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BountyPledgeCallerRaw struct {
	Contract *BountyPledgeCaller // Generic read-only contract binding to access the raw methods on
}

// BountyPledgeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BountyPledgeTransactorRaw struct {
	Contract *BountyPledgeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBountyPledge creates a new instance of BountyPledge, bound to a specific deployed contract.
func NewBountyPledge(address common.Address, backend bind.ContractBackend) (*BountyPledge, error) {
	contract, err := bindBountyPledge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BountyPledge{BountyPledgeCaller: BountyPledgeCaller{contract: contract}, BountyPledgeTransactor: BountyPledgeTransactor{contract: contract}, BountyPledgeFilterer: BountyPledgeFilterer{contract: contract}}, nil
}

// NewBountyPledgeCaller creates a new read-only instance of BountyPledge, bound to a specific deployed contract.
func NewBountyPledgeCaller(address common.Address, caller bind.ContractCaller) (*BountyPledgeCaller, error) {
	contract, err := bindBountyPledge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BountyPledgeCaller{contract: contract}, nil
}

// NewBountyPledgeTransactor creates a new write-only instance of BountyPledge, bound to a specific deployed contract.
func NewBountyPledgeTransactor(address common.Address, transactor bind.ContractTransactor) (*BountyPledgeTransactor, error) {
	contract, err := bindBountyPledge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BountyPledgeTransactor{contract: contract}, nil
}

// NewBountyPledgeFilterer creates a new log filterer instance of BountyPledge, bound to a specific deployed contract.
func NewBountyPledgeFilterer(address common.Address, filterer bind.ContractFilterer) (*BountyPledgeFilterer, error) {
	contract, err := bindBountyPledge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BountyPledgeFilterer{contract: contract}, nil
}

// bindBountyPledge binds a generic wrapper to an already deployed contract.
func bindBountyPledge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BountyPledgeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BountyPledge *BountyPledgeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BountyPledge.Contract.BountyPledgeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BountyPledge *BountyPledgeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BountyPledge.Contract.BountyPledgeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BountyPledge *BountyPledgeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BountyPledge.Contract.BountyPledgeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BountyPledge *BountyPledgeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BountyPledge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BountyPledge *BountyPledgeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BountyPledge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BountyPledge *BountyPledgeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BountyPledge.Contract.contract.Transact(opts, method, params...)
}

// CHILDCHAINID is a free data retrieval call binding the contract method 0x626381a0.
//
// Solidity: function CHILD_CHAIN_ID() view returns(uint256)
func (_BountyPledge *BountyPledgeCaller) CHILDCHAINID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BountyPledge.contract.Call(opts, &out, "CHILD_CHAIN_ID")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CHILDCHAINID is a free data retrieval call binding the contract method 0x626381a0.
//
// Solidity: function CHILD_CHAIN_ID() view returns(uint256)
func (_BountyPledge *BountyPledgeSession) CHILDCHAINID() (*big.Int, error) {
	return _BountyPledge.Contract.CHILDCHAINID(&_BountyPledge.CallOpts)
}

// CHILDCHAINID is a free data retrieval call binding the contract method 0x626381a0.
//
// Solidity: function CHILD_CHAIN_ID() view returns(uint256)
func (_BountyPledge *BountyPledgeCallerSession) CHILDCHAINID() (*big.Int, error) {
	return _BountyPledge.Contract.CHILDCHAINID(&_BountyPledge.CallOpts)
}

// CHILDCHAINIDBYTES is a free data retrieval call binding the contract method 0x0b54817c.
//
// Solidity: function CHILD_CHAIN_ID_BYTES() view returns(bytes)
func (_BountyPledge *BountyPledgeCaller) CHILDCHAINIDBYTES(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _BountyPledge.contract.Call(opts, &out, "CHILD_CHAIN_ID_BYTES")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// CHILDCHAINIDBYTES is a free data retrieval call binding the contract method 0x0b54817c.
//
// Solidity: function CHILD_CHAIN_ID_BYTES() view returns(bytes)
func (_BountyPledge *BountyPledgeSession) CHILDCHAINIDBYTES() ([]byte, error) {
	return _BountyPledge.Contract.CHILDCHAINIDBYTES(&_BountyPledge.CallOpts)
}

// CHILDCHAINIDBYTES is a free data retrieval call binding the contract method 0x0b54817c.
//
// Solidity: function CHILD_CHAIN_ID_BYTES() view returns(bytes)
func (_BountyPledge *BountyPledgeCallerSession) CHILDCHAINIDBYTES() ([]byte, error) {
	return _BountyPledge.Contract.CHILDCHAINIDBYTES(&_BountyPledge.CallOpts)
}

// ERC712VERSION is a free data retrieval call binding the contract method 0x0f7e5970.
//
// Solidity: function ERC712_VERSION() view returns(string)
func (_BountyPledge *BountyPledgeCaller) ERC712VERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BountyPledge.contract.Call(opts, &out, "ERC712_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ERC712VERSION is a free data retrieval call binding the contract method 0x0f7e5970.
//
// Solidity: function ERC712_VERSION() view returns(string)
func (_BountyPledge *BountyPledgeSession) ERC712VERSION() (string, error) {
	return _BountyPledge.Contract.ERC712VERSION(&_BountyPledge.CallOpts)
}

// ERC712VERSION is a free data retrieval call binding the contract method 0x0f7e5970.
//
// Solidity: function ERC712_VERSION() view returns(string)
func (_BountyPledge *BountyPledgeCallerSession) ERC712VERSION() (string, error) {
	return _BountyPledge.Contract.ERC712VERSION(&_BountyPledge.CallOpts)
}

// ROOTCHAINID is a free data retrieval call binding the contract method 0x8acfcaf7.
//
// Solidity: function ROOT_CHAIN_ID() view returns(uint256)
func (_BountyPledge *BountyPledgeCaller) ROOTCHAINID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BountyPledge.contract.Call(opts, &out, "ROOT_CHAIN_ID")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ROOTCHAINID is a free data retrieval call binding the contract method 0x8acfcaf7.
//
// Solidity: function ROOT_CHAIN_ID() view returns(uint256)
func (_BountyPledge *BountyPledgeSession) ROOTCHAINID() (*big.Int, error) {
	return _BountyPledge.Contract.ROOTCHAINID(&_BountyPledge.CallOpts)
}

// ROOTCHAINID is a free data retrieval call binding the contract method 0x8acfcaf7.
//
// Solidity: function ROOT_CHAIN_ID() view returns(uint256)
func (_BountyPledge *BountyPledgeCallerSession) ROOTCHAINID() (*big.Int, error) {
	return _BountyPledge.Contract.ROOTCHAINID(&_BountyPledge.CallOpts)
}

// ROOTCHAINIDBYTES is a free data retrieval call binding the contract method 0x0dd7531a.
//
// Solidity: function ROOT_CHAIN_ID_BYTES() view returns(bytes)
func (_BountyPledge *BountyPledgeCaller) ROOTCHAINIDBYTES(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _BountyPledge.contract.Call(opts, &out, "ROOT_CHAIN_ID_BYTES")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// ROOTCHAINIDBYTES is a free data retrieval call binding the contract method 0x0dd7531a.
//
// Solidity: function ROOT_CHAIN_ID_BYTES() view returns(bytes)
func (_BountyPledge *BountyPledgeSession) ROOTCHAINIDBYTES() ([]byte, error) {
	return _BountyPledge.Contract.ROOTCHAINIDBYTES(&_BountyPledge.CallOpts)
}

// ROOTCHAINIDBYTES is a free data retrieval call binding the contract method 0x0dd7531a.
//
// Solidity: function ROOT_CHAIN_ID_BYTES() view returns(bytes)
func (_BountyPledge *BountyPledgeCallerSession) ROOTCHAINIDBYTES() ([]byte, error) {
	return _BountyPledge.Contract.ROOTCHAINIDBYTES(&_BountyPledge.CallOpts)
}

// GetBounty is a free data retrieval call binding the contract method 0x57f56061.
//
// Solidity: function getBounty(string cid) view returns((uint256,uint256,address,uint8,string,(uint256,uint256,string),uint8))
func (_BountyPledge *BountyPledgeCaller) GetBounty(opts *bind.CallOpts, cid string) (BountyPledgeDataUpgradeableBounty, error) {
	var out []interface{}
	err := _BountyPledge.contract.Call(opts, &out, "getBounty", cid)

	if err != nil {
		return *new(BountyPledgeDataUpgradeableBounty), err
	}

	out0 := *abi.ConvertType(out[0], new(BountyPledgeDataUpgradeableBounty)).(*BountyPledgeDataUpgradeableBounty)

	return out0, err

}

// GetBounty is a free data retrieval call binding the contract method 0x57f56061.
//
// Solidity: function getBounty(string cid) view returns((uint256,uint256,address,uint8,string,(uint256,uint256,string),uint8))
func (_BountyPledge *BountyPledgeSession) GetBounty(cid string) (BountyPledgeDataUpgradeableBounty, error) {
	return _BountyPledge.Contract.GetBounty(&_BountyPledge.CallOpts, cid)
}

// GetBounty is a free data retrieval call binding the contract method 0x57f56061.
//
// Solidity: function getBounty(string cid) view returns((uint256,uint256,address,uint8,string,(uint256,uint256,string),uint8))
func (_BountyPledge *BountyPledgeCallerSession) GetBounty(cid string) (BountyPledgeDataUpgradeableBounty, error) {
	return _BountyPledge.Contract.GetBounty(&_BountyPledge.CallOpts, cid)
}

// GetSupportedAsset is a free data retrieval call binding the contract method 0x3bc430a5.
//
// Solidity: function getSupportedAsset(address asset_) view returns((address,string,uint8,bool))
func (_BountyPledge *BountyPledgeCaller) GetSupportedAsset(opts *bind.CallOpts, asset_ common.Address) (BountyPledgeDataUpgradeableAssetCoin, error) {
	var out []interface{}
	err := _BountyPledge.contract.Call(opts, &out, "getSupportedAsset", asset_)

	if err != nil {
		return *new(BountyPledgeDataUpgradeableAssetCoin), err
	}

	out0 := *abi.ConvertType(out[0], new(BountyPledgeDataUpgradeableAssetCoin)).(*BountyPledgeDataUpgradeableAssetCoin)

	return out0, err

}

// GetSupportedAsset is a free data retrieval call binding the contract method 0x3bc430a5.
//
// Solidity: function getSupportedAsset(address asset_) view returns((address,string,uint8,bool))
func (_BountyPledge *BountyPledgeSession) GetSupportedAsset(asset_ common.Address) (BountyPledgeDataUpgradeableAssetCoin, error) {
	return _BountyPledge.Contract.GetSupportedAsset(&_BountyPledge.CallOpts, asset_)
}

// GetSupportedAsset is a free data retrieval call binding the contract method 0x3bc430a5.
//
// Solidity: function getSupportedAsset(address asset_) view returns((address,string,uint8,bool))
func (_BountyPledge *BountyPledgeCallerSession) GetSupportedAsset(asset_ common.Address) (BountyPledgeDataUpgradeableAssetCoin, error) {
	return _BountyPledge.Contract.GetSupportedAsset(&_BountyPledge.CallOpts, asset_)
}

// GetSupportedAssetBySymbol is a free data retrieval call binding the contract method 0x42151337.
//
// Solidity: function getSupportedAssetBySymbol(string symbol_) view returns((address,string,uint8,bool))
func (_BountyPledge *BountyPledgeCaller) GetSupportedAssetBySymbol(opts *bind.CallOpts, symbol_ string) (BountyPledgeDataUpgradeableAssetCoin, error) {
	var out []interface{}
	err := _BountyPledge.contract.Call(opts, &out, "getSupportedAssetBySymbol", symbol_)

	if err != nil {
		return *new(BountyPledgeDataUpgradeableAssetCoin), err
	}

	out0 := *abi.ConvertType(out[0], new(BountyPledgeDataUpgradeableAssetCoin)).(*BountyPledgeDataUpgradeableAssetCoin)

	return out0, err

}

// GetSupportedAssetBySymbol is a free data retrieval call binding the contract method 0x42151337.
//
// Solidity: function getSupportedAssetBySymbol(string symbol_) view returns((address,string,uint8,bool))
func (_BountyPledge *BountyPledgeSession) GetSupportedAssetBySymbol(symbol_ string) (BountyPledgeDataUpgradeableAssetCoin, error) {
	return _BountyPledge.Contract.GetSupportedAssetBySymbol(&_BountyPledge.CallOpts, symbol_)
}

// GetSupportedAssetBySymbol is a free data retrieval call binding the contract method 0x42151337.
//
// Solidity: function getSupportedAssetBySymbol(string symbol_) view returns((address,string,uint8,bool))
func (_BountyPledge *BountyPledgeCallerSession) GetSupportedAssetBySymbol(symbol_ string) (BountyPledgeDataUpgradeableAssetCoin, error) {
	return _BountyPledge.Contract.GetSupportedAssetBySymbol(&_BountyPledge.CallOpts, symbol_)
}

// GetWinners is a free data retrieval call binding the contract method 0x95a4c525.
//
// Solidity: function getWinners(string cid) view returns(string[])
func (_BountyPledge *BountyPledgeCaller) GetWinners(opts *bind.CallOpts, cid string) ([]string, error) {
	var out []interface{}
	err := _BountyPledge.contract.Call(opts, &out, "getWinners", cid)

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// GetWinners is a free data retrieval call binding the contract method 0x95a4c525.
//
// Solidity: function getWinners(string cid) view returns(string[])
func (_BountyPledge *BountyPledgeSession) GetWinners(cid string) ([]string, error) {
	return _BountyPledge.Contract.GetWinners(&_BountyPledge.CallOpts, cid)
}

// GetWinners is a free data retrieval call binding the contract method 0x95a4c525.
//
// Solidity: function getWinners(string cid) view returns(string[])
func (_BountyPledge *BountyPledgeCallerSession) GetWinners(cid string) ([]string, error) {
	return _BountyPledge.Contract.GetWinners(&_BountyPledge.CallOpts, cid)
}

// GetWorldStatus is a free data retrieval call binding the contract method 0x3169953b.
//
// Solidity: function getWorldStatus() view returns(bool)
func (_BountyPledge *BountyPledgeCaller) GetWorldStatus(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _BountyPledge.contract.Call(opts, &out, "getWorldStatus")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetWorldStatus is a free data retrieval call binding the contract method 0x3169953b.
//
// Solidity: function getWorldStatus() view returns(bool)
func (_BountyPledge *BountyPledgeSession) GetWorldStatus() (bool, error) {
	return _BountyPledge.Contract.GetWorldStatus(&_BountyPledge.CallOpts)
}

// GetWorldStatus is a free data retrieval call binding the contract method 0x3169953b.
//
// Solidity: function getWorldStatus() view returns(bool)
func (_BountyPledge *BountyPledgeCallerSession) GetWorldStatus() (bool, error) {
	return _BountyPledge.Contract.GetWorldStatus(&_BountyPledge.CallOpts)
}

// Getkeccak256 is a free data retrieval call binding the contract method 0xc0f50f68.
//
// Solidity: function getkeccak256(string cidUnit) pure returns(bytes32)
func (_BountyPledge *BountyPledgeCaller) Getkeccak256(opts *bind.CallOpts, cidUnit string) ([32]byte, error) {
	var out []interface{}
	err := _BountyPledge.contract.Call(opts, &out, "getkeccak256", cidUnit)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Getkeccak256 is a free data retrieval call binding the contract method 0xc0f50f68.
//
// Solidity: function getkeccak256(string cidUnit) pure returns(bytes32)
func (_BountyPledge *BountyPledgeSession) Getkeccak256(cidUnit string) ([32]byte, error) {
	return _BountyPledge.Contract.Getkeccak256(&_BountyPledge.CallOpts, cidUnit)
}

// Getkeccak256 is a free data retrieval call binding the contract method 0xc0f50f68.
//
// Solidity: function getkeccak256(string cidUnit) pure returns(bytes32)
func (_BountyPledge *BountyPledgeCallerSession) Getkeccak256(cidUnit string) ([32]byte, error) {
	return _BountyPledge.Contract.Getkeccak256(&_BountyPledge.CallOpts, cidUnit)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_BountyPledge *BountyPledgeCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BountyPledge.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_BountyPledge *BountyPledgeSession) Name() (string, error) {
	return _BountyPledge.Contract.Name(&_BountyPledge.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_BountyPledge *BountyPledgeCallerSession) Name() (string, error) {
	return _BountyPledge.Contract.Name(&_BountyPledge.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BountyPledge *BountyPledgeCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BountyPledge.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BountyPledge *BountyPledgeSession) Owner() (common.Address, error) {
	return _BountyPledge.Contract.Owner(&_BountyPledge.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BountyPledge *BountyPledgeCallerSession) Owner() (common.Address, error) {
	return _BountyPledge.Contract.Owner(&_BountyPledge.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_BountyPledge *BountyPledgeCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BountyPledge.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_BountyPledge *BountyPledgeSession) Symbol() (string, error) {
	return _BountyPledge.Contract.Symbol(&_BountyPledge.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_BountyPledge *BountyPledgeCallerSession) Symbol() (string, error) {
	return _BountyPledge.Contract.Symbol(&_BountyPledge.CallOpts)
}

// AddSupporteAsset is a paid mutator transaction binding the contract method 0xfbbe2f31.
//
// Solidity: function addSupporteAsset(address asset_, string symbol_) returns()
func (_BountyPledge *BountyPledgeTransactor) AddSupporteAsset(opts *bind.TransactOpts, asset_ common.Address, symbol_ string) (*types.Transaction, error) {
	return _BountyPledge.contract.Transact(opts, "addSupporteAsset", asset_, symbol_)
}

// AddSupporteAsset is a paid mutator transaction binding the contract method 0xfbbe2f31.
//
// Solidity: function addSupporteAsset(address asset_, string symbol_) returns()
func (_BountyPledge *BountyPledgeSession) AddSupporteAsset(asset_ common.Address, symbol_ string) (*types.Transaction, error) {
	return _BountyPledge.Contract.AddSupporteAsset(&_BountyPledge.TransactOpts, asset_, symbol_)
}

// AddSupporteAsset is a paid mutator transaction binding the contract method 0xfbbe2f31.
//
// Solidity: function addSupporteAsset(address asset_, string symbol_) returns()
func (_BountyPledge *BountyPledgeTransactorSession) AddSupporteAsset(asset_ common.Address, symbol_ string) (*types.Transaction, error) {
	return _BountyPledge.Contract.AddSupporteAsset(&_BountyPledge.TransactOpts, asset_, symbol_)
}

// ApproveReward is a paid mutator transaction binding the contract method 0x689bccc9.
//
// Solidity: function approveReward(string cid, uint256 amount) returns()
func (_BountyPledge *BountyPledgeTransactor) ApproveReward(opts *bind.TransactOpts, cid string, amount *big.Int) (*types.Transaction, error) {
	return _BountyPledge.contract.Transact(opts, "approveReward", cid, amount)
}

// ApproveReward is a paid mutator transaction binding the contract method 0x689bccc9.
//
// Solidity: function approveReward(string cid, uint256 amount) returns()
func (_BountyPledge *BountyPledgeSession) ApproveReward(cid string, amount *big.Int) (*types.Transaction, error) {
	return _BountyPledge.Contract.ApproveReward(&_BountyPledge.TransactOpts, cid, amount)
}

// ApproveReward is a paid mutator transaction binding the contract method 0x689bccc9.
//
// Solidity: function approveReward(string cid, uint256 amount) returns()
func (_BountyPledge *BountyPledgeTransactorSession) ApproveReward(cid string, amount *big.Int) (*types.Transaction, error) {
	return _BountyPledge.Contract.ApproveReward(&_BountyPledge.TransactOpts, cid, amount)
}

// DisableAsset is a paid mutator transaction binding the contract method 0x70807528.
//
// Solidity: function disableAsset(address asset_) returns()
func (_BountyPledge *BountyPledgeTransactor) DisableAsset(opts *bind.TransactOpts, asset_ common.Address) (*types.Transaction, error) {
	return _BountyPledge.contract.Transact(opts, "disableAsset", asset_)
}

// DisableAsset is a paid mutator transaction binding the contract method 0x70807528.
//
// Solidity: function disableAsset(address asset_) returns()
func (_BountyPledge *BountyPledgeSession) DisableAsset(asset_ common.Address) (*types.Transaction, error) {
	return _BountyPledge.Contract.DisableAsset(&_BountyPledge.TransactOpts, asset_)
}

// DisableAsset is a paid mutator transaction binding the contract method 0x70807528.
//
// Solidity: function disableAsset(address asset_) returns()
func (_BountyPledge *BountyPledgeTransactorSession) DisableAsset(asset_ common.Address) (*types.Transaction, error) {
	return _BountyPledge.Contract.DisableAsset(&_BountyPledge.TransactOpts, asset_)
}

// DisableWorldStatus is a paid mutator transaction binding the contract method 0x0bf4ae70.
//
// Solidity: function disableWorldStatus() returns()
func (_BountyPledge *BountyPledgeTransactor) DisableWorldStatus(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BountyPledge.contract.Transact(opts, "disableWorldStatus")
}

// DisableWorldStatus is a paid mutator transaction binding the contract method 0x0bf4ae70.
//
// Solidity: function disableWorldStatus() returns()
func (_BountyPledge *BountyPledgeSession) DisableWorldStatus() (*types.Transaction, error) {
	return _BountyPledge.Contract.DisableWorldStatus(&_BountyPledge.TransactOpts)
}

// DisableWorldStatus is a paid mutator transaction binding the contract method 0x0bf4ae70.
//
// Solidity: function disableWorldStatus() returns()
func (_BountyPledge *BountyPledgeTransactorSession) DisableWorldStatus() (*types.Transaction, error) {
	return _BountyPledge.Contract.DisableWorldStatus(&_BountyPledge.TransactOpts)
}

// EnableAsset is a paid mutator transaction binding the contract method 0x2aef4024.
//
// Solidity: function enableAsset(address asset_) returns()
func (_BountyPledge *BountyPledgeTransactor) EnableAsset(opts *bind.TransactOpts, asset_ common.Address) (*types.Transaction, error) {
	return _BountyPledge.contract.Transact(opts, "enableAsset", asset_)
}

// EnableAsset is a paid mutator transaction binding the contract method 0x2aef4024.
//
// Solidity: function enableAsset(address asset_) returns()
func (_BountyPledge *BountyPledgeSession) EnableAsset(asset_ common.Address) (*types.Transaction, error) {
	return _BountyPledge.Contract.EnableAsset(&_BountyPledge.TransactOpts, asset_)
}

// EnableAsset is a paid mutator transaction binding the contract method 0x2aef4024.
//
// Solidity: function enableAsset(address asset_) returns()
func (_BountyPledge *BountyPledgeTransactorSession) EnableAsset(asset_ common.Address) (*types.Transaction, error) {
	return _BountyPledge.Contract.EnableAsset(&_BountyPledge.TransactOpts, asset_)
}

// EnableWorldStatus is a paid mutator transaction binding the contract method 0x27b64a31.
//
// Solidity: function enableWorldStatus() returns()
func (_BountyPledge *BountyPledgeTransactor) EnableWorldStatus(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BountyPledge.contract.Transact(opts, "enableWorldStatus")
}

// EnableWorldStatus is a paid mutator transaction binding the contract method 0x27b64a31.
//
// Solidity: function enableWorldStatus() returns()
func (_BountyPledge *BountyPledgeSession) EnableWorldStatus() (*types.Transaction, error) {
	return _BountyPledge.Contract.EnableWorldStatus(&_BountyPledge.TransactOpts)
}

// EnableWorldStatus is a paid mutator transaction binding the contract method 0x27b64a31.
//
// Solidity: function enableWorldStatus() returns()
func (_BountyPledge *BountyPledgeTransactorSession) EnableWorldStatus() (*types.Transaction, error) {
	return _BountyPledge.Contract.EnableWorldStatus(&_BountyPledge.TransactOpts)
}

// IncreaseReward is a paid mutator transaction binding the contract method 0xae1fdcd2.
//
// Solidity: function increaseReward(string cid, uint256 amount) returns()
func (_BountyPledge *BountyPledgeTransactor) IncreaseReward(opts *bind.TransactOpts, cid string, amount *big.Int) (*types.Transaction, error) {
	return _BountyPledge.contract.Transact(opts, "increaseReward", cid, amount)
}

// IncreaseReward is a paid mutator transaction binding the contract method 0xae1fdcd2.
//
// Solidity: function increaseReward(string cid, uint256 amount) returns()
func (_BountyPledge *BountyPledgeSession) IncreaseReward(cid string, amount *big.Int) (*types.Transaction, error) {
	return _BountyPledge.Contract.IncreaseReward(&_BountyPledge.TransactOpts, cid, amount)
}

// IncreaseReward is a paid mutator transaction binding the contract method 0xae1fdcd2.
//
// Solidity: function increaseReward(string cid, uint256 amount) returns()
func (_BountyPledge *BountyPledgeTransactorSession) IncreaseReward(cid string, amount *big.Int) (*types.Transaction, error) {
	return _BountyPledge.Contract.IncreaseReward(&_BountyPledge.TransactOpts, cid, amount)
}

// Initialize is a paid mutator transaction binding the contract method 0xe6a07063.
//
// Solidity: function initialize(string name_, string symbol_, address usdt_, string coinUnit_) returns()
func (_BountyPledge *BountyPledgeTransactor) Initialize(opts *bind.TransactOpts, name_ string, symbol_ string, usdt_ common.Address, coinUnit_ string) (*types.Transaction, error) {
	return _BountyPledge.contract.Transact(opts, "initialize", name_, symbol_, usdt_, coinUnit_)
}

// Initialize is a paid mutator transaction binding the contract method 0xe6a07063.
//
// Solidity: function initialize(string name_, string symbol_, address usdt_, string coinUnit_) returns()
func (_BountyPledge *BountyPledgeSession) Initialize(name_ string, symbol_ string, usdt_ common.Address, coinUnit_ string) (*types.Transaction, error) {
	return _BountyPledge.Contract.Initialize(&_BountyPledge.TransactOpts, name_, symbol_, usdt_, coinUnit_)
}

// Initialize is a paid mutator transaction binding the contract method 0xe6a07063.
//
// Solidity: function initialize(string name_, string symbol_, address usdt_, string coinUnit_) returns()
func (_BountyPledge *BountyPledgeTransactorSession) Initialize(name_ string, symbol_ string, usdt_ common.Address, coinUnit_ string) (*types.Transaction, error) {
	return _BountyPledge.Contract.Initialize(&_BountyPledge.TransactOpts, name_, symbol_, usdt_, coinUnit_)
}

// Pledge is a paid mutator transaction binding the contract method 0x8eafdeb9.
//
// Solidity: function pledge((uint256,uint256,uint8,string,uint256,string) req) payable returns(bytes32)
func (_BountyPledge *BountyPledgeTransactor) Pledge(opts *bind.TransactOpts, req BountyPledgeDataUpgradeablePledgeRequest) (*types.Transaction, error) {
	return _BountyPledge.contract.Transact(opts, "pledge", req)
}

// Pledge is a paid mutator transaction binding the contract method 0x8eafdeb9.
//
// Solidity: function pledge((uint256,uint256,uint8,string,uint256,string) req) payable returns(bytes32)
func (_BountyPledge *BountyPledgeSession) Pledge(req BountyPledgeDataUpgradeablePledgeRequest) (*types.Transaction, error) {
	return _BountyPledge.Contract.Pledge(&_BountyPledge.TransactOpts, req)
}

// Pledge is a paid mutator transaction binding the contract method 0x8eafdeb9.
//
// Solidity: function pledge((uint256,uint256,uint8,string,uint256,string) req) payable returns(bytes32)
func (_BountyPledge *BountyPledgeTransactorSession) Pledge(req BountyPledgeDataUpgradeablePledgeRequest) (*types.Transaction, error) {
	return _BountyPledge.Contract.Pledge(&_BountyPledge.TransactOpts, req)
}

// Pledge2 is a paid mutator transaction binding the contract method 0x83af9bc2.
//
// Solidity: function pledge2(uint256 startTime_, uint256 endTime_, uint8 bountyType_, string cid_, uint256 reward_, string rewardUnit_) payable returns(bytes32)
func (_BountyPledge *BountyPledgeTransactor) Pledge2(opts *bind.TransactOpts, startTime_ *big.Int, endTime_ *big.Int, bountyType_ uint8, cid_ string, reward_ *big.Int, rewardUnit_ string) (*types.Transaction, error) {
	return _BountyPledge.contract.Transact(opts, "pledge2", startTime_, endTime_, bountyType_, cid_, reward_, rewardUnit_)
}

// Pledge2 is a paid mutator transaction binding the contract method 0x83af9bc2.
//
// Solidity: function pledge2(uint256 startTime_, uint256 endTime_, uint8 bountyType_, string cid_, uint256 reward_, string rewardUnit_) payable returns(bytes32)
func (_BountyPledge *BountyPledgeSession) Pledge2(startTime_ *big.Int, endTime_ *big.Int, bountyType_ uint8, cid_ string, reward_ *big.Int, rewardUnit_ string) (*types.Transaction, error) {
	return _BountyPledge.Contract.Pledge2(&_BountyPledge.TransactOpts, startTime_, endTime_, bountyType_, cid_, reward_, rewardUnit_)
}

// Pledge2 is a paid mutator transaction binding the contract method 0x83af9bc2.
//
// Solidity: function pledge2(uint256 startTime_, uint256 endTime_, uint8 bountyType_, string cid_, uint256 reward_, string rewardUnit_) payable returns(bytes32)
func (_BountyPledge *BountyPledgeTransactorSession) Pledge2(startTime_ *big.Int, endTime_ *big.Int, bountyType_ uint8, cid_ string, reward_ *big.Int, rewardUnit_ string) (*types.Transaction, error) {
	return _BountyPledge.Contract.Pledge2(&_BountyPledge.TransactOpts, startTime_, endTime_, bountyType_, cid_, reward_, rewardUnit_)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BountyPledge *BountyPledgeTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BountyPledge.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BountyPledge *BountyPledgeSession) RenounceOwnership() (*types.Transaction, error) {
	return _BountyPledge.Contract.RenounceOwnership(&_BountyPledge.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BountyPledge *BountyPledgeTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _BountyPledge.Contract.RenounceOwnership(&_BountyPledge.TransactOpts)
}

// SetBountyExpired is a paid mutator transaction binding the contract method 0xf04edd87.
//
// Solidity: function setBountyExpired(string cid, uint256 deadLine) returns()
func (_BountyPledge *BountyPledgeTransactor) SetBountyExpired(opts *bind.TransactOpts, cid string, deadLine *big.Int) (*types.Transaction, error) {
	return _BountyPledge.contract.Transact(opts, "setBountyExpired", cid, deadLine)
}

// SetBountyExpired is a paid mutator transaction binding the contract method 0xf04edd87.
//
// Solidity: function setBountyExpired(string cid, uint256 deadLine) returns()
func (_BountyPledge *BountyPledgeSession) SetBountyExpired(cid string, deadLine *big.Int) (*types.Transaction, error) {
	return _BountyPledge.Contract.SetBountyExpired(&_BountyPledge.TransactOpts, cid, deadLine)
}

// SetBountyExpired is a paid mutator transaction binding the contract method 0xf04edd87.
//
// Solidity: function setBountyExpired(string cid, uint256 deadLine) returns()
func (_BountyPledge *BountyPledgeTransactorSession) SetBountyExpired(cid string, deadLine *big.Int) (*types.Transaction, error) {
	return _BountyPledge.Contract.SetBountyExpired(&_BountyPledge.TransactOpts, cid, deadLine)
}

// SetName is a paid mutator transaction binding the contract method 0xc47f0027.
//
// Solidity: function setName(string name_) returns()
func (_BountyPledge *BountyPledgeTransactor) SetName(opts *bind.TransactOpts, name_ string) (*types.Transaction, error) {
	return _BountyPledge.contract.Transact(opts, "setName", name_)
}

// SetName is a paid mutator transaction binding the contract method 0xc47f0027.
//
// Solidity: function setName(string name_) returns()
func (_BountyPledge *BountyPledgeSession) SetName(name_ string) (*types.Transaction, error) {
	return _BountyPledge.Contract.SetName(&_BountyPledge.TransactOpts, name_)
}

// SetName is a paid mutator transaction binding the contract method 0xc47f0027.
//
// Solidity: function setName(string name_) returns()
func (_BountyPledge *BountyPledgeTransactorSession) SetName(name_ string) (*types.Transaction, error) {
	return _BountyPledge.Contract.SetName(&_BountyPledge.TransactOpts, name_)
}

// SetSymbol is a paid mutator transaction binding the contract method 0xb84c8246.
//
// Solidity: function setSymbol(string symbol_) returns()
func (_BountyPledge *BountyPledgeTransactor) SetSymbol(opts *bind.TransactOpts, symbol_ string) (*types.Transaction, error) {
	return _BountyPledge.contract.Transact(opts, "setSymbol", symbol_)
}

// SetSymbol is a paid mutator transaction binding the contract method 0xb84c8246.
//
// Solidity: function setSymbol(string symbol_) returns()
func (_BountyPledge *BountyPledgeSession) SetSymbol(symbol_ string) (*types.Transaction, error) {
	return _BountyPledge.Contract.SetSymbol(&_BountyPledge.TransactOpts, symbol_)
}

// SetSymbol is a paid mutator transaction binding the contract method 0xb84c8246.
//
// Solidity: function setSymbol(string symbol_) returns()
func (_BountyPledge *BountyPledgeTransactorSession) SetSymbol(symbol_ string) (*types.Transaction, error) {
	return _BountyPledge.Contract.SetSymbol(&_BountyPledge.TransactOpts, symbol_)
}

// Transfer is a paid mutator transaction binding the contract method 0x24817ead.
//
// Solidity: function transfer(string bountyCid, address to, uint256 amount, string contributeCid) payable returns()
func (_BountyPledge *BountyPledgeTransactor) Transfer(opts *bind.TransactOpts, bountyCid string, to common.Address, amount *big.Int, contributeCid string) (*types.Transaction, error) {
	return _BountyPledge.contract.Transact(opts, "transfer", bountyCid, to, amount, contributeCid)
}

// Transfer is a paid mutator transaction binding the contract method 0x24817ead.
//
// Solidity: function transfer(string bountyCid, address to, uint256 amount, string contributeCid) payable returns()
func (_BountyPledge *BountyPledgeSession) Transfer(bountyCid string, to common.Address, amount *big.Int, contributeCid string) (*types.Transaction, error) {
	return _BountyPledge.Contract.Transfer(&_BountyPledge.TransactOpts, bountyCid, to, amount, contributeCid)
}

// Transfer is a paid mutator transaction binding the contract method 0x24817ead.
//
// Solidity: function transfer(string bountyCid, address to, uint256 amount, string contributeCid) payable returns()
func (_BountyPledge *BountyPledgeTransactorSession) Transfer(bountyCid string, to common.Address, amount *big.Int, contributeCid string) (*types.Transaction, error) {
	return _BountyPledge.Contract.Transfer(&_BountyPledge.TransactOpts, bountyCid, to, amount, contributeCid)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BountyPledge *BountyPledgeTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _BountyPledge.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BountyPledge *BountyPledgeSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BountyPledge.Contract.TransferOwnership(&_BountyPledge.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BountyPledge *BountyPledgeTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BountyPledge.Contract.TransferOwnership(&_BountyPledge.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0x31fb67c2.
//
// Solidity: function withdraw(string bountyCid) payable returns()
func (_BountyPledge *BountyPledgeTransactor) Withdraw(opts *bind.TransactOpts, bountyCid string) (*types.Transaction, error) {
	return _BountyPledge.contract.Transact(opts, "withdraw", bountyCid)
}

// Withdraw is a paid mutator transaction binding the contract method 0x31fb67c2.
//
// Solidity: function withdraw(string bountyCid) payable returns()
func (_BountyPledge *BountyPledgeSession) Withdraw(bountyCid string) (*types.Transaction, error) {
	return _BountyPledge.Contract.Withdraw(&_BountyPledge.TransactOpts, bountyCid)
}

// Withdraw is a paid mutator transaction binding the contract method 0x31fb67c2.
//
// Solidity: function withdraw(string bountyCid) payable returns()
func (_BountyPledge *BountyPledgeTransactorSession) Withdraw(bountyCid string) (*types.Transaction, error) {
	return _BountyPledge.Contract.Withdraw(&_BountyPledge.TransactOpts, bountyCid)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_BountyPledge *BountyPledgeTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BountyPledge.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_BountyPledge *BountyPledgeSession) Receive() (*types.Transaction, error) {
	return _BountyPledge.Contract.Receive(&_BountyPledge.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_BountyPledge *BountyPledgeTransactorSession) Receive() (*types.Transaction, error) {
	return _BountyPledge.Contract.Receive(&_BountyPledge.TransactOpts)
}

// BountyPledgeAssetActivityIterator is returned from FilterAssetActivity and is used to iterate over the raw logs and unpacked data for AssetActivity events raised by the BountyPledge contract.
type BountyPledgeAssetActivityIterator struct {
	Event *BountyPledgeAssetActivity // Event containing the contract specifics and raw log

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
func (it *BountyPledgeAssetActivityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BountyPledgeAssetActivity)
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
		it.Event = new(BountyPledgeAssetActivity)
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
func (it *BountyPledgeAssetActivityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BountyPledgeAssetActivityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BountyPledgeAssetActivity represents a AssetActivity event raised by the BountyPledge contract.
type BountyPledgeAssetActivity struct {
	Asset   common.Address
	Id      uint8
	Model   uint8
	Symbol  string
	Enabled bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAssetActivity is a free log retrieval operation binding the contract event 0x0d0fdbe5abe58aa6e1f895af84c9940e0f1c6605159211092e73982e8cc4218e.
//
// Solidity: event AssetActivity(address indexed asset, uint8 indexed Id, uint8 indexed model, string symbol, bool enabled)
func (_BountyPledge *BountyPledgeFilterer) FilterAssetActivity(opts *bind.FilterOpts, asset []common.Address, Id []uint8, model []uint8) (*BountyPledgeAssetActivityIterator, error) {

	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}
	var IdRule []interface{}
	for _, IdItem := range Id {
		IdRule = append(IdRule, IdItem)
	}
	var modelRule []interface{}
	for _, modelItem := range model {
		modelRule = append(modelRule, modelItem)
	}

	logs, sub, err := _BountyPledge.contract.FilterLogs(opts, "AssetActivity", assetRule, IdRule, modelRule)
	if err != nil {
		return nil, err
	}
	return &BountyPledgeAssetActivityIterator{contract: _BountyPledge.contract, event: "AssetActivity", logs: logs, sub: sub}, nil
}

// WatchAssetActivity is a free log subscription operation binding the contract event 0x0d0fdbe5abe58aa6e1f895af84c9940e0f1c6605159211092e73982e8cc4218e.
//
// Solidity: event AssetActivity(address indexed asset, uint8 indexed Id, uint8 indexed model, string symbol, bool enabled)
func (_BountyPledge *BountyPledgeFilterer) WatchAssetActivity(opts *bind.WatchOpts, sink chan<- *BountyPledgeAssetActivity, asset []common.Address, Id []uint8, model []uint8) (event.Subscription, error) {

	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}
	var IdRule []interface{}
	for _, IdItem := range Id {
		IdRule = append(IdRule, IdItem)
	}
	var modelRule []interface{}
	for _, modelItem := range model {
		modelRule = append(modelRule, modelItem)
	}

	logs, sub, err := _BountyPledge.contract.WatchLogs(opts, "AssetActivity", assetRule, IdRule, modelRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BountyPledgeAssetActivity)
				if err := _BountyPledge.contract.UnpackLog(event, "AssetActivity", log); err != nil {
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

// ParseAssetActivity is a log parse operation binding the contract event 0x0d0fdbe5abe58aa6e1f895af84c9940e0f1c6605159211092e73982e8cc4218e.
//
// Solidity: event AssetActivity(address indexed asset, uint8 indexed Id, uint8 indexed model, string symbol, bool enabled)
func (_BountyPledge *BountyPledgeFilterer) ParseAssetActivity(log types.Log) (*BountyPledgeAssetActivity, error) {
	event := new(BountyPledgeAssetActivity)
	if err := _BountyPledge.contract.UnpackLog(event, "AssetActivity", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BountyPledgeBountyActivityIterator is returned from FilterBountyActivity and is used to iterate over the raw logs and unpacked data for BountyActivity events raised by the BountyPledge contract.
type BountyPledgeBountyActivityIterator struct {
	Event *BountyPledgeBountyActivity // Event containing the contract specifics and raw log

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
func (it *BountyPledgeBountyActivityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BountyPledgeBountyActivity)
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
		it.Event = new(BountyPledgeBountyActivity)
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
func (it *BountyPledgeBountyActivityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BountyPledgeBountyActivityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BountyPledgeBountyActivity represents a BountyActivity event raised by the BountyPledge contract.
type BountyPledgeBountyActivity struct {
	From         common.Address
	Keccak256CID [32]byte
	Reward       *big.Int
	StartTime    *big.Int
	EndTime      *big.Int
	BountyType   uint8
	Status       uint8
	RewardUnit   string
	Cid          string
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterBountyActivity is a free log retrieval operation binding the contract event 0xd03d4a3a7acae67e86c3f69ea94589cf77eaaaa8cd68a5d8767454bdfaccd173.
//
// Solidity: event BountyActivity(address indexed from, bytes32 indexed keccak256CID, uint256 indexed reward, uint256 startTime, uint256 endTime, uint8 bountyType, uint8 status, string rewardUnit, string cid)
func (_BountyPledge *BountyPledgeFilterer) FilterBountyActivity(opts *bind.FilterOpts, from []common.Address, keccak256CID [][32]byte, reward []*big.Int) (*BountyPledgeBountyActivityIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var keccak256CIDRule []interface{}
	for _, keccak256CIDItem := range keccak256CID {
		keccak256CIDRule = append(keccak256CIDRule, keccak256CIDItem)
	}
	var rewardRule []interface{}
	for _, rewardItem := range reward {
		rewardRule = append(rewardRule, rewardItem)
	}

	logs, sub, err := _BountyPledge.contract.FilterLogs(opts, "BountyActivity", fromRule, keccak256CIDRule, rewardRule)
	if err != nil {
		return nil, err
	}
	return &BountyPledgeBountyActivityIterator{contract: _BountyPledge.contract, event: "BountyActivity", logs: logs, sub: sub}, nil
}

// WatchBountyActivity is a free log subscription operation binding the contract event 0xd03d4a3a7acae67e86c3f69ea94589cf77eaaaa8cd68a5d8767454bdfaccd173.
//
// Solidity: event BountyActivity(address indexed from, bytes32 indexed keccak256CID, uint256 indexed reward, uint256 startTime, uint256 endTime, uint8 bountyType, uint8 status, string rewardUnit, string cid)
func (_BountyPledge *BountyPledgeFilterer) WatchBountyActivity(opts *bind.WatchOpts, sink chan<- *BountyPledgeBountyActivity, from []common.Address, keccak256CID [][32]byte, reward []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var keccak256CIDRule []interface{}
	for _, keccak256CIDItem := range keccak256CID {
		keccak256CIDRule = append(keccak256CIDRule, keccak256CIDItem)
	}
	var rewardRule []interface{}
	for _, rewardItem := range reward {
		rewardRule = append(rewardRule, rewardItem)
	}

	logs, sub, err := _BountyPledge.contract.WatchLogs(opts, "BountyActivity", fromRule, keccak256CIDRule, rewardRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BountyPledgeBountyActivity)
				if err := _BountyPledge.contract.UnpackLog(event, "BountyActivity", log); err != nil {
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

// ParseBountyActivity is a log parse operation binding the contract event 0xd03d4a3a7acae67e86c3f69ea94589cf77eaaaa8cd68a5d8767454bdfaccd173.
//
// Solidity: event BountyActivity(address indexed from, bytes32 indexed keccak256CID, uint256 indexed reward, uint256 startTime, uint256 endTime, uint8 bountyType, uint8 status, string rewardUnit, string cid)
func (_BountyPledge *BountyPledgeFilterer) ParseBountyActivity(log types.Log) (*BountyPledgeBountyActivity, error) {
	event := new(BountyPledgeBountyActivity)
	if err := _BountyPledge.contract.UnpackLog(event, "BountyActivity", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BountyPledgeInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the BountyPledge contract.
type BountyPledgeInitializedIterator struct {
	Event *BountyPledgeInitialized // Event containing the contract specifics and raw log

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
func (it *BountyPledgeInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BountyPledgeInitialized)
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
		it.Event = new(BountyPledgeInitialized)
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
func (it *BountyPledgeInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BountyPledgeInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BountyPledgeInitialized represents a Initialized event raised by the BountyPledge contract.
type BountyPledgeInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_BountyPledge *BountyPledgeFilterer) FilterInitialized(opts *bind.FilterOpts) (*BountyPledgeInitializedIterator, error) {

	logs, sub, err := _BountyPledge.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &BountyPledgeInitializedIterator{contract: _BountyPledge.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_BountyPledge *BountyPledgeFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *BountyPledgeInitialized) (event.Subscription, error) {

	logs, sub, err := _BountyPledge.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BountyPledgeInitialized)
				if err := _BountyPledge.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_BountyPledge *BountyPledgeFilterer) ParseInitialized(log types.Log) (*BountyPledgeInitialized, error) {
	event := new(BountyPledgeInitialized)
	if err := _BountyPledge.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BountyPledgeOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the BountyPledge contract.
type BountyPledgeOwnershipTransferredIterator struct {
	Event *BountyPledgeOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *BountyPledgeOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BountyPledgeOwnershipTransferred)
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
		it.Event = new(BountyPledgeOwnershipTransferred)
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
func (it *BountyPledgeOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BountyPledgeOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BountyPledgeOwnershipTransferred represents a OwnershipTransferred event raised by the BountyPledge contract.
type BountyPledgeOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BountyPledge *BountyPledgeFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BountyPledgeOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BountyPledge.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BountyPledgeOwnershipTransferredIterator{contract: _BountyPledge.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BountyPledge *BountyPledgeFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BountyPledgeOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BountyPledge.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BountyPledgeOwnershipTransferred)
				if err := _BountyPledge.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_BountyPledge *BountyPledgeFilterer) ParseOwnershipTransferred(log types.Log) (*BountyPledgeOwnershipTransferred, error) {
	event := new(BountyPledgeOwnershipTransferred)
	if err := _BountyPledge.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BountyPledgeRewardActivityIterator is returned from FilterRewardActivity and is used to iterate over the raw logs and unpacked data for RewardActivity events raised by the BountyPledge contract.
type BountyPledgeRewardActivityIterator struct {
	Event *BountyPledgeRewardActivity // Event containing the contract specifics and raw log

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
func (it *BountyPledgeRewardActivityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BountyPledgeRewardActivity)
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
		it.Event = new(BountyPledgeRewardActivity)
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
func (it *BountyPledgeRewardActivityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BountyPledgeRewardActivityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BountyPledgeRewardActivity represents a RewardActivity event raised by the BountyPledge contract.
type BountyPledgeRewardActivity struct {
	Keccak256CID [32]byte
	Reward       *big.Int
	RewardUnit   string
	Model        uint8
	Status       uint8
	Cid          string
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRewardActivity is a free log retrieval operation binding the contract event 0xa72c7c8e95772f2f7475676eb66dbb0a320aa2fb673e207fc37f049705003d4f.
//
// Solidity: event RewardActivity(bytes32 indexed keccak256CID, uint256 indexed reward, string rewardUnit, uint8 indexed model, uint8 status, string cid)
func (_BountyPledge *BountyPledgeFilterer) FilterRewardActivity(opts *bind.FilterOpts, keccak256CID [][32]byte, reward []*big.Int, model []uint8) (*BountyPledgeRewardActivityIterator, error) {

	var keccak256CIDRule []interface{}
	for _, keccak256CIDItem := range keccak256CID {
		keccak256CIDRule = append(keccak256CIDRule, keccak256CIDItem)
	}
	var rewardRule []interface{}
	for _, rewardItem := range reward {
		rewardRule = append(rewardRule, rewardItem)
	}

	var modelRule []interface{}
	for _, modelItem := range model {
		modelRule = append(modelRule, modelItem)
	}

	logs, sub, err := _BountyPledge.contract.FilterLogs(opts, "RewardActivity", keccak256CIDRule, rewardRule, modelRule)
	if err != nil {
		return nil, err
	}
	return &BountyPledgeRewardActivityIterator{contract: _BountyPledge.contract, event: "RewardActivity", logs: logs, sub: sub}, nil
}

// WatchRewardActivity is a free log subscription operation binding the contract event 0xa72c7c8e95772f2f7475676eb66dbb0a320aa2fb673e207fc37f049705003d4f.
//
// Solidity: event RewardActivity(bytes32 indexed keccak256CID, uint256 indexed reward, string rewardUnit, uint8 indexed model, uint8 status, string cid)
func (_BountyPledge *BountyPledgeFilterer) WatchRewardActivity(opts *bind.WatchOpts, sink chan<- *BountyPledgeRewardActivity, keccak256CID [][32]byte, reward []*big.Int, model []uint8) (event.Subscription, error) {

	var keccak256CIDRule []interface{}
	for _, keccak256CIDItem := range keccak256CID {
		keccak256CIDRule = append(keccak256CIDRule, keccak256CIDItem)
	}
	var rewardRule []interface{}
	for _, rewardItem := range reward {
		rewardRule = append(rewardRule, rewardItem)
	}

	var modelRule []interface{}
	for _, modelItem := range model {
		modelRule = append(modelRule, modelItem)
	}

	logs, sub, err := _BountyPledge.contract.WatchLogs(opts, "RewardActivity", keccak256CIDRule, rewardRule, modelRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BountyPledgeRewardActivity)
				if err := _BountyPledge.contract.UnpackLog(event, "RewardActivity", log); err != nil {
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

// ParseRewardActivity is a log parse operation binding the contract event 0xa72c7c8e95772f2f7475676eb66dbb0a320aa2fb673e207fc37f049705003d4f.
//
// Solidity: event RewardActivity(bytes32 indexed keccak256CID, uint256 indexed reward, string rewardUnit, uint8 indexed model, uint8 status, string cid)
func (_BountyPledge *BountyPledgeFilterer) ParseRewardActivity(log types.Log) (*BountyPledgeRewardActivity, error) {
	event := new(BountyPledgeRewardActivity)
	if err := _BountyPledge.contract.UnpackLog(event, "RewardActivity", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BountyPledgeTransferActivityIterator is returned from FilterTransferActivity and is used to iterate over the raw logs and unpacked data for TransferActivity events raised by the BountyPledge contract.
type BountyPledgeTransferActivityIterator struct {
	Event *BountyPledgeTransferActivity // Event containing the contract specifics and raw log

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
func (it *BountyPledgeTransferActivityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BountyPledgeTransferActivity)
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
		it.Event = new(BountyPledgeTransferActivity)
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
func (it *BountyPledgeTransferActivityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BountyPledgeTransferActivityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BountyPledgeTransferActivity represents a TransferActivity event raised by the BountyPledge contract.
type BountyPledgeTransferActivity struct {
	From               common.Address
	To                 common.Address
	Keccak256BountyCID [32]byte
	Reward             *big.Int
	Status             uint8
	RewardUnit         string
	BountyCid          string
	ContributeCid      string
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterTransferActivity is a free log retrieval operation binding the contract event 0xba683d94f8ab8e3fc39d318e64f0df170c51344d7e74222116640daaf61baf6b.
//
// Solidity: event TransferActivity(address indexed from, address indexed to, bytes32 indexed keccak256BountyCID, uint256 reward, uint8 status, string rewardUnit, string bountyCid, string contributeCid)
func (_BountyPledge *BountyPledgeFilterer) FilterTransferActivity(opts *bind.FilterOpts, from []common.Address, to []common.Address, keccak256BountyCID [][32]byte) (*BountyPledgeTransferActivityIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var keccak256BountyCIDRule []interface{}
	for _, keccak256BountyCIDItem := range keccak256BountyCID {
		keccak256BountyCIDRule = append(keccak256BountyCIDRule, keccak256BountyCIDItem)
	}

	logs, sub, err := _BountyPledge.contract.FilterLogs(opts, "TransferActivity", fromRule, toRule, keccak256BountyCIDRule)
	if err != nil {
		return nil, err
	}
	return &BountyPledgeTransferActivityIterator{contract: _BountyPledge.contract, event: "TransferActivity", logs: logs, sub: sub}, nil
}

// WatchTransferActivity is a free log subscription operation binding the contract event 0xba683d94f8ab8e3fc39d318e64f0df170c51344d7e74222116640daaf61baf6b.
//
// Solidity: event TransferActivity(address indexed from, address indexed to, bytes32 indexed keccak256BountyCID, uint256 reward, uint8 status, string rewardUnit, string bountyCid, string contributeCid)
func (_BountyPledge *BountyPledgeFilterer) WatchTransferActivity(opts *bind.WatchOpts, sink chan<- *BountyPledgeTransferActivity, from []common.Address, to []common.Address, keccak256BountyCID [][32]byte) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var keccak256BountyCIDRule []interface{}
	for _, keccak256BountyCIDItem := range keccak256BountyCID {
		keccak256BountyCIDRule = append(keccak256BountyCIDRule, keccak256BountyCIDItem)
	}

	logs, sub, err := _BountyPledge.contract.WatchLogs(opts, "TransferActivity", fromRule, toRule, keccak256BountyCIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BountyPledgeTransferActivity)
				if err := _BountyPledge.contract.UnpackLog(event, "TransferActivity", log); err != nil {
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

// ParseTransferActivity is a log parse operation binding the contract event 0xba683d94f8ab8e3fc39d318e64f0df170c51344d7e74222116640daaf61baf6b.
//
// Solidity: event TransferActivity(address indexed from, address indexed to, bytes32 indexed keccak256BountyCID, uint256 reward, uint8 status, string rewardUnit, string bountyCid, string contributeCid)
func (_BountyPledge *BountyPledgeFilterer) ParseTransferActivity(log types.Log) (*BountyPledgeTransferActivity, error) {
	event := new(BountyPledgeTransferActivity)
	if err := _BountyPledge.contract.UnpackLog(event, "TransferActivity", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BountyPledgeWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the BountyPledge contract.
type BountyPledgeWithdrawIterator struct {
	Event *BountyPledgeWithdraw // Event containing the contract specifics and raw log

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
func (it *BountyPledgeWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BountyPledgeWithdraw)
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
		it.Event = new(BountyPledgeWithdraw)
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
func (it *BountyPledgeWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BountyPledgeWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BountyPledgeWithdraw represents a Withdraw event raised by the BountyPledge contract.
type BountyPledgeWithdraw struct {
	Recipient          common.Address
	Keccak256BountyCID [32]byte
	Reward             *big.Int
	Status             uint8
	RewardUnit         string
	BountyCid          string
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x7049445bd308cc9c764c5d6a197d251dd2b1ff724ab13a8d51b17fb4c44eab92.
//
// Solidity: event Withdraw(address indexed recipient, bytes32 indexed keccak256BountyCID, uint256 indexed reward, uint8 status, string rewardUnit, string bountyCid)
func (_BountyPledge *BountyPledgeFilterer) FilterWithdraw(opts *bind.FilterOpts, recipient []common.Address, keccak256BountyCID [][32]byte, reward []*big.Int) (*BountyPledgeWithdrawIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}
	var keccak256BountyCIDRule []interface{}
	for _, keccak256BountyCIDItem := range keccak256BountyCID {
		keccak256BountyCIDRule = append(keccak256BountyCIDRule, keccak256BountyCIDItem)
	}
	var rewardRule []interface{}
	for _, rewardItem := range reward {
		rewardRule = append(rewardRule, rewardItem)
	}

	logs, sub, err := _BountyPledge.contract.FilterLogs(opts, "Withdraw", recipientRule, keccak256BountyCIDRule, rewardRule)
	if err != nil {
		return nil, err
	}
	return &BountyPledgeWithdrawIterator{contract: _BountyPledge.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x7049445bd308cc9c764c5d6a197d251dd2b1ff724ab13a8d51b17fb4c44eab92.
//
// Solidity: event Withdraw(address indexed recipient, bytes32 indexed keccak256BountyCID, uint256 indexed reward, uint8 status, string rewardUnit, string bountyCid)
func (_BountyPledge *BountyPledgeFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *BountyPledgeWithdraw, recipient []common.Address, keccak256BountyCID [][32]byte, reward []*big.Int) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}
	var keccak256BountyCIDRule []interface{}
	for _, keccak256BountyCIDItem := range keccak256BountyCID {
		keccak256BountyCIDRule = append(keccak256BountyCIDRule, keccak256BountyCIDItem)
	}
	var rewardRule []interface{}
	for _, rewardItem := range reward {
		rewardRule = append(rewardRule, rewardItem)
	}

	logs, sub, err := _BountyPledge.contract.WatchLogs(opts, "Withdraw", recipientRule, keccak256BountyCIDRule, rewardRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BountyPledgeWithdraw)
				if err := _BountyPledge.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0x7049445bd308cc9c764c5d6a197d251dd2b1ff724ab13a8d51b17fb4c44eab92.
//
// Solidity: event Withdraw(address indexed recipient, bytes32 indexed keccak256BountyCID, uint256 indexed reward, uint8 status, string rewardUnit, string bountyCid)
func (_BountyPledge *BountyPledgeFilterer) ParseWithdraw(log types.Log) (*BountyPledgeWithdraw, error) {
	event := new(BountyPledgeWithdraw)
	if err := _BountyPledge.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
