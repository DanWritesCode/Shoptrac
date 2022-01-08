package endpoints

import (
	"../config"
	"../logging"
	"../response"
	"fmt"
	"math"
	"math/rand"
	"net/http"
)

type ConfigRequest struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

var facebookState = make(map[string]bool)

func GetFacebook(w http.ResponseWriter, r *http.Request) {
	if config.GetConfig().FacebookAuth == nil {
		logging.GetLogger().Println("Facebook Ads App not configured")
		response.BadRequest(w, "Facebook Ads Config Failed")
		return
	}

	nonce := fmt.Sprintf("%d", rand.Intn(math.MaxInt32))
	facebookState[nonce] = true

	http.Redirect(w, r, fmt.Sprintf("https://www.facebook.com/v12.0/dialog/oauth?client_id=%v&redirect_uri=%v&state=%v&scope=ads_read",
		config.GetConfig().FacebookAuth.ClientId, config.GetConfig().FacebookAuth.RedirectURL, nonce), http.StatusFound)
}

func PostConfig(w http.ResponseWriter, r *http.Request) {

}
