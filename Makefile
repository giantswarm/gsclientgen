PWD := $(shell pwd)
USERID=$(shell id -u)
GROUPID=$(shell id -g)

# API Spec branch name
BRANCH := master

generate: clean
	# pull spec
	mkdir -p ./api-spec
	curl -s https://raw.githubusercontent.com/giantswarm/api-spec/$(BRANCH)/spec.yaml > api-spec/spec.yaml
	curl -s https://raw.githubusercontent.com/giantswarm/api-spec/$(BRANCH)/responses.yaml > api-spec/responses.yaml
	curl -s https://raw.githubusercontent.com/giantswarm/api-spec/$(BRANCH)/parameters.yaml > api-spec/parameters.yaml
	curl -s https://raw.githubusercontent.com/giantswarm/api-spec/$(BRANCH)/definitions.yaml > api-spec/definitions.yaml
	docker run --rm -it \
	  -v ${PWD}:/go/src/github.com/giantswarm/gsclientgen \
		-w /go/src/github.com/giantswarm/gsclientgen/api-spec \
		--user ${USERID}:${GROUPID} \
		quay.io/goswagger/swagger:0.15.0 \
			generate client \
			--spec spec.yaml \
			--name gsclientgen \
			--default-scheme https \
			--target .. \
			--with-flatten=full
	gofmt -s -l -w client
	gofmt -s -l -w models

# removal of all generated files
clean:
	rm -rf client
	rm -rf models
	rm -rf api-spec/*

# validate the spec
validate:
	docker run --rm -it \
	    -v ${PWD}/api-spec:/workdir \
		--user ${USERID}:${GROUPID} \
	    boiyaa/yamllint:1.8.1 ./spec.yaml
	docker run --rm -it \
		-v ${PWD}/api-spec:/swagger-api/yaml \
		--user ${USERID}:${GROUPID} \
		jimschubert/swagger-codegen-cli:2.2.3 generate \
		--input-spec /swagger-api/yaml/spec.yaml \
		--lang swagger \
		--output /tmp/

# validate the code through building
build:
	dep ensure
	go build github.com/giantswarm/gsclientgen/models
	go build github.com/giantswarm/gsclientgen/client
