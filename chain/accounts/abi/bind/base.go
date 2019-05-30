package bind

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	_ "reflect"

	"github.com/Tinachain/Tina/chain"
	"github.com/Tinachain/Tina/chain/accounts/abi"
	"github.com/Tinachain/Tina/chain/boker/protocol"
	"github.com/Tinachain/Tina/chain/common"
	"github.com/Tinachain/Tina/chain/core/types"
	"github.com/Tinachain/Tina/chain/crypto"
	"github.com/Tinachain/Tina/chain/eth"
	"github.com/Tinachain/Tina/chain/log"
	"github.com/Tinachain/Tina/chain/node"
)

// SignerFn is a signer function callback when a contract requires a method to
// sign the transaction before submission.
type SignerFn func(types.Signer, common.Address, *types.Transaction) (*types.Transaction, error)

// CallOpts is the collection of options to fine tune a contract call request.
type CallOpts struct {
	Pending bool            // Whether to operate on the pending state or the last known one
	From    common.Address  // Optional the sender address, otherwise the first account is used
	Context context.Context // Network context to support cancellation and timeouts (nil = no timeout)
}

//创建一个有效的以太坊交易
type TransactOpts struct {
	From     common.Address  // Ethereum account to send the transaction from
	Nonce    *big.Int        // Nonce to use for the transaction execution (nil = use pending state)
	Signer   SignerFn        // Method to use for signing the transaction (mandatory)
	Value    *big.Int        // Funds to transfer along along the transaction (nil = 0 = no funds)
	GasPrice *big.Int        // Gas price to use for the transaction execution (nil = gas price oracle)
	GasLimit *big.Int        // Gas limit to set for the transaction execution (nil = estimate + 10%)
	Context  context.Context // Network context to support cancellation and timeouts (nil = no timeout)
}

//BoundContract定义以太坊合约的基础包装器对象 它包含一组由方法使用的方法更高级别的合同绑定操作。
type BoundContract struct {
	address    common.Address     // Deployment address of the contract on the Ethereum blockchain
	abi        abi.ABI            // Reflect based ABI to access the correct Ethereum methods
	caller     ContractCaller     // Read interface to interact with the blockchain
	transactor ContractTransactor // Write interface to interact with the blockchain
}

var GethNode *node.Node

//NewBoundContract 创建一个通过其调用的低级合约接口并且交易可以通过。
func NewBoundContract(address common.Address,
	abi abi.ABI,
	caller ContractCaller,
	transactor ContractTransactor) *BoundContract {
	return &BoundContract{
		address:    address,
		abi:        abi,
		caller:     caller,
		transactor: transactor,
	}
}

func DeployContract(opts *TransactOpts, abi abi.ABI, bytecode []byte, backend ContractBackend, params ...interface{}) (common.Address, *types.Transaction, *BoundContract, error) {

	log.Info("base.go DeployContract")

	//赋值
	c := NewBoundContract(common.Address{}, abi, backend, backend)

	input, err := c.abi.Pack("", params...)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	tx, err := c.transact(opts, nil, append(bytecode, input...), []byte(""), protocol.Normal, 0)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	c.address = crypto.CreateAddress(opts.From, tx.Nonce())
	return c.address, tx, c, nil
}

func SetBaseContract(opts *TransactOpts, abiJson string, address common.Address, backend ContractBackend) (*types.Transaction, error) {

	log.Info("base.go SetBaseContract")

	var abi abi.ABI
	abiErr := abi.UnmarshalJSON([]byte(abiJson))
	if abiErr != nil {
		return nil, errors.New("SetBaseContract UnmarshalJSON Failed")
	}

	c := NewBoundContract(address, abi, backend, backend)
	tx, err := c.transact(opts, &address, []byte(""), []byte(abiJson), protocol.Base, protocol.SetSystemContract)
	if err != nil {
		return nil, err
	}
	return tx, err
}

//调用合约方法，并将params作为输入值和将输出设置为result
func (c *BoundContract) Call(opts *CallOpts, result interface{}, method string, params ...interface{}) error {

	//判断opts是否为空
	if opts == nil {
		opts = new(CallOpts)
	}

	//打包输入，调用并解压缩结果
	input, err := c.abi.Pack(method, params...)
	if err != nil {
		return err
	}

	var (
		msg    = ethereum.CallMsg{From: opts.From, To: &c.address, Data: input}
		ctx    = ensureContext(opts.Context)
		code   []byte
		output []byte
	)

	if opts.Pending {

		pb, ok := c.caller.(PendingContractCaller)
		if !ok {
			return ErrNoPendingState
		}

		output, err = pb.PendingCallContract(ctx, msg)
		if err == nil && len(output) == 0 {
			// Make sure we have a contract to operate on, and bail out otherwise.
			if code, err = pb.PendingCodeAt(ctx, c.address); err != nil {
				return err
			} else if len(code) == 0 {
				return ErrNoCode
			}
		}

	} else {

		output, err = c.caller.CallContract(ctx, msg, nil)
		if err == nil && len(output) == 0 {

			log.Info("Call", "outputlength", len(output))

			// Make sure we have a contract to operate on, and bail out otherwise.
			if code, err = c.caller.CodeAt(ctx, c.address, nil); err != nil {
				return err
			} else if len(code) == 0 {
				return ErrNoCode
			}
		}

	}
	if err != nil {
		return err
	}
	return c.abi.Unpack(result, method, output)
}

//得到当前分币帐号
func (c *BoundContract) getTokenNoder(opts *TransactOpts) (common.Address, error) {

	var ether *eth.Ethereum
	if err := GethNode.Service(&ether); err != nil {
		return common.Address{}, err
	}

	if ether.BlockChain().CurrentBlock() == nil {
		return common.Address{}, errors.New("failed to lookup token node")
	}

	firstTimer := ether.BlockChain().GetBlockByNumber(0).Time().Int64()
	return ether.BlockChain().CurrentBlock().DposCtx().GetCurrentTokenNoder(firstTimer)
}

//得到当前的验证者帐号
func (c *BoundContract) getProducer(opts *TransactOpts) (common.Address, error) {

	var ether *eth.Ethereum
	if err := GethNode.Service(&ether); err != nil {
		return common.Address{}, err
	}

	if ether.BlockChain().CurrentBlock() == nil {
		return common.Address{}, errors.New("failed to lookup token node")
	}

	firstTimer := ether.BlockChain().GetBlockByNumber(0).Time().Int64()
	return ether.BlockChain().CurrentBlock().DposCtx().GetCurrentProducer(firstTimer)
}

//得到参数类型
func typeof(v interface{}) bool {

	switch v.(type) {
	case string:
		return true
	default:
		return false
	}
}

//使用输入的值作为参数调用合约方法
func (c *BoundContract) Transact(opts *TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {

	log.Info("Create Transact", "method", method)

	//打包合约参数
	input, err := c.abi.Pack(method, params...)
	if err != nil {
		return nil, err
	}

	//判断节点是否已经启动
	if GethNode != nil {

		var e *eth.Ethereum
		if err := GethNode.Service(&e); err != nil {
			return nil, err
		}

		//得到合约类型
		txMajor, err := e.Boker().GetContract(c.address)
		if err != nil {
			return nil, err
		}

		//判断合约类型是否是基础合约
		extra := []byte("")
		if txMajor == protocol.Normal {

			return c.transact(opts, &c.address, input, []byte(""), protocol.Normal, 0)

		} else if txMajor == protocol.Base {

			//用户触发的基础合约（用户触发，但是不收取Gas费用）
			if method == protocol.RegisterCandidateMethod {
				return c.transact(opts, &c.address, input, extra, protocol.Base, protocol.RegisterCandidate)

			} else if method == protocol.VoteCandidateMethod {
				return c.transact(opts, &c.address, input, extra, protocol.Base, protocol.VoteUser)

			} else if method == protocol.CancelVoteMethod {
				return c.transact(opts, &c.address, input, extra, protocol.Base, protocol.VoteCancel)

			} else if method == protocol.RotateVoteMethod {

				//得到当前的分币节点
				tokenNoder, err := c.getTokenNoder(opts)
				if err != nil {
					return nil, errors.New("get rotate vote error")
				}
				if tokenNoder != opts.From {
					return nil, errors.New("current rotate vote not is from account")
				}
				return c.transact(opts, &c.address, input, extra, protocol.Base, protocol.VoteEpoch)
			}
			return nil, errors.New("unknown system contract method name")
		}
	}
	return c.transact(opts, &c.address, input, []byte(""), protocol.Normal, 0)
}

func (c *BoundContract) Transfer(opts *TransactOpts) (*types.Transaction, error) {

	log.Info("****Transfer****")

	var e *eth.Ethereum
	if err := GethNode.Service(&e); err != nil {
		return nil, err
	}

	txMajor, err := e.Boker().GetContract(c.address)
	if err != nil {
		return nil, err
	}

	return c.transact(opts, &c.address, nil, []byte(""), txMajor, 0)
}

func (c *BoundContract) baseTransact(opts *TransactOpts,
	contract *common.Address,
	payload []byte,
	extra []byte,
	txMajor protocol.TxMajor,
	txMinor protocol.TxMinor) (*types.Transaction, error) {

	//判断Value值是否为空
	var err error
	value := opts.Value
	if value == nil {
		value = new(big.Int)
	}

	//判断Nonce值是否为空
	var nonce uint64
	if opts.Nonce == nil {

		//如果Nonce值为空，则初始化一个nonce值来进行初始化
		nonce, err = c.transactor.PendingNonceAt(ensureContext(opts.Context), opts.From)
		if err != nil {
			return nil, fmt.Errorf("failed to retrieve account nonce: %v", err)
		}
	} else {
		nonce = opts.Nonce.Uint64()
	}
	//log.Info("baseTransact", "nonce", nonce)

	/*这里不对Gas和GasLimit进行设置，因为在函数NewBaseTransaction里面已经针对基础业务进行了设置*/
	var rawTx *types.Transaction
	if contract == nil {

		//如果合约尚未创建，则创建合约
		//rawTx = types.NewBaseContractCreation(nonce, value, input)
		return nil, errors.New("not found base contract address")
	} else {

		//合约已经创建，则创建一个交易
		rawTx = types.NewBaseTransaction(txMajor, txMinor, nonce, c.address, value, payload)
	}

	//判断交易是否有签名者
	if opts.Signer == nil {
		return nil, errors.New("no signer to authorize the transaction with")
	}

	//进行签名
	signedTx, err := opts.Signer(types.HomesteadSigner{}, opts.From, rawTx)
	if err != nil {
		return nil, err
	}

	//将交易注入pending池中
	if err := c.transactor.SendTransaction(ensureContext(opts.Context), signedTx); err != nil {
		return nil, err
	}
	return signedTx, nil
}

func (c *BoundContract) normalTransact(opts *TransactOpts,
	contract *common.Address,
	payload []byte,
	extra []byte,
	txMajor protocol.TxMajor,
	txMinor protocol.TxMinor) (*types.Transaction, error) {

	//判断Value值是否为空
	var err error
	value := opts.Value
	if value == nil {
		value = new(big.Int)
	}

	//判断Nonce值是否为空
	var nonce uint64
	if opts.Nonce == nil {

		//如果Nonce值为空，则初始化一个nonce值来进行初始化
		nonce, err = c.transactor.PendingNonceAt(ensureContext(opts.Context), opts.From)
		if err != nil {
			return nil, fmt.Errorf("failed to retrieve account nonce: %v", err)
		}
	} else {
		nonce = opts.Nonce.Uint64()
	}

	//如果GasPrice为空，则设置一个建议的GasPrice
	gasPrice := opts.GasPrice
	if gasPrice == nil {
		gasPrice, err = c.transactor.SuggestGasPrice(ensureContext(opts.Context)) //得到一个建议的GasPrice
		if err != nil {
			return nil, fmt.Errorf("failed to suggest gas price: %v", err)
		}
	}

	//如果GasLimit为空，则设置一个GasLimit
	gasLimit := opts.GasLimit
	if gasLimit == nil {

		//如果合约存在，则根据合约内容评估一个GasLimit
		if contract != nil {

			if code, err := c.transactor.PendingCodeAt(ensureContext(opts.Context), c.address); err != nil {
				return nil, err
			} else if len(code) == 0 {
				return nil, ErrNoCode
			}
		}

		//估算所需要的Gas
		msg := ethereum.CallMsg{From: opts.From, To: contract, Value: value, Data: payload, Extra: extra}
		gasLimit, err = c.transactor.EstimateGas(ensureContext(opts.Context), msg)
		if err != nil {
			return nil, fmt.Errorf("failed to estimate gas needed: %v", err) //估算所需gas失败
		}
	}

	//创建合约交易或者直接产生一个交易
	var rawTx *types.Transaction
	if contract == nil {
		//如果合约尚未创建，则创建合约
		rawTx = types.NewContractCreation(nonce, value, gasLimit, gasPrice, payload)
	} else {
		//合约已经创建，则创建一个交易
		rawTx = types.NewTransaction(txMajor, txMinor, nonce, c.address, value, gasLimit, gasPrice, payload)
	}

	//判断交易是否有签名者
	if opts.Signer == nil {
		return nil, errors.New("no signer to authorize the transaction with")
	}

	//进行签名
	signedTx, err := opts.Signer(types.HomesteadSigner{}, opts.From, rawTx)
	if err != nil {
		return nil, err
	}

	//将交易注入pending池中
	if err := c.transactor.SendTransaction(ensureContext(opts.Context), signedTx); err != nil {
		return nil, err
	}
	return signedTx, nil
}

func (c *BoundContract) transact(opts *TransactOpts,
	contract *common.Address,
	payload []byte,
	extra []byte,
	txMajor protocol.TxMajor,
	txMinor protocol.TxMinor) (*types.Transaction, error) {

	/*根据不同类型计算使用的Gas信息*/
	if txMajor == protocol.Normal {

		//普通交易
		return c.normalTransact(opts, contract, payload, extra, protocol.Normal, 0)
	} else if txMajor == protocol.Base {

		if (txMinor >= protocol.MinMinor) && (txMinor <= protocol.MaxMinor) {
			return c.baseTransact(opts, contract, payload, extra, txMajor, txMinor)
		}
		log.Error("transact", "major", txMajor, "minor", txMinor)
		return nil, errors.New("unknown transaction type")
	} else if txMajor == protocol.Extra {
		return nil, errors.New("unknown transaction type")
	}
	return nil, errors.New("unknown transaction type")
}

//
func ensureContext(ctx context.Context) context.Context {
	if ctx == nil {
		return context.TODO()
	}
	return ctx
}
