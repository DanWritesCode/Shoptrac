package endpoints

import (
	"../data"
	"../response"
	"net/http"
)

func GetCustomers(w http.ResponseWriter, r *http.Request) {

	s := data.Customers{
		TotalCustomers:     50,
		NewCustomers:       45,
		ReturningCustomers: 5,
		HighestSpender:     88.65,
		LowestSpender:      22.48,
		TopCustomerList: []data.Customer{
			{
				Name:        "John Doe",
				Country:     "Canada",
				OrdersMade:  3,
				AmountSpent: 89.75,
			},
		},
	}
	response.JSON(w, 200, s)
}
