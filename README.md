# Simple-Cluster

Simple-Cluster is a simple deployment of a nginx server towards a kubernetes cluster. This kubernetes cluster is assumed to be a local minikube deployment

## Prerequisite

Make sure you have the following things installed before proceding with the installation steps

1. Terraform
2. Docker
   - Make sure buildkit is enabled since the build step for our nginx image will be using buildx
3. Minikube

### Optional Prerequisite

If you are deploying to a local cluster (eg. minikube) make sure you have the minikube cli installed

## Installation

1. In the `images` directory run `make build`

2. In the `deploy` directory first navigate into the `minikube` directory to deploy a minikube cluster locally
   1. Run `terraform init`
   2. Run `terraform apply`
   3. Once minikube is up and running we can add the local nginx image to minikube
      1. Navigate back into the `images` directory and run `make add`

3. Back in the `deploy` directory navigate into the `nginx` directory to deploy the build nginx image

   1. Run `terraform init`
   2. Run `terraform apply`

4. Retrieve the nginx website url with `minikube service nginx --url` in a terminal window

## Steps Taken

1. Downloaded the assesment archive (https://tldp.org/LDP/abs/abs-guide.html.tar.gz)
2. Created file structed and started on Dockerfile
3. Created terraform config for the minikube cluster and nginx deployment
4. Created simple monitor script in bash
5. Created simple monitor script in go
6. Some cleanup

## Visibility

The Nginx service is made visible with the help of the Nodeport in our Minikube cluster. A Nodeport makes a particular port on a node of our Cluster visible to traffic outside the cluster (If this was deployed in a public facing network any internet traffic could stumble on it). The `minikube service nginx --url` command helps us fetch the actual local url such that we can view our website.

Deploying on a different cluster the best practice approach would be by having an ingress controller help navigate traffic to our Nginx instance.