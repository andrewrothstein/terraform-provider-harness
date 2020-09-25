module github.com/eahrend/terraform-harness-provider

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
	github.com/eahrend/terraform-harness-provider/api/client v0.0.0-20200925204633-a8e92db65e5f // indirect
	github.com/eahrend/terraform-harness-provider/provider v0.0.0-20200925204633-a8e92db65e5f
	github.com/fatih/color v1.9.0 // indirect
	github.com/go-logr/logr v0.2.1 // indirect
	github.com/google/gofuzz v1.2.0 // indirect
	github.com/googleapis/gnostic v0.5.1 // indirect
	github.com/hashicorp/go-hclog v0.14.1 // indirect
	github.com/hashicorp/terraform v0.13.3
	github.com/hashicorp/yamux v0.0.0-20200609203250-aecfd211c9ce // indirect
	github.com/imdario/mergo v0.3.11 // indirect
	github.com/json-iterator/go v1.1.10 // indirect
	github.com/mattn/go-colorable v0.1.7 // indirect
	github.com/mitchellh/go-testing-interface v1.14.1 // indirect
	github.com/mitchellh/go-wordwrap v1.0.1 // indirect
	github.com/oklog/run v1.1.0 // indirect
	golang.org/x/net v0.0.0-20200925080053-05aa5d4ee321 // indirect
	golang.org/x/sys v0.0.0-20200923182605-d9f96fdee20d // indirect
	golang.org/x/time v0.0.0-20200630173020-3af7569d3a1e // indirect
	google.golang.org/genproto v0.0.0-20200925023002-c2d885f95484 // indirect
	google.golang.org/grpc v1.32.0 // indirect
	k8s.io/api v0.19.2 // indirect
	k8s.io/klog/v2 v2.3.0 // indirect
	k8s.io/utils v0.0.0-20200912215256-4140de9c8800 // indirect
	sigs.k8s.io/structured-merge-diff/v3 v3.0.0 // indirect
)
