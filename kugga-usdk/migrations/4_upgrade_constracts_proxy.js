
const { upgradeProxy } = require('@openzeppelin/truffle-upgrades');

const BountyPledgeV2 = artifacts.require('BountyPledgeV2');
const BountyPledge = artifacts.require('BountyPledge');

module.exports = async function (deployer) {

  //nft
  const bountyPledgeV1 = await BountyPledge.deployed();
  console.log("Upgraded.bountyPledgeV1.address=", bountyPledgeV1.address);
  const bOwner = await bountyPledgeV1.owner();
  console.log("bOwner=",bOwner);

  //const instanceV3 = await upgradeProxy(existing.address, MyCollectibleERC20_V3, { deployer,fn:initialize_ext,args:[3] });
  const bountyPledgeV2 =await upgradeProxy(bountyPledgeV1.address, BountyPledgeV2, { deployer });
  
  console.log("Upgraded.bountyPledgeV2=", bountyPledgeV2.address);



 

};
