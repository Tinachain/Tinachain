package core

import (
	"errors"
	"math/big"

	"github.com/Tinachain/Tina/chain/boker/api"
	"github.com/Tinachain/Tina/chain/boker/protocol"
	"github.com/Tinachain/Tina/chain/common"
	"github.com/Tinachain/Tina/chain/common/math"
	"github.com/Tinachain/Tina/chain/core/types"
	"github.com/Tinachain/Tina/chain/core/vm"
	"github.com/Tinachain/Tina/chain/log"
	"github.com/Tinachain/Tina/chain/params"
)

var (
	Big0                         = big.NewInt(0)
	errInsufficientBalanceForGas = errors.New("insufficient balance to pay for gas")
)

/*
The State Transitioning Model
状态转换模型

A state transition is a change made when a transaction is applied to the current world state
状态转换 是指用当前的world state来执行交易，并改变当前的world state

The state transitioning model does all all the necessary work to work out a valid new state root.
状态转换做了所有所需的工作来产生一个新的有效的state root

1) Nonce handling	Nonce 处理
2) Pre pay gas		预先支付Gas
3) Create a new state object if the recipient is \0*32		如果接收人是空，那么创建一个新的state object
4) Value transfer	转账
== If contract creation ==
  4a) Attempt to run transaction data		尝试运行输入的数据
  4b) If valid, use result as code for the new state object	如果有效，那么用运行的结果作为新的state object的code
== end ==
5) Run Script section	运行脚本部分
6) Derive new state root	导出新的state root
*/
type StateTransition struct {
	gp         *GasPool //区块剩余的Gas池
	sp         *big.Int //区块剩余的空间池
	msg        Message
	gas        uint64
	gasPrice   *big.Int     // gas的价格
	initialGas *big.Int     // 最开始的gas
	value      *big.Int     // 转账的值
	data       []byte       // 输入数据
	extra      []byte       //扩展字段
	state      vm.StateDB   //StateDB对象
	evm        *vm.EVM      //虚拟机对象
	boker      bokerapi.Api //Tina链的接口对象
}

// Message represents a message sent to a contract.
//发送一个合约的消息
type Message interface {
	From() common.Address
	To() *common.Address
	GasPrice() *big.Int
	Gas() *big.Int
	Value() *big.Int
	Nonce() uint64
	CheckNonce() bool
	Data() []byte
	Extra() []byte
	Major() protocol.TxMajor
	Minor() protocol.TxMinor
}

//计算Gas。
/*
	这一段代码用来计算使用的Gas数量，从以上的算法可以看出，Gas的算法为
	Gas = 创建合约费用（或者交易费用） + 占用字节费用（合约中实际有数据的长度 * 68） + 非占用字节费用（合约中实际没有数据的长度 * 4）
*/
func IntrinsicGas(data []byte, contractCreation, homestead bool) *big.Int {

	//初始化一个igas变量用来保存计算出来的Gas数量
	igas := new(big.Int)

	//判断是否需要进行合约部署费用
	if contractCreation && homestead {

		//部署合约 53000
		igas.SetUint64(params.TxGasContractCreation)
	} else {

		//交易合约 21000
		igas.SetUint64(params.TxGas)
	}

	if len(data) > 0 {

		//过滤掉合约中的空数据（bye = 0），得到数据长度
		var nz int64
		for _, byt := range data {
			if byt != 0 {
				nz++
			}
		}

		//合约中占用的字节数 * 68 = 占用字节费用
		m := big.NewInt(nz)
		m.Mul(m, new(big.Int).SetUint64(params.TxDataNonZeroGas))
		//将创建合约 + 数据长度所需要的费用设置为igas费用
		igas.Add(igas, m)
		//得到合约数据中为空（byt = 0）的数量
		m.SetInt64(int64(len(data)) - nz)
		//将这个数量 * 4 成为一个数值
		m.Mul(m, new(big.Int).SetUint64(params.TxDataZeroGas))
		//将这个数值和原来的数值进行相加
		igas.Add(igas, m)

	}
	return igas
}

// NewStateTransition initialises and returns a new state transition object.
//创建一个交易的状态对象
func NewStateTransition(evm *vm.EVM, msg Message, gp *GasPool, sp *big.Int) *StateTransition {
	return &StateTransition{
		gp:         gp,
		sp:         sp,
		evm:        evm,
		msg:        msg,
		gasPrice:   msg.GasPrice(),
		initialGas: new(big.Int),
		value:      msg.Value(),
		data:       msg.Data(),
		extra:      msg.Extra(),
		state:      evm.StateDB,
	}
}

//普通交易处理
func NormalMessage(evm *vm.EVM,
	msg Message,
	gp *GasPool,
	sp *big.Int,
	dposContext *types.DposContext,
	bokerContext *types.BokerContext,
	boker bokerapi.Api) ([]byte, *big.Int, bool, error) {

	st := NewStateTransition(evm, msg, gp, sp)
	ret, _, gasUsed, failed, err := st.NormalTransitionDb(dposContext, bokerContext, boker)
	return ret, gasUsed, failed, err
}

//系统基础交易
func SystemBaseMessage(evm *vm.EVM,
	msg Message,
	gp *GasPool,
	sp *big.Int,
	dposContext *types.DposContext,
	bokerContext *types.BokerContext,
	boker bokerapi.Api) ([]byte, *big.Int, bool, error) {

	st := NewStateTransition(evm, msg, gp, sp)
	ret, _, _, failed, err := st.SystemBaseTransitionDb(dposContext, bokerContext, boker)
	return ret, new(big.Int).SetInt64(0), failed, err
}

//用户基础交易
func UserBaseMessage(evm *vm.EVM,
	msg Message,
	gp *GasPool,
	sp *big.Int,
	dposContext *types.DposContext,
	bokerContext *types.BokerContext,
	boker bokerapi.Api) ([]byte, *big.Int, bool, error) {

	st := NewStateTransition(evm, msg, gp, sp)
	ret, _, _, failed, err := st.UserBaseTransitionDb(dposContext, bokerContext, boker)
	return ret, new(big.Int).SetInt64(0), failed, err
}

//扩展交易
func ExtraMessage(evm *vm.EVM, msg Message, gp *GasPool, sp *big.Int, dposContext *types.DposContext,
	bokerContext *types.BokerContext, boker bokerapi.Api) ([]byte, *big.Int, bool, error) {

	st := NewStateTransition(evm, msg, gp, sp)
	ret, _, _, failed, err := st.ExtraTransitionDb(dposContext, bokerContext, boker)

	var gas *big.Int = new(big.Int).SetUint64(0)
	gas = gas.Mul(new(big.Int).SetUint64(90000), new(big.Int).SetUint64(50*params.Shannon))
	return ret, gas, failed, err
}

//股权交易
func StockMessage(evm *vm.EVM,
	msg Message,
	gp *GasPool,
	sp *big.Int,
	timer *big.Int,
	dposContext *types.DposContext,
	bokerContext *types.BokerContext,
	boker bokerapi.Api) ([]byte, *big.Int, bool, error) {

	st := NewStateTransition(evm, msg, gp, sp)
	ret, _, _, failed, err := st.StockTransitionDb(msg.Major(), msg.Minor(), timer, dposContext, bokerContext, boker)

	var gas *big.Int = new(big.Int).SetUint64(0)
	gas = gas.Mul(new(big.Int).SetUint64(0), new(big.Int).SetUint64(0))
	return ret, gas, failed, err
}

func systemContractMessage(evm *vm.EVM,
	msg Message,
	gp *GasPool,
	sp *big.Int,
	txMajor protocol.TxMajor,
	txMinor protocol.TxMinor,
	dposContext *types.DposContext,
	bokerContext *types.BokerContext, boker bokerapi.Api) ([]byte, *big.Int, bool, error) {

	st := NewStateTransition(evm, msg, gp, sp)
	ret, _, _, failed, err := st.SystemContractTransitionDb(txMajor, txMinor, dposContext, bokerContext, boker)
	return ret, new(big.Int).SetInt64(0), failed, err
}

func userContractMessage(evm *vm.EVM,
	msg Message,
	gp *GasPool,
	sp *big.Int,
	txMajor protocol.TxMajor,
	txMinor protocol.TxMinor,
	dposContext *types.DposContext,
	bokerContext *types.BokerContext, boker bokerapi.Api) ([]byte, *big.Int, bool, error) {

	st := NewStateTransition(evm, msg, gp, sp)
	ret, _, _, failed, err := st.UserContractTransitionDb(txMajor, txMinor, dposContext, bokerContext, boker)
	return ret, new(big.Int).SetInt64(0), failed, err
}

func validatorMessage(evm *vm.EVM, msg Message, gp *GasPool, sp *big.Int, txMajor protocol.TxMajor, txMinor protocol.TxMinor, dposContext *types.DposContext,
	bokerContext *types.BokerContext, boker bokerapi.Api) ([]byte, *big.Int, bool, error) {

	st := NewStateTransition(evm, msg, gp, sp)
	ret, _, _, failed, err := st.ValidatorTransitionDb(txMajor, txMinor, boker)
	return ret, new(big.Int).SetInt64(0), failed, err
}

//获取交易的from信息
func (st *StateTransition) from() vm.AccountRef {

	f := st.msg.From()
	if !st.state.Exist(f) {
		st.state.CreateAccount(f)
	}

	return vm.AccountRef(f)
}

//获取交易的to信息
func (st *StateTransition) to() vm.AccountRef {
	if st.msg == nil {
		return vm.AccountRef{}
	}
	to := st.msg.To()
	if to == nil {
		return vm.AccountRef{} // contract creation
	}

	reference := vm.AccountRef(*to)
	if !st.state.Exist(*to) {
		st.state.CreateAccount(*to)
	}
	return reference
}

//获取交易中实际使用的Gas信息，从Gas中减去实际费用，得到剩余费用
func (st *StateTransition) useGas(amount uint64) error {
	if st.gas < amount {
		return vm.ErrOutOfGas
	}
	st.gas -= amount

	return nil
}

func (st *StateTransition) buyGas() error {

	//得到消息的Gas
	mgas := st.msg.Gas()
	if mgas.BitLen() > 64 {
		return vm.ErrOutOfGas
	}

	//计算Gas的价格合计 = Gas * GasPrice
	mgval := new(big.Int).Mul(mgas, st.gasPrice)
	var (
		state  = st.state
		sender = st.from()
	)

	//判断用户账户中有足够的费用
	if state.GetBalance(sender.Address()).Cmp(mgval) < 0 {
		return errInsufficientBalanceForGas
	}

	//从GasPool中删除本次交易需要使用的Gas，如果大于0时，则说明交易可以被执行，如果小于0则说明本次交易不能被执行。
	if err := st.gp.SubGas(mgas); err != nil {
		return err
	}
	st.gas += mgas.Uint64()
	st.initialGas.Set(mgas)
	state.SubBalance(sender.Address(), mgval)
	return nil
}

//检测交易大小
func (st *StateTransition) checkTxSize() error {

	return nil
}

//
func (st *StateTransition) preCheck() error {
	msg := st.msg
	sender := st.from()

	// Make sure this transaction's nonce is correct
	if msg.CheckNonce() {
		nonce := st.state.GetNonce(sender.Address())
		if nonce < msg.Nonce() {
			return ErrNonceTooHigh
		} else if nonce > msg.Nonce() {
			return ErrNonceTooLow
		}
	}
	return st.buyGas()
}

func (st *StateTransition) NormalTransitionDb(dposContext *types.DposContext,
	bokerContext *types.BokerContext, boker bokerapi.Api) (ret []byte, requiredGas, usedGas *big.Int, failed bool, err error) {

	//log.Info("(st *StateTransition) NormalTransitionDb")
	if err = st.preCheck(); err != nil {
		return
	}

	msg := st.msg
	sender := st.from()
	homestead := true
	contractCreation := msg.To() == nil

	intrinsicGas := IntrinsicGas(st.data, contractCreation, homestead)
	if intrinsicGas.BitLen() > 64 {
		return nil, nil, nil, false, vm.ErrOutOfGas
	}

	if err = st.useGas(intrinsicGas.Uint64()); err != nil {
		return nil, nil, nil, false, err
	}

	var (
		evm   = st.evm
		vmerr error
	)

	if contractCreation {
		ret, _, st.gas, vmerr = evm.Create(sender, st.data, st.gas, st.value)
	} else {

		st.state.SetNonce(sender.Address(), st.state.GetNonce(sender.Address())+1)
		ret, st.gas, vmerr = evm.Call(sender, st.to().Address(), st.data, st.gas, st.value)
	}
	if vmerr != nil {

		if vmerr == vm.ErrInsufficientBalance {
			return nil, nil, nil, false, vmerr
		}
	}
	requiredGas = new(big.Int).Set(st.gasUsed())

	//退还Gas
	st.refundGas()

	return ret, requiredGas, st.gasUsed(), vmerr != nil, err
}

func (st *StateTransition) SystemBaseTransitionDb(dposContext *types.DposContext,
	bokerContext *types.BokerContext, boker bokerapi.Api) (ret []byte, requiredGas, usedGas *big.Int, failed bool, err error) {

	log.Info("(st *StateTransition) SystemBaseTransitionDb")
	if err = st.preCheck(); err != nil {
		return
	}

	intrinsicGas := IntrinsicGas(st.data, false, true)
	if intrinsicGas.BitLen() > 64 {

		return nil, nil, nil, false, vm.ErrOutOfGas
	}

	var (
		evm   = st.evm
		vmerr error
	)

	st.state.SetNonce(st.from().Address(), st.state.GetNonce(st.from().Address())+1)
	st.gas = protocol.MaxGasPrice.Uint64()
	ret, st.gas, vmerr = evm.Call(st.from(), st.to().Address(), st.data, st.gas, st.value)

	if vmerr != nil {
		log.Debug("VM returned with error", "err", vmerr)
		if vmerr == vm.ErrInsufficientBalance {
			return nil, nil, nil, false, vmerr
		}
	}

	return ret, new(big.Int).SetInt64(0), new(big.Int).SetInt64(0), vmerr != nil, err
}

func (st *StateTransition) UserBaseTransitionDb(dposContext *types.DposContext,
	bokerContext *types.BokerContext, boker bokerapi.Api) (ret []byte, requiredGas, usedGas *big.Int, failed bool, err error) {

	log.Info("(st *StateTransition) UserBaseTransitionDb")
	if err = st.preCheck(); err != nil {
		return
	}

	intrinsicGas := IntrinsicGas(st.data, false, true)
	if intrinsicGas.BitLen() > 64 {

		return nil, nil, nil, false, vm.ErrOutOfGas
	}

	var (
		evm   = st.evm
		vmerr error
	)

	st.state.SetNonce(st.from().Address(), st.state.GetNonce(st.from().Address())+1)
	st.gas = protocol.MaxGasPrice.Uint64()
	ret, st.gas, vmerr = evm.Call(st.from(), st.to().Address(), st.data, st.gas, st.value)

	if vmerr != nil {
		log.Debug("VM returned with error", "err", vmerr)
		if vmerr == vm.ErrInsufficientBalance {
			return nil, nil, nil, false, vmerr
		}
	}

	return ret, new(big.Int).SetInt64(0), new(big.Int).SetInt64(0), vmerr != nil, err
}

func (st *StateTransition) ExtraTransitionDb(dposContext *types.DposContext,
	bokerContext *types.BokerContext, boker bokerapi.Api) (ret []byte, requiredGas, usedGas *big.Int, failed bool, err error) {

	log.Info("(st *StateTransition) ExtraTransitionDb")
	if err = st.preCheck(); err != nil {
		return
	}

	msg := st.msg
	sender := st.from()
	homestead := true
	contractCreation := msg.To() == nil

	intrinsicGas := IntrinsicGas(st.data, contractCreation, homestead)
	if intrinsicGas.BitLen() > 64 {
		return nil, nil, nil, false, vm.ErrOutOfGas
	}

	if err = st.useGas(intrinsicGas.Uint64()); err != nil {
		return nil, nil, nil, false, err
	}

	var (
		evm   = st.evm
		vmerr error
	)

	if contractCreation {

		ret, _, st.gas, vmerr = evm.Create(sender, st.data, st.gas, st.value)
	} else {

		st.state.SetNonce(sender.Address(), st.state.GetNonce(sender.Address())+1)
		ret, st.gas, vmerr = evm.Call(sender, st.to().Address(), st.data, st.gas, st.value)
	}
	if vmerr != nil {

		if vmerr == vm.ErrInsufficientBalance {
			return nil, nil, nil, false, vmerr
		}
	}
	requiredGas = new(big.Int).Set(st.gasUsed())
	st.refundGas()

	return ret, requiredGas, st.gasUsed(), vmerr != nil, err
}

func (st *StateTransition) bokerAssignGas(operation common.Address, now *big.Int, dposContext *types.DposContext, bokerContext *types.BokerContext) (ret []byte, requiredGas, usedGas *big.Int, failed bool, err error) {

	log.Info("(st *StateTransition) bokerAssignGas", "operation", operation.String(), "now", now.Int64())

	address, err := dposContext.GetCurrentNowProducer(0, now.Int64())
	if err != nil {

		log.Error("(st *StateTransition) bokerAssignGas", "err", err)
		return []byte(""), new(big.Int).SetInt64(0), new(big.Int).SetInt64(0), false, err
	}

	if address != operation {

		log.Error("Operation Account Is not Producer", "operation", operation.String(), "producer", address.String())
		return []byte(""), new(big.Int).SetInt64(0), new(big.Int).SetInt64(0), false, errors.New("Operation Account Is not Producer")
	}

	//得到当前股权池中的Gas
	gas := bokerContext.GetGasPool()
	if gas <= 0 {

		log.Info("(st *StateTransition) bokerAssignGas Gas is Zero")
		return []byte(""), new(big.Int).SetInt64(0), new(big.Int).SetInt64(0), false, nil
	}

	//得到总股权数量
	var totalStock uint64 = 0
	stocks := bokerContext.GetStocks()
	for _, v := range stocks {

		if v.State != protocol.Frozen {
			totalStock = totalStock + v.Number
		}
	}

	//得到每个股份应该分配的数量
	singleGas := gas / totalStock

	//开始分币
	var curGas uint64 = 0
	for _, v := range stocks {

		userGas := singleGas * v.Number
		totalGas := curGas + userGas
		if totalGas > gas {
			userGas = gas - curGas
		}
		curGas = curGas + userGas

		//给用户加钱
		st.state.AddBalance(v.Account, new(big.Int).SetUint64(userGas))
	}
	return []byte(""), new(big.Int).SetInt64(0), new(big.Int).SetInt64(0), false, nil
}

func (st *StateTransition) StockTransitionDb(txMajor protocol.TxMajor,
	txMinor protocol.TxMinor,
	timer *big.Int,
	dposContext *types.DposContext,
	bokerContext *types.BokerContext,
	boker bokerapi.Api) (ret []byte, requiredGas, usedGas *big.Int, failed bool, err error) {

	log.Info("(st *StateTransition) StockTransitionDb", "txMajor", txMajor, "txMinor", txMinor)
	if err = st.preCheck(); err != nil {
		return
	}

	sender := st.from()
	st.state.SetNonce(st.from().Address(), st.state.GetNonce(st.from().Address())+1)
	if txMinor == protocol.StockManager {

		bokerContext.SetStockManager(sender.Address(), *st.msg.To(), timer.Int64(), dposContext)
	} else if txMinor == protocol.StockSet {

		bokerContext.SetStock(sender.Address(), *st.msg.To(), st.msg.Value().Uint64())
	} else if txMinor == protocol.StockTransfer {

		bokerContext.TransferStock(sender.Address(), *st.msg.To(), st.msg.Value().Uint64())
	} else if txMinor == protocol.StockClean {

		bokerContext.CleanStock(sender.Address(), *st.msg.To())
	} else if txMinor == protocol.StockFrozen {

		bokerContext.FrozenStock(sender.Address(), *st.msg.To())
	} else if txMinor == protocol.StockUnFrozen {

		bokerContext.UnFrozenStock(sender.Address(), *st.msg.To())
	} else if txMinor == protocol.StockAssignGas {

		st.bokerAssignGas(sender.Address(), timer, dposContext, bokerContext)
	} else {
		return []byte(""), new(big.Int).SetInt64(0), new(big.Int).SetInt64(0), false, errors.New("not know protocol type")
	}

	return []byte(""), new(big.Int).SetInt64(0), new(big.Int).SetInt64(0), false, nil
}

func (st *StateTransition) SystemContractTransitionDb(txMajor protocol.TxMajor,
	txMinor protocol.TxMinor,
	dposContext *types.DposContext,
	bokerContext *types.BokerContext,
	boker bokerapi.Api) (ret []byte, requiredGas, usedGas *big.Int, failed bool, err error) {

	log.Info("(st *StateTransition) SystemContractTransitionDb")
	if err = st.preCheck(); err != nil {
		return
	}

	if txMinor == protocol.SetSystemContract {
		boker.SetSystemContract(*st.msg.To(), st.from().Address(), bokerContext)
	}

	st.state.SetNonce(st.from().Address(), st.state.GetNonce(st.from().Address())+1)
	return []byte(""), new(big.Int).SetInt64(0), new(big.Int).SetInt64(0), false, nil
}

func (st *StateTransition) UserContractTransitionDb(txMajor protocol.TxMajor,
	txMinor protocol.TxMinor,
	dposContext *types.DposContext,
	bokerContext *types.BokerContext,
	boker bokerapi.Api) (ret []byte, requiredGas, usedGas *big.Int, failed bool, err error) {

	log.Info("(st *StateTransition) UserContractTransitionDb")
	if err = st.preCheck(); err != nil {
		return
	}

	if txMinor == protocol.SetSystemContract {
		bokerContext.SetUserBaseContract(*st.msg.To(), st.from().Address())
	}

	st.state.SetNonce(st.from().Address(), st.state.GetNonce(st.from().Address())+1)
	return []byte(""), new(big.Int).SetInt64(0), new(big.Int).SetInt64(0), false, nil
}

func (st *StateTransition) VoteTransitionDb(txMajor protocol.TxMajor, txMinor protocol.TxMinor, boker bokerapi.Api) (ret []byte, requiredGas, usedGas *big.Int, failed bool, err error) {

	log.Info("(st *StateTransition) VoteTransitionDb")
	if err = st.preCheck(); err != nil {
		return
	}
	st.state.SetNonce(st.from().Address(), st.state.GetNonce(st.from().Address())+1)
	return []byte(""), new(big.Int).SetInt64(0), new(big.Int).SetInt64(0), false, nil
}

func (st *StateTransition) ValidatorTransitionDb(txMajor protocol.TxMajor, txMinor protocol.TxMinor, boker bokerapi.Api) (ret []byte, requiredGas, usedGas *big.Int, failed bool, err error) {

	log.Info("(st *StateTransition) ValidatorTransitionDb")
	if err = st.preCheck(); err != nil {
		return
	}
	st.state.SetNonce(st.from().Address(), st.state.GetNonce(st.from().Address())+1)
	return []byte(""), new(big.Int).SetInt64(0), new(big.Int).SetInt64(0), false, nil
}

//退还Gas
func (st *StateTransition) refundGas() {

	sender := st.from()
	remaining := new(big.Int).Mul(new(big.Int).SetUint64(st.gas), st.gasPrice)
	st.state.AddBalance(sender.Address(), remaining)

	uhalf := remaining.Div(st.gasUsed(), common.Big2)
	refund := math.BigMin(uhalf, st.state.GetRefund())
	st.gas += refund.Uint64()

	st.state.AddBalance(sender.Address(), refund.Mul(refund, st.gasPrice))

	//将剩余的Gas归还给发送方
	st.gp.AddGas(new(big.Int).SetUint64(st.gas))
}

func (st *StateTransition) gasUsed() *big.Int {
	return new(big.Int).Sub(st.initialGas, new(big.Int).SetUint64(st.gas))
}
