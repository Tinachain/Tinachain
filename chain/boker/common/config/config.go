package config

import (
	"io/ioutil"
	"os"

	"github.com/Tinachain/Tina/chain/boker/common/goini"
)

type Config struct {
	Listen string
	Upload string

	MysqlIp   string
	MysqlPort string
	MysqlUser string
	MysqlPwd  string
	MysqlDb   string

	RedisIp        string
	RedisPort      string
	RedisPwd       string
	RedisMaxIdle   int
	RedisMaxActive int

	IpfsGateway string
	IpfsApi     string

	AliyunLocation        string
	AliyunEndpoint        string
	AliyunAccessKeyId     string
	AliyunAccessKeySecret string
	AliyunBucketName      string
	AliyunPipelineId      string

	BokerchainRpc               string
	BokerchainManagerAddress    string
	BokerchainAdminKeystoreFile string
	BokerchainAdminKeystore     string
	BokerchainAdminPassword     string
	BokerchainSolFolder         string
}

var _cfg *Config = &Config{}

func GetInstance() *Config {
	return _cfg
}

func Initialize() {
	iniFile := goini.Init("config.ini")

	_cfg.Listen = iniFile.Read_string("common", "listen", "0.0.0.0:8080")
	_cfg.Upload = iniFile.Read_string("common", "upload", "")

	_cfg.MysqlIp = iniFile.Read_string("mysql", "ip", "127.0.0.1")
	_cfg.MysqlPort = iniFile.Read_string("mysql", "port", "3306")
	_cfg.MysqlUser = iniFile.Read_string("mysql", "user", "")
	_cfg.MysqlPwd = iniFile.Read_string("mysql", "pwd", "")
	_cfg.MysqlDb = iniFile.Read_string("mysql", "db", "")

	_cfg.IpfsGateway = iniFile.Read_string("ipfs", "gateway", "")
	_cfg.IpfsApi = iniFile.Read_string("ipfs", "api", "")

	_cfg.RedisIp = iniFile.Read_string("redis", "ip", "127.0.0.1")
	_cfg.RedisPort = iniFile.Read_string("redis", "port", "6379")
	_cfg.RedisPwd = iniFile.Read_string("redis", "pwd", "")
	_cfg.RedisMaxIdle = iniFile.ReadInt("redis", "max_idle", 5)
	_cfg.RedisMaxActive = iniFile.ReadInt("redis", "max_active", 20)

	_cfg.AliyunEndpoint = iniFile.Read_string("aliyun", "endpoint", "http://oss-cn-shanghai.aliyuncs.com")
	_cfg.AliyunAccessKeyId = iniFile.Read_string("aliyun", "accessKeyId", "LTAIXe1slMwQsa8H")
	_cfg.AliyunAccessKeySecret = iniFile.Read_string("aliyun", "accessKeySecret", "hZOUYfDyJS4DKQtnvwKbMf0OAzpUpW")
	_cfg.AliyunBucketName = iniFile.Read_string("aliyun", "bucketName", "bokerchain-test")
	_cfg.AliyunLocation = iniFile.Read_string("aliyun", "location", "oss-cn-shanghai")
	_cfg.AliyunPipelineId = iniFile.Read_string("aliyun", "pipelineId", "d18240d379e244edbb2d614dd3663427")

	_cfg.BokerchainRpc = iniFile.Read_string("bokerchain", "rpc", "http://127.0.0.1:8545")
	_cfg.BokerchainManagerAddress = iniFile.Read_string("bokerchain", "manager_address", "0x7ef24bee9b343fc7f844fc9392ca8ea7dc5c2bca")
	_cfg.BokerchainAdminKeystoreFile = iniFile.Read_string("bokerchain", "admin_keystore_file", "keystore.json")
	fKeystore, _ := os.Open(_cfg.BokerchainAdminKeystoreFile)
	bytes, _ := ioutil.ReadAll(fKeystore)
	_cfg.BokerchainAdminKeystore = string(bytes)
	_cfg.BokerchainAdminPassword = iniFile.Read_string("bokerchain", "admin_password", "123456")
	_cfg.BokerchainSolFolder = iniFile.Read_string("bokerchain", "sol_folder", "")
}
