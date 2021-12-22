package data

type Expenses struct {
	COGS           float64 `json:"cogs,omitempty"`
	Shipping       float64 `json:"shipping,omitempty"`
	TaxesForwarded float64 `json:"taxesForwarded,omitempty"`
	TotalMarketing float64 `json:"totalMarketing,omitempty"`

	Marketing   []Trio `json:"marketing"`
	Recurring   []Trio `json:"recurring"`
	Operational []Trio `json:"operational"`
}
