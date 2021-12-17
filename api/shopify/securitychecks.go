package shopify

import (
	"regexp"
)

func SecurityCheck(fullURL string, nonce string, hmac string, shop string) bool {
	match, err := regexp.MatchString("/\\A(https|http)\\:\\/\\/[a-zA-Z0-9][a-zA-Z0-9\\-]*\\.myshopify\\.com\\//\n", shop)
	if match && err == nil {
		/*if nonce == "" {
		  if hmac2.Equal([]byte(fullURL), []byte(hmac)) {
		    return true
		  }
		}*/
		return true
	}

	return false
}
