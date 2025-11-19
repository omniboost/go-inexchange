package inexchange

import (
	"context"
	"net/http"
	"net/url"

	"github.com/omniboost/go-inexchange/utils"
)

func (c *Client) NewPatchCompanyProfileByIDRequest() PatchCompanyProfileByIDRequest {
	r := PatchCompanyProfileByIDRequest{
		client: c,
		method: http.MethodPatch,
	}

	r.headers = r.NewHeaders()
	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()

	return r
}

type PatchCompanyProfileByIDRequest struct {
	client      *Client
	queryParams *PatchCompanyProfileByIDQueryParams
	pathParams  *PatchCompanyProfileByIDPathParams
	method      string
	headers     *PatchCompanyProfileByIDHeaders
	requestBody PatchCompanyProfileByIDRequestBody
}

func (r PatchCompanyProfileByIDRequest) NewQueryParams() *PatchCompanyProfileByIDQueryParams {
	return &PatchCompanyProfileByIDQueryParams{}
}

type PatchCompanyProfileByIDQueryParams struct{}

func (p PatchCompanyProfileByIDQueryParams) ToURLValues() (url.Values, error) {
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

func (r *PatchCompanyProfileByIDRequest) QueryParams() *PatchCompanyProfileByIDQueryParams {
	return r.queryParams
}

func (r PatchCompanyProfileByIDRequest) NewHeaders() *PatchCompanyProfileByIDHeaders {
	return &PatchCompanyProfileByIDHeaders{}
}

type PatchCompanyProfileByIDHeaders struct {
	PropertyID string `schema:"AC-Property-ID,omitempty"`
	IfMatch    string `schema:"If-Match,omitempty"`
}

func (r *PatchCompanyProfileByIDRequest) Headers() *PatchCompanyProfileByIDHeaders {
	return r.headers
}

func (r PatchCompanyProfileByIDRequest) NewPathParams() *PatchCompanyProfileByIDPathParams {
	return &PatchCompanyProfileByIDPathParams{}
}

type PatchCompanyProfileByIDPathParams struct {
	ID string `schema:"id"`
}

func (p *PatchCompanyProfileByIDPathParams) Params() map[string]string {
	return map[string]string{
		"id": p.ID,
	}
}

func (r *PatchCompanyProfileByIDRequest) PathParams() *PatchCompanyProfileByIDPathParams {
	return r.pathParams
}

func (r *PatchCompanyProfileByIDRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *PatchCompanyProfileByIDRequest) SetMethod(method string) {
	r.method = method
}

func (r *PatchCompanyProfileByIDRequest) Method() string {
	return r.method
}

func (r PatchCompanyProfileByIDRequest) NewRequestBody() PatchCompanyProfileByIDRequestBody {
	return PatchCompanyProfileByIDRequestBody{}
}

type PatchCompanyProfileByIDRequestBody PatchCompany

func (r *PatchCompanyProfileByIDRequest) RequestBody() *PatchCompanyProfileByIDRequestBody {
	return &r.requestBody
}

func (r *PatchCompanyProfileByIDRequest) RequestBodyInterface() interface{} {
	return r.requestBody
}

func (r *PatchCompanyProfileByIDRequest) SetRequestBody(body PatchCompanyProfileByIDRequestBody) {
	r.requestBody = body
}

func (r *PatchCompanyProfileByIDRequest) NewResponseBody() *PatchCompanyProfileByIDResponseBody {
	return &PatchCompanyProfileByIDResponseBody{}
}

type PatchCompanyProfileByIDResponseBody struct{}

func (r *PatchCompanyProfileByIDRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("api-gateway/profiles/v1/company/{{.id}}", r.PathParams())
	return &u
}

func (r *PatchCompanyProfileByIDRequest) Do(ctx context.Context) (PatchCompanyProfileByIDResponseBody, error) {
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
	req.Header.Add("If-Match", r.Headers().IfMatch)

	responseBody := r.NewResponseBody()
	_, err = r.client.Do(req, responseBody)
	return *responseBody, err
}
