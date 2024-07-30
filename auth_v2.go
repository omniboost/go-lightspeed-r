package lightspeed_r

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/omniboost/go-lightspeed-r/utils"
)

func (c *Client) NewAuthRequest() AuthRequest {
	r := AuthRequest{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type AuthRequest struct {
	client      *Client
	queryParams *AuthRequestQueryParams
	pathParams  *AuthRequestPathParams
	method      string
	headers     http.Header
	requestBody AuthRequestBody
}

func (r AuthRequest) NewQueryParams() *AuthRequestQueryParams {
	return &AuthRequestQueryParams{}
}

type AuthRequestQueryParams struct{}

func (p AuthRequestQueryParams) ToURLValues() (url.Values, error) {
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

func (r *AuthRequest) QueryParams() *AuthRequestQueryParams {
	return r.queryParams
}

func (r AuthRequest) NewPathParams() *AuthRequestPathParams {
	return &AuthRequestPathParams{}
}

type AuthRequestPathParams struct{}

func (p *AuthRequestPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *AuthRequest) PathParams() *AuthRequestPathParams {
	return r.pathParams
}

func (r *AuthRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *AuthRequest) SetMethod(method string) {
	r.method = method
}

func (r *AuthRequest) Method() string {
	return r.method
}

func (r AuthRequest) NewRequestBody() AuthRequestBody {
	return AuthRequestBody{}
}

type AuthRequestBody struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	RefreshToken string `json:"refresh_token"`
	GrantType    string `json:"grant_type"`
}

func (r *AuthRequest) RequestBody() *AuthRequestBody {
	return &r.requestBody
}

func (r *AuthRequest) RequestBodyInterface() interface{} {
	formData := url.Values{
		"grant_type":    {"refresh_token"},
		"client_id":     {r.client.ClientID()},
		"client_secret": {r.client.ClientSecret()},
		"refresh_token": {r.client.RefreshToken()},
	}

	return strings.NewReader(formData.Encode())
}

func (r *AuthRequest) SetRequestBody(body AuthRequestBody) {
	r.requestBody = body
}

func (r *AuthRequest) NewResponseBody() *AuthRequestResponseBody {
	return &AuthRequestResponseBody{}
}

type AuthRequestResponseBody BearerToken

func (r *AuthRequest) URL() *url.URL {
	u := r.client.GetAuthEndpointURL("/ls_retail/oauth2/token", r.PathParams())
	return &u
}

func (r *AuthRequest) Do() (AuthRequestResponseBody, error) {
	// Create http request
	req, err := r.client.NewRequest(nil, r)
	if err != nil {
		return *r.NewResponseBody(), err
	}

	req.Header.Set("content-type", "application/x-www-form-urlencoded")

	// Process query parameters
	err = utils.AddQueryParamsToRequest(r.QueryParams(), req, false)
	if err != nil {
		return *r.NewResponseBody(), err
	}

	responseBody := r.NewResponseBody()
	_, err = r.client.Do(req, responseBody)
	return *responseBody, err
}
