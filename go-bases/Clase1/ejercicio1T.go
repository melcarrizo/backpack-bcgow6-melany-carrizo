package main

import (
	"fmt"
)

/*
La Real Academia Española quiere saber cuántas letras tiene una palabra y luego tener cada una de las letras por separado para deletrearla.
Crear una aplicación que tenga una variable con la palabra e imprimir la cantidad de letras que tiene la misma.
Luego imprimí cada una de las letras.

*/

func main() {

	var palabra string
	var cantidad int

	palabra = "prueba"
	cantidad = len(palabra)

	fmt.Printf("Cantidad de letras de la palabra: %d\n", cantidad)

	for i := 0; i < cantidad; i++ {
		fmt.Println(string(palabra[i]))
	}

}
