package client

import (
	"os"
	"testing"
)

func TestGetNewDelegate(t *testing.T) {
	hc := NewClient("https://app.harness.io", os.Getenv("HARNESS_TOKEN"), "2KTQt0X9R82AEBbv9RYn_g")
	r, err := hc.GetNewDelegate("peepeepoopoo", "KUBERNETES_YAML")
	if err != nil {
		t.FailNow()
	}
	t.Log(r)
}
