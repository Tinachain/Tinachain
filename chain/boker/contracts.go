//Tina链增加的特殊账号管理类
package boker

import (
	_ "errors"
	_ "strconv"

	"github.com/Tinachain/Tina/chain/boker/protocol"
	"github.com/Tinachain/Tina/chain/common"
	"github.com/Tinachain/Tina/chain/contracts/boker_interface"
	"github.com/Tinachain/Tina/chain/core/types"
	"github.com/Tinachain/Tina/chain/eth"
	"github.com/Tinachain/Tina/chain/ethdb"
	"github.com/Tinachain/Tina/chain/log"
)

type ContractConfig struct {
	ContractType    protocol.BaseContractType `json:"contractType"`    //基础合约类型
	ContractAddress common.Address            `json:"contractAddress"` //合约地址
}

type ContractConfigs struct {
	baseContractArray []ContractConfig
}

type ContractService struct {
	contract *boker_contract.BokerInterfaceService
}

type BokerContracts struct {
	ethereum *eth.Ethereum
	services ContractService
}

func NewContract(db ethdb.Database, ethereum *eth.Ethereum, bokerProto *types.BokerBackendProto) (bokerContracts *BokerContracts, err error) {

	log.Info("contracts.go NewContract", "root", bokerProto.Root().String())

	bokerContracts = &BokerContracts{
		ethereum: ethereum,
	}

	log.Info("contracts.go NewContract New BokerContext From BokerBackendProto")
	var bokerContext *types.BokerContext = nil
	bokerContext, err = types.NewBokerContextFromProto(db, bokerProto)
	if err != nil {
		return bokerContracts, err
	}

	log.Info("contracts.go NewContract Get System Contract Address")
	var address common.Address
	if address, err = bokerContext.GetSystemContractAddress(); err != nil {

		log.Error("contracts.go NewContract System Contract is`t Exists", "err", err)
		return bokerContracts, nil
	}

	log.Info("contracts.go NewContract New System Contract Service")
	if bokerContracts.services.contract, err = boker_contract.NewBokerInterfaceService(bokerContracts.ethereum, address); err != nil {

		log.Error("contracts.go NewContract New System Contract Service Failed", "err", err)
		return bokerContracts, err
	}
	log.Info("contracts.go NewContract Start System Contract Service")
	bokerContracts.services.contract.Start()

	return bokerContracts, nil
}

func (c *BokerContracts) SetSystemContract(address common.Address, from common.Address, bokerContext *types.BokerContext) (err error) {

	log.Info("(c *BokerContracts) SetSystemContract", "address", address.String(), "from", from.String())

	if err = bokerContext.SetSystemContract(address, from); err != nil {

		log.Error("(c *BokerContracts) SetSystemContract Failed", "err", protocol.ErrNewContractService)
		return err
	}

	c.services.contract, err = boker_contract.NewBokerInterfaceService(c.ethereum, address)
	if err != nil {

		log.Error("(c *BokerContracts) SetSystemContract Failed", "err", err)
		return err
	}

	log.Info("SetContract Start Contract")
	c.services.contract.Start()

	return nil
}

func (c *BokerContracts) GetVotes() error {

	if c.services.contract != nil {

		c.services.contract.GetVotes()
	}
	return nil
}
