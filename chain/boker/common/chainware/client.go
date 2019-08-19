package chainware

import (
	"context"
	"fmt"
	"math/big"
	"strings"

	"github.com/Tinachain/Tina/chain/accounts/abi/bind"
	ethcommon "github.com/Tinachain/Tina/chain/common"
	"github.com/Tinachain/Tina/chain/contracts/boker_interface"
	"github.com/Tinachain/Tina/chain/ethclient"
)

type Client struct {
	Rpc               string
	KeyJson           string
	Password          string
	EthClient         *ethclient.Client
	ManagerAddress    ethcommon.Address
	InterfaceAddress  ethcommon.Address
	From              ethcommon.Address
	Signer            bind.SignerFn
	ContractManager   *boker_contract.BokerManager
	ContractInterface *boker_contract.BokerInterface
}

func NewClient(rpc, managerAddress, interfaceAddress string) (client *Client, err error) {

	client = &Client{
		Rpc: rpc,
	}

	client.EthClient, err = ethclient.Dial(rpc)
	if err != nil {
		return nil, err
	}

	if err := client.AtManager(managerAddress); err != nil {
		return nil, err
	}

	return client, nil
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

	return nil
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

func (client *Client) GetManagerAddress() (address ethcommon.Address) {
	return client.ManagerAddress
}

func (client *Client) GetInterfaceAddress() (address ethcommon.Address) {
	return client.InterfaceAddress
}

func (client *Client) ContractAddresses() ([]ethcommon.Address, error) {
	if nil == client.ContractManager {
		return nil, fmt.Errorf("manager not specified, call AtManager first!")
	}

	return client.ContractManager.ContractAddresses(nil)
}

func (client *Client) GetBalance(address string) (*big.Int, error) {

	balance, err := client.EthClient.BalanceAt(context.Background(), ethcommon.HexToAddress(address), nil)
	if err != nil {
		return nil, err
	}
	return balance, nil
}

func (client *Client) GetPendingNonce() (nonce uint64, err error) {
	return client.EthClient.PendingNonceAt(context.Background(), client.From)
}

func (client *Client) Unlock(keyJson, password string) (err error) {
	opts, err := bind.NewTransactor(strings.NewReader(keyJson), password)
	if err != nil {
		return err
	}
	client.Signer = opts.Signer
	client.From = opts.From
	client.KeyJson = keyJson
	client.Password = password

	return nil
}

func (client *Client) NewOpts() *bind.TransactOpts {
	return &bind.TransactOpts{
		From:   client.From,
		Signer: client.Signer,
	}
}

func (client *Client) NewCallOpts() *bind.CallOpts {
	return &bind.CallOpts{
		From: client.From,
	}
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

func (client *Client) ContractInfo() (string, error) {

	if nil == client.ContractManager {
		return "", fmt.Errorf("manager not specified, call AtManager first!")
	}

	return client.ContractManager.ContractInfo(nil)
}

func (client *Client) GetContractSize() (int64, error) {
	if nil == client.ContractManager {
		return 0, fmt.Errorf("manager not specified, call AtManager first!")
	}

	sz, err := client.ContractManager.GetContractSize(nil)
	if err != nil {
		return 0, err
	}

	return sz.Int64(), nil
}

//*********************Chainware Node Interface

func (client *Client) RegisterCandidate(description, team, name string, tickets *big.Int) (ethcommon.Hash, error) {

	if nil == client.ContractInterface {
		return ethcommon.Hash{}, fmt.Errorf("interface not specified, call AtInterface first!")
	}

	opts := client.NewOpts()
	tx, err := client.ContractInterface.RegisterCandidate(opts, description, team, name, tickets)
	if err != nil {
		return ethcommon.Hash{}, err
	}
	return tx.Hash(), nil
}

/*func (client *Client) Vote(addrvoter, addrcandidate ethcommon.Address, tokens int64) (ethcommon.Hash, error) {

	if nil == client.ContractInterface {
		return ethcommon.Hash{}, fmt.Errorf("interface not specified, call AtInterface first!")
	}

	opts := client.NewOpts()
	tx, err := client.ContractInterface.Vote(opts, addrvoter, addrcandidate, new(big.Int).SetInt64(tokens))
	if err != nil {
		return ethcommon.Hash{}, err
	}
	return tx.Hash(), nil
}*/

func (client *Client) CancelAllVotes(address ethcommon.Address) (ethcommon.Hash, error) {

	if nil == client.ContractInterface {
		return ethcommon.Hash{}, fmt.Errorf("interface not specified, call AtInterface first!")
	}

	opts := client.NewOpts()
	tx, err := client.ContractInterface.CancelAllVotes(opts)
	if err != nil {
		return ethcommon.Hash{}, err
	}
	return tx.Hash(), nil
}

func (client *Client) GetCandidates() (struct {
	Addresses []ethcommon.Address
	Tickets   []*big.Int
}, error) {

	var ret struct {
		Addresses []ethcommon.Address
		Tickets   []*big.Int
	}

	//log4plus.Info("(client *Client) GetCandidates")
	if nil == client.ContractInterface {
		return ret, fmt.Errorf("interface not specified, call AtInterface first!")
	}

	//log4plus.Info("(client *Client) GetCandidates used client.ContractInterface.GetCandidates")
	callOpts := client.NewCallOpts()
	candidates, err := client.ContractInterface.GetCandidates(callOpts)
	if err != nil {
		return ret, err
	}
	return candidates, nil
}

func (client *Client) GetCandidate(address ethcommon.Address) (struct {
	Description string
	Team        string
	Name        string
	Tickets     *big.Int
}, error) {

	var ret struct {
		Description string
		Team        string
		Name        string
		Tickets     *big.Int
	}

	if nil == client.ContractInterface {
		return ret, fmt.Errorf("interface not specified, call AtInterface first!")
	}

	callOpts := client.NewCallOpts()
	candidate, err := client.ContractInterface.GetCandidate(callOpts, address)
	if err != nil {
		return ret, err
	}
	return candidate, nil
}
