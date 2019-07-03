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

// BokerInterfaceABI is the input ABI used to generate the binding from.
const BokerInterfaceABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"getCandidates\",\"outputs\":[{\"name\":\"addresses\",\"type\":\"address[]\"},{\"name\":\"tickets\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"cName\",\"type\":\"string\"}],\"name\":\"contractAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addrCandidate\",\"type\":\"address\"}],\"name\":\"voteCandidate\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"description\",\"type\":\"string\"},{\"name\":\"team\",\"type\":\"string\"},{\"name\":\"name\",\"type\":\"string\"}],\"name\":\"registerCandidate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"key\",\"type\":\"string\"}],\"name\":\"globalConfigString\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"createTime\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"bokerManager\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addrCandidate\",\"type\":\"address\"}],\"name\":\"getCandidate\",\"outputs\":[{\"name\":\"description\",\"type\":\"string\"},{\"name\":\"team\",\"type\":\"string\"},{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"tickets\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addrManager\",\"type\":\"address\"}],\"name\":\"setManager\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isBokerInterface\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"key\",\"type\":\"string\"}],\"name\":\"globalConfigInt\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"cancelAllVotes\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"addrManager\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// BokerInterfaceBin is the compiled bytecode used for deploying new contracts.
const BokerInterfaceBin = `60806040526040805190810160405280600581526020017f312e302e31000000000000000000000000000000000000000000000000000000815250600190805190602001906200005192919062000266565b50426002556001600360146101000a81548160ff0219169083151502179055503480156200007e57600080fd5b5060405160208062001fd48339810180604052810190808051906020019092919050505080336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550620000fd8162000105640100000000026401000000009004565b505062000315565b60008190508073ffffffffffffffffffffffffffffffffffffffff1663519c28826040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b1580156200016f57600080fd5b505af115801562000184573d6000803e3d6000fd5b505050506040513d60208110156200019b57600080fd5b8101908080519060200190929190505050151562000221576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260118152602001807f6e6f7420426f6b65724d616e616765722100000000000000000000000000000081525060200191505060405180910390fd5b80600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10620002a957805160ff1916838001178555620002da565b82800160010185558215620002da579182015b82811115620002d9578251825591602001919060010190620002bc565b5b509050620002e99190620002ed565b5090565b6200031291905b808211156200030e576000816000905550600101620002f4565b5090565b90565b611caf80620003256000396000f3006080604052600436106100db576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff16806306a49fce146100e05780631eb726af1461019457806326bb886d1461023d57806335c3c87614610273578063378298bc146102de57806354fd4d50146103c057806361dcd7ab1461045057806366ebc1c61461047b5780638da5cb5b146104d2578063b7adcf1d14610529578063d0ebdbe7146106c4578063d408efb214610707578063d43c802114610736578063f2fde38b146107b3578063f49f9df8146107f6575b600080fd5b3480156100ec57600080fd5b506100f561080d565b604051808060200180602001838103835285818151815260200191508051906020019060200280838360005b8381101561013c578082015181840152602081019050610121565b50505050905001838103825284818151815260200191508051906020019060200280838360005b8381101561017e578082015181840152602081019050610163565b5050505090500194505050505060405180910390f35b3480156101a057600080fd5b506101fb600480360381019080803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290505050610999565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b610271600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610c47565b005b34801561027f57600080fd5b506102dc600480360381019080803590602001908201803590602001919091929391929390803590602001908201803590602001919091929391929390803590602001908201803590602001919091929391929390505050610eac565b005b3480156102ea57600080fd5b50610345600480360381019080803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290505050611140565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561038557808201518184015260208101905061036a565b50505050905090810190601f1680156103b25780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b3480156103cc57600080fd5b506103d56112d0565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156104155780820151818401526020810190506103fa565b50505050905090810190601f1680156104425780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561045c57600080fd5b5061046561136e565b6040518082815260200191505060405180910390f35b34801561048757600080fd5b50610490611374565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156104de57600080fd5b506104e761139a565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561053557600080fd5b5061056a600480360381019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506113bf565b60405180806020018060200180602001858152602001848103845288818151815260200191508051906020019080838360005b838110156105b857808201518184015260208101905061059d565b50505050905090810190601f1680156105e55780820380516001836020036101000a031916815260200191505b50848103835287818151815260200191508051906020019080838360005b8381101561061e578082015181840152602081019050610603565b50505050905090810190601f16801561064b5780820380516001836020036101000a031916815260200191505b50848103825286818151815260200191508051906020019080838360005b83811015610684578082015181840152602081019050610669565b50505050905090810190601f1680156106b15780820380516001836020036101000a031916815260200191505b5097505050505050505060405180910390f35b3480156106d057600080fd5b50610705600480360381019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506115e9565b005b34801561071357600080fd5b5061071c611650565b604051808215151515815260200191505060405180910390f35b34801561074257600080fd5b5061079d600480360381019080803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290505050611663565b6040518082815260200191505060405180910390f35b3480156107bf57600080fd5b506107f4600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061179d565b005b34801561080257600080fd5b5061080b611804565b005b60608061084e6040805190810160405280600981526020017f426f6b65724e6f64650000000000000000000000000000000000000000000000815250610999565b73ffffffffffffffffffffffffffffffffffffffff166306a49fce6040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401600060405180830381600087803b1580156108b157600080fd5b505af11580156108c5573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f8201168201806040525060408110156108ef57600080fd5b81019080805164010000000081111561090757600080fd5b8281019050602081018481111561091d57600080fd5b815185602082028301116401000000008211171561093a57600080fd5b5050929190602001805164010000000081111561095657600080fd5b8281019050602081018481111561096c57600080fd5b815185602082028301116401000000008211171561098957600080fd5b5050929190505050915091509091565b600080600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663fca1f3c1846040518263ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018080602001828103825283818151815260200191508051906020019080838360005b83811015610a46578082015181840152602081019050610a2b565b50505050905090810190601f168015610a735780820380516001836020036101000a031916815260200191505b5092505050600060405180830381600087803b158015610a9257600080fd5b505af1158015610aa6573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f8201168201806040525060a0811015610ad057600080fd5b81019080805190602001909291908051640100000000811115610af257600080fd5b82810190506020810184811115610b0857600080fd5b8151856001820283011164010000000082111715610b2557600080fd5b5050929190602001805190602001909291908051640100000000811115610b4b57600080fd5b82810190506020810184811115610b6157600080fd5b8151856001820283011164010000000082111715610b7e57600080fd5b505092919060200180519060200190929190505050505092505050600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614151515610c3e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600a8152602001807f616464722069732030210000000000000000000000000000000000000000000081525060200191505060405180910390fd5b80915050919050565b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16635c975abb6040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b158015610ccd57600080fd5b505af1158015610ce1573d6000803e3d6000fd5b505050506040513d6020811015610cf757600080fd5b8101908080519060200190929190505050151515610d7d576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260078152602001807f706175736564210000000000000000000000000000000000000000000000000081525060200191505060405180910390fd5b610dbb6040805190810160405280600981526020017f426f6b65724e6f64650000000000000000000000000000000000000000000000815250610999565b73ffffffffffffffffffffffffffffffffffffffff1663efa5d4313383346040518463ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019350505050600060405180830381600087803b158015610e9157600080fd5b505af1158015610ea5573d6000803e3d6000fd5b5050505050565b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16635c975abb6040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b158015610f3257600080fd5b505af1158015610f46573d6000803e3d6000fd5b505050506040513d6020811015610f5c57600080fd5b8101908080519060200190929190505050151515610fe2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260078152602001807f706175736564210000000000000000000000000000000000000000000000000081525060200191505060405180910390fd5b6110206040805190810160405280600981526020017f426f6b65724e6f64650000000000000000000000000000000000000000000000815250610999565b73ffffffffffffffffffffffffffffffffffffffff166321b3f528338888888888886040518863ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200180602001806020018060200184810384528a8a82818152602001925080828437820191505084810383528888828181526020019250808284378201915050848103825286868281815260200192508082843782019150509a5050505050505050505050600060405180830381600087803b15801561112057600080fd5b505af1158015611134573d6000803e3d6000fd5b50505050505050505050565b6060600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16634a189f35836040518263ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018080602001828103825283818151815260200191508051906020019080838360005b838110156111ec5780820151818401526020810190506111d1565b50505050905090810190601f1680156112195780820380516001836020036101000a031916815260200191505b5092505050600060405180830381600087803b15801561123857600080fd5b505af115801561124c573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f82011682018060405250602081101561127657600080fd5b81019080805164010000000081111561128e57600080fd5b828101905060208101848111156112a457600080fd5b81518560018202830111640100000000821117156112c157600080fd5b50509291905050509050919050565b60018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156113665780601f1061133b57610100808354040283529160200191611366565b820191906000526020600020905b81548152906001019060200180831161134957829003601f168201915b505050505081565b60025481565b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b606080606060006114046040805190810160405280600981526020017f426f6b65724e6f64650000000000000000000000000000000000000000000000815250610999565b73ffffffffffffffffffffffffffffffffffffffff1663b7adcf1d866040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001915050600060405180830381600087803b15801561149e57600080fd5b505af11580156114b2573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f8201168201806040525060808110156114dc57600080fd5b8101908080516401000000008111156114f457600080fd5b8281019050602081018481111561150a57600080fd5b815185600182028301116401000000008211171561152757600080fd5b5050929190602001805164010000000081111561154357600080fd5b8281019050602081018481111561155957600080fd5b815185600182028301116401000000008211171561157657600080fd5b5050929190602001805164010000000081111561159257600080fd5b828101905060208101848111156115a857600080fd5b81518560018202830111640100000000821117156115c557600080fd5b50509291906020018051906020019092919050505093509350935093509193509193565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561164457600080fd5b61164d81611a2c565b50565b600360149054906101000a900460ff1681565b6000600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16633230b078836040518263ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018080602001828103825283818151815260200191508051906020019080838360005b8381101561170f5780820151818401526020810190506116f4565b50505050905090810190601f16801561173c5780820380516001836020036101000a031916815260200191505b5092505050602060405180830381600087803b15801561175b57600080fd5b505af115801561176f573d6000803e3d6000fd5b505050506040513d602081101561178557600080fd5b81019080805190602001909291905050509050919050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156117f857600080fd5b61180181611b89565b50565b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16635c975abb6040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b15801561188a57600080fd5b505af115801561189e573d6000803e3d6000fd5b505050506040513d60208110156118b457600080fd5b810190808051906020019092919050505015151561193a576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260078152602001807f706175736564210000000000000000000000000000000000000000000000000081525060200191505060405180910390fd5b6119786040805190810160405280600981526020017f426f6b65724e6f64650000000000000000000000000000000000000000000000815250610999565b73ffffffffffffffffffffffffffffffffffffffff1663281c231b336040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001915050600060405180830381600087803b158015611a1257600080fd5b505af1158015611a26573d6000803e3d6000fd5b50505050565b60008190508073ffffffffffffffffffffffffffffffffffffffff1663519c28826040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b158015611a9557600080fd5b505af1158015611aa9573d6000803e3d6000fd5b505050506040513d6020811015611abf57600080fd5b81019080805190602001909291905050501515611b44576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260118152602001807f6e6f7420426f6b65724d616e616765722100000000000000000000000000000081525060200191505060405180910390fd5b80600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614151515611bc557600080fd5b8073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550505600a165627a7a723058200f450d834d67ab2c491bb4a3a7fe0599b97b05eaa3c9be83b0cf63886be31be60029`

// DeployBokerInterface deploys a new Ethereum contract, binding an instance of BokerInterface to it.
func DeployBokerInterface(auth *bind.TransactOpts, backend bind.ContractBackend, addrManager common.Address) (common.Address, *types.Transaction, *BokerInterface, error) {
	parsed, err := abi.JSON(strings.NewReader(BokerInterfaceABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BokerInterfaceBin), backend, addrManager)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BokerInterface{BokerInterfaceCaller: BokerInterfaceCaller{contract: contract}, BokerInterfaceTransactor: BokerInterfaceTransactor{contract: contract}}, nil
}

// BokerInterface is an auto generated Go binding around an Ethereum contract.
type BokerInterface struct {
	BokerInterfaceCaller     // Read-only binding to the contract
	BokerInterfaceTransactor // Write-only binding to the contract
}

// BokerInterfaceCaller is an auto generated read-only Go binding around an Ethereum contract.
type BokerInterfaceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BokerInterfaceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BokerInterfaceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BokerInterfaceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BokerInterfaceSession struct {
	Contract     *BokerInterface   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BokerInterfaceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BokerInterfaceCallerSession struct {
	Contract *BokerInterfaceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// BokerInterfaceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BokerInterfaceTransactorSession struct {
	Contract     *BokerInterfaceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// BokerInterfaceRaw is an auto generated low-level Go binding around an Ethereum contract.
type BokerInterfaceRaw struct {
	Contract *BokerInterface // Generic contract binding to access the raw methods on
}

// BokerInterfaceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BokerInterfaceCallerRaw struct {
	Contract *BokerInterfaceCaller // Generic read-only contract binding to access the raw methods on
}

// BokerInterfaceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BokerInterfaceTransactorRaw struct {
	Contract *BokerInterfaceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBokerInterface creates a new instance of BokerInterface, bound to a specific deployed contract.
func NewBokerInterface(address common.Address, backend bind.ContractBackend) (*BokerInterface, error) {
	contract, err := bindBokerInterface(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BokerInterface{BokerInterfaceCaller: BokerInterfaceCaller{contract: contract}, BokerInterfaceTransactor: BokerInterfaceTransactor{contract: contract}}, nil
}

// NewBokerInterfaceCaller creates a new read-only instance of BokerInterface, bound to a specific deployed contract.
func NewBokerInterfaceCaller(address common.Address, caller bind.ContractCaller) (*BokerInterfaceCaller, error) {
	contract, err := bindBokerInterface(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &BokerInterfaceCaller{contract: contract}, nil
}

// NewBokerInterfaceTransactor creates a new write-only instance of BokerInterface, bound to a specific deployed contract.
func NewBokerInterfaceTransactor(address common.Address, transactor bind.ContractTransactor) (*BokerInterfaceTransactor, error) {
	contract, err := bindBokerInterface(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &BokerInterfaceTransactor{contract: contract}, nil
}

// bindBokerInterface binds a generic wrapper to an already deployed contract.
func bindBokerInterface(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BokerInterfaceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BokerInterface *BokerInterfaceRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BokerInterface.Contract.BokerInterfaceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BokerInterface *BokerInterfaceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BokerInterface.Contract.BokerInterfaceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BokerInterface *BokerInterfaceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BokerInterface.Contract.BokerInterfaceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BokerInterface *BokerInterfaceCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BokerInterface.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BokerInterface *BokerInterfaceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BokerInterface.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BokerInterface *BokerInterfaceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BokerInterface.Contract.contract.Transact(opts, method, params...)
}

// BokerManager is a free data retrieval call binding the contract method 0x66ebc1c6.
//
// Solidity: function bokerManager() constant returns(address)
func (_BokerInterface *BokerInterfaceCaller) BokerManager(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BokerInterface.contract.Call(opts, out, "bokerManager")
	return *ret0, err
}

// BokerManager is a free data retrieval call binding the contract method 0x66ebc1c6.
//
// Solidity: function bokerManager() constant returns(address)
func (_BokerInterface *BokerInterfaceSession) BokerManager() (common.Address, error) {
	return _BokerInterface.Contract.BokerManager(&_BokerInterface.CallOpts)
}

// BokerManager is a free data retrieval call binding the contract method 0x66ebc1c6.
//
// Solidity: function bokerManager() constant returns(address)
func (_BokerInterface *BokerInterfaceCallerSession) BokerManager() (common.Address, error) {
	return _BokerInterface.Contract.BokerManager(&_BokerInterface.CallOpts)
}

// ContractAddress is a free data retrieval call binding the contract method 0x1eb726af.
//
// Solidity: function contractAddress(cName string) constant returns(address)
func (_BokerInterface *BokerInterfaceCaller) ContractAddress(opts *bind.CallOpts, cName string) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BokerInterface.contract.Call(opts, out, "contractAddress", cName)
	return *ret0, err
}

// ContractAddress is a free data retrieval call binding the contract method 0x1eb726af.
//
// Solidity: function contractAddress(cName string) constant returns(address)
func (_BokerInterface *BokerInterfaceSession) ContractAddress(cName string) (common.Address, error) {
	return _BokerInterface.Contract.ContractAddress(&_BokerInterface.CallOpts, cName)
}

// ContractAddress is a free data retrieval call binding the contract method 0x1eb726af.
//
// Solidity: function contractAddress(cName string) constant returns(address)
func (_BokerInterface *BokerInterfaceCallerSession) ContractAddress(cName string) (common.Address, error) {
	return _BokerInterface.Contract.ContractAddress(&_BokerInterface.CallOpts, cName)
}

// CreateTime is a free data retrieval call binding the contract method 0x61dcd7ab.
//
// Solidity: function createTime() constant returns(uint256)
func (_BokerInterface *BokerInterfaceCaller) CreateTime(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BokerInterface.contract.Call(opts, out, "createTime")
	return *ret0, err
}

// CreateTime is a free data retrieval call binding the contract method 0x61dcd7ab.
//
// Solidity: function createTime() constant returns(uint256)
func (_BokerInterface *BokerInterfaceSession) CreateTime() (*big.Int, error) {
	return _BokerInterface.Contract.CreateTime(&_BokerInterface.CallOpts)
}

// CreateTime is a free data retrieval call binding the contract method 0x61dcd7ab.
//
// Solidity: function createTime() constant returns(uint256)
func (_BokerInterface *BokerInterfaceCallerSession) CreateTime() (*big.Int, error) {
	return _BokerInterface.Contract.CreateTime(&_BokerInterface.CallOpts)
}

// GetCandidate is a free data retrieval call binding the contract method 0xb7adcf1d.
//
// Solidity: function getCandidate(addrCandidate address) constant returns(description string, team string, name string, tickets uint256)
func (_BokerInterface *BokerInterfaceCaller) GetCandidate(opts *bind.CallOpts, addrCandidate common.Address) (struct {
	Description string
	Team        string
	Name        string
	Tickets     *big.Int
}, error) {
	ret := new(struct {
		Description string
		Team        string
		Name        string
		Tickets     *big.Int
	})
	out := ret
	err := _BokerInterface.contract.Call(opts, out, "getCandidate", addrCandidate)
	return *ret, err
}

// GetCandidate is a free data retrieval call binding the contract method 0xb7adcf1d.
//
// Solidity: function getCandidate(addrCandidate address) constant returns(description string, team string, name string, tickets uint256)
func (_BokerInterface *BokerInterfaceSession) GetCandidate(addrCandidate common.Address) (struct {
	Description string
	Team        string
	Name        string
	Tickets     *big.Int
}, error) {
	return _BokerInterface.Contract.GetCandidate(&_BokerInterface.CallOpts, addrCandidate)
}

// GetCandidate is a free data retrieval call binding the contract method 0xb7adcf1d.
//
// Solidity: function getCandidate(addrCandidate address) constant returns(description string, team string, name string, tickets uint256)
func (_BokerInterface *BokerInterfaceCallerSession) GetCandidate(addrCandidate common.Address) (struct {
	Description string
	Team        string
	Name        string
	Tickets     *big.Int
}, error) {
	return _BokerInterface.Contract.GetCandidate(&_BokerInterface.CallOpts, addrCandidate)
}

// GetCandidates is a free data retrieval call binding the contract method 0x06a49fce.
//
// Solidity: function getCandidates() constant returns(addresses address[], tickets uint256[])
func (_BokerInterface *BokerInterfaceCaller) GetCandidates(opts *bind.CallOpts) (struct {
	Addresses []common.Address
	Tickets   []*big.Int
}, error) {
	ret := new(struct {
		Addresses []common.Address
		Tickets   []*big.Int
	})
	out := ret
	err := _BokerInterface.contract.Call(opts, out, "getCandidates")
	return *ret, err
}

// GetCandidates is a free data retrieval call binding the contract method 0x06a49fce.
//
// Solidity: function getCandidates() constant returns(addresses address[], tickets uint256[])
func (_BokerInterface *BokerInterfaceSession) GetCandidates() (struct {
	Addresses []common.Address
	Tickets   []*big.Int
}, error) {
	return _BokerInterface.Contract.GetCandidates(&_BokerInterface.CallOpts)
}

// GetCandidates is a free data retrieval call binding the contract method 0x06a49fce.
//
// Solidity: function getCandidates() constant returns(addresses address[], tickets uint256[])
func (_BokerInterface *BokerInterfaceCallerSession) GetCandidates() (struct {
	Addresses []common.Address
	Tickets   []*big.Int
}, error) {
	return _BokerInterface.Contract.GetCandidates(&_BokerInterface.CallOpts)
}

// GlobalConfigInt is a free data retrieval call binding the contract method 0xd43c8021.
//
// Solidity: function globalConfigInt(key string) constant returns(uint256)
func (_BokerInterface *BokerInterfaceCaller) GlobalConfigInt(opts *bind.CallOpts, key string) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BokerInterface.contract.Call(opts, out, "globalConfigInt", key)
	return *ret0, err
}

// GlobalConfigInt is a free data retrieval call binding the contract method 0xd43c8021.
//
// Solidity: function globalConfigInt(key string) constant returns(uint256)
func (_BokerInterface *BokerInterfaceSession) GlobalConfigInt(key string) (*big.Int, error) {
	return _BokerInterface.Contract.GlobalConfigInt(&_BokerInterface.CallOpts, key)
}

// GlobalConfigInt is a free data retrieval call binding the contract method 0xd43c8021.
//
// Solidity: function globalConfigInt(key string) constant returns(uint256)
func (_BokerInterface *BokerInterfaceCallerSession) GlobalConfigInt(key string) (*big.Int, error) {
	return _BokerInterface.Contract.GlobalConfigInt(&_BokerInterface.CallOpts, key)
}

// GlobalConfigString is a free data retrieval call binding the contract method 0x378298bc.
//
// Solidity: function globalConfigString(key string) constant returns(string)
func (_BokerInterface *BokerInterfaceCaller) GlobalConfigString(opts *bind.CallOpts, key string) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _BokerInterface.contract.Call(opts, out, "globalConfigString", key)
	return *ret0, err
}

// GlobalConfigString is a free data retrieval call binding the contract method 0x378298bc.
//
// Solidity: function globalConfigString(key string) constant returns(string)
func (_BokerInterface *BokerInterfaceSession) GlobalConfigString(key string) (string, error) {
	return _BokerInterface.Contract.GlobalConfigString(&_BokerInterface.CallOpts, key)
}

// GlobalConfigString is a free data retrieval call binding the contract method 0x378298bc.
//
// Solidity: function globalConfigString(key string) constant returns(string)
func (_BokerInterface *BokerInterfaceCallerSession) GlobalConfigString(key string) (string, error) {
	return _BokerInterface.Contract.GlobalConfigString(&_BokerInterface.CallOpts, key)
}

// IsBokerInterface is a free data retrieval call binding the contract method 0xd408efb2.
//
// Solidity: function isBokerInterface() constant returns(bool)
func (_BokerInterface *BokerInterfaceCaller) IsBokerInterface(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _BokerInterface.contract.Call(opts, out, "isBokerInterface")
	return *ret0, err
}

// IsBokerInterface is a free data retrieval call binding the contract method 0xd408efb2.
//
// Solidity: function isBokerInterface() constant returns(bool)
func (_BokerInterface *BokerInterfaceSession) IsBokerInterface() (bool, error) {
	return _BokerInterface.Contract.IsBokerInterface(&_BokerInterface.CallOpts)
}

// IsBokerInterface is a free data retrieval call binding the contract method 0xd408efb2.
//
// Solidity: function isBokerInterface() constant returns(bool)
func (_BokerInterface *BokerInterfaceCallerSession) IsBokerInterface() (bool, error) {
	return _BokerInterface.Contract.IsBokerInterface(&_BokerInterface.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_BokerInterface *BokerInterfaceCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BokerInterface.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_BokerInterface *BokerInterfaceSession) Owner() (common.Address, error) {
	return _BokerInterface.Contract.Owner(&_BokerInterface.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_BokerInterface *BokerInterfaceCallerSession) Owner() (common.Address, error) {
	return _BokerInterface.Contract.Owner(&_BokerInterface.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() constant returns(string)
func (_BokerInterface *BokerInterfaceCaller) Version(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _BokerInterface.contract.Call(opts, out, "version")
	return *ret0, err
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() constant returns(string)
func (_BokerInterface *BokerInterfaceSession) Version() (string, error) {
	return _BokerInterface.Contract.Version(&_BokerInterface.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() constant returns(string)
func (_BokerInterface *BokerInterfaceCallerSession) Version() (string, error) {
	return _BokerInterface.Contract.Version(&_BokerInterface.CallOpts)
}

// CancelAllVotes is a paid mutator transaction binding the contract method 0xf49f9df8.
//
// Solidity: function cancelAllVotes() returns()
func (_BokerInterface *BokerInterfaceTransactor) CancelAllVotes(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BokerInterface.contract.Transact(opts, "cancelAllVotes")
}

// CancelAllVotes is a paid mutator transaction binding the contract method 0xf49f9df8.
//
// Solidity: function cancelAllVotes() returns()
func (_BokerInterface *BokerInterfaceSession) CancelAllVotes() (*types.Transaction, error) {
	return _BokerInterface.Contract.CancelAllVotes(&_BokerInterface.TransactOpts)
}

// CancelAllVotes is a paid mutator transaction binding the contract method 0xf49f9df8.
//
// Solidity: function cancelAllVotes() returns()
func (_BokerInterface *BokerInterfaceTransactorSession) CancelAllVotes() (*types.Transaction, error) {
	return _BokerInterface.Contract.CancelAllVotes(&_BokerInterface.TransactOpts)
}

// RegisterCandidate is a paid mutator transaction binding the contract method 0x35c3c876.
//
// Solidity: function registerCandidate(description string, team string, name string) returns()
func (_BokerInterface *BokerInterfaceTransactor) RegisterCandidate(opts *bind.TransactOpts, description string, team string, name string) (*types.Transaction, error) {
	return _BokerInterface.contract.Transact(opts, "registerCandidate", description, team, name)
}

// RegisterCandidate is a paid mutator transaction binding the contract method 0x35c3c876.
//
// Solidity: function registerCandidate(description string, team string, name string) returns()
func (_BokerInterface *BokerInterfaceSession) RegisterCandidate(description string, team string, name string) (*types.Transaction, error) {
	return _BokerInterface.Contract.RegisterCandidate(&_BokerInterface.TransactOpts, description, team, name)
}

// RegisterCandidate is a paid mutator transaction binding the contract method 0x35c3c876.
//
// Solidity: function registerCandidate(description string, team string, name string) returns()
func (_BokerInterface *BokerInterfaceTransactorSession) RegisterCandidate(description string, team string, name string) (*types.Transaction, error) {
	return _BokerInterface.Contract.RegisterCandidate(&_BokerInterface.TransactOpts, description, team, name)
}

// SetManager is a paid mutator transaction binding the contract method 0xd0ebdbe7.
//
// Solidity: function setManager(addrManager address) returns()
func (_BokerInterface *BokerInterfaceTransactor) SetManager(opts *bind.TransactOpts, addrManager common.Address) (*types.Transaction, error) {
	return _BokerInterface.contract.Transact(opts, "setManager", addrManager)
}

// SetManager is a paid mutator transaction binding the contract method 0xd0ebdbe7.
//
// Solidity: function setManager(addrManager address) returns()
func (_BokerInterface *BokerInterfaceSession) SetManager(addrManager common.Address) (*types.Transaction, error) {
	return _BokerInterface.Contract.SetManager(&_BokerInterface.TransactOpts, addrManager)
}

// SetManager is a paid mutator transaction binding the contract method 0xd0ebdbe7.
//
// Solidity: function setManager(addrManager address) returns()
func (_BokerInterface *BokerInterfaceTransactorSession) SetManager(addrManager common.Address) (*types.Transaction, error) {
	return _BokerInterface.Contract.SetManager(&_BokerInterface.TransactOpts, addrManager)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_BokerInterface *BokerInterfaceTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _BokerInterface.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_BokerInterface *BokerInterfaceSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _BokerInterface.Contract.TransferOwnership(&_BokerInterface.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_BokerInterface *BokerInterfaceTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _BokerInterface.Contract.TransferOwnership(&_BokerInterface.TransactOpts, _newOwner)
}

// VoteCandidate is a paid mutator transaction binding the contract method 0x26bb886d.
//
// Solidity: function voteCandidate(addrCandidate address) returns()
func (_BokerInterface *BokerInterfaceTransactor) VoteCandidate(opts *bind.TransactOpts, addrCandidate common.Address) (*types.Transaction, error) {
	return _BokerInterface.contract.Transact(opts, "voteCandidate", addrCandidate)
}

// VoteCandidate is a paid mutator transaction binding the contract method 0x26bb886d.
//
// Solidity: function voteCandidate(addrCandidate address) returns()
func (_BokerInterface *BokerInterfaceSession) VoteCandidate(addrCandidate common.Address) (*types.Transaction, error) {
	return _BokerInterface.Contract.VoteCandidate(&_BokerInterface.TransactOpts, addrCandidate)
}

// VoteCandidate is a paid mutator transaction binding the contract method 0x26bb886d.
//
// Solidity: function voteCandidate(addrCandidate address) returns()
func (_BokerInterface *BokerInterfaceTransactorSession) VoteCandidate(addrCandidate common.Address) (*types.Transaction, error) {
	return _BokerInterface.Contract.VoteCandidate(&_BokerInterface.TransactOpts, addrCandidate)
}
