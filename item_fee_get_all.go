package lightspeed_r

import (
	"net/http"
	"net/url"
	"strconv"

	"github.com/omniboost/go-lightspeed-r/utils"
)

func (c *Client) NewItemFeeGetAll() ItemFeeGetAll {
	r := ItemFeeGetAll{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type ItemFeeGetAll struct {
	client      *Client
	queryParams *ItemFeeGetAllQueryParams
	pathParams  *ItemFeeGetAllPathParams
	method      string
	headers     http.Header
	requestBody ItemFeeGetAllBody
}

func (r ItemFeeGetAll) NewQueryParams() *ItemFeeGetAllQueryParams {
	return &ItemFeeGetAllQueryParams{}
}

type ItemFeeGetAllQueryParams struct{}

func (p ItemFeeGetAllQueryParams) ToURLValues() (url.Values, error) {
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

func (r *ItemFeeGetAll) QueryParams() *ItemFeeGetAllQueryParams {
	return r.queryParams
}

func (r ItemFeeGetAll) NewPathParams() *ItemFeeGetAllPathParams {
	return &ItemFeeGetAllPathParams{
		AccountID: r.client.accountID,
	}
}

type ItemFeeGetAllPathParams struct {
	AccountID int `schema:"account_id"`
}

func (p *ItemFeeGetAllPathParams) Params() map[string]string {
	return map[string]string{
		"account_id": strconv.Itoa(p.AccountID),
	}
}

func (r *ItemFeeGetAll) PathParams() *ItemFeeGetAllPathParams {
	return r.pathParams
}

func (r *ItemFeeGetAll) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *ItemFeeGetAll) SetMethod(method string) {
	r.method = method
}

func (r *ItemFeeGetAll) Method() string {
	return r.method
}

func (r ItemFeeGetAll) NewRequestBody() ItemFeeGetAllBody {
	return ItemFeeGetAllBody{}
}

type ItemFeeGetAllBody struct{}

func (r *ItemFeeGetAll) RequestBody() *ItemFeeGetAllBody {
	return nil
}

func (r *ItemFeeGetAll) RequestBodyInterface() interface{} {
	return nil
}

func (r *ItemFeeGetAll) SetRequestBody(body ItemFeeGetAllBody) {
	r.requestBody = body
}

func (r *ItemFeeGetAll) NewResponseBody() *ItemFeeGetAllResponseBody {
	return &ItemFeeGetAllResponseBody{}
}

type ItemFeeGetAllResponseBody ItemFeeResp

func (r *ItemFeeGetAll) URL() *url.URL {
	u := r.client.GetEndpointURL("/API/V3/Account/{{.account_id}}/ItemFee.json", r.PathParams())
	return &u
}

func (r *ItemFeeGetAll) Do() (ItemFeeGetAllResponseBody, error) {
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
