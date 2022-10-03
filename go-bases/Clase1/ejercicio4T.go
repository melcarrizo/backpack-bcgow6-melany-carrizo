package main

import "fmt"

/*
Un empleado de una empresa quiere saber el nombre y edad de uno de sus empleados. Según el siguiente mapa, ayuda  a imprimir la edad de Benjamin.

  var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}

Por otro lado también es necesario:
Saber cuántos de sus empleados son mayores de 21 años.
Agregar un empleado nuevo a la lista, llamado Federico que tiene 25 años.
Eliminar a Pedro del mapa.

*/

func main() {

	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}
	var cantidad int

	fmt.Println("La edad de Benjamin es: ", employees["Benjamin"])

	for _, element := range employees {
		if element > 21 {
			cantidad++
		}
	}

	fmt.Println("Empleados mayores a 21 años: ", cantidad)

	employees["Federico"] = 25

	delete(employees, "Pedro")

}
