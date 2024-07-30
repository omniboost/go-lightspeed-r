package lightspeed_r

import (
	"net/http"
	"net/url"
	"strconv"

	"github.com/omniboost/go-lightspeed-r/utils"
)

func (c *Client) NewCategoryByIdGet() CategoryByIdGet {
	r := CategoryByIdGet{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type CategoryByIdGet struct {
	client      *Client
	queryParams *CategoryByIdGetQueryParams
	pathParams  *CategoryByIdGetPathParams
	method      string
	headers     http.Header
	requestBody CategoryByIdGetBody
}

func (r CategoryByIdGet) NewQueryParams() *CategoryByIdGetQueryParams {
	return &CategoryByIdGetQueryParams{}
}

type CategoryByIdGetQueryParams struct{}

func (p CategoryByIdGetQueryParams) ToURLValues() (url.Values, error) {
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

func (r *CategoryByIdGet) QueryParams() *CategoryByIdGetQueryParams {
	return r.queryParams
}

func (r CategoryByIdGet) NewPathParams() *CategoryByIdGetPathParams {
	return &CategoryByIdGetPathParams{
		AccountID: r.client.accountID,
	}
}

type CategoryByIdGetPathParams struct {
	AccountID  int `schema:"account_id"`
	CategoryID int `schema:"category_id"`
}

func (p *CategoryByIdGetPathParams) Params() map[string]string {
	return map[string]string{
		"account_id":  strconv.Itoa(p.AccountID),
		"category_id": strconv.Itoa(p.CategoryID),
	}
}

func (r *CategoryByIdGet) PathParams() *CategoryByIdGetPathParams {
	return r.pathParams
}

func (r *CategoryByIdGet) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *CategoryByIdGet) SetMethod(method string) {
	r.method = method
}

func (r *CategoryByIdGet) Method() string {
	return r.method
}

func (r CategoryByIdGet) NewRequestBody() CategoryByIdGetBody {
	return CategoryByIdGetBody{}
}

type CategoryByIdGetBody struct{}

func (r *CategoryByIdGet) RequestBody() *CategoryByIdGetBody {
	return nil
}

func (r *CategoryByIdGet) RequestBodyInterface() interface{} {
	return nil
}

func (r *CategoryByIdGet) SetRequestBody(body CategoryByIdGetBody) {
	r.requestBody = body
}

func (r *CategoryByIdGet) NewResponseBody() *CategoryByIdGetResponseBody {
	return &CategoryByIdGetResponseBody{}
}

type CategoryByIdGetResponseBody CategoryResp

func (r *CategoryByIdGet) URL() *url.URL {
	u := r.client.GetEndpointURL("/API/V3/Account/{{.account_id}}/Category/{{.category_id}}.json", r.PathParams())
	return &u
}

func (r *CategoryByIdGet) Do() (CategoryByIdGetResponseBody, error) {
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
