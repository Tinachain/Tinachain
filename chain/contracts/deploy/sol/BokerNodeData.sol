pragma solidity ^0.4.8;

import "./BokerManager.sol";

contract BokerNodeData is BokerManaged {

    using SafeMath for uint256;

    constructor(address addrManager) BokerManaged(addrManager) public {
        candidateArray.push(address(0));
        voterArray.push(address(0));
    }

    //candidates
    struct Candidate {
        address addr;           // address of candidate
        uint256 index;          // index of candidate
        uint256 tickets;        // tickets
        string description;     // description of node
        string team;            // description of team
        string name;            // name of node
    }

    mapping (address => Candidate) public candidates;
    address[] public candidateArray;
    function addCandidate(address addrCandidate, string description, string team, string name, uint256 tickets) external  {

        Candidate storage candidate = candidates[addrCandidate];
        if(candidate.index > 0) {
            return;
        }
        candidate.addr = addrCandidate;
        candidate.index = candidateArray.length;
        candidate.tickets = tickets;
        candidate.description = description;
        candidate.team = team;
        candidate.name = name;
        candidateArray.push(addrCandidate);
    }

    function getCandidates() external view returns(address[] memory addresses, uint256[] memory tickets) {

        uint256 len = candidateArray.length;
        addresses = new address[](len - 1);
        tickets = new uint256[](len - 1);

        for(uint256 index = 1; index < len; index++) {
            address addr = candidateArray[index];
            addresses[index - 1] = addr;
            tickets[index - 1] = candidates[addr].tickets;
        }
    }

    function getCandidate(address addrCandidate) external view  returns(string description, string team, string name, uint256 tickets) {

        Candidate storage candidate = candidates[addrCandidate];
        if(0 == candidate.index){
            return;
        }

        description = candidate.description;
        team = candidate.team;
        name = candidate.name;
        tickets = candidate.tickets;
        return;
    }

    function existCandidate(address addrCandidate) external view returns (bool exist) {

        if(0 != candidates[addrCandidate].index){
            return true;
        }
        return false;
    }

    function increaseCandidateTicket(address addrCandidate, uint256 tokens) external  {

        Candidate storage candidate = candidates[addrCandidate];
        if(0 == candidate.index){
            return;
        }
        uint256 ticket = tokens;
        candidate.tickets = candidate.tickets.add(ticket);
    }

    function decreaseCandidateTicket(address addrCandidate, uint256 tokens)  external{

        Candidate storage candidate = candidates[addrCandidate];
        if(0 == candidate.index){
            return;
        }
        uint256 ticket = tokens;
        require(candidate.tickets >= ticket, "candidate.tickets < amount!");
        candidate.tickets = candidate.tickets.sub(ticket);
    }

    struct Vote {
        uint256 index;                          // index of voter
        uint256 tickets;                        // ammount of tokens voted
    }

    struct Voter {
        address addr;                           // address of voter
        uint256 index;                          // index of voter
        uint256 deposit;                        // deposit tokens
        uint256 unlockTime;                     // time to unlock
        mapping (address=>Vote) votes;          // mapping of vote info, candidate address => Vote
        address[] candidateArray;               // array of candidate address .
    }

    mapping(address => Voter) public voters;
    address[] public voterArray;

    function _findAddVote(Voter storage voter, address addrCandidate) private returns (Vote storage) {

        Vote storage vote = voter.votes[addrCandidate];
        vote.tickets = 0;
        vote.index = voter.candidateArray.length;
        voter.candidateArray.push(addrCandidate);
        return vote;
    }

    function increaseVoterDeposit(address addrVoter, uint256 tokens) external  {

        Voter storage voter = voters[addrVoter];
        if(0 == voter.index){
            voter.addr = addrVoter;
            voter.deposit = 0;
            voter.index = voterArray.length;
            voterArray.push(addrVoter);
            voters[addrVoter].candidateArray.push(address(0));
        }
        voter.deposit = voter.deposit.add(tokens);
    }

    function increaseVoterVote(address addrVoter, address addrCandidate, uint256 tokens) external  {

        Voter storage voter = voters[addrVoter];
        if(0 == voter.index){
            return;
        }
        Vote storage vote = _findAddVote(voter, addrCandidate);
        vote.tickets = vote.tickets.add(tokens);
    }

    function clearVoterVote(address addrVoter, address addrCandidate) external  {

        Voter storage voter = voters[addrVoter];
        if(0 == voter.index){
            return;
        }
        Vote storage vote = voter.votes[addrCandidate];
        vote.tickets = 0;
    }

    function getVoteInfo(address addrVoter) external view returns(address[] addresses, uint256[] tickets, uint256 unlockTime, uint256 deposit) {

        Voter storage voter = voters[addrVoter];
        if(0 == voter.index){
            return;
        }

        unlockTime = voter.unlockTime;
        deposit = voter.deposit;
        uint256 len = voter.candidateArray.length;
        addresses = new address[](len - 1);
        tickets = new uint256[](len - 1);

        for(uint256 index = 1; index < len; index++){
            address addrCandidate = voter.candidateArray[index];
            addresses[index-1] = addrCandidate;
            tickets[index-1] = voter.votes[addrCandidate].tickets;
        }
    }

    function getVoters() external view returns(address[] addresses) {

        uint256 len = 0;

        for(uint index = 1; index < voterArray.length; index++) {
            address addrVoter = voterArray[index];
            if(voters[addrVoter].deposit > 0){
                len++;
            }
        }

        addresses = new address[](len);
        uint256 pos = 0;
        for(index = 1; index < voterArray.length; index++) {

            addrVoter = voterArray[index];
            if(voters[addrVoter].deposit > 0){
                addresses[pos] = addrVoter;
                pos++;
            }
        }
    }

    function setVoterDeposit(address addrVoter, uint256 amount) external  {

        Voter storage voter = voters[addrVoter];
        if(0 == voter.index){
            return;
        }
        voter.deposit = amount;
    }

    //blacklists
    mapping (address => bool) public blackMap;
    address[] blackArray;

    function clearBlacks() external onlyOwner {

        for (uint256 index = 0; index < blackArray.length; index++) {
            delete blackMap[blackArray[index]];
        }
        delete blackArray;
    }

    function insertBlacks(address addresses) external onlyOwner{

        if (!blackMap[addresses]) {
            blackMap[addresses] = true;
            blackArray.push(addresses);
        }
    }

    function deleteBlacks(address addresses) external onlyOwner{

        for (uint256 index = 0; index < blackArray.length; index++) {
            if (blackArray[index] == addresses){
                delete blackMap[blackArray[index]];
                delete blackArray[index];
            }
        }
    }

    function getBlacks() external view returns (address[] addresses){

        uint256 len = 0;
        for (uint256 i = 0; i < blackArray.length; i++) {
            if (blackMap[blackArray[i]]){
                len++;
            }
        }

        addresses = new address[](len);
        uint256 postion = 0;
        for (uint256 index = 0; index < blackArray.length; index++) {
            address tmp_addr = blackArray[index];
            if (blackMap[tmp_addr]){
                addresses[postion] = tmp_addr;
                postion++;
            }
        }
    }

    function tickTimeout() external{
        
        //Todo
    }
}