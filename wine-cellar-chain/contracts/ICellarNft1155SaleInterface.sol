// SPDX-License-Identifier: MIT
pragma solidity >=0.4.22 <0.9.0;

import "@openzeppelin/contracts-upgradeable/token/ERC1155/IERC1155Upgradeable.sol";

interface ICellarNft1155SaleInterface is IERC1155Upgradeable {
    event BurnSingle(
        address indexed operator,
        address indexed from,
        address indexed to,
        uint256 id,
        uint256 value
    );
    event BurnBatch(
        address indexed operator,
        address indexed from,
        address indexed to,
        uint256[] ids,
        uint256[] values
    );
    event Gift(
        address indexed from,
        address indexed to,
        uint256 indexed tokenId,
        uint256 amount
    );
    event GiftBatch(
        address indexed from,
        address indexed to,
        uint256[] tokenIds,
        uint256[] amount
    );
    event Mint1155Item(
        address operator,
        address indexed _to,
        uint256 indexed _tokenId,
        uint256 indexed _amount,
        string _tokenURI
    );
    event MintBatchItems(
        address indexed operator,
        address indexed _to,
        uint256[] _tokenIds,
        uint256[] _amounts,
        string[] _tokenUris
    );
    event MakeOrder(
        address operator,
        address indexed seller,
        uint256 indexed tokenId,
        uint256 amount,
        uint256 indexed orderId,
        uint256 unitPrice,
        uint256 totalPrice,
        uint16 feePercent
    );
    event UpdateOrder(
        address _operator,
        address indexed _seller,
        uint256 indexed _tokenId,
        uint256 _amount,
        uint256 indexed _orderId,
        uint256 _unitPrice,
        uint256 _totalPrice,
        uint16 _feePercent
    );
    event CancelOrder(
        address _operator,
        address indexed _seller,
        uint256 indexed _tokenId,
        uint256 _amount,
        uint256 indexed _orderId,
        uint256 _unitPrice,
        uint256 _totalPrice,
        uint16 _feePercent
    );
    event TradeNft(
        address indexed _buyer,
        address indexed _seller,
        uint256 indexed _tokenId,
        uint256 amount,
        uint256 _tradeId,
        uint256 _orderId
    );

    function name() external view returns (string memory);

    function symbol() external view returns (string memory);

    function tokenURI(uint256 tokenId) external view returns (string memory);

    function uri(uint256 tokenId) external view returns (string memory);

    function burn(
        address account,
        uint256 id,
        uint256 value
    ) external;

    function burnBatch(
        address account,
        uint256[] memory ids,
        uint256[] memory values
    ) external;

    function hasMintRole(address _account) external view returns (bool);

    function grantMintRole(address operator) external;

    function renounceMintRole(address account) external;

    function gift(
        address from,
        address to,
        uint256 id,
        uint256 amount
    ) external;

    function giftBatch(
        address from,
        address to,
        uint256[] memory ids,
        uint256[] memory amounts
    ) external;

    function setPlatformFeePercent(uint16 platformFeePercent) external;

    function getPlatformFeePercent() external view;

    function grantMarket(address market) external;

    function mintItem(
        address to,
        uint256 id,
        uint256 amount,
        string memory tokenURI
    ) external;

    function mintBatchItems(
        address to,
        uint256[] memory ids,
        uint256[] memory amounts,
        string[] memory tokenUris
    ) external;

    function makeOrder(
        uint256 tokenId,
        uint256 amount,
        uint256 unitPrice
    ) external;

    function cancelOrder(uint256 orderId) external;

    function updateOrder(
        uint256 orderId,
        uint256 stockAmount,
        uint256 sellingUnitPrice
    ) external;




    function getOrder(uint256 orderId_)
        external
        view
        returns (
            address seller,
            uint256 tokenId,
            uint256 amount,
            uint16 fee,
            uint256 totalPrice
        );

    function getOrderSeller(uint256 orderId_)
        external
        view
        returns (address seller);

    function tradeNft(
        address buyer,
        uint256 tokenId,
        address from
    ) external returns (uint256 tradeId);
}
