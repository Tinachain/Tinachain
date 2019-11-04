package ethclient

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"

	"github.com/Tinachain/Tina/chain"
	"github.com/Tinachain/Tina/chain/boker/protocol"
	"github.com/Tinachain/Tina/chain/common"
	"github.com/Tinachain/Tina/chain/common/hexutil"
	"github.com/Tinachain/Tina/chain/core/types"
	"github.com/Tinachain/Tina/chain/log"
	"github.com/Tinachain/Tina/chain/rlp"
	"github.com/Tinachain/Tina/chain/rpc"
)

// Client defines typed wrappers for the Ethereum RPC API.
type Client struct {
	c *rpc.Client
}

// Dial connects a client to the given URL.
func Dial(rawurl string) (*Client, error) {
	c, err := rpc.Dial(rawurl)
	if err != nil {
		return nil, err
	}
	return NewClient(c), nil
}

// NewClient creates a client that uses the given RPC client.
func NewClient(c *rpc.Client) *Client {
	return &Client{c}
}

// Blockchain Access

//根据Hash得到区块信息（远程调用的是internal/api 里面的GetBlockByHash函数）
func (ec *Client) BlockByHash(ctx context.Context, hash common.Hash) (*types.Block, error) {
	return ec.getBlock(ctx, "eth_getBlockByHash", hash, true)
}

//根据区块序号得到区块信息（远程调用的是internal/api 里面的GetBlockByNumber函数）,如果number为nil则返回最后的区块
func (ec *Client) BlockByNumber(ctx context.Context, number *big.Int) (*types.Block, error) {
	return ec.getBlock(ctx, "eth_getBlockByNumber", toBlockNumArg(number), true)
}

type rpcBlock struct {
	Hash         common.Hash      `json:"hash"`
	Transactions []rpcTransaction `json:"transactions"`
	UncleHashes  []common.Hash    `json:"uncles"`
}

func (ec *Client) getBlock(ctx context.Context, method string, args ...interface{}) (*types.Block, error) {

	fmt.Println("(ec *Client) getBlock method=%s", method)
	var raw json.RawMessage
	err := ec.c.CallContext(ctx, &raw, method, args...)
	if err != nil {
		return nil, err
	} else if len(raw) == 0 {
		return nil, ethereum.NotFound
	}
	fmt.Println("ec.c.CallContext", "raw", string(raw))

	//解码交易头和交易体
	var head *types.Header
	var body rpcBlock
	if err := json.Unmarshal(raw, &head); err != nil {

		fmt.Println("ec.c.CallContext json.Unmarshal failed err=%s", err.Error())
		return nil, err
	}
	//fmt.Println("ec.c.CallContext", "header", head)

	if err := json.Unmarshal(raw, &body); err != nil {
		return nil, err
	}
	// Quick-verify transaction and uncle lists. This mostly helps with debugging the server.
	if head.UncleHash == types.EmptyUncleHash && len(body.UncleHashes) > 0 {
		return nil, fmt.Errorf("server returned non-empty uncle list but block header indicates no uncles")
	}
	if head.UncleHash != types.EmptyUncleHash && len(body.UncleHashes) == 0 {
		return nil, fmt.Errorf("server returned empty uncle list but block header indicates uncles")
	}
	if head.TxHash == types.EmptyRootHash && len(body.Transactions) > 0 {
		return nil, fmt.Errorf("server returned non-empty transaction list but block header indicates no transactions")
	}
	if head.TxHash != types.EmptyRootHash && len(body.Transactions) == 0 {
		return nil, fmt.Errorf("server returned empty transaction list but block header indicates transactions")
	}
	// Load uncles because they are not included in the block response.
	//加载叔块因为叔块不包含区块返回
	var uncles []*types.Header
	if len(body.UncleHashes) > 0 {
		uncles = make([]*types.Header, len(body.UncleHashes))
		reqs := make([]rpc.BatchElem, len(body.UncleHashes))
		for i := range reqs {
			reqs[i] = rpc.BatchElem{
				Method: "eth_getUncleByBlockHashAndIndex",
				Args:   []interface{}{body.Hash, hexutil.EncodeUint64(uint64(i))},
				Result: &uncles[i],
			}
		}
		if err := ec.c.BatchCallContext(ctx, reqs); err != nil {
			return nil, err
		}
		for i := range reqs {
			if reqs[i].Error != nil {
				return nil, reqs[i].Error
			}
			if uncles[i] == nil {
				return nil, fmt.Errorf("got null header for uncle %d of block %x", i, body.Hash[:])
			}
		}
	}
	// Fill the sender cache of transactions in the block.
	txs := make([]*types.Transaction, len(body.Transactions))
	for i, tx := range body.Transactions {
		setSenderFromServer(tx.tx, tx.From, body.Hash)
		txs[i] = tx.tx
	}
	return types.NewBlockWithHeader(head).WithBody(txs, uncles), nil
}

//通过Hash得到区块头
func (ec *Client) HeaderByHash(ctx context.Context, hash common.Hash) (*types.Header, error) {
	var head *types.Header
	err := ec.c.CallContext(ctx, &head, "eth_getBlockByHash", hash, false)
	if err == nil && head == nil {
		err = ethereum.NotFound
	}
	return head, err
}

//根据区块序号，得到区块头，如果number为空，则返回最后的区块
func (ec *Client) HeaderByNumber(ctx context.Context, number *big.Int) (*types.Header, error) {
	var head *types.Header
	err := ec.c.CallContext(ctx, &head, "eth_getBlockByNumber", toBlockNumArg(number), false)
	if err == nil && head == nil {
		err = ethereum.NotFound
	}
	return head, err
}

type rpcTransaction struct {
	tx *types.Transaction
	txExtraInfo
}

type txExtraInfo struct {
	BlockNumber *string
	BlockHash   common.Hash
	From        common.Address
}

func (tx *rpcTransaction) UnmarshalJSON(msg []byte) error {
	if err := json.Unmarshal(msg, &tx.tx); err != nil {
		return err
	}
	return json.Unmarshal(msg, &tx.txExtraInfo)
}

//根据所给的hash返回相应的交易
func (ec *Client) TransactionByHash(ctx context.Context, hash common.Hash) (tx *types.Transaction, isPending bool, err error) {

	var json *rpcTransaction
	err = ec.c.CallContext(ctx, &json, "eth_getTransactionByHash", hash)
	if err != nil {
		return nil, false, err
	} else if json == nil {
		return nil, false, ethereum.NotFound
	} else if _, r, _ := json.tx.RawSignatureValues(); r == nil {
		return nil, false, fmt.Errorf("server returned transaction without signature")
	}
	setSenderFromServer(json.tx, json.From, json.BlockHash)

	return json.tx, json.BlockNumber == nil, nil
}

// TransactionSender returns the sender address of the given transaction. The transaction
// must be known to the remote node and included in the blockchain at the given block and
// index. The sender is the one derived by the protocol at the time of inclusion.
//
// There is a fast-path for transactions retrieved by TransactionByHash and
// TransactionInBlock. Getting their sender address can be done without an RPC interaction.
func (ec *Client) TransactionSender(ctx context.Context, tx *types.Transaction, block common.Hash, index uint) (common.Address, error) {

	// Try to load the address from the cache.
	sender, err := types.Sender(&senderFromServer{blockhash: block}, tx)
	if err == nil {
		return sender, nil
	}
	var meta struct {
		Hash common.Hash
		From common.Address
	}
	if err = ec.c.CallContext(ctx, &meta, "eth_getTransactionByBlockHashAndIndex", block, hexutil.Uint64(index)); err != nil {
		return common.Address{}, err
	}
	if meta.Hash == (common.Hash{}) || meta.Hash != tx.Hash() {
		return common.Address{}, errors.New("wrong inclusion block/index")
	}
	return meta.From, nil
}

//根据给的区块，获取这个区块中的交易数量
func (ec *Client) TransactionCount(ctx context.Context, blockHash common.Hash) (uint, error) {
	var num hexutil.Uint
	err := ec.c.CallContext(ctx, &num, "eth_getBlockTransactionCountByHash", blockHash)
	return uint(num), err
}

//根据给定的区块，获取这个区块中指定序号的交易
func (ec *Client) TransactionInBlock(ctx context.Context, blockHash common.Hash, index uint) (*types.Transaction, error) {

	var json *rpcTransaction
	err := ec.c.CallContext(ctx, &json, "eth_getTransactionByBlockHashAndIndex", blockHash, hexutil.Uint64(index))
	if err == nil {
		if json == nil {
			return nil, ethereum.NotFound
		} else if _, r, _ := json.tx.RawSignatureValues(); r == nil {
			return nil, fmt.Errorf("server returned transaction without signature")
		}
	}
	setSenderFromServer(json.tx, json.From, json.BlockHash)
	return json.tx, err
}

//根据给定的Hash得到交易收据(注意的是，收据不适合用于正在penging的交易)
func (ec *Client) TransactionReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error) {
	var r *types.Receipt
	err := ec.c.CallContext(ctx, &r, "eth_getTransactionReceipt", txHash)
	if err == nil {
		if r == nil {
			return nil, ethereum.NotFound
		}
	}
	return r, err
}

func toBlockNumArg(number *big.Int) string {
	if number == nil {
		return "latest"
	}
	return hexutil.EncodeBig(number)
}

type rpcProgress struct {
	StartingBlock hexutil.Uint64
	CurrentBlock  hexutil.Uint64
	HighestBlock  hexutil.Uint64
	PulledStates  hexutil.Uint64
	KnownStates   hexutil.Uint64
}

//检索同步算法当前的进度，如果当前没有同步操作，则返回nil
func (ec *Client) SyncProgress(ctx context.Context) (*ethereum.SyncProgress, error) {

	var raw json.RawMessage
	if err := ec.c.CallContext(ctx, &raw, "eth_syncing"); err != nil {
		return nil, err
	}
	// Handle the possible response types
	var syncing bool
	if err := json.Unmarshal(raw, &syncing); err == nil {
		return nil, nil // Not syncing (always false)
	}
	var progress *rpcProgress
	if err := json.Unmarshal(raw, &progress); err != nil {
		return nil, err
	}
	return &ethereum.SyncProgress{
		StartingBlock: uint64(progress.StartingBlock),
		CurrentBlock:  uint64(progress.CurrentBlock),
		HighestBlock:  uint64(progress.HighestBlock),
		PulledStates:  uint64(progress.PulledStates),
		KnownStates:   uint64(progress.KnownStates),
	}, nil
}

//订阅区块链新链头变化通知
func (ec *Client) SubscribeNewHead(ctx context.Context, ch chan<- *types.Header) (ethereum.Subscription, error) {
	return ec.c.EthSubscribe(ctx, ch, "newHeads", map[string]struct{}{})
}

// State Access

//返回此链的网络ID
func (ec *Client) NetworkID(ctx context.Context) (*big.Int, error) {
	version := new(big.Int)
	var ver string
	if err := ec.c.CallContext(ctx, &ver, "net_version"); err != nil {
		return nil, err
	}
	if _, ok := version.SetString(ver, 10); !ok {
		return nil, fmt.Errorf("invalid net_version result %q", ver)
	}
	return version, nil
}

//返回指定账户的余额（单位wei），如果是nil则从最新的块中获取
func (ec *Client) BalanceAt(ctx context.Context, account common.Address, blockNumber *big.Int) (*big.Int, error) {
	var result hexutil.Big
	err := ec.c.CallContext(ctx, &result, "eth_getBalance", account, toBlockNumArg(blockNumber))
	return (*big.Int)(&result), err
}

// StorageAt returns the value of key in the contract storage of the given account.
// The block number can be nil, in which case the value is taken from the latest known block.
func (ec *Client) StorageAt(ctx context.Context, account common.Address, key common.Hash, blockNumber *big.Int) ([]byte, error) {
	var result hexutil.Bytes
	err := ec.c.CallContext(ctx, &result, "eth_getStorageAt", account, key, toBlockNumArg(blockNumber))
	return result, err
}

// CodeAt returns the contract code of the given account.
// The block number can be nil, in which case the code is taken from the latest known block.
func (ec *Client) CodeAt(ctx context.Context, account common.Address, blockNumber *big.Int) ([]byte, error) {
	var result hexutil.Bytes
	err := ec.c.CallContext(ctx, &result, "eth_getCode", account, toBlockNumArg(blockNumber))
	return result, err
}

//从给定的区块序号中获取指定账号的Nonce，如果区块序号为nil则从最新的区块中获取
func (ec *Client) NonceAt(ctx context.Context, account common.Address, blockNumber *big.Int) (uint64, error) {
	var result hexutil.Uint64
	err := ec.c.CallContext(ctx, &result, "eth_getTransactionCount", account, toBlockNumArg(blockNumber))
	return uint64(result), err
}

// Filters

//日志过滤器
func (ec *Client) FilterLogs(ctx context.Context, q ethereum.FilterQuery) ([]types.Log, error) {
	var result []types.Log
	err := ec.c.CallContext(ctx, &result, "eth_getLogs", toFilterArg(q))
	return result, err
}

//订阅日志过滤器
func (ec *Client) SubscribeFilterLogs(ctx context.Context, q ethereum.FilterQuery, ch chan<- types.Log) (ethereum.Subscription, error) {
	return ec.c.EthSubscribe(ctx, ch, "logs", toFilterArg(q))
}

func toFilterArg(q ethereum.FilterQuery) interface{} {
	arg := map[string]interface{}{
		"fromBlock": toBlockNumArg(q.FromBlock),
		"toBlock":   toBlockNumArg(q.ToBlock),
		"address":   q.Addresses,
		"topics":    q.Topics,
	}
	if q.FromBlock == nil {
		arg["fromBlock"] = "0x0"
	}
	return arg
}

// Pending State

//返回处于Pending状态的账户的余额（单位：Wei）
func (ec *Client) PendingBalanceAt(ctx context.Context, account common.Address) (*big.Int, error) {
	var result hexutil.Big
	err := ec.c.CallContext(ctx, &result, "eth_getBalance", account, "pending")
	return (*big.Int)(&result), err
}

// PendingStorageAt returns the value of key in the contract storage of the given account in the pending state.
func (ec *Client) PendingStorageAt(ctx context.Context, account common.Address, key common.Hash) ([]byte, error) {
	var result hexutil.Bytes
	err := ec.c.CallContext(ctx, &result, "eth_getStorageAt", account, key, "pending")
	return result, err
}

// PendingCodeAt returns the contract code of the given account in the pending state.
func (ec *Client) PendingCodeAt(ctx context.Context, account common.Address) ([]byte, error) {
	var result hexutil.Bytes
	err := ec.c.CallContext(ctx, &result, "eth_getCode", account, "pending")
	return result, err
}

// PendingNonceAt returns the account nonce of the given account in the pending state.
// This is the nonce that should be used for the next transaction.
func (ec *Client) PendingNonceAt(ctx context.Context, account common.Address) (uint64, error) {
	var result hexutil.Uint64
	err := ec.c.CallContext(ctx, &result, "eth_getTransactionCount", account, "pending")
	return uint64(result), err
}

//当前全链pinding交易数量
func (ec *Client) PendingTransactionCount(ctx context.Context) (uint, error) {
	var num hexutil.Uint
	err := ec.c.CallContext(ctx, &num, "eth_getBlockTransactionCountByNumber", "pending")
	return uint(num), err
}

//合约调用，执行消息调用交易，该交易直接在VM中执行节点，但从未开采过区块链。
//blockNumber选择调用运行的块高度。 它可以是零，其中代码取自最新的已知块。 注意从很老的状态块可能不可用。
func (ec *Client) CallContract(ctx context.Context, msg ethereum.CallMsg, blockNumber *big.Int) ([]byte, error) {

	log.Info("(ec *Client) CallContract")

	var hex hexutil.Bytes
	err := ec.c.CallContext(ctx, &hex, "eth_call", toCallArg(msg), toBlockNumArg(blockNumber))
	if err != nil {
		return nil, err
	}

	return hex, nil
}

func (ec *Client) PendingCallContract(ctx context.Context, msg ethereum.CallMsg) ([]byte, error) {
	var hex hexutil.Bytes
	err := ec.c.CallContext(ctx, &hex, "eth_call", toCallArg(msg), "pending")
	if err != nil {
		return nil, err
	}
	return hex, nil
}

func (ec *Client) SuggestGasPrice(ctx context.Context) (*big.Int, error) {
	var hex hexutil.Big
	if err := ec.c.CallContext(ctx, &hex, "eth_gasPrice"); err != nil {
		return nil, err
	}
	return (*big.Int)(&hex), nil
}

func (ec *Client) EstimateGas(ctx context.Context, msg ethereum.CallMsg) (*big.Int, error) {
	var hex hexutil.Big
	err := ec.c.CallContext(ctx, &hex, "eth_estimateGas", toCallArg(msg))
	if err != nil {
		return nil, err
	}
	return (*big.Int)(&hex), nil
}

func (ec *Client) SendTransaction(ctx context.Context, tx *types.Transaction) error {

	data, err := rlp.EncodeToBytes(tx)
	if err != nil {
		return err
	}
	log.Info("(ec *Client) SendTransaction", "len", len(data), "data", data)

	tx_tmp := new(types.Transaction)
	if err := rlp.DecodeBytes(data, tx_tmp); err != nil {

		log.Error("SendTransaction", "error", err, "data", data)
		return err
	}

	return ec.c.CallContext(ctx, nil, "eth_sendRawTransaction", common.ToHex(data))
}

func toCallArg(msg ethereum.CallMsg) interface{} {

	arg := map[string]interface{}{
		"from": msg.From,
		"to":   msg.To,
	}
	if len(msg.Data) > 0 {
		arg["data"] = hexutil.Bytes(msg.Data)
	}
	if msg.Value != nil {
		arg["value"] = (*hexutil.Big)(msg.Value)
	}
	if msg.Gas != nil {
		arg["gas"] = (*hexutil.Big)(msg.Gas)
	}
	if msg.GasPrice != nil {
		arg["gasPrice"] = (*hexutil.Big)(msg.GasPrice)
	}
	if len(msg.Extra) > 0 {
		arg["extra"] = hexutil.Bytes(msg.Extra)
	}
	return arg
}

func (ec *Client) GetLastProducerAt(ctx context.Context) ([]byte, error) {

	var result hexutil.Bytes
	err := ec.c.CallContext(ctx, &result, "eth_getLastProducer")
	return result, err
}

func (ec *Client) GetNextProducerAt(ctx context.Context) ([]byte, error) {

	var result hexutil.Bytes
	err := ec.c.CallContext(ctx, &result, "eth_getNextProducer")
	return result, err
}

func (ec *Client) SetSystemBaseContracts(ctx context.Context, address common.Address) (common.Hash, error) {

	log.Info("(ec *Client) SetSystemBaseContracts", "address", address.String())

	var txHash common.Hash
	err := ec.c.CallContext(ctx, &txHash, "eth_setSystemBaseContracts", address)
	if err != nil {
		return common.Hash{}, err
	}

	return txHash, nil
}

func (ec *Client) SetUserBaseContracts(ctx context.Context, address common.Address) (common.Hash, error) {

	log.Info("(ec *Client) SetUserBaseContracts", "address", address.String())

	var txHash common.Hash
	err := ec.c.CallContext(ctx, &txHash, "eth_setUserBaseContracts", address)
	if err != nil {
		return common.Hash{}, err
	}

	return txHash, nil
}

func (ec *Client) CancelUserBaseContracts(ctx context.Context, address common.Address) (common.Hash, error) {

	log.Info("(ec *Client) CancelUserBaseContracts", "address", address.String())

	var txHash common.Hash
	err := ec.c.CallContext(ctx, &txHash, "eth_cancelUserBaseContracts", address)
	if err != nil {
		return common.Hash{}, err
	}
	return txHash, nil
}

func (ec *Client) AddValidator(ctx context.Context, address common.Address, votes uint64) (common.Hash, error) {

	log.Info("(ec *Client) AddValidator", "address", address.String(), "votes", votes)

	var txHash common.Hash
	err := ec.c.CallContext(ctx, &txHash, "eth_addValidator", address, votes)
	if err != nil {
		return common.Hash{}, err
	}
	return txHash, nil
}

func (ec *Client) GetBlockValidator(ctx context.Context, blockNumber *big.Int) ([]byte, error) {

	log.Info("(ec *Client) GetBlockValidator", "blockNumber", blockNumber.Int64())

	var result []byte
	err := ec.c.CallContext(ctx, &result, "eth_getBlockValidator", toBlockNumArg(blockNumber))
	return result, err
}

//股权
func (ec *Client) StockSet(ctx context.Context, address common.Address, number uint64) (common.Hash, error) {

	log.Info("(ec *Client) StockSet", "address", address.String(), "number", number)

	var txHash common.Hash
	err := ec.c.CallContext(ctx, &txHash, "eth_stockSet", address, number)
	if err != nil {
		return common.Hash{}, err
	}
	return txHash, nil
}

func (ec *Client) GetStocks(ctx context.Context) (*protocol.StocksAccount, error) {

	log.Info("(ec *Client) GetStock")

	var stocks *protocol.StocksAccount
	err := ec.c.CallContext(ctx, &stocks, "eth_stocksGet")
	if err != nil {
		return stocks, err
	}
	return stocks, nil
}

func (ec *Client) GetStock(ctx context.Context, address common.Address) (*protocol.StockAccount, error) {

	log.Info("(ec *Client) GetStock", "address", address.String())

	var stock *protocol.StockAccount
	err := ec.c.CallContext(ctx, &stock, "eth_stockGet", address)
	if err != nil {
		return stock, err
	}
	return stock, nil
}

func (ec *Client) StockTransfer(ctx context.Context, from common.Address, to common.Address, number uint64) (common.Hash, error) {

	log.Info("(ec *Client) StockTransfer", "from", from.String(), "to", to.String(), "number", number)

	var txHash common.Hash
	err := ec.c.CallContext(ctx, &txHash, "eth_stockTransfer", from, to, number)
	if err != nil {
		return common.Hash{}, err
	}
	return txHash, nil
}

func (ec *Client) StockClean(ctx context.Context, address common.Address) (common.Hash, error) {

	log.Info("(ec *Client) StockClean", "address", address.String())

	var txHash common.Hash
	err := ec.c.CallContext(ctx, &txHash, "eth_stockClean", address)
	if err != nil {
		return common.Hash{}, err
	}
	return txHash, nil
}

func (ec *Client) StockFrozen(ctx context.Context, address common.Address) (common.Hash, error) {

	log.Info("(ec *Client) StockFrozen", "address", address.String())

	var txHash common.Hash
	err := ec.c.CallContext(ctx, &txHash, "eth_stockFrozen", address)
	if err != nil {
		return common.Hash{}, err
	}
	return txHash, nil
}

func (ec *Client) StockUnFrozen(ctx context.Context, address common.Address) (common.Hash, error) {

	log.Info("(ec *Client) StockUnFrozen", "address", address.String())

	var txHash common.Hash
	err := ec.c.CallContext(ctx, &txHash, "eth_stockUnFrozen", address)
	if err != nil {
		return common.Hash{}, err
	}
	return txHash, nil
}

func (ec *Client) StockGasPool(ctx context.Context) (uint64, error) {

	log.Info("(ec *Client) StockGasPool")

	var gasPool uint64
	err := ec.c.CallContext(ctx, &gasPool, "eth_stockGasPool")
	if err != nil {
		return 0, err
	}
	return gasPool, nil
}

//扩展
func (ec *Client) SetWord(ctx context.Context, word string) (common.Hash, error) {

	log.Info("(ec *Client) SetWord", "len", len(word), "word", word)

	var txHash common.Hash
	err := ec.c.CallContext(ctx, &txHash, "eth_setWord", word)
	if err != nil {
		return common.Hash{}, err
	}
	return txHash, nil
}

func (ec *Client) SetData(ctx context.Context, data []byte) (common.Hash, error) {

	log.Info("(ec *Client) SetData", "len", len(data))

	var txHash common.Hash
	err := ec.c.CallContext(ctx, &txHash, "eth_setData", data)
	if err != nil {
		return common.Hash{}, err
	}
	return txHash, nil
}

func (ec *Client) GetWord(ctx context.Context, hash common.Hash) (string, error) {

	log.Info("(ec *Client) GetWord", "hash", hash)

	var result string
	err := ec.c.CallContext(ctx, &result, "eth_getWord", hash)
	return result, err
}

func (ec *Client) GetData(ctx context.Context, hash common.Hash) ([]byte, error) {

	log.Info("(ec *Client) GetData", "hash", hash)

	var result []byte
	err := ec.c.CallContext(ctx, &result, "eth_getData", hash)
	return result, err
}

//所属
func (ec *Client) SetStockManager(ctx context.Context, address common.Address) (common.Hash, error) {

	log.Info("(ec *Client) SetOwner", "address", address.String())

	var txHash common.Hash
	err := ec.c.CallContext(ctx, &txHash, "eth_setStockManager", address)
	if err != nil {
		return common.Hash{}, err
	}
	return txHash, nil
}

func (ec *Client) GetStockManager(ctx context.Context) (common.Address, error) {

	log.Info("(ec *Client) GetOwner")

	var address common.Address
	err := ec.c.CallContext(ctx, &address, "eth_getStockManager")
	if err != nil {
		return common.Address{}, err
	}
	return address, nil
}
