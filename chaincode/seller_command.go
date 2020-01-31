package chaincode

import (
	"encoding/json"
	"fmt"
	"github.com/dormao/chaincode_foodmanage/actions"
	"github.com/dormao/chaincode_foodmanage/models"
	"github.com/dormao/chaincode_foodmanage/models/consts"
	"github.com/dormao/chaincode_foodmanage/store"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

func (ctx *FoodManageChaincode) processSellerInvoke(seller *models.Seller, fcn string, args []string, stub shim.ChaincodeStubInterface) *models.DataReturns {
	switch fcn {
	case models.OPERATE_ADDPRODUCT: // add_prod
		prodId, err := actions.AddProduct(seller, stub)
		if err != nil {
			return models.WithError(consts.CodeErrorProdNotFound, consts.MsgErrorProdNotFound)
		}
		return models.WithSuccess(consts.MsgOK, prodId)
	case models.OPERATE_UPDATE_PRODUCT: // add_prod <jsonTypeMap>
		Usage := fmt.Sprintf("Usage: %s <Credentials> <ProductId> <Json{each_price,inventory,temperature,description,transport_amount}>", fcn)
		if len(args) != 3 {
			return models.WithError(consts.CodeErrorParams, Usage)
		}
		prod, err := store.GetProductById(args[1], stub)
		if err != nil {
			return models.WithError(consts.CodeErrorProdNotFound, consts.MsgErrorProdNotFound)
		}
		var request = &models.ProductUpdateRequest{}
		err = json.Unmarshal([]byte(args[2]), &request)
		if err != nil {
			return models.WithError(consts.CodeErrorParams, Usage)
		}
		err = actions.UpdateProductInfo(seller, prod, request, stub)
		if err != nil {
			return models.WithError(consts.CodeErrorOperatorNotFound, err.Error())
		}
		return models.WithSuccess(consts.MsgOK, nil)
	case models.OPERATE_TAKEONSELL:
		Usage := fmt.Sprintf("Usage: %s <Credentials> <ProductId>", fcn)
		if len(args) != 2 {
			return models.WithError(consts.CodeErrorParams, Usage)
		}
		prod, err := store.GetProductById(args[1], stub)
		if err != nil {
			return models.WithError(consts.CodeErrorProdNotFound, consts.MsgErrorProdNotFound)
		}
		err = actions.TakeOnSellProduct(seller, prod, stub)
		if err != nil {
			return models.WithError(consts.CodeErrorOperationFail, err.Error())
		}
		return models.WithSuccess(consts.MsgOK, nil)
	case models.OPERATE_TAKEOFFSELL:
		Usage := fmt.Sprintf("Usage: %s <Credentials> <ProductId>", fcn)
		if len(args) != 2 {
			return models.WithError(consts.CodeErrorParams, Usage)
		}
		prod, err := store.GetProductById(args[1], stub)
		if err != nil {
			return models.WithError(consts.CodeErrorProdNotFound, consts.MsgErrorProdNotFound)
		}
		err = actions.TakeOffSellProduct(seller, prod, stub)
		if err != nil {
			return models.WithError(consts.CodeErrorOperationFail, err.Error())
		}
		return models.WithSuccess(consts.MsgOK, nil)
	case models.OPERATE_TRANSMIT:
		Usage := fmt.Sprintf("Usage: %s <Credentials> <TransactionId> <TransporterId>", fcn)
		if len(args) != 3 {
			return models.WithError(consts.CodeErrorParams, Usage)
		}
		tsac, err := store.GetTransactionOrderById(args[1], stub)
		if err != nil {
			return models.WithError(consts.CodeErrorTransactionNotFound, consts.MsgErrorTransactionNotFound)
		}
		tspr, err := store.GetTransporterById(args[2], stub)
		if err != nil {
			return models.WithError(consts.CodeErrorTransporterNotFound, consts.MsgErrorTransporterNotFound)
		}
		torderId, err := actions.TransmitOrder(seller, tsac, tspr, stub)
		if err != nil {
			return models.WithError(consts.CodeErrorOperationFail, err.Error())
		}
		return models.WithSuccess(consts.MsgOK, torderId)
	case models.OPERATE_CANCELORDER:
		Usage := fmt.Sprintf("Usage: %s <Credentials> <TransactionId>", fcn)
		if len(args) != 2 {
			return models.WithError(consts.CodeErrorParams, Usage)
		}
		tsac, err := store.GetTransactionOrderById(args[1], stub)
		if err != nil {
			return models.WithError(consts.CodeErrorTransactionNotFound, consts.MsgErrorTransactionNotFound)
		}
		err = actions.CancelSellTransaction(seller, tsac, stub)
		if err != nil {
			return models.WithError(consts.CodeErrorOperationFail, err.Error())
		}
		return models.WithSuccess(consts.MsgOK, nil)
	default:
		return models.WithError(consts.CodeErrorPermissionDenied, consts.MsgErrorPermissionDenied)
	}
}
