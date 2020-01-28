package store

import (
	"encoding/json"
	"github.com/dmao/gome/chaincode/chaincode_foodmanage/models"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

func GetTransactionOrderById(pid string, stub shim.ChaincodeStubInterface) (*models.TransactionOrder, error) {
	dat, err := stub.GetState(pid)
	if err != nil {
		return nil, err
	}
	prod := &models.TransactionOrder{}
	err = json.Unmarshal(dat, prod)
	if err != nil {
		return nil, err
	}
	return prod, nil
}

func SaveTransactionOrder(order *models.TransactionOrder, stub shim.ChaincodeStubInterface) error {
	return SaveObjectByJson(order.Id, order, stub)
}
