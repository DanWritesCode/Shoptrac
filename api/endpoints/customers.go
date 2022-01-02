package endpoints

import (
	"../data"
	"../database"
	"../response"
	"math"
	"net/http"
)

func GetCustomers(w http.ResponseWriter, r *http.Request) {

	allCustomers, err := database.GetCustomers()
	if err != nil {
		// handle err
		response.BadRequest(w, "Unable to retrieve customers")
		return
	}

	s := data.Customers{}
	s.HighestSpender = 0
	s.LowestSpender = math.MaxFloat64
	i := 0
	s.TopCustomerList = make([]*data.Customer, 10)
	for _, customer := range allCustomers {
		if i < 10 {
			s.TopCustomerList[i] = customer
		}

		if customer.OrdersMade == 1 {
			s.NewCustomers = s.NewCustomers + 1
		} else if customer.OrdersMade > 1 {
			s.ReturningCustomers = s.ReturningCustomers + 1
		}

		s.HighestSpender = math.Max(s.HighestSpender, customer.AmountSpent)
		s.LowestSpender = math.Min(s.LowestSpender, customer.AmountSpent)

		i++
	}

	s.TotalCustomers = i

	response.JSON(w, 200, s)
}
