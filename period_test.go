package period_test

import (
	"testing"

	"github.com/cyrildever/period"
	"gotest.tools/assert"
)

// TestGetDuration ...
func TestGetDuration(t *testing.T) {
	var ref int64 = 10000
	span := uint64(ref)
	err := period.Init(0, span, true)
	if err != nil {
		t.Fatal(err)
	}
	assert.Assert(t, period.IsTestEnvironment())
	d := period.GetDuration().Milliseconds()
	assert.Equal(t, ref, d)
}

// TestNext ...
func TestNext(t *testing.T) {
	err := period.Init(0, 10000, true)
	if _, ok := err.(*period.AlreadyInitializedError); err != nil && !ok {
		t.Fatal(err)
	}
	p, err := period.Get(1)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, p.ID, uint64(1)) // First period is 1

	next := p.Next()
	assert.Equal(t, next.ID, p.ID+1)
}

// TestResetOriginTimestamp ...
func TestResetOriginTimestamp(t *testing.T) {
	err := period.Init(0, 10000, true)
	if _, ok := err.(*period.AlreadyInitializedError); err != nil && !ok {
		t.Fatal(err)
	}
	p, err := period.Get(1)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, p.StartTimestampMillis(), uint64(0))

	var newTimestamp uint64 = 123
	period.ResetOriginTimestamp(newTimestamp)
	assert.Equal(t, p.StartTimestampMillis(), newTimestamp)
}
