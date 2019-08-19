pragma solidity ^0.4.8;

import "./BokerManager.sol";
import "./BokerNodeData.sol";
import "./BokerLog.sol";

// Chainware ERC20
contract TokenERC20 {

    string public name;
    string public symbol;
    uint8 public decimals;
    uint256 public totalSupply;

    mapping (address => uint256) balanceOf;
    mapping (address => mapping (address => uint256)) allowance;

    event Transfer(address indexed from, address indexed to, uint256 value);
    event Burn(address indexed from, uint256 value);
    event Approval(address indexed _owner, address indexed _spender, uint256 _value);

    constructor(
        address initialOwner,
        uint256 initialSupply,
        string tokenName,
        string tokenSymbol,
        uint8  tokenDecimals) public {
        totalSupply = initialSupply * 10 ** uint256(decimals);  // Update total supply with the decimal amount
        balanceOf[initialOwner] = totalSupply;                  // Give the creator all initial tokens
        name = tokenName;                                       // Set the name for display purposes
        symbol = tokenSymbol;                                   // Set the symbol for display purposes
        decimals = tokenDecimals;                               // normally 18
    }

    function _transfer(address _from, address _to, uint _value) internal {

        require(_to != address(0), "to address is 0!");
        require(balanceOf[_from] >= _value, "from balance too low!");
        require(balanceOf[_to] + _value > balanceOf[_to], "to balance overflows!");

        uint previousBalances = balanceOf[_from] + balanceOf[_to];
        balanceOf[_from] -= _value;
        balanceOf[_to] += _value;
        emit Transfer(_from, _to, _value);

        assert(balanceOf[_from] + balanceOf[_to] == previousBalances);
    }

    function transfer(address _to, uint256 _value) public {
        _transfer(tx.origin, _to, _value);
    }

    function transferFrom(address _from, address _to, uint256 _value) public
        returns (bool) {

        require(_value <= allowance[_from][tx.origin], "allowance too low!");
        allowance[_from][tx.origin] -= _value;
        _transfer(_from, _to, _value);
        return true;
    }

    function approve(address _spender, uint256 _value) public
        returns (bool success) {

        allowance[tx.origin][_spender] = _value;
        emit Approval(tx.origin, _spender, _value);
        return true;
    }

    function burn(uint256 _value) public
        returns (bool success) {

        require(balanceOf[tx.origin] >= _value, "tx.origin balance too low!");

        balanceOf[tx.origin] -= _value;            // Subtract from the sender
        totalSupply -= _value;                      // Updates totalSupply
        emit Burn(tx.origin, _value);
        return true;
    }

    function burnFrom(address _from, uint256 _value) public
        returns (bool success) {

        require(balanceOf[_from] >= _value, "from balance too low!");             // Check if the targeted balance is enough
        require(_value <= allowance[_from][tx.origin], "allowance too low!");    // Check allowance
        balanceOf[_from] -= _value;                         // Subtract from the targeted balance
        allowance[_from][tx.origin] -= _value;             // Subtract from the sender's allowance
        totalSupply -= _value;                              // Update totalSupply
        emit Burn(_from, _value);
        return true;
    }
}

contract CWAREToken is BokerManaged, TokenERC20 {

    struct Order {
        bytes32 orderId;
        uint256 index;
        uint256 value;
        string taskIp;
        bool use;
    }

    struct Orders {
        uint256 totalValue;                     // total Token
        bytes32[] orderArray;                   // order array
        mapping (bytes32 => Order) orderMap;    // bytes32 => orderId
    }

    string public name;                         //fancy name: eg ChainWare Token
    uint8 public decimals;                      //How many decimals to show
    string public symbol;                       //An identifier: eg SBX
    string public version = 'CWARE0.1';         //human 0.1 standard. Just an arbitrary versioning scheme.

    mapping (address =>Orders) lockBox;

    event FrozenFunds(address _target, bool _frozen);
    event QueryFunds(address _target, uint256 value, uint256 _locked, bool _frozen);

    function CWAREToken(
        address _initialOwner,
        uint256 _initialAmount,
        string _tokenName,
        uint8 _decimalUnits,
        string _tokenSymbol,
        address addrManager
        ) BokerManaged(addrManager)
        TokenERC20(_initialOwner, _initialAmount, _tokenName, _tokenSymbol, _decimalUnits) public {
        balanceOf[_initialOwner] = _initialAmount;              // Give the creator all initial tokens
        totalSupply = _initialAmount;                           // Update total supply
        name = _tokenName;                                      // Set the name for display purposes
        decimals = _decimalUnits;                               // Amount of decimals for display purposes
        symbol = _tokenSymbol;                                  // Set the symbol for display purposes
    }

    /* Approves and then calls the receiving contract */
    function approveAndCall(address _spender, uint256 _value, bytes _extraData) public
        returns (bool success) {
        
        allowance[tx.origin][_spender] = _value;
        emit Approval(tx.origin, _spender, _value);

        //call the receiveApproval function on the contract you want to be notified. This crafts the function signature manually so one doesn't have to include a contract in here just for this.
        //receiveApproval(address _from, uint256 _value, address _tokenContract, bytes _extraData)
        //it is assumed when one does this that the call *should* succeed, otherwise one would use vanilla approve instead.
        require(_spender.call(bytes4(bytes32(keccak256("receiveApproval(address,uint256,address,bytes)"))), tx.origin, _value, this, _extraData), "");
        return true;
    }

    function mintToken(address target, uint256 mintedAmount) onlyOwner public {
        balanceOf[target] += mintedAmount;
        totalSupply += mintedAmount;
        emit Transfer(0, this, mintedAmount);
        emit Transfer(this, target, mintedAmount);
    }

    function _transferFrom(address _from, address _to, uint _value) public
        returns (bool)  {

        if (!((_to != 0x0) &&
             (balanceOf[_from] > _value)  &&
             (balanceOf[_to] + _value > balanceOf[_to]))) {
            return false;
        }
        balanceOf[_from] -= _value;
        balanceOf[_to] += _value;
        emit Transfer(_from, _to, _value);
        return true;
    }
    
    function lockToken(address target, bytes32 orderId, uint256 lockedAmount, string taskIp) public
        returns(bool) {

        //1. enough funds 2.lockbox empty 3.account is not frozen
        if (balanceOf[target] < lockedAmount) {
            return false;
        }

        //check orderId
        if(lockBox[target].orderMap[orderId].use){
            return false;
        }

        //write orderId
        lockBox[target].orderMap[orderId].orderId = orderId;
        lockBox[target].orderMap[orderId].value = lockedAmount;
        lockBox[target].orderMap[orderId].taskIp = taskIp;
        lockBox[target].orderMap[orderId].index = lockBox[target].orderArray.length;
        lockBox[target].orderMap[orderId].use = true;

        lockBox[target].totalValue += lockedAmount;
        lockBox[target].orderArray.push(orderId);

        balanceOf[target] -= lockedAmount;
        return (true);
    }

    function unLockToken(address target, bytes32 orderId) public
        returns(bool) {

        //check orderId
        if(!lockBox[target].orderMap[orderId].use){
            return false;
        }

        // unlock token;
        uint256 lockedAmount = lockBox[target].orderMap[orderId].value;
        uint256 index = lockBox[target].orderMap[orderId].index;
        balanceOf[target] += lockedAmount;

        lockBox[target].totalValue -= lockedAmount;
        delete lockBox[target].orderMap[orderId];
        delete lockBox[target].orderArray[index];
        return (true);
    }

    function getLockToken(address target, bytes32 orderId) public view
        returns(bool exitst, bytes32 orderID, uint256 value, string taskIp, uint256 index){

        //check orderId
        if(!lockBox[target].orderMap[orderId].use){
            return (false, orderId, 0, "0.0.0.0", 0);
        }

        return (true, lockBox[target].orderMap[orderId].orderId,
                        lockBox[target].orderMap[orderId].value,
                        lockBox[target].orderMap[orderId].taskIp,
                        lockBox[target].orderMap[orderId].index);
    }

    function getAccountBalanceOf(address target) public view onlyOwner
        returns(uint256) {
        return (balanceOf[target]);
    }

    function getLockBoxBalanceOf(address target) public view onlyOwner
        returns(uint256) {
        return (lockBox[target].totalValue);
    }

    function getMyAccountBalanceOf() public view
        returns(uint256) {
        return (balanceOf[tx.origin]);
    }

    function getMyLockBoxBalanceOf() public view
            returns(uint256) {
        return (lockBox[tx.origin].totalValue);
    }
}

contract Chainware is BokerManaged{

    uint32 constant  MAX_WAITING_TIME = 300;        //max wait for 300s
    uint32 constant  PREPARE_TIME = 10;             //computing node estimated prepare time
    uint32 constant  MAX_SEARCH_ASSIGNED_NODE = 10; //max number of search compute nodes

    address public CWARE_Token_Addr;
    uint public numTaskOrder;
    uint public numComputeNode;

    using SafeMath for uint256;

    enum taskState {Unrequested, Confirmed, Assigned, Completed, Canceled}

    struct taskInfo {
        string  codeHash;
        string  dataHash;
        string  taskIp;
    }

    struct taskOrder {
        address client;
        address computeNode;
        taskState state;
        uint32  _state;
        uint64  index;
        uint    createTime;
        uint    assignTime;
        uint    completeTime;
        uint    listPtr;
        uint256    taskGas;
        bytes32 orderID;
        mapping (uint256 => taskInfo) info;
    }

    mapping (bytes32 => taskOrder) public orderBook;
    bytes32[] public taskOrderList;

    struct nodeSet {
        address nodeAddr;
        uint256    hashRate;
        uint    lastHashRate;
        uint    assignCnt;
        bytes32  assignedClient;
        uint    enrollTime;
        bool    enrolled;
        bytes32 nodeID;
        address nextNode;
        address prevNode;
        string  nodeIp;
    }
    mapping (address => nodeSet) public nodeSets;
    address firstNodeSet;
    address lastNodeSet;

    //constructor
    function Chainware(address addrManager, address cware_token) BokerManaged(addrManager) public {
        numTaskOrder = 0;
        numComputeNode = 0;
        CWARE_Token_Addr = cware_token;
        firstNodeSet = 0;
        lastNodeSet = 0;
    }

    //add node to the list, ether first or last
    function addNodetoList(nodeSet storage ns, bool first)

        internal returns(bool) {
        if (ns.enrolled != true) {
            return (false);
        }

        //first
        if (first == true) {

            if (firstNodeSet == 0) {
                lastNodeSet = ns.nodeAddr;
            } else {
                nodeSet storage n = nodeSets[firstNodeSet];
                n.prevNode = ns.nodeAddr;
            }

            ns.nextNode = firstNodeSet;
            ns.prevNode = 0;

            firstNodeSet = ns.nodeAddr;
        }

        // insert to the last
        else {

            ns.nextNode = 0;
            if (firstNodeSet == 0) {
                ns.prevNode = 0;
                firstNodeSet = ns.nodeAddr;
            }
            else {
               n = nodeSets[lastNodeSet];
               n.nextNode = ns.nodeAddr;
               ns.prevNode = lastNodeSet;
            }

            lastNodeSet = ns.nodeAddr;
        }

        return (true);
    }

    function deleteNodeFromList(nodeSet storage ns)
        internal returns(bool) {
        if (ns.enrolled != true) {
            return (false);
        }

        address next = ns.nextNode;
        address prev = ns.prevNode;
        if (prev == 0) {
            firstNodeSet = next;
            nodeSets[next].prevNode = 0;
        }
        else {
            nodeSets[prev].nextNode = next;
        }

        if (next == 0) {
            lastNodeSet = prev;
            nodeSets[prev].nextNode = 0;
        }
        else {
            nodeSets[next].prevNode = prev;
        }

        return (true);
    }

    //define events
    event ClientEvent(bytes32 _id,
        address indexed _client,
        uint64 _indexID,
        address indexed _computeNode,
        uint _taskGas,
        uint _createTime,
        uint32 _state,
        string _reason);

    event NodeEvent(bytes32 _id,
        address indexed _node,
        uint _hashRate,
        uint _enrollTime,
        string _reason);

    event AssignEvent(bytes32 _id,
        address indexed _client,
        uint64 _index,
        address indexed _computeNode,
        uint _taskGas,
        uint _assignTime,
        string _codeHash);

    event CompleteEvent(bytes32 _id,
        address _client,
        uint64 _index,
        address _computeNode,
        uint _completeTime,
        string _resultHash);

    event LogEvent(bytes32 _id,
        address indexed _client,
        address indexed _computeNode,
        string _reason);

    function taskOrderRequest(uint64 index, uint256 taskGas, string codeHash, string taskIp)
        public returns (bool success, bytes32 orderId, string reason) {

        uint nowTime = now;
        bytes32 orderID = keccak256(tx.origin, taskGas, taskState.Confirmed, nowTime);
        taskOrder storage o = orderBook[orderID];
        
        if ((o.state != taskState.Unrequested) && (o.client == tx.origin)) {
            
            emit ClientEvent(o.orderID,
                        o.client,
                        o.index,
                        o.computeNode,
                        o.taskGas,
                        o.createTime,
                        o._state,
                        "task already existed, cancel first");
            success = false;
            orderId = orderID;
            reason = "task already existed, cancel first";
            return (success, orderId, reason);
        }

        if (CWAREToken(CWARE_Token_Addr).lockToken(tx.origin, orderID, taskGas, taskIp) != true) {

            emit ClientEvent(orderID, tx.origin, index, 0, taskGas, 0, 0, "Lock token failed");

            success = false;
            orderId = orderID;
            reason = "Lock token failed";
            return (success, orderId, reason);
        }

        o.client = tx.origin;
        o.index = index;
        o.taskGas = taskGas;
        o.info[0].codeHash = codeHash;
        o.state = taskState.Confirmed;
        o._state = 1;
        o.createTime = nowTime;
        o.info[0].taskIp = taskIp;
        o.orderID = keccak256(o.client, o.taskGas, o.state, nowTime);
        o.listPtr = taskOrderList.push(o.orderID) - 1;
        numTaskOrder ++;

        //then assign to compute node
        bool assigned;
        address nodeAddr;
        (assigned, nodeAddr) = assignTaskOrder(o);
        if (assigned == true) {
            o.state = taskState.Assigned;
            o._state = 2;
            o.computeNode = nodeAddr;
            o.assignTime = nowTime;
        }

        emit ClientEvent(o.orderID, o.client, o.index, o.computeNode, o.taskGas, o.createTime, o._state, "order request confirmed");

        success = true;
        orderId = orderID;
        reason = "order request confirmed";
        return (success, orderId, reason);
    }

    function taskOrderCancel(bytes32 orderID) public
        returns (bool) {

        taskOrder storage o = orderBook[orderID];
        if (o.state == taskState.Unrequested) {
            return (false);
        }
        if (o.client != tx.origin) {
            return (false);
        }
        if (o.state >= taskState.Assigned) {
            return (false);
        }

        o.state = taskState.Canceled;
        o._state = 4;
        //refund the taskGas

        CWAREToken(CWARE_Token_Addr).unLockToken(o.client, orderID);

        if (o.client != address(0)) {
            emit ClientEvent(o.orderID, o.client, o.index, o.computeNode, o.taskGas, o.createTime, o._state, "order canceled");
        }

        uint rowToDelete = orderBook[o.orderID].listPtr;
        bytes32 keyToMove = taskOrderList[taskOrderList.length -1];
        taskOrderList[rowToDelete] = keyToMove;
        orderBook[keyToMove].listPtr = rowToDelete;
        taskOrderList.length --;

        o.client = 0;
        o.state = taskState.Unrequested;
        o._state = 0;
        o.index = 0;
        o.info[0].taskIp = "";
        numTaskOrder --;
        return true;
    }

    // Test Find;
    function getTaskOrder(bytes32 orderID) public view
        returns (bool, address, uint64, uint, uint256, string, string) {

        taskOrder storage o = orderBook[orderID];
        if (o.state == taskState.Unrequested) {
            return (false, address(0), 0, 0, 0, "0.0.0.0", "");
        }

        return (true, o.client, o.index, o.createTime, o.taskGas, o.info[0].taskIp, o.info[0].codeHash);
    }

    function computingNodeEnroll(uint64 hashRate) public
        returns (bool, address) {

        nodeSet storage n = nodeSets[tx.origin];
        uint nowTime = now;

        if (n.enrolled == true) {
            emit NodeEvent(n.nodeID, n.nodeAddr, n.hashRate, n.enrollTime, "node already enrolled");
            return (false, address(0x0));
        }

        n.nodeAddr = tx.origin;
        n.hashRate = hashRate;
        n.enrollTime = nowTime;
        n.enrolled = true;
        n.nodeID = keccak256(n.nodeAddr, hashRate, nowTime);

        addNodetoList(n, true);
        emit NodeEvent(n.nodeID, n.nodeAddr, n.hashRate, n.enrollTime, "node enrolled OK");

        numComputeNode ++;
        return (true, n.nodeAddr);
    }

    function computingNodeDisconnect (address node) public
        returns (bool) {

        nodeSet storage n = nodeSets[node];
        if (n.enrolled != true) {
            emit NodeEvent(n.nodeID, n.nodeAddr, n.hashRate, n.enrollTime, "node is not enrolled");
            return false;
        }

        bytes32 orderId = n.assignedClient;
        if (orderId != 0) {
            taskOrder storage o = orderBook[orderId];
            o.state = taskState.Confirmed;
            o._state = 1;

            bool assigned;
            address newNodeAddr;
            (assigned, newNodeAddr) = assignTaskOrder(o);
            if (assigned == true) {
                o.state = taskState.Assigned;
                o._state = 2;
                o.computeNode = newNodeAddr;
                o.assignTime = now;
            }
        }

        deleteNodeFromList(n);
        emit NodeEvent(n.nodeID, n.nodeAddr, n.hashRate, n.enrollTime, "node removed");
        numComputeNode --;

        n.nodeAddr = 0;
        n.enrolled = false;
        //n.nodeID = 0;

        return true;
    }

    function computingTaskComplete(address node, bytes32 orderID, uint256 realGas, string resultHash) public
        returns (bool) {

        uint nowTime = now;
        taskOrder storage o = orderBook[orderID];
        if ((o.state == taskState.Unrequested) || (o.orderID != orderID)) {
            return false;
        }

        o.state = taskState.Completed;
        o._state = 3;
        emit ClientEvent(o.orderID, o.client, o.index, o.computeNode, o.taskGas, o.createTime, o._state, "task completed");

        CWAREToken(CWARE_Token_Addr).unLockToken(o.client, orderID);

        //if realGas is bigger than estimated gas, use estimated gas
        //otherwise use the real one
        uint256 taskGas = (realGas >= o.taskGas) ? o.taskGas : realGas;
        var payMinerAmount = taskGas.div(5);
        uint payNodeAmount = taskGas - payMinerAmount;
        CWAREToken(CWARE_Token_Addr)._transferFrom(o.client, node, payNodeAmount);

        //miner's rewards fee
        CWAREToken(CWARE_Token_Addr)._transferFrom(o.client, block.coinbase, payMinerAmount);

        //remove this task order out of order book
        uint rowToDelete = orderBook[o.orderID].listPtr;
        bytes32 keyToMove = taskOrderList[taskOrderList.length - 1];
        taskOrderList[rowToDelete] = keyToMove;
        orderBook[keyToMove].listPtr = rowToDelete;
        taskOrderList.length --;
      
/******
        o.client = 0;
        o.orderID = 0;
        o.index = 0;
        o.state = taskState.Unrequested;
        o._state = 0;
*****/
        numTaskOrder --;

        //update compute node
        nodeSet storage n = nodeSets[node];
        n.assignCnt ++;
        n.assignedClient = 0;
        emit NodeEvent(n.nodeID, n.nodeAddr, n.hashRate, n.enrollTime, "node completed 1 task");
        emit CompleteEvent(o.orderID, o.client, o.index, n.nodeAddr, nowTime, resultHash);

        o.client = 0;
        o.orderID = 0;
        o.index = 0;
        o.state = taskState.Unrequested;
        o._state = 0;
        return true;
    }

    function assignTaskOrder(taskOrder storage o) internal
        returns (bool, address) {

        address nodeAddr = firstNodeSet;
        nodeSet storage ns = nodeSets[nodeAddr];
        uint nowTime = now;

        for (uint32 i = 0; i <= MAX_SEARCH_ASSIGNED_NODE; i++) {

            // only count the node which is not assigned
            if ((ns.enrolled == true) && (ns.assignedClient == 0)) {

                uint256 completeTime = (o.taskGas.div(ns.hashRate)).add(PREPARE_TIME);
                if (completeTime <= MAX_WAITING_TIME) {

                    //move this node to end of list
                    deleteNodeFromList(ns);
                    addNodetoList(ns, false);

                    //assigned
                    ns.assignedClient = o.orderID;
                    o.state = taskState.Assigned;
                    o._state = 2;
                    emit AssignEvent(o.orderID, o.client, o.index, nodeAddr, o.taskGas, nowTime, o.info[0].codeHash);
                    return(true, nodeAddr);
                }
            }

            nodeAddr = ns.nextNode;
            if(nodeAddr == 0) {
                return(false, 0);
            }
            ns = nodeSets[nodeAddr];
        }
        return (false, 0);
    }

    function isActiveOrder(bytes32 orderID) public
        returns(bool){

        if (taskOrderList.length == 0) {
            return false;
        }
        if ((orderBook[orderID].state != taskState.Unrequested) && (orderBook[orderID].orderID == orderID)) {
            return true;
        } else {
            return false;
        }
    }

    function getOrderBookCount()
        public returns (uint) {
        return taskOrderList.length;
    }

    function getOneTaskOrder(bytes32 orderID) public view
        returns (bool) {

        for (uint i = 0; i < taskOrderList.length; i++) {
            if (taskOrderList[i] == orderID) {
                taskOrder storage o = orderBook[orderID];
                emit ClientEvent(o.orderID, o.client, o.index, o.computeNode, o.taskGas, o.createTime, o._state, "query order");
                break;
            }
        }
        return (true);
    }

    function getAllTaskOrder() public
        returns (bool) {

        for (uint i = 0; i < taskOrderList.length; i++) {
            taskOrder storage o = orderBook[taskOrderList[i]];
            emit ClientEvent(o.orderID, o.client, o.index, o.computeNode, o.taskGas, o.createTime, o._state, "query oder");
        }
        return true;
    }

    function getAllNodeSet() public
        returns (bool) {

        address nodeAddr = firstNodeSet;
        while (nodeAddr != 0) {
            nodeSet storage n = nodeSets[nodeAddr];
            emit NodeEvent(n.nodeID, n.nodeAddr, n.hashRate, n.enrollTime, "node query");
            nodeAddr = n.nextNode;
        }
        return true;
    }

    function houseHoldAllTaskOrder() public
        returns (bool) {

        uint nowTime = now;
        for (uint i = 0; i < taskOrderList.length; i++) {

            taskOrder storage o = orderBook[taskOrderList[i]];
            emit LogEvent(o.orderID, o.client, o.computeNode, "periodically check");

            // not assigned yet, assign here
            if (o.state == taskState.Confirmed) {

                bool assigned;
                address nodeAddr;
                (assigned, nodeAddr) = assignTaskOrder(o);
                if (assigned == true) {
                    o.state = taskState.Assigned;
                    o._state = 3;
                    o.computeNode = nodeAddr;
                    o.assignTime = nowTime;
                }
            }

            // expired
            if (o.state == taskState.Assigned) {
                if ((nowTime - o.assignTime) > MAX_WAITING_TIME) {
                    computingNodeDisconnect (o.computeNode);
                }
            }

            // should not happen, do garbage collection in case
            if ((o.state == taskState.Completed) || (o.state == taskState.Canceled)) {
                //
            }
        }
        return true;
    }
}
