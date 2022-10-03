package main

import "fmt"

/*
Una empresa marinera necesita calcular el salario de sus empleados basándose en la cantidad
de horas trabajadas por mes y la categoría.

Si es categoría C, su salario es de $1.000 por hora
Si es categoría B, su salario es de $1.500 por hora más un %20 de su salario mensual
Si es de categoría A, su salario es de $3.000 por hora más un %50 de su salario mensual

Se solicita generar una función que reciba por parámetro la cantidad de minutos trabajados
por mes y la categoría, y que devuelva su salario.
*/

func calcularSalario(minutosTrabajados float64, categoria string) float64 {
	var salario float64
	var horas float64

	horas = minutosTrabajados / 60
	switch categoria {
	case "A":
		salario = horas * 3000
		salario *= 1.5
	case "B":
		salario = horas * 1500
		salario *= 1.2
	case "C":
		salario = horas * 1000
	default:
		fmt.Println("Categoría incorrecta")
	}

	return salario

}

func main() {

	fmt.Println("El salario es de: ", calcularSalario(60, "B"))

}
