// SPDX-License-Identifier: MIT
pragma solidity >=0.4.22 <0.9.0;

import "@openzeppelin/contracts/token/ERC1155/ERC1155.sol";
import "./lib/AuthorityUpgradeable.sol";
import "./ERC1155URIStorageUpgradeable.sol";

/**
 * ERC1155 艺术展销中心
 */
contract ArtExhibitionERC1155 is
    AuthorityUpgradeable,
    ERC1155URIStorageUpgradeable
{
    mapping(uint256 => bool) private _hasMinted;

    function initialize(
        string memory name,
        string memory symbol,
        string memory uri
    ) public initializer {
        __ERC1155URIStorage_init(name, symbol, uri);
        _init_unchanged();
    }

    function _init_unchanged() internal onlyInitializing {
        __authorityUpgradeable_init();
    }

    function _checkedMinted(uint256 id) internal view returns (bool) {
        return _hasMinted[id];
    }

    function _checkedBatchMinted(uint256[] memory ids)
        internal
        view
        returns (bool)
    {
        for (uint256 i = 0; i < ids.length; ++i) {
            if (_hasMinted[ids[i]]) {
                return true;
            }
        }

        return false;
    }

    function supportsInterface(bytes4 interfaceId)
        public
        view
        virtual
        override(
            ERC1155URIStorageUpgradeable,
            AccessControlUpgradeable
        )
        returns (bool)
    {
        return
            interfaceId == type(IAccessControlUpgradeable).interfaceId ||
            interfaceId == type(IERC1155ExtUpgradeable).interfaceId ||
            super.supportsInterface(interfaceId);
    }

    /**
     *  设置这个方法，主要是为了部署，直接设置代持合约，
     */
    function setWalletProxyInit(address walletProxy) public onlyOwner {
        require(address(0) != walletProxy);
        _grantRole(MINTER_ROLE, walletProxy); //walletProxy，可以铸币
    }

    /**
     *  申请铸币
     */
    function mintSingleItem(
        address to,
        uint256 tokenId,
        uint256 amount,
        string memory tokenURI
    ) public onlyRole(MINTER_ROLE) {
        require(amount > 0, "ERC1155:minted amount must be greater than zero");

        _mintExt(to, tokenId, amount, tokenURI);

        emit MintSingleItem(_msgSender(), to, tokenId, amount, tokenURI);
    }

    function mintBatchItems(
        address to,
        uint256[] memory tokenIds,
        uint256[] memory amounts,
        string[] memory tokenUris
    ) public onlyRole(MINTER_ROLE) {
        _mintBatchExt(to, tokenIds, amounts, tokenUris);

        emit MintBatchItems(_msgSender(), to, tokenIds, amounts, tokenUris);
    }

    function _mintExt(
        address to,
        uint256 tokenId,
        uint256 amount,
        string memory uri
    ) internal virtual {
        require(!_checkedMinted(tokenId), "ERC1155: minted token is exists.");
        _mint(to, tokenId, amount, "");
        _setTokenURI(tokenId, uri);
        _hasMinted[tokenId] = true;
    }

    /**
     * @dev xref:ROOT:erc1155.adoc#batch-operations[Batched] version of {_mint}.
     *
     * Requirements:
     *
     * - `ids` and `amounts` must have the same length.
     * - If `to` refers to a smart contract, it must implement {IERC1155Receiver-onERC1155BatchReceived} and return the
     * acceptance magic value.
     */
    function _mintBatchExt(
        address to,
        uint256[] memory tokenIds,
        uint256[] memory amounts,
        string[] memory uris
    ) internal virtual {
        require(!_checkedBatchMinted(tokenIds), "ERC1155:token already minted.");

        _mintBatch(to, tokenIds, amounts, "");
        _setBatchTokenURI(tokenIds, uris);
        _setBatchMinted(tokenIds);
    }

    function _setBatchMinted(uint256[] memory tokenIds) internal virtual {
        for (uint256 i = 0; i < tokenIds.length; ++i) {
            _hasMinted[tokenIds[i]] = true;
        }
    }

    function gift(
        address from,
        address to,
        uint256 tokenId,
        uint256 amount
    ) public virtual {
        super.safeTransferFrom(from, to, tokenId, amount, "");

        emit Gift(_msgSender(), from, to, tokenId, amount);
    }

    function giftBatch(
        address from,
        address to,
        uint256[] memory tokenIds,
        uint256[] memory amounts
    ) public virtual {
        super.safeBatchTransferFrom(from, to, tokenIds, amounts, "");

        emit GiftBatch(_msgSender(), from, to, tokenIds, amounts);
    }
}
