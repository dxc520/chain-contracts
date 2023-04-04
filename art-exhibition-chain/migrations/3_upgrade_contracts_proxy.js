const {admin } = require('@openzeppelin/truffle-upgrades');


module.exports = async function (deployer) {

  const instance  =  await admin.getInstance()
  console.log("admin.address=",instance.address);

};
