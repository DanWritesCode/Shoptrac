package endpoints

import (
  "net/http"
  "../data"
  "../response"
)

func GetSummary (w http.ResponseWriter, r *http.Request) {
  s := data.Summary{
    Revenue:        0,
    Profit:         0,
    ProfitMargin:   0,
    Orders:         0,
    AOV:            0,
    COGS:           []data.Expense{},
    Marketing:      []data.Expense{},
    RecurringCosts: []data.Expense{},
  }
  response.JSON(w, 200, s)
}
