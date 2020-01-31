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
  > --name mycc \
  > --version 1.0 \
  > --path github.com/dormao/chaincode_foodmanage \
  > --lang golang
```

## Function & Commands
### Command Tables
| Feature | Function | Args | Returns |
| - | - | - | - |
| Account login | login | (AccountId) (Password) | Access Token |
| Register seller | reg_seller | (Password) | Seller id |
| Seller add product | add_prod | (JsonEncoded Credentials) | Product id |
| Seller update product info | update_prod | (JsonEncoded Credentials) (JsonEncoded ProductInfo) | |
| Seller take product online | sell_on | (JsonEncoded Credentials) (Product id) ||
| Seller take product offline | sell_off | (JsonEncoded Credentials) (Product id) ||
| Seller transmit for a transaction | transmit | (JsonEncoded Credentials) (Transaction id) (Transporter id) | Transport order id |
| Seller cancel an exist order | cancel_order | (JsonEncoded Credentials) (Transaction id) ||
| Register buyer | reg_buyer | (Password) | Buyer id |
| Buyer purchase | buy_prod | (JsonEncoded Credentials) (Product id) (count) | Transaction id |
| Buyer confirm order | confirm | (JsonEncoded Credentials) (Transaction id) ||
| Buyer cancel an exist order | cancel_order | (JsonEncoded Credentials) (Transaction id) ||
| Register Transporter | reg_transporter | (Password) | Transporter id |
| Transporter cancel transport order | cancel_transport | (JsonEncoded Credentials) (Transport id) ||
| Transporter update transport data | update_transport | (JsonEncoded Credentials) (Transport id) (JsonEncoded data) ||
| Transporter complete transport | complete_transport | (JsonEncoded Credentials) (Transport id) ||
### Command Sample (CommandLine)
* Register a seller account
~~~
 {"Function": "reg_seller" , "Args": ["sellerpassword"]}
~~~
* Take the product offline
~~~
 # It will return 500 when buy an offline product
 {"Function": "sell_off" , "Args": ["{"account_id": "myseller_id" , "password": "myseller_pw"}" , "myproduct_id"]}
~~~
* Update the product with credentials & new data
~~~
 # Use Json formatter to view structure 
 {"Function": "update_prod" , "Args": ["{"account_id": "myseller_id" , "password": "myseller_pw"}", "{"each_price":10 , "description": "mydescription" , "inventory": 199 , "transport_amount": 12 , "specified_temperature": 24}"]}
~~~
* Buy a product
~~~
 # Use Json formatter to view structure
 {"Function": "buy_prod" , "Args": ["{"account_id": "mybuyer_id" , "password": "mybuyer_pw"}" , "product_id" , 1]}
~~~
* Update transporting products' temperature
~~~
 # Use Json formatter to view structure
 {"Function": "update_transport" , "Args": ["Args": ["{"account_id": "transporter_id" , "password": "transporter_pw"}" , "{"temperature": 26}"]}
~~~
* You can convert the credentials or structured data to strings as arguments
```javascript
 // JavaScript Code
 credentials = {
    account_id: "myseller_id",
    password: "myseller_password"
 };
 JsonEncoded_Credentials = JSON.stringify(credentials)
```
```golang
 // Golang Code
 credentials := &Credentials{
    AccountId: "myseller_id",
    Password: "myseller_password"
 }
 JsonEncoded_Credentials , err := json.Marshal(credentials)
```

## License
[MIT](https://opensource.org/licenses/MIT)

## Other Language
[Chinese](./README_cn.md)