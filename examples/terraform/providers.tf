provider "harness" {
  source = "github.com/eahrend/terraform-harness-provider"
  clientUrl   = "https://app.harness.io"
  accountID   = "2KTQt0X9R82AEBbv9RYn_g"
  bearerToken = "super secret token yeah buddy"
  kubernetes {
    load_config_file       = false
    host                   = ""
    token                  = ""
    cluster_ca_certificate = ""
  }
}
