package tinachain

import (
	"blockchain/Tinachain/common/config"
	"fmt"
	"testing"
	//	ethcommon "github.com/boker/chain/common"
)

func Test(t *testing.T) {
	fmt.Println("Test")
	config.Initialize()

	addr1, ks, err := NewKeystore("35992401c08a6962eabf36e519e9c2e1f2e333b0189174bd249097d0f0b67bd9", "123456")
	fmt.Println(addr1, ks, err)
}
