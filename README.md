# Tile38 Client
[![Go](https://github.com/xjem/t38c/workflows/Go/badge.svg)](https://github.com/xjem/t38c/actions)
[![Documentation](https://pkg.go.dev/badge/github.com/xjem/t38c)](https://pkg.go.dev/github.com/xjem/t38c?tab=doc)
[![Go Report Card](https://goreportcard.com/badge/github.com/xjem/t38c)](https://goreportcard.com/report/github.com/xjem/t38c)
[![codecov](https://codecov.io/gh/zweihander/tile38-client/branch/master/graph/badge.svg)](https://codecov.io/gh/axvq/tile38-client)
[![license](https://img.shields.io/github/license/axvq/tile38-client.svg)](https://github.com/xjem/t38c/blob/master/LICENSE)

Supported features: [click](TODO.md)

```
go get github.com/xjem/t38c
```

### Basic example

```go
package main

import (
	"fmt"

	"github.com/xjem/t38c"
)

func main() {
	client, err := t38c.New("localhost:9851", t38c.Debug)
	if err != nil {
		panic(err)
	}
	defer client.Close()

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
		Match("truck*").
		Format(t38c.FormatPoints).Do()
	if err != nil {
		panic(err)
	}

	// truck1 {33.5123 -112.2693}
	fmt.Println(response.Points[0].ID, response.Points[0].Point)
}
```
More examples: [click](examples)
