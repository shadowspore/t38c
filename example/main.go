package main

import (
	"github.com/k0kubun/pp"
	t38c "github.com/lostpeer/tile38-client"
)

func main() {
	fencer, err := t38c.NewFence(t38c.NewRedisFencer("localhost:9851"))
	if err != nil {
		panic(err)
	}

	ch, err := fencer.FenceRoam("people", "people", "*", 10000,
		nil,
	)
	if err != nil {
		panic(err)
	}

	for event := range ch {
		pp.Println(event)
	}
}
