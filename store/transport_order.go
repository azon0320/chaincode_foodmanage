package store

import (
	"encoding/json"
	"github.com/dmao/gome/chaincode/chaincode_foodmanage/models"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

func GetTransportOrderById(pid string, stub shim.ChaincodeStubInterface) (*models.TransportOrder, error) {
	dat, err := stub.GetState(pid)
	if err != nil {
		return nil, err
	}
	prod := &models.TransportOrder{}
	err = json.Unmarshal(dat, prod)
	if err != nil {
		return nil, err
	}
	return prod, nil
}

func SaveTransportOrder(order *models.TransportOrder, stub shim.ChaincodeStubInterface) error {
	return SaveObjectByJson(order.Id, order, stub)
}
