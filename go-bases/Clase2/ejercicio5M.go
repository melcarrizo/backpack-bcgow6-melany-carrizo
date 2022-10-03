package main

import (
	"errors"
	"fmt"
)

/*
perro necesitan 10 kg de alimento
gato 5 kg
Hamster 250 gramos.
Tarántula 150 gramos.

Se solicita:
Implementar una función Animal que reciba como parámetro un valor de tipo texto con el animal especificado y que retorne
una función y un mensaje (en caso que no exista el animal)
Una función para cada animal que calcule la cantidad de alimento en base a la cantidad del tipo de animal especificado.
*/

const (
	dog       = "dog"
	cat       = "cat"
	hamster   = "hamster"
	tarantula = "tarantula"
)

func animalDog(cantidad int) float64 {
	return float64(cantidad) * 10
}

func animalCat(cantidad int) float64 {
	return float64(cantidad) * 5
}

func animalHamster(cantidad int) float64 {
	return float64(cantidad) * 0.25
}

func animalTarantula(cantidad int) float64 {
	return float64(cantidad) * 0.15
}

func Animal(nombre string) (func(cantidad int) float64, error) {
	switch nombre {
	case dog:
		return animalDog, nil
	case cat:
		return animalCat, nil
	case hamster:
		return animalHamster, nil
	case tarantula:
		return animalTarantula, nil
	default:
		return nil, errors.New("El nombre de animal no es válido")
	}
}

func main() {
	var amount float64

	animalDog, msg := Animal(dog)
	if msg != nil {
		fmt.Println(msg)
		return
	}

	animalCat, msg := Animal("mono")
	if msg != nil {
		fmt.Println(msg)
		return
	}

	amount += animalDog(5)
	amount += animalCat(8)

	fmt.Println("Se debera comprar: ", amount)

}
