package chaincode

import (
	"fmt"
	"github.com/dormao/chaincode_foodmanage/actions"
	"github.com/dormao/chaincode_foodmanage/models"
	"github.com/dormao/chaincode_foodmanage/models/consts"
	"github.com/dormao/chaincode_foodmanage/store"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"strconv"
)

func (ctx *FoodManageChaincode) processBuyerInvoke(buyer *models.Buyer, fcn string, args []string, stub shim.ChaincodeStubInterface) *models.DataReturns {
	switch fcn {
	case models.OPERATE_PURCHASE:
		Usage := fmt.Sprintf("Usage : %s <Credentials> <ProductId> <PurchaseCount>", fcn)
		if len(args) < 3 {
			return models.WithError(consts.CodeErrorParams, Usage)
		}
		prod, err := store.GetProductById(args[1], stub)
		if err != nil {
			return models.WithError(consts.CodeErrorProdNotFound, consts.MsgErrorProdNotFound)
		}
		count, err := strconv.Atoi(args[2])
		if err != nil {
			return models.WithError(consts.CodeErrorParams, Usage)
		}
		tsacId, err := actions.PurchaseProduct(buyer, prod, uint32(count), stub)
		if err != nil {
			return models.WithError(consts.CodeErrorOperationFail, err.Error())
		}
		return models.WithSuccess(consts.MsgOK, tsacId)
	case models.OPERATE_CONFIRM:
		Usage := fmt.Sprintf("Usage : %s <Credentials> <TransactionId>", fcn)
		if len(args) < 2 {
			return models.WithError(consts.CodeErrorParams, Usage)
		}
		tsac, err := store.GetTransactionOrderById(args[1], stub)
		if err != nil {
			return models.WithError(consts.CodeErrorTransactionNotFound, consts.MsgErrorTransactionNotFound)
		}
		err = actions.ConfirmTransaction(buyer, tsac, stub)
		if err != nil {
			return models.WithError(consts.CodeErrorOperationFail, err.Error())
		}
		return models.WithSuccess(consts.MsgOK, nil)
	case models.OPERATE_CANCELORDER:
		Usage := fmt.Sprintf("Usage : %s <Credentials> <TransactionId>", fcn)
		if len(args) < 2 {
			return models.WithError(consts.CodeErrorParams, Usage)
		}
		tsac, err := store.GetTransactionOrderById(args[1], stub)
		if err != nil {
			return models.WithError(consts.CodeErrorTransactionNotFound, consts.MsgErrorTransactionNotFound)
		}
		err = actions.CancelBuyTransaction(buyer, tsac, stub)
		if err != nil {
			return models.WithError(consts.CodeErrorOperationFail, err.Error())
		}
		return models.WithSuccess(consts.MsgOK, nil)
	default:
		return models.WithError(consts.CodeErrorPermissionDenied, consts.MsgErrorPermissionDenied)
	}
}
