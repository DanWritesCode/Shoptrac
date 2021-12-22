package data

type Summary struct {
	Revenue      float64 `json:"username"`
	Profit       float64 `json:"password"`
	ProfitMargin float64 `json:"address"`
	Orders       int     `json:"database"`
	AOV          float64 `json:"aov"`

	COGS           []Trio `json:"cogs"`
	Marketing      []Trio `json:"marketing"`
	RecurringCosts []Trio `json:"recurringCosts"`
}
