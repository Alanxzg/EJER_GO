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

// Tamanio retorna el número de elementos
func (pila *Pila) Tamanio() int {
	return pila.tamanio
}

// Imprimir muestra el contenido de la pila
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

// Vaciar elimina todos los elementos
func (pila *Pila) Vaciar() {
	pila.tope = nil
	pila.tamanio = 0
}

// ==========================================
// NUEVAS FUNCIONES PARA LOS EJERCICIOS
// ==========================================

// Ejercicio 1: Recibe una pila y devuelve dos pilas (Pares e Impares)
func SepararParesImpares(pila *Pila) (*Pila, *Pila) {
	pares := PilaNueva()
	impares := PilaNueva()

	actual := pila.tope
	for actual != nil {
		// Validamos que el valor de la interfaz sea un entero
		if valorInt, ok := actual.Valor.(int); ok {
			if valorInt%2 == 0 {
				pares.Push(valorInt)
			} else {
				impares.Push(valorInt)
			}
		}
		actual = actual.Siguiente
	}
	return pares, impares
}

// Ejercicio 2: Invierte el orden de los elementos de una pila
func InvertirPila(p *Pila) *Pila {
	nuevaPila := PilaNueva()
	actual := p.tope

	// Al ir recorriendo desde el tope e insertando en la nueva pila,
	// el orden se invierte automáticamente de forma natural.
	for actual != nil {
		nuevaPila.Push(actual.Valor)
		actual = actual.Siguiente
	}
	return nuevaPila
}
