package mockactions

import (
	"encoding/json"
	"fmt"
	"github.com/dormao/chaincode_foodmanage/models"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

func RegTransporter(stub *shim.MockStub) peer.Response{
	resp := stub.MockInvoke(models.AllocateIdS(), [][]byte{
		[]byte(models.UnAuthRegisterTransporter),
		[]byte(TestPassword),
	})
	returns := &models.DataReturns{};json.Unmarshal(resp.Payload, returns)
	Mocktransporter_id = fmt.Sprint(returns.Data)
	return resp
}

func TransporterLogin(stub *shim.MockStub) peer.Response{
	resp := stub.MockInvoke(models.AllocateIdS(), [][]byte{
		[]byte(models.UnAuthLogin),
		[]byte(fmt.Sprint(models.OperatorTransporter)),
		[]byte(Mocktransporter_id),
		[]byte(TestPassword),
	})
	returns := &models.DataReturns{};json.Unmarshal(resp.Payload, returns)
	Mocktransporter_token = fmt.Sprint(returns.Data)
	return resp
}

func TransporterUpdateTemperature(stub *shim.MockStub) peer.Response{
	resp := stub.MockInvoke(models.AllocateIdS(), [][]byte{
		[]byte(models.OPERATE_UPDATE_TRANSPORT),
		[]byte(createCredentialsWithToken(Mocktransporter_token)),
		[]byte(Mocktransport_id),
		jsonEncode(&models.TransportDetails{
			Temperature: MockTransportTemperature,
		}),
	})
	return resp
}

func TransporterCancelTransport(stub *shim.MockStub) peer.Response{
	resp := stub.MockInvoke(models.AllocateIdS(), [][]byte{
		[]byte(models.OPERATE_CANCELTRANSPORT),
		[]byte(createCredentialsWithToken(Mocktransporter_token)),
		[]byte(Mocktransport_id),
	})
	return resp
}

func TransporterCompleteTransport(stub *shim.MockStub) peer.Response{
	resp := stub.MockInvoke(models.AllocateIdS(), [][]byte{
		[]byte(models.OPERATE_COMPLETE_TRANSPORT),
		[]byte(createCredentialsWithToken(Mocktransporter_token)),
		[]byte(Mocktransport_id),
	})
	return resp
}
