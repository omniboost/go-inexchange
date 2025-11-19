package inexchange

import (
	"context"
	"net/http"
	"net/url"

	"github.com/omniboost/go-inexchange/utils"
)

func (c *Client) NewGetUsersMeRequest() GetUsersMeRequest {
	return GetUsersMeRequest{
		client:      c,
		queryParams: c.NewGetUsersMeQueryParams(),
		pathParams:  c.NewGetUsersMePathParams(),
		method:      http.MethodGet,
		headers:     http.Header{},
		requestBody: c.NewGetUsersMeRequestBody(),
	}
}

type GetUsersMeRequest struct {
	client      *Client
	queryParams *GetUsersMeQueryParams
	pathParams  *GetUsersMePathParams
	method      string
	headers     http.Header
	requestBody GetUsersMeRequestBody
}

func (c *Client) NewGetUsersMeQueryParams() *GetUsersMeQueryParams {
	return &GetUsersMeQueryParams{}
}

type GetUsersMeQueryParams struct{}

func (p GetUsersMeQueryParams) ToURLValues() (url.Values, error) {
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

func (r *GetUsersMeRequest) QueryParams() *GetUsersMeQueryParams {
	return r.queryParams
}

func (c *Client) NewGetUsersMePathParams() *GetUsersMePathParams {
	return &GetUsersMePathParams{}
}

type GetUsersMePathParams struct {
}

func (p *GetUsersMePathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GetUsersMeRequest) PathParams() *GetUsersMePathParams {
	return r.pathParams
}

func (r *GetUsersMeRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *GetUsersMeRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetUsersMeRequest) Method() string {
	return r.method
}

func (s *Client) NewGetUsersMeRequestBody() GetUsersMeRequestBody {
	return GetUsersMeRequestBody{}
}

type GetUsersMeRequestBody struct {
}

func (r *GetUsersMeRequest) RequestBody() *GetUsersMeRequestBody {
	return nil
}

func (r *GetUsersMeRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *GetUsersMeRequest) SetRequestBody(body GetUsersMeRequestBody) {
	r.requestBody = body
}

func (r *GetUsersMeRequest) NewResponseBody() *GetUsersMeResponseBody {
	return &GetUsersMeResponseBody{}
}

type GetUsersMeResponseBody UserDetails

func (r *GetUsersMeRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("api-gateway/identity/v1/users/me", r.PathParams())
	return &u
}

func (r *GetUsersMeRequest) Do(ctx context.Context) (GetUsersMeResponseBody, error) {
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
