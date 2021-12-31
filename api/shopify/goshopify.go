package shopify

import (
	"github.com/bold-commerce/go-shopify"
)

var ShopifyClient *goshopify.App

func NewShopify(apiKey, apiSecret, redirectURL string) {
	ShopifyClient = &goshopify.App{
		ApiKey:      apiKey,
		ApiSecret:   apiSecret,
		RedirectUrl: redirectURL,
		Scope:       "read_products,read_orders,read_customers",
	}
}
