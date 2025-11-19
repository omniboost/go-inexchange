package inexchange

import (
	"context"
	"net/http"
	"net/url"

	"github.com/omniboost/go-inexchange/utils"
)

func (c *Client) NewGetConfigurationSupportedTypesRequest() GetConfigurationSupportedTypesRequest {
	return GetConfigurationSupportedTypesRequest{
		client:      c,
		queryParams: c.NewGetConfigurationSupportedTypesQueryParams(),
		pathParams:  c.NewGetConfigurationSupportedTypesPathParams(),
		method:      http.MethodGet,
		headers:     http.Header{},
		requestBody: c.NewGetConfigurationSupportedTypesRequestBody(),
	}
}

type GetConfigurationSupportedTypesRequest struct {
	client      *Client
	queryParams *GetConfigurationSupportedTypesQueryParams
	pathParams  *GetConfigurationSupportedTypesPathParams
	method      string
	headers     http.Header
	requestBody GetConfigurationSupportedTypesRequestBody
}

func (c *Client) NewGetConfigurationSupportedTypesQueryParams() *GetConfigurationSupportedTypesQueryParams {
	return &GetConfigurationSupportedTypesQueryParams{}
}

type GetConfigurationSupportedTypesQueryParams struct {
	PageNumber int32 `schema:"pageNumber,omitempty"`
	PageSize   int32 `schema:"pageSize,omitempty"`
}

func (p GetConfigurationSupportedTypesQueryParams) ToURLValues() (url.Values, error) {
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

func (r *GetConfigurationSupportedTypesRequest) QueryParams() *GetConfigurationSupportedTypesQueryParams {
	return r.queryParams
}

func (c *Client) NewGetConfigurationSupportedTypesPathParams() *GetConfigurationSupportedTypesPathParams {
	return &GetConfigurationSupportedTypesPathParams{}
}

type GetConfigurationSupportedTypesPathParams struct {
}

func (p *GetConfigurationSupportedTypesPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GetConfigurationSupportedTypesRequest) PathParams() *GetConfigurationSupportedTypesPathParams {
	return r.pathParams
}

func (r *GetConfigurationSupportedTypesRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *GetConfigurationSupportedTypesRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetConfigurationSupportedTypesRequest) Method() string {
	return r.method
}

func (s *Client) NewGetConfigurationSupportedTypesRequestBody() GetConfigurationSupportedTypesRequestBody {
	return GetConfigurationSupportedTypesRequestBody{}
}

type GetConfigurationSupportedTypesRequestBody struct {
}

func (r *GetConfigurationSupportedTypesRequest) RequestBody() *GetConfigurationSupportedTypesRequestBody {
	return nil
}

func (r *GetConfigurationSupportedTypesRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *GetConfigurationSupportedTypesRequest) SetRequestBody(body GetConfigurationSupportedTypesRequestBody) {
	r.requestBody = body
}

func (r *GetConfigurationSupportedTypesRequest) NewResponseBody() *GetConfigurationSupportedTypesResponseBody {
	return &GetConfigurationSupportedTypesResponseBody{}
}

type GetConfigurationSupportedTypesResponseBody PaginatedResponse[ConfigurationSupportedType]

func (r *GetConfigurationSupportedTypesRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("api-gateway/aggregator/v1/configuration/supported-types", r.PathParams())
	return &u
}

func (r *GetConfigurationSupportedTypesRequest) Do(ctx context.Context) (GetConfigurationSupportedTypesResponseBody, error) {
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

func (r *GetConfigurationSupportedTypesRequest) All(ctx context.Context) ([]ConfigurationSupportedType, error) {
	supportedTypes := []ConfigurationSupportedType{}
	for {
		resp, err := r.Do(ctx)
		if err != nil {
			return supportedTypes, err
		}

		// Break out of loop when no supportedTypes are found
		if len(resp.Results) == 0 {
			break
		}

		// Add supportedTypes to list
		supportedTypes = append(supportedTypes, resp.Results...)

		if len(supportedTypes) == resp.Paging.TotalCount {
			break
		}

		// Increment page number
		r.QueryParams().PageNumber = r.QueryParams().PageNumber + 1
	}

	return supportedTypes, nil
}
