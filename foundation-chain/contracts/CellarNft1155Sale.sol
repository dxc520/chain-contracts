// SPDX-License-Identifier: MIT
pragma solidity >=0.4.22 <0.9.0;

import "@openzeppelin/contracts-upgradeable/security/ReentrancyGuardUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/utils/CountersUpgradeable.sol";

import "./CellarBaseNftERC1155.sol";

contract CellarNft1155Sale is CellarBaseNftERC1155, ReentrancyGuardUpgradeable {
    using CountersUpgradeable for CountersUpgradeable.Counter;
    CountersUpgradeable.Counter private _orderIds; //订单序列号,如果不要后期可以省略
    CountersUpgradeable.Counter private _tradeIds; //订单序列号,如果不要后期可以省略

    address private _market; //市场地址
    uint256 _currentTradeId; //当前最大交易id
    address private _organiser; //交易费获得者地址
    uint16 _platformFeePercent; //平台费率

    //挂单销售
    struct Order {
        uint256 orderId; //订单ID
        address seller; //谁卖,
        uint256 tokenId; // tokenId
        uint256 stockAmount; //库存数量，当前最大售卖量
        uint16 feePercent; //交易费提成比例
        uint256 unitPrice; //销售价格
        //uint time;//挂单时间
        bool isSelling; //是否处于挂单状态
    }

    //交易信息
    struct SaleDetails {
        uint256 tradeId; //交易序号
        uint256 orderId; //订单ID
        // Order  saleOrder;//订单Id;如果不考虑空间问题，可以转成对象 Oder //这里更改，考虑到实时销售订单完毕后就删除了，转至这里保存
        uint256 purchasePrice; //购买价格
        address buyer; // 购买者
        //uint time ;//购买时间
        bool isSuccess; //交易完成
    }

    //所有被交易的
    mapping(uint256 => mapping(address => uint256)) private _presellBalances; //由于存在挂单，就不是实际转移，方便后期超售管理，增加此状态
    //mapping (uint256=> mapping (address=>bool)) nftForSelling ; //{tokenId=>{owner=>orderID}}正在热销的nft:被销售的tokenId 列表,记录当前在售的nft
    // 挂单销售列表
    //mapping (uint256=>Order) saleOrders; // 销售订单{tokenId=>order}，值记录当前的在销售状态

    mapping(uint256 => Order) saleOrderList; //销售订单(汇总){orderId=>oerder}
    //mapping (uint256=>mapping(address=>uint256[])) sellingNFTSaleOrders; // 当前销售订单{tokenId=>{owner=>orderID[]}}，值记录当前的在销售状态

    // 交易成功列表;
    mapping(uint256 => SaleDetails) tradeDetails; //交易明细{tradeId=>SaleDetails}

    //events
    event Mint1155Item(
        address operator,
        address indexed _to,
        uint256 indexed _tokenId,
        uint256 indexed _amount,
        string _tokenURI
    );
    event MintBatchItems(
        address indexed operator,
        address indexed _to,
        uint256[] _tokenIds,
        uint256[] _amounts,
        string[] _tokenUris
    );
    event MakeOrder(
        address operator,
        address indexed seller,
        uint256 indexed tokenId,
        uint256 amount,
        uint256 indexed orderId,
        uint256 unitPrice,
        uint16 feePercent
    );
    event UpdateOrder(
        address _operator,
        address indexed _seller,
        uint256 indexed _tokenId,
        uint256 _amount,
        uint256 indexed _orderId,
        uint256 _unitPrice,
        uint16 _feePercent
    );
    event CancelOrder(
        address _operator,
        address indexed _seller,
        uint256 indexed _tokenId,
        uint256 _amount,
        uint256 indexed _orderId,
        uint256 _unitPrice,
        uint16 _feePercent
    );
    event TradeNft(
        address indexed _buyer,
        address indexed _seller,
        uint256 indexed _tokenId,
        uint256 amount,
        uint256 _tradeId,
        uint256 _orderId
    );

    function initialize(
        string memory name,
        string memory symbol,
        string memory uri,
        uint16 platformFeePercent
    ) public initializer {
        __ERC1155URIStorage_init(name, symbol, uri);
        _init_base_unchanged();
        _init_sale_unchanged(platformFeePercent);
    }

    function _init_sale_unchanged(uint16 platformFeePercent)
        internal
        onlyInitializing
    {
        __ReentrancyGuard_init();
        _platformFeePercent = platformFeePercent; // 平台抽取2个点的费率
        _organiser = _msgSender();
    }

    /**
     * 设置平台的交易费
     */
    function setPlatformFeePercent(uint16 platformFeePercent)
        public
        isMinterRole
    {
        require(
            platformFeePercent >= 0 && platformFeePercent <= 100,
            "SolarNFT: the platformFeePercent must be in [0,100]."
        );
        _platformFeePercent = platformFeePercent;
    }

    function getPlatformFeePercent() public view returns (uint16) {
        return _platformFeePercent;
    }

    function grantMarket(address market) public onlyOwner returns (bool) {
        _market = market;
        return true;
    }

    function getMarket() public view returns (address) {
        return _market;
    }

    /**
     *  申请铸币
     */
    function mintItem(
        address to,
        uint256 id,
        uint256 amount,
        string memory tokenURI
    ) public isMinterRole {
        require(amount > 0, "ERC1155:minted amount must be greater than zero");

        _mintExt(to, id, amount, tokenURI);

        emit Mint1155Item(_msgSender(), to, id, amount, tokenURI);
    }

    function mintBatchItems(
        address to,
        uint256[] memory ids,
        uint256[] memory amounts,
        string[] memory tokenUris
    ) public isMinterRole {
        _mintBatchExt(to, ids, amounts, tokenUris);

        emit MintBatchItems(_msgSender(), to, ids, amounts, tokenUris);
    }

    /**
     * 出多少钱，从谁那里买
     */
    function makeOrder(
        uint256 tokenId,
        uint256 amount,
        uint256 unitPrice
    ) public nonReentrant returns (uint256) {
        uint256 _orderId = makeCommonOrder(
            _msgSender(),
            tokenId,
            amount,
            unitPrice
        );
        emit MakeOrder(
            _msgSender(),
            _msgSender(),
            tokenId,
            amount,
            _orderId,
            unitPrice,
            _platformFeePercent
        );
        return _orderId;
    }

    function makeOrderProxy(
        address from,
        uint256 tokenId,
        uint256 amount,
        uint256 unitPrice
    ) public nonReentrant returns (uint256) {
        //生成订单ID,或者采用 Hash算法也是可以的
        uint256 _orderId = makeCommonOrder(from, tokenId, amount, unitPrice);
        emit MakeOrder(
            _msgSender(),
            from,
            tokenId,
            amount,
            _orderId,
            unitPrice,
            _platformFeePercent
        );
        return _orderId;
    }

    function makeCommonOrder(
        address from,
        uint256 tokenId,
        uint256 amount,
        uint256 unitPrice
    ) internal returns (uint256) {
        //生成订单ID,或者采用 Hash算法也是可以的

        require(
            _isBalanceAvailable(tokenId, amount),
            "ERC1155Sale: the balance is not enough."
        );

        _orderIds.increment();
        uint256 _orderId = _orderIds.current();
        _currentTradeId = _orderId;

        saleOrderList[_orderId] = Order(
            _orderId,
            from,
            tokenId,
            amount,
            _platformFeePercent,
            unitPrice,
            true
        );

        _presellBalances[tokenId][from] += amount; //累加预售

        super.setApprovalForAll(_market, true);

        return _orderId;
    }

    function cancleOrder(uint256 orderId) public nonReentrant {
        Order memory _order = saleOrderList[orderId];
        require(_order.isSelling, "ERC1155Sale: the order is inValidate.");

        address _from = _order.seller;

        require(
            _from == _msgSender() || isApprovedForAll(_from, _msgSender()),
            "ERC1155Sale: caller is not owner nor approved"
        );

        _removeSaleOrder(orderId);

        emit CancelOrder(
            _msgSender(),
            _order.seller,
            _order.tokenId,
            _order.stockAmount,
            orderId,
            saleOrderList[orderId].unitPrice,
            saleOrderList[orderId].feePercent
        );
    }

    function updateOrder(
        uint256 orderId,
        uint256 stockAmount,
        uint256 sellingUnitPrice
    ) public nonReentrant {
        Order memory _order = saleOrderList[orderId];
        require(_order.isSelling, "ERC1155Sale: the order is inValidate.");

        address _from = _order.seller;

        require(
            _from == _msgSender() || isApprovedForAll(_from, _msgSender()),
            "ERC1155Sale: caller is not owner nor approved"
        );

        saleOrderList[orderId].unitPrice = sellingUnitPrice;
        saleOrderList[orderId].stockAmount = stockAmount;

        emit UpdateOrder(
            _msgSender(),
            _order.seller,
            _order.tokenId,
            _order.stockAmount,
            orderId,
            sellingUnitPrice,
            saleOrderList[orderId].feePercent
        );
    }

    function getOrder(uint256 orderId_)
        public
        view
        returns (
            address seller,
            uint256 tokenId,
            uint256 amount,
            uint16 fee,
            uint256 sellingPrice
        )
    {
        Order memory _order = saleOrderList[orderId_];
        uint256 _tokenId = _order.tokenId;
        address _seller = _order.seller;
        uint256 _amount = _order.stockAmount;
        uint16 _feePercent = _order.feePercent;
        uint256 _sellingPrice = _order.unitPrice;
        return (_seller, _tokenId, _amount, _feePercent, _sellingPrice);
    }

    /**
     * 交易，转移nft
     */
    function tradeNft(address buyer_, uint256 orderId_)
        public
        isMinterRole
        returns (uint256)
    {
        //生成交易ID,或者采用 Hash算法也是可以的
        _tradeIds.increment();
        uint256 _tradeId = _tradeIds.current();
        _currentTradeId = _tradeId;

        Order memory _order = saleOrderList[orderId_];
        uint256 _tokenId = _order.tokenId;
        address _seller = _order.seller;
        uint256 _amount = _order.stockAmount;

        super.safeTransferFrom(_seller, buyer_, _tokenId, _amount, "");

        //状态更新
        SaleDetails memory _saleDetails = SaleDetails({
            tradeId: _tradeId,
            orderId: orderId_,
            purchasePrice: _order.unitPrice,
            buyer: buyer_,
            isSuccess: true
        });

        tradeDetails[_tradeId] = _saleDetails; //交易明细

        _removeSaleOrder(orderId_);

        emit TradeNft(buyer_, _seller, _tokenId, _amount, _tradeId, orderId_);

        return _tradeId;
    }

    function _removeSaleOrder(uint256 orderId) internal {
        Order memory _order = saleOrderList[orderId];
        uint256 _tokenId = _order.tokenId;
        address _seller = _order.seller;
        uint256 _amount = _order.stockAmount;

        _addPresellBalances(_tokenId, _seller, _amount);
        delete saleOrderList[orderId];
    }

    function _addPresellBalances(
        uint256 tokenId,
        address from,
        uint256 amount
    ) internal {
        _presellBalances[tokenId][from] -= amount;
        if (_presellBalances[tokenId][from] <= 0) {
            super.setApprovalForAll(_market, false);
        }
    }

    function _isBalanceAvailable(uint256 tokenId, uint256 amount)
        internal
        view
        returns (bool)
    {
        uint256 _presellAmount = _presellBalances[tokenId][_msgSender()];
        uint256 _maxBalance = balanceOf(_msgSender(), tokenId);
        return _maxBalance - _presellAmount >= amount ? true : false;
    }
}
