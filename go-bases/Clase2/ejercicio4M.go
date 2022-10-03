package main

import (
	"errors"
	"fmt"
)

/*
Se solicita generar una función que indique qué tipo de cálculo se quiere realizar (mínimo, máximo o promedio)
y que devuelva otra función ( y un mensaje en caso que el cálculo no esté definido) que se le puede pasar una
cantidad N de enteros y devuelva el cálculo que se indicó en la función anterior
*/

const (
	minimum = "minimum"
	average = "average"
	maximum = "maximum"
)

func minFunction(num ...int) int {
	var minimo int

	for i, numero := range num {
		if i == 0 {
			minimo = numero
		}

		if numero < minimo {
			minimo = numero
		}
	}

	return minimo
}

func averageFunction(num ...int) int {
	var promedio int

	for _, numero := range num {
		promedio += numero
	}

	return promedio / len(num)
}

func maxFunction(num ...int) int {
	var maximo int

	for i, numero := range num {
		if i == 0 {
			maximo = numero
		}

		if numero > maximo {
			maximo = numero
		}
	}

	return maximo
}

func operation(nombreFuncion string) (func(num ...int) int, error) {
	switch nombreFuncion {
	case "minimum":
		return minFunction, nil
	case "average":
		return averageFunction, nil
	case "maximum":
		return maxFunction, nil
	default:
		return nil, errors.New("La operación no es válida")
	}
}

func main() {

	minFunc, err := operation(minimum)
	if err != nil {
		fmt.Println(err)
		return
	}

	averageFunc, err := operation(average)
	if err != nil {
		fmt.Println(err)
		return
	}

	maxFunc, err := operation(maximum)
	if err != nil {
		fmt.Println(err)
		return
	}

	minValue := minFunc(2, 3, 3, 4, 10, 2, 4, 5)
	averageValue := averageFunc(2, 3, 3, 4, 1, 2, 4, 5)
	maxValue := maxFunc(2, 3, 3, 4, 1, 2, 4, 5)

	fmt.Println("Valor minimo:", minValue)
	fmt.Println("Valor maximo:", maxValue)
	fmt.Println("Valor promedio:", averageValue)

}
