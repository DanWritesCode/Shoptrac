package data

type Revenue struct {
	Date            int64   `json:"date,omitempty"`
	Sales           float64 `json:"sales"`
	ShippingCharged float64 `json:"shippingCharged"`
	TaxesCollected  float64 `json:"taxesCollected"`
	Tips            float64 `json:"tips"`
	Discounts       float64 `json:"discounts"`
	Total           float64 `json:"total"`
}
