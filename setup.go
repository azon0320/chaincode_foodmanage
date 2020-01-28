package main

import (
	"fmt"
	"github.com/dmao/gome/chaincode/chaincode_foodmanage/chaincode"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

func main() {
	//TODO Setup Code
	//shim.SetLoggingLevel(shim.LogDebug)
	err := shim.Start(new(chaincode.FoodManageChaincodeV1))
	if err != nil {
		fmt.Printf("Error Starting FoodManageChaincode v1 %s", err)
	}
}
