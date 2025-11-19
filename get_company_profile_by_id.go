package inexchange

import (
	"context"
	"net/http"
	"net/url"

	"github.com/omniboost/go-inexchange/utils"
)

func (c *Client) NewGetCompanyProfileByIDRequest() GetCompanyProfileByIDRequest {
	return GetCompanyProfileByIDRequest{
		client:      c,
		queryParams: c.NewGetCompanyProfileByIDQueryParams(),
		pathParams:  c.NewGetCompanyProfileByIDPathParams(),
		method:      http.MethodGet,
		headers:     c.NewGetCompanyProfileByIDHeaders(),
		requestBody: c.NewGetCompanyProfileByIDRequestBody(),
	}
}

type GetCompanyProfileByIDRequest struct {
	client      *Client
	queryParams *GetCompanyProfileByIDQueryParams
	pathParams  *GetCompanyProfileByIDPathParams
	method      string
	headers     *GetCompanyProfileByIDHeaders
	requestBody GetCompanyProfileByIDRequestBody
}

func (c *Client) NewGetCompanyProfileByIDQueryParams() *GetCompanyProfileByIDQueryParams {
	return &GetCompanyProfileByIDQueryParams{}
}

type GetCompanyProfileByIDQueryParams struct {
	Extend CommaSeparatedQueryParam `schema:"extend,omitempty"`
}

func (p GetCompanyProfileByIDQueryParams) ToURLValues() (url.Values, error) {
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

func (r *GetCompanyProfileByIDRequest) QueryParams() *GetCompanyProfileByIDQueryParams {
	return r.queryParams
}

func (c *Client) NewGetCompanyProfileByIDHeaders() *GetCompanyProfileByIDHeaders {
	return &GetCompanyProfileByIDHeaders{}
}

type GetCompanyProfileByIDHeaders struct {
	PropertyID string `schema:"AC-Property-ID,omitempty"`
}

func (r *GetCompanyProfileByIDRequest) Headers() *GetCompanyProfileByIDHeaders {
	return r.headers
}

func (c *Client) NewGetCompanyProfileByIDPathParams() *GetCompanyProfileByIDPathParams {
	return &GetCompanyProfileByIDPathParams{}
}

type GetCompanyProfileByIDPathParams struct {
	ID string
}

func (p *GetCompanyProfileByIDPathParams) Params() map[string]string {
	return map[string]string{
		"id": p.ID,
	}
}

func (r *GetCompanyProfileByIDRequest) PathParams() *GetCompanyProfileByIDPathParams {
	return r.pathParams
}

func (r *GetCompanyProfileByIDRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *GetCompanyProfileByIDRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetCompanyProfileByIDRequest) Method() string {
	return r.method
}

func (s *Client) NewGetCompanyProfileByIDRequestBody() GetCompanyProfileByIDRequestBody {
	return GetCompanyProfileByIDRequestBody{}
}

type GetCompanyProfileByIDRequestBody struct {
}

func (r *GetCompanyProfileByIDRequest) RequestBody() *GetCompanyProfileByIDRequestBody {
	return nil
}

func (r *GetCompanyProfileByIDRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *GetCompanyProfileByIDRequest) SetRequestBody(body GetCompanyProfileByIDRequestBody) {
	r.requestBody = body
}

func (r *GetCompanyProfileByIDRequest) NewResponseBody() *GetCompanyProfileByIDResponseBody {
	return &GetCompanyProfileByIDResponseBody{}
}

type GetCompanyProfileByIDResponseBody CompanyProfile

func (r *GetCompanyProfileByIDRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("api-gateway/profiles/v1/company/{{.id}}", r.PathParams())
	return &u
}

func (r *GetCompanyProfileByIDRequest) Do(ctx context.Context) (GetCompanyProfileByIDResponseBody, error) {
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
