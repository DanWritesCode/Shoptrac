package data

type Orders struct {
	Orders            int     `json:"orders"`
	AverageOrderValue float64 `json:"aov"`
	OrderMargin       float64 `json:"margin"`
	Refunds           float64 `json:"refunds"`

	OrderList []*Order `json:"orderList"`
}

type Order struct {
	OrderID        int    `json:"orderId"`
	Date           int64  `json:"date"`
	Items          int    `json:"items"`
	Country        string `json:"country"`
	PaymentGateway string `json:"paymentGateway"`

	Discount    float64 `json:"discount"`
	Subtotal    float64 `json:"subtotal"`
	Shipping    float64 `json:"shippingCharged"`
	Taxes       float64 `json:"taxesCharged"`
	Tips        float64 `json:"tipsCollected,omitempty"`
	TotalAmount float64 `json:"amount"`

	COGS float64 `json:"cogs"`
}
