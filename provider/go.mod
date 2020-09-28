module github.com/eahrend/terraform-harness-provider/provider

go 1.14

replace (
	github.com/googleapis/gnostic => github.com/googleapis/gnostic v0.3.1
	k8s.io/api => k8s.io/api v0.0.0-20190409021203-6e4e0e4f393b
	k8s.io/apimachinery => k8s.io/apimachinery v0.0.0-20190404173353-6a84e37a896d
	k8s.io/client-go => k8s.io/client-go v11.0.1-0.20190409021438-1a26190bd76a+incompatible
)

require (
	github.com/agext/levenshtein v1.2.3 // indirect
	github.com/bmatcuk/doublestar v1.3.2 // indirect
	github.com/eahrend/terraform-harness-provider/api/client v0.0.0-20200928170831-2d57eb52f322
	github.com/gogo/protobuf v1.3.1 // indirect
	github.com/google/go-cmp v0.5.2 // indirect
	github.com/google/gofuzz v1.2.0 // indirect
	github.com/google/uuid v1.1.2 // indirect
	github.com/googleapis/gnostic v0.5.1 // indirect
	github.com/hashicorp/go-multierror v1.1.0 // indirect
	github.com/hashicorp/go-retryablehttp v0.6.7 // indirect
	github.com/hashicorp/go-uuid v1.0.2 // indirect
	github.com/hashicorp/hil v0.0.0-20200423225030-a18a1cd20038 // indirect
	github.com/hashicorp/terraform v0.13.3
	github.com/hashicorp/terraform-plugin-sdk/v2 v2.0.3
	github.com/hashicorp/terraform-svchost v0.0.0-20200729002733-f050f53b9734 // indirect
	github.com/imdario/mergo v0.3.11 // indirect
	github.com/json-iterator/go v1.1.10 // indirect
	github.com/mitchellh/go-wordwrap v1.0.1 // indirect
	github.com/mitchellh/mapstructure v1.3.3 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.1 // indirect
	github.com/spf13/afero v1.4.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/stretchr/testify v1.6.1 // indirect
	github.com/zclconf/go-cty v1.6.1 // indirect
	golang.org/x/crypto v0.0.0-20200820211705-5c72a883971a // indirect
	golang.org/x/net v0.0.0-20200927032502-5d4f70055728 // indirect
	golang.org/x/oauth2 v0.0.0-20200902213428-5d25da1a8d43 // indirect
	golang.org/x/sys v0.0.0-20200926100807-9d91bd62050c // indirect
	golang.org/x/time v0.0.0-20200630173020-3af7569d3a1e // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
	k8s.io/api v0.19.2 // indirect
	k8s.io/apimachinery v0.19.2
	k8s.io/client-go v11.0.1-0.20190409021438-1a26190bd76a+incompatible
	k8s.io/klog v1.0.0 // indirect
	k8s.io/utils v0.0.0-20200912215256-4140de9c8800 // indirect
	sigs.k8s.io/yaml v1.2.0 // indirect
)
