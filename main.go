package main

import "fmt"

func main() {
	fmt.Println("===== CÓDIGO 1: Contar elementos pares =====")
	pila1 := PilaNueva()
	pila1.Push(5)
	pila1.Push(10)
	pila1.Push(7)
	pila1.Push(20)
	pila1.Push(3)
	fmt.Print("Pila:  ")
	pila1.Imprimir()
	fmt.Printf("Resultado: %d\n", ContarPares(pila1))
	fmt.Print("Pila después: ")
	pila1.Imprimir()

	fmt.Println("\n===== CÓDIGO 2: Sumar todos los elementos =====")
	pila2 := PilaNueva()
	pila2.Push(10)
	pila2.Push(20)
	pila2.Push(30)
	pila2.Push(40)
	fmt.Print("Pila:  ")
	pila2.Imprimir()
	fmt.Printf("Resultado: %d\n", SumarElementos(pila2))
	fmt.Print("Pila después: ")
	pila2.Imprimir()
}
