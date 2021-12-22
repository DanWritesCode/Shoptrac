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
			Item:              "Test Item",
			QuantitySold:      10,
			PercentageOfSales: 50,
			AmountSold:        50,
		}, {
			Item:              "Another Test Item",
			QuantitySold:      10,
			PercentageOfSales: 50,
			AmountSold:        50,
		}},
		TopSellingCollections: []data.TopSeller{
			{
				Item:              "Test Item Collection",
				QuantitySold:      20,
				PercentageOfSales: 100,
				AmountSold:        100,
			},
		},
	}
	response.JSON(w, 200, s)
}
