
const { deployProxy } = require('@openzeppelin/truffle-upgrades');


const artExhibitionERC1155 = artifacts.require('ArtExhibitionERC1155');
const walletProxyUpgradeable = artifacts.require('WalletProxyUpgradeable');

const erc1155Name = "ArtsOrg:1155";
const erc1155Symbol = "AO1155";
const walletProxyName="walletProxy";
const walletProxySymbol="WP";
const artExhibitionSuperAdmin="0x8fD86498ff1B750454f5187cc0B2bC6b51d0E07B";
const walletProxySuperAdmin="0x8fD86498ff1B750454f5187cc0B2bC6b51d0E07B";
module.exports = async function (deployer) {

    //step1: nft
     
    const artExhibitionERC1155Instance = await deployProxy( artExhibitionERC1155, [erc1155Name, erc1155Symbol,""], { deployer });
    console.log('artExhibitionERC1155.deployed=', artExhibitionERC1155Instance.address);
    
    //step2:walletProxy
    const walletProxyInstance = await deployProxy( walletProxyUpgradeable, [artExhibitionERC1155Instance.address,walletProxyName,walletProxySymbol], { deployer });
    console.log('walletProxy.deployed=', walletProxyInstance.address);
    
    //step3:init data
    artExhibitionERC1155Instance.setWalletProxyInit(walletProxyInstance.address);
    artExhibitionERC1155Instance.grantAdminRole(artExhibitionSuperAdmin);
    walletProxyInstance.grantAdminRole(walletProxySuperAdmin);
    
};