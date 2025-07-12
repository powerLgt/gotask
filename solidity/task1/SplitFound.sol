// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26;

contract SplitFound {
    function find(uint[] memory arr, uint target) public pure returns (int) {
        uint start;
        uint end = arr.length - 1;
        while (start <= end) {
            uint middle = start + (end - start) / 2;
            if (arr[middle] == target) {
                return int(middle);
            } else if (arr[middle] < target) {
                start = middle + 1;
            } else {
                end = middle - 1;
            }
        }

        return -1;
    }
}