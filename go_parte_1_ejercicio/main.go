package main

import (
	"ejercicios/versiones"
	"fmt"
)

// Una empresa realizó una encuesta de satisfacción entre sus clientes. Cada cliente dejó su opinión en forma de puntaje del 1 al 5 (siendo 5 el mejor).

/*Guardar en una lista los 10 primeros puntajes recibidos.
● Contar cuántas veces aparece cada puntaje.
● Mostrar por pantalla:
	○ La cantidad de veces que se votó cada puntaje.
	○ Si hay más votos positivos (4 o 5) que negativos (1 o 2), imprimí "¡Buen resultado!", si no, imprimí "Resultado mejorable".
● Subir a un repositorio de ustedes en github. */

func main() {
	fmt.Println("Iniciando Version 1")
	versiones.Version1()
	fmt.Println("\n Iniciando Version 2")
	versiones.Version2()
}
