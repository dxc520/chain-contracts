# dev-nft

## 1、合约介绍
本合约遵从 以太坊 ERC721 规范，基于 `openzeppelin`实现NFT相关功能。主要实现合约铸币、转账等交易，以及相关查询。

**合约 开发：** 

- 使用solidity 开发，在 v0.8.0 及以上版本
- 使用truffle 开发框架：编译开发，部署上链
- 使用node 12.8.0版本及以上
- 测试网络使用 rinkeby ，也可以使用 ropsten
- 基数设施使用infura.
- 在solidity中，集成开发，需要安装钱包：`ruffle-hdwallet-provider`
- 在浏览器插件钱包：MetaMask
- 使用golang 开发与合约交互.使用`gEth`工具包生成abi接口以及合约操作代码包。



**install denpen**

- `npm install @openzeppelin/contracts-upgradeable`
- `npm install @truffle/hdwallet-provider`

## 2、升级改造
### 2.1 没有设计模式
遵从ERC721 规范，直接在`openzeppelin`的基础上开发合约。增加自己的铸币逻辑和构造函数参数化等。即可以实现自己的 NFT 铸币等基本的操作。

缺点是： 每次变动，都需要重新部署新合约，因此改动后的合约与之前没有关联性。也即：不可升级。没有延续性


### 2.2 多合约升级模式
使用一定的设计模式，将原本的单合约，按照设计模式拆成多合约。这个模式简称为CD（Controller-Data）模式。将合约分为两类：控制器合约（Controller Contract）与数据合约（Data Contract)。


数据合约 与控制器合约，按照一定的关联关系，依赖注入访问。如下图：

![](./architecture.png)


**具体来说：**



## 3、交易代付





## 4、生成 abi 接口以及 golang 应用接口
- 1. 先 `truffle compile --all` 编译，生成abi 接口。
- 2. 拷贝上面的abi 数据，形成abi文件
- 3. 生成golang 文件。`abigen --abi=Store_sol_Store.abi --pkg=store --out=Store.go`

   
    ```
    abigen --abi=ArtExhibitionERC1155.abi --pkg=ArtExhibition --out=ArtExhibition.go
    abigen --abi=WalletProxyUpgradeable.abi --pkg=WalletProxy --out=WalletProxy.go
    ```

- 4. 修改golang 文件中的有问题的语法。

[参考资料](https://goethereumbook.org/zh/smart-contract-compile/) 





## 合约测试
- 1. 启动容器，进入交互模式。（修改本地挂载路径，映射本地项目空间）

```
docker run -it -v /Users/xcdong/resource/project/Ethereum/contract/dev-nft/:/dev-nft/ trailofbits/eth-security-toolbox
```

- 2. 进入容器，执行命令 

```
slither --solc-remaps @openzeppelin/contracts/=./node_modules/@openzeppelin/contracts/  --filter-paths  "node_modules" contracts/chain_disk.sol
```

