package inexchange_test

import (
	"context"
	"encoding/json"
	"log"
	"testing"

	_ "github.com/joho/godotenv/autoload"
)

func TestPatchCompanyProfileByID(t *testing.T) {
	client := client()

	req := client.NewPatchCompanyProfileByIDRequest()
	req.Headers().PropertyID = "806c52a0-753d-4a03-a4ad-0d5b577a11dc"
	req.Headers().IfMatch = "59"
	req.PathParams().ID = "3bed7328-9cbc-4f7b-b07b-3088db1c7dd3"
	body := req.NewRequestBody()
	body.Details.FullName = "Emma"

	req.SetRequestBody(body)
	resp, err := req.Do(context.Background())
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
