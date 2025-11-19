package inexchange

import (
	"context"
	"net/http"
	"net/url"

	"github.com/omniboost/go-inexchange/utils"
)

func (c *Client) NewGetAccountGroupsRequest() GetAccountGroupsRequest {
	return GetAccountGroupsRequest{
		client:      c,
		queryParams: c.NewGetAccountGroupsQueryParams(),
		pathParams:  c.NewGetAccountGroupsPathParams(),
		method:      http.MethodGet,
		headers:     c.NewGetAccountGroupsHeaders(),
		requestBody: c.NewGetAccountGroupsRequestBody(),
	}
}

type GetAccountGroupsRequest struct {
	client      *Client
	queryParams *GetAccountGroupsQueryParams
	pathParams  *GetAccountGroupsPathParams
	method      string
	headers     *GetAccountGroupsHeaders
	requestBody GetAccountGroupsRequestBody
}

func (c *Client) NewGetAccountGroupsQueryParams() *GetAccountGroupsQueryParams {
	return &GetAccountGroupsQueryParams{}
}

type GetAccountGroupsQueryParams struct {
	ZoneID     string `schema:"zoneId,omitempty"`
	Sort       string `schema:"sort,omitempty"`
	Includes   string `schema:"includes,omitempty"`
	PageNumber int32  `schema:"pageNumber,omitempty"`
	PageSize   int32  `schema:"pageSize,omitempty"`
	Filter     string `schema:"filter,omitempty"`
}

func (p GetAccountGroupsQueryParams) ToURLValues() (url.Values, error) {
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

func (r *GetAccountGroupsRequest) QueryParams() *GetAccountGroupsQueryParams {
	return r.queryParams
}

func (c *Client) NewGetAccountGroupsHeaders() *GetAccountGroupsHeaders {
	return &GetAccountGroupsHeaders{}
}

type GetAccountGroupsHeaders struct {
	TenantID   string `schema:"AC-Tenant-ID,omitempty"`
	PropertyID string `schema:"AC-Property-ID,omitempty"`
}

func (r *GetAccountGroupsRequest) Headers() *GetAccountGroupsHeaders {
	return r.headers
}

func (c *Client) NewGetAccountGroupsPathParams() *GetAccountGroupsPathParams {
	return &GetAccountGroupsPathParams{}
}

type GetAccountGroupsPathParams struct {
}

func (p *GetAccountGroupsPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GetAccountGroupsRequest) PathParams() *GetAccountGroupsPathParams {
	return r.pathParams
}

func (r *GetAccountGroupsRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *GetAccountGroupsRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetAccountGroupsRequest) Method() string {
	return r.method
}

func (s *Client) NewGetAccountGroupsRequestBody() GetAccountGroupsRequestBody {
	return GetAccountGroupsRequestBody{}
}

type GetAccountGroupsRequestBody struct {
}

func (r *GetAccountGroupsRequest) RequestBody() *GetAccountGroupsRequestBody {
	return nil
}

func (r *GetAccountGroupsRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *GetAccountGroupsRequest) SetRequestBody(body GetAccountGroupsRequestBody) {
	r.requestBody = body
}

func (r *GetAccountGroupsRequest) NewResponseBody() *GetAccountGroupsResponseBody {
	return &GetAccountGroupsResponseBody{}
}

type GetAccountGroupsResponseBody PaginatedResponse[AccountGroup]

func (r *GetAccountGroupsRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("api-gateway/configuration/v1/entities/AccountGroup", r.PathParams())
	return &u
}

func (r *GetAccountGroupsRequest) Do(ctx context.Context) (GetAccountGroupsResponseBody, error) {
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

func (r *GetAccountGroupsRequest) All(ctx context.Context) ([]AccountGroup, error) {
	accountGroups := []AccountGroup{}
	for {
		resp, err := r.Do(ctx)
		if err != nil {
			return accountGroups, err
		}

		// Break out of loop when no accountGroups are found
		if len(resp.Results) == 0 {
			break
		}

		// Add accountGroups to list
		accountGroups = append(accountGroups, resp.Results...)

		if len(accountGroups) == resp.Paging.TotalCount {
			break
		}

		// Increment page number
		r.QueryParams().PageNumber = r.QueryParams().PageNumber + 1
	}

	return accountGroups, nil
}
