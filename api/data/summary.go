package data

type Summary struct {
	Revenue      float64 `json:"revenue"`
	Expenses     float64 `json:"expenses"`
	Profit       float64 `json:"profit"`
	ProfitMargin float64 `json:"profitMargin"`
	Orders       int     `json:"orders"`
	AOV          float64 `json:"aov"`

	COGS           []Trio `json:"cogs"`
	Marketing      []Trio `json:"marketing"`
	RecurringCosts []Trio `json:"recurringCosts"`
}
