package endpoints

import (
	"../data"
	"../response"

	"net/http"
)

func GetOrders(w http.ResponseWriter, r *http.Request) {

	s := data.Orders{
		Orders:            20,
		AverageOrderValue: 55.55,
		OrderMargin:       31.12,
		Refunds:           28.95,
		OrderList: []data.Order{{
			OrderID:     1337,
			Items:       1,
			Country:     "Canada",
			TotalAmount: 28.95,
			COGS:        15.55,
		}},
	}
	response.JSON(w, 200, s)
}
