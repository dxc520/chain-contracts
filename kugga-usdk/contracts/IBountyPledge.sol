// SPDX-License-Identifier: MIT
pragma solidity >=0.4.22 <0.9.0;

import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts-upgradeable/utils/cryptography/ECDSAUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/utils/cryptography/EIP712Upgradeable.sol";
import "@openzeppelin/contracts-upgradeable/utils/math/SafeMathUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/utils/CountersUpgradeable.sol";
import "./libraries/Transfers.sol";
import "./libraries/BountyPledgeDataUpgradeable.sol";
import "./ChainConstants.sol";

/**
 * reddit execute
 * https://polygonscan.com/address/0xab7a446621a2159fbff600276963a4113e7c6cfa#code
 */

interface IBountyPledge {
    event BountyActivity(
        address indexed from,
        bytes32 indexed keccak256CID,
        uint256 indexed reward,
        uint256 startTime,
        uint256 endTime,
        uint8 bountyType,
        uint8 status,
        string rewardUnit,
        string cid
    );
    event TransferActivity(
        address indexed from,
        address indexed to,
        bytes32 indexed keccak256BountyCID,
        uint256 reward,
        uint8 status,
        string rewardUnit,
        string bountyCid, // 征稿cid
        string contributeCid //投稿cid
    );
    event Withdraw(
        address indexed recipient,
        bytes32 indexed keccak256BountyCID,
        uint256 indexed reward,
        uint8 status,
        string rewardUnit,
        string bountyCid //征稿cid
    );

    event AssetActivity(
        address indexed asset,
        uint8 indexed Id, //内部编号
        uint8 indexed model, // 1=添加，2=状态更改(禁用/启用)
        string symbol, // 资产简称
        bool enabled //可用状态
    );

    function getBounty(string memory cid)
        external
        view
        returns (BountyPledgeDataUpgradeable.Bounty memory);

    function getkeccak256(string memory cidUnit)
        external
        pure
        returns (bytes32);

    function getWinners(string memory cid)
        external
        view
        returns (string[] memory);

    //deprecated 这个仅仅作为测试
    function setBountyExpired(string memory cid, uint256 deadLine) external;

    function addSupporteAsset(address asset_, string memory symbol_) external;

    function disableAsset(address asset_) external;

    function enableAsset(address asset_) external;

    function getSupportedAsset(address asset_)
        external
        view
        returns (BountyPledgeDataUpgradeable.AssetCoin memory);

    function getSupportedAssetBySymbol(string memory symbol_)
        external
        view
        returns (BountyPledgeDataUpgradeable.AssetCoin memory);

    /**
     * 在质押前：需要 钱包进行一次usdt 授权：approve
     * 质押：
     */
    function pledge2(
        uint256 startTime_,
        uint256 endTime_,
        uint8 bountyType_,
        string memory cid_,
        uint256 reward_,
        string memory rewardUnit_
    ) external payable returns (bytes32);

    /**
     * 在质押前：需要 钱包进行一次usdt 授权：approve
     * 质押：
     */
    function pledge(BountyPledgeDataUpgradeable.PledgeRequest memory req)
        external
        payable
        returns (bytes32);

    /**
     * 分账: 钱包发起交易
     */
    function transfer(
        string memory bountyCid,
        address to,
        uint256 amount,
        string memory contributeCid
    ) external payable;

    /**
     * 体现有严格的要求
     * 1、bounty 的owner 可以体现
     * 2、必须在过期后才能提币
     * 3、还存在提币金额
     */
    function withdraw(string memory bountyCid) external payable;
}
