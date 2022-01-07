package database

import (
	"../data"
	"database/sql"
)

type Shop struct {
	ShopName        string `json:"shopName"`
	ShopAccessToken string `json:"shopAccessToken"`
}

func GetShop() (*Shop, error) {
	shop := Shop{}
	// Query for a value based on a single row.
	rows, err := DB.Query("SELECT `name`, `value` FROM `config` WHERE `name` = ? OR `name` = ?;", "shopAccessToken", "shopName")
	if err != nil {
		return &shop, err
	}
	for rows.Next() {
		var key, val string
		err = rows.Scan(&key, &val)
		if err != nil {
			return &shop, err
		}

		if key == "shopAccessToken" {
			shop.ShopAccessToken = val
		} else if key == "shopName" {
			shop.ShopName = val
		}
	}

	return &shop, nil
}

func GetLastOrder() (*data.Order, error) {
	rev := data.Order{}

	// Query for a value based on a single row.
	if err := DB.QueryRow("SELECT `orderId`, `date`, `items`, `country`, `paymentGateway`, `discount`, `subtotal`, `shipping`, `taxes`, `tips`, `totalAmount`, `cogs` FROM orders ORDER BY date DESC LIMIT 1;").Scan(
		&rev.OrderID, &rev.Date, &rev.Items, &rev.Country, &rev.PaymentGateway, &rev.Discount, &rev.Subtotal, &rev.Shipping, &rev.Taxes, &rev.Tips, &rev.TotalAmount, &rev.COGS); err != nil {
		return nil, err
	}

	return &rev, nil
}

func GetLastDailyRevenue() (*data.Revenue, error) {
	rev := data.Revenue{}

	// Query for a value based on a single row.
	if err := DB.QueryRow("SELECT `date`, sales, shipping, taxes, tips, discounts, total FROM `dailyRevenue` ORDER BY date DESC LIMIT 1;").Scan(
		&rev.Date, &rev.Sales, &rev.ShippingCharged, &rev.TaxesCollected, &rev.Tips, &rev.Discounts, &rev.Total); err != nil {
		return nil, err
	}
	return &rev, nil
}

func GetDailyRevenue(from int64) ([]*data.Revenue, error) {
	revArr := make([]*data.Revenue, 0)

	// Query for a value based on a single row.
	rows, err := DB.Query("SELECT `date`, sales, shipping, taxes, tips, discounts, total FROM `dailyRevenue` WHERE `date` > ? ORDER BY date DESC;", from)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	// Loop through rows, using Scan to assign column data to struct fields.
	err = nil
	for rows.Next() {
		rev := data.Revenue{}
		err2 := rows.Scan(&rev.Date, &rev.Sales, &rev.ShippingCharged, &rev.TaxesCollected, &rev.Tips, &rev.Discounts, &rev.Total)
		if err2 != nil {
			err = err2
		}

		revArr = append(revArr, &rev)
	}

	return revArr, err
}

func GetCustomers() ([]*data.Customer, error) {
	return GetTopCustomersByRevenue(-1)
}

func GetTopCustomersByRevenue(limit int) ([]*data.Customer, error) {
	revArr := make([]*data.Customer, 0)
	var rows *sql.Rows
	var err error

	if limit < 0 {
		// no limit
		rows, err = DB.Query("SELECT `shopifyId`, `name`, `country`, `ordersMade`, amountSpent FROM `customers` ORDER BY amountSpent DESC;")
	} else {
		// limit amount of results to value of limit
		rows, err = DB.Query("SELECT `shopifyId`, `name`, `country`, `ordersMade`, amountSpent FROM `customers` ORDER BY amountSpent DESC LIMIT ?;", limit)
	}

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	// Loop through rows, using Scan to assign column data to struct fields.
	err = nil
	for rows.Next() {
		rev := data.Customer{}
		err2 := rows.Scan(&rev.ShopifyId, &rev.Name, &rev.Country, &rev.OrdersMade, &rev.AmountSpent)
		if err2 != nil {
			err = err2
		}

		revArr = append(revArr, &rev)
	}

	return revArr, err
}

func GetAllProducts() ([]*data.Product, error) {
	retArr := make([]*data.Product, 0)
	rows, err := DB.Query("SELECT `id`, `shopifyVariantId`, `itemName`, `variantName`,`price` FROM products;")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		rev := data.Product{}
		if err := rows.Scan(&rev.ID, &rev.ShopifyVariantId, &rev.ItemName, &rev.VariantName, &rev.Price); err != nil {
			continue
		}
		retArr = append(retArr, &rev)
	}

	return retArr, nil
}

func GetTopSellingProducts(date int64, limit int) ([]*data.TopSeller, error) {
	revArr := make([]*data.TopSeller, 0)

	// quickly grab a stat from the orders table. we'll need this number to do % of Sales calculations later on
	totalItemsOrdered := 0
	if err := DB.QueryRow("SELECT SUM(`items`) FROM `orders` WHERE `date` > ?;", date).Scan(
		&totalItemsOrdered); err != nil {
		return nil, err
	}

	// now for the good part

	// Credits to Thomas / Github: @Period for assisting with this query
	// Still took an hour to write though :/
	rows, err := DB.Query(
		"SELECT `products`.`id`, `products`.`shopifyVariantId`, `products`.`itemName`, `products`.`variantName`, `products`.`price`, SUM(`orderProduct`.`quantity`) AS `unitsSold` FROM `orderProduct` "+
			"INNER JOIN `products` on `products`.`shopifyVariantId` = `orderProduct`.`shopifyVariantId` "+
			"INNER JOIN `orders` on `orders`.`orderId` = `orderProduct`.`shopifyOrderId` "+
			"WHERE `orders`.`date` > ? "+
			"GROUP BY `orderProduct`.`shopifyVariantId` "+
			"ORDER BY `unitsSold` DESC LIMIT ?;",
		date, limit)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	// Loop through rows, using Scan to assign column data to struct fields.
	err = nil
	for rows.Next() {
		rev := data.TopSeller{}
		err2 := rows.Scan(&rev.Product.ID, &rev.Product.ShopifyVariantId, &rev.Product.ItemName, &rev.Product.VariantName, &rev.Product.Price, &rev.QuantitySold)
		if err2 != nil {
			err = err2
		}

		rev.AmountSold = float64(rev.QuantitySold) * rev.Product.Price
		rev.PercentageOfSales = float64(rev.QuantitySold) / float64(totalItemsOrdered) * 100

		revArr = append(revArr, &rev)
	}

	return revArr, err
}

func GetOrdersAfterDate(date int64) ([]*data.Order, error) {
	var revArr []*data.Order

	// Query for a value based on a single row.
	rows, err := DB.Query("SELECT `orderId`, `date`, `items`, `country`, `paymentGateway`, `discount`, `subtotal`, `shipping`, `taxes`, `tips`, `totalAmount`, `cogs` FROM orders WHERE `date` > ?;", date)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		rev := data.Order{}
		if err := rows.Scan(&rev.OrderID, &rev.Date, &rev.Items, &rev.Country, &rev.PaymentGateway, &rev.Discount,
			&rev.Subtotal, &rev.Shipping, &rev.Taxes, &rev.Tips, &rev.TotalAmount, &rev.COGS); err != nil {
			continue
		}
		revArr = append(revArr, &rev)
	}

	return revArr, nil
}
