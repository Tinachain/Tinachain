pragma solidity ^0.4.8;


import "./BokerCommon.sol";
import "./BokerManager.sol";

interface INode {

    //About Node
    function getCandidates() external view  returns(address[] memory addresses, uint256[] memory tickets);
    function getBlacks() external view returns (address[] memory addresses);
    function tickTimeout(uint256 nowTimer) external;
    function rotateVote(uint256 nowTimer) external;
}

contract BokerInterfaceBase is BokerManaged{
    constructor(address addrManager) BokerManaged(addrManager) public {

        testCandidateArray.push(address(0));
    }

    //为了方便测试添加的功能;
    uint256 isTest = 1;
    uint256 public setSize = 0;
    struct TestCandidate {
        uint256 index;
        address addr;
        uint256 tickets;
    }
    mapping (address => TestCandidate) public testCandidates;
    address[] public testCandidateArray;

    function getCandidates() external view returns(address[] memory addresses, uint256[] memory tickets) {

        if(isTest == 0){
            return INode(contractAddress(ContractNode)).getCandidates();
        }else{

            uint256 len = testCandidateArray.length;
            addresses = new address[](len - 1);
            tickets = new uint256[](len - 1);
            for(uint256 index = 1; index < len; index++) {
                address addr = testCandidateArray[index];
                addresses[index - 1] = addr;
                tickets[index - 1] = testCandidates[addr].tickets;
            }
        }
    }

    function getBlacks() external view returns (address[] memory addresses){

        return INode(contractAddress(ContractNode)).getBlacks();
    }

    function tickTimeout(uint256 nowTimer) external{

        return INode(contractAddress(ContractNode)).tickTimeout(nowTimer);
    }

    function rotateVote(uint256 nowTimer) external {

        return INode(contractAddress(ContractNode)).rotateVote(nowTimer);
    }

    function setCandidates(address nodeAddress, uint256 tickets) external returns(bool){

        TestCandidate storage testCandidate = testCandidates[nodeAddress];
        if(testCandidate.index > 0) {
            return false;
        }
        testCandidate.addr = nodeAddress;
        testCandidate.index = testCandidateArray.length;
        testCandidate.tickets = tickets;
        testCandidateArray.push(nodeAddress);

        setSize++;
        return true;
    }

}