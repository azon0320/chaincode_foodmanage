package mockactions

import (
	"encoding/json"
	"github.com/dormao/chaincode_foodmanage/models"
)

const TestPassword = "testpwd"

var(
	Mockseller_id = "none"
	Mockseller_token = ""
	Mockbuyer_id = "none"
	Mockbuyer_token = ""
	Mocktransporter_id = "none"
	Mocktransporter_token = ""

	Mockproduct_id = "none"
	Mocktransaction_id = "none"
	Mocktransport_id = "none"

	MockBuyCount uint32 = 86

	MockTransportTemperature byte = 26
	MockProdEachPrice uint64 = 10
	MockProdDescription string = "Prod Description"
	MockProdInventory uint32 = 87
	MockProdTransportAmount uint64 = 10
	MockProdTemperature byte = 24
)

func createCredentials(id string) []byte{
	dat, _ := json.Marshal(&models.Credentials{
		AccountId: id,
		Password: TestPassword,
	})
	return dat
}

func createCredentialsWithToken(token string) []byte{
	dat, _ := json.Marshal(&models.Credentials{
		Token: token,
	})
	return dat
}

func jsonEncode(v interface{}) []byte{
	dat, _ := json.Marshal(v)
	return dat
}