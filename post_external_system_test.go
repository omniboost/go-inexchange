package inexchange_test

import (
	"context"
	"encoding/json"
	"log"
	"testing"

	_ "github.com/joho/godotenv/autoload"
)

func TestPostExternalSystem(t *testing.T) {
	client := client()

	req := client.NewPostExternalSystemRequest()

	body := req.NewRequestBody()
	body.ZoneID = "6ba2da40-060c-4b64-a153-4113caf7bcf7"
	body.Code = "OBEXACTONLINE"
	body.IsActive = true
	body.ForceToUse = false
	body.Name = []struct {
		Content      string `json:"content,omitempty"`
		LanguageCode string `json:"languageCode,omitempty"`
	}{
		{
			Content:      "OBEXACTONLINE",
			LanguageCode: "EN",
		},
	}

	body.Description = []struct {
		Content      string `json:"content,omitempty"`
		LanguageCode string `json:"languageCode,omitempty"`
	}{
		{
			Content:      "OBEXACTONLINE",
			LanguageCode: "EN",
		},
	}

	req.SetRequestBody(body)
	resp, err := req.Do(context.Background())
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
