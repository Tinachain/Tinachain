pragma solidity ^0.4.8;

contract BokerDefine {

    uint256 constant internal CfgTickTimer = 2 * 60 * 60;   //超时处理时间;

    //配置
    string constant CfgVoteLockup = "VoteLockup";
    string constant CfgVoteUnlockPrecision = "VoteUnlockPrecision";
    string constant CfgVoteCyclePeriod = "VoteCyclePeriod";

    //合约
    string constant ContractManager     = "BokerManager";
    string constant ContractLog         = "BokerLog";
    string constant ContractLogData     = "BokerLogData";
    string constant ContractNode        = "BokerNode";
    string constant ContractNodeData    = "BokerNodeData";
    string constant ContractInterface   = "BokerInterface";
    string constant ContractInterfaceBase = "BokerInterfaceBase";
    
    enum Error {
        Ok,
        AddressIsContract,
        AlreadyCertificated
    }

    enum VoteType {
        Vote,
        Cancel,
        Unlock
    }

}