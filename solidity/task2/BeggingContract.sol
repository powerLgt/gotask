// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26;

contract BeggingContract {
    mapping (address => uint) userPayMoney;
    address owner;
    TopPay[] topPay;
    uint topSmallMoney = type(uint256).max;

    struct TopPay {
        address user;
        uint256 value;
    }

    // 只有在特定时间段内才能捐赠
    uint allowTimeBefore;

    event Donation(address from, uint money);

    modifier onlyOwner() {
        require(owner == msg.sender, "only owner can withdraw");
        _;
    }

    constructor(uint _allowTimeBefore) {
        owner = msg.sender;
        allowTimeBefore = _allowTimeBefore;
    }

    function donate() public payable  {
        require(allowTimeBefore > block.timestamp, "out of time to donate");
        if (msg.value <= 0) {
            return;
        }
        userPayMoney[msg.sender] += msg.value;
        _calcTopPay();
    }

    function withdraw() public onlyOwner {
        payable(owner).transfer(address(this).balance);
    }

    function getDonation(address user) public view returns (uint) {
        return userPayMoney[user];
    }

    function _calcTopPay() private {
        uint senderMoney = userPayMoney[msg.sender];
        uint length = topPay.length;

        // 已上榜的更新
        for (uint i = 0; i < length; i++) {
            if (topPay[i].user == msg.sender) {
                topPay[i].value = senderMoney;
                _updateTopSmallMoney();
                return;
            }
        }

        if (length < 3) {
            topPay.push(TopPay(msg.sender, senderMoney));
            _updateTopSmallMoney();
            return;
        }

        // 不够入榜
        if (topSmallMoney >= senderMoney) {
            return;
        }

        // 踢出top3被挤出去的
        for (uint i = 0; i < length; i++) {
            if (topSmallMoney == topPay[i].value) {
                topPay[i] = TopPay(msg.sender, senderMoney);
                break;
            }
        }

        _updateTopSmallMoney();
    }

    function _updateTopSmallMoney() private {
        // 计算top榜捐献最少的
        uint smallMoney = type(uint256).max;
        for (uint i = 0; i < topPay.length; i++) {
            if (smallMoney > topPay[i].value) {
                smallMoney = topPay[i].value;
            }
        }
        topSmallMoney = smallMoney;
    }

    function getTop3() public view returns (TopPay[] memory) {
        return topPay;
    }

    function getTopSmallMoney() public view returns (uint) {
        return topSmallMoney;
    }
}