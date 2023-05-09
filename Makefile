APP = sentiment-analysis
REPO = swinkstom#Docker hub account
COMMIT_SHA = $(shell git rev-parse --short HEAD)
TAG ?=
SERVICE ?= 	#[frontend, logic, webapp, webapp-go]

IMAGE_NAME = $(APP)-$(SERVICE)
DEPLOYMENT_MANIFEST ?= 
SERVICE_MANIFEST ?= 

build:
	docker build \
		--build-arg COMMIT_SHA=$(COMMIT_SHA)\
		-t $(IMAGE_NAME):$(COMMIT_SHA) $(IMAGE_NAME)/ 


push:
	docker tag $(IMAGE_NAME):$(COMMIT_SHA) swinkstom/$(IMAGE_NAME):$(COMMIT_SHA)
	docker push swinkstom/$(IMAGE_NAME):$(COMMIT_SHA)
	docker tag $(IMAGE_NAME):$(COMMIT_SHA) swinkstom/$(IMAGE_NAME):latest
	docker push swinkstom/$(IMAGE_NAME):latest


patch_manifest: 
	yq '.spec.template.spec.containers[0].image = "$(REPO)/$(IMAGE_NAME):$(COMMIT_SHA)"' -i $(DEPLOYMENT_MANIFEST) 
	yq '.metadata.labels.version = "$(TAG)"' -i $(DEPLOYMENT_MANIFEST) 
	yq '.spec.selector.version = "$(TAG)"' -i $(SERVICE_MANIFEST) 



#	make patch_manifest \
	TAG=2.0.0 SERVICE=frontend \
	DEPLOYMENT_MANIFEST=resource-manifests/deployment-sa-frontend.yaml \
	SERVICE_MANIFEST=resource-manifests/service-sa-frontend-lb.yaml