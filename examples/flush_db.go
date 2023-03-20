//go:build ignore
// +build ignore

package main

import (
	"context"
	"log"

	"github.com/sythang/t38c"
)

/*
*	FLUSHDB Example
*
*	Shows how to erase all data in Tile38
*	database using the FLUSHDB command.
*
 */

func main() {
	// Create a Tile38 client.
	tile38, err := t38c.New(t38c.Config{
		Address: "localhost:9851",
		Debug:   true,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer tile38.Close()

	// Add a point named 'truck1' to a collection named 'first fleet'.
	if err = tile38.Keys.Set("first fleet", "truck1").Point(33.5123, -112.2693).Do(context.TODO()); err != nil {
		log.Fatal(err)
	}

	// Add a point named 'truck2' to a collection named 'second fleet'.
	if err = tile38.Keys.Set("second fleet", "truck2").Point(23.6951, -92.3581).Do(context.TODO()); err != nil {
		log.Fatal(err)
	}

	// Get all keys.
	// Returns ["first fleet","second fleet"].
	_, err = tile38.Keys.Keys(context.TODO(), "*")
	if err != nil {
		log.Fatal(err)
	}

	// Flush ALL data in Tile38 database.
	if err = tile38.Server.FlushDB(context.TODO()); err != nil {
		log.Fatal(err)
	}

	// Get all keys again.
	// Returns [].
	_, err = tile38.Keys.Keys(context.TODO(), "*")
	if err != nil {
		log.Fatal(err)
	}
}
