// SPDX-License-Identifier: MIT
pragma solidity >=0.4.22 <0.9.0;

import "@openzeppelin/contracts-upgradeable/utils/CountersUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/access/AccessControlUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/security/ReentrancyGuardUpgradeable.sol";
import "./CellarBaseNftERC721.sol";


contract CellarNftSale is  OwnableUpgradeable, AccessControlUpgradeable,CellarBaseNftERC721,ReentrancyGuardUpgradeable{
    bytes32 private constant MINTER_ROLE = keccak256("MINTER_ROLE"); //铸币角色

    using CountersUpgradeable for CountersUpgradeable.Counter;
    address private _market; //市场地址
    CountersUpgradeable.Counter private _orderIds; //订单序列号,如果不要后期可以省略
    CountersUpgradeable.Counter private _tradeIds; //订单序列号,如果不要后期可以省略
    uint256 _currentTradeId;//当前最大交易id
    address private  _organiser; //交易费获得者地址
    uint16  _platformFeePercent; //平台费率
    //挂单销售
    struct Order{
        uint256 orderId;//订单ID
        uint256 tokenId; // tokenId
        address seller; //谁卖,合约现在的owner
        uint16  feePercent;//交易费提成比例
        uint256 sellingPrice; //销售价格
        //uint time;//挂单时间
        bool isSelling;//是否处于挂单状态
    }

    //交易信息
    struct SaleDetails{
        uint256 tradeId; //交易序号
        uint256 orderId;//订单ID 
        // Order  saleOrder;//订单Id;如果不考虑空间问题，可以转成对象 Oder //这里更改，考虑到实时销售订单完毕后就删除了，转至这里保存
        uint256 purchasePrice; //购买价格
        address buyer;// 购买者
       // uint time ;//购买时间
        bool isSuccess;//交易完成
    }

    //所有被交易的
    mapping (uint256=>bool) nftForSelling ; //正在热销的nft:被销售的tokenId 列表,记录当前在售的nft
    // 挂单销售列表
    //mapping (uint256=>Order) saleOrders; // 销售订单{tokenId=>order}，值记录当前的在销售状态
    
    mapping (uint256=>Order) saleOrderList; //销售订单(汇总){orderId=>oerder}
    mapping (uint256=>uint256) sellingNFTSaleOrders; // 当前销售订单{tokenId=>orderID}，值记录当前的在销售状态
     
    // 交易成功列表;
    mapping (uint256=>SaleDetails) tradeDetailsLog; //交易明细{tokenId=>SaleDetails}
    mapping (uint256=>uint256[])   tradeDetailsGroup; //tokenId多次交易{tokenId=>[tradeId1,tradeId2 ...]}
    //买家
    mapping(address => uint256[]) private purchasedTrades; //买家交易列表；user-[tradeId1,tradeId2 ...]

    event MakeOrder(address indexed operator,uint256 indexed tokenId, uint256 indexed orderId,uint256 sellingPrice,uint16  feePercent);
    event UpdateOrder(address indexed _operator,uint256 indexed _tokenId,uint256 indexed _orderId,uint256 _sellingPrice,uint16  _feePercent);
    event CancelOrder(address indexed _operator,uint256 indexed _tokenId,uint256 indexed _orderId);
    event Mint721Item(address _operator,address indexed _to, uint256 indexed _tokenId, uint256 indexed _amount,string  _tokenURI);
    event TransferNFT(address indexed _buyer,address indexed _seller, uint256 indexed _tokenId,uint256 _amount,uint256 _tradeId,uint256 _orderId);

    function initialize(string memory name,string memory symbol,uint16  platformFeePercent) initializer public {
       __ERC721_init(name, symbol);
       _init_sale_unchanged(platformFeePercent );
    }

    function _init_sale_unchanged(uint16  platformFeePercent )internal onlyInitializing{
        __Ownable_init();
        __ReentrancyGuard_init();
       _setupRole(MINTER_ROLE, _msgSender());//owner，可以铸币
       _platformFeePercent=platformFeePercent;// 平台抽取2个点的费率
       _organiser=_msgSender();
    }

    modifier isMinterRole {
        require(
            hasRole(MINTER_ROLE, _msgSender()),
            "User must have minter role to mint"
        );
        _;
    }

    modifier isCallerRoleOrOwner(address _operator,uint256 _tokenId) {
        require(
            hasRole(MINTER_ROLE, _msgSender()) || _isApprovedOrOwner(_operator, _tokenId),
            "User must have caller role or is the approvedOwner"
        );
        _;
    }
    /**
     * 查询是否拥有铸币权限
     */
    function hasMintRole(address _account) public view returns (bool) {
        return hasRole(MINTER_ROLE,_account);
    }

    /**
     *  授予用户铸币权限
     */
    function grantMintRole(address operator)  public onlyOwner {
        require(address(0)!=operator);
        _setupRole(MINTER_ROLE, operator);//owner，可以铸币
        //grantRole(MINTER_ROLE, operator);
    }

    /**
     *  收回用户铸币权限
     */
    function renounceMintRole( address account) public onlyOwner {
        renounceRole(MINTER_ROLE, account);
    }
  
    function grantMarket(address market) public onlyOwner returns(bool){
        _market=market;
        return true;
    }

   
    /**
     * 父类中重复定义，需要重写
     */
    function supportsInterface(bytes4 interfaceId) public view virtual override(ERC721Upgradeable, AccessControlUpgradeable)  returns (bool) {
        return interfaceId == type(IAccessControlUpgradeable).interfaceId
            || interfaceId == type(IERC721Upgradeable).interfaceId
            || interfaceId == type(IERC721MetadataUpgradeable).interfaceId
            || super.supportsInterface(interfaceId);
    }

    /**
     * 设置平台的交易费
     */
    function setPlatformFeePercent(uint16 platformFeePercent)public isMinterRole{
        require(platformFeePercent >=0 && platformFeePercent<=100,"SolarNFT: the platformFeePercent must be in [0,100].");
        _platformFeePercent=platformFeePercent;
    }
    function getPlatformFeePercent() public view returns(uint16){
        return _platformFeePercent;
    }
    /**
     *  申请铸币
     */
    function mintItem(address to, uint256 tokenId, string memory tokenURI) public isMinterRole returns(uint256){
        tokenId=_awardItem(to, tokenId, tokenURI);

        emit Mint721Item(_msgSender(),to, tokenId, 1, tokenURI);
        return tokenId;
    }

    /**
     * 挂单前，判断NFT是否正在销售列表中，不能重复挂单，
     */
    function _isSaleNFTAvailable(uint256 tokenId)internal view returns (bool){
       return  nftForSelling[tokenId];
    }

    /**
     * 卖家挂单,返回订单Id
     * 买家挂单，并授权给operator给市场
     */ 
    function makeOrder(uint256 tokenId,uint256 sellingPrice) public  nonReentrant returns(uint256){
 
        require(_isApprovedOrOwner(_msgSender(), tokenId), "SolarNFT: the caller is not owner nor approved");
        require(!_isSaleNFTAvailable(tokenId),"SolarNFT: the nft of token has a selling order, won't repeat.");

        //生成订单ID,或者采用 Hash算法也是可以的
        _orderIds.increment();
        uint256 _orderId=_orderIds.current();
        _currentTradeId=_orderId;
        //
        saleOrderList[_orderId]=Order(_orderId, tokenId, _msgSender(),_platformFeePercent,sellingPrice,true);
        sellingNFTSaleOrders[tokenId]= _orderId; 
        //转让operator到sale合约；But 这里得sale 合约应该 ApproveAll 权限，实际上在后续得transfer已经取得权限，没有必要 再次授权
        //solarNFT.approve(address(this), _tokenId); 

        if(!_isSaleNFTAvailable(tokenId)){
            nftForSelling[tokenId]=true;
        }
        
        super.approve(_market,tokenId);

        emit MakeOrder(_msgSender(), tokenId, _orderId,sellingPrice, _platformFeePercent);

        return _orderId;

    }

    /**
     * 只有合约或者seller可以撤单
     */
    function cancleOrder(uint256 tokenId) public  returns(bool) {
        return cancleOrder(_msgSender(),tokenId);
    }

    function cancleOrder(address operator,uint256 tokenId) public  isCallerRoleOrOwner(operator,tokenId) nonReentrant returns(bool) {
        require(saleOrderList[sellingNFTSaleOrders[tokenId]].isSelling,"SolarNFT: the order is not exists.");
        //require(_operator==saleOrderList[sellingNFTSaleOrders[_tokenId]].seller || solarNFT.isApprovedOrOwner(_operator, _tokenId),"SolarNFTSale: You won't cancle the order.Try to own the access.");
        require(_isApprovedOrOwner(operator, tokenId),"SolarNFT: You won't cancle the order.Try to own the access.");

        //收回operator 权限
        super.approve(address(0), tokenId); 

        uint256 _orderId = sellingNFTSaleOrders[tokenId];
        //remove
        _removeNFTForSale(tokenId);
        _removeSaleOrders(tokenId);

        emit CancelOrder( operator, tokenId,_orderId);

        return true;
    }
    
    /**
     * 谁能改价个呢？应该只有owner
     */
    function updateOrder(uint256 tokenId,uint256 sellingPrice) public  returns(bool) { 
        return updateOrder(_msgSender(),tokenId,sellingPrice);
    }

    function updateOrder(address owner,uint256 tokenId,uint256 sellingPrice) public isCallerRoleOrOwner(owner,tokenId) returns(bool) {
        require(saleOrderList[sellingNFTSaleOrders[tokenId]].isSelling,"SolarNFT: the order is not exists.");
        require(owner==super.ownerOf( tokenId),"SolarNFT: You are not the owner and won't update the order.");

        //更新订单数据
        saleOrderList[sellingNFTSaleOrders[tokenId]].sellingPrice=sellingPrice;
    
        emit  UpdateOrder(owner,  tokenId,sellingNFTSaleOrders[tokenId] ,sellingPrice,saleOrderList[sellingNFTSaleOrders[tokenId]].feePercent);

        return true;
    }

    /**
     * 当 cancle 或者 交易完毕时，需要从实时销售列表中移除掉 某个 nft
     */
    function _removeNFTForSale(uint256 tokenId) internal {
        delete nftForSelling[tokenId];
    }
    
    /**
     * 当 cancle 或者 交易完毕时，需要从实时销售列表中移除掉 某个 nft
     */
    function _removeSaleOrders(uint256 tokenId) internal{
        delete sellingNFTSaleOrders[tokenId];
        //至于销售订单要不要销毁，另计，需要考虑 交易合约
    }

    /**
     * 交易，转移nft
     */
    function transferNFT(address buyer, uint256 tokenId)public isMinterRole returns(uint256,uint256){
        //生成交易ID,或者采用 Hash算法也是可以的
        _tradeIds.increment();
        uint256 _tradeId=_tradeIds.current();
        _currentTradeId=_tradeId;
        address _tokenOwner=getNFTOwnerByTokenId(tokenId);
        super.safeTransferFrom(_tokenOwner, buyer, tokenId);

        //状态更新
        SaleDetails memory _saleDetails = SaleDetails({
            tradeId: _tradeId,
            orderId: sellingNFTSaleOrders[tokenId],
            purchasePrice: saleOrderList[sellingNFTSaleOrders[tokenId]].sellingPrice,
            buyer: buyer,
            isSuccess: true
        });
    
        tradeDetailsLog[_tradeId]=_saleDetails; //交易明细
        tradeDetailsGroup[tokenId].push(_tradeId); //根据tokenID 分组
        purchasedTrades[buyer].push(_tradeId); //买家的交易流水

        //消单(saleOrder)
        _removeNFTForSale(tokenId);
        _removeSaleOrders(tokenId);

        emit TransferNFT(buyer,_tokenOwner,tokenId,1,_tradeId,_saleDetails.orderId);
        return (_tradeId ,_saleDetails.orderId);
    }


    /**
     * 获取订单
     */
    function getOrderSellingPrice(uint256 tokenId)public view returns(uint256){
        return saleOrderList[sellingNFTSaleOrders[tokenId]].sellingPrice;
    }

    /**
     *  获取订单抽成比例
     */
    function getOrderFeePercent(uint256 tokenId)public view returns(uint16){
        return saleOrderList[sellingNFTSaleOrders[tokenId]].feePercent;
    }

    /**
     *  获取销售订单中，token 的owner 
     */
    function getNFTOwnerByTokenId(uint256 tokenId) public  view returns(address){
        return super.ownerOf(tokenId);
    }
 
    function getSaleOrderIdByToken(uint256 tokenId)public view returns(uint256){
        return sellingNFTSaleOrders[tokenId];
    } 

    function getCurrentMaxSaleOrderId()public view returns(uint256){
        //return _tradeIds.current();
        return _currentTradeId;
    }
}