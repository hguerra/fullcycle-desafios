package math_test

import (
	"integracao-continua/math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSomar(t *testing.T) {
	assert.Equal(t, 30, math.Somar(15, 15))
}
