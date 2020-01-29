package models

const PrefixProduct = "prod"

const ShelvesstatusOffsell byte = 0
const ShelvesstatusOnsell byte = 1

type ProductUpdateRequest struct {
	*DataModel
	EachPrice            uint64 `json:"each_price"`
	Description          string `json:"description"`
	Inventory            uint32 `json:"inventory"`
	TransportAmount      uint64 `json:"transport_amount"`
	SpecifiedTemperature byte   `json:"specified_temperature"`
}

type ProductSnapshot struct {
	*ProductUpdateRequest
	SellerId string `json:"seller_id"`
}

type Product struct {
	*ProductSnapshot
	ShelvesStatus byte `json:"shelves_status"`
}

func (prod *Product) CloneSnapshot() *ProductSnapshot {
	return &ProductSnapshot{
		ProductUpdateRequest: &ProductUpdateRequest{
			DataModel:            &DataModel{Id: prod.Id},
			Description:          prod.Description,
			EachPrice:            prod.EachPrice,
			Inventory:            prod.Inventory,
			TransportAmount:      prod.TransportAmount,
			SpecifiedTemperature: prod.SpecifiedTemperature,
		},
		SellerId: prod.SellerId,
	}
}

func NewProduct(sellerId string,
	eachPrice uint64, initialStatus byte,
	description string, initialInventory uint32,
	transportAmount uint64, SpecifiedTemperature byte,
) *Product {
	return &Product{
		ProductSnapshot: &ProductSnapshot{
			ProductUpdateRequest: &ProductUpdateRequest{
				DataModel:            &DataModel{Id: PrefixProduct + AllocateIdS()},
				EachPrice:            eachPrice,
				Inventory:            initialInventory,
				Description:          description,
				TransportAmount:      transportAmount,
				SpecifiedTemperature: SpecifiedTemperature,
			},
			SellerId: sellerId,
		},
		ShelvesStatus: initialStatus,
	}
}
