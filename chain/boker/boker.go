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
}

type BokerBackend struct {
	config       BokerConfig
	ethereum     *eth.Ethereum
	accounts     *BokerAccount
	contracts    *BokerContracts
	transactions *BokerTransaction
	stock        *TinaStock
	blacks       map[common.Address]bool
}

func New() *BokerBackend {

	log.Info("****New Boker****")
	boker := new(BokerBackend)
	boker.ethereum = nil
	boker.accounts = nil
	boker.contracts = nil
	boker.transactions = nil
	boker.blacks = make(map[common.Address]bool)
	boker.loadConfig()
	return boker
}

func (boker *BokerBackend) Init(e *eth.Ethereum, bokerProto *protocol.BokerBackendProto) error {

	log.Info("****Init Boker****")

	//创建类
	boker.ethereum = e

	log.Info("Create Bokerchain Transaction Object")
	boker.transactions = NewTransaction(e)

	log.Info("Create Bokerchain Account Object")
	boker.accounts = NewAccount()

	var err error
	log.Info("Create Bokerchain Contract Object")
	boker.contracts, err = NewContract(e.ChainDb(), e, boker.transactions, bokerProto)
	if err != nil {
		return nil
	}
	return nil
}

//loadConfig 加载json格式的配置文件
func (boker *BokerBackend) loadConfig() error {

	log.Info("****loadConfig****")

	if boker.ethereum != nil {

		boker.config.Dpos = new(params.DposConfig)
		boker.config.Dpos.Validators = make([]common.Address, 0)

		boker.config.Contracts = new(BaseContractConfig)
		boker.config.Contracts.Bases = make([]BaseContract, 0)

		boker.config.Producer = new(ProducerConfig)
		boker.config.BlackList = new(BlackConfig)

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

//GetAccount 根据账号地址，得到账号等级
func (boker *BokerBackend) GetAccount(address common.Address) (protocol.TxMajor, error) {

	if boker.accounts == nil {

		log.Error("Boker GetAccount function accounts Objects is nil")
		return protocol.Normal, nil
	}

	return boker.accounts.GetAccount(address)
}

//GetContract 根据合约地址，得到合约等级
func (boker *BokerBackend) GetContract(address common.Address) (protocol.TxMajor, error) {
	return boker.contracts.GetContract(address)
}

//SetContract 回写合约信息
func (boker *BokerBackend) SetContract(address common.Address, contractType protocol.ContractType, isCancel bool, abiJson string) error {
	return boker.contracts.SetContract(address, contractType, isCancel, abiJson)
}

func (boker *BokerBackend) GetContractAddr(contractType protocol.ContractType) (common.Address, error) {
	return boker.contracts.GetContractAddr(contractType)
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

//SubmitBokerTransaction 设置一个Tina链交易
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

func (boker *BokerBackend) GetContractTrie() (*trie.Trie, *trie.Trie, *trie.Trie) {

	return boker.contracts.GetContractTrie()
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

func (boker *BokerBackend) GetGasPool() uint64 {

	return boker.stock.GetGasPool()
}

func (boker *BokerBackend) AddGasPool(gas uint64) uint64 {

	return boker.stock.AddGasPool(gas)
}

func (boker *BokerBackend) GetOwner() common.Address {

	return boker.stock.GetOwner()
}

func (boker *BokerBackend) SetOwner(operation common.Address, address common.Address) error {

	return boker.stock.SetOwner(operation, address)
}
