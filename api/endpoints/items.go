package endpoints

import (
  "net/http"
  "../data"
  "../response"
)

func GetItems (w http.ResponseWriter, r *http.Request) {
  s := data.Products{
    AvgPrice:              0,
    ItemsPerOrder:         0,
    InventoryValue:        0,
    TopSellingItems:       nil,
    TopSellingCollections: nil,
  }
  response.JSON(w, 200, s)
}
