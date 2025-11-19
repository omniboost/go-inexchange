package inexchange

import (
	"context"
	"net/http"
	"net/url"

	"github.com/omniboost/go-inexchange/utils"
)

func (c *Client) NewPostExternalSystemRequest() PostExternalSystemRequest {
	r := PostExternalSystemRequest{
		client: c,
		method: http.MethodPost,
	}

	r.headers = r.NewHeaders()
	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()

	return r
}

type PostExternalSystemRequest struct {
	client      *Client
	queryParams *PostExternalSystemQueryParams
	pathParams  *PostExternalSystemPathParams
	method      string
	headers     *PostExternalSystemHeaders
	requestBody PostExternalSystemRequestBody
}

func (r PostExternalSystemRequest) NewQueryParams() *PostExternalSystemQueryParams {
	return &PostExternalSystemQueryParams{}
}

type PostExternalSystemQueryParams struct{}

func (p PostExternalSystemQueryParams) ToURLValues() (url.Values, error) {
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

func (r *PostExternalSystemRequest) QueryParams() *PostExternalSystemQueryParams {
	return r.queryParams
}

func (r PostExternalSystemRequest) NewHeaders() *PostExternalSystemHeaders {
	return &PostExternalSystemHeaders{}
}

type PostExternalSystemHeaders struct{}

func (r *PostExternalSystemRequest) Headers() *PostExternalSystemHeaders {
	return r.headers
}

func (r PostExternalSystemRequest) NewPathParams() *PostExternalSystemPathParams {
	return &PostExternalSystemPathParams{}
}

type PostExternalSystemPathParams struct{}

func (p *PostExternalSystemPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *PostExternalSystemRequest) PathParams() *PostExternalSystemPathParams {
	return r.pathParams
}

func (r *PostExternalSystemRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *PostExternalSystemRequest) SetMethod(method string) {
	r.method = method
}

func (r *PostExternalSystemRequest) Method() string {
	return r.method
}

func (r PostExternalSystemRequest) NewRequestBody() PostExternalSystemRequestBody {
	return PostExternalSystemRequestBody{}
}

type PostExternalSystemRequestBody ExternalSystem

func (r *PostExternalSystemRequest) RequestBody() *PostExternalSystemRequestBody {
	return &r.requestBody
}

func (r *PostExternalSystemRequest) RequestBodyInterface() interface{} {
	return r.requestBody
}

func (r *PostExternalSystemRequest) SetRequestBody(body PostExternalSystemRequestBody) {
	r.requestBody = body
}

func (r *PostExternalSystemRequest) NewResponseBody() *PostExternalSystemResponseBody {
	return &PostExternalSystemResponseBody{}
}

type PostExternalSystemResponseBody struct{}

func (r *PostExternalSystemRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("api-gateway/configuration/v1/entities/ExternalSystem", r.PathParams())
	return &u
}

func (r *PostExternalSystemRequest) Do(ctx context.Context) (PostExternalSystemResponseBody, error) {
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
