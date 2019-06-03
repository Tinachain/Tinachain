package config

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	fmt.Println("Test")
	Initialize()
	fmt.Printf("%#v", GetInstance().BokerchainKeystore)
}
