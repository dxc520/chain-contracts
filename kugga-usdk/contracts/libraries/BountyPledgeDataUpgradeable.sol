// SPDX-License-Identifier: MIT
pragma solidity >=0.4.22 <0.9.0;

library BountyPledgeDataUpgradeable {
    struct Bounty {
        uint256 startTime;
        uint256 endTime;
        address owner;
        uint8 bountyType; //single/multiple
        string cid; //描述信息
        Reward reward;
        uint8 status; // 1= 未结束(正常)、2=compelete,3=expired
    }
    struct Reward {
        uint256 amount;
        uint256 balance; //余额
        //uint8 unit; //赏金类型: USDT,MATIC
        string unit; //赏金类型: USDT,MATIC
    }
    struct PledgeRequest {
        uint256 startTime;
        uint256 endTime;
        uint8 bountyType; //1=single/2=multiple
        string cid;
        uint256 reward;
        string rewardUnit; //赏金类型: USDT,MATIC
    }
    struct AssetCoin {
        address asset; // 资产地址
        string symbol; // 资产简称
        uint8 id; //内部编号
        bool enabled; //可用状态
    }
}
