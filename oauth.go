package inexchange

import (
	"net/url"
	"strings"
	"time"

	"github.com/joefitzgerald/passwordcredentials"
	"golang.org/x/oauth2"
)

const (
	scope                = ""
	oauthStateString     = ""
	authorizationTimeout = 60 * time.Second
	tokenTimeout         = 5 * time.Second
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
				AuthURL:  "https://identity.apaleo.com/connect/authorize",
				TokenURL: "https://identity.apaleo.com/connect/token",
			},
		},
	}

	config.SetBaseURL(&BaseURL)
	return config
}

func (c *Oauth2Config) SetBaseURL(baseURL *url.URL) {
	// Strip trailing slash
	baseURL.Path = strings.TrimSuffix(baseURL.Path, "/")

	c.Endpoint = oauth2.Endpoint{
		AuthURL:  baseURL.String() + "/oauth",
		TokenURL: baseURL.String() + "/access_token",
	}
}

type Oauth2PasswordConfig struct {
	passwordcredentials.Config
}

func NewOauth2PasswordConfig() *Oauth2PasswordConfig {
	config := &Oauth2PasswordConfig{
		Config: passwordcredentials.Config{
			ClientID:     "",
			ClientSecret: "",
			Username:     "",
			Password:     "",
			Scopes:       []string{scope},
			Endpoint: oauth2.Endpoint{
				AuthURL:  "https://identity.apaleo.com/connect/authorize",
				TokenURL: "https://identity.apaleo.com/connect/token",
			},
		},
	}

	config.SetBaseURL(&BaseURL)
	return config
}

func (c *Oauth2PasswordConfig) SetBaseURL(baseURL *url.URL) {
	// Strip trailing slash
	baseURL.Path = strings.TrimSuffix(baseURL.Path, "/")

	c.Endpoint = oauth2.Endpoint{
		AuthURL:  baseURL.String() + "/oauth",
		TokenURL: baseURL.String() + "/access_token",
	}
}
