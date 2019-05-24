package bokerapi

import (
	"context"
	"math/big"

	"github.com/Tinachain/Tina/chain/boker/protocol"
	"github.com/Tinachain/Tina/chain/common"
	"github.com/Tinachain/Tina/chain/trie"
)

//Tina链定义的接口
type Api interface {
	//查询帐号得到此账号可以进行的主交易类型
	GetAccount(address common.Address) (protocol.TxMajor, error)
	SetContract(address common.Address,
		contractType protocol.ContractType,
		isCancel bool,
		abiJson string) error //设置合约级别

	SubmitBokerTransaction(ctx context.Context,
		txMajor protocol.TxMajor,
		txMinor protocol.TxMinor,
		to common.Address,
		name []byte,
		extra []byte) error //产生一个设置验证者交易

	IsLocalValidator(address common.Address) bool //判断是否是设置的本地验证者
	GetContractTrie() (*trie.Trie,
		*trie.Trie,
		*trie.Trie) //得到合约树

	GetMethodName(txMinor protocol.TxMinor) (string, string, error) //根据交易类型得到方法名称（只适用于基础合约）
	GetContract(address common.Address) (protocol.TxMajor, error)   //判断合约属性
	GetTeamAccount() common.Address                                 //得到团队账号
}

type SortableAddress struct {
	Address common.Address
	Weight  *big.Int
}
type SortableAddresses []*SortableAddress

func (p SortableAddresses) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
func (p SortableAddresses) Len() int      { return len(p) }
func (p SortableAddresses) Less(i, j int) bool {
	if p[i].Weight.Cmp(p[j].Weight) < 0 {
		return false
	} else if p[i].Weight.Cmp(p[j].Weight) > 0 {
		return true
	} else {
		return p[i].Address.String() < p[j].Address.String()
	}
}
