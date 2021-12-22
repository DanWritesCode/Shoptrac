package data

type Customers struct {
	TotalCustomers     int     `json:"totalCustomers"`
	NewCustomers       int     `json:"newCustomers"`
	ReturningCustomers int     `json:"returningCustomers"`
	HighestSpender     float64 `json:"highestSpender"`
	LowestSpender      float64 `json:"lowestSpender"`

	TopCustomerList []Customer `json:"topCustomerList"`
}

type Customer struct {
	Name        string  `json:"name"`
	Country     string  `json:"country"`
	OrdersMade  int     `json:"ordersMade"`
	AmountSpent float64 `json:"amountSpent"`
}
