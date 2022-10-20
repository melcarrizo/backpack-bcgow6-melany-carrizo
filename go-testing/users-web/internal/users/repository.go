package users

import (
	"fmt"
)

type User struct {
	Id       int     `json:"id"`
	Nombre   string  `json:"nombre"`
	Apellido string  `json:"apellido"`
	Email    string  `json:"email"`
	Fecha    string  `json:"fecha"`
	Edad     int     `json:"edad"`
	Altura   float64 `json:"altura"`
	Activo   bool    `json:"activo"`
}

var ps = []User{
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
var lastID int

// ***Importado**//
type Repository interface {
	GetAll() ([]User, error)
	Store(id int, nombre, apellido, email, fecha string, edad int, altura float64, activo bool) (User, error)
	LastID() (int, error)
	Update(id int, nombre, apellido, email, fecha string, edad int, altura float64, activo bool) (User, error)
}

type repository struct{}

//struct implementa los metodos de la interfaz

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) Store(id int, nombre, apellido, email, fecha string, edad int, altura float64, activo bool) (User, error) {
	p := User{id, nombre, apellido, email, fecha, edad, altura, activo}
	ps = append(ps, p)
	lastID = p.Id

	return p, nil
}

func (r *repository) GetAll() ([]User, error) {
	return ps, nil
}

func (r *repository) LastID() (int, error) {
	return lastID, nil
}

func (r *repository) Update(id int, nombre, apellido, email, fecha string, edad int, altura float64, activo bool) (User, error) {
	p := User{Nombre: nombre, Apellido: apellido, Email: email, Fecha: fecha, Edad: edad, Altura: altura, Activo: activo}
	updated := false
	for i := range ps {
		if ps[i].Id == id {
			p.Id = id
			ps[i] = p
			updated = true
		}
	}
	if !updated {
		return User{}, fmt.Errorf("Usuario %d no encontrado", id)
	}
	return p, nil
}
