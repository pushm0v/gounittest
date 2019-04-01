package toko

import (
	"testing"

	"gotest.tools/assert"
)

func TestMinumanStorage(t *testing.T) {
	ts := NewTokoStorage()

	ts.UpdateStok("some", 3)
	result := ts.GetStok("some")

	assert.Equal(t, 3, result)
}

func TestMinumanStorageIsAvailable(t *testing.T) {
	ts := NewTokoStorage()

	ts.UpdateStok("some", 3)
	result := ts.IsStokAvailable("some", 3)

	assert.Equal(t, true, result)
}

func TestMinumanStorageIsNotAvailable(t *testing.T) {
	ts := NewTokoStorage()

	result := ts.IsStokAvailable("some", 3)

	assert.Equal(t, false, result)
}
