package inexchange_test

import (
	"context"
	"encoding/json"
	"log"
	"testing"

	_ "github.com/joho/godotenv/autoload"
)

func TestBuyerPartiesLookup(t *testing.T) {
	client := client()

	req := client.NewBuyerPartiesLookupRequest()
	req.RequestBody().PartyID = "1"
	req.RequestBody().Name = "Test"

	resp, err := req.Do(context.Background())
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}

