//Tina链增加的特殊账号管理类
package boker

import (
	"errors"

	"github.com/Tinachain/Tina/chain/boker/protocol"
	"github.com/Tinachain/Tina/chain/common"
	_ "github.com/Tinachain/Tina/chain/log"
)

//帐号常量
const SystemAddress = "0xd7fd311c8f97349670963d87f37a68794dfa80ff"

var (
	SystemAccount     = common.HexToAddress("") //系统账户
	CommunityAccount  = common.HexToAddress("") //社区账户
	FoundationAccount = common.HexToAddress("") //基金账户
	TeamAccount       = common.HexToAddress("") //团队账户
)

//Tina链的账号管理
type AcccountLevel struct {
	level protocol.TxMajor
}
type BokerAccount struct {
	ownerAccount      common.Address
	systemAccount     common.Address
	communityAccount  common.Address
	foundationAccount common.Address
	teamAccount       common.Address
}

func NewAccount() *BokerAccount {

	bokerAccount := new(BokerAccount)

	bokerAccount.systemAccount = SystemAccount
	bokerAccount.communityAccount = CommunityAccount
	bokerAccount.foundationAccount = FoundationAccount
	bokerAccount.teamAccount = TeamAccount

	return bokerAccount
}

func (a *BokerAccount) GetAccount(address common.Address) (protocol.TxMajor, error) {

	if address == a.systemAccount || address == a.communityAccount || address == a.foundationAccount || address == a.teamAccount {
		return protocol.Base, nil
	} else {
		return protocol.Normal, nil
	}
}

func (a *BokerAccount) GetTeam() common.Address {

	return a.teamAccount
}

func (a *BokerAccount) SetOwner(txAddress, ownerAddress common.Address) error {

	if a.ownerAccount == common.HexToAddress("") {
		a.ownerAccount = ownerAddress
	} else {
		if txAddress != a.ownerAccount {
			return errors.New("")
		} else {
			a.ownerAccount = ownerAddress
			return nil
		}
	}
	return nil
}

func (a *BokerAccount) SetCommunity(txAddress, communityAddress common.Address) error {

	if a.ownerAccount == common.HexToAddress("") {
		return errors.New("")
	} else {
		if txAddress != a.ownerAccount {
			return errors.New("")
		} else {
			a.communityAccount = communityAddress
			return nil
		}
	}
	return nil

}

func (a *BokerAccount) SetFoundation(txAddress, foundationAddress common.Address) error {

	if a.ownerAccount == common.HexToAddress("") {
		return errors.New("")
	} else {
		if txAddress != a.ownerAccount {
			return errors.New("")
		} else {
			a.foundationAccount = foundationAddress
			return nil
		}
	}
	return nil
}

func (a *BokerAccount) SetTeam(txAddress, teamAddress common.Address) error {
	if a.ownerAccount == common.HexToAddress("") {
		return errors.New("")
	} else {
		if txAddress != a.ownerAccount {
			return errors.New("")
		} else {
			a.teamAccount = teamAddress
			return nil
		}
	}
	return nil
}
