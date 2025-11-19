package inexchange

import (
	"context"
	"net/http"
	"net/url"

	"github.com/omniboost/go-inexchange/utils"
)

func (c *Client) NewGetPropertyLocalDateTimeRequest() GetPropertyLocalDateTimeRequest {
	return GetPropertyLocalDateTimeRequest{
		client:      c,
		queryParams: c.NewGetPropertyLocalDateTimeQueryParams(),
		pathParams:  c.NewGetPropertyLocalDateTimePathParams(),
		method:      http.MethodGet,
		headers:     c.NewGetPropertyLocalDateTimeHeaders(),
		requestBody: c.NewGetPropertyLocalDateTimeRequestBody(),
	}
}

type GetPropertyLocalDateTimeRequest struct {
	client      *Client
	queryParams *GetPropertyLocalDateTimeQueryParams
	pathParams  *GetPropertyLocalDateTimePathParams
	method      string
	headers     *GetPropertyLocalDateTimeHeaders
	requestBody GetPropertyLocalDateTimeRequestBody
}

func (c *Client) NewGetPropertyLocalDateTimeQueryParams() *GetPropertyLocalDateTimeQueryParams {
	return &GetPropertyLocalDateTimeQueryParams{}
}

type GetPropertyLocalDateTimeQueryParams struct{}

func (p GetPropertyLocalDateTimeQueryParams) ToURLValues() (url.Values, error) {
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

func (r *GetPropertyLocalDateTimeRequest) QueryParams() *GetPropertyLocalDateTimeQueryParams {
	return r.queryParams
}

func (c *Client) NewGetPropertyLocalDateTimeHeaders() *GetPropertyLocalDateTimeHeaders {
	return &GetPropertyLocalDateTimeHeaders{}
}

type GetPropertyLocalDateTimeHeaders struct {
	TenantID   string `schema:"AC-Tenant-ID,omitempty"`
	PropertyID string `schema:"AC-Property-ID,omitempty"`
}

func (r *GetPropertyLocalDateTimeRequest) Headers() *GetPropertyLocalDateTimeHeaders {
	return r.headers
}

func (c *Client) NewGetPropertyLocalDateTimePathParams() *GetPropertyLocalDateTimePathParams {
	return &GetPropertyLocalDateTimePathParams{}
}

type GetPropertyLocalDateTimePathParams struct {
}

func (p *GetPropertyLocalDateTimePathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GetPropertyLocalDateTimeRequest) PathParams() *GetPropertyLocalDateTimePathParams {
	return r.pathParams
}

func (r *GetPropertyLocalDateTimeRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *GetPropertyLocalDateTimeRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetPropertyLocalDateTimeRequest) Method() string {
	return r.method
}

func (s *Client) NewGetPropertyLocalDateTimeRequestBody() GetPropertyLocalDateTimeRequestBody {
	return GetPropertyLocalDateTimeRequestBody{}
}

type GetPropertyLocalDateTimeRequestBody struct {
}

func (r *GetPropertyLocalDateTimeRequest) RequestBody() *GetPropertyLocalDateTimeRequestBody {
	return nil
}

func (r *GetPropertyLocalDateTimeRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *GetPropertyLocalDateTimeRequest) SetRequestBody(body GetPropertyLocalDateTimeRequestBody) {
	r.requestBody = body
}

func (r *GetPropertyLocalDateTimeRequest) NewResponseBody() *GetPropertyLocalDateTimeResponseBody {
	return &GetPropertyLocalDateTimeResponseBody{}
}

type GetPropertyLocalDateTimeResponseBody struct {
	LocalPropertyDateTime DateTime `json:"localPropertyDateTime"`
	TimeZoneCode          string   `json:"timeZoneCode"`
}

func (r *GetPropertyLocalDateTimeRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("api-gateway/business-day/v1/property-local-date-time", r.PathParams())
	return &u
}

func (r *GetPropertyLocalDateTimeRequest) Do(ctx context.Context) (GetPropertyLocalDateTimeResponseBody, error) {
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
