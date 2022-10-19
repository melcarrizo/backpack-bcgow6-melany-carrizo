package main

import (
	"fmt"

	"github.com/melcarrizo/backpack-bcgow6-melany-carrizo/go-testing/operaciones"
)

func main() {
	a, b := 10, 5
	sum := operaciones.Sumar(a, b)
	fmt.Println(sum)
}
