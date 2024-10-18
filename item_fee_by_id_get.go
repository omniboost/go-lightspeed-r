package lightspeed_r

import (
	"net/http"
	"net/url"
	"strconv"

	"github.com/omniboost/go-lightspeed-r/utils"
)

func (c *Client) NewItemFeeByIdGet() ItemFeeByIdGet {
	r := ItemFeeByIdGet{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type ItemFeeByIdGet struct {
	client      *Client
	queryParams *ItemFeeByIdGetQueryParams
	pathParams  *ItemFeeByIdGetPathParams
	method      string
	headers     http.Header
	requestBody ItemFeeByIdGetBody
}

func (r ItemFeeByIdGet) NewQueryParams() *ItemFeeByIdGetQueryParams {
	return &ItemFeeByIdGetQueryParams{}
}

type ItemFeeByIdGetQueryParams struct{}

func (p ItemFeeByIdGetQueryParams) ToURLValues() (url.Values, error) {
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

func (r *ItemFeeByIdGet) QueryParams() *ItemFeeByIdGetQueryParams {
	return r.queryParams
}

func (r ItemFeeByIdGet) NewPathParams() *ItemFeeByIdGetPathParams {
	return &ItemFeeByIdGetPathParams{
		AccountID: r.client.accountID,
	}
}

type ItemFeeByIdGetPathParams struct {
	AccountID int `schema:"account_id"`
	ItemFeeID int `schema:"item_fee_id"`
}

func (p *ItemFeeByIdGetPathParams) Params() map[string]string {
	return map[string]string{
		"account_id":  strconv.Itoa(p.AccountID),
		"item_fee_id": strconv.Itoa(p.ItemFeeID),
	}
}

func (r *ItemFeeByIdGet) PathParams() *ItemFeeByIdGetPathParams {
	return r.pathParams
}

func (r *ItemFeeByIdGet) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *ItemFeeByIdGet) SetMethod(method string) {
	r.method = method
}

func (r *ItemFeeByIdGet) Method() string {
	return r.method
}

func (r ItemFeeByIdGet) NewRequestBody() ItemFeeByIdGetBody {
	return ItemFeeByIdGetBody{}
}

type ItemFeeByIdGetBody struct{}

func (r *ItemFeeByIdGet) RequestBody() *ItemFeeByIdGetBody {
	return nil
}

func (r *ItemFeeByIdGet) RequestBodyInterface() interface{} {
	return nil
}

func (r *ItemFeeByIdGet) SetRequestBody(body ItemFeeByIdGetBody) {
	r.requestBody = body
}

func (r *ItemFeeByIdGet) NewResponseBody() *ItemFeeByIdGetResponseBody {
	return &ItemFeeByIdGetResponseBody{}
}

type ItemFeeByIdGetResponseBody ItemFeeResp

func (r *ItemFeeByIdGet) URL() *url.URL {
	u := r.client.GetEndpointURL("/API/V3/Account/{{.account_id}}/ItemFee/{{.item_fee_id}}.json", r.PathParams())
	return &u
}

func (r *ItemFeeByIdGet) Do() (ItemFeeByIdGetResponseBody, error) {
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
