package types

import (
	"errors"
	"math/big"
	"strconv"

	"github.com/Tinachain/Tina/chain/boker/protocol"
	"github.com/Tinachain/Tina/chain/common"
	"github.com/Tinachain/Tina/chain/crypto/sha3"
	"github.com/Tinachain/Tina/chain/ethdb"
	"github.com/Tinachain/Tina/chain/log"
	"github.com/Tinachain/Tina/chain/rlp"
	"github.com/Tinachain/Tina/chain/trie"
)

type BokerBackendProto struct {
	SingleHash      common.Hash `json:"SingleRoot"        gencodec:"required"`
	ContractsHash   common.Hash `json:"ContractsRoot"    gencodec:"required"`
	SingleStockHash common.Hash `json:"singleStockRoot"      gencodec:"required"`
	StocksHash      common.Hash `json:"stocksRoot"      gencodec:"required"`
	OwnerHash       common.Hash `json:"ownerRoot"      gencodec:"required"`
	GasPoolHash     common.Hash `json:"gasPoolRoot"      gencodec:"required"`
}

func (p *BokerBackendProto) Root() (h common.Hash) {

	hw := sha3.NewKeccak256()
	rlp.Encode(hw, p.SingleHash)
	rlp.Encode(hw, p.ContractsHash)
	rlp.Encode(hw, p.SingleStockHash)
	rlp.Encode(hw, p.StocksHash)
	rlp.Encode(hw, p.OwnerHash)
	rlp.Encode(hw, p.GasPoolHash)
	hw.Sum(h[:0])
	return h
}

func ToBokerProto(singleHash, contractsHash, singleStockHash, stocksHash, ownerHash, gasPoolHash common.Hash) *BokerBackendProto {

	return &BokerBackendProto{
		SingleHash:      singleHash,
		ContractsHash:   contractsHash,
		SingleStockHash: singleStockHash,
		StocksHash:      stocksHash,
		OwnerHash:       ownerHash,
		GasPoolHash:     gasPoolHash,
	}
}

type BokerContext struct {
	singleContractsTrie *trie.Trie
	contractsTrie       *trie.Trie
	singleStockTrie     *trie.Trie
	stocksTrie          *trie.Trie
	ownerTrie           *trie.Trie
	gasPoolTrie         *trie.Trie
	db                  ethdb.Database
}

func NewSingleContractTrie(root common.Hash, db ethdb.Database) (*trie.Trie, error) {
	return trie.NewTrieWithPrefix(root, protocol.SingleContractPrefix, db)
}

func NewContractsTrie(root common.Hash, db ethdb.Database) (*trie.Trie, error) {
	return trie.NewTrieWithPrefix(root, protocol.ContractsPrefix, db)
}

func NewSingleStockTrie(root common.Hash, db ethdb.Database) (*trie.Trie, error) {
	return trie.NewTrieWithPrefix(root, protocol.SingleStockPrefix, db)
}

func NewStocksTrie(root common.Hash, db ethdb.Database) (*trie.Trie, error) {
	return trie.NewTrieWithPrefix(root, protocol.StocksPrefix, db)
}

func NewOwnerTrie(root common.Hash, db ethdb.Database) (*trie.Trie, error) {
	return trie.NewTrieWithPrefix(root, protocol.OwnerPrefix, db)
}

func NewGasPoolTrie(root common.Hash, db ethdb.Database) (*trie.Trie, error) {
	return trie.NewTrieWithPrefix(root, protocol.GasPoolPrefix, db)
}

func NewBokerContext(db ethdb.Database) (*BokerContext, error) {

	//log.Info("Create Tinachain Single Stock Trie")
	singleContractsTrie, err := NewSingleContractTrie(common.Hash{}, db)
	if err != nil {
		log.Error("Create Tinachain Single Stock Trie", "err", err)
		return nil, err
	}

	contractsTrie, err := NewContractsTrie(common.Hash{}, db)
	if err != nil {
		log.Error("Create Tinachain Single Stock Trie", "err", err)
		return nil, err
	}

	//log.Info("Create Tinachain Single Stock Trie")
	singleStockTrie, err := NewSingleStockTrie(common.Hash{}, db)
	if err != nil {
		log.Error("Create Tinachain Single Stock Trie", "err", err)
		return nil, err
	}

	//log.Info("Create Tinachain Stocks Trie")
	stocksTrie, err := NewStocksTrie(common.Hash{}, db)
	if err != nil {
		log.Error("Create Tinachain Stocks Trie", "err", err)
		return nil, err
	}

	//log.Info("Create Tinachain Owner Trie")
	ownerTrie, err := NewOwnerTrie(common.Hash{}, db)
	if err != nil {
		log.Error("Create Tinachain Owner Trie", "err", err)
		return nil, err
	}

	log.Info("Create Tinachain Gas Pool Trie")
	gasPoolTrie, err := NewGasPoolTrie(common.Hash{}, db)
	if err != nil {
		log.Error("Create Tinachain Gas Pool Trie", "err", err)
		return nil, err
	}

	return &BokerContext{
		singleContractsTrie: singleContractsTrie,
		contractsTrie:       contractsTrie,
		singleStockTrie:     singleStockTrie,
		stocksTrie:          stocksTrie,
		ownerTrie:           ownerTrie,
		gasPoolTrie:         gasPoolTrie,
		db:                  db,
	}, nil
}

func NewBokerContextFromProto(db ethdb.Database, ctxProto *BokerBackendProto) (*BokerContext, error) {

	//log.Info("Create Tinachain Single Stock Trie")
	singleContractsTrie, err := NewSingleContractTrie(ctxProto.SingleHash, db)
	if err != nil {
		log.Error("Create Tinachain Single Stock Trie", "err", err)
		return nil, err
	}

	contractsTrie, err := NewContractsTrie(ctxProto.ContractsHash, db)
	if err != nil {
		log.Error("Create Tinachain Single Stock Trie", "err", err)
		return nil, err
	}

	//log.Info("Create Tinachain Single Stock Trie", "Hash", ctxProto.SingleStockHash.String())
	singleStockTrie, err := NewSingleStockTrie(ctxProto.SingleStockHash, db)
	if err != nil {
		log.Error("Create Tinachain Single Stock Trie", "err", err)
		return nil, err
	}

	//log.Info("Create Tinachain Stocks Trie", "Hash", ctxProto.StocksHash.String())
	stocksTrie, err := NewStocksTrie(ctxProto.StocksHash, db)
	if err != nil {
		log.Error("Create Tinachain Stocks Trie", "err", err)
		return nil, err
	}

	//log.Info("Create Tinachain Owner Trie", "Hash", ctxProto.OwnerHash.String())
	ownerTrie, err := NewOwnerTrie(ctxProto.OwnerHash, db)
	if err != nil {
		log.Error("Create Tinachain Owner Trie", "err", err)
		return nil, err
	}

	//log.Info("Create Tinachain Gas Pool Trie", "Hash", ctxProto.GasPoolHash.String())
	gasPoolTrie, err := NewGasPoolTrie(ctxProto.GasPoolHash, db)
	if err != nil {
		log.Error("Create Tinachain Gas Pool Trie", "err", err)
		return nil, err
	}

	return &BokerContext{
		singleContractsTrie: singleContractsTrie,
		contractsTrie:       contractsTrie,
		singleStockTrie:     singleStockTrie,
		stocksTrie:          stocksTrie,
		ownerTrie:           ownerTrie,
		gasPoolTrie:         gasPoolTrie,
		db:                  db,
	}, nil
}

func (s *BokerContext) Copy() *BokerContext {

	singleContractsTrie := *s.singleContractsTrie
	contractsTrie := *s.contractsTrie
	singleStockTrie := *s.singleStockTrie
	stocksTrie := *s.stocksTrie
	ownerTrie := *s.ownerTrie
	gasPoolTrie := *s.gasPoolTrie

	return &BokerContext{
		singleContractsTrie: &singleContractsTrie,
		contractsTrie:       &contractsTrie,
		singleStockTrie:     &singleStockTrie,
		stocksTrie:          &stocksTrie,
		ownerTrie:           &ownerTrie,
		gasPoolTrie:         &gasPoolTrie,
	}
}

func (s *BokerContext) Root() (h common.Hash) {

	hw := sha3.NewKeccak256()
	rlp.Encode(hw, s.singleContractsTrie.Hash())
	rlp.Encode(hw, s.contractsTrie.Hash())
	rlp.Encode(hw, s.singleStockTrie.Hash())
	rlp.Encode(hw, s.stocksTrie.Hash())
	rlp.Encode(hw, s.ownerTrie.Hash())
	rlp.Encode(hw, s.gasPoolTrie.Hash())
	hw.Sum(h[:0])
	return h
}

func (s *BokerContext) Snapshot() *BokerContext {
	return s.Copy()
}

func (s *BokerContext) RevertToSnapShot(snapshot *BokerContext) {

	s.singleContractsTrie = snapshot.singleContractsTrie
	s.contractsTrie = snapshot.contractsTrie
	s.singleStockTrie = snapshot.singleStockTrie
	s.stocksTrie = snapshot.stocksTrie
	s.ownerTrie = snapshot.ownerTrie
	s.gasPoolTrie = snapshot.gasPoolTrie
}

func (s *BokerContext) FromProto(dcp *BokerBackendProto) error {

	var err error

	s.singleContractsTrie, err = NewSingleStockTrie(dcp.SingleHash, s.db)
	if err != nil {
		return err
	}

	s.contractsTrie, err = NewStocksTrie(dcp.ContractsHash, s.db)
	if err != nil {
		return err
	}

	s.singleStockTrie, err = NewSingleStockTrie(dcp.SingleStockHash, s.db)
	if err != nil {
		return err
	}

	s.stocksTrie, err = NewStocksTrie(dcp.StocksHash, s.db)
	if err != nil {
		return err
	}

	s.ownerTrie, err = NewOwnerTrie(dcp.OwnerHash, s.db)
	if err != nil {
		return err
	}

	s.gasPoolTrie, err = NewGasPoolTrie(dcp.GasPoolHash, s.db)
	return err
}

func (s *BokerContext) ToProto() *BokerBackendProto {
	return &BokerBackendProto{

		SingleHash:      s.singleContractsTrie.Hash(),
		ContractsHash:   s.contractsTrie.Hash(),
		SingleStockHash: s.singleStockTrie.Hash(),
		StocksHash:      s.stocksTrie.Hash(),
		OwnerHash:       s.ownerTrie.Hash(),
		GasPoolHash:     s.gasPoolTrie.Hash(),
	}
}

func (s *BokerContext) CommitTo(dbw trie.DatabaseWriter) (*BokerBackendProto, error) {

	singleContractsRoot, err := s.singleContractsTrie.CommitTo(dbw)
	if err != nil {
		return nil, err
	}

	contractsRoot, err := s.contractsTrie.CommitTo(dbw)
	if err != nil {
		return nil, err
	}

	singleStockRoot, err := s.singleStockTrie.CommitTo(dbw)
	if err != nil {
		return nil, err
	}

	stocksRoot, err := s.stocksTrie.CommitTo(dbw)
	if err != nil {
		return nil, err
	}

	ownerRoot, err := s.ownerTrie.CommitTo(dbw)
	if err != nil {
		return nil, err
	}

	gasPoolRoot, err := s.gasPoolTrie.CommitTo(dbw)
	if err != nil {
		return nil, err
	}

	return &BokerBackendProto{
		SingleHash:      singleContractsRoot,
		ContractsHash:   contractsRoot,
		SingleStockHash: singleStockRoot,
		StocksHash:      stocksRoot,
		OwnerHash:       ownerRoot,
		GasPoolHash:     gasPoolRoot,
	}, nil
}

func (s *BokerContext) SingleContractsTrie() *trie.Trie           { return s.singleContractsTrie }
func (s *BokerContext) ContractsTrie() *trie.Trie                 { return s.contractsTrie }
func (s *BokerContext) SingleStockTrie() *trie.Trie               { return s.singleStockTrie }
func (s *BokerContext) StocksTrie() *trie.Trie                    { return s.stocksTrie }
func (s *BokerContext) OwnerTrie() *trie.Trie                     { return s.ownerTrie }
func (s *BokerContext) GasPoolTrie() *trie.Trie                   { return s.gasPoolTrie }
func (s *BokerContext) SetContracts(contractsTrie *trie.Trie)     { s.contractsTrie = contractsTrie }
func (s *BokerContext) SetSingleStock(singleStockTrie *trie.Trie) { s.singleStockTrie = singleStockTrie }
func (s *BokerContext) SetStocks(stocksTrie *trie.Trie)           { s.stocksTrie = stocksTrie }
func (s *BokerContext) SetOwner(ownerTrie *trie.Trie)             { s.ownerTrie = ownerTrie }
func (s *BokerContext) SetGasPool(gasPoolTrie *trie.Trie)         { s.gasPoolTrie = gasPoolTrie }
func (s *BokerContext) DB() ethdb.Database                        { return s.db }
func (s *BokerContext) SetSingleContracts(singleContractsTrie *trie.Trie) {
	s.singleContractsTrie = singleContractsTrie
}

//Trie
func (s *BokerContext) setContractsTrie(address []common.Address) error {

	//转换成切片数组
	log.Info("(s *BokerContext) SetContractsTrie")

	var contracts []common.Address = make([]common.Address, 0)
	for _, v := range address {
		contracts = append(contracts, v)
	}
	contractsRLP, err := rlp.EncodeToBytes(contracts)
	if err != nil {

		log.Error("failed to encode contracts to rlp", "error", err)
		return err
	}
	s.contractsTrie.Update(protocol.ContractsPrefix, contractsRLP)
	s.contractsTrie.Commit()
	return nil
}

func (s *BokerContext) getContractsTrie() ([]common.Address, error) {

	log.Info("(s *v) GetContractsTrie")

	var contracts []common.Address
	if s.contractsTrie == nil {
		log.Error("contract Trie is nil")
		return []common.Address{}, protocol.ErrPointerIsNil
	}

	contractsRLP, err := s.contractsTrie.TryGet(protocol.ContractsPrefix)
	if err != nil {
		return []common.Address{}, err
	}

	if string(contractsRLP[:]) != "" {
		if err := rlp.DecodeBytes(contractsRLP, &contracts); err != nil {
			log.Error("failed to decode contracts", "error", err)
			return []common.Address{}, err
		}

		return contracts, nil
	}
	return []common.Address{}, nil
}

func (s *BokerContext) setSingleContractsTrie(address common.Address, contractType protocol.BaseContractType) error {

	log.Info("(s *BokerContext) SetSingleContractsTrie")

	if err := s.singleContractsTrie.TryUpdate(address.Bytes(), []byte(strconv.Itoa(int(contractType)))); err != nil {
		log.Error("SetSystemBaseContract singleContractsTrie.TryUpdate failed", "err", err)
		return err
	}
	s.singleContractsTrie.Commit()
	return nil
}

func (s *BokerContext) getSingleContractsTrie(address common.Address) (protocol.BaseContractType, error) {

	log.Info("(s *BokerContext) GetSingleContractsTrie")

	key := address.Bytes()
	v, err := s.singleContractsTrie.TryGet(key)
	if err != nil {
		return protocol.System, err
	}
	contractType, err := strconv.Atoi(string(v[:]))
	return protocol.BaseContractType(contractType), nil
}

func (s *BokerContext) deleteSingleContractsTrie(address common.Address) error {

	log.Info("(s *BokerContext) DeleteSingleContractsTrie")

	if err := s.singleContractsTrie.TryDelete(address.Bytes()); err != nil {
		return err
	}
	s.singleContractsTrie.Commit()

	return nil
}

func (s *BokerContext) getStocksTrie() []common.Address {

	log.Info("(s *BokerContext) GetStocksTrie")

	if s.stocksTrie == nil {
		log.Error("stocksTrie is nil")
		return []common.Address{}
	}

	stocksRLP, err := s.stocksTrie.TryGet(protocol.StocksPrefix)
	if err != nil {
		return []common.Address{}
	}

	if string(stocksRLP[:]) == "" {
		return []common.Address{}
	}

	var stocks []common.Address
	if err := rlp.DecodeBytes(stocksRLP, &stocks); err != nil {
		log.Error("failed to decode stocksRLP", "error", err)
		return []common.Address{}
	}

	return stocks
}

func (s *BokerContext) setStocksTrie(accounts []common.Address) error {

	log.Info("(s *BokerContext) SetStocksTrie")

	accountsRLP, err := rlp.EncodeToBytes(accounts)
	if err != nil {

		log.Error("failed to encode stocksRLP to rlp", "error", err)
		return err
	}
	s.stocksTrie.Update(protocol.StocksPrefix, accountsRLP)
	s.stocksTrie.Commit()
	return nil
}

func (s *BokerContext) setSingleStockTrie(stock *protocol.StockAccount) error {

	log.Info("(s *BokerContext) SetSingleStockTrie")

	stockRLP, err := rlp.EncodeToBytes(*stock)

	if err = s.singleStockTrie.TryUpdate(stock.Account.Bytes(), stockRLP); err != nil {

		log.Error("SetSingleStockTrie singleTrie.TryUpdate failed", "err", err)
		return err
	}
	s.singleStockTrie.Commit()

	return nil
}

func (s *BokerContext) getSingleStockTrie(address common.Address) (common.Address, uint64, protocol.StockState, error) {

	log.Info("(s *BokerContext) GetSingleStockTrie")

	if s.singleStockTrie == nil {

		log.Error("(s *BokerContext) GetSingleStockTrie is nil")
		return common.Address{}, 0, protocol.Run, protocol.ErrPointerIsNil
	}

	singleStockRLP, err := s.singleStockTrie.TryGet(address.Bytes())
	if err != nil {

		log.Error("(s *BokerContext) GetSingleStockTrie", "err", err)
		return common.Address{}, 0, protocol.Run, err
	}

	if string(singleStockRLP[:]) == "" {

		log.Error("(s *BokerContext) GetSingleStockTrie Stock Account is Nil")
		return common.Address{}, 0, protocol.Run, errors.New("Stock Account is Nil")
	}

	var stock protocol.StockAccount
	if err := rlp.DecodeBytes(singleStockRLP, &stock); err != nil {

		log.Error("failed to decode singleStockRLP", "error", err)
		return common.Address{}, 0, protocol.Run, err
	}

	return stock.Account, stock.Number, stock.State, nil
}

func (s *BokerContext) stockExists(address common.Address) bool {

	log.Info("(s *BokerContext) StockExists")

	if s.singleStockTrie == nil {

		log.Error("(s *BokerContext) StockExists is nil")
		return false
	}

	singleStockRLP, err := s.singleStockTrie.TryGet(address.Bytes())
	if err != nil {

		log.Error("(s *BokerContext) StockExists", "err", err)
		return false
	}

	if string(singleStockRLP[:]) == "" {

		return false
	}

	var stock *protocol.StockAccount
	if err := rlp.DecodeBytes(singleStockRLP, stock); err != nil {

		log.Error("(s *BokerContext) StockExists failed to decode singleStockRLP", "error", err)
		return false
	}

	return true
}

func (s *BokerContext) deleteSingleStockTrie(address common.Address) error {

	log.Info("(s *BokerContext) DeleteSingleStockTrie")

	if err := s.singleStockTrie.TryDelete(address.Bytes()); err != nil {
		return err
	}
	s.singleStockTrie.Commit()

	return nil
}

func (s *BokerContext) setGasPoolTrie(gasPool uint64) error {

	//log.Info("(s *BokerContext) SetGasPoolTrie", "gasPool", gasPool)

	gasPoolRLP, err := rlp.EncodeToBytes(new(big.Int).SetUint64(gasPool))
	if err != nil {

		log.Error("failed to encode gasPoolRLP to rlp", "error", err)
		return err
	}
	s.gasPoolTrie.Update(protocol.GasPoolPrefix, gasPoolRLP)
	s.gasPoolTrie.Commit()

	_, getErr := s.getGasPoolTrie()
	if getErr != nil {

		log.Error("failed to encode gasPoolRLP to rlp", "error", err)
		return err
	}
	//log.Info("(s *BokerContext) SetGasPoolTrie Get Current Gas Pool", "gas", gas)

	return nil
}

func (s *BokerContext) getGasPoolTrie() (uint64, error) {

	if s.gasPoolTrie == nil {
		log.Error("gasPoolTrie is nil")
		return 0, protocol.ErrPointerIsNil
	}

	gasPoolRLP, err := s.gasPoolTrie.TryGet(protocol.GasPoolPrefix)
	if err != nil {
		return 0, err
	}
	if string(gasPoolRLP[:]) != "" {

		var gasPool *big.Int
		if err := rlp.DecodeBytes(gasPoolRLP, &gasPool); err != nil {
			log.Error("failed to decode gasPoolRLP", "error", err)
			return 0, err
		}
		return gasPool.Uint64(), nil
	}
	return 0, nil
}

func (s *BokerContext) setOwnerTrie(owner common.Address) error {

	log.Info("(s *BokerContext) SetOwnerTrie")

	ownerRLP, err := rlp.EncodeToBytes(owner)
	if err != nil {

		log.Error("failed to encode contracts to rlp", "error", err)
		return err
	}
	s.ownerTrie.Update(protocol.OwnerPrefix, ownerRLP)
	s.ownerTrie.Commit()
	return nil
}

func (s *BokerContext) getOwnerTrie() (common.Address, error) {

	log.Info("(s *BokerContext) GetOwnerTrie")

	if s.ownerTrie == nil {
		log.Error("ownerTrie is nil")
		return common.Address{}, protocol.ErrPointerIsNil
	}

	ownerRLP, err := s.ownerTrie.TryGet(protocol.OwnerPrefix)
	if err != nil {
		return common.Address{}, err
	}

	if string(ownerRLP[:]) != "" {
		var ownerAccount common.Address
		if err := rlp.DecodeBytes(ownerRLP, &ownerAccount); err != nil {
			log.Error("failed to decode ownerRLP", "error", err)
			return common.Address{}, err
		}
		//log.Info("(s *BokerContext) getOwnerTrie", "ownerAccount", ownerAccount)
		return ownerAccount, nil
	}
	return common.Address{}, nil
}

//**************

func (s *BokerContext) GetGasPool() uint64 {

	gasPool, err := s.getGasPoolTrie()
	if err != nil {
		return 0
	}
	return gasPool
}

func (s *BokerContext) AddGasPool(gas uint64) uint64 {

	gasPool := s.GetGasPool() + gas
	s.setGasPoolTrie(gasPool)

	//log.Info("(s *BokerContext) AddGasPool", "gas", gas, "gasPool", gasPool)

	return gasPool
}

func (s *BokerContext) GetStockManager() common.Address {

	owner, err := s.getOwnerTrie()
	if err != nil {
		return common.StringToAddress("")
	}
	return owner
}

func (s *BokerContext) SetStockManager(operation common.Address,
	address common.Address,
	now int64,
	dposContext *DposContext) error {

	log.Info("(s *BokerContext) SetStockManager", "operation", operation.String(), "address", address.String())

	owner := s.GetStockManager()
	if owner == common.StringToAddress("") {

		producer, err := dposContext.GetCurrentNowProducer(0, now)
		if err != nil {
			return protocol.ErrInvalidProducer
		}
		if producer != operation {
			return protocol.ErrInvalidProducer
		}

		s.setOwnerTrie(address)
	} else {

		s.setOwnerTrie(address)
	}

	owner = s.GetStockManager()
	log.Info("(s *BokerContext) SetStockManager get Current Stock Manager", "owner", owner.String())

	return nil
}

func (s *BokerContext) GetStock(address common.Address) *protocol.StockAccount {

	stockAddress, number, stockState, err := s.getSingleStockTrie(address)
	if err != nil {
		return nil
	}
	return &protocol.StockAccount{
		Account: stockAddress,
		Number:  number,
		State:   stockState,
	}
}

func (s *BokerContext) GetStocks() []*protocol.StockAccount {

	stockAccounts := s.getStocksTrie()

	var stocks []*protocol.StockAccount
	for _, v := range stockAccounts {

		stockAddress, number, stockState, err := s.getSingleStockTrie(v)
		if err != nil {
			continue
		}

		a := new(protocol.StockAccount)
		a.Account = stockAddress
		a.Number = number
		a.State = stockState
		stocks = append(stocks, a)
	}
	return stocks
}

func (s *BokerContext) SetStock(operation, address common.Address, number uint64) error {

	log.Info("(s *BokerContext) SetStock", "operation", operation.String(), "address", address.String(), "number", number)

	owner := s.GetStockManager()
	if owner != operation {
		return errors.New("Operation Account Is not Owner")
	}

	exists := s.stockExists(address)
	if exists {
		log.Error("(s *BokerContext) SetStock CheckSingleStockTrie Exists")
		return errors.New("Stock Account is not Nil")
	}

	//得到已有股权
	stockAccounts := s.getStocksTrie()
	var stock = new(protocol.StockAccount)
	stock.Account = address
	stock.Number = number
	stock.State = protocol.Run
	stockAccounts = append(stockAccounts, address)

	//写入股权数组
	if err := s.setStocksTrie(stockAccounts); err != nil {
		return errors.New("Stock Account is not Nil")
	}
	stocks := s.getStocksTrie()
	for i, v := range stocks {
		log.Info("(c *BokerContracts) SetStock GetStocks", "i", i, "v", v)
	}
	log.Info("(c *BokerContracts) SetStock", "stocks", len(stocks))

	//写入单个股权
	if err := s.setSingleStockTrie(stock); err != nil {
		return errors.New("Stock Account is not Nil")
	}
	stockAddress, getNumber, _, getErr := s.getSingleStockTrie(address)
	if getErr != nil {

		log.Error("(c *BokerContracts) SetStock GetSingleStockTrie Failed")
		return errors.New("Stock Account is not Nil")
	}
	log.Info("(c *BokerContracts) SetStock", "address", stockAddress, "number", getNumber)

	return nil
}

func (s *BokerContext) TransferStock(from common.Address, to common.Address, number uint64) error {

	fromStock := s.GetStock(from)
	if nil == fromStock {
		return errors.New("From Account Not Found Stock")
	}

	if protocol.Frozen == fromStock.State {
		return errors.New("From Stock State Frozen")
	}

	if fromStock.Number < number {
		return errors.New("From Stock Number Too Low")
	}

	toStock := s.GetStock(to)
	if nil == toStock {

		fromStock.Number = fromStock.Number - number
		s.SetStock(from, to, number)
	}
	return nil
}

func (s *BokerContext) CleanStock(operation common.Address, from common.Address) error {

	owner := s.GetStockManager()
	if owner != operation {
		return errors.New("Operation Account Is not Owner")
	}

	fromStock := s.GetStock(from)
	if nil == fromStock {
		return errors.New("From Account Not Found Stock")
	}

	s.deleteSingleStockTrie(from)

	stockAccounts := s.getStocksTrie()

	var postion int
	for i, v := range stockAccounts {
		if v == from {
			postion = i
			break
		}
	}
	stockAccounts = append(stockAccounts[:postion], stockAccounts[postion+1:]...)
	s.setStocksTrie(stockAccounts)
	return nil
}

func (s *BokerContext) FrozenStock(operation common.Address, from common.Address) error {

	owner := s.GetStockManager()
	if owner != operation {
		return errors.New("Operation Account Is not Owner")
	}

	fromStock := s.GetStock(from)
	if nil == fromStock {
		return errors.New("From Account Not Found Stock")
	}

	fromStock.State = protocol.Frozen
	s.setSingleStockTrie(fromStock)
	return nil
}

func (s *BokerContext) UnFrozenStock(operation, from common.Address) error {

	owner := s.GetStockManager()
	if owner != operation {
		return errors.New("Operation Account Is not Owner")
	}

	fromStock := s.GetStock(from)
	if nil == fromStock {
		return errors.New("From Account Not Found Stock")
	}

	fromStock.State = protocol.Run
	s.setSingleStockTrie(fromStock)

	return nil
}

func (s *BokerContext) ClearStock(operation common.Address) error {

	return nil
}

//*****************
func (s *BokerContext) existContract(address common.Address) (bool, error) {

	contracts, err := s.getContractsTrie()
	if err != nil {
		return false, protocol.ErrNotFoundContract
	}
	for _, v := range contracts {
		if v == address {
			return true, nil
		}
	}
	return false, protocol.ErrNotFoundContract
}

func (s *BokerContext) GetSystemContractAddress() (common.Address, error) {

	contracts, err := s.getContractsTrie()
	if err != nil {
		return common.StringToAddress(""), err
	}
	for _, v := range contracts {

		contractType, err := s.getSingleContractsTrie(v)
		if err != nil {
			return common.StringToAddress(""), err
		}
		if contractType == protocol.System {
			return v, nil
		}
	}
	return common.StringToAddress(""), protocol.ErrNotFoundContract
}

func (s *BokerContext) SetSystemContract(address common.Address, from common.Address) error {

	log.Info("(c *BokerContracts) SetSystemBaseContract", "address", address.String(), "from", from.String())

	contracts, err := s.getContractsTrie()
	if err != nil {
		return err
	}
	for _, v := range contracts {
		if v == address {
			return protocol.ErrContractExist
		}
	}

	contractType, _ := s.getSingleContractsTrie(address)
	if contractType == protocol.System {
		return protocol.ErrContractExist
	}

	contracts = append(contracts, address)
	if err := s.setSingleContractsTrie(address, protocol.System); err != nil {
		return err
	}
	if err := s.setContractsTrie(contracts); err != nil {
		return err
	}

	return nil
}

func (s *BokerContext) SetUserBaseContract(address common.Address, from common.Address) error {

	log.Info("(c *BokerContracts) SetUserBaseContract", "address", address.String(), "from", from.String())

	contracts, err := s.getContractsTrie()
	if err != nil {
		return err
	}
	for _, v := range contracts {
		if v == address {
			return protocol.ErrContractExist
		}
	}

	contracts = append(contracts, address)
	if err := s.setSingleContractsTrie(address, protocol.User); err != nil {
		return err
	}
	if err := s.setContractsTrie(contracts); err != nil {
		return err
	}

	return nil
}

func (s *BokerContext) CancelUserBaseContract(address common.Address, from common.Address) error {

	log.Info("(c *BokerContracts) CancelUserBaseContract")

	contracts, err := s.getContractsTrie()
	if err != nil {
		return err
	}

	var exist bool = false
	var postion int
	for i, v := range contracts {
		if v == address {
			postion = i
			exist = true
			break
		}
	}
	if !exist {
		return protocol.ErrNotFoundContract
	}

	contractType, _ := s.getSingleContractsTrie(address)
	if contractType != protocol.User {
		return protocol.ErrNotFoundContract
	}
	if err := s.deleteSingleContractsTrie(address); err != nil {
		return err
	}

	contracts = append(contracts[:postion], contracts[postion+1:]...)
	if err := s.setContractsTrie(contracts); err != nil {
		return err
	}
	return nil
}

func (s *BokerContext) GetSingleContractsType(address common.Address) (protocol.BaseContractType, error) {
	return s.getSingleContractsTrie(address)
}
