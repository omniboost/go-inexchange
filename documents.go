package inexchange

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/omniboost/go-inexchange/utils"
)

func (c *Client) NewDocumentsRequest() DocumentsRequest {
	return DocumentsRequest{
		client:      c,
		queryParams: c.NewDocumentsQueryParams(),
		pathParams:  c.NewDocumentsPathParams(),
		formParams:  c.NewDocumentsFormParams(),
		method:      http.MethodPost,
		headers:     c.NewDocumentsHeaders(),
		requestBody: c.NewDocumentsRequestBody(),
	}
}

type DocumentsRequest struct {
	client      *Client
	queryParams *DocumentsQueryParams
	pathParams  *DocumentsPathParams
	formParams  *DocumentsFormParams
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

type DocumentsFormParams struct {
	File FormFile
}

func (p DocumentsFormParams) IsMultiPart() bool {
	return true
}

func (p DocumentsFormParams) Files() map[string]FormFile {
	return map[string]FormFile{
		"file": p.File,
	}
}

func (p DocumentsFormParams) Values() url.Values {
	return url.Values{}
}

func (c *Client) NewDocumentsFormParams() *DocumentsFormParams {
	return &DocumentsFormParams{}
}

func (r *DocumentsRequest) FormParams() *DocumentsFormParams {
	return r.formParams
}

func (r *DocumentsRequest) FormParamsInterface() Form {
	return r.formParams
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
}

func (r *DocumentsRequest) RequestBody() *DocumentsRequestBody {
	return &r.requestBody
}

func (r *DocumentsRequest) RequestBodyInterface() any {
	return nil
}

func (r *DocumentsRequest) SetRequestBody(body DocumentsRequestBody) {
	r.requestBody = body
}

func (r *DocumentsRequest) NewResponseBody() *DocumentsResponseBody {
	return &DocumentsResponseBody{}
}

type DocumentsResponseBody struct {
	DocumentURI string `json:"DocumentURI"`
}

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

	req.Header.Add("Filename", r.FormParams().File.Filename)
	req.Header.Set("Content-Disposition", fmt.Sprintf(`attachment; name="File"; filename="%s"`, r.FormParams().File.Filename))

	// Process query parameters
	err = utils.AddQueryParamsToRequest(r.QueryParams(), req, false)
	if err != nil {
		return *r.NewResponseBody(), err
	}

	responseBody := r.NewResponseBody()
	resp, err := r.client.Do(req, responseBody)

	// we need the header location from the response to
	// Location: urn:inexchangedocument:5d1f85b8-da11-4007-bf47-c756735bcb78
	responseBody.DocumentURI = resp.Header.Get("Location")

	return *responseBody, err
}
