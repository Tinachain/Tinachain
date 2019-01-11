## 新增的RPC

# 1：得到当前的出块节点
	eth.getCurrentProducer
	
# 2：得到当前的分币节点
	eth.getCurrentTokenNoder
	
# 3：设置当前的Coinbase私钥
	eth.setPrivateKey
	
# 4：设置基础合约	
	eth.setBaseContracts
	
# 5：取消基础合约
	eth.cancelBaseContracts
	
# 6：添加验证人
	eth.addValidator
	(for example) eth.addValidator("0xa1fb5ba97afb0e1dc743bbdcd8c96ae191780b79", 10000)