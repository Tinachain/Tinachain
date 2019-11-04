package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	"github.com/Tinachain/Tina/chain/boker/agent/business"
	"github.com/Tinachain/Tina/chain/boker/agent/web"
	"github.com/Tinachain/Tina/chain/boker/common"
	log4plus "github.com/Tinachain/Tina/chain/boker/common/log4go"
	ethcommon "github.com/Tinachain/Tina/chain/common"
)

var (
	ConfigName string = "agent.json"
)

type ContractConfig struct {
	Listen           string            `json:"listen"`
	RPC              string            `json:"rpc"`
	Keystore         string            `json:"keystore"`
	Password         string            `json:"password"`
	ManagerAddress   ethcommon.Address `json:"manageraddress"`
	InterfaceAddress ethcommon.Address `json:"interfaceaddress"`
}

type ChainwareBackend struct {
	config   ContractConfig
	Keystore string
}

func New() *ChainwareBackend {

	chainware := new(ChainwareBackend)
	buffer, err := ioutil.ReadFile(ConfigName)
	if err != nil {
		return nil
	}

	config := ContractConfig{}
	err = json.Unmarshal(buffer, &config)
	if err != nil {
		log4plus.Info("main.go Unmarshal failed err=%s", err)
		return nil
	}

	fileName, _ := os.Open(config.Keystore)
	bytes, _ := ioutil.ReadAll(fileName)
	chainware.Keystore = string(bytes)
	chainware.config = config

	log4plus.Info("main.go New config Listen=%s RPC=%s Keystore=%s Password=%s ManagerAddress=%s InterfaceAddress=%s",
		chainware.config.Listen,
		chainware.config.RPC,
		chainware.config.Keystore,
		chainware.config.Password,
		chainware.config.ManagerAddress.String(),
		chainware.config.InterfaceAddress.String())

	return chainware
}

func getExeName() string {
	ret := ""
	ex, err := os.Executable()
	if err == nil {
		ret = filepath.Base(ex)
	}
	return ret
}

func setLog() {
	logJson := "log.json"
	set := false
	if bExist := common.PathExist(logJson); bExist {
		if err := log4plus.SetupLogWithConf(logJson); err == nil {
			set = true
		}
	}

	if !set {
		fileWriter := log4plus.NewFileWriter()
		exeName := getExeName()
		fileWriter.SetPathPattern("./log/" + exeName + "-%Y%M%D.log")
		log4plus.Register(fileWriter)
		log4plus.SetLevel(log4plus.DEBUG)
	}
}

func main() {

	setLog()
	defer log4plus.Close()

	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log4plus.Error("main.go get current dir err=%s", err)
		return
	}
	log4plus.Info("main.go current dir=%s", dir)

	chainware := New()
	if chainware != nil {

		log4plus.Info("main.go initchain")
		if err := business.InitChain(chainware.config.RPC,
			chainware.Keystore,
			chainware.config.Password,
			chainware.config.ManagerAddress.String(),
			chainware.config.InterfaceAddress.String()); err != nil {

			log4plus.Error("main.go initchain Failed err=%s", err)
			return
		}

		w := web.New(chainware.config.RPC)
		log4plus.Info("main.go web.New")
		w.Run(chainware.config.Listen)
		log4plus.Info("main.go web.Run")

	}

	for {
		time.Sleep(1 * time.Second)
	}

}
