
const { admin } = require('@openzeppelin/truffle-upgrades');
 const Migrations = artifacts.require("Migrations");

 module.exports = async function (deployer) {
  let migrationsInstance = await Migrations.deployed();
  console.log("MigrationsInstance=", migrationsInstance.address);

  let migrateOwner = await migrationsInstance.owner();
  console.log("migrateOwner=", migrateOwner);

  //admin-owner
  console.log("admin-owner");
  const instance = await admin.getInstance()
  console.log("admin.address=", instance.address);
 
  console.log("单独修改migration 的owner...然后切换至新用户，继续部署升级");

};

