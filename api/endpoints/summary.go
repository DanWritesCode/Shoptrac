package endpoints

import (
	"../data"
	"../response"
	"net/http"
)

func GetSummary(w http.ResponseWriter, r *http.Request) {
	s := data.Summary{
		Revenue:      51225.25,
		Profit:       1234.56,
		ProfitMargin: 55.5,
		Orders:       524,
		AOV:          112.43,
		COGS: []data.Trio{
			{
				Name:       "Items Sold",
				Amount:     4444.44,
				Percentage: 22.8,
			},
		},
		Marketing: []data.Trio{
			{
				Name:       "Facebook",
				Amount:     2222.22,
				Percentage: 11.4,
			},
		},
		RecurringCosts: []data.Trio{
			{
				Name:       "Shopify Plan",
				Amount:     29.991,
				Percentage: 0.08,
			},
		},
	}
	response.JSON(w, 200, s)
}
