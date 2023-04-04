const { upgradeProxy } = require('@openzeppelin/truffle-upgrades');

const CellarMarketV4 = artifacts.require('CellarMarketV4');
const CellarMarketV3 = artifacts.require('CellarMarketV3');
const CellarNft1155SaleExt = artifacts.require('CellarNft1155Sale');


module.exports = async function (deployer) {
  let  operatorUserAddr = "0xCbc13D27Bac30E6f6E70f769D8360dFd7d3b09F8";

  ///market
  const existingCellarMarketV3 = await CellarMarketV3.deployed();
  console.log("existing.CellarMarketV3.address=", existingCellarMarketV3.address);

  const instanceCellarMarketV4 =await upgradeProxy(existingCellarMarketV3.address, CellarMarketV4, { deployer });
  console.log("Upgraded.instanceCellarMarketV4=", instanceCellarMarketV4.address);

  const cellarNft1155Instance = await CellarNft1155SaleExt.deployed();
  console.log('existing.cellarNft1155.deployed', cellarNft1155Instance.address);

  instanceCellarMarketV4.setSpenderERC1155(cellarNft1155Instance.address)
  instanceCellarMarketV4.grantMintRole(operatorUserAddr);

};
