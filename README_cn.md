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
    > --name mycc \
    > --version 1.0 \
    > --path github.com/dormao/chaincode_foodmanage \
    > --lang golang
```

## 参数与指令
### 指令表
| 功能 | 方法名 | 参数 | 返回值 |
| - | - | - | - |
| 账户登录 | login | (账户ID) (账户密码) | 账户令牌 |
| 注册货商 | reg_seller | (密码) | 货商ID |
| 货商添加商品 | add_prod | (Json格式的身份信息) | 产品ID |
| 货商更新商品信息 | update_prod | (Json格式的身份信息) (Json格式的产品信息) | |
| 货商上架商品 | sell_on | (Json格式的身份信息) (产品ID) ||
| 货商下架商品 | sell_off | (Json格式的身份信息) (产品ID) ||
| 货商发货 | transmit | (Json格式的身份信息) (订单ID) (运输商ID) | 运输订单ID |
| 货商取消订单 | cancel_order | (Json格式的身份信息) (订单ID) ||
| 注册买方 | reg_buyer | (密码) | 买方ID |
| 买方购买商品 | buy_prod | (Json格式的身份信息) (产品ID) (数量) | 订单ID |
| 买方确认收货 | confirm | (Json格式的身份信息) (订单ID) ||
| 买方取消订单 | cancel_order | (Json格式的身份信息) (订单ID) ||
| 注册运输商 | reg_transporter | (密码) | 运输商ID |
| 运输商取消订单 | cancel_transport | (Json格式的身份信息) (运输订单ID) ||
| 运输商更新运输状态 | update_transport | (Json格式的身份信息) (运输订单ID) (Json格式的状态信息) ||
| 运输商到达买家 | complete_transport | (Json格式的身份信息) (运输订单ID) ||
### 指令示例 (docker-cli)
* 注册一个货商
~~~
 {"Function": "reg_seller" , "Args": ["sellerpassword"]}
~~~
* 下架产品
~~~
 # 购买一个下架的产品会得到一个错误
 {"Function": "sell_off" , "Args": ["{"account_id": "myseller_id" , "password": "myseller_pw"}" , "myproduct_id"]}
~~~
* 用Json身份和Json信息更新产品
~~~
 # 用JSON格式化工具更易理解
 {"Function": "update_prod" , "Args": ["{"account_id": "myseller_id" , "password": "myseller_pw"}", "{"each_price":10 , "description": "mydescription" , "inventory": 199 , "transport_amount": 12 , "specified_temperature": 24}"]}
~~~
* 购买产品，生成订单
~~~
 # 用JSON格式化工具更易理解
 {"Function": "buy_prod" , "Args": ["{"account_id": "mybuyer_id" , "password": "mybuyer_pw"}" , "product_id" , 1]}
~~~
* 更新运输订单的温度状态为26
~~~
 # 用JSON格式化工具更易理解
 {"Function": "update_transport" , "Args": ["Args": ["{"account_id": "transporter_id" , "password": "transporter_pw"}" , "{"temperature": 26}"]}
~~~
* 用JSON工具可以很容易地将结构体转换成JSON信息
```javascript
 // JavaScript 代码
 // JSON转换后 "{"account_id": "myseller_id" , "password": "myseller_password"}"
 credentials = {
    account_id: "myseller_id",
    password: "myseller_password"
 };
 JsonEncoded_Credentials = JSON.stringify(credentials)
```
```golang
 // Golang 代码
 credentials := &Credentials{
    AccountId: "myseller_id",
    Password: "myseller_password"
 }
 JsonEncoded_Credentials , err := json.Marshal(credentials)
```

## 许可证
[MIT](https://opensource.org/licenses/MIT)

## 其他语言
[English](./README.md)