package shopify

import (
	"../data"
	"../database"
	"../logging"
	"strings"

	goshopify "github.com/bold-commerce/go-shopify"

	"time"
)

type CustomerListOptions struct {
	SortKey string `url:"sortKey,omitempty"`
	First   int    `url:"first,omitempty"`
}

var importsInProgress = make(map[string]bool)

func IsImportInProgress(shopName string) bool {
	return importsInProgress[shopName]
}

func DataImportProcess(shopName string, client *goshopify.Client) {
	if importsInProgress[shopName] {
		return
	}

	importsInProgress[shopName] = true

	err := ImportCustomers(client)
	if err != nil {
		logging.GetLogger().Println(err.Error())
	}

	err = ImportNewOrders(client)
	if err != nil {
		logging.GetLogger().Println(err.Error())
	}

	err = ImportProducts(client)
	if err != nil {
		logging.GetLogger().Println(err.Error())
	}

	err4 := GenerateDailyRevenue()
	if err4 != nil {
		logging.GetLogger().Println(err4.Error())
	}

	importsInProgress[shopName] = false
}

func ImportCustomers(client *goshopify.Client) error {
	existingCustomers, _ := database.GetCustomers()
	existingCustomersMap := make(map[int64]*data.Customer)
	if existingCustomers != nil {
		for _, customer := range existingCustomers {
			existingCustomersMap[customer.ShopifyId] = customer
		}
	}

	options := CustomerListOptions{
		SortKey: "TOTAL_SPENT",
		First:   10,
	}

	customers, err := client.Customer.List(options)
	if err != nil {
		return err
	}

	var topCustomers []*data.Customer
	for _, customer := range customers {
		if existingCustomersMap[customer.ID] != nil {
			if existingCustomersMap[customer.ID].AmountSpent == customer.TotalSpent.InexactFloat64() &&
				existingCustomersMap[customer.ID].OrdersMade == customer.OrdersCount {
				continue
			} else {
				// TODO: delete customer and re-insert
			}
		}
		topCustomers = append(topCustomers, &data.Customer{
			ShopifyId:   customer.ID,
			Name:        customer.FirstName + " " + customer.LastName,
			Country:     customer.DefaultAddress.Country,
			OrdersMade:  customer.OrdersCount,
			AmountSpent: customer.TotalSpent.InexactFloat64(),
		})
	}

	return database.BulkInsertCustomers(topCustomers)
}

func ImportNewOrders(client *goshopify.Client) error {
	// Check existing orders
	lastOrder, err := database.GetLastOrder()
	if err != nil {
		if strings.Contains(err.Error(), "no rows") {
			lastOrder = &data.Order{Date: 0}
		} else {
			return err
		}
	}

	// add 1 second to the last order's date
	if lastOrder.Date > 0 {
		lastOrder.Date++
	}

	// Only get new orders that are 1 second after the last order's date
	options := goshopify.OrderListOptions{ProcessedAtMin: time.Unix(lastOrder.Date, 0)}

	// Use the options when calling the API.
	orders, err := client.Order.List(options)
	if err != nil {
		return err
	}

	var dbOrders []*data.Order
	var orderProducts []*data.OrderProduct

	for _, order := range orders {
		country := ""
		if order.ShippingAddress != nil {
			country = order.ShippingAddress.Country
		}
		newOrder := data.Order{
			OrderID:        order.OrderNumber,
			Date:           order.CreatedAt.Unix(),
			Items:          len(order.LineItems),
			Country:        country,
			PaymentGateway: order.Gateway,
			Subtotal:       order.TotalLineItemsPrice.InexactFloat64(),
			Shipping:       order.SubtotalPrice.InexactFloat64() - order.TotalLineItemsPrice.InexactFloat64(),
			Taxes:          order.TotalTax.InexactFloat64(),
			Tips:           0,
			TotalAmount:    order.TotalPrice.InexactFloat64(),
			COGS:           0,
		}

		dbOrders = append(dbOrders, &newOrder)

		for _, item := range order.LineItems {
			itemS := data.OrderProduct{
				ShopifyOrderId:   order.OrderNumber,
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

func ImportProducts(client *goshopify.Client) error {
	existingProducts, _ := database.GetAllProducts()
	existingProductsMap := make(map[int64]*data.Product)
	if existingProducts != nil {
		for _, product := range existingProducts {
			existingProductsMap[product.ShopifyVariantId] = product
		}
	}

	products, err := client.Product.List(goshopify.ProductListOptions{})
	if err != nil {
		return err
	}

	var productArr []*data.Product
	for _, product := range products {
		for _, variant := range product.Variants {
			if existingProductsMap[variant.ID] != nil {
				// if nothing changed
				if existingProductsMap[variant.ID].Price == variant.Price.InexactFloat64() &&
					existingProductsMap[variant.ID].VariantName == variant.Title &&
					existingProductsMap[variant.ID].ItemName == product.Title {
					continue
				} else {
					// if something about the product changed and needs updating
					// TODO: delete product from SQL and re-add
				}
			}
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
		if strings.Contains(err.Error(), "no rows") {
			rev = &data.Revenue{Date: 0}
		} else {
			return err
		}
	}

	// search date should be +1 day from last day we have revenue info for
	// unless we dont have any revenue info at all, in which case get all orders > 0
	if rev.Date > 0 {
		rev.Date += 86400
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
		rev.Total += order.TotalAmount
	}

	allRev := make([]*data.Revenue, len(revByDate))
	pos := 0
	for _, rev := range revByDate {
		allRev[pos] = rev
		pos++
	}

	return database.BulkInsertRevenue(allRev)
}
