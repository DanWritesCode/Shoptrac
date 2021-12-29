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

	var topCustomers []data.Customer
	for _, customer := range customers {
		topCustomers = append(topCustomers, data.Customer{
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

	var dbOrders []data.Order
	for _, order := range orders {
		newOrder := data.Order{
			OrderID:     int(order.ID),
			Date:        order.CreatedAt.Unix(),
			Items:       len(order.LineItems),
			Country:     order.ShippingAddress.Country,
			TotalAmount: order.TotalPrice.InexactFloat64(),
			COGS:        0,
		}
		dbOrders = append(dbOrders, newOrder)
	}

	// TODO: for each item in an order, add to sql table mapping item IDs to orders (Order# - ItemID - Quantity)

	return database.BulkInsertOrders(dbOrders)
}

/*func GenerateDailyRevenue() error {
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

  // TODO: group Subtotal, Shipping, Taxes, Tips, Total per day
  var revByDate map[int64]*data.Revenue
  for _, order := range orders {

  }


  // TODO: store it in a data structure & insert to SQL

  return nil
}
*/
