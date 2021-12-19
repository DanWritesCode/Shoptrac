package shopify

import (
	"net/http"
	"time"
)

var Client *http.Client

func SetupClient() {
	Client = &http.Client{
		Timeout: time.Second * 20, // time.Second * <sec for req timeout>
	}
}
