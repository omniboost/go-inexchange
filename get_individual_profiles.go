package inexchange

import (
	"context"
	"net/http"
	"net/url"

	"github.com/omniboost/go-inexchange/utils"
)

func (c *Client) NewGetIndividualProfilesRequest() GetIndividualProfilesRequest {
	return GetIndividualProfilesRequest{
		client:      c,
		queryParams: c.NewGetIndividualProfilesQueryParams(),
		pathParams:  c.NewGetIndividualProfilesPathParams(),
		method:      http.MethodGet,
		headers:     c.NewGetIndividualProfilesHeaders(),
		requestBody: c.NewGetIndividualProfilesRequestBody(),
	}
}

type GetIndividualProfilesRequest struct {
	client      *Client
	queryParams *GetIndividualProfilesQueryParams
	pathParams  *GetIndividualProfilesPathParams
	method      string
	headers     *GetIndividualProfilesHeaders
	requestBody GetIndividualProfilesRequestBody
}

func (c *Client) NewGetIndividualProfilesQueryParams() *GetIndividualProfilesQueryParams {
	return &GetIndividualProfilesQueryParams{}
}

type GetIndividualProfilesQueryParams struct {
	Query      string `schema:"query,omitempty"`
	Sort       string `schema:"sort,omitempty"`
	Fields     string `schema:"fields,omitempty"`
	Extend     string `schema:"extend,omitempty"`
	PageNumber int32  `schema:"pageNumber,omitempty"`
	PageSize   int32  `schema:"pageSize,omitempty"`
	Filter     string `schema:"filter,omitempty"`
}

func (p GetIndividualProfilesQueryParams) ToURLValues() (url.Values, error) {
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

func (r *GetIndividualProfilesRequest) QueryParams() *GetIndividualProfilesQueryParams {
	return r.queryParams
}

func (c *Client) NewGetIndividualProfilesHeaders() *GetIndividualProfilesHeaders {
	return &GetIndividualProfilesHeaders{}
}

type GetIndividualProfilesHeaders struct {
	PropertyID string `schema:"AC-Property-ID,omitempty"`
}

func (r *GetIndividualProfilesRequest) Headers() *GetIndividualProfilesHeaders {
	return r.headers
}

func (c *Client) NewGetIndividualProfilesPathParams() *GetIndividualProfilesPathParams {
	return &GetIndividualProfilesPathParams{}
}

type GetIndividualProfilesPathParams struct {
}

func (p *GetIndividualProfilesPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GetIndividualProfilesRequest) PathParams() *GetIndividualProfilesPathParams {
	return r.pathParams
}

func (r *GetIndividualProfilesRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *GetIndividualProfilesRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetIndividualProfilesRequest) Method() string {
	return r.method
}

func (s *Client) NewGetIndividualProfilesRequestBody() GetIndividualProfilesRequestBody {
	return GetIndividualProfilesRequestBody{}
}

type GetIndividualProfilesRequestBody struct {
}

func (r *GetIndividualProfilesRequest) RequestBody() *GetIndividualProfilesRequestBody {
	return nil
}

func (r *GetIndividualProfilesRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *GetIndividualProfilesRequest) SetRequestBody(body GetIndividualProfilesRequestBody) {
	r.requestBody = body
}

func (r *GetIndividualProfilesRequest) NewResponseBody() *GetIndividualProfilesResponseBody {
	return &GetIndividualProfilesResponseBody{}
}

type GetIndividualProfilesResponseBody PaginatedResponse[IndividualProfile]

func (r *GetIndividualProfilesRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("api-gateway/profiles/v1/individual", r.PathParams())
	return &u
}

func (r *GetIndividualProfilesRequest) Do(ctx context.Context) (GetIndividualProfilesResponseBody, error) {
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

func (r *GetIndividualProfilesRequest) All(ctx context.Context) ([]IndividualProfile, error) {
	individuals := []IndividualProfile{}
	for {
		resp, err := r.Do(ctx)
		if err != nil {
			return individuals, err
		}

		// Break out of loop when no individuals are found
		if len(resp.Results) == 0 {
			break
		}

		// Add individuals to list
		individuals = append(individuals, resp.Results...)

		if len(individuals) == resp.Paging.TotalCount {
			break
		}

		// Increment page number
		r.QueryParams().PageNumber = r.QueryParams().PageNumber + 1
	}

	return individuals, nil
}
