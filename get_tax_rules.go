package inexchange

import (
	"context"
	"net/http"
	"net/url"

	"github.com/omniboost/go-inexchange/utils"
)

func (c *Client) NewGetTaxRulesRequest() GetTaxRulesRequest {
	return GetTaxRulesRequest{
		client:      c,
		queryParams: c.NewGetTaxRulesQueryParams(),
		pathParams:  c.NewGetTaxRulesPathParams(),
		method:      http.MethodGet,
		headers:     c.NewGetTaxRulesHeaders(),
		requestBody: c.NewGetTaxRulesRequestBody(),
	}
}

type GetTaxRulesRequest struct {
	client      *Client
	queryParams *GetTaxRulesQueryParams
	pathParams  *GetTaxRulesPathParams
	method      string
	headers     *GetTaxRulesHeaders
	requestBody GetTaxRulesRequestBody
}

func (c *Client) NewGetTaxRulesQueryParams() *GetTaxRulesQueryParams {
	return &GetTaxRulesQueryParams{}
}

type GetTaxRulesQueryParams struct {
	ZoneID     string `schema:"zoneId,omitempty"`
	Sort       string `schema:"sort,omitempty"`
	Includes   string `schema:"includes,omitempty"`
	PageNumber int32  `schema:"pageNumber,omitempty"`
	PageSize   int32  `schema:"pageSize,omitempty"`
	Filter     string `schema:"filter,omitempty"`
}

func (p GetTaxRulesQueryParams) ToURLValues() (url.Values, error) {
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

func (r *GetTaxRulesRequest) QueryParams() *GetTaxRulesQueryParams {
	return r.queryParams
}

func (c *Client) NewGetTaxRulesHeaders() *GetTaxRulesHeaders {
	return &GetTaxRulesHeaders{}
}

type GetTaxRulesHeaders struct {
	SubsidiaryID string `schema:"AC-Subsidiary-ID,omitempty"`
	TenantID     string `schema:"AC-Tenant-ID,omitempty"`
	PropertyID   string `schema:"AC-Property-ID,omitempty"`
}

func (r *GetTaxRulesRequest) Headers() *GetTaxRulesHeaders {
	return r.headers
}

func (c *Client) NewGetTaxRulesPathParams() *GetTaxRulesPathParams {
	return &GetTaxRulesPathParams{}
}

type GetTaxRulesPathParams struct {
}

func (p *GetTaxRulesPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GetTaxRulesRequest) PathParams() *GetTaxRulesPathParams {
	return r.pathParams
}

func (r *GetTaxRulesRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *GetTaxRulesRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetTaxRulesRequest) Method() string {
	return r.method
}

func (s *Client) NewGetTaxRulesRequestBody() GetTaxRulesRequestBody {
	return GetTaxRulesRequestBody{}
}

type GetTaxRulesRequestBody struct {
}

func (r *GetTaxRulesRequest) RequestBody() *GetTaxRulesRequestBody {
	return nil
}

func (r *GetTaxRulesRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *GetTaxRulesRequest) SetRequestBody(body GetTaxRulesRequestBody) {
	r.requestBody = body
}

func (r *GetTaxRulesRequest) NewResponseBody() *GetTaxRulesResponseBody {
	return &GetTaxRulesResponseBody{}
}

type GetTaxRulesResponseBody PaginatedResponse[TaxRule]

func (r *GetTaxRulesRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("api-gateway/configuration/v1/entities/TaxRule", r.PathParams())
	return &u
}

func (r *GetTaxRulesRequest) Do(ctx context.Context) (GetTaxRulesResponseBody, error) {
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
	if r.Headers().SubsidiaryID != "" {
		req.Header.Add("AC-Subsidiary-ID", r.Headers().SubsidiaryID)
	}

	if r.Headers().PropertyID != "" {
		req.Header.Add("AC-Property-ID", r.Headers().PropertyID)
	}

	if r.Headers().TenantID != "" {
		req.Header.Add("AC-Tenant-ID", r.Headers().TenantID)
	}

	responseBody := r.NewResponseBody()
	_, err = r.client.Do(req, responseBody)
	return *responseBody, err
}

func (r *GetTaxRulesRequest) All(ctx context.Context) ([]TaxRule, error) {
	taxRules := []TaxRule{}
	for {
		resp, err := r.Do(ctx)
		if err != nil {
			return taxRules, err
		}

		// Break out of loop when no taxRules are found
		if len(resp.Results) == 0 {
			break
		}

		// Add taxRules to list
		taxRules = append(taxRules, resp.Results...)

		if len(taxRules) == resp.Paging.TotalCount {
			break
		}

		// Increment page number
		r.QueryParams().PageNumber = r.QueryParams().PageNumber + 1
	}

	return taxRules, nil
}
