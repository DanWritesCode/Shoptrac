package endpoints

import (
	"../data"
	"../database"
	"../response"
	"net/http"
)

func GetItems(w http.ResponseWriter, r *http.Request) {
	s := data.Products{}

	tsp, err := database.GetTopSellingProducts(0, 10)
	if err != nil {
		response.BadRequest(w, "Unable to get products from the database")
		return
	}

	// average price of *top selling* products - later on we can change this to average price of all products
	allItemsPrice := 0.0
	for _, item := range tsp {
		allItemsPrice += item.Product.Price
	}

	s.AvgPrice = allItemsPrice / float64(len(tsp))
	s.TopSellingItems = tsp

	allItems := 0
	orders, err := database.GetOrdersAfterDate(0)
	if err == nil {
		for _, order := range orders {
			allItems += order.Items
		}
		s.ItemsPerOrder = float64(allItems) / float64(len(orders))
	}

	response.JSON(w, 200, s)
}
