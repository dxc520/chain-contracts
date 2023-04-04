// SPDX-License-Identifier: MIT
pragma solidity >=0.4.22 <0.9.0;

import "@openzeppelin/contracts-upgradeable/token/ERC1155/IERC1155Upgradeable.sol";

interface IERC1155ExtUpgradeable is IERC1155Upgradeable {
    event BurnSingle(
        address indexed operator,
        address indexed from,
        address indexed to,
        uint256 tokenId,
        uint256 value
    );
    event BurnBatch(
        address indexed operator,
        address indexed from,
        address indexed to,
        uint256[] tokenIds,
        uint256[] values
    );
    event Gift(
        address operator,
        address indexed from,
        address indexed to,
        uint256 indexed tokenId,
        uint256 amount
    );
    event GiftBatch(
        address operator,
        address indexed from,
        address indexed to,
        uint256[] tokenIds,
        uint256[] amount
    );
    event MintSingleItem(
        address operator,
        address indexed to,
        uint256 indexed tokenId,
        uint256 indexed amount,
        string tokenURI
    );
    event MintBatchItems(
        address indexed operator,
        address indexed to,
        uint256[] tokenIds,
        uint256[] amounts,
        string[] tokenUris
    );

    function name() external view returns (string memory);

    function symbol() external view returns (string memory);

    function tokenURI(uint256 tokenId) external view returns (string memory);

    function burn(
        address account,
        uint256 tokenId,
        uint256 value
    ) external;

    function burnBatch(
        address account,
        uint256[] memory tokenIds,
        uint256[] memory values
    ) external;

    function gift(
        address from,
        address to,
        uint256 tokenId,
        uint256 amount
    ) external;

    function giftBatch(
        address from,
        address to,
        uint256[] memory tokenIds,
        uint256[] memory amounts
    ) external;

    function mintSingleItem(
        address to,
        uint256 tokenId,
        uint256 amount,
        string memory tokenURI
    ) external;

    function mintBatchItems(
        address to,
        uint256[] memory tokenIds,
        uint256[] memory amounts,
        string[] memory tokenUris
    ) external;
}
