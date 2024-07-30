package lightspeed_r

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"path"
	"strconv"
	"strings"
	"text/template"

	"github.com/pkg/errors"
)

const (
	libraryVersion = "0.0.1"
	userAgent      = "go-lightspeed-r" + libraryVersion
	mediaType      = "application/json"
	charset        = "utf-8"
)

var (
	BaseURL = url.URL{
		Scheme: "https",
		Host:   "api.lightspeedapp.com",
		Path:   "",
	}
)

// NewClient returns a new Exact Globe Client client
func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	// oauthConfig := lightspeed_r.NewOauth2Config()
	// oauthConfig.ClientID = clientID
	// oauthConfig.ClientSecret = clientSecret

	// accessToken, err :=
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// tokenString := fmt.Sprintf("AccessToken: %s\n",
	// 	accessToken)

	// fmt.Println(tokenString)

	// token := &oauth2.Token{
	// 	RefreshToken: refreshToken,
	// 	AccessToken:  accessToken,
	// }

	client := &Client{}

	client.SetHTTPClient(httpClient)
	client.SetBaseURL(BaseURL)
	client.SetDebug(false)
	client.SetUserAgent(userAgent)
	client.SetMediaType(mediaType)
	client.SetCharset(charset)

	// if err := client.SetAccountID(); err != nil {
	// 	return nil
	// }

	return client
}

// Client manages communication with Exact Globe Client
type Client struct {
	// HTTP client used to communicate with the Client.
	http *http.Client

	debug   bool
	baseURL url.URL
	authURL url.URL

	// credentials
	refreshToken string
	clientSecret string
	clientID     string
	accessToken  string
	accountID    int

	// User agent for client
	userAgent string

	mediaType             string
	charset               string
	disallowUnknownFields bool

	// Optional function called after every successful request made to the DO Clients
	beforeRequestDo    BeforeRequestDoCallback
	onRequestCompleted RequestCompletionCallback
}

type BeforeRequestDoCallback func(*http.Client, *http.Request, interface{})

// RequestCompletionCallback defines the type of the request callback function
type RequestCompletionCallback func(*http.Request, *http.Response)

func (c *Client) SetHTTPClient(client *http.Client) {
	c.http = client
}

func (c Client) Debug() bool {
	c.debug = true

	return c.debug
}

func (c *Client) SetDebug(debug bool) {
	c.debug = debug
}

func (c Client) AccessToken() string {
	return c.accessToken
}

func (c *Client) SetAccessToken(accessToken string) {
	c.accessToken = accessToken
}

func (c Client) BaseURL() url.URL {
	return c.baseURL
}

func (c *Client) SetBaseURL(baseURL url.URL) {
	c.baseURL = baseURL
}

func (c Client) AuthURL() url.URL {
	return c.authURL
}

func (c *Client) SetAuthURL(authURL url.URL) {
	c.authURL = authURL
}

func (c *Client) SetMediaType(mediaType string) {
	c.mediaType = mediaType
}

func (c Client) MediaType() string {
	return mediaType
}

func (c *Client) SetCharset(charset string) {
	c.charset = charset
}

func (c Client) Charset() string {
	return charset
}

func (c *Client) SetUserAgent(userAgent string) {
	c.userAgent = userAgent
}

func (c Client) UserAgent() string {
	return userAgent
}

func (c *Client) SetDisallowUnknownFields(disallowUnknownFields bool) {
	c.disallowUnknownFields = disallowUnknownFields
}

func (c *Client) SetBeforeRequestDo(fun BeforeRequestDoCallback) {
	c.beforeRequestDo = fun
}

func (c Client) ClientID() string {
	return c.clientID
}

func (c *Client) SetClientID(clientID string) {
	c.clientID = clientID
}

func (c Client) ClientSecret() string {
	return c.clientSecret
}

func (c *Client) SetClientSecret(clientSecret string) {
	c.clientSecret = clientSecret
}

func (c Client) RefreshToken() string {
	return c.refreshToken
}

func (c *Client) SetRefreshToken(refreshToken string) {
	c.refreshToken = refreshToken
}

func (c Client) AccountID() int {
	return c.accountID
}

func (c *Client) SetAccountID(accountID int) {
	c.accountID = accountID
}

func (c *Client) GetBearerToken() (string, error) {
	if c.accessToken == "" {
		var err error
		c.accessToken, err = c.NewBearerToken()
		if err != nil {
			return "", err
		}
	}

	return c.accessToken, nil
}

func (c *Client) NewBearerToken() (string, error) {
	req := c.NewAuthRequest()

	resp, err := req.Do()
	if err != nil {
		return "", err
	}

	return resp.AccessToken, nil
}

func (c *Client) GetAccountID() (int, error) {
	if c.accountID == 0 {
		var err error
		c.accountID, err = c.NewAccountID()
		if err != nil {
			return 0, err
		}
	}

	return c.accountID, nil
	// accountID, err := strconv.Atoi(c.accountID)
	// if err != nil {
	// 	return 0, err
	// }
	// return accountID, nil
}

func (c *Client) NewAccountID() (int, error) {
	req := c.NewAccountGet()

	resp, err := req.Do()
	if err != nil {
		return 0, err
	}

	accountID, err := strconv.Atoi(resp.Account.AccountID)
	if err != nil {
		return 0, err
	}

	return accountID, nil
}

func (c *Client) GetURL(baseUrl url.URL, p string, pathParams PathParams) url.URL {
	clientURL := baseUrl

	parsed, err := url.Parse(p)
	if err != nil {
		log.Fatal(err)
	}
	q := clientURL.Query()
	for k, vv := range parsed.Query() {
		for _, v := range vv {
			q.Add(k, v)
		}
	}
	clientURL.RawQuery = q.Encode()

	clientURL.Path = path.Join(clientURL.Path, parsed.Path)

	tmpl, err := template.New("path").Parse(clientURL.Path)
	if err != nil {
		log.Fatal(err)
	}

	buf := new(bytes.Buffer)
	params := pathParams.Params()
	// params["administration_id"] = c.Administration()
	err = tmpl.Execute(buf, params)
	if err != nil {
		log.Fatal(err)
	}

	clientURL.Path = buf.String()
	return clientURL
}

func (c *Client) GetEndpointURL(p string, pathParams PathParams) url.URL {
	return c.GetURL(c.BaseURL(), p, pathParams)
}

func (c *Client) GetAuthEndpointURL(p string, pathParams PathParams) url.URL {
	return c.GetURL(c.AuthURL(), p, pathParams)
}

func (c *Client) NewRequest(ctx context.Context, req Request) (*http.Request, error) {
	// convert body struct to json
	var body io.Reader
	if req.RequestBodyInterface() != nil {
		if r, ok := req.RequestBodyInterface().(io.Reader); ok {
			body = r
		} else if bb, ok := req.RequestBodyInterface().([]byte); ok {
			body = bytes.NewReader(bb)
		} else {
			buf := new(bytes.Buffer)
			err := json.NewEncoder(buf).Encode(req.RequestBodyInterface())
			if err != nil {
				return nil, err
			}
			body = buf
		}
	}

	// create new http request
	r, err := http.NewRequest(req.Method(), req.URL().String(), body)
	if err != nil {
		return nil, err
	}

	// values := url.Values{}
	// err = utils.AddURLValuesToRequest(values, req, true)
	// if err != nil {
	// 	return nil, err
	// }

	// optionally pass along context
	if ctx != nil {
		r = r.WithContext(ctx)
	}

	// set other headers
	r.Header.Add("Content-Type", fmt.Sprintf("%s; charset=%s", c.MediaType(), c.Charset()))
	r.Header.Add("Accept", c.MediaType())
	r.Header.Add("User-Agent", c.UserAgent())
	r.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.AccessToken()))

	return r, nil
}

// Do sends an Client request and returns the Client response. The Client response is json decoded and stored in the value
// pointed to by v, or returned as an error if an Client error has occurred. If v implements the io.Writer interface,
// the raw response will be written to v, without attempting to decode it.
func (c *Client) Do(req *http.Request, body interface{}) (*http.Response, error) {
	if c.beforeRequestDo != nil {
		c.beforeRequestDo(c.http, req, body)
	}

	if c.debug == true {
		dump, _ := httputil.DumpRequestOut(req, true)
		log.Println(string(dump))
	}

	httpResp, err := c.http.Do(req)
	if err != nil {
		return nil, err
	}

	if c.onRequestCompleted != nil {
		c.onRequestCompleted(req, httpResp)
	}

	// close body io.Reader
	defer func() {
		if rerr := httpResp.Body.Close(); err == nil {
			err = rerr
		}
	}()

	if c.debug == true {
		dump, _ := httputil.DumpResponse(httpResp, true)
		log.Println(string(dump))
	}

	// check the provided interface parameter
	if httpResp == nil {
		return httpResp, nil
	}

	if body == nil {
		return httpResp, err
	}

	if httpResp.ContentLength == 0 {
		return httpResp, nil
	}

	status := &ErrorResponse{Response: httpResp}
	// exResp := &ExceptionResponse{Response: httpResp}
	err = c.Unmarshal(httpResp.Body, []any{body}, []any{status})
	if err != nil {
		return httpResp, err
	}

	if status.Error() != "" {
		return httpResp, status
	}

	// check if the response isn't an error
	err = CheckResponse(httpResp)
	if err != nil {
		return httpResp, err
	}

	return httpResp, nil
}

func (c *Client) Unmarshal(r io.Reader, vv []interface{}, optionalVv []interface{}) error {
	if len(vv) == 0 && len(optionalVv) == 0 {
		return nil
	}

	b, err := io.ReadAll(r)
	if err != nil {
		return err
	}

	for _, v := range vv {
		r := bytes.NewReader(b)
		dec := json.NewDecoder(r)

		err := dec.Decode(v)
		if err != nil && err != io.EOF {
			return errors.WithStack((err))
		}
	}

	for _, v := range optionalVv {
		r := bytes.NewReader(b)
		dec := json.NewDecoder(r)

		_ = dec.Decode(v)
	}

	return nil
}

// CheckResponse checks the Client response for errors, and returns them if
// present. A response is considered an error if it has a status code outside
// the 200 range. Client error responses are expected to have either no response
// body, or a json response body that maps to ErrorResponse. Any other response
// body will be silently ignored.
func CheckResponse(r *http.Response) error {
	errorResponse := &ErrorResponse{Response: r}

	// Don't check content-lenght: a created response, for example, has no body
	// if r.Header.Get("Content-Length") == "0" {
	// 	errorResponse.Errors.Message = r.Status
	// 	return errorResponse
	// }

	if c := r.StatusCode; c >= 200 && c <= 299 {
		return nil
	}

	// read data and copy it back
	data, err := io.ReadAll(r.Body)
	r.Body = io.NopCloser(bytes.NewReader(data))
	if err != nil {
		return errorResponse
	}

	err = checkContentType(r)
	if err != nil {
		return errors.WithStack(err)
	}

	if r.ContentLength == 0 {
		return errors.New("response body is empty")
	}

	// convert json to struct
	if len(data) != 0 {
		err = json.Unmarshal(data, &errorResponse)
		if err != nil {
			return errors.WithStack(err)
		}
	}

	if errorResponse.Message != "" {
		return errorResponse
	}

	return nil
}

// type StatusResponse struct {
// 	// HTTP response that caused this error
// 	Response *http.Response

// 	Page Page `json:"page"`
// }

// func (r *StatusResponse) Error() string {
// 	if r.Page.StatusCode != 0 && (r.Page.StatusCode < 200 || r.Page.StatusCode > 299) {
// 		return fmt.Sprintf("Status %d: %s %s", r.Page.StatusCode, r.Page.Status, r.Page.Message)
// 	}

// 	return ""
// }

type ErrorResponse struct {
	// HTTP response that caused this error
	Response *http.Response

	HttpCode    string `json:"httpCode"`
	HttpMessage string `json:"httpMessage"`
	Message     string `json:"Message"`
	ErrorClass  string `json:"errorClass"`
}

func (r *ErrorResponse) Error() string {
	// if r.HttpCode != 0 && (r.HttpCode < 200 || r.HttpCode > 299) {
	// 	return fmt.Sprintf("Status %d: %s %s %s", r.HttpCode, r.HttpMessage, r.Message, r.ErrorClass)
	// }
	if r.HttpCode != "" {
		return fmt.Sprintf("Status %s: %s %s %s", r.HttpCode, r.HttpMessage, r.Message, r.ErrorClass)
	}
	return ""
	// return fmt.Sprintf("Status %s: %s %s %s", r.HttpCode, r.HttpMessage, r.Message, r.ErrorClass)
}

func checkContentType(response *http.Response) error {
	header := response.Header.Get("Content-Type")
	contentType := strings.Split(header, ";")[0]
	if contentType != mediaType {
		return fmt.Errorf("Expected Content-Type \"%s\", got \"%s\"", mediaType, contentType)
	}

	return nil
}
