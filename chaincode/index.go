package chaincode

import (
	"errors"
	"fmt"
	"github.com/dormao/chaincode_foodmanage/actions"
	"github.com/dormao/chaincode_foodmanage/auth"
	"github.com/dormao/chaincode_foodmanage/models"
	"github.com/dormao/chaincode_foodmanage/store"
	"github.com/dormao/chaincode_foodmanage/util"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
	"strconv"
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

type FoodManageChaincode struct{}

func (ctx *FoodManageChaincode) Init(stub shim.ChaincodeStubInterface) peer.Response {
	var log string = ""
	log += "got init func,"
	_, args := stub.GetFunctionAndParameters()
	initBalance, err := store.GetStateInt(store.GlobalsInitialBalanceIndex, DefaultInitialBalance, stub)
	if err != nil {
		if len(args) > 0 {
			val, err := strconv.Atoi(args[0])
			if err == nil {
				initBalance = val
				log += fmt.Sprintf("argument balance found : %d,", initBalance)
			}
		}
		log += fmt.Sprintf("global initial balance not found,creating as %d,", DefaultInitialBalance)
		err = store.SetStateInt(store.GlobalsInitialBalanceIndex, initBalance, stub)
		if err != nil {
			log += "create balance state failed,"
		}else {
			log += fmt.Sprintf("created balance state : %d,", DefaultInitialBalance)
		}
	}
	log += fmt.Sprintf("default initial balance is %d", initBalance)
	fmt.Println(strings.Replace(log, ",", "\n", strings.IndexAny(log, ",")))
	return shim.Success([]byte(log))
}

func (ctx *FoodManageChaincode) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	fcn, args := stub.GetFunctionAndParameters()

	// TODO DEBUGGER DELETE
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

func (ctx *FoodManageChaincode) processUnAuthenticatedInvoke(fcn string, args []string, stub shim.ChaincodeStubInterface) (peer.Response, bool) {
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

func (ctx *FoodManageChaincode) processAuthenticatedInvoke(operator *models.Operator, fcn string, args []string, stub shim.ChaincodeStubInterface) peer.Response {
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
