// SPDX-License-Identifier: MIT

pragma solidity ^0.8.26;

contract MergeArr {
    function merge(uint[] memory a, uint[] memory b) public pure returns (uint[] memory) {
        uint i;
        uint j;
        uint index;

        uint[] memory out = new uint[](a.length + b.length);

        while (i < a.length && j < b.length) {

            if (a[i] < b[j]) {
                out[index] = a[i];
                i++;
            } else {
                out[index] = b[j];
                j++;
            }
            index++;
        }

        while (i < a.length) {
            out[index] = a[i];
            i++;
            index++;
        }

        while (j < b.length) {
            out[index] = b[j];
            j++;
            index++;
        }

        return out;

    }
}