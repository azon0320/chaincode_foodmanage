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

func (ctx *FoodManageChaincode) processTransporterInvoke(tspr *models.Transporter, fcn string, args []string, stub shim.ChaincodeStubInterface) *models.DataReturns {
	switch fcn {
	case models.OPERATE_COMPLETE_TRANSPORT:
		Usage := fmt.Sprintf("Usage : %s <Credentials> <TransportId>", fcn)
		if len(args) < 2 {
			return models.WithError(consts.CodeErrorParams, Usage)
		}
		torder, err := store.GetTransportOrderById(args[1], stub)
		if err != nil {
			return models.WithError(consts.CodeErrorTransportNotFound, consts.MsgErrorTransportNotFound)
		}
		err = actions.CompleteTransport(tspr, torder, stub)
		if err != nil {
			return models.WithError(consts.CodeErrorOperationFail, err.Error())
		}
		return models.WithSuccess(consts.MsgOK, nil)
	case models.OPERATE_CANCELTRANSPORT:
		Usage := fmt.Sprintf("Usage : %s <Credentials> <TransportId>", fcn)
		if len(args) < 2 {
			return models.WithError(consts.CodeErrorParams, Usage)
		}
		torder, err := store.GetTransportOrderById(args[1], stub)
		if err != nil {
			return models.WithError(consts.CodeErrorTransportNotFound, consts.MsgErrorTransportNotFound)
		}
		err = actions.CancelTransport(tspr, torder, stub)
		if err != nil {
			return models.WithError(consts.CodeErrorOperationFail, err.Error())
		}
		return models.WithSuccess(consts.MsgOK, nil)
	case models.OPERATE_UPDATE_TRANSPORT:
		Usage := fmt.Sprintf("Usage : %s <Credentials> <TransportId> <Json{temperature}>", fcn)
		if len(args) < 2 {
			return models.WithError(consts.CodeErrorParams, Usage)
		}
		torder, err := store.GetTransportOrderById(args[1], stub)
		if err != nil {
			return models.WithError(consts.CodeErrorTransportNotFound, consts.MsgErrorTransportNotFound)
		}
		var details = &models.TransportDetails{}
		err = json.Unmarshal([]byte(args[2]), details)
		if err != nil {
			return models.WithError(consts.CodeErrorParams, Usage)
		}
		err = actions.UpdateDetails(tspr, torder, details, stub)
		if err != nil {
			return models.WithError(consts.CodeErrorOperationFail, err.Error())
		}
		return models.WithSuccess(consts.MsgOK, nil)
	default:
		return models.WithError(consts.CodeErrorPermissionDenied, consts.MsgErrorPermissionDenied)
	}
}
