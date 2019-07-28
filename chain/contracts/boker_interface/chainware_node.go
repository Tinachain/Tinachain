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
const BokerNodeABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"getCandidates\",\"outputs\":[{\"name\":\"addresses\",\"type\":\"address[]\"},{\"name\":\"tickets\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"cName\",\"type\":\"string\"}],\"name\":\"contractAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addrVoter\",\"type\":\"address\"}],\"name\":\"cancelAllVotes\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"key\",\"type\":\"string\"}],\"name\":\"globalConfigString\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"nowTimer\",\"type\":\"uint256\"}],\"name\":\"rotateVote\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"createTime\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"nowTimer\",\"type\":\"uint256\"}],\"name\":\"tickTimeout\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"bokerManager\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addrCandidate\",\"type\":\"address\"}],\"name\":\"getCandidate\",\"outputs\":[{\"name\":\"description\",\"type\":\"string\"},{\"name\":\"team\",\"type\":\"string\"},{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"tickets\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addrManager\",\"type\":\"address\"}],\"name\":\"setManager\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getBlacks\",\"outputs\":[{\"name\":\"addresses\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"key\",\"type\":\"string\"}],\"name\":\"globalConfigInt\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addrVoter\",\"type\":\"address\"},{\"name\":\"addrCandidate\",\"type\":\"address\"},{\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"vote\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addrCandidate\",\"type\":\"address\"},{\"name\":\"description\",\"type\":\"string\"},{\"name\":\"team\",\"type\":\"string\"},{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"tickets\",\"type\":\"uint256\"}],\"name\":\"registerCandidate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"addrManager\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// BokerNodeBin is the compiled bytecode used for deploying new contracts.
const BokerNodeBin = `60806040526040805190810160405280600581526020017f312e302e3100000000000000000000000000000000000000000000000000000081525060019080519060200190620000519291906200024b565b50426002553480156200006357600080fd5b5060405160208062002a0e8339810180604052810190808051906020019092919050505080336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550620000e281620000ea640100000000026401000000009004565b5050620002fa565b60008190508073ffffffffffffffffffffffffffffffffffffffff1663519c28826040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b1580156200015457600080fd5b505af115801562000169573d6000803e3d6000fd5b505050506040513d60208110156200018057600080fd5b8101908080519060200190929190505050151562000206576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260118152602001807f6e6f7420426f6b65724d616e616765722100000000000000000000000000000081525060200191505060405180910390fd5b80600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106200028e57805160ff1916838001178555620002bf565b82800160010185558215620002bf579182015b82811115620002be578251825591602001919060010190620002a1565b5b509050620002ce9190620002d2565b5090565b620002f791905b80821115620002f3576000816000905550600101620002d9565b5090565b90565b612704806200030a6000396000f3006080604052600436106100f1576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff16806306a49fce146100f65780631eb726af146101aa578063281c231b14610253578063378298bc1461029657806343bd7d971461037857806354fd4d50146103a557806361dcd7ab14610435578063662711971461046057806366ebc1c61461048d5780638da5cb5b146104e4578063b7adcf1d1461053b578063d0ebdbe7146106d6578063d38ff7a114610719578063d43c802114610785578063efa5d43114610802578063f2fde38b1461086f578063f581978b146108b2575b600080fd5b34801561010257600080fd5b5061010b610947565b604051808060200180602001838103835285818151815260200191508051906020019060200280838360005b83811015610152578082015181840152602081019050610137565b50505050905001838103825284818151815260200191508051906020019060200280838360005b83811015610194578082015181840152602081019050610179565b5050505090500194505050505060405180910390f35b3480156101b657600080fd5b50610211600480360381019080803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290505050610ad3565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561025f57600080fd5b50610294600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610d81565b005b3480156102a257600080fd5b506102fd600480360381019080803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290505050611173565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561033d578082015181840152602081019050610322565b50505050905090810190601f16801561036a5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561038457600080fd5b506103a360048036038101908080359060200190929190505050611303565b005b3480156103b157600080fd5b506103ba611306565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156103fa5780820151818401526020810190506103df565b50505050905090810190601f1680156104275780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561044157600080fd5b5061044a6113a4565b6040518082815260200191505060405180910390f35b34801561046c57600080fd5b5061048b600480360381019080803590602001909291905050506113aa565b005b34801561049957600080fd5b506104a2611466565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156104f057600080fd5b506104f961148c565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561054757600080fd5b5061057c600480360381019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506114b1565b60405180806020018060200180602001858152602001848103845288818151815260200191508051906020019080838360005b838110156105ca5780820151818401526020810190506105af565b50505050905090810190601f1680156105f75780820380516001836020036101000a031916815260200191505b50848103835287818151815260200191508051906020019080838360005b83811015610630578082015181840152602081019050610615565b50505050905090810190601f16801561065d5780820380516001836020036101000a031916815260200191505b50848103825286818151815260200191508051906020019080838360005b8381101561069657808201518184015260208101905061067b565b50505050905090810190601f1680156106c35780820380516001836020036101000a031916815260200191505b5097505050505050505060405180910390f35b3480156106e257600080fd5b50610717600480360381019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506116db565b005b34801561072557600080fd5b5061072e611742565b6040518080602001828103825283818151815260200191508051906020019060200280838360005b83811015610771578082015181840152602081019050610756565b505050509050019250505060405180910390f35b34801561079157600080fd5b506107ec600480360381019080803590602001908201803590602001908080601f016020809104026020016040519081016040528093929190818152602001838380828437820191505050505050919291929050505061187b565b6040518082815260200191505060405180910390f35b34801561080e57600080fd5b5061086d600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291905050506119b5565b005b34801561087b57600080fd5b506108b0600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611f64565b005b3480156108be57600080fd5b50610945600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190820180359060200191909192939192939080359060200190820180359060200191909192939192939080359060200190820180359060200191909192939192939080359060200190929190505050611fcb565b005b6060806109886040805190810160405280600d81526020017f426f6b65724e6f64654461746100000000000000000000000000000000000000815250610ad3565b73ffffffffffffffffffffffffffffffffffffffff166306a49fce6040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401600060405180830381600087803b1580156109eb57600080fd5b505af11580156109ff573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f820116820180604052506040811015610a2957600080fd5b810190808051640100000000811115610a4157600080fd5b82810190506020810184811115610a5757600080fd5b8151856020820283011164010000000082111715610a7457600080fd5b50509291906020018051640100000000811115610a9057600080fd5b82810190506020810184811115610aa657600080fd5b8151856020820283011164010000000082111715610ac357600080fd5b5050929190505050915091509091565b600080600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663fca1f3c1846040518263ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018080602001828103825283818151815260200191508051906020019080838360005b83811015610b80578082015181840152602081019050610b65565b50505050905090810190601f168015610bad5780820380516001836020036101000a031916815260200191505b5092505050600060405180830381600087803b158015610bcc57600080fd5b505af1158015610be0573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f8201168201806040525060a0811015610c0a57600080fd5b81019080805190602001909291908051640100000000811115610c2c57600080fd5b82810190506020810184811115610c4257600080fd5b8151856001820283011164010000000082111715610c5f57600080fd5b5050929190602001805190602001909291908051640100000000811115610c8557600080fd5b82810190506020810184811115610c9b57600080fd5b8151856001820283011164010000000082111715610cb857600080fd5b505092919060200180519060200190929190505050505092505050600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614151515610d78576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600a8152602001807f616464722069732030210000000000000000000000000000000000000000000081525060200191505060405180910390fd5b80915050919050565b6000606080600080610dc76040805190810160405280600d81526020017f426f6b65724e6f64654461746100000000000000000000000000000000000000815250610ad3565b94508473ffffffffffffffffffffffffffffffffffffffff1663dc1e30da876040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001915050600060405180830381600087803b158015610e6457600080fd5b505af1158015610e78573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f820116820180604052506080811015610ea257600080fd5b810190808051640100000000811115610eba57600080fd5b82810190506020810184811115610ed057600080fd5b8151856020820283011164010000000082111715610eed57600080fd5b50509291906020018051640100000000811115610f0957600080fd5b82810190506020810184811115610f1f57600080fd5b8151856020820283011164010000000082111715610f3c57600080fd5b505092919060200180519060200190929190805190602001909291905050509350935093509350610f6e8685856121d8565b8473ffffffffffffffffffffffffffffffffffffffff1663c90f008a8760006040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050600060405180830381600087803b15801561101257600080fd5b505af1158015611026573d6000803e3d6000fd5b505050506110686040805190810160405280600881526020017f426f6b65724c6f67000000000000000000000000000000000000000000000000815250610ad3565b73ffffffffffffffffffffffffffffffffffffffff1663e7f218f28760006001600281111561109357fe5b856040518563ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001838152602001828152602001945050505050600060405180830381600087803b15801561115357600080fd5b505af1158015611167573d6000803e3d6000fd5b50505050505050505050565b6060600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16634a189f35836040518263ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018080602001828103825283818151815260200191508051906020019080838360005b8381101561121f578082015181840152602081019050611204565b50505050905090810190601f16801561124c5780820380516001836020036101000a031916815260200191505b5092505050600060405180830381600087803b15801561126b57600080fd5b505af115801561127f573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f8201168201806040525060208110156112a957600080fd5b8101908080516401000000008111156112c157600080fd5b828101905060208101848111156112d757600080fd5b81518560018202830111640100000000821117156112f457600080fd5b50509291905050509050919050565b50565b60018054600181600116156101000203166002900480601f01602080910402602001604051908101604052809291908181526020018280546001816001161561010002031660029004801561139c5780601f106113715761010080835404028352916020019161139c565b820191906000526020600020905b81548152906001019060200180831161137f57829003601f168201915b505050505081565b60025481565b6113e86040805190810160405280600d81526020017f426f6b65724e6f64654461746100000000000000000000000000000000000000815250610ad3565b73ffffffffffffffffffffffffffffffffffffffff1663156bf9ba6040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401600060405180830381600087803b15801561144b57600080fd5b505af115801561145f573d6000803e3d6000fd5b5050505050565b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b606080606060006114f66040805190810160405280600d81526020017f426f6b65724e6f64654461746100000000000000000000000000000000000000815250610ad3565b73ffffffffffffffffffffffffffffffffffffffff1663b7adcf1d866040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001915050600060405180830381600087803b15801561159057600080fd5b505af11580156115a4573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f8201168201806040525060808110156115ce57600080fd5b8101908080516401000000008111156115e657600080fd5b828101905060208101848111156115fc57600080fd5b815185600182028301116401000000008211171561161957600080fd5b5050929190602001805164010000000081111561163557600080fd5b8281019050602081018481111561164b57600080fd5b815185600182028301116401000000008211171561166857600080fd5b5050929190602001805164010000000081111561168457600080fd5b8281019050602081018481111561169a57600080fd5b81518560018202830111640100000000821117156116b757600080fd5b50509291906020018051906020019092919050505093509350935093509193509193565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561173657600080fd5b61173f81612418565b50565b60606117826040805190810160405280600d81526020017f426f6b65724e6f64654461746100000000000000000000000000000000000000815250610ad3565b73ffffffffffffffffffffffffffffffffffffffff1663d38ff7a16040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401600060405180830381600087803b1580156117e557600080fd5b505af11580156117f9573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f82011682018060405250602081101561182357600080fd5b81019080805164010000000081111561183b57600080fd5b8281019050602081018481111561185157600080fd5b815185602082028301116401000000008211171561186e57600080fd5b5050929190505050905090565b6000600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16633230b078836040518263ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018080602001828103825283818151815260200191508051906020019080838360005b8381101561192757808201518184015260208101905061190c565b50505050905090810190601f1680156119545780820380516001836020036101000a031916815260200191505b5092505050602060405180830381600087803b15801561197357600080fd5b505af1158015611987573d6000803e3d6000fd5b505050506040513d602081101561199d57600080fd5b81019080805190602001909291905050509050919050565b60006119f56040805190810160405280600d81526020017f426f6b65724e6f64654461746100000000000000000000000000000000000000815250610ad3565b90508073ffffffffffffffffffffffffffffffffffffffff166317132a0f846040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001915050602060405180830381600087803b158015611a9257600080fd5b505af1158015611aa6573d6000803e3d6000fd5b505050506040513d6020811015611abc57600080fd5b81019080805190602001909291905050501515611b41576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260188152602001807f6164647243616e646964617465206e6f7420666f756e6421000000000000000081525060200191505060405180910390fd5b600082111515611bb9576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600c8152602001807f746f6b656e73203c3d203021000000000000000000000000000000000000000081525060200191505060405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff1663424b0baf85846040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050600060405180830381600087803b158015611c5c57600080fd5b505af1158015611c70573d6000803e3d6000fd5b505050508073ffffffffffffffffffffffffffffffffffffffff16630c9f51088585856040518463ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019350505050600060405180830381600087803b158015611d4b57600080fd5b505af1158015611d5f573d6000803e3d6000fd5b505050508073ffffffffffffffffffffffffffffffffffffffff16633e26a5a984846040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050600060405180830381600087803b158015611e0657600080fd5b505af1158015611e1a573d6000803e3d6000fd5b50505050611e5c6040805190810160405280600881526020017f426f6b65724c6f67000000000000000000000000000000000000000000000000815250610ad3565b73ffffffffffffffffffffffffffffffffffffffff1663e7f218f2858560006002811115611e8657fe5b866040518563ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001838152602001828152602001945050505050600060405180830381600087803b158015611f4657600080fd5b505af1158015611f5a573d6000803e3d6000fd5b5050505050505050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515611fbf57600080fd5b611fc881612575565b50565b600073ffffffffffffffffffffffffffffffffffffffff168873ffffffffffffffffffffffffffffffffffffffff1614151515612070576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260128152602001807f6164647243616e6469646174652069732030000000000000000000000000000081525060200191505060405180910390fd5b6120ae6040805190810160405280600d81526020017f426f6b65724e6f64654461746100000000000000000000000000000000000000815250610ad3565b73ffffffffffffffffffffffffffffffffffffffff1663d349273b89898989898989896040518963ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200180602001806020018060200185815260200184810384528b8b82818152602001925080828437820191505084810383528989828181526020019250808284378201915050848103825287878281815260200192508082843782019150509b505050505050505050505050600060405180830381600087803b1580156121b657600080fd5b505af11580156121ca573d6000803e3d6000fd5b505050505050505050505050565b60008060008061221c6040805190810160405280600d81526020017f426f6b65724e6f64654461746100000000000000000000000000000000000000815250610ad3565b9350600092505b855183101561240f57858381518110151561223a57fe5b906020019060200201519150848381518110151561225457fe5b9060200190602002015190508373ffffffffffffffffffffffffffffffffffffffff1663a7a6586a88846040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200192505050600060405180830381600087803b15801561232f57600080fd5b505af1158015612343573d6000803e3d6000fd5b505050508373ffffffffffffffffffffffffffffffffffffffff1663e13e6ffe83836040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050600060405180830381600087803b1580156123ea57600080fd5b505af11580156123fe573d6000803e3d6000fd5b505050508280600101935050612223565b50505050505050565b60008190508073ffffffffffffffffffffffffffffffffffffffff1663519c28826040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b15801561248157600080fd5b505af1158015612495573d6000803e3d6000fd5b505050506040513d60208110156124ab57600080fd5b81019080805190602001909291905050501515612530576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260118152602001807f6e6f7420426f6b65724d616e616765722100000000000000000000000000000081525060200191505060405180910390fd5b80600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415151561261a576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260168152602001807f6e65774f776e657220616464726573732069732030210000000000000000000081525060200191505060405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550505600a165627a7a7230582034e17c59320d0ed9b68985510ab3bda1258a01f212ed66b39c7e0451bac4cef00029`

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

// TickTimeout is a paid mutator transaction binding the contract method 0x66271197.
//
// Solidity: function tickTimeout(nowTimer uint256) returns()
func (_BokerNode *BokerNodeTransactor) TickTimeout(opts *bind.TransactOpts, nowTimer *big.Int) (*types.Transaction, error) {
	return _BokerNode.contract.Transact(opts, "tickTimeout", nowTimer)
}

// TickTimeout is a paid mutator transaction binding the contract method 0x66271197.
//
// Solidity: function tickTimeout(nowTimer uint256) returns()
func (_BokerNode *BokerNodeSession) TickTimeout(nowTimer *big.Int) (*types.Transaction, error) {
	return _BokerNode.Contract.TickTimeout(&_BokerNode.TransactOpts, nowTimer)
}

// TickTimeout is a paid mutator transaction binding the contract method 0x66271197.
//
// Solidity: function tickTimeout(nowTimer uint256) returns()
func (_BokerNode *BokerNodeTransactorSession) TickTimeout(nowTimer *big.Int) (*types.Transaction, error) {
	return _BokerNode.Contract.TickTimeout(&_BokerNode.TransactOpts, nowTimer)
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
