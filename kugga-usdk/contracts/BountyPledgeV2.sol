// SPDX-License-Identifier: MIT
pragma solidity >=0.4.22 <0.9.0;

import "./libraries/BountyPledgeDataUpgradeable.sol";
import "./BountyPledge.sol";

/**
 * reddit execute
 * https://polygonscan.com/address/0xab7a446621a2159fbff600276963a4113e7c6cfa#code
 */

contract BountyPledgeV2 is BountyPledge {
    using Transfers for address;
    event RewardActivity(
        bytes32 indexed keccak256CID, //keccak256(cid)
        uint256 indexed reward,
        string rewardUnit,
        uint8 indexed model, //1=增量，2=全量设置
        uint8 status,
        string cid
    );

    function increaseReward(string memory cid, uint256 amount) external virtual {
        BountyPledgeDataUpgradeable.Bounty memory bounty = _getBounty(cid);
        require(getWorldStatus(),"Pledge: wait for the system recovery ");
        require(amount > 0, "Pledge: invalid reward");
        require(bounty.startTime > 0, "Pledge: bounty not found");
        //bytes memory a = abi.encodePacked("Plege:the bounty is over.",StringsUpgradeable.toString(_bounties[bountyId].status));
        require(bounty.status == 1, "Plege:the bounty is over");
        require(bounty.owner == _msgSender(), "Pledge: not the owner.");

       BountyPledgeDataUpgradeable. AssetCoin memory assetCoin = _checkSupportedAsset(bounty.reward.unit);
        assetCoin.asset.transferIn(amount);

        (bool isSuccess, uint256 newAmount) = SafeMathUpgradeable.tryAdd(
            bounty.reward.amount,
            amount
        );
        require(isSuccess, "Pledge:Fail to increase reward");
        bounty.reward.amount = newAmount;
        bounty.reward.balance += amount; //总量不溢出，这个不会溢出

        _updateBounty(bounty);
        emit RewardActivity(
            getkeccak256(bounty.cid),
            amount,
            bounty.reward.unit,
            1, //1=增量，2=全量设置
            bounty.status,
            bounty.cid
        );
    }

    function approveReward(string memory cid, uint256 amount) external virtual{
        BountyPledgeDataUpgradeable.Bounty memory bounty = _getBounty(cid);
        require(getWorldStatus(),"Pledge: wait for the system recovery ");
        require(amount > 0, "Pledge: invalid reward");
        require(bounty.startTime > 0, "Pledge: bounty not found");
        require(bounty.status == 1, "Plege:the bounty is over");
        require(bounty.owner == _msgSender(), "Pledge: not the owner.");
        require(
            bounty.reward.amount < amount,
            "Pledge: invalid reward,need greater"
        );

        //uint256 balance = amount - bounty.reward.amount; // 不会溢出
        (bool isSuccess, uint256 newBalance) = SafeMathUpgradeable.trySub(
            amount,
            bounty.reward.amount
        );
        require(isSuccess, "Pledge:Fail to approve reward");

        BountyPledgeDataUpgradeable.AssetCoin memory assetCoin = _checkSupportedAsset(bounty.reward.unit);
        assetCoin.asset.transferIn(newBalance);

        bounty.reward.amount = amount;
        bounty.reward.balance += newBalance; //总量不溢出，这个不会溢出

        _updateBounty(bounty);
        emit RewardActivity(
            getkeccak256(bounty.cid),
            amount,
            bounty.reward.unit,
            2, //1=增量，2=全量设置
            bounty.status,
            bounty.cid
        );
    }
}
