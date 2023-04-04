// SPDX-License-Identifier: MIT
pragma solidity >=0.4.22 <0.9.0;

import "./CellarCoinTokenSale.sol";
import "./ICellarNftSaleInterface.sol";
import "./ICellarNft1155SaleInterface.sol";
import "@openzeppelin/contracts/utils/math/SafeMath.sol";
import "@openzeppelin/contracts/utils/Strings.sol";

contract CellarMarket is CellarCoinTokenSale {
    ICellarNftSaleInterface private _cellarNftSale; //ERC721-资产
    address private _organiser; // 平台抽成用户
    ICellarNft1155SaleInterface private _cellarNft1155Sale; //ERC1155-资产

    struct SaleOrder {
        uint256 orderId; // tokenId
        address seller; //谁卖,
        uint256 tokenId; // tokenId
        uint256 stockAmount; //库存数量，当前最大售卖量
        uint16 feePercent; //交易费提成比例
        uint256 totalPrice; //出售总价
    }

    struct AllowanceDetails {
        //uint256 tokenId; //将来可以换成orderId
        address buyer; //买家地址
        //address spnder; //代币转移到销售合约这里是一个相对固定的值
        uint256 price; //订单价钱
        //uint16  feePercent;//费率
        uint256 purchasePrice; // 购买价
    }

    // 根据nft 确权转移ERC20
    mapping(uint256 => AllowanceDetails) allowance721Orders; // {orderId =>AllowanceDetails}
    mapping(uint256 => AllowanceDetails) allowance1155Orders; // {orderId =>AllowanceDetails}

    event PurchasePromise(
        address indexed _buyer,
        address _seller,
        uint256 indexed _tokenId,
        uint256 _saleOrderId,
        uint256 _price,
        uint256 indexed _purchasePrice,
        uint256 _amount,
        string _topic
    );

    event Trade(
        address indexed buyer,
        address indexed seller,
        uint256 indexed tokenId,
        uint256 tradeId,
        uint256 orderId,
        uint256 price,
        uint16 feePercent,
        uint256 sellerIncome,
        uint256 organiserIncome,
        uint256 amount,
        string topic
    );

    function initialize(
        ICellarNftSaleInterface solarNftSale,
        ICellarNft1155SaleInterface cellarNft1155Sale,
        string memory name,
        string memory symbol,
        uint8 decimal
    ) public initializer {
        __ERC20_init(name, symbol);
        _init_base_unchanged(decimal);
        _init_market_unchanged(solarNftSale, cellarNft1155Sale);
    }

    function _init_market_unchanged(
        ICellarNftSaleInterface cellarNftSale,
        ICellarNft1155SaleInterface cellarNft1155Sale
    ) internal onlyInitializing {
        _cellarNftSale = cellarNftSale;
        _cellarNft1155Sale = cellarNft1155Sale;
        _organiser = _msgSender(); //以后如有必要，需要额外指定
    }

    function setSpenderERC721(address spenderERC721) public isMinterRole {
        require(
            spenderERC721 != address(0),
            "Market: the spenderERC721Address must not be zero."
        );
        _cellarNftSale = ICellarNftSaleInterface(spenderERC721);
    }

    function setSpenderERC1155(address spenderERC1155) public isMinterRole {
        require(
            spenderERC1155 != address(0),
            "Market: the spenderERC1155Address must not be zero."
        );
        _cellarNft1155Sale = ICellarNft1155SaleInterface(spenderERC1155);
    }

    function setOrganiser(address organiser) public isMinterRole {
        _organiser = organiser;
    }

    function getSpenderERC721() public view returns (address) {
        return address(_cellarNftSale);
    }

    function getSpenderERC1155() public view returns (address) {
        return address(_cellarNft1155Sale);
    }

    function getOrganiser() public view returns (address) {
        return _organiser;
    }

    /**
     * 买家预购
     */
    function _preorder(
        address seller,
        uint256 _orderId,
        uint256 _tokenId,
        uint256 amount,
        uint256 _price,
        uint256 _purchasePrice,
        string memory _topic
    ) internal returns (bool) {
        AllowanceDetails memory details = AllowanceDetails({
            buyer: _msgSender(),
            price: _price,
            purchasePrice: _purchasePrice
        });

        if (
            keccak256(abi.encodePacked(_topic)) ==
            keccak256(abi.encodePacked("ERC721"))
        ) {
            allowance721Orders[_orderId] = details;
        } else if (
            keccak256(abi.encodePacked(_topic)) ==
            keccak256(abi.encodePacked("ERC1155"))
        ) {
            allowance1155Orders[_orderId] = details;
        }

        //event
        emit PurchasePromise(
            _msgSender(),
            seller,
            _tokenId,
            _orderId,
            _price,
            _purchasePrice,
            amount,
            _topic
        );
        return true;
    }

    //买家购买入口
    function increase721Allowance(uint256 orderId, uint256 purchasePrice)
        public
        virtual
        returns (uint256)
    {
        require(
            purchasePrice > 0,
            "market: the erc721-purchasePrice  must be greater then zero"
        );
        require(
            allowance721Orders[orderId].buyer == address(0) ||
                allowance721Orders[orderId].buyer == _msgSender(),
            string(
                abi.encodePacked(
                    "market: exists ERC721 buyer before you .The buyer is",
                    allowance721Orders[orderId].buyer
                )
            )
        );

        SaleOrder memory _saleOrder;
        _saleOrder.orderId = orderId;

        //uint256 tokenId,address seller,uint256 sellingPrice,uint16  feePercent
        (
            _saleOrder.tokenId,
            _saleOrder.seller,
            _saleOrder.totalPrice,
            _saleOrder.feePercent
        ) = _cellarNftSale.getSellingOrderByOrderId(orderId);
        require(
            purchasePrice >= _saleOrder.totalPrice,
            string(
                abi.encodePacked(
                    "market: the ERC721 purchasePrice must be greater than totalPrice. [",
                    purchasePrice,
                    ",",
                    _saleOrder.totalPrice,
                    "]"
                )
            )
        );
        require(
            _saleOrder.seller != address(0),
            string(
                abi.encodePacked(
                    "market:found no erc721 order in the trade. orderId[",
                    orderId,
                    "]"
                )
            )
        );

        //address seller,uint256 _orderId,uint256 _tokenId,uint256 amount,uint256 _price,uint256 _purchasePrice,string memory _topic
        bool isSuccess = _preorder(
            _saleOrder.seller,
            orderId,
            _saleOrder.tokenId,
            1,
            _saleOrder.totalPrice,
            purchasePrice,
            "ERC721"
        );
        require(isSuccess, "market:preOrder is failed.");
        uint256 tradeId = _tradeNFT721(_saleOrder);

        return tradeId;
    }

    function _tradeNFT721(SaleOrder memory saleOrder)
        internal
        returns (uint256)
    {
        //检查订单情况
        address buyer = allowance721Orders[saleOrder.orderId].buyer;

        require(
            _msgSender() == buyer,
            "Market: current buyer is not the preOrder-owner."
        );

        (bool isAddedSuccess, uint256 confirmedPrice) = SafeMath.tryAdd(
            saleOrder.totalPrice,
            (saleOrder.totalPrice * saleOrder.feePercent) / 100
        );
        require(
            isAddedSuccess,
            "Market:  The amount of ERC721-trade is overflowed"
        );

        require(
            confirmedPrice ==
                allowance721Orders[saleOrder.orderId].purchasePrice,
            string(
                abi.encodePacked(
                    "Market: the erc721-purchase price is changed.[",
                    confirmedPrice,
                    ",",
                    allowance721Orders[saleOrder.orderId].purchasePrice,
                    "]"
                )
            )
        );

        //分账
        bool isSuccess = transfer(saleOrder.seller, saleOrder.totalPrice); //按照订单价，转给卖家
        require(
            isSuccess,
            "Market: fail to  transfer amount from buyer to seller"
        );
        isSuccess = transfer(_organiser, confirmedPrice - saleOrder.totalPrice); // 按照交易提成，给平台交易者
        require(
            isSuccess,
            "Market: fail to  transfer amount from buyer to organiser"
        );
        removeAllowanceOrderItem(saleOrder.orderId, "ERC721");

        //nft 转移
        uint256 tradeId;
        (tradeId, ) = _cellarNftSale.transferNFT(buyer, saleOrder.tokenId);

        //event ,交易日志
        emit Trade(
            buyer,
            saleOrder.seller,
            saleOrder.tokenId,
            tradeId,
            saleOrder.orderId,
            confirmedPrice,
            saleOrder.feePercent,
            saleOrder.totalPrice,
            confirmedPrice - saleOrder.totalPrice,
            1,
            "ERC721"
        );

        return tradeId;
    }

    /**
     * 交易成功后需要清除
     */
    function removeAllowanceOrderItem(uint256 orderId, string memory _topic)
        public
    {
        if (
            keccak256(abi.encodePacked(_topic)) ==
            keccak256(abi.encodePacked("ERC721"))
        ) {
            // require(
            //     allowance721Orders[orderId].buyer == _msgSender() || hasMintRole(_msgSender()),
            //     "SolarCoinToken: You are not be the confirmed buyer of the allowance721Order"
            // );
            delete allowance721Orders[orderId];
        } else if (
            keccak256(abi.encodePacked(_topic)) ==
            keccak256(abi.encodePacked("ERC1155"))
        ) {
            // require(
            //     allowance1155Orders[orderId].buyer == _msgSender() || hasMintRole(_msgSender()),
            //     "SolarCoinToken: You are not be the confirmed buyer of the allowance1155Order"
            // );
            delete allowance1155Orders[orderId];
        }
    }

    //买家购买入口
    function increase1155Allowance(uint256 orderId, uint256 purchasePrice)
        public
        virtual
        returns (uint256)
    {
        require(
            purchasePrice > 0,
            "market: the erc1155-purchasePrice must be greater then zero"
        );

        require(
            allowance1155Orders[orderId].buyer == address(0) ||
                allowance1155Orders[orderId].buyer == _msgSender(),
            string(
                abi.encodePacked(
                    "market: exists ERC1155 buyer before you .the buyer is ",
                    allowance1155Orders[orderId].buyer
                )
            )
        );

        SaleOrder memory _saleOrder;

        (
            _saleOrder.seller,
            _saleOrder.tokenId,
            _saleOrder.stockAmount,
            _saleOrder.feePercent,
            _saleOrder.totalPrice
        ) = _cellarNft1155Sale.getOrder(orderId);
        require(
            purchasePrice >= _saleOrder.totalPrice,
            string(
                abi.encodePacked(
                    "market: the ERC1155 purchasePrice must be greater than totalPrice. [",
                    purchasePrice,
                    ",",
                    _saleOrder.totalPrice,
                    "]"
                )
            )
        );

        require(
            _saleOrder.seller != address(0),
            string(
                abi.encodePacked(
                    "market:found no erc1155 order in the trade. orderId[",
                    orderId,
                    "]"
                )
            )
        );
        //require(_saleOrder.seller==address(0),Strings.toHexString(uint256(uint160(_saleOrder.seller)), 20) );

        //address seller,uint256 _orderId,uint256 _tokenId,uint256 amount,uint256 _price,uint256 _purchasePrice,string memory _topic
        bool isSuccess = _preorder(
            _saleOrder.seller,
            orderId,
            _saleOrder.tokenId,
            _saleOrder.stockAmount,
            _saleOrder.totalPrice,
            purchasePrice,
            "ERC1155"
        );
        require(isSuccess, "market:erc1155-preOrder is failed.");

        uint256 tradeId = _tradeNFT1155(orderId, purchasePrice, _saleOrder);

        return tradeId;
    }

    function _tradeNFT1155(
        uint256 orderId,
        uint256 purchasePrice,
        SaleOrder memory _saleOrder
    ) internal returns (uint256) {
        (bool isAddedSuccess, uint256 confirmedPrice) = SafeMath.tryAdd(
            _saleOrder.totalPrice,
            (_saleOrder.totalPrice * _saleOrder.feePercent) / 100
        );

        require(
            isAddedSuccess,
            "Market:  The amount of ERC1155-trade is overflowed"
        );

        require(
            confirmedPrice == purchasePrice,
            string(
                abi.encodePacked(
                    "Market: the erc1155 purchase price is changed:[",
                    confirmedPrice,
                    ",",
                    purchasePrice,
                    "]"
                )
            )
        );

        // 数据准备
        uint256 commision = purchasePrice - _saleOrder.totalPrice;
        //分账
        bool isSuccess = transfer(_saleOrder.seller, _saleOrder.totalPrice); //按照订单价，转给卖家
        require(
            isSuccess,
            "Market: fail to  transfer amount from buyer to seller"
        );
        isSuccess = transfer(_organiser, commision); // 按照交易提成，给平台交易者
        require(
            isSuccess,
            "Market: fail to  transfer amount from buyer to organiser"
        );
        removeAllowanceOrderItem(_saleOrder.orderId, "ERC1155");

        //nft 转移
        uint256 tradeId = _cellarNft1155Sale.tradeNft(
            _msgSender(),
            _saleOrder.tokenId,
            _saleOrder.seller
        );

        //event ,交易日志
        emit Trade(
            _msgSender(),
            _saleOrder.seller,
            _saleOrder.tokenId,
            tradeId,
            orderId,
            confirmedPrice,
            _saleOrder.feePercent,
            _saleOrder.totalPrice,
            commision,
            0,
            "ERC1155"
        );

        return tradeId;
    }
}
