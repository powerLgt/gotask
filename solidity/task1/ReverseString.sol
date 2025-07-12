// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26;

// ✅ 反转字符串 (Reverse String)
// 题目描述：反转一个字符串。输入 "abcde"，输出 "edcba"
contract ReverseString {
    
    function reverse(string memory str) public pure returns(string memory) {
        bytes memory data = bytes(str);
        uint count = data.length;
        bytes memory out = new bytes(count);
        for(uint i = 0; i < count; i++) {
            out[count - 1  - i] = data[i];
        }
        return string(out);
    }
}