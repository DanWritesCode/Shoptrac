package shopify

import (
	"../data"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

var validNonces = make(map[string]bool)

func SecurityCheck(fullURL string, nonce string, hmac string, shop string) bool {
	reg := regexp.MustCompile(`\A(https|http)://[a-zA-Z0-9][a-zA-Z0-9\-]*\.myshopify\.com/`)
	match := reg.MatchString(shop)
	if match {
		if nonce == "" || validNonces[nonce] == true {
			// TODO: verify HMAC
			// 1) generate hmac object with key being the secret key from config.App
			// 2) prepare the query string, remove the hmac from the fyll URL, remove the http://host? from the URL
			// 3) generate HMAC digest of query string
			// 4) compare to HMAC received. if equal, pass.
			return true
		}

	}

	return false
}

func RequestAccessToken(shopName, clientId, clientSecret, code string) (data.ShopifyAccessTokenResponse, error) {
	sat := data.ShopifyAccessTokenResponse{}
	jsonDat := map[string]string{"client_id": clientId, "client_secret": clientSecret, "code": code}
	cereal, _ := json.Marshal(jsonDat)

	r, _ := http.NewRequest("POST", fmt.Sprintf("%vadmin/oauth/access_token", shopName), bytes.NewBuffer(cereal))
	res, err := Client.Do(r)
	if err != nil {
		return sat, err
	}

	if res.StatusCode == 200 {
		body, _ := ioutil.ReadAll(res.Body)
		err = json.Unmarshal(body, &sat)
		if err != nil {
			return sat, err
		}

		return sat, nil
	} else {
		return sat, errors.New("shopify returned non-200 HTTP code")
	}

}

func AuthorizeNonce(nonce string) {
	validNonces[nonce] = true
}
