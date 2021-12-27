package endpoints

import (
	"../config"
	"../data"
	"../database"
	"../logging"
	"../response"
	"../shopify"
	"encoding/json"
	"fmt"
	"math"
	"math/rand"
	"net/http"
	"net/url"
)

func GetInstall(w http.ResponseWriter, r *http.Request) {
	shopName := r.URL.Query().Get("shop")
	nonce := fmt.Sprintf("%d", rand.Intn(math.MaxInt32))
	shopify.AuthorizeNonce(nonce)

	http.Redirect(w, r, shopify.ShopifyClient.AuthorizeUrl(shopName, nonce), http.StatusFound)
}

func PostInstall(w http.ResponseWriter, r *http.Request) {
	id := data.InstallData{}
	err := json.NewDecoder(r.Body).Decode(&id)
	if err != nil {
		response.Error(w, response.BadRequestError)
		return
	}

	if !shopify.CheckNonce(id.Nonce) {
		response.BadRequest(w, "Shopify Security Check Failed")
		return
	}

	if !shopify.CheckShopDomain(id.Shop) {
		response.BadRequest(w, "Invalid Shopify Domain/Shop Name")
		return
	}

	realURL, _ := url.Parse(id.URL)

	// validate the request
	if ok, authErr := shopify.ShopifyClient.VerifyAuthorizationURL(realURL); ok {
		// request access token from shopify
		sat, err := shopify.RequestAccessToken(id.Shop, config.GetConfig().App.ClientId, config.GetConfig().App.ClientSecret, id.AuthCode)
		if err == nil {
			// insert AT into DB
			err := database.SetDatabaseConfig("shopAccessToken", sat.AccessToken)
			if err != nil {
				logging.GetLogger().Println("Unable to set database config (1) - " + err.Error())
				response.BadRequest(w, "Shopify Installation Failed - Database Insert")
				return
			}

			err = database.SetDatabaseConfig("shopScope", sat.Scope)
			if err != nil {
				logging.GetLogger().Println("Unable to set database config (2) - " + err.Error())
				response.BadRequest(w, "Shopify Installation Failed - Database Insert")
				return
			}

			response.DefaultResponse(w, 200, map[string]string{"message": "Installation Success"})
		} else {
			logging.GetLogger().Println("Unable to obtain access token from Shopify - " + err.Error())
			response.BadRequest(w, "Shopify Installation Failed")
		}
	} else {
		errMsg := ""
		if authErr != nil {
			errMsg = "  Error: " + authErr.Error()
		}
		logging.GetLogger().Println("Spotify Authorization Check Failed. URL: " + id.URL + "  Nonce: " + id.Nonce + "  HMAC: " + id.HMAC + "  Shop: " + id.Shop + "  Host: " + id.Host + "  Error: " + errMsg)
		response.BadRequest(w, "Shopify Authorization Check Failed")
	}
}
