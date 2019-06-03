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

// BokerDappDataABI is the input ABI used to generate the binding from.
const BokerDappDataABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"cName\",\"type\":\"string\"}],\"name\":\"contractAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"key\",\"type\":\"string\"}],\"name\":\"globalConfigString\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"dappCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"dapps\",\"outputs\":[{\"name\":\"dappAddr\",\"type\":\"address\"},{\"name\":\"dappType\",\"type\":\"uint256\"},{\"name\":\"userCount\",\"type\":\"uint256\"},{\"name\":\"monthlySales\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"createTime\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"bokerManager\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"dappAddr\",\"type\":\"address\"},{\"name\":\"dappType\",\"type\":\"uint256\"},{\"name\":\"userCount\",\"type\":\"uint256\"},{\"name\":\"monthlySales\",\"type\":\"uint256\"}],\"name\":\"dappSet\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"dappGetAdresses\",\"outputs\":[{\"name\":\"addrDapps\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addrManager\",\"type\":\"address\"}],\"name\":\"setManager\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"key\",\"type\":\"string\"}],\"name\":\"globalConfigInt\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"dappAddresses\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"addrManager\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// BokerDappDataBin is the compiled bytecode used for deploying new contracts.
const BokerDappDataBin = `60806040526040805190810160405280600581526020017f312e302e3100000000000000000000000000000000000000000000000000000081525060019080519060200190620000519291906200024b565b50426002553480156200006357600080fd5b5060405160208062001a1f8339810180604052810190808051906020019092919050505080336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550620000e281620000ea640100000000026401000000009004565b5050620002fa565b60008190508073ffffffffffffffffffffffffffffffffffffffff1663519c28826040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b1580156200015457600080fd5b505af115801562000169573d6000803e3d6000fd5b505050506040513d60208110156200018057600080fd5b8101908080519060200190929190505050151562000206576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260118152602001807f6e6f7420426f6b65724d616e616765722100000000000000000000000000000081525060200191505060405180910390fd5b80600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106200028e57805160ff1916838001178555620002bf565b82800160010185558215620002bf579182015b82811115620002be578251825591602001919060010190620002a1565b5b509050620002ce9190620002d2565b5090565b620002f791905b80821115620002f3576000816000905550600101620002d9565b5090565b90565b611715806200030a6000396000f3006080604052600436106100d0576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff1680631eb726af146100d5578063378298bc1461017e578063428053441461026057806354fd4d501461028b57806355670f311461031b57806361dcd7ab146103b357806366ebc1c6146103de5780638759bde8146104355780638da5cb5b14610496578063cdf10055146104ed578063d0ebdbe714610559578063d43c80211461059c578063ecead40314610619578063f2fde38b14610686575b600080fd5b3480156100e157600080fd5b5061013c600480360381019080803590602001908201803590602001908080601f01602080910402602001604051908101604052809392919081815260200183838082843782019150505050505091929192905050506106c9565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561018a57600080fd5b506101e5600480360381019080803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290505050610977565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561022557808201518184015260208101905061020a565b50505050905090810190601f1680156102525780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561026c57600080fd5b50610275610b07565b6040518082815260200191505060405180910390f35b34801561029757600080fd5b506102a0610b14565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156102e05780820151818401526020810190506102c5565b50505050905090810190601f16801561030d5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561032757600080fd5b5061035c600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610bb2565b604051808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200184815260200183815260200182815260200194505050505060405180910390f35b3480156103bf57600080fd5b506103c8610c02565b6040518082815260200191505060405180910390f35b3480156103ea57600080fd5b506103f3610c08565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561044157600080fd5b50610494600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291908035906020019092919080359060200190929190505050610c2e565b005b3480156104a257600080fd5b506104ab611145565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156104f957600080fd5b5061050261116a565b6040518080602001828103825283818151815260200191508051906020019060200280838360005b8381101561054557808201518184015260208101905061052a565b505050509050019250505060405180910390f35b34801561056557600080fd5b5061059a600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061124c565b005b3480156105a857600080fd5b50610603600480360381019080803590602001908201803590602001908080601f01602080910402602001604051908101604052809392919081815260200183838082843782019150505050505091929192905050506112b3565b6040518082815260200191505060405180910390f35b34801561062557600080fd5b50610644600480360381019080803590602001909291905050506113ed565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561069257600080fd5b506106c7600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061142b565b005b600080600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663fca1f3c1846040518263ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018080602001828103825283818151815260200191508051906020019080838360005b8381101561077657808201518184015260208101905061075b565b50505050905090810190601f1680156107a35780820380516001836020036101000a031916815260200191505b5092505050600060405180830381600087803b1580156107c257600080fd5b505af11580156107d6573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f8201168201806040525060a081101561080057600080fd5b8101908080519060200190929190805164010000000081111561082257600080fd5b8281019050602081018481111561083857600080fd5b815185600182028301116401000000008211171561085557600080fd5b505092919060200180519060200190929190805164010000000081111561087b57600080fd5b8281019050602081018481111561089157600080fd5b81518560018202830111640100000000821117156108ae57600080fd5b505092919060200180519060200190929190505050505092505050600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415151561096e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600a8152602001807f616464722069732030210000000000000000000000000000000000000000000081525060200191505060405180910390fd5b80915050919050565b6060600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16634a189f35836040518263ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018080602001828103825283818151815260200191508051906020019080838360005b83811015610a23578082015181840152602081019050610a08565b50505050905090810190601f168015610a505780820380516001836020036101000a031916815260200191505b5092505050600060405180830381600087803b158015610a6f57600080fd5b505af1158015610a83573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f820116820180604052506020811015610aad57600080fd5b810190808051640100000000811115610ac557600080fd5b82810190506020810184811115610adb57600080fd5b8151856001820283011164010000000082111715610af857600080fd5b50509291905050509050919050565b6000600580549050905090565b60018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610baa5780601f10610b7f57610100808354040283529160200191610baa565b820191906000526020600020905b815481529060010190602001808311610b8d57829003601f168201915b505050505081565b60046020528060005260406000206000915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060010154908060020154908060030154905084565b60025481565b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663d909fb1b6040805190810160405280600581526020017f61646d696e000000000000000000000000000000000000000000000000000000815250336040518363ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040180806020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001828103825284818151815260200191508051906020019080838360005b83811015610d42578082015181840152602081019050610d27565b50505050905090810190601f168015610d6f5780820380516001836020036101000a031916815260200191505b509350505050602060405180830381600087803b158015610d8f57600080fd5b505af1158015610da3573d6000803e3d6000fd5b505050506040513d6020811015610db957600080fd5b810190808051906020019092919050505080610f6b5750600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663d909fb1b6040805190810160405280600881526020017f636f6e7472616374000000000000000000000000000000000000000000000000815250336040518363ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040180806020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001828103825284818151815260200191508051906020019080838360005b83811015610ee2578082015181840152602081019050610ec7565b50505050905090810190601f168015610f0f5780820380516001836020036101000a031916815260200191505b509350505050602060405180830381600087803b158015610f2f57600080fd5b505af1158015610f43573d6000803e3d6000fd5b505050506040513d6020811015610f5957600080fd5b81019080805190602001909291905050505b1515610fdf576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600f8152602001807f6e6f7420617574686f72697a656421000000000000000000000000000000000081525060200191505060405180910390fd5b600460008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000209050600073ffffffffffffffffffffffffffffffffffffffff168160000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16141561112357848160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060058590806001815401808255809150509060018203906000526020600020016000909192909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550505b8381600101819055508281600201819055508181600301819055505050505050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60606000806005805490509150816040519080825280602002602001820160405280156111a65781602001602082028038833980820191505090505b509250600090505b81811015611247576005818154811015156111c557fe5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1683828151811015156111fe57fe5b9060200190602002019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff168152505080806001019150506111ae565b505090565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156112a757600080fd5b6112b081611492565b50565b6000600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16633230b078836040518263ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018080602001828103825283818151815260200191508051906020019080838360005b8381101561135f578082015181840152602081019050611344565b50505050905090810190601f16801561138c5780820380516001836020036101000a031916815260200191505b5092505050602060405180830381600087803b1580156113ab57600080fd5b505af11580156113bf573d6000803e3d6000fd5b505050506040513d60208110156113d557600080fd5b81019080805190602001909291905050509050919050565b6005818154811015156113fc57fe5b906000526020600020016000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561148657600080fd5b61148f816115ef565b50565b60008190508073ffffffffffffffffffffffffffffffffffffffff1663519c28826040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b1580156114fb57600080fd5b505af115801561150f573d6000803e3d6000fd5b505050506040513d602081101561152557600080fd5b810190808051906020019092919050505015156115aa576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260118152602001807f6e6f7420426f6b65724d616e616765722100000000000000000000000000000081525060200191505060405180910390fd5b80600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415151561162b57600080fd5b8073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550505600a165627a7a723058208cc483b26cf4f1de1d1778ec22055122381aae83e663289041b9783d4c9faf730029`

// DeployBokerDappData deploys a new Ethereum contract, binding an instance of BokerDappData to it.
func DeployBokerDappData(auth *bind.TransactOpts, backend bind.ContractBackend, addrManager common.Address) (common.Address, *types.Transaction, *BokerDappData, error) {
	parsed, err := abi.JSON(strings.NewReader(BokerDappDataABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BokerDappDataBin), backend, addrManager)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BokerDappData{BokerDappDataCaller: BokerDappDataCaller{contract: contract}, BokerDappDataTransactor: BokerDappDataTransactor{contract: contract}}, nil
}

// BokerDappData is an auto generated Go binding around an Ethereum contract.
type BokerDappData struct {
	BokerDappDataCaller     // Read-only binding to the contract
	BokerDappDataTransactor // Write-only binding to the contract
}

// BokerDappDataCaller is an auto generated read-only Go binding around an Ethereum contract.
type BokerDappDataCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BokerDappDataTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BokerDappDataTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BokerDappDataSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BokerDappDataSession struct {
	Contract     *BokerDappData    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BokerDappDataCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BokerDappDataCallerSession struct {
	Contract *BokerDappDataCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// BokerDappDataTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BokerDappDataTransactorSession struct {
	Contract     *BokerDappDataTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// BokerDappDataRaw is an auto generated low-level Go binding around an Ethereum contract.
type BokerDappDataRaw struct {
	Contract *BokerDappData // Generic contract binding to access the raw methods on
}

// BokerDappDataCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BokerDappDataCallerRaw struct {
	Contract *BokerDappDataCaller // Generic read-only contract binding to access the raw methods on
}

// BokerDappDataTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BokerDappDataTransactorRaw struct {
	Contract *BokerDappDataTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBokerDappData creates a new instance of BokerDappData, bound to a specific deployed contract.
func NewBokerDappData(address common.Address, backend bind.ContractBackend) (*BokerDappData, error) {
	contract, err := bindBokerDappData(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BokerDappData{BokerDappDataCaller: BokerDappDataCaller{contract: contract}, BokerDappDataTransactor: BokerDappDataTransactor{contract: contract}}, nil
}

// NewBokerDappDataCaller creates a new read-only instance of BokerDappData, bound to a specific deployed contract.
func NewBokerDappDataCaller(address common.Address, caller bind.ContractCaller) (*BokerDappDataCaller, error) {
	contract, err := bindBokerDappData(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &BokerDappDataCaller{contract: contract}, nil
}

// NewBokerDappDataTransactor creates a new write-only instance of BokerDappData, bound to a specific deployed contract.
func NewBokerDappDataTransactor(address common.Address, transactor bind.ContractTransactor) (*BokerDappDataTransactor, error) {
	contract, err := bindBokerDappData(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &BokerDappDataTransactor{contract: contract}, nil
}

// bindBokerDappData binds a generic wrapper to an already deployed contract.
func bindBokerDappData(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BokerDappDataABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BokerDappData *BokerDappDataRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BokerDappData.Contract.BokerDappDataCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BokerDappData *BokerDappDataRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BokerDappData.Contract.BokerDappDataTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BokerDappData *BokerDappDataRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BokerDappData.Contract.BokerDappDataTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BokerDappData *BokerDappDataCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BokerDappData.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BokerDappData *BokerDappDataTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BokerDappData.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BokerDappData *BokerDappDataTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BokerDappData.Contract.contract.Transact(opts, method, params...)
}

// BokerManager is a free data retrieval call binding the contract method 0x66ebc1c6.
//
// Solidity: function bokerManager() constant returns(address)
func (_BokerDappData *BokerDappDataCaller) BokerManager(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BokerDappData.contract.Call(opts, out, "bokerManager")
	return *ret0, err
}

// BokerManager is a free data retrieval call binding the contract method 0x66ebc1c6.
//
// Solidity: function bokerManager() constant returns(address)
func (_BokerDappData *BokerDappDataSession) BokerManager() (common.Address, error) {
	return _BokerDappData.Contract.BokerManager(&_BokerDappData.CallOpts)
}

// BokerManager is a free data retrieval call binding the contract method 0x66ebc1c6.
//
// Solidity: function bokerManager() constant returns(address)
func (_BokerDappData *BokerDappDataCallerSession) BokerManager() (common.Address, error) {
	return _BokerDappData.Contract.BokerManager(&_BokerDappData.CallOpts)
}

// ContractAddress is a free data retrieval call binding the contract method 0x1eb726af.
//
// Solidity: function contractAddress(cName string) constant returns(address)
func (_BokerDappData *BokerDappDataCaller) ContractAddress(opts *bind.CallOpts, cName string) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BokerDappData.contract.Call(opts, out, "contractAddress", cName)
	return *ret0, err
}

// ContractAddress is a free data retrieval call binding the contract method 0x1eb726af.
//
// Solidity: function contractAddress(cName string) constant returns(address)
func (_BokerDappData *BokerDappDataSession) ContractAddress(cName string) (common.Address, error) {
	return _BokerDappData.Contract.ContractAddress(&_BokerDappData.CallOpts, cName)
}

// ContractAddress is a free data retrieval call binding the contract method 0x1eb726af.
//
// Solidity: function contractAddress(cName string) constant returns(address)
func (_BokerDappData *BokerDappDataCallerSession) ContractAddress(cName string) (common.Address, error) {
	return _BokerDappData.Contract.ContractAddress(&_BokerDappData.CallOpts, cName)
}

// CreateTime is a free data retrieval call binding the contract method 0x61dcd7ab.
//
// Solidity: function createTime() constant returns(uint256)
func (_BokerDappData *BokerDappDataCaller) CreateTime(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BokerDappData.contract.Call(opts, out, "createTime")
	return *ret0, err
}

// CreateTime is a free data retrieval call binding the contract method 0x61dcd7ab.
//
// Solidity: function createTime() constant returns(uint256)
func (_BokerDappData *BokerDappDataSession) CreateTime() (*big.Int, error) {
	return _BokerDappData.Contract.CreateTime(&_BokerDappData.CallOpts)
}

// CreateTime is a free data retrieval call binding the contract method 0x61dcd7ab.
//
// Solidity: function createTime() constant returns(uint256)
func (_BokerDappData *BokerDappDataCallerSession) CreateTime() (*big.Int, error) {
	return _BokerDappData.Contract.CreateTime(&_BokerDappData.CallOpts)
}

// DappAddresses is a free data retrieval call binding the contract method 0xecead403.
//
// Solidity: function dappAddresses( uint256) constant returns(address)
func (_BokerDappData *BokerDappDataCaller) DappAddresses(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BokerDappData.contract.Call(opts, out, "dappAddresses", arg0)
	return *ret0, err
}

// DappAddresses is a free data retrieval call binding the contract method 0xecead403.
//
// Solidity: function dappAddresses( uint256) constant returns(address)
func (_BokerDappData *BokerDappDataSession) DappAddresses(arg0 *big.Int) (common.Address, error) {
	return _BokerDappData.Contract.DappAddresses(&_BokerDappData.CallOpts, arg0)
}

// DappAddresses is a free data retrieval call binding the contract method 0xecead403.
//
// Solidity: function dappAddresses( uint256) constant returns(address)
func (_BokerDappData *BokerDappDataCallerSession) DappAddresses(arg0 *big.Int) (common.Address, error) {
	return _BokerDappData.Contract.DappAddresses(&_BokerDappData.CallOpts, arg0)
}

// DappCount is a free data retrieval call binding the contract method 0x42805344.
//
// Solidity: function dappCount() constant returns(uint256)
func (_BokerDappData *BokerDappDataCaller) DappCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BokerDappData.contract.Call(opts, out, "dappCount")
	return *ret0, err
}

// DappCount is a free data retrieval call binding the contract method 0x42805344.
//
// Solidity: function dappCount() constant returns(uint256)
func (_BokerDappData *BokerDappDataSession) DappCount() (*big.Int, error) {
	return _BokerDappData.Contract.DappCount(&_BokerDappData.CallOpts)
}

// DappCount is a free data retrieval call binding the contract method 0x42805344.
//
// Solidity: function dappCount() constant returns(uint256)
func (_BokerDappData *BokerDappDataCallerSession) DappCount() (*big.Int, error) {
	return _BokerDappData.Contract.DappCount(&_BokerDappData.CallOpts)
}

// DappGetAdresses is a free data retrieval call binding the contract method 0xcdf10055.
//
// Solidity: function dappGetAdresses() constant returns(addrDapps address[])
func (_BokerDappData *BokerDappDataCaller) DappGetAdresses(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _BokerDappData.contract.Call(opts, out, "dappGetAdresses")
	return *ret0, err
}

// DappGetAdresses is a free data retrieval call binding the contract method 0xcdf10055.
//
// Solidity: function dappGetAdresses() constant returns(addrDapps address[])
func (_BokerDappData *BokerDappDataSession) DappGetAdresses() ([]common.Address, error) {
	return _BokerDappData.Contract.DappGetAdresses(&_BokerDappData.CallOpts)
}

// DappGetAdresses is a free data retrieval call binding the contract method 0xcdf10055.
//
// Solidity: function dappGetAdresses() constant returns(addrDapps address[])
func (_BokerDappData *BokerDappDataCallerSession) DappGetAdresses() ([]common.Address, error) {
	return _BokerDappData.Contract.DappGetAdresses(&_BokerDappData.CallOpts)
}

// Dapps is a free data retrieval call binding the contract method 0x55670f31.
//
// Solidity: function dapps( address) constant returns(dappAddr address, dappType uint256, userCount uint256, monthlySales uint256)
func (_BokerDappData *BokerDappDataCaller) Dapps(opts *bind.CallOpts, arg0 common.Address) (struct {
	DappAddr     common.Address
	DappType     *big.Int
	UserCount    *big.Int
	MonthlySales *big.Int
}, error) {
	ret := new(struct {
		DappAddr     common.Address
		DappType     *big.Int
		UserCount    *big.Int
		MonthlySales *big.Int
	})
	out := ret
	err := _BokerDappData.contract.Call(opts, out, "dapps", arg0)
	return *ret, err
}

// Dapps is a free data retrieval call binding the contract method 0x55670f31.
//
// Solidity: function dapps( address) constant returns(dappAddr address, dappType uint256, userCount uint256, monthlySales uint256)
func (_BokerDappData *BokerDappDataSession) Dapps(arg0 common.Address) (struct {
	DappAddr     common.Address
	DappType     *big.Int
	UserCount    *big.Int
	MonthlySales *big.Int
}, error) {
	return _BokerDappData.Contract.Dapps(&_BokerDappData.CallOpts, arg0)
}

// Dapps is a free data retrieval call binding the contract method 0x55670f31.
//
// Solidity: function dapps( address) constant returns(dappAddr address, dappType uint256, userCount uint256, monthlySales uint256)
func (_BokerDappData *BokerDappDataCallerSession) Dapps(arg0 common.Address) (struct {
	DappAddr     common.Address
	DappType     *big.Int
	UserCount    *big.Int
	MonthlySales *big.Int
}, error) {
	return _BokerDappData.Contract.Dapps(&_BokerDappData.CallOpts, arg0)
}

// GlobalConfigInt is a free data retrieval call binding the contract method 0xd43c8021.
//
// Solidity: function globalConfigInt(key string) constant returns(uint256)
func (_BokerDappData *BokerDappDataCaller) GlobalConfigInt(opts *bind.CallOpts, key string) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BokerDappData.contract.Call(opts, out, "globalConfigInt", key)
	return *ret0, err
}

// GlobalConfigInt is a free data retrieval call binding the contract method 0xd43c8021.
//
// Solidity: function globalConfigInt(key string) constant returns(uint256)
func (_BokerDappData *BokerDappDataSession) GlobalConfigInt(key string) (*big.Int, error) {
	return _BokerDappData.Contract.GlobalConfigInt(&_BokerDappData.CallOpts, key)
}

// GlobalConfigInt is a free data retrieval call binding the contract method 0xd43c8021.
//
// Solidity: function globalConfigInt(key string) constant returns(uint256)
func (_BokerDappData *BokerDappDataCallerSession) GlobalConfigInt(key string) (*big.Int, error) {
	return _BokerDappData.Contract.GlobalConfigInt(&_BokerDappData.CallOpts, key)
}

// GlobalConfigString is a free data retrieval call binding the contract method 0x378298bc.
//
// Solidity: function globalConfigString(key string) constant returns(string)
func (_BokerDappData *BokerDappDataCaller) GlobalConfigString(opts *bind.CallOpts, key string) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _BokerDappData.contract.Call(opts, out, "globalConfigString", key)
	return *ret0, err
}

// GlobalConfigString is a free data retrieval call binding the contract method 0x378298bc.
//
// Solidity: function globalConfigString(key string) constant returns(string)
func (_BokerDappData *BokerDappDataSession) GlobalConfigString(key string) (string, error) {
	return _BokerDappData.Contract.GlobalConfigString(&_BokerDappData.CallOpts, key)
}

// GlobalConfigString is a free data retrieval call binding the contract method 0x378298bc.
//
// Solidity: function globalConfigString(key string) constant returns(string)
func (_BokerDappData *BokerDappDataCallerSession) GlobalConfigString(key string) (string, error) {
	return _BokerDappData.Contract.GlobalConfigString(&_BokerDappData.CallOpts, key)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_BokerDappData *BokerDappDataCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BokerDappData.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_BokerDappData *BokerDappDataSession) Owner() (common.Address, error) {
	return _BokerDappData.Contract.Owner(&_BokerDappData.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_BokerDappData *BokerDappDataCallerSession) Owner() (common.Address, error) {
	return _BokerDappData.Contract.Owner(&_BokerDappData.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() constant returns(string)
func (_BokerDappData *BokerDappDataCaller) Version(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _BokerDappData.contract.Call(opts, out, "version")
	return *ret0, err
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() constant returns(string)
func (_BokerDappData *BokerDappDataSession) Version() (string, error) {
	return _BokerDappData.Contract.Version(&_BokerDappData.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() constant returns(string)
func (_BokerDappData *BokerDappDataCallerSession) Version() (string, error) {
	return _BokerDappData.Contract.Version(&_BokerDappData.CallOpts)
}

// DappSet is a paid mutator transaction binding the contract method 0x8759bde8.
//
// Solidity: function dappSet(dappAddr address, dappType uint256, userCount uint256, monthlySales uint256) returns()
func (_BokerDappData *BokerDappDataTransactor) DappSet(opts *bind.TransactOpts, dappAddr common.Address, dappType *big.Int, userCount *big.Int, monthlySales *big.Int) (*types.Transaction, error) {
	return _BokerDappData.contract.Transact(opts, "dappSet", dappAddr, dappType, userCount, monthlySales)
}

// DappSet is a paid mutator transaction binding the contract method 0x8759bde8.
//
// Solidity: function dappSet(dappAddr address, dappType uint256, userCount uint256, monthlySales uint256) returns()
func (_BokerDappData *BokerDappDataSession) DappSet(dappAddr common.Address, dappType *big.Int, userCount *big.Int, monthlySales *big.Int) (*types.Transaction, error) {
	return _BokerDappData.Contract.DappSet(&_BokerDappData.TransactOpts, dappAddr, dappType, userCount, monthlySales)
}

// DappSet is a paid mutator transaction binding the contract method 0x8759bde8.
//
// Solidity: function dappSet(dappAddr address, dappType uint256, userCount uint256, monthlySales uint256) returns()
func (_BokerDappData *BokerDappDataTransactorSession) DappSet(dappAddr common.Address, dappType *big.Int, userCount *big.Int, monthlySales *big.Int) (*types.Transaction, error) {
	return _BokerDappData.Contract.DappSet(&_BokerDappData.TransactOpts, dappAddr, dappType, userCount, monthlySales)
}

// SetManager is a paid mutator transaction binding the contract method 0xd0ebdbe7.
//
// Solidity: function setManager(addrManager address) returns()
func (_BokerDappData *BokerDappDataTransactor) SetManager(opts *bind.TransactOpts, addrManager common.Address) (*types.Transaction, error) {
	return _BokerDappData.contract.Transact(opts, "setManager", addrManager)
}

// SetManager is a paid mutator transaction binding the contract method 0xd0ebdbe7.
//
// Solidity: function setManager(addrManager address) returns()
func (_BokerDappData *BokerDappDataSession) SetManager(addrManager common.Address) (*types.Transaction, error) {
	return _BokerDappData.Contract.SetManager(&_BokerDappData.TransactOpts, addrManager)
}

// SetManager is a paid mutator transaction binding the contract method 0xd0ebdbe7.
//
// Solidity: function setManager(addrManager address) returns()
func (_BokerDappData *BokerDappDataTransactorSession) SetManager(addrManager common.Address) (*types.Transaction, error) {
	return _BokerDappData.Contract.SetManager(&_BokerDappData.TransactOpts, addrManager)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_BokerDappData *BokerDappDataTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _BokerDappData.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_BokerDappData *BokerDappDataSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _BokerDappData.Contract.TransferOwnership(&_BokerDappData.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_BokerDappData *BokerDappDataTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _BokerDappData.Contract.TransferOwnership(&_BokerDappData.TransactOpts, _newOwner)
}
