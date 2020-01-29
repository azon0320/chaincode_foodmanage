package main

import (
	"fmt"
	"github.com/dormao/chaincode_foodmanage/chaincode"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

func main() {
	//TODO Setup Code
	fmt.Println("Start the FoodManage Chaincode")
	err := shim.Start(new(chaincode.FoodManageChaincode))
	if err != nil {
		fmt.Printf("Error Starting FoodManageChaincode v1 %s", err)
	}
}
