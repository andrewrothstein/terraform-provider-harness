module github.com/eahrend/terraform-harness-provider

go 1.13

require (
	github.com/eahrend/terraform-harness-provider/api/client v0.0.0-20200924221147-8781e686c52e // indirect
	github.com/eahrend/terraform-harness-provider/provider v0.0.0-20200924221147-8781e686c52e
	github.com/fatih/color v1.9.0 // indirect
	github.com/go-logr/logr v0.2.1 // indirect
	github.com/gogo/protobuf v1.3.1 // indirect
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
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.1 // indirect
	github.com/oklog/run v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	golang.org/x/net v0.0.0-20200925080053-05aa5d4ee321 // indirect
	golang.org/x/time v0.0.0-20200630173020-3af7569d3a1e // indirect
	google.golang.org/genproto v0.0.0-20200925023002-c2d885f95484 // indirect
	google.golang.org/grpc v1.32.0 // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
	k8s.io/api v0.19.2
	k8s.io/apimachinery v0.19.2
	k8s.io/client-go v11.0.1-0.20190409021438-1a26190bd76a+incompatible
	k8s.io/klog v1.0.0 // indirect
	k8s.io/utils v0.0.0-20200912215256-4140de9c8800 // indirect
	sigs.k8s.io/yaml v1.2.0 // indirect
)

replace (
	k8s.io/api => k8s.io/api v0.0.0-20190409021203-6e4e0e4f393b
	k8s.io/apimachinery => k8s.io/apimachinery v0.0.0-20190404173353-6a84e37a896d
	k8s.io/client-go => k8s.io/client-go v11.0.1-0.20190409021438-1a26190bd76a+incompatible
)
