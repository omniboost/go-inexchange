package inexchange

import (
	"context"
	"net/http"
	"net/url"

	"github.com/omniboost/go-inexchange/utils"
)

func (c *Client) NewGetPropertyVisibleZonesRequest() GetPropertyVisibleZonesRequest {
	return GetPropertyVisibleZonesRequest{
		client:      c,
		queryParams: c.NewGetPropertyVisibleZonesQueryParams(),
		pathParams:  c.NewGetPropertyVisibleZonesPathParams(),
		method:      http.MethodGet,
		headers:     c.NewGetPropertyVisibleZonesHeaders(),
		requestBody: c.NewGetPropertyVisibleZonesRequestBody(),
	}
}

type GetPropertyVisibleZonesRequest struct {
	client      *Client
	queryParams *GetPropertyVisibleZonesQueryParams
	pathParams  *GetPropertyVisibleZonesPathParams
	method      string
	headers     *GetPropertyVisibleZonesHeaders
	requestBody GetPropertyVisibleZonesRequestBody
}

func (c *Client) NewGetPropertyVisibleZonesQueryParams() *GetPropertyVisibleZonesQueryParams {
	return &GetPropertyVisibleZonesQueryParams{}
}

type GetPropertyVisibleZonesQueryParams struct{}

func (p GetPropertyVisibleZonesQueryParams) ToURLValues() (url.Values, error) {
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

func (r *GetPropertyVisibleZonesRequest) QueryParams() *GetPropertyVisibleZonesQueryParams {
	return r.queryParams
}

func (c *Client) NewGetPropertyVisibleZonesHeaders() *GetPropertyVisibleZonesHeaders {
	return &GetPropertyVisibleZonesHeaders{}
}

type GetPropertyVisibleZonesHeaders struct {
	TenantID string `schema:"AC-Tenant-ID,omitempty"`
}

func (r *GetPropertyVisibleZonesRequest) Headers() *GetPropertyVisibleZonesHeaders {
	return r.headers
}

func (c *Client) NewGetPropertyVisibleZonesPathParams() *GetPropertyVisibleZonesPathParams {
	return &GetPropertyVisibleZonesPathParams{}
}

type GetPropertyVisibleZonesPathParams struct {
	ID string `schema:"id"`
}

func (p *GetPropertyVisibleZonesPathParams) Params() map[string]string {
	return map[string]string{
		"id": p.ID,
	}
}

func (r *GetPropertyVisibleZonesRequest) PathParams() *GetPropertyVisibleZonesPathParams {
	return r.pathParams
}

func (r *GetPropertyVisibleZonesRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *GetPropertyVisibleZonesRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetPropertyVisibleZonesRequest) Method() string {
	return r.method
}

func (s *Client) NewGetPropertyVisibleZonesRequestBody() GetPropertyVisibleZonesRequestBody {
	return GetPropertyVisibleZonesRequestBody{}
}

type GetPropertyVisibleZonesRequestBody struct {
}

func (r *GetPropertyVisibleZonesRequest) RequestBody() *GetPropertyVisibleZonesRequestBody {
	return nil
}

func (r *GetPropertyVisibleZonesRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *GetPropertyVisibleZonesRequest) SetRequestBody(body GetPropertyVisibleZonesRequestBody) {
	r.requestBody = body
}

func (r *GetPropertyVisibleZonesRequest) NewResponseBody() *GetPropertyVisibleZonesResponseBody {
	return &GetPropertyVisibleZonesResponseBody{}
}

type GetPropertyVisibleZonesResponseBody []Zone

func (r *GetPropertyVisibleZonesRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("api-gateway/configuration/v1/properties/{{.id}}/visible-zones", r.PathParams())
	return &u
}

func (r *GetPropertyVisibleZonesRequest) Do(ctx context.Context) (GetPropertyVisibleZonesResponseBody, error) {
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

	responseBody := r.NewResponseBody()
	_, err = r.client.Do(req, responseBody)
	return *responseBody, err
}
