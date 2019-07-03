package boker_contract

//go: 生成abi和bin文件 solc BokerAssignToken.sol BokerAssignTokenData.sol BokerAssignTokenDefine.sol BokerAssignTokenEventHandler.sol BokerAssignTokenImpl.sol ../BokerCommon.sol
//go: 生成go文件 abigen --abi BokerAssignToken.sol:BokerAssignToken.abi --bin BokerAssignToken.sol:BokerAssignToken.bin  --pkg assigntoken --out contract.go

import (
	"context"
	"math/big"
	"time"

	"github.com/Tinachain/Tina/chain/accounts/abi/bind"
	"github.com/Tinachain/Tina/chain/boker/protocol"
	"github.com/Tinachain/Tina/chain/common"
	_ "github.com/Tinachain/Tina/chain/common/hexutil"
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
	currentEpoch   *big.Int            //当前周期序号
	currentTimeOut *big.Int            //当前超时序号
	addr           common.Address      //合约地址
	ethereum       *eth.Ethereum       //以太坊对象
	tickQuit       chan chan error     //tick退出chan
	epochQuit      chan chan error     //epoch退出chan
	blackQuit      chan chan error     //epoch退出chan
	timeOutQuit    chan chan error     //超时检测 chan
	quit           chan chan error     //退出chan
	start          bool                //是否已经启动
	startTimer     int64               //启动时间
}

var InterfaceService *BokerInterfaceService = nil

//创建接口服务
func NewBokerInterfaceService(ethereum *eth.Ethereum, address common.Address) (*BokerInterfaceService, error) {

	var s *BokerInterfaceService = new(BokerInterfaceService)

	bokerInterface, err := NewBokerInterfaceBase(address, eth.NewContractBackend(ethereum.ApiBackend))
	if err != nil {
		return nil, err
	}
	s.bokerInterface = bokerInterface
	s.addr = address
	s.ethereum = ethereum
	s.tickQuit = make(chan chan error)
	s.epochQuit = make(chan chan error)
	s.blackQuit = make(chan chan error)
	s.timeOutQuit = make(chan chan error)
	s.quit = make(chan chan error)
	s.start = false
	s.currentEpoch = big.NewInt(0)
	s.currentTimeOut = big.NewInt(0)
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
	go s.tick()
	go s.getEpoch()
	go s.checkBlackList()
	go s.timeOut()
}

func (s *BokerInterfaceService) Stop() {

	s.start = false

	errTick := make(chan error)
	s.tickQuit <- errTick

	errEpoch := make(chan error)
	s.epochQuit <- errEpoch

	errBlack := make(chan error)
	s.blackQuit <- errBlack

	errTimeout := make(chan error)
	s.timeOutQuit <- errTimeout
}

func (s *BokerInterfaceService) tick() {

	timer := time.NewTimer(protocol.BokerInterval * 1)
	defer timer.Stop()

	for {
		select {
		case <-timer.C:
			s.business()
			timer.Reset(protocol.BokerInterval * 1)
		case errc := <-s.quit:
			errc <- nil
			return
		}
	}
}

func (s *BokerInterfaceService) tickVotes() {

	//调用时钟函数，判断是否周期发生改变
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	callOpts := &bind.CallOpts{Context: ctx}
	defer cancel()

	epochBool, err := s.bokerInterface.BokerInterfaceBaseCaller.TickVote(callOpts)
	if err != nil {

		if err == bind.ErrNoCode {
			log.Error("tickVote method not found", "Contract", s.addr)
		} else {
			log.Error("Failed to tickVote", "err", err)
		}
		return

	} else {

		//调用转换票数函数
		if epochBool {

			opts := s.createTransactOpts()
			if opts.Nonce == nil {
				nonce := s.ethereum.TxPool().State().GetNonce(opts.From)
				opts.Nonce = new(big.Int).SetUint64(nonce)
			}
			log.Info("Create RotateVote TransactOpts", "From", opts.From.String(), "Nonce", opts.Nonce)

			_, err := s.bokerInterface.BokerInterfaceBaseTransactor.RotateVote(opts)
			if err != nil {

				if err == bind.ErrNoCode {
					log.Error("rotateVote method not found", "Contract", s.addr)
				} else {
					log.Error("Failed to rotateVote", "err", err)
				}
				return
			}
		}
	}
}

//定期获取周期
func (s *BokerInterfaceService) getEpoch() {

	timer := time.NewTimer(protocol.BokerInterval * 5)
	defer timer.Stop()

	for {
		select {
		case <-timer.C:
			s.vote()
			timer.Reset(protocol.BokerInterval * 5)

		case errc := <-s.epochQuit:
			errc <- nil
			return
		}
	}
}

func (s *BokerInterfaceService) vote() {

	log.Info("(s *BokerInterfaceService) vote")

	//判断投票;
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	callOpts := &bind.CallOpts{Context: ctx}
	defer cancel()

	//获取周期
	epochIndex, err := s.bokerInterface.BokerInterfaceBaseCaller.GetVoteRound(callOpts)
	if err != nil {

		if err == bind.ErrNoCode {
			log.Error("(s *BokerInterfaceService) vote GetVoteRound method not found", "Contract", s.addr)
		} else {
			log.Error("(s *BokerInterfaceService) vote Failed to GetVoteRound", "err", err)
		}
		return
	} else {

		//判断轮数是否和当前记录的是否一致，如果不一致，则重新获取数据
		if epochIndex != s.currentEpoch && epochIndex.Int64() == s.currentEpoch.Int64()+1 {

			//调用获取候选人列表数据
			candidateArray, err := s.bokerInterface.BokerInterfaceBaseCaller.GetCandidates(callOpts)
			if err != nil {
				if err == bind.ErrNoCode {
					log.Error("GetCandidates method not found", "Contract", s.addr)
				} else {
					log.Error("Failed to GetCandidates", "err", err)
				}
				return

			} else {

				//判断候选人数组长度和候选人票数数组长度是否一致
				if len(candidateArray.Tickets) != len(candidateArray.Addresses) {

					log.Error("Failed Candidate Address Size Not equal to Candidate Vote Size")
					return
				}

				//检索以太坊服务依赖项以访问区块链
				if s.ethereum != nil {
					s.ethereum.BlockChain().CurrentBlock().DposCtx().SetValidatorVotes(candidateArray.Addresses, candidateArray.Tickets)
				}
			}

			s.currentEpoch.Add(s.currentEpoch, big.NewInt(1))
		}
	}

}

func (s *BokerInterfaceService) timeOut() {

	for {

		time.Sleep(time.Duration(5) * protocol.BokerInterval)

		blocks := s.ethereum.BlockChain().GetBlockByNumber(0)
		if blocks == nil {
			continue
		}
		firstTimer := blocks.Time().Int64()
		offset := time.Now().Unix() - firstTimer

		if offset%protocol.TimeoutInterval == 0 {

			log.Info("(s *BokerInterfaceService) timeOut", "Now", time.Now().Unix(), "firstTimer", firstTimer)
			if err := s.getCurrentProducer(); err != nil {
				log.Error("(s *BokerInterfaceService) timeOut Is`t Producer", "err", err)
				return
			}

			opts := s.createTransactOpts()
			if opts.Nonce == nil {
				nonce := s.ethereum.TxPool().State().GetNonce(opts.From)
				opts.Nonce = new(big.Int).SetUint64(nonce)
			}
			log.Info("(s *BokerInterfaceService) timeOut Create Timeout TransactOpts", "From", opts.From.String(), "Nonce", opts.Nonce)

			_, err := s.bokerInterface.BokerInterfaceBaseTransactor.TickTimeout(opts)
			if err != nil {
				if err == bind.ErrNoCode {
					log.Info("(s *BokerInterfaceService) timeOut Address Not Found", "Contract", s.addr)
				} else {
					log.Error("(s *BokerInterfaceService) timeOut Failed", "err", err)
				}
				return
			} else {
				log.Info("(s *BokerInterfaceService) timeOut End")
			}

		}
	}
}

//通证分配函数
func (s *BokerInterfaceService) business() {

	//判断出块节点是否是当前节点
	if err := s.getCurrentProducer(); err != nil {
		return
	}

	//执行投票
	//log.Info("(s *BokerInterfaceService) business Tick Vote")
	s.tickVotes()
}

func (s *BokerInterfaceService) checkBlackList() {

	timer := time.NewTimer(protocol.BlackInterval * 5)
	defer timer.Stop()

	for {
		select {
		case <-timer.C:
			s.getBlacks()
			timer.Reset(protocol.BlackInterval * 5)
		case errc := <-s.quit:
			errc <- nil
			return
		}
	}
}

func (s *BokerInterfaceService) getBlacks() {

	log.Info("(s *BokerInterfaceService) getBlacks")

	//得到黑名单;
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*1)
	callOpts := &bind.CallOpts{Context: ctx}
	defer cancel()

	blacks, err := s.bokerInterface.BokerInterfaceBaseCaller.GetBlacks(callOpts)
	if err != nil {

		if err == bind.ErrNoCode {
			log.Error("(s *BokerInterfaceService) getBlacks method not found", "Contract", s.addr)
		} else {
			log.Error("(s *BokerInterfaceService) getBlacks Failed to GetCandidates", "err", err)
		}
		return

	} else {

		//将黑名单放入到队列中;
		s.ethereum.Boker().SetBlacks(blacks)
	}
}

func (s *BokerInterfaceService) GetTokenAddr() common.Address {

	return s.addr
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

func (s *BokerInterfaceService) getCurrentTokenNoder() error {

	if s.ethereum != nil {

		//得到第一个区块
		blocks := s.ethereum.BlockChain().GetBlockByNumber(0)
		if blocks == nil {
			return protocol.ErrInvalidSystem
		}
		//得到第一个区块的时间
		firstTimer := blocks.Time().Int64()

		//得到当前的出块节点
		tokenNoder, err := s.ethereum.BlockChain().CurrentBlock().DposCtx().GetCurrentTokenNoder(firstTimer)
		if err != nil {
			return protocol.ErrInvalidTokenNoder
		}

		//得到当前挖矿节点
		coinbase, err := s.ethereum.Coinbase()
		if err != nil {
			return protocol.ErrInvalidCoinbase
		}

		//将当前出块节点和当前节点进行比较，如果是当前出块节点，则允许继续进行处理
		if tokenNoder == coinbase {
			return nil
		}
	}
	return protocol.ErrInvalidSystem
}

func (s *BokerInterfaceService) IsStart() bool { return s.start }
