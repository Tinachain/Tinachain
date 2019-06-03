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

	"github.com/Tinachain/Tina/chain/contracts/deploy/common/tinachain"

	"github.com/Tinachain/Tina/chain/contracts/deploy/common/config"

	"github.com/Tinachain/Tina/chain/contracts/deploy/common"

	log "common/log4go"

	ethcommon "github.com/Tinachain/Tina/chain/common"
)

//版本号
var (
	Version   = "1.0.6"
	BuildTime = C.GoString(C.build_time())
)

func TLog(format string, args ...interface{}) {
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

func Deploy() {
	client, err := tinachain.NewClient(config.GetInstance().BokerchainRpc)
	if err != nil {
		TLog(err.Error())
		return
	}

	err = client.Unlock(config.GetInstance().BokerchainAdminKeystore, config.GetInstance().BokerchainAdminPassword)
	if err != nil {
		TLog(err.Error())
		return
	}
	TLog("connected to tinachain")

	fJs, err := os.Create("contract.js")
	if err != nil {
		TLog(err.Error())
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
	var CWARETokenAddress ethcommon.Address
	for _, entry := range deployCfg.Deploy {
		//		TLog("***********Deploying File %s ***********", entry.File)
		filePath := config.GetInstance().BokerchainSolFolder + "/" + entry.File
		if !common.PathExist(filePath) {
			TLog("%s not found\n", filePath)
			continue
		}
		compiledContracts, err := client.ContractCompile(filePath)
		if err != nil {
			TLog(err.Error())
			continue
		}
		for _, contractName := range entry.Contracts {
			TLog("Contract %s deploying...", contractName)
			contract := compiledContracts[contractName]
			if nil == contract {
				TLog("Contract %s nod found", contractName)
				continue
			}
			if contractName == tinachain.ContractManager {
				err = client.ContractDeploy(contract)
				if err != nil {
					TLog("Contract %s deploy fail err=%s", contractName, err.Error())
					return
				}
				managerAddress = contract.Address
				fJs.WriteString(fmt.Sprintf(format, contract.Name, contract.Abi, contract.Address.String(), contract.Name, contract.Name))
				TLog("ManagerAddress = %v", managerAddress.String())
				TLog("Contract %s deploy ok!", contract.Name)
				err = client.AtManager(managerAddress.String())
				if err != nil {
					TLog("AtManager fail err=%s", err.Error())
					return
				}
				TLog("AtManager ok!")
			} else if contractName == tinachain.ContractInterfaceBase {
				if managerAddress.String() == tinachain.ZeroAddressString {
					TLog("manager address empty!")
					return
				}
				err = client.ContractDeploy(contract, managerAddress)
				if err != nil {
					TLog("Contract %s deploy fail err=%s", err.Error())
					return
				}
				TLog("Contract %s deploye ok!", contract.Name)
				fJs.WriteString(fmt.Sprintf(format, contract.Name, contract.Abi, contract.Address.String(), contract.Name, contract.Name))
				TLog("Contract %s manager address setting...", contract.Name)
				err = client.SetContract(contract.Name, contract.Address)
				if err != nil {
					TLog("Contract setContract fail err=%s", err.Error())
					return
				}
				TLog("Contract %s setContract ok!", contract.Name)

				contractOld := GetInterfaceBaseContract()
				if "" != contractOld {
					TLog("CancelBaseContracts contractOld=%s", contractOld)
					err = client.CancelBaseContracts(tinachain.HexToAddress(contractOld))
					if err != nil {
						TLog("CancelBaseContracts err=%s address=%s", err.Error(), contractOld)
						return
					}
				}

				TLog("SetBaseContracts address=%s", contract.Address.String())
				err = client.SetBaseContracts(contract.Address, contract.Abi)
				if err != nil {
					TLog("SetBaseContracts err=%s address=%s", err.Error(), contract.Address.String())
					return
				}

				TLog("SetBaseContracts ok! address=%s", contract.Address.String())
				SaveInterfaceBaseContract(contract.Address.String())
			} else if "RecipientSuccess" == contractName {
				err = client.ContractDeploy(contract)
				if err != nil {
					TLog("Contract %s deploy fail err=%s", err.Error())
					return
				}
				TLog("Contract %s deploye ok!", contract.Name)
				fJs.WriteString(fmt.Sprintf(format, contract.Name, contract.Abi, contract.Address.String(), contract.Name, contract.Name))
			} else if "CWAREToken" == contractName {
				err = client.ContractDeploy(
					contract,
					ethcommon.HexToAddress("0x1ceed12b103d9e76fea7de5410f08684eb5ef113"),
					big.NewInt(100000),
					"tina",
					uint8(9),
					"tina")
				if err != nil {
					TLog("Contract %s deploy fail err=%s", err.Error())
					return
				}
				CWARETokenAddress = contract.Address
				TLog("Contract %s deploye ok!", contract.Name)
				fJs.WriteString(fmt.Sprintf(format, contract.Name, contract.Abi, contract.Address.String(), contract.Name, contract.Name))
			} else if "Chainware" == contractName {
				err = client.ContractDeploy(contract, CWARETokenAddress)
				if err != nil {
					TLog("Contract %s deploy fail err=%s", err.Error())
					return
				}
				TLog("Contract %s deploye ok!", contract.Name)
				fJs.WriteString(fmt.Sprintf(format, contract.Name, contract.Abi, contract.Address.String(), contract.Name, contract.Name))
			} else {
				if managerAddress.String() == tinachain.ZeroAddressString {
					TLog("manager address empty!")
					return
				}
				err = client.ContractDeploy(contract, managerAddress)
				if err != nil {
					TLog("Contract %s deploy fail err=%s", err.Error())
					return
				}
				TLog("Contract %s deploye ok!", contract.Name)
				fJs.WriteString(fmt.Sprintf(format, contract.Name, contract.Abi, contract.Address.String(), contract.Name, contract.Name))
				TLog("Contract %s manager address setting...", contract.Name)
				err = client.SetContract(contract.Name, contract.Address)
				if err != nil {
					TLog("Contract setContract fail err=%s", err.Error())
					return
				}
				TLog("Contract %s setContract ok!", contract.Name)
			}

		}
	}

	return
}

func main() {
	checkVer := flag.Bool("V", false, "is ok")
	flag.Parse()
	if *checkVer {
		verString := "Bokerchain Deploy Version: " + Version + "\r\n"
		verString += "Compile At:" + BuildTime + "\r\n"
		fmt.Println(verString)
		return
	}

	//init config
	config.Initialize()

	//init log
	common.InitLog()
	defer log.Close()

	Deploy()
}
