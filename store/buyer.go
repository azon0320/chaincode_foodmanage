package store

import (
	"encoding/json"
	"github.com/dmao/gome/chaincode/chaincode_foodmanage/models"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

func GetBuyerById(pid string, stub shim.ChaincodeStubInterface) (*models.Buyer, error) {
	dat, err := stub.GetState(pid)
	if err != nil {
		return nil, err
	}
	prod := &models.Buyer{}
	err = json.Unmarshal(dat, prod)
	if err != nil {
		return nil, err
	}
	return prod, nil
}

func SaveBuyer(product *models.Buyer, stub shim.ChaincodeStubInterface) error {
	return SaveObjectByJson(product.Id, product, stub)
}
