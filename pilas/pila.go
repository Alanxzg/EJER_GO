package main

import "fmt"

// Pila representa una pila de nodos
type Stack struct {
	Peak *Nodo
	Size int
}

func NewStack() *Stack {
	return &Stack{
		Peak: nil,
		Size: 0,
	}
}

func (stack *Stack) push(value interface{}) {
	nuevoNodo := &Nodo{
		Value: value,
		Next:  stack.Peak,
	}
	stack.Peak = nuevoNodo
	stack.Size++
}

// Pop elimina y retorna el elemento del tope de la pila
func (stack *Stack) Pop() (interface{}, error) {
	if stack.IsEmpty() {
		return nil, fmt.Errorf("error: pila vacía")
	}
	valor := stack.Peak.Value
	stack.Peak = stack.Peak.Next
	stack.Size--
	return valor, nil
}

// Peek retorna el elemento del tope sin eliminarlo
func (stack *Stack) Peek() (interface{}, error) {
	if stack.IsEmpty() {
		return nil, fmt.Errorf("error: pila vacía")
	}
	return stack.Peak.Value, nil
}

// IsEmpty verifica si la pila está vacía
func (stack *Stack) IsEmpty() bool {
	return stack.Peak == nil
}

// Size retorna el número de elementos en la pila
func (stack *Stack) GetSize() int {
	return stack.Size
}

// Imprimir muestra todos los elementos de la pila
func (stack *Stack) Imprimir() {
	if stack.IsEmpty() {
		fmt.Println("Pila vacía")
		return
	}
	actual := stack.Peak
	fmt.Print("Tope -> ")
	for actual != nil {
		fmt.Printf("%v -> ", actual.Value)
		actual = actual.Next
	}
	fmt.Println("nil")
}

// Vaciar elimina todos los elementos de la pila
func (stack *Stack) Vaciar() {
	stack.Peak = nil
	stack.Size = 0
}

func (stack *Stack) InvertirPila() {
	var prev, siguiente *Nodo
	actual := stack.Peak
	for actual != nil {
		siguiente = actual.Next
		actual.Next = prev
		prev = actual
		actual = siguiente
	}
	stack.Peak = prev
}

func (stack *Stack) ParEImpar() (pares, impares []int) {
	actual := stack.Peak
	for actual != nil {
		if num := actual.Value.(int); num != 0 {
			if num%2 == 0 {
				pares = append(pares, num)
			} else {
				impares = append(impares, num)
			}
		}
		actual = actual.Next
	}
	return pares, impares
}
