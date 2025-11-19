package inexchange

import (
	"context"
	"io"
	"mime"
	"net/http"
	"net/url"
	"strings"

	"github.com/omniboost/go-inexchange/utils"
)

func (c *Client) NewGetCashieringFinalInvoiceContentByFolioIDRequest() GetCashieringFinalInvoiceContentByFolioIDRequest {
	return GetCashieringFinalInvoiceContentByFolioIDRequest{
		client:      c,
		queryParams: c.NewGetCashieringFinalInvoiceContentByFolioIDQueryParams(),
		pathParams:  c.NewGetCashieringFinalInvoiceContentByFolioIDPathParams(),
		method:      http.MethodGet,
		headers:     c.NewGetCashieringFinalInvoiceContentByFolioIDHeaders(),
		requestBody: c.NewGetCashieringFinalInvoiceContentByFolioIDRequestBody(),
	}
}

type GetCashieringFinalInvoiceContentByFolioIDRequest struct {
	client      *Client
	queryParams *GetCashieringFinalInvoiceContentByFolioIDQueryParams
	pathParams  *GetCashieringFinalInvoiceContentByFolioIDPathParams
	method      string
	headers     *GetCashieringFinalInvoiceContentByFolioIDHeaders
	requestBody GetCashieringFinalInvoiceContentByFolioIDRequestBody
}

func (c *Client) NewGetCashieringFinalInvoiceContentByFolioIDQueryParams() *GetCashieringFinalInvoiceContentByFolioIDQueryParams {
	return &GetCashieringFinalInvoiceContentByFolioIDQueryParams{}
}

type GetCashieringFinalInvoiceContentByFolioIDQueryParams struct {
	Extend string `schema:"extend,omitempty"`
}

func (p GetCashieringFinalInvoiceContentByFolioIDQueryParams) ToURLValues() (url.Values, error) {
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

func (r *GetCashieringFinalInvoiceContentByFolioIDRequest) QueryParams() *GetCashieringFinalInvoiceContentByFolioIDQueryParams {
	return r.queryParams
}

func (c *Client) NewGetCashieringFinalInvoiceContentByFolioIDHeaders() *GetCashieringFinalInvoiceContentByFolioIDHeaders {
	return &GetCashieringFinalInvoiceContentByFolioIDHeaders{}
}

type GetCashieringFinalInvoiceContentByFolioIDHeaders struct {
	TenantID   string `schema:"AC-Tenant-ID,omitempty"`
	PropertyID string `schema:"AC-Property-ID,omitempty"`
}

func (r *GetCashieringFinalInvoiceContentByFolioIDRequest) Headers() *GetCashieringFinalInvoiceContentByFolioIDHeaders {
	return r.headers
}

func (c *Client) NewGetCashieringFinalInvoiceContentByFolioIDPathParams() *GetCashieringFinalInvoiceContentByFolioIDPathParams {
	return &GetCashieringFinalInvoiceContentByFolioIDPathParams{}
}

type GetCashieringFinalInvoiceContentByFolioIDPathParams struct {
	AccountID string `schema:"account_id"`
	FolioID   string `schema:"folio_id"`
}

func (p *GetCashieringFinalInvoiceContentByFolioIDPathParams) Params() map[string]string {
	return map[string]string{
		"account_id": p.AccountID,
		"folio_id":   p.FolioID,
	}
}

func (r *GetCashieringFinalInvoiceContentByFolioIDRequest) PathParams() *GetCashieringFinalInvoiceContentByFolioIDPathParams {
	return r.pathParams
}

func (r *GetCashieringFinalInvoiceContentByFolioIDRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *GetCashieringFinalInvoiceContentByFolioIDRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetCashieringFinalInvoiceContentByFolioIDRequest) Method() string {
	return r.method
}

func (s *Client) NewGetCashieringFinalInvoiceContentByFolioIDRequestBody() GetCashieringFinalInvoiceContentByFolioIDRequestBody {
	return GetCashieringFinalInvoiceContentByFolioIDRequestBody{}
}

type GetCashieringFinalInvoiceContentByFolioIDRequestBody struct {
}

func (r *GetCashieringFinalInvoiceContentByFolioIDRequest) RequestBody() *GetCashieringFinalInvoiceContentByFolioIDRequestBody {
	return nil
}

func (r *GetCashieringFinalInvoiceContentByFolioIDRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *GetCashieringFinalInvoiceContentByFolioIDRequest) SetRequestBody(body GetCashieringFinalInvoiceContentByFolioIDRequestBody) {
	r.requestBody = body
}

func (r *GetCashieringFinalInvoiceContentByFolioIDRequest) NewResponseBody() *GetCashieringFinalInvoiceContentByFolioIDResponseBody {
	return &GetCashieringFinalInvoiceContentByFolioIDResponseBody{}
}

type GetCashieringFinalInvoiceContentByFolioIDResponseBody PDFResponseBody

func (r *GetCashieringFinalInvoiceContentByFolioIDRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("api-gateway/cashiering/v1/accounts/{{.account_id}}/folios/{{.folio_id}}/correspondences/invoice/final/content", r.PathParams())
	return &u
}

func extractFilename(contentDisposition string) string {
	_, params, err := mime.ParseMediaType(contentDisposition)
	if err != nil {
		return ""
	}
	filename := params["filename"]
	if filename == "" {
		// Fallback: probeer filename* (RFC 5987)
		filename = params["filename*"]
		if strings.HasPrefix(filename, "UTF-8''") {
			filename = strings.TrimPrefix(filename, "UTF-8''")
		}
	}
	return filename
}

func (r *GetCashieringFinalInvoiceContentByFolioIDRequest) Do(ctx context.Context) (GetCashieringFinalInvoiceContentByFolioIDResponseBody, error) {
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

	if r.Headers().PropertyID != "" {
		req.Header.Add("AC-Property-ID", r.Headers().PropertyID)
	}

	httpResp, err := r.client.Do(req, nil)
	if err != nil {
		return *r.NewResponseBody(), err
	}
	defer httpResp.Body.Close()

	content, err := io.ReadAll(httpResp.Body)
	if err != nil {
		return *r.NewResponseBody(), err
	}

	// Assign to the response body
	responseBody := r.NewResponseBody()
	responseBody.Content = content
	responseBody.ContentType = httpResp.Header.Get("Content-Type")
	responseBody.FileName = extractFilename(httpResp.Header.Get("Content-Disposition"))

	return *responseBody, err
}
