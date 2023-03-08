resource "kubernetes_deployment" "nginx" {
  metadata {
    name      = "nginx"
    namespace = "default"
  }

  spec {
    replicas = 1
    selector {
      match_labels = {
        app = "nginx"
      }
    }

    template {
      metadata {
        labels = {
          app = "nginx"
        }
      }
      spec {
        container {
          name              = "nginx"
          image             = "local/nginx:0.0.0" # If deployed to a remote cluster, this needs to be changed to the appropriate image tag
          image_pull_policy = "Never"             # If Deployed to a remote cluster change this to "IfNotPresent"
          port {
            container_port = 80
          }
        }
      }
    }
  }
}

resource "kubernetes_service" "nginx" {
  metadata {
    name      = "nginx"
    namespace = "default"
  }

  spec {
    type = "NodePort"
    selector = {
      app = "nginx"
    }
    port {
      port        = 80
      name        = "nginx"
      target_port = 80
    }
  }
}
