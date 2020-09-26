package provider

import (
	"bytes"
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	apimachineryschema "k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"log"
)

var k8sPrefix = "kubernetes.0."

func newKubeClient(configData *schema.ResourceData) (*kubernetes.Clientset, *rest.Config, error) {
	overrides := &clientcmd.ConfigOverrides{}
	loader := &clientcmd.ClientConfigLoadingRules{}
	clusterCaCertificate := configData.Get(fmt.Sprintf("%scluster_ca_certificate", k8sPrefix)).(string)
	overrides.ClusterInfo.CertificateAuthorityData = bytes.NewBufferString(clusterCaCertificate).Bytes()
	hostString := configData.Get(fmt.Sprintf("%scluster_ca_certificate", k8sPrefix)).(string)
	// hard coding TLS true cause our server only has a true
	host, _, err := rest.DefaultServerURL(hostString, "", apimachineryschema.GroupVersion{}, true)
	if err != nil {
		return nil, nil, err
	}
	overrides.ClusterInfo.Server = host.String()
	overrides.AuthInfo.Token = configData.Get(fmt.Sprintf("%stoken", k8sPrefix)).(string)
	fmt.Println(overrides)
	client := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(loader, overrides)
	if client == nil {
		return nil, nil, fmt.Errorf("failed to initialize kubernetes config")
	}
	log.Printf("[INFO] Successfully initialized kubernetes config")
	c, err := client.ClientConfig()
	if err != nil {
		return nil, nil, err
	}
	k, err := kubernetes.NewForConfig(c)
	if err != nil {
		return nil, nil, err
	}
	return k, c, nil

}
