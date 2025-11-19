package inexchange

import (
	"context"
	"net/http"
	"net/url"

	"github.com/omniboost/go-inexchange/utils"
)

func (c *Client) NewInvoicesOutboundDocumentIDRequest() InvoicesOutboundDocumentIDRequest {
	return InvoicesOutboundDocumentIDRequest{
		client:      c,
		queryParams: c.NewInvoicesOutboundDocumentIDQueryParams(),
		pathParams:  c.NewInvoicesOutboundDocumentIDPathParams(),
		method:      http.MethodPost,
		headers:     c.NewInvoicesOutboundDocumentIDHeaders(),
		requestBody: c.NewInvoicesOutboundDocumentIDRequestBody(),
	}
}

type InvoicesOutboundDocumentIDRequest struct {
	client      *Client
	queryParams *InvoicesOutboundDocumentIDQueryParams
	pathParams  *InvoicesOutboundDocumentIDPathParams
	method      string
	headers     *InvoicesOutboundDocumentIDHeaders
	requestBody InvoicesOutboundDocumentIDRequestBody
}

func (c *Client) NewInvoicesOutboundDocumentIDQueryParams() *InvoicesOutboundDocumentIDQueryParams {
	return &InvoicesOutboundDocumentIDQueryParams{}
}

type InvoicesOutboundDocumentIDQueryParams struct {
	StatusCode string `schema:"statusCode,omitempty"`
	ParentID   string `schema:"parentId,omitempty"`
	Recursive  bool   `schema:"recursive,omitempty"`
	Sort       string `schema:"sort,omitempty"`
	PageNumber int32  `schema:"pageNumber,omitempty"`
	PageSize   int32  `schema:"pageSize,omitempty"`
}

func (p InvoicesOutboundDocumentIDQueryParams) ToURLValues() (url.Values, error) {
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

func (r *InvoicesOutboundDocumentIDRequest) QueryParams() *InvoicesOutboundDocumentIDQueryParams {
	return r.queryParams
}

func (c *Client) NewInvoicesOutboundDocumentIDHeaders() *InvoicesOutboundDocumentIDHeaders {
	return &InvoicesOutboundDocumentIDHeaders{}
}

type InvoicesOutboundDocumentIDHeaders struct {
	TenantID string `schema:"AC-Tenant-ID,omitempty"`
}

func (r *InvoicesOutboundDocumentIDRequest) Headers() *InvoicesOutboundDocumentIDHeaders {
	return r.headers
}

func (c *Client) NewInvoicesOutboundDocumentIDPathParams() *InvoicesOutboundDocumentIDPathParams {
	return &InvoicesOutboundDocumentIDPathParams{}
}

type InvoicesOutboundDocumentIDPathParams struct {
	DocumentID string `schema:"document_id"`
}

func (p *InvoicesOutboundDocumentIDPathParams) Params() map[string]string {
	return map[string]string{
		"document_id": p.DocumentID,
	}
}

func (r *InvoicesOutboundDocumentIDRequest) PathParams() *InvoicesOutboundDocumentIDPathParams {
	return r.pathParams
}

func (r *InvoicesOutboundDocumentIDRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *InvoicesOutboundDocumentIDRequest) SetMethod(method string) {
	r.method = method
}

func (r *InvoicesOutboundDocumentIDRequest) Method() string {
	return r.method
}

func (s *Client) NewInvoicesOutboundDocumentIDRequestBody() InvoicesOutboundDocumentIDRequestBody {
	return InvoicesOutboundDocumentIDRequestBody{}
}

type InvoicesOutboundDocumentIDRequestBody struct {
	IncludeFileInfo bool `json:"includeFileInfo,omitempty"`
}

func (r *InvoicesOutboundDocumentIDRequest) RequestBody() *InvoicesOutboundDocumentIDRequestBody {
	return &r.requestBody
}

func (r *InvoicesOutboundDocumentIDRequest) RequestBodyInterface() any {
	return r.requestBody
}

func (r *InvoicesOutboundDocumentIDRequest) SetRequestBody(body InvoicesOutboundDocumentIDRequestBody) {
	r.requestBody = body
}

func (r *InvoicesOutboundDocumentIDRequest) NewResponseBody() *InvoicesOutboundDocumentIDResponseBody {
	return &InvoicesOutboundDocumentIDResponseBody{}
}

type InvoicesOutboundDocumentIDResponseBody Document

func (r *InvoicesOutboundDocumentIDRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("/invoices/outbound/{{.document_id}}", r.PathParams())
	return &u
}

func (r *InvoicesOutboundDocumentIDRequest) Do(ctx context.Context) (InvoicesOutboundDocumentIDResponseBody, error) {
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

	responseBody := r.NewResponseBody()
	_, err = r.client.Do(req, responseBody)
	return *responseBody, err
}
