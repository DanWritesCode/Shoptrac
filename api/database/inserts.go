package database

import (
	"../data"
)

func BulkInsertOrders(orders []data.Order) error {
	stmt, err := DB.Prepare("INSERT INTO orders (id, `orderId`, `date`, `items`, `country`, `paymentGateway`, `subtotal`, `shipping`, `taxes`, `tips`, `totalAmount`, `cogs`) VALUES (NULL, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);")
	defer stmt.Close()
	if err != nil {
		return err
	}

	for _, order := range orders {
		_, err = stmt.Exec(order.OrderID, order.Date, order.Items, order.Country, order.PaymentGateway, order.Subtotal, order.Shipping, order.Taxes, order.Tips, order.TotalAmount, order.COGS)
		if err != nil {
			return err
		}
	}

	return nil
}

func BulkInsertCustomers(customers []data.Customer) error {
	stmt, err := DB.Prepare("INSERT INTO customers (id, `name`, `country`, `ordersMade`, `amountSpent`) VALUES (NULL, ?, ?, ?, ?);")
	defer stmt.Close()
	if err != nil {
		return err
	}

	for _, customer := range customers {
		_, err = stmt.Exec(customer.Name, customer.Country, customer.OrdersMade, customer.AmountSpent)
		if err != nil {
			return err
		}
	}

	return nil
}

// used for both products and collections
func BulkInsertProduct(ts []data.TopSeller) error {
	stmt, err := DB.Prepare("INSERT INTO product (id, `name`, `quantity`, `percentage`, `amount`) VALUES (NULL, ?, ?, ?, ?);")
	defer stmt.Close()
	if err != nil {
		return err
	}

	for _, t := range ts {
		_, err = stmt.Exec(t.Item, t.QuantitySold, t.PercentageOfSales, t.AmountSold)
		if err != nil {
			return err
		}
	}

	return nil
}

func BulkInsertRevenue(rev []data.Revenue) error {
	stmt, err := DB.Prepare("INSERT INTO dailyRevenue (id, `date`, `sales`, `shipping`, `taxes`, `tips`, `discounts`) VALUES (NULL, ?, ?, ?, ?, ?, ?);")
	defer stmt.Close()
	if err != nil {
		return err
	}

	for _, r := range rev {
		_, err = stmt.Exec(r.Date, r.Sales, r.ShippingCharged, r.TaxesCollected, r.Tips, r.Discounts)
		if err != nil {
			return err
		}
	}

	return nil
}
