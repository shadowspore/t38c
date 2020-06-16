package t38c

import (
	"fmt"
	"strings"
	"testing"
)

func TestSearch(t *testing.T) {
	search := &Search{}
	tests := []struct {
		Cmd      *tileCmd
		Expected string
	}{
		{
			Cmd: search.Nearby("fleet", 10, 20, 30).
				Where("speed", 10, 20).
				Wherein("speed", 10, 20, 30).
				Match("abc*").
				Cursor(10).
				Format(FormatIDs).
				Limit(5).toCmd(),
			Expected: "NEARBY fleet WHERE speed 10 20 WHEREIN speed 3 10 20 30 MATCH abc* CURSOR 10 LIMIT 5 IDS POINT 10 20 30",
		},
		{
			Cmd: search.Intersects("fleet").
				Tile(10, 20, 30).
				Match("abc*").toCmd(),
			Expected: "INTERSECTS fleet MATCH abc* TILE 10 20 30",
		},
	}

	for _, test := range tests {
		err := compare(test.Cmd, test.Expected)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func TestGeofence(t *testing.T) {
	geofence := &Geofence{}
	tests := []struct {
		Cmd      *tileCmd
		Expected string
	}{
		{
			Cmd: geofence.Nearby("fleet", 10, 20, 30).
				Actions(Enter, Exit, Cross).
				Clip().
				Commands(Set, Del).
				Cursor(5).
				Format(FormatHashes(5)).toCmd(),
			Expected: "NEARBY fleet CLIP CURSOR 5 FENCE DETECT enter,exit,cross COMMANDS set,del HASHES 5 POINT 10 20 30",
		},
		{
			Cmd: geofence.Roam("agent", "target", "*", 100).
				Distance().
				Wherein("price", 20, 30).toCmd(),
			Expected: "NEARBY agent DISTANCE WHEREIN price 2 20 30 FENCE ROAM target * 100",
		},
	}

	for _, test := range tests {
		err := compare(test.Cmd, test.Expected)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func TestKeys(t *testing.T) {
	keys := &Keys{}
	tests := []struct {
		Cmd      *tileCmd
		Expected string
	}{
		{
			Cmd: keys.Set("agent", "47").
				PointZ(0, 0, -20).
				Field("age", 55).
				Expiration(60 * 60 * 24 * 365).toCmd(),
			Expected: "SET agent 47 EX 31536000 FIELD age 55 POINT 0 0 -20",
		},
		{
			Cmd: keys.FSet("agent", "47").
				Field("cash", 100500).
				IfExists().toCmd(),
			Expected: "FSET agent 47 XX cash 100500",
		},
	}

	for _, test := range tests {
		err := compare(test.Cmd, test.Expected)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func compare(cmd *tileCmd, expectedCmd string) error {
	actual := append([]string{cmd.Name}, cmd.Args...)
	expected := strings.Split(expectedCmd, " ")
	if len(actual) != len(expected) {
		return fmt.Errorf("not equal: bad length (%d, %d)\n"+
			"expected: %s\n"+
			"actual  : %s\n", len(expected), len(actual), expected, actual)
	}

	for i := 0; i < len(actual); i++ {
		if actual[i] != expected[i] {
			return fmt.Errorf("not equal:\n"+
				"expected: %s\n"+
				"actual  : %s\n", expected, actual)
		}
	}
	return nil
}
