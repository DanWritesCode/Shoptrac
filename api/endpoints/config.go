package endpoints

import (
	"../config"
	"../database"
	"../logging"
	"../response"
	"../shopify"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"math/rand"
	"net/http"
	"time"
)

type FacebookRequest struct {
	Code  string `json:"code"`
	State string `json:"state"`
}

type FacebookResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int64  `json:"expires_in"`
}

type ConfigRequest struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// manage the nonces/states associated with each oAuth request
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

func PostFacebook(w http.ResponseWriter, r *http.Request) {
	fbConf := config.GetConfig().FacebookAuth

	var fbr FacebookRequest
	err := json.NewDecoder(r.Body).Decode(&fbr)
	if err != nil || fbr.Code == "" || fbr.State == "" {
		response.BadRequest(w, "Required JSON fields missing from request")
		return
	}

	if facebookState[fbr.State] != true {
		response.BadRequest(w, "Invalid Nonce/State")
		return
	}

	res, err := shopify.Client.Get(fmt.Sprintf("https://graph.facebook.com/v12.0/oauth/access_token?client_id=%v&redirect_uri=%v&client_secret=%v&code=%v",
		fbConf.ClientId, fbConf.RedirectURL, fbConf.ClientSecret, fbr.Code))
	if err != nil {
		logging.GetLogger().Println("Could not obtain access token from facebook's graph API: " + err.Error())
		response.Error(w, response.BadRequestError)
		return
	}

	body, _ := ioutil.ReadAll(res.Body)

	if res.StatusCode == 200 {
		var fbResponse FacebookResponse
		err = json.Unmarshal(body, &fbResponse)
		if err != nil {
			logging.GetLogger().Println("Could not unmarshal request from facebook api: " + err.Error())
			response.Error(w, response.BadRequestError)
			return
		}

		_ = database.SetDatabaseConfig("facebookAccessToken", fbResponse.AccessToken)
		_ = database.SetDatabaseConfig("facebookTokenType", fbResponse.TokenType)
		_ = database.SetDatabaseConfig("facebookExpiresAt", fmt.Sprintf("%d", time.Now().Unix()+fbResponse.ExpiresIn))

		response.DefaultResponse(w, 200, map[string]string{"message": "Facebook Connection Success"})
		facebookState[fbr.State] = false
	} else {
		logging.GetLogger().Println("Could not obtain access token from facebook's graph API: " + string(body))
		response.Error(w, response.BadRequestError)
		return
	}

}

func PostConfig(w http.ResponseWriter, r *http.Request) {

}
