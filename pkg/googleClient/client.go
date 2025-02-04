package googleClient

import (
	"encoding/json"
	"fmt"

	"github.com/parnurzeal/gorequest"
)

const (
	tokenURL = "https://oauth2.googleapis.com/token"
)

type GoogleClient interface {
	ExchangeCode(code string) (*ExchangeCodeResponse, error)
	RefreshToken(refreshToken string) (*ExchangeCodeResponse, error)
}

type googleClient struct {
	gorequest    *gorequest.SuperAgent
	clientID     string
	clientSecret string
	redirectURL  string
}

func NewGoogleClient(clientID, clientSecret, redirectURL string) GoogleClient {
	return &googleClient{
		gorequest:    gorequest.New(),
		clientID:     clientID,
		clientSecret: clientSecret,
		redirectURL:  redirectURL,
	}
}

func (c *googleClient) ExchangeCode(code string) (*ExchangeCodeResponse, error) {
	req := ExchangeCodeRequest{
		ClientID:     c.clientID,
		ClientSecret: c.clientSecret,
		RedirectURL:  c.redirectURL,
		Code:         code,
		GrantType:    "authorization_code",
	}
	res, body, errList := c.gorequest.Post(tokenURL).Send(req).End()
	if errList != nil {
		return nil, errList[0]
	}

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("google oauth api failed to exchange code: %d", res.StatusCode)
	}

	var response ExchangeCodeResponse
	err := json.Unmarshal([]byte(body), &response)
	return &response, err
}

func (c *googleClient) RefreshToken(refreshToken string) (*ExchangeCodeResponse, error) {
	req := RefreshTokenRequest{
		ClientID:     c.clientID,
		ClientSecret: c.clientSecret,
		RefreshToken: refreshToken,
		GrantType:    "refresh_token",
	}

	res, body, errList := c.gorequest.Post(tokenURL).Send(req).End()
	if errList != nil {
		return nil, errList[0]
	}

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("google oauth api failed to refresh token: %d, body: %s", res.StatusCode, body)
	}

	var response ExchangeCodeResponse
	err := json.Unmarshal([]byte(body), &response)
	return &response, err
}
