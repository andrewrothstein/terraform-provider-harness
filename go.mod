module github.com/eahrend/terraform-harness-provider

go 1.14

replace (
	k8s.io/api => k8s.io/api v0.0.0-20190409021203-6e4e0e4f393b
	k8s.io/apimachinery => k8s.io/apimachinery v0.0.0-20190404173353-6a84e37a896d
	k8s.io/client-go => k8s.io/client-go v11.0.1-0.20190409021438-1a26190bd76a+incompatible
	github.com/googleapis/gnostic => github.com/googleapis/gnostic v0.3.1
)

require (
	github.com/eahrend/terraform-harness-provider/provider v0.0.0-20200925211759-2e9f8e8b4d97
	github.com/hashicorp/terraform v0.13.3
	k8s.io/client-go v11.0.1-0.20190409021438-1a26190bd76a+incompatible // indirect
)
