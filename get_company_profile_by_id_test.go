package inexchange_test

import (
	"context"
	"encoding/json"
	"log"
	"testing"

	_ "github.com/joho/godotenv/autoload"
)

func TestGetCompanyProfileByID(t *testing.T) {
	client := client()

	req := client.NewGetCompanyProfileByIDRequest()
	req.Headers().PropertyID = "f142d311-7fd8-46c8-afbe-6fd50c03244d"
	req.PathParams().ID = "e5bd9019-3c49-42cd-8cc2-15d9e3ccf176"
	req.QueryParams().Extend = []string{"LocalExternalAccountReceivable"}

	resp, err := req.Do(context.Background())
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
