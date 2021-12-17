package endpoints

import (
	"../data"
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

	if shopify.SecurityCheck(id.URL, id.Nonce, id.HMAC, id.ShopName) {
		// request access token from shopify
	} else {
		response.BadRequest(w, "Shopify Security Check Failed")
	}
}
