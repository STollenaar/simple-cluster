provider "minikube" {
  kubernetes_version = "v1.26.1"
}

resource "minikube_cluster" "docker" {
  driver       = "docker"
  cluster_name = "simple-cluster"
}

output "minikube" {
  value     = minikube_cluster.docker
  sensitive = true
}
