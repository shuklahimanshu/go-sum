package maths

import (
	"testing"

	"github.com/6ameDev/go-sum/app"
	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	assert.Equal(t, 15, maths.Add(10, 5), "10 + 5 = 15")
	assert.Equal(t, 5, maths.Add(0, 5), "0 + 5 = 5")
	assert.Equal(t, 1, maths.Add(-4, 5), "-4 + 5 = 1")
}
