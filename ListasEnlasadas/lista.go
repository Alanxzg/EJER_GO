package main

import "fmt"

// listas enlazadas simples
type ListasEnlazadas struct {
	Head *Nodo
}

// insertar el inicio de la lista
func (l *ListasEnlazadas) Agregar(valor int) {
	nuevoNodo := &Nodo{Valor: valor}
	nuevoNodo.Siguiente = l.Head
	l.Head = nuevoNodo
}

func (lis *ListasEnlazadas) InsertarInicio(valor int) {
	nuevoNodo := &Nodo{Valor: valor}
	nuevoNodo.Siguiente = lis.Head
	lis.Head = nuevoNodo
}

func (lis *ListasEnlazadas) InsertarFinal(valor int) {
	nuevoNodo := &Nodo{Valor: valor}
	if lis.Head == nil {
		lis.Head = nuevoNodo
	} else {
		actual := lis.Head
		for actual.Siguiente != nil {
			actual = actual.Siguiente
		}
		actual.Siguiente = nuevoNodo
	}
}

// imprimir la lista
func (lis *ListasEnlazadas) Imprimir() {
	actual := lis.Head
	for actual != nil {
		fmt.Print(actual.Valor, " ")
		actual = actual.Siguiente
	}
	fmt.Println()
}

// buscar, vificar, analizar si un valor existe en la lista
func (lis *ListasEnlazadas) Buscar(valor int) bool {
	actual := lis.Head
	for actual != nil {
		if actual.Valor == valor {
			return true
		}
		actual = actual.Siguiente
	}
	return false
}
