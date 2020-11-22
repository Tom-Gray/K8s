APP = sentiment-analysis
REPO = docker.pkg.github.com/tom-gray/$(APP)
COMMIT_SHA = $(shell git rev-parse --short HEAD)
SERVICE ?= 	#[fronted, logic, webapp]

IMAGE_NAME = $(APP)-$(SERVICE)



build:
	docker build \
		--build-arg COMMIT_SHA=$(COMMIT_SHA)\
		-t $(IMAGE_NAME):$(COMMIT_SHA) $(IMAGE_NAME)/ 


push:
	docker tag $(IMAGE_NAME):$(COMMIT_SHA) $(REPO)/$(IMAGE_NAME):$(COMMIT_SHA)
	docker push $(REPO)/$(IMAGE_NAME):$(COMMIT_SHA)
	docker tag $(IMAGE_NAME):$(COMMIT_SHA) $(REPO)/$(IMAGE_NAME):latest
	docker push $(REPO)/$(IMAGE_NAME):latest


compile-webapp:
	mvn