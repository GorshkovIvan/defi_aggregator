// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package stakingrewards

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

// StakingrewardsABI is the input ABI used to generate the binding from.
const StakingrewardsABI = "[{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"earned\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"exit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"getReward\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getRewardForDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lastTimeRewardApplicable\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rewardPerToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"stake\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Stakingrewards is an auto generated Go binding around an Ethereum contract.
type Stakingrewards struct {
	StakingrewardsCaller     // Read-only binding to the contract
	StakingrewardsTransactor // Write-only binding to the contract
	StakingrewardsFilterer   // Log filterer for contract events
}

// StakingrewardsCaller is an auto generated read-only Go binding around an Ethereum contract.
type StakingrewardsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingrewardsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StakingrewardsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingrewardsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StakingrewardsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingrewardsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StakingrewardsSession struct {
	Contract     *Stakingrewards   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StakingrewardsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StakingrewardsCallerSession struct {
	Contract *StakingrewardsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// StakingrewardsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StakingrewardsTransactorSession struct {
	Contract     *StakingrewardsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// StakingrewardsRaw is an auto generated low-level Go binding around an Ethereum contract.
type StakingrewardsRaw struct {
	Contract *Stakingrewards // Generic contract binding to access the raw methods on
}

// StakingrewardsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StakingrewardsCallerRaw struct {
	Contract *StakingrewardsCaller // Generic read-only contract binding to access the raw methods on
}

// StakingrewardsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StakingrewardsTransactorRaw struct {
	Contract *StakingrewardsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStakingrewards creates a new instance of Stakingrewards, bound to a specific deployed contract.
func NewStakingrewards(address common.Address, backend bind.ContractBackend) (*Stakingrewards, error) {
	contract, err := bindStakingrewards(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Stakingrewards{StakingrewardsCaller: StakingrewardsCaller{contract: contract}, StakingrewardsTransactor: StakingrewardsTransactor{contract: contract}, StakingrewardsFilterer: StakingrewardsFilterer{contract: contract}}, nil
}

// NewStakingrewardsCaller creates a new read-only instance of Stakingrewards, bound to a specific deployed contract.
func NewStakingrewardsCaller(address common.Address, caller bind.ContractCaller) (*StakingrewardsCaller, error) {
	contract, err := bindStakingrewards(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StakingrewardsCaller{contract: contract}, nil
}

// NewStakingrewardsTransactor creates a new write-only instance of Stakingrewards, bound to a specific deployed contract.
func NewStakingrewardsTransactor(address common.Address, transactor bind.ContractTransactor) (*StakingrewardsTransactor, error) {
	contract, err := bindStakingrewards(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StakingrewardsTransactor{contract: contract}, nil
}

// NewStakingrewardsFilterer creates a new log filterer instance of Stakingrewards, bound to a specific deployed contract.
func NewStakingrewardsFilterer(address common.Address, filterer bind.ContractFilterer) (*StakingrewardsFilterer, error) {
	contract, err := bindStakingrewards(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StakingrewardsFilterer{contract: contract}, nil
}

// bindStakingrewards binds a generic wrapper to an already deployed contract.
func bindStakingrewards(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StakingrewardsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Stakingrewards *StakingrewardsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Stakingrewards.Contract.StakingrewardsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Stakingrewards *StakingrewardsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Stakingrewards.Contract.StakingrewardsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Stakingrewards *StakingrewardsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Stakingrewards.Contract.StakingrewardsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Stakingrewards *StakingrewardsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Stakingrewards.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Stakingrewards *StakingrewardsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Stakingrewards.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Stakingrewards *StakingrewardsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Stakingrewards.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Stakingrewards *StakingrewardsCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Stakingrewards.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Stakingrewards *StakingrewardsSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Stakingrewards.Contract.BalanceOf(&_Stakingrewards.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Stakingrewards *StakingrewardsCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Stakingrewards.Contract.BalanceOf(&_Stakingrewards.CallOpts, account)
}

// Earned is a free data retrieval call binding the contract method 0x008cc262.
//
// Solidity: function earned(address account) view returns(uint256)
func (_Stakingrewards *StakingrewardsCaller) Earned(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Stakingrewards.contract.Call(opts, &out, "earned", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Earned is a free data retrieval call binding the contract method 0x008cc262.
//
// Solidity: function earned(address account) view returns(uint256)
func (_Stakingrewards *StakingrewardsSession) Earned(account common.Address) (*big.Int, error) {
	return _Stakingrewards.Contract.Earned(&_Stakingrewards.CallOpts, account)
}

// Earned is a free data retrieval call binding the contract method 0x008cc262.
//
// Solidity: function earned(address account) view returns(uint256)
func (_Stakingrewards *StakingrewardsCallerSession) Earned(account common.Address) (*big.Int, error) {
	return _Stakingrewards.Contract.Earned(&_Stakingrewards.CallOpts, account)
}

// GetRewardForDuration is a free data retrieval call binding the contract method 0x1c1f78eb.
//
// Solidity: function getRewardForDuration() view returns(uint256)
func (_Stakingrewards *StakingrewardsCaller) GetRewardForDuration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Stakingrewards.contract.Call(opts, &out, "getRewardForDuration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRewardForDuration is a free data retrieval call binding the contract method 0x1c1f78eb.
//
// Solidity: function getRewardForDuration() view returns(uint256)
func (_Stakingrewards *StakingrewardsSession) GetRewardForDuration() (*big.Int, error) {
	return _Stakingrewards.Contract.GetRewardForDuration(&_Stakingrewards.CallOpts)
}

// GetRewardForDuration is a free data retrieval call binding the contract method 0x1c1f78eb.
//
// Solidity: function getRewardForDuration() view returns(uint256)
func (_Stakingrewards *StakingrewardsCallerSession) GetRewardForDuration() (*big.Int, error) {
	return _Stakingrewards.Contract.GetRewardForDuration(&_Stakingrewards.CallOpts)
}

// LastTimeRewardApplicable is a free data retrieval call binding the contract method 0x80faa57d.
//
// Solidity: function lastTimeRewardApplicable() view returns(uint256)
func (_Stakingrewards *StakingrewardsCaller) LastTimeRewardApplicable(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Stakingrewards.contract.Call(opts, &out, "lastTimeRewardApplicable")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastTimeRewardApplicable is a free data retrieval call binding the contract method 0x80faa57d.
//
// Solidity: function lastTimeRewardApplicable() view returns(uint256)
func (_Stakingrewards *StakingrewardsSession) LastTimeRewardApplicable() (*big.Int, error) {
	return _Stakingrewards.Contract.LastTimeRewardApplicable(&_Stakingrewards.CallOpts)
}

// LastTimeRewardApplicable is a free data retrieval call binding the contract method 0x80faa57d.
//
// Solidity: function lastTimeRewardApplicable() view returns(uint256)
func (_Stakingrewards *StakingrewardsCallerSession) LastTimeRewardApplicable() (*big.Int, error) {
	return _Stakingrewards.Contract.LastTimeRewardApplicable(&_Stakingrewards.CallOpts)
}

// RewardPerToken is a free data retrieval call binding the contract method 0xcd3daf9d.
//
// Solidity: function rewardPerToken() view returns(uint256)
func (_Stakingrewards *StakingrewardsCaller) RewardPerToken(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Stakingrewards.contract.Call(opts, &out, "rewardPerToken")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardPerToken is a free data retrieval call binding the contract method 0xcd3daf9d.
//
// Solidity: function rewardPerToken() view returns(uint256)
func (_Stakingrewards *StakingrewardsSession) RewardPerToken() (*big.Int, error) {
	return _Stakingrewards.Contract.RewardPerToken(&_Stakingrewards.CallOpts)
}

// RewardPerToken is a free data retrieval call binding the contract method 0xcd3daf9d.
//
// Solidity: function rewardPerToken() view returns(uint256)
func (_Stakingrewards *StakingrewardsCallerSession) RewardPerToken() (*big.Int, error) {
	return _Stakingrewards.Contract.RewardPerToken(&_Stakingrewards.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Stakingrewards *StakingrewardsCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Stakingrewards.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Stakingrewards *StakingrewardsSession) TotalSupply() (*big.Int, error) {
	return _Stakingrewards.Contract.TotalSupply(&_Stakingrewards.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Stakingrewards *StakingrewardsCallerSession) TotalSupply() (*big.Int, error) {
	return _Stakingrewards.Contract.TotalSupply(&_Stakingrewards.CallOpts)
}

// Exit is a paid mutator transaction binding the contract method 0xe9fad8ee.
//
// Solidity: function exit() returns()
func (_Stakingrewards *StakingrewardsTransactor) Exit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Stakingrewards.contract.Transact(opts, "exit")
}

// Exit is a paid mutator transaction binding the contract method 0xe9fad8ee.
//
// Solidity: function exit() returns()
func (_Stakingrewards *StakingrewardsSession) Exit() (*types.Transaction, error) {
	return _Stakingrewards.Contract.Exit(&_Stakingrewards.TransactOpts)
}

// Exit is a paid mutator transaction binding the contract method 0xe9fad8ee.
//
// Solidity: function exit() returns()
func (_Stakingrewards *StakingrewardsTransactorSession) Exit() (*types.Transaction, error) {
	return _Stakingrewards.Contract.Exit(&_Stakingrewards.TransactOpts)
}

// GetReward is a paid mutator transaction binding the contract method 0x3d18b912.
//
// Solidity: function getReward() returns()
func (_Stakingrewards *StakingrewardsTransactor) GetReward(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Stakingrewards.contract.Transact(opts, "getReward")
}

// GetReward is a paid mutator transaction binding the contract method 0x3d18b912.
//
// Solidity: function getReward() returns()
func (_Stakingrewards *StakingrewardsSession) GetReward() (*types.Transaction, error) {
	return _Stakingrewards.Contract.GetReward(&_Stakingrewards.TransactOpts)
}

// GetReward is a paid mutator transaction binding the contract method 0x3d18b912.
//
// Solidity: function getReward() returns()
func (_Stakingrewards *StakingrewardsTransactorSession) GetReward() (*types.Transaction, error) {
	return _Stakingrewards.Contract.GetReward(&_Stakingrewards.TransactOpts)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 amount) returns()
func (_Stakingrewards *StakingrewardsTransactor) Stake(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Stakingrewards.contract.Transact(opts, "stake", amount)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 amount) returns()
func (_Stakingrewards *StakingrewardsSession) Stake(amount *big.Int) (*types.Transaction, error) {
	return _Stakingrewards.Contract.Stake(&_Stakingrewards.TransactOpts, amount)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 amount) returns()
func (_Stakingrewards *StakingrewardsTransactorSession) Stake(amount *big.Int) (*types.Transaction, error) {
	return _Stakingrewards.Contract.Stake(&_Stakingrewards.TransactOpts, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_Stakingrewards *StakingrewardsTransactor) Withdraw(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Stakingrewards.contract.Transact(opts, "withdraw", amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_Stakingrewards *StakingrewardsSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _Stakingrewards.Contract.Withdraw(&_Stakingrewards.TransactOpts, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_Stakingrewards *StakingrewardsTransactorSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _Stakingrewards.Contract.Withdraw(&_Stakingrewards.TransactOpts, amount)
}
