package provider

import (
	"bytes"
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"io"
	k8sMeta "k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer/yaml"
	yamlutil "k8s.io/apimachinery/pkg/util/yaml"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/restmapper"
	"log"
)

func resourceDelegateItem() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"delegate_name": {
				Type:         schema.TypeString,
				Required:     true,
				Description:  "The name of the resource, also acts as it's unique ID",
				ForceNew:     true,
				ValidateFunc: validateName,
			},
			"delegate_install": {
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
	log.Printf("[DEBUG] starting resourceCreateItem")
	meta := m.(Meta)
	log.Printf("[DEBUG] getting delegate name")
	delegateName := d.Get("delegate_name").(string)
	log.Printf("[DEBUG] delegate_name: %s", delegateName)
	log.Printf("[DEBUG] getting install type")
	installType := d.Get("install_type").(string)
	log.Printf("[DEBUG] install type: %s", installType)
	b, err := meta.harnessClient.GetNewDelegate(delegateName, installType)
	if err != nil {
		log.Printf("[ERROR] error getting new delegate: %s", err.Error())
		return err
	}
	log.Printf("[DEBUG] getting rest config from meta")
	rc := meta.restConfig
	log.Printf("[DEBUG] getting new dynamic config")
	dd, err := dynamic.NewForConfig(rc)
	if err != nil {
		log.Printf("[ERROR] failed to get new dynamic config: %s", err.Error())
		return err
	}
	log.Printf("[DEBUG] successfully created dynamic config")
	log.Printf("[DEBUG] creating yamlorjsondecorder")
	decoder := yamlutil.NewYAMLOrJSONDecoder(bytes.NewReader([]byte(b)), 1000)
	for {
		var rawObj runtime.RawExtension
		if err = decoder.Decode(&rawObj); err != nil {
			break
		}
		obj, gvk, err := yaml.NewDecodingSerializer(unstructured.UnstructuredJSONScheme).Decode(rawObj.Raw, nil, nil)
		unstructuredMap, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
		if err != nil {
			return err
		}
		unstructuredObj := &unstructured.Unstructured{Object: unstructuredMap}
		gr, err := restmapper.GetAPIGroupResources(meta.kubeClient.Discovery())
		if err != nil {
			return err
		}

		mapper := restmapper.NewDiscoveryRESTMapper(gr)
		mapping, err := mapper.RESTMapping(gvk.GroupKind(), gvk.Version)
		if err != nil {
			return err
		}
		var dri dynamic.ResourceInterface
		if mapping.Scope.Name() == k8sMeta.RESTScopeNameNamespace {
			if unstructuredObj.GetNamespace() == "" {
				unstructuredObj.SetNamespace("default")
			}
			dri = dd.Resource(mapping.Resource).Namespace(unstructuredObj.GetNamespace())
		} else {
			dri = dd.Resource(mapping.Resource)
		}
		if _, err := dri.Create(unstructuredObj, metav1.CreateOptions{}); err != nil {
			return err
		}
	}
	if err != io.EOF {
		log.Printf("[ERROR] failed to apply manifest: %s", err.Error())
		return err
	}
	return nil
}

// uhhh, not sure.
func resourceReadItem(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceUpdateItem(d *schema.ResourceData, m interface{}) error {
	return resourceCreateItem(d, m)
}

func resourceDeleteItem(d *schema.ResourceData, m interface{}) error {
	meta := m.(Meta)
	kc := meta.kubeClient
	return kc.CoreV1().Namespaces().Delete("harness-delegate", &metav1.DeleteOptions{})
}

func resourceExistsItem(d *schema.ResourceData, m interface{}) (bool, error) {
	// pretty much we can check for namespace/stateful set here.
	meta := m.(Meta)
	kc := meta.kubeClient
	_, err := kc.CoreV1().Namespaces().Get("harness-delegate", metav1.GetOptions{})
	if err != nil {
		// I know this can create false promises, I'll fix this later
		return false, nil
	}
	return true, nil
}

func validateName(v interface{}, k string) (ws []string, es []error) {
	var errs []error
	var warns []string
	/*
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
	*/
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
