package util

import (
	"encoding/json"
	"fmt"
	"github.com/dormao/chaincode_foodmanage/models"
)

func GetCredentialsFromString(creden string) (*models.Credentials, error) {
	var credentials *models.Credentials = &models.Credentials{}
	err := json.Unmarshal([]byte(creden), credentials)
	if err != nil {
		return nil, err
	}
	return credentials, nil
}

func Index_ArrayS(args []string, needle string) int {
	for key, value := range args {
		if value == needle {
			return key
		}
	}
	return -1
}

func GetTxTimeNanos(sec int64, nano int32) int64{
	return sec * 1e9 + int64(nano)
}

func GetTxTimeMillisS(sec int64, nano int32) string{
	return fmt.Sprint(sec) + fmt.Sprint(nano / 1e6)
}

func GetTxTimeMillis(sec int64, nano int32) int64{
	return GetTxTimeNanos(sec, nano) / 1e6
}