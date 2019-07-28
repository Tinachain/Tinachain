//Tina链增加的特殊账号管理类
package boker

import (
	"context"
	"encoding/json"
	"io/ioutil"

	"github.com/Tinachain/Tina/chain/boker/protocol"
	"github.com/Tinachain/Tina/chain/common"
	"github.com/Tinachain/Tina/chain/core/types"
	"github.com/Tinachain/Tina/chain/eth"
	"github.com/Tinachain/Tina/chain/log"
	"github.com/Tinachain/Tina/chain/params"
	"github.com/Tinachain/Tina/chain/trie"
)

const JsonFileName = "boker.json" //Tina链配置

//Tina链中基础合约相关配置信息
type BaseContract struct {
	ContractType    uint64         `json:"contracttype"`    //基础合约类型
	DeployAddress   common.Address `json:"deployaddress"`   //部署账号
	ContractAddress common.Address `json:"contractaddress"` //合约地址
}

type BaseContractConfig struct {
	Bases []BaseContract `json:"bases,omitempty"`
}

type BlackConfig struct {
	Blacks []common.Address `json:"blacks,omitempty"`
}

type StockConfig struct {
	Account common.Address `json:"account"`
	Number  uint64         `json:"number"`
	State   int            `json:"state"`
}

type StocksConfig struct {
	Stocks []StockConfig `json:"bases,omitempty"`
}

//Tina链用来分配通证的账号和私钥信息
type ProducerConfig struct {
	Coinbase common.Address `json:"coinbase"` //挖矿的Coinbase
	Password string         `json:"password"` //Coinbase的密码
}

type BokerConfig struct {
	BlackList *BlackConfig        `json:"blacklist,omitempty"` //黑名单的配置信息
	Dpos      *params.DposConfig  `json:"dpos,omitempty"`      //Dpos的配置信息
	Contracts *BaseContractConfig `json:"contracts,omitempty"` //基础合约配置信息
	Producer  *ProducerConfig     `json:"producer,omitempty"`  //出块节点使用的信息
	Stocks    *StocksConfig       `json:"stocks,omitempty"`    //股权配置
}

type BokerBackend struct {
	config       BokerConfig
	ethereum     *eth.Ethereum
	contracts    *BokerContracts
	transactions *BokerTransaction
	stocks       *TinaStock
	blacks       map[common.Address]bool
}

func New() *BokerBackend {

	log.Info("****New Boker****")
	boker := new(BokerBackend)
	boker.ethereum = nil
	boker.contracts = nil
	boker.transactions = nil
	boker.blacks = make(map[common.Address]bool)
	boker.loadConfig()
	return boker
}

func (boker *BokerBackend) Init(e *eth.Ethereum, bokerProto *protocol.BokerBackendProto) error {

	log.Info("(boker *BokerBackend) Init")

	//创建类
	boker.ethereum = e

	log.Info("Create Bokerchain Transaction Object")
	boker.transactions = NewTransaction(e)

	var err error
	log.Info("Create Bokerchain Contract Object")
	boker.contracts, err = NewContract(e.ChainDb(), e, boker.transactions, bokerProto)
	if err != nil {
		return nil
	}

	boker.stocks = NewStock(e.ChainDb(), e, bokerProto)

	return nil
}

//loadConfig 加载json格式的配置文件
func (boker *BokerBackend) loadConfig() error {

	log.Info("(boker *BokerBackend) loadConfig")

	if boker.ethereum != nil {

		boker.config.Dpos = new(params.DposConfig)
		boker.config.Dpos.Validators = make([]common.Address, 0)

		boker.config.Contracts = new(BaseContractConfig)
		boker.config.Contracts.Bases = make([]BaseContract, 0)

		boker.config.Producer = new(ProducerConfig)
		boker.config.BlackList = new(BlackConfig)

		boker.config.Stocks = new(StocksConfig)

		//读取文件
		buffer, err := ioutil.ReadFile(JsonFileName)
		if err != nil {
			log.Error("Bokerchain Read boker.json File Error", "Error", err)
			return err
		}

		config := BokerConfig{}
		err = json.Unmarshal(buffer, &config)
		if err != nil {
			log.Error("Bokerchain Unmarshal boker.json File Error", "Error", err)
			return err
		}

		//清空原来的数据
		boker.config.Dpos.Validators = make([]common.Address, 0)
		for _, v := range config.Dpos.Validators {
			boker.config.Dpos.Validators = append(boker.config.Dpos.Validators, v)
		}
		log.Info("Load Bokerchain Dpos Validators", "Size", len(boker.config.Dpos.Validators))

		boker.config.Contracts.Bases = make([]BaseContract, 0)
		for _, v := range config.Contracts.Bases {
			boker.config.Contracts.Bases = append(boker.config.Contracts.Bases, v)
		}
		log.Info("Load Bokerchain Base Contracts", "Size", len(boker.config.Contracts.Bases))

		boker.config.BlackList.Blacks = make([]common.Address, 0)
		for _, v := range config.BlackList.Blacks {
			boker.config.BlackList.Blacks = append(boker.config.BlackList.Blacks, v)
		}
		log.Info("(boker *BokerBackend) Blacks", "Size", len(boker.config.BlackList.Blacks))

		boker.config.Producer.Coinbase = config.Producer.Coinbase
		boker.config.Producer.Password = config.Producer.Password
		log.Info("Load Bokerchain Producer", "Coinbase", boker.config.Producer.Coinbase, "Password", boker.config.Producer.Password)

		return nil
	}
	return protocol.ErrSystem
}

func (boker *BokerBackend) SetSystemContract(address common.Address, from common.Address) error {
	return boker.contracts.SetSystemContract(address, from)
}

func (boker *BokerBackend) SetUserBaseContract(address common.Address, from common.Address) error {
	return boker.contracts.SetUserBaseContract(address, from)
}

func (boker *BokerBackend) CancelUserBaseContract(address common.Address, from common.Address) error {
	return boker.contracts.CancelUserBaseContract(address, from)
}

func (boker *BokerBackend) IsSystemBaseContract(address common.Address) bool {
	return boker.contracts.IsSystemBaseContract(address)
}

func (boker *BokerBackend) IsUserBaseContract(address common.Address) bool {
	return boker.contracts.IsUserBaseContract(address)
}

func (boker *BokerBackend) IsLocalValidator(address common.Address) bool {

	if boker.ethereum.GetLocalValidator() == address {
		return true
	} else {
		return false
	}
}

func (boker *BokerBackend) GetBlacks() []common.Address {

	var blacks []common.Address
	for k, _ := range boker.blacks {
		blacks = append(blacks, k)
	}
	return blacks
}

func (boker *BokerBackend) SetBlacks(address []common.Address) error {

	boker.blacks = nil
	boker.blacks = make(map[common.Address]bool)
	for _, v := range address {
		boker.blacks[v] = true
	}
	return nil
}

func (boker *BokerBackend) CheckBlackAddress(address common.Address) bool {

	if _, ok := boker.blacks[address]; ok {
		return true
	}
	return false
}

//SubmitBokerTransaction
func (boker *BokerBackend) SubmitBokerTransaction(ctx context.Context,
	txMajor protocol.TxMajor,
	txMinor protocol.TxMinor,
	to common.Address,
	name []byte,
	extra []byte,
	encryption uint8) (*types.Transaction, error) {

	return boker.transactions.SubmitBokerTransaction(ctx, txMajor, txMinor, to, name, extra, encryption)
}

func (boker *BokerBackend) EncoderContext(context []byte, key []byte) (error, []byte) {
	return boker.transactions.EncoderContext(context, key)
}

func (boker *BokerBackend) DecoderContext(context []byte, key []byte) (error, []byte) {
	return boker.transactions.DecoderContext(context, key)
}

func (boker *BokerBackend) CommitTrie() (*protocol.BokerBackendProto, error) {

	//提交基础合约交易
	contractsRoot, err := boker.contracts.contractsTrie.CommitTo(boker.ethereum.ChainDb())
	if err != nil {
		return nil, err
	}

	return &protocol.BokerBackendProto{
		ContractsHash: contractsRoot,
	}, nil
}

func (boker *BokerBackend) GetContractTrie() (*trie.Trie, *trie.Trie) {

	return boker.contracts.GetContractTrie()
}

func (boker *BokerBackend) GetVotes() error {

	return boker.contracts.GetVotes()
}

func (boker *BokerBackend) GetMethodName(txMinor protocol.TxMinor) (string, string, error) {

	if txMinor < protocol.SetValidator {
		return "", "", protocol.ErrTxType
	}

	switch txMinor {

	case protocol.SetValidator: //设置验证者
		return "", "", nil

	case protocol.RegisterCandidate: //注册成为候选人
		return "", protocol.RegisterCandidateMethod, nil

	case protocol.VoteUser: //用户投票
		return "", protocol.VoteCandidateMethod, nil

	case protocol.VoteEpoch: //产生当前的出块节点
		return "", protocol.RotateVoteMethod, nil

	default:
		return "", "", protocol.ErrTxType
	}
}

//
func (boker *BokerBackend) GetOwner() common.Address {

	return boker.stocks.GetOwner()
}

func (boker *BokerBackend) SetOwner(operation common.Address, address common.Address) error {

	return boker.stocks.SetOwner(operation, address)
}

func (boker *BokerBackend) GetGasPool() uint64 {

	return boker.stocks.GetGasPool()
}

func (boker *BokerBackend) AddGasPool(gas uint64) uint64 {

	return boker.stocks.AddGasPool(gas)
}

func (boker *BokerBackend) GetStock(address common.Address) *protocol.StockAccount {

	return boker.stocks.GetStock(address)
}

func (boker *BokerBackend) GetStocks() ([]*protocol.StockAccount, error) {

	return boker.stocks.GetStocks()
}

func (boker *BokerBackend) SetStock(operation common.Address, address common.Address, number uint64) error {

	return boker.stocks.SetStock(operation, address, number)
}

func (boker *BokerBackend) TransferStock(operation common.Address, from common.Address, to common.Address, number uint64) error {

	return boker.stocks.TransferStock(operation, from, to, number)
}

func (boker *BokerBackend) CleanStock(operation common.Address, from common.Address) error {

	return boker.stocks.CleanStock(operation, from)
}

func (boker *BokerBackend) FrozenStock(operation common.Address, from common.Address) error {

	return boker.stocks.FrozenStock(operation, from)
}

func (boker *BokerBackend) UnFrozenStock(operation common.Address, from common.Address) error {

	return boker.stocks.UnFrozenStock(operation, from)
}

func (boker *BokerBackend) ClearStock(operation common.Address) error {

	return boker.stocks.ClearStock(operation)
}

func (boker *BokerBackend) GetStockTrie() (*trie.Trie, *trie.Trie, *trie.Trie, *trie.Trie) {

	return boker.stocks.GetStockTrie()
}
