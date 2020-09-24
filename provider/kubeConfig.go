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

func newKubeClient(configData *schema.ResourceData) (*kubernetes.Clientset, *rest.Config) {
	overrides := &clientcmd.ConfigOverrides{}
	loader := &clientcmd.ClientConfigLoadingRules{}
	clusterCaCertificate := configData.Get(fmt.Sprintf("%scluster_ca_certificate", k8sPrefix)).(string)
	overrides.ClusterInfo.CertificateAuthorityData = bytes.NewBufferString(clusterCaCertificate).Bytes()
	hostString := configData.Get(fmt.Sprintf("%scluster_ca_certificate", k8sPrefix)).(string)
	// hard coding TLS true cause our server only has a true
	host, _, err := rest.DefaultServerURL(hostString, "", apimachineryschema.GroupVersion{}, true)
	if err != nil {
		return nil, nil
	}
	overrides.ClusterInfo.Server = host.String()
	overrides.AuthInfo.Token = configData.Get(fmt.Sprintf("%stoken", k8sPrefix)).(string)
	client := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(loader, overrides)
	if client == nil {
		log.Printf("[ERROR] Failed to initialize kubernetes config")
		return nil, nil
	}
	log.Printf("[INFO] Successfully initialized kubernetes config")
	c, _ := client.ClientConfig()
	k, _ := kubernetes.NewForConfig(c)
	return k, c

}
