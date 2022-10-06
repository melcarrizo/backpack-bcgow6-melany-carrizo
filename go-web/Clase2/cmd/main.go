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
1. Se debe validar todos los campos enviados en la petición, todos los campos son
requeridos
2. En caso que algún campo no esté completo se debe retornar un código de error 400
con el mensaje “el campo %s es requerido”.
(En %s debe ir el nombre del campo que no está completo).
1. Al momento de enviar la petición se debe validar que un token sea enviado
2. Se debe validar ese token en nuestro código (el token puede estar hardcodeado).
3. En caso que el token enviado no sea correcto debemos retornar un error 401 y un
mensaje que “no tiene permisos para realizar la petición solicitada”.
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
		token := c.GetHeader("token")
		if token != "123456" || token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "token invalido"})
			return
		}

		var req User
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if req.Altura <= 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Campo altura es requerido"})
			return
		}

		if req.Edad <= 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Campo edad es requerido"})
			return
		}

		if req.Apellido == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Campo apellido es requerido"})
			return
		}

		if req.Fecha == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Campo fecha es requerido"})
			return
		}

		if req.Nombre == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Campo nombre es requerido"})
			return
		}

		req.Id = len(users) + 1
		users = append(users, req)
		c.JSON(http.StatusOK, req)
	}
}
