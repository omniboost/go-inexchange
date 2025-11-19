package inexchange

import (
	"context"
	"net/http"
	"net/url"

	"github.com/omniboost/go-inexchange/utils"
)

func (c *Client) NewGetCurrentBusinessDateRequest() GetCurrentBusinessDateRequest {
	return GetCurrentBusinessDateRequest{
		client:      c,
		queryParams: c.NewGetCurrentBusinessDateQueryParams(),
		pathParams:  c.NewGetCurrentBusinessDatePathParams(),
		method:      http.MethodGet,
		headers:     c.NewGetCurrentBusinessDateHeaders(),
		requestBody: c.NewGetCurrentBusinessDateRequestBody(),
	}
}

type GetCurrentBusinessDateRequest struct {
	client      *Client
	queryParams *GetCurrentBusinessDateQueryParams
	pathParams  *GetCurrentBusinessDatePathParams
	method      string
	headers     *GetCurrentBusinessDateHeaders
	requestBody GetCurrentBusinessDateRequestBody
}

func (c *Client) NewGetCurrentBusinessDateQueryParams() *GetCurrentBusinessDateQueryParams {
	return &GetCurrentBusinessDateQueryParams{}
}

type GetCurrentBusinessDateQueryParams struct{}

func (p GetCurrentBusinessDateQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	encoder.RegisterEncoder(Date{}, utils.EncodeSchemaMarshaler)
	encoder.RegisterEncoder(DateTime{}, utils.EncodeSchemaMarshaler)
	encoder.RegisterEncoder(CommaSeparatedQueryParam{}, utils.EncodeSchemaMarshaler)
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *GetCurrentBusinessDateRequest) QueryParams() *GetCurrentBusinessDateQueryParams {
	return r.queryParams
}

func (c *Client) NewGetCurrentBusinessDateHeaders() *GetCurrentBusinessDateHeaders {
	return &GetCurrentBusinessDateHeaders{}
}

type GetCurrentBusinessDateHeaders struct {
	TenantID   string `schema:"AC-Tenant-ID,omitempty"`
	PropertyID string `schema:"AC-Property-ID,omitempty"`
}

func (r *GetCurrentBusinessDateRequest) Headers() *GetCurrentBusinessDateHeaders {
	return r.headers
}

func (c *Client) NewGetCurrentBusinessDatePathParams() *GetCurrentBusinessDatePathParams {
	return &GetCurrentBusinessDatePathParams{}
}

type GetCurrentBusinessDatePathParams struct {
}

func (p *GetCurrentBusinessDatePathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GetCurrentBusinessDateRequest) PathParams() *GetCurrentBusinessDatePathParams {
	return r.pathParams
}

func (r *GetCurrentBusinessDateRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *GetCurrentBusinessDateRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetCurrentBusinessDateRequest) Method() string {
	return r.method
}

func (s *Client) NewGetCurrentBusinessDateRequestBody() GetCurrentBusinessDateRequestBody {
	return GetCurrentBusinessDateRequestBody{}
}

type GetCurrentBusinessDateRequestBody struct {
}

func (r *GetCurrentBusinessDateRequest) RequestBody() *GetCurrentBusinessDateRequestBody {
	return nil
}

func (r *GetCurrentBusinessDateRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *GetCurrentBusinessDateRequest) SetRequestBody(body GetCurrentBusinessDateRequestBody) {
	r.requestBody = body
}

func (r *GetCurrentBusinessDateRequest) NewResponseBody() *GetCurrentBusinessDateResponseBody {
	return &GetCurrentBusinessDateResponseBody{}
}

type GetCurrentBusinessDateResponseBody struct {
	BusinessDate DateTime `json:"businessDate"`
}

func (r *GetCurrentBusinessDateRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("api-gateway/business-day/v1/current-date", r.PathParams())
	return &u
}

func (r *GetCurrentBusinessDateRequest) Do(ctx context.Context) (GetCurrentBusinessDateResponseBody, error) {
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

	// Add headers
	if r.Headers().TenantID != "" {
		req.Header.Add("AC-Tenant-ID", r.Headers().TenantID)
	}

	if r.Headers().PropertyID != "" {
		req.Header.Add("AC-Property-ID", r.Headers().PropertyID)
	}

	responseBody := r.NewResponseBody()
	_, err = r.client.Do(req, responseBody)
	return *responseBody, err
}
