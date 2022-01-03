package endpoints

import (
	"../data"
	"../database"
	"../response"

	"net/http"
)

func GetOrders(w http.ResponseWriter, r *http.Request) {

	s := data.Orders{}
	s.OrderList = make([]*data.Order, 0)

	ord, err := database.GetOrdersAfterDate(0)
	if err != nil {
		response.BadRequest(w, "Unable to get orders from the database")
		return
	}

	revenue := 0.0
	expenses := 0.0
	for _, o := range ord {
		s.OrderList = append(s.OrderList, o)
		s.Orders += 1

		revenue += o.TotalAmount
		expenses += o.Taxes // taxes get forwarded and thus get counted as expenses
	}

	s.AverageOrderValue = revenue / float64(s.Orders)
	s.OrderMargin = (revenue - expenses) / revenue

	// TODO refunds

	response.JSON(w, 200, s)
}
