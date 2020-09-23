package client

import (
	"net/http"
)

type HarnessClient struct {
	clientURL string
	token     string
	accountID string
	client    *http.Client
}

func NewClient(clientURL, token, accountID string) *HarnessClient {
	httpClient := &http.Client{}
	return &HarnessClient{
		client:    httpClient,
		clientURL: clientURL,
		token:     token,
		accountID: accountID,
	}
}
