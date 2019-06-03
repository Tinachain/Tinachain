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

// BokerTokenPowerDataABI is the input ABI used to generate the binding from.
const BokerTokenPowerDataABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"setAssignCycleBegin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"cName\",\"type\":\"string\"}],\"name\":\"contractAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"key\",\"type\":\"string\"}],\"name\":\"globalConfigString\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"val\",\"type\":\"uint256\"}],\"name\":\"tokenAssignedTotalSet\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"createTime\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"bokerManager\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"dappType\",\"type\":\"uint256\"}],\"name\":\"bindDappConfigGet\",\"outputs\":[{\"name\":\"uppers\",\"type\":\"int256[]\"},{\"name\":\"powers\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"tokenAssignedTotal\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"dappType\",\"type\":\"uint256\"},{\"name\":\"index\",\"type\":\"uint256\"},{\"name\":\"upper\",\"type\":\"int256\"},{\"name\":\"power\",\"type\":\"uint256\"}],\"name\":\"bindDappConfigSet\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"dappAddr\",\"type\":\"address\"}],\"name\":\"bindDappGetPower\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"assignCycleBegin\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addrManager\",\"type\":\"address\"}],\"name\":\"setManager\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"key\",\"type\":\"string\"}],\"name\":\"globalConfigInt\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"addrManager\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// BokerTokenPowerDataBin is the compiled bytecode used for deploying new contracts.
const BokerTokenPowerDataBin = `60806040526040805190810160405280600581526020017f312e302e31000000000000000000000000000000000000000000000000000000815250600190805190602001906200005192919062000b20565b5042600255600060045560006005553480156200006d57600080fd5b5060405160208062002a0a8339810180604052810190808051906020019092919050505060008082336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550620000ef81620009bf640100000000026401000000009004565b504260048190555060019150600660008381526020019081526020016000209050818160000181905550806001016040805190810160405280600a81526020016103e8600102815250908060018154018082558091505090600182039060005260206000209060020201600090919290919091506000820151816000015560208201518160010155505050806001016040805190810160405280603281526020016103e8600302815250908060018154018082558091505090600182039060005260206000209060020201600090919290919091506000820151816000015560208201518160010155505050806001016040805190810160405280606481526020016103e86005028152509080600181540180825580915050906001820390600052602060002090600202016000909192909190915060008201518160000155602082015181600101555050508060010160408051908101604052806101f481526020016103e8600a028152509080600181540180825580915050906001820390600052602060002090600202016000909192909190915060008201518160000155602082015181600101555050508060010160408051908101604052806103e881526020016103e8600f0281525090806001815401808255809150509060018203906000526020600020906002020160009091929091909150600082015181600001556020820151816001015550505080600101604080519081016040528061138881526020016103e860140281525090806001815401808255809150509060018203906000526020600020906002020160009091929091909150600082015181600001556020820151816001015550505080600101604080519081016040528061271081526020016103e8601e028152509080600181540180825580915050906001820390600052602060002090600202016000909192909190915060008201518160000155602082015181600101555050508060010160408051908101604052807fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81526020016103e860280281525090806001815401808255809150509060018203906000526020600020906002020160009091929091909150600082015181600001556020820151816001015550505060029150600660008381526020019081526020016000209050818160000181905550806001016040805190810160405280600a81526020016103e8600102815250908060018154018082558091505090600182039060005260206000209060020201600090919290919091506000820151816000015560208201518160010155505050806001016040805190810160405280603281526020016103e8600302815250908060018154018082558091505090600182039060005260206000209060020201600090919290919091506000820151816000015560208201518160010155505050806001016040805190810160405280606481526020016103e860050281525090806001815401808255809150509060018203906000526020600020906002020160009091929091909150600082015181600001556020820151816001015550505080600101604080519081016040528060c881526020016103e8600a028152509080600181540180825580915050906001820390600052602060002090600202016000909192909190915060008201518160000155602082015181600101555050508060010160408051908101604052806101f481526020016103e86014028152509080600181540180825580915050906001820390600052602060002090600202016000909192909190915060008201518160000155602082015181600101555050508060010160408051908101604052806103e881526020016103e8601e028152509080600181540180825580915050906001820390600052602060002090600202016000909192909190915060008201518160000155602082015181600101555050508060010160408051908101604052807fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81526020016103e860280281525090806001815401808255809150509060018203906000526020600020906002020160009091929091909150600082015181600001556020820151816001015550505060039150600660008381526020019081526020016000209050818160000181905550806001016040805190810160405280600a81526020016103e8600102815250908060018154018082558091505090600182039060005260206000209060020201600090919290919091506000820151816000015560208201518160010155505050806001016040805190810160405280603281526020016103e8600502815250908060018154018082558091505090600182039060005260206000209060020201600090919290919091506000820151816000015560208201518160010155505050806001016040805190810160405280606481526020016103e8600a0281525090806001815401808255809150509060018203906000526020600020906002020160009091929091909150600082015181600001556020820151816001015550505080600101604080519081016040528060c881526020016103e86014028152509080600181540180825580915050906001820390600052602060002090600202016000909192909190915060008201518160000155602082015181600101555050508060010160408051908101604052806101f481526020016103e8601e028152509080600181540180825580915050906001820390600052602060002090600202016000909192909190915060008201518160000155602082015181600101555050508060010160408051908101604052807fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81526020016103e860280281525090806001815401808255809150509060018203906000526020600020906002020160009091929091909150600082015181600001556020820151816001015550505050505062000bcf565b60008190508073ffffffffffffffffffffffffffffffffffffffff1663519c28826040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b15801562000a2957600080fd5b505af115801562000a3e573d6000803e3d6000fd5b505050506040513d602081101562000a5557600080fd5b8101908080519060200190929190505050151562000adb576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260118152602001807f6e6f7420426f6b65724d616e616765722100000000000000000000000000000081525060200191505060405180910390fd5b80600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1062000b6357805160ff191683800117855562000b94565b8280016001018555821562000b94579182015b8281111562000b9357825182559160200191906001019062000b76565b5b50905062000ba3919062000ba7565b5090565b62000bcc91905b8082111562000bc857600081600090555060010162000bae565b5090565b90565b611e2b8062000bdf6000396000f3006080604052600436106100e6576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff168063061a74cd146100eb5780631eb726af14610118578063378298bc146101c15780633fe59a8c146102a357806354fd4d50146102d057806361dcd7ab1461036057806366ebc1c61461038b578063692202e1146103e25780636b3ee14a146104ac578063897df680146104d75780638da5cb5b146105225780639781009114610579578063c38ce9f8146105d0578063d0ebdbe7146105fb578063d43c80211461063e578063f2fde38b146106bb575b600080fd5b3480156100f757600080fd5b50610116600480360381019080803590602001909291905050506106fe565b005b34801561012457600080fd5b5061017f600480360381019080803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290505050610ab7565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156101cd57600080fd5b50610228600480360381019080803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290505050610d65565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561026857808201518184015260208101905061024d565b50505050905090810190601f1680156102955780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b3480156102af57600080fd5b506102ce60048036038101908080359060200190929190505050610ef5565b005b3480156102dc57600080fd5b506102e56112ae565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561032557808201518184015260208101905061030a565b50505050905090810190601f1680156103525780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561036c57600080fd5b5061037561134c565b6040518082815260200191505060405180910390f35b34801561039757600080fd5b506103a0611352565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156103ee57600080fd5b5061040d60048036038101908080359060200190929190505050611378565b604051808060200180602001838103835285818151815260200191508051906020019060200280838360005b83811015610454578082015181840152602081019050610439565b50505050905001838103825284818151815260200191508051906020019060200280838360005b8381101561049657808201518184015260208101905061047b565b5050505090500194505050505060405180910390f35b3480156104b857600080fd5b506104c161149d565b6040518082815260200191505060405180910390f35b3480156104e357600080fd5b50610520600480360381019080803590602001909291908035906020019092919080359060200190929190803590602001909291905050506114a3565b005b34801561052e57600080fd5b5061053761170e565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561058557600080fd5b506105ba600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611733565b6040518082815260200191505060405180910390f35b3480156105dc57600080fd5b506105e561199a565b6040518082815260200191505060405180910390f35b34801561060757600080fd5b5061063c600480360381019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506119a0565b005b34801561064a57600080fd5b506106a5600480360381019080803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290505050611a07565b6040518082815260200191505060405180910390f35b3480156106c757600080fd5b506106fc600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611b41565b005b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663d909fb1b6040805190810160405280600581526020017f61646d696e000000000000000000000000000000000000000000000000000000815250336040518363ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040180806020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001828103825284818151815260200191508051906020019080838360005b838110156108105780820151818401526020810190506107f5565b50505050905090810190601f16801561083d5780820380516001836020036101000a031916815260200191505b509350505050602060405180830381600087803b15801561085d57600080fd5b505af1158015610871573d6000803e3d6000fd5b505050506040513d602081101561088757600080fd5b810190808051906020019092919050505080610a395750600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663d909fb1b6040805190810160405280600881526020017f636f6e7472616374000000000000000000000000000000000000000000000000815250336040518363ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040180806020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001828103825284818151815260200191508051906020019080838360005b838110156109b0578082015181840152602081019050610995565b50505050905090810190601f1680156109dd5780820380516001836020036101000a031916815260200191505b509350505050602060405180830381600087803b1580156109fd57600080fd5b505af1158015610a11573d6000803e3d6000fd5b505050506040513d6020811015610a2757600080fd5b81019080805190602001909291905050505b1515610aad576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600f8152602001807f6e6f7420617574686f72697a656421000000000000000000000000000000000081525060200191505060405180910390fd5b8060048190555050565b600080600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663fca1f3c1846040518263ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018080602001828103825283818151815260200191508051906020019080838360005b83811015610b64578082015181840152602081019050610b49565b50505050905090810190601f168015610b915780820380516001836020036101000a031916815260200191505b5092505050600060405180830381600087803b158015610bb057600080fd5b505af1158015610bc4573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f8201168201806040525060a0811015610bee57600080fd5b81019080805190602001909291908051640100000000811115610c1057600080fd5b82810190506020810184811115610c2657600080fd5b8151856001820283011164010000000082111715610c4357600080fd5b5050929190602001805190602001909291908051640100000000811115610c6957600080fd5b82810190506020810184811115610c7f57600080fd5b8151856001820283011164010000000082111715610c9c57600080fd5b505092919060200180519060200190929190505050505092505050600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614151515610d5c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600a8152602001807f616464722069732030210000000000000000000000000000000000000000000081525060200191505060405180910390fd5b80915050919050565b6060600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16634a189f35836040518263ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018080602001828103825283818151815260200191508051906020019080838360005b83811015610e11578082015181840152602081019050610df6565b50505050905090810190601f168015610e3e5780820380516001836020036101000a031916815260200191505b5092505050600060405180830381600087803b158015610e5d57600080fd5b505af1158015610e71573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f820116820180604052506020811015610e9b57600080fd5b810190808051640100000000811115610eb357600080fd5b82810190506020810184811115610ec957600080fd5b8151856001820283011164010000000082111715610ee657600080fd5b50509291905050509050919050565b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663d909fb1b6040805190810160405280600581526020017f61646d696e000000000000000000000000000000000000000000000000000000815250336040518363ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040180806020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001828103825284818151815260200191508051906020019080838360005b83811015611007578082015181840152602081019050610fec565b50505050905090810190601f1680156110345780820380516001836020036101000a031916815260200191505b509350505050602060405180830381600087803b15801561105457600080fd5b505af1158015611068573d6000803e3d6000fd5b505050506040513d602081101561107e57600080fd5b8101908080519060200190929190505050806112305750600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663d909fb1b6040805190810160405280600881526020017f636f6e7472616374000000000000000000000000000000000000000000000000815250336040518363ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040180806020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001828103825284818151815260200191508051906020019080838360005b838110156111a757808201518184015260208101905061118c565b50505050905090810190601f1680156111d45780820380516001836020036101000a031916815260200191505b509350505050602060405180830381600087803b1580156111f457600080fd5b505af1158015611208573d6000803e3d6000fd5b505050506040513d602081101561121e57600080fd5b81019080805190602001909291905050505b15156112a4576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600f8152602001807f6e6f7420617574686f72697a656421000000000000000000000000000000000081525060200191505060405180910390fd5b8060058190555050565b60018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156113445780601f1061131957610100808354040283529160200191611344565b820191906000526020600020905b81548152906001019060200180831161132757829003601f168201915b505050505081565b60025481565b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6060806000806000806006600088815260200190815260200160002093506000846000015414156113a857611494565b83600101805490509250826040519080825280602002602001820160405280156113e15781602001602082028038833980820191505090505b509550826040519080825280602002602001820160405280156114135781602001602082028038833980820191505090505b509450600091505b8282101561149357836001018281548110151561143457fe5b906000526020600020906002020190508060000154868381518110151561145757fe5b90602001906020020181815250508060010154858381518110151561147857fe5b9060200190602002018181525050818060010192505061141b565b5b50505050915091565b60055481565b6000600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16634d4810836040805190810160405280600581526020017f61646d696e000000000000000000000000000000000000000000000000000000815250336040518363ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040180806020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001828103825284818151815260200191508051906020019080838360005b838110156115b757808201518184015260208101905061159c565b50505050905090810190601f1680156115e45780820380516001836020036101000a031916815260200191505b509350505050600060405180830381600087803b15801561160457600080fd5b505af1158015611618573d6000803e3d6000fd5b50505050600660008681526020019081526020016000209050600081600001541415611648578481600001819055505b8060010180549050841015156116b85780600101604080519081016040528085815260200184815250908060018154018082558091505090600182039060005260206000209060020201600090919290919091506000820151816000015560208201518160010155505050611707565b8281600101858154811015156116ca57fe5b9060005260206000209060020201600001819055508181600101858154811015156116f157fe5b9060005260206000209060020201600101819055505b5050505050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000806000806000806000806000806117806040805190810160405280600d81526020017f426f6b6572446170704461746100000000000000000000000000000000000000815250610ab7565b73ffffffffffffffffffffffffffffffffffffffff166355670f318c6040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001915050608060405180830381600087803b15801561181a57600080fd5b505af115801561182e573d6000803e3d6000fd5b505050506040513d608081101561184457600080fd5b810190808051906020019092919080519060200190929190805190602001909291908051906020019092919050505099509950995050600089141561188c576000995061198c565b600660008a815260200190815260200160002095506000866000015414156118b7576000995061198c565b60009450600093506000925088600114156118d4578692506118d8565b8792505b600091505b85600101805490508210156119875785600101828154811015156118fd57fe5b9060005260206000209060020201905080600001549350837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1415611956578483121515611951578060010154995061198c565b611977565b84831215801561196557508383125b15611976578060010154995061198c565b5b83945081806001019250506118dd565b600099505b505050505050505050919050565b60045481565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156119fb57600080fd5b611a0481611ba8565b50565b6000600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16633230b078836040518263ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018080602001828103825283818151815260200191508051906020019080838360005b83811015611ab3578082015181840152602081019050611a98565b50505050905090810190601f168015611ae05780820380516001836020036101000a031916815260200191505b5092505050602060405180830381600087803b158015611aff57600080fd5b505af1158015611b13573d6000803e3d6000fd5b505050506040513d6020811015611b2957600080fd5b81019080805190602001909291905050509050919050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515611b9c57600080fd5b611ba581611d05565b50565b60008190508073ffffffffffffffffffffffffffffffffffffffff1663519c28826040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b158015611c1157600080fd5b505af1158015611c25573d6000803e3d6000fd5b505050506040513d6020811015611c3b57600080fd5b81019080805190602001909291905050501515611cc0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260118152602001807f6e6f7420426f6b65724d616e616765722100000000000000000000000000000081525060200191505060405180910390fd5b80600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614151515611d4157600080fd5b8073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550505600a165627a7a72305820d195929f0ca5089c0803da58cc259ae3f71cce48e62c57ffd25f404839aae3e50029`

// DeployBokerTokenPowerData deploys a new Ethereum contract, binding an instance of BokerTokenPowerData to it.
func DeployBokerTokenPowerData(auth *bind.TransactOpts, backend bind.ContractBackend, addrManager common.Address) (common.Address, *types.Transaction, *BokerTokenPowerData, error) {
	parsed, err := abi.JSON(strings.NewReader(BokerTokenPowerDataABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BokerTokenPowerDataBin), backend, addrManager)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BokerTokenPowerData{BokerTokenPowerDataCaller: BokerTokenPowerDataCaller{contract: contract}, BokerTokenPowerDataTransactor: BokerTokenPowerDataTransactor{contract: contract}}, nil
}

// BokerTokenPowerData is an auto generated Go binding around an Ethereum contract.
type BokerTokenPowerData struct {
	BokerTokenPowerDataCaller     // Read-only binding to the contract
	BokerTokenPowerDataTransactor // Write-only binding to the contract
}

// BokerTokenPowerDataCaller is an auto generated read-only Go binding around an Ethereum contract.
type BokerTokenPowerDataCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BokerTokenPowerDataTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BokerTokenPowerDataTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BokerTokenPowerDataSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BokerTokenPowerDataSession struct {
	Contract     *BokerTokenPowerData // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// BokerTokenPowerDataCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BokerTokenPowerDataCallerSession struct {
	Contract *BokerTokenPowerDataCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// BokerTokenPowerDataTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BokerTokenPowerDataTransactorSession struct {
	Contract     *BokerTokenPowerDataTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// BokerTokenPowerDataRaw is an auto generated low-level Go binding around an Ethereum contract.
type BokerTokenPowerDataRaw struct {
	Contract *BokerTokenPowerData // Generic contract binding to access the raw methods on
}

// BokerTokenPowerDataCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BokerTokenPowerDataCallerRaw struct {
	Contract *BokerTokenPowerDataCaller // Generic read-only contract binding to access the raw methods on
}

// BokerTokenPowerDataTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BokerTokenPowerDataTransactorRaw struct {
	Contract *BokerTokenPowerDataTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBokerTokenPowerData creates a new instance of BokerTokenPowerData, bound to a specific deployed contract.
func NewBokerTokenPowerData(address common.Address, backend bind.ContractBackend) (*BokerTokenPowerData, error) {
	contract, err := bindBokerTokenPowerData(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BokerTokenPowerData{BokerTokenPowerDataCaller: BokerTokenPowerDataCaller{contract: contract}, BokerTokenPowerDataTransactor: BokerTokenPowerDataTransactor{contract: contract}}, nil
}

// NewBokerTokenPowerDataCaller creates a new read-only instance of BokerTokenPowerData, bound to a specific deployed contract.
func NewBokerTokenPowerDataCaller(address common.Address, caller bind.ContractCaller) (*BokerTokenPowerDataCaller, error) {
	contract, err := bindBokerTokenPowerData(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &BokerTokenPowerDataCaller{contract: contract}, nil
}

// NewBokerTokenPowerDataTransactor creates a new write-only instance of BokerTokenPowerData, bound to a specific deployed contract.
func NewBokerTokenPowerDataTransactor(address common.Address, transactor bind.ContractTransactor) (*BokerTokenPowerDataTransactor, error) {
	contract, err := bindBokerTokenPowerData(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &BokerTokenPowerDataTransactor{contract: contract}, nil
}

// bindBokerTokenPowerData binds a generic wrapper to an already deployed contract.
func bindBokerTokenPowerData(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BokerTokenPowerDataABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BokerTokenPowerData *BokerTokenPowerDataRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BokerTokenPowerData.Contract.BokerTokenPowerDataCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BokerTokenPowerData *BokerTokenPowerDataRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BokerTokenPowerData.Contract.BokerTokenPowerDataTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BokerTokenPowerData *BokerTokenPowerDataRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BokerTokenPowerData.Contract.BokerTokenPowerDataTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BokerTokenPowerData *BokerTokenPowerDataCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BokerTokenPowerData.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BokerTokenPowerData *BokerTokenPowerDataTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BokerTokenPowerData.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BokerTokenPowerData *BokerTokenPowerDataTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BokerTokenPowerData.Contract.contract.Transact(opts, method, params...)
}

// AssignCycleBegin is a free data retrieval call binding the contract method 0xc38ce9f8.
//
// Solidity: function assignCycleBegin() constant returns(uint256)
func (_BokerTokenPowerData *BokerTokenPowerDataCaller) AssignCycleBegin(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BokerTokenPowerData.contract.Call(opts, out, "assignCycleBegin")
	return *ret0, err
}

// AssignCycleBegin is a free data retrieval call binding the contract method 0xc38ce9f8.
//
// Solidity: function assignCycleBegin() constant returns(uint256)
func (_BokerTokenPowerData *BokerTokenPowerDataSession) AssignCycleBegin() (*big.Int, error) {
	return _BokerTokenPowerData.Contract.AssignCycleBegin(&_BokerTokenPowerData.CallOpts)
}

// AssignCycleBegin is a free data retrieval call binding the contract method 0xc38ce9f8.
//
// Solidity: function assignCycleBegin() constant returns(uint256)
func (_BokerTokenPowerData *BokerTokenPowerDataCallerSession) AssignCycleBegin() (*big.Int, error) {
	return _BokerTokenPowerData.Contract.AssignCycleBegin(&_BokerTokenPowerData.CallOpts)
}

// BindDappConfigGet is a free data retrieval call binding the contract method 0x692202e1.
//
// Solidity: function bindDappConfigGet(dappType uint256) constant returns(uppers int256[], powers uint256[])
func (_BokerTokenPowerData *BokerTokenPowerDataCaller) BindDappConfigGet(opts *bind.CallOpts, dappType *big.Int) (struct {
	Uppers []*big.Int
	Powers []*big.Int
}, error) {
	ret := new(struct {
		Uppers []*big.Int
		Powers []*big.Int
	})
	out := ret
	err := _BokerTokenPowerData.contract.Call(opts, out, "bindDappConfigGet", dappType)
	return *ret, err
}

// BindDappConfigGet is a free data retrieval call binding the contract method 0x692202e1.
//
// Solidity: function bindDappConfigGet(dappType uint256) constant returns(uppers int256[], powers uint256[])
func (_BokerTokenPowerData *BokerTokenPowerDataSession) BindDappConfigGet(dappType *big.Int) (struct {
	Uppers []*big.Int
	Powers []*big.Int
}, error) {
	return _BokerTokenPowerData.Contract.BindDappConfigGet(&_BokerTokenPowerData.CallOpts, dappType)
}

// BindDappConfigGet is a free data retrieval call binding the contract method 0x692202e1.
//
// Solidity: function bindDappConfigGet(dappType uint256) constant returns(uppers int256[], powers uint256[])
func (_BokerTokenPowerData *BokerTokenPowerDataCallerSession) BindDappConfigGet(dappType *big.Int) (struct {
	Uppers []*big.Int
	Powers []*big.Int
}, error) {
	return _BokerTokenPowerData.Contract.BindDappConfigGet(&_BokerTokenPowerData.CallOpts, dappType)
}

// BindDappGetPower is a free data retrieval call binding the contract method 0x97810091.
//
// Solidity: function bindDappGetPower(dappAddr address) constant returns(uint256)
func (_BokerTokenPowerData *BokerTokenPowerDataCaller) BindDappGetPower(opts *bind.CallOpts, dappAddr common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BokerTokenPowerData.contract.Call(opts, out, "bindDappGetPower", dappAddr)
	return *ret0, err
}

// BindDappGetPower is a free data retrieval call binding the contract method 0x97810091.
//
// Solidity: function bindDappGetPower(dappAddr address) constant returns(uint256)
func (_BokerTokenPowerData *BokerTokenPowerDataSession) BindDappGetPower(dappAddr common.Address) (*big.Int, error) {
	return _BokerTokenPowerData.Contract.BindDappGetPower(&_BokerTokenPowerData.CallOpts, dappAddr)
}

// BindDappGetPower is a free data retrieval call binding the contract method 0x97810091.
//
// Solidity: function bindDappGetPower(dappAddr address) constant returns(uint256)
func (_BokerTokenPowerData *BokerTokenPowerDataCallerSession) BindDappGetPower(dappAddr common.Address) (*big.Int, error) {
	return _BokerTokenPowerData.Contract.BindDappGetPower(&_BokerTokenPowerData.CallOpts, dappAddr)
}

// BokerManager is a free data retrieval call binding the contract method 0x66ebc1c6.
//
// Solidity: function bokerManager() constant returns(address)
func (_BokerTokenPowerData *BokerTokenPowerDataCaller) BokerManager(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BokerTokenPowerData.contract.Call(opts, out, "bokerManager")
	return *ret0, err
}

// BokerManager is a free data retrieval call binding the contract method 0x66ebc1c6.
//
// Solidity: function bokerManager() constant returns(address)
func (_BokerTokenPowerData *BokerTokenPowerDataSession) BokerManager() (common.Address, error) {
	return _BokerTokenPowerData.Contract.BokerManager(&_BokerTokenPowerData.CallOpts)
}

// BokerManager is a free data retrieval call binding the contract method 0x66ebc1c6.
//
// Solidity: function bokerManager() constant returns(address)
func (_BokerTokenPowerData *BokerTokenPowerDataCallerSession) BokerManager() (common.Address, error) {
	return _BokerTokenPowerData.Contract.BokerManager(&_BokerTokenPowerData.CallOpts)
}

// ContractAddress is a free data retrieval call binding the contract method 0x1eb726af.
//
// Solidity: function contractAddress(cName string) constant returns(address)
func (_BokerTokenPowerData *BokerTokenPowerDataCaller) ContractAddress(opts *bind.CallOpts, cName string) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BokerTokenPowerData.contract.Call(opts, out, "contractAddress", cName)
	return *ret0, err
}

// ContractAddress is a free data retrieval call binding the contract method 0x1eb726af.
//
// Solidity: function contractAddress(cName string) constant returns(address)
func (_BokerTokenPowerData *BokerTokenPowerDataSession) ContractAddress(cName string) (common.Address, error) {
	return _BokerTokenPowerData.Contract.ContractAddress(&_BokerTokenPowerData.CallOpts, cName)
}

// ContractAddress is a free data retrieval call binding the contract method 0x1eb726af.
//
// Solidity: function contractAddress(cName string) constant returns(address)
func (_BokerTokenPowerData *BokerTokenPowerDataCallerSession) ContractAddress(cName string) (common.Address, error) {
	return _BokerTokenPowerData.Contract.ContractAddress(&_BokerTokenPowerData.CallOpts, cName)
}

// CreateTime is a free data retrieval call binding the contract method 0x61dcd7ab.
//
// Solidity: function createTime() constant returns(uint256)
func (_BokerTokenPowerData *BokerTokenPowerDataCaller) CreateTime(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BokerTokenPowerData.contract.Call(opts, out, "createTime")
	return *ret0, err
}

// CreateTime is a free data retrieval call binding the contract method 0x61dcd7ab.
//
// Solidity: function createTime() constant returns(uint256)
func (_BokerTokenPowerData *BokerTokenPowerDataSession) CreateTime() (*big.Int, error) {
	return _BokerTokenPowerData.Contract.CreateTime(&_BokerTokenPowerData.CallOpts)
}

// CreateTime is a free data retrieval call binding the contract method 0x61dcd7ab.
//
// Solidity: function createTime() constant returns(uint256)
func (_BokerTokenPowerData *BokerTokenPowerDataCallerSession) CreateTime() (*big.Int, error) {
	return _BokerTokenPowerData.Contract.CreateTime(&_BokerTokenPowerData.CallOpts)
}

// GlobalConfigInt is a free data retrieval call binding the contract method 0xd43c8021.
//
// Solidity: function globalConfigInt(key string) constant returns(uint256)
func (_BokerTokenPowerData *BokerTokenPowerDataCaller) GlobalConfigInt(opts *bind.CallOpts, key string) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BokerTokenPowerData.contract.Call(opts, out, "globalConfigInt", key)
	return *ret0, err
}

// GlobalConfigInt is a free data retrieval call binding the contract method 0xd43c8021.
//
// Solidity: function globalConfigInt(key string) constant returns(uint256)
func (_BokerTokenPowerData *BokerTokenPowerDataSession) GlobalConfigInt(key string) (*big.Int, error) {
	return _BokerTokenPowerData.Contract.GlobalConfigInt(&_BokerTokenPowerData.CallOpts, key)
}

// GlobalConfigInt is a free data retrieval call binding the contract method 0xd43c8021.
//
// Solidity: function globalConfigInt(key string) constant returns(uint256)
func (_BokerTokenPowerData *BokerTokenPowerDataCallerSession) GlobalConfigInt(key string) (*big.Int, error) {
	return _BokerTokenPowerData.Contract.GlobalConfigInt(&_BokerTokenPowerData.CallOpts, key)
}

// GlobalConfigString is a free data retrieval call binding the contract method 0x378298bc.
//
// Solidity: function globalConfigString(key string) constant returns(string)
func (_BokerTokenPowerData *BokerTokenPowerDataCaller) GlobalConfigString(opts *bind.CallOpts, key string) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _BokerTokenPowerData.contract.Call(opts, out, "globalConfigString", key)
	return *ret0, err
}

// GlobalConfigString is a free data retrieval call binding the contract method 0x378298bc.
//
// Solidity: function globalConfigString(key string) constant returns(string)
func (_BokerTokenPowerData *BokerTokenPowerDataSession) GlobalConfigString(key string) (string, error) {
	return _BokerTokenPowerData.Contract.GlobalConfigString(&_BokerTokenPowerData.CallOpts, key)
}

// GlobalConfigString is a free data retrieval call binding the contract method 0x378298bc.
//
// Solidity: function globalConfigString(key string) constant returns(string)
func (_BokerTokenPowerData *BokerTokenPowerDataCallerSession) GlobalConfigString(key string) (string, error) {
	return _BokerTokenPowerData.Contract.GlobalConfigString(&_BokerTokenPowerData.CallOpts, key)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_BokerTokenPowerData *BokerTokenPowerDataCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BokerTokenPowerData.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_BokerTokenPowerData *BokerTokenPowerDataSession) Owner() (common.Address, error) {
	return _BokerTokenPowerData.Contract.Owner(&_BokerTokenPowerData.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_BokerTokenPowerData *BokerTokenPowerDataCallerSession) Owner() (common.Address, error) {
	return _BokerTokenPowerData.Contract.Owner(&_BokerTokenPowerData.CallOpts)
}

// TokenAssignedTotal is a free data retrieval call binding the contract method 0x6b3ee14a.
//
// Solidity: function tokenAssignedTotal() constant returns(uint256)
func (_BokerTokenPowerData *BokerTokenPowerDataCaller) TokenAssignedTotal(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BokerTokenPowerData.contract.Call(opts, out, "tokenAssignedTotal")
	return *ret0, err
}

// TokenAssignedTotal is a free data retrieval call binding the contract method 0x6b3ee14a.
//
// Solidity: function tokenAssignedTotal() constant returns(uint256)
func (_BokerTokenPowerData *BokerTokenPowerDataSession) TokenAssignedTotal() (*big.Int, error) {
	return _BokerTokenPowerData.Contract.TokenAssignedTotal(&_BokerTokenPowerData.CallOpts)
}

// TokenAssignedTotal is a free data retrieval call binding the contract method 0x6b3ee14a.
//
// Solidity: function tokenAssignedTotal() constant returns(uint256)
func (_BokerTokenPowerData *BokerTokenPowerDataCallerSession) TokenAssignedTotal() (*big.Int, error) {
	return _BokerTokenPowerData.Contract.TokenAssignedTotal(&_BokerTokenPowerData.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() constant returns(string)
func (_BokerTokenPowerData *BokerTokenPowerDataCaller) Version(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _BokerTokenPowerData.contract.Call(opts, out, "version")
	return *ret0, err
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() constant returns(string)
func (_BokerTokenPowerData *BokerTokenPowerDataSession) Version() (string, error) {
	return _BokerTokenPowerData.Contract.Version(&_BokerTokenPowerData.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() constant returns(string)
func (_BokerTokenPowerData *BokerTokenPowerDataCallerSession) Version() (string, error) {
	return _BokerTokenPowerData.Contract.Version(&_BokerTokenPowerData.CallOpts)
}

// BindDappConfigSet is a paid mutator transaction binding the contract method 0x897df680.
//
// Solidity: function bindDappConfigSet(dappType uint256, index uint256, upper int256, power uint256) returns()
func (_BokerTokenPowerData *BokerTokenPowerDataTransactor) BindDappConfigSet(opts *bind.TransactOpts, dappType *big.Int, index *big.Int, upper *big.Int, power *big.Int) (*types.Transaction, error) {
	return _BokerTokenPowerData.contract.Transact(opts, "bindDappConfigSet", dappType, index, upper, power)
}

// BindDappConfigSet is a paid mutator transaction binding the contract method 0x897df680.
//
// Solidity: function bindDappConfigSet(dappType uint256, index uint256, upper int256, power uint256) returns()
func (_BokerTokenPowerData *BokerTokenPowerDataSession) BindDappConfigSet(dappType *big.Int, index *big.Int, upper *big.Int, power *big.Int) (*types.Transaction, error) {
	return _BokerTokenPowerData.Contract.BindDappConfigSet(&_BokerTokenPowerData.TransactOpts, dappType, index, upper, power)
}

// BindDappConfigSet is a paid mutator transaction binding the contract method 0x897df680.
//
// Solidity: function bindDappConfigSet(dappType uint256, index uint256, upper int256, power uint256) returns()
func (_BokerTokenPowerData *BokerTokenPowerDataTransactorSession) BindDappConfigSet(dappType *big.Int, index *big.Int, upper *big.Int, power *big.Int) (*types.Transaction, error) {
	return _BokerTokenPowerData.Contract.BindDappConfigSet(&_BokerTokenPowerData.TransactOpts, dappType, index, upper, power)
}

// SetAssignCycleBegin is a paid mutator transaction binding the contract method 0x061a74cd.
//
// Solidity: function setAssignCycleBegin(time uint256) returns()
func (_BokerTokenPowerData *BokerTokenPowerDataTransactor) SetAssignCycleBegin(opts *bind.TransactOpts, time *big.Int) (*types.Transaction, error) {
	return _BokerTokenPowerData.contract.Transact(opts, "setAssignCycleBegin", time)
}

// SetAssignCycleBegin is a paid mutator transaction binding the contract method 0x061a74cd.
//
// Solidity: function setAssignCycleBegin(time uint256) returns()
func (_BokerTokenPowerData *BokerTokenPowerDataSession) SetAssignCycleBegin(time *big.Int) (*types.Transaction, error) {
	return _BokerTokenPowerData.Contract.SetAssignCycleBegin(&_BokerTokenPowerData.TransactOpts, time)
}

// SetAssignCycleBegin is a paid mutator transaction binding the contract method 0x061a74cd.
//
// Solidity: function setAssignCycleBegin(time uint256) returns()
func (_BokerTokenPowerData *BokerTokenPowerDataTransactorSession) SetAssignCycleBegin(time *big.Int) (*types.Transaction, error) {
	return _BokerTokenPowerData.Contract.SetAssignCycleBegin(&_BokerTokenPowerData.TransactOpts, time)
}

// SetManager is a paid mutator transaction binding the contract method 0xd0ebdbe7.
//
// Solidity: function setManager(addrManager address) returns()
func (_BokerTokenPowerData *BokerTokenPowerDataTransactor) SetManager(opts *bind.TransactOpts, addrManager common.Address) (*types.Transaction, error) {
	return _BokerTokenPowerData.contract.Transact(opts, "setManager", addrManager)
}

// SetManager is a paid mutator transaction binding the contract method 0xd0ebdbe7.
//
// Solidity: function setManager(addrManager address) returns()
func (_BokerTokenPowerData *BokerTokenPowerDataSession) SetManager(addrManager common.Address) (*types.Transaction, error) {
	return _BokerTokenPowerData.Contract.SetManager(&_BokerTokenPowerData.TransactOpts, addrManager)
}

// SetManager is a paid mutator transaction binding the contract method 0xd0ebdbe7.
//
// Solidity: function setManager(addrManager address) returns()
func (_BokerTokenPowerData *BokerTokenPowerDataTransactorSession) SetManager(addrManager common.Address) (*types.Transaction, error) {
	return _BokerTokenPowerData.Contract.SetManager(&_BokerTokenPowerData.TransactOpts, addrManager)
}

// TokenAssignedTotalSet is a paid mutator transaction binding the contract method 0x3fe59a8c.
//
// Solidity: function tokenAssignedTotalSet(val uint256) returns()
func (_BokerTokenPowerData *BokerTokenPowerDataTransactor) TokenAssignedTotalSet(opts *bind.TransactOpts, val *big.Int) (*types.Transaction, error) {
	return _BokerTokenPowerData.contract.Transact(opts, "tokenAssignedTotalSet", val)
}

// TokenAssignedTotalSet is a paid mutator transaction binding the contract method 0x3fe59a8c.
//
// Solidity: function tokenAssignedTotalSet(val uint256) returns()
func (_BokerTokenPowerData *BokerTokenPowerDataSession) TokenAssignedTotalSet(val *big.Int) (*types.Transaction, error) {
	return _BokerTokenPowerData.Contract.TokenAssignedTotalSet(&_BokerTokenPowerData.TransactOpts, val)
}

// TokenAssignedTotalSet is a paid mutator transaction binding the contract method 0x3fe59a8c.
//
// Solidity: function tokenAssignedTotalSet(val uint256) returns()
func (_BokerTokenPowerData *BokerTokenPowerDataTransactorSession) TokenAssignedTotalSet(val *big.Int) (*types.Transaction, error) {
	return _BokerTokenPowerData.Contract.TokenAssignedTotalSet(&_BokerTokenPowerData.TransactOpts, val)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_BokerTokenPowerData *BokerTokenPowerDataTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _BokerTokenPowerData.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_BokerTokenPowerData *BokerTokenPowerDataSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _BokerTokenPowerData.Contract.TransferOwnership(&_BokerTokenPowerData.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_BokerTokenPowerData *BokerTokenPowerDataTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _BokerTokenPowerData.Contract.TransferOwnership(&_BokerTokenPowerData.TransactOpts, _newOwner)
}
