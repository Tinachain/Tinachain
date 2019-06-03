package deploy

import (
	"blockchain/Tinachain/common/tinachain"
)

func DoDeployDefault(client *tinachain.Client, compiledContract *tinachain.ContractInfo) error {
	return client.ContractDeploy(compiledContract)
}

var deployer_default *Deployer = &Deployer{
	Name:     "default",
	Type:     0,
	DoDeploy: DoDeployDefault,
}
