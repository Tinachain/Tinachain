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
	"blockchain/BokerChain/common"
	"blockchain/BokerChain/common/bokerchain"
	"blockchain/BokerChain/common/config"
	"bufio"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"strings"

	log "common/log4go"

	ethcommon "github.com/Bokerchain/Boker/chain/common"
)

//版本号
var (
	Version   = "1.0.6"
	BuildTime = C.GoString(C.build_time())
)

func GetContractsToDeploy(solPath string) (contracts map[string]*bokerchain.ContractInfo, err error) {
	contracts = make(map[string]*bokerchain.ContractInfo)
	solListFile := solPath + "/" + "sol_list.txt"
	file, err := os.Open(solListFile)
	if err != nil {
		return nil, fmt.Errorf("open solListFile fail err=%s", solListFile, err.Error())
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		if line != "" {
			solFileName := path.Base(line)
			contractName := strings.TrimSuffix(solFileName, ".sol")
			contract := bokerchain.BokerchainContractMap[contractName]
			if contract == nil {
				return nil, fmt.Errorf("file %s not recognized! contractName=%s", line, contractName)
			}

			contract.SolFile = solPath + "/" + solFileName
			if !common.PathExist(contract.SolFile) {
				return nil, fmt.Errorf("file %s not found!", contract.SolFile)
			}
			contracts[contractName] = contract
		}

		if err != nil {
			if err != io.EOF {
				return nil, err
			}
			break
		}
	}
	return contracts, err
}

func DoDeployContract(client *bokerchain.Client, contractInfo *bokerchain.ContractInfo, addressManager *ethcommon.Address) (err error) {
	addressManagerString := ""
	if addressManager != nil {
		addressManagerString = addressManager.String()
	}
	log.Info("%s deploying addressManager=%s", contractInfo.Name, addressManagerString)
	var address ethcommon.Address
	if addressManager != nil {
		address, err = client.ContractDeploy(contractInfo, *addressManager)
	} else {
		address, err = client.ContractDeploy(contractInfo)
	}
	if err != nil {
		return fmt.Errorf("bokerchain.ContractDeploy err=%s", err.Error())
	}
	contractInfo.Address = address
	log.Info("%s deployed address=%s", contractInfo.Name, address.String())

	if addressManager != nil {
		log.Info("%s manager address setting...", contractInfo.Name)
		err = client.SetContract(contractInfo.Name, contractInfo.Address)
		if err != nil {
			return fmt.Errorf("client.SetContract err=%s contract=%s", err.Error(), contractInfo.Name)
		}
		log.Info("%s manager address set ok", contractInfo.Name)
	}

	return nil
}

func GenerateClientJs(client *bokerchain.Client) error {
	fJs, err := os.Create("bokerchain.js")
	if err != nil {
		return err
	}
	defer fJs.Close()

	format := `
bokerchain.%s = web3.eth.contract(%s).at('%s', function(error, contract){
    if(!error){
        console.log('%s at : ' + contract.address);
        }
    else
        console.error('%s at error : ' + error);
})

`
	fJs.WriteString("var bokerchain = new Object();\n")
	for _, contractInfo := range bokerchain.BokerchainContractMap {
		contractInfo.SolFile = config.GetInstance().BokerchainSolFolder + "/" + contractInfo.Name + ".sol"

		if contractInfo.Abi == "" {
			err := client.ContractCompile(contractInfo)
			if err != nil {
				return fmt.Errorf("ContractCompile err=%s name=%s", err.Error(), contractInfo.Name)
			}
		}
		contractAddress := contractInfo.Address
		if contractAddress == bokerchain.ZeroAddress {
			contractFind, err := client.FindContract(contractInfo.Name)
			if err != nil {
				fmt.Errorf("FindContract err=%s name=%s", err.Error(), contractInfo.Name)
				continue
			}
			contractInfo.Address = contractFind.Address
			contractAddress = contractInfo.Address
		}

		content := fmt.Sprintf(format, contractInfo.Name, contractInfo.Abi, contractAddress.String(), contractInfo.Name, contractInfo.Name)
		fJs.WriteString(content)
	}
	return nil
}

func GetInterfaceBaseContract() string {
	contractAddress := ""
	interfaceBaseFile := bokerchain.ContractInterfaceBase + ".contract"
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
	interfaceBaseFile := bokerchain.ContractInterfaceBase + ".contract"
	file, err := os.Create(interfaceBaseFile)
	if err != nil {
		return err
	}
	defer file.Close()
	file.WriteString(address)
	return nil
}

func DeployContract(rpc, keyJson, password, solPath string) (contracts map[string]*bokerchain.ContractInfo, err error) {
	contracts, err = GetContractsToDeploy(solPath)
	if err != nil {
		return nil, fmt.Errorf("getBokerContracts err=%s", err.Error())
	}

	client, err := bokerchain.NewClient(rpc)
	if err != nil {
		return nil, fmt.Errorf("NewClient %v", err.Error())
	}

	err = client.Unlock(keyJson, password)
	if err != nil {
		return nil, fmt.Errorf("Unlock %v", err.Error())
	}

	//deploy BokerManager.sol first
	contractManager := contracts[bokerchain.ContractManager]
	if contractManager != nil {
		err = DoDeployContract(client, contractManager, nil)
		if err != nil {
			return nil, fmt.Errorf("DoDeployContract err=%s solFile=%s", err.Error(), contractManager.SolFile)
		}
	} else {
		if config.GetInstance().BokerchainManagerAddress == "" {
			return nil, fmt.Errorf("BokerManager address is nil!")
		}
		contractManager = bokerchain.BokerchainContractMap[bokerchain.ContractManager]
		contractManager.Address = bokerchain.HexToAddress(config.GetInstance().BokerchainManagerAddress)
	}
	err = client.AtManager(contractManager.Address.String())
	if err != nil {
		return nil, fmt.Errorf("client.AtManager err=%s address=%s", err.Error(), contractManager.Address.String())
	}

	var bokerInterfaceBase *bokerchain.ContractInfo = nil
	for _, contract := range contracts {
		if bokerchain.ContractManager == contract.Name {
			continue
		}
		err = DoDeployContract(client, contract, &contractManager.Address)
		if err != nil {
			return nil, fmt.Errorf("DoDeployContract err=%s solFile=%s", err.Error(), contract.SolFile)
		}

		if bokerchain.ContractInterfaceBase == contract.Name {
			bokerInterfaceBase = contract
		}
	}
	client.AtInterface()
	client.AtInterfaceBase()

	//set base contract
	if bokerInterfaceBase != nil {
		contractOld := GetInterfaceBaseContract()
		if "" != contractOld {
			log.Info("CancelBaseContracts contractOld=%s", contractOld)
			err = client.CancelBaseContracts(bokerchain.HexToAddress(contractOld))
			if err != nil {
				return nil, fmt.Errorf("CancelBaseContracts err=%s address=%s", err.Error(), contractOld)
			}
		}

		log.Info("SetBaseContracts bokerInterfaceBase.Address=%s", bokerInterfaceBase.Address.String())
		err = client.SetBaseContracts(bokerInterfaceBase.Address, bokerInterfaceBase.Abi)
		if err != nil {
			return nil, fmt.Errorf("SetBaseContracts err=%s address=%s", err.Error(), bokerInterfaceBase.Address.String())
		}
		SaveInterfaceBaseContract(bokerInterfaceBase.Address.String())
	}

	info, _ := client.ContractInfo()
	log.Info("ContractInfo %s", info)

	err = GenerateClientJs(client)
	if err != nil {
		return nil, fmt.Errorf("GenerateClientJs err=%s", err.Error())
	}

	return contracts, nil
}

func DoTest() (tested bool) {

	testFile := config.GetInstance().BokerchainSolFolder + "/" + "test.deploy"
	file, err := os.Open(testFile)
	if err != nil {
		return false
	}
	defer file.Close()

	var contracts []*bokerchain.ContractInfo

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		if line != "" {
			solFileName := path.Base(line)
			contractName := strings.TrimSuffix(solFileName, ".sol")
			contractInfo := bokerchain.NewContractInfo(contractName)
			contractInfo.SolFile = config.GetInstance().BokerchainSolFolder + "/" + solFileName
			contracts = append(contracts, contractInfo)
		}

		if err != nil {
			if err != io.EOF {
				return true
			}
			break
		}
	}
	if len(contracts) <= 0 {
		return true
	}

	fmt.Println("test.deploy")

	client, err := bokerchain.NewClient(config.GetInstance().BokerchainRpc)
	if err != nil {
		fmt.Println(err)
		return true
	}

	err = client.Unlock(config.GetInstance().BokerchainAdminKeystore, config.GetInstance().BokerchainAdminPassword)
	if err != nil {
		fmt.Println(err)
		return true
	}

	fmt.Println("connected to bokerchain")

	for _, contractInfo := range contracts {
		fmt.Printf("deploying %s\n", contractInfo.Name)

		if !common.PathExist(contractInfo.SolFile) {
			fmt.Printf("%s not found\n", contractInfo.SolFile)
			return true
		}
		contractInfo.Address, err = client.ContractDeploy(contractInfo)
		if err != nil {
			fmt.Println(err)
			return true
		}
		log.Info("%s deployed address=%s", contractInfo.Name, contractInfo.Address.String())
	}

	fJs, err := os.Create("test.js")
	if err != nil {
		fmt.Println(err)
		return true
	}
	defer fJs.Close()

	format := `
test.%s = web3.eth.contract(%s).at('%s', function(error, contract){
    if(!error){
        console.log('%s at : ' + contract.address);
        }
    else
        console.error('%s at error : ' + error);
})

`
	fJs.WriteString("var test = new Object();\n")
	for _, contractInfo := range contracts {
		contractAddress := contractInfo.Address.String()
		content := fmt.Sprintf(format, contractInfo.Name, contractInfo.Abi, contractAddress, contractInfo.Name, contractInfo.Name)
		fJs.WriteString(content)
	}

	return true
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

	if DoTest() {
		return
	}

	log.Info("bokerchain deploy started...")
	contracts, err := DeployContract(
		config.GetInstance().BokerchainRpc,
		config.GetInstance().BokerchainAdminKeystore,
		config.GetInstance().BokerchainAdminPassword,
		config.GetInstance().BokerchainSolFolder)
	if err != nil {
		log.Error("DeployContract err=%s", err.Error())
	}
	log.Info("%d contracts deployed", len(contracts))
}
