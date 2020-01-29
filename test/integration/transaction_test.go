package integration

import (
	"github.com/dormao/chaincode_foodmanage/test"
	"github.com/dormao/chaincode_foodmanage/test/mockactions"
	"github.com/hyperledger/fabric/protos/peer"
	"testing"
)

/*
 * This transaction mocks a scene for this:
 * 1. Seller create a product on sell
 * 2. The inventory is not enough
 */
func TestTransactionBuy01(t *testing.T) {
	mockactions.MockProdInventory = 87
	mockactions.MockBuyCount = 88
	mkstub := test.LaunchMock()
	mockactions.RegSeller(mkstub)
	mockactions.RegBuyer(mkstub)
	mockactions.SellerAddProd(mkstub)
	mockactions.SellerUpdateProd(mkstub)
	mockactions.SellerSellOnProd(mkstub)
	r := mockactions.BuyerBuy(mkstub)
	t.Log(r.String())
}

/*
 * This transaction mocks a scene for this:
 * 1. Seller create a product on sell
 * 2. The balance is not enough
 */
func TestTransactionBuy02(t *testing.T) {
	mockactions.MockProdInventory = 87
	mockactions.MockProdEachPrice = 10
	mockactions.MockBuyCount = 86
	mkstub := test.LaunchMock(50)
	mockactions.RegSeller(mkstub)
	mockactions.RegBuyer(mkstub)
	mockactions.SellerAddProd(mkstub)
	mockactions.SellerUpdateProd(mkstub)
	mockactions.SellerSellOnProd(mkstub)
	r := mockactions.BuyerBuy(mkstub)
	t.Log(r.String())
}

/*
 * This transaction mocks a scene for this:
 * 1. Seller create a product but not on sell
 * 2. Buyer will be prohibited buying
 */
func TestTransactionBuy03(t *testing.T) {
	mkstub := test.LaunchMock()
	mockactions.RegSeller(mkstub)
	mockactions.RegBuyer(mkstub)
	mockactions.SellerAddProd(mkstub)
	mockactions.SellerUpdateProd(mkstub)
	mockactions.SellerSellOffProd(mkstub)
	r := mockactions.BuyerBuy(mkstub)
	t.Log(r.String())
}

/*
 * This transaction mocks a scene for this:
 * 1. Transaction created successfully
 */
func TestTransactionBuy04(t *testing.T) {
	mkstub := test.LaunchMock()
	mockactions.RegSeller(mkstub)
	mockactions.RegBuyer(mkstub)
	mockactions.SellerAddProd(mkstub)
	mockactions.SellerUpdateProd(mkstub)
	mockactions.SellerSellOnProd(mkstub)
	r := mockactions.BuyerBuy(mkstub)
	t.Log(r.String())
}

/*
 * This transaction mocks a scene for this:
 * 1. Transmit but transporter not found
 */
func TestTransactionTransmit01(t *testing.T) {
	mkstub := test.LaunchMock()
	mockactions.RegSeller(mkstub)
	mockactions.RegBuyer(mkstub)
	mockactions.SellerAddProd(mkstub)
	mockactions.SellerUpdateProd(mkstub)
	mockactions.SellerSellOnProd(mkstub)
	mockactions.BuyerBuy(mkstub)
	r := mockactions.SellerTransmitProd(mkstub)
	t.Log(r.String())
}

/*
 * This transaction mocks a scene for this:
 * 1. Transmitted successfully and return the TransportOrder id
 */
func TestTransactionTransmit02(t *testing.T) {
	mkstub := test.LaunchMock()
	mockactions.RegSeller(mkstub)
	mockactions.RegBuyer(mkstub)
	mockactions.RegTransporter(mkstub)
	mockactions.SellerAddProd(mkstub)
	mockactions.SellerUpdateProd(mkstub)
	mockactions.SellerSellOnProd(mkstub)
	mockactions.BuyerBuy(mkstub)
	r := mockactions.SellerTransmitProd(mkstub)
	t.Log(r.String())
}

/*
 * This transaction mocks a scene for this:
 * 1. Transaction created
 * 2. Not transmit
 * 3. Seller close the order
 */
func TestSellerCloseTransaction01(t *testing.T) {
	mkstub := test.LaunchMock()
	mockactions.RegSeller(mkstub)
	mockactions.RegBuyer(mkstub)
	mockactions.RegTransporter(mkstub)
	mockactions.SellerAddProd(mkstub)
	mockactions.SellerUpdateProd(mkstub)
	mockactions.SellerSellOnProd(mkstub)
	mockactions.BuyerBuy(mkstub)
	var resp peer.Response
	resp = mockactions.SellerCancelTransaction(mkstub)
	test.LogResponse(&resp, t)
}

// DEBUGGER
func TestConfirm01(t *testing.T) {
	mkstub := test.LaunchMock()
	mockactions.RegSeller(mkstub)
	mockactions.RegBuyer(mkstub)
	mockactions.RegTransporter(mkstub)
	mockactions.SellerAddProd(mkstub)
	mockactions.SellerUpdateProd(mkstub)
	mockactions.SellerSellOnProd(mkstub)
	mockactions.BuyerBuy(mkstub)
	mockactions.SellerTransmitProd(mkstub)
	var resp peer.Response
	resp = mockactions.TransporterUpdateTemperature(mkstub)
	test.LogResponse(&resp, t)
	mockactions.TransporterCancelTransport(mkstub)
	test.LogResponse(&resp, t)
	resp = mockactions.SellerCancelTransaction(mkstub)
	test.LogResponse(&resp, t)
	resp = mockactions.BuyerCancelTransaction(mkstub)
	test.LogResponse(&resp, t)
	resp = mockactions.BuyerConfirm(mkstub)
	test.LogResponse(&resp, t)
}