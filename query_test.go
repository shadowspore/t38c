package t38c

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInwQueryBuilder(t *testing.T) {
	client := &Client{}
	tests := []struct {
		Query    InwQueryBuilder
		Expected string
	}{
		{
			Query: client.Nearby("fleet", 10, 20, 30).
				Where("speed", 10, 20).
				Wherein("speed", 10, 20, 30).
				Match("abc*").
				Cursor(10).
				Format(FormatIDs).
				Limit(5),
			Expected: "NEARBY fleet WHERE speed 10 20 WHEREIN speed 3 10 20 30 MATCH abc* CURSOR 10 LIMIT 5 IDS POINT 10 20 30",
		},
		{
			Query: client.Intersects("fleet").
				Tile(10, 20, 30).
				Match("abc*"),
			Expected: "INTERSECTS fleet MATCH abc* TILE 10 20 30",
		},
	}

	for _, test := range tests {
		cmd := *test.Query.toCmd()
		actual := append([]string{cmd.Name}, cmd.Args...)
		expected := strings.Split(test.Expected, " ")
		assert.Equal(t, expected, actual)
	}
}

func TestGeofenceQueryBuilder(t *testing.T) {
	client := &Client{}
	tests := []struct {
		Query    GeofenceQueryBuilder
		Expected string
	}{
		{
			Query: client.GeofenceNearby("fleet", 10, 20, 30).
				Actions(Enter, Exit, Cross).
				Clip().
				Commands(Set, Del).
				Cursor(5).
				Format(FormatHashes(5)),
			Expected: "NEARBY fleet CLIP CURSOR 5 FENCE DETECT enter,exit,cross COMMANDS set,del HASHES 5 POINT 10 20 30",
		},
		{
			Query: client.GeofenceRoam("agent", "target", "*", 100).
				Distance().
				Wherein("price", 20, 30),
			Expected: "NEARBY agent DISTANCE WHEREIN price 2 20 30 FENCE ROAM target * 100",
		},
	}

	for _, test := range tests {
		cmd := *test.Query.toCmd()
		actual := append([]string{cmd.Name}, cmd.Args...)
		expected := strings.Split(test.Expected, " ")
		assert.Equal(t, expected, actual)
	}
}

func TestSetQueryBuilder(t *testing.T) {
	client := &Client{}
	tests := []struct {
		Query    SetQueryBuilder
		Expected string
	}{
		{
			Query: client.Set("agent", "47").
				PointZ(0, 0, -20).
				Field("age", 55).
				Expiration(60 * 60 * 24 * 365),
			Expected: "SET agent 47 EX 31536000 FIELD age 55 POINT 0 0 -20",
		},
	}

	for _, test := range tests {
		cmd := *test.Query.toCmd()
		actual := append([]string{cmd.Name}, cmd.Args...)
		expected := strings.Split(test.Expected, " ")
		assert.Equal(t, expected, actual)
	}
}
