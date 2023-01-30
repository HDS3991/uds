package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUniqueArray(t *testing.T) {
	arr := []int{1, 2, 3, 3, 3, 3, 4, 5, 5, 5, 6, 7}
	expected := []int{1, 2, 3, 4, 5, 6, 7}
	assert.Equal(t, expected, UniqueArray(arr))
}

func TestIsContain(t *testing.T) {
	arr := []int{1, 2, 3, 3, 3, 3, 4, 5, 5, 5, 6, 7}
	assert.Equal(t, true, IsContain(arr, 3))
}
