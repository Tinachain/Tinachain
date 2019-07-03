pragma solidity ^0.4.8;

import "./BokerCommon.sol";
import "./BokerManager.sol";

interface INode {
    function registerCandidate(address addrCandidate, string description, string team, string name) external;
    function vote(address addrVoter, address addrCandidate, uint256 tokens) external;
    function cancelAllVotes(address addrVoter) external;
    function getCandidates() external view returns (address[] memory addresses, uint256[] memory tickets);
    function getCandidate(address addrCandidate) external view returns (string description, string team, string name, uint256 tickets);
}

interface IUser {
    function getInvitedFriendsCount(address addrUser) external view returns (uint256);
    function userCount() external view returns (uint256 total);
    function giveTipTo(address addrTo) external payable;
}

interface Ilog {

    function userPowerLogGet(address addrUser, uint256 page, uint256 pageSize) external view returns (
        address[] addrDapps, uint256[] powerTypes, int256[] powerChanges, uint256[] reasons, uint256[] param1s, uint256[] times);

    function userFinanceLogGet(address addrUser, uint256 page, uint256 pageSize) external view returns (
        int256[] tokensChanges, uint256[] tokensAfters, uint256[] reasons, uint256[] times, uint256[] indexes);

    function userTipLogGet(address addrUser, uint256 page, uint256 pageSize) external view returns (
        address[] addrUsers, int256[] tokensChanges, uint256[] tokensAfters, uint256[] times);
}

contract BokerInterface is BokerManaged{
    // Ensure that we are pointing to the right contract in our set call.
    bool public isBokerInterface = true;
    constructor(address addrManager) BokerManaged(addrManager) public {
    }

    /** @dev Register to be candidate of verifier.
    * @param description description of node
    * @param team description of team
    * @param name name of node
    */
    function registerCandidate(string description, string team, string name) external whenNotPaused {
        INode(contractAddress(ContractNode)).registerCandidate(msg.sender, description, team, name);
    }

    /** @dev Vote for candidate.
    * @param addrCandidate Address of candidate.
    */
    function voteCandidate(address addrCandidate) external payable whenNotPaused {
        INode(contractAddress(ContractNode)).vote(msg.sender, addrCandidate, msg.value);
    }

    /** @dev Cancel all votes of voter.
    */
    function cancelAllVotes() external whenNotPaused {
        INode(contractAddress(ContractNode)).cancelAllVotes(msg.sender);
    }

    /** @dev Get all candidates with tickets.
    * @return addresses The addresses of candidates.
    * @return tickets The tickets of candidates.
    */
    function getCandidates() external view returns(address[] memory addresses, uint256[] memory tickets) {
        return INode(contractAddress(ContractNode)).getCandidates();
    }

    /** @dev Get candidate.
    * @param addrCandidate Address of candidate.
    * @return description description of node
    * @return team description of team
    * @return name name of node
    * @return tickets tickets of node
    */
    function getCandidate(address addrCandidate) external view  returns(string description, string team, string name, uint256 tickets) {
        return INode(contractAddress(ContractNode)).getCandidate(addrCandidate);
    }
}