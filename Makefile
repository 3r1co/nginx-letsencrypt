IMAGE_NAME=r.3r1.co/letsencrypt-helper

all: compile build push deploy

compile:
	docker run --rm -it -v $(shell pwd)/letsencrypt:/go/src/myapp -w /go/src/myapp golang:1.7 sh -c '\
		go get ./... && \
		go build -o letsencrypt-helper \
	'
build:
	docker build -t $(IMAGE_NAME) -f letsencrypt/Dockerfile letsencrypt/

push:
	docker push $(IMAGE_NAME)
	
deploy:
	docker-compose down && docker-compose up
