// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26;

/**
    ✅ 创建一个名为Voting的合约，包含以下功能：
    一个mapping来存储候选人的得票数
    一个vote函数，允许用户投票给某个候选人
    一个getVotes函数，返回某个候选人的得票数
    一个resetVotes函数，重置所有候选人的得票数
*/
contract Voting {
    mapping(address => mapping(address => uint)) userVoters;
    mapping(address => uint) userVoteCount;
    mapping(address => uint) userVoteVersion;    // 投票版本号, 需要被投票人的投票版本号等于当前版本号，票数才算有效，避免重置时批量操作

    function vote(address voteTo) public {
        if (userVoteVersion[voteTo] == 0) {
            userVoteVersion[voteTo] = 1;
        }
        // 不能投票给自己
        require(voteTo != msg.sender, "you can't vote to yourself");
        // 已经投票过
        require(userVoters[voteTo][msg.sender] != userVoteVersion[voteTo], "you had voted to this user");
        userVoters[voteTo][msg.sender] = userVoteVersion[voteTo];
        userVoteCount[voteTo]++;
    }

    function getVotes(address voteTo) public view returns (uint) {
        return userVoteCount[voteTo];
    }

    function resetVotes(address voteTo) public {
        if (userVoteCount[voteTo] > 0) {
            userVoteCount[voteTo] = 0;
            userVoteVersion[voteTo]++;
        }
    }
}