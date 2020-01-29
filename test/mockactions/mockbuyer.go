package mockactions

import (
	"fmt"
	"github.com/dormao/chaincode_foodmanage/chaincode"
	"github.com/dormao/chaincode_foodmanage/models"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

func RegBuyer(stub *shim.MockStub) peer.Response{
	resp := stub.MockInvoke(models.AllocateIdS(), [][]byte{
		[]byte(chaincode.UnAuthRegisterBuyer),
		[]byte(TestPassword),
	})
	buyer_id = string(resp.GetPayload())
	return resp
}

func BuyerBuy(stub *shim.MockStub) peer.Response{
	resp := stub.MockInvoke(models.AllocateIdS(), [][]byte{
		[]byte(chaincode.OPERATE_PURCHASE),
		[]byte(createCredentials(buyer_id)),
		[]byte(product_id),
		[]byte(fmt.Sprint(MockBuyCount)),
	})
	transaction_id = string(resp.GetPayload())
	return resp
}

func BuyerConfirm(stub *shim.MockStub) peer.Response{
	resp := stub.MockInvoke(models.AllocateIdS(), [][]byte{
		[]byte(chaincode.OPERATE_CONFIRM),
		[]byte(createCredentials(buyer_id)),
		[]byte(transaction_id),
	})
	return resp
}

func BuyerCancelTransaction(stub *shim.MockStub) peer.Response{
	resp := stub.MockInvoke(models.AllocateIdS(), [][]byte{
		[]byte(chaincode.OPERATE_CANCELORDER),
		[]byte(createCredentials(buyer_id)),
		[]byte(transaction_id),
	})
	return resp
}