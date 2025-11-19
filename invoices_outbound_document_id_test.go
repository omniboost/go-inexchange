package inexchange_test

import (
	"context"
	"encoding/json"
	"log"
	"testing"

	_ "github.com/joho/godotenv/autoload"
)

func TestInvoicesOutboundDocumentID(t *testing.T) {
	client := client()

	req := client.NewInvoicesOutboundDocumentIDRequest()
	req.PathParams().DocumentID = "74c54e26-c985-43c6-b284-79db55d93438"
	req.RequestBody().IncludeFileInfo = true

	resp, err := req.Do(context.Background())
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
