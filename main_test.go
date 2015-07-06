package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	// Using stock testing framework
	var tests = []struct {
		a, b, actual int
	}{
		{3, 4, 7},
		{0, 0, 0},
		{0, 5, 5},
	}

	for _, test := range tests {
		got := add(test.a, test.b)
		if got != test.actual {
			t.Error("add(%q, %q) == %q, want %q", test.a, test.b, got, test.actual)
		}
	}

	// Using Testify testing framework
	assert.True(t, true, "True is true!")
	assert.Equal(t, 7, add(3, 4), "3 + 4 is 7")
	assert.Equal(t, 0, add(0, 0), "0 + 0 is 0")
	assert.Equal(t, 5, add(0, 5), "0 + 5 is 5")
}
