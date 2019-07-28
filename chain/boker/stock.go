package boker

import (
	"errors"

	"github.com/Tinachain/Tina/chain/boker/protocol"
	"github.com/Tinachain/Tina/chain/common"
	"github.com/Tinachain/Tina/chain/eth"
	"github.com/Tinachain/Tina/chain/ethdb"
	"github.com/Tinachain/Tina/chain/log"
	"github.com/Tinachain/Tina/chain/rlp"
	"github.com/Tinachain/Tina/chain/trie"
)

type TinaStock struct {
	singleStockTrie *trie.Trie
	stocksTrie      *trie.Trie
	ownerTrie       *trie.Trie
	gasPoolTrie     *trie.Trie
	gasPool         uint64
	ownerAccount    common.Address
	stockAccounts   map[common.Address]*protocol.StockAccount
}

func NewStock(db ethdb.Database, ethereum *eth.Ethereum, bokerProto *protocol.BokerBackendProto) *TinaStock {

	log.Info("****NewStock****")

	log.Info("Create Tinachain Stocks Trie", "Hash", bokerProto.StocksHash.String())

	singleStockTrie, err := NewSingleStockTrie(bokerProto.GasPoolHash, db)
	if err != nil {
		return nil
	}

	stocksTrie, err := NewStocksTrie(bokerProto.StocksHash, db)
	if err != nil {
		return nil
	}

	ownerTrie, err := NewOwnerTrie(bokerProto.OwnerHash, db)
	if err != nil {
		return nil
	}

	gasPoolTrie, err := NewGasPoolTrie(bokerProto.GasPoolHash, db)
	if err != nil {
		return nil
	}

	return &TinaStock{
		stocksTrie:      stocksTrie,
		ownerTrie:       ownerTrie,
		gasPoolTrie:     gasPoolTrie,
		singleStockTrie: singleStockTrie,
		gasPool:         0,
		ownerAccount:    common.Address{},
		stockAccounts:   make(map[common.Address]*protocol.StockAccount),
	}
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

func (s *TinaStock) loadTrie() error {

	log.Info("(s *TinaStock) loadTrie")

	gasPool, err := s.getGasPoolTrie()
	if err != nil {
		return err
	}
	s.gasPool = gasPool

	ownerAccount, err := s.getOwnerTrie()
	if err != nil {
		return err
	}
	s.ownerAccount = ownerAccount

	stocksAddress, err := s.getStocksTrie()
	if err != nil {
		return err
	}
	for _, v := range stocksAddress {

		var err error
		stock := new(protocol.StockAccount)
		stock.Account, stock.Number, stock.State, err = s.getSingleStockTrie(v)
		if err != nil {
			return err
		}
		s.stockAccounts[stock.Account] = stock
	}

	return nil
}

func (s *TinaStock) GetOwner() common.Address {

	return s.ownerAccount
}

func (s *TinaStock) SetOwner(operation common.Address, address common.Address) error {

	if s.ownerAccount != operation {
		return errors.New("Operation Account Is not Owner")
	}

	if err := s.ownerTrie.TryUpdate(address.Bytes(), []byte("")); err != nil {
		log.Error("SetOwner ownerTrie.TryUpdate failed", "err", err)
		return err
	}
	s.ownerAccount = address

	return nil
}

func (s *TinaStock) setOwnerTrie() error {

	log.Info("(s *TinaStock) setOwnerTrie")

	ownerRLP, err := rlp.EncodeToBytes(s.ownerAccount)
	if err != nil {

		log.Error("failed to encode contracts to rlp", "error", err)
		return err
	}
	s.ownerTrie.Update(protocol.OwnerPrefix, ownerRLP)
	return nil
}

func (s *TinaStock) getOwnerTrie() (common.Address, error) {

	log.Info("(s *TinaStock) getOwnerTrie")

	if s.ownerTrie == nil {
		log.Error("ownerTrie is nil")
		return common.Address{}, protocol.ErrPointerIsNil
	}

	ownerRLP, err := s.gasPoolTrie.TryGet(protocol.OwnerPrefix)
	if err != nil {
		return common.Address{}, err
	}

	var ownerAccount common.Address
	if err := rlp.DecodeBytes(ownerRLP, &ownerAccount); err != nil {
		log.Error("failed to decode ownerRLP", "error", err)
		return common.Address{}, err
	}
	return ownerAccount, nil
}

func (s *TinaStock) GetGasPool() uint64 {

	return s.gasPool
}

func (s *TinaStock) AddGasPool(gas uint64) uint64 {

	s.gasPool = s.gasPool + gas
	s.setGasPoolTrie()
	return s.gasPool
}

func (s *TinaStock) setGasPoolTrie() error {

	log.Info("(s *TinaStock) setGasPoolTrie")

	gasPoolRLP, err := rlp.EncodeToBytes(s.gasPool)
	if err != nil {

		log.Error("failed to encode gasPoolRLP to rlp", "error", err)
		return err
	}
	s.gasPoolTrie.Update(protocol.GasPoolPrefix, gasPoolRLP)
	return nil
}

func (s *TinaStock) getGasPoolTrie() (uint64, error) {

	log.Info("(s *TinaStock) getGasPoolTrie")

	if s.gasPoolTrie == nil {
		log.Error("gasPoolTrie is nil")
		return 0, protocol.ErrPointerIsNil
	}

	gasPoolRLP, err := s.gasPoolTrie.TryGet(protocol.GasPoolPrefix)
	if err != nil {
		return 0, err
	}

	var gasPool uint64
	if err := rlp.DecodeBytes(gasPoolRLP, &gasPool); err != nil {
		log.Error("failed to decode gasPoolRLP", "error", err)
		return 0, err
	}
	return gasPool, nil
}

func (s *TinaStock) GetStock(address common.Address) *protocol.StockAccount {

	if account, ok := s.stockAccounts[address]; ok {
		return account
	}
	return nil
}

func (s *TinaStock) GetStocks() ([]*protocol.StockAccount, error) {

	var accounts []*protocol.StockAccount
	for _, v := range s.stockAccounts {
		accounts = append(accounts, v)
	}
	return accounts, nil
}

func (s *TinaStock) SetStock(operation common.Address, address common.Address, number uint64) error {

	if operation != s.ownerAccount {
		return errors.New("Operation Account Is not Owner")
	}

	stock := s.GetStock(address)
	if stock != nil {
		return errors.New("Stock Account is not Nil")
	}

	stock = new(protocol.StockAccount)
	stock.Account = address
	stock.Number = number
	stock.State = protocol.Run
	s.stockAccounts[address] = stock

	if err := s.setSingleStockTrie(stock); err != nil {
		return errors.New("Stock Account is not Nil")
	}
	if err := s.setStocksTrie(); err != nil {
		return errors.New("Stock Account is not Nil")
	}

	return nil
}

func (s *TinaStock) setSingleStockTrie(stock *protocol.StockAccount) error {

	log.Info("(s *TinaStock) setSingleStockTrie")

	stockRLP, err := rlp.EncodeToBytes(stock)

	if err = s.singleStockTrie.TryUpdate(stock.Account.Bytes(), stockRLP); err != nil {
		log.Error("setSingleStockTrie singleTrie.TryUpdate failed", "err", err)
		return err
	}

	return nil
}

func (s *TinaStock) getSingleStockTrie(address common.Address) (common.Address, uint64, protocol.StockState, error) {

	log.Info("(s *TinaStock) getSingleStockTrie")

	if s.singleStockTrie == nil {
		log.Error("singleStockTrie is nil")
		return common.Address{}, 0, protocol.Run, protocol.ErrPointerIsNil
	}

	singleStockRLP, err := s.singleStockTrie.TryGet(address.Bytes())
	if err != nil {
		return common.Address{}, 0, protocol.Run, err
	}

	var stock *protocol.StockAccount
	if err := rlp.DecodeBytes(singleStockRLP, stock); err != nil {
		log.Error("failed to decode singleStockRLP", "error", err)
		return common.Address{}, 0, protocol.Run, err
	}

	return stock.Account, stock.Number, stock.State, nil
}

func (s *TinaStock) setStocksTrie() error {

	log.Info("(s *TinaStock) setStocksTrie")

	var stocks []common.Address = make([]common.Address, 0)
	for k, _ := range s.stockAccounts {
		stocks = append(stocks, k)
	}

	stocksRLP, err := rlp.EncodeToBytes(stocks)
	if err != nil {

		log.Error("failed to encode stocksRLP to rlp", "error", err)
		return err
	}
	s.gasPoolTrie.Update(protocol.StocksPrefix, stocksRLP)
	return nil
}

func (s *TinaStock) getStocksTrie() ([]common.Address, error) {

	log.Info("(s *TinaStock) getStocksTrie")

	if s.stocksTrie == nil {
		log.Error("stocksTrie is nil")
		return []common.Address{}, protocol.ErrPointerIsNil
	}

	stocksRLP, err := s.stocksTrie.TryGet(protocol.StocksPrefix)
	if err != nil {
		return []common.Address{}, err
	}

	var stocks []common.Address
	if err := rlp.DecodeBytes(stocksRLP, &stocks); err != nil {
		log.Error("failed to decode stocksRLP", "error", err)
		return []common.Address{}, err
	}

	return stocks, nil
}

func (s *TinaStock) TransferStock(operation common.Address, from common.Address, to common.Address, number uint64) error {

	//检测操作账号是否是Owner
	if operation != s.ownerAccount {
		return errors.New("Operation Account Is not Owner")
	}

	//检测from身上是否有相应的股权
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

	//判断to账号是否存在
	toStock := s.GetStock(to)
	if nil == toStock {

		fromStock.Number = fromStock.Number - number
		s.SetStock(operation, to, number)
	} else {

		fromStock.Number = fromStock.Number - number
		toStock.Number = toStock.Number + number
	}
	return nil
}

func (s *TinaStock) CleanStock(operation common.Address, from common.Address) error {

	//检测操作账号是否是Owner
	if operation != s.ownerAccount {
		return errors.New("Operation Account Is not Owner")
	}

	fromStock := s.GetStock(from)
	if nil == fromStock {
		return errors.New("From Account Not Found Stock")
	}

	delete(s.stockAccounts, from)
	return nil
}

func (s *TinaStock) FrozenStock(operation common.Address, from common.Address) error {

	//检测操作账号是否是Owner
	if operation != s.ownerAccount {
		return errors.New("Operation Account Is not Owner")
	}

	fromStock := s.GetStock(from)
	if nil == fromStock {
		return errors.New("From Account Not Found Stock")
	}

	fromStock.State = protocol.Frozen
	return nil
}

func (s *TinaStock) UnFrozenStock(operation common.Address, from common.Address) error {

	//检测操作账号是否是Owner
	if operation != s.ownerAccount {
		return errors.New("Operation Account Is not Owner")
	}

	fromStock := s.GetStock(from)
	if nil == fromStock {
		return errors.New("From Account Not Found Stock")
	}

	fromStock.State = protocol.Run
	return nil
}

func (s *TinaStock) ClearStock(operation common.Address) error {

	return nil
}

func (s *TinaStock) GetStockTrie() (*trie.Trie, *trie.Trie, *trie.Trie, *trie.Trie) {

	return s.singleStockTrie, s.stocksTrie, s.ownerTrie, s.gasPoolTrie

}
