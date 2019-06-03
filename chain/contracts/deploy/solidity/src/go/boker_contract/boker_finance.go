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

// BokerFinanceABI is the input ABI used to generate the binding from.
const BokerFinanceABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"cName\",\"type\":\"string\"}],\"name\":\"contractAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"key\",\"type\":\"string\"}],\"name\":\"globalConfigString\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"createTime\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"bokerManager\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addrManager\",\"type\":\"address\"}],\"name\":\"setManager\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"key\",\"type\":\"string\"}],\"name\":\"globalConfigInt\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addrTo\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"reason\",\"type\":\"uint256\"}],\"name\":\"grantTokenTo\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addrFrom\",\"type\":\"address\"},{\"name\":\"reason\",\"type\":\"uint256\"}],\"name\":\"receiveTokenFrom\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"addrManager\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// BokerFinanceBin is the compiled bytecode used for deploying new contracts.
const BokerFinanceBin = `60806040526040805190810160405280600581526020017f312e302e3100000000000000000000000000000000000000000000000000000081525060019080519060200190620000519291906200024b565b50426002553480156200006357600080fd5b50604051602080620020ec8339810180604052810190808051906020019092919050505080336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550620000e281620000ea640100000000026401000000009004565b5050620002fa565b60008190508073ffffffffffffffffffffffffffffffffffffffff1663519c28826040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b1580156200015457600080fd5b505af115801562000169573d6000803e3d6000fd5b505050506040513d60208110156200018057600080fd5b8101908080519060200190929190505050151562000206576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260118152602001807f6e6f7420426f6b65724d616e616765722100000000000000000000000000000081525060200191505060405180910390fd5b80600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106200028e57805160ff1916838001178555620002bf565b82800160010185558215620002bf579182015b82811115620002be578251825591602001919060010190620002a1565b5b509050620002ce9190620002d2565b5090565b620002f791905b80821115620002f3576000816000905550600101620002d9565b5090565b90565b611de2806200030a6000396000f3006080604052600436106100ba576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff1680631eb726af1461023a578063378298bc146102e35780633ccfd60b146103c557806354fd4d50146103dc57806361dcd7ab1461046c57806366ebc1c6146104975780638da5cb5b146104ee578063d0ebdbe714610545578063d43c802114610588578063f2fde38b14610605578063f4d2cbb414610648578063f9fdc7c01461069f575b6100f86040805190810160405280600881526020017f426f6b65724c6f670000000000000000000000000000000000000000000000008152506106df565b73ffffffffffffffffffffffffffffffffffffffff1663854bb3223330343373ffffffffffffffffffffffffffffffffffffffff16313073ffffffffffffffffffffffffffffffffffffffff16316000600781111561015357fe5b6040518763ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018581526020018481526020018381526020018281526020019650505050505050600060405180830381600087803b15801561022057600080fd5b505af1158015610234573d6000803e3d6000fd5b50505050005b34801561024657600080fd5b506102a1600480360381019080803590602001908201803590602001908080601f01602080910402602001604051908101604052809392919081815260200183838082843782019150505050505091929192905050506106df565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156102ef57600080fd5b5061034a600480360381019080803590602001908201803590602001908080601f016020809104026020016040519081016040528093929190818152602001838380828437820191505050505050919291929050505061098d565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561038a57808201518184015260208101905061036f565b50505050905090810190601f1680156103b75780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b3480156103d157600080fd5b506103da610b1d565b005b3480156103e857600080fd5b506103f1610dd6565b6040518080602001828103825283818151815260200191508051906020019080838360005b83811015610431578082015181840152602081019050610416565b50505050905090810190601f16801561045e5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561047857600080fd5b50610481610e74565b6040518082815260200191505060405180910390f35b3480156104a357600080fd5b506104ac610e7a565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156104fa57600080fd5b50610503610ea0565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561055157600080fd5b50610586600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610ec5565b005b34801561059457600080fd5b506105ef600480360381019080803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290505050610f2c565b6040518082815260200191505060405180910390f35b34801561061157600080fd5b50610646600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611066565b005b34801561065457600080fd5b5061069d600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190803590602001909291905050506110cd565b005b6106dd600480360381019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291908035906020019092919050505061163a565b005b600080600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663fca1f3c1846040518263ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018080602001828103825283818151815260200191508051906020019080838360005b8381101561078c578082015181840152602081019050610771565b50505050905090810190601f1680156107b95780820380516001836020036101000a031916815260200191505b5092505050600060405180830381600087803b1580156107d857600080fd5b505af11580156107ec573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f8201168201806040525060a081101561081657600080fd5b8101908080519060200190929190805164010000000081111561083857600080fd5b8281019050602081018481111561084e57600080fd5b815185600182028301116401000000008211171561086b57600080fd5b505092919060200180519060200190929190805164010000000081111561089157600080fd5b828101905060208101848111156108a757600080fd5b81518560018202830111640100000000821117156108c457600080fd5b505092919060200180519060200190929190505050505092505050600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614151515610984576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600a8152602001807f616464722069732030210000000000000000000000000000000000000000000081525060200191505060405180910390fd5b80915050919050565b6060600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16634a189f35836040518263ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018080602001828103825283818151815260200191508051906020019080838360005b83811015610a39578082015181840152602081019050610a1e565b50505050905090810190601f168015610a665780820380516001836020036101000a031916815260200191505b5092505050600060405180830381600087803b158015610a8557600080fd5b505af1158015610a99573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f820116820180604052506020811015610ac357600080fd5b810190808051640100000000811115610adb57600080fd5b82810190506020810184811115610af157600080fd5b8151856001820283011164010000000082111715610b0e57600080fd5b50509291905050509050919050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515610b7a57600080fd5b3073ffffffffffffffffffffffffffffffffffffffff163190506000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166108fc3073ffffffffffffffffffffffffffffffffffffffff16319081150290604051600060405180830381858888f19350505050158015610c12573d6000803e3d6000fd5b50610c516040805190810160405280600881526020017f426f6b65724c6f670000000000000000000000000000000000000000000000008152506106df565b73ffffffffffffffffffffffffffffffffffffffff1663854bb322306000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff16843073ffffffffffffffffffffffffffffffffffffffff16316000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163160066007811115610cee57fe5b6040518763ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018581526020018481526020018381526020018281526020019650505050505050600060405180830381600087803b158015610dbb57600080fd5b505af1158015610dcf573d6000803e3d6000fd5b5050505050565b60018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610e6c5780601f10610e4157610100808354040283529160200191610e6c565b820191906000526020600020905b815481529060010190602001808311610e4f57829003601f168201915b505050505081565b60025481565b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515610f2057600080fd5b610f2981611b5f565b50565b6000600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16633230b078836040518263ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018080602001828103825283818151815260200191508051906020019080838360005b83811015610fd8578082015181840152602081019050610fbd565b50505050905090810190601f1680156110055780820380516001836020036101000a031916815260200191505b5092505050602060405180830381600087803b15801561102457600080fd5b505af1158015611038573d6000803e3d6000fd5b505050506040513d602081101561104e57600080fd5b81019080805190602001909291905050509050919050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156110c157600080fd5b6110ca81611cbc565b50565b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663d909fb1b6040805190810160405280600581526020017f61646d696e000000000000000000000000000000000000000000000000000000815250336040518363ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040180806020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001828103825284818151815260200191508051906020019080838360005b838110156111df5780820151818401526020810190506111c4565b50505050905090810190601f16801561120c5780820380516001836020036101000a031916815260200191505b509350505050602060405180830381600087803b15801561122c57600080fd5b505af1158015611240573d6000803e3d6000fd5b505050506040513d602081101561125657600080fd5b8101908080519060200190929190505050806114085750600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663d909fb1b6040805190810160405280600881526020017f636f6e7472616374000000000000000000000000000000000000000000000000815250336040518363ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040180806020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001828103825284818151815260200191508051906020019080838360005b8381101561137f578082015181840152602081019050611364565b50505050905090810190601f1680156113ac5780820380516001836020036101000a031916815260200191505b509350505050602060405180830381600087803b1580156113cc57600080fd5b505af11580156113e0573d6000803e3d6000fd5b505050506040513d60208110156113f657600080fd5b81019080805190602001909291905050505b151561147c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600f8152602001807f6e6f7420617574686f72697a656421000000000000000000000000000000000081525060200191505060405180910390fd5b8273ffffffffffffffffffffffffffffffffffffffff166108fc839081150290604051600060405180830381858888f193505050501580156114c2573d6000803e3d6000fd5b506115016040805190810160405280600881526020017f426f6b65724c6f670000000000000000000000000000000000000000000000008152506106df565b73ffffffffffffffffffffffffffffffffffffffff1663854bb3223085853073ffffffffffffffffffffffffffffffffffffffff16318873ffffffffffffffffffffffffffffffffffffffff1631876040518763ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018581526020018481526020018381526020018281526020019650505050505050600060405180830381600087803b15801561161d57600080fd5b505af1158015611631573d6000803e3d6000fd5b50505050505050565b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663d909fb1b6040805190810160405280600581526020017f61646d696e000000000000000000000000000000000000000000000000000000815250336040518363ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040180806020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001828103825284818151815260200191508051906020019080838360005b8381101561174c578082015181840152602081019050611731565b50505050905090810190601f1680156117795780820380516001836020036101000a031916815260200191505b509350505050602060405180830381600087803b15801561179957600080fd5b505af11580156117ad573d6000803e3d6000fd5b505050506040513d60208110156117c357600080fd5b8101908080519060200190929190505050806119755750600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663d909fb1b6040805190810160405280600881526020017f636f6e7472616374000000000000000000000000000000000000000000000000815250336040518363ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040180806020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001828103825284818151815260200191508051906020019080838360005b838110156118ec5780820151818401526020810190506118d1565b50505050905090810190601f1680156119195780820380516001836020036101000a031916815260200191505b509350505050602060405180830381600087803b15801561193957600080fd5b505af115801561194d573d6000803e3d6000fd5b505050506040513d602081101561196357600080fd5b81019080805190602001909291905050505b15156119e9576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600f8152602001807f6e6f7420617574686f72697a656421000000000000000000000000000000000081525060200191505060405180910390fd5b611a276040805190810160405280600881526020017f426f6b65724c6f670000000000000000000000000000000000000000000000008152506106df565b73ffffffffffffffffffffffffffffffffffffffff1663854bb3228330348673ffffffffffffffffffffffffffffffffffffffff16313073ffffffffffffffffffffffffffffffffffffffff1631876040518763ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018581526020018481526020018381526020018281526020019650505050505050600060405180830381600087803b158015611b4357600080fd5b505af1158015611b57573d6000803e3d6000fd5b505050505050565b60008190508073ffffffffffffffffffffffffffffffffffffffff1663519c28826040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b158015611bc857600080fd5b505af1158015611bdc573d6000803e3d6000fd5b505050506040513d6020811015611bf257600080fd5b81019080805190602001909291905050501515611c77576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260118152602001807f6e6f7420426f6b65724d616e616765722100000000000000000000000000000081525060200191505060405180910390fd5b80600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614151515611cf857600080fd5b8073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550505600a165627a7a723058208572c16e3d1f7228036b3ac59216b2cdceeb7f6f3781a711cd9e8a55431f31ae0029`

// DeployBokerFinance deploys a new Ethereum contract, binding an instance of BokerFinance to it.
func DeployBokerFinance(auth *bind.TransactOpts, backend bind.ContractBackend, addrManager common.Address) (common.Address, *types.Transaction, *BokerFinance, error) {
	parsed, err := abi.JSON(strings.NewReader(BokerFinanceABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BokerFinanceBin), backend, addrManager)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BokerFinance{BokerFinanceCaller: BokerFinanceCaller{contract: contract}, BokerFinanceTransactor: BokerFinanceTransactor{contract: contract}}, nil
}

// BokerFinance is an auto generated Go binding around an Ethereum contract.
type BokerFinance struct {
	BokerFinanceCaller     // Read-only binding to the contract
	BokerFinanceTransactor // Write-only binding to the contract
}

// BokerFinanceCaller is an auto generated read-only Go binding around an Ethereum contract.
type BokerFinanceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BokerFinanceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BokerFinanceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BokerFinanceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BokerFinanceSession struct {
	Contract     *BokerFinance     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BokerFinanceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BokerFinanceCallerSession struct {
	Contract *BokerFinanceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// BokerFinanceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BokerFinanceTransactorSession struct {
	Contract     *BokerFinanceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// BokerFinanceRaw is an auto generated low-level Go binding around an Ethereum contract.
type BokerFinanceRaw struct {
	Contract *BokerFinance // Generic contract binding to access the raw methods on
}

// BokerFinanceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BokerFinanceCallerRaw struct {
	Contract *BokerFinanceCaller // Generic read-only contract binding to access the raw methods on
}

// BokerFinanceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BokerFinanceTransactorRaw struct {
	Contract *BokerFinanceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBokerFinance creates a new instance of BokerFinance, bound to a specific deployed contract.
func NewBokerFinance(address common.Address, backend bind.ContractBackend) (*BokerFinance, error) {
	contract, err := bindBokerFinance(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BokerFinance{BokerFinanceCaller: BokerFinanceCaller{contract: contract}, BokerFinanceTransactor: BokerFinanceTransactor{contract: contract}}, nil
}

// NewBokerFinanceCaller creates a new read-only instance of BokerFinance, bound to a specific deployed contract.
func NewBokerFinanceCaller(address common.Address, caller bind.ContractCaller) (*BokerFinanceCaller, error) {
	contract, err := bindBokerFinance(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &BokerFinanceCaller{contract: contract}, nil
}

// NewBokerFinanceTransactor creates a new write-only instance of BokerFinance, bound to a specific deployed contract.
func NewBokerFinanceTransactor(address common.Address, transactor bind.ContractTransactor) (*BokerFinanceTransactor, error) {
	contract, err := bindBokerFinance(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &BokerFinanceTransactor{contract: contract}, nil
}

// bindBokerFinance binds a generic wrapper to an already deployed contract.
func bindBokerFinance(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BokerFinanceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BokerFinance *BokerFinanceRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BokerFinance.Contract.BokerFinanceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BokerFinance *BokerFinanceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BokerFinance.Contract.BokerFinanceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BokerFinance *BokerFinanceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BokerFinance.Contract.BokerFinanceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BokerFinance *BokerFinanceCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BokerFinance.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BokerFinance *BokerFinanceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BokerFinance.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BokerFinance *BokerFinanceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BokerFinance.Contract.contract.Transact(opts, method, params...)
}

// BokerManager is a free data retrieval call binding the contract method 0x66ebc1c6.
//
// Solidity: function bokerManager() constant returns(address)
func (_BokerFinance *BokerFinanceCaller) BokerManager(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BokerFinance.contract.Call(opts, out, "bokerManager")
	return *ret0, err
}

// BokerManager is a free data retrieval call binding the contract method 0x66ebc1c6.
//
// Solidity: function bokerManager() constant returns(address)
func (_BokerFinance *BokerFinanceSession) BokerManager() (common.Address, error) {
	return _BokerFinance.Contract.BokerManager(&_BokerFinance.CallOpts)
}

// BokerManager is a free data retrieval call binding the contract method 0x66ebc1c6.
//
// Solidity: function bokerManager() constant returns(address)
func (_BokerFinance *BokerFinanceCallerSession) BokerManager() (common.Address, error) {
	return _BokerFinance.Contract.BokerManager(&_BokerFinance.CallOpts)
}

// ContractAddress is a free data retrieval call binding the contract method 0x1eb726af.
//
// Solidity: function contractAddress(cName string) constant returns(address)
func (_BokerFinance *BokerFinanceCaller) ContractAddress(opts *bind.CallOpts, cName string) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BokerFinance.contract.Call(opts, out, "contractAddress", cName)
	return *ret0, err
}

// ContractAddress is a free data retrieval call binding the contract method 0x1eb726af.
//
// Solidity: function contractAddress(cName string) constant returns(address)
func (_BokerFinance *BokerFinanceSession) ContractAddress(cName string) (common.Address, error) {
	return _BokerFinance.Contract.ContractAddress(&_BokerFinance.CallOpts, cName)
}

// ContractAddress is a free data retrieval call binding the contract method 0x1eb726af.
//
// Solidity: function contractAddress(cName string) constant returns(address)
func (_BokerFinance *BokerFinanceCallerSession) ContractAddress(cName string) (common.Address, error) {
	return _BokerFinance.Contract.ContractAddress(&_BokerFinance.CallOpts, cName)
}

// CreateTime is a free data retrieval call binding the contract method 0x61dcd7ab.
//
// Solidity: function createTime() constant returns(uint256)
func (_BokerFinance *BokerFinanceCaller) CreateTime(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BokerFinance.contract.Call(opts, out, "createTime")
	return *ret0, err
}

// CreateTime is a free data retrieval call binding the contract method 0x61dcd7ab.
//
// Solidity: function createTime() constant returns(uint256)
func (_BokerFinance *BokerFinanceSession) CreateTime() (*big.Int, error) {
	return _BokerFinance.Contract.CreateTime(&_BokerFinance.CallOpts)
}

// CreateTime is a free data retrieval call binding the contract method 0x61dcd7ab.
//
// Solidity: function createTime() constant returns(uint256)
func (_BokerFinance *BokerFinanceCallerSession) CreateTime() (*big.Int, error) {
	return _BokerFinance.Contract.CreateTime(&_BokerFinance.CallOpts)
}

// GlobalConfigInt is a free data retrieval call binding the contract method 0xd43c8021.
//
// Solidity: function globalConfigInt(key string) constant returns(uint256)
func (_BokerFinance *BokerFinanceCaller) GlobalConfigInt(opts *bind.CallOpts, key string) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BokerFinance.contract.Call(opts, out, "globalConfigInt", key)
	return *ret0, err
}

// GlobalConfigInt is a free data retrieval call binding the contract method 0xd43c8021.
//
// Solidity: function globalConfigInt(key string) constant returns(uint256)
func (_BokerFinance *BokerFinanceSession) GlobalConfigInt(key string) (*big.Int, error) {
	return _BokerFinance.Contract.GlobalConfigInt(&_BokerFinance.CallOpts, key)
}

// GlobalConfigInt is a free data retrieval call binding the contract method 0xd43c8021.
//
// Solidity: function globalConfigInt(key string) constant returns(uint256)
func (_BokerFinance *BokerFinanceCallerSession) GlobalConfigInt(key string) (*big.Int, error) {
	return _BokerFinance.Contract.GlobalConfigInt(&_BokerFinance.CallOpts, key)
}

// GlobalConfigString is a free data retrieval call binding the contract method 0x378298bc.
//
// Solidity: function globalConfigString(key string) constant returns(string)
func (_BokerFinance *BokerFinanceCaller) GlobalConfigString(opts *bind.CallOpts, key string) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _BokerFinance.contract.Call(opts, out, "globalConfigString", key)
	return *ret0, err
}

// GlobalConfigString is a free data retrieval call binding the contract method 0x378298bc.
//
// Solidity: function globalConfigString(key string) constant returns(string)
func (_BokerFinance *BokerFinanceSession) GlobalConfigString(key string) (string, error) {
	return _BokerFinance.Contract.GlobalConfigString(&_BokerFinance.CallOpts, key)
}

// GlobalConfigString is a free data retrieval call binding the contract method 0x378298bc.
//
// Solidity: function globalConfigString(key string) constant returns(string)
func (_BokerFinance *BokerFinanceCallerSession) GlobalConfigString(key string) (string, error) {
	return _BokerFinance.Contract.GlobalConfigString(&_BokerFinance.CallOpts, key)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_BokerFinance *BokerFinanceCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BokerFinance.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_BokerFinance *BokerFinanceSession) Owner() (common.Address, error) {
	return _BokerFinance.Contract.Owner(&_BokerFinance.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_BokerFinance *BokerFinanceCallerSession) Owner() (common.Address, error) {
	return _BokerFinance.Contract.Owner(&_BokerFinance.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() constant returns(string)
func (_BokerFinance *BokerFinanceCaller) Version(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _BokerFinance.contract.Call(opts, out, "version")
	return *ret0, err
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() constant returns(string)
func (_BokerFinance *BokerFinanceSession) Version() (string, error) {
	return _BokerFinance.Contract.Version(&_BokerFinance.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() constant returns(string)
func (_BokerFinance *BokerFinanceCallerSession) Version() (string, error) {
	return _BokerFinance.Contract.Version(&_BokerFinance.CallOpts)
}

// GrantTokenTo is a paid mutator transaction binding the contract method 0xf4d2cbb4.
//
// Solidity: function grantTokenTo(addrTo address, amount uint256, reason uint256) returns()
func (_BokerFinance *BokerFinanceTransactor) GrantTokenTo(opts *bind.TransactOpts, addrTo common.Address, amount *big.Int, reason *big.Int) (*types.Transaction, error) {
	return _BokerFinance.contract.Transact(opts, "grantTokenTo", addrTo, amount, reason)
}

// GrantTokenTo is a paid mutator transaction binding the contract method 0xf4d2cbb4.
//
// Solidity: function grantTokenTo(addrTo address, amount uint256, reason uint256) returns()
func (_BokerFinance *BokerFinanceSession) GrantTokenTo(addrTo common.Address, amount *big.Int, reason *big.Int) (*types.Transaction, error) {
	return _BokerFinance.Contract.GrantTokenTo(&_BokerFinance.TransactOpts, addrTo, amount, reason)
}

// GrantTokenTo is a paid mutator transaction binding the contract method 0xf4d2cbb4.
//
// Solidity: function grantTokenTo(addrTo address, amount uint256, reason uint256) returns()
func (_BokerFinance *BokerFinanceTransactorSession) GrantTokenTo(addrTo common.Address, amount *big.Int, reason *big.Int) (*types.Transaction, error) {
	return _BokerFinance.Contract.GrantTokenTo(&_BokerFinance.TransactOpts, addrTo, amount, reason)
}

// ReceiveTokenFrom is a paid mutator transaction binding the contract method 0xf9fdc7c0.
//
// Solidity: function receiveTokenFrom(addrFrom address, reason uint256) returns()
func (_BokerFinance *BokerFinanceTransactor) ReceiveTokenFrom(opts *bind.TransactOpts, addrFrom common.Address, reason *big.Int) (*types.Transaction, error) {
	return _BokerFinance.contract.Transact(opts, "receiveTokenFrom", addrFrom, reason)
}

// ReceiveTokenFrom is a paid mutator transaction binding the contract method 0xf9fdc7c0.
//
// Solidity: function receiveTokenFrom(addrFrom address, reason uint256) returns()
func (_BokerFinance *BokerFinanceSession) ReceiveTokenFrom(addrFrom common.Address, reason *big.Int) (*types.Transaction, error) {
	return _BokerFinance.Contract.ReceiveTokenFrom(&_BokerFinance.TransactOpts, addrFrom, reason)
}

// ReceiveTokenFrom is a paid mutator transaction binding the contract method 0xf9fdc7c0.
//
// Solidity: function receiveTokenFrom(addrFrom address, reason uint256) returns()
func (_BokerFinance *BokerFinanceTransactorSession) ReceiveTokenFrom(addrFrom common.Address, reason *big.Int) (*types.Transaction, error) {
	return _BokerFinance.Contract.ReceiveTokenFrom(&_BokerFinance.TransactOpts, addrFrom, reason)
}

// SetManager is a paid mutator transaction binding the contract method 0xd0ebdbe7.
//
// Solidity: function setManager(addrManager address) returns()
func (_BokerFinance *BokerFinanceTransactor) SetManager(opts *bind.TransactOpts, addrManager common.Address) (*types.Transaction, error) {
	return _BokerFinance.contract.Transact(opts, "setManager", addrManager)
}

// SetManager is a paid mutator transaction binding the contract method 0xd0ebdbe7.
//
// Solidity: function setManager(addrManager address) returns()
func (_BokerFinance *BokerFinanceSession) SetManager(addrManager common.Address) (*types.Transaction, error) {
	return _BokerFinance.Contract.SetManager(&_BokerFinance.TransactOpts, addrManager)
}

// SetManager is a paid mutator transaction binding the contract method 0xd0ebdbe7.
//
// Solidity: function setManager(addrManager address) returns()
func (_BokerFinance *BokerFinanceTransactorSession) SetManager(addrManager common.Address) (*types.Transaction, error) {
	return _BokerFinance.Contract.SetManager(&_BokerFinance.TransactOpts, addrManager)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_BokerFinance *BokerFinanceTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _BokerFinance.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_BokerFinance *BokerFinanceSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _BokerFinance.Contract.TransferOwnership(&_BokerFinance.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_BokerFinance *BokerFinanceTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _BokerFinance.Contract.TransferOwnership(&_BokerFinance.TransactOpts, _newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_BokerFinance *BokerFinanceTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BokerFinance.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_BokerFinance *BokerFinanceSession) Withdraw() (*types.Transaction, error) {
	return _BokerFinance.Contract.Withdraw(&_BokerFinance.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_BokerFinance *BokerFinanceTransactorSession) Withdraw() (*types.Transaction, error) {
	return _BokerFinance.Contract.Withdraw(&_BokerFinance.TransactOpts)
}
