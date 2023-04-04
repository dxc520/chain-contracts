const Migrations = artifacts.require("Migrations");

module.exports = function (deployer) {
  let instance = deployer.deploy(Migrations);
  console.log("Migrate.instance=",instance);

};