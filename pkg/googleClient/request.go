package googleClient

type ExchangeCodeRequest struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	RedirectURL  string `json:"redirect_uri"`
	Code         string `json:"code"`
	GrantType    string `json:"grant_type"`
}
