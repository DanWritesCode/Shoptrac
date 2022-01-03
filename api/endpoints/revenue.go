package endpoints

import (
	"../data"
	"../database"
	"../response"
	"net/http"
)

func GetRevenue(w http.ResponseWriter, r *http.Request) {
	s := data.Revenue{}

	revs, err := database.GetDailyRevenue(0)
	if err != nil {
		response.BadRequest(w, "Unable to get revenue numbers from the database")
		return
	}

	for _, rev := range revs {
		s.Sales += rev.Sales
		s.ShippingCharged += rev.ShippingCharged
		s.TaxesCollected += rev.TaxesCollected
		s.Tips += rev.Tips
		s.Discounts += rev.Discounts
		s.Total += rev.Total
	}

	response.JSON(w, 200, s)
}
