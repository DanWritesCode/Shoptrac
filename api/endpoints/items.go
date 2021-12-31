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
				ID:               0,
				ShopifyVariantId: 0,
				ItemName:         "Item",
				VariantName:      "Variant 1",
				Price:            10,
			},
			QuantitySold:      10,
			PercentageOfSales: 50,
			AmountSold:        60,
		}, {
			Product: data.Product{
				ID:               0,
				ShopifyVariantId: 0,
				ItemName:         "Item 2",
				VariantName:      "Variant 2",
				Price:            20,
			},
			QuantitySold:      10,
			PercentageOfSales: 50,
			AmountSold:        60,
		}},
	}
	response.JSON(w, 200, s)
}
