package data

type ExpensesSummary struct {
	PaymentProcessing float64 `json:"paymentProcessing,omitempty"`
	COGS              float64 `json:"cogs,omitempty"`
	Shipping          float64 `json:"shipping,omitempty"`
	TaxesForwarded    float64 `json:"taxesForwarded,omitempty"`
	TotalMarketing    float64 `json:"totalMarketing,omitempty"`

	Marketing []*Trio `json:"marketing"`
	Recurring []*Trio `json:"recurring"`
}

type Expense struct {
	ID       int     `json:"id,omitempty"`
	Category string  `json:"category"`
	Name     string  `json:"name"`
	Date     int64   `json:"date"`
	Amount   float64 `json:"amount"`
}

type GroupedExpenses struct {
  COGS      []*Trio `json:"cogs"`
  Marketing []*Trio `json:"marketing"`
  Recurring []*Trio `json:"recurring"`
}

func GroupExpenses(totalRev float64, expenses []*Expense) *GroupedExpenses {
  // i love golang so much
  mm :=  make(map[string]map[string]float64) // magic map
  mm["MARKETING"] = make(map[string]float64)
  mm["RECURRING"] = make(map[string]float64)
  mm["COGS"]      = make(map[string]float64)

  for _, e := range expenses {
    if mm[e.Category] != nil {
      mm[e.Category][e.Name] += e.Amount
    }
  }

  ge := GroupedExpenses{}
  for k, v := range mm {
    t := make([]*Trio, 0)
    for kk, vv := range v {
      t = append(t, &Trio{
        Name:       kk,
        Amount:     vv,
        Percentage: vv / totalRev * 100,
      })
    }
    if k == "MARKETING" {
      ge.Marketing = t
    } else if k == "RECURRING" {
      ge.Recurring = t
    } else if k == "COGS" {
      ge.COGS = t
    }
  }

  return &ge
}
