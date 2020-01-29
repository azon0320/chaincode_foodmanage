package mockactions

import (
	"github.com/dormao/chaincode_foodmanage/chaincode"
	"github.com/dormao/chaincode_foodmanage/models"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

func RegSeller(stub *shim.MockStub) peer.Response{
	resp := stub.MockInvoke(models.AllocateIdS(), [][]byte{
		[]byte(chaincode.UnAuthRegisterSeller),
		[]byte(TestPassword),
	})
	seller_id = string(resp.GetPayload())
	return resp
}

func SellerAddProd(stub *shim.MockStub) peer.Response{
	resp := stub.MockInvoke(models.AllocateIdS(), [][]byte{
		[]byte(chaincode.OPERATE_ADDPRODUCT),
		[]byte(createCredentials(seller_id)),
	})
	product_id = string(resp.GetPayload())
	return resp
}

func SellerUpdateProd(stub *shim.MockStub) peer.Response{
	resp := stub.MockInvoke(models.AllocateIdS(), [][]byte{
		[]byte(chaincode.OPERATE_UPDATE_PRODUCT),
		[]byte(createCredentials(seller_id)),
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
		[]byte(chaincode.OPERATE_TAKEONSELL),
		[]byte(createCredentials(seller_id)),
		[]byte(product_id),
	})
	return resp
}

func SellerSellOffProd(stub *shim.MockStub) peer.Response{
	resp := stub.MockInvoke(models.AllocateIdS(), [][]byte{
		[]byte(chaincode.OPERATE_TAKEOFFSELL),
		[]byte(createCredentials(seller_id)),
		[]byte(product_id),
	})
	return resp
}

func SellerTransmitProd(stub *shim.MockStub) peer.Response{
	resp := stub.MockInvoke(models.AllocateIdS(), [][]byte{
		[]byte(chaincode.OPERATE_TRANSMIT),
		[]byte(createCredentials(seller_id)),
		[]byte(transaction_id),
		[]byte(transporter_id),
	})
	transport_id = string(resp.GetPayload())
	return resp
}

func SellerCancelTransaction(stub *shim.MockStub) peer.Response{
	resp := stub.MockInvoke(models.AllocateIdS(), [][]byte{
		[]byte(chaincode.OPERATE_CANCELORDER),
		[]byte(createCredentials(seller_id)),
		[]byte(transaction_id),
	})
	return resp
}