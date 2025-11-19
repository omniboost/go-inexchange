package inexchange

import (
	"context"
	"net/http"
	"net/url"

	"github.com/omniboost/go-inexchange/utils"
)

func (c *Client) NewGetTransactionCodesRequest() GetTransactionCodesRequest {
	return GetTransactionCodesRequest{
		client:      c,
		queryParams: c.NewGetTransactionCodesQueryParams(),
		pathParams:  c.NewGetTransactionCodesPathParams(),
		method:      http.MethodGet,
		headers:     c.NewGetTransactionCodesHeaders(),
		requestBody: c.NewGetTransactionCodesRequestBody(),
	}
}

type GetTransactionCodesRequest struct {
	client      *Client
	queryParams *GetTransactionCodesQueryParams
	pathParams  *GetTransactionCodesPathParams
	method      string
	headers     *GetTransactionCodesHeaders
	requestBody GetTransactionCodesRequestBody
}

func (c *Client) NewGetTransactionCodesQueryParams() *GetTransactionCodesQueryParams {
	return &GetTransactionCodesQueryParams{}
}

type GetTransactionCodesQueryParams struct {
	ZoneID     string `schema:"zoneId,omitempty"`
	Sort       string `schema:"sort,omitempty"`
	Includes   string `schema:"includes,omitempty"`
	PageNumber int32  `schema:"pageNumber,omitempty"`
	PageSize   int32  `schema:"pageSize,omitempty"`
	Filter     string `schema:"filter,omitempty"`
}

func (p GetTransactionCodesQueryParams) ToURLValues() (url.Values, error) {
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

func (r *GetTransactionCodesRequest) QueryParams() *GetTransactionCodesQueryParams {
	return r.queryParams
}

func (c *Client) NewGetTransactionCodesHeaders() *GetTransactionCodesHeaders {
	return &GetTransactionCodesHeaders{}
}

type GetTransactionCodesHeaders struct {
	PropertyID string `schema:"AC-Property-ID,omitempty"`
}

func (r *GetTransactionCodesRequest) Headers() *GetTransactionCodesHeaders {
	return r.headers
}

func (c *Client) NewGetTransactionCodesPathParams() *GetTransactionCodesPathParams {
	return &GetTransactionCodesPathParams{}
}

type GetTransactionCodesPathParams struct {
}

func (p *GetTransactionCodesPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GetTransactionCodesRequest) PathParams() *GetTransactionCodesPathParams {
	return r.pathParams
}

func (r *GetTransactionCodesRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *GetTransactionCodesRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetTransactionCodesRequest) Method() string {
	return r.method
}

func (s *Client) NewGetTransactionCodesRequestBody() GetTransactionCodesRequestBody {
	return GetTransactionCodesRequestBody{}
}

type GetTransactionCodesRequestBody struct {
}

func (r *GetTransactionCodesRequest) RequestBody() *GetTransactionCodesRequestBody {
	return nil
}

func (r *GetTransactionCodesRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *GetTransactionCodesRequest) SetRequestBody(body GetTransactionCodesRequestBody) {
	r.requestBody = body
}

func (r *GetTransactionCodesRequest) NewResponseBody() *GetTransactionCodesResponseBody {
	return &GetTransactionCodesResponseBody{}
}

type GetTransactionCodesResponseBody PaginatedResponse[TransactionCode]

func (r *GetTransactionCodesRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("api-gateway/configuration/v1/entities/TransactionCode", r.PathParams())
	return &u
}

func (r *GetTransactionCodesRequest) Do(ctx context.Context) (GetTransactionCodesResponseBody, error) {
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
	if r.Headers().PropertyID != "" {
		req.Header.Add("AC-Property-ID", r.Headers().PropertyID)
	}

	responseBody := r.NewResponseBody()
	_, err = r.client.Do(req, responseBody)
	return *responseBody, err
}

func (r *GetTransactionCodesRequest) All(ctx context.Context) ([]TransactionCode, error) {
	transactionCodes := []TransactionCode{}
	for {
		resp, err := r.Do(ctx)
		if err != nil {
			return transactionCodes, err
		}

		// Break out of loop when no transactionCodes are found
		if len(resp.Results) == 0 {
			break
		}

		// Add transactionCodes to list
		transactionCodes = append(transactionCodes, resp.Results...)

		if len(transactionCodes) == resp.Paging.TotalCount {
			break
		}

		// Increment page number
		r.QueryParams().PageNumber = r.QueryParams().PageNumber + 1
	}

	return transactionCodes, nil
}
