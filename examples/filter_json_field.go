package main

import (
	"context"
	geojson "github.com/paulmach/go.geojson"
	"github.com/xjem/t38c"
	"log"
)

func main() {
	tile38, err := t38c.New(t38c.Config{
		Address: "localhost:9851",
		Debug:   true,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer tile38.Close()
	feature := geojson.NewPointFeature([]float64{33.5123, -112.2693})
	feature.Properties = map[string]interface{}{
		"speed": 55,
		"name":  "Carol",
		"age":   "23",
	}
	f2 := geojson.NewPointFeature([]float64{33.553653, -112.112222})
	f2.Properties = map[string]interface{}{
		"speed": 40,
		"name":  "Andy",
		"age":   "25",
	}
	tile38.Keys.Set("fleet", "carol").Feature(feature).Do(context.Background())
	tile38.Keys.Set("fleet", "andy").Feature(f2).Do(context.Background())

	// references: https://tile38.com/topics/filter-expressions
	tile38.Search.Scan("fleet").RawQuery("properties.age == 25 && properties.speed > 50").Do(context.Background())
}
