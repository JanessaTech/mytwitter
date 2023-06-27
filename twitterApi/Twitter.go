// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package twitterapi

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
	_ = abi.ConvertType
)

// TwitterMetaData contains all meta data concerning the Twitter contract.
var TwitterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_userManager\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"who\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"content\",\"type\":\"string\"}],\"name\":\"postTracker\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"content\",\"type\":\"string\"}],\"name\":\"post\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"postNumbers\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"readFrom\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"readPost\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5060405162000fa638038062000fa6833981810160405281019061003491906100de565b80600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505061010b565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006100ab82610080565b9050919050565b6100bb816100a0565b81146100c657600080fd5b50565b6000815190506100d8816100b2565b92915050565b6000602082840312156100f4576100f361007b565b5b6000610102848285016100c9565b91505092915050565b610e8b806200011b6000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c80638ee93cf314610046578063d362308d14610062578063fd36206914610080575b600080fd5b610060600480360381019061005b9190610710565b6100b2565b005b61006a610326565b6040516100779190610772565b60405180910390f35b61009a60048036038101906100959190610817565b61036f565b6040516100a9939291906108d6565b60405180910390f35b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663ef8f7def336040518263ffffffff1660e01b815260040161010d919061092a565b602060405180830381865afa15801561012a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061014e919061097d565b61018d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610184906109f6565b60405180910390fd5b6000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060405180604001604052804281526020018381525090806001815401808255809150506001900390600052602060002090600202016000909190919091506000820151816000015560208201518160010190816102289190610c22565b5050506000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663908c38fd336040518263ffffffff1660e01b8152600401610288919061092a565b600060405180830381865afa1580156102a5573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f820116820180604052508101906102ce9190610d64565b505090503373ffffffffffffffffffffffffffffffffffffffff167fcce0e40e3f29aadd87f707a8710f24c35c7f6ffc02ae7b30e1ee5ce609d9e523828460405161031a929190610def565b60405180910390a25050565b60008060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002080549050905090565b60608060008060008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208054905084106103eb576000604051806020016040528060008152509060405180602001604052806000815250909250925092506105af565b60008060008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020858154811061043d5761043c610e26565b5b90600052602060002090600202016040518060400160405290816000820154815260200160018201805461047090610a45565b80601f016020809104026020016040519081016040528092919081815260200182805461049c90610a45565b80156104e95780601f106104be576101008083540402835291602001916104e9565b820191906000526020600020905b8154815290600101906020018083116104cc57829003601f168201915b50505050508152505090506000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663908c38fd886040518263ffffffff1660e01b8152600401610551919061092a565b600060405180830381865afa15801561056e573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f820116820180604052508101906105979190610d64565b50509050808260200151836000015194509450945050505b9250925092565b6000604051905090565b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b61061d826105d4565b810181811067ffffffffffffffff8211171561063c5761063b6105e5565b5b80604052505050565b600061064f6105b6565b905061065b8282610614565b919050565b600067ffffffffffffffff82111561067b5761067a6105e5565b5b610684826105d4565b9050602081019050919050565b82818337600083830152505050565b60006106b36106ae84610660565b610645565b9050828152602081018484840111156106cf576106ce6105cf565b5b6106da848285610691565b509392505050565b600082601f8301126106f7576106f66105ca565b5b81356107078482602086016106a0565b91505092915050565b600060208284031215610726576107256105c0565b5b600082013567ffffffffffffffff811115610744576107436105c5565b5b610750848285016106e2565b91505092915050565b6000819050919050565b61076c81610759565b82525050565b60006020820190506107876000830184610763565b92915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006107b88261078d565b9050919050565b6107c8816107ad565b81146107d357600080fd5b50565b6000813590506107e5816107bf565b92915050565b6107f481610759565b81146107ff57600080fd5b50565b600081359050610811816107eb565b92915050565b6000806040838503121561082e5761082d6105c0565b5b600061083c858286016107d6565b925050602061084d85828601610802565b9150509250929050565b600081519050919050565b600082825260208201905092915050565b60005b83811015610891578082015181840152602081019050610876565b60008484015250505050565b60006108a882610857565b6108b28185610862565b93506108c2818560208601610873565b6108cb816105d4565b840191505092915050565b600060608201905081810360008301526108f0818661089d565b90508181036020830152610904818561089d565b90506109136040830184610763565b949350505050565b610924816107ad565b82525050565b600060208201905061093f600083018461091b565b92915050565b60008115159050919050565b61095a81610945565b811461096557600080fd5b50565b60008151905061097781610951565b92915050565b600060208284031215610993576109926105c0565b5b60006109a184828501610968565b91505092915050565b7f4974206973206e6f74206c6f67696e6564000000000000000000000000000000600082015250565b60006109e0601183610862565b91506109eb826109aa565b602082019050919050565b60006020820190508181036000830152610a0f816109d3565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b60006002820490506001821680610a5d57607f821691505b602082108103610a7057610a6f610a16565b5b50919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b600060088302610ad87fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82610a9b565b610ae28683610a9b565b95508019841693508086168417925050509392505050565b6000819050919050565b6000610b1f610b1a610b1584610759565b610afa565b610759565b9050919050565b6000819050919050565b610b3983610b04565b610b4d610b4582610b26565b848454610aa8565b825550505050565b600090565b610b62610b55565b610b6d818484610b30565b505050565b5b81811015610b9157610b86600082610b5a565b600181019050610b73565b5050565b601f821115610bd657610ba781610a76565b610bb084610a8b565b81016020851015610bbf578190505b610bd3610bcb85610a8b565b830182610b72565b50505b505050565b600082821c905092915050565b6000610bf960001984600802610bdb565b1980831691505092915050565b6000610c128383610be8565b9150826002028217905092915050565b610c2b82610857565b67ffffffffffffffff811115610c4457610c436105e5565b5b610c4e8254610a45565b610c59828285610b95565b600060209050601f831160018114610c8c5760008415610c7a578287015190505b610c848582610c06565b865550610cec565b601f198416610c9a86610a76565b60005b82811015610cc257848901518255600182019150602085019450602081019050610c9d565b86831015610cdf5784890151610cdb601f891682610be8565b8355505b6001600288020188555050505b505050505050565b6000610d07610d0284610660565b610645565b905082815260208101848484011115610d2357610d226105cf565b5b610d2e848285610873565b509392505050565b600082601f830112610d4b57610d4a6105ca565b5b8151610d5b848260208601610cf4565b91505092915050565b600080600060608486031215610d7d57610d7c6105c0565b5b600084015167ffffffffffffffff811115610d9b57610d9a6105c5565b5b610da786828701610d36565b935050602084015167ffffffffffffffff811115610dc857610dc76105c5565b5b610dd486828701610d36565b9250506040610de586828701610968565b9150509250925092565b60006040820190508181036000830152610e09818561089d565b90508181036020830152610e1d818461089d565b90509392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fdfea2646970667358221220d92301ccb110da2cfc28f7c8af947337a6e1378f81536d028cb7eedfe0a6471364736f6c63430008130033",
}

// TwitterABI is the input ABI used to generate the binding from.
// Deprecated: Use TwitterMetaData.ABI instead.
var TwitterABI = TwitterMetaData.ABI

// TwitterBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TwitterMetaData.Bin instead.
var TwitterBin = TwitterMetaData.Bin

// DeployTwitter deploys a new Ethereum contract, binding an instance of Twitter to it.
func DeployTwitter(auth *bind.TransactOpts, backend bind.ContractBackend, _userManager common.Address) (common.Address, *types.Transaction, *Twitter, error) {
	parsed, err := TwitterMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TwitterBin), backend, _userManager)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Twitter{TwitterCaller: TwitterCaller{contract: contract}, TwitterTransactor: TwitterTransactor{contract: contract}, TwitterFilterer: TwitterFilterer{contract: contract}}, nil
}

// Twitter is an auto generated Go binding around an Ethereum contract.
type Twitter struct {
	TwitterCaller     // Read-only binding to the contract
	TwitterTransactor // Write-only binding to the contract
	TwitterFilterer   // Log filterer for contract events
}

// TwitterCaller is an auto generated read-only Go binding around an Ethereum contract.
type TwitterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TwitterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TwitterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TwitterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TwitterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TwitterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TwitterSession struct {
	Contract     *Twitter          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TwitterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TwitterCallerSession struct {
	Contract *TwitterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// TwitterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TwitterTransactorSession struct {
	Contract     *TwitterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// TwitterRaw is an auto generated low-level Go binding around an Ethereum contract.
type TwitterRaw struct {
	Contract *Twitter // Generic contract binding to access the raw methods on
}

// TwitterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TwitterCallerRaw struct {
	Contract *TwitterCaller // Generic read-only contract binding to access the raw methods on
}

// TwitterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TwitterTransactorRaw struct {
	Contract *TwitterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTwitter creates a new instance of Twitter, bound to a specific deployed contract.
func NewTwitter(address common.Address, backend bind.ContractBackend) (*Twitter, error) {
	contract, err := bindTwitter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Twitter{TwitterCaller: TwitterCaller{contract: contract}, TwitterTransactor: TwitterTransactor{contract: contract}, TwitterFilterer: TwitterFilterer{contract: contract}}, nil
}

// NewTwitterCaller creates a new read-only instance of Twitter, bound to a specific deployed contract.
func NewTwitterCaller(address common.Address, caller bind.ContractCaller) (*TwitterCaller, error) {
	contract, err := bindTwitter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TwitterCaller{contract: contract}, nil
}

// NewTwitterTransactor creates a new write-only instance of Twitter, bound to a specific deployed contract.
func NewTwitterTransactor(address common.Address, transactor bind.ContractTransactor) (*TwitterTransactor, error) {
	contract, err := bindTwitter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TwitterTransactor{contract: contract}, nil
}

// NewTwitterFilterer creates a new log filterer instance of Twitter, bound to a specific deployed contract.
func NewTwitterFilterer(address common.Address, filterer bind.ContractFilterer) (*TwitterFilterer, error) {
	contract, err := bindTwitter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TwitterFilterer{contract: contract}, nil
}

// bindTwitter binds a generic wrapper to an already deployed contract.
func bindTwitter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TwitterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Twitter *TwitterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Twitter.Contract.TwitterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Twitter *TwitterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Twitter.Contract.TwitterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Twitter *TwitterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Twitter.Contract.TwitterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Twitter *TwitterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Twitter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Twitter *TwitterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Twitter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Twitter *TwitterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Twitter.Contract.contract.Transact(opts, method, params...)
}

// PostNumbers is a free data retrieval call binding the contract method 0xd362308d.
//
// Solidity: function postNumbers() view returns(uint256)
func (_Twitter *TwitterCaller) PostNumbers(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Twitter.contract.Call(opts, &out, "postNumbers")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PostNumbers is a free data retrieval call binding the contract method 0xd362308d.
//
// Solidity: function postNumbers() view returns(uint256)
func (_Twitter *TwitterSession) PostNumbers() (*big.Int, error) {
	return _Twitter.Contract.PostNumbers(&_Twitter.CallOpts)
}

// PostNumbers is a free data retrieval call binding the contract method 0xd362308d.
//
// Solidity: function postNumbers() view returns(uint256)
func (_Twitter *TwitterCallerSession) PostNumbers() (*big.Int, error) {
	return _Twitter.Contract.PostNumbers(&_Twitter.CallOpts)
}

// ReadPost is a free data retrieval call binding the contract method 0xfd362069.
//
// Solidity: function readPost(address readFrom, uint256 index) view returns(string, string, uint256)
func (_Twitter *TwitterCaller) ReadPost(opts *bind.CallOpts, readFrom common.Address, index *big.Int) (string, string, *big.Int, error) {
	var out []interface{}
	err := _Twitter.contract.Call(opts, &out, "readPost", readFrom, index)

	if err != nil {
		return *new(string), *new(string), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	out1 := *abi.ConvertType(out[1], new(string)).(*string)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return out0, out1, out2, err

}

// ReadPost is a free data retrieval call binding the contract method 0xfd362069.
//
// Solidity: function readPost(address readFrom, uint256 index) view returns(string, string, uint256)
func (_Twitter *TwitterSession) ReadPost(readFrom common.Address, index *big.Int) (string, string, *big.Int, error) {
	return _Twitter.Contract.ReadPost(&_Twitter.CallOpts, readFrom, index)
}

// ReadPost is a free data retrieval call binding the contract method 0xfd362069.
//
// Solidity: function readPost(address readFrom, uint256 index) view returns(string, string, uint256)
func (_Twitter *TwitterCallerSession) ReadPost(readFrom common.Address, index *big.Int) (string, string, *big.Int, error) {
	return _Twitter.Contract.ReadPost(&_Twitter.CallOpts, readFrom, index)
}

// Post is a paid mutator transaction binding the contract method 0x8ee93cf3.
//
// Solidity: function post(string content) returns()
func (_Twitter *TwitterTransactor) Post(opts *bind.TransactOpts, content string) (*types.Transaction, error) {
	return _Twitter.contract.Transact(opts, "post", content)
}

// Post is a paid mutator transaction binding the contract method 0x8ee93cf3.
//
// Solidity: function post(string content) returns()
func (_Twitter *TwitterSession) Post(content string) (*types.Transaction, error) {
	return _Twitter.Contract.Post(&_Twitter.TransactOpts, content)
}

// Post is a paid mutator transaction binding the contract method 0x8ee93cf3.
//
// Solidity: function post(string content) returns()
func (_Twitter *TwitterTransactorSession) Post(content string) (*types.Transaction, error) {
	return _Twitter.Contract.Post(&_Twitter.TransactOpts, content)
}

// TwitterPostTrackerIterator is returned from FilterPostTracker and is used to iterate over the raw logs and unpacked data for PostTracker events raised by the Twitter contract.
type TwitterPostTrackerIterator struct {
	Event *TwitterPostTracker // Event containing the contract specifics and raw log

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
func (it *TwitterPostTrackerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TwitterPostTracker)
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
		it.Event = new(TwitterPostTracker)
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
func (it *TwitterPostTrackerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TwitterPostTrackerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TwitterPostTracker represents a PostTracker event raised by the Twitter contract.
type TwitterPostTracker struct {
	From    common.Address
	Who     string
	Content string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPostTracker is a free log retrieval operation binding the contract event 0xcce0e40e3f29aadd87f707a8710f24c35c7f6ffc02ae7b30e1ee5ce609d9e523.
//
// Solidity: event postTracker(address indexed from, string who, string content)
func (_Twitter *TwitterFilterer) FilterPostTracker(opts *bind.FilterOpts, from []common.Address) (*TwitterPostTrackerIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _Twitter.contract.FilterLogs(opts, "postTracker", fromRule)
	if err != nil {
		return nil, err
	}
	return &TwitterPostTrackerIterator{contract: _Twitter.contract, event: "postTracker", logs: logs, sub: sub}, nil
}

// WatchPostTracker is a free log subscription operation binding the contract event 0xcce0e40e3f29aadd87f707a8710f24c35c7f6ffc02ae7b30e1ee5ce609d9e523.
//
// Solidity: event postTracker(address indexed from, string who, string content)
func (_Twitter *TwitterFilterer) WatchPostTracker(opts *bind.WatchOpts, sink chan<- *TwitterPostTracker, from []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _Twitter.contract.WatchLogs(opts, "postTracker", fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TwitterPostTracker)
				if err := _Twitter.contract.UnpackLog(event, "postTracker", log); err != nil {
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

// ParsePostTracker is a log parse operation binding the contract event 0xcce0e40e3f29aadd87f707a8710f24c35c7f6ffc02ae7b30e1ee5ce609d9e523.
//
// Solidity: event postTracker(address indexed from, string who, string content)
func (_Twitter *TwitterFilterer) ParsePostTracker(log types.Log) (*TwitterPostTracker, error) {
	event := new(TwitterPostTracker)
	if err := _Twitter.contract.UnpackLog(event, "postTracker", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
