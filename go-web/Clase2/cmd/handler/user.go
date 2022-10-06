package handler

import (
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/melcarrizo/backpack-bcgow6-melany-carrizo/go-web/Clase2/internal/users"
)

// Validacion de datos

type request struct {
	Nombre   string  `json:"nombre"`
	Apellido string  `json:"apellido"`
	Email    string  `json:"email"`
	Fecha    string  `json:"fecha"`
	Edad     int     `json:"edad"`
	Altura   float64 `json:"altura"`
	Activo   bool    `json:"activo"`
}

type User struct {
	service users.Service
}

func NewUser(p users.Service) *User {
	return &User{
		service: p,
	}
}

func (c *User) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(401, gin.H{
				"error": "token inválido",
			})
			return
		}

		p, err := c.service.GetAll()
		if err != nil {
			ctx.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(200, p)
	}
}

func (c *User) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(401, gin.H{"error": "token inválido"})
			return
		}
		var req request
		if err := ctx.Bind(&req); err != nil {
			ctx.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}
		p, err := c.service.Store(req.Nombre, req.Apellido, req.Email, req.Fecha, req.Edad, req.Altura, req.Activo)
		if err != nil {
			ctx.JSON(404, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(200, p)
	}
}

func (c *User) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(401, gin.H{"error": "token inválido"})
			return
		}
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(400, gin.H{"error": "invalid ID"})
			return
		}
		var req request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}
		if req.Altura <= 0 {
			ctx.JSON(400, gin.H{"error": "Campo altura es requerido"})
			return
		}

		if req.Edad <= 0 {
			ctx.JSON(400, gin.H{"error": "Campo edad es requerido"})
			return
		}

		if req.Apellido == "" {
			ctx.JSON(400, gin.H{"error": "Campo apellido es requerido"})
			return
		}

		if req.Fecha == "" {
			ctx.JSON(400, gin.H{"error": "Campo fecha es requerido"})
			return
		}

		if req.Nombre == "" {
			ctx.JSON(400, gin.H{"error": "Campo nombre es requerido"})
			return
		}

		p, err := c.service.Update(int(id), req.Nombre, req.Apellido, req.Email, req.Fecha, req.Edad, req.Altura, req.Activo)
		if err != nil {
			ctx.JSON(404, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(200, p)
	}
}
