// SPDX-License-Identifier: MIT
pragma solidity >=0.4.22 <0.9.0;

import "@openzeppelin/contracts-upgradeable/token/ERC20/ERC20Upgradeable.sol";
import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/access/AccessControlUpgradeable.sol";

contract CellarCoinBaseERC20 is
    OwnableUpgradeable,
    AccessControlUpgradeable,
    ERC20Upgradeable
{
    bytes32 public constant MINTER_ROLE = keccak256("MINTER_ROLE"); //铸币角色
    uint8 _decimal;

    event TransferDirectSolar(
        address indexed _owner,
        address indexed _recipient,
        uint256 indexed _amount
    );

    modifier isMinterRole() {
        require(
            hasRole(MINTER_ROLE, _msgSender()),
            "User must have minter role to mint"
        );
        _;
    }

    function _init_base_unchanged(uint8 decimal) internal onlyInitializing {
        _decimal = decimal;
        __Ownable_init();
        _setupRole(MINTER_ROLE, _msgSender()); //owner，可以铸币
    }

    /**
     * 重新独立decimal,需要重写逻辑
     */
    function setDecimal(uint8 decimal) public isMinterRole {
        require(
            decimal >= 0 && decimal < 19,
            "SolarCoinERC20:decimal must be in [0,18]"
        );
        _decimal = decimal;
    }

    function decimals() public view virtual override returns (uint8) {
        return _decimal;
    }

    /**
     * 查询是否拥有铸币权限
     */
    function hasMintRole(address account) public view returns (bool) {
        return hasRole(MINTER_ROLE, account);
    }

    /**
     *  授予用户铸币权限
     */
    function grantMintRole(address operator) public onlyOwner {
        require(address(0) != operator);
        _setupRole(MINTER_ROLE, operator); //owner，可以铸币
        //grantRole(MINTER_ROLE, operator);//owner，可以铸币
    }

    /**
     *  收回用户铸币权限
     */
    function renounceMintRole(address account) public onlyOwner {
        renounceRole(MINTER_ROLE, account);
    }

    /**
     * 增加Access 控制，有权限才能mint 代币
     */
    function mint(address account, uint256 amount)
        public
        isMinterRole
        returns (bool)
    {
        _mint(account, amount);

        return true;
    }

    /**
     * 重写转移，主要是为了增加 event，方便后期定向日志筛选
     */
    function transfer(address recipient, uint256 amount)
        public
        virtual
        override
        returns (bool)
    {
        bool isSuccess = super.transfer(recipient, amount);
        emit TransferDirectSolar(_msgSender(), recipient, amount);
        return isSuccess;
    }
}
