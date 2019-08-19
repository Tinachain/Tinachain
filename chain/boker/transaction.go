//Tina链主动产生交易
package boker

import (
	"bytes"
	"context"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"

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

//补码
func PKCS7Padding(ciphertext []byte, blocksize int) []byte {
	padding := blocksize - len(ciphertext)%blocksize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

//去码
func PKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

func (t *BokerTransaction) EncoderContext(context []byte, k []byte) (error, []byte) {
	// 分组秘钥
	block, _ := aes.NewCipher(k)
	// 获取秘钥块的长度
	blockSize := block.BlockSize()
	// 补全码
	context = PKCS7Padding(context, blockSize)
	// 加密模式
	blockMode := cipher.NewCBCEncrypter(block, k[:blockSize])
	// 创建数组
	cryted := make([]byte, len(context))
	// 加密
	blockMode.CryptBlocks(cryted, context)

	return nil, []byte(base64.StdEncoding.EncodeToString(cryted))
}

func (t *BokerTransaction) DecoderContext(context []byte, k []byte) (error, []byte) {

	// 转成字节数组
	crytedByte, _ := base64.StdEncoding.DecodeString(string(context[:]))

	// 分组秘钥
	block, _ := aes.NewCipher(k)
	// 获取秘钥块的长度
	blockSize := block.BlockSize()
	// 加密模式
	blockMode := cipher.NewCBCDecrypter(block, k[:blockSize])
	// 创建数组
	orig := make([]byte, len(crytedByte))
	// 解密
	blockMode.CryptBlocks(orig, crytedByte)
	// 去补全码
	orig = PKCS7UnPadding(orig)
	return nil, []byte(string(orig))
}

func (t *BokerTransaction) SubmitBokerTransaction(ctx context.Context,
	txMajor protocol.TxMajor,
	txMinor protocol.TxMinor,
	from common.Address,
	to common.Address,
	name []byte,
	extra []byte,
	value *big.Int,
	encryption uint8) (*types.Transaction, error) {

	if t.ethereum != nil {

		//设置参数（其中有些参数可以通过调用设置默认设置来进行获取）
		args := ethapi.SendTxArgs{
			From:       from,
			Major:      txMajor,
			Minor:      txMinor,
			Nonce:      nil,
			To:         &to,
			Gas:        nil,
			GasPrice:   nil,
			Value:      (*hexutil.Big)(value),
			Name:       hexutil.Bytes(name),
			Encryption: encryption,
			Data:       hexutil.Bytes(extra),
			Extra:      hexutil.Bytes(extra),
		}

		//查找包含所请求签名者的钱包
		account := accounts.Account{Address: args.From}

		//根据帐号得到钱包信息
		wallet, err := t.ethereum.AccountManager().Find(account)
		if err != nil {
			log.Error("SubmitBokerTransaction AccountManager Find", "error", err)
			return nil, err
		}

		//设置默认设置
		if err := args.SetDefaults(ctx, t.ethereum.ApiBackend); err != nil {
			log.Error("SubmitBokerTransaction SetDefaults", "error", err)
			return nil, err
		}
		log.Info("SubmitBokerTransaction SetDefaults",
			"Nonce", args.Nonce.String(),
			"Major", args.Major,
			"Minor", args.Minor,
			"Name", args.Name,
			"Extra", args.Extra)

		var chainID *big.Int
		var tx *types.Transaction

		//根据交易类型进行区分
		if protocol.SystemBase == txMajor || protocol.UserBase == txMajor {

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
				args.Name,
				args.Extra,
				args.Encryption)
			if config := t.ethereum.ApiBackend.ChainConfig(); config.IsEIP155(t.ethereum.ApiBackend.CurrentBlock().Number()) {

				chainID = config.ChainId
			}
			log.Info("SubmitBokerTransaction tx", "Major", tx.Major(), "Miner", tx.Minor(), "Extra", tx.Extra())

		} else if protocol.Stock == txMajor {

			tx = types.NewStockTransaction(args.Major,
				args.Minor,
				(uint64)(*args.Nonce),
				(common.Address)(*args.To),
				(*big.Int)(args.Value),
				new(big.Int).SetUint64(defaultGas),
				new(big.Int).SetUint64(defaultGasPrice),
				args.Name,
				args.Extra,
				args.Encryption)
			if config := t.ethereum.ApiBackend.ChainConfig(); config.IsEIP155(t.ethereum.ApiBackend.CurrentBlock().Number()) {

				chainID = config.ChainId
			}
			log.Info("SubmitBokerTransaction tx", "Major", tx.Major(), "Miner", tx.Minor(), "Extra", tx.Extra())
		} else {

			log.Error("SubmitBokerTransaction txType Not Found")
			return nil, errors.New("SubmitBokerTransaction txType Not Found")
		}

		log.Info("SubmitBokerTransaction NewTransaction", "Value", tx.Value, "tx.Hash", tx.Hash().String())

		//对该笔交易签名来确保该笔交易的真实有效性
		signed, err := wallet.SignTxWithPassphrase(account, t.ethereum.Password(), tx, chainID)
		if err != nil {

			log.Error("SubmitBokerTransaction SignTxWithPassphrase", "error", err)
			return nil, err
		}

		log.Info("SubmitBokerTransaction SignTxWithPassphrase",
			"Gas", signed.Gas(),
			"GasPrice", signed.GasPrice(),
			"Value", signed.Value,
			"Nonce", signed.Nonce(),
			"To", signed.To(),
			"tx.Hash", signed.Hash().String())

		if _, err := ethapi.SubmitTransaction(ctx, t.ethereum.ApiBackend, signed); err != nil {

			log.Error("SubmitBokerTransaction SubmitTransaction", "error", err)
			return nil, err
		}

		return signed, nil
	}
	return nil, protocol.ErrInvalidSystem
}
