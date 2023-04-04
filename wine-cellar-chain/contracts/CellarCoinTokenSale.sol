// SPDX-License-Identifier: MIT
pragma solidity >=0.4.22 <0.9.0;

import "./CellarCoinBaseERC20.sol";

contract CellarCoinTokenSale is CellarCoinBaseERC20 {
    /**
     * 铸币交易信息
     */
    struct DealInfo {
        address account; //铸币账户owner
        uint256 amount; //铸币数量
        uint256 datetime; //铸币时间
    }

    DealInfo[] private dealInfoList; //记录铸币交易信息，这里没有考虑增发

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
        dealInfoList.push(DealInfo(_recipient, _amount, block.timestamp)); //记录铸币的交易日志

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
