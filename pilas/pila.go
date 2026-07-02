package main

import "fmt"

type Pila struct {
	tope   *Nodo
	tamaño int
}

//PUSH

func (p *Pila) Push(valor interface{}) {
	nuevoNodo := NuevoNodo(valor)
	nuevoNodo.Siguiente = p.tope
	p.tope = nuevoNodo
	p.tamaño++
}

//POP

func (pila *Pila) Pop() (interface{}, error) {
	if pila.EstaVacia() {
		return nil, fmt.Errorf("la pila está vacía")
	}
}

func (pila *Pila) EstaVacia() bool {
	return pila.tope == nil
}
