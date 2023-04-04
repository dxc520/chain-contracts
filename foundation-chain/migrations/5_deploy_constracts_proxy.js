
const { deployProxy } = require('@openzeppelin/truffle-upgrades');

const CellarNft1155Sale = artifacts.require('CellarNft1155Sale');
const CellarMarketExt = artifacts.require('CellarMarketV3');

var erc1155Name = "WinsArtsOrg:1155";
var erc1155Symbol = "WAO1155";

module.exports = async function (deployer) {
    let  operatorUserAddr = "0xCbc13D27Bac30E6f6E70f769D8360dFd7d3b09F8";

    //step1: nft
    const cellar1155NFTSaleInstance = await deployProxy( CellarNft1155Sale, [erc1155Name, erc1155Symbol,"",2], { deployer });
    console.log('cellar1155NFTSale.deployed', cellar1155NFTSaleInstance.address);
    
    //step2:coin
    const cellarMarketInstance = await CellarMarketExt.deployed();
    console.log('CellarMarket.deployed', cellarMarketInstance.address);
 
    //step3:init data
     
    cellar1155NFTSaleInstance.grantMintRole(cellarMarketInstance.address);
    cellar1155NFTSaleInstance.grantMarket(cellarMarketInstance.address);
    cellar1155NFTSaleInstance.setPlatformFeePercent(3);
    cellar1155NFTSaleInstance.grantMintRole(operatorUserAddr);

};
