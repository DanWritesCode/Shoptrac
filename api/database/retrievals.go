package database

import (
	"../data"
)

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
	rows, err := DB.Query("SELECT orderId, `date`, amount, items, country, cogs FROM orders WHERE `date` > ?;", date)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		rev := data.Order{}
		if err := rows.Scan(&rev.OrderID, &rev.Date, &rev.TotalAmount,
			&rev.Items, &rev.Country, &rev.COGS); err != nil {
			continue
		}
		revArr = append(revArr, &rev)
	}

	return revArr, nil
}
