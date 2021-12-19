package data

type InstallData struct {
	URL      string `json:"url,omitempty"`
	AuthCode string `json:"authCode,omitempty"`
	HMAC     string `json:"hmac,omitempty"`
	Shop     string `json:"shop,omitempty"`
	Host     string `json:"host,omitempty"`
	Nonce    string `json:"nonce,omitempty"`
}

type ShopifyAccessTokenResponse struct {
	AccessToken string `json:"access_token,omitempty"`
	Scope       string `json:"scope,omitempty"`
}
