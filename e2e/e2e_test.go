package e2e

import (
	"os"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/sythang/t38c"
)

func TestE2E(t *testing.T) {
	skipE2E(t)

	client, err := t38c.New(t38c.Config{
		Address: os.Getenv("T38C_TEST_ADDR"),
		Debug:   true,
	})
	require.NoError(t, err)

	t.Run("TestKeys", func(t *testing.T) {
		testKeys(t, client)
	})
	t.Run("TestWithin", func(t *testing.T) {
		testWithin(t, client)
	})
	t.Run("TestGeofence", func(t *testing.T) {
		testGeofence(t, client)
	})
}

func skipE2E(tb testing.TB) {
	env := "T38C_TEST_E2E"

	tb.Helper()
	if ok, _ := strconv.ParseBool(os.Getenv(env)); !ok {
		tb.Skipf("Skipped. Set %s=1 to enable e2e test.", env)
	}
}
