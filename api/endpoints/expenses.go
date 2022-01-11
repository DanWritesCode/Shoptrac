package endpoints

import (
	"../data"
	"../response"
	"net/http"
)

func GetExpenses(w http.ResponseWriter, r *http.Request) {

	s := data.ExpensesSummary{
		PaymentProcessing: 515.55,
		COGS:              747.69,
		Shipping:          512.33,
		TaxesForwarded:    115.74,
		Marketing: []data.Trio{
			{
				Name:       "Facebook",
				Amount:     1499.22,
				Percentage: 75,
			},
			{
				Name:       "Google Ads",
				Amount:     333.52,
				Percentage: 25,
			},
		},
		Recurring: []data.Trio{
			{
				Name:       "Shopify Plan",
				Amount:     29.99,
				Percentage: 8.3,
			},
		},
	}
	response.JSON(w, 200, s)
}
