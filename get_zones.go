package inexchange

import (
	"context"
	"net/http"
	"net/url"

	"github.com/omniboost/go-inexchange/utils"
)

func (c *Client) NewGetZonesRequest() GetZonesRequest {
	return GetZonesRequest{
		client:      c,
		queryParams: c.NewGetZonesQueryParams(),
		pathParams:  c.NewGetZonesPathParams(),
		method:      http.MethodGet,
		headers:     c.NewGetZonesHeaders(),
		requestBody: c.NewGetZonesRequestBody(),
	}
}

type GetZonesRequest struct {
	client      *Client
	queryParams *GetZonesQueryParams
	pathParams  *GetZonesPathParams
	method      string
	headers     *GetZonesHeaders
	requestBody GetZonesRequestBody
}

func (c *Client) NewGetZonesQueryParams() *GetZonesQueryParams {
	return &GetZonesQueryParams{}
}

type GetZonesQueryParams struct {
	OwningPropertyID   string `schema:"OwningPropertyID,omitempty"`
	OwningSubsidiaryID string `schema:"OwningSubsidiaryID,omitempty"`
	PageNumber         int32  `schema:"pageNumber,omitempty"`
	PageSize           int32  `schema:"pageSize,omitempty"`
}

func (p GetZonesQueryParams) ToURLValues() (url.Values, error) {
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

func (r *GetZonesRequest) QueryParams() *GetZonesQueryParams {
	return r.queryParams
}

func (c *Client) NewGetZonesHeaders() *GetZonesHeaders {
	return &GetZonesHeaders{}
}

type GetZonesHeaders struct {
	TenantID string `schema:"AC-Tenant-ID,omitempty"`
}

func (r *GetZonesRequest) Headers() *GetZonesHeaders {
	return r.headers
}

func (c *Client) NewGetZonesPathParams() *GetZonesPathParams {
	return &GetZonesPathParams{}
}

type GetZonesPathParams struct {
}

func (p *GetZonesPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GetZonesRequest) PathParams() *GetZonesPathParams {
	return r.pathParams
}

func (r *GetZonesRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *GetZonesRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetZonesRequest) Method() string {
	return r.method
}

func (s *Client) NewGetZonesRequestBody() GetZonesRequestBody {
	return GetZonesRequestBody{}
}

type GetZonesRequestBody struct {
}

func (r *GetZonesRequest) RequestBody() *GetZonesRequestBody {
	return nil
}

func (r *GetZonesRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *GetZonesRequest) SetRequestBody(body GetZonesRequestBody) {
	r.requestBody = body
}

func (r *GetZonesRequest) NewResponseBody() *GetZonesResponseBody {
	return &GetZonesResponseBody{}
}

type GetZonesResponseBody PaginatedResponse[Zone]

func (r *GetZonesRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("api-gateway/configuration/v1/zones", r.PathParams())
	return &u
}

func (r *GetZonesRequest) Do(ctx context.Context) (GetZonesResponseBody, error) {
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

	responseBody := r.NewResponseBody()
	_, err = r.client.Do(req, responseBody)
	return *responseBody, err
}

func (r *GetZonesRequest) All(ctx context.Context) ([]Zone, error) {
	zones := []Zone{}
	for {
		resp, err := r.Do(ctx)
		if err != nil {
			return zones, err
		}

		// Break out of loop when no zones are found
		if len(resp.Results) == 0 {
			break
		}

		// Add zones to list
		zones = append(zones, resp.Results...)

		if len(zones) == resp.Paging.TotalCount {
			break
		}

		// Increment page number
		r.QueryParams().PageNumber = r.QueryParams().PageNumber + 1
	}

	return zones, nil
}
