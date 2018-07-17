[![GoDoc](https://godoc.org/github.com/giantswarm/gsclientgen?status.svg)](https://godoc.org/github.com/giantswarm/gsclientgen)
[![Go Report Card](https://goreportcard.com/badge/github.com/giantswarm/gsclientgen)](https://goreportcard.com/report/github.com/giantswarm/gsclientgen)
[![IRC Channel](https://img.shields.io/badge/irc-%23giantswarm-blue.svg)](https://kiwiirc.com/client/irc.freenode.net/#giantswarm)

# Giant Swarm Golang Client (generated)

Experimental Go client for the Giant Swarm API, auto-generated based on an OAI/Swagger specification using Swagger Codegen.

Note: This client currently covers only a part of the API. Expect lots of breaking changes within the code. Use at your own risk.

Documentation can be found in the sub folder `docs`.

## Usage

Here is a simplistic example of how to use the client for listing clusters:

```go
package main

import (
	"fmt"

	"github.com/giantswarm/gsclientgen/client"
	"github.com/giantswarm/gsclientgen/client/clusters"

	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

func main() {
	// You'll need the correct host name and probably "https" instead of "http"
	tp := httptransport.New("localhost:8000", "", []string{"http"})

	// You'll need a proper token
	token := "some-example-token"

	params := clusters.NewGetClustersParams().WithAuthorization("giantswarm " + token)
	c := client.New(tp, strfmt.Default)

	response, err := c.Clusters.GetClusters(params, nil)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("First cluster ID is '%s'\n", response.Payload[0].ID)
}
```

## Development

The source API specification can be found in `api-spec/oai-spec.yaml`. Changes here will recflect in changes of the generated code.

To generate client library code after changes in above file, run:

```nohighlight
make validate
make generate
```
