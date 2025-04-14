package matematica

import "fmt"

// Constante pública
const Pi = 3.14159

// Constante privada
const e = 2.71828

// Variable pública
var Version = "1.0"

// Variable privada
var autor = "Ayrton"

// Función pública
func Sumar(a, b int) int {
	return a + b
}

// Función privada
func restar(a, b int) int {
	return a - b
}

// Función pública
func MostrarAutor() {
	fmt.Println("Autor del paquete:", autor)
}

type Calculadora struct {
	memoria int
}

func (c *Calculadora) Sumar(a, b int) int {
	c.memoria = a + b
	return c.memoria
}

func (c *Calculadora) MostrarMemoria() int {
	return c.memoria
}
