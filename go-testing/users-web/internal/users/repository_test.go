package users

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAll(t *testing.T) {
	//arrange
	repo := NewRepository()
	esperado := []User{
		{
			Nombre:   "Ivan",
			Apellido: "Buhaje",
			Email:    "ib@gmail.com",
			Fecha:    "20-10-2022",
			Edad:     22,
			Altura:   1.78,
			Activo:   true,
		},
	}

	//act
	resultado, _ := repo.GetAll()

	//assert
	assert.Equal(t, esperado, resultado)
}
