package inexchange

import (
	"context"
	"net/http"
	"net/url"

	"github.com/omniboost/go-inexchange/utils"
)

func (c *Client) NewGetSubsidiariesRequest() GetSubsidiariesRequest {
	return GetSubsidiariesRequest{
		client:      c,
		queryParams: c.NewGetSubsidiariesQueryParams(),
		pathParams:  c.NewGetSubsidiariesPathParams(),
		method:      http.MethodGet,
		headers:     c.NewGetSubsidiariesHeaders(),
		requestBody: c.NewGetSubsidiariesRequestBody(),
	}
}

type GetSubsidiariesRequest struct {
	client      *Client
	queryParams *GetSubsidiariesQueryParams
	pathParams  *GetSubsidiariesPathParams
	method      string
	headers     *GetSubsidiariesHeaders
	requestBody GetSubsidiariesRequestBody
}

func (c *Client) NewGetSubsidiariesQueryParams() *GetSubsidiariesQueryParams {
	return &GetSubsidiariesQueryParams{}
}

type GetSubsidiariesQueryParams struct {
	StatusCode string `schema:"statusCode,omitempty"`
	ParentID   string `schema:"parentId,omitempty"`
	Recursive  bool   `schema:"recursive,omitempty"`
	Sort       string `schema:"sort,omitempty"`
	PageNumber int32  `schema:"pageNumber,omitempty"`
	PageSize   int32  `schema:"pageSize,omitempty"`
}

func (p GetSubsidiariesQueryParams) ToURLValues() (url.Values, error) {
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

func (r *GetSubsidiariesRequest) QueryParams() *GetSubsidiariesQueryParams {
	return r.queryParams
}

func (c *Client) NewGetSubsidiariesHeaders() *GetSubsidiariesHeaders {
	return &GetSubsidiariesHeaders{}
}

type GetSubsidiariesHeaders struct {
	TenantID string `schema:"AC-Tenant-ID,omitempty"`
}

func (r *GetSubsidiariesRequest) Headers() *GetSubsidiariesHeaders {
	return r.headers
}

func (c *Client) NewGetSubsidiariesPathParams() *GetSubsidiariesPathParams {
	return &GetSubsidiariesPathParams{}
}

type GetSubsidiariesPathParams struct {
}

func (p *GetSubsidiariesPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GetSubsidiariesRequest) PathParams() *GetSubsidiariesPathParams {
	return r.pathParams
}

func (r *GetSubsidiariesRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *GetSubsidiariesRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetSubsidiariesRequest) Method() string {
	return r.method
}

func (s *Client) NewGetSubsidiariesRequestBody() GetSubsidiariesRequestBody {
	return GetSubsidiariesRequestBody{}
}

type GetSubsidiariesRequestBody struct {
}

func (r *GetSubsidiariesRequest) RequestBody() *GetSubsidiariesRequestBody {
	return nil
}

func (r *GetSubsidiariesRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *GetSubsidiariesRequest) SetRequestBody(body GetSubsidiariesRequestBody) {
	r.requestBody = body
}

func (r *GetSubsidiariesRequest) NewResponseBody() *GetSubsidiariesResponseBody {
	return &GetSubsidiariesResponseBody{}
}

type GetSubsidiariesResponseBody []SubsidiaryListItem

func (r *GetSubsidiariesRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("api-gateway/configuration/v1/subsidiaries", r.PathParams())
	return &u
}

func (r *GetSubsidiariesRequest) Do(ctx context.Context) (GetSubsidiariesResponseBody, error) {
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
