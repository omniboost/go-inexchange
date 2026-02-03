package inexchange

import (
	"context"
	"net/http"
	"net/url"

	"github.com/omniboost/go-inexchange/utils"
)

func (c *Client) NewNetworkSetupRequestRequest() NetworkSetupRequestRequest {
	return NetworkSetupRequestRequest{
		client:      c,
		queryParams: c.NewNetworkSetupRequestQueryParams(),
		pathParams:  c.NewNetworkSetupRequestPathParams(),
		method:      http.MethodPost,
		headers:     c.NewNetworkSetupRequestHeaders(),
		requestBody: c.NewNetworkSetupRequestRequestBody(),
	}
}

type NetworkSetupRequestRequest struct {
	client      *Client
	queryParams *NetworkSetupRequestQueryParams
	pathParams  *NetworkSetupRequestPathParams
	method      string
	headers     *NetworkSetupRequestHeaders
	requestBody NetworkSetupRequestRequestBody
}

func (c *Client) NewNetworkSetupRequestQueryParams() *NetworkSetupRequestQueryParams {
	return &NetworkSetupRequestQueryParams{}
}

type NetworkSetupRequestQueryParams struct {
}

func (p NetworkSetupRequestQueryParams) ToURLValues() (url.Values, error) {
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

func (r *NetworkSetupRequestRequest) QueryParams() *NetworkSetupRequestQueryParams {
	return r.queryParams
}

func (c *Client) NewNetworkSetupRequestHeaders() *NetworkSetupRequestHeaders {
	return &NetworkSetupRequestHeaders{}
}

type NetworkSetupRequestHeaders struct {
	TenantID string `schema:"AC-Tenant-ID,omitempty"`
}

func (r *NetworkSetupRequestRequest) Headers() *NetworkSetupRequestHeaders {
	return r.headers
}

func (c *Client) NewNetworkSetupRequestPathParams() *NetworkSetupRequestPathParams {
	return &NetworkSetupRequestPathParams{}
}

type NetworkSetupRequestPathParams struct{}

func (p *NetworkSetupRequestPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *NetworkSetupRequestRequest) PathParams() *NetworkSetupRequestPathParams {
	return r.pathParams
}

func (r *NetworkSetupRequestRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *NetworkSetupRequestRequest) SetMethod(method string) {
	r.method = method
}

func (r *NetworkSetupRequestRequest) Method() string {
	return r.method
}

func (s *Client) NewNetworkSetupRequestRequestBody() NetworkSetupRequestRequestBody {
	return NetworkSetupRequestRequestBody{}
}

type NetworkSetupRequestRequestBody struct {
	Operator struct {
		Name         string `json:"Name,omitempty"`
		ContactName  string `json:"ContactName,omitempty"`
		ContactPhone string `json:"ContactPhone,omitempty"`
		ContactEmail string `json:"ContactEmail,omitempty"`
	} `json:"Operator,omitempty"`

	Name                      string   `json:"Name,omitempty"`
	Department                string   `json:"Department,omitempty"`
	StreetName                string   `json:"StreetName,omitempty"`
	PostBox                   string   `json:"PostBox,omitempty"`
	City                      string   `json:"City,omitempty"`
	PostalZone                string   `json:"PostalZone,omitempty"`
	CountryCode               string   `json:"CountryCode,omitempty"`
	PhoneNo                   string   `json:"PhoneNo,omitempty"`
	FaxNo                     string   `json:"FaxNo,omitempty"`
	Email                     string   `json:"Email,omitempty"`
	GLN                       string   `json:"GLN,omitempty"`
	OrgNo                     string   `json:"OrgNo,omitempty"`
	VatNo                     string   `json:"VatNo,omitempty"`
	AttachmentUri             string   `json:"AttachmentUri,omitempty"`
	Processes                 []string `json:"Processes,omitempty"`
	NetworkSetupFeedbackEmail string   `json:"NetworkSetupFeedbackEmail,omitempty"`
}

func (r *NetworkSetupRequestRequest) RequestBody() *NetworkSetupRequestRequestBody {
	return &r.requestBody
}

func (r *NetworkSetupRequestRequest) RequestBodyInterface() any {
	return r.requestBody
}

func (r *NetworkSetupRequestRequest) SetRequestBody(body NetworkSetupRequestRequestBody) {
	r.requestBody = body
}

func (r *NetworkSetupRequestRequest) NewResponseBody() *NetworkSetupRequestResponseBody {
	return &NetworkSetupRequestResponseBody{}
}

type NetworkSetupRequestResponseBody struct{}

func (r *NetworkSetupRequestRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("/network/setup", r.PathParams())
	return &u
}

func (r *NetworkSetupRequestRequest) Do(ctx context.Context) (NetworkSetupRequestResponseBody, error) {
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
