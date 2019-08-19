PWD := $(shell pwd)
USERID=$(shell id -u)
GROUPID=$(shell id -g)

# API Spec branch name
BRANCH := openapi3

generate: clean
	# pull spec
	git clone --single-branch --branch ${BRANCH} --depth=1 https://github.com/giantswarm/api-spec/

	docker run --rm -it \
			-v ${PWD}:/local \
			-w /go/src/github.com/giantswarm/gsclientgen/api-spec \
			--user ${USERID}:${GROUPID} \
			openapitools/openapi-generator-cli generate \
				-i /local/api-spec/spec/spec.yaml \
				-o /local/src \
				-g go

# removal of all generated files
clean:
	rm -rf api-spec
	rm -rf src

# validate the code through building
build:
	dep ensure
	go build github.com/giantswarm/gsclientgen/models
	go build github.com/giantswarm/gsclientgen/client
