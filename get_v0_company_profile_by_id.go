package inexchange

import (
	"context"
	"net/http"
	"net/url"

	"github.com/omniboost/go-inexchange/utils"
)

func (c *Client) NewGetV0CompanyProfileByIDRequest() GetV0CompanyProfileByIDRequest {
	return GetV0CompanyProfileByIDRequest{
		client:      c,
		queryParams: c.NewGetV0CompanyProfileByIDQueryParams(),
		pathParams:  c.NewGetV0CompanyProfileByIDPathParams(),
		method:      http.MethodGet,
		headers:     c.NewGetV0CompanyProfileByIDHeaders(),
		requestBody: c.NewGetV0CompanyProfileByIDRequestBody(),
	}
}

type GetV0CompanyProfileByIDRequest struct {
	client      *Client
	queryParams *GetV0CompanyProfileByIDQueryParams
	pathParams  *GetV0CompanyProfileByIDPathParams
	method      string
	headers     *GetV0CompanyProfileByIDHeaders
	requestBody GetV0CompanyProfileByIDRequestBody
}

func (c *Client) NewGetV0CompanyProfileByIDQueryParams() *GetV0CompanyProfileByIDQueryParams {
	return &GetV0CompanyProfileByIDQueryParams{}
}

type GetV0CompanyProfileByIDQueryParams struct {
	Extend CommaSeparatedQueryParam `schema:"extend,omitempty"`
}

func (p GetV0CompanyProfileByIDQueryParams) ToURLValues() (url.Values, error) {
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

func (r *GetV0CompanyProfileByIDRequest) QueryParams() *GetV0CompanyProfileByIDQueryParams {
	return r.queryParams
}

func (c *Client) NewGetV0CompanyProfileByIDHeaders() *GetV0CompanyProfileByIDHeaders {
	return &GetV0CompanyProfileByIDHeaders{}
}

type GetV0CompanyProfileByIDHeaders struct {
	PropertyID string `schema:"AC-Property-ID,omitempty"`
}

func (r *GetV0CompanyProfileByIDRequest) Headers() *GetV0CompanyProfileByIDHeaders {
	return r.headers
}

func (c *Client) NewGetV0CompanyProfileByIDPathParams() *GetV0CompanyProfileByIDPathParams {
	return &GetV0CompanyProfileByIDPathParams{}
}

type GetV0CompanyProfileByIDPathParams struct {
	ID string
}

func (p *GetV0CompanyProfileByIDPathParams) Params() map[string]string {
	return map[string]string{
		"id": p.ID,
	}
}

func (r *GetV0CompanyProfileByIDRequest) PathParams() *GetV0CompanyProfileByIDPathParams {
	return r.pathParams
}

func (r *GetV0CompanyProfileByIDRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *GetV0CompanyProfileByIDRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetV0CompanyProfileByIDRequest) Method() string {
	return r.method
}

func (s *Client) NewGetV0CompanyProfileByIDRequestBody() GetV0CompanyProfileByIDRequestBody {
	return GetV0CompanyProfileByIDRequestBody{}
}

type GetV0CompanyProfileByIDRequestBody struct {
}

func (r *GetV0CompanyProfileByIDRequest) RequestBody() *GetV0CompanyProfileByIDRequestBody {
	return nil
}

func (r *GetV0CompanyProfileByIDRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *GetV0CompanyProfileByIDRequest) SetRequestBody(body GetV0CompanyProfileByIDRequestBody) {
	r.requestBody = body
}

func (r *GetV0CompanyProfileByIDRequest) NewResponseBody() *GetV0CompanyProfileByIDResponseBody {
	return &GetV0CompanyProfileByIDResponseBody{}
}

type GetV0CompanyProfileByIDResponseBody CompanyProfileV0

func (r *GetV0CompanyProfileByIDRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("api-gateway/profiles/v0/company/{{.id}}", r.PathParams())
	return &u
}

func (r *GetV0CompanyProfileByIDRequest) Do(ctx context.Context) (GetV0CompanyProfileByIDResponseBody, error) {
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
