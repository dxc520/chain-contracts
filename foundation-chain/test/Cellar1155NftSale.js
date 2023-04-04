const CellarNft1155Sale = artifacts.require("CellarNft1155Sale");

contract('CellarNft1155Sale', (accounts) => {

    console.log("the CellarNft1155Sale wallet accounts=",accounts);

    let _mintOwner =accounts[0];//address(0x8fD86498ff1B750454f5187cc0B2bC6b51d0E07B);
    let _tokenURI = "https://pnode.solarfs.io/dn/file/1e00000000000e57f41272545f3403c31eef506c78a37399f3456e9ce936756204061c9fb5281fa541a2b2983cb5d005d77f6edfeaffe57ac9c79fee3762969e/meta8.json";
    let _tokenId= 123;
    let _mintAmount=1000;//铸币数量
    let _burnAmount=20;//销毁数量
    let _giftAmount=100;//赠与数量
    let _giftTo ="0x4293859c2429c5554174D4ea9b336559933847D2";
    let _CellarNft1155SaleInstance;

    beforeEach(async function () {
         _CellarNft1155SaleInstance =await CellarNft1155Sale.deployed();
    });

    it('get contract name', async () => {
        let  _name = await _CellarNft1155SaleInstance.name();
        console.log("_name= ",_name);
    });

    it('get contract symbol', async () => {
        let  _symbol = await _CellarNft1155SaleInstance.symbol();
        console.log("_symbol= ",_symbol);
    });

    it('mintErc1155Nft,then check the owner amount', async () => {
        console.log("_mintOwner= ",_mintOwner);
        
        await _CellarNft1155SaleInstance.mintItem(_mintOwner, _tokenId,_mintAmount, _tokenURI);

        let _ownNftAmount = await _CellarNft1155SaleInstance.balanceOf(_mintOwner,_tokenId);
        console.log("_mintAmount=",_ownNftAmount.toNumber());
        assert.equal(_mintAmount,_ownNftAmount,"after mint.The owner own nft amount is different.");
    });

    it('giftErc1155Nft,then check the owner amount', async () => {
        let _ownNftAmount = await _CellarNft1155SaleInstance.balanceOf(_mintOwner,_tokenId);
        console.log("_mintAmount=",_ownNftAmount.toNumber());
        assert.equal(_mintAmount,_ownNftAmount,"after mint.The owner own nft amount is different.");

        console.log("gift.amount=",_giftAmount);
        await _CellarNft1155SaleInstance.gift(_mintOwner,_giftTo,_tokenId,_giftAmount);

        let _ownerOriginAmount = await _CellarNft1155SaleInstance.balanceOf(_mintOwner,_tokenId);
        console.log("after gift : _originOwnerAmount=",_ownerOriginAmount.toNumber());
        assert.equal(_ownerOriginAmount,_mintAmount-_giftAmount,"after gift.The origin owner nft amount is not expected.");


        let _ownerGiftAmount = await _CellarNft1155SaleInstance.balanceOf(_giftTo,_tokenId);
        console.log("after gift : _giftOwnerAmount=",_ownerGiftAmount.toNumber());
        assert.equal(_ownerGiftAmount,_giftAmount,"after gift.The gift owner nft amount is not expected.");

    });

    it('BurnErc1155Nft,then check the owner amount', async () => {
        
        let _ownerOriginAmount = await _CellarNft1155SaleInstance.balanceOf(_mintOwner,_tokenId);
        console.log("after gift : _originOwnerAmount=",_ownerOriginAmount.toNumber());
        assert.equal(_ownerOriginAmount,_mintAmount-_giftAmount,"after gift.The origin owner nft amount is not expected.");

        console.log("exec burn.amount=",_burnAmount);
        await _CellarNft1155SaleInstance.burn(_mintOwner,_tokenId,_burnAmount);

        _ownerOriginAmount = await _CellarNft1155SaleInstance.balanceOf(_mintOwner,_tokenId);
        console.log("after gift : _originOwnerAmount=",_ownerOriginAmount.toNumber());
        assert.equal(_ownerOriginAmount,_mintAmount-_giftAmount-_burnAmount,"after burn.The origin owner nft amount is not expected.");

    });

});