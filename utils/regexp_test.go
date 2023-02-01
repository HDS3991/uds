package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckPostcode(t *testing.T) {
	assert.Equal(t, true, CheckPostcode("110000"))
	assert.Equal(t, false, CheckPostcode("1210000"))
	assert.Equal(t, false, CheckPostcode("11000"))
	assert.Equal(t, false, CheckPostcode("北京"))
	assert.Equal(t, false, CheckPostcode("110000s"))
	assert.Equal(t, true, CheckPostcode(110000))
	var a float64 = 110001
	assert.Equal(t, true, CheckPostcode(a))
	b := 110001.000
	assert.Equal(t, true, CheckPostcode(b))
}
