package main

import "fmt"

/*
Para ello es necesario generar una estructura Alumnos con las variables Nombre, Apellido,
DNI, Fecha y que tenga un método detalle
*/

type Alumno struct {
	Nombre   string
	Apellido string
	DNI      int64
	Fecha    string
}

func (v Alumno) detalle() {
	fmt.Println("Nombre: ", v.Nombre)
	fmt.Println("Apellido: ", v.Apellido)
	fmt.Println("DNI:", v.DNI)
	fmt.Println("Fecha: ", v.Fecha)
}

func main() {
	alumno := Alumno{"Ivan", "Buhajeruk", 41626347, "02-10-2022"}

	alumno.detalle()
}
