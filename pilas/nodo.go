package main

// Nodo representa un elemento de la pila
type Nodo struct {
	Valor     interface{} // Puede almacenar cualquier tipo
	Siguiente *Nodo
}

// NuevoNodo crea un nuevo nodo con el valor dado
func NuevoNodo(valor interface{}) *Nodo {
	return &Nodo{
		Valor:     valor,
		Siguiente: nil,
	}
}
