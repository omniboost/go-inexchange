package inexchange

import (
	"context"
	"net/http"
	"net/url"

	"github.com/omniboost/go-inexchange/utils"
)

func (c *Client) NewGetV0IndividualProfileByIDRequest() GetV0IndividualProfileByIDRequest {
	return GetV0IndividualProfileByIDRequest{
		client:      c,
		queryParams: c.NewGetV0IndividualProfileByIDQueryParams(),
		pathParams:  c.NewGetV0IndividualProfileByIDPathParams(),
		method:      http.MethodGet,
		headers:     c.NewGetV0IndividualProfileByIDHeaders(),
		requestBody: c.NewGetV0IndividualProfileByIDRequestBody(),
	}
}

type GetV0IndividualProfileByIDRequest struct {
	client      *Client
	queryParams *GetV0IndividualProfileByIDQueryParams
	pathParams  *GetV0IndividualProfileByIDPathParams
	method      string
	headers     *GetV0IndividualProfileByIDHeaders
	requestBody GetV0IndividualProfileByIDRequestBody
}

func (c *Client) NewGetV0IndividualProfileByIDQueryParams() *GetV0IndividualProfileByIDQueryParams {
	return &GetV0IndividualProfileByIDQueryParams{}
}

type GetV0IndividualProfileByIDQueryParams struct {
	Extend []string `schema:"extend,omitempty"`
}

func (p GetV0IndividualProfileByIDQueryParams) ToURLValues() (url.Values, error) {
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

func (r *GetV0IndividualProfileByIDRequest) QueryParams() *GetV0IndividualProfileByIDQueryParams {
	return r.queryParams
}

func (c *Client) NewGetV0IndividualProfileByIDHeaders() *GetV0IndividualProfileByIDHeaders {
	return &GetV0IndividualProfileByIDHeaders{}
}

type GetV0IndividualProfileByIDHeaders struct {
	PropertyID string `schema:"AC-Property-ID,omitempty"`
}

func (r *GetV0IndividualProfileByIDRequest) Headers() *GetV0IndividualProfileByIDHeaders {
	return r.headers
}

func (c *Client) NewGetV0IndividualProfileByIDPathParams() *GetV0IndividualProfileByIDPathParams {
	return &GetV0IndividualProfileByIDPathParams{}
}

type GetV0IndividualProfileByIDPathParams struct {
	ID string
}

func (p *GetV0IndividualProfileByIDPathParams) Params() map[string]string {
	return map[string]string{
		"id": p.ID,
	}
}

func (r *GetV0IndividualProfileByIDRequest) PathParams() *GetV0IndividualProfileByIDPathParams {
	return r.pathParams
}

func (r *GetV0IndividualProfileByIDRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *GetV0IndividualProfileByIDRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetV0IndividualProfileByIDRequest) Method() string {
	return r.method
}

func (s *Client) NewGetV0IndividualProfileByIDRequestBody() GetV0IndividualProfileByIDRequestBody {
	return GetV0IndividualProfileByIDRequestBody{}
}

type GetV0IndividualProfileByIDRequestBody struct {
}

func (r *GetV0IndividualProfileByIDRequest) RequestBody() *GetV0IndividualProfileByIDRequestBody {
	return nil
}

func (r *GetV0IndividualProfileByIDRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *GetV0IndividualProfileByIDRequest) SetRequestBody(body GetV0IndividualProfileByIDRequestBody) {
	r.requestBody = body
}

func (r *GetV0IndividualProfileByIDRequest) NewResponseBody() *GetV0IndividualProfileByIDResponseBody {
	return &GetV0IndividualProfileByIDResponseBody{}
}

type GetV0IndividualProfileByIDResponseBody IndividualProfileV0

func (r *GetV0IndividualProfileByIDRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("api-gateway/profiles/v0/individual/{{.id}}", r.PathParams())
	return &u
}

func (r *GetV0IndividualProfileByIDRequest) Do(ctx context.Context) (GetV0IndividualProfileByIDResponseBody, error) {
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
