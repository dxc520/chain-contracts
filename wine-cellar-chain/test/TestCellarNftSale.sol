// SPDX-License-Identifier: MIT
pragma solidity >=0.4.22 <0.9.0;

import "truffle/Assert.sol";
import "truffle/DeployedAddresses.sol";
import "../contracts/CellarNftSale.sol";

contract TestCellarNftSale {
/*
      function testMintItem( )public {
            string memory name="testNft";
            string memory symbol="simple";
            uint16  platformFeePercent=3;

            CellarNftSale erc721NftSale= new CellarNftSale();
            erc721NftSale.initialize(name, symbol, platformFeePercent);

            address _to =address(0x99A731E4d5f706BC6D8bB415Df12ad9794Dbb7EF);
            //Assert.equal(_owner,_to,"the owner is different ");

            string memory _tokenURI = "https://pnode.solarfs.io/dn/file/1e00000000000e57f41272545f3403c31eef506c78a37399f3456e9ce936756204061c9fb5281fa541a2b2983cb5d005d77f6edfeaffe57ac9c79fee3762969e/meta8.json";
            uint256 _tokenId= 123;

            erc721NftSale.mintItem(_to, _tokenId, _tokenURI);
            
            //Assert.equal(erc721NftSale.ownerOf(_tokenId),tx.origin,"The owner is not equal.");
            Assert.equal(erc721NftSale.ownerOf(_tokenId),_to,"The owner is not equal.");

      }

      function testGift( )public {
            string memory name="testNft";
            string memory symbol="simple";
            uint16  platformFeePercent=3;

            CellarNftSale erc721NftSale= new CellarNftSale();
            erc721NftSale.initialize(name, symbol, platformFeePercent);

            address _owner =address(0x99A731E4d5f706BC6D8bB415Df12ad9794Dbb7EF);
            //Assert.equal(_owner,_to,"the owner is different ");

            string memory _tokenURI = "https://pnode.solarfs.io/dn/file/1e00000000000e57f41272545f3403c31eef506c78a37399f3456e9ce936756204061c9fb5281fa541a2b2983cb5d005d77f6edfeaffe57ac9c79fee3762969e/meta8.json";
            uint256 _tokenId= 124;

            erc721NftSale.mintItem(_owner, _tokenId, _tokenURI);
            
            Assert.equal(erc721NftSale.ownerOf(_tokenId),_owner,"after mint,The owner is not expected.");

            //address _giftTo = address(0x8fD86498ff1B750454f5187cc0B2bC6b51d0E07B);

            //erc721NftSale.gift(_owner, _giftTo, _tokenId);
            //Assert.equal(erc721NftSale.ownerOf(_tokenId),tx.origin,"The owner is not equal.");
            //Assert.equal(erc721NftSale.ownerOf(_tokenId),_giftTo,"after gift.The owner is not expected.");

      }
*/
}