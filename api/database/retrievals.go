package database

import (
	"../data"
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

func GetLastDailyRevenue() (*data.Revenue, error) {
	rev := data.Revenue{}

	// Query for a value based on a single row.
	if err := DB.QueryRow("SELECT `date`, sales, shipping, taxes, tips, discounts FROM `dailyRevenue` ORDER BY date DESC LIMIT 1;").Scan(
		&rev.Date, &rev.Sales, &rev.ShippingCharged, &rev.TaxesCollected, &rev.Tips, &rev.Discounts); err != nil {
		return nil, err
	}
	return &rev, nil
}

func GetOrdersAfterDate(date int64) ([]*data.Order, error) {
	var revArr []*data.Order

	// Query for a value based on a single row.
	rows, err := DB.Query("SELECT `orderId`, `date`, `items`, `country`, `paymentGateway`, `subtotal`, `shipping`, `taxes`, `tips`, `totalAmount`, `cogs` FROM orders WHERE `date` > ?;", date)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		rev := data.Order{}
		if err := rows.Scan(&rev.OrderID, &rev.Date, &rev.Items, &rev.Country, &rev.PaymentGateway,
			&rev.Subtotal, &rev.Shipping, &rev.Taxes, &rev.Tips, &rev.TotalAmount, &rev.COGS); err != nil {
			continue
		}
		revArr = append(revArr, &rev)
	}

	return revArr, nil
}
