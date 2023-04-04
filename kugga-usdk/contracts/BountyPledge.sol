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
import "./IBountyPledge.sol";

/**
 * reddit execute
 * https://polygonscan.com/address/0xab7a446621a2159fbff600276963a4113e7c6cfa#code
 */

error TimedForwarder__VerificationFailed();
error TimedForwarder__RequestExpired(uint256 expiry, uint256 actual);

contract BountyPledge is
    Initializable,
    EIP712Upgradeable,
    OwnableUpgradeable,
    ChainConstants,IBountyPledge
{
    using ECDSAUpgradeable for bytes32;
    using Transfers for address;
    using CountersUpgradeable for CountersUpgradeable.Counter;

    bool private _theWorldStatus;//如果 正常使用=true;禁用=false;
    string private _name;
    string private _symbol;
    //mapping(address => uint256) private _nonces;
    CountersUpgradeable.Counter _assetIdIndexer; //asset 的Id 自动增长
    mapping(bytes32 => BountyPledgeDataUpgradeable.Bounty) internal _bounties; //bytes32:keccak256(cid)=>Bounty
    mapping(bytes32 => string[]) internal _bountWinners; //bytes32:keccak256(cid)=>cid[]
    mapping(bytes32 =>bool) internal _contributeWinnerIndexer;//投稿获胜索引
    //代币资产索引
    mapping(address => BountyPledgeDataUpgradeable.AssetCoin) internal _supportedAssetCoins; // 支持支付的资产 assetAddress => assetCoin
    mapping(bytes32 => address) internal _supportedAssetSymbolIndexer; // 资产索引:bytes32:keccak256(unit) => assetAddress
   // mapping(uint8 => address) _supportedAssetIndexer; // 资产索引:assetId => assetAddress
    address[] private _assetCoinList; //资产列表

    function initialize(
        string calldata name_,
        string calldata symbol_,
        address usdt_,
        string calldata coinUnit_
    ) external initializer {
        __BountyPledge_init_unchained(name_, symbol_, usdt_, coinUnit_);
        __Ownable_init();
        __EIP712_init(name_, ERC712_VERSION);
    }

    function __BountyPledge_init_unchained(
        string calldata name_,
        string calldata symbol_,
        address usdt_,
        string calldata coinUnit
    ) internal onlyInitializing {
        _theWorldStatus = true; //
        _name = name_;
        _symbol = symbol_;
        _newAsset(usdt_, coinUnit);
    }

    function name() external view returns (string memory) {
        return _name;
    }

    function setName(string memory name_) external onlyOwner {
        _setName(name_);
    }

    function _setName(string memory name_) internal virtual {
        _name = name_;
    }

    function symbol() external view returns (string memory) {
        return _symbol;
    }

    function setSymbol(string memory symbol_) external onlyOwner {
        _setSymbol(symbol_);
    }

    function _setSymbol(string memory symbol_) internal virtual {
        _symbol = symbol_;
    }

    function getWorldStatus()public view returns(bool){
        return _theWorldStatus;
    }
    function enableWorldStatus() public onlyOwner{
        _theWorldStatus =true;
    }
    function disableWorldStatus() public onlyOwner {
        _theWorldStatus =false;
    }

    function getBounty(string memory cid) external view returns (BountyPledgeDataUpgradeable.Bounty memory) {
        return _getBounty(cid);
    }

    function getkeccak256(string memory cidUnit) public pure returns (bytes32) {
        return keccak256(abi.encodePacked(cidUnit));
    }

    //bounty 的获取与更新
    function _getBounty(string memory cid)
        internal
        view
        virtual
        returns (BountyPledgeDataUpgradeable.Bounty memory)
    {
        bytes32 keccak256CID = getkeccak256(cid);
        return _bounties[keccak256CID];
    }

    function _updateBounty(BountyPledgeDataUpgradeable.Bounty memory bounty)
        internal
        virtual
        returns (bool)
    {
        bytes32 keccak256CID = getkeccak256(bounty.cid);
        _bounties[keccak256CID] = bounty;
        return true;
    }

    function getWinners(string memory cid)
        external
        view
        returns (string[] memory)
    {
        return _bountWinners[getkeccak256(cid)];
    }

    //deprecated 这个仅仅作为测试
    function setBountyExpired(string memory cid, uint256 deadLine)
        external
        virtual
        onlyOwner
    {
        require(
            _bounties[getkeccak256(cid)].startTime > 0,
            "Pledge:not exists"
        );
        _bounties[getkeccak256(cid)].endTime = deadLine;
    }

    function addSupporteAsset(address asset_, string memory symbol_)
        external
        virtual
        onlyOwner
    {
        //需要验证asset 是 ERC20 类型的。后续再做，这里仅仅是为了方便测试
        //资产是否已经加过
        require(_theWorldStatus,"Pledge: wait for the system recovery ");
        require(asset_ != address(0x0), "Pledge:Not support Zero asset");
        require(
            _supportedAssetCoins[asset_].id < 1,
            "Pledge: Exists supported asset "
        );

        //添加资产
        _newAsset(asset_, symbol_);
    }

    /**
     *  添加资产
     */
    function _newAsset(address asset_, string memory symbol_)
        internal
        returns (BountyPledgeDataUpgradeable.AssetCoin memory)
    {
        _assetIdIndexer.increment();
        uint8 coinIdx = uint8(_assetIdIndexer.current());
        BountyPledgeDataUpgradeable.AssetCoin memory asset = BountyPledgeDataUpgradeable.AssetCoin({
            asset: asset_,
            symbol: symbol_,
            id: coinIdx,
            enabled: true
        });

        _supportedAssetCoins[asset_] = asset; // 资产地址到资产的对应关系
        _supportedAssetSymbolIndexer[getkeccak256(symbol_) ] = asset_; //维护关系：根据Id 查资产
        _assetCoinList.push(asset_); //添加进资产列表，所有的资产列表

        emit AssetActivity(
            asset.asset,
            asset.id, //内部编号
            1, // 1=添加，2=状态更改(禁用/启用)
            asset.symbol, // 资产简称
            asset.enabled //可用状态
        );
        return asset;
    }

    function disableAsset(address asset_) external virtual onlyOwner {
        require(_theWorldStatus,"Pledge: wait for the system recovery ");
        require(asset_ != address(0x0), "Pledge:Not support Zero asset");
        require(
            _supportedAssetCoins[asset_].id > 0,
            "Pledge: Not found the asset"
        );
        //资产不用
        _supportedAssetCoins[asset_].enabled = false;
        emit AssetActivity(
            _supportedAssetCoins[asset_].asset,
            _supportedAssetCoins[asset_].id, //内部编号
            2, // 1=添加，2=状态更改(禁用/启用)
            _supportedAssetCoins[asset_].symbol, // 资产简称
            _supportedAssetCoins[asset_].enabled //可用状态
        );
    }

    function enableAsset(address asset_) external virtual onlyOwner {
        require(_theWorldStatus,"Pledge: wait for the system recovery ");
        require(asset_ != address(0x0), "Pledge:Not support Zero asset");
        require(
            _supportedAssetCoins[asset_].id > 0,
            "Pledge: Not found the asset"
        );
        //资产再次启用
        _supportedAssetCoins[asset_].enabled = true;
        emit AssetActivity(
            _supportedAssetCoins[asset_].asset,
            _supportedAssetCoins[asset_].id, //内部编号
            2, // 1=添加，2=状态更改(禁用/启用)
            _supportedAssetCoins[asset_].symbol, // 资产简称
            _supportedAssetCoins[asset_].enabled //可用状态
        );
    }

    function getSupportedAsset(address asset_)
        external
        view
        returns (BountyPledgeDataUpgradeable.AssetCoin memory)
    {
        return _getSupportedAsset(asset_);
    }

    function _getSupportedAsset(address asset_)
        internal
        view
        returns (BountyPledgeDataUpgradeable.AssetCoin memory)
    {
        return _supportedAssetCoins[asset_];
    }

    function getSupportedAssetBySymbol(string memory symbol_)
        external
        view
        returns (BountyPledgeDataUpgradeable.AssetCoin memory)
    {
        BountyPledgeDataUpgradeable.AssetCoin memory _asset = _getSupportedAssetBySymbol(symbol_);
        //require(_supportedAssetIndexer[assetId] !=address(0x0),"Plege:Not found the asset indexer.");//可能存在本币支付
        require(_asset.id > 0, "Pledge: Not found the asset");

        return _asset;
    }

    function _getSupportedAssetBySymbol(string memory symbol_)
        internal
        view
        returns (BountyPledgeDataUpgradeable.AssetCoin memory)
    {
        return _supportedAssetCoins[_supportedAssetSymbolIndexer[getkeccak256(symbol_)]];
    }

    function _updateAsset(BountyPledgeDataUpgradeable.AssetCoin memory assertCoin)
        internal
        virtual
        returns (bool)
    {
        _supportedAssetCoins[assertCoin.asset] = assertCoin;
        return true;
    }

    function _checkSupportedAsset(string memory symbol_)
        internal
        virtual
        returns (BountyPledgeDataUpgradeable.AssetCoin memory)
    {
        BountyPledgeDataUpgradeable.AssetCoin memory assetCoin = _getSupportedAssetBySymbol(symbol_);
        //require(assetCoin.asset !=address(0x0),"Plege: Not support the asset");  //可能存在本币支付
        require(
            assetCoin.id > 0,
            "Pledge: Not support the asset or invalid asset"
        );
        require(assetCoin.enabled, "Pledge: The asset is disabled");
        return assetCoin;
    }

      /**
     * 在质押前：需要 钱包进行一次usdt 授权：approve
     * 质押：
     */
    function pledge2( uint256 startTime_,
        uint256 endTime_,
        uint8 bountyType_,
        string memory cid_,
        uint256 reward_,
        string memory rewardUnit_)
        external
        payable
        virtual
        returns (bytes32)
    {
        BountyPledgeDataUpgradeable.PledgeRequest memory req = BountyPledgeDataUpgradeable.PledgeRequest({
              startTime:startTime_,
          endTime:endTime_,
          bountyType:bountyType_,
            cid:cid_,
          reward:reward_,
            rewardUnit:rewardUnit_
        });
         return pledge(req);
    }

    /**
     * 在质押前：需要 钱包进行一次usdt 授权：approve
     * 质押：
     */
    function pledge(BountyPledgeDataUpgradeable.PledgeRequest memory req)
        public
        payable
        virtual
        returns (bytes32)
    {
        bytes32 keccak256Cid = getkeccak256(req.cid);
        require(_theWorldStatus,"Pledge: wait for the system recovery ");
        require(_bounties[keccak256Cid].startTime < 1 ,"Pledge: exists the same bounty"); //cid去重
        //验证是否支持此资产，支付
        BountyPledgeDataUpgradeable.AssetCoin memory assetCoin = _checkSupportedAsset(req.rewardUnit);
        require(
            req.reward > 0,
            "Pledge: guarantee funds should be greater than zero"
        );
        //require(req.startTime >= block.timestamp, "Pledge: startTime is expired");
        require(
            req.startTime < req.endTime && req.endTime > block.timestamp,
            "Pledge: invalid duration time"
        );
        //验资，质押合约能转么 ？如果不验证，就直接失败了也可以
        //uint256 allowanceAmount  = _USDTERC20.getAllowance(_msgSender());
        //require (allowanceAmount>=bounty_.reward.amount,"Pledge:insufficient allowance");
        BountyPledgeDataUpgradeable.Bounty memory _bounty = _newBounty(req);
        _bounties[keccak256Cid] = _bounty;
        //直接转入质押合约
       assetCoin.asset.transferIn(req.reward); //如果是合约地址，考虑回调，后续的 状态设置没有累加
        //index.increment
        
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
     * 新建一个 bounty
     */
    function _newBounty(BountyPledgeDataUpgradeable.PledgeRequest memory req)
        internal
        virtual
        returns (BountyPledgeDataUpgradeable.Bounty memory)
    {
        // Reward storage _reward = ;
        BountyPledgeDataUpgradeable.Bounty memory _bounty = BountyPledgeDataUpgradeable.Bounty({
            startTime: req.startTime,
            endTime: req.endTime,
            owner: _msgSender(),
            bountyType: req.bountyType,
            cid: req.cid,
            reward: BountyPledgeDataUpgradeable.Reward({
                amount: req.reward,
                balance: req.reward,
                unit: req.rewardUnit
            }),
            status: 1
        });
        return _bounty;
    }

    /**
     * 分账: 钱包发起交易
     */
    function transfer(
        string memory bountyCid,
        address to,
        uint256 amount,
        string memory contributeCid
    ) public payable virtual {
        bytes32 keccak256BountyCid = getkeccak256(bountyCid);
        require(_theWorldStatus,"Pledge: wait for the system recovery ");
        require(!_contributeWinnerIndexer[getkeccak256(contributeCid)],"Pledge:reduplicative contribute or winner");
        require(_bounties[keccak256BountyCid].startTime > 0, "Pledge: bounty not found");
        
        //bytes memory a = abi.encodePacked("Plege:the bounty is over.",StringsUpgradeable.toString(_bounties[bountyId].status));
        require(_bounties[keccak256BountyCid].status == 1, "Plege:The bounty is over");

        require(
            block.timestamp <= _bounties[keccak256BountyCid].endTime,
            "Pledge: The bounty activity is over"
        );

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

        BountyPledgeDataUpgradeable.AssetCoin memory assetCoin = _checkSupportedAsset(
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
        _contributeWinnerIndexer[getkeccak256(contributeCid)]=true;//文章获胜

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

    /**
     * 体现有严格的要求
     * 1、bounty 的owner 可以体现
     * 2、必须在过期后才能提币
     * 3、还存在提币金额
     */
    function withdraw(string memory bountyCid) public payable virtual {
        bytes32 keccak256BountyCid = getkeccak256(bountyCid);
        require(_theWorldStatus,"Pledge: wait for the system recovery ");
        //require(_bounties[bountyId].status ==1,"Plege:the bounty is over.");
        require(_bounties[keccak256BountyCid].startTime > 0, "Pledge: bounty not found");
        require(
            _bounties[keccak256BountyCid].owner == _msgSender(),
            "Pledge: Not the owner of the bounty"
        );
        require(
            block.timestamp >= _bounties[keccak256BountyCid].endTime,
            "Pledge: Wait for the expired time"
        );
        require(
            _bounties[keccak256BountyCid].reward.balance > 0,
            "Pledge: not sufficient funds"
        );

        BountyPledgeDataUpgradeable.AssetCoin memory assetCoin = _checkSupportedAsset(
            _bounties[keccak256BountyCid].reward.unit
        );

        uint256 _balance = _bounties[keccak256BountyCid].reward.balance;
        _bounties[keccak256BountyCid].reward.balance = 0;
        _bounties[keccak256BountyCid].status = 3; // 过期 ，保留数据
        assetCoin.asset.transferOut(_bounties[keccak256BountyCid].owner, _balance);

        emit Withdraw(_msgSender(), keccak256BountyCid, _balance,_bounties[keccak256BountyCid].status,_bounties[keccak256BountyCid].reward.unit, bountyCid);
       // emit BountyStatus( keccak256BountyCid, _bounties[keccak256BountyCid].status,  3,  bountyCid ); //model:1=bounty（投稿）,2=contribution(征稿),3= withdraw(提款)
    }

    receive() external payable {
        // this built-in function doesn't require any calldata,
        // it will get called if the data field is empty and
        // the value field is not empty.
        // this allows the smart contract to receive ether just like a
        // regular user account controlled by a private key would.
    }
    /*
    function destroy() public onlyOwner{
        selfdestruct(payable(_msgSender())); // 销毁合约
    }
    */
}
