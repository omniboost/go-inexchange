package inexchange_test

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"os"

	"github.com/omniboost/go-inexchange"
)

func client() *inexchange.Client {
	clientID := os.Getenv("OAUTH_CLIENT_ID")
	clientSecret := os.Getenv("OAUTH_CLIENT_SECRET")
	tokenURL := os.Getenv("OAUTH_TOKEN_URL")
	inexchangeBaseURL := os.Getenv("INEXCHANGE_BASE_URL")
	inexchangeTenant := os.Getenv("INEXCHANGE_TENANT")
	inexchangeUsername := os.Getenv("INEXCHANGE_USERNAME")
	inexchangePassword := os.Getenv("INEXCHANGE_PASSWORD")

	oauthConfig := inexchange.NewOauth2PasswordConfig()
	oauthConfig.ClientID = clientID
	oauthConfig.ClientSecret = clientSecret
	oauthConfig.Username = fmt.Sprintf("%s@%s", inexchangeUsername, inexchangeTenant)
	oauthConfig.Password = inexchangePassword

	// set alternative token url
	if tokenURL != "" {
		oauthConfig.Endpoint.TokenURL = tokenURL
	}

	// get http client with automatic oauth logic
	httpClient := oauthConfig.Client(context.Background())

	client := inexchange.NewClient(httpClient)
	client.SetDebug(true)
	client.SetDisallowUnknownFields(true)

	if inexchangeBaseURL != "" {
		u, err := url.Parse(inexchangeBaseURL)
		if err != nil {
			log.Fatal(err)
		}
		client.SetBaseURL(*u)
	}
	return client
}
