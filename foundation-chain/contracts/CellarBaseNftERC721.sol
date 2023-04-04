// SPDX-License-Identifier: MIT
pragma solidity >=0.4.22 <0.9.0;

import "@openzeppelin/contracts-upgradeable/token/ERC721/extensions/ERC721URIStorageUpgradeable.sol";

contract CellarBaseNftERC721 is ERC721URIStorageUpgradeable {

    event Gift(address indexed from, address indexed to, uint256 indexed tokenId);

    
    function _init_base_unchanged( )internal onlyInitializing{
      
    }


    /**
    * 内部使用，铸币NFT
    */
    function _awardItem(address _to, uint256 _tokenId, string memory _tokenURI) internal virtual  returns (uint256){
        _safeMint(_to, _tokenId);
        _setTokenURI(_tokenId, _tokenURI);
        return _tokenId;
    }

    function approve(address to, uint256 tokenId) public virtual override {
        address owner = ERC721Upgradeable.ownerOf(tokenId);
        require(to != owner, "ERC721: approval to current owner");

        require(_msgSender() == owner || getApproved(tokenId) == _msgSender() || isApprovedForAll(owner, _msgSender()),
            "ERC721: approve caller is not owner nor approved for all"
        );

        _approve(to, tokenId);
    }

    
    function gift(address from, address to, uint256 tokenId)external{
        ERC721Upgradeable.safeTransferFrom( from,  to,  tokenId);
        emit Gift(from,to,tokenId);
    }

}