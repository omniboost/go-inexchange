package inexchange_test

import (
	"context"
	"encoding/json"
	"log"
	"testing"

	_ "github.com/joho/godotenv/autoload"
)

func TestGetCashieringFinalInvoiceContentByFolioID(t *testing.T) {
	client := client()

	req := client.NewGetCashieringFinalInvoiceContentByFolioIDRequest()
	req.Headers().PropertyID = "806c52a0-753d-4a03-a4ad-0d5b577a11dc"
	req.PathParams().AccountID = "4188cb66-acd9-4651-980c-c1d4abcb1624"
	req.PathParams().FolioID = "e65f65f9-a13e-40d0-b5aa-1e1aac2bd6a7"

	resp, err := req.Do(context.Background())
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
