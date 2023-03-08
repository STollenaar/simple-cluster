# Simple-Cluster

Simple-Cluster is a simple deployment of a nginx server towards a kubernetes cluster.

## Prerequisite

Make sure you have the following things installed before preceding with the installation steps

1. Terraform
2. Docker
   - If deploying to an ARM based cluster, enable buildx and have a build agent that supports linux/arm

### Optional Prerequisite

If you are deploying to a remote cluster, make sure you have valid login credentials to you favorite container registry

## Installation

1. In the images directory run `make build`
   - In the event of pushing this build images to a remote container registry provide the registry url with `REGISTRY="REGISTRY_URL"`
   - (Only applicable for remote container regitries) If you're already logged in with your remote container registry you can also provide the push option with `OUT=--push`
   - You can also modify the to be build image tag with `TAG=x.x.x`




## Steps Taken

1. Downloaded the assesment archive (https://tldp.org/LDP/abs/abs-guide.html.tar.gz)
2. Created file structed and started on Dockerfile
