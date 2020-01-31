package store

import (
	"encoding/json"
	"github.com/dormao/chaincode_foodmanage/models"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

func GetTokenMapByToken(token string, stub shim.ChaincodeStubInterface) (*models.TokenMap, error) {
	dat, err := stub.GetState(token)
	if err != nil {
		return nil,err
	}
	tokenmap := &models.TokenMap{}
	err = json.Unmarshal(dat, tokenmap)
	if err != nil {
		return nil, err
	}
	return tokenmap, nil
}

func SaveTokenMap(tokenMap *models.TokenMap, stub shim.ChaincodeStubInterface) error{
	return SaveObjectByJson(tokenMap.Token, tokenMap, stub)
}

func DeleteTokenMap(tkey string, stub shim.ChaincodeStubInterface) error{
	return stub.DelState(tkey)
}