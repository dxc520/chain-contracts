// SPDX-License-Identifier: MIT
pragma solidity >=0.4.22 <0.9.0;

import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/access/AccessControlUpgradeable.sol";

/**
 * 权限控制基类
 */
contract AuthorityUpgradeable is OwnableUpgradeable, AccessControlUpgradeable {
    enum AuthorityRole {
        None,
        Admin_Role,
        Manager_Role,
        Mint_Role,
        Gift_role
    } //枚举的范围 [0,256),即uint8类型；返回值从0 开始，为了避免解析event获取值混淆，None=0，占位而已；
    bytes32 internal constant ADMIN_ROLE = keccak256("ADMINT_ROLE"); //只能被owner管理
    bytes32 internal constant MANAGER_ROLE = keccak256("MANAGER_ROLE"); //只能有Admin角色管控
    bytes32 internal constant MINTER_ROLE = keccak256("MINTER_ROLE"); //只能由manager角色管控
    bytes32 internal constant GIFT_ROLE = keccak256("GIFT_ROLE"); //只能由manager角色管控

    mapping(uint256 => bool) private _hasMinted;

    event GrantRole(
        address contratAdress,
        address indexed operator,
        address indexed account,
        AuthorityRole indexed role
    );
    event RevokeRole(
        address contratAdress,
        address indexed operator,
        address indexed account,
        AuthorityRole indexed role
    );

    function __authorityUpgradeable_init() internal onlyInitializing {
        _authorityUpgradeable_init_unchanged();
    }

    function _authorityUpgradeable_init_unchanged() internal onlyInitializing {
        __Ownable_init();
        //_setupRole(MINTER_ROLE, _msgSender()); //owner，可以铸币
        _setRoleAdmin(MANAGER_ROLE, ADMIN_ROLE); //管理者role 由admin 管理
        _setRoleAdmin(MINTER_ROLE, MANAGER_ROLE); //铸币role 由 manager 管理
        _setRoleAdmin(GIFT_ROLE, MANAGER_ROLE); //赠予role 由 manager 管理
    }

    /**
     * 查询是否拥有铸币权限
     */
    function hasAdminRole(address account) public view returns (bool) {
        return hasRole(ADMIN_ROLE, account);
    }

    function hasManagerRole(address account) public view returns (bool) {
        return hasRole(MANAGER_ROLE, account);
    }

    function hasMintRole(address account) public view returns (bool) {
        return hasRole(MINTER_ROLE, account);
    }

    function hasGiftRole(address account) public view returns (bool) {
        return hasRole(GIFT_ROLE, account);
    }

    /**
     *  SuperAdmin管理
     */
    function grantAdminRole(address account) public onlyOwner {
        require(address(0) != account);
        _grantRole(ADMIN_ROLE, account); //owner，可以铸币
        emit GrantRole(
            address(this),
            _msgSender(),
            account,
            AuthorityRole.Admin_Role
        );
    }

    function revokeAdminRole(address account) public onlyOwner {
        _revokeRole(ADMIN_ROLE, account);
        emit RevokeRole(
            address(this),
            _msgSender(),
            account,
            AuthorityRole.Admin_Role
        );
    }

    /**
     *  Manager 权限管理
     */
    function grantManagerRole(address account) public onlyRole(ADMIN_ROLE) {
        require(address(0) != account);
        _grantRole(MANAGER_ROLE, account); //owner，可以铸币
        emit GrantRole(
            address(this),
            _msgSender(),
            account,
            AuthorityRole.Manager_Role
        );
    }

    function revokeManagerRole(address account) public onlyRole(ADMIN_ROLE) {
        _revokeRole(MANAGER_ROLE, account);
        emit RevokeRole(
            address(this),
            _msgSender(),
            account,
            AuthorityRole.Manager_Role
        );
    }

    /**
     *  MintRole 权限管理
     */
    function grantMintRole(address account) public onlyRole(MANAGER_ROLE) {
        require(address(0) != account);
        _grantRole(MINTER_ROLE, account); //owner，可以铸币
        emit GrantRole(
            address(this),
            _msgSender(),
            account,
            AuthorityRole.Mint_Role
        );
    }

    function revokeMintRole(address account) public onlyRole(MANAGER_ROLE) {
        _revokeRole(MINTER_ROLE, account);
        emit RevokeRole(
            address(this),
            _msgSender(),
            account,
            AuthorityRole.Mint_Role
        );
    }

    /**
     *  MintRole 权限管理
     */
    function grantGiftRole(address account) public onlyRole(MANAGER_ROLE) {
        require(address(0) != account);
        _grantRole(GIFT_ROLE, account); //owner，可以铸币
        emit GrantRole(
            address(this),
            _msgSender(),
            account,
            AuthorityRole.Gift_role
        );
    }

    function revokeGiftRole(address account) public onlyRole(MANAGER_ROLE) {
        _revokeRole(GIFT_ROLE, account);
        emit RevokeRole(
            address(this),
            _msgSender(),
            account,
            AuthorityRole.Gift_role
        );
    }
}
