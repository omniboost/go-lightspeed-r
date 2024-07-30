package lightspeed_r

import (
	"net/http"
	"net/url"
	"strconv"

	"github.com/omniboost/go-lightspeed-r/utils"
)

func (c *Client) NewTaxClassesGetAll() TaxClassesGetAll {
	r := TaxClassesGetAll{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type TaxClassesGetAll struct {
	client      *Client
	queryParams *TaxClassesGetAllQueryParams
	pathParams  *TaxClassesGetAllPathParams
	method      string
	headers     http.Header
	requestBody TaxClassesGetAllBody
}

func (r TaxClassesGetAll) NewQueryParams() *TaxClassesGetAllQueryParams {
	return &TaxClassesGetAllQueryParams{}
}

type TaxClassesGetAllQueryParams struct{}

func (p TaxClassesGetAllQueryParams) ToURLValues() (url.Values, error) {
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

func (r *TaxClassesGetAll) QueryParams() *TaxClassesGetAllQueryParams {
	return r.queryParams
}

func (r TaxClassesGetAll) NewPathParams() *TaxClassesGetAllPathParams {
	return &TaxClassesGetAllPathParams{
		AccountID: r.client.accountID,
	}
}

type TaxClassesGetAllPathParams struct {
	AccountID int `schema:"account_id"`
}

func (p *TaxClassesGetAllPathParams) Params() map[string]string {
	return map[string]string{
		"account_id": strconv.Itoa(p.AccountID),
	}
}

func (r *TaxClassesGetAll) PathParams() *TaxClassesGetAllPathParams {
	return r.pathParams
}

func (r *TaxClassesGetAll) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *TaxClassesGetAll) SetMethod(method string) {
	r.method = method
}

func (r *TaxClassesGetAll) Method() string {
	return r.method
}

func (r TaxClassesGetAll) NewRequestBody() TaxClassesGetAllBody {
	return TaxClassesGetAllBody{}
}

type TaxClassesGetAllBody struct{}

func (r *TaxClassesGetAll) RequestBody() *TaxClassesGetAllBody {
	return nil
}

func (r *TaxClassesGetAll) RequestBodyInterface() interface{} {
	return nil
}

func (r *TaxClassesGetAll) SetRequestBody(body TaxClassesGetAllBody) {
	r.requestBody = body
}

func (r *TaxClassesGetAll) NewResponseBody() *TaxClassesGetAllResponseBody {
	return &TaxClassesGetAllResponseBody{}
}

type TaxClassesGetAllResponseBody TaxClassesResp

func (r *TaxClassesGetAll) URL() *url.URL {
	u := r.client.GetEndpointURL("/API/V3/Account/{{.account_id}}/TaxClass.json", r.PathParams())
	return &u
}

func (r *TaxClassesGetAll) Do() (TaxClassesGetAllResponseBody, error) {
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
