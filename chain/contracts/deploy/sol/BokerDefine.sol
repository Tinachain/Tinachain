pragma solidity ^0.4.8;

contract BokerDefine {

    uint256 constant internal bobby = 1 ether;      //一个bobby币
    uint256 constant internal POWER = 1000;         //一个算力，
    uint256 constant internal CfgTickTimer = 2 * 60 * 60;   //超时处理时间;

    //配置
    string constant CfgRegisterPowerAdd = "RegisterPowerAdd";
    string constant CfgInviteCountMax = "InviteCountMax";
    string constant CfgInvitedPowerAdd = "InvitedPowerAdd";
    string constant CfgInvitorPowerAdd = "InvitorPowerAdd";
    string constant CfgLoginDailyPowerAdd = "LoginDailyPowerAdd";
    string constant CfgCertificationPowerAdd = "CertificationPowerAdd";
    string constant CfgAssignPeriod = "AssignPeriod";
    string constant CfgAssignTokenPerCycle = "AssignTokenPerCycle";
    string constant CfgAssignTokenTotal = "AssignTokenTotal";
    string constant CfgUploadCountMax = "UploadCountMax";
    string constant CfgPowerWatchOwnerRatio = "PowerWatchOwnerRatio";
    string constant CfgAssignTokenLongtermRatio = "AssignTokenLongtermRatio";
    string constant CfgVoteLockup = "VoteLockup";
    string constant CfgVoteUnlockPrecision = "VoteUnlockPrecision";
    string constant CfgVoteCyclePeriod = "VoteCyclePeriod";

    //角色
    string constant RoleContract = "contract";      //角色：合约
    string constant RoleAdmin = "admin";            //角色：管理
    string constant RoleDapp = "dapp";              //角色：渠道

    //合约
    string constant ContractManager = "BokerManager";
    string constant ContractLog = "BokerLog";
    string constant ContractLogData = "BokerLogData";
    string constant ContractNode = "BokerNode";
    string constant ContractNodeData = "BokerNodeData";
    string constant ContractInterface = "BokerInterface";
    string constant ContractInterfaceBase = "BokerInterfaceBase";
    string constant ContractChainware = "Chainware";
    string constant ContractCWAREToken = "CWAREToken";
    string constant ContractTokenERC20 = "TokenERC20";
    string constant ContractRecipientSuccess = "RecipientSuccess";

    enum Error {

        Ok,
        AddressIsContract,
        EventNotSupported,
        AlreadyRegistered,
        AlreadyDailyLogined,
        AlreadyBindDapp,
        AlreadyCertificated
    }

    enum UserEventType {
        Register,
        LoginDaily,
        BindDapp,
        Watch,
        Upload,
        Certification,
        End                      // end of event type, event type value should less than End
    }

    enum FinanceReason {
        Transfer,               //普通转账
        Mine,                   //挖矿所得
        AssignToken,            //分币
        Vote,                   //投票
        VoteCancel,             //取消投票
        VoteUnlock,             //投票解锁
        FinanceWithdraw,        //
        Tip                     //打赏
    }

    enum VoteType {
        Vote,
        Cancel,
        Unlock
    }

}