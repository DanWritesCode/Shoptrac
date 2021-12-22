package data

type Revenue struct {
	Sales           float64 `json:"sales,omitempty"`
	ShippingCharged float64 `json:"shippingCharged,omitempty"`
	TaxesCollected  float64 `json:"taxesCollected,omitempty"`
	Tips            float64 `json:"tips,omitempty"`
}
