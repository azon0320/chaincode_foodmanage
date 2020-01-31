package auth

import (
	"errors"
	"github.com/dormao/chaincode_foodmanage/models"
	"github.com/dormao/chaincode_foodmanage/store"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"math"
)

const MaxAttemptFails = 8
const AttemptCooldown = .5 * 60 * 1000
const DefaultTokenExpire = 15 * 60 * 1000

func AttemptWithPassword(credentials *models.Credentials, stub shim.ChaincodeStubInterface) (*models.Operator, error) {
	operator := store.GetOperator(credentials.AccountId, stub)
	if operator == nil {
		return nil, errors.New("account not found")
	}
	if operator.AttemptFails > MaxAttemptFails {
		if math.Abs(float64(operator.LastAttempts-models.CurrentTimeMillis())) < AttemptCooldown {
			return nil, errors.New("too much attempts")
		} else {
			operator.AttemptFails = 0
			operator.LastAttempts = 0
		}
	}
	var err error = nil
	if operator.Password != credentials.Password {
		err = errors.New("attempt failed")
		operator.AttemptFails++
	} else {
		if operator.LastLog == 0 {
			// Activate the account
			operator.Frozen = false
		}
		operator.AttemptFails = 0
		operator.LastLog = models.CurrentTimeMillis()
		tokenMap := models.NewTokenMap(
			models.GenerateTokenWithTime(models.GetTxTimeNanos()),
			operator.Id)
		err = store.SaveTokenMap(tokenMap, stub)
		if err != nil {
			return nil,err
		}else{
			operator.Token = tokenMap.Token
		}
	}
	operator.LastAttempts = models.CurrentTimeMillis()
	err = store.SaveOperator(operator, stub)
	if err == nil {
		return operator, nil
	} else {
		return nil, err
	}
}

func AttemptWithToken(credentials *models.Credentials, stub shim.ChaincodeStubInterface) (*models.Operator, error){
	token, err := store.GetTokenMapByToken(credentials.Token, stub)
	if err != nil {
		return nil,err
	}
	if math.Abs(float64(models.CurrentTimeMillis()) - float64(token.CreateTime)) > DefaultTokenExpire {
		_ = store.DeleteTokenMap(token.Token, stub)
		return nil, errors.New("token expired")
	}
	operator := store.GetOperator(token.AccountId, stub)
	if operator == nil {
		return nil, errors.New("account not found")
	}
	return operator, nil
}

func AttemptWithCredentials(credentials *models.Credentials, stub shim.ChaincodeStubInterface) (*models.Operator, error){
	if credentials.Token != "" {
		return AttemptWithToken(credentials, stub)
	}else {
		return AttemptWithPassword(credentials, stub)
	}
}

