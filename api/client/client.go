package client

import (
	"archive/tar"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
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

// Untar takes a destination path and a reader; a tar reader loops over the tarfile
// creating the file structure at 'dst' along the way, and writing any files
func Untar(dst string, r io.Reader) error {

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
		log.Printf("Kubernetes URL: %s", hd.Resource["kuberentesUrl"])
		return hd.Resource["kubernetesUrl"], nil
	default:
	}
	return "", fmt.Errorf("no install type found for %s", installType)
}

//https://app.harness.io/gateway/api/setup/delegates/downloadUrl
func (c *HarnessClient) GetNewDelegate(delegateName, installType string) ([]byte, error) {
	if installType == "" {
		return nil, fmt.Errorf("empty string not allowed for install type")
	}
	url := fmt.Sprintf("%s/gateway/api/setup/delegates/downloadUrl?accountId=%s", c.clientURL, c.accountID)
	req, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.token))
	resp, err := c.client.Do(req)
	defer resp.Body.Close()

	log.Printf("Harness Response Code: %s", resp.Status)

	hc := HarnessDelegateResponse{}
	err = json.NewDecoder(resp.Body).Decode(&hc)
	if err != nil {
		return nil, err
	}
	downloadURL, err := hc.GetURLByInstallType(installType)
	if err != nil {
		return nil, err
	}
	downloadURL = fmt.Sprintf("%s&delegateName=%s", downloadURL, delegateName)
	log.Printf("DownloadURL: %s", downloadURL)
	downloadRequest, err := http.NewRequest(http.MethodGet, downloadURL, nil)
	if err != nil {
		return nil, err
	}
	downloadResp, err := c.client.Do(downloadRequest)
	if err != nil {
		return nil, err
	}
	defer downloadResp.Body.Close()
	out, err := os.Create("./harness-delegate-kubernetes.tar.gz")
	if err != nil {
		return nil, err
	}
	defer out.Close()
	// Write the body to file
	_, err = io.Copy(out, downloadResp.Body)
	if err != nil {
		return nil, err
	}
	r, err := os.Open("./harness-delegate-kubernetes.tar.gz")
	if err != nil {
		return nil, err
	}
	defer os.Remove("./harness-delegate-kubernetes.tar.gz")
	err = Untar(".", r)
	defer os.RemoveAll("./harness-delegate-kubernetes")
	if err != nil {
		return nil, err
	}
	b, err := ioutil.ReadFile("./harness-delegate-kubernetes/harness-delegate.yaml")
	if err != nil {
		return nil, err
	}
	return b, err
}
