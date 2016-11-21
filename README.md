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

Some usage examples:

```go
client := goclient.NewDefaultApi()
myToken = ""

// get an auth token (aka "Login")
requestBody := goclient.LoginBody{Password: base64EncodedPass}
loginResponse, _, err := client.UserLogin("email@example.com", requestBody)
if err != nil {
	log.Fatal(err)
}
if loginResponse.StatusCode == 10000 {
	myToken = loginResponse.Data.Id
	fmt.Printf("Successfully logged in. Token is %s.\n", loginResponse.Data.Id)
}

// list organizations the user is member of
authHeader := "giantswarm " + myToken
orgsResponse, _, err := client.GetUserOrganizations(authHeader)
if err != nil {
	log.Fatal(err)
}
if orgsResponse.StatusCode == 10000 {
	var organizations = orgsResponse.Data
	for _, orgName := orgsResponse.Data {
		fmt.Println(orgName)
	}
}

// log out
logoutResponse, _, err := client.UserLogout(authHeader)
if err != nil {
	log.Fatal(err)
}
if logoutResponse.StatusCode == 10007 {
	myToken = ""
	fmt.Println("Successfully logged out")
}
```

## Development

The source API specification can be found in `api-spec/oai-spec.yaml`. Changes here will recflect in changes of the generated code.

To generate client library code after changes in above file, run:

```nohighlight
make generate
```

