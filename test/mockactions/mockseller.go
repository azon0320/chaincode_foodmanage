package mockactions

import (
	"encoding/json"
	"fmt"
	"github.com/dormao/chaincode_foodmanage/models"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

func RegSeller(stub *shim.MockStub) peer.Response{
	resp := stub.MockInvoke(models.AllocateIdS(), [][]byte{
		[]byte(models.UnAuthRegisterSeller),
		[]byte(TestPassword),
	})
	returns := &models.DataReturns{};json.Unmarshal(resp.Payload, returns)
	seller_id = fmt.Sprint(returns.Data)
	return resp
}

func SellerLogin(stub *shim.MockStub) peer.Response{
	resp := stub.MockInvoke(models.AllocateIdS(), [][]byte{
		[]byte(models.UnAuthLogin),
		[]byte(seller_id),
		[]byte(TestPassword),
	})
	returns := &models.DataReturns{};json.Unmarshal(resp.Payload, returns)
	seller_token = fmt.Sprint(returns.Data)
	return resp
}

func SellerAddProd(stub *shim.MockStub) peer.Response{
	resp := stub.MockInvoke(models.AllocateIdS(), [][]byte{
		[]byte(models.OPERATE_ADDPRODUCT),
		[]byte(createCredentialsWithToken(seller_token)),
	})
	returns := &models.DataReturns{};json.Unmarshal(resp.Payload, returns)
	println("Product Id : " + product_id)
	product_id = fmt.Sprint(returns.Data)
	return resp
}

func SellerUpdateProd(stub *shim.MockStub) peer.Response{
	resp := stub.MockInvoke(models.AllocateIdS(), [][]byte{
		[]byte(models.OPERATE_UPDATE_PRODUCT),
		[]byte(createCredentialsWithToken(seller_token)),
		[]byte(product_id),
		[]byte(jsonEncode(&models.ProductUpdateRequest{
			EachPrice: MockProdEachPrice,
			Description: MockProdDescription,
			Inventory: MockProdInventory,
			TransportAmount: MockProdTransportAmount,
			SpecifiedTemperature: MockProdTemperature,
		})),
	})
	return resp
}

func SellerSellOnProd(stub *shim.MockStub) peer.Response{
	resp := stub.MockInvoke(models.AllocateIdS(), [][]byte{
		[]byte(models.OPERATE_TAKEONSELL),
		[]byte(createCredentialsWithToken(seller_token)),
		[]byte(product_id),
	})
	return resp
}

func SellerSellOffProd(stub *shim.MockStub) peer.Response{
	resp := stub.MockInvoke(models.AllocateIdS(), [][]byte{
		[]byte(models.OPERATE_TAKEOFFSELL),
		[]byte(createCredentialsWithToken(seller_token)),
		[]byte(product_id),
	})
	return resp
}

func SellerTransmitProd(stub *shim.MockStub) peer.Response{
	resp := stub.MockInvoke(models.AllocateIdS(), [][]byte{
		[]byte(models.OPERATE_TRANSMIT),
		[]byte(createCredentialsWithToken(seller_token)),
		[]byte(transaction_id),
		[]byte(transporter_id),
	})
	returns := &models.DataReturns{};json.Unmarshal(resp.Payload, returns)
	transport_id = fmt.Sprint(returns.Data)
	return resp
}

func SellerCancelTransaction(stub *shim.MockStub) peer.Response{
	resp := stub.MockInvoke(models.AllocateIdS(), [][]byte{
		[]byte(models.OPERATE_CANCELORDER),
		[]byte(createCredentialsWithToken(seller_token)),
		[]byte(transaction_id),
	})
	return resp
}