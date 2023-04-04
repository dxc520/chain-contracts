// SPDX-License-Identifier: MIT
pragma solidity >=0.4.22 <0.9.0;

import "@openzeppelin/contracts-upgradeable/token/ERC20/ERC20Upgradeable.sol";
import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/access/AccessControlUpgradeable.sol";

contract CellarCoinBaseERC20 is OwnableUpgradeable,AccessControlUpgradeable,ERC20Upgradeable{

  bytes32 public constant MINTER_ROLE = keccak256("MINTER_ROLE"); //铸币角色
  uint8 _decimal;
    
  struct AllowanceDetails{
      //uint256 tokenId; //将来可以换成orderId
      address buyer; //买家地址
      //address spnder; //代币转移到销售合约这里是一个相对固定的值
      uint256 price;//订单价钱
      //uint16  feePercent;//费率
      uint256 purchasePrice;// 购买价
  }

    // 根据nft 确权转移ERC20
  mapping (uint256=>AllowanceDetails) allowanceOrders; // {tokenId =>AllowanceDetails}

  event PurchasePromise(address indexed _buyer,address _spender,uint256 indexed _tokenId,uint256 indexed _price,uint256  _purchasePrice);
  event TransferDirectSolar(address indexed _owner,address indexed _recipient,uint256 indexed _amount);

  modifier isMinterRole {
        require(
            hasRole(MINTER_ROLE, _msgSender()),
            "User must have minter role to mint"
        );
        _;
    }

    function _init_base_unchanged(uint8 decimal)internal onlyInitializing{
        _decimal=decimal;
        __Ownable_init();
        _setupRole(MINTER_ROLE, _msgSender());//owner，可以铸币
    }

    /**
     * 重新独立decimal,需要重写逻辑
     */
    function setDecimal(uint8 decimal ) public isMinterRole{
        require(decimal>=0 && decimal<19,"SolarCoinERC20:decimal must be in [0,18]");
        _decimal=decimal;
    }

    /**
     * 查询是否拥有铸币权限
     */
    function hasMintRole(address account) public view returns (bool) {
        return hasRole(MINTER_ROLE,account);
    }
    
    /**
     *  授予用户铸币权限
     */
    function grantMintRole(address operator)  public onlyOwner {
        require(address(0)!=operator);
        _setupRole(MINTER_ROLE, operator);//owner，可以铸币
        //grantRole(MINTER_ROLE, operator);//owner，可以铸币
    }

    /**
     *  收回用户铸币权限
     */
    function renounceMintRole( address account)public onlyOwner {
        renounceRole(MINTER_ROLE, account);
    }

    /**
     * 增加Access 控制，有权限才能mint 代币
     */
    function mint(address account, uint256 amount) public isMinterRole returns (bool) {
        _mint(account, amount);
       
      return true;
    }

    /**
     * 重写转移，主要是为了增加 event，方便后期定向日志筛选
     */
    function transfer(address recipient, uint256 amount) public virtual override returns (bool) {
        bool isSuccess = super.transfer(recipient,amount);
        emit TransferDirectSolar(_msgSender(),recipient,amount);
        return isSuccess;
    }

    /**
     * 买家预购
     */
    function _preorder(uint256 _tokenId,uint256 _price,uint256 _purchasePrice) internal returns (bool) {
         AllowanceDetails memory details = AllowanceDetails({
            buyer:_msgSender(),
            price:_price,
            purchasePrice:_purchasePrice
        });

        allowanceOrders[_tokenId]=details;
        //event 
        emit PurchasePromise(_msgSender(),address(this),  _tokenId,  _price, _purchasePrice);
       return true;
    }

    /**
     * 获取授权订单价格
     */
    function getAllowanceOrderPrice(uint256 _tokenId) public view  returns(uint256,uint256){
        return  (allowanceOrders[_tokenId].price,allowanceOrders[_tokenId].purchasePrice);
    }

    function getAllowanceOrderBuyer(uint256 _tokenId) public view  returns(address){
        return  allowanceOrders[_tokenId].buyer;
    }
    /**
     * 交易成功后需要清除
     */
    function removeAllowanceOrderItem(uint256 _tokenId) public {
        require( allowanceOrders[_tokenId].buyer == _msgSender() || hasMintRole(_msgSender())  ,"SolarCoinToken: You are not be the confirmed buyer of the allowanceOrder" );
        delete allowanceOrders[_tokenId];
    }

}