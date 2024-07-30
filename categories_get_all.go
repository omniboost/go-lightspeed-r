package lightspeed_r

import (
	"net/http"
	"net/url"
	"strconv"

	"github.com/omniboost/go-lightspeed-r/utils"
)

func (c *Client) NewCategoriesGetAll() CategoriesGetAll {
	r := CategoriesGetAll{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type CategoriesGetAll struct {
	client      *Client
	queryParams *CategoriesGetAllQueryParams
	pathParams  *CategoriesGetAllPathParams
	method      string
	headers     http.Header
	requestBody CategoriesGetAllBody
}

func (r CategoriesGetAll) NewQueryParams() *CategoriesGetAllQueryParams {
	return &CategoriesGetAllQueryParams{}
}

type CategoriesGetAllQueryParams struct{}

func (p CategoriesGetAllQueryParams) ToURLValues() (url.Values, error) {
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

func (r *CategoriesGetAll) QueryParams() *CategoriesGetAllQueryParams {
	return r.queryParams
}

func (r CategoriesGetAll) NewPathParams() *CategoriesGetAllPathParams {
	return &CategoriesGetAllPathParams{
		AccountID: r.client.accountID,
	}
}

type CategoriesGetAllPathParams struct {
	AccountID int `schema:"account_id"`
}

func (p *CategoriesGetAllPathParams) Params() map[string]string {
	return map[string]string{
		"account_id": strconv.Itoa(p.AccountID),
	}
}

func (r *CategoriesGetAll) PathParams() *CategoriesGetAllPathParams {
	return r.pathParams
}

func (r *CategoriesGetAll) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *CategoriesGetAll) SetMethod(method string) {
	r.method = method
}

func (r *CategoriesGetAll) Method() string {
	return r.method
}

func (r CategoriesGetAll) NewRequestBody() CategoriesGetAllBody {
	return CategoriesGetAllBody{}
}

type CategoriesGetAllBody struct{}

func (r *CategoriesGetAll) RequestBody() *CategoriesGetAllBody {
	return nil
}

func (r *CategoriesGetAll) RequestBodyInterface() interface{} {
	return nil
}

func (r *CategoriesGetAll) SetRequestBody(body CategoriesGetAllBody) {
	r.requestBody = body
}

func (r *CategoriesGetAll) NewResponseBody() *CategoriesGetAllResponseBody {
	return &CategoriesGetAllResponseBody{}
}

type CategoriesGetAllResponseBody CategoriesResp

func (r *CategoriesGetAll) URL() *url.URL {
	u := r.client.GetEndpointURL("/API/V3/Account/{{.account_id}}/Category.json", r.PathParams())
	return &u
}

func (r *CategoriesGetAll) Do() (CategoriesGetAllResponseBody, error) {
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
