# Simple Makefile to create our wanted nginx image

.PHONY: build


TAG ?=0.0.0
REGISTRY?=local
IMAGE = $(REGISTRY)/nginx
OUT=--load


# Going to build with buildx
PLATFORMS ?=linux/amd64,linux/arm

# required to enable buildx
export DOCKER_CLI_EXPERIMENTAL=enabled

build: 
	docker buildx build --platform=${PLATFORMS} $(OUT) --pull --tag $(IMAGE):$(TAG) nginx

add:
	minikube image load $(IMAGE):$(TAG)