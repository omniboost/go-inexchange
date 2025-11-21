package inexchange_test

import (
	"context"
	"encoding/json"
	"log"
	"testing"

	_ "github.com/joho/godotenv/autoload"
)

func TestDocuments(t *testing.T) {
	client := client()

	req := client.NewDocumentsRequest()
	req.RequestBody().File = "UEsDBBQACAgIAK6ZkVIAAAAAAAAAAAAAAAAJABwAZmlsZTEudHh0VVQJAAPO1lFfztZRXXgLAAEEAAAAAAQAAAAAAABQSwMEFAACAAgAqpmRUgAAAAAAAAAAAAAAAAgAGABoAZmlsZTIudHh0VVQJAAPO1lFfztZRXXgLAAEEAAAAAAQAAAAAAABQSwECHgMUAAICAgArpmRUgAAAAAAAAAAAAAAAkAGAAAAAAAAAAAApIEAAAAAZmlsZTEudHh0VVQFAAPO1lFfXgLAAEEAAAAAAQAAAAAAABQSwECHgMUAAICAgAqpmRUgAAAAAAAAAAAAAAAgAGAAAAAAAAAAAApIEAAAAAZmlsZTIudHh0VVQFAAPO1lFfXgLAAEEAAAAAAQAAAAAAABQSwUGAAAAAAIAAgC9AAAAUQAAAAAA"

	resp, err := req.Do(context.Background())
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}

