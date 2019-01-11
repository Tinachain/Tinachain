# 流程说明

## 	交易调用流程

	* SendTransaction() 声明 (accounts/abi/bind/backend.go)
	* SendTransaction()->block.AddTx() 位置 (accounts/abi/bind/backends/simulated.go)
	* AddTx()->ApplyTransaction() 位置 (core/chain_makers.go)
	* ApplyTransaction()->binaryTransaction() 位置 (core/state_processer.go)
	* binaryTransaction()->BinaryMessage() 位置 (core/state_processer.go)
	* BinaryMessage()->st.TransitionDb() 位置 (core/state_transition.go)
	* TransitionDb()->evm.Call() 位置 (core/state_transition.go)
	* Call()->run() 位置 (core/vm/evm.go)

## RPC添加流程

	* Client(ethclient/ethclient.go)中添加需要处理的RPC指令 例如 "eth_getCurrentTokenNoder" eth_ 后第一个字母小写
	* PublicBlockChainAPI(internal/ethapi/api.go)中新增 GetCurrentTokenNoder 函数名为eth_ 后名，第一个字母大写
	* 在添加指向(internal/web3ext/web3ext.go)
		new web3._extend.Method({
			name: 'getCurrentTokenNoder',
			call: 'eth_getCurrentTokenNoder',
			params: 0,
		}),		
		
## 加载RPC提供的API接口

	* 函数GetAPIs用来加载所有支持的API接口函数(eth/backend.go)
	* ethapi.GetAPIs 用来加载所有公共服务的API接口
	
		NewPublicEthereumAPI, 
		NewPublicBlockChainAPI, 
		NewPublicTransactionPoolAPI, 
		NewPublicTxPoolAPI, 
		NewPublicDebugAPI, 
		NewPrivateDebugAPI, 
		NewPublicAccountAPI, 
		NewPrivateAccountAPI
		
	* 加载共识引擎提供的API函数接口
		
		apis = append(apis, s.engine.APIs(s.BlockChain())...)
		
	* 加载提供的调试类API接口
		NewPublicEthereumAPI,
		NewPublicMinerAPI,
		NewPublicDownloaderAPI,
		NewPrivateMinerAPI,
		NewPublicFilterAPI,
		NewPrivateAdminAPI,
		NewPublicDebugAPI,
		NewPrivateDebugAPI		
		
		
## 挖矿流程
	
	* newWorker(miner/worker.go)创建矿工
	* start(miner/worker.go)启动矿工挖矿
	* go self.mintLoop() 启动循环挖矿
	* 间隔一秒调用一次 self.mintBlock(now.Unix())
	* 获取当前共识机制 代码 self.engine.(*dpos.Dpos)
	* 判断当前节点是否是验证者节点 即 出块节点 代码 engine.UpdateProducer(self.chain.CurrentBlock(), self.coinbase)
	* 判断是否是正常的出块时间 代码 engine.CheckDeadline(self.chain.CurrentBlock(), now)
	* 检测当前是否是出块节点 代码 engine.CheckProducer(self.chain.CurrentBlock(), now)
	* 开始挖矿 代码 self.createNewWork()
	
## 	调用方式
	* package ethclient
	
## 类之间的调用关系
	* PublicBlockChainAPI(internal/ethapi/api.go) ->NewContractBackend(eth/bind.go) 以及 Service: NewContractBackend (internal/ethapi/backend.go)
	  在eth/backend.go文件中 apis := ethapi.GetAPIs(s.ApiBackend) 类进行调用的
	* PublicEthereumAPI和PublicBlockChainAPI产生方法相同
	* PublicTransactionPoolAPI和PublicBlockChainAPI产生方法相同

	* 此处注意设置Boker的接口的时间

## 新增RPC接口
	PublicBlockChainAPI
		* GetCurrentProducer		
		* GetCurrentTokenNoder
		* SetPrivateKey
		* SetBaseContracts
		* CancelBaseContracts
	
	RPC函数接口
		* eth_getCurrentProducer
		* eth_getCurrentTokenNoder
		* eth_setPrivateKey
		* eth_setBaseContracts
		* eth_cancelBaseContracts


