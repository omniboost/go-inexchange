package inexchange

import (
	"context"
	"net/http"
	"net/url"

	"github.com/omniboost/go-inexchange/utils"
)

func (c *Client) NewDocumentsOutboundRequest() DocumentsOutboundRequest {
	return DocumentsOutboundRequest{
		client:      c,
		queryParams: c.NewDocumentsOutboundQueryParams(),
		pathParams:  c.NewDocumentsOutboundPathParams(),
		method:      http.MethodPost,
		headers:     c.NewDocumentsOutboundHeaders(),
		requestBody: c.NewDocumentsOutboundRequestBody(),
	}
}

type DocumentsOutboundRequest struct {
	client      *Client
	queryParams *DocumentsOutboundQueryParams
	pathParams  *DocumentsOutboundPathParams
	method      string
	headers     *DocumentsOutboundHeaders
	requestBody DocumentsOutboundRequestBody
}

func (c *Client) NewDocumentsOutboundQueryParams() *DocumentsOutboundQueryParams {
	return &DocumentsOutboundQueryParams{}
}

type DocumentsOutboundQueryParams struct {
}

func (p DocumentsOutboundQueryParams) ToURLValues() (url.Values, error) {
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

func (r *DocumentsOutboundRequest) QueryParams() *DocumentsOutboundQueryParams {
	return r.queryParams
}

func (c *Client) NewDocumentsOutboundHeaders() *DocumentsOutboundHeaders {
	return &DocumentsOutboundHeaders{}
}

type DocumentsOutboundHeaders struct {
	TenantID string `schema:"AC-Tenant-ID,omitempty"`
}

func (r *DocumentsOutboundRequest) Headers() *DocumentsOutboundHeaders {
	return r.headers
}

func (c *Client) NewDocumentsOutboundPathParams() *DocumentsOutboundPathParams {
	return &DocumentsOutboundPathParams{}
}

type DocumentsOutboundPathParams struct{}

func (p *DocumentsOutboundPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *DocumentsOutboundRequest) PathParams() *DocumentsOutboundPathParams {
	return r.pathParams
}

func (r *DocumentsOutboundRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *DocumentsOutboundRequest) SetMethod(method string) {
	r.method = method
}

func (r *DocumentsOutboundRequest) Method() string {
	return r.method
}

func (s *Client) NewDocumentsOutboundRequestBody() DocumentsOutboundRequestBody {
	return DocumentsOutboundRequestBody{}
}

type DocumentsOutboundRequestBody struct {
	SendDocumentAs       DocumentsOutboundSendDocumentAs       `json:"SendDocumentAs"`
	RecipientInformation DocumentsOutboundRecipientInformation `json:"RecipientInformation"`
	Document             DocumentsOutboundDocument             `json:"Document"`
}

func (r *DocumentsOutboundRequest) RequestBody() *DocumentsOutboundRequestBody {
	return &r.requestBody
}

func (r *DocumentsOutboundRequest) RequestBodyInterface() any {
	return r.requestBody
}

func (r *DocumentsOutboundRequest) SetRequestBody(body DocumentsOutboundRequestBody) {
	r.requestBody = body
}

func (r *DocumentsOutboundRequest) NewResponseBody() *DocumentsOutboundResponseBody {
	return &DocumentsOutboundResponseBody{}
}

type DocumentsOutboundResponseBody struct {
	Location string `json:"location"`
}

func (r *DocumentsOutboundRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("/documents/outbound", r.PathParams())
	return &u
}

func (r *DocumentsOutboundRequest) Do(ctx context.Context) (DocumentsOutboundResponseBody, error) {
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
	resp, err := r.client.Do(req, responseBody)
	if err != nil {
		return *responseBody, err
	}

	// we need the header location from the response to
	// Location: http://test-api:81/v1/api/documents/outbound/5cde2c1c-593f-4f19-a9b3-9a431efcce14
	responseBody.Location = resp.Header.Get("Location")

	return *responseBody, nil
}
