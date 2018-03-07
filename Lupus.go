// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

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

// DogsOfRomeABI is the input ABI used to generate the binding from.
const DogsOfRomeABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"question\",\"type\":\"string\"}],\"name\":\"askScoobyDoo\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"usage_fee\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"counter\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"romulus\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"asker\",\"type\":\"address\"},{\"name\":\"question\",\"type\":\"uint256\"},{\"name\":\"answer\",\"type\":\"string\"},{\"name\":\"questionBlock\",\"type\":\"uint256\"}],\"name\":\"letRomulusReply\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lastQBlock\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newRate\",\"type\":\"uint256\"}],\"name\":\"setUsageRate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"balance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"question\",\"type\":\"string\"}],\"name\":\"askRomulus\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_auth\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"asker\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"nonce\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"question\",\"type\":\"string\"}],\"name\":\"Romulus\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"asker\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"nonce\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"question\",\"type\":\"string\"}],\"name\":\"Scooby\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"asker\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"nonce\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"answer\",\"type\":\"string\"}],\"name\":\"Remus\",\"type\":\"event\"}]"

// DogsOfRomeBin is the compiled bytecode used for deploying new contracts.
const DogsOfRomeBin = `0x606060405260038054600160a060020a03191633600160a060020a0316179055341561002a57600080fd5b6040516020806106628339810160405280805160028054600160a060020a031916600160a060020a0392909216919091179055505067016345785d8a0000600555436001556105e48061007e6000396000f3006060604052600436106100ae5763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416631235ee2581146100b35780633ccfd60b146100fb578063563757e51461010e57806361bc221a1461013357806364e14cce1461014657806367a1ac32146101755780638da5cb5b146101dc5780639dac6c5f146101ef578063aa274f6914610202578063e3d670d714610218578063fc22f9ed14610237575b600080fd5b6100f960046024813581810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284375094965061027d95505050505050565b005b341561010657600080fd5b6100f961033d565b341561011957600080fd5b610121610388565b60405190815260200160405180910390f35b341561013e57600080fd5b61012161038e565b341561015157600080fd5b610159610394565b604051600160a060020a03909116815260200160405180910390f35b341561018057600080fd5b6100f960048035600160a060020a03169060248035919060649060443590810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284375094965050933593506103a392505050565b34156101e757600080fd5b610159610484565b34156101fa57600080fd5b610121610493565b341561020d57600080fd5b6100f9600435610499565b341561022357600080fd5b610121600160a060020a03600435166104b9565b6100f960046024813581810190830135806020601f820181900481020160405190810160405281815292919060208401838380828437509496506104cb95505050505050565b610285610543565b33600160a060020a03167f3b23f12c7fa973944b1e0f9b9528682c91ef9c2e51597b53f66a3872dca801d36000548360405182815260406020820181815290820183818151815260200191508051906020019080838360005b838110156102f65780820151838201526020016102de565b50505050905090810190601f1680156103235780820380516001836020036101000a031916815260200191505b50935050505060405180910390a250600080546001019055565b600160a060020a033316600081815260046020526040808220805492905590919082156108fc0290839051600060405180830381858888f19350505050151561038557600080fd5b50565b60055481565b60005481565b600254600160a060020a031681565b60025433600160a060020a039081169116146103be57600080fd5b83600160a060020a03167ffdb046d2db5929b5bc1f8c23c348c991db6b9f0cd5d26a1826b48ca62cd85b07848460405182815260406020820181815290820183818151815260200191508051906020019080838360005b8381101561042d578082015183820152602001610415565b50505050905090810190601f16801561045a5780820380516001836020036101000a031916815260200191505b50935050505060405180910390a260015481116104795760015461047b565b805b60015550505050565b600354600160a060020a031681565b60015481565b60035433600160a060020a039081169116146104b457600080fd5b600555565b60046020526000908152604090205481565b6104d3610543565b33600160a060020a03167f85a71fb24520052471c503f4da51cda8036cb15f8e06c7c2a2c5801c313ff0b4600054836040518281526040602082018181529082018381815181526020019150805190602001908083836000838110156102f65780820151838201526020016102de565b600554600160a060020a03331660009081526004602052604090205461056f903463ffffffff6105a216565b101561057a57600080fd5b600160a060020a03331660009081526004602052604090208054340180825560055490039055565b6000828201838110156105b157fe5b93925050505600a165627a7a723058204a158f9287422b2fdc87924d36b4147e260d221a9f7a0b04e9b7f469849e0da80029`

// DeployDogsOfRome deploys a new Ethereum contract, binding an instance of DogsOfRome to it.
func DeployDogsOfRome(auth *bind.TransactOpts, backend bind.ContractBackend, _auth common.Address) (common.Address, *types.Transaction, *DogsOfRome, error) {
	parsed, err := abi.JSON(strings.NewReader(DogsOfRomeABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(DogsOfRomeBin), backend, _auth)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DogsOfRome{DogsOfRomeCaller: DogsOfRomeCaller{contract: contract}, DogsOfRomeTransactor: DogsOfRomeTransactor{contract: contract}, DogsOfRomeFilterer: DogsOfRomeFilterer{contract: contract}}, nil
}

// DogsOfRome is an auto generated Go binding around an Ethereum contract.
type DogsOfRome struct {
	DogsOfRomeCaller     // Read-only binding to the contract
	DogsOfRomeTransactor // Write-only binding to the contract
	DogsOfRomeFilterer   // Log filterer for contract events
}

// DogsOfRomeCaller is an auto generated read-only Go binding around an Ethereum contract.
type DogsOfRomeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DogsOfRomeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DogsOfRomeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DogsOfRomeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DogsOfRomeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DogsOfRomeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DogsOfRomeSession struct {
	Contract     *DogsOfRome       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DogsOfRomeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DogsOfRomeCallerSession struct {
	Contract *DogsOfRomeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// DogsOfRomeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DogsOfRomeTransactorSession struct {
	Contract     *DogsOfRomeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// DogsOfRomeRaw is an auto generated low-level Go binding around an Ethereum contract.
type DogsOfRomeRaw struct {
	Contract *DogsOfRome // Generic contract binding to access the raw methods on
}

// DogsOfRomeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DogsOfRomeCallerRaw struct {
	Contract *DogsOfRomeCaller // Generic read-only contract binding to access the raw methods on
}

// DogsOfRomeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DogsOfRomeTransactorRaw struct {
	Contract *DogsOfRomeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDogsOfRome creates a new instance of DogsOfRome, bound to a specific deployed contract.
func NewDogsOfRome(address common.Address, backend bind.ContractBackend) (*DogsOfRome, error) {
	contract, err := bindDogsOfRome(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DogsOfRome{DogsOfRomeCaller: DogsOfRomeCaller{contract: contract}, DogsOfRomeTransactor: DogsOfRomeTransactor{contract: contract}, DogsOfRomeFilterer: DogsOfRomeFilterer{contract: contract}}, nil
}

// NewDogsOfRomeCaller creates a new read-only instance of DogsOfRome, bound to a specific deployed contract.
func NewDogsOfRomeCaller(address common.Address, caller bind.ContractCaller) (*DogsOfRomeCaller, error) {
	contract, err := bindDogsOfRome(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DogsOfRomeCaller{contract: contract}, nil
}

// NewDogsOfRomeTransactor creates a new write-only instance of DogsOfRome, bound to a specific deployed contract.
func NewDogsOfRomeTransactor(address common.Address, transactor bind.ContractTransactor) (*DogsOfRomeTransactor, error) {
	contract, err := bindDogsOfRome(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DogsOfRomeTransactor{contract: contract}, nil
}

// NewDogsOfRomeFilterer creates a new log filterer instance of DogsOfRome, bound to a specific deployed contract.
func NewDogsOfRomeFilterer(address common.Address, filterer bind.ContractFilterer) (*DogsOfRomeFilterer, error) {
	contract, err := bindDogsOfRome(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DogsOfRomeFilterer{contract: contract}, nil
}

// bindDogsOfRome binds a generic wrapper to an already deployed contract.
func bindDogsOfRome(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DogsOfRomeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DogsOfRome *DogsOfRomeRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DogsOfRome.Contract.DogsOfRomeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DogsOfRome *DogsOfRomeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DogsOfRome.Contract.DogsOfRomeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DogsOfRome *DogsOfRomeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DogsOfRome.Contract.DogsOfRomeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DogsOfRome *DogsOfRomeCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DogsOfRome.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DogsOfRome *DogsOfRomeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DogsOfRome.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DogsOfRome *DogsOfRomeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DogsOfRome.Contract.contract.Transact(opts, method, params...)
}

// Balance is a free data retrieval call binding the contract method 0xe3d670d7.
//
// Solidity: function balance( address) constant returns(uint256)
func (_DogsOfRome *DogsOfRomeCaller) Balance(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DogsOfRome.contract.Call(opts, out, "balance", arg0)
	return *ret0, err
}

// Balance is a free data retrieval call binding the contract method 0xe3d670d7.
//
// Solidity: function balance( address) constant returns(uint256)
func (_DogsOfRome *DogsOfRomeSession) Balance(arg0 common.Address) (*big.Int, error) {
	return _DogsOfRome.Contract.Balance(&_DogsOfRome.CallOpts, arg0)
}

// Balance is a free data retrieval call binding the contract method 0xe3d670d7.
//
// Solidity: function balance( address) constant returns(uint256)
func (_DogsOfRome *DogsOfRomeCallerSession) Balance(arg0 common.Address) (*big.Int, error) {
	return _DogsOfRome.Contract.Balance(&_DogsOfRome.CallOpts, arg0)
}

// Counter is a free data retrieval call binding the contract method 0x61bc221a.
//
// Solidity: function counter() constant returns(uint256)
func (_DogsOfRome *DogsOfRomeCaller) Counter(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DogsOfRome.contract.Call(opts, out, "counter")
	return *ret0, err
}

// Counter is a free data retrieval call binding the contract method 0x61bc221a.
//
// Solidity: function counter() constant returns(uint256)
func (_DogsOfRome *DogsOfRomeSession) Counter() (*big.Int, error) {
	return _DogsOfRome.Contract.Counter(&_DogsOfRome.CallOpts)
}

// Counter is a free data retrieval call binding the contract method 0x61bc221a.
//
// Solidity: function counter() constant returns(uint256)
func (_DogsOfRome *DogsOfRomeCallerSession) Counter() (*big.Int, error) {
	return _DogsOfRome.Contract.Counter(&_DogsOfRome.CallOpts)
}

// LastQBlock is a free data retrieval call binding the contract method 0x9dac6c5f.
//
// Solidity: function lastQBlock() constant returns(uint256)
func (_DogsOfRome *DogsOfRomeCaller) LastQBlock(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DogsOfRome.contract.Call(opts, out, "lastQBlock")
	return *ret0, err
}

// LastQBlock is a free data retrieval call binding the contract method 0x9dac6c5f.
//
// Solidity: function lastQBlock() constant returns(uint256)
func (_DogsOfRome *DogsOfRomeSession) LastQBlock() (*big.Int, error) {
	return _DogsOfRome.Contract.LastQBlock(&_DogsOfRome.CallOpts)
}

// LastQBlock is a free data retrieval call binding the contract method 0x9dac6c5f.
//
// Solidity: function lastQBlock() constant returns(uint256)
func (_DogsOfRome *DogsOfRomeCallerSession) LastQBlock() (*big.Int, error) {
	return _DogsOfRome.Contract.LastQBlock(&_DogsOfRome.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_DogsOfRome *DogsOfRomeCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DogsOfRome.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_DogsOfRome *DogsOfRomeSession) Owner() (common.Address, error) {
	return _DogsOfRome.Contract.Owner(&_DogsOfRome.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_DogsOfRome *DogsOfRomeCallerSession) Owner() (common.Address, error) {
	return _DogsOfRome.Contract.Owner(&_DogsOfRome.CallOpts)
}

// Romulus is a free data retrieval call binding the contract method 0x64e14cce.
//
// Solidity: function romulus() constant returns(address)
func (_DogsOfRome *DogsOfRomeCaller) Romulus(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DogsOfRome.contract.Call(opts, out, "romulus")
	return *ret0, err
}

// Romulus is a free data retrieval call binding the contract method 0x64e14cce.
//
// Solidity: function romulus() constant returns(address)
func (_DogsOfRome *DogsOfRomeSession) Romulus() (common.Address, error) {
	return _DogsOfRome.Contract.Romulus(&_DogsOfRome.CallOpts)
}

// Romulus is a free data retrieval call binding the contract method 0x64e14cce.
//
// Solidity: function romulus() constant returns(address)
func (_DogsOfRome *DogsOfRomeCallerSession) Romulus() (common.Address, error) {
	return _DogsOfRome.Contract.Romulus(&_DogsOfRome.CallOpts)
}

// Usage_fee is a free data retrieval call binding the contract method 0x563757e5.
//
// Solidity: function usage_fee() constant returns(uint256)
func (_DogsOfRome *DogsOfRomeCaller) Usage_fee(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DogsOfRome.contract.Call(opts, out, "usage_fee")
	return *ret0, err
}

// Usage_fee is a free data retrieval call binding the contract method 0x563757e5.
//
// Solidity: function usage_fee() constant returns(uint256)
func (_DogsOfRome *DogsOfRomeSession) Usage_fee() (*big.Int, error) {
	return _DogsOfRome.Contract.Usage_fee(&_DogsOfRome.CallOpts)
}

// Usage_fee is a free data retrieval call binding the contract method 0x563757e5.
//
// Solidity: function usage_fee() constant returns(uint256)
func (_DogsOfRome *DogsOfRomeCallerSession) Usage_fee() (*big.Int, error) {
	return _DogsOfRome.Contract.Usage_fee(&_DogsOfRome.CallOpts)
}

// AskRomulus is a paid mutator transaction binding the contract method 0xfc22f9ed.
//
// Solidity: function askRomulus(question string) returns()
func (_DogsOfRome *DogsOfRomeTransactor) AskRomulus(opts *bind.TransactOpts, question string) (*types.Transaction, error) {
	return _DogsOfRome.contract.Transact(opts, "askRomulus", question)
}

// AskRomulus is a paid mutator transaction binding the contract method 0xfc22f9ed.
//
// Solidity: function askRomulus(question string) returns()
func (_DogsOfRome *DogsOfRomeSession) AskRomulus(question string) (*types.Transaction, error) {
	return _DogsOfRome.Contract.AskRomulus(&_DogsOfRome.TransactOpts, question)
}

// AskRomulus is a paid mutator transaction binding the contract method 0xfc22f9ed.
//
// Solidity: function askRomulus(question string) returns()
func (_DogsOfRome *DogsOfRomeTransactorSession) AskRomulus(question string) (*types.Transaction, error) {
	return _DogsOfRome.Contract.AskRomulus(&_DogsOfRome.TransactOpts, question)
}

// AskScoobyDoo is a paid mutator transaction binding the contract method 0x1235ee25.
//
// Solidity: function askScoobyDoo(question string) returns()
func (_DogsOfRome *DogsOfRomeTransactor) AskScoobyDoo(opts *bind.TransactOpts, question string) (*types.Transaction, error) {
	return _DogsOfRome.contract.Transact(opts, "askScoobyDoo", question)
}

// AskScoobyDoo is a paid mutator transaction binding the contract method 0x1235ee25.
//
// Solidity: function askScoobyDoo(question string) returns()
func (_DogsOfRome *DogsOfRomeSession) AskScoobyDoo(question string) (*types.Transaction, error) {
	return _DogsOfRome.Contract.AskScoobyDoo(&_DogsOfRome.TransactOpts, question)
}

// AskScoobyDoo is a paid mutator transaction binding the contract method 0x1235ee25.
//
// Solidity: function askScoobyDoo(question string) returns()
func (_DogsOfRome *DogsOfRomeTransactorSession) AskScoobyDoo(question string) (*types.Transaction, error) {
	return _DogsOfRome.Contract.AskScoobyDoo(&_DogsOfRome.TransactOpts, question)
}

// LetRomulusReply is a paid mutator transaction binding the contract method 0x67a1ac32.
//
// Solidity: function letRomulusReply(asker address, question uint256, answer string, questionBlock uint256) returns()
func (_DogsOfRome *DogsOfRomeTransactor) LetRomulusReply(opts *bind.TransactOpts, asker common.Address, question *big.Int, answer string, questionBlock *big.Int) (*types.Transaction, error) {
	return _DogsOfRome.contract.Transact(opts, "letRomulusReply", asker, question, answer, questionBlock)
}

// LetRomulusReply is a paid mutator transaction binding the contract method 0x67a1ac32.
//
// Solidity: function letRomulusReply(asker address, question uint256, answer string, questionBlock uint256) returns()
func (_DogsOfRome *DogsOfRomeSession) LetRomulusReply(asker common.Address, question *big.Int, answer string, questionBlock *big.Int) (*types.Transaction, error) {
	return _DogsOfRome.Contract.LetRomulusReply(&_DogsOfRome.TransactOpts, asker, question, answer, questionBlock)
}

// LetRomulusReply is a paid mutator transaction binding the contract method 0x67a1ac32.
//
// Solidity: function letRomulusReply(asker address, question uint256, answer string, questionBlock uint256) returns()
func (_DogsOfRome *DogsOfRomeTransactorSession) LetRomulusReply(asker common.Address, question *big.Int, answer string, questionBlock *big.Int) (*types.Transaction, error) {
	return _DogsOfRome.Contract.LetRomulusReply(&_DogsOfRome.TransactOpts, asker, question, answer, questionBlock)
}

// SetUsageRate is a paid mutator transaction binding the contract method 0xaa274f69.
//
// Solidity: function setUsageRate(newRate uint256) returns()
func (_DogsOfRome *DogsOfRomeTransactor) SetUsageRate(opts *bind.TransactOpts, newRate *big.Int) (*types.Transaction, error) {
	return _DogsOfRome.contract.Transact(opts, "setUsageRate", newRate)
}

// SetUsageRate is a paid mutator transaction binding the contract method 0xaa274f69.
//
// Solidity: function setUsageRate(newRate uint256) returns()
func (_DogsOfRome *DogsOfRomeSession) SetUsageRate(newRate *big.Int) (*types.Transaction, error) {
	return _DogsOfRome.Contract.SetUsageRate(&_DogsOfRome.TransactOpts, newRate)
}

// SetUsageRate is a paid mutator transaction binding the contract method 0xaa274f69.
//
// Solidity: function setUsageRate(newRate uint256) returns()
func (_DogsOfRome *DogsOfRomeTransactorSession) SetUsageRate(newRate *big.Int) (*types.Transaction, error) {
	return _DogsOfRome.Contract.SetUsageRate(&_DogsOfRome.TransactOpts, newRate)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_DogsOfRome *DogsOfRomeTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DogsOfRome.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_DogsOfRome *DogsOfRomeSession) Withdraw() (*types.Transaction, error) {
	return _DogsOfRome.Contract.Withdraw(&_DogsOfRome.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_DogsOfRome *DogsOfRomeTransactorSession) Withdraw() (*types.Transaction, error) {
	return _DogsOfRome.Contract.Withdraw(&_DogsOfRome.TransactOpts)
}

// DogsOfRomeRemusIterator is returned from FilterRemus and is used to iterate over the raw logs and unpacked data for Remus events raised by the DogsOfRome contract.
type DogsOfRomeRemusIterator struct {
	Event *DogsOfRomeRemus // Event containing the contract specifics and raw log

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
func (it *DogsOfRomeRemusIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DogsOfRomeRemus)
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
		it.Event = new(DogsOfRomeRemus)
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
func (it *DogsOfRomeRemusIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DogsOfRomeRemusIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DogsOfRomeRemus represents a Remus event raised by the DogsOfRome contract.
type DogsOfRomeRemus struct {
	Asker  common.Address
	Nonce  *big.Int
	Answer string
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRemus is a free log retrieval operation binding the contract event 0xfdb046d2db5929b5bc1f8c23c348c991db6b9f0cd5d26a1826b48ca62cd85b07.
//
// Solidity: event Remus(asker indexed address, nonce uint256, answer string)
func (_DogsOfRome *DogsOfRomeFilterer) FilterRemus(opts *bind.FilterOpts, asker []common.Address) (*DogsOfRomeRemusIterator, error) {

	var askerRule []interface{}
	for _, askerItem := range asker {
		askerRule = append(askerRule, askerItem)
	}

	logs, sub, err := _DogsOfRome.contract.FilterLogs(opts, "Remus", askerRule)
	if err != nil {
		return nil, err
	}
	return &DogsOfRomeRemusIterator{contract: _DogsOfRome.contract, event: "Remus", logs: logs, sub: sub}, nil
}

// WatchRemus is a free log subscription operation binding the contract event 0xfdb046d2db5929b5bc1f8c23c348c991db6b9f0cd5d26a1826b48ca62cd85b07.
//
// Solidity: event Remus(asker indexed address, nonce uint256, answer string)
func (_DogsOfRome *DogsOfRomeFilterer) WatchRemus(opts *bind.WatchOpts, sink chan<- *DogsOfRomeRemus, asker []common.Address) (event.Subscription, error) {

	var askerRule []interface{}
	for _, askerItem := range asker {
		askerRule = append(askerRule, askerItem)
	}

	logs, sub, err := _DogsOfRome.contract.WatchLogs(opts, "Remus", askerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DogsOfRomeRemus)
				if err := _DogsOfRome.contract.UnpackLog(event, "Remus", log); err != nil {
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

// DogsOfRomeRomulusIterator is returned from FilterRomulus and is used to iterate over the raw logs and unpacked data for Romulus events raised by the DogsOfRome contract.
type DogsOfRomeRomulusIterator struct {
	Event *DogsOfRomeRomulus // Event containing the contract specifics and raw log

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
func (it *DogsOfRomeRomulusIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DogsOfRomeRomulus)
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
		it.Event = new(DogsOfRomeRomulus)
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
func (it *DogsOfRomeRomulusIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DogsOfRomeRomulusIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DogsOfRomeRomulus represents a Romulus event raised by the DogsOfRome contract.
type DogsOfRomeRomulus struct {
	Asker    common.Address
	Nonce    *big.Int
	Question string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRomulus is a free log retrieval operation binding the contract event 0x85a71fb24520052471c503f4da51cda8036cb15f8e06c7c2a2c5801c313ff0b4.
//
// Solidity: event Romulus(asker indexed address, nonce uint256, question string)
func (_DogsOfRome *DogsOfRomeFilterer) FilterRomulus(opts *bind.FilterOpts, asker []common.Address) (*DogsOfRomeRomulusIterator, error) {

	var askerRule []interface{}
	for _, askerItem := range asker {
		askerRule = append(askerRule, askerItem)
	}

	logs, sub, err := _DogsOfRome.contract.FilterLogs(opts, "Romulus", askerRule)
	if err != nil {
		return nil, err
	}
	return &DogsOfRomeRomulusIterator{contract: _DogsOfRome.contract, event: "Romulus", logs: logs, sub: sub}, nil
}

// WatchRomulus is a free log subscription operation binding the contract event 0x85a71fb24520052471c503f4da51cda8036cb15f8e06c7c2a2c5801c313ff0b4.
//
// Solidity: event Romulus(asker indexed address, nonce uint256, question string)
func (_DogsOfRome *DogsOfRomeFilterer) WatchRomulus(opts *bind.WatchOpts, sink chan<- *DogsOfRomeRomulus, asker []common.Address) (event.Subscription, error) {

	var askerRule []interface{}
	for _, askerItem := range asker {
		askerRule = append(askerRule, askerItem)
	}

	logs, sub, err := _DogsOfRome.contract.WatchLogs(opts, "Romulus", askerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DogsOfRomeRomulus)
				if err := _DogsOfRome.contract.UnpackLog(event, "Romulus", log); err != nil {
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

// DogsOfRomeScoobyIterator is returned from FilterScooby and is used to iterate over the raw logs and unpacked data for Scooby events raised by the DogsOfRome contract.
type DogsOfRomeScoobyIterator struct {
	Event *DogsOfRomeScooby // Event containing the contract specifics and raw log

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
func (it *DogsOfRomeScoobyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DogsOfRomeScooby)
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
		it.Event = new(DogsOfRomeScooby)
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
func (it *DogsOfRomeScoobyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DogsOfRomeScoobyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DogsOfRomeScooby represents a Scooby event raised by the DogsOfRome contract.
type DogsOfRomeScooby struct {
	Asker    common.Address
	Nonce    *big.Int
	Question string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterScooby is a free log retrieval operation binding the contract event 0x3b23f12c7fa973944b1e0f9b9528682c91ef9c2e51597b53f66a3872dca801d3.
//
// Solidity: event Scooby(asker indexed address, nonce uint256, question string)
func (_DogsOfRome *DogsOfRomeFilterer) FilterScooby(opts *bind.FilterOpts, asker []common.Address) (*DogsOfRomeScoobyIterator, error) {

	var askerRule []interface{}
	for _, askerItem := range asker {
		askerRule = append(askerRule, askerItem)
	}

	logs, sub, err := _DogsOfRome.contract.FilterLogs(opts, "Scooby", askerRule)
	if err != nil {
		return nil, err
	}
	return &DogsOfRomeScoobyIterator{contract: _DogsOfRome.contract, event: "Scooby", logs: logs, sub: sub}, nil
}

// WatchScooby is a free log subscription operation binding the contract event 0x3b23f12c7fa973944b1e0f9b9528682c91ef9c2e51597b53f66a3872dca801d3.
//
// Solidity: event Scooby(asker indexed address, nonce uint256, question string)
func (_DogsOfRome *DogsOfRomeFilterer) WatchScooby(opts *bind.WatchOpts, sink chan<- *DogsOfRomeScooby, asker []common.Address) (event.Subscription, error) {

	var askerRule []interface{}
	for _, askerItem := range asker {
		askerRule = append(askerRule, askerItem)
	}

	logs, sub, err := _DogsOfRome.contract.WatchLogs(opts, "Scooby", askerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DogsOfRomeScooby)
				if err := _DogsOfRome.contract.UnpackLog(event, "Scooby", log); err != nil {
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

// MiniSafeABI is the input ABI used to generate the binding from.
const MiniSafeABI = "[]"

// MiniSafeBin is the compiled bytecode used for deploying new contracts.
const MiniSafeBin = `0x60606040523415600e57600080fd5b603580601b6000396000f3006060604052600080fd00a165627a7a7230582046f3e0cb66ea808a028a743a701960e91407b6eb7d5834aeec2facbab338ce2e0029`

// DeployMiniSafe deploys a new Ethereum contract, binding an instance of MiniSafe to it.
func DeployMiniSafe(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MiniSafe, error) {
	parsed, err := abi.JSON(strings.NewReader(MiniSafeABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MiniSafeBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MiniSafe{MiniSafeCaller: MiniSafeCaller{contract: contract}, MiniSafeTransactor: MiniSafeTransactor{contract: contract}, MiniSafeFilterer: MiniSafeFilterer{contract: contract}}, nil
}

// MiniSafe is an auto generated Go binding around an Ethereum contract.
type MiniSafe struct {
	MiniSafeCaller     // Read-only binding to the contract
	MiniSafeTransactor // Write-only binding to the contract
	MiniSafeFilterer   // Log filterer for contract events
}

// MiniSafeCaller is an auto generated read-only Go binding around an Ethereum contract.
type MiniSafeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MiniSafeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MiniSafeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MiniSafeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MiniSafeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MiniSafeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MiniSafeSession struct {
	Contract     *MiniSafe         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MiniSafeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MiniSafeCallerSession struct {
	Contract *MiniSafeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// MiniSafeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MiniSafeTransactorSession struct {
	Contract     *MiniSafeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// MiniSafeRaw is an auto generated low-level Go binding around an Ethereum contract.
type MiniSafeRaw struct {
	Contract *MiniSafe // Generic contract binding to access the raw methods on
}

// MiniSafeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MiniSafeCallerRaw struct {
	Contract *MiniSafeCaller // Generic read-only contract binding to access the raw methods on
}

// MiniSafeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MiniSafeTransactorRaw struct {
	Contract *MiniSafeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMiniSafe creates a new instance of MiniSafe, bound to a specific deployed contract.
func NewMiniSafe(address common.Address, backend bind.ContractBackend) (*MiniSafe, error) {
	contract, err := bindMiniSafe(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MiniSafe{MiniSafeCaller: MiniSafeCaller{contract: contract}, MiniSafeTransactor: MiniSafeTransactor{contract: contract}, MiniSafeFilterer: MiniSafeFilterer{contract: contract}}, nil
}

// NewMiniSafeCaller creates a new read-only instance of MiniSafe, bound to a specific deployed contract.
func NewMiniSafeCaller(address common.Address, caller bind.ContractCaller) (*MiniSafeCaller, error) {
	contract, err := bindMiniSafe(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MiniSafeCaller{contract: contract}, nil
}

// NewMiniSafeTransactor creates a new write-only instance of MiniSafe, bound to a specific deployed contract.
func NewMiniSafeTransactor(address common.Address, transactor bind.ContractTransactor) (*MiniSafeTransactor, error) {
	contract, err := bindMiniSafe(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MiniSafeTransactor{contract: contract}, nil
}

// NewMiniSafeFilterer creates a new log filterer instance of MiniSafe, bound to a specific deployed contract.
func NewMiniSafeFilterer(address common.Address, filterer bind.ContractFilterer) (*MiniSafeFilterer, error) {
	contract, err := bindMiniSafe(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MiniSafeFilterer{contract: contract}, nil
}

// bindMiniSafe binds a generic wrapper to an already deployed contract.
func bindMiniSafe(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MiniSafeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MiniSafe *MiniSafeRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MiniSafe.Contract.MiniSafeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MiniSafe *MiniSafeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MiniSafe.Contract.MiniSafeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MiniSafe *MiniSafeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MiniSafe.Contract.MiniSafeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MiniSafe *MiniSafeCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MiniSafe.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MiniSafe *MiniSafeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MiniSafe.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MiniSafe *MiniSafeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MiniSafe.Contract.contract.Transact(opts, method, params...)
}
