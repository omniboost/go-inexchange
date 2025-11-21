package inexchange

import (
	"context"
	"net/http"
	"net/url"

	"github.com/omniboost/go-inexchange/utils"
)

func (c *Client) NewDocumentsRequest() DocumentsRequest {
	return DocumentsRequest{
		client:      c,
		queryParams: c.NewDocumentsQueryParams(),
		pathParams:  c.NewDocumentsPathParams(),
		method:      http.MethodPost,
		headers:     c.NewDocumentsHeaders(),
		requestBody: c.NewDocumentsRequestBody(),
	}
}

type DocumentsRequest struct {
	client      *Client
	queryParams *DocumentsQueryParams
	pathParams  *DocumentsPathParams
	method      string
	headers     *DocumentsHeaders
	requestBody DocumentsRequestBody
}

func (c *Client) NewDocumentsQueryParams() *DocumentsQueryParams {
	return &DocumentsQueryParams{}
}

type DocumentsQueryParams struct {
}

func (p DocumentsQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	encoder.RegisterEncoder(Date{}, utils.EncodeSchemaMarshaler)
	encoder.RegisterEncoder(DateTime{}, utils.EncodeSchemaMarshaler)
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *DocumentsRequest) QueryParams() *DocumentsQueryParams {
	return r.queryParams
}

func (c *Client) NewDocumentsHeaders() *DocumentsHeaders {
	return &DocumentsHeaders{}
}

type DocumentsHeaders struct {
	TenantID string `schema:"AC-Tenant-ID,omitempty"`
}

func (r *DocumentsRequest) Headers() *DocumentsHeaders {
	return r.headers
}

func (c *Client) NewDocumentsPathParams() *DocumentsPathParams {
	return &DocumentsPathParams{}
}

type DocumentsPathParams struct{}

func (p *DocumentsPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *DocumentsRequest) PathParams() *DocumentsPathParams {
	return r.pathParams
}

func (r *DocumentsRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *DocumentsRequest) SetMethod(method string) {
	r.method = method
}

func (r *DocumentsRequest) Method() string {
	return r.method
}

func (s *Client) NewDocumentsRequestBody() DocumentsRequestBody {
	return DocumentsRequestBody{}
}

type DocumentsRequestBody struct {
	File string `json:"Filerrrr,omitempty"`
}

func (r *DocumentsRequest) RequestBody() *DocumentsRequestBody {
	return &r.requestBody
}

func (r *DocumentsRequest) RequestBodyInterface() any {
	return r.requestBody
}

func (r *DocumentsRequest) SetRequestBody(body DocumentsRequestBody) {
	r.requestBody = body
}

func (r *DocumentsRequest) NewResponseBody() *DocumentsResponseBody {
	return &DocumentsResponseBody{}
}

type DocumentsResponseBody Document

func (r *DocumentsRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("/documents", r.PathParams())
	return &u
}

func (r *DocumentsRequest) Do(ctx context.Context) (DocumentsResponseBody, error) {
	// Create http request
	req, err := r.client.NewRequest(ctx, r)
	if err != nil {
		return *r.NewResponseBody(), err
	}

	req.Header.Set("Content-Disposition", `attachment; name="File"; filename="Invoice_0072: 58130.xml"`)
	req.Header.Set("Content-Type", "application/xml")

	// Process query parameters
	err = utils.AddQueryParamsToRequest(r.QueryParams(), req, false)
	if err != nil {
		return *r.NewResponseBody(), err
	}

	responseBody := r.NewResponseBody()
	_, err = r.client.Do(req, responseBody)
	return *responseBody, err
}
