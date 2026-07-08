package main

type Nodo struct {
	Valor     interface{}
	Siguiente *Nodo
}

func NuevoNodo(valor interface{}) *Nodo {
	return &Nodo{
		Valor:     valor,
		Siguiente: nil,
	}
}
