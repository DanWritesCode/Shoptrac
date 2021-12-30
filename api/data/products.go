package data

type Products struct {
	AvgPrice       float64 `json:"avgPrice"`
	ItemsPerOrder  float64 `json:"itemsPerOrder"`
	InventoryValue float64 `json:"inventoryValue"`

	TopSellingItems []TopSeller `json:"topSellingItems"`
	// TopSellingCollections   []TopSeller `json:"topSellingCollections"`
}

type TopSeller struct {
	Product           Product `json:"product"`
	QuantitySold      int     `json:"quantitySold"`
	PercentageOfSales float64 `json:"percentageSales"`
	AmountSold        float64 `json:"amountSold"`
}

type Product struct {
	ID               int     `json:"id,omitempty"`
	ShopifyVariantId int64   `json:"shopifyVariantId"`
	ItemName         string  `json:"itemName"`
	VariantName      string  `json:"variantName"`
	Price            float64 `json:"price"`
}

type OrderProduct struct {
	ShopifyOrderId   int   `json:"shopifyOrderId"`
	ShopifyVariantId int64 `json:"ShopifyVariantId"`
	Quantity         int   `json:"quantity"`
}
