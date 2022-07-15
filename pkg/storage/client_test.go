package storage

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClient(t *testing.T) {
	c := NewClient("/gravity", "localhost:2379")
	assert.NotNil(t, c)
	assert.Panics(t, func() {
		NewClient("/gravity")
	})
}
