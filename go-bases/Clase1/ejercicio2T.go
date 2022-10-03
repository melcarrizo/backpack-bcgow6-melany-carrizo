package main

import "fmt"

/*
Un banco quiere otorgar préstamos a sus clientes, pero no todos pueden acceder a los mismos.
Para ello tiene ciertas reglas para saber a qué cliente se le puede otorgar.
Solo le otorga préstamos a clientes cuya edad sea mayor a 22 años, se encuentren empleados y tengan más de
un año de antigüedad en su trabajo. Dentro de los préstamos que otorga no les cobrará interés a los que su
sueldo es mejor a $100.000.
Es necesario realizar una aplicación que tenga estas variables y que imprima un mensaje de acuerdo a cada caso.
*/

func main() {

	var edad int
	var empleado bool
	var antiguedad int
	var sueldo int

	edad = 23
	empleado = true
	antiguedad = 15
	sueldo = 15000

	if edad > 22 && empleado == true && antiguedad > 1 {
		if sueldo > 100000 {
			fmt.Println("Puede acceder al préstamo sin interés.")
		} else {
			fmt.Println("Puede acceder al préstamo con interés.")
		}
	} else {
		fmt.Println("No cumple los requisitos para solicitar préstamo.")
	}

}
