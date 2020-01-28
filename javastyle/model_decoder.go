package javastyle

import (
	"encoding/json"
	"github.com/dormao/chaincode_foodmanage/models"
)

/*
 * No error will be thrown on Java Style Value Returns
 * Generally used on a Released ChainCode
 * Otherwise, use Logger or throw errors directly
 * TODO Add Debug Logger
 */
func ProductFromJson(dat []byte) *models.Product {
	var prod *models.Product = &models.Product{}
	err := json.Unmarshal(dat, prod)
	if err != nil {
		return nil
	}
	return prod
}

func TransportOrderFromJson(dat []byte) *models.TransportOrder {
	var prod *models.TransportOrder = &models.TransportOrder{}
	err := json.Unmarshal(dat, prod)
	if err != nil {
		return nil
	}
	return prod
}

func TransactionOrderFromJson(dat []byte) *models.TransactionOrder {
	var prod *models.TransactionOrder = &models.TransactionOrder{}
	err := json.Unmarshal(dat, prod)
	if err != nil {
		return nil
	}
	return prod
}

func AuthenticateFromJson(dat []byte) *models.Authenticate {
	var prod *models.Authenticate = &models.Authenticate{}
	err := json.Unmarshal(dat, prod)
	if err != nil {
		return nil
	}
	return prod
}

func BalanceHolderFromJson(dat []byte) *models.BalanceHolder {
	var prod *models.BalanceHolder = &models.BalanceHolder{}
	err := json.Unmarshal(dat, prod)
	if err != nil {
		return nil
	}
	return prod
}
