// SPDX-License-Identifier: MIT
pragma solidity >=0.4.22 <0.9.0;

import "./CellarCoinTokenSale.sol";
import "./ICellarNftSaleInterface.sol";

contract CellarMarketV2 is CellarCoinTokenSale {
    ICellarNftSaleInterface private _cellarNftSale; //ERC721-资产
    address private _organiser;// 平台抽成用户

    event Trade(address indexed buyer, address indexed seller, uint256 indexed tokenId,uint256 tradeId,uint256 orderId,uint256 price,uint16 feePercent,uint256 sellerIncome,uint256 organiserIncome);

    function initialize(ICellarNftSaleInterface  solarNftSale,string memory name,string memory symbol,uint8 decimal) initializer public {
       __ERC20_init(name, symbol);
       _init_base_unchanged(decimal);
       _init_market_unchanged(solarNftSale);
    }

    function _init_market_unchanged(ICellarNftSaleInterface  cellarNftSale)internal onlyInitializing{
        _cellarNftSale=cellarNftSale;
        _organiser=_msgSender(); //以后如有必要，需要额外指定

    }

    function setSpenderERC721(address spenderERC721)public isMinterRole{
        require(spenderERC721 !=address(0),"SolarTokenMarket: the spenderERC721Address must not be zero.");
        _cellarNftSale=ICellarNftSaleInterface(spenderERC721);
    }
    function getSpenderERC721()public view returns(address){
        return address(_cellarNftSale);
    }
    //买家购买入口
    function increaseAllowance(uint256 _tokenId,uint256 _price,uint256 _purchasePrice) public virtual returns (uint256) {
        require(_price >0,"SolarToken: the price must be greater then zero.");
        require(_purchasePrice >0,"SolarToken: the allowance amount  must be greater then zero");
        require(_purchasePrice >=_price ,"SolarToken: the allowance amount must not be less than price. ");
        require(allowanceOrders[_tokenId].buyer == address(0),"SolarToken: exists buyer before you .");

        bool isSuccess= _preorder( _tokenId, _price, _purchasePrice);
        require(isSuccess,"SolarToken:preOrder is failed.");
        uint256 tradeId= _tradeNFT(_tokenId);

        return tradeId;
    }

    function _tradeNFT(uint256 _tokenId) internal  returns(uint256){

        //检查订单情况
        address buyer = getAllowanceOrderBuyer(_tokenId);
        require( _msgSender() == buyer,"Market: current buyer is not the preOrder-owner.");
        uint256 orderPrice;
        uint256 purchasePrice;
        
        (orderPrice,purchasePrice)=getAllowanceOrderPrice(_tokenId);
        require(orderPrice > 0,"Market: the buyer selling price is not allowed zero.");
        require(purchasePrice > 0,"Market: the buyer confirmed price is not allowed zero.");

        address seller = _cellarNftSale.getNFTOwnerByTokenId(_tokenId);
        uint256 sellingPrice =_cellarNftSale.getOrderSellingPrice(_tokenId);
        require (orderPrice==sellingPrice,"Market: the selling price is changed.");

        uint16 feePercent= _cellarNftSale.getOrderFeePercent(_tokenId);
        
        uint256 confirmedPrice = sellingPrice*(100+feePercent)/100;
        require (confirmedPrice==purchasePrice,"Market: the purchase price is changed.");

       // 数据准备
        uint256 commision = (purchasePrice * feePercent) / 100;
        //分账
        bool isSuccess = transfer(seller, sellingPrice); //按照订单价，转给卖家
        require(isSuccess,"Market: fail to  transfer amount from buyer to seller");
        isSuccess = transfer(_organiser, commision);// 按照交易提成，给平台交易者
        require(isSuccess,"Market: fail to  transfer amount from buyer to organiser");
        removeAllowanceOrderItem(_tokenId);

        //nft 转移
        uint256 tradeId;
        uint256 orderId;
        
        (tradeId,orderId)=_cellarNftSale.transferNFT(buyer, _tokenId);

        //event ,交易日志
        emit Trade(buyer, seller, _tokenId,tradeId,orderId,sellingPrice, feePercent,sellingPrice - commision,commision);

        return tradeId;
    }
    
    function setOrganiser(address organiser)public isMinterRole{
        _organiser=organiser;
    }
    

    function getOrganiser()public view returns(address){
        return _organiser ;
    }


}