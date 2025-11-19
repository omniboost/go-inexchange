package inexchange

import (
	"context"
	"net/http"
	"net/url"

	"github.com/omniboost/go-inexchange/utils"
)

func (c *Client) NewGetAggregatedConfigurationRequest() GetAggregatedConfigurationRequest {
	return GetAggregatedConfigurationRequest{
		client:      c,
		queryParams: c.NewGetAggregatedConfigurationQueryParams(),
		pathParams:  c.NewGetAggregatedConfigurationPathParams(),
		method:      http.MethodGet,
		headers:     c.NewGetAggregatedConfigurationHeaders(),
		requestBody: c.NewGetAggregatedConfigurationRequestBody(),
	}
}

type GetAggregatedConfigurationRequest struct {
	client      *Client
	queryParams *GetAggregatedConfigurationQueryParams
	pathParams  *GetAggregatedConfigurationPathParams
	method      string
	headers     *GetAggregatedConfigurationHeaders
	requestBody GetAggregatedConfigurationRequestBody
}

func (c *Client) NewGetAggregatedConfigurationQueryParams() *GetAggregatedConfigurationQueryParams {
	return &GetAggregatedConfigurationQueryParams{}
}

type GetAggregatedConfigurationQueryParams struct {
	Types         CommaSeparatedQueryParam `schema:"types,omitempty"`
	Country       string                   `schema:"country,omitempty"`
	PageNumber    int32                    `schema:"pageNumber,omitempty"`
	PageSize      int32                    `schema:"pageSize,omitempty"`
	IsActive      bool                     `schema:"isActive,omitempty"`
	UpdatedAt     DateTime                 `schema:"updatedAt,omitempty"`
	IncludeCopies bool                     `schema:"includeCopies,omitempty"`
}

func (p GetAggregatedConfigurationQueryParams) ToURLValues() (url.Values, error) {
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

func (r *GetAggregatedConfigurationRequest) QueryParams() *GetAggregatedConfigurationQueryParams {
	return r.queryParams
}

func (c *Client) NewGetAggregatedConfigurationHeaders() *GetAggregatedConfigurationHeaders {
	return &GetAggregatedConfigurationHeaders{}
}

type GetAggregatedConfigurationHeaders struct {
	PropertyID string `schema:"AC-Property-ID,omitempty"`
}

func (r *GetAggregatedConfigurationRequest) Headers() *GetAggregatedConfigurationHeaders {
	return r.headers
}

func (c *Client) NewGetAggregatedConfigurationPathParams() *GetAggregatedConfigurationPathParams {
	return &GetAggregatedConfigurationPathParams{}
}

type GetAggregatedConfigurationPathParams struct {
}

func (p *GetAggregatedConfigurationPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GetAggregatedConfigurationRequest) PathParams() *GetAggregatedConfigurationPathParams {
	return r.pathParams
}

func (r *GetAggregatedConfigurationRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *GetAggregatedConfigurationRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetAggregatedConfigurationRequest) Method() string {
	return r.method
}

func (s *Client) NewGetAggregatedConfigurationRequestBody() GetAggregatedConfigurationRequestBody {
	return GetAggregatedConfigurationRequestBody{}
}

type GetAggregatedConfigurationRequestBody struct {
}

func (r *GetAggregatedConfigurationRequest) RequestBody() *GetAggregatedConfigurationRequestBody {
	return nil
}

func (r *GetAggregatedConfigurationRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *GetAggregatedConfigurationRequest) SetRequestBody(body GetAggregatedConfigurationRequestBody) {
	r.requestBody = body
}

func (r *GetAggregatedConfigurationRequest) NewResponseBody() *GetAggregatedConfigurationResponseBody {
	return &GetAggregatedConfigurationResponseBody{}
}

type GetAggregatedConfigurationResponseBody AggregatedPaginatedResponse

func (r *GetAggregatedConfigurationRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("api-gateway/aggregator/v1/configuration/extract", r.PathParams())
	return &u
}

func (r *GetAggregatedConfigurationRequest) Do(ctx context.Context) (GetAggregatedConfigurationResponseBody, error) {
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
	req.Header.Add("AC-Property-ID", r.Headers().PropertyID)

	responseBody := r.NewResponseBody()
	_, err = r.client.Do(req, responseBody)
	return *responseBody, err
}

func (r *GetAggregatedConfigurationRequest) All(ctx context.Context) (AggregatedExtractResult, error) {
	result := AggregatedExtractResult{}
	for {
		resp, err := r.Do(ctx)
		if err != nil {
			return result, err
		}

		// Cast to extract result
		response := AggregatedPaginatedResponse(resp)

		// Break out of loop when no companies are found
		if response.TotalCount() == 0 {
			break
		}

		// Add results to list
		result.TaxRule = append(result.TaxRule, resp.Results.TaxRule...)
		result.TransactionCodes = append(result.TransactionCodes, resp.Results.TransactionCodes...)
		result.TransactionGroups = append(result.TransactionGroups, resp.Results.TransactionGroups...)
		result.TransactionSubGroups = append(result.TransactionSubGroups, resp.Results.TransactionSubGroups...)

		if response.TotalCount() == resp.Paging.TotalCount {
			break
		}

		// Increment page number
		r.QueryParams().PageNumber = r.QueryParams().PageNumber + 1
	}

	return result, nil
}
