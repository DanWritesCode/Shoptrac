package database

import (
	"../data"
)

func BulkInsertOrders(orders []*data.Order) error {
	stmt, err := DB.Prepare("INSERT INTO orders (id, `orderId`, `date`, `items`, `country`, `paymentGateway`, `subtotal`, `shipping`, `taxes`, `tips`, `totalAmount`, `cogs`) VALUES (NULL, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);")
	defer stmt.Close()
	if err != nil {
		return err
	}

	for _, order := range orders {
		_, _ = stmt.Exec(order.OrderID, order.Date, order.Items, order.Country, order.PaymentGateway, order.Subtotal, order.Shipping, order.Taxes, order.Tips, order.TotalAmount, order.COGS)
	}

	return nil
}

func BulkInsertCustomers(customers []*data.Customer) error {
	stmt, err := DB.Prepare("INSERT INTO customers (id, `shopifyId`, `name`, `country`, `ordersMade`, `amountSpent`) VALUES (NULL, ?, ?, ?, ?, ?);")
	defer stmt.Close()
	if err != nil {
		return err
	}

	for _, customer := range customers {
		_, err = stmt.Exec(customer.ShopifyId, customer.Name, customer.Country, customer.OrdersMade, customer.AmountSpent)
		if err != nil {
			return err
		}
	}

	return nil
}

func BulkInsertOrderProduct(ts []*data.OrderProduct) error {
	stmt, err := DB.Prepare("INSERT INTO orderProduct (`shopifyOrderId`, `shopifyVariantId`,`quantity`) VALUES (?, ?, ?);")
	defer stmt.Close()
	if err != nil {
		return err
	}

	err = nil
	for _, t := range ts {
		_, err = stmt.Exec(t.ShopifyOrderId, t.ShopifyVariantId, t.Quantity)
		// no immediate error handling. attempt all insertions first
	}

	// this will return either the error or nil
	return err
}

func BulkInsertProducts(ts []*data.Product) error {
	stmt, err := DB.Prepare("INSERT INTO products (id, `shopifyVariantId`, `itemName`, `variantName`,`price`) VALUES (NULL, ?, ?, ?, ?);")
	defer stmt.Close()
	if err != nil {
		return err
	}

	for _, t := range ts {
		_, err = stmt.Exec(t.ShopifyVariantId, t.ItemName, t.VariantName, t.Price)
		if err != nil {
			return err
		}
	}

	return nil
}

func BulkInsertRevenue(rev []*data.Revenue) error {
	stmt, err := DB.Prepare("INSERT INTO dailyRevenue (id, `date`, `sales`, `shipping`, `taxes`, `tips`, `discounts`, `total`) VALUES (NULL, ?, ?, ?, ?, ?, ?, ?);")
	defer stmt.Close()
	if err != nil {
		return err
	}

	for _, r := range rev {
		_, err = stmt.Exec(r.Date, r.Sales, r.ShippingCharged, r.TaxesCollected, r.Tips, r.Discounts, r.Total)
		if err != nil {
			return err
		}
	}

	return nil
}
