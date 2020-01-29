package models

const PREFIX_TRANSACTION_ORDER = "tsac"

const TRANSACTION_ORDER_STATUS_UNTRANSMIT byte = 0
const TRANSACTION_ORDER_STATUS_TRANSPORTING byte = 1
const TRANSACTION_ORDER_STATUS_UNCONFIRM byte = 2
const TRANSACTION_ORDER_STATUS_FAILED byte = 3
const TRANSACTION_ORDER_STATUS_COMPLETED byte = 4

type TransactionOrder struct {
	*DataModel
	Snapshot         *ProductSnapshot `json:"snapshot"`
	TransportOrderId string           `json:"transport_order_id"`
	BuyerId          string           `json:"buyer_id"`
	ProductCount     uint32           `json:"product_count"`
	OrderStatus      byte             `json:"order_status"`
}

func NewTransactionOrder(
	snapshot *ProductSnapshot, buyerId string, productCount uint32,
) *TransactionOrder {
	order := &TransactionOrder{
		DataModel: &DataModel{Id: PREFIX_TRANSACTION_ORDER + AllocateIdS()},
		Snapshot:  snapshot, TransportOrderId: "",
		BuyerId:      buyerId,
		ProductCount: productCount,
	}
	order.OrderStatus = TRANSACTION_ORDER_STATUS_UNTRANSMIT
	order.TransportOrderId = ""
	return order
}
