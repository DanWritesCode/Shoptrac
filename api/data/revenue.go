package data

type Revenue struct {
	Date            int64   `json:"date,omitempty"`
	Sales           float64 `json:"sales,omitempty"`
	ShippingCharged float64 `json:"shippingCharged,omitempty"`
	TaxesCollected  float64 `json:"taxesCollected,omitempty"`
	Tips            float64 `json:"tips,omitempty"`
	Discounts       float64 `json:"discounts,omitempty"`
}
