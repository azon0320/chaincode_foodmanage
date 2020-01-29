package test

import (
	"github.com/dormao/chaincode_foodmanage/test/mockactions"
	"testing"
)

func TestBuyerReg(t *testing.T) {
	mkstub := launchMock()
	r := mockactions.RegBuyer(mkstub)
	t.Log(r.String())
}
