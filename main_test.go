package main

import (
	"testing"

	"github.com/6amedev/go_sum/sum"
	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	assert.Equal(t, 7, sum.Add(3, 4), "3 + 4 is 7")
	assert.Equal(t, 0, sum.Add(0, 0), "0 + 0 is 0")
	assert.Equal(t, 5, sum.Add(0, 5), "0 + 5 is 5")

	assert.Equal(t, -1, sum.Sub(3, 4), "3 - 4 is -1")
	assert.Equal(t, 0, sum.Sub(0, 0), "0 - 0 is 0")
	assert.Equal(t, -5, sum.Sub(0, 5), "0 - 5 is -5")
}
