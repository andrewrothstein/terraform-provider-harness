resource "delegate" "new_delegate" {
  delegateName    = "us-east1-dev"
  delegateInstall = "KUBERNETES_YAML"
}

/*******************************
Supported Delegate Install Types
SHELL_SCRIPT
DOCKER_IMAGE
KUBERNETES_YAML
HELM_VALUES_YAML
ECS_TASK_SPEC
*******************************/
