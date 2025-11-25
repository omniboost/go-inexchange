package inexchange

import (
	"context"
	"net/http"
	"net/url"

	"github.com/omniboost/go-inexchange/utils"
)

func (c *Client) NewBuyerPartiesLookupRequest() BuyerPartiesLookupRequest {
	return BuyerPartiesLookupRequest{
		client:      c,
		queryParams: c.NewBuyerPartiesLookupQueryParams(),
		pathParams:  c.NewBuyerPartiesLookupPathParams(),
		method:      http.MethodPost,
		headers:     c.NewBuyerPartiesLookupHeaders(),
		requestBody: c.NewBuyerPartiesLookupRequestBody(),
	}
}

type BuyerPartiesLookupRequest struct {
	client      *Client
	queryParams *BuyerPartiesLookupQueryParams
	pathParams  *BuyerPartiesLookupPathParams
	method      string
	headers     *BuyerPartiesLookupHeaders
	requestBody BuyerPartiesLookupRequestBody
}

func (c *Client) NewBuyerPartiesLookupQueryParams() *BuyerPartiesLookupQueryParams {
	return &BuyerPartiesLookupQueryParams{}
}

type BuyerPartiesLookupQueryParams struct {
	StatusCode string `schema:"statusCode,omitempty"`
	ParentID   string `schema:"parentId,omitempty"`
	Recursive  bool   `schema:"recursive,omitempty"`
	Sort       string `schema:"sort,omitempty"`
	PageNumber int32  `schema:"pageNumber,omitempty"`
	PageSize   int32  `schema:"pageSize,omitempty"`
}

func (p BuyerPartiesLookupQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	encoder.RegisterEncoder(Date{}, utils.EncodeSchemaMarshaler)
	encoder.RegisterEncoder(DateTime{}, utils.EncodeSchemaMarshaler)
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *BuyerPartiesLookupRequest) QueryParams() *BuyerPartiesLookupQueryParams {
	return r.queryParams
}

func (c *Client) NewBuyerPartiesLookupHeaders() *BuyerPartiesLookupHeaders {
	return &BuyerPartiesLookupHeaders{}
}

type BuyerPartiesLookupHeaders struct {
	TenantID string `schema:"AC-Tenant-ID,omitempty"`
}

func (r *BuyerPartiesLookupRequest) Headers() *BuyerPartiesLookupHeaders {
	return r.headers
}

func (c *Client) NewBuyerPartiesLookupPathParams() *BuyerPartiesLookupPathParams {
	return &BuyerPartiesLookupPathParams{}
}

type BuyerPartiesLookupPathParams struct {
}

func (p *BuyerPartiesLookupPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *BuyerPartiesLookupRequest) PathParams() *BuyerPartiesLookupPathParams {
	return r.pathParams
}

func (r *BuyerPartiesLookupRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *BuyerPartiesLookupRequest) SetMethod(method string) {
	r.method = method
}

func (r *BuyerPartiesLookupRequest) Method() string {
	return r.method
}

func (s *Client) NewBuyerPartiesLookupRequestBody() BuyerPartiesLookupRequestBody {
	return BuyerPartiesLookupRequestBody{}
}

type BuyerPartiesLookupRequestBody struct {
	// Customer number
	PartyID string `json:"PartyId,omitempty"`
	// Company name
	Name string `json:"name,omitempty"`
	// Department name
	Department string `json:"department,omitempty"`
	// Street name
	StreetName string `json:"streetName,omitempty"`
	// Post box
	PostBox string `json:"postBox,omitempty"`
	// Postal zone
	PostalZone string `json:"postalZone,omitempty"`
	// City
	City string `json:"city,omitempty"`
	// Country code, eg. SE, DK, FI etc
	CountryCode string `json:"countryCode,omitempty"`
	// Phone number
	PhoneNo string `json:"phoneNo,omitempty"`
	// Fax number
	FaxNo string `json:"faxNo,omitempty"`
	// Email
	Email string `json:"email,omitempty"`
	// GLN
	GLN string `json:"gln,omitempty"`
	// Org number
	OrgNo string `json:"orgNo,omitempty"`
	// Vat number
	VATNo string `json:"vatNo,omitempty"`
	// Peppol participant identifier, schema:idvalue
	PeppolParticipantIdentifier string `json:"peppolParticipantIdentifier,omitempty"`
	// Postal code. Use PostalZone instead. Will only be used if PostalZone isn’t set.
	PostalCode string `json:"postalCode,omitempty"`
}

func (r *BuyerPartiesLookupRequest) RequestBody() *BuyerPartiesLookupRequestBody {
	return &r.requestBody
}

func (r *BuyerPartiesLookupRequest) RequestBodyInterface() any {
	return r.requestBody
}

func (r *BuyerPartiesLookupRequest) SetRequestBody(body BuyerPartiesLookupRequestBody) {
	r.requestBody = body
}

func (r *BuyerPartiesLookupRequest) NewResponseBody() *BuyerPartiesLookupResponseBody {
	return &BuyerPartiesLookupResponseBody{}
}

type BuyerPartiesLookupResponseBody struct {
	Parties          Parties `json:"parties"`
	PotentialMatches int     `json:"potentialMatches"`
	ExactHit         bool    `json:"exactHit"`
}

func (r *BuyerPartiesLookupRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("/buyerparties/lookup", r.PathParams())
	return &u
}

func (r *BuyerPartiesLookupRequest) Do(ctx context.Context) (BuyerPartiesLookupResponseBody, error) {
	// Create http request
	req, err := r.client.NewRequest(ctx, r)
	if err != nil {
		return *r.NewResponseBody(), err
	}

	// Process query parameters
	err = utils.AddQueryParamsToRequest(r.QueryParams(), req, false)
	if err != nil {
		return *r.NewResponseBody(), err
	}

	responseBody := r.NewResponseBody()
	_, err = r.client.Do(req, responseBody)
	return *responseBody, err
}
