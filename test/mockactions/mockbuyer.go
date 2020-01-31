package mockactions

import (
	"fmt"
	"github.com/dormao/chaincode_foodmanage/models"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

func RegBuyer(stub *shim.MockStub) peer.Response{
	resp := stub.MockInvoke(models.AllocateIdS(), [][]byte{
		[]byte(models.UnAuthRegisterBuyer),
		[]byte(TestPassword),
	})
	buyer_id = string(resp.GetPayload())
	return resp
}

func BuyerLogin(stub *shim.MockStub) peer.Response{
	resp := stub.MockInvoke(models.AllocateIdS(), [][]byte{
		[]byte(models.UnAuthLogin),
		[]byte(buyer_id),
		[]byte(TestPassword),
	})
	buyer_token = string(resp.GetPayload())
	return resp
}

func BuyerBuy(stub *shim.MockStub) peer.Response{
	resp := stub.MockInvoke(models.AllocateIdS(), [][]byte{
		[]byte(models.OPERATE_PURCHASE),
		[]byte(createCredentialsWithToken(buyer_token)),
		[]byte(product_id),
		[]byte(fmt.Sprint(MockBuyCount)),
	})
	transaction_id = string(resp.GetPayload())
	return resp
}

func BuyerConfirm(stub *shim.MockStub) peer.Response{
	resp := stub.MockInvoke(models.AllocateIdS(), [][]byte{
		[]byte(models.OPERATE_CONFIRM),
		[]byte(createCredentialsWithToken(buyer_token)),
		[]byte(transaction_id),
	})
	return resp
}

func BuyerCancelTransaction(stub *shim.MockStub) peer.Response{
	resp := stub.MockInvoke(models.AllocateIdS(), [][]byte{
		[]byte(models.OPERATE_CANCELORDER),
		[]byte(createCredentialsWithToken(buyer_token)),
		[]byte(transaction_id),
	})
	return resp
}