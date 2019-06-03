// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package boker_contract

import (
	"math/big"
	"strings"

	"github.com/Tinachain/Tina/chain/accounts/abi"
	"github.com/Tinachain/Tina/chain/accounts/abi/bind"
	"github.com/Tinachain/Tina/chain/common"
	"github.com/Tinachain/Tina/chain/core/types"
)

// BokerDappABI is the input ABI used to generate the binding from.
const BokerDappABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"cName\",\"type\":\"string\"}],\"name\":\"contractAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"key\",\"type\":\"string\"}],\"name\":\"globalConfigString\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"createTime\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"bokerManager\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"dappAddr\",\"type\":\"address\"},{\"name\":\"dappType\",\"type\":\"uint256\"},{\"name\":\"userCount\",\"type\":\"uint256\"},{\"name\":\"monthlySales\",\"type\":\"uint256\"}],\"name\":\"dappSet\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"dappGetAdresses\",\"outputs\":[{\"name\":\"addrDapps\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addrManager\",\"type\":\"address\"}],\"name\":\"setManager\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"key\",\"type\":\"string\"}],\"name\":\"globalConfigInt\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"addrManager\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// BokerDappBin is the compiled bytecode used for deploying new contracts.
const BokerDappBin = `60806040526040805190810160405280600581526020017f312e302e3100000000000000000000000000000000000000000000000000000081525060019080519060200190620000519291906200024b565b50426002553480156200006357600080fd5b50604051602080620018308339810180604052810190808051906020019092919050505080336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550620000e281620000ea640100000000026401000000009004565b5050620002fa565b60008190508073ffffffffffffffffffffffffffffffffffffffff1663519c28826040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b1580156200015457600080fd5b505af115801562000169573d6000803e3d6000fd5b505050506040513d60208110156200018057600080fd5b8101908080519060200190929190505050151562000206576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260118152602001807f6e6f7420426f6b65724d616e616765722100000000000000000000000000000081525060200191505060405180910390fd5b80600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106200028e57805160ff1916838001178555620002bf565b82800160010185558215620002bf579182015b82811115620002be578251825591602001919060010190620002a1565b5b509050620002ce9190620002d2565b5090565b620002f791905b80821115620002f3576000816000905550600101620002d9565b5090565b90565b611526806200030a6000396000f3006080604052600436106100af576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff1680631eb726af146100b4578063378298bc1461015d57806354fd4d501461023f57806361dcd7ab146102cf57806366ebc1c6146102fa5780638759bde8146103515780638da5cb5b146103b2578063cdf1005514610409578063d0ebdbe714610475578063d43c8021146104b8578063f2fde38b14610535575b600080fd5b3480156100c057600080fd5b5061011b600480360381019080803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290505050610578565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561016957600080fd5b506101c4600480360381019080803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290505050610826565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156102045780820151818401526020810190506101e9565b50505050905090810190601f1680156102315780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561024b57600080fd5b506102546109b6565b6040518080602001828103825283818151815260200191508051906020019080838360005b83811015610294578082015181840152602081019050610279565b50505050905090810190601f1680156102c15780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b3480156102db57600080fd5b506102e4610a54565b6040518082815260200191505060405180910390f35b34801561030657600080fd5b5061030f610a5a565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561035d57600080fd5b506103b0600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291908035906020019092919080359060200190929190505050610a80565b005b3480156103be57600080fd5b506103c7610f3d565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561041557600080fd5b5061041e610f62565b6040518080602001828103825283818151815260200191508051906020019060200280838360005b83811015610461578082015181840152602081019050610446565b505050509050019250505060405180910390f35b34801561048157600080fd5b506104b6600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061109b565b005b3480156104c457600080fd5b5061051f600480360381019080803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290505050611102565b6040518082815260200191505060405180910390f35b34801561054157600080fd5b50610576600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061123c565b005b600080600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663fca1f3c1846040518263ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018080602001828103825283818151815260200191508051906020019080838360005b8381101561062557808201518184015260208101905061060a565b50505050905090810190601f1680156106525780820380516001836020036101000a031916815260200191505b5092505050600060405180830381600087803b15801561067157600080fd5b505af1158015610685573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f8201168201806040525060a08110156106af57600080fd5b810190808051906020019092919080516401000000008111156106d157600080fd5b828101905060208101848111156106e757600080fd5b815185600182028301116401000000008211171561070457600080fd5b505092919060200180519060200190929190805164010000000081111561072a57600080fd5b8281019050602081018481111561074057600080fd5b815185600182028301116401000000008211171561075d57600080fd5b505092919060200180519060200190929190505050505092505050600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415151561081d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600a8152602001807f616464722069732030210000000000000000000000000000000000000000000081525060200191505060405180910390fd5b80915050919050565b6060600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16634a189f35836040518263ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018080602001828103825283818151815260200191508051906020019080838360005b838110156108d25780820151818401526020810190506108b7565b50505050905090810190601f1680156108ff5780820380516001836020036101000a031916815260200191505b5092505050600060405180830381600087803b15801561091e57600080fd5b505af1158015610932573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f82011682018060405250602081101561095c57600080fd5b81019080805164010000000081111561097457600080fd5b8281019050602081018481111561098a57600080fd5b81518560018202830111640100000000821117156109a757600080fd5b50509291905050509050919050565b60018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610a4c5780601f10610a2157610100808354040283529160200191610a4c565b820191906000526020600020905b815481529060010190602001808311610a2f57829003601f168201915b505050505081565b60025481565b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663d909fb1b6040805190810160405280600581526020017f61646d696e000000000000000000000000000000000000000000000000000000815250336040518363ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040180806020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001828103825284818151815260200191508051906020019080838360005b83811015610b92578082015181840152602081019050610b77565b50505050905090810190601f168015610bbf5780820380516001836020036101000a031916815260200191505b509350505050602060405180830381600087803b158015610bdf57600080fd5b505af1158015610bf3573d6000803e3d6000fd5b505050506040513d6020811015610c0957600080fd5b810190808051906020019092919050505080610dbb5750600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663d909fb1b6040805190810160405280600881526020017f636f6e7472616374000000000000000000000000000000000000000000000000815250336040518363ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040180806020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001828103825284818151815260200191508051906020019080838360005b83811015610d32578082015181840152602081019050610d17565b50505050905090810190601f168015610d5f5780820380516001836020036101000a031916815260200191505b509350505050602060405180830381600087803b158015610d7f57600080fd5b505af1158015610d93573d6000803e3d6000fd5b505050506040513d6020811015610da957600080fd5b81019080805190602001909291905050505b1515610e2f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600f8152602001807f6e6f7420617574686f72697a656421000000000000000000000000000000000081525060200191505060405180910390fd5b610e6d6040805190810160405280600d81526020017f426f6b6572446170704461746100000000000000000000000000000000000000815250610578565b73ffffffffffffffffffffffffffffffffffffffff16638759bde8858585856040518563ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001848152602001838152602001828152602001945050505050600060405180830381600087803b158015610f1f57600080fd5b505af1158015610f33573d6000803e3d6000fd5b5050505050505050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6060610fa26040805190810160405280600d81526020017f426f6b6572446170704461746100000000000000000000000000000000000000815250610578565b73ffffffffffffffffffffffffffffffffffffffff1663cdf100556040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401600060405180830381600087803b15801561100557600080fd5b505af1158015611019573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f82011682018060405250602081101561104357600080fd5b81019080805164010000000081111561105b57600080fd5b8281019050602081018481111561107157600080fd5b815185602082028301116401000000008211171561108e57600080fd5b5050929190505050905090565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156110f657600080fd5b6110ff816112a3565b50565b6000600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16633230b078836040518263ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018080602001828103825283818151815260200191508051906020019080838360005b838110156111ae578082015181840152602081019050611193565b50505050905090810190601f1680156111db5780820380516001836020036101000a031916815260200191505b5092505050602060405180830381600087803b1580156111fa57600080fd5b505af115801561120e573d6000803e3d6000fd5b505050506040513d602081101561122457600080fd5b81019080805190602001909291905050509050919050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561129757600080fd5b6112a081611400565b50565b60008190508073ffffffffffffffffffffffffffffffffffffffff1663519c28826040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b15801561130c57600080fd5b505af1158015611320573d6000803e3d6000fd5b505050506040513d602081101561133657600080fd5b810190808051906020019092919050505015156113bb576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260118152602001807f6e6f7420426f6b65724d616e616765722100000000000000000000000000000081525060200191505060405180910390fd5b80600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415151561143c57600080fd5b8073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550505600a165627a7a72305820802bc5be2b810017aa59732b2516072d4bf55cab6c8eb44b38cdfa5b404102290029`

// DeployBokerDapp deploys a new Ethereum contract, binding an instance of BokerDapp to it.
func DeployBokerDapp(auth *bind.TransactOpts, backend bind.ContractBackend, addrManager common.Address) (common.Address, *types.Transaction, *BokerDapp, error) {
	parsed, err := abi.JSON(strings.NewReader(BokerDappABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BokerDappBin), backend, addrManager)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BokerDapp{BokerDappCaller: BokerDappCaller{contract: contract}, BokerDappTransactor: BokerDappTransactor{contract: contract}}, nil
}

// BokerDapp is an auto generated Go binding around an Ethereum contract.
type BokerDapp struct {
	BokerDappCaller     // Read-only binding to the contract
	BokerDappTransactor // Write-only binding to the contract
}

// BokerDappCaller is an auto generated read-only Go binding around an Ethereum contract.
type BokerDappCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BokerDappTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BokerDappTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BokerDappSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BokerDappSession struct {
	Contract     *BokerDapp        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BokerDappCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BokerDappCallerSession struct {
	Contract *BokerDappCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// BokerDappTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BokerDappTransactorSession struct {
	Contract     *BokerDappTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// BokerDappRaw is an auto generated low-level Go binding around an Ethereum contract.
type BokerDappRaw struct {
	Contract *BokerDapp // Generic contract binding to access the raw methods on
}

// BokerDappCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BokerDappCallerRaw struct {
	Contract *BokerDappCaller // Generic read-only contract binding to access the raw methods on
}

// BokerDappTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BokerDappTransactorRaw struct {
	Contract *BokerDappTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBokerDapp creates a new instance of BokerDapp, bound to a specific deployed contract.
func NewBokerDapp(address common.Address, backend bind.ContractBackend) (*BokerDapp, error) {
	contract, err := bindBokerDapp(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BokerDapp{BokerDappCaller: BokerDappCaller{contract: contract}, BokerDappTransactor: BokerDappTransactor{contract: contract}}, nil
}

// NewBokerDappCaller creates a new read-only instance of BokerDapp, bound to a specific deployed contract.
func NewBokerDappCaller(address common.Address, caller bind.ContractCaller) (*BokerDappCaller, error) {
	contract, err := bindBokerDapp(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &BokerDappCaller{contract: contract}, nil
}

// NewBokerDappTransactor creates a new write-only instance of BokerDapp, bound to a specific deployed contract.
func NewBokerDappTransactor(address common.Address, transactor bind.ContractTransactor) (*BokerDappTransactor, error) {
	contract, err := bindBokerDapp(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &BokerDappTransactor{contract: contract}, nil
}

// bindBokerDapp binds a generic wrapper to an already deployed contract.
func bindBokerDapp(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BokerDappABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BokerDapp *BokerDappRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BokerDapp.Contract.BokerDappCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BokerDapp *BokerDappRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BokerDapp.Contract.BokerDappTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BokerDapp *BokerDappRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BokerDapp.Contract.BokerDappTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BokerDapp *BokerDappCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BokerDapp.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BokerDapp *BokerDappTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BokerDapp.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BokerDapp *BokerDappTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BokerDapp.Contract.contract.Transact(opts, method, params...)
}

// BokerManager is a free data retrieval call binding the contract method 0x66ebc1c6.
//
// Solidity: function bokerManager() constant returns(address)
func (_BokerDapp *BokerDappCaller) BokerManager(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BokerDapp.contract.Call(opts, out, "bokerManager")
	return *ret0, err
}

// BokerManager is a free data retrieval call binding the contract method 0x66ebc1c6.
//
// Solidity: function bokerManager() constant returns(address)
func (_BokerDapp *BokerDappSession) BokerManager() (common.Address, error) {
	return _BokerDapp.Contract.BokerManager(&_BokerDapp.CallOpts)
}

// BokerManager is a free data retrieval call binding the contract method 0x66ebc1c6.
//
// Solidity: function bokerManager() constant returns(address)
func (_BokerDapp *BokerDappCallerSession) BokerManager() (common.Address, error) {
	return _BokerDapp.Contract.BokerManager(&_BokerDapp.CallOpts)
}

// ContractAddress is a free data retrieval call binding the contract method 0x1eb726af.
//
// Solidity: function contractAddress(cName string) constant returns(address)
func (_BokerDapp *BokerDappCaller) ContractAddress(opts *bind.CallOpts, cName string) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BokerDapp.contract.Call(opts, out, "contractAddress", cName)
	return *ret0, err
}

// ContractAddress is a free data retrieval call binding the contract method 0x1eb726af.
//
// Solidity: function contractAddress(cName string) constant returns(address)
func (_BokerDapp *BokerDappSession) ContractAddress(cName string) (common.Address, error) {
	return _BokerDapp.Contract.ContractAddress(&_BokerDapp.CallOpts, cName)
}

// ContractAddress is a free data retrieval call binding the contract method 0x1eb726af.
//
// Solidity: function contractAddress(cName string) constant returns(address)
func (_BokerDapp *BokerDappCallerSession) ContractAddress(cName string) (common.Address, error) {
	return _BokerDapp.Contract.ContractAddress(&_BokerDapp.CallOpts, cName)
}

// CreateTime is a free data retrieval call binding the contract method 0x61dcd7ab.
//
// Solidity: function createTime() constant returns(uint256)
func (_BokerDapp *BokerDappCaller) CreateTime(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BokerDapp.contract.Call(opts, out, "createTime")
	return *ret0, err
}

// CreateTime is a free data retrieval call binding the contract method 0x61dcd7ab.
//
// Solidity: function createTime() constant returns(uint256)
func (_BokerDapp *BokerDappSession) CreateTime() (*big.Int, error) {
	return _BokerDapp.Contract.CreateTime(&_BokerDapp.CallOpts)
}

// CreateTime is a free data retrieval call binding the contract method 0x61dcd7ab.
//
// Solidity: function createTime() constant returns(uint256)
func (_BokerDapp *BokerDappCallerSession) CreateTime() (*big.Int, error) {
	return _BokerDapp.Contract.CreateTime(&_BokerDapp.CallOpts)
}

// DappGetAdresses is a free data retrieval call binding the contract method 0xcdf10055.
//
// Solidity: function dappGetAdresses() constant returns(addrDapps address[])
func (_BokerDapp *BokerDappCaller) DappGetAdresses(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _BokerDapp.contract.Call(opts, out, "dappGetAdresses")
	return *ret0, err
}

// DappGetAdresses is a free data retrieval call binding the contract method 0xcdf10055.
//
// Solidity: function dappGetAdresses() constant returns(addrDapps address[])
func (_BokerDapp *BokerDappSession) DappGetAdresses() ([]common.Address, error) {
	return _BokerDapp.Contract.DappGetAdresses(&_BokerDapp.CallOpts)
}

// DappGetAdresses is a free data retrieval call binding the contract method 0xcdf10055.
//
// Solidity: function dappGetAdresses() constant returns(addrDapps address[])
func (_BokerDapp *BokerDappCallerSession) DappGetAdresses() ([]common.Address, error) {
	return _BokerDapp.Contract.DappGetAdresses(&_BokerDapp.CallOpts)
}

// GlobalConfigInt is a free data retrieval call binding the contract method 0xd43c8021.
//
// Solidity: function globalConfigInt(key string) constant returns(uint256)
func (_BokerDapp *BokerDappCaller) GlobalConfigInt(opts *bind.CallOpts, key string) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BokerDapp.contract.Call(opts, out, "globalConfigInt", key)
	return *ret0, err
}

// GlobalConfigInt is a free data retrieval call binding the contract method 0xd43c8021.
//
// Solidity: function globalConfigInt(key string) constant returns(uint256)
func (_BokerDapp *BokerDappSession) GlobalConfigInt(key string) (*big.Int, error) {
	return _BokerDapp.Contract.GlobalConfigInt(&_BokerDapp.CallOpts, key)
}

// GlobalConfigInt is a free data retrieval call binding the contract method 0xd43c8021.
//
// Solidity: function globalConfigInt(key string) constant returns(uint256)
func (_BokerDapp *BokerDappCallerSession) GlobalConfigInt(key string) (*big.Int, error) {
	return _BokerDapp.Contract.GlobalConfigInt(&_BokerDapp.CallOpts, key)
}

// GlobalConfigString is a free data retrieval call binding the contract method 0x378298bc.
//
// Solidity: function globalConfigString(key string) constant returns(string)
func (_BokerDapp *BokerDappCaller) GlobalConfigString(opts *bind.CallOpts, key string) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _BokerDapp.contract.Call(opts, out, "globalConfigString", key)
	return *ret0, err
}

// GlobalConfigString is a free data retrieval call binding the contract method 0x378298bc.
//
// Solidity: function globalConfigString(key string) constant returns(string)
func (_BokerDapp *BokerDappSession) GlobalConfigString(key string) (string, error) {
	return _BokerDapp.Contract.GlobalConfigString(&_BokerDapp.CallOpts, key)
}

// GlobalConfigString is a free data retrieval call binding the contract method 0x378298bc.
//
// Solidity: function globalConfigString(key string) constant returns(string)
func (_BokerDapp *BokerDappCallerSession) GlobalConfigString(key string) (string, error) {
	return _BokerDapp.Contract.GlobalConfigString(&_BokerDapp.CallOpts, key)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_BokerDapp *BokerDappCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BokerDapp.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_BokerDapp *BokerDappSession) Owner() (common.Address, error) {
	return _BokerDapp.Contract.Owner(&_BokerDapp.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_BokerDapp *BokerDappCallerSession) Owner() (common.Address, error) {
	return _BokerDapp.Contract.Owner(&_BokerDapp.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() constant returns(string)
func (_BokerDapp *BokerDappCaller) Version(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _BokerDapp.contract.Call(opts, out, "version")
	return *ret0, err
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() constant returns(string)
func (_BokerDapp *BokerDappSession) Version() (string, error) {
	return _BokerDapp.Contract.Version(&_BokerDapp.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() constant returns(string)
func (_BokerDapp *BokerDappCallerSession) Version() (string, error) {
	return _BokerDapp.Contract.Version(&_BokerDapp.CallOpts)
}

// DappSet is a paid mutator transaction binding the contract method 0x8759bde8.
//
// Solidity: function dappSet(dappAddr address, dappType uint256, userCount uint256, monthlySales uint256) returns()
func (_BokerDapp *BokerDappTransactor) DappSet(opts *bind.TransactOpts, dappAddr common.Address, dappType *big.Int, userCount *big.Int, monthlySales *big.Int) (*types.Transaction, error) {
	return _BokerDapp.contract.Transact(opts, "dappSet", dappAddr, dappType, userCount, monthlySales)
}

// DappSet is a paid mutator transaction binding the contract method 0x8759bde8.
//
// Solidity: function dappSet(dappAddr address, dappType uint256, userCount uint256, monthlySales uint256) returns()
func (_BokerDapp *BokerDappSession) DappSet(dappAddr common.Address, dappType *big.Int, userCount *big.Int, monthlySales *big.Int) (*types.Transaction, error) {
	return _BokerDapp.Contract.DappSet(&_BokerDapp.TransactOpts, dappAddr, dappType, userCount, monthlySales)
}

// DappSet is a paid mutator transaction binding the contract method 0x8759bde8.
//
// Solidity: function dappSet(dappAddr address, dappType uint256, userCount uint256, monthlySales uint256) returns()
func (_BokerDapp *BokerDappTransactorSession) DappSet(dappAddr common.Address, dappType *big.Int, userCount *big.Int, monthlySales *big.Int) (*types.Transaction, error) {
	return _BokerDapp.Contract.DappSet(&_BokerDapp.TransactOpts, dappAddr, dappType, userCount, monthlySales)
}

// SetManager is a paid mutator transaction binding the contract method 0xd0ebdbe7.
//
// Solidity: function setManager(addrManager address) returns()
func (_BokerDapp *BokerDappTransactor) SetManager(opts *bind.TransactOpts, addrManager common.Address) (*types.Transaction, error) {
	return _BokerDapp.contract.Transact(opts, "setManager", addrManager)
}

// SetManager is a paid mutator transaction binding the contract method 0xd0ebdbe7.
//
// Solidity: function setManager(addrManager address) returns()
func (_BokerDapp *BokerDappSession) SetManager(addrManager common.Address) (*types.Transaction, error) {
	return _BokerDapp.Contract.SetManager(&_BokerDapp.TransactOpts, addrManager)
}

// SetManager is a paid mutator transaction binding the contract method 0xd0ebdbe7.
//
// Solidity: function setManager(addrManager address) returns()
func (_BokerDapp *BokerDappTransactorSession) SetManager(addrManager common.Address) (*types.Transaction, error) {
	return _BokerDapp.Contract.SetManager(&_BokerDapp.TransactOpts, addrManager)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_BokerDapp *BokerDappTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _BokerDapp.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_BokerDapp *BokerDappSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _BokerDapp.Contract.TransferOwnership(&_BokerDapp.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_BokerDapp *BokerDappTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _BokerDapp.Contract.TransferOwnership(&_BokerDapp.TransactOpts, _newOwner)
}
