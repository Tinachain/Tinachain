package boker_contract

import (
	"context"
	"math/big"
	"time"

	"github.com/Tinachain/Tina/chain/accounts/abi/bind"
	"github.com/Tinachain/Tina/chain/boker/protocol"
	"github.com/Tinachain/Tina/chain/common"
	"github.com/Tinachain/Tina/chain/eth"
	"github.com/Tinachain/Tina/chain/log"
)

//新增合约类型
type InterfaceServiceState uint8

const (
	None InterfaceServiceState = iota
	Runing
)

//Tina链合约接口
type BokerInterfaceService struct {
	bokerInterface *BokerInterfaceBase //Tina链合约接口
	addr           common.Address      //合约地址
	ethereum       *eth.Ethereum       //以太坊对象
	quit           chan chan error     //退出chan
	start          bool                //是否已经启动
	startTimer     int64               //启动时间
}

var InterfaceService *BokerInterfaceService = nil

func NewBokerInterfaceService(ethereum *eth.Ethereum, address common.Address) (*BokerInterfaceService, error) {

	var s *BokerInterfaceService = new(BokerInterfaceService)

	bokerInterface, err := NewBokerInterfaceBase(address, eth.NewContractBackend(ethereum.ApiBackend))
	if err != nil {
		return nil, err
	}
	s.bokerInterface = bokerInterface
	s.addr = address
	s.ethereum = ethereum
	s.quit = make(chan chan error)
	s.start = false
	s.startTimer = time.Now().Unix()

	InterfaceService = s

	return s, nil
}

func (s *BokerInterfaceService) DeleteBokerInterfaceService() error {

	errc := make(chan error)
	s.quit <- errc
	return <-errc
}

func BokerInterfaceServiceState() InterfaceServiceState {
	if nil == InterfaceService {
		return None
	} else {
		return Runing
	}
}

func (s *BokerInterfaceService) createTransactOpts() *bind.TransactOpts {

	if coinbase, err := s.ethereum.Coinbase(); err == nil {

		opts := bind.NewPasswordTransactor(s.ethereum, coinbase)
		return opts
	}
	return nil
}

func (s *BokerInterfaceService) Start() {

	s.start = true

	go s.pollVote()
	go s.pollBlacks()
	go s.pollGas()
}

func (s *BokerInterfaceService) Stop() {

	s.start = false
}

//发起投票转换交易
func (s *BokerInterfaceService) pollVote() {

	var lastTxTime int64 = 0
	for {

		time.Sleep(time.Duration(500) * protocol.MillisecondTimer)

		if lastTxTime != 0 && lastTxTime == time.Now().Unix() {
			continue
		}

		if s.ethereum != nil {

			blocks := s.ethereum.BlockChain().GetBlockByNumber(0)
			if blocks == nil {
				continue
			}

			now := time.Now().Unix()
			firstTimer := blocks.Time().Int64()
			offset := now - firstTimer

			if offset%protocol.EpochInterval == 0 {

				lastTxTime = now

				opts := s.createTransactOpts()
				tx, err := s.bokerInterface.BokerInterfaceBaseTransactor.RotateVote(opts, new(big.Int).SetInt64(now))
				if err != nil {

					if err == bind.ErrNoCode {
						log.Error("(s *BokerInterfaceService) pollVote RotateVote method not found", "Contract", s.addr)
					} else {
						log.Error("(s *BokerInterfaceService) pollVote Failed to RotateVote", "err", err)
					}
					continue
				}
				log.Info("(s *BokerInterfaceService) pollVote", "tx", tx.Hash(), "tx Timer", tx.Time())
			}
		}
	}
}

func (s *BokerInterfaceService) GetVotes() {

	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*1)
	callOpts := &bind.CallOpts{Context: ctx}
	defer cancel()

	candidateArray, err := s.bokerInterface.BokerInterfaceBaseCaller.GetCandidates(callOpts)
	if err != nil {

		if err == bind.ErrNoCode {
			log.Error("GetCandidates method not found", "Contract", s.addr)
		} else {
			log.Error("Failed to GetCandidates", "err", err)
		}
		return
	} else {

		if len(candidateArray.Tickets) == 0 {

			log.Error("Failed Candidate Address Size Is Zero")
			return
		}

		if len(candidateArray.Tickets) != len(candidateArray.Addresses) {

			log.Error("Failed Candidate Address Size Not equal to Candidate Vote Size")
			return
		}

		if s.ethereum != nil {
			s.ethereum.BlockChain().CurrentBlock().DposCtx().SetValidatorVotes(candidateArray.Addresses, candidateArray.Tickets)
		}
	}
}

// 黑名单
func (s *BokerInterfaceService) pollBlacks() {

	var lastTxTime int64 = 0
	for {

		time.Sleep(time.Duration(500) * protocol.MillisecondTimer)

		if lastTxTime != 0 && lastTxTime == time.Now().Unix() {
			continue
		}

		if s.ethereum != nil {

			blocks := s.ethereum.BlockChain().GetBlockByNumber(0)
			if blocks == nil {
				continue
			}

			now := time.Now().Unix()
			firstTimer := blocks.Time().Int64()
			offset := now - firstTimer

			if offset%protocol.BlacksInterval == 0 {

				lastTxTime = now
				s.getBlacks()
			}
		}
	}
}

func (s *BokerInterfaceService) getBlacks() {

	log.Info("(s *BokerInterfaceService) getBlacks")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
	callOpts := &bind.CallOpts{Context: ctx}
	defer cancel()

	blacks, err := s.bokerInterface.BokerInterfaceBaseCaller.GetBlacks(callOpts)
	if err != nil {

		if err == bind.ErrNoCode {
			log.Error("(s *BokerInterfaceService) getBlacks method not found", "Contract", s.addr)
		} else {
			log.Error("(s *BokerInterfaceService) getBlacks Failed to GetCandidates", "err", err)
		}
	} else {

		s.ethereum.Boker().SetBlacks(blacks)
	}
}

func (s *BokerInterfaceService) pollGas() {

	var lastTxTime int64 = 0
	for {

		time.Sleep(time.Duration(500) * protocol.MillisecondTimer)

		if lastTxTime != 0 && lastTxTime == time.Now().Unix() {
			continue
		}

		if s.ethereum != nil {

			blocks := s.ethereum.BlockChain().GetBlockByNumber(0)
			if blocks == nil {
				continue
			}

			now := time.Now().Unix()
			firstTimer := blocks.Time().Int64()
			offset := now - firstTimer

			if offset%protocol.GasInterval == 0 {

				lastTxTime = now

				ctx, cancel := context.WithTimeout(context.Background(), time.Duration(1)*protocol.MillisecondTimer)
				defer cancel()

				var coinbase, producer common.Address
				var err error

				coinbase, err = s.ethereum.Coinbase()
				if err != nil {
					continue
				}

				producer, err = s.ethereum.BlockChain().CurrentBlock().DposCtx().GetCurrentProducer(firstTimer)
				if err != nil {
					continue
				}

				if coinbase != producer {
					continue
				}

				tx, _ := s.ethereum.Boker().SubmitBokerTransaction(ctx, protocol.Stock, protocol.StockAssignGas, coinbase, common.Address{}, []byte(""), []byte(""), new(big.Int).SetUint64(0), 0)
				log.Info("(s *BokerInterfaceService) pollGas", "tx", tx.Hash(), "tx Timer", tx.Time())
			}
		}
	}
}

func (s *BokerInterfaceService) getCurrentProducer() error {

	if s.ethereum != nil {

		//得到第一个区块
		blocks := s.ethereum.BlockChain().GetBlockByNumber(0)
		if blocks == nil {
			return protocol.ErrInvalidSystem
		}
		//得到第一个区块的时间
		firstTimer := blocks.Time().Int64()

		//得到当前的出块节点
		producer, err := s.ethereum.BlockChain().CurrentBlock().DposCtx().GetCurrentProducer(firstTimer)
		if err != nil {
			return protocol.ErrInvalidProducer
		}

		//得到当前挖矿节点
		coinbase, err := s.ethereum.Coinbase()
		if err != nil {
			return protocol.ErrInvalidCoinbase
		}

		//将当前出块节点和当前节点进行比较，如果是当前出块节点，则允许继续进行处理
		if producer == coinbase {
			return nil
		}
	}
	return protocol.ErrInvalidSystem
}

func (s *BokerInterfaceService) IsStart() bool { return s.start }
