package inexchange

import (
	"context"
	"net/http"
	"net/url"

	"github.com/omniboost/go-inexchange/utils"
)

func (c *Client) NewGetProfileDashboardRequest() GetProfileDashboardRequest {
	return GetProfileDashboardRequest{
		client:      c,
		queryParams: c.NewGetProfileDashboardQueryParams(),
		pathParams:  c.NewGetProfileDashboardPathParams(),
		method:      http.MethodGet,
		headers:     c.NewGetProfileDashboardHeaders(),
		requestBody: c.NewGetProfileDashboardRequestBody(),
	}
}

type GetProfileDashboardRequest struct {
	client      *Client
	queryParams *GetProfileDashboardQueryParams
	pathParams  *GetProfileDashboardPathParams
	method      string
	headers     *GetProfileDashboardHeaders
	requestBody GetProfileDashboardRequestBody
}

func (c *Client) NewGetProfileDashboardQueryParams() *GetProfileDashboardQueryParams {
	return &GetProfileDashboardQueryParams{}
}

type GetProfileDashboardQueryParams struct {
	Query        string   `schema:"query,omitempty"`
	Sort         string   `schema:"sort,omitempty"`
	Aggregation  string   `schema:"aggregation,omitempty"`
	ProfileRoles []string `schema:"profileRoles,omitempty"`
	Extend       string   `schema:"extend,omitempty"`
	PageNumber   int32    `schema:"pageNumber,omitempty"`
	PageSize     int32    `schema:"pageSize,omitempty"`
	Filter       string   `schema:"filter,omitempty"`
}

func (p GetProfileDashboardQueryParams) ToURLValues() (url.Values, error) {
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

func (r *GetProfileDashboardRequest) QueryParams() *GetProfileDashboardQueryParams {
	return r.queryParams
}

func (c *Client) NewGetProfileDashboardHeaders() *GetProfileDashboardHeaders {
	return &GetProfileDashboardHeaders{}
}

type GetProfileDashboardHeaders struct {
	TenantID   string `schema:"AC-Tenant-ID,omitempty"`
	PropertyID string `schema:"AC-Property-ID,omitempty"`
}

func (r *GetProfileDashboardRequest) Headers() *GetProfileDashboardHeaders {
	return r.headers
}

func (c *Client) NewGetProfileDashboardPathParams() *GetProfileDashboardPathParams {
	return &GetProfileDashboardPathParams{}
}

type GetProfileDashboardPathParams struct {
}

func (p *GetProfileDashboardPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GetProfileDashboardRequest) PathParams() *GetProfileDashboardPathParams {
	return r.pathParams
}

func (r *GetProfileDashboardRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *GetProfileDashboardRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetProfileDashboardRequest) Method() string {
	return r.method
}

func (s *Client) NewGetProfileDashboardRequestBody() GetProfileDashboardRequestBody {
	return GetProfileDashboardRequestBody{}
}

type GetProfileDashboardRequestBody struct {
}

func (r *GetProfileDashboardRequest) RequestBody() *GetProfileDashboardRequestBody {
	return nil
}

func (r *GetProfileDashboardRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *GetProfileDashboardRequest) SetRequestBody(body GetProfileDashboardRequestBody) {
	r.requestBody = body
}

func (r *GetProfileDashboardRequest) NewResponseBody() *GetProfileDashboardResponseBody {
	return &GetProfileDashboardResponseBody{}
}

type GetProfileDashboardResponseBody PaginatedResponse[ProfileResponse]

func (r *GetProfileDashboardRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("api-gateway/profiles/v1/views/dashboard", r.PathParams())
	return &u
}

func (r *GetProfileDashboardRequest) Do(ctx context.Context) (GetProfileDashboardResponseBody, error) {
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

func (r *GetProfileDashboardRequest) All(ctx context.Context) ([]ProfileResponse, error) {
	profileResponses := []ProfileResponse{}
	for {
		resp, err := r.Do(ctx)
		if err != nil {
			return profileResponses, err
		}

		// Break out of loop when no profileResponses are found
		if len(resp.Results) == 0 {
			break
		}

		// Add profileResponses to list
		profileResponses = append(profileResponses, resp.Results...)

		if len(profileResponses) == resp.Paging.TotalCount {
			break
		}

		// Increment page number
		r.QueryParams().PageNumber = r.QueryParams().PageNumber + 1
	}

	return profileResponses, nil
}
