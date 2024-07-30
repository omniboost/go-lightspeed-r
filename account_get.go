package lightspeed_r

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-lightspeed-r/utils"
)

func (c *Client) NewAccountGet() AccountGet {
	r := AccountGet{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type AccountGet struct {
	client      *Client
	queryParams *AccountGetQueryParams
	pathParams  *AccountGetPathParams
	method      string
	headers     http.Header
	requestBody AccountGetBody
}

func (r AccountGet) NewQueryParams() *AccountGetQueryParams {
	return &AccountGetQueryParams{}
}

type AccountGetQueryParams struct{}

func (p AccountGetQueryParams) ToURLValues() (url.Values, error) {
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

func (r *AccountGet) QueryParams() *AccountGetQueryParams {
	return r.queryParams
}

func (r AccountGet) NewPathParams() *AccountGetPathParams {
	return &AccountGetPathParams{}
}

type AccountGetPathParams struct{}

func (p *AccountGetPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *AccountGet) PathParams() *AccountGetPathParams {
	return r.pathParams
}

func (r *AccountGet) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *AccountGet) SetMethod(method string) {
	r.method = method
}

func (r *AccountGet) Method() string {
	return r.method
}

func (r AccountGet) NewRequestBody() AccountGetBody {
	return AccountGetBody{}
}

type AccountGetBody struct{}

func (r *AccountGet) RequestBody() *AccountGetBody {
	return nil
}

func (r *AccountGet) RequestBodyInterface() interface{} {
	return nil
}

func (r *AccountGet) SetRequestBody(body AccountGetBody) {
	r.requestBody = body
}

func (r *AccountGet) NewResponseBody() *AccountGetResponseBody {
	return &AccountGetResponseBody{}
}

type AccountGetResponseBody AccountResp

func (r *AccountGet) URL() *url.URL {
	u := r.client.GetEndpointURL("/API/V3/Account.json", r.PathParams())
	return &u
}

func (r *AccountGet) Do() (AccountGetResponseBody, error) {
	// Create http request
	req, err := r.client.NewRequest(nil, r)
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
