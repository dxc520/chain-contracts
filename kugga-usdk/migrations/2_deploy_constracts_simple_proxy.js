
const { deployProxy } = require('@openzeppelin/truffle-upgrades');

const UChildERC20 = artifacts.require('UChildERC20Upgradeable');
 
const BountyPledge = artifacts.require('BountyPledge');

const bountyName = "BountyPledge";
const bountySymbol = "BP";

const usdcMumbai = '0xE097d6B3100777DC31B34dC2c58fB524C2e76921'; //mumbai 测试网的 usdc代币
const usdcSymbolMumbai = "USDC";
const usdc4kMumbai = '0xe235AF2460FDCfDDfb1892754A801947dF30D0CD'; //mumbai 测试网的 usdc4k代币
const usdc4kSymbolMumbai = "USDC4K";
const usdc4kSolarChain = '0xE5E7F44298ffed88eddfa0b8D495ffa53caE8355'; //mumbai 测试网的 usdc4k代币
const usdc4kSymbolSolarChain = "USDC4K";
const usdkSolarChain = "0xEc468b9fef39866513DA4545e52C8fb037AC611b";//SolarChain
const usdkSymbolSolarChain ="USDK";
const usdkMumbai = "0xf03a1cd9502fe5965756828DFc5B200226c2e922"; //Polygon
const usdkSymbolMumbai = "USDK"
const usdtPolygonMainnet = "0xc2132d05d31c914a87c6611c10748aeb04b58e8f";//polygon-MainNet 主网的USDT
const usdtSymbolPolygonMainnet ="USDT";
//使用哪个配置哪个 todo fix
const coinErc20 = usdkMumbai;
const coinErc20Symbol = usdkSymbolMumbai;
 
module.exports = async function (deployer) {
    //step1:coin
  /* 
    const usdkInstance = await deployProxy(UChildERC20, [erc20Name, erc20Symbol,decimals,childChainManager], { deployer });
    console.log('UsdkERC20.deployed', usdkInstance.address);
 */
    //const bpInstance = await deployProxy(BountyPledge, [bountyName, bountySymbol,usdkInstance.address], { deployer });
    const bpInstance = await deployProxy(BountyPledge, [bountyName, bountySymbol,coinErc20,coinErc20Symbol], { deployer });

    console.log('bountyPledge.deployed', bpInstance.address);
 

    //step2:init data
    /**
    const depositRole =  await usdkInstance.DEPOSITOR_ROLE();
    console.log("depositRole=",depositRole);

    let tx1=  await usdkInstance.grantRole(depositRole,admin);
    console.log("tx.grantRole=",tx1.tx);
    let tx2 = await usdkInstance.deposit2(admin ,654321000000);
    console.log("tx.deposit2.admin=",tx2.tx);


    let amountData = web3.eth.abi.encodeParameters(['uint256'],[3000000]);
    let txMint2 = await usdkInstance.deposit(user2, amountData);
    console.log("tx.usdk.mint=", txMint2.tx);
    
    //让质押合约可操作usdt
    let allowance = 2000000;
    let transferAmountAvg = 500000;
    let tx4 = await usdkInstance.approve(bpInstance.address,allowance);
    console.log("tx.increaseAllowance=",tx4.tx);
 
    //let bountyId = Math.floor(new Date().getTime() / 1000);
    let startTime=Math.floor(new Date().getTime() / 1000);
    let endTime0=Math.floor(new Date('2023-1-26 19:45:23').getTime() / 1000);
    let endTime=Math.floor(startTime +3600);
    console.log("startTime=",startTime,";endTime=",endTime,";endTimeTmp=",endTime0);

    bountyReq ={
        startTime:startTime,
        endTime:endTime,
        bountyType:2,//1=single,2=multiple
        cid:"QM02dfdfierewrewlerwr",
        amount:allowance,
        unit:1,//1=usdt,2=matic,
         
    };

    //质押
    let tx6 = await bpInstance.pledge(bountyReq);
    console.log("tx.pledge=",tx6.tx);

    let currentIdx = (await bpInstance.getCurrentBountyIdx()).toNumber();
    console.log("currentIdx=",currentIdx);

    //分账
    let bountyId = currentIdx; 
    console.log("bountyId=",bountyId);
    let tx7 = await bpInstance.transfer(bountyId,pz,transferAmountAvg,"QM:reply-"+pz+"-"+1);
    console.log("tx7.transfer.1=",tx7.tx);
    tx7 = await bpInstance.transfer(bountyId,user2,transferAmountAvg,"QM:reply-"+user2+"-"+1);
    console.log("tx7.transfer.2=",tx7.tx);

    // getWinner
    let winners = await bpInstance.getWinners(bountyId);
    console.log("winners=",winners);

    //expire
    let expiredTime  = Math.floor(new Date().getTime() / 1000);
    let tx8 =  await bpInstance.setBountyExpired(bountyId,expiredTime);
    console.log("tx8.setBountyExpired=",tx8.tx);

    //withdraw
    let tx9 = await bpInstance.withdraw(bountyId);
    console.log("tx9.withdraw=",tx9.tx);

    let bountyStatus = await bpInstance.getBounty(bountyId);
    console.log("tx10.bountyStatus=",bountyStatus);

    */
};