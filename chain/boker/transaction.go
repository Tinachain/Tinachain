//Tina链主动产生交易
package boker

import (
	"context"
	"errors"
	"math/big"

	"github.com/Tinachain/Tina/chain/accounts"
	"github.com/Tinachain/Tina/chain/boker/protocol"
	"github.com/Tinachain/Tina/chain/common"
	"github.com/Tinachain/Tina/chain/common/hexutil"
	"github.com/Tinachain/Tina/chain/core/types"
	"github.com/Tinachain/Tina/chain/eth"
	"github.com/Tinachain/Tina/chain/internal/ethapi"
	"github.com/Tinachain/Tina/chain/log"
	"github.com/Tinachain/Tina/chain/params"
)

const (
	defaultGas      = 90000
	defaultGasPrice = 50 * params.Shannon
)

//Tina链的基础合约管理
type BokerTransaction struct {
	ethereum *eth.Ethereum
}

func NewTransaction(ethereum *eth.Ethereum) *BokerTransaction {

	return &BokerTransaction{
		ethereum: ethereum,
	}
}

func (t *BokerTransaction) SubmitBokerTransaction(ctx context.Context,
	txMajor protocol.TxMajor,
	txMinor protocol.TxMinor,
	to common.Address,
	extra []byte) error {

	if t.ethereum != nil {

		//得到From账号
		from, err := t.ethereum.ApiBackend.Coinbase()
		if err != nil {
			log.Error("bokerTransaction CoinBase", "error", err)
			return err
		}

		//设置参数（其中有些参数可以通过调用设置默认设置来进行获取）
		args := ethapi.SendTxArgs{
			From:     from,
			Major:    txMajor,
			Minor:    txMinor,
			Nonce:    nil,
			To:       &to,
			Gas:      nil,
			GasPrice: nil,
			Value:    nil,
			Data:     hexutil.Bytes(extra),
			Extra:    hexutil.Bytes(extra),
		}

		//查找包含所请求签名者的钱包
		account := accounts.Account{Address: args.From}

		//根据帐号得到钱包信息
		wallet, err := t.ethereum.AccountManager().Find(account)
		if err != nil {
			log.Error("SubmitBokerTransaction AccountManager Find", "error", err)
			return err
		}

		//设置默认设置
		if err := args.SetDefaults(ctx, t.ethereum.ApiBackend); err != nil {
			log.Error("SubmitBokerTransaction SetDefaults", "error", err)
			return err
		}
		log.Info("SubmitBokerTransaction SetDefaults", "Nonce", args.Nonce.String(), "Major", args.Major, "Minor", args.Minor, "Extra", args.Extra)

		var chainID *big.Int
		var tx *types.Transaction

		//根据交易类型进行区分
		if protocol.Base == txMajor {

			tx = types.NewBaseTransaction(args.Major,
				args.Minor,
				(uint64)(*args.Nonce),
				(common.Address)(*args.To),
				(*big.Int)(args.Value),
				[]byte(""))
			if config := t.ethereum.ApiBackend.ChainConfig(); config.IsEIP155(t.ethereum.ApiBackend.CurrentBlock().Number()) {

				chainID = config.ChainId
			}

		} else if protocol.Extra == txMajor {

			tx = types.NewExtraTransaction(args.Major,
				args.Minor,
				(uint64)(*args.Nonce),
				(common.Address)(*args.To),
				(*big.Int)(args.Value),
				new(big.Int).SetUint64(defaultGas),
				new(big.Int).SetUint64(defaultGasPrice),
				args.Extra)
			if config := t.ethereum.ApiBackend.ChainConfig(); config.IsEIP155(t.ethereum.ApiBackend.CurrentBlock().Number()) {

				chainID = config.ChainId
			}
			log.Info("SubmitBokerTransaction tx", "Major", tx.Major(), "Miner", tx.Minor(), "Word", tx.Word(), "Extra", tx.Extra())

		} else {

			log.Error("SubmitBokerTransaction txType Not Found")
			return errors.New("SubmitBokerTransaction txType Not Found")
		}

		//对该笔交易签名来确保该笔交易的真实有效性
		signed, err := wallet.SignTxWithPassphrase(account, t.ethereum.Password(), tx, chainID)
		if err != nil {

			log.Error("SubmitBokerTransaction SignTxWithPassphrase", "error", err)
			return err
		}

		log.Info("SubmitBokerTransaction SignTxWithPassphrase", "Gas", signed.Gas(), "GasPrice", signed.GasPrice(), "Value", signed.Value)

		if _, err := ethapi.SubmitTransaction(ctx, t.ethereum.ApiBackend, signed); err != nil {

			log.Error("SubmitBokerTransaction SubmitTransaction", "error", err)
			return err
		}

		return nil
	}
	return protocol.ErrInvalidSystem
}
