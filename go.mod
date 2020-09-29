module github.com/eahrend/terraform-harness-provider

go 1.14

replace (
	github.com/googleapis/gnostic => github.com/googleapis/gnostic v0.3.1
	k8s.io/api => k8s.io/api v0.0.0-20190409021203-6e4e0e4f393b
	k8s.io/apimachinery => k8s.io/apimachinery v0.0.0-20190404173353-6a84e37a896d
	k8s.io/client-go => k8s.io/client-go v11.0.1-0.20190409021438-1a26190bd76a+incompatible
)

require (
	github.com/eahrend/terraform-harness-provider/provider v0.0.0-20200929011740-f259bac8e7fe
	github.com/fatih/color v1.9.0 // indirect
	github.com/hashicorp/go-hclog v0.14.1 // indirect
	github.com/hashicorp/terraform v0.13.3
	github.com/hashicorp/yamux v0.0.0-20200609203250-aecfd211c9ce // indirect
	github.com/mattn/go-colorable v0.1.7 // indirect
	github.com/mitchellh/go-testing-interface v1.14.1 // indirect
	github.com/oklog/run v1.1.0 // indirect
	google.golang.org/genproto v0.0.0-20200925023002-c2d885f95484 // indirect
	google.golang.org/grpc v1.32.0 // indirect
)
