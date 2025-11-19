package inexchange

import (
	"context"
	"net/http"
	"net/url"

	"github.com/omniboost/go-inexchange/utils"
)

func (c *Client) NewGetPropertyLinkedZonesRequest() GetPropertyLinkedZonesRequest {
	return GetPropertyLinkedZonesRequest{
		client:      c,
		queryParams: c.NewGetPropertyLinkedZonesQueryParams(),
		pathParams:  c.NewGetPropertyLinkedZonesPathParams(),
		method:      http.MethodGet,
		headers:     c.NewGetPropertyLinkedZonesHeaders(),
		requestBody: c.NewGetPropertyLinkedZonesRequestBody(),
	}
}

type GetPropertyLinkedZonesRequest struct {
	client      *Client
	queryParams *GetPropertyLinkedZonesQueryParams
	pathParams  *GetPropertyLinkedZonesPathParams
	method      string
	headers     *GetPropertyLinkedZonesHeaders
	requestBody GetPropertyLinkedZonesRequestBody
}

func (c *Client) NewGetPropertyLinkedZonesQueryParams() *GetPropertyLinkedZonesQueryParams {
	return &GetPropertyLinkedZonesQueryParams{}
}

type GetPropertyLinkedZonesQueryParams struct{}

func (p GetPropertyLinkedZonesQueryParams) ToURLValues() (url.Values, error) {
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

func (r *GetPropertyLinkedZonesRequest) QueryParams() *GetPropertyLinkedZonesQueryParams {
	return r.queryParams
}

func (c *Client) NewGetPropertyLinkedZonesHeaders() *GetPropertyLinkedZonesHeaders {
	return &GetPropertyLinkedZonesHeaders{}
}

type GetPropertyLinkedZonesHeaders struct {
	TenantID string `schema:"AC-Tenant-ID,omitempty"`
}

func (r *GetPropertyLinkedZonesRequest) Headers() *GetPropertyLinkedZonesHeaders {
	return r.headers
}

func (c *Client) NewGetPropertyLinkedZonesPathParams() *GetPropertyLinkedZonesPathParams {
	return &GetPropertyLinkedZonesPathParams{}
}

type GetPropertyLinkedZonesPathParams struct {
	ID string `schema:"id"`
}

func (p *GetPropertyLinkedZonesPathParams) Params() map[string]string {
	return map[string]string{
		"id": p.ID,
	}
}

func (r *GetPropertyLinkedZonesRequest) PathParams() *GetPropertyLinkedZonesPathParams {
	return r.pathParams
}

func (r *GetPropertyLinkedZonesRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *GetPropertyLinkedZonesRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetPropertyLinkedZonesRequest) Method() string {
	return r.method
}

func (s *Client) NewGetPropertyLinkedZonesRequestBody() GetPropertyLinkedZonesRequestBody {
	return GetPropertyLinkedZonesRequestBody{}
}

type GetPropertyLinkedZonesRequestBody struct {
}

func (r *GetPropertyLinkedZonesRequest) RequestBody() *GetPropertyLinkedZonesRequestBody {
	return nil
}

func (r *GetPropertyLinkedZonesRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *GetPropertyLinkedZonesRequest) SetRequestBody(body GetPropertyLinkedZonesRequestBody) {
	r.requestBody = body
}

func (r *GetPropertyLinkedZonesRequest) NewResponseBody() *GetPropertyLinkedZonesResponseBody {
	return &GetPropertyLinkedZonesResponseBody{}
}

type GetPropertyLinkedZonesResponseBody []Zone

func (r *GetPropertyLinkedZonesRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("api-gateway/configuration/v1/properties/{{.id}}/linked-zones", r.PathParams())
	return &u
}

func (r *GetPropertyLinkedZonesRequest) Do(ctx context.Context) (GetPropertyLinkedZonesResponseBody, error) {
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
