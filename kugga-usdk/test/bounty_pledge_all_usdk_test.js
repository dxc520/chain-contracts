const usdtErc20 = artifacts.require("UChildERC20Upgradeable");
//const bountyPledge = artifacts.require("BountyPledge");
const bountyPledge = artifacts.require("BountyPledgeV2");
contract('bountyPledge', (accounts) => {

    console.log("the bountyPledge accounts=", accounts);

    let from = accounts[0];//address(0x8fD86498ff1B750454f5187cc0B2bC6b51d0E07B);
    let to = accounts[1];
    let user2 = accounts[2];
    let _usdtErc20Instance;
    let _bountyPledgeInstance;
    //const usdAsset = '0xE097d6B3100777DC31B34dC2c58fB524C2e76921'; //mumbai 测试网的 usdc代币
    //const usdAssetSymbol = "USDC";
    const usdAsset = '0x536274C4a091D502928ecE0397962c84D9d18b1b'; //SolarChain:usdk
    const usdAssetSymbol ="usdk1";
    const erc20Symbol = "USDK";

    let admin = "0x8fd86498ff1b750454f5187cc0b2bc6b51d0e07b";
    const erc20Abi = [
        {
            "anonymous": false,
            "inputs": [
                {
                    "indexed": true,
                    "internalType": "address",
                    "name": "owner",
                    "type": "address"
                },
                {
                    "indexed": true,
                    "internalType": "address",
                    "name": "spender",
                    "type": "address"
                },
                {
                    "indexed": false,
                    "internalType": "uint256",
                    "name": "value",
                    "type": "uint256"
                }
            ],
            "name": "Approval",
            "type": "event"
        },
        {
            "anonymous": false,
            "inputs": [
                {
                    "indexed": true,
                    "internalType": "address",
                    "name": "_owner",
                    "type": "address"
                },
                {
                    "indexed": true,
                    "internalType": "address",
                    "name": "_recipient",
                    "type": "address"
                },
                {
                    "indexed": true,
                    "internalType": "uint256",
                    "name": "_amount",
                    "type": "uint256"
                }
            ],
            "name": "Burn",
            "type": "event"
        },
        {
            "anonymous": false,
            "inputs": [
                {
                    "indexed": true,
                    "internalType": "address",
                    "name": "_from",
                    "type": "address"
                },
                {
                    "indexed": true,
                    "internalType": "address",
                    "name": "_buyer",
                    "type": "address"
                },
                {
                    "indexed": true,
                    "internalType": "uint256",
                    "name": "_amount",
                    "type": "uint256"
                }
            ],
            "name": "MintTokens",
            "type": "event"
        },
        {
            "anonymous": false,
            "inputs": [
                {
                    "indexed": true,
                    "internalType": "address",
                    "name": "previousOwner",
                    "type": "address"
                },
                {
                    "indexed": true,
                    "internalType": "address",
                    "name": "newOwner",
                    "type": "address"
                }
            ],
            "name": "OwnershipTransferred",
            "type": "event"
        },
        {
            "anonymous": false,
            "inputs": [
                {
                    "indexed": true,
                    "internalType": "bytes32",
                    "name": "role",
                    "type": "bytes32"
                },
                {
                    "indexed": true,
                    "internalType": "bytes32",
                    "name": "previousAdminRole",
                    "type": "bytes32"
                },
                {
                    "indexed": true,
                    "internalType": "bytes32",
                    "name": "newAdminRole",
                    "type": "bytes32"
                }
            ],
            "name": "RoleAdminChanged",
            "type": "event"
        },
        {
            "anonymous": false,
            "inputs": [
                {
                    "indexed": true,
                    "internalType": "bytes32",
                    "name": "role",
                    "type": "bytes32"
                },
                {
                    "indexed": true,
                    "internalType": "address",
                    "name": "account",
                    "type": "address"
                },
                {
                    "indexed": true,
                    "internalType": "address",
                    "name": "sender",
                    "type": "address"
                }
            ],
            "name": "RoleGranted",
            "type": "event"
        },
        {
            "anonymous": false,
            "inputs": [
                {
                    "indexed": true,
                    "internalType": "bytes32",
                    "name": "role",
                    "type": "bytes32"
                },
                {
                    "indexed": true,
                    "internalType": "address",
                    "name": "account",
                    "type": "address"
                },
                {
                    "indexed": true,
                    "internalType": "address",
                    "name": "sender",
                    "type": "address"
                }
            ],
            "name": "RoleRevoked",
            "type": "event"
        },
        {
            "anonymous": false,
            "inputs": [
                {
                    "indexed": true,
                    "internalType": "address",
                    "name": "from",
                    "type": "address"
                },
                {
                    "indexed": true,
                    "internalType": "address",
                    "name": "to",
                    "type": "address"
                },
                {
                    "indexed": false,
                    "internalType": "uint256",
                    "name": "value",
                    "type": "uint256"
                }
            ],
            "name": "Transfer",
            "type": "event"
        },
        {
            "anonymous": false,
            "inputs": [
                {
                    "indexed": true,
                    "internalType": "address",
                    "name": "_owner",
                    "type": "address"
                },
                {
                    "indexed": true,
                    "internalType": "address",
                    "name": "_recipient",
                    "type": "address"
                },
                {
                    "indexed": true,
                    "internalType": "uint256",
                    "name": "_amount",
                    "type": "uint256"
                }
            ],
            "name": "TransferDirectSolar",
            "type": "event"
        },
        {
            "anonymous": false,
            "inputs": [
                {
                    "indexed": true,
                    "internalType": "address",
                    "name": "_from",
                    "type": "address"
                },
                {
                    "indexed": true,
                    "internalType": "address",
                    "name": "_recipient",
                    "type": "address"
                },
                {
                    "indexed": true,
                    "internalType": "uint256",
                    "name": "_amount",
                    "type": "uint256"
                }
            ],
            "name": "TransferFrom",
            "type": "event"
        },
        {
            "inputs": [],
            "name": "DEFAULT_ADMIN_ROLE",
            "outputs": [
                {
                    "internalType": "bytes32",
                    "name": "",
                    "type": "bytes32"
                }
            ],
            "stateMutability": "view",
            "type": "function"
        },
        {
            "inputs": [],
            "name": "MINTER_ROLE",
            "outputs": [
                {
                    "internalType": "bytes32",
                    "name": "",
                    "type": "bytes32"
                }
            ],
            "stateMutability": "view",
            "type": "function"
        },
        {
            "inputs": [
                {
                    "internalType": "address",
                    "name": "owner",
                    "type": "address"
                },
                {
                    "internalType": "address",
                    "name": "spender",
                    "type": "address"
                }
            ],
            "name": "allowance",
            "outputs": [
                {
                    "internalType": "uint256",
                    "name": "",
                    "type": "uint256"
                }
            ],
            "stateMutability": "view",
            "type": "function"
        },
        {
            "inputs": [
                {
                    "internalType": "address",
                    "name": "spender",
                    "type": "address"
                },
                {
                    "internalType": "uint256",
                    "name": "amount",
                    "type": "uint256"
                }
            ],
            "name": "approve",
            "outputs": [
                {
                    "internalType": "bool",
                    "name": "",
                    "type": "bool"
                }
            ],
            "stateMutability": "nonpayable",
            "type": "function"
        },
        {
            "inputs": [
                {
                    "internalType": "address",
                    "name": "account",
                    "type": "address"
                }
            ],
            "name": "balanceOf",
            "outputs": [
                {
                    "internalType": "uint256",
                    "name": "",
                    "type": "uint256"
                }
            ],
            "stateMutability": "view",
            "type": "function"
        },
        {
            "inputs": [
                {
                    "internalType": "address",
                    "name": "account",
                    "type": "address"
                },
                {
                    "internalType": "uint256",
                    "name": "amount",
                    "type": "uint256"
                }
            ],
            "name": "burn",
            "outputs": [],
            "stateMutability": "nonpayable",
            "type": "function"
        },
        {
            "inputs": [],
            "name": "decimals",
            "outputs": [
                {
                    "internalType": "uint8",
                    "name": "",
                    "type": "uint8"
                }
            ],
            "stateMutability": "view",
            "type": "function"
        },
        {
            "inputs": [
                {
                    "internalType": "address",
                    "name": "spender",
                    "type": "address"
                },
                {
                    "internalType": "uint256",
                    "name": "subtractedValue",
                    "type": "uint256"
                }
            ],
            "name": "decreaseAllowance",
            "outputs": [
                {
                    "internalType": "bool",
                    "name": "",
                    "type": "bool"
                }
            ],
            "stateMutability": "nonpayable",
            "type": "function"
        },
        {
            "inputs": [
                {
                    "internalType": "bytes32",
                    "name": "role",
                    "type": "bytes32"
                }
            ],
            "name": "getRoleAdmin",
            "outputs": [
                {
                    "internalType": "bytes32",
                    "name": "",
                    "type": "bytes32"
                }
            ],
            "stateMutability": "view",
            "type": "function"
        },
        {
            "inputs": [
                {
                    "internalType": "address",
                    "name": "operator",
                    "type": "address"
                }
            ],
            "name": "grantMintRole",
            "outputs": [],
            "stateMutability": "nonpayable",
            "type": "function"
        },
        {
            "inputs": [
                {
                    "internalType": "bytes32",
                    "name": "role",
                    "type": "bytes32"
                },
                {
                    "internalType": "address",
                    "name": "account",
                    "type": "address"
                }
            ],
            "name": "grantRole",
            "outputs": [],
            "stateMutability": "nonpayable",
            "type": "function"
        },
        {
            "inputs": [
                {
                    "internalType": "address",
                    "name": "account",
                    "type": "address"
                }
            ],
            "name": "hasMintRole",
            "outputs": [
                {
                    "internalType": "bool",
                    "name": "",
                    "type": "bool"
                }
            ],
            "stateMutability": "view",
            "type": "function"
        },
        {
            "inputs": [
                {
                    "internalType": "bytes32",
                    "name": "role",
                    "type": "bytes32"
                },
                {
                    "internalType": "address",
                    "name": "account",
                    "type": "address"
                }
            ],
            "name": "hasRole",
            "outputs": [
                {
                    "internalType": "bool",
                    "name": "",
                    "type": "bool"
                }
            ],
            "stateMutability": "view",
            "type": "function"
        },
        {
            "inputs": [
                {
                    "internalType": "address",
                    "name": "spender",
                    "type": "address"
                },
                {
                    "internalType": "uint256",
                    "name": "addedValue",
                    "type": "uint256"
                }
            ],
            "name": "increaseAllowance",
            "outputs": [
                {
                    "internalType": "bool",
                    "name": "",
                    "type": "bool"
                }
            ],
            "stateMutability": "nonpayable",
            "type": "function"
        },
        {
            "inputs": [
                {
                    "internalType": "address",
                    "name": "_recipient",
                    "type": "address"
                },
                {
                    "internalType": "uint256",
                    "name": "_amount",
                    "type": "uint256"
                }
            ],
            "name": "mint",
            "outputs": [
                {
                    "internalType": "bool",
                    "name": "",
                    "type": "bool"
                }
            ],
            "stateMutability": "nonpayable",
            "type": "function"
        },
        {
            "inputs": [],
            "name": "name",
            "outputs": [
                {
                    "internalType": "string",
                    "name": "",
                    "type": "string"
                }
            ],
            "stateMutability": "view",
            "type": "function"
        },
        {
            "inputs": [],
            "name": "owner",
            "outputs": [
                {
                    "internalType": "address",
                    "name": "",
                    "type": "address"
                }
            ],
            "stateMutability": "view",
            "type": "function"
        },
        {
            "inputs": [
                {
                    "internalType": "address",
                    "name": "account",
                    "type": "address"
                }
            ],
            "name": "renounceMintRole",
            "outputs": [],
            "stateMutability": "nonpayable",
            "type": "function"
        },
        {
            "inputs": [],
            "name": "renounceOwnership",
            "outputs": [],
            "stateMutability": "nonpayable",
            "type": "function"
        },
        {
            "inputs": [
                {
                    "internalType": "bytes32",
                    "name": "role",
                    "type": "bytes32"
                },
                {
                    "internalType": "address",
                    "name": "account",
                    "type": "address"
                }
            ],
            "name": "renounceRole",
            "outputs": [],
            "stateMutability": "nonpayable",
            "type": "function"
        },
        {
            "inputs": [
                {
                    "internalType": "bytes32",
                    "name": "role",
                    "type": "bytes32"
                },
                {
                    "internalType": "address",
                    "name": "account",
                    "type": "address"
                }
            ],
            "name": "revokeRole",
            "outputs": [],
            "stateMutability": "nonpayable",
            "type": "function"
        },
        {
            "inputs": [
                {
                    "internalType": "uint8",
                    "name": "decimal",
                    "type": "uint8"
                }
            ],
            "name": "setDecimal",
            "outputs": [],
            "stateMutability": "nonpayable",
            "type": "function"
        },
        {
            "inputs": [
                {
                    "internalType": "bytes4",
                    "name": "interfaceId",
                    "type": "bytes4"
                }
            ],
            "name": "supportsInterface",
            "outputs": [
                {
                    "internalType": "bool",
                    "name": "",
                    "type": "bool"
                }
            ],
            "stateMutability": "view",
            "type": "function"
        },
        {
            "inputs": [],
            "name": "symbol",
            "outputs": [
                {
                    "internalType": "string",
                    "name": "",
                    "type": "string"
                }
            ],
            "stateMutability": "view",
            "type": "function"
        },
        {
            "inputs": [],
            "name": "totalSupply",
            "outputs": [
                {
                    "internalType": "uint256",
                    "name": "",
                    "type": "uint256"
                }
            ],
            "stateMutability": "view",
            "type": "function"
        },
        {
            "inputs": [
                {
                    "internalType": "address",
                    "name": "recipient",
                    "type": "address"
                },
                {
                    "internalType": "uint256",
                    "name": "amount",
                    "type": "uint256"
                }
            ],
            "name": "transfer",
            "outputs": [
                {
                    "internalType": "bool",
                    "name": "",
                    "type": "bool"
                }
            ],
            "stateMutability": "nonpayable",
            "type": "function"
        },
        {
            "inputs": [
                {
                    "internalType": "address",
                    "name": "newOwner",
                    "type": "address"
                }
            ],
            "name": "transferOwnership",
            "outputs": [],
            "stateMutability": "nonpayable",
            "type": "function"
        },
        {
            "inputs": [
                {
                    "internalType": "string",
                    "name": "name",
                    "type": "string"
                },
                {
                    "internalType": "string",
                    "name": "symbol",
                    "type": "string"
                },
                {
                    "internalType": "uint8",
                    "name": "decimal",
                    "type": "uint8"
                }
            ],
            "name": "initialize",
            "outputs": [],
            "stateMutability": "nonpayable",
            "type": "function"
        },
        {
            "inputs": [
                {
                    "internalType": "address",
                    "name": "_from",
                    "type": "address"
                },
                {
                    "internalType": "address",
                    "name": "_recipient",
                    "type": "address"
                },
                {
                    "internalType": "uint256",
                    "name": "_amount",
                    "type": "uint256"
                }
            ],
            "name": "transferFrom",
            "outputs": [
                {
                    "internalType": "bool",
                    "name": "",
                    "type": "bool"
                }
            ],
            "stateMutability": "nonpayable",
            "type": "function"
        }
    ];


    beforeEach(async function () {
        _usdtErc20Instance = await usdtErc20.deployed();
        _bountyPledgeInstance = await bountyPledge.deployed();
        //_bountyPledgeV2Instance = await bountyPledgeV2.deployed();
        console.log("\n+++++++++++++++++++++++\n");
        console.log("_usdtErc20Instance=", _usdtErc20Instance.address);
        console.log("_bountyPledgeInstance=", _bountyPledgeInstance.address);

        let _name = await _bountyPledgeInstance.name();
        console.log("_name= ", _name);
        let _symbol = await _bountyPledgeInstance.symbol();
        console.log("_symbol= ", _symbol);

        //console.log("_bountyPledgeV2Instance=",_bountyPledgeV2Instance.address);

    });
    it('sign-message',async()=>{

        let fromPrivateKey ="406e5d71d5978ebe6818f4489ec6307c7c2ac4eb72ddfa967971b11677583f12";
        console.log("\n签名：web3.eth.accounts.sign ：");
        let digestHash =  web3.utils.keccak256("Login to Kugga\n2022-12-27 10:54:34\n890987687");
        console.log("digestHash=",digestHash);


        let signaturet1 = await web3.eth.accounts.sign(digestHash, fromPrivateKey);
        console.log("signature[accounts.sign]=",signaturet1);
       
        console.log("\n签名：web3.eth.sign:");

        let signaturet4 = await web3.eth.sign(digestHash,from);      
        console.log("signature[eth.sign]=",signaturet4);


    });
 
    it('pledge', async () => {
       
        console.log("pledge:...");
        let allowanceAmount = 3000000; //3 个
        let transferAmountAvg = 1000000; //1个
        //approve
        let tx1 = await _usdtErc20Instance.approve(_bountyPledgeInstance.address, allowanceAmount);
        console.log("tx1.increaseAllowance=", tx1.tx);


        //pledge
        let startTime = Math.floor(new Date().getTime() / 1000);
        let endTime = Math.floor(new Date('2023-1-26 19:45:23').getTime() / 1000);

        /*
            bountyReq ={
            startTime:startTime,
            endTime:endTime,
            bountyType:2,//1=single,2=multiple
            cid:"QM02dfdfierewr2354ler",
            amount:allowance,
            unit:2,//1=usdt,2=matic,
             
        };
        */
        let cid = "QM02dfdfierewr2354ler";
        let bountyReq = [startTime, endTime, 2, cid, allowanceAmount, erc20Symbol];


        let sha3Cid = web3.utils.keccak256(cid);
        let chainSha3Cid =  await _bountyPledgeInstance.getkeccak256(cid);
        console.log("sha3Cid=",sha3Cid);
        console.log("chainSha3Cid-1=",chainSha3Cid);
        console.log("chainSha3Cid-2=",web3.utils.toHex(chainSha3Cid));

        

        let tx2 = await _bountyPledgeInstance.pledge(bountyReq);
        console.log("tx2.pledge=", tx2.tx);

        let bountyId = cid;
        console.log("current.bountyId=", bountyId);

        let bountyStatus1 = await _bountyPledgeInstance.getBounty(bountyId);
        console.log("tx3.bountyStatus-1=", bountyStatus1);



        //分账

        let tx4 = await _bountyPledgeInstance.transfer(bountyId, to, transferAmountAvg, "QM:reply-" + to + "-1");
        console.log("tx4.transfer.1=", tx4.tx);
        let bountyStatus2 = await _bountyPledgeInstance.getBounty(bountyId);
        console.log("tx4.bountyStatus-2=", bountyStatus2);
        //提高赏金
        console.log("提高赏金:增量");
        let tx5 = await _usdtErc20Instance.increaseAllowance(_bountyPledgeInstance.address, allowanceAmount);
        console.log("tx5.increaseAllowance=", tx5.tx);
        let tx6 = await _bountyPledgeInstance.increaseReward(bountyId, allowanceAmount);
        console.log("tx6.increaseReward=", tx6.tx);

        let bountyStatus3 = await _bountyPledgeInstance.getBounty(bountyId);
        console.log("tx6.bountyStatus-3=", bountyStatus3);


        let tx7 = await _bountyPledgeInstance.transfer(bountyId, user2, transferAmountAvg, "QM:reply-" + user2 + "-1");
        console.log("tx7.transfer.2=", tx7.tx);

        let winners = await _bountyPledgeInstance.getWinners(bountyId);
        console.log("winners=", winners);

        //模拟过期，这个method 在正式部署的时候，将废弃
        let expiredTime = Math.floor(new Date().getTime() / 1000);
        let tx8 = await _bountyPledgeInstance.setBountyExpired(bountyId, expiredTime);
        console.log("tx8.setBountyExpired=", tx8.tx);
        let baseAssetBountyId = bountyId;
        /*
                let tx9 = await _bountyPledgeInstance.withdraw(bountyId);
                console.log("tx9.withdraw=", tx9.tx);
        */
        let bountyStatus = await _bountyPledgeInstance.getBounty(bountyId);
        console.log("tx10.bountyStatus=", bountyStatus);


        console.log("\n\n************ Now add supported asset coin **********\n\n");

        let tx20 = await _bountyPledgeInstance.addSupporteAsset(usdAsset, usdAssetSymbol);
        console.log("tx20.change.coin=", tx20.tx);

        let supportedAsset = await _bountyPledgeInstance.getSupportedAsset(usdAsset);
        console.log("supportedAsset=", supportedAsset);
        console.log("supportedAsset.id=", supportedAsset.id);
        console.log("supportedAsset.symbol=", supportedAsset.symbol);

        var usdContract = new web3.eth.Contract(erc20Abi, usdAsset, { from: admin });
        console.log("newAssetCoin : approve is begin....");

        await usdContract.methods.approve(_bountyPledgeInstance.address, allowanceAmount).send({ from: admin })
            .on('transactionHash', function (hash) {
                console.log("hash=", hash);
            })
            .on('receipt', function (receipt) {
                console.log("receipt.blockHash=", receipt.blockHash, ";receipt.transactionHash=", receipt.transactionHash);
            })
            /*.on('confirmation', function (confirmationNumber, receipt) {
              console.log("confirmationNumber=", confirmationNumber, ";receipt.transactionHash=", receipt.transactionHash,";receipt.blockHash=",receipt.blockHash);
            })
            */
            .on('error', function (error, receipt) {
                console.log("error=", error, ";receipt=", receipt);
            });

        console.log("newAssetCoin : approve is over....");

        //pledge
        startTime = Math.floor(new Date().getTime() / 1000);
        endTime = Math.floor(new Date('2023-1-26 19:45:23').getTime() / 1000);

        let cid2 = "QM02dfdfierewqsdz65789abler";
        bountyReq = [startTime, endTime, 2, cid2, allowanceAmount, supportedAsset.symbol];

        let tx21 = await _bountyPledgeInstance.pledge(bountyReq);
        console.log("tx21.pledge=", tx21.tx);

        //分账
        bountyId = cid2;
        let bountyStatus21 = await _bountyPledgeInstance.getBounty(bountyId);
        console.log("bountyStatus21",bountyStatus21);
        let tx22 = await _bountyPledgeInstance.transfer(bountyId, to, transferAmountAvg, "QM:reply-" + to + "-1");
        console.log("tx22.transfer.1=", tx22.tx);
        let tx23 = await _bountyPledgeInstance.transfer(bountyId, user2, transferAmountAvg, "QM:reply-" + user2 + "-1");
        console.log("tx23.transfer.2=", tx23.tx);

        let bountyStatus22 = await _bountyPledgeInstance.getBounty(bountyId);
        console.log("bountyStatus21",bountyStatus22);
        winners = await _bountyPledgeInstance.getWinners(bountyId);
        console.log("winners=", winners);

        //模拟过期，这个method 在正式部署的时候，将废弃
        expiredTime = Math.floor(new Date().getTime() / 1000);
        let tx24 = await _bountyPledgeInstance.setBountyExpired(bountyId, expiredTime);
        console.log("tx24.setBountyExpired=", tx24.tx);


        console.log("\n\n**********统一提款:各提各的********\n\n");

        bountyStatus = await _bountyPledgeInstance.getBounty(baseAssetBountyId);
        console.log("tx29.0.bountyStatus=", bountyStatus);
        bountyStatus = await _bountyPledgeInstance.getBounty(bountyId);
        console.log("tx29.0.bountyStatus=", bountyStatus);

        console.log("\n++++++++withdraw++++++++\n");
        let tx19 = await _bountyPledgeInstance.withdraw(baseAssetBountyId); //原来的交易提取
        console.log("tx19.withdraw=", tx19.tx);

        let tx29 = await _bountyPledgeInstance.withdraw(bountyId); //后来的提取
        console.log("tx29.withdraw=", tx29.tx);

        bountyStatus = await _bountyPledgeInstance.getBounty(baseAssetBountyId);
        console.log("tx29.0.bountyStatus=", bountyStatus);
        bountyStatus = await _bountyPledgeInstance.getBounty(bountyId);
        console.log("tx29.0.bountyStatus=", bountyStatus);

    });
    
    //// other it

});