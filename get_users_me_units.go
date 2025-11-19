package inexchange

import (
	"context"
	"net/http"
	"net/url"

	"github.com/omniboost/go-inexchange/utils"
)

func (c *Client) NewGetUsersMeUnitsRequest() GetUsersMeUnitsRequest {
	return GetUsersMeUnitsRequest{
		client:      c,
		queryParams: c.NewGetUsersMeUnitsQueryParams(),
		pathParams:  c.NewGetUsersMeUnitsPathParams(),
		method:      http.MethodGet,
		headers:     http.Header{},
		requestBody: c.NewGetUsersMeUnitsRequestBody(),
	}
}

type GetUsersMeUnitsRequest struct {
	client      *Client
	queryParams *GetUsersMeUnitsQueryParams
	pathParams  *GetUsersMeUnitsPathParams
	method      string
	headers     http.Header
	requestBody GetUsersMeUnitsRequestBody
}

func (c *Client) NewGetUsersMeUnitsQueryParams() *GetUsersMeUnitsQueryParams {
	return &GetUsersMeUnitsQueryParams{}
}

type GetUsersMeUnitsQueryParams struct {
	PageNumber int32  `schema:"pageNumber,omitempty"`
	PageSize   int32  `schema:"pageSize,omitempty"`
	Sort       string `schema:"sort,omitempty"`
	Filter     string `schema:"filter,omitempty"`
}

func (p GetUsersMeUnitsQueryParams) ToURLValues() (url.Values, error) {
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

func (r *GetUsersMeUnitsRequest) QueryParams() *GetUsersMeUnitsQueryParams {
	return r.queryParams
}

func (c *Client) NewGetUsersMeUnitsPathParams() *GetUsersMeUnitsPathParams {
	return &GetUsersMeUnitsPathParams{}
}

type GetUsersMeUnitsPathParams struct {
}

func (p *GetUsersMeUnitsPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GetUsersMeUnitsRequest) PathParams() *GetUsersMeUnitsPathParams {
	return r.pathParams
}

func (r *GetUsersMeUnitsRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *GetUsersMeUnitsRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetUsersMeUnitsRequest) Method() string {
	return r.method
}

func (s *Client) NewGetUsersMeUnitsRequestBody() GetUsersMeUnitsRequestBody {
	return GetUsersMeUnitsRequestBody{}
}

type GetUsersMeUnitsRequestBody struct {
}

func (r *GetUsersMeUnitsRequest) RequestBody() *GetUsersMeUnitsRequestBody {
	return nil
}

func (r *GetUsersMeUnitsRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *GetUsersMeUnitsRequest) SetRequestBody(body GetUsersMeUnitsRequestBody) {
	r.requestBody = body
}

func (r *GetUsersMeUnitsRequest) NewResponseBody() *GetUsersMeUnitsResponseBody {
	return &GetUsersMeUnitsResponseBody{}
}

type GetUsersMeUnitsResponseBody PaginatedResponse[UnitAssignmentListItem]

func (r *GetUsersMeUnitsRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("api-gateway/permission-management/users/me/units", r.PathParams())
	return &u
}

func (r *GetUsersMeUnitsRequest) Do(ctx context.Context) (GetUsersMeUnitsResponseBody, error) {
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

func (r *GetUsersMeUnitsRequest) All(ctx context.Context) ([]UnitAssignmentListItem, error) {
	units := []UnitAssignmentListItem{}
	for {
		resp, err := r.Do(ctx)
		if err != nil {
			return units, err
		}

		// Break out of loop when no units are found
		if len(resp.Results) == 0 {
			break
		}

		// Add units to list
		units = append(units, resp.Results...)

		if len(units) == resp.Paging.TotalCount {
			break
		}

		// Increment page number
		r.QueryParams().PageNumber = r.QueryParams().PageNumber + 1
	}

	return units, nil
}
