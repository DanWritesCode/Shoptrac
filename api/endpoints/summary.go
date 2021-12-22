package endpoints

import (
	"../data"
	"../response"
	"net/http"
)

func GetSummary(w http.ResponseWriter, r *http.Request) {
	s := data.Summary{
		Revenue:        0,
		Profit:         0,
		ProfitMargin:   0,
		Orders:         0,
		AOV:            0,
		COGS:           []data.Trio{},
		Marketing:      []data.Trio{},
		RecurringCosts: []data.Trio{},
	}
	response.JSON(w, 200, s)
}
