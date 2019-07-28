pragma solidity ^0.4.8;

library SafeMath {
    function mul(uint256 a, uint256 b) internal pure returns (uint256) {
        
        if (a == 0) {
            return 0;
        }

        uint256 c = a * b;
        assert(c / a == b);
        return c;
    }

    function div(uint256 a, uint256 b) internal pure returns (uint256) {

        return a / b;
    }

    function sub(uint256 a, uint256 b) internal pure returns (uint256) {

        assert(b <= a);
        return a - b;
    }

    function add(uint256 a, uint256 b) internal pure returns (uint256) {

        uint256 c = a + b;
        assert(c >= a);
        return c;
    }
}

contract Ownable {

    address public owner;
    event OwnershipTransferred(

        address indexed previousOwner,
        address indexed newOwner
    );

    constructor() public {

        owner = msg.sender;
    }

    modifier onlyOwner() {

        require(msg.sender == owner);
        _;
    }

    function transferOwnership(address _newOwner) public onlyOwner {

        _transferOwnership(_newOwner);
    }

    function _transferOwnership(address _newOwner) internal {

        require(_newOwner != address(0), "newOwner address is 0!");
        emit OwnershipTransferred(owner, _newOwner);
        owner = _newOwner;
    }
}

contract Config is Ownable{

    using Uint256Util for uint256;
    using StringUtil for string;

    struct ConfigEntry {

        uint256 index;
        string name;
        string valueString;
        uint256 valueInt;
    }

    mapping (string=>ConfigEntry) configEntries;
    string[] public configNames;

    function configSetString(string name, string value) public onlyOwner {

        ConfigEntry storage entry = configEntries[name];
        if (entry.index <= 0) {

            entry.index = configNames.length;
            entry.name = name;
            configNames.push(name);
        }
        entry.valueString = value;
        entry.valueInt = 0;
    }

    function configGetString(string name) public view returns (string) {

        return configEntries[name].valueString;
    }

    function configSetInt(string name, uint256 value) public onlyOwner {

        ConfigEntry storage entry = configEntries[name];
        if (entry.index <= 0) {
            entry.index = configNames.length;
            entry.name = name;
            configNames.push(name);
        }
        entry.valueInt = value;
        entry.valueString = "";
    }

    function configGetInt(string name) public view returns (uint256) {

        return configEntries[name].valueInt;
    }

    function configInfo() public view returns (string) {

        string memory str = "";
        for (uint256 index = 0; index < configNames.length; index++) {

            ConfigEntry storage entry = configEntries[configNames[index]];
            string memory value = entry.valueString;
            if (value.empty()) {
                value = entry.valueInt.toString();
            }
            str = str.concat(entry.name, "=", value, "|");
        }
        return str;
    }
}

library TimeUtil {

    using SafeMath for uint256;
    function isSameDay(uint256 time1, uint256 time2) internal pure returns (bool) {

        //GMT +8
        uint256 timeZone = uint256(8).mul(3600);
        uint256 secondsOneDay = uint256(24).mul(3600);
        uint256 day1 = time1.add(timeZone).div(secondsOneDay);
        uint256 day2 = time2.add(timeZone).div(secondsOneDay);
        if(day1 == day2){

            return true;
        }
        return false;
    }
}

library Int256Util {

     function abs(int256 value) internal pure returns (int256) {
        if(value < 0) {
            value = -value;
        }
        return value;
    }
}

library Uint256Util {

     function toString(uint i) internal pure returns (string){
        if (i == 0) return "0";
        uint j = i;
        uint len;
        while (j != 0){
            len++;
            j /= 10;
        }

        bytes memory bstr = new bytes(len);
        uint k = len - 1;
        while (i != 0){
            bstr[k--] = byte(48 + i % 10);
            i /= 10;
        }
        return string(bstr);
    }

    function diff(uint256 a, uint256 b) internal pure returns (uint256) {

        if(a > b){
            return a - b;
        }

        else {
            return b - a;
        }
    }
}

library PageUtil {

    using SafeMath for uint256;
    function pageRange(uint256 total, uint256 page, uint256 pageSize) internal pure returns (uint256 start, uint256 end) {

        if(total <= 0) {
            return (0, 0);
        }

        if(page <= 0) {
            page = 1;
        }

        if(pageSize == 0) {
            start = 0;
            end = total - 1;
            return (start, end);
        }

	    start = (page - 1).mul(pageSize);
        if(start >= total) {
            page = total/pageSize + 1;
            if(0 == total%pageSize) {
                page = page - 1;
            }
            start = (page - 1).mul(pageSize);
        }

        end = start + pageSize - 1;
        if(end >= total) {
            end = total - 1;
        }

        return (start, end);
    }
}

library AddressUtil {

    function codeSize(address _addr) internal view returns(uint256 _size) {
        assembly {  _size := extcodesize(_addr) }
    }

    function isContract(address _addr) internal view returns (bool) {
        uint256 size;
        assembly { size := extcodesize(_addr) }
        return size > 0;
    }
}

library StringUtil {

    function empty(string _a) internal pure returns (bool) {
        return (bytes(_a).length == 0);
    }

    function compare(string _a, string _b) internal pure returns (int) {

        bytes memory a = bytes(_a);
        bytes memory b = bytes(_b);
        uint minLength = a.length;
        if (b.length < minLength) 
        {
            minLength = b.length;
        }
            
        
        for (uint i = 0; i < minLength; i ++)

            if (a[i] < b[i])

                return -1;

            else if (a[i] > b[i])

                return 1;

        if (a.length < b.length)

            return -1;

        else if (a.length > b.length)

            return 1;

        else

            return 0;

    }



    function indexOf(string _haystack, string _needle) internal pure returns (int) {

        bytes memory h = bytes(_haystack);
        bytes memory n = bytes(_needle);
        if(h.length < 1 || n.length < 1 || (n.length > h.length))
            return -1;
        else if(h.length > (2**128 -1))
            return -1;
        else
        {
            uint subindex = 0;
            for (uint i = 0; i < h.length; i ++)
            {
                if (h[i] == n[0])
                {
                    subindex = 1;
                    while(subindex < n.length && (i + subindex) < h.length && h[i + subindex] == n[subindex])
                    {
                        subindex++;
                    }
                    if(subindex == n.length)
                        return int(i);
                }
            }
            return -1;
        }
    }

    function concat(string _a, string _b, string _c, string _d, string _e) internal pure returns (string) {
        bytes memory _ba = bytes(_a);
        bytes memory _bb = bytes(_b);
        bytes memory _bc = bytes(_c);
        bytes memory _bd = bytes(_d);
        bytes memory _be = bytes(_e);
        string memory abcde = new string(_ba.length + _bb.length + _bc.length + _bd.length + _be.length);
        bytes memory babcde = bytes(abcde);

        uint k = 0;
        for (uint i = 0; i < _ba.length; i++) babcde[k++] = _ba[i];
        for (i = 0; i < _bb.length; i++) babcde[k++] = _bb[i];
        for (i = 0; i < _bc.length; i++) babcde[k++] = _bc[i];
        for (i = 0; i < _bd.length; i++) babcde[k++] = _bd[i];
        for (i = 0; i < _be.length; i++) babcde[k++] = _be[i];
        return string(babcde);
    }

    function concat(string _a, string _b, string _c, string _d) internal pure returns (string) {

        return concat(_a, _b, _c, _d, "");
    }

    function concat(string _a, string _b, string _c) internal pure returns (string) {

        return concat(_a, _b, _c, "", "");
    }

    function concat(string _a, string _b) internal pure returns (string) {

        return concat(_a, _b, "", "", "");
    }

    // parseInt
    function parseInt(string _a) internal pure returns (uint) {

        return parseInt(_a, 0);
    }

    // parseInt(parseFloat*10^_b)
    function parseInt(string _a, uint _b) internal pure returns (uint) {
        bytes memory bresult = bytes(_a);
        uint mint = 0;
        bool decimals = false;
        for (uint i=0; i<bresult.length; i++){
            if ((bresult[i] >= 48)&&(bresult[i] <= 57)){
                if (decimals){
                   if (_b == 0) break;
                    else _b--;
                }
                mint *= 10;
                mint += uint(bresult[i]) - 48;
            } else if (bresult[i] == 46) decimals = true;
        }
        if (_b > 0) mint *= 10**_b;
        return mint;
    }
}