package main

import "fmt"

// Realizar una aplicación que contenga una variable con el número del mes.

func main() {

	var meses = map[string]int{
		"Enero":      1,
		"Febrero":    2,
		"Marzo":      3,
		"Abril":      4,
		"Mayo":       5,
		"Junio":      6,
		"Julio":      7,
		"Agosto":     8,
		"Septiembre": 9,
		"Octubre":    10,
		"Noviembre":  11,
		"Diciembre":  12}

	fmt.Println(meses["Enero"])

}
