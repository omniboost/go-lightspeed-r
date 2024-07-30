package lightspeed_r

import (
	"net/http"
	"net/url"
	"strconv"

	"github.com/omniboost/go-lightspeed-r/utils"
)

func (c *Client) NewPaymentTypesGetAll() PaymentTypesGetAll {
	r := PaymentTypesGetAll{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type PaymentTypesGetAll struct {
	client      *Client
	queryParams *PaymentTypesGetAllQueryParams
	pathParams  *PaymentTypesGetAllPathParams
	method      string
	headers     http.Header
	requestBody PaymentTypesGetAllBody
}

func (r PaymentTypesGetAll) NewQueryParams() *PaymentTypesGetAllQueryParams {
	return &PaymentTypesGetAllQueryParams{}
}

type PaymentTypesGetAllQueryParams struct{}

func (p PaymentTypesGetAllQueryParams) ToURLValues() (url.Values, error) {
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

func (r *PaymentTypesGetAll) QueryParams() *PaymentTypesGetAllQueryParams {
	return r.queryParams
}

func (r PaymentTypesGetAll) NewPathParams() *PaymentTypesGetAllPathParams {
	return &PaymentTypesGetAllPathParams{
		AccountID: r.client.accountID,
	}
}

type PaymentTypesGetAllPathParams struct {
	AccountID int `schema:"account_id"`
}

func (p *PaymentTypesGetAllPathParams) Params() map[string]string {
	return map[string]string{
		"account_id": strconv.Itoa(p.AccountID),
	}
}

func (r *PaymentTypesGetAll) PathParams() *PaymentTypesGetAllPathParams {
	return r.pathParams
}

func (r *PaymentTypesGetAll) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *PaymentTypesGetAll) SetMethod(method string) {
	r.method = method
}

func (r *PaymentTypesGetAll) Method() string {
	return r.method
}

func (r PaymentTypesGetAll) NewRequestBody() PaymentTypesGetAllBody {
	return PaymentTypesGetAllBody{}
}

type PaymentTypesGetAllBody struct{}

func (r *PaymentTypesGetAll) RequestBody() *PaymentTypesGetAllBody {
	return nil
}

func (r *PaymentTypesGetAll) RequestBodyInterface() interface{} {
	return nil
}

func (r *PaymentTypesGetAll) SetRequestBody(body PaymentTypesGetAllBody) {
	r.requestBody = body
}

func (r *PaymentTypesGetAll) NewResponseBody() *PaymentTypesGetAllResponseBody {
	return &PaymentTypesGetAllResponseBody{}
}

type PaymentTypesGetAllResponseBody PaymentTypesResp

func (r *PaymentTypesGetAll) URL() *url.URL {
	u := r.client.GetEndpointURL("/API/V3/Account/{{.account_id}}/PaymentType.json", r.PathParams())
	return &u
}

func (r *PaymentTypesGetAll) Do() (PaymentTypesGetAllResponseBody, error) {
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
