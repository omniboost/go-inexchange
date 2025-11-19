package inexchange_test

import (
	"context"
	"encoding/json"
	"log"
	"testing"

	_ "github.com/joho/godotenv/autoload"
	"github.com/omniboost/go-inexchange"
)

func TestGetV0TravelAgentProfileByID(t *testing.T) {
	client := client()

	req := client.NewGetV0TravelAgentProfileByIDRequest()
	req.Headers().PropertyID = "f142d311-7fd8-46c8-afbe-6fd50c03244d"
	req.PathParams().ID = "d65615a6-1120-4f5e-876f-c7c5c2340c91"
	req.QueryParams().Extend = inexchange.CommaSeparatedQueryParam{"LocalExternalAccountReceivable"}

	resp, err := req.Do(context.Background())
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
