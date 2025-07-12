// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26;

contract RomanToInt {
    mapping(string => uint256) private romanMap;
    
    constructor() {
        romanMap["I"] = 1;
        romanMap["V"] = 5;
        romanMap["X"] = 10;
        romanMap["L"] = 50;
        romanMap["C"] = 100;
        romanMap["D"] = 500;
        romanMap["M"] = 1000;
        romanMap["IV"] = 4;
        romanMap["IX"] = 9;
        romanMap["XL"] = 40;
        romanMap["XC"] = 90;
        romanMap["CD"] = 400;
        romanMap["CM"] = 900;
    }
    
    function romanToInt(string memory s) public view returns (uint256) {
        bytes memory roman = bytes(s);
        uint total = 0;
        uint i = 0;
        
        while (i < roman.length) {
            // 尝试匹配双字符组合
            if (i + 1 < roman.length) {
                string memory doubleChar = string(abi.encodePacked(roman[i], roman[i+1]));
                if (romanMap[doubleChar] != 0) {
                    total += romanMap[doubleChar];
                    i += 2;
                    continue;
                }
            }
            
            // 单字符匹配
            string memory singleChar = string(abi.encodePacked(roman[i]));
            total += romanMap[singleChar];
            i++;
        }
        
        return total;
    }
}