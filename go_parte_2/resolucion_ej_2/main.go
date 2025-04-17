package main

import (
	"ejercicio2/dipositivo"
)

func main() {
	l := []dipositivo.Controlable{
		dipositivo.New("A"),
		dipositivo.New("B"),
		dipositivo.New("C"),
		dipositivo.New("D"),
		dipositivo.New("E"),
		dipositivo.New("F"),
		dipositivo.New("G"),
		dipositivo.New("H"),
		dipositivo.New("I"),
		dipositivo.New("J"),
	}

	for _, controlable := range l {
		controlable.EstadoActual()
	}
}
