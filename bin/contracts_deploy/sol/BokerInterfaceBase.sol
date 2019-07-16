pragma solidity ^0.4.8;


import "./BokerCommon.sol";
import "./BokerManager.sol";

interface INode {

    //About Node
    function getVoteRound() external view  returns(uint256 round);
    function checkVote() external view returns (bool);
    function rotateVote() external;
    function getCandidates() external view  returns(address[] memory addresses, uint256[] memory tickets);
    function getBlacks() external view returns (address[] memory addresses);
    function tickTimeout() external;
}


contract BokerInterfaceBase is BokerManaged{

    constructor(address addrManager) BokerManaged(addrManager) public {
    }

    function () public payable {

    }

    function getVoteRound() external view  returns(uint256 round) {

        return INode(contractAddress(ContractNode)).getVoteRound();
    }

    function getCandidates() external view returns(address[] memory addresses, uint256[] memory tickets) {

        return INode(contractAddress(ContractNode)).getCandidates();
    }


    function tickVote() external view returns (bool) {

        return INode(contractAddress(ContractNode)).checkVote();
    }

    function rotateVote() external whenNotPaused  {

        INode(contractAddress(ContractNode)).rotateVote();
    }

    function getBlacks() external view returns (address[] memory addresses){

        return INode(contractAddress(ContractNode)).getBlacks();
    }

    function tickTimeout() external{

        return INode(contractAddress(ContractNode)).tickTimeout();
    }

}