package provider

import (
	"github.com/eahrend/terraform-harness-provider/api/client"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"clientUrl": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("HARNESS_CLIENT_URL", "https://app.harness.io"),
			},
			"accountID": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("HARNESS_ACCOUNT_ID", ""),
			},
			"token": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("HARNESS_TOKEN", ""),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"delegate": resourceDelegateItem(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	clientURL := d.Get("clientUrl").(string)
	accountID := d.Get("accountID").(string)
	token := d.Get("token").(string)
	return client.NewClient(clientURL, token, accountID), nil
}
