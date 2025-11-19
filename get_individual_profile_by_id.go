package inexchange

import (
	"context"
	"net/http"
	"net/url"

	"github.com/omniboost/go-inexchange/utils"
)

func (c *Client) NewGetIndividualProfileByIDRequest() GetIndividualProfileByIDRequest {
	return GetIndividualProfileByIDRequest{
		client:      c,
		queryParams: c.NewGetIndividualProfileByIDQueryParams(),
		pathParams:  c.NewGetIndividualProfileByIDPathParams(),
		method:      http.MethodGet,
		headers:     c.NewGetIndividualProfileByIDHeaders(),
		requestBody: c.NewGetIndividualProfileByIDRequestBody(),
	}
}

type GetIndividualProfileByIDRequest struct {
	client      *Client
	queryParams *GetIndividualProfileByIDQueryParams
	pathParams  *GetIndividualProfileByIDPathParams
	method      string
	headers     *GetIndividualProfileByIDHeaders
	requestBody GetIndividualProfileByIDRequestBody
}

func (c *Client) NewGetIndividualProfileByIDQueryParams() *GetIndividualProfileByIDQueryParams {
	return &GetIndividualProfileByIDQueryParams{}
}

type GetIndividualProfileByIDQueryParams struct {
	Extend CommaSeparatedQueryParam `schema:"extend,omitempty"`
}

func (p GetIndividualProfileByIDQueryParams) ToURLValues() (url.Values, error) {
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

func (r *GetIndividualProfileByIDRequest) QueryParams() *GetIndividualProfileByIDQueryParams {
	return r.queryParams
}

func (c *Client) NewGetIndividualProfileByIDHeaders() *GetIndividualProfileByIDHeaders {
	return &GetIndividualProfileByIDHeaders{}
}

type GetIndividualProfileByIDHeaders struct {
	PropertyID string `schema:"AC-Property-ID,omitempty"`
}

func (r *GetIndividualProfileByIDRequest) Headers() *GetIndividualProfileByIDHeaders {
	return r.headers
}

func (c *Client) NewGetIndividualProfileByIDPathParams() *GetIndividualProfileByIDPathParams {
	return &GetIndividualProfileByIDPathParams{}
}

type GetIndividualProfileByIDPathParams struct {
	ID string
}

func (p *GetIndividualProfileByIDPathParams) Params() map[string]string {
	return map[string]string{
		"id": p.ID,
	}
}

func (r *GetIndividualProfileByIDRequest) PathParams() *GetIndividualProfileByIDPathParams {
	return r.pathParams
}

func (r *GetIndividualProfileByIDRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *GetIndividualProfileByIDRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetIndividualProfileByIDRequest) Method() string {
	return r.method
}

func (s *Client) NewGetIndividualProfileByIDRequestBody() GetIndividualProfileByIDRequestBody {
	return GetIndividualProfileByIDRequestBody{}
}

type GetIndividualProfileByIDRequestBody struct {
}

func (r *GetIndividualProfileByIDRequest) RequestBody() *GetIndividualProfileByIDRequestBody {
	return nil
}

func (r *GetIndividualProfileByIDRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *GetIndividualProfileByIDRequest) SetRequestBody(body GetIndividualProfileByIDRequestBody) {
	r.requestBody = body
}

func (r *GetIndividualProfileByIDRequest) NewResponseBody() *GetIndividualProfileByIDResponseBody {
	return &GetIndividualProfileByIDResponseBody{}
}

type GetIndividualProfileByIDResponseBody IndividualProfile

func (r *GetIndividualProfileByIDRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("api-gateway/profiles/v1/individual/{{.id}}", r.PathParams())
	return &u
}

func (r *GetIndividualProfileByIDRequest) Do(ctx context.Context) (GetIndividualProfileByIDResponseBody, error) {
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
