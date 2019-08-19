// Copyright 2015 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package core

import (
	_ "bytes"
	"errors"
	"math/big"

	"github.com/Tinachain/Tina/chain/boker/api"
	"github.com/Tinachain/Tina/chain/boker/protocol"
	"github.com/Tinachain/Tina/chain/common"
	"github.com/Tinachain/Tina/chain/consensus"
	"github.com/Tinachain/Tina/chain/consensus/misc"
	"github.com/Tinachain/Tina/chain/core/state"
	"github.com/Tinachain/Tina/chain/core/types"
	"github.com/Tinachain/Tina/chain/core/vm"
	"github.com/Tinachain/Tina/chain/crypto"
	"github.com/Tinachain/Tina/chain/log"
	"github.com/Tinachain/Tina/chain/params"
)

//状态处理器，负责一个从一个节点到另一个节点
type StateProcessor struct {
	config *params.ChainConfig //链配置选项
	bc     *BlockChain         //规范块链
	engine consensus.Engine    //共识引擎
	boker  bokerapi.Api        //Tina链的接口
}

//初始化一个新的状态处理器。
func NewStateProcessor(config *params.ChainConfig, bc *BlockChain, engine consensus.Engine) *StateProcessor {
	return &StateProcessor{
		config: config,
		bc:     bc,
		engine: engine,
		boker:  bc.Boker(),
	}
}

func (p *StateProcessor) Process(block *types.Block, statedb *state.StateDB, cfg vm.Config) (types.Receipts, []*types.Log, *big.Int, error) {

	var (
		receipts     types.Receipts
		totalUsedGas = big.NewInt(0)
		header       = block.Header()
		allLogs      []*types.Log
		gp           = new(GasPool).AddGas(block.GasLimit())
		sp           = big.NewInt(protocol.MaxBlockSize)
	)

	//根据任何硬叉规范改变块和状态
	if p.config.DAOForkSupport && p.config.DAOForkBlock != nil && p.config.DAOForkBlock.Cmp(block.Number()) == 0 {
		misc.ApplyDAOHardFork(statedb)
	}

	//得到区块中所有的交易，并将这些交易使用Dpos引擎进行执行。
	for i, tx := range block.Transactions() {

		//设置当前statedb状态,以便后面evm创建交易日志
		statedb.Prepare(tx.Hash(), block.Hash(), i)

		if p.bc.Boker() != nil {
			log.Info("(p *StateProcessor) Process boker notis nil")
		}

		receipt, _, err := ApplyTransaction(p.config, block.DposCtx(), block.BokerCtx(), p.bc, nil, gp, sp, statedb, header, tx, totalUsedGas, cfg, p.bc.Boker())
		if err != nil {
			return nil, nil, nil, err
		}

		//执行完毕的交易回执放入到回执数组中
		receipts = append(receipts, receipt)
		allLogs = append(allLogs, receipt.Logs...)
	}

	//执行完块中所有的交易，应用任何共识引擎特定的附加功能（例如块奖励）
	p.engine.Finalize(p.bc, header, statedb, block.Transactions(), block.Uncles(), receipts, block.DposCtx(), block.BokerCtx(), p.bc.Boker())

	//返回执行成功的回执数组/日志/以及总的使用Gas的数量
	return receipts, allLogs, totalUsedGas, nil
}

func normalTransaction(config *params.ChainConfig,
	dposContext *types.DposContext,
	bokerContext *types.BokerContext,
	bc *BlockChain,
	author *common.Address,
	gp *GasPool,
	sp *big.Int,
	statedb *state.StateDB,
	header *types.Header,
	tx *types.Transaction,
	usedGas *big.Int,
	cfg vm.Config,
	msg types.Message,
	boker bokerapi.Api) (*types.Receipt, *big.Int, error) {

	msg, err := tx.AsMessage(types.MakeSigner(config, header.Number))
	if err != nil {
		return nil, nil, err
	}

	context := NewEVMContext(msg, header, bc, author)
	vmenv := vm.NewEVM(context, statedb, config, cfg)
	_, gas, failed, err := NormalMessage(vmenv, msg, gp, sp, dposContext, bokerContext, boker)
	if err != nil {
		return nil, nil, err
	}

	//用待处理的更改更新状态
	var root []byte
	if config.IsByzantium(header.Number) {
		statedb.Finalise(true)
	} else {
		root = statedb.IntermediateRoot(config.IsEIP158(header.Number)).Bytes()
	}
	usedGas.Add(usedGas, gas)

	//为交易创建一个新收据，存储tx使用的中间根和gas基于eip阶段，我们传递了根触发删除帐户。
	receipt := types.NewReceipt(root, failed, usedGas)
	receipt.TxHash = tx.Hash()
	receipt.GasUsed = new(big.Int).Set(gas)

	//如果交易创建了合同，则将创建地址存储在收据中
	if msg.To() == nil {
		receipt.ContractAddress = crypto.CreateAddress(vmenv.Context.Origin, tx.Nonce())
		log.Info("binaryTransaction", "address", receipt.ContractAddress.String(), "tx.Hash", tx.Hash().String())
	}

	//设置收据日志并创建一个用于过滤的布尔值
	receipt.Logs = statedb.GetLogs(tx.Hash())
	receipt.Bloom = types.CreateBloom(types.Receipts{receipt})

	return receipt, gas, err
}

func setSystemContractTransaction(config *params.ChainConfig,
	dposContext *types.DposContext,
	bokerContext *types.BokerContext,
	bc *BlockChain,
	author *common.Address,
	gp *GasPool,
	sp *big.Int,
	statedb *state.StateDB,
	header *types.Header,
	tx *types.Transaction,
	usedGas *big.Int,
	cfg vm.Config,
	msg types.Message,
	boker bokerapi.Api) (*types.Receipt, *big.Int, error) {

	log.Info("state_processor.go setSystemContractTransaction")

	msg, err := tx.AsMessage(types.MakeSigner(config, header.Number))
	if err != nil {
		log.Error("contractSetTransaction tx.AsMessage", "msg", msg, "err", err)
		return nil, nil, err
	}

	context := NewEVMContext(msg, header, bc, author)
	vmenv := vm.NewEVM(context, statedb, config, cfg)
	_, gas, failed, err := systemContractMessage(vmenv, msg, gp, sp, msg.Major(), msg.Minor(), dposContext, bokerContext, boker)
	if err != nil {
		log.Error("contractSetTransaction contractMessage", "gas", gas, "failed", failed, "err", err)
		return nil, nil, err
	}

	var root []byte
	if config.IsByzantium(header.Number) {
		statedb.Finalise(true)
	} else {
		root = statedb.IntermediateRoot(config.IsEIP158(header.Number)).Bytes()
	}
	usedGas.Add(usedGas, gas)

	receipt := types.NewReceipt(root, failed, usedGas)
	receipt.TxHash = tx.Hash()
	receipt.GasUsed = new(big.Int).Set(gas)
	receipt.Logs = statedb.GetLogs(tx.Hash())
	receipt.Bloom = types.CreateBloom(types.Receipts{receipt})
	return receipt, gas, err
}

func systemBaseTransaction(config *params.ChainConfig,
	dposContext *types.DposContext,
	bokerContext *types.BokerContext,
	bc *BlockChain,
	author *common.Address,
	gp *GasPool,
	sp *big.Int,
	statedb *state.StateDB,
	header *types.Header,
	tx *types.Transaction,
	usedGas *big.Int,
	cfg vm.Config,
	msg types.Message,
	boker bokerapi.Api) (*types.Receipt, *big.Int, error) {

	log.Info("state_processor.go systemBaseTransaction")

	msg, err := tx.AsMessage(types.MakeSigner(config, header.Number))
	if err != nil {

		log.Error("systemBaseTransaction AsMessage", "err", err)
		return nil, nil, err
	}
	log.Info("systemBaseTransaction", "Major", tx.Major(), "Minor", tx.Minor(), "Time", header.Time.Int64())

	context := NewEVMContext(msg, header, bc, author)
	vmenv := vm.NewEVM(context, statedb, config, cfg)
	_, gas, failed, err := SystemBaseMessage(vmenv, msg, gp, sp, dposContext, bokerContext, boker)
	if err != nil {
		log.Error("baseTransaction failed", "err", err)
		return nil, nil, err
	}

	var root []byte
	if config.IsByzantium(header.Number) {
		statedb.Finalise(true)
	} else {
		root = statedb.IntermediateRoot(config.IsEIP158(header.Number)).Bytes()
	}
	usedGas.Add(usedGas, gas)

	receipt := types.NewReceipt(root, failed, usedGas)
	receipt.TxHash = tx.Hash()
	receipt.GasUsed = new(big.Int).Set(gas)
	receipt.Logs = statedb.GetLogs(tx.Hash())
	receipt.Bloom = types.CreateBloom(types.Receipts{receipt})

	log.Info("state_processor.go systemBaseTransaction", "gas", gas, "err", err)
	return receipt, gas, err
}

func userBaseTransaction(config *params.ChainConfig,
	dposContext *types.DposContext,
	bokerContext *types.BokerContext,
	bc *BlockChain,
	author *common.Address,
	gp *GasPool,
	sp *big.Int,
	statedb *state.StateDB,
	header *types.Header,
	tx *types.Transaction,
	usedGas *big.Int,
	cfg vm.Config,
	msg types.Message,
	boker bokerapi.Api) (*types.Receipt, *big.Int, error) {

	log.Info("state_processor.go userBaseTransaction")

	msg, err := tx.AsMessage(types.MakeSigner(config, header.Number))
	if err != nil {

		log.Error("userBaseTransaction AsMessage", "err", err)
		return nil, nil, err
	}
	log.Info("userBaseTransaction", "Major", tx.Major(), "Minor", tx.Minor(), "Time", header.Time.Int64())

	context := NewEVMContext(msg, header, bc, author)
	vmenv := vm.NewEVM(context, statedb, config, cfg)
	_, gas, failed, err := UserBaseMessage(vmenv, msg, gp, sp, dposContext, bokerContext, boker)
	if err != nil {
		log.Error("userBaseTransaction failed", "err", err)
		return nil, nil, err
	}

	var root []byte
	if config.IsByzantium(header.Number) {
		statedb.Finalise(true)
	} else {
		root = statedb.IntermediateRoot(config.IsEIP158(header.Number)).Bytes()
	}
	usedGas.Add(usedGas, gas)

	receipt := types.NewReceipt(root, failed, usedGas)
	receipt.TxHash = tx.Hash()
	receipt.GasUsed = new(big.Int).Set(gas)
	receipt.Logs = statedb.GetLogs(tx.Hash())
	receipt.Bloom = types.CreateBloom(types.Receipts{receipt})

	log.Info("state_processor.go userBaseTransaction", "gas", gas, "err", err)
	return receipt, gas, err
}

func extraTransaction(config *params.ChainConfig,
	dposContext *types.DposContext,
	bokerContext *types.BokerContext,
	bc *BlockChain,
	author *common.Address,
	gp *GasPool,
	sp *big.Int,
	statedb *state.StateDB,
	header *types.Header,
	tx *types.Transaction,
	usedGas *big.Int,
	cfg vm.Config,
	msg types.Message,
	boker bokerapi.Api) (*types.Receipt, *big.Int, error) {

	log.Info("state_processor.go extraTransaction")

	msg, err := tx.AsMessage(types.MakeSigner(config, header.Number))
	if err != nil {

		log.Error("extraTransaction AsMessage", "err", err)
		return nil, nil, err
	}
	log.Info("extraTransaction", "Major", tx.Major(), "Minor", tx.Minor(), "Time", header.Time.Int64(), "Extra", tx.Extra())

	context := NewEVMContext(msg, header, bc, author)
	vmenv := vm.NewEVM(context, statedb, config, cfg)
	_, gas, failed, err := ExtraMessage(vmenv, msg, gp, sp, dposContext, bokerContext, boker)
	if err != nil {
		log.Error("extraTransaction failed", "err", err)
		return nil, nil, err
	}

	var root []byte
	if config.IsByzantium(header.Number) {
		statedb.Finalise(true)
	} else {
		root = statedb.IntermediateRoot(config.IsEIP158(header.Number)).Bytes()
	}
	usedGas.Add(usedGas, gas)

	receipt := types.NewReceipt(root, failed, usedGas)
	receipt.TxHash = tx.Hash()
	receipt.GasUsed = new(big.Int).Set(gas)
	receipt.Logs = statedb.GetLogs(tx.Hash())
	receipt.Bloom = types.CreateBloom(types.Receipts{receipt})

	log.Info("state_processor.go extraTransaction", "gas", gas)
	return receipt, gas, err
}

func validatorTransaction(config *params.ChainConfig,
	dposContext *types.DposContext,
	bokerContext *types.BokerContext,
	bc *BlockChain,
	author *common.Address,
	gp *GasPool,
	sp *big.Int,
	statedb *state.StateDB,
	header *types.Header,
	tx *types.Transaction,
	usedGas *big.Int,
	cfg vm.Config,
	msg types.Message,
	boker bokerapi.Api) (*types.Receipt, *big.Int, error) {

	log.Info("state_processor.go validatorTransaction")

	msg, err := tx.AsMessage(types.MakeSigner(config, header.Number))
	if err != nil {
		return nil, nil, err
	}

	context := NewEVMContext(msg, header, bc, author)
	vmenv := vm.NewEVM(context, statedb, config, cfg)

	firstBlock := bc.GetBlockByNumber(0)
	if firstBlock == nil {
		return nil, nil, errors.New("not found first block")
	}
	producer, err := dposContext.GetProducer(header.Time.Int64(), firstBlock.Time().Int64())

	log.Info("validatorTransaction", "Time", header.Time.Int64(), "firstTimer", firstBlock.Time().Int64(), "err", err)

	if protocol.ErrEpochTrieNil == err || protocol.ErrInvalidProducer == err {

		log.Info("validatorTransaction", "Number", bc.CurrentBlock().Number().Int64())
		if bc.CurrentBlock().Number().Int64() == 0 {

			_, gas, failed, err := validatorMessage(vmenv, msg, gp, sp, msg.Major(), msg.Minor(), dposContext, bokerContext, boker)
			if err != nil {
				return nil, nil, err
			}

			//设置验证者
			dposContext.Clean()
			dposContext.InsertValidator(*msg.To(), protocol.SetValidatorVotes)

			root := statedb.IntermediateRoot(false).Bytes()
			usedGas.Add(usedGas, gas)

			receipt := types.NewReceipt(root, failed, usedGas)
			receipt.TxHash = tx.Hash()
			receipt.GasUsed = new(big.Int).Set(gas)
			receipt.Logs = statedb.GetLogs(tx.Hash())
			receipt.Bloom = types.CreateBloom(types.Receipts{receipt})

			return receipt, gas, err
		}
		return nil, nil, errors.New("current block number is`t zero")
	}

	if producer != msg.From() {
		return nil, nil, errors.New("from address not assign token producer")
	}

	_, gas, failed, err := validatorMessage(vmenv, msg, gp, sp, msg.Major(), msg.Minor(), dposContext, bokerContext, boker)
	if err != nil {
		return nil, nil, err
	}

	//设置验证者
	//log.Info("validatorTransaction validatorMessage", "txType", msg.TxType())
	dposContext.Clean()
	dposContext.InsertValidator(*msg.To(), protocol.SetValidatorVotes)

	var root []byte
	if config.IsByzantium(header.Number) {
		statedb.Finalise(true)
	} else {
		root = statedb.IntermediateRoot(config.IsEIP158(header.Number)).Bytes()
	}
	usedGas.Add(usedGas, gas)
	receipt := types.NewReceipt(root, failed, usedGas)
	receipt.TxHash = tx.Hash()
	receipt.GasUsed = new(big.Int).Set(gas)
	receipt.Logs = statedb.GetLogs(tx.Hash())
	receipt.Bloom = types.CreateBloom(types.Receipts{receipt})
	return receipt, gas, err
}

func stockTransaction(config *params.ChainConfig,
	dposContext *types.DposContext,
	bokerContext *types.BokerContext,
	bc *BlockChain,
	author *common.Address,
	gp *GasPool,
	sp *big.Int,
	statedb *state.StateDB,
	header *types.Header,
	tx *types.Transaction,
	usedGas *big.Int,
	cfg vm.Config,
	msg types.Message,
	boker bokerapi.Api) (*types.Receipt, *big.Int, error) {

	log.Info("state_processor.go stockTransaction")
	msg, err := tx.AsMessage(types.MakeSigner(config, header.Number))
	if err != nil {

		log.Error("state_processor.go stockTransaction AsMessage", "err", err)
		return nil, nil, err
	}
	log.Info("state_processor.go stockTransaction", "Major", tx.Major(), "Minor", tx.Minor(), "Time", header.Time.Int64(), "Extra", tx.Extra())

	if nil == msg.To {
		return nil, nil, errors.New("Stock Transaction To is nil")
	}

	context := NewEVMContext(msg, header, bc, author)
	vmenv := vm.NewEVM(context, statedb, config, cfg)
	_, gas, failed, err := StockMessage(vmenv, msg, gp, sp, tx.Time(), dposContext, bokerContext, boker)
	if err != nil {
		log.Error("state_processor.go stockTransaction failed", "err", err)
		return nil, nil, err
	}

	var root []byte
	if config.IsByzantium(header.Number) {
		statedb.Finalise(true)
	} else {
		root = statedb.IntermediateRoot(config.IsEIP158(header.Number)).Bytes()
	}
	usedGas.Add(usedGas, gas)

	receipt := types.NewReceipt(root, failed, usedGas)
	receipt.TxHash = tx.Hash()
	receipt.GasUsed = new(big.Int).Set(gas)
	receipt.Logs = statedb.GetLogs(tx.Hash())
	receipt.Bloom = types.CreateBloom(types.Receipts{receipt})

	log.Info("state_processor.go stockTransaction", "gas", gas)
	return receipt, gas, err
}

//执行交易
func ApplyTransaction(config *params.ChainConfig,
	dposContext *types.DposContext,
	bokerContext *types.BokerContext,
	bc *BlockChain,
	author *common.Address,
	gp *GasPool,
	sp *big.Int,
	statedb *state.StateDB,
	header *types.Header,
	tx *types.Transaction,
	usedGas *big.Int,
	cfg vm.Config,
	boker bokerapi.Api) (*types.Receipt, *big.Int, error) {

	msg, err := tx.AsMessage(types.MakeSigner(config, header.Number))
	if err != nil {
		return nil, nil, err
	}
	log.Info("state_processor.go ApplyTransaction", "Number", header.Number.String(), "Major", msg.Major(), "Minor", msg.Minor(), "from", msg.From(), "extra", tx.Extra())

	if msg.Major() == protocol.Normal {

		return normalTransaction(config, dposContext, bokerContext, bc, author, gp, sp, statedb, header, tx, usedGas, cfg, msg, boker)
	} else if msg.Major() == protocol.SystemBase {

		if msg.To() == nil {
			return nil, nil, protocol.ErrToIsNil
		}

		switch msg.Minor() {

		case protocol.SetSystemContract:

			return setSystemContractTransaction(config, dposContext, bokerContext, bc, author, gp, sp, statedb, header, tx, usedGas, cfg, msg, boker)
		case protocol.VoteUser, protocol.VoteEpoch, protocol.RegisterCandidate:

			return systemBaseTransaction(config, dposContext, bokerContext, bc, author, gp, sp, statedb, header, tx, usedGas, cfg, msg, boker)
		case protocol.SetValidator:

			return validatorTransaction(config, dposContext, bokerContext, bc, author, gp, sp, statedb, header, tx, usedGas, cfg, msg, boker)
		default:

			return nil, nil, protocol.ErrInvalidType
		}
	} else if msg.Major() == protocol.UserBase {
		return userBaseTransaction(config, dposContext, bokerContext, bc, author, gp, sp, statedb, header, tx, usedGas, cfg, msg, boker)

	} else if protocol.Stock == msg.Major() {

		return stockTransaction(config, dposContext, bokerContext, bc, author, gp, sp, statedb, header, tx, usedGas, cfg, msg, boker)

	} else if protocol.Extra == msg.Major() {

		return extraTransaction(config, dposContext, bokerContext, bc, author, gp, sp, statedb, header, tx, usedGas, cfg, msg, boker)
	}
	return nil, nil, errors.New("Not Found Tx Major")
}

func (p *StateProcessor) SetBoker(boker bokerapi.Api) { p.boker = boker }
