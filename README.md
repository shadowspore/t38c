[![Go Report Card](https://goreportcard.com/badge/github.com/lostpeer/tile38-client)](https://goreportcard.com/report/github.com/lostpeer/tile38-client)

# tile38-client - Tile38 client for Go

Most features are supported (see [TODO.md](TODO.md))

## Basic example

more examples in [examples](examples) folder

```go
package main

import t38c "github.com/lostpeer/tile38-client"

func main() {
	client, err := t38c.New("localhost:9851", t38c.Debug())
	if err != nil {
		panic(err)
	}

	client.Set("fleet", "truck1", t38c.SetPoint(33.5123, -112.2693))
	client.Set("fleet", "truck2", t38c.SetPoint(33.4626, -112.1695),
		// optional params
		t38c.Field("speed", 20),
		t38c.Expiration(20),
	)

	client.Search(
		// search 6 kilometers around a point. returns one truck.
		t38c.Nearby("fleet", 33.462, -112.268, 6000,
			// optional params
			t38c.Where("speed", 0, 100),
			t38c.Match("truck*"),
		),
	)
}
```