package mockactions

import (
	"github.com/dormao/chaincode_foodmanage/chaincode"
	"github.com/dormao/chaincode_foodmanage/models"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

func RegTransporter(stub *shim.MockStub) peer.Response{
	resp := stub.MockInvoke(models.AllocateIdS(), [][]byte{
		[]byte(chaincode.UnAuthRegisterTransporter),
		[]byte(TestPassword),
	})
	transporter_id = string(resp.GetPayload())
	return resp
}

func TransporterUpdateTemperature(stub *shim.MockStub) peer.Response{
	resp := stub.MockInvoke(models.AllocateIdS(), [][]byte{
		[]byte(chaincode.OPERATE_UPDATE_TRANSPORT),
		[]byte(createCredentials(transporter_id)),
		[]byte(transport_id),
		jsonEncode(&models.TransportDetails{
			Temperature: MockTransportTemperature,
		}),
	})
	return resp
}

func TransporterCancelTransport(stub *shim.MockStub) peer.Response{
	resp := stub.MockInvoke(models.AllocateIdS(), [][]byte{
		[]byte(chaincode.OPERATE_CANCELTRANSPORT),
		[]byte(createCredentials(transporter_id)),
		[]byte(transport_id),
	})
	return resp
}

func TransporterCompleteTransport(stub *shim.MockStub) peer.Response{
	resp := stub.MockInvoke(models.AllocateIdS(), [][]byte{
		[]byte(chaincode.OPERATE_COMPLETE_TRANSPORT),
		[]byte(createCredentials(transporter_id)),
		[]byte(transport_id),
	})
	return resp
}
