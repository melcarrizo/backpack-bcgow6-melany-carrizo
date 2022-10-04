package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	Id       int
	Nombre   string
	Apellido string
	Email    string
	Edad     int
	Altura   float64
	Activo   bool
	Fecha    string
}

var users = []User{
	{
		Id:       1111,
		Nombre:   "Mauro",
		Apellido: "Rodriguez",
		Email:    "mrod@gmail.com",
		Edad:     33,
		Altura:   1.43,
		Activo:   false,
		Fecha:    "23-07-2022"},
}

func GetAll(ctx *gin.Context) {
	var filtrados []*User

	for _, u := range users {
		if ctx.Query("Id") == "u.Id" {
			filtrados = append(filtrados, &u)
		}
	}

	//ctx.String(http.StatusOK, "ID: %s Nombre: %s", Id, Nombre)
	ctx.JSON(http.StatusOK, filtrados)
}

func main() {

	router := gin.Default()

	// EJERCICIO 2
	router.GET("/ep", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Hola Melany!",
		})
	})

	//EJERCICIO 3
	router.GET("/users", GetAll)

	router.Run()

}

// endpoint: http://localhost:8080/ep
// endpoint: http://localhost:8080/users
