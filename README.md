# Simple-Cluster

Simple-Cluster is a simple deployment of a nginx server towards a kubernetes cluster.

## Prerequisite

Make sure you have the following things installed before preceding with the installation steps

1. Terraform
2. Docker
   - If deploying to an ARM based cluster, enable buildx and have a build agent that supports linux/arm

### Optional Prerequisite

If you are deploying to a remote cluster, make sure you have valid login credentials to you favorite container registry
If you are deploying to a local cluster (eg. minikube) make sure you have the minikube cli installed

## Installation

1. In the `images` directory run `make build`
   - In the event of pushing this build images to a remote container registry provide the registry url with `REGISTRY="REGISTRY_URL"`
   - (Only applicable for remote container regitries) If you're already logged in with your remote container registry you can also provide the push option with `OUT=--push`
   - You can also modify the to be build image tag with `TAG=x.x.x`

2. In the `deploy` directory first navigate into the `minikube` directory to deploy a minikube cluster locally if deploying on a local cluster
   1. Run `terraform init`
   2. Run `terraform apply`
   3. Once minikube is up and running we can add the local nginx image to minikube if we didn't push it to a container registry.
      1. Navigate back into the `images` directory and run `make add`

The Following step assumes this is all deployed into minikube from the previoous step. Changes to the `main.tf` and `nginx.tf` files are
to get this to work with a remote cluster (eg. AWS EKS)

1. Back in the `deploy` directory navigate into the `nginx` directory to deploy the build nginx image

   1. Run `terraform init`
   2. Run `terraform apply`

2. Retrieve the nginx website url with `minikube service nginx --url` in a terminal window

## Steps Taken

1. Downloaded the assesment archive (https://tldp.org/LDP/abs/abs-guide.html.tar.gz)
2. Created file structed and started on Dockerfile
3. Created terraform config for the minikube cluster and nginx deployment
