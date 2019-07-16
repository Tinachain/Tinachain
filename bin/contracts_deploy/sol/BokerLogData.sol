pragma solidity ^0.4.8;

import "./BokerManager.sol";

contract BokerLogData is BokerManaged {

    using PageUtil for uint256;
    constructor(address addrManager) BokerManaged(addrManager) public {
    }

    // 通用日志
    struct CommonLogEntry {
        string data;
        uint256 time;
    }
    CommonLogEntry[] commonLogs;

    function commonLogAdd(string data, uint256 time) public {
        commonLogs.push(CommonLogEntry(data, time));
    }

    function commonLogGet(int256 idx) public view returns (string data, uint256 time) {

        if(commonLogs.length <= 0) {
            return;
        }

        uint256 uidx = 0;
        if(idx < 0) {
            uidx = commonLogs.length - 1;
        }else{
            uidx = uint256(idx);
        }

        if(uidx >= commonLogs.length) {
            uidx = commonLogs.length - 1;
        }

        data = commonLogs[uidx].data;
        time = commonLogs[uidx].time;
    }


    //投票日志
    struct VoteLogEntry {
        uint256   voteType;
        uint256 tokens;
        address addrCandidate;
        uint256 time;
    }

    struct VoteLog {
        address addrUser;
        VoteLogEntry[] entries;
    }

    mapping (address=>VoteLog) voteLogs;

    function voteLogAdd(address addrUser, address addrCandidate, uint256 voteType, uint256 tokens) public {

        VoteLog storage voteLog = voteLogs[addrUser];
        voteLog.addrUser = addrUser;
        voteLog.entries.push(VoteLogEntry(uint256(voteType), tokens, addrCandidate, now));
    }

    function voteLogGet(address addrUser, uint256 page, uint256 pageSize) public view returns (

        uint256[] voteTypes, uint256[] tokenses, address[] addrCandidates, uint256[] times) {
        VoteLog storage voteLog = voteLogs[addrUser];
        if(voteLog.addrUser == address(0)) {
            return;
        }

        (uint256 start, uint256 end) = voteLog.entries.length.pageRange(page, pageSize);
        uint256 len = end - start + 1;
        voteTypes = new uint256[](len);
        tokenses = new uint256[](len);
        addrCandidates = new address[](len);
        times = new uint256[](len);

        for (uint256 index = start; index <= end; index++) {

            VoteLogEntry storage entry = voteLog.entries[index];
            voteTypes[index-start] = entry.voteType;
            tokenses[index-start] = entry.tokens;
            addrCandidates[index-start] = entry.addrCandidate;
            times[index-start] = entry.time;
        }
    }

    //vote rotate 日志
    struct VoteRotateLog {
        uint256 round;
        uint256 time;
    }

    VoteRotateLog[] voteRotateLogs;

    function voteRotateLogAdd(uint256 round) public {

        voteRotateLogs.push(VoteRotateLog(round, now));
    }

    function voteRotateLogGetIndex(uint256 index) public view returns (uint256 round, uint256 time) {

        require(index < voteRotateLogs.length, "index exceeds voteRotateLogs length");
        VoteRotateLog storage voteRotateLog = voteRotateLogs[index];
        round = voteRotateLog.round;
        time = voteRotateLog.time;
    }

    function voteRotateLogGet(uint256 page, uint256 pageSize) public view returns (uint256[] rounds, uint256[] times) {

        (uint256 start, uint256 end) = voteRotateLogs.length.pageRange(page, pageSize);
        uint256 len = end - start + 1;
        rounds = new uint256[](len);
        times = new uint256[](len);

        for (uint256 index = start; index <= end; index++) {

            VoteRotateLog storage voteRotateLog = voteRotateLogs[index];
            rounds[index-start] = voteRotateLog.round;
            times[index-start] = voteRotateLog.time;
        }
    }
}