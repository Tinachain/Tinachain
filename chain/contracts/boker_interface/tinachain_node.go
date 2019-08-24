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

// BokerNodeABI is the input ABI used to generate the binding from.
const BokerNodeABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"getCandidates\",\"outputs\":[{\"name\":\"addresses\",\"type\":\"address[]\"},{\"name\":\"tickets\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"cName\",\"type\":\"string\"}],\"name\":\"contractAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addrVoter\",\"type\":\"address\"}],\"name\":\"cancelAllVotes\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"key\",\"type\":\"string\"}],\"name\":\"globalConfigString\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"nowTimer\",\"type\":\"uint256\"}],\"name\":\"rotateVote\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"createTime\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"bokerManager\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addrCandidate\",\"type\":\"address\"}],\"name\":\"getCandidate\",\"outputs\":[{\"name\":\"description\",\"type\":\"string\"},{\"name\":\"team\",\"type\":\"string\"},{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"tickets\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addrManager\",\"type\":\"address\"}],\"name\":\"setManager\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getBlacks\",\"outputs\":[{\"name\":\"addresses\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"key\",\"type\":\"string\"}],\"name\":\"globalConfigInt\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addrVoter\",\"type\":\"address\"},{\"name\":\"addrCandidate\",\"type\":\"address\"},{\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"vote\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addrCandidate\",\"type\":\"address\"},{\"name\":\"description\",\"type\":\"string\"},{\"name\":\"team\",\"type\":\"string\"},{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"tickets\",\"type\":\"uint256\"}],\"name\":\"registerCandidate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"addrManager\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// BokerNodeBin is the compiled bytecode used for deploying new contracts.
const BokerNodeBin = `60806040526040805190810160405280600581526020017f312e302e3100000000000000000000000000000000000000000000000000000081525060019080519060200190620000519291906200024b565b50426002553480156200006357600080fd5b506040516020806200291a8339810180604052810190808051906020019092919050505080336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550620000e281620000ea640100000000026401000000009004565b5050620002fa565b60008190508073ffffffffffffffffffffffffffffffffffffffff1663519c28826040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b1580156200015457600080fd5b505af115801562000169573d6000803e3d6000fd5b505050506040513d60208110156200018057600080fd5b8101908080519060200190929190505050151562000206576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260118152602001807f6e6f7420426f6b65724d616e616765722100000000000000000000000000000081525060200191505060405180910390fd5b80600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106200028e57805160ff1916838001178555620002bf565b82800160010185558215620002bf579182015b82811115620002be578251825591602001919060010190620002a1565b5b509050620002ce9190620002d2565b5090565b620002f791905b80821115620002f3576000816000905550600101620002d9565b5090565b90565b612610806200030a6000396000f3006080604052600436106100e6576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff16806306a49fce146100eb5780631eb726af1461019f578063281c231b14610248578063378298bc1461028b57806343bd7d971461036d57806354fd4d501461039a57806361dcd7ab1461042a57806366ebc1c6146104555780638da5cb5b146104ac578063b7adcf1d14610503578063d0ebdbe71461069e578063d38ff7a1146106e1578063d43c80211461074d578063efa5d431146107ca578063f2fde38b14610837578063f581978b1461087a575b600080fd5b3480156100f757600080fd5b5061010061090f565b604051808060200180602001838103835285818151815260200191508051906020019060200280838360005b8381101561014757808201518184015260208101905061012c565b50505050905001838103825284818151815260200191508051906020019060200280838360005b8381101561018957808201518184015260208101905061016e565b5050505090500194505050505060405180910390f35b3480156101ab57600080fd5b50610206600480360381019080803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290505050610a9b565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561025457600080fd5b50610289600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610d49565b005b34801561029757600080fd5b506102f2600480360381019080803590602001908201803590602001908080601f016020809104026020016040519081016040528093929190818152602001838380828437820191505050505050919291929050505061113b565b6040518080602001828103825283818151815260200191508051906020019080838360005b83811015610332578082015181840152602081019050610317565b50505050905090810190601f16801561035f5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561037957600080fd5b50610398600480360381019080803590602001909291905050506112cb565b005b3480156103a657600080fd5b506103af6112ce565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156103ef5780820151818401526020810190506103d4565b50505050905090810190601f16801561041c5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561043657600080fd5b5061043f61136c565b6040518082815260200191505060405180910390f35b34801561046157600080fd5b5061046a611372565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156104b857600080fd5b506104c1611398565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561050f57600080fd5b50610544600480360381019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506113bd565b60405180806020018060200180602001858152602001848103845288818151815260200191508051906020019080838360005b83811015610592578082015181840152602081019050610577565b50505050905090810190601f1680156105bf5780820380516001836020036101000a031916815260200191505b50848103835287818151815260200191508051906020019080838360005b838110156105f85780820151818401526020810190506105dd565b50505050905090810190601f1680156106255780820380516001836020036101000a031916815260200191505b50848103825286818151815260200191508051906020019080838360005b8381101561065e578082015181840152602081019050610643565b50505050905090810190601f16801561068b5780820380516001836020036101000a031916815260200191505b5097505050505050505060405180910390f35b3480156106aa57600080fd5b506106df600480360381019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506115e7565b005b3480156106ed57600080fd5b506106f661164e565b6040518080602001828103825283818151815260200191508051906020019060200280838360005b8381101561073957808201518184015260208101905061071e565b505050509050019250505060405180910390f35b34801561075957600080fd5b506107b4600480360381019080803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290505050611787565b6040518082815260200191505060405180910390f35b3480156107d657600080fd5b50610835600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291905050506118c1565b005b34801561084357600080fd5b50610878600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611e70565b005b34801561088657600080fd5b5061090d600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190820180359060200191909192939192939080359060200190820180359060200191909192939192939080359060200190820180359060200191909192939192939080359060200190929190505050611ed7565b005b6060806109506040805190810160405280600d81526020017f426f6b65724e6f64654461746100000000000000000000000000000000000000815250610a9b565b73ffffffffffffffffffffffffffffffffffffffff166306a49fce6040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401600060405180830381600087803b1580156109b357600080fd5b505af11580156109c7573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f8201168201806040525060408110156109f157600080fd5b810190808051640100000000811115610a0957600080fd5b82810190506020810184811115610a1f57600080fd5b8151856020820283011164010000000082111715610a3c57600080fd5b50509291906020018051640100000000811115610a5857600080fd5b82810190506020810184811115610a6e57600080fd5b8151856020820283011164010000000082111715610a8b57600080fd5b5050929190505050915091509091565b600080600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663fca1f3c1846040518263ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018080602001828103825283818151815260200191508051906020019080838360005b83811015610b48578082015181840152602081019050610b2d565b50505050905090810190601f168015610b755780820380516001836020036101000a031916815260200191505b5092505050600060405180830381600087803b158015610b9457600080fd5b505af1158015610ba8573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f8201168201806040525060a0811015610bd257600080fd5b81019080805190602001909291908051640100000000811115610bf457600080fd5b82810190506020810184811115610c0a57600080fd5b8151856001820283011164010000000082111715610c2757600080fd5b5050929190602001805190602001909291908051640100000000811115610c4d57600080fd5b82810190506020810184811115610c6357600080fd5b8151856001820283011164010000000082111715610c8057600080fd5b505092919060200180519060200190929190505050505092505050600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614151515610d40576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600a8152602001807f616464722069732030210000000000000000000000000000000000000000000081525060200191505060405180910390fd5b80915050919050565b6000606080600080610d8f6040805190810160405280600d81526020017f426f6b65724e6f64654461746100000000000000000000000000000000000000815250610a9b565b94508473ffffffffffffffffffffffffffffffffffffffff1663dc1e30da876040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001915050600060405180830381600087803b158015610e2c57600080fd5b505af1158015610e40573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f820116820180604052506080811015610e6a57600080fd5b810190808051640100000000811115610e8257600080fd5b82810190506020810184811115610e9857600080fd5b8151856020820283011164010000000082111715610eb557600080fd5b50509291906020018051640100000000811115610ed157600080fd5b82810190506020810184811115610ee757600080fd5b8151856020820283011164010000000082111715610f0457600080fd5b505092919060200180519060200190929190805190602001909291905050509350935093509350610f368685856120e4565b8473ffffffffffffffffffffffffffffffffffffffff1663c90f008a8760006040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050600060405180830381600087803b158015610fda57600080fd5b505af1158015610fee573d6000803e3d6000fd5b505050506110306040805190810160405280600881526020017f426f6b65724c6f67000000000000000000000000000000000000000000000000815250610a9b565b73ffffffffffffffffffffffffffffffffffffffff1663e7f218f28760006001600281111561105b57fe5b856040518563ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001838152602001828152602001945050505050600060405180830381600087803b15801561111b57600080fd5b505af115801561112f573d6000803e3d6000fd5b50505050505050505050565b6060600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16634a189f35836040518263ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018080602001828103825283818151815260200191508051906020019080838360005b838110156111e75780820151818401526020810190506111cc565b50505050905090810190601f1680156112145780820380516001836020036101000a031916815260200191505b5092505050600060405180830381600087803b15801561123357600080fd5b505af1158015611247573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f82011682018060405250602081101561127157600080fd5b81019080805164010000000081111561128957600080fd5b8281019050602081018481111561129f57600080fd5b81518560018202830111640100000000821117156112bc57600080fd5b50509291905050509050919050565b50565b60018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156113645780601f1061133957610100808354040283529160200191611364565b820191906000526020600020905b81548152906001019060200180831161134757829003601f168201915b505050505081565b60025481565b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b606080606060006114026040805190810160405280600d81526020017f426f6b65724e6f64654461746100000000000000000000000000000000000000815250610a9b565b73ffffffffffffffffffffffffffffffffffffffff1663b7adcf1d866040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001915050600060405180830381600087803b15801561149c57600080fd5b505af11580156114b0573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f8201168201806040525060808110156114da57600080fd5b8101908080516401000000008111156114f257600080fd5b8281019050602081018481111561150857600080fd5b815185600182028301116401000000008211171561152557600080fd5b5050929190602001805164010000000081111561154157600080fd5b8281019050602081018481111561155757600080fd5b815185600182028301116401000000008211171561157457600080fd5b5050929190602001805164010000000081111561159057600080fd5b828101905060208101848111156115a657600080fd5b81518560018202830111640100000000821117156115c357600080fd5b50509291906020018051906020019092919050505093509350935093509193509193565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561164257600080fd5b61164b81612324565b50565b606061168e6040805190810160405280600d81526020017f426f6b65724e6f64654461746100000000000000000000000000000000000000815250610a9b565b73ffffffffffffffffffffffffffffffffffffffff1663d38ff7a16040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401600060405180830381600087803b1580156116f157600080fd5b505af1158015611705573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f82011682018060405250602081101561172f57600080fd5b81019080805164010000000081111561174757600080fd5b8281019050602081018481111561175d57600080fd5b815185602082028301116401000000008211171561177a57600080fd5b5050929190505050905090565b6000600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16633230b078836040518263ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018080602001828103825283818151815260200191508051906020019080838360005b83811015611833578082015181840152602081019050611818565b50505050905090810190601f1680156118605780820380516001836020036101000a031916815260200191505b5092505050602060405180830381600087803b15801561187f57600080fd5b505af1158015611893573d6000803e3d6000fd5b505050506040513d60208110156118a957600080fd5b81019080805190602001909291905050509050919050565b60006119016040805190810160405280600d81526020017f426f6b65724e6f64654461746100000000000000000000000000000000000000815250610a9b565b90508073ffffffffffffffffffffffffffffffffffffffff166317132a0f846040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001915050602060405180830381600087803b15801561199e57600080fd5b505af11580156119b2573d6000803e3d6000fd5b505050506040513d60208110156119c857600080fd5b81019080805190602001909291905050501515611a4d576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260188152602001807f6164647243616e646964617465206e6f7420666f756e6421000000000000000081525060200191505060405180910390fd5b600082111515611ac5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600c8152602001807f746f6b656e73203c3d203021000000000000000000000000000000000000000081525060200191505060405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff1663424b0baf85846040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050600060405180830381600087803b158015611b6857600080fd5b505af1158015611b7c573d6000803e3d6000fd5b505050508073ffffffffffffffffffffffffffffffffffffffff16630c9f51088585856040518463ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019350505050600060405180830381600087803b158015611c5757600080fd5b505af1158015611c6b573d6000803e3d6000fd5b505050508073ffffffffffffffffffffffffffffffffffffffff16633e26a5a984846040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050600060405180830381600087803b158015611d1257600080fd5b505af1158015611d26573d6000803e3d6000fd5b50505050611d686040805190810160405280600881526020017f426f6b65724c6f67000000000000000000000000000000000000000000000000815250610a9b565b73ffffffffffffffffffffffffffffffffffffffff1663e7f218f2858560006002811115611d9257fe5b866040518563ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001838152602001828152602001945050505050600060405180830381600087803b158015611e5257600080fd5b505af1158015611e66573d6000803e3d6000fd5b5050505050505050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515611ecb57600080fd5b611ed481612481565b50565b600073ffffffffffffffffffffffffffffffffffffffff168873ffffffffffffffffffffffffffffffffffffffff1614151515611f7c576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260128152602001807f6164647243616e6469646174652069732030000000000000000000000000000081525060200191505060405180910390fd5b611fba6040805190810160405280600d81526020017f426f6b65724e6f64654461746100000000000000000000000000000000000000815250610a9b565b73ffffffffffffffffffffffffffffffffffffffff1663d349273b89898989898989896040518963ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200180602001806020018060200185815260200184810384528b8b82818152602001925080828437820191505084810383528989828181526020019250808284378201915050848103825287878281815260200192508082843782019150509b505050505050505050505050600060405180830381600087803b1580156120c257600080fd5b505af11580156120d6573d6000803e3d6000fd5b505050505050505050505050565b6000806000806121286040805190810160405280600d81526020017f426f6b65724e6f64654461746100000000000000000000000000000000000000815250610a9b565b9350600092505b855183101561231b57858381518110151561214657fe5b906020019060200201519150848381518110151561216057fe5b9060200190602002015190508373ffffffffffffffffffffffffffffffffffffffff1663a7a6586a88846040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200192505050600060405180830381600087803b15801561223b57600080fd5b505af115801561224f573d6000803e3d6000fd5b505050508373ffffffffffffffffffffffffffffffffffffffff1663e13e6ffe83836040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050600060405180830381600087803b1580156122f657600080fd5b505af115801561230a573d6000803e3d6000fd5b50505050828060010193505061212f565b50505050505050565b60008190508073ffffffffffffffffffffffffffffffffffffffff1663519c28826040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b15801561238d57600080fd5b505af11580156123a1573d6000803e3d6000fd5b505050506040513d60208110156123b757600080fd5b8101908080519060200190929190505050151561243c576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260118152602001807f6e6f7420426f6b65724d616e616765722100000000000000000000000000000081525060200191505060405180910390fd5b80600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614151515612526576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260168152602001807f6e65774f776e657220616464726573732069732030210000000000000000000081525060200191505060405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550505600a165627a7a723058204e041c00e4d474ae07f0d1fd0d5aa5b79d35c3f527397d87dde973bb919c1a430029`

// DeployBokerNode deploys a new Ethereum contract, binding an instance of BokerNode to it.
func DeployBokerNode(auth *bind.TransactOpts, backend bind.ContractBackend, addrManager common.Address) (common.Address, *types.Transaction, *BokerNode, error) {
	parsed, err := abi.JSON(strings.NewReader(BokerNodeABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BokerNodeBin), backend, addrManager)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BokerNode{BokerNodeCaller: BokerNodeCaller{contract: contract}, BokerNodeTransactor: BokerNodeTransactor{contract: contract}}, nil
}

// BokerNode is an auto generated Go binding around an Ethereum contract.
type BokerNode struct {
	BokerNodeCaller     // Read-only binding to the contract
	BokerNodeTransactor // Write-only binding to the contract
}

// BokerNodeCaller is an auto generated read-only Go binding around an Ethereum contract.
type BokerNodeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BokerNodeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BokerNodeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BokerNodeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BokerNodeSession struct {
	Contract     *BokerNode        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BokerNodeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BokerNodeCallerSession struct {
	Contract *BokerNodeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// BokerNodeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BokerNodeTransactorSession struct {
	Contract     *BokerNodeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// BokerNodeRaw is an auto generated low-level Go binding around an Ethereum contract.
type BokerNodeRaw struct {
	Contract *BokerNode // Generic contract binding to access the raw methods on
}

// BokerNodeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BokerNodeCallerRaw struct {
	Contract *BokerNodeCaller // Generic read-only contract binding to access the raw methods on
}

// BokerNodeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BokerNodeTransactorRaw struct {
	Contract *BokerNodeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBokerNode creates a new instance of BokerNode, bound to a specific deployed contract.
func NewBokerNode(address common.Address, backend bind.ContractBackend) (*BokerNode, error) {
	contract, err := bindBokerNode(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BokerNode{BokerNodeCaller: BokerNodeCaller{contract: contract}, BokerNodeTransactor: BokerNodeTransactor{contract: contract}}, nil
}

// NewBokerNodeCaller creates a new read-only instance of BokerNode, bound to a specific deployed contract.
func NewBokerNodeCaller(address common.Address, caller bind.ContractCaller) (*BokerNodeCaller, error) {
	contract, err := bindBokerNode(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &BokerNodeCaller{contract: contract}, nil
}

// NewBokerNodeTransactor creates a new write-only instance of BokerNode, bound to a specific deployed contract.
func NewBokerNodeTransactor(address common.Address, transactor bind.ContractTransactor) (*BokerNodeTransactor, error) {
	contract, err := bindBokerNode(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &BokerNodeTransactor{contract: contract}, nil
}

// bindBokerNode binds a generic wrapper to an already deployed contract.
func bindBokerNode(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BokerNodeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BokerNode *BokerNodeRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BokerNode.Contract.BokerNodeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BokerNode *BokerNodeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BokerNode.Contract.BokerNodeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BokerNode *BokerNodeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BokerNode.Contract.BokerNodeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BokerNode *BokerNodeCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BokerNode.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BokerNode *BokerNodeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BokerNode.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BokerNode *BokerNodeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BokerNode.Contract.contract.Transact(opts, method, params...)
}

// BokerManager is a free data retrieval call binding the contract method 0x66ebc1c6.
//
// Solidity: function bokerManager() constant returns(address)
func (_BokerNode *BokerNodeCaller) BokerManager(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BokerNode.contract.Call(opts, out, "bokerManager")
	return *ret0, err
}

// BokerManager is a free data retrieval call binding the contract method 0x66ebc1c6.
//
// Solidity: function bokerManager() constant returns(address)
func (_BokerNode *BokerNodeSession) BokerManager() (common.Address, error) {
	return _BokerNode.Contract.BokerManager(&_BokerNode.CallOpts)
}

// BokerManager is a free data retrieval call binding the contract method 0x66ebc1c6.
//
// Solidity: function bokerManager() constant returns(address)
func (_BokerNode *BokerNodeCallerSession) BokerManager() (common.Address, error) {
	return _BokerNode.Contract.BokerManager(&_BokerNode.CallOpts)
}

// ContractAddress is a free data retrieval call binding the contract method 0x1eb726af.
//
// Solidity: function contractAddress(cName string) constant returns(address)
func (_BokerNode *BokerNodeCaller) ContractAddress(opts *bind.CallOpts, cName string) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BokerNode.contract.Call(opts, out, "contractAddress", cName)
	return *ret0, err
}

// ContractAddress is a free data retrieval call binding the contract method 0x1eb726af.
//
// Solidity: function contractAddress(cName string) constant returns(address)
func (_BokerNode *BokerNodeSession) ContractAddress(cName string) (common.Address, error) {
	return _BokerNode.Contract.ContractAddress(&_BokerNode.CallOpts, cName)
}

// ContractAddress is a free data retrieval call binding the contract method 0x1eb726af.
//
// Solidity: function contractAddress(cName string) constant returns(address)
func (_BokerNode *BokerNodeCallerSession) ContractAddress(cName string) (common.Address, error) {
	return _BokerNode.Contract.ContractAddress(&_BokerNode.CallOpts, cName)
}

// CreateTime is a free data retrieval call binding the contract method 0x61dcd7ab.
//
// Solidity: function createTime() constant returns(uint256)
func (_BokerNode *BokerNodeCaller) CreateTime(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BokerNode.contract.Call(opts, out, "createTime")
	return *ret0, err
}

// CreateTime is a free data retrieval call binding the contract method 0x61dcd7ab.
//
// Solidity: function createTime() constant returns(uint256)
func (_BokerNode *BokerNodeSession) CreateTime() (*big.Int, error) {
	return _BokerNode.Contract.CreateTime(&_BokerNode.CallOpts)
}

// CreateTime is a free data retrieval call binding the contract method 0x61dcd7ab.
//
// Solidity: function createTime() constant returns(uint256)
func (_BokerNode *BokerNodeCallerSession) CreateTime() (*big.Int, error) {
	return _BokerNode.Contract.CreateTime(&_BokerNode.CallOpts)
}

// GetBlacks is a free data retrieval call binding the contract method 0xd38ff7a1.
//
// Solidity: function getBlacks() constant returns(addresses address[])
func (_BokerNode *BokerNodeCaller) GetBlacks(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _BokerNode.contract.Call(opts, out, "getBlacks")
	return *ret0, err
}

// GetBlacks is a free data retrieval call binding the contract method 0xd38ff7a1.
//
// Solidity: function getBlacks() constant returns(addresses address[])
func (_BokerNode *BokerNodeSession) GetBlacks() ([]common.Address, error) {
	return _BokerNode.Contract.GetBlacks(&_BokerNode.CallOpts)
}

// GetBlacks is a free data retrieval call binding the contract method 0xd38ff7a1.
//
// Solidity: function getBlacks() constant returns(addresses address[])
func (_BokerNode *BokerNodeCallerSession) GetBlacks() ([]common.Address, error) {
	return _BokerNode.Contract.GetBlacks(&_BokerNode.CallOpts)
}

// GetCandidate is a free data retrieval call binding the contract method 0xb7adcf1d.
//
// Solidity: function getCandidate(addrCandidate address) constant returns(description string, team string, name string, tickets uint256)
func (_BokerNode *BokerNodeCaller) GetCandidate(opts *bind.CallOpts, addrCandidate common.Address) (struct {
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
	err := _BokerNode.contract.Call(opts, out, "getCandidate", addrCandidate)
	return *ret, err
}

// GetCandidate is a free data retrieval call binding the contract method 0xb7adcf1d.
//
// Solidity: function getCandidate(addrCandidate address) constant returns(description string, team string, name string, tickets uint256)
func (_BokerNode *BokerNodeSession) GetCandidate(addrCandidate common.Address) (struct {
	Description string
	Team        string
	Name        string
	Tickets     *big.Int
}, error) {
	return _BokerNode.Contract.GetCandidate(&_BokerNode.CallOpts, addrCandidate)
}

// GetCandidate is a free data retrieval call binding the contract method 0xb7adcf1d.
//
// Solidity: function getCandidate(addrCandidate address) constant returns(description string, team string, name string, tickets uint256)
func (_BokerNode *BokerNodeCallerSession) GetCandidate(addrCandidate common.Address) (struct {
	Description string
	Team        string
	Name        string
	Tickets     *big.Int
}, error) {
	return _BokerNode.Contract.GetCandidate(&_BokerNode.CallOpts, addrCandidate)
}

// GetCandidates is a free data retrieval call binding the contract method 0x06a49fce.
//
// Solidity: function getCandidates() constant returns(addresses address[], tickets uint256[])
func (_BokerNode *BokerNodeCaller) GetCandidates(opts *bind.CallOpts) (struct {
	Addresses []common.Address
	Tickets   []*big.Int
}, error) {
	ret := new(struct {
		Addresses []common.Address
		Tickets   []*big.Int
	})
	out := ret
	err := _BokerNode.contract.Call(opts, out, "getCandidates")
	return *ret, err
}

// GetCandidates is a free data retrieval call binding the contract method 0x06a49fce.
//
// Solidity: function getCandidates() constant returns(addresses address[], tickets uint256[])
func (_BokerNode *BokerNodeSession) GetCandidates() (struct {
	Addresses []common.Address
	Tickets   []*big.Int
}, error) {
	return _BokerNode.Contract.GetCandidates(&_BokerNode.CallOpts)
}

// GetCandidates is a free data retrieval call binding the contract method 0x06a49fce.
//
// Solidity: function getCandidates() constant returns(addresses address[], tickets uint256[])
func (_BokerNode *BokerNodeCallerSession) GetCandidates() (struct {
	Addresses []common.Address
	Tickets   []*big.Int
}, error) {
	return _BokerNode.Contract.GetCandidates(&_BokerNode.CallOpts)
}

// GlobalConfigInt is a free data retrieval call binding the contract method 0xd43c8021.
//
// Solidity: function globalConfigInt(key string) constant returns(uint256)
func (_BokerNode *BokerNodeCaller) GlobalConfigInt(opts *bind.CallOpts, key string) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BokerNode.contract.Call(opts, out, "globalConfigInt", key)
	return *ret0, err
}

// GlobalConfigInt is a free data retrieval call binding the contract method 0xd43c8021.
//
// Solidity: function globalConfigInt(key string) constant returns(uint256)
func (_BokerNode *BokerNodeSession) GlobalConfigInt(key string) (*big.Int, error) {
	return _BokerNode.Contract.GlobalConfigInt(&_BokerNode.CallOpts, key)
}

// GlobalConfigInt is a free data retrieval call binding the contract method 0xd43c8021.
//
// Solidity: function globalConfigInt(key string) constant returns(uint256)
func (_BokerNode *BokerNodeCallerSession) GlobalConfigInt(key string) (*big.Int, error) {
	return _BokerNode.Contract.GlobalConfigInt(&_BokerNode.CallOpts, key)
}

// GlobalConfigString is a free data retrieval call binding the contract method 0x378298bc.
//
// Solidity: function globalConfigString(key string) constant returns(string)
func (_BokerNode *BokerNodeCaller) GlobalConfigString(opts *bind.CallOpts, key string) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _BokerNode.contract.Call(opts, out, "globalConfigString", key)
	return *ret0, err
}

// GlobalConfigString is a free data retrieval call binding the contract method 0x378298bc.
//
// Solidity: function globalConfigString(key string) constant returns(string)
func (_BokerNode *BokerNodeSession) GlobalConfigString(key string) (string, error) {
	return _BokerNode.Contract.GlobalConfigString(&_BokerNode.CallOpts, key)
}

// GlobalConfigString is a free data retrieval call binding the contract method 0x378298bc.
//
// Solidity: function globalConfigString(key string) constant returns(string)
func (_BokerNode *BokerNodeCallerSession) GlobalConfigString(key string) (string, error) {
	return _BokerNode.Contract.GlobalConfigString(&_BokerNode.CallOpts, key)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_BokerNode *BokerNodeCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BokerNode.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_BokerNode *BokerNodeSession) Owner() (common.Address, error) {
	return _BokerNode.Contract.Owner(&_BokerNode.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_BokerNode *BokerNodeCallerSession) Owner() (common.Address, error) {
	return _BokerNode.Contract.Owner(&_BokerNode.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() constant returns(string)
func (_BokerNode *BokerNodeCaller) Version(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _BokerNode.contract.Call(opts, out, "version")
	return *ret0, err
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() constant returns(string)
func (_BokerNode *BokerNodeSession) Version() (string, error) {
	return _BokerNode.Contract.Version(&_BokerNode.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() constant returns(string)
func (_BokerNode *BokerNodeCallerSession) Version() (string, error) {
	return _BokerNode.Contract.Version(&_BokerNode.CallOpts)
}

// CancelAllVotes is a paid mutator transaction binding the contract method 0x281c231b.
//
// Solidity: function cancelAllVotes(addrVoter address) returns()
func (_BokerNode *BokerNodeTransactor) CancelAllVotes(opts *bind.TransactOpts, addrVoter common.Address) (*types.Transaction, error) {
	return _BokerNode.contract.Transact(opts, "cancelAllVotes", addrVoter)
}

// CancelAllVotes is a paid mutator transaction binding the contract method 0x281c231b.
//
// Solidity: function cancelAllVotes(addrVoter address) returns()
func (_BokerNode *BokerNodeSession) CancelAllVotes(addrVoter common.Address) (*types.Transaction, error) {
	return _BokerNode.Contract.CancelAllVotes(&_BokerNode.TransactOpts, addrVoter)
}

// CancelAllVotes is a paid mutator transaction binding the contract method 0x281c231b.
//
// Solidity: function cancelAllVotes(addrVoter address) returns()
func (_BokerNode *BokerNodeTransactorSession) CancelAllVotes(addrVoter common.Address) (*types.Transaction, error) {
	return _BokerNode.Contract.CancelAllVotes(&_BokerNode.TransactOpts, addrVoter)
}

// RegisterCandidate is a paid mutator transaction binding the contract method 0xf581978b.
//
// Solidity: function registerCandidate(addrCandidate address, description string, team string, name string, tickets uint256) returns()
func (_BokerNode *BokerNodeTransactor) RegisterCandidate(opts *bind.TransactOpts, addrCandidate common.Address, description string, team string, name string, tickets *big.Int) (*types.Transaction, error) {
	return _BokerNode.contract.Transact(opts, "registerCandidate", addrCandidate, description, team, name, tickets)
}

// RegisterCandidate is a paid mutator transaction binding the contract method 0xf581978b.
//
// Solidity: function registerCandidate(addrCandidate address, description string, team string, name string, tickets uint256) returns()
func (_BokerNode *BokerNodeSession) RegisterCandidate(addrCandidate common.Address, description string, team string, name string, tickets *big.Int) (*types.Transaction, error) {
	return _BokerNode.Contract.RegisterCandidate(&_BokerNode.TransactOpts, addrCandidate, description, team, name, tickets)
}

// RegisterCandidate is a paid mutator transaction binding the contract method 0xf581978b.
//
// Solidity: function registerCandidate(addrCandidate address, description string, team string, name string, tickets uint256) returns()
func (_BokerNode *BokerNodeTransactorSession) RegisterCandidate(addrCandidate common.Address, description string, team string, name string, tickets *big.Int) (*types.Transaction, error) {
	return _BokerNode.Contract.RegisterCandidate(&_BokerNode.TransactOpts, addrCandidate, description, team, name, tickets)
}

// RotateVote is a paid mutator transaction binding the contract method 0x43bd7d97.
//
// Solidity: function rotateVote(nowTimer uint256) returns()
func (_BokerNode *BokerNodeTransactor) RotateVote(opts *bind.TransactOpts, nowTimer *big.Int) (*types.Transaction, error) {
	return _BokerNode.contract.Transact(opts, "rotateVote", nowTimer)
}

// RotateVote is a paid mutator transaction binding the contract method 0x43bd7d97.
//
// Solidity: function rotateVote(nowTimer uint256) returns()
func (_BokerNode *BokerNodeSession) RotateVote(nowTimer *big.Int) (*types.Transaction, error) {
	return _BokerNode.Contract.RotateVote(&_BokerNode.TransactOpts, nowTimer)
}

// RotateVote is a paid mutator transaction binding the contract method 0x43bd7d97.
//
// Solidity: function rotateVote(nowTimer uint256) returns()
func (_BokerNode *BokerNodeTransactorSession) RotateVote(nowTimer *big.Int) (*types.Transaction, error) {
	return _BokerNode.Contract.RotateVote(&_BokerNode.TransactOpts, nowTimer)
}

// SetManager is a paid mutator transaction binding the contract method 0xd0ebdbe7.
//
// Solidity: function setManager(addrManager address) returns()
func (_BokerNode *BokerNodeTransactor) SetManager(opts *bind.TransactOpts, addrManager common.Address) (*types.Transaction, error) {
	return _BokerNode.contract.Transact(opts, "setManager", addrManager)
}

// SetManager is a paid mutator transaction binding the contract method 0xd0ebdbe7.
//
// Solidity: function setManager(addrManager address) returns()
func (_BokerNode *BokerNodeSession) SetManager(addrManager common.Address) (*types.Transaction, error) {
	return _BokerNode.Contract.SetManager(&_BokerNode.TransactOpts, addrManager)
}

// SetManager is a paid mutator transaction binding the contract method 0xd0ebdbe7.
//
// Solidity: function setManager(addrManager address) returns()
func (_BokerNode *BokerNodeTransactorSession) SetManager(addrManager common.Address) (*types.Transaction, error) {
	return _BokerNode.Contract.SetManager(&_BokerNode.TransactOpts, addrManager)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_BokerNode *BokerNodeTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _BokerNode.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_BokerNode *BokerNodeSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _BokerNode.Contract.TransferOwnership(&_BokerNode.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_BokerNode *BokerNodeTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _BokerNode.Contract.TransferOwnership(&_BokerNode.TransactOpts, _newOwner)
}

// Vote is a paid mutator transaction binding the contract method 0xefa5d431.
//
// Solidity: function vote(addrVoter address, addrCandidate address, tokens uint256) returns()
func (_BokerNode *BokerNodeTransactor) Vote(opts *bind.TransactOpts, addrVoter common.Address, addrCandidate common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _BokerNode.contract.Transact(opts, "vote", addrVoter, addrCandidate, tokens)
}

// Vote is a paid mutator transaction binding the contract method 0xefa5d431.
//
// Solidity: function vote(addrVoter address, addrCandidate address, tokens uint256) returns()
func (_BokerNode *BokerNodeSession) Vote(addrVoter common.Address, addrCandidate common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _BokerNode.Contract.Vote(&_BokerNode.TransactOpts, addrVoter, addrCandidate, tokens)
}

// Vote is a paid mutator transaction binding the contract method 0xefa5d431.
//
// Solidity: function vote(addrVoter address, addrCandidate address, tokens uint256) returns()
func (_BokerNode *BokerNodeTransactorSession) Vote(addrVoter common.Address, addrCandidate common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _BokerNode.Contract.Vote(&_BokerNode.TransactOpts, addrVoter, addrCandidate, tokens)
}
