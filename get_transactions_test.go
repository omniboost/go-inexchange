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

func TestGetTransactions(t *testing.T) {
	client := client()

	req := client.NewGetTransactionsRequest()
	req.Headers().PropertyID = "c6a95ecb-be35-4619-875b-cae5d0df5fd2"

	resp, err := req.Do(context.Background())
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}

func TestGetTransactionsAll(t *testing.T) {
	client := client()

	req := client.NewGetTransactionsRequest()
	req.Headers().PropertyID = "50abb838-0a4e-4bea-afb4-28523f8f8daa"
	req.QueryParams().DateType = "BusinessDate"
	req.QueryParams().Extend = inexchange.CommaSeparatedQueryParam{"TaxTransactions", "ReservationDetails"}
	req.QueryParams().LedgerTypes = inexchange.CommaSeparatedQueryParam{"Guest", "Deposit"}
	req.QueryParams().StartDate = inexchange.Date{time.Date(
		2024, 11, 1, 0, 0, 0, 0, time.UTC,
	)}

	req.QueryParams().EndDate = inexchange.Date{time.Date(
		2024, 11, 2, 0, 0, 0, 0, time.UTC,
	)}

	resp, err := req.All(context.Background())
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
