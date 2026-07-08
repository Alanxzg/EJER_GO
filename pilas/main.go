package main

import "fmt"

func main() {

	Stack := NewStack()

	fmt.Println("CREACION DE UNA PILA")

	fmt.Print("Agregando datos al Stack... ")

	Stack.push(10)
	Stack.push(20)
	Stack.push(30)
	Stack.push(40)
	Stack.push(50)
	Stack.push(60)
	Stack.push(70)
	Stack.push(80)
	Stack.push(90)
	Stack.push(100)

	fmt.Println("Datos agregados correctamente")
	Stack.Imprimir()

	fmt.Println("Buscando pares e impares")
	Stack.ParEImpar()
	Impares, Pares := Stack.ParEImpar()
	fmt.Println("Pares:", Pares)
	fmt.Println("Impares:", Impares)

	Stack.InvertirPila()
	fmt.Println("Pila invertida: ")
	Stack.Imprimir()

	Stack.GetSize()
	fmt.Println("Tamaño de la pila:", Stack.GetSize())

}
