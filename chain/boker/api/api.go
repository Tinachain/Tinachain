package bokerapi

import (
	"context"
	"math/big"

	"github.com/Tinachain/Tina/chain/boker/protocol"
	"github.com/Tinachain/Tina/chain/common"
	"github.com/Tinachain/Tina/chain/core/types"
	"github.com/Tinachain/Tina/chain/trie"
)

//Tina链定义的接口
type Api interface {
	//查询帐号得到此账号可以进行的主交易类型
	GetAccount(address common.Address) (protocol.TxMajor, error)
	SetContract(address common.Address, contractType protocol.ContractType, isCancel bool, abiJson string) error //设置合约级别

	IsLocalValidator(address common.Address) bool          //判断是否是设置的本地验证者
	GetContractTrie() (*trie.Trie, *trie.Trie, *trie.Trie) //得到合约树

	GetMethodName(txMinor protocol.TxMinor) (string, string, error) //根据交易类型得到方法名称（只适用于基础合约）
	GetContract(address common.Address) (protocol.TxMajor, error)   //判断合约属性
	EncoderContext(context []byte, key []byte) (error, []byte)      //对内容进行编码
	DecoderContext(context []byte, key []byte) (error, []byte)      //对内容进行解码

	//黑名单
	GetBlacks() []common.Address                   //得到当前黑名单
	SetBlacks(address []common.Address) error      //设置当前的黑名单
	CheckBlackAddress(address common.Address) bool //检测地址是否在黑名单中

	//股权相关操作
	GetOwner() common.Address                                                                            //得到所有者账号
	SetOwner(operation common.Address, address common.Address) error                                     //设置所有者账号
	GetGasPool() uint64                                                                                  //得到待分配币尺数量
	AddGasPool(gas uint64) uint64                                                                        //增加待分配币尺数量
	GetStock(address common.Address) *StockAccount                                                       //得到股权信息
	GetStocks() ([]*StockAccount, error)                                                                 //得到所有股权信息
	SetStock(operation common.Address, address common.Address, number uint64) error                      //设置股权
	TransferStock(operation common.Address, from common.Address, to common.Address, number uint64) error //转移股权
	CleanStock(operation common.Address, from common.Address) error                                      //清空股权
	FrozenStock(operation common.Address, from common.Address) error                                     //冻结股权
	UnFrozenStock(operation common.Address, from common.Address) error                                   //解冻股权
	ClearStock(operation common.Address) error                                                           //清空所有人的股权
	GetStockTrie() *trie.Trie                                                                            //得到股权树

	//链交易
	SubmitBokerTransaction(ctx context.Context, txMajor protocol.TxMajor, txMinor protocol.TxMinor, to common.Address, name []byte, extra []byte, encryption uint8) (*types.Transaction, error) //产生一个设置验证者交易
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
