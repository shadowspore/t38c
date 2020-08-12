package t38c

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChannels(t *testing.T) {
	channels := &Channels{}
	geofence := &Geofence{}
	tests := []struct {
		Cmd      cmd
		Expected string
	}{
		{
			Cmd:      channels.SetChan("foo", geofence.Roam("cat", "dog", "*", 50).Actions(Enter, Exit)).toCmd(),
			Expected: "SETCHAN foo NEARBY cat FENCE DETECT enter,exit ROAM dog * 50",
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.Expected, test.Cmd.String())
	}
}

func TestHooks(t *testing.T) {
	hooks := &Hooks{}
	geofence := &Geofence{}
	tests := []struct {
		Cmd      cmd
		Expected string
	}{
		{
			Cmd:      hooks.SetHook("foo", "localhost:1337", geofence.Roam("cat", "dog", "*", 50).Actions(Enter, Exit)).toCmd(),
			Expected: "SETHOOK foo localhost:1337 NEARBY cat FENCE DETECT enter,exit ROAM dog * 50",
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.Expected, test.Cmd.String())
	}
}

func TestSearch(t *testing.T) {
	search := &Search{}
	tests := []struct {
		Cmd      cmd
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
		{
			Cmd:      search.Within("foo").Get("objID").Distance().WhereEval("bar").toCmd(),
			Expected: "WITHIN foo WHEREEVAL bar 0 DISTANCE GET objID",
		},
		{
			Cmd:      search.Search("foo").Asc().Cursor(5).Limit(5).Match("bar*").FormatIDs().toCmd(),
			Expected: "SEARCH foo MATCH bar* ASC CURSOR 5 LIMIT 5 IDS",
		},
		{
			Cmd:      search.Search("foo").Desc().Where("bar", 0, 20).FormatCount().toCmd(),
			Expected: "SEARCH foo WHERE bar 0 20 DESC COUNT",
		},
		{
			Cmd:      search.Scan("foo").Cursor(5).Limit(5).Asc().Match("bar").Wherein("baz", 10, 12, 13).NoFields().toCmd(),
			Expected: "SCAN foo WHEREIN baz 3 10 12 13 MATCH bar ASC NOFIELDS CURSOR 5 LIMIT 5",
		},
		{
			Cmd:      search.Scan("foo").Where("field", 0, 20).Limit(5).Desc().Match("bar").Format(FormatPoints).toCmd(),
			Expected: "SCAN foo WHERE field 0 20 MATCH bar DESC LIMIT 5 POINTS",
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.Expected, test.Cmd.String())
	}
}

func TestGeofence(t *testing.T) {
	geofence := &Geofence{}
	tests := []struct {
		Cmd      cmd
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
				Wherein("price", 20, 30).
				WhereEval("foo", "arg1", "arg2").toCmd(),
			Expected: "NEARBY agent WHEREIN price 2 20 30 WHEREEVAL foo 2 arg1 arg2 DISTANCE FENCE ROAM target * 100",
		},
		{
			Cmd: geofence.Within("foo").
				Bounds(10, 20, 30, 40).
				WhereEvalSHA("sha-hash", "arg1", "arg2").
				NoFields().
				Limit(1).toCmd(),
			Expected: "WITHIN foo WHEREEVALSHA sha-hash 2 arg1 arg2 NOFIELDS LIMIT 1 FENCE BOUNDS 10 20 30 40",
		},
		{
			Cmd:      geofence.Intersects("foo").Circle(10, 20, 30).Sparse(5).Where("param", 0, 100).Match("*").toCmd(),
			Expected: "INTERSECTS foo WHERE param 0 100 MATCH * SPARSE 5 FENCE CIRCLE 10 20 30",
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.Expected, test.Cmd.String())
	}
}

func TestKeys(t *testing.T) {
	keys := &Keys{}
	tests := []struct {
		Cmd      cmd
		Expected string
	}{
		{
			Cmd: keys.Set("agent", "49").
				PointZ(0, 0, -20).
				Field("age", 55).
				IfNotExists().
				Expiration(60 * 60 * 24 * 365).toCmd(),
			Expected: "SET agent 49 NX EX 31536000 FIELD age 55 POINT 0 0 -20",
		},
		{
			Cmd: keys.Set("agent", "47").
				PointZ(0, 0, -20).
				Field("age", 55).
				IfExists().Field("foo", 10).toCmd(),
			Expected: "SET agent 47 XX FIELD age 55 FIELD foo 10 POINT 0 0 -20",
		},
		{
			Cmd: keys.FSet("agent", "47").
				Field("cash", 100500).
				IfExists().toCmd(),
			Expected: "FSET agent 47 XX cash 100500",
		},
		{
			Cmd:      keys.JSet("foo", "bar", "some.field", "some-value").Raw().toCmd(),
			Expected: "JSET foo bar some.field some-value RAW",
		},
		{
			Cmd:      keys.JSet("foo", "bar", "some.field", "some-value").Str().toCmd(),
			Expected: "JSET foo bar some.field some-value STR",
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.Expected, test.Cmd.String())
	}
}
