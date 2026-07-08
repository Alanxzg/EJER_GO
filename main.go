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

	fmt.Println("\n===== CÓDIGO 3: Encontrar el elemento máximo =====")
	pila3 := PilaNueva()
	pila3.Push(45)
	pila3.Push(12)
	pila3.Push(78)
	pila3.Push(23)
	pila3.Push(56)
	fmt.Print("Pila:  ")
	pila3.Imprimir()
	maximo, err := ElementoMaximo(pila3)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Resultado: %d\n", maximo)
	}
	fmt.Print("Pila después: ")
	pila3.Imprimir()

	fmt.Println("\n===== CÓDIGO 4: ¿Pila ordenada ascendente? =====")
	pila4 := PilaNueva()
	pila4.Push(1)
	pila4.Push(3)
	pila4.Push(5)
	pila4.Push(7)
	fmt.Print("Pila:  ")
	pila4.Imprimir()
	fmt.Printf("Resultado: %t\n", EstaOrdenada(pila4))
	fmt.Print("Pila después: ")
	pila4.Imprimir()
}
