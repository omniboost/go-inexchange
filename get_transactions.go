package inexchange

import (
	"context"
	"net/http"
	"net/url"

	"github.com/omniboost/go-inexchange/utils"
)

func (c *Client) NewGetTransactionsRequest() GetTransactionsRequest {
	return GetTransactionsRequest{
		client:      c,
		queryParams: c.NewGetTransactionsQueryParams(),
		pathParams:  c.NewGetTransactionsPathParams(),
		method:      http.MethodGet,
		headers:     c.NewGetTransactionsHeaders(),
		requestBody: c.NewGetTransactionsRequestBody(),
	}
}

type GetTransactionsRequest struct {
	client      *Client
	queryParams *GetTransactionsQueryParams
	pathParams  *GetTransactionsPathParams
	method      string
	headers     *GetTransactionsHeaders
	requestBody GetTransactionsRequestBody
}

func (c *Client) NewGetTransactionsQueryParams() *GetTransactionsQueryParams {
	return &GetTransactionsQueryParams{}
}

type GetTransactionsQueryParams struct {
	StartDate     Date                     `schema:"startDate,omitempty"`
	EndDate       Date                     `schema:"endDate,omitempty"`
	DateType      string                   `schema:"dateType,omitempty"`
	Extend        CommaSeparatedQueryParam `schema:"extend,omitempty"`
	PageSize      int32                    `schema:"pageSize,omitempty"`
	SeekReference string                   `schema:"seekReference,omitempty"`
	LedgerTypes   CommaSeparatedQueryParam `schema:"ledgerTypes,omitempty"`
}

func (p GetTransactionsQueryParams) ToURLValues() (url.Values, error) {
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

func (r *GetTransactionsRequest) QueryParams() *GetTransactionsQueryParams {
	return r.queryParams
}

func (c *Client) NewGetTransactionsHeaders() *GetTransactionsHeaders {
	return &GetTransactionsHeaders{}
}

type GetTransactionsHeaders struct {
	PropertyID string `schema:"AC-Property-ID,omitempty"`
}

func (r *GetTransactionsRequest) Headers() *GetTransactionsHeaders {
	return r.headers
}

func (c *Client) NewGetTransactionsPathParams() *GetTransactionsPathParams {
	return &GetTransactionsPathParams{}
}

type GetTransactionsPathParams struct {
}

func (p *GetTransactionsPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GetTransactionsRequest) PathParams() *GetTransactionsPathParams {
	return r.pathParams
}

func (r *GetTransactionsRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *GetTransactionsRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetTransactionsRequest) Method() string {
	return r.method
}

func (s *Client) NewGetTransactionsRequestBody() GetTransactionsRequestBody {
	return GetTransactionsRequestBody{}
}

type GetTransactionsRequestBody struct {
}

func (r *GetTransactionsRequest) RequestBody() *GetTransactionsRequestBody {
	return nil
}

func (r *GetTransactionsRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *GetTransactionsRequest) SetRequestBody(body GetTransactionsRequestBody) {
	r.requestBody = body
}

func (r *GetTransactionsRequest) NewResponseBody() *GetTransactionsResponseBody {
	return &GetTransactionsResponseBody{}
}

type GetTransactionsResponseBody struct {
	Results Transactions `json:"results"`
	Paging  struct {
		PageSize          int    `json:"pageSize"`
		NextSeekReference string `json:"nextSeekReference"`
	} `json:"paging"`
}

func (r *GetTransactionsRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("api-gateway/data-api/v1/transactions", r.PathParams())
	return &u
}

func (r *GetTransactionsRequest) Do(ctx context.Context) (GetTransactionsResponseBody, error) {
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
	if r.Headers().PropertyID != "" {
		req.Header.Add("AC-Property-ID", r.Headers().PropertyID)
	}

	responseBody := r.NewResponseBody()
	_, err = r.client.Do(req, responseBody)
	return *responseBody, err
}

func (r *GetTransactionsRequest) All(ctx context.Context) (Transactions, error) {
	transactions := Transactions{}
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

		if resp.Paging.NextSeekReference == "" {
			break
		}

		// Increment page number
		r.QueryParams().SeekReference = resp.Paging.NextSeekReference
	}

	return transactions, nil
}
