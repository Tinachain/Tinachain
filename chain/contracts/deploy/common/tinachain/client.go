package tinachain

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"strings"

	"github.com/Tinachain/Tina/chain/accounts/abi"
	"github.com/Tinachain/Tina/chain/accounts/abi/bind"
	ethcommon "github.com/Tinachain/Tina/chain/common"
	"github.com/Tinachain/Tina/chain/common/compiler"
	"github.com/Tinachain/Tina/chain/core/types"
	"github.com/Tinachain/Tina/chain/ethclient"

	"github.com/Tinachain/Tina/chain/contracts/deploy/solidity/src/go/boker_contract"
)

type Client struct {
	Rpc       string
	KeyJson   string
	Password  string
	EthClient *ethclient.Client

	ManagerAddress   ethcommon.Address
	InterfaceAddress ethcommon.Address

	ContractManager       *boker_contract.BokerManager
	ContractInterface     *boker_contract.BokerInterface
	ContractInterfaceBase *boker_contract.BokerInterfaceBase

	Opts *bind.TransactOpts
}

func NewClient(rpc string) (client *Client, err error) {
	client = &Client{
		Rpc: rpc,
	}

	client.EthClient, err = ethclient.Dial(rpc)
	if err != nil {
		return nil, err
	}

	return client, nil
}

func (client *Client) ContractCompile(solFile string) (contracts map[string]*ContractInfo, err error) {
	contracts = make(map[string]*ContractInfo)

	compiledContracts, err := compiler.CompileSolidity("solc", solFile)
	if err != nil {
		return nil, fmt.Errorf("compiler.CompileSolidity err=%v", err)
	}

	for compiledName, compiledContract := range compiledContracts {
		name := strings.Split(compiledName, ":")[1]
		contractInfo := NewContractInfo(solFile, name)
		abiBytes, _ := json.Marshal(compiledContract.Info.AbiDefinition)
		contractInfo.Abi = string(abiBytes)
		contractInfo.Bin = compiledContract.Code
		contracts[name] = contractInfo
	}

	return contracts, nil
}

func (client *Client) ContractDeploy(contract *ContractInfo, params ...interface{}) error {
	var parsed abi.ABI
	var err error
	var address ethcommon.Address
	var tx *types.Transaction

	parsed, err = abi.JSON(strings.NewReader(contract.Abi))
	if err != nil {
		return fmt.Errorf("abi.JSON err=%s abi=%s", err.Error(), contract.Abi)
	}
	address, tx, _, err = bind.DeployContract(client.Opts, parsed, ethcommon.FromHex(contract.Bin), client.EthClient, params...)
	if err != nil {
		return fmt.Errorf("bind.DeployContract err=%s abi=%s bin=%s", err.Error(), contract.Abi, contract.Bin)
	}
	fmt.Println("ContractDeploy tx=", tx.Hash().String(), "nounce=", tx.Nonce())
	_, err = bind.WaitDeployed(context.Background(), client.EthClient, tx)
	if err != nil {
		return fmt.Errorf("bind.WaitMined err=%v", err)
	}
	contract.Address = address

	return nil
}

func (client *Client) SetContract(contractName string, addrContract ethcommon.Address) error {
	if nil == client.ContractManager {
		return fmt.Errorf("manager not specified, call AtManager first!")
	}

	tx, err := client.ContractManager.SetContract(client.Opts, contractName, addrContract)
	if err != nil {
		return fmt.Errorf("client.ContractManager.SetContract: %v", err)
	}
	ctx := context.Background()
	_, err = bind.WaitMined(ctx, client.EthClient, tx)
	if err != nil {
		return fmt.Errorf("bind.WaitMined :%v", err)
	}
	return nil
}

func (client *Client) ContractInfo() (string, error) {
	if nil == client.ContractManager {
		return "", fmt.Errorf("manager not specified, call AtManager first!")
	}

	return client.ContractManager.ContractInfo(nil)
}

func (client *Client) FindContract(contractName string) (contract *ContractInfo, err error) {
	if nil == client.ContractManager {
		return nil, fmt.Errorf("manager not specified, call AtManager first!")
	}

	idx, name, addr, version, createTime, err := client.ContractManager.FindContract(nil, contractName)
	if err != nil {
		return nil, fmt.Errorf("client.ContractManager.FindContract: %v", err)
	}
	if name == "" {
		return nil, fmt.Errorf("contract not found")
	}

	contract = NewContractInfo("", contractName)
	contract.Idx = idx.Int64()
	contract.Address = addr
	contract.Version = version
	contract.CreateTime = createTime.Int64()
	return contract, nil
}

func (client *Client) AtInterface() (err error) {
	contractInterface, err := client.FindContract(ContractInterface)
	if err != nil {
		return fmt.Errorf("client.FindContract ContractInterface err=%s", err.Error())
	}

	if contractInterface.Address == ZeroAddress {
		return fmt.Errorf("contractInterface address is empty")
	}

	client.ContractInterface, err = boker_contract.NewBokerInterface(contractInterface.Address, client.EthClient)
	if err != nil {
		return err
	}
	client.InterfaceAddress = contractInterface.Address

	return nil
}

func (client *Client) AtInterfaceBase() (err error) {
	contractInterfaceBase, err := client.FindContract(ContractInterfaceBase)
	if err != nil {
		return fmt.Errorf("client.FindContract ContractInterfaceBase err=%s", err.Error())
	}

	if contractInterfaceBase.Address == ZeroAddress {
		return fmt.Errorf("contractInterfaceBase address is empty")
	}

	client.ContractInterfaceBase, err = boker_contract.NewBokerInterfaceBase(contractInterfaceBase.Address, client.EthClient)
	if err != nil {
		return err
	}

	return nil
}

func (client *Client) SetBaseContracts(address ethcommon.Address, abiJson string) (err error) {
	hash, err := client.EthClient.SetBaseContracts(context.Background(), address, 1, abiJson)
	if err != nil {
		return err
	}
	fmt.Printf("tx.hash=%s\n", hash.String())

	_, err = bind.WaitTransaction(context.Background(), client.EthClient, hash)
	if err != nil {
		return err
	}

	return nil
}

func (client *Client) GetBalance(address string) (*big.Int, error) {
	balance, err := client.EthClient.BalanceAt(context.Background(), ethcommon.HexToAddress(address), nil)
	if err != nil {
		return nil, err
	}
	return balance, nil
}

func (client *Client) Unlock(keyJson, password string) (err error) {
	client.Opts, err = bind.NewTransactor(strings.NewReader(keyJson), password)
	if err != nil {
		return err
	}
	client.Opts.GasLimit = big.NewInt(4700000)
	client.KeyJson = keyJson
	client.Password = password
	return nil
}

func (client *Client) AtManager(address string) (err error) {
	client.ContractManager, err = boker_contract.NewBokerManager(ethcommon.HexToAddress(address), client.EthClient)
	if err != nil {
		return fmt.Errorf("NewBokerManager err=%s", err.Error())
	}

	ok, err := client.ContractManager.IsBokerManager(nil)
	if err != nil {
		if err.Error() == "abi: unmarshalling empty output" {
			return fmt.Errorf("not BokerInterface address")
		}
		return fmt.Errorf("IsBokerManager err=%s", err.Error())
	}
	if !ok {
		return fmt.Errorf("not BokerManager address")
	}
	client.ManagerAddress = ethcommon.HexToAddress(address)

	client.AtInterface()
	client.AtInterfaceBase()

	return nil
}

func (client *Client) SetPicture(picture string) (err error) {
	hash, err := client.EthClient.SetPicture(context.Background(), picture)
	if err != nil {
		return err
	}

	_, err = bind.WaitTransaction(context.Background(), client.EthClient, hash)
	if err != nil {
		return err
	}
	return nil
}

func (client *Client) GetPicture(hash ethcommon.Hash, savePath string) (err error) {
	err = client.EthClient.GetPicture(context.Background(), hash, savePath)
	if err != nil {
		return err
	}
	return nil
}

func (client *Client) CancelBaseContracts(address ethcommon.Address) (err error) {
	hash, err := client.EthClient.CancelBaseContracts(context.Background(), address, 1)
	if err != nil {
		return err
	}

	_, err = bind.WaitTransaction(context.Background(), client.EthClient, hash)
	if err != nil {
		return err
	}
	return nil
}

func (client *Client) NetworkID() (id *big.Int, err error) {
	id, err = client.EthClient.NetworkID(context.Background())
	if err != nil {
		return id, err
	}
	return id, err
}
