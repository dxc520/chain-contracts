
const { deployProxy } = require('@openzeppelin/truffle-upgrades');

const CellarNftSale = artifacts.require('CellarNftSale');
const CellarMarket = artifacts.require('CellarMarket');

var erc721Name = "WinsArtsOrg:721";
var erc721Symbol = "WAO721";
var erc20Name = "WinsArtsOrg:ArtToken";
var erc20Symbol = "WAOT";


module.exports = async function (deployer) {
  
    //step1: nft
    const cellarNFTSaleInstance = await deployProxy( CellarNftSale, [erc721Name, erc721Symbol,2], { deployer });
    console.log('CellarNftSale.deployed', cellarNFTSaleInstance.address);
    
    //step2:coin
    const cellarMarketInstance = await deployProxy(CellarMarket, [cellarNFTSaleInstance.address,erc20Name, erc20Symbol,0], { deployer });
    console.log('CellarMarket.deployed', cellarMarketInstance.address);
 
    //step3:init data
    cellarNFTSaleInstance.grantMintRole(cellarMarketInstance.address);
    cellarNFTSaleInstance.grantMarket(cellarMarketInstance.address);
    cellarNFTSaleInstance.setPlatformFeePercent(3);
   
};
