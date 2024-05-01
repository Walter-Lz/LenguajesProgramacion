package main

import (
	"fmt"
)

type producto struct {
	nombre   string
	cantidad int
	precio   int
}
type listaProductos []producto

var lProductos listaProductos

const existenciaMinima int = 10 //la existencia mínima es el número mínimo debajo de el cual se deben tomar eventuales desiciones

func (l *listaProductos) agregarProducto(nombre string, cantidad int, precio int) {
	productoExistente, err := l.buscarProducto(nombre)
	if err == 0 && productoExistente != nil { // aumentar el stock en caso de existencia
		productoExistente.cantidad += cantidad
		productoExistente.precio = precio
	} else { // Agregar nuevo producto al stock
		*l = append(*l, producto{nombre: nombre, cantidad: cantidad, precio: precio})
	}

}

func (l *listaProductos) Agregar_Multiples_Producto(productos ...producto) {
	*l = append(*l, productos...) // Agregar los productos al final de la lista
}

func (l *listaProductos) buscarProducto(nombre string) (*producto, int) {
	for i, p := range *l {
		if p.nombre == nombre {
			return &(*l)[i], 0 // Producto encontrado, error = 0 - encontrado
		}
	}
	return nil, -1 // Producto no encontrado, error = -1
}

func (l *listaProductos) venderProducto(nombre string, cant int) {
	productoExistente, err := l.buscarProducto(nombre)
	var lProductosAuxiliar listaProductos
	if err == 0 && cant > 0 {
		if productoExistente.cantidad > cant {
			productoExistente.cantidad = productoExistente.cantidad - cant
			fmt.Println("Se han vendido", cant, "unidades del producto:", nombre)
		} else {
			if productoExistente.cantidad > 0 {
				fmt.Println("Solo se han vendido", productoExistente.cantidad, "unidades del producto:", nombre)
				fmt.Println("Eliminando el producto:", nombre, "de la lista de Productos.")
			} else {
				fmt.Println("No hay existencias del producto:", nombre)
			}
			for _, p := range *l {
				if p.nombre != nombre {
					lProductosAuxiliar = append(lProductosAuxiliar, p)
				}
			}
			*l = lProductosAuxiliar // Actualizar la lista original con la lista auxiliar
		}
	}
}

func (l *listaProductos) modificarPrecio(nombre string, precio int) {
	productoExistente, err := l.buscarProducto(nombre)
	if err == 0 && productoExistente != nil { // aumentar el stock en caso de existencia
		fmt.Println("Modificar precio del producto:", nombre, "nuevo precio:", precio)
		productoExistente.precio = precio
	} else {
		fmt.Println("El producto: ", nombre, ", no se encuentra en la lista de productos")
	}

}

func llenarDatos() {
	lProductos.agregarProducto("arroz", 15, 2500)
	lProductos.agregarProducto("frijoles", 4, 2000)
	lProductos.agregarProducto("leche", 8, 1200)
	lProductos.agregarProducto("Manteca", 5, 1000)
	lProductos.agregarProducto("cacao", 10, 1700)
	lProductos.agregarProducto("café", 12, 4500)
	lProductos.agregarProducto("azucar", 12, 4500)
	lProductos.agregarProducto("cacao", 10, 1700)
}
func (l *listaProductos) listarProductosMínimos() listaProductos {
	var lProductosAuxiliarMinima listaProductos
	var i int
	for i = 0; i < len(*l); i++ {
		if (*l)[i].cantidad <= existenciaMinima {
			lProductosAuxiliarMinima = append(lProductosAuxiliarMinima, producto{nombre: (*l)[i].nombre, cantidad: (*l)[i].cantidad, precio: (*l)[i].precio})
		}
	}
	// retorna una nueva lista con productos con existencia mínima
	return lProductosAuxiliarMinima
}

func main() {
	llenarDatos()
	fmt.Println("Lista de productos:", lProductos)
	fmt.Println("Realizando la venta de productos...")
	fmt.Println("")
	lProductos.venderProducto("arroz", 5)
	lProductos.venderProducto("cacao", 8)
	lProductos.venderProducto("leche", 9)
	fmt.Println("Lista de productos actualizada...")
	fmt.Println(lProductos)
	fmt.Println("")
	fmt.Println("Lista de productos con existencias mínimas...")
	fmt.Println(lProductos.listarProductosMínimos())
	fmt.Println("")
	fmt.Println("Haciendo uso de la función de agregar múltiples productos")
	lProductos.Agregar_Multiples_Producto(
		producto{nombre: "harina", cantidad: 10, precio: 3500},
		producto{nombre: "fideos", cantidad: 20, precio: 2500},
		producto{nombre: "refresco", cantidad: 15, precio: 4000},
	)
	fmt.Println(lProductos)
	fmt.Println("")
	lProductos.modificarPrecio("cacao", 3000)
	fmt.Println(lProductos)
}

/*						Ejecución del código

Lista de productos: [{arroz 15 2500} {frijoles 4 2000} {leche 8 1200} {Manteca 5 1000} {cacao 20 1700} {café 12 4500} {azucar 12 4500}]
Realizando la venta de productos...

Se han vendido 5 unidades del producto: arroz
Se han vendido 8 unidades del producto: cacao
Solo se han vendido 8 unidades del producto: leche
Eliminando el producto: leche de la lista de Productos.
Lista de productos actualizada...
[{arroz 10 2500} {frijoles 4 2000} {Manteca 5 1000} {cacao 12 1700} {café 12 4500} {azucar 12 4500}]

Lista de productos con existencias mínimas...
[{arroz 10 2500} {frijoles 4 2000} {Manteca 5 1000}]

Haciendo uso de la función de agregar múltiples productos
[{arroz 10 2500} {frijoles 4 2000} {Manteca 5 1000} {cacao 12 1700} {café 12 4500} {azucar 12 4500} {harina 10 3500} {fideos 20 2500} {refresco 15 4000}]

Modificar precio del producto: cacao nuevo precio: 3000
[{arroz 10 2500} {frijoles 4 2000} {Manteca 5 1000} {cacao 12 3000} {café 12 4500} {azucar 12 4500} {harina 10 3500} {fideos 20 2500} {refresco 15 4000}]

Process finished with the exit code 0

*/
