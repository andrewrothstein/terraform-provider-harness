package client

import (
	"encoding/json"
	"fmt"
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

type HarnessDelegateResponse struct {
	Resource map[string]string `json:"resource"`
}

func (hd HarnessDelegateResponse) GetURLByInstallType(installType string) (string, error) {
	switch inst := installType; inst {
	case "KUBERNETES_YAML":
		return hd.Resource["kubernetesUrl"], nil
	default:
	}
	return "", fmt.Errorf("no install type found for %s", installType)
}

func (c *HarnessClient) GetNewDelegate(delegateName, installType string) (string, error) {
	if installType == "" {
		return "", fmt.Errorf("empty string not allowed for install type")
	}
	newDelegateInitURL := fmt.Sprintf("%s/gateway/api/setup/delegates/downloadUrl", c.clientURL)
	req, err := http.NewRequest(http.MethodGet, newDelegateInitURL, nil)
	if err != nil {
		return "", err
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.token))
	q := req.URL.Query()
	q.Add("accountId", c.accountID)
	req.URL.RawQuery = q.Encode()
	resp, err := c.client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	hc := HarnessDelegateResponse{}
	err = json.NewDecoder(resp.Body).Decode(&hc)
	if err != nil {
		return "", err
	}
	downloadURL, err := hc.GetURLByInstallType(installType)
	if err != nil {
		return "", err
	}
	downloadRequest, err := http.NewRequest(http.MethodGet, downloadURL, nil)
	if err != nil {
		return "", err
	}
	downloadQuery := downloadRequest.URL.Query()
	downloadQuery.Add("delegateName", delegateName)
	downloadRequest.URL.RawQuery = downloadQuery.Encode()
	downloadResp, err := c.client.Do(downloadRequest)
	if err != nil {
		return "", err
	}
	return downloadResp.Status, err
}
