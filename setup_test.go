package inexchange_test

import (
	"log"
	"net/url"
	"os"

	"github.com/omniboost/go-inexchange"
)

func client() *inexchange.Client {
	inexchangeBaseURL := os.Getenv("INEXCHANGE_BASE_URL")
	inexchangeClientToken := os.Getenv("INEXCHANGE_CLIENT_TOKEN")

	client := inexchange.NewClient(nil)
	client.SetClientToken(inexchangeClientToken)
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
