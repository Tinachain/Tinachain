
### Tina链（Tinachain）
Tina链一款采用Dpos共识机制设计的区块链产品。
Tina链为了避免沦为无实际应用的发币工具，因此Tina链中自带了股权机制。
Tina链的设计初衷是为了方便用户在区块链上永久保存文字、图片、文章以及文件保存而开发的垂直型区块链平台，用户可以使用Tina链提供的接口将自己的相关信息永久保存到Tina链上。
同时为了保证信息的安全和保密，Tina链对于信息的保存提供了可选加密功能。从而保证了用户信息的安全性。

Tinachain is a vertical blockchain platform specifically designed for preservation of text, images, articles and files. With rpc interfaces provided by Tinachain, it is convenient for users to save personal information permanently. Meanwhile, Tinachain provides optional encryption function for users to ensure the security and privacy of their information.

### Tina链交易结构
* Tina链的交易结构格式

    type txdata struct {
    
        Major        protocol.TxMajor `json:"major"   gencodec:"required"`          
    	Minor        protocol.TxMinor `json:"minor"   gencodec:"required"`          
		AccountNonce uint64           `json:"nonce"    gencodec:"required"`         
		Price        *big.Int         `json:"gasPrice" gencodec:"required"`         
		GasLimit     *big.Int         `json:"gas"      gencodec:"required"`         
		Time         *big.Int         `json:"timestamp"        gencodec:"required"` 
		Recipient    *common.Address  `json:"to"       rlp:"nil"`                   
		Amount       *big.Int         `json:"value"    gencodec:"required"`         
		Payload      []byte           `json:"input"    gencodec:"required"`         
		Name         []byte           `json:"name"    gencodec:"required"`          
		Encryption   uint8            `json:"encryption"    gencodec:"required"`    
		Extra        []byte           `json:"extra"    gencodec:"required"`         
		Ip           []byte           `json:"ip"    gencodec:"required"`            

		V *big.Int `json:"v" gencodec:"required"`
		R *big.Int `json:"r" gencodec:"required"`
		S *big.Int `json:"s" gencodec:"required"`

		// This is only used when marshaling to JSON.
		Hash *common.Hash `json:"hash" rlp:"-"`
    }

* Tina链的RPC交易结构

    type RPCTransaction struct {
	
    	Major            protocol.TxMajor `json:"major"`
		MajorNotes       string           `json:"majorNotes"`
		Minor            protocol.TxMinor `json:"minor"`
		MinorNotes       string           `json:"minorNotes"`
		BlockHash        common.Hash      `json:"blockHash"`
		BlockNumber      *hexutil.Big     `json:"blockNumber"`
		From             common.Address   `json:"from"`
		Gas              *hexutil.Big     `json:"gas"`
		GasPrice         *hexutil.Big     `json:"gasPrice"`
		Hash             common.Hash      `json:"hash"`
		Input            hexutil.Bytes    `json:"input"`
		Name             string           `json:"name"`
		Encryption       uint8            `json:"encryption"`
		Extra            hexutil.Bytes    `json:"extra"`
		Ip               string           `json:"ip"`
		Nonce            hexutil.Uint64   `json:"nonce"`
		To               *common.Address  `json:"to"`
		TransactionIndex hexutil.Uint     `json:"transactionIndex"`
		Value            *hexutil.Big     `json:"value"`
		V                *hexutil.Big     `json:"v"`
		R                *hexutil.Big     `json:"r"`
		S                *hexutil.Big     `json:"s"`
    }
	
### 源码编译

编译Tinachain需要Go语言版本在1.9.0以上，并且运行

    make geth

或者编译所有

    make all

### 配置

### genesis.json
```json
{
    "config": {
		"chainId": 0,
		"homesteadBlock": 0,
		"eip155Block": 0,
		"eip158Block": 0
	},
	"alloc": {},
	"difficulty": "0x000001",
	"extraData": "",
	"gasLimit": "0xffffffff"
}

```

### Tina链启动方式

第一步：同步时钟（由于采用DPOS共识机制，因此对于时钟同步非常重要）

	crontab -e

添加内容

	*/10 * * * * /usr/sbin/ntpdate 1.cn.pool.ntp.org


第二步：初始化创世文件

    geth --datadir "/projects/tinachain/node" init genesis.json


第三步：启动geth

    nohup /projects/tinachain/geth --nodiscover  \
    --maxpeers 3 \
    --identity "tinachain" \
    --rpc \
    --rpcaddr 0.0.0.0 \
    --rpccorsdomain "*" \
    --datadir "/projects/tinachain/node" \
    --port 30404 \
    --rpcapi "db,eth,net,web3" \
    --networkid 98765 &


第四步：进入geth控制台

    geth attach ipc:/projects/tinachain/node/geth.ipc

第五步：创建账号

    personal.newAccount()


第六步：设置帐号解锁（这里使用假定账号、密码）

    personal.unlockAccount(account, password, 0)

第七步：设置自己为验证人

    miner.setLocalValidator()

第八步：设置验证人（这里使用假定账号、票数）

    eth.addValidator(account, 10000)

第九步：启动挖矿

    miner.start()
	
	
### 新增 RPC 指令

1：得到最后一个出块超级节点

    GetLastProducerAt
    
2：得到下一个出块超级节点

    GetNextProducerAt
    
3：设置系统基础合约（系统调用，免Gas）

    SetSystemBaseContracts
    
4：设置用户基础合约（免Gas）
    
    SetUserBaseContracts
    
5：取消用户基础合约（免Gas）

    CancelUserBaseContracts
    
6：添加超级节点（创世超级节点使用）
    
    AddValidator
    
7：根据区块编号得到出口超级节点

    GetBlockValidator
    
8：设置股权

    StockSet
    
9：得到指定账号的股权数量

    GetStock
    
10：股权交易

    StockTransfer
    
11：股权清空
    
    StockClean
    
12：股权冻结

    StockFrozen
    
13：股权解冻

    StockUnFrozen
    
14：得到当前股权池中Gas总数量（用于分给股权用户的Gas数量）
    
    StockGasPool
    
15：设置股权管理账号

    SetStockManager
    
16：设置文字（文字上链）

    SetWord
    
17：设置数据（数据上链）
    
    SetData
    


#### 文字上链

1：将文字上链

    eth.setWord("This is Test")
    
2：从链上获取文字（交易Hash）

    eth.getWord("0x8c52e28e4819d28f215563635e7a4971cddd5ad48d21424383c832ff03b12800")


#### 图片上链

1：将图片上链（其中：/projects/tina/1.jpg 是节点上地址）

    eth.setPicture("/projects/tina/1.jpg")
	
2：从链上获取图片（其中：交易Hash、保存图片的节点目录，文件名使用链上的文件名，例如设置/projects/tina则保存的图片地址为：/projects/tina/1.jpg）

    eth.getPicture("0x26635445ae6e1f20bc2a7ed5be45c3a0b7e847e1c79167c9b1564fe77ef72094", "/projects/tina")

