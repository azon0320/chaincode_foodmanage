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
	}
	operator.LastAttempts = models.CurrentTimeMillis()
	err = store.SaveOperator(operator, stub)
	if err == nil {
		return operator, nil
	} else {
		return nil, err
	}
}
