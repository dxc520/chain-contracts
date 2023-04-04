// SPDX-License-Identifier: MIT
pragma solidity >=0.4.22 <0.9.0;

import "@openzeppelin/contracts-upgradeable/security/ReentrancyGuardUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/utils/CountersUpgradeable.sol";
import "@openzeppelin/contracts/utils/math/SafeMath.sol";
import "@openzeppelin/contracts/utils/Strings.sol";
import "./CellarBaseNftERC1155.sol";

/**
 *  这里的 ERC1155  交易，挂单后不允许再次挂单，单可以修改；即不允许同时挂单N个。
 *   那么简化操作： 取消挂单等一系列的操作，可以用tokenId取消，
 *
 */

contract CellarNft1155Sale is CellarBaseNftERC1155, ReentrancyGuardUpgradeable {
    using CountersUpgradeable for CountersUpgradeable.Counter;
    CountersUpgradeable.Counter private _orderIds; //订单序列号,如果不要后期可以省略
    CountersUpgradeable.Counter private _tradeIds; //订单序列号,如果不要后期可以省略

    address private _market; //市场地址
    address private _organiser; //交易费获得者地址
    uint16 _platformFeePercent; //平台费率

    //挂单销售
    struct Order {
        uint256 orderId; //订单ID
        address seller; //谁卖,
        uint256 tokenId; // tokenId
        uint256 stockAmount; //库存数量，当前最大售卖量
        uint16 feePercent; //交易费提成比例
        uint256 unitPrice; //销售单价
        uint256 sellingTotalPrice; //出售总价
        //uint time;//挂单时间
        bool isSelling; //是否处于挂单状态
    }

    //交易信息
    struct SaleDetails {
        uint256 tradeId; //交易序号
        uint256 orderId; //订单ID
        // Order  saleOrder;//订单Id;如果不考虑空间问题，可以转成对象 Oder //这里更改，考虑到实时销售订单完毕后就删除了，转至这里保存
        uint256 purchasePrice; //购买价格,这里反应的是订单价格(总价)
        address buyer; // 购买者
        //uint time ;//购买时间
        bool isSuccess; //交易完成
    }

    //所有被交易的
    mapping(uint256 => mapping(address => uint256)) private _presellBalances; //{tokenId=>{owner=>amount}}由于存在挂单，就不是实际转移，方便后期超售管理，增加此状态
    mapping(uint256 => mapping(address => uint256)) private _nftForSelling; //{tokenId=>{owner=>orderID}}正在热销的nft:被销售的tokenId 列表,记录当前在售的nft
    // 挂单销售列表
    //mapping (uint256=>Order) saleOrders; // 销售订单{tokenId=>order}，值记录当前的在销售状态

    //这里的销售订单，存在假设前提：就是ERC1155的同一token不支持同段时间多次未完成的挂单。否则，这里的数据结构需要调整。
    mapping(uint256 => Order) _saleOrderList; //销售订单(){orderId=>order}
    //mapping(uint256=>uint256) sellingNFTSaleOrders; // 当前销售订单{tokenId=>orderID}，值记录当前的在销售状态
    //mapping (uint256=>mapping(address=>uint256[])) sellingNFTSaleOrders; // 当前销售订单{tokenId=>{owner=>orderID[]}}，值记录当前的在销售状态

    // 交易成功列表;
    mapping(uint256 => SaleDetails) _tradeDetails; //交易明细{tradeId=>SaleDetails}

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
        uint256 totalPrice,
        uint16 feePercent
    );
    event UpdateOrder(
        address _operator,
        address indexed _seller,
        uint256 indexed _tokenId,
        uint256 _amount,
        uint256 indexed _orderId,
        uint256 _unitPrice,
        uint256 _totalPrice,
        uint16 _feePercent
    );
    event CancelOrder(
        address _operator,
        address indexed _seller,
        uint256 indexed _tokenId,
        uint256 _amount,
        uint256 indexed _orderId,
        uint256 _unitPrice,
        uint256 _totalPrice,
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
        uint256 totalPrice
    ) public nonReentrant returns (uint256) {
        return
            makeCommonOrder(
                _msgSender(),
                _msgSender(),
                tokenId,
                amount,
                totalPrice
            );
    }

    function makeCommonOrder(
        address operator,
        address from,
        uint256 tokenId,
        uint256 amount,
        uint256 totalPrice
    ) internal returns (uint256 orderId) {
        //生成订单ID,或者采用 Hash算法也是可以的

        require(
            _isBalanceAvailable(tokenId, amount),
            "ERC1155Sale: the balance is not enough."
        );

        uint256 _orderId = getNextOrderId();

        _saleOrderList[_orderId] = Order(
            _orderId,
            from,
            tokenId,
            amount,
            _platformFeePercent,
            0,
            totalPrice,
            true
        );

        _presellBalances[tokenId][from] += amount; //累加预售

        super.setApprovalForAll(_market, true);

        _nftForSelling[tokenId][from] = _orderId;

        emit MakeOrder(
            operator,
            from,
            tokenId,
            amount,
            _orderId,
            0,
            totalPrice,
            _platformFeePercent
        );

        return _orderId;
    }

    function getNextOrderId() internal returns (uint256 orderId) {
        _orderIds.increment();
        uint256 _orderId = _orderIds.current();
        bytes32 result = keccak256(abi.encodePacked("ERC1155", _orderId));
        orderId = uint256(result);
    }

    function cancelOrder(uint256 orderId) public nonReentrant {
        Order memory _order = _saleOrderList[orderId];
        require(_order.isSelling, "ERC1155Sale: the order is inValidate.");

        address _from = _order.seller;

        require(
            _from == _msgSender() || isApprovedForAll(_from, _msgSender()),
            "ERC1155Sale: caller is not owner nor approved"
        );

        _removeSaleOrder(orderId, false);

        emit CancelOrder(
            _msgSender(),
            _order.seller,
            _order.tokenId,
            _order.stockAmount,
            orderId,
            _saleOrderList[orderId].unitPrice,
            _saleOrderList[orderId].sellingTotalPrice,
            _saleOrderList[orderId].feePercent
        );
    }

    function updateOrder(
        uint256 orderId,
        uint256 stockAmount,
        uint256 totalPrice
    ) public nonReentrant {
        Order memory _order = _saleOrderList[orderId];
        require(_order.isSelling, "ERC1155Sale: the order is inValidate.");

        address _from = _order.seller;

        require(
            _from == _msgSender() || isApprovedForAll(_from, _msgSender()),
            "ERC1155Sale: caller is not owner nor approved"
        );

        //_saleOrderList[orderId].unitPrice = sellingUnitPrice;
        _saleOrderList[orderId].stockAmount = stockAmount;
        _saleOrderList[orderId].sellingTotalPrice = totalPrice;

        emit UpdateOrder(
            _msgSender(),
            _order.seller,
            _order.tokenId,
            _order.stockAmount,
            orderId,
            0,
            totalPrice,
            _order.feePercent
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
            uint256 totalPrice
        )
    {
        Order memory _order = _saleOrderList[orderId_];

        return (
            _order.seller,
            _order.tokenId,
            _order.stockAmount,
            _order.feePercent,
            _order.sellingTotalPrice
        );
    }

    function getOrderSeller(uint256 orderId_)
        public
        view
        returns (address seller)
    {
        Order memory _order = _saleOrderList[orderId_];

        return _order.seller;
    }

    function tradeNft(
        address buyer,
        uint256 tokenId,
        address from
    ) public isMinterRole returns (uint256) {
        uint256 _orderId = _nftForSelling[tokenId][from];
        require(_orderId > 0, "ERC1155Sale:found no order in trade.");
        return tradeNftByOrder(buyer, _orderId);
    }

    /**
     * 交易，转移nft
     */
    function tradeNftByOrder(address buyer_, uint256 orderId_)
        internal
        returns (uint256)
    {
        //生成交易ID,或者采用 Hash算法也是可以的
        _tradeIds.increment();
        uint256 _tradeId = _tradeIds.current();

        Order memory _order = _saleOrderList[orderId_];
        uint256 _tokenId = _order.tokenId;
        address _seller = _order.seller;
        uint256 _amount = _order.stockAmount;

        super.safeTransferFrom(_seller, buyer_, _tokenId, _amount, "");

        //状态更新
        SaleDetails memory _saleDetails = SaleDetails({
            tradeId: _tradeId,
            orderId: orderId_,
            purchasePrice: _order.sellingTotalPrice,
            buyer: buyer_,
            isSuccess: true
        });

        _tradeDetails[_tradeId] = _saleDetails; //交易明细

        _removeSaleOrder(orderId_, true);

        emit TradeNft(buyer_, _seller, _tokenId, _amount, _tradeId, orderId_);

        return _tradeId;
    }

    function _removeSaleOrder(uint256 orderId, bool isSaleRetained) internal {
        Order memory _order = _saleOrderList[orderId];
        uint256 _tokenId = _order.tokenId;
        address _seller = _order.seller;
        uint256 _amount = _order.stockAmount;

        _reducePresellBalances(_tokenId, _seller, _amount);
        if (!isSaleRetained) {
            delete _saleOrderList[orderId];
        }
        delete _nftForSelling[_tokenId][_seller];
    }

    function _reducePresellBalances(
        uint256 tokenId,
        address from,
        uint256 amount
    ) internal {
        _presellBalances[tokenId][from] -= amount;
        /*if (_presellBalances[tokenId][from] <= 0) {
            super.setApprovalForAll(_market, false);
        }
        */
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

    function getOrderId(uint256 tokenId,address seller) public view returns (uint256) {
        return _nftForSelling[tokenId][seller];
    }
}
