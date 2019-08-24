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

// BokerInterfaceBaseABI is the input ABI used to generate the binding from.
const BokerInterfaceBaseABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"getCandidates\",\"outputs\":[{\"name\":\"addresses\",\"type\":\"address[]\"},{\"name\":\"tickets\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"testCandidateArray\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"testCandidates\",\"outputs\":[{\"name\":\"index\",\"type\":\"uint256\"},{\"name\":\"addr\",\"type\":\"address\"},{\"name\":\"tickets\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"cName\",\"type\":\"string\"}],\"name\":\"contractAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"key\",\"type\":\"string\"}],\"name\":\"globalConfigString\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"nowTimer\",\"type\":\"uint256\"}],\"name\":\"rotateVote\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"createTime\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"bokerManager\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"nodeAddress\",\"type\":\"address\"},{\"name\":\"tickets\",\"type\":\"uint256\"}],\"name\":\"setCandidates\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addrManager\",\"type\":\"address\"}],\"name\":\"setManager\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getBlacks\",\"outputs\":[{\"name\":\"addresses\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"key\",\"type\":\"string\"}],\"name\":\"globalConfigInt\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"setSize\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"addrManager\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// BokerInterfaceBaseBin is the compiled bytecode used for deploying new contracts.
const BokerInterfaceBaseBin = `60806040526040805190810160405280600581526020017f312e302e310000000000000000000000000000000000000000000000000000008152506001908051906020019062000051929190620002bc565b5042600255600160045560006005553480156200006d57600080fd5b5060405160208062001c3b8339810180604052810190808051906020019092919050505080336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550620000ec816200015b640100000000026401000000009004565b506007600090806001815401808255809150509060018203906000526020600020016000909192909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050506200036b565b60008190508073ffffffffffffffffffffffffffffffffffffffff1663519c28826040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b158015620001c557600080fd5b505af1158015620001da573d6000803e3d6000fd5b505050506040513d6020811015620001f157600080fd5b8101908080519060200190929190505050151562000277576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260118152602001807f6e6f7420426f6b65724d616e616765722100000000000000000000000000000081525060200191505060405180910390fd5b80600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10620002ff57805160ff191683800117855562000330565b8280016001018555821562000330579182015b828111156200032f57825182559160200191906001019062000312565b5b5090506200033f919062000343565b5090565b6200036891905b80821115620003645760008160009055506001016200034a565b5090565b90565b6118c0806200037b6000396000f3006080604052600436106100e6576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff16806306a49fce146100eb57806308cae5d31461019f57806316439a661461020c5780631eb726af1461029d578063378298bc1461034657806343bd7d971461042857806354fd4d501461045557806361dcd7ab146104e557806366ebc1c6146105105780638265e61e146105675780638da5cb5b146105cc578063d0ebdbe714610623578063d38ff7a114610666578063d43c8021146106d2578063ee90c99f1461074f578063f2fde38b1461077a575b600080fd5b3480156100f757600080fd5b506101006107bd565b604051808060200180602001838103835285818151815260200191508051906020019060200280838360005b8381101561014757808201518184015260208101905061012c565b50505050905001838103825284818151815260200191508051906020019060200280838360005b8381101561018957808201518184015260208101905061016e565b5050505090500194505050505060405180910390f35b3480156101ab57600080fd5b506101ca60048036038101908080359060200190929190505050610ad9565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561021857600080fd5b5061024d600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610b17565b604051808481526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001828152602001935050505060405180910390f35b3480156102a957600080fd5b50610304600480360381019080803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290505050610b61565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561035257600080fd5b506103ad600480360381019080803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290505050610e0f565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156103ed5780820151818401526020810190506103d2565b50505050905090810190601f16801561041a5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561043457600080fd5b5061045360048036038101908080359060200190929190505050610f9f565b005b34801561046157600080fd5b5061046a611066565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156104aa57808201518184015260208101905061048f565b50505050905090810190601f1680156104d75780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b3480156104f157600080fd5b506104fa611104565b6040518082815260200191505060405180910390f35b34801561051c57600080fd5b5061052561110a565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561057357600080fd5b506105b2600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050611130565b604051808215151515815260200191505060405180910390f35b3480156105d857600080fd5b506105e1611268565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561062f57600080fd5b50610664600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061128d565b005b34801561067257600080fd5b5061067b6112f4565b6040518080602001828103825283818151815260200191508051906020019060200280838360005b838110156106be5780820151818401526020810190506106a3565b505050509050019250505060405180910390f35b3480156106de57600080fd5b50610739600480360381019080803590602001908201803590602001908080601f016020809104026020016040519081016040528093929190818152602001838380828437820191505050505050919291929050505061142d565b6040518082815260200191505060405180910390f35b34801561075b57600080fd5b50610764611567565b6040518082815260200191505060405180910390f35b34801561078657600080fd5b506107bb600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061156d565b005b60608060008060008060045414156109595761080d6040805190810160405280600981526020017f426f6b65724e6f64650000000000000000000000000000000000000000000000815250610b61565b73ffffffffffffffffffffffffffffffffffffffff166306a49fce6040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401600060405180830381600087803b15801561087057600080fd5b505af1158015610884573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f8201168201806040525060408110156108ae57600080fd5b8101908080516401000000008111156108c657600080fd5b828101905060208101848111156108dc57600080fd5b81518560208202830111640100000000821117156108f957600080fd5b5050929190602001805164010000000081111561091557600080fd5b8281019050602081018481111561092b57600080fd5b815185602082028301116401000000008211171561094857600080fd5b505092919050505094509450610ad2565b6007805490509250600183036040519080825280602002602001820160405280156109935781602001602082028038833980820191505090505b509450600183036040519080825280602002602001820160405280156109c85781602001602082028038833980820191505090505b509350600191505b82821015610ad1576007828154811015156109e757fe5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050808560018403815181101515610a2657fe5b9060200190602002019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff1681525050600660008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600201548460018403815181101515610ab657fe5b906020019060200201818152505081806001019250506109d0565b5b5050509091565b600781815481101515610ae857fe5b906000526020600020016000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60066020528060005260406000206000915090508060000154908060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060020154905083565b600080600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663fca1f3c1846040518263ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018080602001828103825283818151815260200191508051906020019080838360005b83811015610c0e578082015181840152602081019050610bf3565b50505050905090810190601f168015610c3b5780820380516001836020036101000a031916815260200191505b5092505050600060405180830381600087803b158015610c5a57600080fd5b505af1158015610c6e573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f8201168201806040525060a0811015610c9857600080fd5b81019080805190602001909291908051640100000000811115610cba57600080fd5b82810190506020810184811115610cd057600080fd5b8151856001820283011164010000000082111715610ced57600080fd5b5050929190602001805190602001909291908051640100000000811115610d1357600080fd5b82810190506020810184811115610d2957600080fd5b8151856001820283011164010000000082111715610d4657600080fd5b505092919060200180519060200190929190505050505092505050600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614151515610e06576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600a8152602001807f616464722069732030210000000000000000000000000000000000000000000081525060200191505060405180910390fd5b80915050919050565b6060600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16634a189f35836040518263ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018080602001828103825283818151815260200191508051906020019080838360005b83811015610ebb578082015181840152602081019050610ea0565b50505050905090810190601f168015610ee85780820380516001836020036101000a031916815260200191505b5092505050600060405180830381600087803b158015610f0757600080fd5b505af1158015610f1b573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f820116820180604052506020811015610f4557600080fd5b810190808051640100000000811115610f5d57600080fd5b82810190506020810184811115610f7357600080fd5b8151856001820283011164010000000082111715610f9057600080fd5b50509291905050509050919050565b610fdd6040805190810160405280600981526020017f426f6b65724e6f64650000000000000000000000000000000000000000000000815250610b61565b73ffffffffffffffffffffffffffffffffffffffff166343bd7d97826040518263ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040180828152602001915050600060405180830381600087803b15801561104b57600080fd5b505af115801561105f573d6000803e3d6000fd5b5050505050565b60018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156110fc5780601f106110d1576101008083540402835291602001916110fc565b820191906000526020600020905b8154815290600101906020018083116110df57829003601f168201915b505050505081565b60025481565b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600080600660008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020905060008160000154111561118a5760009150611261565b838160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600780549050816000018190555082816002018190555060078490806001815401808255809150509060018203906000526020600020016000909192909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050600560008154809291906001019190505550600191505b5092915050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156112e857600080fd5b6112f1816115d4565b50565b60606113346040805190810160405280600981526020017f426f6b65724e6f64650000000000000000000000000000000000000000000000815250610b61565b73ffffffffffffffffffffffffffffffffffffffff1663d38ff7a16040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401600060405180830381600087803b15801561139757600080fd5b505af11580156113ab573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f8201168201806040525060208110156113d557600080fd5b8101908080516401000000008111156113ed57600080fd5b8281019050602081018481111561140357600080fd5b815185602082028301116401000000008211171561142057600080fd5b5050929190505050905090565b6000600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16633230b078836040518263ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018080602001828103825283818151815260200191508051906020019080838360005b838110156114d95780820151818401526020810190506114be565b50505050905090810190601f1680156115065780820380516001836020036101000a031916815260200191505b5092505050602060405180830381600087803b15801561152557600080fd5b505af1158015611539573d6000803e3d6000fd5b505050506040513d602081101561154f57600080fd5b81019080805190602001909291905050509050919050565b60055481565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156115c857600080fd5b6115d181611731565b50565b60008190508073ffffffffffffffffffffffffffffffffffffffff1663519c28826040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b15801561163d57600080fd5b505af1158015611651573d6000803e3d6000fd5b505050506040513d602081101561166757600080fd5b810190808051906020019092919050505015156116ec576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260118152602001807f6e6f7420426f6b65724d616e616765722100000000000000000000000000000081525060200191505060405180910390fd5b80600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16141515156117d6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260168152602001807f6e65774f776e657220616464726573732069732030210000000000000000000081525060200191505060405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550505600a165627a7a723058208d0790718b05775eff5c0c0ee6a656fadd9eb7464d1e5c6e3589759f7e5305fc0029`

// DeployBokerInterfaceBase deploys a new Ethereum contract, binding an instance of BokerInterfaceBase to it.
func DeployBokerInterfaceBase(auth *bind.TransactOpts, backend bind.ContractBackend, addrManager common.Address) (common.Address, *types.Transaction, *BokerInterfaceBase, error) {
	parsed, err := abi.JSON(strings.NewReader(BokerInterfaceBaseABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BokerInterfaceBaseBin), backend, addrManager)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BokerInterfaceBase{BokerInterfaceBaseCaller: BokerInterfaceBaseCaller{contract: contract}, BokerInterfaceBaseTransactor: BokerInterfaceBaseTransactor{contract: contract}}, nil
}

// BokerInterfaceBase is an auto generated Go binding around an Ethereum contract.
type BokerInterfaceBase struct {
	BokerInterfaceBaseCaller     // Read-only binding to the contract
	BokerInterfaceBaseTransactor // Write-only binding to the contract
}

// BokerInterfaceBaseCaller is an auto generated read-only Go binding around an Ethereum contract.
type BokerInterfaceBaseCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BokerInterfaceBaseTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BokerInterfaceBaseTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BokerInterfaceBaseSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BokerInterfaceBaseSession struct {
	Contract     *BokerInterfaceBase // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// BokerInterfaceBaseCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BokerInterfaceBaseCallerSession struct {
	Contract *BokerInterfaceBaseCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// BokerInterfaceBaseTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BokerInterfaceBaseTransactorSession struct {
	Contract     *BokerInterfaceBaseTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// BokerInterfaceBaseRaw is an auto generated low-level Go binding around an Ethereum contract.
type BokerInterfaceBaseRaw struct {
	Contract *BokerInterfaceBase // Generic contract binding to access the raw methods on
}

// BokerInterfaceBaseCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BokerInterfaceBaseCallerRaw struct {
	Contract *BokerInterfaceBaseCaller // Generic read-only contract binding to access the raw methods on
}

// BokerInterfaceBaseTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BokerInterfaceBaseTransactorRaw struct {
	Contract *BokerInterfaceBaseTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBokerInterfaceBase creates a new instance of BokerInterfaceBase, bound to a specific deployed contract.
func NewBokerInterfaceBase(address common.Address, backend bind.ContractBackend) (*BokerInterfaceBase, error) {
	contract, err := bindBokerInterfaceBase(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BokerInterfaceBase{BokerInterfaceBaseCaller: BokerInterfaceBaseCaller{contract: contract}, BokerInterfaceBaseTransactor: BokerInterfaceBaseTransactor{contract: contract}}, nil
}

// NewBokerInterfaceBaseCaller creates a new read-only instance of BokerInterfaceBase, bound to a specific deployed contract.
func NewBokerInterfaceBaseCaller(address common.Address, caller bind.ContractCaller) (*BokerInterfaceBaseCaller, error) {
	contract, err := bindBokerInterfaceBase(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &BokerInterfaceBaseCaller{contract: contract}, nil
}

// NewBokerInterfaceBaseTransactor creates a new write-only instance of BokerInterfaceBase, bound to a specific deployed contract.
func NewBokerInterfaceBaseTransactor(address common.Address, transactor bind.ContractTransactor) (*BokerInterfaceBaseTransactor, error) {
	contract, err := bindBokerInterfaceBase(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &BokerInterfaceBaseTransactor{contract: contract}, nil
}

// bindBokerInterfaceBase binds a generic wrapper to an already deployed contract.
func bindBokerInterfaceBase(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BokerInterfaceBaseABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BokerInterfaceBase *BokerInterfaceBaseRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BokerInterfaceBase.Contract.BokerInterfaceBaseCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BokerInterfaceBase *BokerInterfaceBaseRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BokerInterfaceBase.Contract.BokerInterfaceBaseTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BokerInterfaceBase *BokerInterfaceBaseRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BokerInterfaceBase.Contract.BokerInterfaceBaseTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BokerInterfaceBase *BokerInterfaceBaseCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BokerInterfaceBase.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BokerInterfaceBase *BokerInterfaceBaseTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BokerInterfaceBase.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BokerInterfaceBase *BokerInterfaceBaseTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BokerInterfaceBase.Contract.contract.Transact(opts, method, params...)
}

// BokerManager is a free data retrieval call binding the contract method 0x66ebc1c6.
//
// Solidity: function bokerManager() constant returns(address)
func (_BokerInterfaceBase *BokerInterfaceBaseCaller) BokerManager(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BokerInterfaceBase.contract.Call(opts, out, "bokerManager")
	return *ret0, err
}

// BokerManager is a free data retrieval call binding the contract method 0x66ebc1c6.
//
// Solidity: function bokerManager() constant returns(address)
func (_BokerInterfaceBase *BokerInterfaceBaseSession) BokerManager() (common.Address, error) {
	return _BokerInterfaceBase.Contract.BokerManager(&_BokerInterfaceBase.CallOpts)
}

// BokerManager is a free data retrieval call binding the contract method 0x66ebc1c6.
//
// Solidity: function bokerManager() constant returns(address)
func (_BokerInterfaceBase *BokerInterfaceBaseCallerSession) BokerManager() (common.Address, error) {
	return _BokerInterfaceBase.Contract.BokerManager(&_BokerInterfaceBase.CallOpts)
}

// ContractAddress is a free data retrieval call binding the contract method 0x1eb726af.
//
// Solidity: function contractAddress(cName string) constant returns(address)
func (_BokerInterfaceBase *BokerInterfaceBaseCaller) ContractAddress(opts *bind.CallOpts, cName string) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BokerInterfaceBase.contract.Call(opts, out, "contractAddress", cName)
	return *ret0, err
}

// ContractAddress is a free data retrieval call binding the contract method 0x1eb726af.
//
// Solidity: function contractAddress(cName string) constant returns(address)
func (_BokerInterfaceBase *BokerInterfaceBaseSession) ContractAddress(cName string) (common.Address, error) {
	return _BokerInterfaceBase.Contract.ContractAddress(&_BokerInterfaceBase.CallOpts, cName)
}

// ContractAddress is a free data retrieval call binding the contract method 0x1eb726af.
//
// Solidity: function contractAddress(cName string) constant returns(address)
func (_BokerInterfaceBase *BokerInterfaceBaseCallerSession) ContractAddress(cName string) (common.Address, error) {
	return _BokerInterfaceBase.Contract.ContractAddress(&_BokerInterfaceBase.CallOpts, cName)
}

// CreateTime is a free data retrieval call binding the contract method 0x61dcd7ab.
//
// Solidity: function createTime() constant returns(uint256)
func (_BokerInterfaceBase *BokerInterfaceBaseCaller) CreateTime(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BokerInterfaceBase.contract.Call(opts, out, "createTime")
	return *ret0, err
}

// CreateTime is a free data retrieval call binding the contract method 0x61dcd7ab.
//
// Solidity: function createTime() constant returns(uint256)
func (_BokerInterfaceBase *BokerInterfaceBaseSession) CreateTime() (*big.Int, error) {
	return _BokerInterfaceBase.Contract.CreateTime(&_BokerInterfaceBase.CallOpts)
}

// CreateTime is a free data retrieval call binding the contract method 0x61dcd7ab.
//
// Solidity: function createTime() constant returns(uint256)
func (_BokerInterfaceBase *BokerInterfaceBaseCallerSession) CreateTime() (*big.Int, error) {
	return _BokerInterfaceBase.Contract.CreateTime(&_BokerInterfaceBase.CallOpts)
}

// GetBlacks is a free data retrieval call binding the contract method 0xd38ff7a1.
//
// Solidity: function getBlacks() constant returns(addresses address[])
func (_BokerInterfaceBase *BokerInterfaceBaseCaller) GetBlacks(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _BokerInterfaceBase.contract.Call(opts, out, "getBlacks")
	return *ret0, err
}

// GetBlacks is a free data retrieval call binding the contract method 0xd38ff7a1.
//
// Solidity: function getBlacks() constant returns(addresses address[])
func (_BokerInterfaceBase *BokerInterfaceBaseSession) GetBlacks() ([]common.Address, error) {
	return _BokerInterfaceBase.Contract.GetBlacks(&_BokerInterfaceBase.CallOpts)
}

// GetBlacks is a free data retrieval call binding the contract method 0xd38ff7a1.
//
// Solidity: function getBlacks() constant returns(addresses address[])
func (_BokerInterfaceBase *BokerInterfaceBaseCallerSession) GetBlacks() ([]common.Address, error) {
	return _BokerInterfaceBase.Contract.GetBlacks(&_BokerInterfaceBase.CallOpts)
}

// GetCandidates is a free data retrieval call binding the contract method 0x06a49fce.
//
// Solidity: function getCandidates() constant returns(addresses address[], tickets uint256[])
func (_BokerInterfaceBase *BokerInterfaceBaseCaller) GetCandidates(opts *bind.CallOpts) (struct {
	Addresses []common.Address
	Tickets   []*big.Int
}, error) {
	ret := new(struct {
		Addresses []common.Address
		Tickets   []*big.Int
	})
	out := ret
	err := _BokerInterfaceBase.contract.Call(opts, out, "getCandidates")
	return *ret, err
}

// GetCandidates is a free data retrieval call binding the contract method 0x06a49fce.
//
// Solidity: function getCandidates() constant returns(addresses address[], tickets uint256[])
func (_BokerInterfaceBase *BokerInterfaceBaseSession) GetCandidates() (struct {
	Addresses []common.Address
	Tickets   []*big.Int
}, error) {
	return _BokerInterfaceBase.Contract.GetCandidates(&_BokerInterfaceBase.CallOpts)
}

// GetCandidates is a free data retrieval call binding the contract method 0x06a49fce.
//
// Solidity: function getCandidates() constant returns(addresses address[], tickets uint256[])
func (_BokerInterfaceBase *BokerInterfaceBaseCallerSession) GetCandidates() (struct {
	Addresses []common.Address
	Tickets   []*big.Int
}, error) {
	return _BokerInterfaceBase.Contract.GetCandidates(&_BokerInterfaceBase.CallOpts)
}

// GlobalConfigInt is a free data retrieval call binding the contract method 0xd43c8021.
//
// Solidity: function globalConfigInt(key string) constant returns(uint256)
func (_BokerInterfaceBase *BokerInterfaceBaseCaller) GlobalConfigInt(opts *bind.CallOpts, key string) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BokerInterfaceBase.contract.Call(opts, out, "globalConfigInt", key)
	return *ret0, err
}

// GlobalConfigInt is a free data retrieval call binding the contract method 0xd43c8021.
//
// Solidity: function globalConfigInt(key string) constant returns(uint256)
func (_BokerInterfaceBase *BokerInterfaceBaseSession) GlobalConfigInt(key string) (*big.Int, error) {
	return _BokerInterfaceBase.Contract.GlobalConfigInt(&_BokerInterfaceBase.CallOpts, key)
}

// GlobalConfigInt is a free data retrieval call binding the contract method 0xd43c8021.
//
// Solidity: function globalConfigInt(key string) constant returns(uint256)
func (_BokerInterfaceBase *BokerInterfaceBaseCallerSession) GlobalConfigInt(key string) (*big.Int, error) {
	return _BokerInterfaceBase.Contract.GlobalConfigInt(&_BokerInterfaceBase.CallOpts, key)
}

// GlobalConfigString is a free data retrieval call binding the contract method 0x378298bc.
//
// Solidity: function globalConfigString(key string) constant returns(string)
func (_BokerInterfaceBase *BokerInterfaceBaseCaller) GlobalConfigString(opts *bind.CallOpts, key string) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _BokerInterfaceBase.contract.Call(opts, out, "globalConfigString", key)
	return *ret0, err
}

// GlobalConfigString is a free data retrieval call binding the contract method 0x378298bc.
//
// Solidity: function globalConfigString(key string) constant returns(string)
func (_BokerInterfaceBase *BokerInterfaceBaseSession) GlobalConfigString(key string) (string, error) {
	return _BokerInterfaceBase.Contract.GlobalConfigString(&_BokerInterfaceBase.CallOpts, key)
}

// GlobalConfigString is a free data retrieval call binding the contract method 0x378298bc.
//
// Solidity: function globalConfigString(key string) constant returns(string)
func (_BokerInterfaceBase *BokerInterfaceBaseCallerSession) GlobalConfigString(key string) (string, error) {
	return _BokerInterfaceBase.Contract.GlobalConfigString(&_BokerInterfaceBase.CallOpts, key)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_BokerInterfaceBase *BokerInterfaceBaseCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BokerInterfaceBase.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_BokerInterfaceBase *BokerInterfaceBaseSession) Owner() (common.Address, error) {
	return _BokerInterfaceBase.Contract.Owner(&_BokerInterfaceBase.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_BokerInterfaceBase *BokerInterfaceBaseCallerSession) Owner() (common.Address, error) {
	return _BokerInterfaceBase.Contract.Owner(&_BokerInterfaceBase.CallOpts)
}

// SetSize is a free data retrieval call binding the contract method 0xee90c99f.
//
// Solidity: function setSize() constant returns(uint256)
func (_BokerInterfaceBase *BokerInterfaceBaseCaller) SetSize(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BokerInterfaceBase.contract.Call(opts, out, "setSize")
	return *ret0, err
}

// SetSize is a free data retrieval call binding the contract method 0xee90c99f.
//
// Solidity: function setSize() constant returns(uint256)
func (_BokerInterfaceBase *BokerInterfaceBaseSession) SetSize() (*big.Int, error) {
	return _BokerInterfaceBase.Contract.SetSize(&_BokerInterfaceBase.CallOpts)
}

// SetSize is a free data retrieval call binding the contract method 0xee90c99f.
//
// Solidity: function setSize() constant returns(uint256)
func (_BokerInterfaceBase *BokerInterfaceBaseCallerSession) SetSize() (*big.Int, error) {
	return _BokerInterfaceBase.Contract.SetSize(&_BokerInterfaceBase.CallOpts)
}

// TestCandidateArray is a free data retrieval call binding the contract method 0x08cae5d3.
//
// Solidity: function testCandidateArray( uint256) constant returns(address)
func (_BokerInterfaceBase *BokerInterfaceBaseCaller) TestCandidateArray(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BokerInterfaceBase.contract.Call(opts, out, "testCandidateArray", arg0)
	return *ret0, err
}

// TestCandidateArray is a free data retrieval call binding the contract method 0x08cae5d3.
//
// Solidity: function testCandidateArray( uint256) constant returns(address)
func (_BokerInterfaceBase *BokerInterfaceBaseSession) TestCandidateArray(arg0 *big.Int) (common.Address, error) {
	return _BokerInterfaceBase.Contract.TestCandidateArray(&_BokerInterfaceBase.CallOpts, arg0)
}

// TestCandidateArray is a free data retrieval call binding the contract method 0x08cae5d3.
//
// Solidity: function testCandidateArray( uint256) constant returns(address)
func (_BokerInterfaceBase *BokerInterfaceBaseCallerSession) TestCandidateArray(arg0 *big.Int) (common.Address, error) {
	return _BokerInterfaceBase.Contract.TestCandidateArray(&_BokerInterfaceBase.CallOpts, arg0)
}

// TestCandidates is a free data retrieval call binding the contract method 0x16439a66.
//
// Solidity: function testCandidates( address) constant returns(index uint256, addr address, tickets uint256)
func (_BokerInterfaceBase *BokerInterfaceBaseCaller) TestCandidates(opts *bind.CallOpts, arg0 common.Address) (struct {
	Index   *big.Int
	Addr    common.Address
	Tickets *big.Int
}, error) {
	ret := new(struct {
		Index   *big.Int
		Addr    common.Address
		Tickets *big.Int
	})
	out := ret
	err := _BokerInterfaceBase.contract.Call(opts, out, "testCandidates", arg0)
	return *ret, err
}

// TestCandidates is a free data retrieval call binding the contract method 0x16439a66.
//
// Solidity: function testCandidates( address) constant returns(index uint256, addr address, tickets uint256)
func (_BokerInterfaceBase *BokerInterfaceBaseSession) TestCandidates(arg0 common.Address) (struct {
	Index   *big.Int
	Addr    common.Address
	Tickets *big.Int
}, error) {
	return _BokerInterfaceBase.Contract.TestCandidates(&_BokerInterfaceBase.CallOpts, arg0)
}

// TestCandidates is a free data retrieval call binding the contract method 0x16439a66.
//
// Solidity: function testCandidates( address) constant returns(index uint256, addr address, tickets uint256)
func (_BokerInterfaceBase *BokerInterfaceBaseCallerSession) TestCandidates(arg0 common.Address) (struct {
	Index   *big.Int
	Addr    common.Address
	Tickets *big.Int
}, error) {
	return _BokerInterfaceBase.Contract.TestCandidates(&_BokerInterfaceBase.CallOpts, arg0)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() constant returns(string)
func (_BokerInterfaceBase *BokerInterfaceBaseCaller) Version(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _BokerInterfaceBase.contract.Call(opts, out, "version")
	return *ret0, err
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() constant returns(string)
func (_BokerInterfaceBase *BokerInterfaceBaseSession) Version() (string, error) {
	return _BokerInterfaceBase.Contract.Version(&_BokerInterfaceBase.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() constant returns(string)
func (_BokerInterfaceBase *BokerInterfaceBaseCallerSession) Version() (string, error) {
	return _BokerInterfaceBase.Contract.Version(&_BokerInterfaceBase.CallOpts)
}

// RotateVote is a paid mutator transaction binding the contract method 0x43bd7d97.
//
// Solidity: function rotateVote(nowTimer uint256) returns()
func (_BokerInterfaceBase *BokerInterfaceBaseTransactor) RotateVote(opts *bind.TransactOpts, nowTimer *big.Int) (*types.Transaction, error) {
	return _BokerInterfaceBase.contract.Transact(opts, "rotateVote", nowTimer)
}

// RotateVote is a paid mutator transaction binding the contract method 0x43bd7d97.
//
// Solidity: function rotateVote(nowTimer uint256) returns()
func (_BokerInterfaceBase *BokerInterfaceBaseSession) RotateVote(nowTimer *big.Int) (*types.Transaction, error) {
	return _BokerInterfaceBase.Contract.RotateVote(&_BokerInterfaceBase.TransactOpts, nowTimer)
}

// RotateVote is a paid mutator transaction binding the contract method 0x43bd7d97.
//
// Solidity: function rotateVote(nowTimer uint256) returns()
func (_BokerInterfaceBase *BokerInterfaceBaseTransactorSession) RotateVote(nowTimer *big.Int) (*types.Transaction, error) {
	return _BokerInterfaceBase.Contract.RotateVote(&_BokerInterfaceBase.TransactOpts, nowTimer)
}

// SetCandidates is a paid mutator transaction binding the contract method 0x8265e61e.
//
// Solidity: function setCandidates(nodeAddress address, tickets uint256) returns(bool)
func (_BokerInterfaceBase *BokerInterfaceBaseTransactor) SetCandidates(opts *bind.TransactOpts, nodeAddress common.Address, tickets *big.Int) (*types.Transaction, error) {
	return _BokerInterfaceBase.contract.Transact(opts, "setCandidates", nodeAddress, tickets)
}

// SetCandidates is a paid mutator transaction binding the contract method 0x8265e61e.
//
// Solidity: function setCandidates(nodeAddress address, tickets uint256) returns(bool)
func (_BokerInterfaceBase *BokerInterfaceBaseSession) SetCandidates(nodeAddress common.Address, tickets *big.Int) (*types.Transaction, error) {
	return _BokerInterfaceBase.Contract.SetCandidates(&_BokerInterfaceBase.TransactOpts, nodeAddress, tickets)
}

// SetCandidates is a paid mutator transaction binding the contract method 0x8265e61e.
//
// Solidity: function setCandidates(nodeAddress address, tickets uint256) returns(bool)
func (_BokerInterfaceBase *BokerInterfaceBaseTransactorSession) SetCandidates(nodeAddress common.Address, tickets *big.Int) (*types.Transaction, error) {
	return _BokerInterfaceBase.Contract.SetCandidates(&_BokerInterfaceBase.TransactOpts, nodeAddress, tickets)
}

// SetManager is a paid mutator transaction binding the contract method 0xd0ebdbe7.
//
// Solidity: function setManager(addrManager address) returns()
func (_BokerInterfaceBase *BokerInterfaceBaseTransactor) SetManager(opts *bind.TransactOpts, addrManager common.Address) (*types.Transaction, error) {
	return _BokerInterfaceBase.contract.Transact(opts, "setManager", addrManager)
}

// SetManager is a paid mutator transaction binding the contract method 0xd0ebdbe7.
//
// Solidity: function setManager(addrManager address) returns()
func (_BokerInterfaceBase *BokerInterfaceBaseSession) SetManager(addrManager common.Address) (*types.Transaction, error) {
	return _BokerInterfaceBase.Contract.SetManager(&_BokerInterfaceBase.TransactOpts, addrManager)
}

// SetManager is a paid mutator transaction binding the contract method 0xd0ebdbe7.
//
// Solidity: function setManager(addrManager address) returns()
func (_BokerInterfaceBase *BokerInterfaceBaseTransactorSession) SetManager(addrManager common.Address) (*types.Transaction, error) {
	return _BokerInterfaceBase.Contract.SetManager(&_BokerInterfaceBase.TransactOpts, addrManager)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_BokerInterfaceBase *BokerInterfaceBaseTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _BokerInterfaceBase.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_BokerInterfaceBase *BokerInterfaceBaseSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _BokerInterfaceBase.Contract.TransferOwnership(&_BokerInterfaceBase.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_BokerInterfaceBase *BokerInterfaceBaseTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _BokerInterfaceBase.Contract.TransferOwnership(&_BokerInterfaceBase.TransactOpts, _newOwner)
}
