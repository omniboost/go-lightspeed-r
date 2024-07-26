package lightspeed_r_test

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"os"
	"testing"

	lightspeed_r "github.com/omniboost/go-lightspeed-r"
	"golang.org/x/oauth2"
)

var (
	client *lightspeed_r.Client
)

func TestMain(m *testing.M) {
	baseURLString := os.Getenv("BASE_URL")
	clientID := os.Getenv("LIGHTSPEED_CLIENT_ID")
	clientSecret := os.Getenv("LIGHTSPEED_CLIENT_SECRET")
	refreshToken := os.Getenv("LIGHTSPEED_REFRESH_TOKEN")
	tokenURL := os.Getenv("LIGHTSPEED_TOKEN_URL")
	debug := os.Getenv("DEBUG")

	oauthConfig := lightspeed_r.NewOauth2Config()
	oauthConfig.ClientID = clientID
	oauthConfig.ClientSecret = clientSecret

	if tokenURL != "" {
		oauthConfig.Endpoint.TokenURL = tokenURL
	}

	getAccessToken, err := lightspeed_r.GetBerearToken(clientID, clientSecret, refreshToken, oauthConfig.Endpoint.TokenURL)
	if err != nil {
		fmt.Println(err)
	}

	tokenString := fmt.Sprintf("AccessToken: %s\nTokenType: %s\nRefreshToken: %s\nExpiry: %v\n",
		getAccessToken.AccessToken, getAccessToken.TokenType, getAccessToken.RefreshToken, getAccessToken.ExpiresIn)

	fmt.Println(tokenString)

	token := &oauth2.Token{
		RefreshToken: refreshToken,
		AccessToken:  getAccessToken.AccessToken,
	}

	httpClient := oauthConfig.Client(context.Background(), token)

	client = lightspeed_r.NewClient(httpClient)
	// client.SetToken(token)
	if debug != "" {
		client.SetDebug(true)
	}

	if baseURLString != "" {
		baseURL, err := url.Parse(baseURLString)
		if err != nil {
			log.Fatal(err)
		}
		client.SetBaseURL(*baseURL)
	}

	resp, err := client.GetAccountId()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("AccountID: %v\n", resp)
	m.Run()
}
