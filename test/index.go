package test

import (
	"fmt"
	"github.com/dormao/chaincode_foodmanage/chaincode"
	"github.com/dormao/chaincode_foodmanage/models"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
	"net/http"
	"strconv"
	"testing"
)

func launchMock(arg ...uint64) *shim.MockStub{
	var foodcc *chaincode.FoodManageChaincode = new(chaincode.FoodManageChaincode)
	var stub *shim.MockStub = shim.NewMockStub("testStub", foodcc)
	var bal uint64
	if len(arg) > 0 {
		bal = arg[0]
	}else {
		bal = models.DefaultInitialBalance
	}
	stub.MockInit(models.AllocateIdS(), [][]byte{
		[]byte("unused function name"),
		[]byte(strconv.Itoa(int(bal))),
	})
	return stub
}
func LaunchMock(arg ...uint64) *shim.MockStub{
	return launchMock(arg...)
}

func LogResponse(response *peer.Response, t *testing.T){
	fmt.Println(response.String())
	if response.Status == http.StatusOK {
		t.Log(response.String())
	}else {
		t.Error(response.String())
	}
}