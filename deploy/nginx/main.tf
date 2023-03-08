terraform {
  backend "local" {

  }
}

data "terraform_remote_state" "minikube_cluster" {
  backend = "local"
  config = {
    path = "../minikube/terraform.tfstate"
  }
}

locals {
  minikube = data.terraform_remote_state.minikube_cluster.outputs.minikube
}

provider "kubernetes" {
  host                   = local.minikube.host
  cluster_ca_certificate = local.minikube.cluster_ca_certificate
  client_certificate     = local.minikube.client_certificate
  client_key             = local.minikube.client_key
}
