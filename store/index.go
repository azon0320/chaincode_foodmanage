package store

import (
	"encoding/json"
	"github.com/dormao/chaincode_foodmanage/models"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"strconv"
)

const GlobalsTransportIndex = "globals_transport_index"
const GlobalsTransactionIndex = "globals_transaction_index"
const GlobalsProductIndex = "globals_product_index"
const GlobalsBuyerIndex = "globals_buyer_index"
const GlobalsSellerIndex = "globals_seller_index"
const GlobalsTransporterIndex = "globals_transporter_index"

const GlobalsInitialBalanceIndex = "globals_initial_balance_index"

func SaveObjectByJson(key string, obj interface{}, stub shim.ChaincodeStubInterface) error {
	dat, err := json.Marshal(obj)
	if err != nil {
		return err
	}
	return stub.PutState(key, dat)
}

func GetStateInt(key string, defaultVal int, stub shim.ChaincodeStubInterface) (int,error) {
	value, err := stub.GetState(key)
	if err != nil {
		return defaultVal, err
	}
	intValue, err2 := strconv.Atoi(string(value))
	if err2 != nil {
		return defaultVal, err2
	}
	return intValue, nil
}

func SetStateInt(key string, value int, stub shim.ChaincodeStubInterface) error {
	return stub.PutState(key, []byte(strconv.Itoa(value)))
}

func GetOperator(id string, stub shim.ChaincodeStubInterface) *models.Operator {
	dat, err := stub.GetState(id)
	if err != nil {
		return nil
	}
	operator := &models.Operator{}
	err = json.Unmarshal(dat, operator)
	if err != nil {
		return nil
	}
	return operator
}

func SaveOperator(operator *models.Operator, stub shim.ChaincodeStubInterface) error {
	return SaveObjectByJson(operator.Id, operator, stub)
}
