pragma solidity ^0.4.8;

import "./BokerCommon.sol";
import "./BokerManager.sol";

//Chainware Node Interface
interface INode {
    function registerCandidate(address addrCandidate, string description, string team, string name, uint256 tickets) external;
    function vote(address addrVoter, address addrCandidate, uint256 tokens) external;
    function cancelAllVotes(address addrVoter) external;
    function getCandidates() external view returns (address[] memory addresses, uint256[] memory tickets);
    function getCandidate(address addrCandidate) external view returns (string description, string team, string name, uint256 tickets);
}

//CWAREToken Interface
interface ICWAREToken{
    function transfer(address to, uint256 value) public;
    function transferFrom(address from, address to, uint256 value) public returns (bool success);
    function approveAndCall(address spender, uint256 value, bytes extraData) public returns (bool success);
    function mintToken(address target, uint256 mintedAmount) public;
    function lockToken(address target, bytes32 orderId, uint256 lockedAmount, string taskIp) public returns(bool success);
    function unLockToken(address target, bytes32 orderId) public returns(bool);
    function getLockToken(address target, bytes32 orderId) public view
        returns(bool exitst,
        bytes32 orderID,
        uint256 value,
        string taskIp,
        uint256 index);
    function getAccountBalanceOf(address target) public view returns(uint256);
    function getLockBoxBalanceOf(address target) public view returns(uint256);
    function getMyAccountBalanceOf() public view returns(uint256);
    function getMyLockBoxBalanceOf() public view returns(uint256);
}

//Chainware Interface
interface IChainware{
    function taskOrderRequest(uint64 index, uint256 taskGas, string codeHash, string taskIp) public
        returns (bool success,
        bytes32 orderId,
        string reason);

    function taskOrderCancel(bytes32 orderID) public returns (bool success);
    function computingNodeEnroll(uint64 hashRate) public returns (bool, address);
    function computingNodeDisconnect (address node) public returns (bool success);
    function computingTaskComplete(address node, bytes32 orderID, uint256 realGas, string resultHash) public returns (bool);
    function isActiveOrder(bytes32 orderID) public returns(bool);
    function getOrderBookCount() public returns (uint);
    function getOneTaskOrder(bytes32 orderID) public view returns (bool);
    function getAllTaskOrder() public returns (bool);
    function getAllNodeSet() public returns (bool);
    function houseHoldAllTaskOrder() public returns (bool);

    // Test;
    function getTaskOrder(bytes32 orderID) public view
        returns (bool exists,
        address client,
        uint64 index,
        uint createTime,
        uint256 taskGas,
        string taskIp,
        string codeHash);
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

    //CWAREToken Interface
    function transfer(address to, uint256 value) public{
        return ICWAREToken(contractAddress(ContractCWAREToken)).transfer(to, value);
    }

    function transferFrom(address from, address to, uint256 value) public returns (bool){
        return ICWAREToken(contractAddress(ContractCWAREToken)).transferFrom(from, to, value);
    }

    function approveAndCall(address spender, uint256 value, bytes extraData) public returns  (bool success){
        return ICWAREToken(contractAddress(ContractCWAREToken)).approveAndCall(spender, value, extraData);
    }

    function mintToken(address target, uint256 mintedAmount) public{
        return ICWAREToken(contractAddress(ContractCWAREToken)).mintToken(target, mintedAmount);
    }

    function lockToken(address target, bytes32 orderId, uint256 lockedAmount, string taskIp) public returns(bool){
        return ICWAREToken(contractAddress(ContractCWAREToken)).lockToken(target, orderId, lockedAmount, taskIp);
    }

    function unLockToken(address target, bytes32 orderId) public returns(bool){
        return ICWAREToken(contractAddress(ContractCWAREToken)).unLockToken(target, orderId);
    }

    function getLockToken(address target, bytes32 orderId) public returns(bool exitst, bytes32 orderID, uint256 value, string taskIp, uint256 index){
        return ICWAREToken(contractAddress(ContractCWAREToken)).getLockToken(target, orderId);
    }

    function getAccountBalanceOf(address target) public view returns(uint256){
        return ICWAREToken(contractAddress(ContractCWAREToken)).getAccountBalanceOf(target);
    }

    function getLockBoxBalanceOf(address target) public view returns(uint256){
        return ICWAREToken(contractAddress(ContractCWAREToken)).getLockBoxBalanceOf(target);
    }

    function getMyAccountBalanceOf() public view returns(uint256){
        return ICWAREToken(contractAddress(ContractCWAREToken)).getMyAccountBalanceOf();
    }

    function getMyLockBoxBalanceOf() public view returns(uint256){
        return ICWAREToken(contractAddress(ContractCWAREToken)).getMyLockBoxBalanceOf();
    }

    //Chainware Interface
    function taskOrderRequest(uint64 index, uint256 taskGas, string codeHash, string taskIp) public 
        returns (bool success,
                bytes32 orderId,
                string reason){
        
        return IChainware(contractAddress(ContractChainware)).taskOrderRequest(index, taskGas, codeHash, taskIp);
    }

    function taskOrderCancel(bytes32 orderID) public returns (bool){
        return IChainware(contractAddress(ContractChainware)).taskOrderCancel(orderID);
    }

    function computingNodeEnroll(uint64 hashRate) public returns (bool, address){
        return IChainware(contractAddress(ContractChainware)).computingNodeEnroll(hashRate);
    }

    function computingNodeDisconnect (address node) public returns (bool){
        return IChainware(contractAddress(ContractChainware)).computingNodeDisconnect(node);
    }

    function computingTaskComplete(address node, bytes32 orderID, uint256 realGas, string resultHash) public 
        returns (bool){

        return IChainware(contractAddress(ContractChainware)).computingTaskComplete(node, orderID, realGas, resultHash);
    }

    function isActiveOrder(bytes32 orderID) public returns  (bool){
        return IChainware(contractAddress(ContractChainware)).isActiveOrder(orderID);
    }

    function getOrderBookCount() public returns (uint){
        return IChainware(contractAddress(ContractChainware)).getOrderBookCount();
    }

    function getOneTaskOrder(bytes32 orderID) public view returns (bool){
        return IChainware(contractAddress(ContractChainware)).getOneTaskOrder(orderID);
    }

    function getAllTaskOrder() public returns (bool){
        return IChainware(contractAddress(ContractChainware)).getAllTaskOrder();
    }

    function getAllNodeSet() public returns (bool){
        return IChainware(contractAddress(ContractChainware)).getAllNodeSet();
    }

    function houseHoldAllTaskOrder() public returns (bool){
        return IChainware(contractAddress(ContractChainware)).houseHoldAllTaskOrder();
    }

    function getTaskOrder(bytes32 orderID) public 
        returns (bool exists,
                address client,
                uint64 index,
                uint createTime,
                uint256 taskGas,
                string taskIp,
                string codeHash){
        
        (exists, client, index, createTime, taskGas, taskIp, codeHash) = IChainware(contractAddress(ContractChainware)).getTaskOrder(orderID);
        return (exists, client, index, createTime, taskGas, taskIp, codeHash);
    }




}