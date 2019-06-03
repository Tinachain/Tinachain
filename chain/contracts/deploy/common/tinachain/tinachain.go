package tinachain

import (
	"crypto/ecdsa"
	crand "crypto/rand"
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/Tinachain/Tina/chain/accounts/keystore"
	ethcommon "github.com/Tinachain/Tina/chain/common"
	"github.com/Tinachain/Tina/chain/crypto"
	"github.com/pborman/uuid"
)

var (
	ZeroAddress       ethcommon.Address = ethcommon.Address{}
	ZeroAddressString string            = ZeroAddress.String()
)

var clientDefault *Client

func Initialize(rpc string, keyJson, password string, managerAddress string) (err error) {
	clientDefault, err = NewClient(rpc)
	if err != nil {
		return fmt.Errorf("NewClient %v", err.Error())
	}

	//Unlock
	if keyJson != "" {
		err = clientDefault.Unlock(keyJson, password)
		if err != nil {
			return fmt.Errorf("Unlock %v", err.Error())
		}
	}

	return nil
}

func HexToAddress(address string) ethcommon.Address {
	return ethcommon.HexToAddress(address)
}

func IsAddress(address string) bool {
	if !ethcommon.IsHexAddress(address) {
		return false
	}

	return (ethcommon.FromHex(address) != nil)
}

func BigIntsToInt64s(bigints []*big.Int) (int64s []int64) {
	for _, bigint := range bigints {
		int64s = append(int64s, bigint.Int64())
	}
	return int64s
}

func parseAddressByKey(key *keystore.Key) (address string, privateKey string) {
	address = hex.EncodeToString(key.Address.Bytes())
	privateKey = hex.EncodeToString(key.PrivateKey.D.Bytes())
	return address, privateKey
}

func ParseAddressByJson(keystoreJson string, password string) (address string, privateKey string, err error) {
	key, err := keystore.DecryptKey([]byte(keystoreJson), password)
	if err != nil {
		return "", "", err
	}

	address, privateKey = parseAddressByKey(key)
	return address, privateKey, nil
}

func NewKeystore(privateKey string, password string) (address string, keystoreJson string, err error) {
	pkBytes, err := hex.DecodeString(privateKey)
	if err != nil {
		return "", "", err
	}
	privateKeyECDSA, err := crypto.ToECDSA(pkBytes)
	if err != nil {
		return "", "", err
	}
	id := uuid.NewRandom()
	key := &keystore.Key{
		Id:         id,
		Address:    crypto.PubkeyToAddress(privateKeyECDSA.PublicKey),
		PrivateKey: privateKeyECDSA,
	}
	keystoreBytes, err := keystore.EncryptKey(key, password, keystore.StandardScryptN, keystore.StandardScryptP)
	if err != nil {
		return "", "", err
	}

	address = hex.EncodeToString(key.Address.Bytes())
	keystoreJson = string(keystoreBytes)
	return address, keystoreJson, nil
}

func ParseAddressByPrivatekey(privateKey string) (address string, err error) {
	pkBytes, err := hex.DecodeString(privateKey)
	if err != nil {
		return "", err
	}
	privateKeyECDSA, err := crypto.ToECDSA(pkBytes)
	if err != nil {
		return "", err
	}
	address = hex.EncodeToString(crypto.PubkeyToAddress(privateKeyECDSA.PublicKey).Bytes())
	return address, nil
}

func NewAddress(password string) (address string, keystoreJson string, err error) {
	privateKeyECDSA, err := ecdsa.GenerateKey(crypto.S256(), crand.Reader)
	if err != nil {
		return "", "", err
	}
	id := uuid.NewRandom()
	key := &keystore.Key{
		Id:         id,
		Address:    crypto.PubkeyToAddress(privateKeyECDSA.PublicKey),
		PrivateKey: privateKeyECDSA,
	}
	keystoreBytes, err := keystore.EncryptKey(key, password, keystore.StandardScryptN, keystore.StandardScryptP)
	if err != nil {
		return "", "", err
	}

	address = hex.EncodeToString(key.Address.Bytes())
	keystoreJson = string(keystoreBytes)
	return address, keystoreJson, nil
}

//func ContractDeploy(contractInfo *ContractInfo, params ...interface{}) (ethcommon.Address, error) {
//	if clientDefault == nil {
//		return ethcommon.Address{}, fmt.Errorf("clientDefault not initialized!")
//	}
//	return clientDefault.ContractDeploy(contractInfo, params...)
//}

func GetBalance(address string) (*big.Int, error) {
	if clientDefault == nil {
		return nil, fmt.Errorf("clientDefault not initialized!")
	}
	return clientDefault.GetBalance(address)
}
