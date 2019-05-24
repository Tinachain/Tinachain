package main

import (
	"blockchain/BokerChain/common/config"
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/Bokerchain/Boker/chain/common/compiler"
	//	ethcommon "github.com/boker/chain/common"
)

func GenerateManagerJs(contractName, solAbi, solBin string) error {
	format := `var %s = web3.eth.contract(%s).new(
		'0x1e0E01e42BE2E15F86A045b765E4ae8E0c644F15',
	   {
	     from: web3.eth.accounts[0],
	     data: '%s',
	     gas: '3000000'
	   }, function (e, contract){
	    console.log(e, contract);
	    if (typeof contract.address !== 'undefined') {
	         console.log('Contract mined! address: ' + contract.address + ' transactionHash: ' + contract.transactionHash);
	    }
	 })
	`
	//	content := fmt.Sprintf(format, contractName, solAbi, contractName, contractName, solBin)
	content := fmt.Sprintf(format, contractName, solAbi, solBin)
	fJs, err := os.Create("/projects/ethereum/geth/Deploy" + contractName + ".js")
	//	fJs, err := os.Create("Deploy" + contractName + ".js")
	if err != nil {
		return err
	}
	defer fJs.Close()
	_, err = fJs.WriteString(content)
	if err != nil {
		return err
	}
	return nil
}

func GenerateContractJs(solFile, contractName, param1 string) error {
	params := ""
	if param1 != "" {
		params = fmt.Sprintf("'%s',", param1)
	}
	format := `var %s = web3.eth.contract(%s).new(
		%s
	   {
	     from: web3.eth.accounts[0],
	     data: '%s',
	     gas: '7000000'
	   }, function (e, contract){
	    console.log(e, contract);
	    if (typeof contract.address !== 'undefined') {
	         console.log('Contract mined! address: ' + contract.address + ' transactionHash: ' + contract.transactionHash);
	    }
	 })
	`

	contracts, err := compiler.CompileSolidity("solc", solFile)
	if err != nil {
		return fmt.Errorf("compiler.CompileSolidity err=%v", err)
	}

	contractFullName := solFile + ":" + contractName
	contract, exist := contracts[contractFullName]
	if !exist {
		return fmt.Errorf("contractName not found!")
	}
	abiBytes, _ := json.Marshal(contract.Info.AbiDefinition)
	abiString := string(abiBytes)
	content := fmt.Sprintf(format, contractName, abiString, params, contract.Code)
	fJs, err := os.Create("Deploy" + contractName + ".js")
	if err != nil {
		return err
	}
	defer fJs.Close()
	_, err = fJs.WriteString(content)
	if err != nil {
		return err
	}
	return nil
}

func Test(t *testing.T) {
	fmt.Println("Test")

	config.Initialize()

	fmt.Println(GetInterfaceBaseContract())
	fmt.Println(SaveInterfaceBaseContract("haha"))
	fmt.Println(GetInterfaceBaseContract())
}
