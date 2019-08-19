package business

import (
	"fmt"

	"github.com/Tinachain/Tina/chain/boker/common/chainware"
)

var (
	ChainClient *chainware.Client = nil
)

func InitChain(rpc, keyJson, password, managerAddress, interfaceAddress string) (err error) {

	ChainClient, err = chainware.NewClient(rpc, managerAddress, interfaceAddress)
	if err != nil {
		return fmt.Errorf("NewClient %v", err.Error())
	}

	err = ChainClient.Unlock(keyJson, password)
	if err != nil {
		return fmt.Errorf("Unlock %v", err.Error())
	}

	if managerAddress != "" {
		err = ChainClient.AtManager(managerAddress)
		if err != nil {
			return fmt.Errorf("client.AtManager err=%s address=%s", err.Error(), managerAddress)
		}
	}

	return nil
}
