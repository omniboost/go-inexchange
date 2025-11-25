package inexchange

import (
	"context"
	"net/http"
	"net/url"

	"github.com/omniboost/go-inexchange/utils"
)

func (c *Client) NewCompaniesRegisterRequest() CompaniesRegisterRequest {
	return CompaniesRegisterRequest{
		client:      c,
		queryParams: c.NewCompaniesRegisterQueryParams(),
		pathParams:  c.NewCompaniesRegisterPathParams(),
		method:      http.MethodPost,
		headers:     c.NewCompaniesRegisterHeaders(),
		requestBody: c.NewCompaniesRegisterRequestBody(),
	}
}

type CompaniesRegisterRequest struct {
	client      *Client
	queryParams *CompaniesRegisterQueryParams
	pathParams  *CompaniesRegisterPathParams
	method      string
	headers     *CompaniesRegisterHeaders
	requestBody CompaniesRegisterRequestBody
}

func (c *Client) NewCompaniesRegisterQueryParams() *CompaniesRegisterQueryParams {
	return &CompaniesRegisterQueryParams{}
}

type CompaniesRegisterQueryParams struct {
}

func (p CompaniesRegisterQueryParams) ToURLValues() (url.Values, error) {
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

func (r *CompaniesRegisterRequest) QueryParams() *CompaniesRegisterQueryParams {
	return r.queryParams
}

func (c *Client) NewCompaniesRegisterHeaders() *CompaniesRegisterHeaders {
	return &CompaniesRegisterHeaders{}
}

type CompaniesRegisterHeaders struct {
	TenantID string `schema:"AC-Tenant-ID,omitempty"`
}

func (r *CompaniesRegisterRequest) Headers() *CompaniesRegisterHeaders {
	return r.headers
}

func (c *Client) NewCompaniesRegisterPathParams() *CompaniesRegisterPathParams {
	return &CompaniesRegisterPathParams{}
}

type CompaniesRegisterPathParams struct{}

func (p *CompaniesRegisterPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *CompaniesRegisterRequest) PathParams() *CompaniesRegisterPathParams {
	return r.pathParams
}

func (r *CompaniesRegisterRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *CompaniesRegisterRequest) SetMethod(method string) {
	r.method = method
}

func (r *CompaniesRegisterRequest) Method() string {
	return r.method
}

func (s *Client) NewCompaniesRegisterRequestBody() CompaniesRegisterRequestBody {
	return CompaniesRegisterRequestBody{}
}

type CompaniesRegisterRequestBody struct {
	ErpID           string   `json:"ErpId"`
	ErpProduct      string   `json:"ErpProduct"`
	Name            string   `json:"Name"`
	Email           string   `json:"Email"`
	CountryCode     string   `json:"CountryCode"`
	AltName         string   `json:"AltName"`
	StreetName      string   `json:"StreetName"`
	Department      string   `json:"Department"`
	PostBox         string   `json:"PostBox"`
	City            string   `json:"City"`
	PostalZone      string   `json:"PostalZone"`
	PostalCode      string   `json:"PostalCode"`
	LanguageCode    string   `json:"LanguageCode"`
	PhoneNo         string   `json:"PhoneNo"`
	FaxNo           string   `json:"FaxNo"`
	OrgNo           string   `json:"OrgNo"`
	VAtNo           string   `json:"VatNo"`
	GLN             string   `json:"GLN"`
	IsVatRegistered bool     `json:"IsVatRegistered"`
	Processes       []string `json:"Processes"`
}

func (r *CompaniesRegisterRequest) RequestBody() *CompaniesRegisterRequestBody {
	return &r.requestBody
}

func (r *CompaniesRegisterRequest) RequestBodyInterface() any {
	return r.requestBody
}

func (r *CompaniesRegisterRequest) SetRequestBody(body CompaniesRegisterRequestBody) {
	r.requestBody = body
}

func (r *CompaniesRegisterRequest) NewResponseBody() *CompaniesRegisterResponseBody {
	return &CompaniesRegisterResponseBody{}
}

type CompaniesRegisterResponseBody struct{}

func (r *CompaniesRegisterRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("/companies/register", r.PathParams())
	return &u
}

func (r *CompaniesRegisterRequest) Do(ctx context.Context) (CompaniesRegisterResponseBody, error) {
	// Create http request
	req, err := r.client.NewRequest(ctx, r)
	if err != nil {
		return *r.NewResponseBody(), err
	}

	req.Header.Del("ClientToken") // Not needed for this endpoint

	// Process query parameters
	err = utils.AddQueryParamsToRequest(r.QueryParams(), req, false)
	if err != nil {
		return *r.NewResponseBody(), err
	}

	responseBody := r.NewResponseBody()
	_, err = r.client.Do(req, responseBody)
	return *responseBody, err
}
