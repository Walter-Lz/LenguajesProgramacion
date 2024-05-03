package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"sort"
	"strconv"
	"strings"
)

type infoCliente struct {
	nombre string
	correo string
	edad   int32
}
type listaClientes []infoCliente

var lClientes listaClientes

// Consulta #1
func (l *listaClientes) agregarCliente(nombre string, correo string, edad int32) {
	*l = append(*l, infoCliente{nombre: nombre, correo: correo, edad: edad})
}

// Funciones genéricas en este caso de tipo Any
func filter(list any, f func(any) bool) []any {
	result := make([]any, 0)
	switch reflect.TypeOf(list).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(list)
		for i := 0; i < s.Len(); i++ {
			if f(s.Index(i).Interface()) {
				result = append(result, s.Index(i).Interface())
			}
		}
	}
	return result
}

// Función map1 genérica de tipo Any
func map1(list any, f func(any) any) []any {
	result := make([]any, 0)
	switch reflect.TypeOf(list).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(list)
		for i := 0; i < s.Len(); i++ {
			result = append(result, f(s.Index(i).Interface()))
		}
	}
	return result
}

// Función Reduce genérica de tipo interface
func Reduce(slice interface{}, initial reflect.Value, reducer func(acc, elem reflect.Value) reflect.Value) reflect.Value {
	sliceValue := reflect.ValueOf(slice)
	if sliceValue.Kind() != reflect.Slice {
		return reflect.ValueOf(nil)
	}
	result := initial
	for i := 0; i < sliceValue.Len(); i++ {
		result = reducer(result, sliceValue.Index(i))
	}
	return result
}

// ------------------Consulta #2
func isApellidoInCorreo(cliente infoCliente, apellido string) bool {
	return strings.Contains(strings.ToLower(cliente.correo), strings.ToLower(apellido))
}
func listaClientes_ApellidoEnCorreo(listaC *listaClientes, apellido string) listaClientes {
	filtered := filter(*listaC, func(cliente interface{}) bool {
		c := cliente.(infoCliente)
		return isApellidoInCorreo(c, apellido)
	})
	if len(filtered) == 0 {
		return listaClientes{} // Devolver una lista vacía si no hay elementos después de filtrar
	}
	var clientesFiltrados listaClientes
	for _, cliente := range filtered {
		clientesFiltrados = append(clientesFiltrados, cliente.(infoCliente))
	}
	return clientesFiltrados
}

// ---------------------

// ------------------Consulta #3
func esCorreoCostaRica(correo string) bool {
	return strings.HasSuffix(strings.ToLower(correo), ".cr")
}

func cantidadCorreosCostaRica(clientes *listaClientes) int64 {
	// Mapear los correos de los clientes a booleanos que indican si son de Costa Rica
	correosCostaRica := map1(*clientes, func(cliente any) any {
		c := cliente.(infoCliente)
		return esCorreoCostaRica(c.correo)
	})
	// Filtrar los booleanos que representan los correos de Costa Rica
	correosCostaRicaFiltrados := filter(correosCostaRica, func(correo any) bool {
		return correo.(bool)
	})
	// Reducir la lista de booleanos a un solo valor que representa la cantidad total de correos de Costa Rica
	cantidad := Reduce(correosCostaRicaFiltrados, reflect.ValueOf(0), func(acc, elem reflect.Value) reflect.Value { // 	if elem.Interface().(bool) {
		if elem.Interface().(bool) {
			return reflect.ValueOf(acc.Int() + 1)
		}
		return acc
	}).Int() // va Convertir el resultado a int64
	return cantidad
}

// ------------------Consulta #4
// funcion para verificar el formato del correo
func contieneNombreEnCorreo(cliente infoCliente) bool {
	correoIniciales := strings.Split(cliente.correo, "@")[0]
	nombre := strings.Fields(cliente.nombre)[0]
	primeraLetra := strings.ToLower(string(nombre[0]))
	apellido1 := strings.ToLower(strings.Fields(cliente.nombre)[1])
	return strings.Contains(strings.ToLower(primeraLetra+apellido1), strings.ToLower(correoIniciales))
}

// Función para generar una sugerencia de correo
func sugerirCorreoConNombre(cliente infoCliente) infoCliente {
	partesNombre := strings.Fields(cliente.nombre)
	nombre := ""
	apellido := ""
	if len(partesNombre) >= 1 {
		nombre = partesNombre[0]
	}
	if len(partesNombre) >= 2 {
		apellido = partesNombre[1]
	}
	sugerencia := ""
	if len(nombre) > 0 && len(apellido) > 0 {
		sugerencia = strings.ToLower(string(nombre[0])) + strings.ToLower(apellido)
	}
	for len(sugerencia) < 8 { // Longitud total de 8 caracteres para el correo sugerido
		sugerencia += strconv.Itoa(rand.Intn(10)) // Agregar un dígito aleatorio
	}
	nuevoCorreo := sugerencia + "@" + strings.Split(cliente.correo, "@")[1]
	cliente.correo = nuevoCorreo

	return cliente
}

// Función clientesSugerenciasCorreos que utiliza map y filter genéricas para aplicar sugerencias de correo a los clientes
func clientesSugerenciasCorreos(clientes listaClientes) listaClientes {

	clientesParaSugerencia := filter(clientes, func(cliente interface{}) bool {
		return !contieneNombreEnCorreo(cliente.(infoCliente))
	})
	fmt.Println("Clientes Sugeridos: ", clientesParaSugerencia)
	// Aplicar map1 con la función de sugerirCorreoConNombre
	clientesConSugerencias := map1(clientesParaSugerencia, func(cliente interface{}) interface{} {
		return sugerirCorreoConNombre(cliente.(infoCliente))
	})
	clientesOriginales := filter(clientes, func(cliente interface{}) bool {
		return contieneNombreEnCorreo(cliente.(infoCliente))
	})
	// Si hay clientes con sugerencias, combinarlos con los clientes originales
	if len(clientesConSugerencias) > 0 {
		// Convertir el resultado de map1 a listaClientes
		var listadoClientes listaClientes

		for _, cliente := range clientesOriginales {
			listadoClientes = append(listadoClientes, cliente.(infoCliente))
		}
		for _, cliente := range clientesConSugerencias {
			listadoClientes = append(listadoClientes, cliente.(infoCliente))
		}
		return listadoClientes
	}

	// Si no hay clientes con sugerencias, devolver lista vacía
	return listaClientes{}
}

// ------------------Consulta #5
func correosOrdenadosAlfabeticos(clientes listaClientes) listaClientes {
	// Realizar una copia de la lista original para evitar modificarla
	clientesOrdenados := make(listaClientes, len(clientes))
	copy(clientesOrdenados, clientes)
	// Extraer los correos de los clientes en una lista separada
	correos := make([]string, len(clientes))
	for i, cliente := range clientesOrdenados {
		correos[i] = cliente.correo
	}
	// Ordenar los correos alfabéticamente
	sort.Strings(correos)
	// Crear un mapa para buscar clientes por correo
	clientePorCorreo := make(map[string]infoCliente)
	for _, cliente := range clientesOrdenados {
		clientePorCorreo[cliente.correo] = cliente
	}
	// Crear una nueva lista ordenada con la información de los clientes
	clientesOrdenadosNuevos := make(listaClientes, len(clientes))
	for i, correo := range correos {
		clientesOrdenadosNuevos[i] = clientePorCorreo[correo]
	}
	return clientesOrdenadosNuevos
}

func main() {
	fmt.Println("------------------#1------------------------")
	lClientes.agregarCliente("Oscar Viquez", "oviquez@tec.ac.cr", 44)
	lClientes.agregarCliente("Pedro Perez", "elsegundo@gmail.com", 30)
	lClientes.agregarCliente("Maria Lopez", "mlopez@hotmail.com", 18)
	lClientes.agregarCliente("Juan Rodriguez", "jrodriguez@gmail.com", 25)
	lClientes.agregarCliente("Luisa Gonzalez", "muyseguro@tec.ac.cr", 67)
	lClientes.agregarCliente("Marco Rojas", "loquesea@hotmail.com", 47)
	lClientes.agregarCliente("Marta Saborio", "msaborio@gmail.com", 33)
	lClientes.agregarCliente("Camila Segura", "csegura@ice.co.cr", 19)
	lClientes.agregarCliente("Fernando Rojas", "frojas@estado.gov", 27)
	lClientes.agregarCliente("Rosa Ramirez", "risuenna@gmail.com", 50)
	fmt.Println("Lista clientes: ", lClientes)

	fmt.Println("------------------#2------------------------")
	apellido := "Rojas"
	clientesConApellidoEnCorreo := listaClientes_ApellidoEnCorreo(&lClientes, apellido) // Pasar puntero a la lista
	fmt.Printf("Clientes cuyos correos contienen el apellido '%s':\n", apellido)
	if clientesConApellidoEnCorreo != nil && len(clientesConApellidoEnCorreo) > 0 {
		for _, cliente := range clientesConApellidoEnCorreo {
			fmt.Printf("Nombre: %s, Correo: %s\n", cliente.nombre, cliente.correo)
		}
	} else {
		fmt.Printf("No se encontraron clientes cuyos correos contengan el apellido '%s'.\n", apellido)
	}

	fmt.Println("------------------#3------------------------")
	ClientesCostaRica := cantidadCorreosCostaRica(&lClientes)
	fmt.Printf("Cantidad de clientes con correos de dominio de Costa Rica: %d\n", ClientesCostaRica)

	fmt.Println("------------------#4------------------------")
	clientesConSugerencias := clientesSugerenciasCorreos(lClientes)
	if len(clientesConSugerencias) > 0 {
		// Imprimir la lista actualizada de clientes
		fmt.Println("Nueva Lista de clientes con los correos sugeridos:")
		fmt.Println(clientesConSugerencias)
	} else {
		fmt.Println("No hay clientes con sugerencia de correo")
	}

	fmt.Println("------------------#5------------------------")
	correosOrdenados := correosOrdenadosAlfabeticos(lClientes)

	// Imprimir los correos ordenados
	fmt.Println("Correos ordenados alfabéticamente:")
	for _, cliente := range correosOrdenados {
		fmt.Printf("Cliente: %s, Correo: %s, Edad: %d\n", cliente.nombre, cliente.correo, cliente.edad)
	}
}

/*       Ejecución del código

------------------#1------------------------
Lista clientes:  [{Oscar Viquez oviquez@tec.ac.cr 44} {Pedro Perez elsegundo@gmail.com 30} {Maria Lopez mlopez@hotmail.com 18} {Juan Rodriguez jrodriguez@gmail.com 25} {Luisa Gonzalez muyseguro@tec.ac.cr 67} {Marco Rojas loquesea@hotmail.com 47} {Marta Saborio msaborio@gmail.com 33} {Camila Segura csegura@ice.co.cr 19} {Fernando Rojas frojas@estado.gov 27} {Rosa Ramirez risuenna@gmail.com 50}]
------------------#2------------------------
Clientes cuyos correos contienen el apellido 'Rojas':
Nombre: Fernando Rojas, Correo: frojas@estado.gov
------------------#3------------------------
Cantidad de clientes con correos de dominio de Costa Rica: 3
------------------#4------------------------
Clientes Sugeridos:  [{Pedro Perez elsegundo@gmail.com 30} {Luisa Gonzalez muyseguro@tec.ac.cr 67} {Marco Rojas loquesea@hotmail.com 47} {Rosa Ramirez risuenna@gmail.com 50}]
Nueva Lista de clientes con los correos sugeridos:
[{Oscar Viquez oviquez@tec.ac.cr 44} {Maria Lopez mlopez@hotmail.com 18} {Juan Rodriguez jrodriguez@gmail.com 25} {Marta Saborio msaborio@gmail.com 33} {Camila Segura csegura@ice.co.cr 19} {Fernando Rojas frojas@estado.gov 27} {Pedro Perez pperez46@gmail.com 30} {Luisa Gonzalez lgonzalez@tec.ac.cr 67} {Marco Rojas mrojas25@hotmail.com 47} {Rosa Ramirez rramirez@gmail.com 50}]
------------------#5------------------------
Correos ordenados alfabéticamente:
Cliente: Camila Segura, Correo: csegura@ice.co.cr, Edad: 19
Cliente: Pedro Perez, Correo: elsegundo@gmail.com, Edad: 30
Cliente: Fernando Rojas, Correo: frojas@estado.gov, Edad: 27
Cliente: Juan Rodriguez, Correo: jrodriguez@gmail.com, Edad: 25
Cliente: Marco Rojas, Correo: loquesea@hotmail.com, Edad: 47
Cliente: Maria Lopez, Correo: mlopez@hotmail.com, Edad: 18
Cliente: Marta Saborio, Correo: msaborio@gmail.com, Edad: 33
Cliente: Luisa Gonzalez, Correo: muyseguro@tec.ac.cr, Edad: 67
Cliente: Oscar Viquez, Correo: oviquez@tec.ac.cr, Edad: 44
Cliente: Rosa Ramirez, Correo: risuenna@gmail.com, Edad: 50

Process finished with the exit code 0
*/
