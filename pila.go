package main

import "fmt"

type Pila struct {
	tope    *Nodo
	tamanio int
}

func PilaNueva() *Pila {
	return &Pila{
		tope:    nil,
		tamanio: 0,
	}
}

func (p *Pila) Push(valor interface{}) {
	nuevoNodo := NuevoNodo(valor)
	nuevoNodo.Siguiente = p.tope
	p.tope = nuevoNodo
	p.tamanio++
}

func (pila *Pila) Pop() (interface{}, error) {
	if pila.EstaVacia() {
		return nil, fmt.Errorf("la pila está vacía")
	}
	valor := pila.tope.Valor
	pila.tope = pila.tope.Siguiente
	pila.tamanio--
	return valor, nil
}

func (pila *Pila) EstaVacia() bool {
	return pila.tope == nil
}

func (pila *Pila) Tamanio() int {
	return pila.tamanio
}

func (pila *Pila) Imprimir() {
	if pila.EstaVacia() {
		fmt.Println("Pila vacía")
		return
	}
	actual := pila.tope
	fmt.Print("Tope -> ")
	for actual != nil {
		fmt.Printf("%v -> ", actual.Valor)
		actual = actual.Siguiente
	}
	fmt.Println("nil")
}

func (pila *Pila) Vaciar() {
	pila.tope = nil
	pila.tamanio = 0
}

func ContarPares(pila *Pila) int {
	contador := 0
	actual := pila.tope
	for actual != nil {
		if valor, ok := actual.Valor.(int); ok {
			if valor%2 == 0 {
				contador++
			}
		}
		actual = actual.Siguiente
	}
	return contador
}

func SumarElementos(pila *Pila) int {
	suma := 0
	actual := pila.tope
	for actual != nil {
		if valor, ok := actual.Valor.(int); ok {
			suma += valor
		}
		actual = actual.Siguiente
	}
	return suma
}

func ElementoMaximo(pila *Pila) (int, error) {
	if pila.EstaVacia() {
		return 0, fmt.Errorf("la pila está vacía")
	}
	actual := pila.tope
	maximo := 0
	encontrado := false
	for actual != nil {
		if valor, ok := actual.Valor.(int); ok {
			if !encontrado || valor > maximo {
				maximo = valor
				encontrado = true
			}
		}
		actual = actual.Siguiente
	}
	if !encontrado {
		return 0, fmt.Errorf("la pila no contiene enteros")
	}
	return maximo, nil
}

func EstaOrdenada(pila *Pila) bool {
	actual := pila.tope
	if actual == nil || actual.Siguiente == nil {
		return true
	}
	for actual.Siguiente != nil {
		valorActual, ok1 := actual.Valor.(int)
		valorSiguiente, ok2 := actual.Siguiente.Valor.(int)
		if !ok1 || !ok2 {
			return false
		}
		if valorActual < valorSiguiente {
			return false
		}
		actual = actual.Siguiente
	}
	return true
}
