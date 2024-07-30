package lightspeed_r

import (
	"net/http"
	"net/url"
	"strconv"

	"github.com/omniboost/go-lightspeed-r/utils"
)

func (c *Client) NewTaxCategoriesGetAll() TaxCategoriesGetAll {
	r := TaxCategoriesGetAll{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type TaxCategoriesGetAll struct {
	client      *Client
	queryParams *TaxCategoriesGetAllQueryParams
	pathParams  *TaxCategoriesGetAllPathParams
	method      string
	headers     http.Header
	requestBody TaxCategoriesGetAllBody
}

func (r TaxCategoriesGetAll) NewQueryParams() *TaxCategoriesGetAllQueryParams {
	return &TaxCategoriesGetAllQueryParams{}
}

type TaxCategoriesGetAllQueryParams struct{}

func (p TaxCategoriesGetAllQueryParams) ToURLValues() (url.Values, error) {
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

func (r *TaxCategoriesGetAll) QueryParams() *TaxCategoriesGetAllQueryParams {
	return r.queryParams
}

func (r TaxCategoriesGetAll) NewPathParams() *TaxCategoriesGetAllPathParams {
	return &TaxCategoriesGetAllPathParams{
		AccountID: r.client.accountID,
	}
}

type TaxCategoriesGetAllPathParams struct {
	AccountID int `schema:"account_id"`
}

func (p *TaxCategoriesGetAllPathParams) Params() map[string]string {
	return map[string]string{
		"account_id": strconv.Itoa(p.AccountID),
	}
}

func (r *TaxCategoriesGetAll) PathParams() *TaxCategoriesGetAllPathParams {
	return r.pathParams
}

func (r *TaxCategoriesGetAll) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *TaxCategoriesGetAll) SetMethod(method string) {
	r.method = method
}

func (r *TaxCategoriesGetAll) Method() string {
	return r.method
}

func (r TaxCategoriesGetAll) NewRequestBody() TaxCategoriesGetAllBody {
	return TaxCategoriesGetAllBody{}
}

type TaxCategoriesGetAllBody struct{}

func (r *TaxCategoriesGetAll) RequestBody() *TaxCategoriesGetAllBody {
	return nil
}

func (r *TaxCategoriesGetAll) RequestBodyInterface() interface{} {
	return nil
}

func (r *TaxCategoriesGetAll) SetRequestBody(body TaxCategoriesGetAllBody) {
	r.requestBody = body
}

func (r *TaxCategoriesGetAll) NewResponseBody() *TaxCategoriesGetAllResponseBody {
	return &TaxCategoriesGetAllResponseBody{}
}

type TaxCategoriesGetAllResponseBody TaxCategoriesResp

func (r *TaxCategoriesGetAll) URL() *url.URL {
	u := r.client.GetEndpointURL("/API/V3/Account/{{.account_id}}/TaxCategory.json", r.PathParams())
	return &u
}

func (r *TaxCategoriesGetAll) Do() (TaxCategoriesGetAllResponseBody, error) {
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
