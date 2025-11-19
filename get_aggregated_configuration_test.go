package inexchange_test

import (
	"context"
	"encoding/json"
	"log"
	"testing"

	_ "github.com/joho/godotenv/autoload"
)

func TestGetAggregatedConfiguration(t *testing.T) {
	client := client()

	req := client.NewGetAggregatedConfigurationRequest()
	req.Headers().PropertyID = "c6a95ecb-be35-4619-875b-cae5d0df5fd2"
	req.QueryParams().Types = []string{"TaxRule"}

	resp, err := req.Do(context.Background())
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}

func TestGetAggregatedConfigurationAll(t *testing.T) {
	client := client()

	req := client.NewGetAggregatedConfigurationRequest()
	req.Headers().PropertyID = "c6a95ecb-be35-4619-875b-cae5d0df5fd2"
	req.QueryParams().Types = []string{"TaxRule", "TransactionSubGroup"}

	resp, err := req.All(context.Background())
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
