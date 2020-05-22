package tests

import (
	"context"
	"testing"

	t38c "github.com/b3q/tile38-client"
	"github.com/b3q/tile38-client/transport"
	"github.com/stretchr/testify/assert"
)

func BenchmarkGeofence(b *testing.B) {
	client, err := t38c.NewWithExecutor(transport.NewMock(), false)
	assert.Nil(b, err)
	for i := 0; i < b.N; i++ {
		client.GeofenceIntersects("fleet").
			Bounds(-20, -20, 20.5, 29.6).
			Actions(t38c.Enter, t38c.Exit).
			Format(t38c.FormatPoints).
			Match("abc*").Wherein("speed", 5).
			Do(context.Background(), func(ev *t38c.GeofenceEvent) {})
	}
}
