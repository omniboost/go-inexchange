package inexchange

import (
	"context"
	"net/http"
	"net/url"

	"github.com/omniboost/go-inexchange/utils"
)

func (c *Client) NewPatchTravelAgentProfileByIDRequest() PatchTravelAgentProfileByIDRequest {
	r := PatchTravelAgentProfileByIDRequest{
		client: c,
		method: http.MethodPatch,
	}

	r.headers = r.NewHeaders()
	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()

	return r
}

type PatchTravelAgentProfileByIDRequest struct {
	client      *Client
	queryParams *PatchTravelAgentProfileByIDQueryParams
	pathParams  *PatchTravelAgentProfileByIDPathParams
	method      string
	headers     *PatchTravelAgentProfileByIDHeaders
	requestBody PatchTravelAgentProfileByIDRequestBody
}

func (r PatchTravelAgentProfileByIDRequest) NewQueryParams() *PatchTravelAgentProfileByIDQueryParams {
	return &PatchTravelAgentProfileByIDQueryParams{}
}

type PatchTravelAgentProfileByIDQueryParams struct{}

func (p PatchTravelAgentProfileByIDQueryParams) ToURLValues() (url.Values, error) {
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

func (r *PatchTravelAgentProfileByIDRequest) QueryParams() *PatchTravelAgentProfileByIDQueryParams {
	return r.queryParams
}

func (r PatchTravelAgentProfileByIDRequest) NewHeaders() *PatchTravelAgentProfileByIDHeaders {
	return &PatchTravelAgentProfileByIDHeaders{}
}

type PatchTravelAgentProfileByIDHeaders struct {
	PropertyID string `schema:"AC-Property-ID,omitempty"`
	IfMatch    string `schema:"If-Match,omitempty"`
}

func (r *PatchTravelAgentProfileByIDRequest) Headers() *PatchTravelAgentProfileByIDHeaders {
	return r.headers
}

func (r PatchTravelAgentProfileByIDRequest) NewPathParams() *PatchTravelAgentProfileByIDPathParams {
	return &PatchTravelAgentProfileByIDPathParams{}
}

type PatchTravelAgentProfileByIDPathParams struct {
	ID string `schema:"id"`
}

func (p *PatchTravelAgentProfileByIDPathParams) Params() map[string]string {
	return map[string]string{
		"id": p.ID,
	}
}

func (r *PatchTravelAgentProfileByIDRequest) PathParams() *PatchTravelAgentProfileByIDPathParams {
	return r.pathParams
}

func (r *PatchTravelAgentProfileByIDRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *PatchTravelAgentProfileByIDRequest) SetMethod(method string) {
	r.method = method
}

func (r *PatchTravelAgentProfileByIDRequest) Method() string {
	return r.method
}

func (r PatchTravelAgentProfileByIDRequest) NewRequestBody() PatchTravelAgentProfileByIDRequestBody {
	return PatchTravelAgentProfileByIDRequestBody{}
}

type PatchTravelAgentProfileByIDRequestBody PatchTravelAgent

func (r *PatchTravelAgentProfileByIDRequest) RequestBody() *PatchTravelAgentProfileByIDRequestBody {
	return &r.requestBody
}

func (r *PatchTravelAgentProfileByIDRequest) RequestBodyInterface() interface{} {
	return r.requestBody
}

func (r *PatchTravelAgentProfileByIDRequest) SetRequestBody(body PatchTravelAgentProfileByIDRequestBody) {
	r.requestBody = body
}

func (r *PatchTravelAgentProfileByIDRequest) NewResponseBody() *PatchTravelAgentProfileByIDResponseBody {
	return &PatchTravelAgentProfileByIDResponseBody{}
}

type PatchTravelAgentProfileByIDResponseBody struct{}

func (r *PatchTravelAgentProfileByIDRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("api-gateway/profiles/v1/travelagent/{{.id}}", r.PathParams())
	return &u
}

func (r *PatchTravelAgentProfileByIDRequest) Do(ctx context.Context) (PatchTravelAgentProfileByIDResponseBody, error) {
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
