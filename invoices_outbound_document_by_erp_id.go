package inexchange

import (
	"context"
	"net/http"
	"net/url"

	"github.com/omniboost/go-inexchange/utils"
)

func (c *Client) NewInvoicesOutboundDocumentByERPIDRequest() InvoicesOutboundDocumentByERPIDRequest {
	return InvoicesOutboundDocumentByERPIDRequest{
		client:      c,
		queryParams: c.NewInvoicesOutboundDocumentByERPIDQueryParams(),
		pathParams:  c.NewInvoicesOutboundDocumentByERPIDPathParams(),
		method:      http.MethodPost,
		headers:     c.NewInvoicesOutboundDocumentByERPIDHeaders(),
		requestBody: c.NewInvoicesOutboundDocumentByERPIDRequestBody(),
	}
}

type InvoicesOutboundDocumentByERPIDRequest struct {
	client      *Client
	queryParams *InvoicesOutboundDocumentByERPIDQueryParams
	pathParams  *InvoicesOutboundDocumentByERPIDPathParams
	method      string
	headers     *InvoicesOutboundDocumentByERPIDHeaders
	requestBody InvoicesOutboundDocumentByERPIDRequestBody
}

func (c *Client) NewInvoicesOutboundDocumentByERPIDQueryParams() *InvoicesOutboundDocumentByERPIDQueryParams {
	return &InvoicesOutboundDocumentByERPIDQueryParams{}
}

type InvoicesOutboundDocumentByERPIDQueryParams struct {
	StatusCode string `schema:"statusCode,omitempty"`
	ParentID   string `schema:"parentId,omitempty"`
	Recursive  bool   `schema:"recursive,omitempty"`
	Sort       string `schema:"sort,omitempty"`
	PageNumber int32  `schema:"pageNumber,omitempty"`
	PageSize   int32  `schema:"pageSize,omitempty"`
}

func (p InvoicesOutboundDocumentByERPIDQueryParams) ToURLValues() (url.Values, error) {
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

func (r *InvoicesOutboundDocumentByERPIDRequest) QueryParams() *InvoicesOutboundDocumentByERPIDQueryParams {
	return r.queryParams
}

func (c *Client) NewInvoicesOutboundDocumentByERPIDHeaders() *InvoicesOutboundDocumentByERPIDHeaders {
	return &InvoicesOutboundDocumentByERPIDHeaders{}
}

type InvoicesOutboundDocumentByERPIDHeaders struct {
	TenantID string `schema:"AC-Tenant-ID,omitempty"`
}

func (r *InvoicesOutboundDocumentByERPIDRequest) Headers() *InvoicesOutboundDocumentByERPIDHeaders {
	return r.headers
}

func (c *Client) NewInvoicesOutboundDocumentByERPIDPathParams() *InvoicesOutboundDocumentByERPIDPathParams {
	return &InvoicesOutboundDocumentByERPIDPathParams{}
}

type InvoicesOutboundDocumentByERPIDPathParams struct {
	DocumentERPID string `schema:"document_erp_id"`
}

func (p *InvoicesOutboundDocumentByERPIDPathParams) Params() map[string]string {
	return map[string]string{
		"document_erp_id": p.DocumentERPID,
	}
}

func (r *InvoicesOutboundDocumentByERPIDRequest) PathParams() *InvoicesOutboundDocumentByERPIDPathParams {
	return r.pathParams
}

func (r *InvoicesOutboundDocumentByERPIDRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *InvoicesOutboundDocumentByERPIDRequest) SetMethod(method string) {
	r.method = method
}

func (r *InvoicesOutboundDocumentByERPIDRequest) Method() string {
	return r.method
}

func (s *Client) NewInvoicesOutboundDocumentByERPIDRequestBody() InvoicesOutboundDocumentByERPIDRequestBody {
	return InvoicesOutboundDocumentByERPIDRequestBody{}
}

type InvoicesOutboundDocumentByERPIDRequestBody struct {
	IncludeFileInfo bool `json:"includeFileInfo,omitempty"`
}

func (r *InvoicesOutboundDocumentByERPIDRequest) RequestBody() *InvoicesOutboundDocumentByERPIDRequestBody {
	return &r.requestBody
}

func (r *InvoicesOutboundDocumentByERPIDRequest) RequestBodyInterface() any {
	return r.requestBody
}

func (r *InvoicesOutboundDocumentByERPIDRequest) SetRequestBody(body InvoicesOutboundDocumentByERPIDRequestBody) {
	r.requestBody = body
}

func (r *InvoicesOutboundDocumentByERPIDRequest) NewResponseBody() *InvoicesOutboundDocumentByERPIDResponseBody {
	return &InvoicesOutboundDocumentByERPIDResponseBody{}
}

type InvoicesOutboundDocumentByERPIDResponseBody OutboundDocument

func (r *InvoicesOutboundDocumentByERPIDRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("/invoices/outbound/byerpid/{{.document_erp_id}}", r.PathParams())
	return &u
}

func (r *InvoicesOutboundDocumentByERPIDRequest) Do(ctx context.Context) (InvoicesOutboundDocumentByERPIDResponseBody, error) {
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
