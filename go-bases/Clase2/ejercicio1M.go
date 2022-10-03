package main

import "fmt"

/*
Una empresa de chocolates necesita calcular el impuesto de sus empleados al momento de depositar el sueldo,
para cumplir el objetivo es necesario crear una función que devuelva el impuesto de un salario.
Teniendo en cuenta que si la persona gana más de $50.000 se le descontará un 17% del sueldo y si gana más
de $150.000 se le descontará además un 10%.
*/

func calcularImpuesto(sueldo float64) float64 {
	if sueldo > 150000 {
		return sueldo * 0.27
	}

	if sueldo > 50000 {
		return sueldo * 0.17
	}

	return 0
}

func main() {

	var sueldoEjemplo float64

	sueldoEjemplo = 200000

	fmt.Println("Impuesto: ", calcularImpuesto(sueldoEjemplo))

}
