package protocol

import (
	"encoding/hex"
	"errors"
	"fmt"
	_ "io/ioutil"
	"math/big"
	_ "net/http"
	"reflect"
	"strings"
	"time"

	"github.com/Tinachain/Tina/chain/accounts/abi"
	"github.com/Tinachain/Tina/chain/common"
	"github.com/Tinachain/Tina/chain/log"
)

const (
	ExtraVanity        = 32   //扩展字段的前缀字节数量
	ExtraSeal          = 65   //扩展字段的后缀字节数量
	InmemorySignatures = 4096 //保留在内存中的最近块签名的数量
)

const (
	MaxValidatorSize = 1 //DPOS的验证者数量
	ConsensusSize    = 1 //共识确认验证者数量
)

const (
	BlacksInterval   = int64(300)       //黑名单周期
	EpochInterval    = int64(86400)     //验证者周期
	GasInterval      = int64(300)       //股权周期
	BlockInterval    = int64(5)         //打包周期
	MinuteTimer      = time.Minute      //分
	SecondTimer      = time.Second      //秒
	MillisecondTimer = time.Millisecond //毫秒

)

//主要交易类型
type TxMajor uint8

const (
	Normal     TxMajor = iota //普通交易类型
	SystemBase                //基础交易类型（不使用Gas的交易）
	UserBase                  //用户基础交易
	Extra                     //扩展交易类型（可以在区块中存放文件类型的）
	Stock                     //股权类型，用来定义用户股票权益信息（没有Gas消费）
)

//次要交易类型
type TxMinor uint8

//基础交易的次要类型
const (
	MinMinor          TxMinor = iota
	SetValidator              //设置验证者
	SetSystemContract         //设置系统基础合约
	RegisterCandidate         //注册成为候选人(用户注册为候选人)
	VoteUser                  //用户投票
	VoteCancel                //用户取消投票
	VoteEpoch                 //产生当前的出块节点(在每次周期产生的时候触发)
	Timeout                   //超时处理
	MaxMinor                  //最大值
)

//用户基础交易的次要类型
const (
	SetUserContract TxMinor = iota
	CancelUserContract
)

//扩展交易的次要类型
const (
	Word TxMinor = iota
	Data
)

//股权类型的次要类型
const (
	StockManager   TxMinor = iota //股权管理者
	StockSet                      //设置股权
	StockTransfer                 //转移部分股权
	StockClean                    //销毁股权(将股权从某个账号中强行销毁)
	StockFrozen                   //冻结股权(股权冻结，则此股权将无法获得Gas收益)
	StockUnFrozen                 //解冻股权
	StockAssignGas                //股权分币
)

//股权产生方式
type StockCreate uint8

const (
	New      StockCreate = iota //设置股权
	Transfer                    //转移部分股权
)

//新增合约类型
type BaseContractType uint8

const (
	System BaseContractType = iota //系统基础合约
	User                           //个人基础合约
)

type StockState uint8

const (
	Run StockState = iota
	Frozen
)

var (
	TinaUnit           *big.Int = big.NewInt(1e+12) //Tina的单位
	TransferUnit       *big.Int = big.NewInt(1e+17) //转账单位(这个数值仅用于每次给指定账号，方便指定账号给用户分配通证)
	TransferMultiple   *big.Int = big.NewInt(165)   //转账倍数
	SetValidatorVotes  *big.Int = big.NewInt(10000)
	MaxGasPrice        *big.Int = new(big.Int).SetUint64(0xffffffffffffffff) //最大的GasPrice
	MaxGasLimit        *big.Int = new(big.Int).SetUint64(0)                  //最大的GasLimit
	TimeOfFirstBlock            = int64(0)                                   //创世区块的时间偏移量
	ConfirmedBlockHead          = []byte("confirmed-block-head")
	MaxWordSize                 = int64(1 * 1024 * 1024)
	MaxDataSize                 = int(1 * 1024 * 1024)
	MaxExtraSize                = int64(5 * 1024 * 1024)
	MaxNormalSize               = common.StorageSize(32 * 1024)
	MaxBlockSize                = int64(5 * 1024 * 1024)
)

var (
	//用户触发的合约方法名
	RegisterCandidateMethod = "registerCandidate" //申请候选节点
	VoteCandidateMethod     = "voteCandidate"     //投票候选节点
	CancelVoteMethod        = "cancelAllVotes"    //取消所有投票
)

var (
	//基础链触发的基础合约
	RotateVoteMethod    = "rotateVote"    //产生当前的出块节点(在每次周期产生的时候触发)
	TickCandidateMethod = "tickVote"      //投票时钟
	GetCandidateMethod  = "getCandidates" //获取候选人结果
)

//周期验证者相关
var (
	EpochPrefix      = []byte("epoch")      //存放周期信息
	ValidatorPrefix  = []byte("validator")  //存放验证者投票信息
	VotePrefix       = []byte("vote")       //存放投票数量
	ValidatorsPrefix = []byte("validators") //存放所有的验证者列表

)

//合约相关
var (
	SingleContractPrefix = []byte("single")    //存放单个合约信息
	ContractsPrefix      = []byte("contracts") //存放所有合约信息
)

//股权相关
var (
	SingleStockPrefix = []byte("stock")
	StocksPrefix      = []byte("stocks")
	OwnerPrefix       = []byte("owner")
	GasPoolPrefix     = []byte("gasPool")
)

var (
	ErrNilBlockHeader             = errors.New("nil block header returned")                       //区块头为空
	ErrUnknownBlock               = errors.New("unknown block")                                   //未知区块
	ErrInvalidProducer            = errors.New("invalid current producer")                        //出块节点出错
	ErrInvalidProducerTime        = errors.New("invalid time to mint the block")                  //不正确的出块时间
	ErrInvalidCoinbase            = errors.New("invalid current mining coinbase")                 //当前挖矿账号错误
	ErrInvalidSystem              = errors.New("invalid current system")                          //当前系统的投票合约出错
	ErrMismatchSignerAndValidator = errors.New("mismatch block signer and validator")             //签名者和区块头中的验证者不是同一个
	ErrNoSigner                   = errors.New("missing signing methods")                         //缺少签名方法
	ErrInvalidType                = errors.New("invalid transaction type")                        //无效的交易类型
	ErrInvalidAddress             = errors.New("invalid transaction payload address")             //无效的交易有效负载地
	ErrInvalidAction              = errors.New("invalid transaction payload action")              //无效的事务有效负载操
	ErrLoadConfig                 = errors.New("load bokerchain config error")                    //加载配置信息出错
	ErrNotFoundAddress            = errors.New("not found bokerchain contract address")           //没有找到合约地址
	ErrNotFoundType               = errors.New("not found bokerchain contract type")              //没有找到合约类型
	ErrWriteJson                  = errors.New("write bokerchain json file error")                //写保存基础合约的Json格式出错
	ErrOpenFile                   = errors.New("open bokerchain json file error")                 //打开基础合约保存文件出错
	ErrWriteFile                  = errors.New("bokerchain write file error")                     //写基础合约保存文件出错
	ErrContractExist              = errors.New("bokerchain contract aleady exist")                //写基础合约保存文件出错
	ErrSystem                     = errors.New("system error")                                    //系统错误
	ErrNotFoundContract           = errors.New("not found bokerchain contract")                   //没有找到合约
	ErrNotFoundAccount            = errors.New("not found bokerchain account")                    //没有找到合约
	ErrNewContractService         = errors.New("create bokerchain base contract err")             //没有找到合约
	ErrSaveContractTrie           = errors.New("save contract trie err")                          //没有找到合约
	ErrLevel                      = errors.New("account level error")                             //没有找到合约
	ErrPointerIsNil               = errors.New("Trie Pointer is Nil")                             //Hash树指针是nil
	ErrTransactionType            = errors.New("Error Transaction Type")                          //交易类型错误
	ErrSpecialAccount             = errors.New("Current Account is`t BokerChain Special Account") //当前账号不是指定的特殊账号
	ErrValidatorsIsFull           = errors.New("Current Validators is Full")                      //当前验证者数量已满
	ErrExistsValidators           = errors.New("Current Validators Exists")                       //当前存在验证者
	ErrGenesisBlock               = errors.New("not genesis block")                               //区块需要不为0，即最近区块不是创世区块，证明已经工作
	ErrDecodeValidators           = errors.New("failed to decode validators")                     //解码验证者失败
	ErrEncodeValidators           = errors.New("failed to encode validators")                     //编码验证者失败
	ErrSetEpochTrieFail           = errors.New("failed set epoch trie")                           //设置周期树失败
	ErrEpochTrieNil               = errors.New("failed to producers length is zero")              //出块节点长度为0
	ErrToIsNil                    = errors.New("setValidator block header to is nil")             //设置验证者区块头为nil
	ErrTxType                     = errors.New("failed to tx type")                               //交易类型失败
	ErrIsnStock                   = errors.New("not is stock account")                            //不是股权账号
	ErrIsnOwner                   = errors.New("coinbase not is owner of chain")
	ErrStockLow                   = errors.New("account stock too low")
)

type StockRewards struct {
	Timer  int64
	Number uint64
}
type StockAccount struct {
	Account common.Address
	Number  uint64
	State   StockState
}

type StocksAccount struct {
	Stock []*StockAccount
}

//设置Tina链配置
type BokerConfig struct {
	Address common.Address
}

//Abi函数参数信息
type ParamJson struct {
	Name  string `json:"name"`  //参数名称
	Type  string `json:"type"`  //参数类型
	Value string `json:"value"` //参数值
}
type MethodJson struct {
	Method string      `json:"method"`           //方法名称
	Params []ParamJson `json:"params,omitempty"` //方法参数列表
}

func getInterface(input abi.Argument) reflect.Value {

	//设置参数类型
	var param reflect.Value

	switch input.Type.T {
	case abi.SliceTy:

		//log.Info("DecodeAbi abi.SliceTy")
		var paramType = []byte{}
		param = reflect.New(reflect.TypeOf(paramType))
	case abi.StringTy:

		//log.Info("DecodeAbi abi.StringTy")
		var paramType = string("")
		param = reflect.New(reflect.TypeOf(paramType))
	case abi.IntTy:

		switch input.Type.Type {

		case reflect.TypeOf(int8(0)):

			var paramType = int8(0)
			param = reflect.New(reflect.TypeOf(paramType))
		case reflect.TypeOf(int16(0)):

			var paramType = int16(0)
			param = reflect.New(reflect.TypeOf(paramType))
		case reflect.TypeOf(int32(0)):

			var paramType = int32(0)
			param = reflect.New(reflect.TypeOf(paramType))
		case reflect.TypeOf(int64(0)):

			var paramType = int64(0)
			param = reflect.New(reflect.TypeOf(paramType))
		case reflect.TypeOf(&big.Int{}):

			var paramType = big.NewInt(1)
			param = reflect.New(reflect.TypeOf(paramType))
		}
	case abi.UintTy:

		//log.Info("DecodeAbi abi.UintTy", "input.Type.Type", input.Type.Type)
		switch input.Type.Type {

		case reflect.TypeOf(uint8(0)):

			var paramType = uint8(0)
			param = reflect.New(reflect.TypeOf(paramType))
		case reflect.TypeOf(uint16(0)):

			var paramType = uint16(0)
			param = reflect.New(reflect.TypeOf(paramType))
		case reflect.TypeOf(uint32(0)):

			var paramType = uint32(0)
			param = reflect.New(reflect.TypeOf(paramType))
		case reflect.TypeOf(uint64(0)):

			var paramType = uint64(0)
			param = reflect.New(reflect.TypeOf(paramType))
		case reflect.TypeOf(&big.Int{}):

			var paramType = big.NewInt(1)
			param = reflect.New(reflect.TypeOf(paramType))
		}
	case abi.BoolTy:

		//log.Info("DecodeAbi abi.BoolTy")
		var paramType = bool(true)
		param = reflect.New(reflect.TypeOf(paramType))
	case abi.AddressTy:

		//log.Info("DecodeAbi abi.AddressTy")
		var paramType = common.Address{}
		param = reflect.New(reflect.TypeOf(paramType))
	case abi.BytesTy:

		//log.Info("DecodeAbi abi.BytesTy")
		var paramType = common.Hex2Bytes("0100000000000000000000000000000000000000000000000000000000000000")
		param = reflect.New(reflect.TypeOf(paramType))
	case abi.FunctionTy:

		//log.Info("DecodeAbi abi.FunctionTy")
		var paramType = common.Hex2Bytes("0100000000000000000000000000000000000000000000000000000000000000")
		param = reflect.New(reflect.TypeOf(paramType))
	}
	return param
}

func GetParamCount(abiJson string, name string) int {

	abiDecoder, err := abi.JSON(strings.NewReader(abiJson))
	if err != nil {
		return 0
	}
	return len(abiDecoder.Methods[name].Inputs)
}

func DecodeAbi(abiJson string, name string, payload string) (MethodJson, error) {

	const definition = `[{"constant":false,"inputs":[{"name":"","type":"int256"},{"name":"str","type":"string"}],"name":"test","outputs":[{"name":"","type":"string"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[],"name":"show","outputs":[{"name":"","type":"string"}],"payable":false,"stateMutability":"view","type":"function"},{"payable":false,"stateMutability":"nonpayable","type":"fallback"}]`
	payload = "069fd0a30000000000000000000000000000000000000000000000000000000000000005000000000000000000000000000000000000000000000000000000000000004000000000000000000000000000000000000000000000000000000000000000046861686100000000000000000000000000000000000000000000000000000000"
	name = "test"

	//解析Abi格式成为Json格式
	abiDecoder, err := abi.JSON(strings.NewReader(definition))
	if err != nil {
		return MethodJson{}, err
	}
	log.Info("DecodeAbi", "Methods Inputs count", len(abiDecoder.Methods[name].Inputs))

	//剔除最前面的0x标记
	var decodeString string = ""
	hexFlag := strings.Index(payload, "0x")
	if hexFlag == -1 {
		decodeString = payload
	} else {
		decodeString = payload[2:]
	}

	//将字符串转换成[]Byte
	decodeBytes, err := hex.DecodeString(decodeString)
	if err != nil {
		return MethodJson{}, err
	}
	log.Info("DecodeAbi", "decodeBytes", decodeBytes)

	//根据函数的名称，设置函数的输入参数信息
	method, ok := abiDecoder.Methods[name]
	if !ok {
		return MethodJson{}, errors.New("")
	}

	//写入获取参数类型
	params := make([]interface{}, 0)
	for i := 0; i < len(method.Inputs); i++ {

		input := method.Inputs[i]

		//设置参数类型
		var param reflect.Value

		switch input.Type.T {
		case abi.SliceTy:

			log.Info("DecodeAbi abi.SliceTy")
			var paramType = []byte{}
			param = reflect.New(reflect.TypeOf(paramType))
		case abi.StringTy:

			log.Info("DecodeAbi abi.StringTy")
			var paramType = string("")
			param = reflect.New(reflect.TypeOf(paramType))
		case abi.IntTy:

			log.Info("DecodeAbi abi.IntTy", "input.Type.Type", input.Type.Type)
			switch input.Type.Type {

			case reflect.TypeOf(int8(0)):

				var paramType = int8(0)
				param = reflect.New(reflect.TypeOf(paramType))
			case reflect.TypeOf(int16(0)):

				var paramType = int16(0)
				param = reflect.New(reflect.TypeOf(paramType))
			case reflect.TypeOf(int32(0)):

				var paramType = int32(0)
				param = reflect.New(reflect.TypeOf(paramType))
			case reflect.TypeOf(int64(0)):

				var paramType = int64(0)
				param = reflect.New(reflect.TypeOf(paramType))
			case reflect.TypeOf(&big.Int{}):

				var paramType = big.NewInt(1)
				param = reflect.New(reflect.TypeOf(paramType))
			}
		case abi.UintTy:

			log.Info("DecodeAbi abi.UintTy", "input.Type.Type", input.Type.Type)
			switch input.Type.Type {

			case reflect.TypeOf(uint8(0)):

				var paramType = uint8(0)
				param = reflect.New(reflect.TypeOf(paramType))
			case reflect.TypeOf(uint16(0)):

				var paramType = uint16(0)
				param = reflect.New(reflect.TypeOf(paramType))
			case reflect.TypeOf(uint32(0)):

				var paramType = uint32(0)
				param = reflect.New(reflect.TypeOf(paramType))
			case reflect.TypeOf(uint64(0)):

				var paramType = uint64(0)
				param = reflect.New(reflect.TypeOf(paramType))
			case reflect.TypeOf(&big.Int{}):

				var paramType = big.NewInt(1)
				param = reflect.New(reflect.TypeOf(paramType))
			}
		case abi.BoolTy:

			log.Info("DecodeAbi abi.BoolTy")
			var paramType = bool(true)
			param = reflect.New(reflect.TypeOf(paramType))
		case abi.AddressTy:

			log.Info("DecodeAbi abi.AddressTy")
			var paramType = common.Address{}
			param = reflect.New(reflect.TypeOf(paramType))
		case abi.BytesTy:

			log.Info("DecodeAbi abi.BytesTy")
			var paramType = common.Hex2Bytes("0100000000000000000000000000000000000000000000000000000000000000")
			param = reflect.New(reflect.TypeOf(paramType))
		case abi.FunctionTy:

			log.Info("DecodeAbi abi.FunctionTy")
			var paramType = common.Hex2Bytes("0100000000000000000000000000000000000000000000000000000000000000")
			param = reflect.New(reflect.TypeOf(paramType))
		}
		params = append(params, param.Interface())
	}

	//解码
	if err := abiDecoder.InputUnpack(params, name, decodeBytes[4:]); err != nil {
		log.Error("DecodeAbi ", "err", err)
		return MethodJson{}, err
	}

	//将返回的信息放入到Json格式中
	json := MethodJson{}
	json.Method = name
	json.Params = make([]ParamJson, 0)

	for i := 0; i < len(params); i++ {

		valueOf := reflect.ValueOf(params[i])
		out := valueOf.Elem().Interface()
		s := fmt.Sprintf("%v", out)

		param := ParamJson{
			Name:  abiDecoder.Methods[name].Inputs[i].Name,
			Type:  abiDecoder.Methods[name].Inputs[i].Type.String(),
			Value: s,
		}
		json.Params = append(json.Params, param)
	}

	return json, nil
}

//此处需要注意，由于使用了IP地址，所以会造成创世区块的Hash每个节点计算的结果不一致，从而造成无法进行互联的情况。
func GetExternalIp() string {

	/*resp, err := http.Get("http://myexternalip.com/raw")
	if err != nil {
		return ""
	}
	defer resp.Body.Close()
	content, _ := ioutil.ReadAll(resp.Body)
	return string(content)*/
	return ""
}
