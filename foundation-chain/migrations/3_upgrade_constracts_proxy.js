
const { upgradeProxy } = require('@openzeppelin/truffle-upgrades');

const CellarNftSaleV1 = artifacts.require('CellarNftSale');
const CellarMarketV1 = artifacts.require('CellarMarket');
const CellarNftSaleV2 = artifacts.require('CellarNftSaleV2');
const CellarMarketV2 = artifacts.require('CellarMarketV2');


module.exports = async function (deployer) {
  
  let  operatorUserAddr = "0xCbc13D27Bac30E6f6E70f769D8360dFd7d3b09F8";
  //nft
  const existingCellarNftSaleV1 = await CellarNftSaleV1.deployed();
  console.log("Upgraded.find.CellarNftSaleV1.address=", existingCellarNftSaleV1.address);

  //const instanceV3 = await upgradeProxy(existing.address, MyCollectibleERC20_V3, { deployer,fn:initialize_ext,args:[3] });
  const instanceCellarNftSaleV2 =await upgradeProxy(existingCellarNftSaleV1.address, CellarNftSaleV2, { deployer });
  console.log("Upgraded.instanceCellarNftSaleV2=", instanceCellarNftSaleV2.address);

  instanceCellarNftSaleV2.grantMintRole(operatorUserAddr);
///market
  const existingCellarMarketV1 = await CellarMarketV1.deployed();
  console.log("Upgraded.find.existingCellarMarketV1.address=", existingCellarMarketV1.address);

  const instanceCellarMarketV2 =await upgradeProxy(existingCellarMarketV1.address, CellarMarketV2, { deployer });
  console.log("Upgraded.instanceCellarMarketV2=", instanceCellarMarketV2.address);



 

};
