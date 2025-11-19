package inexchange_test

import (
	"context"
	"encoding/json"
	"log"
	"testing"

	_ "github.com/joho/godotenv/autoload"
)

func TestGetTravelAgentProfileByID(t *testing.T) {
	client := client()

	req := client.NewGetTravelAgentProfileByIDRequest()
	req.Headers().PropertyID = "cb5774da-e261-4421-90f6-070525fdb007"
	req.PathParams().ID = "28d2ed1f-5dcf-4783-b4da-3dbebad08d3e"
	// req.QueryParams().Extend = []string{"LocalExternalAccountReceivable"}

	resp, err := req.Do(context.Background())
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
