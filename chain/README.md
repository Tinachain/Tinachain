
### Tina链（Tinachain）
Tina链是一个专门为服务于文字、图片、文章以及文件保存而开发的垂直型区块链平台，用户可以使用Tina链提供的接口将自己的相关信息永久保存到Tina链上。同时为了保证信息的安全和保密，Tina链对于信息的保存提供了可选加密功能。从而保证了用户信息的安全性。

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
    	Word         []byte           `json:"word"    gencodec:"required"`          
    	Extra        []byte           `json:"extra"    gencodec:"required"`         
    	Ip           []byte           `json:"ip"    gencodec:"required"`            
    	V *big.Int `json:"v" gencodec:"required"`
    	R *big.Int `json:"r" gencodec:"required"`
    	S *big.Int `json:"s" gencodec:"required"`
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
    	Word             string           `json:"word"`
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
		"byzantiumBlock": 0,
		"eip155Block": 0,
		"eip158Block": 0
	},
	"alloc": {},
	"difficulty": "0x000001",
	"extraData": "",
	"gasLimit": "0xffffffff"
}

```
**注意：`"byzantiumBlock": 0` 这个配置，否则合约调用有可能出错。**

### boker.json
```json
{
    "dpos":{
        "validators":[
            "0xdd165ba267593d2acc837fc507c2e94e802817d9"
        ]
    },
    "contracts":{
        "bases":[
            {
                "contracttype":2,
                "deployaddress":"0xdd165ba267593d2acc837fc507c2e94e802817d9",
                "contractaddress":"0xd7fd311c8f97349670963d87f37a68794dfa80ff"
            }
        ]
    },
    "producer":{
        "coinbase":"0x1aa228dde26f02e1cc5551cb4f1d74d0e998d24a",
        "password":"123456"
    }
}
```

### Tina链启动方式

第一步：同步时钟（由于采用DPOS共识机制，因此对于时钟同步非常重要）

	crontab -e

添加内容

	*/10 * * * * /usr/sbin/ntpdate 1.cn.pool.ntp.org


第二步：初始化创世文件

    geth --datadir "/projects/tina/node" init genesis.json


第三步：启动geth

    nohup geth --nodiscover  \
    --maxpeers 3 \
    --identity "tinachain" \
    --rpc \
    --rpcaddr 0.0.0.0 \
    --rpccorsdomain "*" \
    --rpccorsdomain "http://localhost:8000" \
    --datadir "/projects/tina/node" \
    --port 30304 \
    --rpcapi "db,eth,net,web3" \
    --networkid 96579 &


第四步：进入geth控制台

    geth attach ipc:/projects/tina/node/geth.ipc

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
	
	
### RPC 指令

#### 文字上链

1：将文字上链

    eth.setWord("This is Test")
    
2：从链上获取文字（交易Hash）

    eth.getWord("0x26635445ae6e1f20bc2a7ed5be45c3a0b7e847e1c79167c9b1564fe77ef72094")


#### 图片上链

1：将图片上链（其中：/projects/tina/1.jpg 是节点上地址）

    eth.setPicture("/projects/tina/1.jpg")
	
2：从链上获取图片（其中：交易Hash、保存图片的节点目录，文件名使用链上的文件名，例如设置/projects/tina则保存的图片地址为：/projects/tina/1.jpg）

    eth.getPicture("0x26635445ae6e1f20bc2a7ed5be45c3a0b7e847e1c79167c9b1564fe77ef72094", "/projects/tina")
