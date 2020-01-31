package integration

import (
	"github.com/dormao/chaincode_foodmanage/test"
	"github.com/dormao/chaincode_foodmanage/test/mockactions"
	"github.com/hyperledger/fabric/protos/peer"
	"testing"
)

/*
 * This is the completed transaction cycle
 * The full steps is:
 * 1. Seller , Buyer , Transporter register accounts
 * 2. Seller add && update the product , then set it on sell
 * 3. Buyer purchase the product
 * 4. Seller transmit the product to transporter
 * 5. Transporter update the product details(temperature) while transporting
 * 6. Transporter arrive at the buyer, wait for buyer confirm
 * 7. The Buyer confirm , transaction finished
 */
func TestLifeCycle(t *testing.T) {
	var resp peer.Response
	mkstub := test.LaunchMock()

	resp = mockactions.RegSeller(mkstub)
	test.LogResponse(&resp, t)

	resp = mockactions.RegBuyer(mkstub)
	test.LogResponse(&resp, t)

	resp = mockactions.RegTransporter(mkstub)
	test.LogResponse(&resp, t)

	resp = mockactions.SellerLogin(mkstub)
	test.LogResponse(&resp, t)

	resp = mockactions.BuyerLogin(mkstub)
	test.LogResponse(&resp, t)

	resp = mockactions.TransporterLogin(mkstub)
	test.LogResponse(&resp, t)

	resp = mockactions.SellerAddProd(mkstub)
	test.LogResponse(&resp, t)

	resp = mockactions.SellerUpdateProd(mkstub)
	test.LogResponse(&resp, t)

	resp = mockactions.SellerSellOnProd(mkstub)
	test.LogResponse(&resp, t)

	resp = mockactions.BuyerBuy(mkstub)
	test.LogResponse(&resp, t)

	resp = mockactions.SellerTransmitProd(mkstub)
	test.LogResponse(&resp, t)

	resp = mockactions.TransporterUpdateTemperature(mkstub)
	test.LogResponse(&resp, t)

	resp = mockactions.TransporterUpdateTemperature(mkstub)
	test.LogResponse(&resp, t)

	resp = mockactions.TransporterUpdateTemperature(mkstub)
	test.LogResponse(&resp, t)

	resp = mockactions.TransporterUpdateTemperature(mkstub)
	test.LogResponse(&resp, t)

	resp = mockactions.TransporterCompleteTransport(mkstub)
	test.LogResponse(&resp, t)

	resp = mockactions.BuyerConfirm(mkstub)
	test.LogResponse(&resp, t)

	// You will receive an error of (confirmed order) here
	//resp = mockactions.BuyerConfirm(mkstub)
	//test.LogResponse(&resp, t)
}
