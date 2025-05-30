package lightspeed_r

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

func GetBerearToken(clientID, clientSecret, refreshToken, tokenURL string) (*BearerToken, error) {
	q := url.Values{}
	q.Set("grant_type", "refresh_token")
	q.Add("client_id", clientID)
	q.Add("client_secret", clientSecret)
	q.Add("refresh_token", refreshToken)

	req, err := http.NewRequest("POST", tokenURL, bytes.NewBufferString(q.Encode()))
	if err != nil {
		return nil, err
	}

	fmt.Println(bytes.NewBufferString(q.Encode()))

	req.Header.Set("accept", "application/json")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded;charset=UTF-8")

	httpClient := &http.Client{}
	res, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	bearerToken := BearerToken{}
	err = json.NewDecoder(res.Body).Decode(&bearerToken)
	if err != nil {
		return nil, err
	}

	return &bearerToken, nil
}
