package inexchange

import (
	"context"
	"net/http"
	"net/url"

	"github.com/omniboost/go-inexchange/utils"
)

func (c *Client) NewGetFolioTransactionsOfAccountByIDRequest() GetFolioTransactionsOfAccountByIDRequest {
	return GetFolioTransactionsOfAccountByIDRequest{
		client:      c,
		queryParams: c.NewGetFolioTransactionsOfAccountByIDQueryParams(),
		pathParams:  c.NewGetFolioTransactionsOfAccountByIDPathParams(),
		method:      http.MethodGet,
		headers:     c.NewGetFolioTransactionsOfAccountByIDHeaders(),
		requestBody: c.NewGetFolioTransactionsOfAccountByIDRequestBody(),
	}
}

type GetFolioTransactionsOfAccountByIDRequest struct {
	client      *Client
	queryParams *GetFolioTransactionsOfAccountByIDQueryParams
	pathParams  *GetFolioTransactionsOfAccountByIDPathParams
	method      string
	headers     *GetFolioTransactionsOfAccountByIDHeaders
	requestBody GetFolioTransactionsOfAccountByIDRequestBody
}

func (c *Client) NewGetFolioTransactionsOfAccountByIDQueryParams() *GetFolioTransactionsOfAccountByIDQueryParams {
	return &GetFolioTransactionsOfAccountByIDQueryParams{}
}

type GetFolioTransactionsOfAccountByIDQueryParams struct {
	PageNumber int32  `schema:"pageNumber,omitempty"`
	PageSize   int32  `schema:"pageSize,omitempty"`
	Search     string `schema:"search,omitempty"`
	Sort       string `schema:"sort,omitempty"`
}

func (p GetFolioTransactionsOfAccountByIDQueryParams) ToURLValues() (url.Values, error) {
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

func (r *GetFolioTransactionsOfAccountByIDRequest) QueryParams() *GetFolioTransactionsOfAccountByIDQueryParams {
	return r.queryParams
}

func (c *Client) NewGetFolioTransactionsOfAccountByIDHeaders() *GetFolioTransactionsOfAccountByIDHeaders {
	return &GetFolioTransactionsOfAccountByIDHeaders{}
}

type GetFolioTransactionsOfAccountByIDHeaders struct {
	TenantID   string `schema:"AC-Tenant-ID,omitempty"`
	PropertyID string `schema:"AC-Property-ID,omitempty"`
}

func (r *GetFolioTransactionsOfAccountByIDRequest) Headers() *GetFolioTransactionsOfAccountByIDHeaders {
	return r.headers
}

func (c *Client) NewGetFolioTransactionsOfAccountByIDPathParams() *GetFolioTransactionsOfAccountByIDPathParams {
	return &GetFolioTransactionsOfAccountByIDPathParams{}
}

type GetFolioTransactionsOfAccountByIDPathParams struct {
	AccountID string `schema:"AccountID"`
	FolioID   string `schema:"FolioID"`
}

func (p *GetFolioTransactionsOfAccountByIDPathParams) Params() map[string]string {
	return map[string]string{
		"accountId": p.AccountID,
		"folioId":   p.FolioID,
	}
}

func (r *GetFolioTransactionsOfAccountByIDRequest) PathParams() *GetFolioTransactionsOfAccountByIDPathParams {
	return r.pathParams
}

func (r *GetFolioTransactionsOfAccountByIDRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *GetFolioTransactionsOfAccountByIDRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetFolioTransactionsOfAccountByIDRequest) Method() string {
	return r.method
}

func (s *Client) NewGetFolioTransactionsOfAccountByIDRequestBody() GetFolioTransactionsOfAccountByIDRequestBody {
	return GetFolioTransactionsOfAccountByIDRequestBody{}
}

type GetFolioTransactionsOfAccountByIDRequestBody struct {
}

func (r *GetFolioTransactionsOfAccountByIDRequest) RequestBody() *GetFolioTransactionsOfAccountByIDRequestBody {
	return nil
}

func (r *GetFolioTransactionsOfAccountByIDRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *GetFolioTransactionsOfAccountByIDRequest) SetRequestBody(body GetFolioTransactionsOfAccountByIDRequestBody) {
	r.requestBody = body
}

func (r *GetFolioTransactionsOfAccountByIDRequest) NewResponseBody() *GetFolioTransactionsOfAccountByIDResponseBody {
	return &GetFolioTransactionsOfAccountByIDResponseBody{}
}

type GetFolioTransactionsOfAccountByIDResponseBody PaginatedResponse[FolioTransaction]

func (r *GetFolioTransactionsOfAccountByIDRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("api-gateway/cashiering/v1/accounts/{{.accountId}}/folios/{{.folioId}}/transactions", r.PathParams())
	return &u
}

func (r *GetFolioTransactionsOfAccountByIDRequest) Do(ctx context.Context) (GetFolioTransactionsOfAccountByIDResponseBody, error) {
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

func (r *GetFolioTransactionsOfAccountByIDRequest) All(ctx context.Context) ([]FolioTransaction, error) {
	transactions := []FolioTransaction{}

	// Set the page number to 1
	r.QueryParams().PageNumber = 1

	for {
		resp, err := r.Do(ctx)
		if err != nil {
			return transactions, err
		}

		// Break out of loop when no transactions are found
		if len(resp.Results) == 0 {
			break
		}

		// Add transactions to list
		transactions = append(transactions, resp.Results...)

		if len(transactions) == resp.Paging.TotalCount {
			break
		}

		// Increment page number
		r.QueryParams().PageNumber = r.QueryParams().PageNumber + 1
	}

	return transactions, nil
}
