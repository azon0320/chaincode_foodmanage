package actions

import (
	"errors"
	"github.com/dmao/gome/chaincode/chaincode_foodmanage/models"
	"github.com/dmao/gome/chaincode/chaincode_foodmanage/store"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

/*
 * TODO Avoid Compile in Windows because of the gcc
 * TODO Be sure the code is tagged when compile in Windows
 */
// -----  BEGIN ORDER  -----
func CancelTransport(tspr *models.Transporter, torder *models.TransportOrder, stub shim.ChaincodeStubInterface) error {
	if torder.TransporterId != tspr.Id {
		return errors.New("not the transport order owner")
	}
	if torder.OrderWasted {
		return errors.New("order has been wasted")
	}
	torder.OrderWasted = true
	return store.SaveTransportOrder(torder, stub)
}

func UpdateDetails(tspr *models.Transporter, torder *models.TransportOrder, newer *models.TransportDetails, stub shim.ChaincodeStubInterface) error {
	tsac, err := store.GetTransactionOrderById(torder.TransactionId, stub)
	if err != nil {
		return err
	}
	if torder.TransporterId != tspr.Id {
		return errors.New("not the transport order owner")
	} else if tsac.OrderStatus != models.TRANSACTION_ORDER_STATUS_TRANSPORTING {
		if tsac.OrderStatus == models.TRANSACTION_ORDER_STATUS_UNTRANSMIT {
			return errors.New("transaction order is not transmitted")
		} else {
			return errors.New("transport has been completed")
		}
	}
	// TODO calculate the temperature offset for extra cost
	torder.Details = newer
	return store.SaveTransportOrder(torder, stub)
}

func CompleteTransport(tspr *models.Transporter, torder *models.TransportOrder, stub shim.ChaincodeStubInterface) error {
	tsac, err := store.GetTransactionOrderById(torder.TransactionId, stub)
	if err != nil {
		return err
	}
	if torder.TransporterId != tspr.Id {
		return errors.New("not the transport order owner")
	} else if tsac.OrderStatus != models.TRANSACTION_ORDER_STATUS_TRANSPORTING {
		if tsac.OrderStatus == models.TRANSACTION_ORDER_STATUS_UNTRANSMIT {
			return errors.New("transaction order is not transmitted")
		} else {
			return errors.New("transport has been completed")
		}
	}
	tsac.OrderStatus = models.TRANSACTION_ORDER_STATUS_UNCONFIRM
	return store.SaveTransactionOrder(tsac, stub)
}

func RegisterTransporter(password string, initialBalance uint64, stub shim.ChaincodeStubInterface) (string, error){
	tspr := &models.Transporter{Operator: models.NewOperator(models.OperatorTransporter,password, initialBalance)}
	err := store.SaveTransporter(tspr, stub)
	return tspr.Id, err
}