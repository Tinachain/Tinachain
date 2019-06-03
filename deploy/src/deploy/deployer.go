package deploy

import (
	"blockchain/Tinachain/common/config"
	"blockchain/Tinachain/common/tinachain"
	"encoding/json"
	"io/ioutil"
	"os"
)

const (
	DeployerTypeDefault = iota
	DeployerTypeManaged
)

var (
	deployerNameMap = make(map[string]*Deployer)
	deployerTypeMap = make(map[int]*Deployer)
)

func init() {
	deployerTypeMap[DeployerTypeDefault] = deployer_default
}

type DeployContract struct {
	Name string
	Type int
}

type DeployFile struct {
	File      string
	Contracts []DeployContract
}

type DeployConfig struct {
	Deploy []DeployFile
}

type DoDeploy func(*tinachain.Client, *tinachain.ContractInfo) error
type Deployer struct {
	Name     string
	Type     int
	DoDeploy DoDeploy
}

func (deployer *Deployer) Deploy(client *tinachain.Client, compiledContract *tinachain.ContractInfo) error {
	if deployer.DoDeploy != nil {
		return deployer.DoDeploy(client, compiledContract)
	}
	return nil
}

func Deploy(client *tinachain.Client, compiledContract *tinachain.ContractInfo, deployContract *DeployContract) error {
	deployer := GetDeployer(deployContract)
	return deployer.Deploy(client, compiledContract)
}

func GetDeployer(deployContract *DeployContract) (deployer *Deployer) {
	deployer = deployerNameMap[deployContract.Name]
	if deployer != nil {
		return deployer
	}
	return deployerTypeMap[deployContract.Type]
}

func LoadDeployConfig() (cfg *DeployConfig, err error) {
	cfg = &DeployConfig{}

	cfgFile, err := os.Open(config.GetInstance().BokerchainSolFolder + "/" + "deploy.json")
	if err != nil {
		return nil, err
	}
	defer cfgFile.Close()

	cfgBytes, _ := ioutil.ReadAll(cfgFile)
	err = json.Unmarshal(cfgBytes, cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
