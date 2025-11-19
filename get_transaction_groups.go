package inexchange

import (
	"context"
	"net/http"
	"net/url"

	"github.com/omniboost/go-inexchange/utils"
)

func (c *Client) NewGetTransactionGroupsRequest() GetTransactionGroupsRequest {
	return GetTransactionGroupsRequest{
		client:      c,
		queryParams: c.NewGetTransactionGroupsQueryParams(),
		pathParams:  c.NewGetTransactionGroupsPathParams(),
		method:      http.MethodGet,
		headers:     c.NewGetTransactionGroupsHeaders(),
		requestBody: c.NewGetTransactionGroupsRequestBody(),
	}
}

type GetTransactionGroupsRequest struct {
	client      *Client
	queryParams *GetTransactionGroupsQueryParams
	pathParams  *GetTransactionGroupsPathParams
	method      string
	headers     *GetTransactionGroupsHeaders
	requestBody GetTransactionGroupsRequestBody
}

func (c *Client) NewGetTransactionGroupsQueryParams() *GetTransactionGroupsQueryParams {
	return &GetTransactionGroupsQueryParams{}
}

type GetTransactionGroupsQueryParams struct {
	ZoneID     string `schema:"zoneId,omitempty"`
	Sort       string `schema:"sort,omitempty"`
	Includes   string `schema:"includes,omitempty"`
	PageNumber int32  `schema:"pageNumber,omitempty"`
	PageSize   int32  `schema:"pageSize,omitempty"`
	Filter     string `schema:"filter,omitempty"`
}

func (p GetTransactionGroupsQueryParams) ToURLValues() (url.Values, error) {
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

func (r *GetTransactionGroupsRequest) QueryParams() *GetTransactionGroupsQueryParams {
	return r.queryParams
}

func (c *Client) NewGetTransactionGroupsHeaders() *GetTransactionGroupsHeaders {
	return &GetTransactionGroupsHeaders{}
}

type GetTransactionGroupsHeaders struct {
	TenantID   string `schema:"AC-Tenant-ID,omitempty"`
	PropertyID string `schema:"AC-Property-ID,omitempty"`
}

func (r *GetTransactionGroupsRequest) Headers() *GetTransactionGroupsHeaders {
	return r.headers
}

func (c *Client) NewGetTransactionGroupsPathParams() *GetTransactionGroupsPathParams {
	return &GetTransactionGroupsPathParams{}
}

type GetTransactionGroupsPathParams struct {
}

func (p *GetTransactionGroupsPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GetTransactionGroupsRequest) PathParams() *GetTransactionGroupsPathParams {
	return r.pathParams
}

func (r *GetTransactionGroupsRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *GetTransactionGroupsRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetTransactionGroupsRequest) Method() string {
	return r.method
}

func (s *Client) NewGetTransactionGroupsRequestBody() GetTransactionGroupsRequestBody {
	return GetTransactionGroupsRequestBody{}
}

type GetTransactionGroupsRequestBody struct {
}

func (r *GetTransactionGroupsRequest) RequestBody() *GetTransactionGroupsRequestBody {
	return nil
}

func (r *GetTransactionGroupsRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *GetTransactionGroupsRequest) SetRequestBody(body GetTransactionGroupsRequestBody) {
	r.requestBody = body
}

func (r *GetTransactionGroupsRequest) NewResponseBody() *GetTransactionGroupsResponseBody {
	return &GetTransactionGroupsResponseBody{}
}

type GetTransactionGroupsResponseBody PaginatedResponse[TransactionGroup]

func (r *GetTransactionGroupsRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("api-gateway/configuration/v1/entities/TransactionGroup", r.PathParams())
	return &u
}

func (r *GetTransactionGroupsRequest) Do(ctx context.Context) (GetTransactionGroupsResponseBody, error) {
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

func (r *GetTransactionGroupsRequest) All(ctx context.Context) ([]TransactionGroup, error) {
	transactionGroups := []TransactionGroup{}
	for {
		resp, err := r.Do(ctx)
		if err != nil {
			return transactionGroups, err
		}

		// Break out of loop when no transactionGroups are found
		if len(resp.Results) == 0 {
			break
		}

		// Add transactionGroups to list
		transactionGroups = append(transactionGroups, resp.Results...)

		if len(transactionGroups) == resp.Paging.TotalCount {
			break
		}

		// Increment page number
		r.QueryParams().PageNumber = r.QueryParams().PageNumber + 1
	}

	return transactionGroups, nil
}
