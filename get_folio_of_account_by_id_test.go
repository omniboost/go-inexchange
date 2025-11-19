package inexchange_test

import (
	"context"
	"encoding/json"
	"log"
	"testing"

	_ "github.com/joho/godotenv/autoload"
)

func TestGetFolioOfAccountByID(t *testing.T) {
	client := client()

	req := client.NewGetFolioOfAccountByIDRequest()
	req.Headers().PropertyID = "806c52a0-753d-4a03-a4ad-0d5b577a11dc"
	req.PathParams().AccountID = "167b3e65-145d-429d-a158-d5d3e731d0b5"
	req.PathParams().FolioID = "85f26777-53f6-49b6-89ce-56b3a4523e26"

	resp, err := req.Do(context.Background())
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
