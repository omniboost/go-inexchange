package inexchange

import (
	"context"
	"net/http"
	"net/url"

	"github.com/omniboost/go-inexchange/utils"
)

func (c *Client) NewGetFolioTransactionOfAccountByIDRequest() GetFolioTransactionOfAccountByIDRequest {
	return GetFolioTransactionOfAccountByIDRequest{
		client:      c,
		queryParams: c.NewGetFolioTransactionOfAccountByIDQueryParams(),
		pathParams:  c.NewGetFolioTransactionOfAccountByIDPathParams(),
		method:      http.MethodGet,
		headers:     c.NewGetFolioTransactionOfAccountByIDHeaders(),
		requestBody: c.NewGetFolioTransactionOfAccountByIDRequestBody(),
	}
}

type GetFolioTransactionOfAccountByIDRequest struct {
	client      *Client
	queryParams *GetFolioTransactionOfAccountByIDQueryParams
	pathParams  *GetFolioTransactionOfAccountByIDPathParams
	method      string
	headers     *GetFolioTransactionOfAccountByIDHeaders
	requestBody GetFolioTransactionOfAccountByIDRequestBody
}

func (c *Client) NewGetFolioTransactionOfAccountByIDQueryParams() *GetFolioTransactionOfAccountByIDQueryParams {
	return &GetFolioTransactionOfAccountByIDQueryParams{}
}

type GetFolioTransactionOfAccountByIDQueryParams struct {
	PageNumber int32  `schema:"pageNumber,omitempty"`
	PageSize   int32  `schema:"pageSize,omitempty"`
	Search     string `schema:"search,omitempty"`
	Sort       string `schema:"sort,omitempty"`
}

func (p GetFolioTransactionOfAccountByIDQueryParams) ToURLValues() (url.Values, error) {
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

func (r *GetFolioTransactionOfAccountByIDRequest) QueryParams() *GetFolioTransactionOfAccountByIDQueryParams {
	return r.queryParams
}

func (c *Client) NewGetFolioTransactionOfAccountByIDHeaders() *GetFolioTransactionOfAccountByIDHeaders {
	return &GetFolioTransactionOfAccountByIDHeaders{}
}

type GetFolioTransactionOfAccountByIDHeaders struct {
	TenantID   string `schema:"AC-Tenant-ID,omitempty"`
	PropertyID string `schema:"AC-Property-ID,omitempty"`
}

func (r *GetFolioTransactionOfAccountByIDRequest) Headers() *GetFolioTransactionOfAccountByIDHeaders {
	return r.headers
}

func (c *Client) NewGetFolioTransactionOfAccountByIDPathParams() *GetFolioTransactionOfAccountByIDPathParams {
	return &GetFolioTransactionOfAccountByIDPathParams{}
}

type GetFolioTransactionOfAccountByIDPathParams struct {
	AccountID     string `schema:"AccountID"`
	FolioID       string `schema:"FolioID"`
	TransactionID string `schema:"TransactionID"`
}

func (p *GetFolioTransactionOfAccountByIDPathParams) Params() map[string]string {
	return map[string]string{
		"accountId":     p.AccountID,
		"folioId":       p.FolioID,
		"transactionId": p.TransactionID,
	}
}

func (r *GetFolioTransactionOfAccountByIDRequest) PathParams() *GetFolioTransactionOfAccountByIDPathParams {
	return r.pathParams
}

func (r *GetFolioTransactionOfAccountByIDRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *GetFolioTransactionOfAccountByIDRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetFolioTransactionOfAccountByIDRequest) Method() string {
	return r.method
}

func (s *Client) NewGetFolioTransactionOfAccountByIDRequestBody() GetFolioTransactionOfAccountByIDRequestBody {
	return GetFolioTransactionOfAccountByIDRequestBody{}
}

type GetFolioTransactionOfAccountByIDRequestBody struct {
}

func (r *GetFolioTransactionOfAccountByIDRequest) RequestBody() *GetFolioTransactionOfAccountByIDRequestBody {
	return nil
}

func (r *GetFolioTransactionOfAccountByIDRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *GetFolioTransactionOfAccountByIDRequest) SetRequestBody(body GetFolioTransactionOfAccountByIDRequestBody) {
	r.requestBody = body
}

func (r *GetFolioTransactionOfAccountByIDRequest) NewResponseBody() *GetFolioTransactionOfAccountByIDResponseBody {
	return &GetFolioTransactionOfAccountByIDResponseBody{}
}

type GetFolioTransactionOfAccountByIDResponseBody FolioTransaction

func (r *GetFolioTransactionOfAccountByIDRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("api-gateway/cashiering/v1/accounts/{{.accountId}}/folios/{{.folioId}}/transactions/{{.transactionId}}", r.PathParams())
	return &u
}

func (r *GetFolioTransactionOfAccountByIDRequest) Do(ctx context.Context) (GetFolioTransactionOfAccountByIDResponseBody, error) {
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
