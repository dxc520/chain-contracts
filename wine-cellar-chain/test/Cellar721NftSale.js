const CellarNftSale = artifacts.require("CellarNftSale");

contract('CellarNftSale', (accounts) => {

    console.log("the CellarNftSale wallet accounts=",accounts);

    let _mintOwner =accounts[0];//address(0x8fD86498ff1B750454f5187cc0B2bC6b51d0E07B);
    let _tokenURI = "https://pnode.solarfs.io/dn/file/1e00000000000e57f41272545f3403c31eef506c78a37399f3456e9ce936756204061c9fb5281fa541a2b2983cb5d005d77f6edfeaffe57ac9c79fee3762969e/meta8.json";
    let _tokenId= 123;
    let _giftTo ="0x4293859c2429c5554174D4ea9b336559933847D2";
    let _CellarNftSaleInstance;

    beforeEach(async function () {
         _CellarNftSaleInstance =await CellarNftSale.deployed();
    });

    it('get contract name', async () => {
        let  _name = await _CellarNftSaleInstance.name();
        console.log("_name= ",_name);
    });

    it('get contract symbol', async () => {
        let  _symbol = await _CellarNftSaleInstance.symbol();
        console.log("_symbol= ",_symbol);
    });

    it('mintErc721Nft,then check the owner', async () => {
        console.log("_mintOwner= ",_mintOwner);
        
        await _CellarNftSaleInstance.mintItem(_mintOwner, _tokenId, _tokenURI);

        let _ownerOf = await _CellarNftSaleInstance.ownerOf(_tokenId);
        console.log("_ownerOf",_ownerOf);
        assert.equal(_ownerOf,_mintOwner,"after mint.The owner is not equal.");
    });

    it('giftErc721Nft,then check the owner', async () => {
        let _giftTokenId = _tokenId;
        let _ownerOf = await _CellarNftSaleInstance.ownerOf(_giftTokenId);
        console.log("_ownerOf=",_ownerOf);
        assert.equal(_ownerOf,_mintOwner,"after mint.The owner is not equal.");

        console.log("_giftTo= ",_giftTo);
        await _CellarNftSaleInstance.gift(_mintOwner,_giftTo,_giftTokenId);

        let _newOwnerOf = await _CellarNftSaleInstance.ownerOf(_giftTokenId);
        console.log("_newOwnerOf=",_newOwnerOf);
        assert.equal(_newOwnerOf,_giftTo,"after gift.The owner is not equal.");

    });


});