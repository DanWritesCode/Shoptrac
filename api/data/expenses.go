package data

type ExpensesSummary struct {
	PaymentProcessing float64 `json:"paymentProcessing,omitempty"`
	COGS              float64 `json:"cogs,omitempty"`
	Shipping          float64 `json:"shipping,omitempty"`
	TaxesForwarded    float64 `json:"taxesForwarded,omitempty"`
	TotalMarketing    float64 `json:"totalMarketing,omitempty"`

	Marketing []Trio `json:"marketing"`
	Recurring []Trio `json:"recurring"`
}

type Expense struct {
	ID       int     `json:"id,omitempty"`
	Category string  `json:"category"`
	Name     string  `json:"name"`
	Date     int64   `json:"date"`
	Amount   float64 `json:"amount"`
}
