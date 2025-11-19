package inexchange

import (
	"context"
	"net/http"
	"net/url"

	"github.com/omniboost/go-inexchange/utils"
)

func (c *Client) NewGetFolioOfAccountByIDRequest() GetFolioOfAccountByIDRequest {
	return GetFolioOfAccountByIDRequest{
		client:      c,
		queryParams: c.NewGetFolioOfAccountByIDQueryParams(),
		pathParams:  c.NewGetFolioOfAccountByIDPathParams(),
		method:      http.MethodGet,
		headers:     c.NewGetFolioOfAccountByIDHeaders(),
		requestBody: c.NewGetFolioOfAccountByIDRequestBody(),
	}
}

type GetFolioOfAccountByIDRequest struct {
	client      *Client
	queryParams *GetFolioOfAccountByIDQueryParams
	pathParams  *GetFolioOfAccountByIDPathParams
	method      string
	headers     *GetFolioOfAccountByIDHeaders
	requestBody GetFolioOfAccountByIDRequestBody
}

func (c *Client) NewGetFolioOfAccountByIDQueryParams() *GetFolioOfAccountByIDQueryParams {
	return &GetFolioOfAccountByIDQueryParams{}
}

type GetFolioOfAccountByIDQueryParams struct{}

func (p GetFolioOfAccountByIDQueryParams) ToURLValues() (url.Values, error) {
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

func (r *GetFolioOfAccountByIDRequest) QueryParams() *GetFolioOfAccountByIDQueryParams {
	return r.queryParams
}

func (c *Client) NewGetFolioOfAccountByIDHeaders() *GetFolioOfAccountByIDHeaders {
	return &GetFolioOfAccountByIDHeaders{}
}

type GetFolioOfAccountByIDHeaders struct {
	TenantID   string `schema:"AC-Tenant-ID,omitempty"`
	PropertyID string `schema:"AC-Property-ID,omitempty"`
}

func (r *GetFolioOfAccountByIDRequest) Headers() *GetFolioOfAccountByIDHeaders {
	return r.headers
}

func (c *Client) NewGetFolioOfAccountByIDPathParams() *GetFolioOfAccountByIDPathParams {
	return &GetFolioOfAccountByIDPathParams{}
}

type GetFolioOfAccountByIDPathParams struct {
	AccountID string `schema:"AccountID"`
	FolioID   string `schema:"FolioID"`
}

func (p *GetFolioOfAccountByIDPathParams) Params() map[string]string {
	return map[string]string{
		"accountId": p.AccountID,
		"folioId":   p.FolioID,
	}
}

func (r *GetFolioOfAccountByIDRequest) PathParams() *GetFolioOfAccountByIDPathParams {
	return r.pathParams
}

func (r *GetFolioOfAccountByIDRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *GetFolioOfAccountByIDRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetFolioOfAccountByIDRequest) Method() string {
	return r.method
}

func (s *Client) NewGetFolioOfAccountByIDRequestBody() GetFolioOfAccountByIDRequestBody {
	return GetFolioOfAccountByIDRequestBody{}
}

type GetFolioOfAccountByIDRequestBody struct {
}

func (r *GetFolioOfAccountByIDRequest) RequestBody() *GetFolioOfAccountByIDRequestBody {
	return nil
}

func (r *GetFolioOfAccountByIDRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *GetFolioOfAccountByIDRequest) SetRequestBody(body GetFolioOfAccountByIDRequestBody) {
	r.requestBody = body
}

func (r *GetFolioOfAccountByIDRequest) NewResponseBody() *GetFolioOfAccountByIDResponseBody {
	return &GetFolioOfAccountByIDResponseBody{}
}

type GetFolioOfAccountByIDResponseBody FolioItem

func (r *GetFolioOfAccountByIDRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("api-gateway/cashiering/v1/accounts/{{.accountId}}/folios/{{.folioId}}", r.PathParams())
	return &u
}

func (r *GetFolioOfAccountByIDRequest) Do(ctx context.Context) (GetFolioOfAccountByIDResponseBody, error) {
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
