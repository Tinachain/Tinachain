package account

import (
	"time"

	"github.com/Tinachain/Tina/chain/contracts/deploy/common/tinachain"
)

type AccountInfo struct {
	BcAddress  string    "bcAddress"
	Keystore   string    "keystore"
	CreateTime time.Time "createTime"

	Password string
}

func NewAccount(accountInfo *AccountInfo) (err tinachain.Error) {
	//create boker chain account
	var e error
	accountInfo.BcAddress, accountInfo.Keystore, e = tinachain.NewAddress(accountInfo.Password)
	if e != nil {
		return tinachain.NewError(tinachain.ErrorServerError, "NewAddress err=%s", e.Error())
	}

	accountInfo.CreateTime = time.Now()
	return nil
}
