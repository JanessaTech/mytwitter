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

// TwitterUserMetaData contains all meta data concerning the TwitterUser contract.
var TwitterUserMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"userAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"pswd1\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"pswd2\",\"type\":\"string\"}],\"name\":\"userLoginTracker\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"userAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"action\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"password\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"userNumber\",\"type\":\"uint256\"}],\"name\":\"userTracker\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"userAddress\",\"type\":\"address\"}],\"name\":\"userUnloginTracker\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"userAddress\",\"type\":\"address\"}],\"name\":\"checkIfUserLogined\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"userAddress\",\"type\":\"address\"}],\"name\":\"checkUserInfo\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"checkUserNumbers\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"userAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"password\",\"type\":\"string\"}],\"name\":\"login\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"userAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"password\",\"type\":\"string\"}],\"name\":\"register\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"userAddress\",\"type\":\"address\"}],\"name\":\"unlogin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"userAddress\",\"type\":\"address\"}],\"name\":\"unregister\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5033600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550611bc6806100616000396000f3fe608060405234801561001057600080fd5b506004361061007d5760003560e01c80638d59cc021161005b5780638d59cc02146100d6578063908c38fd146100f2578063be6ba1f214610124578063ef8f7def146101425761007d565b80632ec2c246146100825780632f19d50b1461009e5780638b0b3bc0146100ba575b600080fd5b61009c60048036038101906100979190610f7c565b610172565b005b6100b860048036038101906100b391906110ef565b6103da565b005b6100d460048036038101906100cf9190610f7c565b610799565b005b6100f060048036038101906100eb919061114b565b610913565b005b61010c60048036038101906101079190610f7c565b610be3565b60405161011b93929190611270565b60405180910390f35b61012c610e16565b60405161013991906112ce565b60405180910390f35b61015c60048036038101906101579190610f7c565b610e20565b60405161016991906112e9565b60405180910390f35b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610202576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016101f990611350565b60405180910390fd5b60008060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000180546102509061139f565b905011156103d757604051806020016040528060008152506000808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000190816102b5919061157c565b50604051806020016040528060008152506000808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206001019081610313919061157c565b5060008060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060020160006101000a81548160ff021916908315150217905550600160008154809291906103819061167d565b91905055508073ffffffffffffffffffffffffffffffffffffffff167f699b600a189c6432050a208e515648f4b2b9de0017ba7daf5321f6c0160a02746001546040516103ce9190611718565b60405180910390a25b50565b8160008060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000180546104299061139f565b90501161046b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610462906117b8565b60405180910390fd5b60008060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000180546104b99061139f565b905011156107945760008060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060405180606001604052908160008201805461051c9061139f565b80601f01602080910402602001604051908101604052809291908181526020018280546105489061139f565b80156105955780601f1061056a57610100808354040283529160200191610595565b820191906000526020600020905b81548152906001019060200180831161057857829003601f168201915b505050505081526020016001820180546105ae9061139f565b80601f01602080910402602001604051908101604052809291908181526020018280546105da9061139f565b80156106275780601f106105fc57610100808354040283529160200191610627565b820191906000526020600020905b81548152906001019060200180831161060a57829003601f168201915b505050505081526020016002820160009054906101000a900460ff16151515158152505090508260405160200161065e9190611814565b6040516020818303038152906040528051906020012081602001516040516020016106899190611814565b60405160208183030381529060405280519060200120036107575760016000808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060020160006101000a81548160ff0219169083151502179055508373ffffffffffffffffffffffffffffffffffffffff167f2394534d4bf17a42e5eb41fa37f127ac858aef333f0b138a98abd12187f7084084836020015160405161074a92919061182b565b60405180910390a2610792565b6040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610789906118d4565b60405180910390fd5b505b505050565b8060008060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000180546107e89061139f565b90501161082a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610821906117b8565b60405180910390fd5b61083382610e20565b610872576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161086990611940565b60405180910390fd5b60008060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060020160006101000a81548160ff0219169083151502179055508173ffffffffffffffffffffffffffffffffffffffff167f731ee42bed785b2d551eff7ed95419d92fdc48f6f18af20a463a4f2e860cf90d60405160405180910390a25050565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146109a3576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161099a90611350565b60405180910390fd5b8260008060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000180546109f29061139f565b905014610a34576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a2b906119ac565b60405180910390fd5b82826000825111610a7a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a7190611a18565b60405180910390fd5b6000815111610abe576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ab590611a84565b60405180910390fd5b60006040518060600160405280878152602001868152602001600015158152509050806000808973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000820151816000019081610b34919061157c565b506020820151816001019081610b4a919061157c565b5060408201518160020160006101000a81548160ff02191690831515021790555090505060016000815480929190610b8190611aa4565b91905055508673ffffffffffffffffffffffffffffffffffffffff167f699b600a189c6432050a208e515648f4b2b9de0017ba7daf5321f6c0160a02748787600154604051610bd293929190611b38565b60405180910390a250505050505050565b6060806000806000808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000018054610c359061139f565b90501115610de45760008060008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020604051806060016040529081600082018054610c989061139f565b80601f0160208091040260200160405190810160405280929190818152602001828054610cc49061139f565b8015610d115780601f10610ce657610100808354040283529160200191610d11565b820191906000526020600020905b815481529060010190602001808311610cf457829003601f168201915b50505050508152602001600182018054610d2a9061139f565b80601f0160208091040260200160405190810160405280929190818152602001828054610d569061139f565b8015610da35780601f10610d7857610100808354040283529160200191610da3565b820191906000526020600020905b815481529060010190602001808311610d8657829003601f168201915b505050505081526020016002820160009054906101000a900460ff161515151581525050905080600001518160200151826040015193509350935050610e0f565b6000604051806020016040528060008152509060405180602001604052806000815250909250925092505b9193909250565b6000600154905090565b60008160008060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000018054610e719061139f565b905011610eb3576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610eaa906117b8565b60405180910390fd5b6000808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060020160009054906101000a900460ff16915050919050565b6000604051905090565b600080fd5b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000610f4982610f1e565b9050919050565b610f5981610f3e565b8114610f6457600080fd5b50565b600081359050610f7681610f50565b92915050565b600060208284031215610f9257610f91610f14565b5b6000610fa084828501610f67565b91505092915050565b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b610ffc82610fb3565b810181811067ffffffffffffffff8211171561101b5761101a610fc4565b5b80604052505050565b600061102e610f0a565b905061103a8282610ff3565b919050565b600067ffffffffffffffff82111561105a57611059610fc4565b5b61106382610fb3565b9050602081019050919050565b82818337600083830152505050565b600061109261108d8461103f565b611024565b9050828152602081018484840111156110ae576110ad610fae565b5b6110b9848285611070565b509392505050565b600082601f8301126110d6576110d5610fa9565b5b81356110e684826020860161107f565b91505092915050565b6000806040838503121561110657611105610f14565b5b600061111485828601610f67565b925050602083013567ffffffffffffffff81111561113557611134610f19565b5b611141858286016110c1565b9150509250929050565b60008060006060848603121561116457611163610f14565b5b600061117286828701610f67565b935050602084013567ffffffffffffffff81111561119357611192610f19565b5b61119f868287016110c1565b925050604084013567ffffffffffffffff8111156111c0576111bf610f19565b5b6111cc868287016110c1565b9150509250925092565b600081519050919050565b600082825260208201905092915050565b60005b838110156112105780820151818401526020810190506111f5565b60008484015250505050565b6000611227826111d6565b61123181856111e1565b93506112418185602086016111f2565b61124a81610fb3565b840191505092915050565b60008115159050919050565b61126a81611255565b82525050565b6000606082019050818103600083015261128a818661121c565b9050818103602083015261129e818561121c565b90506112ad6040830184611261565b949350505050565b6000819050919050565b6112c8816112b5565b82525050565b60006020820190506112e360008301846112bf565b92915050565b60006020820190506112fe6000830184611261565b92915050565b7f4f6e6c792061646d696e2063616e20646f2074686973206f7065726174696f6e600082015250565b600061133a6020836111e1565b915061134582611304565b602082019050919050565b600060208201905081810360008301526113698161132d565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b600060028204905060018216806113b757607f821691505b6020821081036113ca576113c9611370565b5b50919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b6000600883026114327fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff826113f5565b61143c86836113f5565b95508019841693508086168417925050509392505050565b6000819050919050565b600061147961147461146f846112b5565b611454565b6112b5565b9050919050565b6000819050919050565b6114938361145e565b6114a761149f82611480565b848454611402565b825550505050565b600090565b6114bc6114af565b6114c781848461148a565b505050565b5b818110156114eb576114e06000826114b4565b6001810190506114cd565b5050565b601f82111561153057611501816113d0565b61150a846113e5565b81016020851015611519578190505b61152d611525856113e5565b8301826114cc565b50505b505050565b600082821c905092915050565b600061155360001984600802611535565b1980831691505092915050565b600061156c8383611542565b9150826002028217905092915050565b611585826111d6565b67ffffffffffffffff81111561159e5761159d610fc4565b5b6115a8825461139f565b6115b38282856114ef565b600060209050601f8311600181146115e657600084156115d4578287015190505b6115de8582611560565b865550611646565b601f1984166115f4866113d0565b60005b8281101561161c578489015182556001820191506020850194506020810190506115f7565b868310156116395784890151611635601f891682611542565b8355505b6001600288020188555050505b505050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000611688826112b5565b91506000820361169b5761169a61164e565b5b600182039050919050565b7f756e726567697374657200000000000000000000000000000000000000000000600082015250565b60006116dc600a836111e1565b91506116e7826116a6565b602082019050919050565b50565b60006117026000836111e1565b915061170d826116f2565b600082019050919050565b60006080820190508181036000830152611731816116cf565b90508181036020830152611744816116f5565b90508181036040830152611757816116f5565b905061176660608301846112bf565b92915050565b7f61646472657373206973206e6f74207265676973746572656400000000000000600082015250565b60006117a26019836111e1565b91506117ad8261176c565b602082019050919050565b600060208201905081810360008301526117d181611795565b9050919050565b600081905092915050565b60006117ee826111d6565b6117f881856117d8565b93506118088185602086016111f2565b80840191505092915050565b600061182082846117e3565b915081905092915050565b60006040820190508181036000830152611845818561121c565b90508181036020830152611859818461121c565b90509392505050565b7f6661696c656420746f206c6f67696e2062656361757365206f6620746865206d60008201527f697374636865642070617373776f726400000000000000000000000000000000602082015250565b60006118be6030836111e1565b91506118c982611862565b604082019050919050565b600060208201905081810360008301526118ed816118b1565b9050919050565b7f6661696c656420746f20756e6c6f67696e000000000000000000000000000000600082015250565b600061192a6011836111e1565b9150611935826118f4565b602082019050919050565b600060208201905081810360008301526119598161191d565b9050919050565b7f6164647265737320697320726567697374657265640000000000000000000000600082015250565b60006119966015836111e1565b91506119a182611960565b602082019050919050565b600060208201905081810360008301526119c581611989565b9050919050565b7f6e616d652063616e6e6f7420626520656d707479000000000000000000000000600082015250565b6000611a026014836111e1565b9150611a0d826119cc565b602082019050919050565b60006020820190508181036000830152611a31816119f5565b9050919050565b7f70617373776f72642063616e6e6f7420626520656d7074790000000000000000600082015250565b6000611a6e6018836111e1565b9150611a7982611a38565b602082019050919050565b60006020820190508181036000830152611a9d81611a61565b9050919050565b6000611aaf826112b5565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203611ae157611ae061164e565b5b600182019050919050565b7f7265676973746572000000000000000000000000000000000000000000000000600082015250565b6000611b226008836111e1565b9150611b2d82611aec565b602082019050919050565b60006080820190508181036000830152611b5181611b15565b90508181036020830152611b65818661121c565b90508181036040830152611b79818561121c565b9050611b8860608301846112bf565b94935050505056fea264697066735822122091a468d261d6c94f64a0a86685ab57277e232b11b50f0210feced98c7ce8ba8f64736f6c63430008130033",
}

// TwitterUserABI is the input ABI used to generate the binding from.
// Deprecated: Use TwitterUserMetaData.ABI instead.
var TwitterUserABI = TwitterUserMetaData.ABI

// TwitterUserBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TwitterUserMetaData.Bin instead.
var TwitterUserBin = TwitterUserMetaData.Bin

// DeployTwitterUser deploys a new Ethereum contract, binding an instance of TwitterUser to it.
func DeployTwitterUser(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TwitterUser, error) {
	parsed, err := TwitterUserMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TwitterUserBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TwitterUser{TwitterUserCaller: TwitterUserCaller{contract: contract}, TwitterUserTransactor: TwitterUserTransactor{contract: contract}, TwitterUserFilterer: TwitterUserFilterer{contract: contract}}, nil
}

// TwitterUser is an auto generated Go binding around an Ethereum contract.
type TwitterUser struct {
	TwitterUserCaller     // Read-only binding to the contract
	TwitterUserTransactor // Write-only binding to the contract
	TwitterUserFilterer   // Log filterer for contract events
}

// TwitterUserCaller is an auto generated read-only Go binding around an Ethereum contract.
type TwitterUserCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TwitterUserTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TwitterUserTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TwitterUserFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TwitterUserFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TwitterUserSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TwitterUserSession struct {
	Contract     *TwitterUser      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TwitterUserCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TwitterUserCallerSession struct {
	Contract *TwitterUserCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// TwitterUserTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TwitterUserTransactorSession struct {
	Contract     *TwitterUserTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// TwitterUserRaw is an auto generated low-level Go binding around an Ethereum contract.
type TwitterUserRaw struct {
	Contract *TwitterUser // Generic contract binding to access the raw methods on
}

// TwitterUserCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TwitterUserCallerRaw struct {
	Contract *TwitterUserCaller // Generic read-only contract binding to access the raw methods on
}

// TwitterUserTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TwitterUserTransactorRaw struct {
	Contract *TwitterUserTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTwitterUser creates a new instance of TwitterUser, bound to a specific deployed contract.
func NewTwitterUser(address common.Address, backend bind.ContractBackend) (*TwitterUser, error) {
	contract, err := bindTwitterUser(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TwitterUser{TwitterUserCaller: TwitterUserCaller{contract: contract}, TwitterUserTransactor: TwitterUserTransactor{contract: contract}, TwitterUserFilterer: TwitterUserFilterer{contract: contract}}, nil
}

// NewTwitterUserCaller creates a new read-only instance of TwitterUser, bound to a specific deployed contract.
func NewTwitterUserCaller(address common.Address, caller bind.ContractCaller) (*TwitterUserCaller, error) {
	contract, err := bindTwitterUser(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TwitterUserCaller{contract: contract}, nil
}

// NewTwitterUserTransactor creates a new write-only instance of TwitterUser, bound to a specific deployed contract.
func NewTwitterUserTransactor(address common.Address, transactor bind.ContractTransactor) (*TwitterUserTransactor, error) {
	contract, err := bindTwitterUser(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TwitterUserTransactor{contract: contract}, nil
}

// NewTwitterUserFilterer creates a new log filterer instance of TwitterUser, bound to a specific deployed contract.
func NewTwitterUserFilterer(address common.Address, filterer bind.ContractFilterer) (*TwitterUserFilterer, error) {
	contract, err := bindTwitterUser(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TwitterUserFilterer{contract: contract}, nil
}

// bindTwitterUser binds a generic wrapper to an already deployed contract.
func bindTwitterUser(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TwitterUserMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TwitterUser *TwitterUserRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TwitterUser.Contract.TwitterUserCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TwitterUser *TwitterUserRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TwitterUser.Contract.TwitterUserTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TwitterUser *TwitterUserRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TwitterUser.Contract.TwitterUserTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TwitterUser *TwitterUserCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TwitterUser.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TwitterUser *TwitterUserTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TwitterUser.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TwitterUser *TwitterUserTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TwitterUser.Contract.contract.Transact(opts, method, params...)
}

// CheckIfUserLogined is a free data retrieval call binding the contract method 0xef8f7def.
//
// Solidity: function checkIfUserLogined(address userAddress) view returns(bool)
func (_TwitterUser *TwitterUserCaller) CheckIfUserLogined(opts *bind.CallOpts, userAddress common.Address) (bool, error) {
	var out []interface{}
	err := _TwitterUser.contract.Call(opts, &out, "checkIfUserLogined", userAddress)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckIfUserLogined is a free data retrieval call binding the contract method 0xef8f7def.
//
// Solidity: function checkIfUserLogined(address userAddress) view returns(bool)
func (_TwitterUser *TwitterUserSession) CheckIfUserLogined(userAddress common.Address) (bool, error) {
	return _TwitterUser.Contract.CheckIfUserLogined(&_TwitterUser.CallOpts, userAddress)
}

// CheckIfUserLogined is a free data retrieval call binding the contract method 0xef8f7def.
//
// Solidity: function checkIfUserLogined(address userAddress) view returns(bool)
func (_TwitterUser *TwitterUserCallerSession) CheckIfUserLogined(userAddress common.Address) (bool, error) {
	return _TwitterUser.Contract.CheckIfUserLogined(&_TwitterUser.CallOpts, userAddress)
}

// CheckUserInfo is a free data retrieval call binding the contract method 0x908c38fd.
//
// Solidity: function checkUserInfo(address userAddress) view returns(string, string, bool)
func (_TwitterUser *TwitterUserCaller) CheckUserInfo(opts *bind.CallOpts, userAddress common.Address) (string, string, bool, error) {
	var out []interface{}
	err := _TwitterUser.contract.Call(opts, &out, "checkUserInfo", userAddress)

	if err != nil {
		return *new(string), *new(string), *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	out1 := *abi.ConvertType(out[1], new(string)).(*string)
	out2 := *abi.ConvertType(out[2], new(bool)).(*bool)

	return out0, out1, out2, err

}

// CheckUserInfo is a free data retrieval call binding the contract method 0x908c38fd.
//
// Solidity: function checkUserInfo(address userAddress) view returns(string, string, bool)
func (_TwitterUser *TwitterUserSession) CheckUserInfo(userAddress common.Address) (string, string, bool, error) {
	return _TwitterUser.Contract.CheckUserInfo(&_TwitterUser.CallOpts, userAddress)
}

// CheckUserInfo is a free data retrieval call binding the contract method 0x908c38fd.
//
// Solidity: function checkUserInfo(address userAddress) view returns(string, string, bool)
func (_TwitterUser *TwitterUserCallerSession) CheckUserInfo(userAddress common.Address) (string, string, bool, error) {
	return _TwitterUser.Contract.CheckUserInfo(&_TwitterUser.CallOpts, userAddress)
}

// CheckUserNumbers is a free data retrieval call binding the contract method 0xbe6ba1f2.
//
// Solidity: function checkUserNumbers() view returns(uint256)
func (_TwitterUser *TwitterUserCaller) CheckUserNumbers(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TwitterUser.contract.Call(opts, &out, "checkUserNumbers")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CheckUserNumbers is a free data retrieval call binding the contract method 0xbe6ba1f2.
//
// Solidity: function checkUserNumbers() view returns(uint256)
func (_TwitterUser *TwitterUserSession) CheckUserNumbers() (*big.Int, error) {
	return _TwitterUser.Contract.CheckUserNumbers(&_TwitterUser.CallOpts)
}

// CheckUserNumbers is a free data retrieval call binding the contract method 0xbe6ba1f2.
//
// Solidity: function checkUserNumbers() view returns(uint256)
func (_TwitterUser *TwitterUserCallerSession) CheckUserNumbers() (*big.Int, error) {
	return _TwitterUser.Contract.CheckUserNumbers(&_TwitterUser.CallOpts)
}

// Login is a paid mutator transaction binding the contract method 0x2f19d50b.
//
// Solidity: function login(address userAddress, string password) returns()
func (_TwitterUser *TwitterUserTransactor) Login(opts *bind.TransactOpts, userAddress common.Address, password string) (*types.Transaction, error) {
	return _TwitterUser.contract.Transact(opts, "login", userAddress, password)
}

// Login is a paid mutator transaction binding the contract method 0x2f19d50b.
//
// Solidity: function login(address userAddress, string password) returns()
func (_TwitterUser *TwitterUserSession) Login(userAddress common.Address, password string) (*types.Transaction, error) {
	return _TwitterUser.Contract.Login(&_TwitterUser.TransactOpts, userAddress, password)
}

// Login is a paid mutator transaction binding the contract method 0x2f19d50b.
//
// Solidity: function login(address userAddress, string password) returns()
func (_TwitterUser *TwitterUserTransactorSession) Login(userAddress common.Address, password string) (*types.Transaction, error) {
	return _TwitterUser.Contract.Login(&_TwitterUser.TransactOpts, userAddress, password)
}

// Register is a paid mutator transaction binding the contract method 0x8d59cc02.
//
// Solidity: function register(address userAddress, string name, string password) returns()
func (_TwitterUser *TwitterUserTransactor) Register(opts *bind.TransactOpts, userAddress common.Address, name string, password string) (*types.Transaction, error) {
	return _TwitterUser.contract.Transact(opts, "register", userAddress, name, password)
}

// Register is a paid mutator transaction binding the contract method 0x8d59cc02.
//
// Solidity: function register(address userAddress, string name, string password) returns()
func (_TwitterUser *TwitterUserSession) Register(userAddress common.Address, name string, password string) (*types.Transaction, error) {
	return _TwitterUser.Contract.Register(&_TwitterUser.TransactOpts, userAddress, name, password)
}

// Register is a paid mutator transaction binding the contract method 0x8d59cc02.
//
// Solidity: function register(address userAddress, string name, string password) returns()
func (_TwitterUser *TwitterUserTransactorSession) Register(userAddress common.Address, name string, password string) (*types.Transaction, error) {
	return _TwitterUser.Contract.Register(&_TwitterUser.TransactOpts, userAddress, name, password)
}

// Unlogin is a paid mutator transaction binding the contract method 0x8b0b3bc0.
//
// Solidity: function unlogin(address userAddress) returns()
func (_TwitterUser *TwitterUserTransactor) Unlogin(opts *bind.TransactOpts, userAddress common.Address) (*types.Transaction, error) {
	return _TwitterUser.contract.Transact(opts, "unlogin", userAddress)
}

// Unlogin is a paid mutator transaction binding the contract method 0x8b0b3bc0.
//
// Solidity: function unlogin(address userAddress) returns()
func (_TwitterUser *TwitterUserSession) Unlogin(userAddress common.Address) (*types.Transaction, error) {
	return _TwitterUser.Contract.Unlogin(&_TwitterUser.TransactOpts, userAddress)
}

// Unlogin is a paid mutator transaction binding the contract method 0x8b0b3bc0.
//
// Solidity: function unlogin(address userAddress) returns()
func (_TwitterUser *TwitterUserTransactorSession) Unlogin(userAddress common.Address) (*types.Transaction, error) {
	return _TwitterUser.Contract.Unlogin(&_TwitterUser.TransactOpts, userAddress)
}

// Unregister is a paid mutator transaction binding the contract method 0x2ec2c246.
//
// Solidity: function unregister(address userAddress) returns()
func (_TwitterUser *TwitterUserTransactor) Unregister(opts *bind.TransactOpts, userAddress common.Address) (*types.Transaction, error) {
	return _TwitterUser.contract.Transact(opts, "unregister", userAddress)
}

// Unregister is a paid mutator transaction binding the contract method 0x2ec2c246.
//
// Solidity: function unregister(address userAddress) returns()
func (_TwitterUser *TwitterUserSession) Unregister(userAddress common.Address) (*types.Transaction, error) {
	return _TwitterUser.Contract.Unregister(&_TwitterUser.TransactOpts, userAddress)
}

// Unregister is a paid mutator transaction binding the contract method 0x2ec2c246.
//
// Solidity: function unregister(address userAddress) returns()
func (_TwitterUser *TwitterUserTransactorSession) Unregister(userAddress common.Address) (*types.Transaction, error) {
	return _TwitterUser.Contract.Unregister(&_TwitterUser.TransactOpts, userAddress)
}

// TwitterUserUserLoginTrackerIterator is returned from FilterUserLoginTracker and is used to iterate over the raw logs and unpacked data for UserLoginTracker events raised by the TwitterUser contract.
type TwitterUserUserLoginTrackerIterator struct {
	Event *TwitterUserUserLoginTracker // Event containing the contract specifics and raw log

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
func (it *TwitterUserUserLoginTrackerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TwitterUserUserLoginTracker)
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
		it.Event = new(TwitterUserUserLoginTracker)
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
func (it *TwitterUserUserLoginTrackerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TwitterUserUserLoginTrackerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TwitterUserUserLoginTracker represents a UserLoginTracker event raised by the TwitterUser contract.
type TwitterUserUserLoginTracker struct {
	UserAddress common.Address
	Pswd1       string
	Pswd2       string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterUserLoginTracker is a free log retrieval operation binding the contract event 0x2394534d4bf17a42e5eb41fa37f127ac858aef333f0b138a98abd12187f70840.
//
// Solidity: event userLoginTracker(address indexed userAddress, string pswd1, string pswd2)
func (_TwitterUser *TwitterUserFilterer) FilterUserLoginTracker(opts *bind.FilterOpts, userAddress []common.Address) (*TwitterUserUserLoginTrackerIterator, error) {

	var userAddressRule []interface{}
	for _, userAddressItem := range userAddress {
		userAddressRule = append(userAddressRule, userAddressItem)
	}

	logs, sub, err := _TwitterUser.contract.FilterLogs(opts, "userLoginTracker", userAddressRule)
	if err != nil {
		return nil, err
	}
	return &TwitterUserUserLoginTrackerIterator{contract: _TwitterUser.contract, event: "userLoginTracker", logs: logs, sub: sub}, nil
}

// WatchUserLoginTracker is a free log subscription operation binding the contract event 0x2394534d4bf17a42e5eb41fa37f127ac858aef333f0b138a98abd12187f70840.
//
// Solidity: event userLoginTracker(address indexed userAddress, string pswd1, string pswd2)
func (_TwitterUser *TwitterUserFilterer) WatchUserLoginTracker(opts *bind.WatchOpts, sink chan<- *TwitterUserUserLoginTracker, userAddress []common.Address) (event.Subscription, error) {

	var userAddressRule []interface{}
	for _, userAddressItem := range userAddress {
		userAddressRule = append(userAddressRule, userAddressItem)
	}

	logs, sub, err := _TwitterUser.contract.WatchLogs(opts, "userLoginTracker", userAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TwitterUserUserLoginTracker)
				if err := _TwitterUser.contract.UnpackLog(event, "userLoginTracker", log); err != nil {
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

// ParseUserLoginTracker is a log parse operation binding the contract event 0x2394534d4bf17a42e5eb41fa37f127ac858aef333f0b138a98abd12187f70840.
//
// Solidity: event userLoginTracker(address indexed userAddress, string pswd1, string pswd2)
func (_TwitterUser *TwitterUserFilterer) ParseUserLoginTracker(log types.Log) (*TwitterUserUserLoginTracker, error) {
	event := new(TwitterUserUserLoginTracker)
	if err := _TwitterUser.contract.UnpackLog(event, "userLoginTracker", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TwitterUserUserTrackerIterator is returned from FilterUserTracker and is used to iterate over the raw logs and unpacked data for UserTracker events raised by the TwitterUser contract.
type TwitterUserUserTrackerIterator struct {
	Event *TwitterUserUserTracker // Event containing the contract specifics and raw log

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
func (it *TwitterUserUserTrackerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TwitterUserUserTracker)
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
		it.Event = new(TwitterUserUserTracker)
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
func (it *TwitterUserUserTrackerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TwitterUserUserTrackerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TwitterUserUserTracker represents a UserTracker event raised by the TwitterUser contract.
type TwitterUserUserTracker struct {
	UserAddress common.Address
	Action      string
	Name        string
	Password    string
	UserNumber  *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterUserTracker is a free log retrieval operation binding the contract event 0x699b600a189c6432050a208e515648f4b2b9de0017ba7daf5321f6c0160a0274.
//
// Solidity: event userTracker(address indexed userAddress, string action, string name, string password, uint256 userNumber)
func (_TwitterUser *TwitterUserFilterer) FilterUserTracker(opts *bind.FilterOpts, userAddress []common.Address) (*TwitterUserUserTrackerIterator, error) {

	var userAddressRule []interface{}
	for _, userAddressItem := range userAddress {
		userAddressRule = append(userAddressRule, userAddressItem)
	}

	logs, sub, err := _TwitterUser.contract.FilterLogs(opts, "userTracker", userAddressRule)
	if err != nil {
		return nil, err
	}
	return &TwitterUserUserTrackerIterator{contract: _TwitterUser.contract, event: "userTracker", logs: logs, sub: sub}, nil
}

// WatchUserTracker is a free log subscription operation binding the contract event 0x699b600a189c6432050a208e515648f4b2b9de0017ba7daf5321f6c0160a0274.
//
// Solidity: event userTracker(address indexed userAddress, string action, string name, string password, uint256 userNumber)
func (_TwitterUser *TwitterUserFilterer) WatchUserTracker(opts *bind.WatchOpts, sink chan<- *TwitterUserUserTracker, userAddress []common.Address) (event.Subscription, error) {

	var userAddressRule []interface{}
	for _, userAddressItem := range userAddress {
		userAddressRule = append(userAddressRule, userAddressItem)
	}

	logs, sub, err := _TwitterUser.contract.WatchLogs(opts, "userTracker", userAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TwitterUserUserTracker)
				if err := _TwitterUser.contract.UnpackLog(event, "userTracker", log); err != nil {
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

// ParseUserTracker is a log parse operation binding the contract event 0x699b600a189c6432050a208e515648f4b2b9de0017ba7daf5321f6c0160a0274.
//
// Solidity: event userTracker(address indexed userAddress, string action, string name, string password, uint256 userNumber)
func (_TwitterUser *TwitterUserFilterer) ParseUserTracker(log types.Log) (*TwitterUserUserTracker, error) {
	event := new(TwitterUserUserTracker)
	if err := _TwitterUser.contract.UnpackLog(event, "userTracker", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TwitterUserUserUnloginTrackerIterator is returned from FilterUserUnloginTracker and is used to iterate over the raw logs and unpacked data for UserUnloginTracker events raised by the TwitterUser contract.
type TwitterUserUserUnloginTrackerIterator struct {
	Event *TwitterUserUserUnloginTracker // Event containing the contract specifics and raw log

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
func (it *TwitterUserUserUnloginTrackerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TwitterUserUserUnloginTracker)
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
		it.Event = new(TwitterUserUserUnloginTracker)
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
func (it *TwitterUserUserUnloginTrackerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TwitterUserUserUnloginTrackerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TwitterUserUserUnloginTracker represents a UserUnloginTracker event raised by the TwitterUser contract.
type TwitterUserUserUnloginTracker struct {
	UserAddress common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterUserUnloginTracker is a free log retrieval operation binding the contract event 0x731ee42bed785b2d551eff7ed95419d92fdc48f6f18af20a463a4f2e860cf90d.
//
// Solidity: event userUnloginTracker(address indexed userAddress)
func (_TwitterUser *TwitterUserFilterer) FilterUserUnloginTracker(opts *bind.FilterOpts, userAddress []common.Address) (*TwitterUserUserUnloginTrackerIterator, error) {

	var userAddressRule []interface{}
	for _, userAddressItem := range userAddress {
		userAddressRule = append(userAddressRule, userAddressItem)
	}

	logs, sub, err := _TwitterUser.contract.FilterLogs(opts, "userUnloginTracker", userAddressRule)
	if err != nil {
		return nil, err
	}
	return &TwitterUserUserUnloginTrackerIterator{contract: _TwitterUser.contract, event: "userUnloginTracker", logs: logs, sub: sub}, nil
}

// WatchUserUnloginTracker is a free log subscription operation binding the contract event 0x731ee42bed785b2d551eff7ed95419d92fdc48f6f18af20a463a4f2e860cf90d.
//
// Solidity: event userUnloginTracker(address indexed userAddress)
func (_TwitterUser *TwitterUserFilterer) WatchUserUnloginTracker(opts *bind.WatchOpts, sink chan<- *TwitterUserUserUnloginTracker, userAddress []common.Address) (event.Subscription, error) {

	var userAddressRule []interface{}
	for _, userAddressItem := range userAddress {
		userAddressRule = append(userAddressRule, userAddressItem)
	}

	logs, sub, err := _TwitterUser.contract.WatchLogs(opts, "userUnloginTracker", userAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TwitterUserUserUnloginTracker)
				if err := _TwitterUser.contract.UnpackLog(event, "userUnloginTracker", log); err != nil {
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

// ParseUserUnloginTracker is a log parse operation binding the contract event 0x731ee42bed785b2d551eff7ed95419d92fdc48f6f18af20a463a4f2e860cf90d.
//
// Solidity: event userUnloginTracker(address indexed userAddress)
func (_TwitterUser *TwitterUserFilterer) ParseUserUnloginTracker(log types.Log) (*TwitterUserUserUnloginTracker, error) {
	event := new(TwitterUserUserUnloginTracker)
	if err := _TwitterUser.contract.UnpackLog(event, "userUnloginTracker", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
