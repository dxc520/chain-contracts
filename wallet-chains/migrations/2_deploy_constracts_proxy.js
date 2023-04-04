
const { deployProxy } = require('@openzeppelin/truffle-upgrades');

const WalletERC20 = artifacts.require('WalletERC20');
 

const erc20Name = "WinsArtsOrg:ArtToken";
const erc20Symbol = "WAOT";

const operatorUserAddr = "0xE5a5cF7Ce53FFe043e3223bd620c8E5a9C068319";//dev
//const operatorUserAddr = "0x408a4a7e74d2e457EbEAabd628e5B5C206A6A407"; //test
module.exports = async function (deployer) {


    //step1:coin
    const walletInstance = await deployProxy(WalletERC20, [erc20Name, erc20Symbol,0], { deployer });
    console.log('WalletERC20.deployed', walletInstance.address);
 
    //step2:init data
    walletInstance.grantMintRole(operatorUserAddr);
    console.log("coin ERC20 is over ... ");
 


    

    
};