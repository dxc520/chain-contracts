// SPDX-License-Identifier: MIT
pragma solidity >=0.4.22 <0.9.0;

import "./BaseERC20.sol";

contract WalletERC20 is BaseERC20 {
   
    event MintTokens(
        address indexed _from,
        address indexed _buyer,
        uint256 indexed _amount
    );
    event TransferTokens(
        address indexed _from,
        address indexed _recipient,
        uint256 indexed _amount
    );

 function initialize(
        string memory name,
        string memory symbol,
        uint8 decimal
    ) public initializer {
        __ERC20_init(name, symbol);
        _init_base_unchanged(decimal);
    }


    /**
     * 向系统申请铸币
     */
    function mintTokens(address _recipient, uint256 _amount)
        public
        isMinterRole
    {
        require(address(0) != _recipient, "won't 0 address"); //严格上来说是否合约持有
        //todo _numberOfTokens 大小合理性限制

        _mint(_recipient, _amount);
 
        emit MintTokens(address(0), _recipient, _amount);
    }

    /**
     * proxy转移,需要approve
     */
    function transferTokens(
        address _from,
        address _recipient,
        uint256 _amount
    ) public returns (bool) {
        bool flag = super.transferFrom(_from, _recipient, _amount);

        emit TransferTokens(_from, _recipient, _amount);

        return flag;
    }
}
