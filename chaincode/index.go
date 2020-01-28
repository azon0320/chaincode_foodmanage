package chaincode

import (
	"errors"
	"fmt"
	"github.com/dmao/gome/chaincode/chaincode_foodmanage/actions"
	"github.com/dmao/gome/chaincode/chaincode_foodmanage/auth"
	"github.com/dmao/gome/chaincode/chaincode_foodmanage/models"
	"github.com/dmao/gome/chaincode/chaincode_foodmanage/store"
	"github.com/dmao/gome/chaincode/chaincode_foodmanage/util"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
	"strings"
)

// -----  BEFORE AUTHENTICATE  -----
const UnAuthRegisterSeller = "reg_seller"
const UnAuthRegisterBuyer = "reg_buyer"
const UnAuthRegisterTransporter = "reg_transporter"

// -----  BEGIN SELLER  -----
const OPERATE_ADDPRODUCT = "add_prod"
const OPERATE_UPDATE_PRODUCT = "update_prod"
const OPERATE_TAKEONSELL = "sell_on"
const OPERATE_TAKEOFFSELL = "sell_off"
const OPERATE_TRANSMIT = "transmit"
const OPERATE_CANCELORDER = "cancel_order"

// -----  BEGIN BUYER  -----
const OPERATE_PURCHASE = "buy_prod"
const OPERATE_CONFIRM = "confirm"

//const OPERATE_CANCELORDER = "cancel_order"

// -----  BEGIN TRANSPORTER  -----
const OPERATE_CANCELTRANSPORT = "cancel_transport"
const OPERATE_UPDATE_TRANSPORT = "update_transport"
const OPERATE_COMPLETE_TRANSPORT = "complete_transport"

const DefaultInitialBalance = 1000

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

type FoodManageChaincodeV1 struct{}

func (ctx *FoodManageChaincodeV1) Init(stub shim.ChaincodeStubInterface) peer.Response {
	var log string = ""
	log += "got init func, "
	initBalance, err := store.GetStateInt(store.GlobalsInitialBalanceIndex, DefaultInitialBalance, stub)
	if err != nil {
		log += fmt.Sprintf("global initial balance not found, creating as %d, ", DefaultInitialBalance)
		err = store.SetStateInt(store.GlobalsInitialBalanceIndex, DefaultInitialBalance, stub)
		if err != nil {
			log += "create balance state failed, "
		}
	}
	log += fmt.Sprintf("default initial balance is %d", initBalance)
	return shim.Success([]byte(log))
}

func (ctx *FoodManageChaincodeV1) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	fcn, args := stub.GetFunctionAndParameters()
	response, ok := ctx.processUnAuthenticatedInvoke(fcn, args, stub)
	if ok {
		 return response
	}
	if len(args) < 1 {
		return shim.Error("authenticate failed")
	}
	credentials, err := util.GetCredentialsFromString(args[0])
	if err != nil {
		return shim.Error(err.Error())
	}
	operator, err2 := auth.AttemptWithPassword(credentials, stub)
	if err2 != nil {
		return shim.Error(err2.Error())
	}
	return ctx.processAuthenticatedInvoke(operator, fcn, args, stub)
}

func (ctx *FoodManageChaincodeV1) processUnAuthenticatedInvoke(fcn string, args []string, stub shim.ChaincodeStubInterface) (peer.Response, bool) {
	switch fcn {
	case UnAuthRegisterSeller, UnAuthRegisterBuyer, UnAuthRegisterTransporter:
		Usage := fmt.Sprintf("Usage : %s <Password>", fcn)
		if len(args) < 1 {
			return shim.Error(Usage), true
		}
		password := args[0]
		if strings.TrimSpace(password) == "" {
			return shim.Error(Usage), true
		}
		initBal,_ := store.GetStateInt(store.GlobalsInitialBalanceIndex, DefaultInitialBalance, stub)
		id, err := "", errors.New("unexpected register function")
		switch fcn {
		case UnAuthRegisterSeller:
			id, err = actions.RegisterSeller(password, uint64(initBal), stub)
		case UnAuthRegisterBuyer:
			id, err = actions.RegisterBuyer(password, uint64(initBal), stub)
		case UnAuthRegisterTransporter:
			id, err = actions.RegisterTransporter(password, uint64(initBal), stub)
		}
		if err != nil {
			return shim.Error(err.Error()), true
		}
		return shim.Success([]byte(id)), true
	}
	return shim.Error("unauth function invalid"), false
}

func (ctx *FoodManageChaincodeV1) processAuthenticatedInvoke(operator *models.Operator, fcn string, args []string, stub shim.ChaincodeStubInterface) peer.Response {
	switch operator.OperatorType {
	case models.OperatorSeller:
		return ctx.processSellerInvoke(&models.Seller{Operator: operator}, fcn, args, stub)
	case models.OperatorBuyer:
		return ctx.processBuyerInvoke(&models.Buyer{Operator: operator}, fcn, args, stub)
	case models.OperatorTransporter:
		return ctx.processTransporterInvoke(&models.Transporter{Operator: operator}, fcn, args, stub)
	default:
		return shim.Error("invalid identity")
	}
}
