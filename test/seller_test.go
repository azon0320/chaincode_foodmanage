package test

import (
	"github.com/dormao/chaincode_foodmanage/test/mockactions"
	"testing"
)

func TestSellerReg(t *testing.T) {
	resp := mockactions.RegSeller(launchMock())
	t.Log(resp.String())
}

func TestSellerAddProd(t *testing.T) {
	mkstub := launchMock()
	mockactions.RegSeller(mkstub)
	resp := mockactions.SellerAddProd(mkstub)
	t.Log(resp.String())
}

func TestSellerUpdateProd(t *testing.T) {
	mkstub := launchMock()
	mockactions.RegSeller(mkstub)
	mockactions.SellerAddProd(mkstub)
	resp := mockactions.SellerUpdateProd(mkstub)
	t.Log(resp.String())
}

func TestSellerSellOnProd(t *testing.T) {
	mkstub := launchMock()
	mockactions.RegSeller(mkstub)
	mockactions.SellerAddProd(mkstub)
	resp := mockactions.SellerSellOnProd(mkstub)
	t.Log(resp.String())
}

func TestSellerSellOnProdTwice(t *testing.T) {
	mkstub := launchMock()
	mockactions.RegSeller(mkstub)
	mockactions.SellerAddProd(mkstub)
	resp := mockactions.SellerSellOnProd(mkstub)
	t.Log(resp.String())
	resp = mockactions.SellerSellOnProd(mkstub)
	t.Log(resp.String())
}

func TestSellerSellOffProd(t *testing.T) {
	mkstub := launchMock()
	mockactions.RegSeller(mkstub)
	mockactions.SellerAddProd(mkstub)
	mockactions.SellerSellOnProd(mkstub)
	resp := mockactions.SellerSellOffProd(mkstub)
	t.Log(resp.String())
}

func TestSellerSellOffProdTwice(t *testing.T) {
	mkstub := launchMock()
	mockactions.RegSeller(mkstub)
	mockactions.SellerAddProd(mkstub)
	resp := mockactions.SellerSellOffProd(mkstub)
	t.Log(resp.String())
}

