// SPDX-License-Identifier: MIT
pragma solidity >=0.4.22 <0.9.0;

import "@openzeppelin/contracts-upgradeable/token/ERC1155/IERC1155ReceiverUpgradeable.sol";
import "./lib/AuthorityUpgradeable.sol";
import "./ERC1155URIStorageUpgradeable.sol";

contract WalletProxyUpgradeable is
    AuthorityUpgradeable,
    IERC1155ReceiverUpgradeable
{
    IERC1155ExtUpgradeable private _erc1155;

    string private _name; // Token name
    string private _symbol; // Token symbol
    bytes4 private _recRetval;
    bytes4 private _batRetval;

    mapping(uint256 => bool) private _hasMinted;

    event Received(
        address operator,
        address from,
        uint256 tokenId,
        uint256 value,
        bytes data,
        uint256 gas
    );
    event BatchReceived(
        address operator,
        address from,
        uint256[] tokenIds,
        uint256[] values,
        bytes data,
        uint256 gas
    );

    //Solidity 变量中 memory 、calldata 2 个表示作用非常类似，都是函数内部临时变量，它们最大的区别就是 calldata 是不可修改的，在某些只读的情况比较省 Gas.
    function initialize(IERC1155ExtUpgradeable erc1155,string calldata name,string calldata symbol) public initializer {
        _init_unchanged();
        _erc1155 = erc1155;
        _name=name;
        _symbol=symbol;
    }

    function _init_unchanged() internal onlyInitializing {
        __authorityUpgradeable_init();
         _recRetval = 0xf23a6e61;
        _batRetval = 0xbc197c81;
    }

    function mintSingleItem(
        uint256 tokenId,
        uint256 amount,
        string memory tokenURI
    ) public onlyRole(MINTER_ROLE) {
        _erc1155.mintSingleItem(address(this), tokenId, amount, tokenURI);
    }

    function mintBatchItems(
        uint256[] memory tokenIds,
        uint256[] memory amounts,
        string[] memory tokenUris
    ) public onlyRole(MINTER_ROLE) {
        _erc1155.mintBatchItems(address(this), tokenIds, amounts, tokenUris);
    }

    function gift(
        address to,
        uint256 tokenId,
        uint256 amount
    ) public virtual onlyRole(GIFT_ROLE) {
        _erc1155.gift(address(this), to, tokenId, amount);
    }

    function giftBatch(
        address to,
        uint256[] memory tokenIds,
        uint256[] memory amounts
    ) public virtual onlyRole(GIFT_ROLE) {
        _erc1155.giftBatch(address(this), to, tokenIds, amounts);
    }

    function onERC1155Received(
        address operator,
        address from,
        uint256 tokenId,
        uint256 value,
        bytes calldata data
    ) external override returns (bytes4) {
        emit Received(operator, from, tokenId, value, data, gasleft());
        return _recRetval;
        //return this.onERC1155Received.selector;
    }

    function onERC1155BatchReceived(
        address operator,
        address from,
        uint256[] calldata tokenIds,
        uint256[] calldata values,
        bytes calldata data
    ) external override returns (bytes4) {
        emit BatchReceived(operator, from, tokenIds, values, data, gasleft());
        return _batRetval;
        //return this.onERC1155BatchReceived.selector;
    }
}
