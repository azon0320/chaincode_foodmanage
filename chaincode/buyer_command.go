package chaincode

import (
	"fmt"
	"github.com/dormao/chaincode_foodmanage/actions"
	"github.com/dormao/chaincode_foodmanage/models"
	"github.com/dormao/chaincode_foodmanage/store"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
	"strconv"
)

func (ctx *FoodManageChaincode) processBuyerInvoke(buyer *models.Buyer, fcn string, args []string, stub shim.ChaincodeStubInterface) peer.Response {
	switch fcn {
	case OPERATE_PURCHASE:
		Usage := fmt.Sprintf("Usage : %s <Credentials> <ProductId> <PurchaseCount>", fcn)
		if len(args) < 3 {
			return shim.Error(Usage)
		}
		prod, err := store.GetProductById(args[1], stub)
		if err != nil {
			return shim.Error(err.Error())
		}
		count, err := strconv.Atoi(args[2])
		if err != nil {
			return shim.Error(Usage)
		}
		tsacId, err := actions.PurchaseProduct(buyer, prod, uint32(count), stub)
		if err != nil {
			return shim.Error(err.Error())
		}
		return shim.Success([]byte(tsacId))
	case OPERATE_CONFIRM:
		Usage := fmt.Sprintf("Usage : %s <Credentials> <TransactionId>", fcn)
		if len(args) < 2 {
			return shim.Error(Usage)
		}
		tsac, err := store.GetTransactionOrderById(args[1], stub)
		if err != nil {
			return shim.Error(err.Error())
		}
		err = actions.ConfirmTransaction(buyer, tsac, stub)
		if err != nil {
			return shim.Error(err.Error())
		}
		return shim.Success([]byte{})
	case OPERATE_CANCELORDER:
		Usage := fmt.Sprintf("Usage : %s <Credentials> <TransactionId>", fcn)
		if len(args) < 2 {
			return shim.Error(Usage)
		}
		tsac, err := store.GetTransactionOrderById(args[1], stub)
		if err != nil {
			return shim.Error(err.Error())
		}
		err = actions.CancelBuyTransaction(buyer, tsac, stub)
		if err != nil {
			return shim.Error(err.Error())
		}
		return shim.Success([]byte{})
	default:
		return shim.Error("permission denied")
	}
}
