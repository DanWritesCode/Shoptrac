package endpoints

import (
	"../data"
	"../response"
	"net/http"
)

func GetExpenses(w http.ResponseWriter, r *http.Request) {

	s := data.Expenses{
		COGS:           747.69,
		Shipping:       512.33,
		TaxesForwarded: 115.74,
		TotalMarketing: 1832.74,
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
		Operational: []data.Trio{
			{
				Name:       "Refunds",
				Amount:     88.11,
				Percentage: 5.8,
			},
			{
				Name:       "Payment Processing Fees (PayPal)",
				Amount:     11.48,
				Percentage: 1.4,
			},
			{
				Name:       "Payment Processing Fees (Shopify)",
				Amount:     52.82,
				Percentage: 5.2,
			},
		},
	}
	response.JSON(w, 200, s)
}
