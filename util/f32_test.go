package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Absf32(t *testing.T) {
	assert.Equal(t, float32(1), Absf32(-1))
	assert.Equal(t, float32(1), Absf32(1))
}
