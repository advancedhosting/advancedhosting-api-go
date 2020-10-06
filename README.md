# AH: A Go library for the Advanced Hosting API


The AH API documentation is available [here](https://api.websa.com/api-docs/).

## Install
```
go get github.com/advancedhosters/advancedhosting-api-go
```

## Usage

```go
package main

import (
	"context"
	"log"

	"github.com/advancedhosting/advancedhosting-api-go/ah"
)

func main() {
	clientOptions := &ah.ClientOptions{
		Token: "ACCESS_TOKEN",
	}

	client, err := ah.NewAPIClient(clientOptions)

	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	instances, _, err := client.Instances.List(ctx, nil)

	if err != nil {
		panic(err)
	}
	log.Printf("%v", instances)
}

```

