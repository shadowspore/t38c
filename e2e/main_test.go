package e2e

import (
	"log"
	"os"
	"testing"
	"time"

	t38c "github.com/axvq/tile38-client"
	"github.com/ory/dockertest/v3"
	geojson "github.com/paulmach/go.geojson"
	"github.com/stretchr/testify/assert"
)

var client *t38c.Client

func TestMain(m *testing.M) {
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	// pulls an image, creates a container based on it and runs it
	resource, err := pool.Run("tile38/tile38", "latest", nil)
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	// exponential backoff-retry, because the application in the container might not be ready to accept connections yet
	if err := pool.Retry(func() error {
		var err error
		client, err = t38c.New("localhost:"+resource.GetPort("9851/tcp"), t38c.Debug)
		if err != nil {
			return err
		}
		return client.Ping()
	}); err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	code := m.Run()

	// You can't defer this because os.Exit doesn't care for defer
	if err := pool.Purge(resource); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}

	os.Exit(code)
}

func TestKeys(t *testing.T) {
	if err := client.Keys.Set("foo", "bar").Point(0, 0).Do(); err != nil {
		t.Error(err)
	}

	resp, err := client.Keys.Get("foo", "bar", false)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, geojson.NewPointGeometry([]float64{0, 0}), resp.Object.Geometry)

	if err := client.Keys.Set("foo", "baz").Bounds(0, 0, 20, 20).Field("age", 20).Expiration(5).Do(); err != nil {
		t.Error(err)
	}

	resp, err = client.Keys.Get("foo", "baz", true)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, geojson.NewPolygonGeometry([][][]float64{{
		{0, 0},
		{20, 0},
		{20, 20},
		{0, 20},
		{0, 0},
	}}), resp.Object.Geometry)
	time.Sleep(time.Second * 5)
	_, err = client.Keys.Get("foo", "baz", false)
	assert.Error(t, err)
}
