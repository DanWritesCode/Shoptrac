package shopify

import (
	"../data"
	"../database"

	goshopify "github.com/bold-commerce/go-shopify"

	"time"
)

type CustomerListOptions struct {
	SortKey string `url:"sortKey,omitempty"`
	First   int    `url:"first,omitempty"`
}

func ImportCustomers() error {
	// 0) TODO: get count of customers (new, returning, and total) and store it somewhere in the database (?)

	options := CustomerListOptions{
		SortKey: "TOTAL_SPENT",
		First:   10,
	}

	customers, err := ShopifyClient.Client.Customer.List(options)
	if err != nil {
		return err
	}

	var topCustomers []*data.Customer
	for _, customer := range customers {
		topCustomers = append(topCustomers, &data.Customer{
			Name:        customer.FirstName + " " + customer.LastName,
			Country:     customer.DefaultAddress.Country,
			OrdersMade:  customer.OrdersCount,
			AmountSpent: customer.TotalSpent.InexactFloat64(),
		})
	}

	return database.BulkInsertCustomers(topCustomers)
}

func ImportOrders() error {
	// Create standard CountOptions
	date := time.Now().Add(time.Hour * 24 * -30)
	options := goshopify.OrderListOptions{ProcessedAtMin: date}

	// Use the options when calling the API.
	orders, err := ShopifyClient.Client.Order.List(options)
	if err != nil {
		return err
	}

	var dbOrders []*data.Order
	var orderProducts []*data.OrderProduct

	for _, order := range orders {
		newOrder := data.Order{
			OrderID:     order.AppID,
			Date:        order.CreatedAt.Unix(),
			Items:       len(order.LineItems),
			Country:     order.ShippingAddress.Country,
			TotalAmount: order.TotalPrice.InexactFloat64(),
			COGS:        0,
		}

		dbOrders = append(dbOrders, &newOrder)

		for _, item := range order.LineItems {
			itemS := data.OrderProduct{
				ShopifyOrderId:   order.AppID,
				ShopifyVariantId: item.VariantID,
				Quantity:         item.Quantity,
			}

			orderProducts = append(orderProducts, &itemS)
		}
	}

	err = database.BulkInsertOrders(dbOrders)
	if err != nil {
		return err
	}

	return database.BulkInsertOrderProduct(orderProducts)
}

func ImportProducts() error {
	products, err := ShopifyClient.Client.Product.List(goshopify.ProductListOptions{})
	if err != nil {
		return err
	}

	var productArr []*data.Product
	for _, product := range products {
		for _, variant := range product.Variants {
			p := data.Product{
				ShopifyVariantId: variant.ID,
				ItemName:         product.Title,
				VariantName:      variant.Title,
				Price:            variant.Price.InexactFloat64(),
			}

			productArr = append(productArr, &p)
		}
	}

	return database.BulkInsertProducts(productArr)
}

func GenerateDailyRevenue() error {
	// check the daily revenue database to find last day revenue has been generated for
	rev, err := database.GetLastDailyRevenue()
	if err != nil {
		return err
	}

	// get the orders from sql (WHERE date > last date of known revenue)
	orders, err2 := database.GetOrdersAfterDate(rev.Date)
	if err2 != nil {
		return err2
	}

	// Group Subtotal, Shipping, Taxes, Tips, Total per day
	revByDate := make(map[int64]*data.Revenue)
	for _, order := range orders {
		date := time.Unix(order.Date, 0).Truncate(24 * time.Hour).Unix() // get start of day epoch
		rev := &data.Revenue{}
		if revByDate[date] != nil {
			rev = revByDate[date]
		} else {
			revByDate[date] = rev
		}

		rev.Date = date
		rev.Sales += order.Subtotal
		rev.ShippingCharged += order.Shipping
		rev.TaxesCollected += order.Taxes
		rev.Tips += order.Tips
	}

	allRev := make([]*data.Revenue, len(revByDate))
	for i, rev := range revByDate {
		allRev[i] = rev
	}

	return database.BulkInsertRevenue(allRev)
}
