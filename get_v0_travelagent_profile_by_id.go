package inexchange

import (
	"context"
	"net/http"
	"net/url"

	"github.com/omniboost/go-inexchange/utils"
)

func (c *Client) NewGetV0TravelAgentProfileByIDRequest() GetV0TravelAgentProfileByIDRequest {
	return GetV0TravelAgentProfileByIDRequest{
		client:      c,
		queryParams: c.NewGetV0TravelAgentProfileByIDQueryParams(),
		pathParams:  c.NewGetV0TravelAgentProfileByIDPathParams(),
		method:      http.MethodGet,
		headers:     c.NewGetV0TravelAgentProfileByIDHeaders(),
		requestBody: c.NewGetV0TravelAgentProfileByIDRequestBody(),
	}
}

type GetV0TravelAgentProfileByIDRequest struct {
	client      *Client
	queryParams *GetV0TravelAgentProfileByIDQueryParams
	pathParams  *GetV0TravelAgentProfileByIDPathParams
	method      string
	headers     *GetV0TravelAgentProfileByIDHeaders
	requestBody GetV0TravelAgentProfileByIDRequestBody
}

func (c *Client) NewGetV0TravelAgentProfileByIDQueryParams() *GetV0TravelAgentProfileByIDQueryParams {
	return &GetV0TravelAgentProfileByIDQueryParams{}
}

type GetV0TravelAgentProfileByIDQueryParams struct {
	Extend []string `schema:"extend,omitempty"`
}

func (p GetV0TravelAgentProfileByIDQueryParams) ToURLValues() (url.Values, error) {
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

func (r *GetV0TravelAgentProfileByIDRequest) QueryParams() *GetV0TravelAgentProfileByIDQueryParams {
	return r.queryParams
}

func (c *Client) NewGetV0TravelAgentProfileByIDHeaders() *GetV0TravelAgentProfileByIDHeaders {
	return &GetV0TravelAgentProfileByIDHeaders{}
}

type GetV0TravelAgentProfileByIDHeaders struct {
	PropertyID string `schema:"AC-Property-ID,omitempty"`
}

func (r *GetV0TravelAgentProfileByIDRequest) Headers() *GetV0TravelAgentProfileByIDHeaders {
	return r.headers
}

func (c *Client) NewGetV0TravelAgentProfileByIDPathParams() *GetV0TravelAgentProfileByIDPathParams {
	return &GetV0TravelAgentProfileByIDPathParams{}
}

type GetV0TravelAgentProfileByIDPathParams struct {
	ID string
}

func (p *GetV0TravelAgentProfileByIDPathParams) Params() map[string]string {
	return map[string]string{
		"id": p.ID,
	}
}

func (r *GetV0TravelAgentProfileByIDRequest) PathParams() *GetV0TravelAgentProfileByIDPathParams {
	return r.pathParams
}

func (r *GetV0TravelAgentProfileByIDRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *GetV0TravelAgentProfileByIDRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetV0TravelAgentProfileByIDRequest) Method() string {
	return r.method
}

func (s *Client) NewGetV0TravelAgentProfileByIDRequestBody() GetV0TravelAgentProfileByIDRequestBody {
	return GetV0TravelAgentProfileByIDRequestBody{}
}

type GetV0TravelAgentProfileByIDRequestBody struct {
}

func (r *GetV0TravelAgentProfileByIDRequest) RequestBody() *GetV0TravelAgentProfileByIDRequestBody {
	return nil
}

func (r *GetV0TravelAgentProfileByIDRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *GetV0TravelAgentProfileByIDRequest) SetRequestBody(body GetV0TravelAgentProfileByIDRequestBody) {
	r.requestBody = body
}

func (r *GetV0TravelAgentProfileByIDRequest) NewResponseBody() *GetV0TravelAgentProfileByIDResponseBody {
	return &GetV0TravelAgentProfileByIDResponseBody{}
}

type GetV0TravelAgentProfileByIDResponseBody TravelAgentProfileV0

func (r *GetV0TravelAgentProfileByIDRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("api-gateway/profiles/v0/travel-agent/{{.id}}", r.PathParams())
	return &u
}

func (r *GetV0TravelAgentProfileByIDRequest) Do(ctx context.Context) (GetV0TravelAgentProfileByIDResponseBody, error) {
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
