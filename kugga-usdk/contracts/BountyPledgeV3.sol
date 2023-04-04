// SPDX-License-Identifier: MIT
pragma solidity >=0.4.22 <0.9.0;

import "./libraries/BountyPledgeDataUpgradeable.sol";
import "./BountyPledgeV2.sol";
import "./libraries/Transfers.sol";

/**
 * reddit execute
 * https://polygonscan.com/address/0xab7a446621a2159fbff600276963a4113e7c6cfa#code
 */

contract BountyPledgeV3 is BountyPledgeV2 {
    using Transfers for address;

    address internal _organizer; //平台组织者、承办单位，需要分账的
    uint8 internal _commissionRate; //平台抽成 (10=10/100)

    function getCommissionRate() public view returns (uint8) {
        return _commissionRate;
    }

    function setCommissionRate(uint8 commissionRate) public onlyOwner {
        require(
            commissionRate >= 0 && commissionRate <= 100,
            "Pledge:invalid commonsion rate.[0,100]"
        );
        _commissionRate = commissionRate;
    }

    function getOrganizer() public view returns (address) {
        return _organizer;
    }

    function setOrganizer(address organizer) public onlyOwner {
        require(organizer != address(0), "Pledge: invalid zero address");
        _organizer = organizer;
    }

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
    ) external payable virtual override returns (bytes32) {
        BountyPledgeDataUpgradeable.PledgeRequest
            memory req = BountyPledgeDataUpgradeable.PledgeRequest({
                startTime: startTime_,
                endTime: endTime_,
                bountyType: bountyType_,
                cid: cid_,
                reward: reward_,
                rewardUnit: rewardUnit_
            });
        return pledge(req);
    }

    /**
     * 在质押前：需要 钱包进行一次usdt 授权：approve
     * 质押：
     */
    function pledge(
        BountyPledgeDataUpgradeable.PledgeRequest memory req
    ) public payable virtual override returns (bytes32) {
        bytes32 keccak256Cid = getkeccak256(req.cid);
        require(getWorldStatus(), "Pledge: wait for the system recovery ");
        require(
            _bounties[keccak256Cid].startTime < 1,
            "Pledge: exists the same bounty"
        ); //cid去重
        //验证是否支持此资产，支付
        BountyPledgeDataUpgradeable.AssetCoin
            memory assetCoin = _checkSupportedAsset(req.rewardUnit);
        require(
            req.reward > 0,
            "Pledge: guarantee funds should be greater than zero"
        );
        //require(req.startTime >= block.timestamp, "Pledge: startTime is expired");
        require(
            req.startTime < req.endTime && req.endTime > block.timestamp,
            "Pledge: invalid duration time"
        );
        require(_organizer!=address(0),"Pledge: First Setting organizer.");
        //验资，质押合约能转么 ？如果不验证，就直接失败了也可以
        //uint256 allowanceAmount  = _USDTERC20.getAllowance(_msgSender());
        //require (allowanceAmount>=bounty_.reward.amount,"Pledge:insufficient allowance");
        BountyPledgeDataUpgradeable.Bounty memory _bounty = _newBounty(req);
        _bounties[keccak256Cid] = _bounty;
        //直接转入质押合约
        assetCoin.asset.transferIn(req.reward); //如果是合约地址，考虑回调，后续的 状态设置没有累加
        //平台账户抽成
        //int256 commission = 
        (bool b, uint256 commission) = SafeMathUpgradeable.tryDiv(req.reward, uint256(11));
        require(b,"Pledge:Fail to calculate commission");
        if (commission>0){ //如果佣金大于0 ，就继续转
            assetCoin.asset.transferOut(_organizer,commission);
        }

        //todo emit event
        emit BountyActivity(
            _bounty.owner,
            keccak256Cid,
            _bounty.reward.amount,
            _bounty.startTime,
            _bounty.endTime,
            _bounty.bountyType,
            _bounty.status,
            _bounty.reward.unit,
            _bounty.cid
        );
        //emit BountyStatus(keccak256Cid, _bounty.status, 1,_bounty.cid); //model:1=bounty（投稿）,2=contribution(征稿),3= withdraw(提款)

        return keccak256Cid;
    }

    /**
     * 分账: 钱包发起交易
     */
    function transfer(
        string memory bountyCid,
        address to,
        uint256 amount,
        string memory contributeCid
    ) public payable virtual override {
        bytes32 keccak256BountyCid = getkeccak256(bountyCid);
        require(getWorldStatus(), "Pledge: wait for the system recovery ");
        require(
            !_contributeWinnerIndexer[getkeccak256(contributeCid)],
            "Pledge:reduplicative contribute or winner"
        );
        require(
            _bounties[keccak256BountyCid].startTime > 0,
            "Pledge: bounty not found"
        );

        //bytes memory a = abi.encodePacked("Plege:the bounty is over.",StringsUpgradeable.toString(_bounties[bountyId].status));
        require(
            _bounties[keccak256BountyCid].status == 1,
            "Plege:The bounty is over"
        );

        // require(
        //     block.timestamp <= _bounties[keccak256BountyCid].endTime,
        //     "Pledge: The bounty activity is over"
        // );

        //bytes memory b = abi.encodePacked("Pledge: not the owner.",StringsUpgradeable.toHexString(_bounties[bountyId].owner));
        require(
            _bounties[keccak256BountyCid].owner == _msgSender(),
            "Pledge: not the owner."
        );
        require(
            amount > 0,
            "Pledge: transfer amount must be greater than zero"
        );
        require(
            _bounties[keccak256BountyCid].reward.balance >= amount,
            "Pledge: not sufficient funds"
        );

        BountyPledgeDataUpgradeable.AssetCoin
            memory assetCoin = _checkSupportedAsset(
                _bounties[keccak256BountyCid].reward.unit
            );

        (bool isSuccess, uint256 newBalance) = SafeMathUpgradeable.trySub(
            _bounties[keccak256BountyCid].reward.balance,
            amount
        );
        require(isSuccess, "Pledge:Fail to transfer.payment failed");

        //修改质押金额
        _bounties[keccak256BountyCid].reward.balance = newBalance;
        //获胜列表
        _bountWinners[keccak256BountyCid].push(contributeCid); //保存cid
        _contributeWinnerIndexer[getkeccak256(contributeCid)] = true; //文章获胜

        if (newBalance == 0) {
            _bounties[keccak256BountyCid].status = 2; // 结束
        }
        //扣款
        assetCoin.asset.transferOut(to, amount);

        emit TransferActivity(
            _bounties[keccak256BountyCid].owner,
            to,
            keccak256BountyCid,
            amount,
            _bounties[keccak256BountyCid].status,
            _bounties[keccak256BountyCid].reward.unit,
            bountyCid,
            contributeCid
        );
        // emit BountyStatus(  keccak256BountyCid,  _bounties[keccak256BountyCid].status, 2,  bountyCid ); //model:1=bounty（投稿）,2=contribution(征稿),3= withdraw(提款)
    }
}
