package main

import (
	"errors"
	"fmt"
)

type ListasEnlazadas struct {
	cabeza *Nodo
}

func (l *ListasEnlazadas) Agregar(valor int) {
	nuevoNodo := &Nodo{Valor: valor}
	nuevoNodo.Siguiente = l.cabeza
	l.cabeza = nuevoNodo
}

func (lis *ListasEnlazadas) InsertarInicio(valor int) {
	nuevoNodo := &Nodo{Valor: valor}
	nuevoNodo.Siguiente = lis.cabeza
	lis.cabeza = nuevoNodo
}

func (lis *ListasEnlazadas) InsertarFinal(valor int) {
	nuevoNodo := &Nodo{Valor: valor}
	if lis.cabeza == nil {
		lis.cabeza = nuevoNodo
	} else {
		actual := lis.cabeza
		for actual.Siguiente != nil {
			actual = actual.Siguiente
		}
		actual.Siguiente = nuevoNodo
	}
}

func (lis *ListasEnlazadas) Imprimir() {
	actual := lis.cabeza
	for actual != nil {
		fmt.Print(actual.Valor, " ")
		actual = actual.Siguiente
	}
	fmt.Println()
}

func (lis *ListasEnlazadas) Buscar(valor int) bool {
	actual := lis.cabeza
	for actual != nil {
		if actual.Valor == valor {
			return true
		}
		actual = actual.Siguiente
	}
	return false
}

func (lis *ListasEnlazadas) Eliminar(valor int) {
	if lis.cabeza == nil {
		return
	}
	if lis.cabeza.Valor == valor {
		lis.cabeza = lis.cabeza.Siguiente
		return
	}
	actual := lis.cabeza
	for actual.Siguiente != nil && actual.Siguiente.Valor != valor {
		actual = actual.Siguiente
	}
	if actual.Siguiente != nil {
		actual.Siguiente = actual.Siguiente.Siguiente
	}
}

// ejercicio 1: encontrar la longitud de la lista
func (lis *ListasEnlazadas) Longitud() int {
	contador := 0
	actual := lis.cabeza
	for actual != nil {
		contador++
		actual = actual.Siguiente
	}
	return contador
}

// ejercicio 2: encontrar el último dato
func (lis *ListasEnlazadas) Ultimo() (int, error) {
	if lis.cabeza == nil {
		return 0, errors.New("lista vacía")
	}
	actual := lis.cabeza
	for actual.Siguiente != nil {
		actual = actual.Siguiente
	}
	return actual.Valor, nil
}

// ejercicio 3: insertar después del penúltimo dato
func (lis *ListasEnlazadas) InsertarDespuesPenultimo(nuevo int) error {
	if lis.cabeza == nil || lis.cabeza.Siguiente == nil {
		return errors.New("lista debe tener al menos 2 elementos")
	}
	actual := lis.cabeza
	for actual.Siguiente.Siguiente != nil {
		actual = actual.Siguiente
	}
	// actual es el penúltimo
	nuevoNodo := &Nodo{Valor: nuevo}
	nuevoNodo.Siguiente = actual.Siguiente
	actual.Siguiente = nuevoNodo
	return nil
}

func (lis *ListasEnlazadas) InsertarDespues(buscado, nuevo int) {
	actual := lis.cabeza
	for actual != nil {
		if actual.Valor == buscado {
			nuevoNodo := &Nodo{Valor: nuevo}
			nuevoNodo.Siguiente = actual.Siguiente
			actual.Siguiente = nuevoNodo
			return
		}
		actual = actual.Siguiente
	}
}

func (lis *ListasEnlazadas) Invertir() {
	var prev, siguiente *Nodo
	actual := lis.cabeza
	for actual != nil {
		siguiente = actual.Siguiente
		actual.Siguiente = prev
		prev = actual
		actual = siguiente
	}
	lis.cabeza = prev
}

func (lis *ListasEnlazadas) EliminarDuplicados() {
	if lis.cabeza == nil {
		return
	}
	vistos := make(map[int]bool)
	vistos[lis.cabeza.Valor] = true
	actual := lis.cabeza
	for actual.Siguiente != nil {
		if vistos[actual.Siguiente.Valor] {
			actual.Siguiente = actual.Siguiente.Siguiente
		} else {
			vistos[actual.Siguiente.Valor] = true
			actual = actual.Siguiente
		}
	}

}

// ejercicio 4: Buscar el elemento central de la lista
func (lis *ListasEnlazadas) ElementoCentral() (int, error) {
	if lis.cabeza == nil {
		return 0, errors.New("lista vacía")
	}
	lento := lis.cabeza
	rapido := lis.cabeza
	for rapido != nil && rapido.Siguiente != nil {
		lento = lento.Siguiente
		rapido = rapido.Siguiente.Siguiente
	}
	return lento.Valor, nil
}
