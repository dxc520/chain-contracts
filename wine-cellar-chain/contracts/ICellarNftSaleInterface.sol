// SPDX-License-Identifier: MIT
pragma solidity >=0.4.22 <0.9.0;

import "@openzeppelin/contracts-upgradeable/token/ERC721/IERC721Upgradeable.sol";

interface ICellarNftSaleInterface is IERC721Upgradeable {
    event Gift(
        address indexed from,
        address indexed to,
        uint256 indexed tokenId
    );
    event MakeOrder(
        address _operator,
        address indexed _seller,
        uint256 indexed tokenId,
        uint256 indexed orderId,
        uint256 sellingPrice,
        uint16 feePercent
    );
    event UpdateOrder(
        address _operator,
        address indexed _seller,
        uint256 indexed _tokenId,
        uint256 indexed _orderId,
        uint256 _sellingPrice,
        uint16 _feePercent
    );
    event CancelOrder(
        address _operator,
        address indexed _seller,
        uint256 indexed _tokenId,
        uint256 indexed _orderId
    );
    event Mint721Item(
        address _operator,
        address indexed _to,
        uint256 indexed _tokenId,
        uint256 indexed _amount,
        string _tokenURI
    );
    event TransferNFT(
        address indexed _buyer,
        address indexed _seller,
        uint256 indexed _tokenId,
        uint256 _amount,
        uint256 _tradeId,
        uint256 _orderId
    ); 
    
    function gift(
        address from,
        address to,
        uint256 tokenId
    ) external;

    function hasMintRole(address _account) external view returns (bool);

    function grantMintRole(address operator) external;

    function renounceMintRole(address account) external;

    function grantMarket(address market) external;

    function setPlatformFeePercent(uint16 platformFeePercent) external;

    function getPlatformFeePercent() external view returns (uint16);

    function mintItem(
        address to,
        uint256 tokenId,
        string memory tokenURI
    ) external returns (uint256);

    function makeOrder(uint256 tokenId, uint256 sellingPrice)
        external
        returns (uint256);

    function cancleOrder(uint256 tokenId) external returns (bool);

    function updateOrder(uint256 tokenId, uint256 sellingPrice)
        external
        returns (bool);

    function transferNFT(address buyer, uint256 tokenId)
        external
        returns (uint256, uint256);

    function getSellingOrderByTokenId(uint256 tokenId)
        external
        view
        returns (
            uint256 orderId,
            address seller,
            uint256 sellingPrice,
            uint16 feePercent
        );

    function getSellingOrderByOrderId(uint256 orderId)
        external
        view
        returns (
            uint256 tokenId,
            address seller,
            uint256 sellingPrice,
            uint16 feePercent
        );

    function getNFTOwnerByTokenId(uint256 tokenId)
        external
        view
        returns (address);

    function getSaleOrderIdByToken(uint256 tokenId)
        external
        view
        returns (uint256);

    function getCurrentMaxSaleOrderId() external view returns (uint256);
}
