package client

import (
	"archive/tar"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
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

func untar(dst string, r io.Reader) error {

	gzr, err := gzip.NewReader(r)
	if err != nil {
		return err
	}
	defer gzr.Close()

	tr := tar.NewReader(gzr)

	for {
		header, err := tr.Next()

		switch {

		// if no more files are found return
		case err == io.EOF:
			return nil

		// return any other error
		case err != nil:
			return err

		// if the header is nil, just skip it (not sure how this happens)
		case header == nil:
			continue
		}

		// the target location where the dir/file should be created
		target := filepath.Join(dst, header.Name)

		// the following switch could also be done using fi.Mode(), not sure if there
		// a benefit of using one vs. the other.
		// fi := header.FileInfo()

		// check the file type
		switch header.Typeflag {

		// if its a dir and it doesn't exist create it
		case tar.TypeDir:
			if _, err := os.Stat(target); err != nil {
				if err := os.MkdirAll(target, 0755); err != nil {
					return err
				}
			}

		// if it's a file create it
		case tar.TypeReg:
			f, err := os.OpenFile(target, os.O_CREATE|os.O_RDWR, os.FileMode(header.Mode))
			if err != nil {
				return err
			}

			// copy over contents
			if _, err := io.Copy(f, tr); err != nil {
				return err
			}

			// manually close here after each file operation; defering would cause each file close
			// to wait until all operations have completed.
			f.Close()
		}
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

func (c *HarnessClient) GetNewDelegate(delegateName, installType string) ([]byte, error) {
	if installType == "" {
		return nil, fmt.Errorf("empty string not allowed for install type")
	}
	newDelegateInitURL := fmt.Sprintf("%s/gateway/api/setup/delegates/downloadUrl", c.clientURL)
	req, err := http.NewRequest(http.MethodGet, newDelegateInitURL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.token))
	q := req.URL.Query()
	q.Add("accountId", c.accountID)
	req.URL.RawQuery = q.Encode()
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	hc := HarnessDelegateResponse{}
	err = json.NewDecoder(resp.Body).Decode(&hc)
	if err != nil {
		return nil, err
	}
	downloadURL, err := hc.GetURLByInstallType(installType)
	if err != nil {
		return nil, err
	}
	downloadRequest, err := http.NewRequest(http.MethodGet, downloadURL, nil)
	if err != nil {
		return nil, err
	}
	downloadQuery := downloadRequest.URL.Query()
	downloadQuery.Add("delegate_name", delegateName)
	downloadRequest.URL.RawQuery = downloadQuery.Encode()
	downloadResp, err := c.client.Do(downloadRequest)
	if err != nil {
		return nil, err
	}
	defer downloadResp.Body.Close()
	err = untar(".", downloadResp.Body)
	if err != nil {
		return nil, err
	}
	b, err := ioutil.ReadFile("./harness-delegate-kubernetes/harness-delegate.yaml")
	if err != nil {
		return nil, err
	}
	err = os.RemoveAll("./harness-delegate/kubernetes")
	return b, err
}
