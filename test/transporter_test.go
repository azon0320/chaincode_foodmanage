package test

import (
	"github.com/dormao/chaincode_foodmanage/test/mockactions"
	"testing"
)

func TestTransporterReg(t *testing.T) {
	mkstub := launchMock()
	r := mockactions.RegTransporter(mkstub)
	t.Log(r.String())
}
