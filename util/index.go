package util

import (
	"encoding/json"
	"github.com/dmao/gome/chaincode/chaincode_foodmanage/models"
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
