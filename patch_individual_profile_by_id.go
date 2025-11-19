package inexchange

import (
	"context"
	"net/http"
	"net/url"

	"github.com/omniboost/go-inexchange/utils"
)

func (c *Client) NewPatchIndividualProfileByIDRequest() PatchIndividualProfileByIDRequest {
	r := PatchIndividualProfileByIDRequest{
		client: c,
		method: http.MethodPatch,
	}

	r.headers = r.NewHeaders()
	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()

	return r
}

type PatchIndividualProfileByIDRequest struct {
	client      *Client
	queryParams *PatchIndividualProfileByIDQueryParams
	pathParams  *PatchIndividualProfileByIDPathParams
	method      string
	headers     *PatchIndividualProfileByIDHeaders
	requestBody PatchIndividualProfileByIDRequestBody
}

func (r PatchIndividualProfileByIDRequest) NewQueryParams() *PatchIndividualProfileByIDQueryParams {
	return &PatchIndividualProfileByIDQueryParams{}
}

type PatchIndividualProfileByIDQueryParams struct{}

func (p PatchIndividualProfileByIDQueryParams) ToURLValues() (url.Values, error) {
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

func (r *PatchIndividualProfileByIDRequest) QueryParams() *PatchIndividualProfileByIDQueryParams {
	return r.queryParams
}

func (r PatchIndividualProfileByIDRequest) NewHeaders() *PatchIndividualProfileByIDHeaders {
	return &PatchIndividualProfileByIDHeaders{}
}

type PatchIndividualProfileByIDHeaders struct {
	PropertyID string `schema:"AC-Property-ID,omitempty"`
	IfMatch    string `schema:"If-Match,omitempty"`
}

func (r *PatchIndividualProfileByIDRequest) Headers() *PatchIndividualProfileByIDHeaders {
	return r.headers
}

func (r PatchIndividualProfileByIDRequest) NewPathParams() *PatchIndividualProfileByIDPathParams {
	return &PatchIndividualProfileByIDPathParams{}
}

type PatchIndividualProfileByIDPathParams struct {
	ID string `schema:"id"`
}

func (p *PatchIndividualProfileByIDPathParams) Params() map[string]string {
	return map[string]string{
		"id": p.ID,
	}
}

func (r *PatchIndividualProfileByIDRequest) PathParams() *PatchIndividualProfileByIDPathParams {
	return r.pathParams
}

func (r *PatchIndividualProfileByIDRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *PatchIndividualProfileByIDRequest) SetMethod(method string) {
	r.method = method
}

func (r *PatchIndividualProfileByIDRequest) Method() string {
	return r.method
}

func (r PatchIndividualProfileByIDRequest) NewRequestBody() PatchIndividualProfileByIDRequestBody {
	return PatchIndividualProfileByIDRequestBody{}
}

type PatchIndividualProfileByIDRequestBody PatchIndividual

func (r *PatchIndividualProfileByIDRequest) RequestBody() *PatchIndividualProfileByIDRequestBody {
	return &r.requestBody
}

func (r *PatchIndividualProfileByIDRequest) RequestBodyInterface() interface{} {
	return r.requestBody
}

func (r *PatchIndividualProfileByIDRequest) SetRequestBody(body PatchIndividualProfileByIDRequestBody) {
	r.requestBody = body
}

func (r *PatchIndividualProfileByIDRequest) NewResponseBody() *PatchIndividualProfileByIDResponseBody {
	return &PatchIndividualProfileByIDResponseBody{}
}

type PatchIndividualProfileByIDResponseBody struct{}

func (r *PatchIndividualProfileByIDRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("api-gateway/profiles/v1/individual/{{.id}}", r.PathParams())
	return &u
}

func (r *PatchIndividualProfileByIDRequest) Do(ctx context.Context) (PatchIndividualProfileByIDResponseBody, error) {
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
