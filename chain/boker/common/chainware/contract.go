package chainware

import (
	ethcommon "github.com/Tinachain/Tina/chain/common"
)

type ContractInfo struct {
	Idx        int64 // index of conftract
	Name       string
	Address    ethcommon.Address // address of contract
	Version    string            // version of contract
	CreateTime int64             // create time of contract

	SolFile string
	Abi     string
	Bin     string
}

func NewContractInfo(solFile, name string) *ContractInfo {
	return &ContractInfo{
		Name:    name,
		SolFile: solFile,
	}
}

const (
	ContractManager        = "BokerManager"
	ContractDapp           = "BokerDapp"
	ContractDappData       = "BokerDappData"
	ContractFile           = "BokerFile"
	ContractFileData       = "BokerFileData"
	ContractFinance        = "BokerFinance"
	ContractLog            = "BokerLog"
	ContractLogData        = "BokerLogData"
	ContractTokenPower     = "BokerTokenPower"
	ContractTokenPowerData = "BokerTokenPowerData"
	ContractUser           = "BokerUser"
	ContractUserData       = "BokerUserData"
	ContractNode           = "BokerNode"
	ContractNodeData       = "BokerNodeData"
	ContractDataTransfer   = "BokerDataTransfer"
	ContractInterface      = "BokerInterface"
	ContractInterfaceBase  = "BokerInterfaceBase"
)
