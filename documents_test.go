package inexchange_test

import (
	"context"
	"encoding/json"
	"log"
	"strings"
	"testing"

	_ "github.com/joho/godotenv/autoload"
	"github.com/omniboost/go-inexchange"
)

func TestDocuments(t *testing.T) {
	client := client()

	req := client.NewDocumentsRequest()
	req.FormParams().File = inexchange.FormFile{
		Filename: "test.csv",
		Content:  strings.NewReader("asdafadaf"),
	}

	resp, err := req.Do(context.Background())
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}

