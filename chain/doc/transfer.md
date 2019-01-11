# 调用方式

## 	获取到链的配置信息
	
	var ethereum *eth.Ethereum
	if err := stack.Service(&ethereum); err != nil {
		utils.Fatalf("ethereum service not running: %v", err)
	}
	chainConfig := ethereum.BlockChain().Config()
	
	
## 主动创建交易

	DeployKey, err := crypto.HexToECDSA(tina.ethereum.BlockChain().Config().Producer.PrivateKey)
	DeployAddr := crypto.PubkeyToAddress(DeployKey.PublicKey)
	DeployBalance := big.NewInt(0)
	DeployBalance.SetInt64(tina.ethereum.BlockChain().Config().Producer.Balance)

	//构造backend和opts
	backend := backends.NewSimulatedBackend(core.GenesisAlloc{DeployAddr: {Balance: DeployBalance}})
	opts := bind.NewKeyedTransactor(DeployKey)

	//得到Nonce
	nonce, err := backend.PendingNonceAt(opts.Context, tina.ethereum.BlockChain().Config().Coinbase)
		
	//交易
	var input []byte
	rawTx := types.NewBaseTransaction(transactTypes, nonce, address, value, input)
	
	