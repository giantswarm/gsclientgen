[![GoDoc](https://godoc.org/github.com/giantswarm/gsclientgen?status.svg)](https://godoc.org/github.com/giantswarm/gsclientgen)
[![Go Report Card](https://goreportcard.com/badge/github.com/giantswarm/gsclientgen)](https://goreportcard.com/report/github.com/giantswarm/gsclientgen)

# Giant Swarm Go Client (generated)

Experimental Go/Golang client for the Giant Swarm API, auto-generated based on the [OAI/Swagger specification](https://github.com/giantswarm/api-spec) using go-swagger.

## Usage

Here is a simplistic example of how to use the client for listing clusters:

```go
package main

import (
	"fmt"

	"github.com/giantswarm/gsclientgen/v2/client"
	"github.com/giantswarm/gsclientgen/v2/client/clusters"

	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

func main() {
	// You'll need the correct host name and probably "https" instead of "http"
	tp := httptransport.New("localhost:8000", "", []string{"http"})

	// You'll need a proper token
	token := "some-example-token"
	auth := httptransport.APIKeyAuth("Authorization", "header", "giantswarm " + token))

	params := clusters.NewGetClustersParams()
	c := client.New(tp, strfmt.Default)

	response, err := c.Clusters.GetClusters(params, auth)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("First cluster ID is '%s'\n", response.Payload[0].ID)
}
```

## Development

To pull the latest `api-spec` master and generate documentation and Go code:

```nohighlight
$ make generate
```

To work on a different branch of the `api-spec`, edit the `BRANCH` variable in the top of `Makefile`.

To double-check the spec's validity do this:

```nohighlight
$ make validate
```
