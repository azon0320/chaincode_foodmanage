# Food Transport Chaincode
##### FoodTransport Manager Chaincode on Hyperledger Fabric v1.0.0

## Compatiable Libraries and Environments
* Golang (1.10.1)(no go mod)
* Hyperledger Fabric (v1.0.0)
```
    git clone https://github.com/hyperledger/fabric
    cd fabric
    git checkout v1.0.0
```

## Testing launch
Chaincode will run in STANDALONE, such as docker container.

To test chaincode without restart or updating the network, use MOCK
```
    var mycc = new(chaincode.FoodManageChaincode{})
    var mockstub = shim.NewMockStub("test stub", mycc)
    // more mock details please go to ./test/
```
The library github.com/miekg/pkcs11 needs libltdl-dev which the Mock code includes.
```
    Related packages: libltdl3-dev , libltdl-dev
    Please use (yum) || (apt) to install it
```

## Standalone launch
Use E2E_CLI install the chaincode
```
 # Example with no TLS
    [docker-cli]$ peer chaincode install \
    > --name mycc
    > --version 1.0
    > --path github.com/dormao/chaincode_foodmanage
    > --lang golang
```

## License
[MIT](https://opensource.org/licenses/MIT)

[Chinese](./README_cn.md)