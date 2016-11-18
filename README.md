[![Travis-CI Build Badge](https://api.travis-ci.org/giantswarm/go-client-gen.svg?branch=master)](https://travis-ci.org/giantswarm/go-client-gen)

# go-client-gen

Experimental Go client for the Giant Swarm API, auto-generated based on an OAI/Swagger specification using Swagger Codegen.

This client currently covers only a part of the API.

Documentation can be found in the sub folder `docs`.

## Usage

```nohighlight
go get github.com/go-resty/resty
go get github.com/giantswarm/go-client-gen
```

In your Go package, import like this:

```go
import goclient "github.com/giantswarm/go-client-gen"
```

## Development

The source API specification can be found in `api-spec/oai-spec.yaml`. Changes here will recflect in changes of the generated code.

To generate client library code after changes in above file, run:

```nohighlight
make generate
```

