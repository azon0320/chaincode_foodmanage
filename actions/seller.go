package actions

import (
	"errors"
	"github.com/dormao/chaincode_foodmanage/models"
	"github.com/dormao/chaincode_foodmanage/store"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

/*
 * TODO Avoid Compile in Windows because of the gcc
 * TODO Be sure the code is tagged when compile in Windows
 */

// -----  BEGIN PRODUCT  -----
func AddProduct(seller *models.Seller, stub shim.ChaincodeStubInterface) (string, error) {
	var prod *models.Product = models.NewProduct(
		seller.Id, 0, models.ShelvesstatusOffsell,
		"", 0, 0, 0)
	return prod.Id, store.SaveProduct(prod, stub)
}

func UpdateProductInfo(seller *models.Seller, prod *models.Product, newer *models.ProductUpdateRequest, stub shim.ChaincodeStubInterface) error {
	if prod.SellerId != seller.Id {
		return errors.New("not the product owner")
	} else if prod.ShelvesStatus == models.ShelvesstatusOnsell {
		return errors.New("product is on sell")
	}
	prod.EachPrice = newer.EachPrice
	prod.Inventory = newer.Inventory
	prod.SpecifiedTemperature = newer.SpecifiedTemperature
	prod.Description = newer.Description
	prod.TransportAmount = newer.TransportAmount
	return store.SaveProduct(prod, stub)
}

func TakeOnSellProduct(seller *models.Seller, prod *models.Product, stub shim.ChaincodeStubInterface) error {
	if prod.SellerId != seller.Id {
		return errors.New("not the product owner")
	}
	if prod.ShelvesStatus == models.ShelvesstatusOnsell {
		return errors.New("already on sell")
	}
	prod.ShelvesStatus = models.ShelvesstatusOnsell
	return store.SaveProduct(prod, stub)
}

func TakeOffSellProduct(seller *models.Seller, prod *models.Product, stub shim.ChaincodeStubInterface) error {
	if prod.SellerId != seller.Id {
		return errors.New("not the product owner")
	}
	if prod.ShelvesStatus == models.ShelvesstatusOffsell {
		return errors.New("already off sell")
	}
	prod.ShelvesStatus = models.ShelvesstatusOffsell
	return store.SaveProduct(prod, stub)
}

// -----  BEGIN TRANSACTION  -----
func TransmitOrder(
	seller *models.Seller, tsac *models.TransactionOrder, tspr *models.Transporter,
	stub shim.ChaincodeStubInterface) (string, error) {
	if tsac.Snapshot.SellerId != seller.Id {
		return "", errors.New("not the product owner")
	}
	if tsac.OrderStatus != models.TRANSACTION_ORDER_STATUS_UNTRANSMIT {
		return "", errors.New("target is transmitted")
	}
	if tsac.TransportOrderId != "" {
		// 未发往运输却出现了运输订单
		//TODO WARNING
	}
	torder := models.NewTransportOrder(
		tspr.Id, tsac.Id, tsac.Snapshot,
		&models.TransportDetails{
			Temperature: tsac.Snapshot.SpecifiedTemperature,
		})
	tsac.TransportOrderId = torder.Id
	tsac.OrderStatus = models.TRANSACTION_ORDER_STATUS_TRANSPORTING
	err := store.SaveTransactionOrder(tsac, stub)
	if err != nil {
		return "", err
	}
	return torder.Id, store.SaveTransportOrder(torder, stub)
}

func CancelSellTransaction(seller *models.Seller, order *models.TransactionOrder, stub shim.ChaincodeStubInterface) error {
	if order.Snapshot.SellerId != seller.Id {
		return errors.New("not the product owner")
	}
	if order.TransportOrderId != "" {
		transportOrder, _ := store.GetTransportOrderById(order.TransportOrderId, stub)
		if transportOrder != nil && !transportOrder.OrderWasted {
			return errors.New("order has been transmitted")
		}
	}
	// TODO pay back to the buyer
	// TODO pay back to the transporter
	order.OrderStatus = models.TRANSACTION_ORDER_STATUS_FAILED
	return store.SaveTransactionOrder(order, stub)
}

func RegisterSeller(password string, initialBalance uint64, stub shim.ChaincodeStubInterface) (string, error){
	seller := &models.Seller{Operator: models.NewOperator(models.OperatorBuyer,password, initialBalance)}
	err := store.SaveSeller(seller, stub)
	return seller.Id, err
}