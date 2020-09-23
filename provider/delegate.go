package provider

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"regexp"
)

func resourceDelegateItem() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"delegateName": {
				Type:         schema.TypeString,
				Required:     true,
				Description:  "The name of the resource, also acts as it's unique ID",
				ForceNew:     true,
				ValidateFunc: validateName,
			},
			"delegateInstall": {
				Type:         schema.TypeString,
				Required:     true,
				Description:  "A description of an item",
				ValidateFunc: validateInstall,
			},
		},
		Create: resourceCreateItem,
		Read:   resourceReadItem,
		Update: resourceUpdateItem,
		Delete: resourceDeleteItem,
		Exists: resourceExistsItem,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
	}
}

func resourceCreateItem(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceReadItem(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceUpdateItem(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceDeleteItem(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceExistsItem(d *schema.ResourceData, m interface{}) (bool, error) {
	return false, nil
}

func validateName(v interface{}, k string) (ws []string, es []error) {
	var errs []error
	var warns []string
	value, ok := v.(string)
	if !ok {
		errs = append(errs, fmt.Errorf("expected name to be string"))
		return warns, errs
	}
	whiteSpace := regexp.MustCompile(`\s+`)
	if whiteSpace.Match([]byte(value)) {
		errs = append(errs, fmt.Errorf("name cannot contain whitespace. Got %s", value))
		return warns, errs
	}
	return warns, errs
}

func validateInstall(v interface{}, k string) (ws []string, es []error) {
	var errs []error
	var warns []string
	value, ok := v.(string)
	if !ok {
		errs = append(errs, fmt.Errorf("install type should be a string"))
		return warns, errs
	}
	// starting with just a specific install type that works for me
	switch installType := value; installType {
	case KUBERNETES_YAML:
		fmt.Println("KUBERNETES_YAML install")
	/*
		case SHELL_SCRIPT:
			fmt.Println("SHELL_SCRIPT install")
		case DOCKER_IMAGE:
			fmt.Println("DOCKER_IMAGE install")
		case HELM_VALUES_YAML:
			fmt.Println("HELM_VALUES_YAML install")
		case ECS_TASK_SPEC:
			fmt.Println("ECS_TASK_SPEC install")
	*/
	default:
		errs = append(errs, fmt.Errorf("unsupported install type %s,expected delegate install to be one of the following: SHELL_SCRIPT\nDOCKER_IMAGE\nKUBERNETES_YAML\nHELM_VALUES_YAML\nECS_TASK_SPEC", installType))
	}
	return warns, errs
}
