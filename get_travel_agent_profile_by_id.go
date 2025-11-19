package inexchange

import (
	"context"
	"net/http"
	"net/url"

	"github.com/omniboost/go-inexchange/utils"
)

func (c *Client) NewGetTravelAgentProfileByIDRequest() GetTravelAgentProfileByIDRequest {
	return GetTravelAgentProfileByIDRequest{
		client:      c,
		queryParams: c.NewGetTravelAgentProfileByIDQueryParams(),
		pathParams:  c.NewGetTravelAgentProfileByIDPathParams(),
		method:      http.MethodGet,
		headers:     c.NewGetTravelAgentProfileByIDHeaders(),
		requestBody: c.NewGetTravelAgentProfileByIDRequestBody(),
	}
}

type GetTravelAgentProfileByIDRequest struct {
	client      *Client
	queryParams *GetTravelAgentProfileByIDQueryParams
	pathParams  *GetTravelAgentProfileByIDPathParams
	method      string
	headers     *GetTravelAgentProfileByIDHeaders
	requestBody GetTravelAgentProfileByIDRequestBody
}

func (c *Client) NewGetTravelAgentProfileByIDQueryParams() *GetTravelAgentProfileByIDQueryParams {
	return &GetTravelAgentProfileByIDQueryParams{}
}

type GetTravelAgentProfileByIDQueryParams struct {
	Extend CommaSeparatedQueryParam `schema:"extend,omitempty"`
}

func (p GetTravelAgentProfileByIDQueryParams) ToURLValues() (url.Values, error) {
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

func (r *GetTravelAgentProfileByIDRequest) QueryParams() *GetTravelAgentProfileByIDQueryParams {
	return r.queryParams
}

func (c *Client) NewGetTravelAgentProfileByIDHeaders() *GetTravelAgentProfileByIDHeaders {
	return &GetTravelAgentProfileByIDHeaders{}
}

type GetTravelAgentProfileByIDHeaders struct {
	PropertyID string `schema:"AC-Property-ID,omitempty"`
}

func (r *GetTravelAgentProfileByIDRequest) Headers() *GetTravelAgentProfileByIDHeaders {
	return r.headers
}

func (c *Client) NewGetTravelAgentProfileByIDPathParams() *GetTravelAgentProfileByIDPathParams {
	return &GetTravelAgentProfileByIDPathParams{}
}

type GetTravelAgentProfileByIDPathParams struct {
	ID string
}

func (p *GetTravelAgentProfileByIDPathParams) Params() map[string]string {
	return map[string]string{
		"id": p.ID,
	}
}

func (r *GetTravelAgentProfileByIDRequest) PathParams() *GetTravelAgentProfileByIDPathParams {
	return r.pathParams
}

func (r *GetTravelAgentProfileByIDRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *GetTravelAgentProfileByIDRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetTravelAgentProfileByIDRequest) Method() string {
	return r.method
}

func (s *Client) NewGetTravelAgentProfileByIDRequestBody() GetTravelAgentProfileByIDRequestBody {
	return GetTravelAgentProfileByIDRequestBody{}
}

type GetTravelAgentProfileByIDRequestBody struct {
}

func (r *GetTravelAgentProfileByIDRequest) RequestBody() *GetTravelAgentProfileByIDRequestBody {
	return nil
}

func (r *GetTravelAgentProfileByIDRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *GetTravelAgentProfileByIDRequest) SetRequestBody(body GetTravelAgentProfileByIDRequestBody) {
	r.requestBody = body
}

func (r *GetTravelAgentProfileByIDRequest) NewResponseBody() *GetTravelAgentProfileByIDResponseBody {
	return &GetTravelAgentProfileByIDResponseBody{}
}

type GetTravelAgentProfileByIDResponseBody TravelAgentProfile

func (r *GetTravelAgentProfileByIDRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("api-gateway/profiles/v1/travel-agent/{{.id}}", r.PathParams())
	return &u
}

func (r *GetTravelAgentProfileByIDRequest) Do(ctx context.Context) (GetTravelAgentProfileByIDResponseBody, error) {
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
