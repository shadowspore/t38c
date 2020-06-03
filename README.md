[![Go Report Card](https://goreportcard.com/badge/github.com/b3q/tile38-client)](https://goreportcard.com/report/github.com/b3q/tile38-client)
[![Documentation](https://godoc.org/github.com/b3q/tile38-client?status.svg)](https://pkg.go.dev/github.com/b3q/tile38-client?tab=doc)
# Tile38 Client

Supported features: [click](TODO.md)

```
go get github.com/b3q/tile38-client
```

### Basic example

```go
package main

import t38c "github.com/b3q/tile38-client"

func main() {
	client, err := t38c.New("localhost:9851", t38c.Debug)
	if err != nil {
		panic(err)
	}

	if err := client.Keys.Set("fleet", "truck1").Point(33.5123, -112.2693).Do(); err != nil {
		panic(err)
	}

	if err := client.Keys.Set("fleet", "truck2").Point(33.4626, -112.1695).
		// optional params
		Field("speed", 20).
		Expiration(20).
		Do(); err != nil {
		panic(err)
	}

	// search 6 kilometers around a point. returns one truck.
	response, err := client.Search.Nearby("fleet", 33.462, -112.268, 6000).
		Where("speed", 0, 100).
		Match("truck*").Do()
	if err != nil {
		panic(err)
	}
}
```
More examples: [click](examples)