package main

import "fmt"

/*
Las empresas tienen 3 tipos de productos:
Pequeño, Mediano y Grande. (Se espera que sean muchos más)
Existen costos adicionales por mantener el producto en el almacén de la tienda, y costos de envío.

Sus costos adicionales son:
Pequeño: El costo del producto (sin costo adicional)
Mediano: El costo del producto + un 3% por mantenerlo en existencia en el almacén de la tienda.
Grande: El costo del producto + un 6%  por mantenimiento, y un costo adicional  por envío de $2500.

Requerimientos:
Crear una estructura “tienda” que guarde una lista de productos.
Crear una estructura “producto” que guarde el tipo de producto, nombre y precio
Crear una interface “Producto” que tenga el método “CalcularCosto”
Crear una interface “Ecommerce” que tenga los métodos “Total” y “Agregar”.
Se requiere una función “nuevoProducto” que reciba el tipo de producto, su nombre y precio y devuelva un Producto.
Se requiere una función “nuevaTienda” que devuelva un Ecommerce.
Interface Producto:
El método “CalcularCosto” debe calcular el costo adicional según el tipo de producto.
Interface Ecommerce:
 - El método “Total” debe retornar el precio total en base al costo total de los productos y los adicionales si los hubiera.
 - El método “Agregar” debe recibir un producto y añadirlo a la lista de la tienda
*/

var (
	pequenio string = "pequeño"
	mediano  string = "mediano"
	grande   string = "grande"
)

type producto struct {
	tipo   string
	nombre string
	precio float64
}

type tienda struct {
	productos []producto
}

type Producto interface {
	CalcularCosto() float64
}

type Ecommerce interface {
	Total() float64
	Agregar(nuevo producto)
}

func (p producto) CalcularCosto() float64 {
	switch p.tipo {
	case pequenio:
		return p.precio
	case mediano:
		return p.precio * 1.03
	case grande:
		return p.precio * 1.06
	default:
		return 0
	}
}

func nuevoProducto(tipo, nombre string, precio float64) Producto {

	return &producto{tipo: tipo, nombre: nombre, precio: precio}
}

func nuevaTienda() Ecommerce {
	return &tienda{}
}

func (t tienda) Agregar(nuevo producto) {
	t.productos = append(t.productos, nuevo)
}

func (t tienda) Total() float64 {
	var total float64

	for i := 0; i < len(t.productos); i++ {
		total += t.productos[i].CalcularCosto()
	}

	return total
}

func main() {

	ejemploProducto := producto{tipo: mediano, nombre: "capuccino", precio: 100}
	ejemploTienda := nuevaTienda()

	ejemploTienda.Agregar(ejemploProducto)

	total := ejemploTienda.Total()
	fmt.Println("Total ", total)

}
