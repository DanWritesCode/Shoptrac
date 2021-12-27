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

func CheckShopDomain(shop string) bool {
	reg := regexp.MustCompile(`\A[a-zA-Z0-9][a-zA-Z0-9\-]*\.myshopify\.com`)
	return reg.MatchString(shop)
}

func RequestAccessToken(shopName, clientId, clientSecret, code string) (data.ShopifyAccessTokenResponse, error) {
	sat := data.ShopifyAccessTokenResponse{}
	jsonDat := map[string]string{"client_id": clientId, "client_secret": clientSecret, "code": code}
	cereal, _ := json.Marshal(jsonDat)

	r, _ := http.NewRequest("POST", fmt.Sprintf("https://%v/admin/oauth/access_token", shopName), bytes.NewBuffer(cereal))
	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("User-Agent", "StonksUp App")

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

func CheckNonce(nonce string) bool {
	return validNonces[nonce] == true
}

func AuthorizeNonce(nonce string) {
	validNonces[nonce] = true
}
