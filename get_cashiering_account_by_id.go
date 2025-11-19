package inexchange

import (
	"context"
	"net/http"
	"net/url"

	"github.com/omniboost/go-inexchange/utils"
)

func (c *Client) NewGetCashieringAccountByIDRequest() GetCashieringAccountByIDRequest {
	return GetCashieringAccountByIDRequest{
		client:      c,
		queryParams: c.NewGetCashieringAccountByIDQueryParams(),
		pathParams:  c.NewGetCashieringAccountByIDPathParams(),
		method:      http.MethodGet,
		headers:     c.NewGetCashieringAccountByIDHeaders(),
		requestBody: c.NewGetCashieringAccountByIDRequestBody(),
	}
}

type GetCashieringAccountByIDRequest struct {
	client      *Client
	queryParams *GetCashieringAccountByIDQueryParams
	pathParams  *GetCashieringAccountByIDPathParams
	method      string
	headers     *GetCashieringAccountByIDHeaders
	requestBody GetCashieringAccountByIDRequestBody
}

func (c *Client) NewGetCashieringAccountByIDQueryParams() *GetCashieringAccountByIDQueryParams {
	return &GetCashieringAccountByIDQueryParams{}
}

type GetCashieringAccountByIDQueryParams struct {
	Extend string `schema:"extend,omitempty"`
}

func (p GetCashieringAccountByIDQueryParams) ToURLValues() (url.Values, error) {
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

func (r *GetCashieringAccountByIDRequest) QueryParams() *GetCashieringAccountByIDQueryParams {
	return r.queryParams
}

func (c *Client) NewGetCashieringAccountByIDHeaders() *GetCashieringAccountByIDHeaders {
	return &GetCashieringAccountByIDHeaders{}
}

type GetCashieringAccountByIDHeaders struct {
	TenantID   string `schema:"AC-Tenant-ID,omitempty"`
	PropertyID string `schema:"AC-Property-ID,omitempty"`
}

func (r *GetCashieringAccountByIDRequest) Headers() *GetCashieringAccountByIDHeaders {
	return r.headers
}

func (c *Client) NewGetCashieringAccountByIDPathParams() *GetCashieringAccountByIDPathParams {
	return &GetCashieringAccountByIDPathParams{}
}

type GetCashieringAccountByIDPathParams struct {
	ID string `schema:"id"`
}

func (p *GetCashieringAccountByIDPathParams) Params() map[string]string {
	return map[string]string{
		"id": p.ID,
	}
}

func (r *GetCashieringAccountByIDRequest) PathParams() *GetCashieringAccountByIDPathParams {
	return r.pathParams
}

func (r *GetCashieringAccountByIDRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *GetCashieringAccountByIDRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetCashieringAccountByIDRequest) Method() string {
	return r.method
}

func (s *Client) NewGetCashieringAccountByIDRequestBody() GetCashieringAccountByIDRequestBody {
	return GetCashieringAccountByIDRequestBody{}
}

type GetCashieringAccountByIDRequestBody struct {
}

func (r *GetCashieringAccountByIDRequest) RequestBody() *GetCashieringAccountByIDRequestBody {
	return nil
}

func (r *GetCashieringAccountByIDRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *GetCashieringAccountByIDRequest) SetRequestBody(body GetCashieringAccountByIDRequestBody) {
	r.requestBody = body
}

func (r *GetCashieringAccountByIDRequest) NewResponseBody() *GetCashieringAccountByIDResponseBody {
	return &GetCashieringAccountByIDResponseBody{}
}

type GetCashieringAccountByIDResponseBody CashieringAccount

func (r *GetCashieringAccountByIDRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("api-gateway/cashiering/v1/accounts/{{.id}}", r.PathParams())
	return &u
}

func (r *GetCashieringAccountByIDRequest) Do(ctx context.Context) (GetCashieringAccountByIDResponseBody, error) {
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
