package inexchange

import (
	"context"
	"net/http"
	"net/url"

	"github.com/omniboost/go-inexchange/utils"
)

func (c *Client) NewGetTransactionSubGroupsRequest() GetTransactionSubGroupsRequest {
	return GetTransactionSubGroupsRequest{
		client:      c,
		queryParams: c.NewGetTransactionSubGroupsQueryParams(),
		pathParams:  c.NewGetTransactionSubGroupsPathParams(),
		method:      http.MethodGet,
		headers:     c.NewGetTransactionSubGroupsHeaders(),
		requestBody: c.NewGetTransactionSubGroupsRequestBody(),
	}
}

type GetTransactionSubGroupsRequest struct {
	client      *Client
	queryParams *GetTransactionSubGroupsQueryParams
	pathParams  *GetTransactionSubGroupsPathParams
	method      string
	headers     *GetTransactionSubGroupsHeaders
	requestBody GetTransactionSubGroupsRequestBody
}

func (c *Client) NewGetTransactionSubGroupsQueryParams() *GetTransactionSubGroupsQueryParams {
	return &GetTransactionSubGroupsQueryParams{}
}

type GetTransactionSubGroupsQueryParams struct {
	ZoneID     string `schema:"zoneId,omitempty"`
	Sort       string `schema:"sort,omitempty"`
	Includes   string `schema:"includes,omitempty"`
	PageNumber int32  `schema:"pageNumber,omitempty"`
	PageSize   int32  `schema:"pageSize,omitempty"`
	Filter     string `schema:"filter,omitempty"`
}

func (p GetTransactionSubGroupsQueryParams) ToURLValues() (url.Values, error) {
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

func (r *GetTransactionSubGroupsRequest) QueryParams() *GetTransactionSubGroupsQueryParams {
	return r.queryParams
}

func (c *Client) NewGetTransactionSubGroupsHeaders() *GetTransactionSubGroupsHeaders {
	return &GetTransactionSubGroupsHeaders{}
}

type GetTransactionSubGroupsHeaders struct {
	TenantID   string `schema:"AC-Tenant-ID,omitempty"`
	PropertyID string `schema:"AC-Property-ID,omitempty"`
}

func (r *GetTransactionSubGroupsRequest) Headers() *GetTransactionSubGroupsHeaders {
	return r.headers
}

func (c *Client) NewGetTransactionSubGroupsPathParams() *GetTransactionSubGroupsPathParams {
	return &GetTransactionSubGroupsPathParams{}
}

type GetTransactionSubGroupsPathParams struct {
}

func (p *GetTransactionSubGroupsPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GetTransactionSubGroupsRequest) PathParams() *GetTransactionSubGroupsPathParams {
	return r.pathParams
}

func (r *GetTransactionSubGroupsRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *GetTransactionSubGroupsRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetTransactionSubGroupsRequest) Method() string {
	return r.method
}

func (s *Client) NewGetTransactionSubGroupsRequestBody() GetTransactionSubGroupsRequestBody {
	return GetTransactionSubGroupsRequestBody{}
}

type GetTransactionSubGroupsRequestBody struct {
}

func (r *GetTransactionSubGroupsRequest) RequestBody() *GetTransactionSubGroupsRequestBody {
	return nil
}

func (r *GetTransactionSubGroupsRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *GetTransactionSubGroupsRequest) SetRequestBody(body GetTransactionSubGroupsRequestBody) {
	r.requestBody = body
}

func (r *GetTransactionSubGroupsRequest) NewResponseBody() *GetTransactionSubGroupsResponseBody {
	return &GetTransactionSubGroupsResponseBody{}
}

type GetTransactionSubGroupsResponseBody PaginatedResponse[TransactionSubGroup]

func (r *GetTransactionSubGroupsRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("api-gateway/configuration/v1/entities/TransactionSubGroup", r.PathParams())
	return &u
}

func (r *GetTransactionSubGroupsRequest) Do(ctx context.Context) (GetTransactionSubGroupsResponseBody, error) {
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

func (r *GetTransactionSubGroupsRequest) All(ctx context.Context) ([]TransactionSubGroup, error) {
	transactionSubGroups := []TransactionSubGroup{}
	for {
		resp, err := r.Do(ctx)
		if err != nil {
			return transactionSubGroups, err
		}

		// Break out of loop when no transactionSubGroups are found
		if len(resp.Results) == 0 {
			break
		}

		// Add transactionSubGroups to list
		transactionSubGroups = append(transactionSubGroups, resp.Results...)

		if len(transactionSubGroups) == resp.Paging.TotalCount {
			break
		}

		// Increment page number
		r.QueryParams().PageNumber = r.QueryParams().PageNumber + 1
	}

	return transactionSubGroups, nil
}
