
const { upgradeProxy } = require('@openzeppelin/truffle-upgrades');

const BountyPledgeV2 = artifacts.require('BountyPledgeV2');
const BountyPledgeV3 = artifacts.require('BountyPledgeV3');

const organizer = "0x8fd86498ff1b750454f5187cc0b2bc6b51d0e07b";
const rate =10;
module.exports = async function (deployer) {

  //nft
  const bountyPledgeV2 = await BountyPledgeV2.deployed();
  console.log("Upgraded.bountyPledgeV2.address=", bountyPledgeV2.address);
  const bOwner = await bountyPledgeV2.owner();
  console.log("bOwner=",bOwner);

  //const instanceV3 = await upgradeProxy(existing.address, MyCollectibleERC20_V3, { deployer,fn:initialize_ext,args:[3] });
  const bountyPledgeV3 =await upgradeProxy(bountyPledgeV2.address, BountyPledgeV3, { deployer });
  
  console.log("Upgraded.bountyPledgeV3=", bountyPledgeV3.address);

  let tx1= await bountyPledgeV3.setCommissionRate(rate);
  console.log("tx1=",tx1.tx);
  let tx2= await bountyPledgeV3.setOrganizer(organizer);
  console.log("tx2=",tx2.tx);

};
