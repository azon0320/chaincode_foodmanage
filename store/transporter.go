package store

import (
	"encoding/json"
	"github.com/dormao/chaincode_foodmanage/models"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

func GetTransporterById(pid string, stub shim.ChaincodeStubInterface) (*models.Transporter, error) {
	dat, err := stub.GetState(pid)
	if err != nil {
		return nil, err
	}
	prod := &models.Transporter{}
	err = json.Unmarshal(dat, prod)
	if err != nil {
		return nil, err
	}
	return prod, nil
}

func SaveTransporter(product *models.Transporter, stub shim.ChaincodeStubInterface) error {
	return SaveObjectByJson(product.Id, product, stub)
}
