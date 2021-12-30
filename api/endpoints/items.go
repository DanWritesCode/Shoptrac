package endpoints

import (
	"../data"
	"../response"
	"net/http"
)

func GetItems(w http.ResponseWriter, r *http.Request) {
	s := data.Products{
		AvgPrice:       0,
		ItemsPerOrder:  0,
		InventoryValue: 0,
		TopSellingItems: []data.TopSeller{{
			Product: data.Product{
				Item:    "Item 1",
				Variant: "Variant 1",
				Price:   0,
			},
			QuantitySold:      10,
			PercentageOfSales: 50,
			AmountSold:        60,
		}, {
			Product: data.Product{
				Item:    "Item 2",
				Variant: "Variant 2",
				Price:   0,
			},
			QuantitySold:      10,
			PercentageOfSales: 50,
			AmountSold:        60,
		}},
	}
	response.JSON(w, 200, s)
}
