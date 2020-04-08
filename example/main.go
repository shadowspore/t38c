package main

import (
	"github.com/k0kubun/pp"
	t38c "github.com/zerobounty/tile38-client"
)

func main() {
	fencer, err := t38c.NewFence(t38c.NewRedisFencer("localhost:9851"), true)
	if err != nil {
		panic(err)
	}

	ch, err := fencer.FenceRoam(
		t38c.NewFenceRequest("people", "people", "*", 10000).WithOptions(t38c.Match("kekka*")),
	)
	if err != nil {
		panic(err)
	}

	for event := range ch {
		pp.Println(event)
	}
}
