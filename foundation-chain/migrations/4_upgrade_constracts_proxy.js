
const { upgradeProxy } = require('@openzeppelin/truffle-upgrades');

const CellarMarketV2 = artifacts.require('CellarMarketV2');
const CellarMarketV3 = artifacts.require('CellarMarketV3');


module.exports = async function (deployer) {

  
///market
  const existingCellarMarketV2 = await CellarMarketV2.deployed();
  console.log("Upgraded.find.existingCellarMarketV2.address=", existingCellarMarketV2.address);

  const instanceCellarMarketV3 =await upgradeProxy(existingCellarMarketV2.address, CellarMarketV3, { deployer });
  console.log("Upgraded.instanceCellarMarketV3=", instanceCellarMarketV3.address);
  instanceCellarMarketV3.setDecimal(0);

 

};
