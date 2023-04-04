// SPDX-License-Identifier: MIT
pragma solidity >=0.4.22 <0.9.0;

import "./ArtExhibitionERC1155.sol";

contract HackAttackV1 {
     IERC1155ExtUpgradeable private erc1155Sale;

    event receiveEvent(address indexed sender,uint256 indexed value );
    event fallbackEvent(address indexed sender,uint256 indexed value );


    constructor(address _erc1155Addr) {
        erc1155Sale = IERC1155ExtUpgradeable(_erc1155Addr);
    }

    function MintERC1155(
        uint256 id,
        uint256 amount,
        string memory tokenURI
    ) public {
       
        erc1155Sale.mintSingleItem(address(this), id, amount, tokenURI);
    }

 
    function GiftERC1155(
        address from,
        address to,
        uint256 id,
        uint256 amount
    ) public {
        erc1155Sale.gift(from, to, id, amount);
    }

    function BurnERC1155(
        address account,
        uint256 tokenId,
        uint256 amount)public {
            
            erc1155Sale.burn(account, tokenId, amount);
    }
    receive() external payable  {
       emit receiveEvent(msg.sender,msg.value);
    }
    
    fallback()external payable{
        emit fallbackEvent(msg.sender, msg.value);
    }
}
