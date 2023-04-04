
const { deployProxy } = require('@openzeppelin/truffle-upgrades');

const CellarNftSale = artifacts.require('CellarNftSale');
const CellarNft1155Sale = artifacts.require('CellarNft1155Sale');
const CellarMarket = artifacts.require('CellarMarket');

const erc721Name = "WinsArtsOrg:721";
const erc721Symbol = "WAO721";
const erc20Name = "WinsArtsOrg:ArtToken";
const erc20Symbol = "WAOT";
const erc1155Name = "WinsArtsOrg:1155";
const erc1155Symbol = "WAO1155";
const operatorUserAddr = "0xB0618ED597cdd2D1d1646a1fAA51678b969c5e6b";
const feePercent = 2;

module.exports = async function (deployer) {

    //step1: nft
    const cellarNFTSaleInstance = await deployProxy( CellarNftSale, [erc721Name, erc721Symbol,feePercent], { deployer });
    console.log('CellarNft721Sale.deployed', cellarNFTSaleInstance.address);
    
    const cellar1155NFTSaleInstance = await deployProxy( CellarNft1155Sale, [erc1155Name, erc1155Symbol,"",feePercent], { deployer });
    console.log('cellar1155NFTSale.deployed', cellar1155NFTSaleInstance.address);
    

    //step2:coin
    const cellarMarketInstance = await deployProxy(CellarMarket, [cellarNFTSaleInstance.address,cellar1155NFTSaleInstance.address,erc20Name, erc20Symbol,0], { deployer });
    console.log('CellarMarketERC20.deployed', cellarMarketInstance.address);
 
    //step3:init data
    cellarNFTSaleInstance.grantMintRole(cellarMarketInstance.address);
    cellarNFTSaleInstance.grantMarket(cellarMarketInstance.address);
    cellarNFTSaleInstance.grantMintRole(operatorUserAddr);

    cellar1155NFTSaleInstance.grantMintRole(cellarMarketInstance.address);
    cellar1155NFTSaleInstance.grantMarket(cellarMarketInstance.address);
    cellar1155NFTSaleInstance.grantMintRole(operatorUserAddr);
   
    cellarMarketInstance.grantMintRole(operatorUserAddr);
};