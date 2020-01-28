package chaincode

import (
	"encoding/json"
	"fmt"
	"github.com/dmao/gome/chaincode/chaincode_foodmanage/actions"
	"github.com/dmao/gome/chaincode/chaincode_foodmanage/models"
	"github.com/dmao/gome/chaincode/chaincode_foodmanage/store"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

func (ctx *FoodManageChaincodeV1) processTransporterInvoke(tspr *models.Transporter, fcn string, args []string, stub shim.ChaincodeStubInterface) peer.Response {
	switch fcn {
	case OPERATE_COMPLETE_TRANSPORT:
		Usage := fmt.Sprintf("Usage : %s <Credentials> <TransportId>", fcn)
		if len(args) < 2 {
			return shim.Error(Usage)
		}
		torder, err := store.GetTransportOrderById(args[1], stub)
		if err != nil {
			return shim.Error(err.Error())
		}
		err = actions.CompleteTransport(tspr, torder, stub)
		if err != nil {
			return shim.Error(err.Error())
		}
		return shim.Success([]byte{})
	case OPERATE_CANCELTRANSPORT:
		Usage := fmt.Sprintf("Usage : %s <Credentials> <TransportId>", fcn)
		if len(args) < 2 {
			return shim.Error(Usage)
		}
		torder, err := store.GetTransportOrderById(args[1], stub)
		if err != nil {
			return shim.Error(err.Error())
		}
		err = actions.CancelTransport(tspr, torder, stub)
		if err != nil {
			return shim.Error(err.Error())
		}
		return shim.Success([]byte{})
	case OPERATE_UPDATE_TRANSPORT:
		Usage := fmt.Sprintf("Usage : %s <Credentials> <TransportId> <Json{temperature}>", fcn)
		if len(args) < 2 {
			return shim.Error(Usage)
		}
		torder, err := store.GetTransportOrderById(args[1], stub)
		if err != nil {
			return shim.Error(err.Error())
		}
		var details = &models.TransportDetails{}
		err = json.Unmarshal([]byte(args[2]), details)
		if err != nil {
			return shim.Error(err.Error() + Usage)
		}
		err = actions.UpdateDetails(tspr, torder, details, stub)
		if err != nil {
			return shim.Error(err.Error())
		}
		return shim.Success([]byte{})
	default:
		return shim.Error("permission denied")
	}
}
