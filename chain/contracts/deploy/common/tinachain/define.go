package tinachain

import (
	"math/big"
)

const (
	UserPowerReasonRegister            = iota //注册
	UserPowerReasonInvited                    //被邀请，填写邀请码
	UserPowerReasonInvitor                    //邀请别人
	UserPowerReasonLoginDaily                 //每日登录
	UserPowerReasonBindDapp                   //绑定渠道
	UserPowerReasonCertification              //实名认证
	UserPowerReasonWatch                      //观看视频
	UserPowerReasonVideoWatched               //视频被观看
	UserPowerReasonUpload                     //上传视频
	UserPowerReasonShorttermPowerReset        //临时算力重置
)

var (
	BobbyAccuracy = big.NewInt(1e18) //bobby币精度, 小数点后的位数
	PowerAccuracy = big.NewInt(1e3)  //算力精度，小数点后的位数

	PowerTypeLongterm  = big.NewInt(0)
	PowerTypeShortterm = big.NewInt(1)
)

const (
	DefaultPassword = "12345678"
)
