pragma solidity ^0.4.8;

import "./BokerManager.sol";
import "./BokerNodeData.sol";
import "./BokerLog.sol";

contract BokerNode is BokerManaged {

    using SafeMath for uint256;
    using Uint256Util for uint256;

    constructor(address addrManager) BokerManaged(addrManager) public {

    }

    function registerCandidate(address addrCandidate, string description, string team, string name, uint256 tickets) external  {

        require(addrCandidate != address(0), "addrCandidate is 0");
        BokerNodeData(contractAddress(ContractNodeData)).addCandidate(addrCandidate, description, team, name, tickets);
    }

    function vote(address addrVoter, address addrCandidate, uint256 tokens) external  {

        BokerNodeData nodeData = BokerNodeData(contractAddress(ContractNodeData));
        require(nodeData.existCandidate(addrCandidate), "addrCandidate not found!");
        require(tokens > 0, "tokens <= 0!");

        nodeData.increaseVoterDeposit(addrVoter, tokens);
        nodeData.increaseVoterVote(addrVoter, addrCandidate, tokens);
        nodeData.increaseCandidateTicket(addrCandidate, tokens);

        BokerLog(contractAddress(ContractLog)).voteLogAdd(addrVoter, addrCandidate, uint256(VoteType.Vote), tokens);
    }

    function _cancelAllVotes(address addrVoter, address[] memory addresses, uint256[] memory tickets) private {

        BokerNodeData nodeData = BokerNodeData(contractAddress(ContractNodeData));
        for(uint256 index = 0; index < addresses.length; index++){
            address addrCandidate = addresses[index];
            uint256 ticket = tickets[index];
            nodeData.clearVoterVote(addrVoter, addrCandidate);
            nodeData.decreaseCandidateTicket(addrCandidate, ticket);
        }
    }

    function cancelAllVotes(address addrVoter) external  {

        BokerNodeData nodeData = BokerNodeData(contractAddress(ContractNodeData));
        (address[] memory addresses, uint256[] memory tickets, uint256 unlockTime, uint256 deposit) = nodeData.getVoteInfo(addrVoter);
        unlockTime;
        _cancelAllVotes(addrVoter, addresses, tickets);
        nodeData.setVoterDeposit(addrVoter, 0);
        BokerLog(contractAddress(ContractLog)).voteLogAdd(addrVoter, address(0), uint256(VoteType.Cancel), deposit);
    }

    function getCandidates() external view returns(address[] memory addresses, uint256[] memory tickets) {

        return BokerNodeData(contractAddress(ContractNodeData)).getCandidates();
    }

    function getBlacks() external view returns(address[] memory addresses) {

        return BokerNodeData(contractAddress(ContractNodeData)).getBlacks();
    }

    function getCandidate(address addrCandidate) external view  returns(string description, string team, string name, uint256 tickets) {

        return BokerNodeData(contractAddress(ContractNodeData)).getCandidate(addrCandidate);
    }

    function rotateVote(uint256 nowTimer) external {

        // TODO
    }
}