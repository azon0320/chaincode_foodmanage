package models

const PREFIX_TRANSPORT_ORDER = "tspo"

type TransportDetails struct {
	Temperature byte `json:"temperature"`
}

type TransportOrder struct {
	*DataModel
	Snapshot      *ProductSnapshot  `json:"snapshot"`
	OrderWasted   bool              `json:"order_wasted"`
	TransporterId string            `json:"transporter_id"`
	TransactionId string            `json:"transaction_id"`
	Details       *TransportDetails `json:"details"`
}

func NewTransportOrder(
	transporterId string,
	transactionId string,
	snapshot *ProductSnapshot,
	initialDetails *TransportDetails,
) *TransportOrder {
	return &TransportOrder{
		OrderWasted:   false,
		DataModel:     &DataModel{Id: PREFIX_TRANSPORT_ORDER + AllocateIdS()},
		Snapshot:      snapshot,
		TransporterId: transporterId,
		TransactionId: transactionId,
		Details:       initialDetails,
	}
}
