module github.com/eahrend/terraform-harness-provider/provider

go 1.13

require (
	github.com/eahrend/terraform-harness-provider/api/client v0.0.0-20200923202512-401023d8a8dd
	github.com/hashicorp/terraform v0.13.3
	github.com/mitchellh/go-homedir v1.1.0
	golang.org/x/net v0.0.0-20200923182212-328152dc79b1 // indirect
	golang.org/x/sys v0.0.0-20200923182605-d9f96fdee20d // indirect
	k8s.io/apimachinery v0.0.0-20190913080033-27d36303b655
	k8s.io/client-go v11.0.0+incompatible
)
