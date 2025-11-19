package inexchange

import (
	"context"
	"net/http"
	"net/url"

	"github.com/omniboost/go-inexchange/utils"
)

func (c *Client) NewGetCompanyProfilesRequest() GetCompanyProfilesRequest {
	return GetCompanyProfilesRequest{
		client:      c,
		queryParams: c.NewGetCompanyProfilesQueryParams(),
		pathParams:  c.NewGetCompanyProfilesPathParams(),
		method:      http.MethodGet,
		headers:     c.NewGetCompanyProfilesHeaders(),
		requestBody: c.NewGetCompanyProfilesRequestBody(),
	}
}

type GetCompanyProfilesRequest struct {
	client      *Client
	queryParams *GetCompanyProfilesQueryParams
	pathParams  *GetCompanyProfilesPathParams
	method      string
	headers     *GetCompanyProfilesHeaders
	requestBody GetCompanyProfilesRequestBody
}

func (c *Client) NewGetCompanyProfilesQueryParams() *GetCompanyProfilesQueryParams {
	return &GetCompanyProfilesQueryParams{}
}

type GetCompanyProfilesQueryParams struct {
	Query      string `schema:"query,omitempty"`
	Sort       string `schema:"sort,omitempty"`
	Fields     string `schema:"fields,omitempty"`
	PageNumber int32  `schema:"pageNumber,omitempty"`
	PageSize   int32  `schema:"pageSize,omitempty"`
	Filter     string `schema:"filter,omitempty"`
}

func (p GetCompanyProfilesQueryParams) ToURLValues() (url.Values, error) {
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

func (r *GetCompanyProfilesRequest) QueryParams() *GetCompanyProfilesQueryParams {
	return r.queryParams
}

func (c *Client) NewGetCompanyProfilesHeaders() *GetCompanyProfilesHeaders {
	return &GetCompanyProfilesHeaders{}
}

type GetCompanyProfilesHeaders struct {
	PropertyID string `schema:"AC-Property-ID,omitempty"`
}

func (r *GetCompanyProfilesRequest) Headers() *GetCompanyProfilesHeaders {
	return r.headers
}

func (c *Client) NewGetCompanyProfilesPathParams() *GetCompanyProfilesPathParams {
	return &GetCompanyProfilesPathParams{}
}

type GetCompanyProfilesPathParams struct {
}

func (p *GetCompanyProfilesPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GetCompanyProfilesRequest) PathParams() *GetCompanyProfilesPathParams {
	return r.pathParams
}

func (r *GetCompanyProfilesRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *GetCompanyProfilesRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetCompanyProfilesRequest) Method() string {
	return r.method
}

func (s *Client) NewGetCompanyProfilesRequestBody() GetCompanyProfilesRequestBody {
	return GetCompanyProfilesRequestBody{}
}

type GetCompanyProfilesRequestBody struct {
}

func (r *GetCompanyProfilesRequest) RequestBody() *GetCompanyProfilesRequestBody {
	return nil
}

func (r *GetCompanyProfilesRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *GetCompanyProfilesRequest) SetRequestBody(body GetCompanyProfilesRequestBody) {
	r.requestBody = body
}

func (r *GetCompanyProfilesRequest) NewResponseBody() *GetCompanyProfilesResponseBody {
	return &GetCompanyProfilesResponseBody{}
}

type GetCompanyProfilesResponseBody PaginatedResponse[CompanyProfile]

func (r *GetCompanyProfilesRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("api-gateway/profiles/v1/company", r.PathParams())
	return &u
}

func (r *GetCompanyProfilesRequest) Do(ctx context.Context) (GetCompanyProfilesResponseBody, error) {
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
	req.Header.Add("AC-Property-ID", r.Headers().PropertyID)

	responseBody := r.NewResponseBody()
	_, err = r.client.Do(req, responseBody)
	return *responseBody, err
}

func (r *GetCompanyProfilesRequest) All(ctx context.Context) ([]CompanyProfile, error) {
	companies := []CompanyProfile{}
	for {
		resp, err := r.Do(ctx)
		if err != nil {
			return companies, err
		}

		// Break out of loop when no companies are found
		if len(resp.Results) == 0 {
			break
		}

		// Add companies to list
		companies = append(companies, resp.Results...)

		if len(companies) == resp.Paging.TotalCount {
			break
		}

		// Increment page number
		r.QueryParams().PageNumber = r.QueryParams().PageNumber + 1
	}

	return companies, nil
}
