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
// -----  BEGIN PRODUCT  -----
func PurchaseProduct(byer *models.Buyer, prod *models.Product, count uint32, stub shim.ChaincodeStubInterface) (string, error) {
	if byer.Frozen {
		return "", errors.New("account has been frozen")
	} else if prod.ShelvesStatus != models.ShelvesstatusOffsell {
		return "", errors.New("product is now discontinued")
	} else if prod.Inventory < count {
		return "", errors.New("no enough inventory")
	} else if count < 1 {
		return "", errors.New("count must be greater than 0")
	}
	productSum := uint64(count) * prod.EachPrice
	if productSum > byer.Balance {
		return "", errors.New("no enough balance")
	}
	var transaction *models.TransactionOrder = models.NewTransactionOrder(
		prod.ProductSnapshot, byer.Id, count)
	prod.Inventory -= count
	byer.Balance -= productSum
	err := store.SaveTransactionOrder(transaction, stub)
	if err != nil {
		return "", err
	}
	err = store.SaveBuyer(byer, stub)
	if err != nil {
		return "", err
	}
	return transaction.Id, store.SaveProduct(prod, stub)
}

// -----  BEGIN TRANSACTION  -----
func ConfirmTransaction(byer *models.Buyer, transactionOrder *models.TransactionOrder, stub shim.ChaincodeStubInterface) error {
	if byer.Frozen {
		return errors.New("account has been frozen")
	} else if transactionOrder.BuyerId != byer.Id {
		return errors.New("not the transport order buyer")
	} else if transactionOrder.OrderStatus != models.TRANSACTION_ORDER_STATUS_UNCONFIRM {
		switch transactionOrder.OrderStatus {
		case models.TRANSACTION_ORDER_STATUS_UNTRANSMIT:
			return errors.New("transaction order is not transmitted")
		case models.TRANSACTION_ORDER_STATUS_TRANSPORTING:
			return errors.New("transaction order is transporting")
		case models.TRANSACTION_ORDER_STATUS_FAILED:
			return errors.New("transaction order has been closed")
		case models.TRANSACTION_ORDER_STATUS_COMPLETED:
			return errors.New("transaction order has been completed")
		default:
			return errors.New("unknown failure")
		}
	}
	transactionOrder.OrderStatus = models.TRANSACTION_ORDER_STATUS_COMPLETED
	seller, err2 := store.GetSellerById(transactionOrder.Snapshot.SellerId, stub)
	if err2 != nil {
		return err2
	}
	prodSum := transactionOrder.Snapshot.EachPrice * uint64(transactionOrder.ProductCount)
	seller.Balance += prodSum
	err := store.SaveTransactionOrder(transactionOrder, stub)
	if err != nil {
		return err
	}
	return store.SaveSeller(seller, stub)
}

func CancelBuyTransaction(byer *models.Buyer, transactionOrder *models.TransactionOrder, stub shim.ChaincodeStubInterface) error {
	if byer.Frozen {
		return errors.New("account has been frozen")
	} else if transactionOrder.BuyerId != byer.Id {
		return errors.New("not the transport order buyer")
	} else if transactionOrder.OrderStatus != models.TRANSACTION_ORDER_STATUS_UNTRANSMIT {
		switch transactionOrder.OrderStatus {
		case models.TRANSACTION_ORDER_STATUS_TRANSPORTING:
			fallthrough
		case models.TRANSACTION_ORDER_STATUS_UNCONFIRM:
			return errors.New("cancel is prohibited after transmitted")
		case models.TRANSACTION_ORDER_STATUS_COMPLETED:
			return errors.New("transaction order has been completed")
		default:
			return errors.New("unknown failure")
		}
	}
	prod, err3 := store.GetProductById(transactionOrder.Snapshot.Id, stub)
	if err3 != nil {
		return err3
	}
	productSum := transactionOrder.Snapshot.EachPrice * uint64(transactionOrder.ProductCount)
	byer.Balance += productSum
	prod.Inventory += transactionOrder.ProductCount
	transactionOrder.OrderStatus = models.TRANSACTION_ORDER_STATUS_FAILED
	err := store.SaveBuyer(byer, stub)
	if err != nil {
		return err
	}
	err = store.SaveProduct(prod, stub)
	if err != nil {
		return err
	}
	return store.SaveTransactionOrder(transactionOrder, stub)
}

func RegisterBuyer(password string, initialBalance uint64, stub shim.ChaincodeStubInterface) (string, error){
	buyer := &models.Buyer{Operator: models.NewOperator(models.OperatorBuyer,password, initialBalance)}
	err := store.SaveBuyer(buyer, stub)
	return buyer.Id, err
}