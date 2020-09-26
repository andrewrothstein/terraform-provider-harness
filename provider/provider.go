package provider

import (
	"fmt"
	"github.com/eahrend/terraform-harness-provider/api/client"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"log"
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
	clientURL := d.Get("client_url").(string)
	log.Printf("Client URL: %s", clientURL)
	accountID := d.Get("account_id").(string)
	log.Printf("Account ID: %s", accountID)
	token := d.Get("token").(string)
	log.Printf("token: %s", token)
	kClient, restConfig, err := newKubeClient(d)
	if err != nil {
		log.Printf("provider configure error %v", err)
		return nil, err
	}
	m.kubeClient = kClient
	m.restConfig = restConfig
	m.harnessClient = client.NewClient(clientURL, token, accountID)
	if m.harnessClient == nil {
		return m, fmt.Errorf("harness client nil")
	}
	log.Println("[INFO] Harness Provider Configured successfully!")
	return m, nil
}
