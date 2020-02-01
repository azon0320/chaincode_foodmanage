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
	Mockseller_id = fmt.Sprint(returns.Data)
	return resp
}

func SellerLogin(stub *shim.MockStub) peer.Response{
	resp := stub.MockInvoke(models.AllocateIdS(), [][]byte{
		[]byte(models.UnAuthLogin),
		[]byte(fmt.Sprint(models.OperatorSeller)),
		[]byte(Mockseller_id),
		[]byte(TestPassword),
	})
	returns := &models.DataReturns{};json.Unmarshal(resp.Payload, returns)
	Mockseller_token = fmt.Sprint(returns.Data)
	return resp
}

func SellerAddProd(stub *shim.MockStub) peer.Response{
	resp := stub.MockInvoke(models.AllocateIdS(), [][]byte{
		[]byte(models.OPERATE_ADDPRODUCT),
		[]byte(createCredentialsWithToken(Mockseller_token)),
	})
	returns := &models.DataReturns{};json.Unmarshal(resp.Payload, returns)
	//println("Product Id : " + Mockproduct_id)
	Mockproduct_id = fmt.Sprint(returns.Data)
	return resp
}

func SellerUpdateProd(stub *shim.MockStub) peer.Response{
	resp := stub.MockInvoke(models.AllocateIdS(), [][]byte{
		[]byte(models.OPERATE_UPDATE_PRODUCT),
		[]byte(createCredentialsWithToken(Mockseller_token)),
		[]byte(Mockproduct_id),
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
		[]byte(createCredentialsWithToken(Mockseller_token)),
		[]byte(Mockproduct_id),
	})
	return resp
}

func SellerSellOffProd(stub *shim.MockStub) peer.Response{
	resp := stub.MockInvoke(models.AllocateIdS(), [][]byte{
		[]byte(models.OPERATE_TAKEOFFSELL),
		[]byte(createCredentialsWithToken(Mockseller_token)),
		[]byte(Mockproduct_id),
	})
	return resp
}

func SellerTransmitProd(stub *shim.MockStub) peer.Response{
	resp := stub.MockInvoke(models.AllocateIdS(), [][]byte{
		[]byte(models.OPERATE_TRANSMIT),
		[]byte(createCredentialsWithToken(Mockseller_token)),
		[]byte(Mocktransaction_id),
		[]byte(Mocktransporter_id),
	})
	returns := &models.DataReturns{};json.Unmarshal(resp.Payload, returns)
	Mocktransport_id = fmt.Sprint(returns.Data)
	return resp
}

func SellerCancelTransaction(stub *shim.MockStub) peer.Response{
	resp := stub.MockInvoke(models.AllocateIdS(), [][]byte{
		[]byte(models.OPERATE_CANCELORDER),
		[]byte(createCredentialsWithToken(Mockseller_token)),
		[]byte(Mocktransaction_id),
	})
	return resp
}