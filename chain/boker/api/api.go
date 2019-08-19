package bokerapi

import (
	"context"
	"math/big"

	"github.com/Tinachain/Tina/chain/boker/protocol"
	"github.com/Tinachain/Tina/chain/common"
	"github.com/Tinachain/Tina/chain/core/types"
)

//Tina链定义的接口
type Api interface {
	SetSystemContract(address common.Address, from common.Address, bokerContext *types.BokerContext) error
	IsLocalValidator(address common.Address) bool //判断是否是设置的本地验证者
	GetVotes() error

	//黑名单
	GetBlacks() []common.Address                   //得到当前黑名单
	SetBlacks(address []common.Address) error      //设置当前的黑名单
	CheckBlackAddress(address common.Address) bool //检测地址是否在黑名单中

	//链交易
	SubmitBokerTransaction(ctx context.Context, txMajor protocol.TxMajor, txMinor protocol.TxMinor, from, to common.Address, name, extra []byte, value *big.Int, encryption uint8) (*types.Transaction, error) //产生一个设置验证者交易
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
