
## Tina链（Tinachain）
Tina链是一个专门为服务于文章保存而开发的垂直型区块链平台，用户可以非常方便的使用Tina链所提供的接口将自己的文章永久保存在区块链上。<br/>
Bokerchain is a public blockchain platform that serves the vertical area of audio & video. It is convenient for different intelligent terminal devices to access to Bokerchain. We can form a Video Application Union by providing SDK for various video APP, meeting the need for copyright protection, data sharing and benefit protection among all apps in Video Application Union. In this situation, we may make data among podcasts, advertisers and users more transparent, also maximizing the benefits.<br/>
Our goal is to acheive video sharing, benefits sharing and user resources sharing, benefiting podcasts, advertisers and our users while providing entertainment.

## Tina链系统架构（System architecture）
![Image text](https://github.com/Tinachain/Tina/blob/master/image/Architecture.jpg)

## Tina链交易结构
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

## Tina链启动方式

第一步：初始化创世文件

    geth --datadir "/projects/tina/node" init genesis.json


第二步：启动geth

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


第三步：进入geth控制台

    geth attach ipc:/projects/tina/node/geth.ipc

第四步：创建账号

    personal.newAccount()


第五步：设置帐号解锁（这里使用假定账号、密码）

    personal.unlockAccount(account, password, 0)

第六步：设置自己为验证人

    miner.setLocalValidator()

第七步：设置验证人（这里使用假定账号、票数）

    eth.addValidator(account, 10000)

第八步：启动挖矿

    miner.start()





## 关于我们（About us）

### Tinachain Co-Founder
* WeChat: [区什么块什么链啊](Blockchain_fxh7622) 
* Twitter: [区什么块什么链啊](https://twitter.com/chain_fxh7622) 
* Twitter: [后青春期的诗](https://twitter.com/chain_stayreal)


## 目录（Folders）

### [chain](https://github.com/Bokerchain/Boker/tree/master/chain)
基础链代码，实现DPOS<br/>
Main chain code, implementing DPOS.

### [contracts](https://github.com/Bokerchain/Boker/tree/master/contracts)
用Solidity写的基础合约代码<br/>
Basic contract code in solidity.

### [explorer](https://github.com/Bokerchain/Boker/tree/master/explorer)
Tina链的区块链浏览器以及用户文章上传页面<br/>
Basic contract code in solidity.
