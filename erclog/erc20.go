// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package erclog

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

// ErclogABI is the input ABI used to generate the binding from.
const ErclogABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"tokenOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"}]"

// Erclog is an auto generated Go binding around an Ethereum contract.
type Erclog struct {
	ErclogCaller     // Read-only binding to the contract
	ErclogTransactor // Write-only binding to the contract
	ErclogFilterer   // Log filterer for contract events
}

// ErclogCaller is an auto generated read-only Go binding around an Ethereum contract.
type ErclogCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ErclogTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ErclogTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ErclogFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ErclogFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ErclogSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ErclogSession struct {
	Contract     *Erclog           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ErclogCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ErclogCallerSession struct {
	Contract *ErclogCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ErclogTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ErclogTransactorSession struct {
	Contract     *ErclogTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ErclogRaw is an auto generated low-level Go binding around an Ethereum contract.
type ErclogRaw struct {
	Contract *Erclog // Generic contract binding to access the raw methods on
}

// ErclogCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ErclogCallerRaw struct {
	Contract *ErclogCaller // Generic read-only contract binding to access the raw methods on
}

// ErclogTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ErclogTransactorRaw struct {
	Contract *ErclogTransactor // Generic write-only contract binding to access the raw methods on
}

// NewErclog creates a new instance of Erclog, bound to a specific deployed contract.
func NewErclog(address common.Address, backend bind.ContractBackend) (*Erclog, error) {
	contract, err := bindErclog(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Erclog{ErclogCaller: ErclogCaller{contract: contract}, ErclogTransactor: ErclogTransactor{contract: contract}, ErclogFilterer: ErclogFilterer{contract: contract}}, nil
}

// NewErclogCaller creates a new read-only instance of Erclog, bound to a specific deployed contract.
func NewErclogCaller(address common.Address, caller bind.ContractCaller) (*ErclogCaller, error) {
	contract, err := bindErclog(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ErclogCaller{contract: contract}, nil
}

// NewErclogTransactor creates a new write-only instance of Erclog, bound to a specific deployed contract.
func NewErclogTransactor(address common.Address, transactor bind.ContractTransactor) (*ErclogTransactor, error) {
	contract, err := bindErclog(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ErclogTransactor{contract: contract}, nil
}

// NewErclogFilterer creates a new log filterer instance of Erclog, bound to a specific deployed contract.
func NewErclogFilterer(address common.Address, filterer bind.ContractFilterer) (*ErclogFilterer, error) {
	contract, err := bindErclog(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ErclogFilterer{contract: contract}, nil
}

// bindErclog binds a generic wrapper to an already deployed contract.
func bindErclog(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ErclogABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Erclog *ErclogRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Erclog.Contract.ErclogCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Erclog *ErclogRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Erclog.Contract.ErclogTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Erclog *ErclogRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Erclog.Contract.ErclogTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Erclog *ErclogCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Erclog.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Erclog *ErclogTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Erclog.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Erclog *ErclogTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Erclog.Contract.contract.Transact(opts, method, params...)
}

// ErclogApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Erclog contract.
type ErclogApprovalIterator struct {
	Event *ErclogApproval // Event containing the contract specifics and raw log

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
func (it *ErclogApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ErclogApproval)
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
		it.Event = new(ErclogApproval)
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
func (it *ErclogApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ErclogApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ErclogApproval represents a Approval event raised by the Erclog contract.
type ErclogApproval struct {
	TokenOwner common.Address
	Spender    common.Address
	Tokens     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed tokenOwner, address indexed spender, uint256 tokens)
func (_Erclog *ErclogFilterer) FilterApproval(opts *bind.FilterOpts, tokenOwner []common.Address, spender []common.Address) (*ErclogApprovalIterator, error) {

	var tokenOwnerRule []interface{}
	for _, tokenOwnerItem := range tokenOwner {
		tokenOwnerRule = append(tokenOwnerRule, tokenOwnerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Erclog.contract.FilterLogs(opts, "Approval", tokenOwnerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ErclogApprovalIterator{contract: _Erclog.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed tokenOwner, address indexed spender, uint256 tokens)
func (_Erclog *ErclogFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ErclogApproval, tokenOwner []common.Address, spender []common.Address) (event.Subscription, error) {

	var tokenOwnerRule []interface{}
	for _, tokenOwnerItem := range tokenOwner {
		tokenOwnerRule = append(tokenOwnerRule, tokenOwnerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Erclog.contract.WatchLogs(opts, "Approval", tokenOwnerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ErclogApproval)
				if err := _Erclog.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed tokenOwner, address indexed spender, uint256 tokens)
func (_Erclog *ErclogFilterer) ParseApproval(log types.Log) (*ErclogApproval, error) {
	event := new(ErclogApproval)
	if err := _Erclog.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ErclogTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Erclog contract.
type ErclogTransferIterator struct {
	Event *ErclogTransfer // Event containing the contract specifics and raw log

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
func (it *ErclogTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ErclogTransfer)
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
		it.Event = new(ErclogTransfer)
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
func (it *ErclogTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ErclogTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ErclogTransfer represents a Transfer event raised by the Erclog contract.
type ErclogTransfer struct {
	From   common.Address
	To     common.Address
	Tokens *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 tokens)
func (_Erclog *ErclogFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ErclogTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Erclog.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ErclogTransferIterator{contract: _Erclog.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 tokens)
func (_Erclog *ErclogFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ErclogTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Erclog.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ErclogTransfer)
				if err := _Erclog.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 tokens)
func (_Erclog *ErclogFilterer) ParseTransfer(log types.Log) (*ErclogTransfer, error) {
	event := new(ErclogTransfer)
	if err := _Erclog.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}
