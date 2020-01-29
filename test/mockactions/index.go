package mockactions

import (
	"encoding/json"
	"github.com/dormao/chaincode_foodmanage/models"
)

const TestPassword = "testpwd"

var(
	seller_id = "none"
	buyer_id = "none"
	transporter_id = "none"

	product_id = "none"
	transaction_id = "none"
	transport_id = "none"

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

func jsonEncode(v interface{}) []byte{
	dat, _ := json.Marshal(v)
	return dat
}