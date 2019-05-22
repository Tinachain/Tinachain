
### Tina链（Tinachain）
Tina链是一个专门为服务于文字、图片、文章以及文件保存而开发的垂直型区块链平台，用户可以使用Tina链提供的接口将自己的相关信息永久保存到Tina链上。同时为了保证信息的安全和保密，Tina链对于信息的保存提供了可选加密功能。从而保证了用户信息的安全性。
Bokerchain is a public blockchain platform that serves the vertical area of audio & video. It is convenient for different intelligent terminal devices to access to Bokerchain. We can form a Video Application Union by providing SDK for various video APP, meeting the need for copyright protection, data sharing and benefit protection among all apps in Video Application Union. In this situation, we may make data among podcasts, advertisers and users more transparent, also maximizing the benefits.<br/>
Our goal is to acheive video sharing, benefits sharing and user resources sharing, benefiting podcasts, advertisers and our users while providing entertainment.

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
	
## Building the source

Building Bokerchain requires both a Go (version 1.7.3 or later) and a C compiler. You can install them using your favourite package manager. Once the dependencies are installed, run

    make geth

or, to build the full suite of utilities:

    make all

## Executables

The Bokerchain project comes with several wrappers/executables found in the `cmd` directory.

| Command    | Description |
|:----------:|-------------|
| **`geth`** | Our main Ethereum CLI client. It is the entry point into the Ethereum network (main-, test- or private net), capable of running as a full node (default) archive node (retaining all historical state) or a light node (retrieving data live). It can be used by other processes as a gateway into the Ethereum network via JSON RPC endpoints exposed on top of HTTP, WebSocket and/or IPC transports. `geth --help` and the [CLI Wiki page](https://github.com/ethereum/chain/wiki/Command-Line-Options) for command line options. |
| `abigen` | Source code generator to convert Ethereum contract definitions into easy to use, compile-time type-safe Go packages. It operates on plain [Ethereum contract ABIs](https://github.com/ethereum/wiki/wiki/Ethereum-Contract-ABI) with expanded functionality if the contract bytecode is also available. However it also accepts Solidity source files, making development much more streamlined. Please see our [Native DApps](https://github.com/ethereum/chain/wiki/Native-DApps:-Go-bindings-to-Ethereum-contracts) wiki page for details. |

## Configuration

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
**Note: `"byzantiumBlock": 0` should be the config value, otherwise contract call may malfunction!**

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
