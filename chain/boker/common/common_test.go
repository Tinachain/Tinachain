package common

import (
	"fmt"
	"math/big"
	"testing"
)

func BigIntsToInt64s(bigints []*big.Int) (int64s []int64) {
	for _, bigint := range bigints {
		int64s = append(int64s, bigint.Int64())
	}
	return int64s
}

func commonTest() {
	fmt.Println(BigIntsToInt64s([]*big.Int{big.NewInt(5), big.NewInt(6)}))
}

func Test(t *testing.T) {
	commonTest()
}
