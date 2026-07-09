package main

import "fmt"

func main() {
	lista := ListasEnlazadas{}
	lista.InsertarFinal(10)
	lista.InsertarFinal(20)
	lista.InsertarFinal(5)
	lista.InsertarInicio(50)
	lista.Imprimir() // 50 10 20 5

	lista.Eliminar(20)
	lista.Imprimir() // 50 10 5

	// ejercicio 1: longitud
	listaLongitud := lista.Longitud()
	fmt.Println("Longitud:", listaLongitud)

	// ejercicio 2: último valor
	ultimo, err := lista.Ultimo()
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Último:", ultimo)
	}

	// ejercicio 3: insertar después del penúltimo
	lista.InsertarDespuesPenultimo(99)
	lista.Imprimir() // 50 10 99 5

	// ejercicio 4: elemento central
	central, err := lista.ElementoCentral()
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Elemento central:", central)
	}

}
