pragma solidity ^0.4.8;

import "./BokerManager.sol";
import "./BokerNodeData.sol";
import "./BokerLog.sol";

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

    constructor(address initialOwner,
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

        // Prevent transfer to 0x0 address. Use burn() instead
        require(_to != 0x0);

        // Check if the sender has enough
        require(balanceOf[_from] >= _value);

        // Check for overflows
        require(balanceOf[_to] + _value > balanceOf[_to]);

        // Save this for an assertion in the future
        uint previousBalances = balanceOf[_from] + balanceOf[_to];

        // Subtract from the sender
        balanceOf[_from] -= _value;

        // Add the same to the recipient
        balanceOf[_to] += _value;
        Transfer(_from, _to, _value);

        // Asserts are used to use static analysis to find bugs in your code. They should never fail
        assert(balanceOf[_from] + balanceOf[_to] == previousBalances);
    }

    /**
     * Transfer tokens
     *
     * Send `_value` tokens to `_to` from your account
     *
     * @param _to The address of the recipient
     * @param _value the amount to send
     */
    function transfer(address _to, uint256 _value) public {
        _transfer(msg.sender, _to, _value);
    }

    /**
     * Transfer tokens from other address
     *
     * Send `_value` tokens to `_to` in behalf of `_from`
     *
     * @param _from The address of the sender
     * @param _to The address of the recipient
     * @param _value the amount to send
     */
    function transferFrom(address _from, address _to, uint256 _value) public returns (bool success) {
        require(_value <= allowance[_from][msg.sender]);     // Check allowance
        allowance[_from][msg.sender] -= _value;
        _transfer(_from, _to, _value);
        return true;
    }

    /**
     * Set allowance for other address
     *
     * Allows `_spender` to spend no more than `_value` tokens in your behalf
     *
     * @param _spender The address authorized to spend
     * @param _value the max amount they can spend
     */
    function approve(address _spender, uint256 _value) public
        returns (bool success) {
        allowance[msg.sender][_spender] = _value;
        Approval(msg.sender, _spender, _value);
        return true;
    }

    /**
     * Set allowance for other address and notify
     *
     * Allows `_spender` to spend no more than `_value` tokens in your behalf, and then ping the contract about it
     *
     * @param _spender The address authorized to spend
     * @param _value the max amount they can spend
     * @param _extraData some extra information to send to the approved contract
     */

    function approveAndCall(address _spender, uint256 _value, bytes _extraData)
        public
        returns (bool success) {
        RecipientSuccess spender = RecipientSuccess(_spender);
        if (approve(_spender, _value)) {
            spender.receiveApproval(msg.sender, _value, this, _extraData);
            return true;
        }
    }

    /**
     * Destroy tokens
     *
     * Remove `_value` tokens from the system irreversibly
     *
     * @param _value the amount of money to burn
     */
    function burn(uint256 _value) public returns (bool success) {
        require(balanceOf[msg.sender] >= _value);   // Check if the sender has enough
        balanceOf[msg.sender] -= _value;            // Subtract from the sender
        totalSupply -= _value;                      // Updates totalSupply
        Burn(msg.sender, _value);
        return true;
    }

    /**
     * Destroy tokens from other ccount
     *
     * Remove `_value` tokens from the system irreversibly on behalf of `_from`.
     *
     * @param _from the address of the sender
     * @param _value the amount of money to burn
     */
    function burnFrom(address _from, uint256 _value) public returns (bool success) {
        require(balanceOf[_from] >= _value);                // Check if the targeted balance is enough
        require(_value <= allowance[_from][msg.sender]);    // Check allowance
        balanceOf[_from] -= _value;                         // Subtract from the targeted balance
        allowance[_from][msg.sender] -= _value;             // Subtract from the sender's allowance
        totalSupply -= _value;                              // Update totalSupply
        Burn(_from, _value);
        return true;
    }
}

/******************************************
 * This is an example contract that helps test the functionality of the 
 * approveAndCall() functionality of HumanStandardToken.sol.
 * This one assumes successful receival of approval.
 ********************************************/
contract RecipientSuccess {
  /* A Generic receiving function for contracts that accept tokens */
  address public from;
  uint256 public value;
  address public tokenContract;
  bytes public extraData;

  event ReceivedApproval(uint256 _value);

  function receiveApproval(address _from, uint256 _value, address _tokenContract, bytes _extraData) public {
    from = _from;
    value = _value;
    tokenContract = _tokenContract;
    extraData = _extraData;
    ReceivedApproval(_value);
  }
}

contract CWAREToken is BokerManaged, TokenERC20 {

    struct Task {
        uint256 index;
        uint256 value;
        string taskIp;
    }

    struct Tasks {
        uint256 totalValue;                 //total Token
        bytes32[] taskArray;                // task array
        mapping (bytes32 => Task) taskMap;  // bytes32 => taskId
    }

    string public name;                         //fancy name: eg ChainWare Token
    uint8 public decimals;                      //How many decimals to show
    string public symbol;                       //An identifier: eg SBX
    string public version = 'CWARE0.1';         //human 0.1 standard. Just an arbitrary versioning scheme.

    mapping (address =>Tasks) lockBox;
    
    event FrozenFunds(address _target, bool _frozen);
    event QueryFunds(address _target, uint256 value, uint256 _locked, bool _frozen);
    
    function CWAREToken(
        address _initialOwner,
        uint256 _initialAmount,
        string _tokenName,
        uint8 _decimalUnits,
        string _tokenSymbol,
        address addrManager
        ) TokenERC20(_initialOwner, _initialAmount, _tokenName, _tokenSymbol, _decimalUnits)
          BokerManaged(addrManager) public {
        balanceOf[_initialOwner] = _initialAmount;              // Give the creator all initial tokens
        totalSupply = _initialAmount;                           // Update total supply
        name = _tokenName;                                      // Set the name for display purposes
        decimals = _decimalUnits;                               // Amount of decimals for display purposes
        symbol = _tokenSymbol;                                  // Set the symbol for display purposes
    }

    /* Approves and then calls the receiving contract */
    function approveAndCall(address _spender, uint256 _value, bytes _extraData) public returns (bool success) {
        allowance[msg.sender][_spender] = _value;
        Approval(msg.sender, _spender, _value);

        //call the receiveApproval function on the contract you want to be notified. This crafts the function signature manually so one doesn't have to include a contract in here just for this.
        //receiveApproval(address _from, uint256 _value, address _tokenContract, bytes _extraData)
        //it is assumed when one does this that the call *should* succeed, otherwise one would use vanilla approve instead.
        require(_spender.call(bytes4(bytes32(sha3("receiveApproval(address,uint256,address,bytes)"))), msg.sender, _value, this, _extraData));
        return true;
    }


    /// @notice Create `mintedAmount` tokens and send it to `target`
    /// @param target Address to receive the tokens
    /// @param mintedAmount the amount of tokens it will receive
    function mintToken(address target, uint256 mintedAmount) onlyOwner public {
        balanceOf[target] += mintedAmount;
        totalSupply += mintedAmount;
        Transfer(0, this, mintedAmount);
        Transfer(this, target, mintedAmount);
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
        Transfer(_from, _to, _value);
        return true;
    }

    function lockToken(address target, bytes32 taskId, uint256 lockedAmount, string taskIp) public returns(bool) {

        //1. enough funds 2.lockbox empty 3.account is not frozen
        if (balanceOf[target] < lockedAmount) {
            return false;
        }

        //check taskId
        if(lockBox[target].taskMap[taskId].value > 0){
            return false;
        }

        //write taskId
        lockBox[target].taskMap[taskId].value = lockedAmount;
        lockBox[target].taskMap[taskId].taskIp = taskIp;
        lockBox[target].taskMap[taskId].index = lockBox[target].taskArray.length;
        lockBox[target].totalValue += lockedAmount;
        lockBox[target].taskArray.push(taskId);

        balanceOf[target] -= lockedAmount;
        return (true);
    }

    function unLockToken(address target, bytes32 taskId) public returns(bool) {

        // unlock token;
        var lockedAmount = lockBox[target].taskMap[taskId].value;
        var index = lockBox[target].taskMap[taskId].index;
        balanceOf[target] += lockedAmount;
        lockBox[target].totalValue -= lockedAmount;
        delete lockBox[target].taskMap[taskId];
        delete lockBox[target].taskArray[index];
        return (true);
    }

    function getAccountBalanceOf(address target) public view onlyOwner returns(uint256) {
        return (balanceOf[target]);
    }

    function getLockBoxBalanceOf(address target) public view onlyOwner returns(uint256) {
        return (lockBox[target].totalValue);
    }

    function getMyAccountBalanceOf() public view returns(uint256) {
        return (balanceOf[msg.sender]);
    }

    function getMyLockBoxBalanceOf() public view
            returns(uint256) {
        return (lockBox[msg.sender].totalValue);
    }
}

contract Chainware is BokerManaged{

    uint32 constant  MAX_WAITING_TIME = 300;        //max wait for 300s
    uint32 constant  PREPARE_TIME = 10;             //computing node estimated prepare time
    uint32 constant  MAX_SEARCH_ASSIGNED_NODE = 10; //max number of search compute nodes
                                                    // to assign a task

    address public CWARE_Token_Addr;
    uint public numTaskOrder;
    uint public numComputeNode;

    enum taskState {Unrequested, Confirmed, Assigned, Completed, Canceled}

    struct taskInfo {
        string  codeHash;
        string  dataHash;
        string  taskIp;
        bytes32 taskId;
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
                var n = nodeSets[firstNodeSet];
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

    //constructor
    function Chainware(address cware_token, address addrManager) BokerManaged(addrManager) public {
        numTaskOrder = 0;
        numComputeNode = 0;
        CWARE_Token_Addr = cware_token;
        firstNodeSet = 0;
        lastNodeSet = 0;
    }
    
    using SafeMath for uint256;

    //define events
    event ClientEvent(bytes32 _id, address indexed _client, uint64 _indexID, address indexed _computeNode, uint _taskGas, uint _createTime, uint32 _state, string _reason);
    event NodeEvent(bytes32 _id, address indexed _node, uint _hashRate, uint _enrollTime, string _reason);
    event AssignEvent(bytes32 _id, address indexed _client, uint64 _index, address indexed _computeNode, uint _taskGas, uint _assignTime, string _codeHash);
    event CompleteEvent(bytes32 _id, address _client, uint64 _index, address _computeNode, uint _completeTime, string _resultHash);
    event LogEvent(bytes32 _id, address indexed _client, address indexed _computeNode, string _reason);

    ////////////////From Clients /////////////////////

    function getTaskId(uint64 index)
        public returns (bytes32){
            return keccak256(msg.sender, index);
        }
    
    /// @notice Create `taskOrderRequest` create task order for request
    /// @param taskGas the task computing amount
    /// @param codeHash metadata from client
    // dataHash not used yet
    //Request
    function taskOrderRequest(uint64 index, uint256 taskGas, string codeHash, string taskIp)
        public returns (bool) {
        bytes32 orderID = keccak256(msg.sender, taskGas, taskState.Confirmed, now);
        var o = orderBook[orderID];
        
        if ((o.state != taskState.Unrequested) && (o.client == msg.sender)) {
            //already has active serverice, cancel first
            ClientEvent(o.orderID, o.client, o.index,  o.computeNode, o.taskGas, o.createTime, o._state, "task already existed, cancel first");
            return false;
        }

        //create taskId
        bytes32 taskId = getTaskId(index);

        //Lock the taskGas, must be successful.
        if (CWAREToken(CWARE_Token_Addr).lockToken(msg.sender, taskId, taskGas, taskIp) != true) {
            ClientEvent(orderID, msg.sender, index, 0, taskGas, 0, 0, "Lock token failed");
            return false;
        }

        o.client = msg.sender;
        o.index = index;
        o.taskGas = taskGas;
        o.info[0].codeHash = codeHash;
        o.state = taskState.Confirmed;
        o._state = 1;
        o.createTime = now;
        o.info[0].taskId = taskId;
        o.info[0].taskIp = taskIp;
        o.orderID = sha3(o.client, o.taskGas, o.state, now);
        //push to the list
        o.listPtr = taskOrderList.push(o.orderID) - 1;
        numTaskOrder ++;

        //then assign to compute node
        var (assigned, nodeAddr) = assignTaskOrder(o);
        if (assigned == true) {
            o.state = taskState.Assigned;
            o._state = 2;
            o.computeNode = nodeAddr;
            o.assignTime = now;
        }

        ClientEvent(o.orderID, o.client, o.index, o.computeNode, o.taskGas, o.createTime, o._state, "order request confirmed");
        return true;
    }

    // Cancel
    /// @notice Cancel `taskOrderCancel` cancel task order from client
    function taskOrderCancel(bytes32 orderID) public returns (bool) { 
        
        //send out the request
        var o = orderBook[orderID];
        if (o.state == taskState.Unrequested) {
            //not the owner
            return false;
        }

        if (o.client != msg.sender) {
            //only client can cancel order
            return (false);
        }

        if (o.state >= taskState.Assigned) {
            //too late
            return (false);
        }

        o.state = taskState.Canceled;
        o._state = 4;
        //refund the taskGas

        CWAREToken(CWARE_Token_Addr).unLockToken(o.client, o.info[0].taskId);

        //send event out
        if (o.client != 0) {
            ClientEvent(o.orderID, o.client, o.index, o.computeNode, o.taskGas, o.createTime, o._state, "order canceled");
        }

        //remove from the link list
        uint rowToDelete = orderBook[o.orderID].listPtr;
        bytes32 keyToMove = taskOrderList[taskOrderList.length -1];
        taskOrderList[rowToDelete] = keyToMove;
        orderBook[keyToMove].listPtr = rowToDelete;
        taskOrderList.length --;

        o.client = 0;
        o.state = taskState.Unrequested;
        o._state = 0;
        o.index = 0;
        o.info[0].taskId = 0;
        o.info[0].taskIp = "";
        numTaskOrder --;
        return true;
    }


    /////////// Call from computing node ////////////////
    /// @notice  `computingNodeEnroll` computing node enrolk to chainware
    /// @param hashRate the computing node computing power
    function computingNodeEnroll(uint64 hashRate) public returns (bool) {

        var n = nodeSets[msg.sender];

        // no need to enroll twice
        if (n.enrolled == true) {
            NodeEvent(n.nodeID, n.nodeAddr, n.hashRate, n.enrollTime, "node already enrolled");
            return false;
        }

        n.nodeAddr = msg.sender;
        n.hashRate = hashRate;
        n.enrollTime = now;
        n.enrolled = true;
        n.nodeID = sha3(n.nodeAddr, hashRate, now);

        addNodetoList(n, true);

        NodeEvent(n.nodeID, n.nodeAddr, n.hashRate, n.enrollTime, "node enrolled OK");
        numComputeNode ++;
        return true;
    }

    //once computing node is assigned, offchain heart beat mechnism is introduced
    //to guarantee the work is going on. The disconnected computing node which is
    //detected call for this method
    /// @notice `CcomputingNodeDisconnect` disconnect for a specific node
    /// @param node the disconnected computing node address
    function computingNodeDisconnect (address node) public returns (bool) {
        var n = nodeSets[node];
        if (n.enrolled != true) {
            NodeEvent(n.nodeID, n.nodeAddr, n.hashRate, n.enrollTime, "node is not enrolled");
            return false;
        }

        var client = n.assignedClient;
        if (client != 0) {
            var o = orderBook[client];
            o.state = taskState.Confirmed;
            o._state = 1;
            var (assigned, newNodeAddr) = assignTaskOrder(o);
            if (assigned == true) {
                o.state = taskState.Assigned;
                o._state = 2;
                o.computeNode = newNodeAddr;
                o.assignTime = now;
            }
        }

        deleteNodeFromList(n);
        NodeEvent(n.nodeID, n.nodeAddr, n.hashRate, n.enrollTime, "node removed");
        numComputeNode --;

        n.nodeAddr = 0;
        n.enrolled = false;
        //n.nodeID = 0;

        return true;
    }

    //this method is called after the computing node complete the
    //computing task
    /// @notice Create `computingTaskComplete` task order of computing is completed
    /// @param node the computing node address
    /// @param orderID the task orderID
    /// @param realGas the true gas of task execution
    /// @param resultHash the computing result hash in ipfs
    function computingTaskComplete(address node, bytes32 orderID, uint256 realGas, string resultHash) public returns (bool) {

        var o = orderBook[orderID];
        if ((o.state == taskState.Unrequested) || (o.orderID != orderID)) {
            //something wrong or re-assigned
            return false;
        }

        o.state = taskState.Completed;
        o._state = 3;
        ClientEvent(o.orderID, o.client, o.index, o.computeNode, o.taskGas, o.createTime, o._state, "task completed");

        //unlock first
        CWAREToken(CWARE_Token_Addr).unLockToken(o.client, o.info[0].taskId);

        //if realGas is bigger than estimated gas, use estimated gas
        //otherwise use the real one
        var taskGas = (realGas >= o.taskGas) ? o.taskGas : realGas;
        var payMinerAmount = taskGas.div(5);
        var payNodeAmount = taskGas - payMinerAmount;
        CWAREToken(CWARE_Token_Addr)._transferFrom(o.client, node, payNodeAmount);
        //miner's rewards fee
        CWAREToken(CWARE_Token_Addr)._transferFrom(o.client, block.coinbase, payMinerAmount);

        //remove this task order out of order book
        uint rowToDelete = orderBook[o.orderID].listPtr;
        bytes32 keyToMove = taskOrderList[taskOrderList.length -1];
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
        var n = nodeSets[node];
        n.assignCnt ++;
        n.assignedClient = 0;
        NodeEvent(n.nodeID, n.nodeAddr, n.hashRate, n.enrollTime, "node completed 1 task");
        CompleteEvent(o.orderID, o.client, o.index, n.nodeAddr, now, resultHash);

        o.client = 0;
        o.orderID = 0;
        o.index = 0;
        o.state = taskState.Unrequested;
        o._state = 0;
        return true;
    }

    /////////////// House hold of Task Order/////////////
    //
    //internal method, called by task order is created and periodic check
    /// @notice Create `assignTaskOrder` assign a task order to computing node
    /// @param o task order storage o
    function assignTaskOrder(taskOrder storage o) internal returns (bool, address) {

        var nodeAddr = firstNodeSet;
        var ns = nodeSets[nodeAddr];

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
                    AssignEvent(o.orderID, o.client, o.index, nodeAddr, o.taskGas, now, o.info[0].codeHash);
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

    /// @notice Create `taskOrderRequest` create task order for request
    /// @param orderID task orderID
    function isActiveOrder(bytes32 orderID) internal
        returns(bool){
        if (taskOrderList.length == 0) {
            return false;
        }
        if ((orderBook[orderID].state != taskState.Unrequested) &&
            (orderBook[orderID].orderID == orderID)) {
            return true;
        } else {
            return false;
        }
    }
 
    /// @notice Create `taskOrderRequest` create task order for request
    function getOrderBookCount()
        public returns (uint) {
        return taskOrderList.length;
    }
  
    //query the specific task order in order book
    /// @notice Create `getOneTaskOrder` get one task order
    /// @param orderID the task computing node address
    function getOneTaskOrder(bytes32 orderID) public view returns (bool) {
        for (uint i=0; i<taskOrderList.length; i ++) {
            if (taskOrderList[i] == orderID) {
                var o = orderBook[orderID];
                ClientEvent(o.orderID, o.client, o.index, o.computeNode, o.taskGas, o.createTime, o._state, "query order");
                break;
            }
        }
        return (true);
    }
  
    //query all task node
    /// @notice Create `taskOrderRequest` create task order for request
    function getAllTaskOrder() public returns (bool) {

        for (uint i=0; i<taskOrderList.length; i++) {
            var o = orderBook[taskOrderList[i]];
            ClientEvent(o.orderID, o.client, o.index, o.computeNode, o.taskGas, o.createTime, o._state, "query oder");
        }

        return true;
    }

    //query all node sets
    /// @notice Create `taskOrderRequest` create task order for request
    function getAllNodeSet() public returns (bool) {

        var nodeAddr = firstNodeSet;
        while (nodeAddr != 0) {
            var n = nodeSets[nodeAddr];
            NodeEvent(n.nodeID, n.nodeAddr, n.hashRate, n.enrollTime, "node query");
            nodeAddr = n.nextNode;
        }

        return true;
    }

    // periodically household all active task order
    // This is the slow check. every 5 minutes???
    // fast check is on offchain and call computeNodeDisconnect
    /// @notice Create `taskOrderRequest` create task order for request
    function houseHoldAllTaskOrder() public returns (bool) {
        
        for (uint i=0; i<taskOrderList.length; i++) {
            var o = orderBook[taskOrderList[i]];
            LogEvent(o.orderID, o.client, o.computeNode, "periodically check");

            // not assigned yet, assign here
            if (o.state == taskState.Confirmed) {
                var (assigned, nodeAddr) = assignTaskOrder(o);
                if (assigned == true) {
                    o.state = taskState.Assigned;
                    o._state = 3;
                    o.computeNode = nodeAddr;
                    o.assignTime = now;
                }
            }

            // expired
            if (o.state == taskState.Assigned) {
                if ((now - o.assignTime) > MAX_WAITING_TIME) {
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
