# Simple Makefile to create our wanted nginx image

.PHONY: build


TAG ?=0.0.0
REGISTRY?=local
IMAGE = $(REGISTRY)/nginx
OUT=--load


# Going to build with buildx. There is a way to build into multiple platforms at the same time
# I could not do this at the time.
# PLATFORMS ?=linux/amd64,linux/arm64
PLATFORMS ?=linux/amd64

# required to enable buildx
export DOCKER_CLI_EXPERIMENTAL=enabled

build: 
	docker buildx build --platform=${PLATFORMS} $(OUT) --pull --tag $(IMAGE):$(TAG) nginx

add:
	minikube image load $(IMAGE):$(TAG)