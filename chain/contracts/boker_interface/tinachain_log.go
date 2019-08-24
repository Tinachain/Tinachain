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

// BokerLogABI is the input ABI used to generate the binding from.
const BokerLogABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"cName\",\"type\":\"string\"}],\"name\":\"contractAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"data\",\"type\":\"string\"},{\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"commonLogAdd\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"key\",\"type\":\"string\"}],\"name\":\"globalConfigString\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"page\",\"type\":\"uint256\"},{\"name\":\"pageSize\",\"type\":\"uint256\"}],\"name\":\"voteRotateLogGet\",\"outputs\":[{\"name\":\"rounds\",\"type\":\"uint256[]\"},{\"name\":\"times\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"createTime\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"bokerManager\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addrUser\",\"type\":\"address\"},{\"name\":\"page\",\"type\":\"uint256\"},{\"name\":\"pageSize\",\"type\":\"uint256\"}],\"name\":\"voteLogGet\",\"outputs\":[{\"name\":\"voteTypes\",\"type\":\"uint256[]\"},{\"name\":\"tokenses\",\"type\":\"uint256[]\"},{\"name\":\"addrCandidates\",\"type\":\"address[]\"},{\"name\":\"times\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"voteRotateLogGetIndex\",\"outputs\":[{\"name\":\"round\",\"type\":\"uint256\"},{\"name\":\"time\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"round\",\"type\":\"uint256\"}],\"name\":\"voteRotateLogAdd\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addrManager\",\"type\":\"address\"}],\"name\":\"setManager\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"key\",\"type\":\"string\"}],\"name\":\"globalConfigInt\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addrUser\",\"type\":\"address\"},{\"name\":\"addrCandidate\",\"type\":\"address\"},{\"name\":\"voteType\",\"type\":\"uint256\"},{\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"voteLogAdd\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"addrManager\",\"type\":\"address\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// BokerLogBin is the compiled bytecode used for deploying new contracts.
const BokerLogBin = `60806040526040805190810160405280600581526020017f312e302e3100000000000000000000000000000000000000000000000000000081525060019080519060200190620000519291906200023d565b504260025560405160208062001d6a8339810180604052810190808051906020019092919050505080336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550620000d481620000dc640100000000026401000000009004565b5050620002ec565b60008190508073ffffffffffffffffffffffffffffffffffffffff1663519c28826040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b1580156200014657600080fd5b505af11580156200015b573d6000803e3d6000fd5b505050506040513d60208110156200017257600080fd5b81019080805190602001909291905050501515620001f8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260118152602001807f6e6f7420426f6b65724d616e616765722100000000000000000000000000000081525060200191505060405180910390fd5b80600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106200028057805160ff1916838001178555620002b1565b82800160010185558215620002b1579182015b82811115620002b057825182559160200191906001019062000293565b5b509050620002c09190620002c4565b5090565b620002e991905b80821115620002e5576000816000905550600101620002cb565b5090565b90565b611a6e80620002fc6000396000f3006080604052600436106100db576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff1680631eb726af146100e057806322df411414610189578063378298bc146101ce57806354fd4d50146102b05780636069c8f41461034057806361dcd7ab1461041457806366ebc1c61461043f5780638da5cb5b14610496578063b25532e0146104ed578063bef80e4414610671578063c7de6c91146106b9578063d0ebdbe7146106d9578063d43c80211461071c578063e7f218f214610799578063f2fde38b14610803575b600080fd5b3480156100ec57600080fd5b50610147600480360381019080803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290505050610846565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561019557600080fd5b506101cc60048036038101908080359060200190820180359060200191909192939192939080359060200190929190505050610af4565b005b3480156101da57600080fd5b50610235600480360381019080803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290505050610bdf565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561027557808201518184015260208101905061025a565b50505050905090810190601f1680156102a25780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b3480156102bc57600080fd5b506102c5610d6f565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156103055780820151818401526020810190506102ea565b50505050905090810190601f1680156103325780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561034c57600080fd5b506103756004803603810190808035906020019092919080359060200190929190505050610e0d565b604051808060200180602001838103835285818151815260200191508051906020019060200280838360005b838110156103bc5780820151818401526020810190506103a1565b50505050905001838103825284818151815260200191508051906020019060200280838360005b838110156103fe5780820151818401526020810190506103e3565b5050505090500194505050505060405180910390f35b34801561042057600080fd5b50610429610faf565b6040518082815260200191505060405180910390f35b34801561044b57600080fd5b50610454610fb5565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156104a257600080fd5b506104ab610fdb565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156104f957600080fd5b50610542600480360381019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291908035906020019092919080359060200190929190505050611000565b6040518080602001806020018060200180602001858103855289818151815260200191508051906020019060200280838360005b83811015610591578082015181840152602081019050610576565b50505050905001858103845288818151815260200191508051906020019060200280838360005b838110156105d35780820151818401526020810190506105b8565b50505050905001858103835287818151815260200191508051906020019060200280838360005b838110156106155780820151818401526020810190506105fa565b50505050905001858103825286818151815260200191508051906020019060200280838360005b8381101561065757808201518184015260208101905061063c565b505050509050019850505050505050505060405180910390f35b34801561067d57600080fd5b5061069c6004803603810190808035906020019092919050505061127d565b604051808381526020018281526020019250505060405180910390f35b6106d760048036038101908080359060200190929190505050611379565b005b3480156106e557600080fd5b5061071a600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611440565b005b34801561072857600080fd5b50610783600480360381019080803590602001908201803590602001908080601f01602080910402602001604051908101604052809392919081815260200183838082843782019150505050505091929192905050506114a7565b6040518082815260200191505060405180910390f35b610801600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190803590602001909291905050506115e1565b005b34801561080f57600080fd5b50610844600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061171b565b005b600080600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663fca1f3c1846040518263ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018080602001828103825283818151815260200191508051906020019080838360005b838110156108f35780820151818401526020810190506108d8565b50505050905090810190601f1680156109205780820380516001836020036101000a031916815260200191505b5092505050600060405180830381600087803b15801561093f57600080fd5b505af1158015610953573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f8201168201806040525060a081101561097d57600080fd5b8101908080519060200190929190805164010000000081111561099f57600080fd5b828101905060208101848111156109b557600080fd5b81518560018202830111640100000000821117156109d257600080fd5b50509291906020018051906020019092919080516401000000008111156109f857600080fd5b82810190506020810184811115610a0e57600080fd5b8151856001820283011164010000000082111715610a2b57600080fd5b505092919060200180519060200190929190505050505092505050600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614151515610aeb576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600a8152602001807f616464722069732030210000000000000000000000000000000000000000000081525060200191505060405180910390fd5b80915050919050565b610b326040805190810160405280600c81526020017f426f6b65724c6f67446174610000000000000000000000000000000000000000815250610846565b73ffffffffffffffffffffffffffffffffffffffff166322df41148484846040518463ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808060200183815260200182810382528585828181526020019250808284378201915050945050505050600060405180830381600087803b158015610bc257600080fd5b505af1158015610bd6573d6000803e3d6000fd5b50505050505050565b6060600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16634a189f35836040518263ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018080602001828103825283818151815260200191508051906020019080838360005b83811015610c8b578082015181840152602081019050610c70565b50505050905090810190601f168015610cb85780820380516001836020036101000a031916815260200191505b5092505050600060405180830381600087803b158015610cd757600080fd5b505af1158015610ceb573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f820116820180604052506020811015610d1557600080fd5b810190808051640100000000811115610d2d57600080fd5b82810190506020810184811115610d4357600080fd5b8151856001820283011164010000000082111715610d6057600080fd5b50509291905050509050919050565b60018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610e055780601f10610dda57610100808354040283529160200191610e05565b820191906000526020600020905b815481529060010190602001808311610de857829003601f168201915b505050505081565b606080610e4e6040805190810160405280600c81526020017f426f6b65724c6f67446174610000000000000000000000000000000000000000815250610846565b73ffffffffffffffffffffffffffffffffffffffff16636069c8f485856040518363ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018083815260200182815260200192505050600060405180830381600087803b158015610ec457600080fd5b505af1158015610ed8573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f820116820180604052506040811015610f0257600080fd5b810190808051640100000000811115610f1a57600080fd5b82810190506020810184811115610f3057600080fd5b8151856020820283011164010000000082111715610f4d57600080fd5b50509291906020018051640100000000811115610f6957600080fd5b82810190506020810184811115610f7f57600080fd5b8151856020820283011164010000000082111715610f9c57600080fd5b5050929190505050915091509250929050565b60025481565b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6060806060806110446040805190810160405280600c81526020017f426f6b65724c6f67446174610000000000000000000000000000000000000000815250610846565b73ffffffffffffffffffffffffffffffffffffffff1663b25532e08888886040518463ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018381526020018281526020019350505050600060405180830381600087803b1580156110ee57600080fd5b505af1158015611102573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f82011682018060405250608081101561112c57600080fd5b81019080805164010000000081111561114457600080fd5b8281019050602081018481111561115a57600080fd5b815185602082028301116401000000008211171561117757600080fd5b5050929190602001805164010000000081111561119357600080fd5b828101905060208101848111156111a957600080fd5b81518560208202830111640100000000821117156111c657600080fd5b505092919060200180516401000000008111156111e257600080fd5b828101905060208101848111156111f857600080fd5b815185602082028301116401000000008211171561121557600080fd5b5050929190602001805164010000000081111561123157600080fd5b8281019050602081018481111561124757600080fd5b815185602082028301116401000000008211171561126457600080fd5b5050929190505050935093509350935093509350935093565b6000806112be6040805190810160405280600c81526020017f426f6b65724c6f67446174610000000000000000000000000000000000000000815250610846565b73ffffffffffffffffffffffffffffffffffffffff1663bef80e44846040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808281526020019150506040805180830381600087803b15801561132b57600080fd5b505af115801561133f573d6000803e3d6000fd5b505050506040513d604081101561135557600080fd5b81019080805190602001909291908051906020019092919050505091509150915091565b6113b76040805190810160405280600c81526020017f426f6b65724c6f67446174610000000000000000000000000000000000000000815250610846565b73ffffffffffffffffffffffffffffffffffffffff1663c7de6c91826040518263ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040180828152602001915050600060405180830381600087803b15801561142557600080fd5b505af1158015611439573d6000803e3d6000fd5b5050505050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561149b57600080fd5b6114a481611782565b50565b6000600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16633230b078836040518263ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018080602001828103825283818151815260200191508051906020019080838360005b83811015611553578082015181840152602081019050611538565b50505050905090810190601f1680156115805780820380516001836020036101000a031916815260200191505b5092505050602060405180830381600087803b15801561159f57600080fd5b505af11580156115b3573d6000803e3d6000fd5b505050506040513d60208110156115c957600080fd5b81019080805190602001909291905050509050919050565b61161f6040805190810160405280600c81526020017f426f6b65724c6f67446174610000000000000000000000000000000000000000815250610846565b73ffffffffffffffffffffffffffffffffffffffff1663e7f218f2858585856040518563ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001838152602001828152602001945050505050600060405180830381600087803b1580156116fd57600080fd5b505af1158015611711573d6000803e3d6000fd5b5050505050505050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561177657600080fd5b61177f816118df565b50565b60008190508073ffffffffffffffffffffffffffffffffffffffff1663519c28826040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b1580156117eb57600080fd5b505af11580156117ff573d6000803e3d6000fd5b505050506040513d602081101561181557600080fd5b8101908080519060200190929190505050151561189a576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260118152602001807f6e6f7420426f6b65724d616e616765722100000000000000000000000000000081525060200191505060405180910390fd5b80600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614151515611984576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260168152602001807f6e65774f776e657220616464726573732069732030210000000000000000000081525060200191505060405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550505600a165627a7a7230582044e3873e8b7d791f74f5fd581a77a113528e11af8cb16fe72bd668f76fd626980029`

// DeployBokerLog deploys a new Ethereum contract, binding an instance of BokerLog to it.
func DeployBokerLog(auth *bind.TransactOpts, backend bind.ContractBackend, addrManager common.Address) (common.Address, *types.Transaction, *BokerLog, error) {
	parsed, err := abi.JSON(strings.NewReader(BokerLogABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BokerLogBin), backend, addrManager)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BokerLog{BokerLogCaller: BokerLogCaller{contract: contract}, BokerLogTransactor: BokerLogTransactor{contract: contract}}, nil
}

// BokerLog is an auto generated Go binding around an Ethereum contract.
type BokerLog struct {
	BokerLogCaller     // Read-only binding to the contract
	BokerLogTransactor // Write-only binding to the contract
}

// BokerLogCaller is an auto generated read-only Go binding around an Ethereum contract.
type BokerLogCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BokerLogTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BokerLogTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BokerLogSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BokerLogSession struct {
	Contract     *BokerLog         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BokerLogCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BokerLogCallerSession struct {
	Contract *BokerLogCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// BokerLogTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BokerLogTransactorSession struct {
	Contract     *BokerLogTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// BokerLogRaw is an auto generated low-level Go binding around an Ethereum contract.
type BokerLogRaw struct {
	Contract *BokerLog // Generic contract binding to access the raw methods on
}

// BokerLogCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BokerLogCallerRaw struct {
	Contract *BokerLogCaller // Generic read-only contract binding to access the raw methods on
}

// BokerLogTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BokerLogTransactorRaw struct {
	Contract *BokerLogTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBokerLog creates a new instance of BokerLog, bound to a specific deployed contract.
func NewBokerLog(address common.Address, backend bind.ContractBackend) (*BokerLog, error) {
	contract, err := bindBokerLog(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BokerLog{BokerLogCaller: BokerLogCaller{contract: contract}, BokerLogTransactor: BokerLogTransactor{contract: contract}}, nil
}

// NewBokerLogCaller creates a new read-only instance of BokerLog, bound to a specific deployed contract.
func NewBokerLogCaller(address common.Address, caller bind.ContractCaller) (*BokerLogCaller, error) {
	contract, err := bindBokerLog(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &BokerLogCaller{contract: contract}, nil
}

// NewBokerLogTransactor creates a new write-only instance of BokerLog, bound to a specific deployed contract.
func NewBokerLogTransactor(address common.Address, transactor bind.ContractTransactor) (*BokerLogTransactor, error) {
	contract, err := bindBokerLog(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &BokerLogTransactor{contract: contract}, nil
}

// bindBokerLog binds a generic wrapper to an already deployed contract.
func bindBokerLog(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BokerLogABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BokerLog *BokerLogRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BokerLog.Contract.BokerLogCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BokerLog *BokerLogRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BokerLog.Contract.BokerLogTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BokerLog *BokerLogRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BokerLog.Contract.BokerLogTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BokerLog *BokerLogCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BokerLog.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BokerLog *BokerLogTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BokerLog.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BokerLog *BokerLogTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BokerLog.Contract.contract.Transact(opts, method, params...)
}

// BokerManager is a free data retrieval call binding the contract method 0x66ebc1c6.
//
// Solidity: function bokerManager() constant returns(address)
func (_BokerLog *BokerLogCaller) BokerManager(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BokerLog.contract.Call(opts, out, "bokerManager")
	return *ret0, err
}

// BokerManager is a free data retrieval call binding the contract method 0x66ebc1c6.
//
// Solidity: function bokerManager() constant returns(address)
func (_BokerLog *BokerLogSession) BokerManager() (common.Address, error) {
	return _BokerLog.Contract.BokerManager(&_BokerLog.CallOpts)
}

// BokerManager is a free data retrieval call binding the contract method 0x66ebc1c6.
//
// Solidity: function bokerManager() constant returns(address)
func (_BokerLog *BokerLogCallerSession) BokerManager() (common.Address, error) {
	return _BokerLog.Contract.BokerManager(&_BokerLog.CallOpts)
}

// ContractAddress is a free data retrieval call binding the contract method 0x1eb726af.
//
// Solidity: function contractAddress(cName string) constant returns(address)
func (_BokerLog *BokerLogCaller) ContractAddress(opts *bind.CallOpts, cName string) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BokerLog.contract.Call(opts, out, "contractAddress", cName)
	return *ret0, err
}

// ContractAddress is a free data retrieval call binding the contract method 0x1eb726af.
//
// Solidity: function contractAddress(cName string) constant returns(address)
func (_BokerLog *BokerLogSession) ContractAddress(cName string) (common.Address, error) {
	return _BokerLog.Contract.ContractAddress(&_BokerLog.CallOpts, cName)
}

// ContractAddress is a free data retrieval call binding the contract method 0x1eb726af.
//
// Solidity: function contractAddress(cName string) constant returns(address)
func (_BokerLog *BokerLogCallerSession) ContractAddress(cName string) (common.Address, error) {
	return _BokerLog.Contract.ContractAddress(&_BokerLog.CallOpts, cName)
}

// CreateTime is a free data retrieval call binding the contract method 0x61dcd7ab.
//
// Solidity: function createTime() constant returns(uint256)
func (_BokerLog *BokerLogCaller) CreateTime(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BokerLog.contract.Call(opts, out, "createTime")
	return *ret0, err
}

// CreateTime is a free data retrieval call binding the contract method 0x61dcd7ab.
//
// Solidity: function createTime() constant returns(uint256)
func (_BokerLog *BokerLogSession) CreateTime() (*big.Int, error) {
	return _BokerLog.Contract.CreateTime(&_BokerLog.CallOpts)
}

// CreateTime is a free data retrieval call binding the contract method 0x61dcd7ab.
//
// Solidity: function createTime() constant returns(uint256)
func (_BokerLog *BokerLogCallerSession) CreateTime() (*big.Int, error) {
	return _BokerLog.Contract.CreateTime(&_BokerLog.CallOpts)
}

// GlobalConfigInt is a free data retrieval call binding the contract method 0xd43c8021.
//
// Solidity: function globalConfigInt(key string) constant returns(uint256)
func (_BokerLog *BokerLogCaller) GlobalConfigInt(opts *bind.CallOpts, key string) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BokerLog.contract.Call(opts, out, "globalConfigInt", key)
	return *ret0, err
}

// GlobalConfigInt is a free data retrieval call binding the contract method 0xd43c8021.
//
// Solidity: function globalConfigInt(key string) constant returns(uint256)
func (_BokerLog *BokerLogSession) GlobalConfigInt(key string) (*big.Int, error) {
	return _BokerLog.Contract.GlobalConfigInt(&_BokerLog.CallOpts, key)
}

// GlobalConfigInt is a free data retrieval call binding the contract method 0xd43c8021.
//
// Solidity: function globalConfigInt(key string) constant returns(uint256)
func (_BokerLog *BokerLogCallerSession) GlobalConfigInt(key string) (*big.Int, error) {
	return _BokerLog.Contract.GlobalConfigInt(&_BokerLog.CallOpts, key)
}

// GlobalConfigString is a free data retrieval call binding the contract method 0x378298bc.
//
// Solidity: function globalConfigString(key string) constant returns(string)
func (_BokerLog *BokerLogCaller) GlobalConfigString(opts *bind.CallOpts, key string) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _BokerLog.contract.Call(opts, out, "globalConfigString", key)
	return *ret0, err
}

// GlobalConfigString is a free data retrieval call binding the contract method 0x378298bc.
//
// Solidity: function globalConfigString(key string) constant returns(string)
func (_BokerLog *BokerLogSession) GlobalConfigString(key string) (string, error) {
	return _BokerLog.Contract.GlobalConfigString(&_BokerLog.CallOpts, key)
}

// GlobalConfigString is a free data retrieval call binding the contract method 0x378298bc.
//
// Solidity: function globalConfigString(key string) constant returns(string)
func (_BokerLog *BokerLogCallerSession) GlobalConfigString(key string) (string, error) {
	return _BokerLog.Contract.GlobalConfigString(&_BokerLog.CallOpts, key)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_BokerLog *BokerLogCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BokerLog.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_BokerLog *BokerLogSession) Owner() (common.Address, error) {
	return _BokerLog.Contract.Owner(&_BokerLog.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_BokerLog *BokerLogCallerSession) Owner() (common.Address, error) {
	return _BokerLog.Contract.Owner(&_BokerLog.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() constant returns(string)
func (_BokerLog *BokerLogCaller) Version(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _BokerLog.contract.Call(opts, out, "version")
	return *ret0, err
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() constant returns(string)
func (_BokerLog *BokerLogSession) Version() (string, error) {
	return _BokerLog.Contract.Version(&_BokerLog.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() constant returns(string)
func (_BokerLog *BokerLogCallerSession) Version() (string, error) {
	return _BokerLog.Contract.Version(&_BokerLog.CallOpts)
}

// VoteLogGet is a free data retrieval call binding the contract method 0xb25532e0.
//
// Solidity: function voteLogGet(addrUser address, page uint256, pageSize uint256) constant returns(voteTypes uint256[], tokenses uint256[], addrCandidates address[], times uint256[])
func (_BokerLog *BokerLogCaller) VoteLogGet(opts *bind.CallOpts, addrUser common.Address, page *big.Int, pageSize *big.Int) (struct {
	VoteTypes      []*big.Int
	Tokenses       []*big.Int
	AddrCandidates []common.Address
	Times          []*big.Int
}, error) {
	ret := new(struct {
		VoteTypes      []*big.Int
		Tokenses       []*big.Int
		AddrCandidates []common.Address
		Times          []*big.Int
	})
	out := ret
	err := _BokerLog.contract.Call(opts, out, "voteLogGet", addrUser, page, pageSize)
	return *ret, err
}

// VoteLogGet is a free data retrieval call binding the contract method 0xb25532e0.
//
// Solidity: function voteLogGet(addrUser address, page uint256, pageSize uint256) constant returns(voteTypes uint256[], tokenses uint256[], addrCandidates address[], times uint256[])
func (_BokerLog *BokerLogSession) VoteLogGet(addrUser common.Address, page *big.Int, pageSize *big.Int) (struct {
	VoteTypes      []*big.Int
	Tokenses       []*big.Int
	AddrCandidates []common.Address
	Times          []*big.Int
}, error) {
	return _BokerLog.Contract.VoteLogGet(&_BokerLog.CallOpts, addrUser, page, pageSize)
}

// VoteLogGet is a free data retrieval call binding the contract method 0xb25532e0.
//
// Solidity: function voteLogGet(addrUser address, page uint256, pageSize uint256) constant returns(voteTypes uint256[], tokenses uint256[], addrCandidates address[], times uint256[])
func (_BokerLog *BokerLogCallerSession) VoteLogGet(addrUser common.Address, page *big.Int, pageSize *big.Int) (struct {
	VoteTypes      []*big.Int
	Tokenses       []*big.Int
	AddrCandidates []common.Address
	Times          []*big.Int
}, error) {
	return _BokerLog.Contract.VoteLogGet(&_BokerLog.CallOpts, addrUser, page, pageSize)
}

// VoteRotateLogGet is a free data retrieval call binding the contract method 0x6069c8f4.
//
// Solidity: function voteRotateLogGet(page uint256, pageSize uint256) constant returns(rounds uint256[], times uint256[])
func (_BokerLog *BokerLogCaller) VoteRotateLogGet(opts *bind.CallOpts, page *big.Int, pageSize *big.Int) (struct {
	Rounds []*big.Int
	Times  []*big.Int
}, error) {
	ret := new(struct {
		Rounds []*big.Int
		Times  []*big.Int
	})
	out := ret
	err := _BokerLog.contract.Call(opts, out, "voteRotateLogGet", page, pageSize)
	return *ret, err
}

// VoteRotateLogGet is a free data retrieval call binding the contract method 0x6069c8f4.
//
// Solidity: function voteRotateLogGet(page uint256, pageSize uint256) constant returns(rounds uint256[], times uint256[])
func (_BokerLog *BokerLogSession) VoteRotateLogGet(page *big.Int, pageSize *big.Int) (struct {
	Rounds []*big.Int
	Times  []*big.Int
}, error) {
	return _BokerLog.Contract.VoteRotateLogGet(&_BokerLog.CallOpts, page, pageSize)
}

// VoteRotateLogGet is a free data retrieval call binding the contract method 0x6069c8f4.
//
// Solidity: function voteRotateLogGet(page uint256, pageSize uint256) constant returns(rounds uint256[], times uint256[])
func (_BokerLog *BokerLogCallerSession) VoteRotateLogGet(page *big.Int, pageSize *big.Int) (struct {
	Rounds []*big.Int
	Times  []*big.Int
}, error) {
	return _BokerLog.Contract.VoteRotateLogGet(&_BokerLog.CallOpts, page, pageSize)
}

// VoteRotateLogGetIndex is a free data retrieval call binding the contract method 0xbef80e44.
//
// Solidity: function voteRotateLogGetIndex(index uint256) constant returns(round uint256, time uint256)
func (_BokerLog *BokerLogCaller) VoteRotateLogGetIndex(opts *bind.CallOpts, index *big.Int) (struct {
	Round *big.Int
	Time  *big.Int
}, error) {
	ret := new(struct {
		Round *big.Int
		Time  *big.Int
	})
	out := ret
	err := _BokerLog.contract.Call(opts, out, "voteRotateLogGetIndex", index)
	return *ret, err
}

// VoteRotateLogGetIndex is a free data retrieval call binding the contract method 0xbef80e44.
//
// Solidity: function voteRotateLogGetIndex(index uint256) constant returns(round uint256, time uint256)
func (_BokerLog *BokerLogSession) VoteRotateLogGetIndex(index *big.Int) (struct {
	Round *big.Int
	Time  *big.Int
}, error) {
	return _BokerLog.Contract.VoteRotateLogGetIndex(&_BokerLog.CallOpts, index)
}

// VoteRotateLogGetIndex is a free data retrieval call binding the contract method 0xbef80e44.
//
// Solidity: function voteRotateLogGetIndex(index uint256) constant returns(round uint256, time uint256)
func (_BokerLog *BokerLogCallerSession) VoteRotateLogGetIndex(index *big.Int) (struct {
	Round *big.Int
	Time  *big.Int
}, error) {
	return _BokerLog.Contract.VoteRotateLogGetIndex(&_BokerLog.CallOpts, index)
}

// CommonLogAdd is a paid mutator transaction binding the contract method 0x22df4114.
//
// Solidity: function commonLogAdd(data string, time uint256) returns()
func (_BokerLog *BokerLogTransactor) CommonLogAdd(opts *bind.TransactOpts, data string, time *big.Int) (*types.Transaction, error) {
	return _BokerLog.contract.Transact(opts, "commonLogAdd", data, time)
}

// CommonLogAdd is a paid mutator transaction binding the contract method 0x22df4114.
//
// Solidity: function commonLogAdd(data string, time uint256) returns()
func (_BokerLog *BokerLogSession) CommonLogAdd(data string, time *big.Int) (*types.Transaction, error) {
	return _BokerLog.Contract.CommonLogAdd(&_BokerLog.TransactOpts, data, time)
}

// CommonLogAdd is a paid mutator transaction binding the contract method 0x22df4114.
//
// Solidity: function commonLogAdd(data string, time uint256) returns()
func (_BokerLog *BokerLogTransactorSession) CommonLogAdd(data string, time *big.Int) (*types.Transaction, error) {
	return _BokerLog.Contract.CommonLogAdd(&_BokerLog.TransactOpts, data, time)
}

// SetManager is a paid mutator transaction binding the contract method 0xd0ebdbe7.
//
// Solidity: function setManager(addrManager address) returns()
func (_BokerLog *BokerLogTransactor) SetManager(opts *bind.TransactOpts, addrManager common.Address) (*types.Transaction, error) {
	return _BokerLog.contract.Transact(opts, "setManager", addrManager)
}

// SetManager is a paid mutator transaction binding the contract method 0xd0ebdbe7.
//
// Solidity: function setManager(addrManager address) returns()
func (_BokerLog *BokerLogSession) SetManager(addrManager common.Address) (*types.Transaction, error) {
	return _BokerLog.Contract.SetManager(&_BokerLog.TransactOpts, addrManager)
}

// SetManager is a paid mutator transaction binding the contract method 0xd0ebdbe7.
//
// Solidity: function setManager(addrManager address) returns()
func (_BokerLog *BokerLogTransactorSession) SetManager(addrManager common.Address) (*types.Transaction, error) {
	return _BokerLog.Contract.SetManager(&_BokerLog.TransactOpts, addrManager)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_BokerLog *BokerLogTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _BokerLog.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_BokerLog *BokerLogSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _BokerLog.Contract.TransferOwnership(&_BokerLog.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_BokerLog *BokerLogTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _BokerLog.Contract.TransferOwnership(&_BokerLog.TransactOpts, _newOwner)
}

// VoteLogAdd is a paid mutator transaction binding the contract method 0xe7f218f2.
//
// Solidity: function voteLogAdd(addrUser address, addrCandidate address, voteType uint256, tokens uint256) returns()
func (_BokerLog *BokerLogTransactor) VoteLogAdd(opts *bind.TransactOpts, addrUser common.Address, addrCandidate common.Address, voteType *big.Int, tokens *big.Int) (*types.Transaction, error) {
	return _BokerLog.contract.Transact(opts, "voteLogAdd", addrUser, addrCandidate, voteType, tokens)
}

// VoteLogAdd is a paid mutator transaction binding the contract method 0xe7f218f2.
//
// Solidity: function voteLogAdd(addrUser address, addrCandidate address, voteType uint256, tokens uint256) returns()
func (_BokerLog *BokerLogSession) VoteLogAdd(addrUser common.Address, addrCandidate common.Address, voteType *big.Int, tokens *big.Int) (*types.Transaction, error) {
	return _BokerLog.Contract.VoteLogAdd(&_BokerLog.TransactOpts, addrUser, addrCandidate, voteType, tokens)
}

// VoteLogAdd is a paid mutator transaction binding the contract method 0xe7f218f2.
//
// Solidity: function voteLogAdd(addrUser address, addrCandidate address, voteType uint256, tokens uint256) returns()
func (_BokerLog *BokerLogTransactorSession) VoteLogAdd(addrUser common.Address, addrCandidate common.Address, voteType *big.Int, tokens *big.Int) (*types.Transaction, error) {
	return _BokerLog.Contract.VoteLogAdd(&_BokerLog.TransactOpts, addrUser, addrCandidate, voteType, tokens)
}

// VoteRotateLogAdd is a paid mutator transaction binding the contract method 0xc7de6c91.
//
// Solidity: function voteRotateLogAdd(round uint256) returns()
func (_BokerLog *BokerLogTransactor) VoteRotateLogAdd(opts *bind.TransactOpts, round *big.Int) (*types.Transaction, error) {
	return _BokerLog.contract.Transact(opts, "voteRotateLogAdd", round)
}

// VoteRotateLogAdd is a paid mutator transaction binding the contract method 0xc7de6c91.
//
// Solidity: function voteRotateLogAdd(round uint256) returns()
func (_BokerLog *BokerLogSession) VoteRotateLogAdd(round *big.Int) (*types.Transaction, error) {
	return _BokerLog.Contract.VoteRotateLogAdd(&_BokerLog.TransactOpts, round)
}

// VoteRotateLogAdd is a paid mutator transaction binding the contract method 0xc7de6c91.
//
// Solidity: function voteRotateLogAdd(round uint256) returns()
func (_BokerLog *BokerLogTransactorSession) VoteRotateLogAdd(round *big.Int) (*types.Transaction, error) {
	return _BokerLog.Contract.VoteRotateLogAdd(&_BokerLog.TransactOpts, round)
}
