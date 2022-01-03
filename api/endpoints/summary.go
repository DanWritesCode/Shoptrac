package endpoints

import (
	"../data"
	"../database"
	"../response"
	"net/http"
)

func GetSummary(w http.ResponseWriter, r *http.Request) {
	s := data.Summary{}

	totalRev := 0.0
	totalExpenses := 0.0
	rev, err := database.GetDailyRevenue(0)
	if err != nil {
		response.BadRequest(w, "Unable to get revenue numbers from the database")
		return
	}

	for _, drev := range rev {
		totalRev += drev.Sales
	}
	s.Revenue = totalRev

	totalOrderRev := 0.0
	ord, err := database.GetOrdersAfterDate(0)
	if err != nil {
		response.BadRequest(w, "Unable to get orders from the database")
		return
	}

	for _, o := range ord {
		totalOrderRev += o.TotalAmount
		totalExpenses += o.Taxes // taxes collected will be automatically considered expenses as they are not kept by the store (hopefully)
	}

	s.Profit = s.Revenue - totalExpenses
	s.ProfitMargin = s.Profit / s.Revenue
	s.Orders = len(ord)
	s.AOV = s.Revenue / float64(len(ord))

	s.COGS = make([]*data.Trio, 0)
	s.Marketing = make([]*data.Trio, 0)
	s.RecurringCosts = make([]*data.Trio, 0)

	response.JSON(w, 200, s)
}
