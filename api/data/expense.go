package data

type Expense struct {
  ExpenseName string `json:"name,omitempty"`
  Amount float64 `json:"amount,omitempty"`
  Percentage  float64 `json:"percentage,omitempty"`
}
