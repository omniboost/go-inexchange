package inexchange

import (
	"context"
	"net/http"
	"net/url"

	"github.com/omniboost/go-inexchange/utils"
)

func (c *Client) NewGetTravelAgentProfilesRequest() GetTravelAgentProfilesRequest {
	return GetTravelAgentProfilesRequest{
		client:      c,
		queryParams: c.NewGetTravelAgentProfilesQueryParams(),
		pathParams:  c.NewGetTravelAgentProfilesPathParams(),
		method:      http.MethodGet,
		headers:     c.NewGetTravelAgentProfilesHeaders(),
		requestBody: c.NewGetTravelAgentProfilesRequestBody(),
	}
}

type GetTravelAgentProfilesRequest struct {
	client      *Client
	queryParams *GetTravelAgentProfilesQueryParams
	pathParams  *GetTravelAgentProfilesPathParams
	method      string
	headers     *GetTravelAgentProfilesHeaders
	requestBody GetTravelAgentProfilesRequestBody
}

func (c *Client) NewGetTravelAgentProfilesQueryParams() *GetTravelAgentProfilesQueryParams {
	return &GetTravelAgentProfilesQueryParams{}
}

type GetTravelAgentProfilesQueryParams struct {
	Query      string `schema:"query,omitempty"`
	Sort       string `schema:"sort,omitempty"`
	Fields     string `schema:"fields,omitempty"`
	PageNumber int32  `schema:"pageNumber,omitempty"`
	PageSize   int32  `schema:"pageSize,omitempty"`
	Filter     string `schema:"filter,omitempty"`
}

func (p GetTravelAgentProfilesQueryParams) ToURLValues() (url.Values, error) {
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

func (r *GetTravelAgentProfilesRequest) QueryParams() *GetTravelAgentProfilesQueryParams {
	return r.queryParams
}

func (c *Client) NewGetTravelAgentProfilesHeaders() *GetTravelAgentProfilesHeaders {
	return &GetTravelAgentProfilesHeaders{}
}

type GetTravelAgentProfilesHeaders struct {
	PropertyID string `schema:"AC-Property-ID,omitempty"`
}

func (r *GetTravelAgentProfilesRequest) Headers() *GetTravelAgentProfilesHeaders {
	return r.headers
}

func (c *Client) NewGetTravelAgentProfilesPathParams() *GetTravelAgentProfilesPathParams {
	return &GetTravelAgentProfilesPathParams{}
}

type GetTravelAgentProfilesPathParams struct {
}

func (p *GetTravelAgentProfilesPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GetTravelAgentProfilesRequest) PathParams() *GetTravelAgentProfilesPathParams {
	return r.pathParams
}

func (r *GetTravelAgentProfilesRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *GetTravelAgentProfilesRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetTravelAgentProfilesRequest) Method() string {
	return r.method
}

func (s *Client) NewGetTravelAgentProfilesRequestBody() GetTravelAgentProfilesRequestBody {
	return GetTravelAgentProfilesRequestBody{}
}

type GetTravelAgentProfilesRequestBody struct {
}

func (r *GetTravelAgentProfilesRequest) RequestBody() *GetTravelAgentProfilesRequestBody {
	return nil
}

func (r *GetTravelAgentProfilesRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *GetTravelAgentProfilesRequest) SetRequestBody(body GetTravelAgentProfilesRequestBody) {
	r.requestBody = body
}

func (r *GetTravelAgentProfilesRequest) NewResponseBody() *GetTravelAgentProfilesResponseBody {
	return &GetTravelAgentProfilesResponseBody{}
}

type GetTravelAgentProfilesResponseBody PaginatedResponse[TravelAgentProfile]

func (r *GetTravelAgentProfilesRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("api-gateway/profiles/v1/travel-agent", r.PathParams())
	return &u
}

func (r *GetTravelAgentProfilesRequest) Do(ctx context.Context) (GetTravelAgentProfilesResponseBody, error) {
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

func (r *GetTravelAgentProfilesRequest) All(ctx context.Context) ([]TravelAgentProfile, error) {
	travelAgencies := []TravelAgentProfile{}
	for {
		resp, err := r.Do(ctx)
		if err != nil {
			return travelAgencies, err
		}

		// Break out of loop when no travelAgencies are found
		if len(resp.Results) == 0 {
			break
		}

		// Add travelAgencies to list
		travelAgencies = append(travelAgencies, resp.Results...)

		if len(travelAgencies) == resp.Paging.TotalCount {
			break
		}

		// Increment page number
		r.QueryParams().PageNumber = r.QueryParams().PageNumber + 1
	}

	return travelAgencies, nil
}
