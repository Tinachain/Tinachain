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

// BokerLogDataABI is the input ABI used to generate the binding from.
const BokerLogDataABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"cName\",\"type\":\"string\"}],\"name\":\"contractAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"data\",\"type\":\"string\"},{\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"commonLogAdd\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"key\",\"type\":\"string\"}],\"name\":\"globalConfigString\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"page\",\"type\":\"uint256\"},{\"name\":\"pageSize\",\"type\":\"uint256\"}],\"name\":\"voteRotateLogGet\",\"outputs\":[{\"name\":\"rounds\",\"type\":\"uint256[]\"},{\"name\":\"times\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"createTime\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"bokerManager\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"idx\",\"type\":\"int256\"}],\"name\":\"commonLogGet\",\"outputs\":[{\"name\":\"data\",\"type\":\"string\"},{\"name\":\"time\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addrUser\",\"type\":\"address\"},{\"name\":\"page\",\"type\":\"uint256\"},{\"name\":\"pageSize\",\"type\":\"uint256\"}],\"name\":\"voteLogGet\",\"outputs\":[{\"name\":\"voteTypes\",\"type\":\"uint256[]\"},{\"name\":\"tokenses\",\"type\":\"uint256[]\"},{\"name\":\"addrCandidates\",\"type\":\"address[]\"},{\"name\":\"times\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"voteRotateLogGetIndex\",\"outputs\":[{\"name\":\"round\",\"type\":\"uint256\"},{\"name\":\"time\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"round\",\"type\":\"uint256\"}],\"name\":\"voteRotateLogAdd\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addrManager\",\"type\":\"address\"}],\"name\":\"setManager\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"key\",\"type\":\"string\"}],\"name\":\"globalConfigInt\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addrUser\",\"type\":\"address\"},{\"name\":\"addrCandidate\",\"type\":\"address\"},{\"name\":\"voteType\",\"type\":\"uint256\"},{\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"voteLogAdd\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"addrManager\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// BokerLogDataBin is the compiled bytecode used for deploying new contracts.
const BokerLogDataBin = `60806040526040805190810160405280600581526020017f312e302e3100000000000000000000000000000000000000000000000000000081525060019080519060200190620000519291906200024b565b50426002553480156200006357600080fd5b50604051602080620020488339810180604052810190808051906020019092919050505080336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550620000e281620000ea640100000000026401000000009004565b5050620002fa565b60008190508073ffffffffffffffffffffffffffffffffffffffff1663519c28826040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b1580156200015457600080fd5b505af115801562000169573d6000803e3d6000fd5b505050506040513d60208110156200018057600080fd5b8101908080519060200190929190505050151562000206576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260118152602001807f6e6f7420426f6b65724d616e616765722100000000000000000000000000000081525060200191505060405180910390fd5b80600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106200028e57805160ff1916838001178555620002bf565b82800160010185558215620002bf579182015b82811115620002be578251825591602001919060010190620002a1565b5b509050620002ce9190620002d2565b5090565b620002f791905b80821115620002f3576000816000905550600101620002d9565b5090565b90565b611d3e806200030a6000396000f3006080604052600436106100e6576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff1680631eb726af146100eb57806322df411414610194578063378298bc1461020757806354fd4d50146102e95780636069c8f41461037957806361dcd7ab1461044d57806366ebc1c6146104785780636844bc94146104cf5780638da5cb5b1461057c578063b25532e0146105d3578063bef80e4414610757578063c7de6c911461079f578063d0ebdbe7146107cc578063d43c80211461080f578063e7f218f21461088c578063f2fde38b14610903575b600080fd5b3480156100f757600080fd5b50610152600480360381019080803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290505050610946565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156101a057600080fd5b50610205600480360381019080803590602001908201803590602001908080601f016020809104026020016040519081016040528093929190818152602001838380828437820191505050505050919291929080359060200190929190505050610bf4565b005b34801561021357600080fd5b5061026e600480360381019080803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290505050610c64565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156102ae578082015181840152602081019050610293565b50505050905090810190601f1680156102db5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b3480156102f557600080fd5b506102fe610df4565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561033e578082015181840152602081019050610323565b50505050905090810190601f16801561036b5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561038557600080fd5b506103ae6004803603810190808035906020019092919080359060200190929190505050610e92565b604051808060200180602001838103835285818151815260200191508051906020019060200280838360005b838110156103f55780820151818401526020810190506103da565b50505050905001838103825284818151815260200191508051906020019060200280838360005b8381101561043757808201518184015260208101905061041c565b5050505090500194505050505060405180910390f35b34801561045957600080fd5b50610462610fb2565b6040518082815260200191505060405180910390f35b34801561048457600080fd5b5061048d610fb8565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156104db57600080fd5b506104fa60048036038101908080359060200190929190505050610fde565b6040518080602001838152602001828103825284818151815260200191508051906020019080838360005b83811015610540578082015181840152602081019050610525565b50505050905090810190601f16801561056d5780820380516001836020036101000a031916815260200191505b50935050505060405180910390f35b34801561058857600080fd5b50610591611117565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156105df57600080fd5b50610628600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291908035906020019092919050505061113c565b6040518080602001806020018060200180602001858103855289818151815260200191508051906020019060200280838360005b8381101561067757808201518184015260208101905061065c565b50505050905001858103845288818151815260200191508051906020019060200280838360005b838110156106b957808201518184015260208101905061069e565b50505050905001858103835287818151815260200191508051906020019060200280838360005b838110156106fb5780820151818401526020810190506106e0565b50505050905001858103825286818151815260200191508051906020019060200280838360005b8381101561073d578082015181840152602081019050610722565b505050509050019850505050505050505060405180910390f35b34801561076357600080fd5b50610782600480360381019080803590602001909291905050506113ff565b604051808381526020018281526020019250505060405180910390f35b3480156107ab57600080fd5b506107ca600480360381019080803590602001909291905050506114d9565b005b3480156107d857600080fd5b5061080d600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611535565b005b34801561081b57600080fd5b50610876600480360381019080803590602001908201803590602001908080601f016020809104026020016040519081016040528093929190818152602001838380828437820191505050505050919291929050505061159c565b6040518082815260200191505060405180910390f35b34801561089857600080fd5b50610901600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190803590602001909291905050506116d6565b005b34801561090f57600080fd5b50610944600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611832565b005b600080600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663fca1f3c1846040518263ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018080602001828103825283818151815260200191508051906020019080838360005b838110156109f35780820151818401526020810190506109d8565b50505050905090810190601f168015610a205780820380516001836020036101000a031916815260200191505b5092505050600060405180830381600087803b158015610a3f57600080fd5b505af1158015610a53573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f8201168201806040525060a0811015610a7d57600080fd5b81019080805190602001909291908051640100000000811115610a9f57600080fd5b82810190506020810184811115610ab557600080fd5b8151856001820283011164010000000082111715610ad257600080fd5b5050929190602001805190602001909291908051640100000000811115610af857600080fd5b82810190506020810184811115610b0e57600080fd5b8151856001820283011164010000000082111715610b2b57600080fd5b505092919060200180519060200190929190505050505092505050600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614151515610beb576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600a8152602001807f616464722069732030210000000000000000000000000000000000000000000081525060200191505060405180910390fd5b80915050919050565b6004604080519081016040528084815260200183815250908060018154018082558091505090600182039060005260206000209060020201600090919290919091506000820151816000019080519060200190610c52929190611c6d565b50602082015181600101555050505050565b6060600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16634a189f35836040518263ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018080602001828103825283818151815260200191508051906020019080838360005b83811015610d10578082015181840152602081019050610cf5565b50505050905090810190601f168015610d3d5780820380516001836020036101000a031916815260200191505b5092505050600060405180830381600087803b158015610d5c57600080fd5b505af1158015610d70573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f820116820180604052506020811015610d9a57600080fd5b810190808051640100000000811115610db257600080fd5b82810190506020810184811115610dc857600080fd5b8151856001820283011164010000000082111715610de557600080fd5b50509291905050509050919050565b60018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610e8a5780601f10610e5f57610100808354040283529160200191610e8a565b820191906000526020600020905b815481529060010190602001808311610e6d57829003601f168201915b505050505081565b6060806000806000806000610eb789896006805490506118999092919063ffffffff16565b94509450600185850301925082604051908082528060200260200182016040528015610ef25781602001602082028038833980820191505090505b50965082604051908082528060200260200182016040528015610f245781602001602082028038833980820191505090505b5095508491505b8382111515610fa657600682815481101515610f4357fe5b90600052602060002090600202019050806000015487868403815181101515610f6857fe5b9060200190602002018181525050806001015486868403815181101515610f8b57fe5b90602001906020020181815250508180600101925050610f2b565b50505050509250929050565b60025481565b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60606000806000600480549050111515610ff757611111565b600090506000841215611014576001600480549050039050611018565b8390505b600480549050811015156110325760016004805490500390505b60048181548110151561104157fe5b90600052602060002090600202016000018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156110e65780601f106110bb576101008083540402835291602001916110e6565b820191906000526020600020905b8154815290600101906020018083116110c957829003601f168201915b505050505092506004818154811015156110fc57fe5b90600052602060002090600202016001015491505b50915091565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b606080606080600080600080600080600560008e73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000209550600073ffffffffffffffffffffffffffffffffffffffff168660000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614156111ea576113f0565b6112068c8c88600101805490506118999092919063ffffffff16565b945094506001858503019250826040519080825280602002602001820160405280156112415781602001602082028038833980820191505090505b509950826040519080825280602002602001820160405280156112735781602001602082028038833980820191505090505b509850826040519080825280602002602001820160405280156112a55781602001602082028038833980820191505090505b509750826040519080825280602002602001820160405280156112d75781602001602082028038833980820191505090505b5096508491505b83821115156113ef5785600101828154811015156112f857fe5b9060005260206000209060040201905080600001548a86840381518110151561131d57fe5b906020019060200201818152505080600101548986840381518110151561134057fe5b90602001906020020181815250508060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff168886840381518110151561138357fe5b9060200190602002019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250508060030154878684038151811015156113d457fe5b906020019060200201818152505081806001019250506112de565b5b50505050505093509350935093565b6000806000600680549050841015156114a6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260238152602001807f696e646578206578636565647320766f7465526f746174654c6f6773206c656e81526020017f677468000000000000000000000000000000000000000000000000000000000081525060400191505060405180910390fd5b6006848154811015156114b557fe5b90600052602060002090600202019050806000015492508060010154915050915091565b600660408051908101604052808381526020014281525090806001815401808255809150509060018203906000526020600020906002020160009091929091909150600082015181600001556020820151816001015550505050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561159057600080fd5b61159981611972565b50565b6000600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16633230b078836040518263ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018080602001828103825283818151815260200191508051906020019080838360005b8381101561164857808201518184015260208101905061162d565b50505050905090810190601f1680156116755780820380516001836020036101000a031916815260200191505b5092505050602060405180830381600087803b15801561169457600080fd5b505af11580156116a8573d6000803e3d6000fd5b505050506040513d60208110156116be57600080fd5b81019080805190602001909291905050509050919050565b6000600560008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000209050848160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550806001016080604051908101604052808581526020018481526020018673ffffffffffffffffffffffffffffffffffffffff1681526020014281525090806001815401808255809150509060018203906000526020600020906004020160009091929091909150600082015181600001556020820151816001015560408201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550606082015181600301555050505050505050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561188d57600080fd5b61189681611acf565b50565b6000806000851115156118b8576000808191508090509150915061196a565b6000841115156118c757600193505b60008314156118e5576000915060018503905081819150915061196a565b6118fb8360018603611c3290919063ffffffff16565b9150848210151561194b576001838681151561191357fe5b04019350828581151561192257fe5b0660001415611932576001840393505b6119488360018603611c3290919063ffffffff16565b91505b60018383010390508481101515611963576001850390505b8181915091505b935093915050565b60008190508073ffffffffffffffffffffffffffffffffffffffff1663519c28826040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b1580156119db57600080fd5b505af11580156119ef573d6000803e3d6000fd5b505050506040513d6020811015611a0557600080fd5b81019080805190602001909291905050501515611a8a576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260118152602001807f6e6f7420426f6b65724d616e616765722100000000000000000000000000000081525060200191505060405180910390fd5b80600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614151515611b74576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260168152602001807f6e65774f776e657220616464726573732069732030210000000000000000000081525060200191505060405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b6000806000841415611c475760009150611c66565b8284029050828482811515611c5857fe5b04141515611c6257fe5b8091505b5092915050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10611cae57805160ff1916838001178555611cdc565b82800160010185558215611cdc579182015b82811115611cdb578251825591602001919060010190611cc0565b5b509050611ce99190611ced565b5090565b611d0f91905b80821115611d0b576000816000905550600101611cf3565b5090565b905600a165627a7a72305820992ba3c316814ae1407a842b81966d543a721321705367824e7e518a3329d1e80029`

// DeployBokerLogData deploys a new Ethereum contract, binding an instance of BokerLogData to it.
func DeployBokerLogData(auth *bind.TransactOpts, backend bind.ContractBackend, addrManager common.Address) (common.Address, *types.Transaction, *BokerLogData, error) {
	parsed, err := abi.JSON(strings.NewReader(BokerLogDataABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BokerLogDataBin), backend, addrManager)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BokerLogData{BokerLogDataCaller: BokerLogDataCaller{contract: contract}, BokerLogDataTransactor: BokerLogDataTransactor{contract: contract}}, nil
}

// BokerLogData is an auto generated Go binding around an Ethereum contract.
type BokerLogData struct {
	BokerLogDataCaller     // Read-only binding to the contract
	BokerLogDataTransactor // Write-only binding to the contract
}

// BokerLogDataCaller is an auto generated read-only Go binding around an Ethereum contract.
type BokerLogDataCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BokerLogDataTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BokerLogDataTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BokerLogDataSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BokerLogDataSession struct {
	Contract     *BokerLogData     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BokerLogDataCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BokerLogDataCallerSession struct {
	Contract *BokerLogDataCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// BokerLogDataTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BokerLogDataTransactorSession struct {
	Contract     *BokerLogDataTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// BokerLogDataRaw is an auto generated low-level Go binding around an Ethereum contract.
type BokerLogDataRaw struct {
	Contract *BokerLogData // Generic contract binding to access the raw methods on
}

// BokerLogDataCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BokerLogDataCallerRaw struct {
	Contract *BokerLogDataCaller // Generic read-only contract binding to access the raw methods on
}

// BokerLogDataTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BokerLogDataTransactorRaw struct {
	Contract *BokerLogDataTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBokerLogData creates a new instance of BokerLogData, bound to a specific deployed contract.
func NewBokerLogData(address common.Address, backend bind.ContractBackend) (*BokerLogData, error) {
	contract, err := bindBokerLogData(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BokerLogData{BokerLogDataCaller: BokerLogDataCaller{contract: contract}, BokerLogDataTransactor: BokerLogDataTransactor{contract: contract}}, nil
}

// NewBokerLogDataCaller creates a new read-only instance of BokerLogData, bound to a specific deployed contract.
func NewBokerLogDataCaller(address common.Address, caller bind.ContractCaller) (*BokerLogDataCaller, error) {
	contract, err := bindBokerLogData(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &BokerLogDataCaller{contract: contract}, nil
}

// NewBokerLogDataTransactor creates a new write-only instance of BokerLogData, bound to a specific deployed contract.
func NewBokerLogDataTransactor(address common.Address, transactor bind.ContractTransactor) (*BokerLogDataTransactor, error) {
	contract, err := bindBokerLogData(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &BokerLogDataTransactor{contract: contract}, nil
}

// bindBokerLogData binds a generic wrapper to an already deployed contract.
func bindBokerLogData(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BokerLogDataABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BokerLogData *BokerLogDataRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BokerLogData.Contract.BokerLogDataCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BokerLogData *BokerLogDataRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BokerLogData.Contract.BokerLogDataTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BokerLogData *BokerLogDataRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BokerLogData.Contract.BokerLogDataTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BokerLogData *BokerLogDataCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BokerLogData.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BokerLogData *BokerLogDataTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BokerLogData.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BokerLogData *BokerLogDataTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BokerLogData.Contract.contract.Transact(opts, method, params...)
}

// BokerManager is a free data retrieval call binding the contract method 0x66ebc1c6.
//
// Solidity: function bokerManager() constant returns(address)
func (_BokerLogData *BokerLogDataCaller) BokerManager(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BokerLogData.contract.Call(opts, out, "bokerManager")
	return *ret0, err
}

// BokerManager is a free data retrieval call binding the contract method 0x66ebc1c6.
//
// Solidity: function bokerManager() constant returns(address)
func (_BokerLogData *BokerLogDataSession) BokerManager() (common.Address, error) {
	return _BokerLogData.Contract.BokerManager(&_BokerLogData.CallOpts)
}

// BokerManager is a free data retrieval call binding the contract method 0x66ebc1c6.
//
// Solidity: function bokerManager() constant returns(address)
func (_BokerLogData *BokerLogDataCallerSession) BokerManager() (common.Address, error) {
	return _BokerLogData.Contract.BokerManager(&_BokerLogData.CallOpts)
}

// CommonLogGet is a free data retrieval call binding the contract method 0x6844bc94.
//
// Solidity: function commonLogGet(idx int256) constant returns(data string, time uint256)
func (_BokerLogData *BokerLogDataCaller) CommonLogGet(opts *bind.CallOpts, idx *big.Int) (struct {
	Data string
	Time *big.Int
}, error) {
	ret := new(struct {
		Data string
		Time *big.Int
	})
	out := ret
	err := _BokerLogData.contract.Call(opts, out, "commonLogGet", idx)
	return *ret, err
}

// CommonLogGet is a free data retrieval call binding the contract method 0x6844bc94.
//
// Solidity: function commonLogGet(idx int256) constant returns(data string, time uint256)
func (_BokerLogData *BokerLogDataSession) CommonLogGet(idx *big.Int) (struct {
	Data string
	Time *big.Int
}, error) {
	return _BokerLogData.Contract.CommonLogGet(&_BokerLogData.CallOpts, idx)
}

// CommonLogGet is a free data retrieval call binding the contract method 0x6844bc94.
//
// Solidity: function commonLogGet(idx int256) constant returns(data string, time uint256)
func (_BokerLogData *BokerLogDataCallerSession) CommonLogGet(idx *big.Int) (struct {
	Data string
	Time *big.Int
}, error) {
	return _BokerLogData.Contract.CommonLogGet(&_BokerLogData.CallOpts, idx)
}

// ContractAddress is a free data retrieval call binding the contract method 0x1eb726af.
//
// Solidity: function contractAddress(cName string) constant returns(address)
func (_BokerLogData *BokerLogDataCaller) ContractAddress(opts *bind.CallOpts, cName string) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BokerLogData.contract.Call(opts, out, "contractAddress", cName)
	return *ret0, err
}

// ContractAddress is a free data retrieval call binding the contract method 0x1eb726af.
//
// Solidity: function contractAddress(cName string) constant returns(address)
func (_BokerLogData *BokerLogDataSession) ContractAddress(cName string) (common.Address, error) {
	return _BokerLogData.Contract.ContractAddress(&_BokerLogData.CallOpts, cName)
}

// ContractAddress is a free data retrieval call binding the contract method 0x1eb726af.
//
// Solidity: function contractAddress(cName string) constant returns(address)
func (_BokerLogData *BokerLogDataCallerSession) ContractAddress(cName string) (common.Address, error) {
	return _BokerLogData.Contract.ContractAddress(&_BokerLogData.CallOpts, cName)
}

// CreateTime is a free data retrieval call binding the contract method 0x61dcd7ab.
//
// Solidity: function createTime() constant returns(uint256)
func (_BokerLogData *BokerLogDataCaller) CreateTime(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BokerLogData.contract.Call(opts, out, "createTime")
	return *ret0, err
}

// CreateTime is a free data retrieval call binding the contract method 0x61dcd7ab.
//
// Solidity: function createTime() constant returns(uint256)
func (_BokerLogData *BokerLogDataSession) CreateTime() (*big.Int, error) {
	return _BokerLogData.Contract.CreateTime(&_BokerLogData.CallOpts)
}

// CreateTime is a free data retrieval call binding the contract method 0x61dcd7ab.
//
// Solidity: function createTime() constant returns(uint256)
func (_BokerLogData *BokerLogDataCallerSession) CreateTime() (*big.Int, error) {
	return _BokerLogData.Contract.CreateTime(&_BokerLogData.CallOpts)
}

// GlobalConfigInt is a free data retrieval call binding the contract method 0xd43c8021.
//
// Solidity: function globalConfigInt(key string) constant returns(uint256)
func (_BokerLogData *BokerLogDataCaller) GlobalConfigInt(opts *bind.CallOpts, key string) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BokerLogData.contract.Call(opts, out, "globalConfigInt", key)
	return *ret0, err
}

// GlobalConfigInt is a free data retrieval call binding the contract method 0xd43c8021.
//
// Solidity: function globalConfigInt(key string) constant returns(uint256)
func (_BokerLogData *BokerLogDataSession) GlobalConfigInt(key string) (*big.Int, error) {
	return _BokerLogData.Contract.GlobalConfigInt(&_BokerLogData.CallOpts, key)
}

// GlobalConfigInt is a free data retrieval call binding the contract method 0xd43c8021.
//
// Solidity: function globalConfigInt(key string) constant returns(uint256)
func (_BokerLogData *BokerLogDataCallerSession) GlobalConfigInt(key string) (*big.Int, error) {
	return _BokerLogData.Contract.GlobalConfigInt(&_BokerLogData.CallOpts, key)
}

// GlobalConfigString is a free data retrieval call binding the contract method 0x378298bc.
//
// Solidity: function globalConfigString(key string) constant returns(string)
func (_BokerLogData *BokerLogDataCaller) GlobalConfigString(opts *bind.CallOpts, key string) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _BokerLogData.contract.Call(opts, out, "globalConfigString", key)
	return *ret0, err
}

// GlobalConfigString is a free data retrieval call binding the contract method 0x378298bc.
//
// Solidity: function globalConfigString(key string) constant returns(string)
func (_BokerLogData *BokerLogDataSession) GlobalConfigString(key string) (string, error) {
	return _BokerLogData.Contract.GlobalConfigString(&_BokerLogData.CallOpts, key)
}

// GlobalConfigString is a free data retrieval call binding the contract method 0x378298bc.
//
// Solidity: function globalConfigString(key string) constant returns(string)
func (_BokerLogData *BokerLogDataCallerSession) GlobalConfigString(key string) (string, error) {
	return _BokerLogData.Contract.GlobalConfigString(&_BokerLogData.CallOpts, key)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_BokerLogData *BokerLogDataCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BokerLogData.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_BokerLogData *BokerLogDataSession) Owner() (common.Address, error) {
	return _BokerLogData.Contract.Owner(&_BokerLogData.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_BokerLogData *BokerLogDataCallerSession) Owner() (common.Address, error) {
	return _BokerLogData.Contract.Owner(&_BokerLogData.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() constant returns(string)
func (_BokerLogData *BokerLogDataCaller) Version(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _BokerLogData.contract.Call(opts, out, "version")
	return *ret0, err
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() constant returns(string)
func (_BokerLogData *BokerLogDataSession) Version() (string, error) {
	return _BokerLogData.Contract.Version(&_BokerLogData.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() constant returns(string)
func (_BokerLogData *BokerLogDataCallerSession) Version() (string, error) {
	return _BokerLogData.Contract.Version(&_BokerLogData.CallOpts)
}

// VoteLogGet is a free data retrieval call binding the contract method 0xb25532e0.
//
// Solidity: function voteLogGet(addrUser address, page uint256, pageSize uint256) constant returns(voteTypes uint256[], tokenses uint256[], addrCandidates address[], times uint256[])
func (_BokerLogData *BokerLogDataCaller) VoteLogGet(opts *bind.CallOpts, addrUser common.Address, page *big.Int, pageSize *big.Int) (struct {
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
	err := _BokerLogData.contract.Call(opts, out, "voteLogGet", addrUser, page, pageSize)
	return *ret, err
}

// VoteLogGet is a free data retrieval call binding the contract method 0xb25532e0.
//
// Solidity: function voteLogGet(addrUser address, page uint256, pageSize uint256) constant returns(voteTypes uint256[], tokenses uint256[], addrCandidates address[], times uint256[])
func (_BokerLogData *BokerLogDataSession) VoteLogGet(addrUser common.Address, page *big.Int, pageSize *big.Int) (struct {
	VoteTypes      []*big.Int
	Tokenses       []*big.Int
	AddrCandidates []common.Address
	Times          []*big.Int
}, error) {
	return _BokerLogData.Contract.VoteLogGet(&_BokerLogData.CallOpts, addrUser, page, pageSize)
}

// VoteLogGet is a free data retrieval call binding the contract method 0xb25532e0.
//
// Solidity: function voteLogGet(addrUser address, page uint256, pageSize uint256) constant returns(voteTypes uint256[], tokenses uint256[], addrCandidates address[], times uint256[])
func (_BokerLogData *BokerLogDataCallerSession) VoteLogGet(addrUser common.Address, page *big.Int, pageSize *big.Int) (struct {
	VoteTypes      []*big.Int
	Tokenses       []*big.Int
	AddrCandidates []common.Address
	Times          []*big.Int
}, error) {
	return _BokerLogData.Contract.VoteLogGet(&_BokerLogData.CallOpts, addrUser, page, pageSize)
}

// VoteRotateLogGet is a free data retrieval call binding the contract method 0x6069c8f4.
//
// Solidity: function voteRotateLogGet(page uint256, pageSize uint256) constant returns(rounds uint256[], times uint256[])
func (_BokerLogData *BokerLogDataCaller) VoteRotateLogGet(opts *bind.CallOpts, page *big.Int, pageSize *big.Int) (struct {
	Rounds []*big.Int
	Times  []*big.Int
}, error) {
	ret := new(struct {
		Rounds []*big.Int
		Times  []*big.Int
	})
	out := ret
	err := _BokerLogData.contract.Call(opts, out, "voteRotateLogGet", page, pageSize)
	return *ret, err
}

// VoteRotateLogGet is a free data retrieval call binding the contract method 0x6069c8f4.
//
// Solidity: function voteRotateLogGet(page uint256, pageSize uint256) constant returns(rounds uint256[], times uint256[])
func (_BokerLogData *BokerLogDataSession) VoteRotateLogGet(page *big.Int, pageSize *big.Int) (struct {
	Rounds []*big.Int
	Times  []*big.Int
}, error) {
	return _BokerLogData.Contract.VoteRotateLogGet(&_BokerLogData.CallOpts, page, pageSize)
}

// VoteRotateLogGet is a free data retrieval call binding the contract method 0x6069c8f4.
//
// Solidity: function voteRotateLogGet(page uint256, pageSize uint256) constant returns(rounds uint256[], times uint256[])
func (_BokerLogData *BokerLogDataCallerSession) VoteRotateLogGet(page *big.Int, pageSize *big.Int) (struct {
	Rounds []*big.Int
	Times  []*big.Int
}, error) {
	return _BokerLogData.Contract.VoteRotateLogGet(&_BokerLogData.CallOpts, page, pageSize)
}

// VoteRotateLogGetIndex is a free data retrieval call binding the contract method 0xbef80e44.
//
// Solidity: function voteRotateLogGetIndex(index uint256) constant returns(round uint256, time uint256)
func (_BokerLogData *BokerLogDataCaller) VoteRotateLogGetIndex(opts *bind.CallOpts, index *big.Int) (struct {
	Round *big.Int
	Time  *big.Int
}, error) {
	ret := new(struct {
		Round *big.Int
		Time  *big.Int
	})
	out := ret
	err := _BokerLogData.contract.Call(opts, out, "voteRotateLogGetIndex", index)
	return *ret, err
}

// VoteRotateLogGetIndex is a free data retrieval call binding the contract method 0xbef80e44.
//
// Solidity: function voteRotateLogGetIndex(index uint256) constant returns(round uint256, time uint256)
func (_BokerLogData *BokerLogDataSession) VoteRotateLogGetIndex(index *big.Int) (struct {
	Round *big.Int
	Time  *big.Int
}, error) {
	return _BokerLogData.Contract.VoteRotateLogGetIndex(&_BokerLogData.CallOpts, index)
}

// VoteRotateLogGetIndex is a free data retrieval call binding the contract method 0xbef80e44.
//
// Solidity: function voteRotateLogGetIndex(index uint256) constant returns(round uint256, time uint256)
func (_BokerLogData *BokerLogDataCallerSession) VoteRotateLogGetIndex(index *big.Int) (struct {
	Round *big.Int
	Time  *big.Int
}, error) {
	return _BokerLogData.Contract.VoteRotateLogGetIndex(&_BokerLogData.CallOpts, index)
}

// CommonLogAdd is a paid mutator transaction binding the contract method 0x22df4114.
//
// Solidity: function commonLogAdd(data string, time uint256) returns()
func (_BokerLogData *BokerLogDataTransactor) CommonLogAdd(opts *bind.TransactOpts, data string, time *big.Int) (*types.Transaction, error) {
	return _BokerLogData.contract.Transact(opts, "commonLogAdd", data, time)
}

// CommonLogAdd is a paid mutator transaction binding the contract method 0x22df4114.
//
// Solidity: function commonLogAdd(data string, time uint256) returns()
func (_BokerLogData *BokerLogDataSession) CommonLogAdd(data string, time *big.Int) (*types.Transaction, error) {
	return _BokerLogData.Contract.CommonLogAdd(&_BokerLogData.TransactOpts, data, time)
}

// CommonLogAdd is a paid mutator transaction binding the contract method 0x22df4114.
//
// Solidity: function commonLogAdd(data string, time uint256) returns()
func (_BokerLogData *BokerLogDataTransactorSession) CommonLogAdd(data string, time *big.Int) (*types.Transaction, error) {
	return _BokerLogData.Contract.CommonLogAdd(&_BokerLogData.TransactOpts, data, time)
}

// SetManager is a paid mutator transaction binding the contract method 0xd0ebdbe7.
//
// Solidity: function setManager(addrManager address) returns()
func (_BokerLogData *BokerLogDataTransactor) SetManager(opts *bind.TransactOpts, addrManager common.Address) (*types.Transaction, error) {
	return _BokerLogData.contract.Transact(opts, "setManager", addrManager)
}

// SetManager is a paid mutator transaction binding the contract method 0xd0ebdbe7.
//
// Solidity: function setManager(addrManager address) returns()
func (_BokerLogData *BokerLogDataSession) SetManager(addrManager common.Address) (*types.Transaction, error) {
	return _BokerLogData.Contract.SetManager(&_BokerLogData.TransactOpts, addrManager)
}

// SetManager is a paid mutator transaction binding the contract method 0xd0ebdbe7.
//
// Solidity: function setManager(addrManager address) returns()
func (_BokerLogData *BokerLogDataTransactorSession) SetManager(addrManager common.Address) (*types.Transaction, error) {
	return _BokerLogData.Contract.SetManager(&_BokerLogData.TransactOpts, addrManager)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_BokerLogData *BokerLogDataTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _BokerLogData.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_BokerLogData *BokerLogDataSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _BokerLogData.Contract.TransferOwnership(&_BokerLogData.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_BokerLogData *BokerLogDataTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _BokerLogData.Contract.TransferOwnership(&_BokerLogData.TransactOpts, _newOwner)
}

// VoteLogAdd is a paid mutator transaction binding the contract method 0xe7f218f2.
//
// Solidity: function voteLogAdd(addrUser address, addrCandidate address, voteType uint256, tokens uint256) returns()
func (_BokerLogData *BokerLogDataTransactor) VoteLogAdd(opts *bind.TransactOpts, addrUser common.Address, addrCandidate common.Address, voteType *big.Int, tokens *big.Int) (*types.Transaction, error) {
	return _BokerLogData.contract.Transact(opts, "voteLogAdd", addrUser, addrCandidate, voteType, tokens)
}

// VoteLogAdd is a paid mutator transaction binding the contract method 0xe7f218f2.
//
// Solidity: function voteLogAdd(addrUser address, addrCandidate address, voteType uint256, tokens uint256) returns()
func (_BokerLogData *BokerLogDataSession) VoteLogAdd(addrUser common.Address, addrCandidate common.Address, voteType *big.Int, tokens *big.Int) (*types.Transaction, error) {
	return _BokerLogData.Contract.VoteLogAdd(&_BokerLogData.TransactOpts, addrUser, addrCandidate, voteType, tokens)
}

// VoteLogAdd is a paid mutator transaction binding the contract method 0xe7f218f2.
//
// Solidity: function voteLogAdd(addrUser address, addrCandidate address, voteType uint256, tokens uint256) returns()
func (_BokerLogData *BokerLogDataTransactorSession) VoteLogAdd(addrUser common.Address, addrCandidate common.Address, voteType *big.Int, tokens *big.Int) (*types.Transaction, error) {
	return _BokerLogData.Contract.VoteLogAdd(&_BokerLogData.TransactOpts, addrUser, addrCandidate, voteType, tokens)
}

// VoteRotateLogAdd is a paid mutator transaction binding the contract method 0xc7de6c91.
//
// Solidity: function voteRotateLogAdd(round uint256) returns()
func (_BokerLogData *BokerLogDataTransactor) VoteRotateLogAdd(opts *bind.TransactOpts, round *big.Int) (*types.Transaction, error) {
	return _BokerLogData.contract.Transact(opts, "voteRotateLogAdd", round)
}

// VoteRotateLogAdd is a paid mutator transaction binding the contract method 0xc7de6c91.
//
// Solidity: function voteRotateLogAdd(round uint256) returns()
func (_BokerLogData *BokerLogDataSession) VoteRotateLogAdd(round *big.Int) (*types.Transaction, error) {
	return _BokerLogData.Contract.VoteRotateLogAdd(&_BokerLogData.TransactOpts, round)
}

// VoteRotateLogAdd is a paid mutator transaction binding the contract method 0xc7de6c91.
//
// Solidity: function voteRotateLogAdd(round uint256) returns()
func (_BokerLogData *BokerLogDataTransactorSession) VoteRotateLogAdd(round *big.Int) (*types.Transaction, error) {
	return _BokerLogData.Contract.VoteRotateLogAdd(&_BokerLogData.TransactOpts, round)
}
