pragma solidity ^0.4.8;

import "./BokerCommon.sol";
import "./BokerManager.sol";

//Tinachain Node Interface
interface INode {
    // Transaction;
    function registerCandidate(address addrCandidate, string description, string team, string name, uint256 tickets) external;
    function vote(address addrVoter, address addrCandidate, uint256 tokens) external;
    function cancelAllVotes(address addrVoter) external;

    // View
    function getCandidates() external view returns (address[] memory addresses, uint256[] memory tickets);
    function getCandidate(address addrCandidate) external view returns (string description, string team, string name, uint256 tickets);
}

contract BokerInterface is BokerManaged{

    bool public isBokerInterface = true;
    constructor(address addrManager) BokerManaged(addrManager) public {
    }

    //INode Interface
    function registerCandidate(string description, string team, string name, uint256 tickets) external whenNotPaused {
        INode(contractAddress(ContractNode)).registerCandidate(msg.sender, description, team, name, tickets);
    }

    function voteCandidate(address addrCandidate) external payable whenNotPaused {
        INode(contractAddress(ContractNode)).vote(msg.sender, addrCandidate, msg.value);
    }

    function cancelAllVotes() external whenNotPaused {
        INode(contractAddress(ContractNode)).cancelAllVotes(msg.sender);
    }

    function getCandidates() external view returns(address[] memory addresses, uint256[] memory tickets) {
        return INode(contractAddress(ContractNode)).getCandidates();
    }

    function getCandidate(address addrCandidate) external view  returns(string description, string team, string name, uint256 tickets) {
        return INode(contractAddress(ContractNode)).getCandidate(addrCandidate);
    }
}