package inexchange_test

import (
	"context"
	"encoding/json"
	"log"
	"testing"

	_ "github.com/joho/godotenv/autoload"
	"github.com/omniboost/go-inexchange"
)

func TestGetV0IndividualProfileByID(t *testing.T) {
	client := client()

	req := client.NewGetV0IndividualProfileByIDRequest()
	req.Headers().PropertyID = "f142d311-7fd8-46c8-afbe-6fd50c03244d"
	req.PathParams().ID = "066c2ecd-8ea5-4be2-a601-a8f894a66c21"
	req.QueryParams().Extend = inexchange.CommaSeparatedQueryParam{"LocalExternalAccountReceivable"}

	resp, err := req.Do(context.Background())
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
