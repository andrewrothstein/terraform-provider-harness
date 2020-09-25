module github.com/eahrend/terraform-harness-provider/provider

go 1.14

replace (
	github.com/googleapis/gnostic => github.com/googleapis/gnostic v0.3.1
	golang.org/x/sys => golang.org/x/sys v0.0.0-20190813064441-fde4db37ae7a // pinned to release-branch.go1.13
	golang.org/x/tools => golang.org/x/tools v0.0.0-20190821162956-65e3620a7ae7 // pinned to release-branch.go1.13
	k8s.io/api => k8s.io/api v0.0.0-20200214081623-ecbd4af0fc33
	k8s.io/apimachinery => k8s.io/apimachinery v0.0.0-20200214081019-7490b3ed6e92
	k8s.io/client-go => k8s.io/client-go v0.0.0-20200214082307-e38a84523341
	k8s.io/code-generator => k8s.io/code-generator v0.0.0-20200214080538-dc8f3adce97c
)

require (
	github.com/agext/levenshtein v1.2.3 // indirect
	github.com/bmatcuk/doublestar v1.3.2 // indirect
	github.com/eahrend/terraform-harness-provider/api/client v0.0.0-20200925204633-a8e92db65e5f
	github.com/google/go-cmp v0.5.2 // indirect
	github.com/google/gofuzz v1.2.0 // indirect
	github.com/google/uuid v1.1.2 // indirect
	github.com/googleapis/gnostic v0.5.1 // indirect
	github.com/hashicorp/go-multierror v1.1.0 // indirect
	github.com/hashicorp/go-retryablehttp v0.6.7 // indirect
	github.com/hashicorp/go-uuid v1.0.2 // indirect
	github.com/hashicorp/go-version v1.2.1 // indirect
	github.com/hashicorp/hil v0.0.0-20200423225030-a18a1cd20038 // indirect
	github.com/hashicorp/terraform v0.13.3
	github.com/hashicorp/terraform-svchost v0.0.0-20200729002733-f050f53b9734 // indirect
	github.com/imdario/mergo v0.3.11 // indirect
	github.com/json-iterator/go v1.1.10 // indirect
	github.com/mitchellh/go-wordwrap v1.0.1 // indirect
	github.com/mitchellh/mapstructure v1.3.3 // indirect
	github.com/spf13/afero v1.4.0 // indirect
	github.com/vmihailenco/msgpack v4.0.4+incompatible // indirect
	github.com/zclconf/go-cty v1.6.1 // indirect
	golang.org/x/crypto v0.0.0-20200820211705-5c72a883971a // indirect
	golang.org/x/net v0.0.0-20200925080053-05aa5d4ee321 // indirect
	golang.org/x/oauth2 v0.0.0-20200902213428-5d25da1a8d43 // indirect
	golang.org/x/sys v0.0.0-20200923182605-d9f96fdee20d // indirect
	golang.org/x/time v0.0.0-20200630173020-3af7569d3a1e // indirect
	k8s.io/api v0.19.2 // indirect
	k8s.io/apimachinery v0.19.2
	k8s.io/client-go v11.0.0+incompatible
	k8s.io/utils v0.0.0-20200912215256-4140de9c8800 // indirect
	sigs.k8s.io/structured-merge-diff/v3 v3.0.0 // indirect
)
