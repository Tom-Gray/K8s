APP = sentiment-analysis
REPO = swinkstom #Docker hub account
COMMIT_SHA = $(shell git rev-parse --short HEAD)
SERVICE ?= 	#[frontend, logic, webapp, webapp-go]

IMAGE_NAME = $(APP)-$(SERVICE)


build:
	docker build \
		--build-arg COMMIT_SHA=$(COMMIT_SHA)\
		-t $(IMAGE_NAME):$(COMMIT_SHA) $(IMAGE_NAME)/ 


push:
	docker tag $(IMAGE_NAME):$(COMMIT_SHA) swinkstom/$(IMAGE_NAME):$(COMMIT_SHA)
	docker push swinkstom/$(IMAGE_NAME):$(COMMIT_SHA)
	docker tag $(IMAGE_NAME):$(COMMIT_SHA) swinkstom/$(IMAGE_NAME):latest
	docker push swinkstom/$(IMAGE_NAME):latest
