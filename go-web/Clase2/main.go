package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
Se debe implementar la funcionalidad para crear la entidad. pasa eso se deben seguir los
siguientes pasos:
1. Crea un endpoint mediante POST el cual reciba la entidad.
2. Se debe tener un array de la entidad en memoria (a nivel global), en el cual se
deberán ir guardando todas las peticiones que se vayan realizando.
3. Al momento de realizar la petición se debe generar el ID. Para generar el ID se debe
buscar el ID del último registro generado, incrementarlo en 1 y asignarlo a nuestro
nuevo registro (sin tener una variable de último ID a nivel global).
*/

type User struct {
	Id       int     `json:"id"`
	Nombre   string  `json:"nombre"`
	Apellido string  `json:"apellido"`
	Email    string  `json:"email"`
	Edad     int     `json:"edad"`
	Altura   float64 `json:"altura"`
	Activo   bool    `json:"activo"`
	Fecha    string  `json:"fecha"`
}

var (
	users []User
)

func main() {
	router := gin.Default()

	router.POST("/users", Store())

	router.Run()
}

func Store() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req User
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		req.Id = len(users) + 1
		users = append(users, req)
		c.JSON(http.StatusOK, req)
	}
}
