//Tina链增加的特殊账号管理类
package boker

import (
	"errors"

	"github.com/Tinachain/Tina/chain/boker/protocol"
	"github.com/Tinachain/Tina/chain/common"
	_ "github.com/Tinachain/Tina/chain/log"
)

type StockAccount struct {
	Account common.Address
	Number  uint64
	State   protocol.StockState
}

type TinaStock struct {
	gasPool       uint64
	ownerAccount  common.Address
	stockAccounts map[common.Address]*StockAccount
}

func NewStock() *TinaStock {

	return &TinaStock{
		gasPool:       0,
		ownerAccount:  common.Address{},
		stockAccounts: make(map[common.Address]*StockAccount),
	}
}

func (s *TinaStock) GetOwner() common.Address {

	return s.ownerAccount
}

func (s *TinaStock) SetOwner(operation common.Address, address common.Address) error {

	if s.ownerAccount != operation {
		return errors.New("Operation Account Is not Owner")
	}
	s.ownerAccount = address
	return nil
}

func (s *TinaStock) GetGasPool() uint64 {

	return s.gasPool
}

func (s *TinaStock) AddGasPool(gas uint64) uint64 {

	s.gasPool = s.gasPool + gas
	return s.gasPool
}

func (s *TinaStock) GetStock(address common.Address) *StockAccount {

	if account, ok := s.stockAccounts[address]; ok {
		return account
	}
	return nil
}

func (s *TinaStock) GetStocks() ([]*StockAccount, error) {

	var accounts []*StockAccount
	for _, v := range s.stockAccounts {
		accounts = append(accounts, v)
	}
	return accounts, nil
}

func (s *TinaStock) SetStock(operation common.Address, address common.Address, number uint64) error {

	//检测操作账号是否是Owner
	if operation != s.ownerAccount {
		return errors.New("Operation Account Is not Owner")
	}

	stock := s.GetStock(address)
	if stock != nil {
		return errors.New("Stock Account is not Nil")
	}

	stock = new(StockAccount)
	stock.Account = address
	stock.Number = number
	stock.State = protocol.Run

	s.stockAccounts[address] = stock
	return nil
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

}
