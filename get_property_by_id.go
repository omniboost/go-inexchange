package inexchange

import (
	"context"
	"net/http"
	"net/url"

	"github.com/omniboost/go-inexchange/utils"
)

func (c *Client) NewGetPropertyByIDRequest() GetPropertyByIDRequest {
	return GetPropertyByIDRequest{
		client:      c,
		queryParams: c.NewGetPropertyByIDQueryParams(),
		pathParams:  c.NewGetPropertyByIDPathParams(),
		method:      http.MethodGet,
		headers:     http.Header{},
		requestBody: c.NewGetPropertyByIDRequestBody(),
	}
}

type GetPropertyByIDRequest struct {
	client      *Client
	queryParams *GetPropertyByIDQueryParams
	pathParams  *GetPropertyByIDPathParams
	method      string
	headers     http.Header
	requestBody GetPropertyByIDRequestBody
}

func (c *Client) NewGetPropertyByIDQueryParams() *GetPropertyByIDQueryParams {
	return &GetPropertyByIDQueryParams{}
}

type GetPropertyByIDQueryParams struct {
	StatusCode   string `schema:"statusCode,omitempty"`
	SubsidiaryID string `schema:"subsidiaryId,omitempty"`
	Recursive    bool   `schema:"recursive,omitempty"`
	Sort         string `schema:"sort,omitempty"`
	PageNumber   int32  `schema:"pageNumber,omitempty"`
	PageSize     int32  `schema:"pageSize,omitempty"`
	Filter       string `schema:"filter,omitempty"`
}

func (p GetPropertyByIDQueryParams) ToURLValues() (url.Values, error) {
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

func (r *GetPropertyByIDRequest) QueryParams() *GetPropertyByIDQueryParams {
	return r.queryParams
}

func (c *Client) NewGetPropertyByIDPathParams() *GetPropertyByIDPathParams {
	return &GetPropertyByIDPathParams{}
}

type GetPropertyByIDPathParams struct {
	ID string
}

func (p *GetPropertyByIDPathParams) Params() map[string]string {
	return map[string]string{
		"id": p.ID,
	}
}

func (r *GetPropertyByIDRequest) PathParams() *GetPropertyByIDPathParams {
	return r.pathParams
}

func (r *GetPropertyByIDRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *GetPropertyByIDRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetPropertyByIDRequest) Method() string {
	return r.method
}

func (s *Client) NewGetPropertyByIDRequestBody() GetPropertyByIDRequestBody {
	return GetPropertyByIDRequestBody{}
}

type GetPropertyByIDRequestBody struct {
}

func (r *GetPropertyByIDRequest) RequestBody() *GetPropertyByIDRequestBody {
	return nil
}

func (r *GetPropertyByIDRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *GetPropertyByIDRequest) SetRequestBody(body GetPropertyByIDRequestBody) {
	r.requestBody = body
}

func (r *GetPropertyByIDRequest) NewResponseBody() *GetPropertyByIDResponseBody {
	return &GetPropertyByIDResponseBody{}
}

type GetPropertyByIDResponseBody Property

func (r *GetPropertyByIDRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("api-gateway/configuration/v1/properties/{{.id}}", r.PathParams())
	return &u
}

func (r *GetPropertyByIDRequest) Do(ctx context.Context) (GetPropertyByIDResponseBody, error) {
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
