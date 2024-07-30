package lightspeed_r

import (
	"net/http"
	"net/url"
	"strconv"

	"github.com/omniboost/go-lightspeed-r/utils"
)

func (c *Client) NewTaxClassByIdGet() TaxClassByIdGet {
	r := TaxClassByIdGet{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type TaxClassByIdGet struct {
	client      *Client
	queryParams *TaxClassByIdGetQueryParams
	pathParams  *TaxClassByIdGetPathParams
	method      string
	headers     http.Header
	requestBody TaxClassByIdGetBody
}

func (r TaxClassByIdGet) NewQueryParams() *TaxClassByIdGetQueryParams {
	return &TaxClassByIdGetQueryParams{}
}

type TaxClassByIdGetQueryParams struct{}

func (p TaxClassByIdGetQueryParams) ToURLValues() (url.Values, error) {
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

func (r *TaxClassByIdGet) QueryParams() *TaxClassByIdGetQueryParams {
	return r.queryParams
}

func (r TaxClassByIdGet) NewPathParams() *TaxClassByIdGetPathParams {
	return &TaxClassByIdGetPathParams{
		AccountID: r.client.accountID,
	}
}

type TaxClassByIdGetPathParams struct {
	AccountID  int `schema:"account_id"`
	TaxClassID int `schema:"tax_class_id"`
}

func (p *TaxClassByIdGetPathParams) Params() map[string]string {
	return map[string]string{
		"account_id":   strconv.Itoa(p.AccountID),
		"tax_class_id": strconv.Itoa(p.TaxClassID),
	}
}

func (r *TaxClassByIdGet) PathParams() *TaxClassByIdGetPathParams {
	return r.pathParams
}

func (r *TaxClassByIdGet) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *TaxClassByIdGet) SetMethod(method string) {
	r.method = method
}

func (r *TaxClassByIdGet) Method() string {
	return r.method
}

func (r TaxClassByIdGet) NewRequestBody() TaxClassByIdGetBody {
	return TaxClassByIdGetBody{}
}

type TaxClassByIdGetBody struct{}

func (r *TaxClassByIdGet) RequestBody() *TaxClassByIdGetBody {
	return nil
}

func (r *TaxClassByIdGet) RequestBodyInterface() interface{} {
	return nil
}

func (r *TaxClassByIdGet) SetRequestBody(body TaxClassByIdGetBody) {
	r.requestBody = body
}

func (r *TaxClassByIdGet) NewResponseBody() *TaxClassByIdGetResponseBody {
	return &TaxClassByIdGetResponseBody{}
}

type TaxClassByIdGetResponseBody TaxClassResp

func (r *TaxClassByIdGet) URL() *url.URL {
	u := r.client.GetEndpointURL("/API/V3/Account/{{.account_id}}/TaxClass/{{.tax_class_id}}.json", r.PathParams())
	return &u
}

func (r *TaxClassByIdGet) Do() (TaxClassByIdGetResponseBody, error) {
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
