package timestamp_test

import (
	"testing"

	"github.com/cyrildever/period/timestamp"
	"gotest.tools/assert"
)

// TestForceAndInTestEnvironment ...
func TestForceAndInTestEnvironment(t *testing.T) {
	ts, err := timestamp.CurrentMillis(true)
	if err != nil {
		t.Fatal(err)
	}
	assert.Assert(t, ts > 0)

	_, err = timestamp.CurrentMillis(false)
	assert.Error(t, err, "ntp.Time has failed: timestamp can't be returned")

	timestamp.InTestEnvironment = true
	ts, err = timestamp.CurrentMillis()
	if err != nil {
		t.Fatal(err)
	}
	assert.Assert(t, ts > 0)
}
