package inexchange

import (
	"context"
	"net/http"
	"net/url"

	"github.com/omniboost/go-inexchange/utils"
)

func (c *Client) NewGetCustomersRequest() GetCustomersRequest {
	return GetCustomersRequest{
		client:      c,
		queryParams: c.NewGetCustomersQueryParams(),
		pathParams:  c.NewGetCustomersPathParams(),
		method:      http.MethodGet,
		headers:     c.NewGetCustomersHeaders(),
		requestBody: c.NewGetCustomersRequestBody(),
	}
}

type GetCustomersRequest struct {
	client      *Client
	queryParams *GetCustomersQueryParams
	pathParams  *GetCustomersPathParams
	method      string
	headers     *GetCustomersHeaders
	requestBody GetCustomersRequestBody
}

func (c *Client) NewGetCustomersQueryParams() *GetCustomersQueryParams {
	return &GetCustomersQueryParams{}
}

type GetCustomersQueryParams struct {
	StatusCode string `schema:"statusCode,omitempty"`
	Sort       string `schema:"sort,omitempty"`
	PageNumber int32  `schema:"pageNumber,omitempty"`
	PageSize   int32  `schema:"pageSize,omitempty"`
	Filter     string `schema:"filter,omitempty"`
}

func (p GetCustomersQueryParams) ToURLValues() (url.Values, error) {
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

func (r *GetCustomersRequest) QueryParams() *GetCustomersQueryParams {
	return r.queryParams
}

func (c *Client) NewGetCustomersHeaders() *GetCustomersHeaders {
	return &GetCustomersHeaders{}
}

type GetCustomersHeaders struct{}

func (r *GetCustomersRequest) Headers() *GetCustomersHeaders {
	return r.headers
}

func (c *Client) NewGetCustomersPathParams() *GetCustomersPathParams {
	return &GetCustomersPathParams{}
}

type GetCustomersPathParams struct {
}

func (p *GetCustomersPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GetCustomersRequest) PathParams() *GetCustomersPathParams {
	return r.pathParams
}

func (r *GetCustomersRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *GetCustomersRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetCustomersRequest) Method() string {
	return r.method
}

func (s *Client) NewGetCustomersRequestBody() GetCustomersRequestBody {
	return GetCustomersRequestBody{}
}

type GetCustomersRequestBody struct {
}

func (r *GetCustomersRequest) RequestBody() *GetCustomersRequestBody {
	return nil
}

func (r *GetCustomersRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *GetCustomersRequest) SetRequestBody(body GetCustomersRequestBody) {
	r.requestBody = body
}

func (r *GetCustomersRequest) NewResponseBody() *GetCustomersResponseBody {
	return &GetCustomersResponseBody{}
}

type GetCustomersResponseBody PaginatedResponse[Customer]

func (r *GetCustomersRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("api-gateway/configuration/v1/customers", r.PathParams())
	return &u
}

func (r *GetCustomersRequest) Do(ctx context.Context) (GetCustomersResponseBody, error) {
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

	responseBody := r.NewResponseBody()
	_, err = r.client.Do(req, responseBody)
	return *responseBody, err
}

func (r *GetCustomersRequest) All(ctx context.Context) ([]Customer, error) {
	customers := []Customer{}
	for {
		resp, err := r.Do(ctx)
		if err != nil {
			return customers, err
		}

		// Break out of loop when no customers are found
		if len(resp.Results) == 0 {
			break
		}

		// Add customers to list
		customers = append(customers, resp.Results...)

		if len(customers) == resp.Paging.TotalCount {
			break
		}

		// Increment page number
		r.QueryParams().PageNumber = r.QueryParams().PageNumber + 1
	}

	return customers, nil
}
