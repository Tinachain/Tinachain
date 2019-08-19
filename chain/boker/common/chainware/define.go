package chainware

const (
	BobbyAccuracy int64 = 1e18 //bobby币精度, 小数点后的位数
	BowerAccuracy int64 = 1e3  //算力精度，小数点后的位数

	DefaultPassword string = "12345678"
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

const (
	UserEventRegister      = iota //注册
	UserEventLoginDaily           //每日登录
	UserEventBindDapp             //绑定渠道
	UserEventWatch                //观看视频
	UserEventUpload               //上传视频
	UserEventCertification        //实名认证
)

const (
	FinanceReasonTransfer        = iota //普通转账
	FinanceReasonMine                   //挖矿所得
	FinanceReasonAssignToken            //分币
	FinanceReasonVote                   //投票
	FinanceReasonVoteCancel             //取消投票
	FinanceReasonVoteUnlock             //投票解锁
	FinanceReasonFinanceWithdraw        //退款
	FinanceReasonTip                    //打赏
	FinanceReasonExchange               //兑换
)

const (
	PowerTypeLongterm  = iota //长期算力
	PowerTypeShortterm        //临时算力
)

const (
	ConfigExchangeRate = "exchange_rate"
)
