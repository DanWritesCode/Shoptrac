package endpoints

import (
	"../database"
	"../response"
	"../shopify"
	"net/http"
)

func GetSync(w http.ResponseWriter, r *http.Request) {
	var lastUpdate int64
	lastUpdate = 0

	rev, err := database.GetLastDailyRevenue()
	if err == nil {
		lastUpdate = rev.Date
	}

	response.JSON(w, 200, map[string]int64{"lastUpdate": lastUpdate})
}

func PostSync(w http.ResponseWriter, r *http.Request) {
	shop, err := database.GetShop()
	if err != nil {
		response.BadRequest(w, "Could not obtain Shopify access token for re-sync. Try re-installing the app.")
		return
	}

	if shopify.IsImportInProgress(shop.ShopName) {
		response.BadRequest(w, "A sync is already in progress!")
		return
	}

	// start the import process asynchronously
	go shopify.DataImportProcess(shop.ShopName, shopify.ShopifyClient.NewClient(shop.ShopName, shop.ShopAccessToken))

	response.DefaultResponse(w, 200, map[string]string{"message": "Successfully requested sync"})
}
