pragma solidity ^0.4.8;

import "./BokerManager.sol";
import "./BokerLogData.sol";

contract BokerLog is BokerManaged {

    constructor(address addrManager) BokerManaged(addrManager) public payable {
    }

    function commonLogAdd(string data, uint256 time) external {

        BokerLogData(contractAddress(ContractLogData)).commonLogAdd(data, time);
    }

    function voteLogAdd(address addrUser, address addrCandidate, uint256 voteType, uint256 tokens) public payable {

        BokerLogData(contractAddress(ContractLogData)).voteLogAdd(addrUser, addrCandidate, voteType, tokens);
    }

    function voteLogGet(address addrUser, uint256 page, uint256 pageSize) public view returns (

        uint256[] voteTypes, uint256[] tokenses, address[] addrCandidates, uint256[] times) {
        return BokerLogData(contractAddress(ContractLogData)).voteLogGet(addrUser, page, pageSize);
    }

    function voteRotateLogAdd(uint256 round) public payable {

        BokerLogData(contractAddress(ContractLogData)).voteRotateLogAdd(round);
    }

    function voteRotateLogGetIndex(uint256 index) public view returns (uint256 round, uint256 time) {

        return BokerLogData(contractAddress(ContractLogData)).voteRotateLogGetIndex(index);
    }

    function voteRotateLogGet(uint256 page, uint256 pageSize) public view returns (uint256[] rounds, uint256[] times) {

        return BokerLogData(contractAddress(ContractLogData)).voteRotateLogGet(page, pageSize);
    }
}

