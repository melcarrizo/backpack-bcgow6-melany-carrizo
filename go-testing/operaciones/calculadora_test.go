package operaciones

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRestar(t *testing.T) {
	n1 := 10
	n2 := 5
	esperado := 8

	resultado := Restar(n1, n2)

	assert.Equal(t, esperado, resultado, "El resultado no es el esperado.")

}
