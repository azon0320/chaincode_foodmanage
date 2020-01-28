package store

import (
	"encoding/json"
	"github.com/dmao/gome/chaincode/chaincode_foodmanage/models"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

func GetProductById(pid string, stub shim.ChaincodeStubInterface) (*models.Product, error) {
	dat, err := stub.GetState(pid)
	if err != nil {
		return nil, err
	}
	prod := &models.Product{}
	err = json.Unmarshal(dat, prod)
	if err != nil {
		return nil, err
	}
	return prod, nil
}

func SaveProduct(product *models.Product, stub shim.ChaincodeStubInterface) error {
	return SaveObjectByJson(product.Id, product, stub)
}
