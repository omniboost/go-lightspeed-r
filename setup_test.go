package lightspeed_r_test

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"testing"

	lightspeed_r "github.com/omniboost/go-lightspeed-r"
)

var (
	client *lightspeed_r.Client
)

func TestMain(m *testing.M) {
	baseURLString := os.Getenv("BASE_URL")
	clientID := os.Getenv("LIGHTSPEED_CLIENT_ID")
	clientSecret := os.Getenv("LIGHTSPEED_CLIENT_SECRET")
	refreshToken := os.Getenv("LIGHTSPEED_REFRESH_TOKEN")
	authURLString := os.Getenv("LIGHTSPEED_AUTH_URL")
	debug := os.Getenv("DEBUG")

	oauthConfig := lightspeed_r.NewOauth2Config()
	oauthConfig.ClientID = clientID
	oauthConfig.ClientSecret = clientSecret

	// if tokenURL != "" {
	// 	oauthConfig.Endpoint.TokenURL = tokenURL
	// }

	// accessToken, err := client.BearerToken()
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// tokenString := fmt.Sprintf("AccessToken: %s\n",
	// 	accessToken)

	// fmt.Println(tokenString)

	// getAccessToken, err := lightspeed_r.GetBerearToken(clientID, clientSecret, refreshToken, oauthConfig.Endpoint.TokenURL)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// tokenString := fmt.Sprintf("AccessToken: %s\nTokenType: %s\nRefreshToken: %s\nExpiry: %v\n",
	// 	getAccessToken.AccessToken, getAccessToken.TokenType, getAccessToken.RefreshToken, getAccessToken.ExpiresIn)

	// fmt.Println(tokenString)

	// token := &oauth2.Token{
	// 	RefreshToken: refreshToken,
	// 	AccessToken:  getAccessToken.AccessToken,
	// }

	// httpClient := oauthConfig.Client(context.Background(), token)

	client = lightspeed_r.NewClient(nil)
	client.SetClientID(clientID)
	client.SetClientSecret(clientSecret)
	client.SetRefreshToken(refreshToken)

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

	if authURLString != "" {
		authURL, err := url.Parse(authURLString)
		if err != nil {
			log.Fatal(err)
		}
		client.SetAuthURL(*authURL)
	}

	accessToken, err := client.GetBearerToken()
	if err != nil {
		fmt.Println(err)
	}

	tokenString := fmt.Sprintf("AccessToken: %s\n",
		accessToken)

	fmt.Println(tokenString)

	accountID, err := client.GetAccountID()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("AccountID: %v\n", accountID)
	m.Run()
}
