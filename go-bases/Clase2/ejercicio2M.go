package main

import "fmt"

/*
Un colegio necesita calcular el promedio (por alumno) de sus calificaciones.
Se solicita generar una función en la cual se le pueda pasar N cantidad de enteros y devuelva el
promedio y un error en caso que uno de los números ingresados sea negativo
*/

func calcularPromedio(notas ...int) (float64, string) {
	var promedio float64

	for _, nota := range notas {
		if nota > 0 {
			promedio += float64(nota)
		} else {
			return 0, fmt.Sprintf("Nota inváida")
		}

	}
	promedio /= float64(len(notas))

	return promedio, ""
}

func main() {
	promedio, mensaje := calcularPromedio(2, 3, -2, 3, 2)

	if promedio == 0 {
		fmt.Println(mensaje)
	} else {
		fmt.Println("El promedio es de: ", promedio)
	}

}
