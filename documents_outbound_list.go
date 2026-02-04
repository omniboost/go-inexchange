package inexchange

import (
	"context"
	"net/http"
	"net/url"

	"github.com/omniboost/go-inexchange/utils"
)

func (c *Client) NewDocumentsOutboundListRequest() DocumentsOutboundListRequest {
	return DocumentsOutboundListRequest{
		client:      c,
		queryParams: c.NewDocumentsOutboundListQueryParams(),
		pathParams:  c.NewDocumentsOutboundListPathParams(),
		method:      http.MethodPost,
		headers:     c.NewDocumentsOutboundListHeaders(),
		requestBody: c.NewDocumentsOutboundListRequestBody(),
	}
}

type DocumentsOutboundListRequest struct {
	client      *Client
	queryParams *DocumentsOutboundListQueryParams
	pathParams  *DocumentsOutboundListPathParams
	method      string
	headers     *DocumentsOutboundListHeaders
	requestBody DocumentsOutboundListRequestBody
}

func (c *Client) NewDocumentsOutboundListQueryParams() *DocumentsOutboundListQueryParams {
	return &DocumentsOutboundListQueryParams{}
}

type DocumentsOutboundListQueryParams struct {
	StatusCode string `schema:"statusCode,omitempty"`
	ParentID   string `schema:"parentId,omitempty"`
	Recursive  bool   `schema:"recursive,omitempty"`
	Sort       string `schema:"sort,omitempty"`
	PageNumber int32  `schema:"pageNumber,omitempty"`
	PageSize   int32  `schema:"pageSize,omitempty"`
}

func (p DocumentsOutboundListQueryParams) ToURLValues() (url.Values, error) {
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

func (r *DocumentsOutboundListRequest) QueryParams() *DocumentsOutboundListQueryParams {
	return r.queryParams
}

func (c *Client) NewDocumentsOutboundListHeaders() *DocumentsOutboundListHeaders {
	return &DocumentsOutboundListHeaders{}
}

type DocumentsOutboundListHeaders struct {
	TenantID string `schema:"AC-Tenant-ID,omitempty"`
}

func (r *DocumentsOutboundListRequest) Headers() *DocumentsOutboundListHeaders {
	return r.headers
}

func (c *Client) NewDocumentsOutboundListPathParams() *DocumentsOutboundListPathParams {
	return &DocumentsOutboundListPathParams{}
}

type DocumentsOutboundListPathParams struct {
}

func (p *DocumentsOutboundListPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *DocumentsOutboundListRequest) PathParams() *DocumentsOutboundListPathParams {
	return r.pathParams
}

func (r *DocumentsOutboundListRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *DocumentsOutboundListRequest) SetMethod(method string) {
	r.method = method
}

func (r *DocumentsOutboundListRequest) Method() string {
	return r.method
}

func (s *Client) NewDocumentsOutboundListRequestBody() DocumentsOutboundListRequestBody {
	return DocumentsOutboundListRequestBody{}
}

type DocumentsOutboundListRequestBody struct {
	Take             int      `json:"take,omitempty"`
	Skip             int      `json:"skip,omitempty"`
	CreatedFrom      DateTime `json:"createdFrom,omitzero"`
	CreatedTo        DateTime `json:"createdTo,omitzero"`
	UpdatedAfter     DateTime `json:"updatedAfter,omitzero"`
	DocumentType     string   `json:"documentType,omitempty"`
	Status           string   `json:"status,omitempty"`
	IgnoreStatuses   []string `json:"ignoreStatuses,omitempty"`
	IncludeFileInfo  bool     `json:"includeFileInfo,omitempty"`
	IncludeErrorInfo bool     `json:"includeErrorInfo,omitempty"`
}

func (r *DocumentsOutboundListRequest) RequestBody() *DocumentsOutboundListRequestBody {
	return &r.requestBody
}

func (r *DocumentsOutboundListRequest) RequestBodyInterface() any {
	return r.requestBody
}

func (r *DocumentsOutboundListRequest) SetRequestBody(body DocumentsOutboundListRequestBody) {
	r.requestBody = body
}

func (r *DocumentsOutboundListRequest) NewResponseBody() *DocumentsOutboundListResponseBody {
	return &DocumentsOutboundListResponseBody{}
}

type DocumentsOutboundListResponseBody struct {
	Documents  OutboundDocuments `json:"documents"`
	TotalCount int               `json:"totalCount"`
}

func (r *DocumentsOutboundListRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("/documents/outbound/list", r.PathParams())
	return &u
}

func (r *DocumentsOutboundListRequest) Do(ctx context.Context) (DocumentsOutboundListResponseBody, error) {
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
