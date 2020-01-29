# Food Transport Chaincode
##### 基于 Hyperledger Fabric v1.0.0 的食品运输管理(智能合约)

## 兼容的测试环境与版本
* Golang (1.10.1)(没有go mod)
* Hyperledger Fabric (v1.0.0)
```
    git clone https://github.com/hyperledger/fabric
    cd fabric
    git checkout v1.0.0
```

## 测试环境部署
一般情况下链代码会部署在网络中，每次修改都要更新代码甚至重启网络

用 模拟区块(Mock Stub) 来做开发环境
```
    var mycc = new(chaincode.FoodManageChaincode{})
    var mockstub = shim.NewMockStub("test stub", mycc)
    // more mock details please go to ./test/
```
模拟区块需要github.com/miekg/pkcs11，而这个库需要安装 libltdl-dev，用下面的方法安装这个 libltdl-dev
```
    相关软件包: libltdl3-dev , libltdl-dev
    用 (yum) || (apt) 这些指令安装
```

## 独立部署
示例：用 E2E_CLI 部署
```
 # 这个指令不包含TLS，关闭节点的TLS请在节点docker-compose的配置里面修改
    [docker-cli]$ peer chaincode install \
    > --name mycc
    > --version 1.0
    > --path github.com/dormao/chaincode_foodmanage
    > --lang golang
```

## 许可证
[MIT](https://opensource.org/licenses/MIT)

## 其他语言
[English](./README.md)