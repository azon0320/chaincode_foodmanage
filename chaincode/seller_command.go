package chaincode

import (
	"encoding/json"
	"fmt"
	"github.com/dormao/chaincode_foodmanage/actions"
	"github.com/dormao/chaincode_foodmanage/models"
	"github.com/dormao/chaincode_foodmanage/store"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

func (ctx *FoodManageChaincode) processSellerInvoke(seller *models.Seller, fcn string, args []string, stub shim.ChaincodeStubInterface) peer.Response {
	switch fcn {
	case OPERATE_ADDPRODUCT: // add_prod
		prodId, err := actions.AddProduct(seller, stub)
		if err != nil {
			return shim.Error(err.Error())
		}
		return shim.Success([]byte(prodId))
	case OPERATE_UPDATE_PRODUCT: // add_prod <jsonTypeMap>
		Usage := fmt.Sprintf("Usage: %s <Credentials> <ProductId> <Json{each_price,inventory,temperature,description,transport_amount}>", fcn)
		if len(args) != 3 {
			return shim.Error(Usage)
		}
		prod, err := store.GetProductById(args[1], stub)
		if err != nil {
			return shim.Error(err.Error())
		}
		var request *models.ProductUpdateRequest = &models.ProductUpdateRequest{}
		err = json.Unmarshal([]byte(args[2]), &request)
		if err != nil {
			return shim.Error(Usage)
		}
		err = actions.UpdateProductInfo(seller, prod, request, stub)
		if err != nil {
			return shim.Error(err.Error())
		}
		return shim.Success([]byte{})
	case OPERATE_TAKEONSELL:
		Usage := fmt.Sprintf("Usage: %s <Credentials> <ProductId>", fcn)
		if len(args) != 2 {
			return shim.Error(Usage)
		}
		prod, err := store.GetProductById(args[1], stub)
		if err != nil {
			return shim.Error(err.Error())
		}
		err = actions.TakeOnSellProduct(seller, prod, stub)
		if err != nil {
			return shim.Error(err.Error())
		}
		return shim.Success([]byte{})
	case OPERATE_TAKEOFFSELL:
		Usage := fmt.Sprintf("Usage: %s <Credentials> <ProductId>", fcn)
		if len(args) != 2 {
			return shim.Error(Usage)
		}
		prod, err := store.GetProductById(args[1], stub)
		if err != nil {
			return shim.Error(err.Error())
		}
		err = actions.TakeOffSellProduct(seller, prod, stub)
		if err != nil {
			return shim.Error(err.Error())
		}
		return shim.Success([]byte{})
	case OPERATE_TRANSMIT:
		Usage := fmt.Sprintf("Usage: %s <Credentials> <TransactionId> <TransporterId>", fcn)
		if len(args) != 3 {
			return shim.Error(Usage)
		}
		tsac, err := store.GetTransactionOrderById(args[1], stub)
		if err != nil {
			return shim.Error("transaction not found")
		}
		tspr, err := store.GetTransporterById(args[2], stub)
		if err != nil {
			return shim.Error("transporter not found")
		}
		torderId, err := actions.TransmitOrder(seller, tsac, tspr, stub)
		if err != nil {
			return shim.Error(err.Error())
		}
		return shim.Success([]byte(torderId))
	case OPERATE_CANCELORDER:
		Usage := fmt.Sprintf("Usage: %s <Credentials> <TransactionId>", fcn)
		if len(args) != 2 {
			return shim.Error(Usage)
		}
		tsac, err := store.GetTransactionOrderById(args[1], stub)
		if err != nil {
			return shim.Error(err.Error())
		}
		err = actions.CancelSellTransaction(seller, tsac, stub)
		if err != nil {
			return shim.Error(err.Error())
		}
		return shim.Success([]byte{})
	default:
		return shim.Error("permission denied")
	}
}
