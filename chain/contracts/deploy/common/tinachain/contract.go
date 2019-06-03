package tinachain

import (
	ethcommon "github.com/Tinachain/Tina/chain/common"
)

//type ContractInfo struct {
//	Idx        int64 // index of conftract
//	Name       string
//	Address    ethcommon.Address // address of contract
//	Version    string            // version of contract
//	CreateTime int64             // create time of contract

//	SolFile string
//	Abi     string
//	Bin     string
//}

type ContractInfo struct {
	Idx        int64 // index of conftract
	Name       string
	Address    ethcommon.Address // address of contract
	Version    string            // version of contract
	CreateTime int64             // create time of contract

	SolFile string
	Abi     string
	Bin     string
	Type    int
}

//type ContractFile struct {
//	SolFile string

//	Contracts map[string]*ContractInfo
//}

func NewContractInfo(solFile, name string) *ContractInfo {
	return &ContractInfo{
		Name:    name,
		SolFile: solFile,
	}
}

//func NewContractFile(solFile string) *ContractFile {
//	return &ContractFile{
//		SolFile:   solFile,
//		Contracts: make(map[string]*ContractInfo),
//	}
//}

//func (contractFile *ContractFile) NewContractInfo(name string) *ContractInfo {
//	contractInfo := &ContractInfo{
//		Name: name,
//	}

//	contractFile.Contracts[name] = contractInfo
//	return contractInfo
//}

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

var (
//	BokerchainContractMap = map[string]*ContractInfo{
//	//		ContractManager:        NewContractInfo(ContractManager),
//	//		ContractDapp:           NewContractInfo(ContractDapp),
//	//		ContractDappData:       NewContractInfo(ContractDappData),
//	//		ContractFile:           NewContractInfo(ContractFile),
//	//		ContractFileData:       NewContractInfo(ContractFileData),
//	//		ContractFinance:        NewContractInfo(ContractFinance),
//	//		ContractLog:            NewContractInfo(ContractLog),
//	//		ContractLogData:        NewContractInfo(ContractLogData),
//	//		ContractTokenPower:     NewContractInfo(ContractTokenPower),
//	//		ContractTokenPowerData: NewContractInfo(ContractTokenPowerData),
//	//		ContractUser:           NewContractInfo(ContractUser),
//	//		ContractUserData:       NewContractInfo(ContractUserData),
//	//		ContractNode:           NewContractInfo(ContractNode),
//	//		ContractNodeData:       NewContractInfo(ContractNodeData),
//	//		ContractDataTransfer:   NewContractInfo(ContractDataTransfer),
//	//		ContractInterface:      NewContractInfo(ContractInterface),
//	//		ContractInterfaceBase:  NewContractInfo(ContractInterfaceBase),
//	}
)
