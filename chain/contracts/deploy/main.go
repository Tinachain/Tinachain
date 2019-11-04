package main

/*
const char* build_time(void)
{
    static const char* psz_build_time = "["__DATE__ "  " __TIME__ "]";
    return psz_build_time;

}
*/
import "C"
import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"math/big"
	"os"
	"time"

	ethcommon "github.com/Tinachain/Tina/chain/common"
	"github.com/Tinachain/Tina/chain/contracts/deploy/common"
	"github.com/Tinachain/Tina/chain/contracts/deploy/common/config"
	"github.com/Tinachain/Tina/chain/contracts/deploy/common/tinachain"
)

//版本号
var (
	Version   = "1.0.6"
	BuildTime = C.GoString(C.build_time())
)

func Log(format string, args ...interface{}) {
	fmt.Printf("[%s]", time.Now().Format("2006-01-02 15:04:05"))
	fmt.Printf(format, args...)
	fmt.Printf("\n")
}

func GetInterfaceBaseContract() string {
	contractAddress := ""
	interfaceBaseFile := tinachain.ContractInterfaceBase + ".contract"
	file, err := os.Open(interfaceBaseFile)
	if err != nil {
		return ""
	}
	defer file.Close()

	addressBytes, _ := ioutil.ReadAll(file)
	contractAddress = string(addressBytes)
	return contractAddress
}

func SaveInterfaceBaseContract(address string) error {
	interfaceBaseFile := tinachain.ContractInterfaceBase + ".contract"
	file, err := os.Create(interfaceBaseFile)
	if err != nil {
		return err
	}
	defer file.Close()
	file.WriteString(address)
	return nil
}

type DeployConfigEntry struct {
	File      string
	Contracts []string
}

type DeployConfig struct {
	Deploy []DeployConfigEntry
}

type CWARETokenConfigEntry struct {
	Owner    string
	Supply   uint64
	Name     string
	Decimals uint8
	Symbol   string
}

type ChainwareConfig struct {
	CWAREToken CWARETokenConfigEntry
}

var interfaceBaseAddress ethcommon.Address

func LoadDeployConfig() *DeployConfig {

	cfg := &DeployConfig{}

	cfgFile, err := os.Open(config.GetInstance().BokerchainSolFolder + "/" + "deploy.json")
	if err != nil {
		return cfg
	}
	defer cfgFile.Close()

	cfgBytes, _ := ioutil.ReadAll(cfgFile)

	json.Unmarshal(cfgBytes, cfg)
	return cfg
}

func LoadChainwareConfig() *ChainwareConfig {

	cfg := &ChainwareConfig{}

	cfgFile, err := os.Open(config.GetInstance().BokerchainSolFolder + "/" + "chainware.json")
	if err != nil {
		return cfg
	}
	defer cfgFile.Close()

	cfgBytes, _ := ioutil.ReadAll(cfgFile)

	json.Unmarshal(cfgBytes, cfg)
	return cfg
}

func Deploy() {

	Log("Deploy Contracts to Tinachain")

	client, err := tinachain.NewClient(config.GetInstance().BokerchainRpc)
	if err != nil {
		Log(err.Error())
		return
	}

	err = client.Unlock(config.GetInstance().BokerchainAdminKeystore, config.GetInstance().BokerchainAdminPassword)
	if err != nil {
		Log(err.Error())
		return
	}

	Log("Deploy Program connected to Tinachain account %s\n", config.GetInstance().BokerchainAdminKeystore)

	fJs, err := os.Create("contract.js")
	if err != nil {
		Log(err.Error())
		return
	}
	defer fJs.Close()
	format := `
Ware.%s = web3.eth.contract(%s).at('%s', function(error, contract){
    if(!error){
        console.log('%s at : ' + contract.address);
        }
    else
        console.error('%s at error : ' + error);
})

				`
	fJs.WriteString("var Ware = new Object();\n")

	deployCfg := LoadDeployConfig()
	var managerAddress ethcommon.Address
	if config.GetInstance().BokerchainManagerAddress != "" {
		managerAddress = tinachain.HexToAddress(config.GetInstance().BokerchainManagerAddress)
	}
	chainwareCfg := LoadChainwareConfig()

	var CWARETokenAddress ethcommon.Address

	for _, entry := range deployCfg.Deploy {

		filePath := config.GetInstance().BokerchainSolFolder + "/" + entry.File
		if !common.PathExist(filePath) {

			Log("Deploy Program Not Found Contracts File %s", filePath)
			continue
		}

		compiledContracts, err := client.ContractCompile(filePath)
		if err != nil {
			Log(err.Error())
			continue
		}

		for _, contractName := range entry.Contracts {

			Log("Deploying Contract %s", contractName)
			contract := compiledContracts[contractName]
			if nil == contract {

				Log("Not Found Contract %s", contractName)
				continue
			}

			if contractName == tinachain.ContractManager {

				err = client.ContractDeploy(contract)
				if err != nil {

					Log("Deploy Contract %s Fail Err %s", contractName, err.Error())
					return
				}
				managerAddress = contract.Address
				fJs.WriteString(fmt.Sprintf(format, contract.Name, contract.Abi, contract.Address.String(), contract.Name, contract.Name))

				Log("Deployed Contract %s OK Address %s", contractName, managerAddress.String())
				err = client.AtManager(managerAddress.String())
				if err != nil {

					Log("Deploy Program Set Manager Contract Address %s Fail Err %s", managerAddress.String(), err.Error())
					return
				}
				Log("Set Manager Contract Address %s OK\n", managerAddress.String())

			} else if contractName == tinachain.ContractInterfaceBase {

				if managerAddress.String() == tinachain.ZeroAddressString {

					Log("Manager Address is Empty")
					return
				}

				err = client.ContractDeploy(contract, managerAddress)
				if err != nil {

					Log("Contract %s Deploy Fail Err %s", contractName, err.Error())
					return
				}
				Log("Contract %s Deploy Ok", contractName)

				fJs.WriteString(fmt.Sprintf(format, contract.Name, contract.Abi, contract.Address.String(), contract.Name, contract.Name))
				Log("Contract %s Manager Address Setring...", contractName)
				err = client.SetContract(contract.Name, contract.Address)
				if err != nil {

					Log("Contract %s Manager Address Set Fail", contractName, err.Error())
					return
				}
				Log("Contract %s SetContract Ok", contract.Name)

				contractOld := GetInterfaceBaseContract()
				if "" != contractOld {

					/*Log("CancelBaseContracts Old Contract %s", contractOld)
					err = client.CancelBaseContracts(tinachain.HexToAddress(contractOld))
					if err != nil {

						Log("CancelBaseContracts Err %s Address %s", err.Error(), contractOld)
						return
					}*/
				}

				Log("SetBaseContracts Address %s", contract.Address.String())
				err = client.SetSystemBaseContracts(contract.Address)
				if err != nil {

					Log("SetBaseContracts Address %s Err %s", contract.Address.String(), err.Error())
					return
				}

				Log("SetBaseContracts Address %s Ok", contract.Address.String())
				SaveInterfaceBaseContract(contract.Address.String())

				interfaceBaseAddress = contract.Address

			} else if "CWAREToken" == contractName {

				if managerAddress.String() == tinachain.ZeroAddressString {

					Log("Manager Address is Empty")
					return
				}

				var ownerAddress ethcommon.Address
				ownerAddress = ethcommon.HexToAddress(chainwareCfg.CWAREToken.Owner)
				Log("Deploy Contract %s %s", chainwareCfg.CWAREToken.Owner, ownerAddress.String())

				// Chainware  ERC20;
				err = client.ContractDeploy(
					contract,
					ownerAddress,
					new(big.Int).SetUint64(chainwareCfg.CWAREToken.Supply),
					chainwareCfg.CWAREToken.Name,
					chainwareCfg.CWAREToken.Decimals,
					chainwareCfg.CWAREToken.Symbol,
					managerAddress)
				if err != nil {

					Log("Deploy Contract %s Fail Err %s", contractName, err.Error())
					return
				}

				CWARETokenAddress = contract.Address
				Log("Deploy Contract %s Ok", contract.Name)
				fJs.WriteString(fmt.Sprintf(format, contract.Name, contract.Abi, contract.Address.String(), contract.Name, contract.Name))

				Log("Set Contract %s Address %s", contract.Name, contract.Address.String())
				err = client.SetContract(contract.Name, contract.Address)
				if err != nil {

					Log("Set Contract %s Address %s Fail Err %s", contract.Name, contract.Address, err.Error())
					return
				}
				Log("Set Contract %s Address %s Ok", contract.Name, contract.Address.String())

			} else if "Chainware" == contractName {

				if managerAddress.String() == tinachain.ZeroAddressString {

					Log("Manager Address is Empty")
					return
				}

				//Chainware Chainware
				err = client.ContractDeploy(contract, managerAddress, CWARETokenAddress)
				if err != nil {

					Log("Deploy Contract %s Fail Err %s", contractName, err.Error())
					return
				}
				Log("Deploy Contract %s Ok", contract.Name)

				fJs.WriteString(fmt.Sprintf(format, contract.Name, contract.Abi, contract.Address.String(), contract.Name, contract.Name))
				Log("SetContract %s Manager Address Setting...", contract.Name)
				err = client.SetContract(contract.Name, contract.Address)
				if err != nil {

					Log("SetContract %s Manager Address Fail Err %s", contract.Name, err.Error())
					return
				}
				Log("SetContract %s Manager Address %s Ok \n", contract.Name, contract.Address.String())

			} else {
				if managerAddress.String() == tinachain.ZeroAddressString {

					Log("Manager Address Is Empty")
					return
				}

				err = client.ContractDeploy(contract, managerAddress)
				if err != nil {

					Log("Deploy Contract %s Fail Err %s", contractName, err.Error())
					return
				}
				Log("Deploy Contract %s Ok", contract.Name)

				fJs.WriteString(fmt.Sprintf(format, contract.Name, contract.Abi, contract.Address.String(), contract.Name, contract.Name))
				Log("SetContract %s Manager Address Setting...", contract.Name)
				err = client.SetContract(contract.Name, contract.Address)
				if err != nil {

					Log("SetContract %s Manager Address Fail Err %s", contract.Name, err.Error())
					return
				}
				Log("SetContract %s Manager Address %s Ok \n", contract.Name, contract.Address.String())
			}

		}
	}

	return
}

func main() {
	checkVer := flag.Bool("V", false, "is ok")
	flag.Parse()
	if *checkVer {
		verString := "Chainware Deploy Version: " + Version + "\r\n"
		verString += "Compile At:" + BuildTime + "\r\n"
		fmt.Println(verString)
		return
	}

	//init config
	config.Initialize()

	//init log
	common.InitLog()
	//defer log.Close()

	Deploy()
}
