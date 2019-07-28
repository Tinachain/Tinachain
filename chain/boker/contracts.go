//Tina链增加的特殊账号管理类
package boker

import (
	"errors"
	"strconv"

	"github.com/Tinachain/Tina/chain/boker/protocol"
	"github.com/Tinachain/Tina/chain/common"
	"github.com/Tinachain/Tina/chain/contracts/boker_interface"
	"github.com/Tinachain/Tina/chain/eth"
	"github.com/Tinachain/Tina/chain/ethdb"
	"github.com/Tinachain/Tina/chain/log"
	"github.com/Tinachain/Tina/chain/rlp"
	"github.com/Tinachain/Tina/chain/trie"
)

//基本合约配置信息
type ContractConfig struct {
	ContractType    protocol.BaseContractType `json:"contractType"`    //基础合约类型
	ContractAddress common.Address            `json:"contractAddress"` //合约地址
}

type ContractConfigs struct {
	baseContractArray []ContractConfig //基础合约
}

type ContractService struct {
	contract *boker_contract.BokerInterfaceService //合约服务类
}

//Tina链的基础合约管理
type BokerContracts struct {
	singleTrie    *trie.Trie                                   //基础合约Hash树，Key为基础合约的Address Value为合约的类型
	contractsTrie *trie.Trie                                   //所有基础合约整体保存树，这里保存所有的基础合约信息，但是不保存基础合约的类型信息
	ethereum      *eth.Ethereum                                //以太坊对象
	transactions  *BokerTransaction                            //Tina链交易对象
	db            ethdb.Database                               //数据库
	contracts     map[common.Address]protocol.BaseContractType //基础合约的Map
	services      ContractService                              //合约服务类
}

func NewContract(db ethdb.Database, ethereum *eth.Ethereum, transactions *BokerTransaction, bokerProto *protocol.BokerBackendProto) (*BokerContracts, error) {

	log.Info("****NewContract****")

	log.Info("Create Bokerchain Single Contract Trie", "Hash", bokerProto.SingleHash.String())
	singleContractTrie, errTrie := NewSingleContractTrie(bokerProto.SingleHash, db)
	if errTrie != nil {
		return nil, errTrie
	}

	log.Info("Create Bokerchain Contracts Trie", "Hash", bokerProto.ContractsHash.String())
	contractsTrie, errTrie := NewContractsTrie(bokerProto.ContractsHash, db)
	if errTrie != nil {
		return nil, errTrie
	}

	//创建对象
	c := new(BokerContracts)
	c.singleTrie = singleContractTrie
	c.contractsTrie = contractsTrie
	c.contracts = make(map[common.Address]protocol.BaseContractType)
	c.db = db
	c.ethereum = ethereum
	c.transactions = transactions

	//从树中加载合约信息
	log.Info("Load Bokerchain Base Contract Config")
	var err error
	if err = c.loadTrieContract(); err != nil {
		log.Error("Load Bokerchain Base Contract Trie", "error", err)
		return c, nil
	}

	//创建投票合约服务
	log.Info("Check Bokerchain System Contract Exists")
	var address common.Address
	address, err = c.getContractAddress(protocol.System)
	if err != nil {

		log.Debug("Bokerchain System Contract is`t Exists")
		return c, nil
	}

	log.Info("New Bokerchain System Contract Service")
	c.services.contract, err = boker_contract.NewBokerInterfaceService(c.ethereum, address)
	if err != nil {

		return c, nil
	}
	log.Info("Start Bokerchain System Contract Service")
	c.services.contract.Start()

	return c, nil
}

//基础合约详细信息树
func NewSingleContractTrie(root common.Hash, db ethdb.Database) (*trie.Trie, error) {
	return trie.NewTrieWithPrefix(root, protocol.SingleContractPrefix, db)
}

//创建基础合约列表树
func NewContractsTrie(root common.Hash, db ethdb.Database) (*trie.Trie, error) {
	return trie.NewTrieWithPrefix(root, protocol.ContractsPrefix, db)
}

//设置合约到Hash树中
func (c *BokerContracts) SetSystemContract(address common.Address, from common.Address) error {

	//设置基础合约
	log.Info("(c *BokerContracts) SetSystemBaseContract", "address", address.String(), "from", from.String())

	exist, err := c.existContract(address)
	if exist {
		log.Error("SetSystemBaseContract existContract Contract Existed")
		return protocol.ErrContractExist
	}

	//设置单个树的信息
	if err = c.singleTrie.TryUpdate(address.Bytes(), []byte(strconv.Itoa(int(protocol.System)))); err != nil {
		log.Error("SetSystemBaseContract singleTrie.TryUpdate failed", "err", err)
		return err
	}
	c.contracts[address] = protocol.System

	//更新基础合约列表保存树
	if err = c.setContractsTrie(); err != nil {
		log.Error("SetSystemBaseContract setContractsTrie failed", "err", err)
		return protocol.ErrSaveContractTrie
	}

	c.services.contract, err = boker_contract.NewBokerInterfaceService(c.ethereum, address)
	if err != nil {
		log.Error("SetSystemBaseContract Start NewAssignTokenService failed", "err", err)
		return protocol.ErrNewContractService
	}

	log.Info("SetContract Start Contract")
	c.services.contract.Start()

	return nil
}

func (c *BokerContracts) SetUserBaseContract(address common.Address, from common.Address) error {

	//设置基础合约
	log.Info("(c *BokerContracts) SetUserBaseContract", "address", address.String(), "from", from.String())

	exist, err := c.existContract(address)
	if exist {
		log.Error("SetSystemBaseContract existContract Contract Existed")
		return protocol.ErrContractExist
	}

	if err = c.singleTrie.TryUpdate(address.Bytes(), []byte(strconv.Itoa(int(protocol.User)))); err != nil {
		log.Error("SetSystemBaseContract singleTrie.TryUpdate failed", "err", err)
		return err
	}
	c.contracts[address] = protocol.User

	if err = c.setContractsTrie(); err != nil {
		log.Error("SetSystemBaseContract setContractsTrie failed", "err", err)
		return protocol.ErrSaveContractTrie
	}

	return nil
}

func (c *BokerContracts) CancelUserBaseContract(address common.Address, from common.Address) error {

	log.Info("(c *BokerContracts) CancelUserBaseContract")
	ret := c.IsUserBaseContract(address)
	if !ret {
		return errors.New("Address is`t User Base Contract")
	}

	if err := c.singleTrie.TryDelete(address.Bytes()); err != nil {
		return err
	}
	delete(c.contracts, address)

	//更新基础合约列表保存树
	if err := c.setContractsTrie(); err != nil {
		return protocol.ErrSaveContractTrie
	}

	return nil
}

func (c *BokerContracts) IsUserBaseContract(address common.Address) bool {

	user, err := c.readContractType(address)
	if err != nil {
		return false
	}

	if user == protocol.User {
		return true
	}
	return false
}

func (c *BokerContracts) IsSystemBaseContract(address common.Address) bool {

	system, err := c.readContractType(address)
	if err != nil {
		return false
	}

	if system == protocol.System {
		return true
	}
	return false
}

//将所有的合约地址进行设置
func (c *BokerContracts) setContractsTrie() error {

	//转换成切片数组
	log.Info("(c *BokerContracts) setContractsTrie")

	var contracts []common.Address = make([]common.Address, 0)
	for k, _ := range c.contracts {
		contracts = append(contracts, k)
	}
	contractsRLP, err := rlp.EncodeToBytes(contracts)
	if err != nil {

		log.Error("failed to encode contracts to rlp", "error", err)
		return err
	}
	c.contractsTrie.Update(protocol.ContractsPrefix, contractsRLP)
	return nil
}

//得到所有的合约地址
func (c *BokerContracts) getContractsTrie() ([]common.Address, error) {

	log.Info("(c *BokerContracts) getContractsTrie")

	var contracts []common.Address
	if c.contractsTrie == nil {
		log.Error("contract Trie is nil")
		return []common.Address{}, protocol.ErrPointerIsNil
	}

	contractsRLP, err := c.contractsTrie.TryGet(protocol.ContractsPrefix)
	if err != nil {
		return []common.Address{}, err
	}

	if err := rlp.DecodeBytes(contractsRLP, &contracts); err != nil {
		log.Error("failed to decode contracts", "error", err)
		return []common.Address{}, err
	}

	return contracts, nil
}

//从Hash树中获取合约类型
func (c *BokerContracts) readContractType(address common.Address) (protocol.BaseContractType, error) {

	log.Info("(c *BokerContracts) readContractType")

	//根据地址获取数据
	key := address.Bytes()
	v, err := c.singleTrie.TryGet(key)
	if err != nil {
		return protocol.System, err
	}

	//转换成交易类型
	contractType, err := strconv.Atoi(string(v[:]))
	return protocol.BaseContractType(contractType), nil
}

//根据合约类型得到合约地址
func (c *BokerContracts) getContractAddress(contractType protocol.BaseContractType) (common.Address, error) {

	if len(c.contracts) <= 0 {
		return common.Address{}, protocol.ErrNotFoundContract
	}

	for k, v := range c.contracts {

		if v == contractType {
			return k, nil
		}
	}
	return common.Address{}, protocol.ErrNotFoundContract
}

//加载基础合约信息
func (c *BokerContracts) loadTrieContract() error {

	log.Info("(c *BokerContracts) loadTrieContract")

	//获取所有合约
	contracts, err := c.getContractsTrie()
	if err != nil {
		return err
	}
	log.Info("Load Boker Trie Contract", "Size", len(contracts))
	if len(contracts) <= 0 {
		return nil
	}

	//根据所有合约得到合约类型
	for _, address := range contracts {
		contractType, err := c.readContractType(address)
		if err != nil {
			continue
		}
		c.contracts[address] = contractType
	}
	return nil
}

//查找合约账户
/*func (c *BokerContracts) GetContract(address common.Address) (protocol.TxMajor, error) {

	if len(c.contracts) > 0 {
		_, exist := c.contracts[address]
		if exist {
			return protocol.Base, nil
		}
	}
	return protocol.Normal, nil
}*/

//判断此合约是否已经存在
func (c *BokerContracts) existContract(address common.Address) (bool, error) {

	_, exist := c.contracts[address]
	if exist {
		return true, nil
	}
	return false, protocol.ErrNotFoundContract
}

func (c *BokerContracts) GetContractTrie() (*trie.Trie, *trie.Trie) {

	return c.singleTrie, c.contractsTrie
}

func (c *BokerContracts) GetVotes() error {

	if c.services.contract != nil {

		c.services.contract.GetVotes()
	}
	return nil
}
