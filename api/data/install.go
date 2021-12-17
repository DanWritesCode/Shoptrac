package data

type InstallData struct {
	URL      string `json:"url,omitempty"`
	AuthCode string `json:"authCode,omitempty"`
	HMAC     string `json:"hmac,omitempty"`
	ShopName string `json:"shopName,omitempty"`
	HostName string `json:"hostName,omitempty"`
	Nonce    string `json:"nonce,omitempty"`
}
