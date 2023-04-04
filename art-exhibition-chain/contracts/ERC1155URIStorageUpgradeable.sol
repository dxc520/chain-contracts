// SPDX-License-Identifier: MIT
pragma solidity >=0.4.22 <0.9.0;

import "@openzeppelin/contracts-upgradeable/token/ERC1155/ERC1155Upgradeable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "./IERC1155ExtUpgradeable.sol";

abstract contract ERC1155URIStorageUpgradeable is
    Initializable,
    ERC1155Upgradeable,IERC1155ExtUpgradeable
{
    // Token name
    string private _name;

    // Token symbol
    string private _symbol;

    // Optional mapping for token URIs
    mapping(uint256 => string) private _tokenURIs;


    function __ERC1155URIStorage_init(
        string memory name_,
        string memory symbol_,
        string memory uri_
    ) internal onlyInitializing {
        __ERC1155_init(uri_);
        __ERC1155__ERC1155URIStorage_init_unchained(name_, symbol_);
    }

    function __ERC1155__ERC1155URIStorage_init_unchained(
        string memory name_,
        string memory symbol_
    ) internal onlyInitializing {
        _name = name_;
        _symbol = symbol_;
    }

    /**
     * @dev See {IERC721Metadata-name}.
     */
    function name() public view virtual returns (string memory) {
        return _name;
    }

    /**
     * @dev See {IERC721Metadata-symbol}.
     */
    function symbol() public view virtual returns (string memory) {
        return _symbol;
    }

    /**
     * @dev See {IERC721Metadata-tokenURI}.
     */
    function tokenURI(uint256 tokenId)
        public
        view
        virtual
        returns (string memory)
    {
        //todo 校验 token 存在

        string memory _tokenURI = _tokenURIs[tokenId];
        string memory base = _baseURI();

        // If there is no base URI, return the token URI.
        if (bytes(base).length == 0) {
            return _tokenURI;
        }
        // If both are set, concatenate the baseURI and tokenURI (via abi.encodePacked).
        if (bytes(_tokenURI).length > 0) {
            return string(abi.encodePacked(base, _tokenURI));
        }

        return "undefined";
    }

    function _baseURI() internal view virtual returns (string memory) {
        return "";
    }

    /**
     * @dev Sets `_tokenURI` as the tokenURI of `tokenId`.
     *
     * Requirements:
     *
     * - `tokenId` must exist.
     */
    function _setTokenURI(uint256 tokenId, string memory uri_)
        internal
        virtual
    {
        _tokenURIs[tokenId] = uri_;
    }

    function _setBatchTokenURI(uint256[] memory tokenIds, string[] memory uris)
        internal
        virtual
    {
        require(
            tokenIds.length == uris.length,
            "ERC1155URIStrorage: tokenIds and uris length mismatch"
        );
        for (uint256 i = 0; i < tokenIds.length; ++i) {
            _setTokenURI(tokenIds[i], uris[i]);
        }
    }

    function burn(
        address account,
        uint256 id,
        uint256 value
    ) public virtual {
        require(value > 0, "ERC1155:burned amount must be greater than zero");
        require(
            account == _msgSender() || isApprovedForAll(account, _msgSender()),
            "ERC1155: caller is not owner nor approved"
        );
        super._burn(account, id, value);

        emit BurnSingle(_msgSender(), account, address(0), id, value);
    }

    function burnBatch(
        address account,
        uint256[] memory ids,
        uint256[] memory values
    ) public virtual {
        require(
            account == _msgSender() || isApprovedForAll(account, _msgSender()),
            "ERC1155: caller is not owner nor approved"
        );

        super._burnBatch(account, ids, values);
        emit BurnBatch(_msgSender(), account, address(0), ids, values);
    }


    function supportsInterface(bytes4 interfaceId) public view virtual override(ERC1155Upgradeable,IERC165Upgradeable) returns (bool) {
        return
            interfaceId == type(IERC1155ExtUpgradeable).interfaceId ||
            super.supportsInterface(interfaceId);
    }
    
}
