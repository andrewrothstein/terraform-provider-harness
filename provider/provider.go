package provider

import (
	"fmt"
	"github.com/eahrend/terraform-harness-provider/api/client"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"sync"
)

type Meta struct {
	data          *schema.ResourceData
	kubeClient    *kubernetes.Clientset
	harnessClient *client.HarnessClient
	restConfig    *rest.Config
	// Used to lock some operations
	sync.Mutex
}

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"client_url": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("HARNESS_CLIENT_URL", "https://app.harness.io"),
			},
			"account_id": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("HARNESS_ACCOUNT_ID", ""),
			},
			"token": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("HARNESS_TOKEN", ""),
			},
			"kubernetes": {
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Description: "Kubernetes configuration.",
				Elem:        kubernetesResource(),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"delegate": resourceDelegateItem(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	m := &Meta{data: d}
	clientURL := d.Get("clientUrl").(string)
	accountID := d.Get("accountID").(string)
	token := d.Get("token").(string)
	kClient, restConfig := newKubeClient(d)
	if kClient == nil || restConfig == nil {
		return m, fmt.Errorf("k8s clients nil")
	}
	m.kubeClient = kClient
	m.restConfig = restConfig
	m.harnessClient = client.NewClient(clientURL, token, accountID)
	if m.harnessClient == nil {
		return m, fmt.Errorf("harness client nil")
	}
	return m, nil
}
