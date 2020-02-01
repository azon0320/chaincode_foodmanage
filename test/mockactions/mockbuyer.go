package mockactions

import (
	"encoding/json"
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
	returns := &models.DataReturns{};json.Unmarshal(resp.Payload, returns)
	Mockbuyer_id = fmt.Sprint(returns.Data)
	return resp
}

func BuyerLogin(stub *shim.MockStub) peer.Response{
	resp := stub.MockInvoke(models.AllocateIdS(), [][]byte{
		[]byte(models.UnAuthLogin),
		[]byte(fmt.Sprint(models.OperatorBuyer)),
		[]byte(Mockbuyer_id),
		[]byte(TestPassword),
	})
	returns := &models.DataReturns{};json.Unmarshal(resp.Payload, returns)
	Mockbuyer_token = fmt.Sprint(returns.Data)
	return resp
}

func BuyerBuy(stub *shim.MockStub) peer.Response{
	resp := stub.MockInvoke(models.AllocateIdS(), [][]byte{
		[]byte(models.OPERATE_PURCHASE),
		[]byte(createCredentialsWithToken(Mockbuyer_token)),
		[]byte(Mockproduct_id),
		[]byte(fmt.Sprint(MockBuyCount)),
	})
	returns := &models.DataReturns{};json.Unmarshal(resp.Payload, returns)
	Mocktransaction_id = fmt.Sprint(returns.Data)
	return resp
}

func BuyerConfirm(stub *shim.MockStub) peer.Response{
	resp := stub.MockInvoke(models.AllocateIdS(), [][]byte{
		[]byte(models.OPERATE_CONFIRM),
		[]byte(createCredentialsWithToken(Mockbuyer_token)),
		[]byte(Mocktransaction_id),
	})
	return resp
}

func BuyerCancelTransaction(stub *shim.MockStub) peer.Response{
	resp := stub.MockInvoke(models.AllocateIdS(), [][]byte{
		[]byte(models.OPERATE_CANCELORDER),
		[]byte(createCredentialsWithToken(Mockbuyer_token)),
		[]byte(Mocktransaction_id),
	})
	return resp
}