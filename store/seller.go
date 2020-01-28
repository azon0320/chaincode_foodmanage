package store

import (
	"encoding/json"
	"github.com/dmao/gome/chaincode/chaincode_foodmanage/models"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

func GetSellerById(pid string, stub shim.ChaincodeStubInterface) (*models.Seller, error) {
	dat, err := stub.GetState(pid)
	if err != nil {
		return nil, err
	}
	prod := &models.Seller{}
	err = json.Unmarshal(dat, prod)
	if err != nil {
		return nil, err
	}
	return prod, nil
}

func SaveSeller(product *models.Seller, stub shim.ChaincodeStubInterface) error {
	return SaveObjectByJson(product.Id, product, stub)
}
