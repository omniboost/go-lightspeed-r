package lightspeed_r

import (
	"net/http"
	"net/url"
	"strconv"

	"github.com/omniboost/go-lightspeed-r/utils"
)

func (c *Client) NewTaxCategoryByIdGet() TaxCategoryByIdGet {
	r := TaxCategoryByIdGet{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type TaxCategoryByIdGet struct {
	client      *Client
	queryParams *TaxCategoryByIdGetQueryParams
	pathParams  *TaxCategoryByIdGetPathParams
	method      string
	headers     http.Header
	requestBody TaxCategoryByIdGetBody
}

func (r TaxCategoryByIdGet) NewQueryParams() *TaxCategoryByIdGetQueryParams {
	return &TaxCategoryByIdGetQueryParams{}
}

type TaxCategoryByIdGetQueryParams struct{}

func (p TaxCategoryByIdGetQueryParams) ToURLValues() (url.Values, error) {
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

func (r *TaxCategoryByIdGet) QueryParams() *TaxCategoryByIdGetQueryParams {
	return r.queryParams
}

func (r TaxCategoryByIdGet) NewPathParams() *TaxCategoryByIdGetPathParams {
	return &TaxCategoryByIdGetPathParams{
		AccountID: r.client.accountID,
	}
}

type TaxCategoryByIdGetPathParams struct {
	AccountID     int `schema:"account_id"`
	TaxCategoryID int `schema:"tax_category_id"`
}

func (p *TaxCategoryByIdGetPathParams) Params() map[string]string {
	return map[string]string{
		"account_id":      strconv.Itoa(p.AccountID),
		"tax_category_id": strconv.Itoa(p.TaxCategoryID),
	}
}

func (r *TaxCategoryByIdGet) PathParams() *TaxCategoryByIdGetPathParams {
	return r.pathParams
}

func (r *TaxCategoryByIdGet) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *TaxCategoryByIdGet) SetMethod(method string) {
	r.method = method
}

func (r *TaxCategoryByIdGet) Method() string {
	return r.method
}

func (r TaxCategoryByIdGet) NewRequestBody() TaxCategoryByIdGetBody {
	return TaxCategoryByIdGetBody{}
}

type TaxCategoryByIdGetBody struct{}

func (r *TaxCategoryByIdGet) RequestBody() *TaxCategoryByIdGetBody {
	return nil
}

func (r *TaxCategoryByIdGet) RequestBodyInterface() interface{} {
	return nil
}

func (r *TaxCategoryByIdGet) SetRequestBody(body TaxCategoryByIdGetBody) {
	r.requestBody = body
}

func (r *TaxCategoryByIdGet) NewResponseBody() *TaxCategoryByIdGetResponseBody {
	return &TaxCategoryByIdGetResponseBody{}
}

type TaxCategoryByIdGetResponseBody TaxCategoryResp

func (r *TaxCategoryByIdGet) URL() *url.URL {
	u := r.client.GetEndpointURL("/API/V3/Account/{{.account_id}}/TaxCategory/{{.tax_category_id}}.json", r.PathParams())
	return &u
}

func (r *TaxCategoryByIdGet) Do() (TaxCategoryByIdGetResponseBody, error) {
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
