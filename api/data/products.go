package data

type Products struct {
  AvgPrice float64 `json:"avgPrice"`
  ItemsPerOrder float64 `json:"itemsPerOrder"`
  InventoryValue  float64 `json:"inventoryValue"`

  TopSellingItems   []TopSeller `json:"topSellingItems"`
  TopSellingCollections   []TopSeller `json:"topSellingCollections"`
}

type TopSeller struct {
  Item string `json:"item"`
  QuantitySold int `json:"quantitySold"`
  PercentageOfSales float64 `json:"percentageSales"`
  AmountSold      float64   `json:"amountSold"`
}
