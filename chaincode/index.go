package chaincode

import (
	"errors"
	"fmt"
	"github.com/dormao/chaincode_foodmanage/actions"
	"github.com/dormao/chaincode_foodmanage/auth"
	"github.com/dormao/chaincode_foodmanage/models"
	"github.com/dormao/chaincode_foodmanage/models/consts"
	"github.com/dormao/chaincode_foodmanage/store"
	"github.com/dormao/chaincode_foodmanage/util"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
	"strconv"
	"strings"
)

/*
var (
	OperationMap = map[byte][]string{
		models.OperatorSeller: []string{
			OPERATE_ADDPRODUCT,
			OPERATE_UPDATE_PRODUCT,
			OPERATE_TAKEONSELL,
			OPERATE_TAKEOFFSELL,
			OPERATE_CANCELORDER,
		},
		models.OperatorBuyer: []string{
			OPERATE_PURCHASE,
			OPERATE_CANCELORDER,
			OPERATE_CONFIRM,
		},
		models.OperatorTransporter: []string{
			OPERATE_CANCELTRANSPORT,
			OPERATE_UPDATE_TRANSPORT,
			OPERATE_COMPLETE_TRANSPORT,
		},
	}
)
*/

type FoodManageChaincode struct{}

func (ctx *FoodManageChaincode) Init(stub shim.ChaincodeStubInterface) peer.Response {
	var log string = ""
	log += "got init func,"
	_, args := stub.GetFunctionAndParameters()
	initBalance, err := store.GetStateInt(store.GlobalsInitialBalanceIndex, models.DefaultInitialBalance, stub)
	if err != nil {
		if len(args) > 0 {
			val, err := strconv.Atoi(args[0])
			if err == nil {
				initBalance = val
				log += fmt.Sprintf("argument balance found : %d,", initBalance)
			}
		}
		log += fmt.Sprintf("global initial balance not found,creating as %d,", models.DefaultInitialBalance)
		err = store.SetStateInt(store.GlobalsInitialBalanceIndex, initBalance, stub)
		if err != nil {
			log += "create balance state failed,"
		}else {
			log += fmt.Sprintf("created balance state : %d,", models.DefaultInitialBalance)
		}
	}
	log += fmt.Sprintf("default initial balance is %d", initBalance)
	fmt.Println(strings.Replace(log, ",", "\n", strings.IndexAny(log, ",")))
	return shim.Success([]byte(log))
}

func (ctx *FoodManageChaincode) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	time, err := stub.GetTxTimestamp()
	if err != nil {
		return shim.Error(err.Error())
	}
	models.UpdateTxNanos(util.GetTxTimeNanos(time.GetSeconds(), time.GetNanos()))

	fcn, args := stub.GetFunctionAndParameters()

	// TODO DEBUGGER DELETE
	fmt.Println(fmt.Sprintf("Begin Process (txNano : %d)", models.GetTxTimeNanos()))
	fmt.Println(fmt.Sprintf("Fcn = %s", fcn))
	for k, v := range args{
		fmt.Println(fmt.Sprintf("Arg[%d] = %s", k, v))
	}
	fmt.Println()

	response, ok := ctx.processUnAuthenticatedInvoke(fcn, args, stub)
	if ok {
		 return response
	}
	if len(args) < 1 {
		return shim.Success(util.JsonEncode(models.WithError(consts.CodeErrorUnAuth, consts.MsgErrorUnAuth)))
	}
	credentials, err := util.GetCredentialsFromString(args[0])
	if err != nil {
		return shim.Success(util.JsonEncode(models.WithError(consts.CodeErrorAuthFail, consts.MsgErrorAuthFail)))
	}
	operator, err2 := auth.AttemptWithCredentials(credentials, stub)
	if err2 != nil {
		return shim.Success(util.JsonEncode(models.WithError(consts.CodeErrorAuthFail, err2.Error())))
	}
	return ctx.processAuthenticatedInvoke(operator, fcn, args, stub)
}

func (ctx *FoodManageChaincode) processUnAuthenticatedInvoke(fcn string, args []string, stub shim.ChaincodeStubInterface) (peer.Response, bool) {
	switch fcn {
	case models.UnAuthRegisterSeller, models.UnAuthRegisterBuyer, models.UnAuthRegisterTransporter:
		Usage := fmt.Sprintf("Usage : %s <Password>", fcn)
		if len(args) < 1 {
			return shim.Success(util.JsonEncode(models.WithError(consts.CodeErrorParams, Usage))), true
		}
		password := args[0]
		if strings.TrimSpace(password) == "" {
			return shim.Success(util.JsonEncode(models.WithError(consts.CodeErrorParams, Usage))), true
		}
		initBal,_ := store.GetStateInt(store.GlobalsInitialBalanceIndex, models.DefaultInitialBalance, stub)
		id, err := "", errors.New("unexpected register function")
		switch fcn {
		case models.UnAuthRegisterSeller:
			id, err = actions.RegisterSeller(password, uint64(initBal), stub)
		case models.UnAuthRegisterBuyer:
			id, err = actions.RegisterBuyer(password, uint64(initBal), stub)
		case models.UnAuthRegisterTransporter:
			id, err = actions.RegisterTransporter(password, uint64(initBal), stub)
		}
		if err != nil {
			return shim.Success(util.JsonEncode(models.WithError(consts.CodeErrorParams, err.Error()))), true
		}
		return shim.Success(util.JsonEncode(models.WithSuccess(consts.MsgOK, id))), true
	case models.UnAuthLogin:
		Usage := fmt.Sprintf("Usage : %s <AccountId> <Password>", fcn)
		if len(args) < 2 {
			return shim.Success(util.JsonEncode(models.WithError(consts.CodeErrorParams, Usage))), true
		}
		operator, err := auth.AttemptWithPassword(
			&models.Credentials{AccountId:args[0], Password:args[1], Token:""}, stub)
		if err != nil {
			return shim.Success(util.JsonEncode(models.WithError(consts.CodeErrorAuthFail, err.Error()))), true
		}
		return shim.Success(util.JsonEncode(models.WithSuccess(consts.MsgOK, operator.Token))), true
	case models.UnAuthPing:
		return shim.Success(util.JsonEncode(models.WithSuccess(consts.MsgOK, "pong"))), true
	}
	return shim.Success(util.JsonEncode(models.WithError(consts.CodeErrorAuthFail, "unauth function invalid"))), false
}

func (ctx *FoodManageChaincode) processAuthenticatedInvoke(operator *models.Operator, fcn string, args []string, stub shim.ChaincodeStubInterface) peer.Response {
	var Returns *models.DataReturns = nil
	switch operator.OperatorType {
	case models.OperatorSeller:
		Returns = ctx.processSellerInvoke(&models.Seller{Operator: operator}, fcn, args, stub)
	case models.OperatorBuyer:
		Returns = ctx.processBuyerInvoke(&models.Buyer{Operator: operator}, fcn, args, stub)
	case models.OperatorTransporter:
		Returns = ctx.processTransporterInvoke(&models.Transporter{Operator: operator}, fcn, args, stub)
	}
	if Returns == nil {
		return shim.Success(util.JsonEncode(models.WithError(consts.CodeErrorAuthFail, "invalid identity")))
	}else {
		return shim.Success(util.JsonEncode(Returns))
	}
}
