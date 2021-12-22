package data

type Orders struct {
	Orders            int     `json:"orders"`
	AverageOrderValue float64 `json:"aov"`
	OrderMargin       float64 `json:"margin"`
	Refunds           float64 `json:"refunds"`

	OrderList []Order `json:"orderList"`
}

type Order struct {
	Items        int     `json:"items"`
	CustomerName string  `json:"customerName"`
	Country      string  `json:"country"`
	Amount       float64 `json:"amount"`
	COGS         float64 `json:"cogs"`
}
