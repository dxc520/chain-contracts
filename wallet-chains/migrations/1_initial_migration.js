const Migrations = artifacts.require("Migrations");

var deployedName = "WinsArtsOrg:Deployer";
var deployedSymbol ="WAODeploy";

module.exports = function (deployer) {
  deployer.deploy(Migrations);
};
