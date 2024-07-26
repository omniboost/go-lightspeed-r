package lightspeed_r

import (
	"golang.org/x/oauth2"
)

const (
	scope = ""
)

type Oauth2Config struct {
	oauth2.Config
}

func NewOauth2Config() *Oauth2Config {
	config := &Oauth2Config{
		Config: oauth2.Config{
			RedirectURL:  "",
			ClientID:     "",
			ClientSecret: "",
			Scopes:       []string{scope},
			Endpoint: oauth2.Endpoint{
				TokenURL:  "https://oauth-proxy.omniboost.io/ls_retail/oauth2/token",
				AuthStyle: oauth2.AuthStyleAutoDetect,
			},
		},
	}

	return config
}
