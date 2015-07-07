package maths

import (
	"testing"

	"github.com/6amedev/go_sum/app"
	"github.com/stretchr/testify/assert"
)

func TestSub(t *testing.T) {
	assert.Equal(t, 5, maths.Sub(10, 5), "10 - 5 = 5")
	assert.Equal(t, -5, maths.Sub(0, 5), "0 - 5 = -5")
	assert.Equal(t, -9, maths.Sub(-4, 5), "-4 - 5 = -9")
}
