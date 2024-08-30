package adiq

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/ronymmoura/adiq-recurrence-check/internal/util"
)

type AuthResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int32  `json:"expires_in"`
}

const authUrl = "https://authorization.adiq.io/authorize/clientcredentials"

func Auth(config util.Config) (accessToken string, err error) {
	body := []byte(`{"grant_type": "client_credentials"}`)
	req, err := http.NewRequest("POST", authUrl, bytes.NewBuffer(body))
	req.Header.Set("Authorization", "Basic "+config.AdiqKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return
	}

	defer res.Body.Close()

	resBody, _ := io.ReadAll(res.Body)

	var token AuthResponse
	err = json.Unmarshal(resBody, &token)
	if err != nil {
		return
	}

	accessToken = token.AccessToken

	return
}
