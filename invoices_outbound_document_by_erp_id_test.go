package inexchange_test

import (
	"context"
	"encoding/json"
	"log"
	"testing"

	_ "github.com/joho/godotenv/autoload"
)

func TestInvoicesOutboundDocumentByERPID(t *testing.T) {
	client := client()

	req := client.NewInvoicesOutboundDocumentByERPIDRequest()
	req.PathParams().DocumentERPID = "6036bd29-0a7c-49ac-9ed2-b3a200dfabbe"
	req.RequestBody().IncludeFileInfo = true

	resp, err := req.Do(context.Background())
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
