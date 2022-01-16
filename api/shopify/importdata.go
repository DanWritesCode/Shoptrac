package shopify

import (
	"../data"
	"../database"
	"../logging"

	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
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

	err5 := ImportFacebookAdExpenses()
	if err5 != nil {
		//fmt.Println(err5.Error())
		logging.GetLogger().Println(err5.Error())
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
				// re-insert the customer after deleting
				err = database.DeleteCustomer(customer.ID)
				if err != nil {
					continue
				}
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
	// TODO - how to handle an existing order changing?

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
			Discount:       order.TotalDiscounts.InexactFloat64(),
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
					// re-insert the product after deleting
					err = database.DeleteProduct(variant.ID)
					if err != nil {
						continue
					}
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

type FacebookInsights struct {
	Data []struct {
		Spend     string `json:"spend"`
		DateStart string `json:"date_start"`
		DateStop  string `json:"date_stop"`
	} `json:"data"`
	Paging struct {
		Cursors struct {
			Before string `json:"before"`
			After  string `json:"after"`
		} `json:"cursors"`
		Next string `json:"next,omitempty"`
	} `json:"paging"`
}

func ImportFacebookAdExpenses() error {
	expense, err := database.GetLatestExpenseByName("Facebook")
	if err != nil {
		return err
	}

	var last int64
	if expense != nil {
		last = expense.Date
	} else {
		// current timestamp minus 60 days
		// essentially get last 60 days of FB data
		// TODO: a setting here that checks if the user is able to get extended historical data throughout the application
		//       if so, then increase the range we pull from FB
		//last = time.Now().Unix()-94610000 // 1096 days / 36 months
		last = time.Now().Unix() - 5184000 // 60 days
	}

	if last <= (time.Now().Unix() - (26280 * 60 * 60)) {
		last = time.Now().Unix() - (26280 * 60 * 60)
	}

	since := time.Unix(last, 0).Truncate(24 * time.Hour).Format("2006-01-02")
	until := time.Now().Truncate(24 * time.Hour).Format("2006-01-02")
	if since == until {
		// nothing to update, we're all caught up
		// TODO: i suppose technically we could update today's information ...
		return nil
	}

	// TODO check for ad account expiry
	fbAddAc, _ := database.GetDatabaseConfig("facebookAdAccountId")
	fbToken, _ := database.GetDatabaseConfig("facebookAccessToken")

	req, err := http.NewRequest("GET", fmt.Sprintf("https://graph.facebook.com/v12.0/act_%v/insights?fields=spend&level=account&time_range={since:'%v',until:'%v'}&time_increment=1&limit=1096", fbAddAc, since, until), nil)
	if err != nil {
		return err
	}
	req.Header.Add("Authorization", "Bearer "+fbToken)
	resp, err := Client.Do(req)
	if err != nil {
		return err
	}
	body, _ := ioutil.ReadAll(resp.Body)

	if resp.StatusCode == 200 {

		fbR := FacebookInsights{}
		err = json.Unmarshal(body, &fbR)
		if err != nil {
			return err
		}

		expensesArr := make([]*data.Expense, 0)

		for _, fbRDay := range fbR.Data {
			dateSplat, err := time.Parse("2006-01-02", fbRDay.DateStart)
			if err != nil {
				continue
			}

			spend, _ := strconv.ParseFloat(fbRDay.Spend, 64)
			exp := data.Expense{
				Category: "MARKETING",
				Name:     "Facebook",
				Date:     dateSplat.Unix(),
				Amount:   spend,
			}

			expensesArr = append(expensesArr, &exp)
		}

		err2 := database.BulkInsertExpenses(expensesArr)
		if err2 != nil {
			return err2
		}
	}

	return nil
}

func GenerateDailyRevenue() error {
	// TODO - how to handle daily revenue changing due to an order changing?

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
