
const { admin } = require('@openzeppelin/truffle-upgrades');
const BountyPledge = artifacts.require('BountyPledge');
const Migrations = artifacts.require("Migrations");

const newOwner = "0xD61c47B00BB534993f6A79A6561c7a6184aDF616";
module.exports = async function (deployer) {
  let migrationsInstance = await Migrations.deployed();
  console.log("MigrationsInstance=", migrationsInstance.address);

  let migrateOwner = await migrationsInstance.owner();
  console.log("migrateOwner=", migrateOwner);

  const bountyPledgeV1 = await BountyPledge.deployed();
  console.log("Upgraded.bountyPledgeV1.address=", bountyPledgeV1.address);

  let bpOwner = await bountyPledgeV1.owner();
  console.log("bp.owner1=", bpOwner);




  //admin-owner
  console.log("admin-owner");
  const instance = await admin.getInstance()
  console.log("admin.address=", instance.address);

  const ship = await admin.transferProxyAdminOwnership(newOwner)
  console.log("Upgraded.ship=", ship);

  const instance2 = await admin.getInstance()
  console.log("admin2.address=", instance2.address);


  //bounty-pledge--owner
  console.log("bounty-pledge--owner\n");
  bpOwner = await bountyPledgeV1.owner();
  console.log("bp.owner1=", bpOwner);

  let transOwner = await bountyPledgeV1.transferOwnership(newOwner);
  console.log("transOwner.tx=", transOwner.tx);

  bpOwner = await bountyPledgeV1.owner();
  console.log("bp.owner2=", bpOwner);

  //migrate--onwer
  console.log("migrate owner...\n");
  //let transMigrateOwner = await migrationsInstance.transferOwnership(newOwner);
  //console.log("transMigrateOwner.tx=", transMigrateOwner.tx);

  migrateOwner = await migrationsInstance.owner();
  console.log("migrateOwner=", migrateOwner);



  console.log("单独修改migration 的owner...然后切换至新用户，继续部署升级");

};
