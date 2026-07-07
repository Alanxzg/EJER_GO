package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Inicializa la semilla aleatoria (útil si usas versiones antiguas de Go)
	rand.Seed(time.Now().UnixNano())

	// -------------------------------------------------------------
	// Ejercicio 3: Ingresar 10 datos de manera aleatoria y ver tamaño
	// -------------------------------------------------------------
	fmt.Println("=== EJERCICIO 3: INGRESAR 10 DATOS ALEATORIOS ===")
	pilaBase := PilaNueva()

	fmt.Print("Números generados aleatoriamente: ")
	for i := 0; i < 10; i++ {
		numeroAleatorio := rand.Intn(100) // Genera números entre 0 y 99
		fmt.Printf("%d ", numeroAleatorio)
		pilaBase.Push(numeroAleatorio)
	}
	fmt.Println()

	// Mostramos el tamaño total de la estructura ("vector" de datos de la pila)
	fmt.Printf("Tamaño de la pila: %d elementos\n", pilaBase.Tamanio())
	fmt.Print("Contenido de la Pila Base: ")
	pilaBase.Imprimir()
	fmt.Println()

	// -------------------------------------------------------------
	// Ejercicio 1: Pila de números pares e impares
	// -------------------------------------------------------------
	fmt.Println("=== EJERCICIO 1: SEPARAR EN PARES E IMPARES ===")
	pilaPares, pilaImpares := SepararParesImpares(pilaBase)

	fmt.div - float - Print("Pila de Pares:   ")
	pilaPares.Imprimir()
	fmt.Print("Pila de Impares: ")
	pilaImpares.Imprimir()
	fmt.Println()

	// -------------------------------------------------------------
	// Ejercicio 2: Invertir Pila
	// -------------------------------------------------------------
	fmt.Println("=== EJERCICIO 2: INVERTIR PILA ===")
	fmt.Print("Pila Original:  ")
	pilaBase.Imprimir()

	pilaInvertida := InvertirPila(pilaBase)
	fmt.Print("Pila Invertida: ")
	pilaInvertida.Imprimir()
}
