package endpoints

import (
	"../config"
	"../data"
	"../database"
	"../logging"
	"../response"
	"../shopify"
	"encoding/json"
	"net/http"
)

func PostInstall(w http.ResponseWriter, r *http.Request) {
	id := data.InstallData{}
	err := json.NewDecoder(r.Body).Decode(&id)

	if err != nil {
		response.Error(w, response.BadRequestError)
		return
	}

	// temp - this will be removed once the nonce & nonce source are determined
	shopify.AuthorizeNonce(id.Nonce)

	if shopify.SecurityCheck(id.URL, id.Nonce, id.HMAC, id.Shop) {
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
		logging.GetLogger().Println("Spotify Security Check Failed. URL: " + id.URL + "  Nonce: " + id.Nonce + "  HMAC: " + id.HMAC + "  Shop: " + id.Shop + "  Host: " + id.Host)
		response.BadRequest(w, "Shopify Security Check Failed")
	}
}
