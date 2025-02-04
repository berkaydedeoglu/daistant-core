package response

import "fmt"

type GoogleOAuthURLResponse struct {
	BaseURL      string `json:"base_url"`
	Scope        string `json:"-"`
	ClientID     string `json:"-"`
	RedirectURI  string `json:"-"`
	ResponseType string `json:"-"`
	AccessType   string `json:"-"`
	State        string `json:"-"`
}

func (r *GoogleOAuthURLResponse) SetOAuthURL(userId uint) *GoogleOAuthURLResponse {
	r.State = fmt.Sprintf("%d", userId)
	url := fmt.Sprintf("%s?access_type=%s&scope=%s&client_id=%s&redirect_uri=%s&response_type=%s&state=%s",
		r.BaseURL,
		r.AccessType,
		r.Scope,
		r.ClientID,
		r.RedirectURI,
		r.ResponseType,
		r.State,
	)

	r.BaseURL = url
	return r
}
