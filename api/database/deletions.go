package database

func DeleteCustomer(shopifyCustomerId int64) error {
	_, err := DB.Exec("DELETE FROM `customers` WHERE `shopifyId` = ?;", shopifyCustomerId)
	if err != nil {
		return err
	}
	return nil
}

func DeleteProduct(shopifyVariantId int64) error {
	_, err := DB.Exec("DELETE FROM `products` WHERE `shopifyVariantId` = ?;", shopifyVariantId)
	if err != nil {
		return err
	}
	return nil
}
