package lightspeed_r

import (
	"net/http"
	"net/url"
	"strconv"

	"github.com/omniboost/go-lightspeed-r/utils"
)

func (c *Client) NewPaymentTypeByIdGet() PaymentTypeByIdGet {
	r := PaymentTypeByIdGet{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type PaymentTypeByIdGet struct {
	client      *Client
	queryParams *PaymentTypeByIdGetQueryParams
	pathParams  *PaymentTypeByIdGetPathParams
	method      string
	headers     http.Header
	requestBody PaymentTypeByIdGetBody
}

func (r PaymentTypeByIdGet) NewQueryParams() *PaymentTypeByIdGetQueryParams {
	return &PaymentTypeByIdGetQueryParams{}
}

type PaymentTypeByIdGetQueryParams struct{}

func (p PaymentTypeByIdGetQueryParams) ToURLValues() (url.Values, error) {
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

func (r *PaymentTypeByIdGet) QueryParams() *PaymentTypeByIdGetQueryParams {
	return r.queryParams
}

func (r PaymentTypeByIdGet) NewPathParams() *PaymentTypeByIdGetPathParams {
	return &PaymentTypeByIdGetPathParams{
		AccountID: r.client.accountID,
	}
}

type PaymentTypeByIdGetPathParams struct {
	AccountID     int `schema:"account_id"`
	PaymentTypeID int `schema:"payment_type_id"`
}

func (p *PaymentTypeByIdGetPathParams) Params() map[string]string {
	return map[string]string{
		"account_id":      strconv.Itoa(p.AccountID),
		"payment_type_id": strconv.Itoa(p.PaymentTypeID),
	}
}

func (r *PaymentTypeByIdGet) PathParams() *PaymentTypeByIdGetPathParams {
	return r.pathParams
}

func (r *PaymentTypeByIdGet) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *PaymentTypeByIdGet) SetMethod(method string) {
	r.method = method
}

func (r *PaymentTypeByIdGet) Method() string {
	return r.method
}

func (r PaymentTypeByIdGet) NewRequestBody() PaymentTypeByIdGetBody {
	return PaymentTypeByIdGetBody{}
}

type PaymentTypeByIdGetBody struct{}

func (r *PaymentTypeByIdGet) RequestBody() *PaymentTypeByIdGetBody {
	return nil
}

func (r *PaymentTypeByIdGet) RequestBodyInterface() interface{} {
	return nil
}

func (r *PaymentTypeByIdGet) SetRequestBody(body PaymentTypeByIdGetBody) {
	r.requestBody = body
}

func (r *PaymentTypeByIdGet) NewResponseBody() *PaymentTypeByIdGetResponseBody {
	return &PaymentTypeByIdGetResponseBody{}
}

type PaymentTypeByIdGetResponseBody PaymentTypeResp

func (r *PaymentTypeByIdGet) URL() *url.URL {
	u := r.client.GetEndpointURL("/API/V3/Account/{{.account_id}}/PaymentType/{{.payment_type_id}}.json", r.PathParams())
	return &u
}

func (r *PaymentTypeByIdGet) Do() (PaymentTypeByIdGetResponseBody, error) {
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
