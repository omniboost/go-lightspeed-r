package lightspeed_r

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/omniboost/go-lightspeed-r/utils"
)

func (c *Client) NewSalesGet() SalesGet {
	r := SalesGet{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type SalesGet struct {
	client      *Client
	queryParams *SalesGetQueryParams
	pathParams  *SalesGetPathParams
	method      string
	headers     http.Header
	requestBody SalesGetBody
}

func (r SalesGet) NewQueryParams() *SalesGetQueryParams {
	return &SalesGetQueryParams{
		LoadRelations: LoadRelations{},
		Sort:          "completeTime",
	}
}

type SalesGetQueryParams struct {
	LoadRelations LoadRelations `schema:"load_relations,omitempty"`
	Sort          string        `schema:"sort,omitempty"`
}

func (p SalesGetQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	encoder.RegisterEncoder(Date{}, utils.EncodeSchemaMarshaler)
	encoder.RegisterEncoder(DateTime{}, utils.EncodeSchemaMarshaler)
	encoder.RegisterEncoder(LoadRelations{}, utils.EncodeSchemaMarshaler)
	params := url.Values{}

	fmt.Println(p.LoadRelations)

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	// fmt.Println(params)
	return params, nil
}

func (r *SalesGet) QueryParams() *SalesGetQueryParams {
	return r.queryParams
}

func (r SalesGet) NewPathParams() *SalesGetPathParams {
	return &SalesGetPathParams{
		AccountID: r.client.accountID,
	}
}

type SalesGetPathParams struct {
	AccountID int `schema:"account_id"`
}

func (p *SalesGetPathParams) Params() map[string]string {
	return map[string]string{
		"account_id": strconv.Itoa(p.AccountID),
	}
}

func (r *SalesGet) PathParams() *SalesGetPathParams {
	return r.pathParams
}

func (r *SalesGet) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *SalesGet) SetMethod(method string) {
	r.method = method
}

func (r *SalesGet) Method() string {
	return r.method
}

func (r SalesGet) NewRequestBody() SalesGetBody {
	return SalesGetBody{}
}

type SalesGetBody struct{}

func (r *SalesGet) RequestBody() *SalesGetBody {
	return nil
}

func (r *SalesGet) RequestBodyInterface() interface{} {
	return nil
}

func (r *SalesGet) SetRequestBody(body SalesGetBody) {
	r.requestBody = body
}

func (r *SalesGet) NewResponseBody() *SalesGetResponseBody {
	return &SalesGetResponseBody{}
}

type SalesGetResponseBody SalesResp

func (r *SalesGet) URL() *url.URL {
	u := r.client.GetEndpointURL("/API/V3/Account/{{.account_id}}/Sale.json", r.PathParams())
	return &u
}

func (r *SalesGet) Do() (SalesGetResponseBody, error) {
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
