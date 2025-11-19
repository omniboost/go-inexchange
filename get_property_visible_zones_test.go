package inexchange_test

import (
	"context"
	"encoding/json"
	"log"
	"testing"

	_ "github.com/joho/godotenv/autoload"
)

func TestGetPropertyVisibleZones(t *testing.T) {
	client := client()

	req := client.NewGetPropertyVisibleZonesRequest()
	req.PathParams().ID = "806c52a0-753d-4a03-a4ad-0d5b577a11dc"

	resp, err := req.Do(context.Background())
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
