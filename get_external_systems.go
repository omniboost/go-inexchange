package inexchange

import (
	"context"
	"net/http"
	"net/url"

	"github.com/omniboost/go-inexchange/utils"
)

func (c *Client) NewGetExternalSystemsRequest() GetExternalSystemsRequest {
	return GetExternalSystemsRequest{
		client:      c,
		queryParams: c.NewGetExternalSystemsQueryParams(),
		pathParams:  c.NewGetExternalSystemsPathParams(),
		method:      http.MethodGet,
		headers:     c.NewGetExternalSystemsHeaders(),
		requestBody: c.NewGetExternalSystemsRequestBody(),
	}
}

type GetExternalSystemsRequest struct {
	client      *Client
	queryParams *GetExternalSystemsQueryParams
	pathParams  *GetExternalSystemsPathParams
	method      string
	headers     *GetExternalSystemsHeaders
	requestBody GetExternalSystemsRequestBody
}

func (c *Client) NewGetExternalSystemsQueryParams() *GetExternalSystemsQueryParams {
	return &GetExternalSystemsQueryParams{}
}

type GetExternalSystemsQueryParams struct {
	ZoneID     string `schema:"zoneId,omitempty"`
	Sort       string `schema:"sort,omitempty"`
	Includes   string `schema:"includes,omitempty"`
	PageNumber int32  `schema:"pageNumber,omitempty"`
	PageSize   int32  `schema:"pageSize,omitempty"`
	Filter     string `schema:"filter,omitempty"`
}

func (p GetExternalSystemsQueryParams) ToURLValues() (url.Values, error) {
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

func (r *GetExternalSystemsRequest) QueryParams() *GetExternalSystemsQueryParams {
	return r.queryParams
}

func (c *Client) NewGetExternalSystemsHeaders() *GetExternalSystemsHeaders {
	return &GetExternalSystemsHeaders{}
}

type GetExternalSystemsHeaders struct {
	TenantID   string `schema:"AC-Tenant-ID,omitempty"`
	PropertyID string `schema:"AC-Property-ID,omitempty"`
}

func (r *GetExternalSystemsRequest) Headers() *GetExternalSystemsHeaders {
	return r.headers
}

func (c *Client) NewGetExternalSystemsPathParams() *GetExternalSystemsPathParams {
	return &GetExternalSystemsPathParams{}
}

type GetExternalSystemsPathParams struct {
}

func (p *GetExternalSystemsPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GetExternalSystemsRequest) PathParams() *GetExternalSystemsPathParams {
	return r.pathParams
}

func (r *GetExternalSystemsRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *GetExternalSystemsRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetExternalSystemsRequest) Method() string {
	return r.method
}

func (s *Client) NewGetExternalSystemsRequestBody() GetExternalSystemsRequestBody {
	return GetExternalSystemsRequestBody{}
}

type GetExternalSystemsRequestBody struct {
}

func (r *GetExternalSystemsRequest) RequestBody() *GetExternalSystemsRequestBody {
	return nil
}

func (r *GetExternalSystemsRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *GetExternalSystemsRequest) SetRequestBody(body GetExternalSystemsRequestBody) {
	r.requestBody = body
}

func (r *GetExternalSystemsRequest) NewResponseBody() *GetExternalSystemsResponseBody {
	return &GetExternalSystemsResponseBody{}
}

type GetExternalSystemsResponseBody PaginatedResponse[ExternalSystem]

func (r *GetExternalSystemsRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("api-gateway/configuration/v1/entities/ExternalSystem", r.PathParams())
	return &u
}

func (r *GetExternalSystemsRequest) Do(ctx context.Context) (GetExternalSystemsResponseBody, error) {
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

func (r *GetExternalSystemsRequest) All(ctx context.Context) ([]ExternalSystem, error) {
	externalSystems := []ExternalSystem{}
	for {
		resp, err := r.Do(ctx)
		if err != nil {
			return externalSystems, err
		}

		// Break out of loop when no externalSystems are found
		if len(resp.Results) == 0 {
			break
		}

		// Add externalSystems to list
		externalSystems = append(externalSystems, resp.Results...)

		if len(externalSystems) == resp.Paging.TotalCount {
			break
		}

		// Increment page number
		r.QueryParams().PageNumber = r.QueryParams().PageNumber + 1
	}

	return externalSystems, nil
}
