package inexchange_test

import (
	"context"
	"encoding/json"
	"log"
	"testing"
	"time"

	_ "github.com/joho/godotenv/autoload"
	"github.com/omniboost/go-inexchange"
)

func TestDocumentsOutboundList(t *testing.T) {
	client := client()

	req := client.NewDocumentsOutboundListRequest()
	req.RequestBody().Take = 10
	req.RequestBody().CreatedFrom = inexchange.DateTime{Time: time.Date(2025, 8, 24, 14, 15, 22, 0, time.UTC)}
	req.RequestBody().CreatedTo = inexchange.DateTime{Time: time.Date(2025, 8, 30, 14, 15, 22, 0, time.UTC)}
	req.RequestBody().IncludeFileInfo = true
	req.RequestBody().IncludeErrorInfo = true

	resp, err := req.Do(context.Background())
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
