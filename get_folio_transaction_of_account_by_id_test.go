package inexchange_test

import (
	"context"
	"encoding/json"
	"log"
	"testing"

	_ "github.com/joho/godotenv/autoload"
)

func TestGetFolioTransactionOfAccountByID(t *testing.T) {
	client := client()

	req := client.NewGetFolioTransactionOfAccountByIDRequest()
	req.Headers().PropertyID = "50abb838-0a4e-4bea-afb4-28523f8f8daa"
	req.PathParams().AccountID = "3519c26e-0b0d-4020-b9b3-20acaa482763"
	req.PathParams().FolioID = "dcd6ace7-dce6-4aeb-a7b8-3660f7b9753e"
	req.PathParams().TransactionID = "cef19de0-d295-44df-88fd-df4d638a7872"

	resp, err := req.Do(context.Background())
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
