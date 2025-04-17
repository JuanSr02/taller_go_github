package main

import (
	"errors"
	"fmt"
	"github.com/fatih/color"
	"go_parte_2/matematica"
	"strings"
)

func main() {
	// ------- manejo de dependencias -------
	fmt.Println("Hello World")
	color.Green("Hello World")

	// ------- manejo de punteros -------
	// En Go, los punteros nos permiten trabajar directamente con la dirección de memoria de una variable.
	// En lugar de copiar un valor, podemos acceder y modificar el valor original desde otra parte del programa.
	//
	// ¿Por qué es útil esto? Porque cuando pasamos variables a funciones, por defecto Go trabaja
	// con copias. Si queremos modificar el valor original desde dentro de una función, necesitamos pasar un puntero.
	//
	// Un puntero es simplemente una variable que apunta a otra. Lo declaramos usando el operador *,
	// y obtenemos la dirección de una variable con &.
	litString := "Hello world"
	var punteroString *string

	fmt.Println(litString, punteroString)

	punteroString = &litString
	fmt.Println(punteroString)

	// desreferencia de un puntero
	fmt.Println(*punteroString)

	//litString = nil
	punteroString = nil

	// que pasa si se desreferencia un nil?
	//fmt.Println(*punteroString)

	p := "HELLO WORLD"
	litFunc(p)
	fmt.Println("Con literales: ", p)

	punteroFunc(&p)
	fmt.Println("Con punteros: ", p)

	// ------- manejo de errores -------
	// En Go, el manejo de errores no se hace con excepciones como en otros lenguajes. Acá, los errores se manejan
	// de forma explícita, retornando un valor del tipo error desde las funciones.
	//
	// Esto obliga al programador a tratar los errores conscientemente, lo cual es una decisión de diseño pensada
	// para mejorar la robustez del software.
	//
	// El tipo error en Go es una interfaz incorporada que tiene un único método: Error() string.
	// Es decir, cualquier tipo que implemente ese método puede ser tratado como un error.
	//
	// Cuando una función puede fallar, se define así: func hacerAlgo() (Resultado, error). Y cuando la usamos,
	// debemos verificar si el error está presente: res, err := hacerAlgo(); if err != nil { // manejar error }.
	//
	// Esto es muy idiomático en Go. El chequeo de errores se hace inmediatamente después de llamar a la
	// función, lo que mantiene el código predecible y seguro.
	//
	// Además, vos mismo podés crear tus propios errores personalizados, ya sea con errors.New(...),
	// fmt.Errorf(...), o implementando la interfaz error con un struct propio.
	//
	// En resumen: el manejo de errores en Go es simple, explícito y confiable. Y entender la interfaz
	// error nos permite extenderlo de forma flexible y clara.
	var err error
	fmt.Println("error -> ", err)

	err = errors.New("this is an error")
	fmt.Println("error -> ", err)

	manejoErr(err)
	fmt.Println("error -> ", err)

	manejoErrPuntero(&err)
	fmt.Println("error -> ", err)

	// sin error
	s, err := manejaString("hello")
	if err != nil {
		fmt.Println("tuve un error llamando a manejaString. error -> ", err)
	} else {
		fmt.Println("resultado de manejaString -> ", s)
	}

	// con error
	s, err = manejaString("hello world")
	if err != nil {
		fmt.Println("tuve un error llamando a manejaString. error -> ", err)
	} else {
		fmt.Println("resultado de manejaString -> ", s)
	}

	// ------- estructuras -------
	var persona1 persona
	//fmt.Println("persona1 -> ", persona1)
	fmt.Printf("persona1 -> %#v\n", persona1)

	persona1 = persona{}
	fmt.Printf("persona1 -> %#v\n", persona1)

	persona1 = persona{
		Nombre:   "Ayrton",
		Apellido: "Marini",
		Apodo:    "Chiche",
		Edad:     31,
	}
	fmt.Printf("persona1 -> %#v\n", persona1)

	persona1.Edad = 99
	persona1.Nombre = "NN"
	persona1.Apellido = "NN"
	fmt.Printf("persona1 -> %#v\n", persona1)

	fmt.Printf("nombre y apellido -> %s\n", persona1.NombreApellido())

	persona1.ModificarEdad(31)
	fmt.Printf("persona1 -> %#v\n", persona1)

	persona1.Nombre = "Ayrton"
	persona1.Apellido = "Marini"
	presentacion(&persona1)

	otraStruct := OtraStruct{}
	presentacion(&otraStruct)

	// packaging
	// En Go, la visibilidad se controla con la primera letra del nombre.
	// Si empieza con mayúscula, es pública (exportada). Si empieza con minúscula, es privada al paquete.
	//
	// Esto aplica a funciones, structs, métodos y atributos. Es una forma simple y efectiva de encapsular lógica
	// y proteger tu código de uso indebido desde otros paquetes.
	//
	// Además, usar paquetes ayuda a organizar el código y crear componentes reutilizables.

	resultado := matematica.Sumar(10, 5)
	fmt.Println("Resultado de suma:", resultado)

	fmt.Println("Valor de Pi:", matematica.Pi)
	fmt.Println("Versión del paquete:", matematica.Version)

	matematica.MostrarAutor()

	// fmt.Println(matematica.e)          // ❌ Error: e no es accesible
	// fmt.Println(matematica.autor)      // ❌ Error: autor no es accesible
	// fmt.Println(matematica.restar(10, 5)) // ❌ Error: restar no es accesible

	calc := matematica.Calculadora{}

	res := calc.Sumar(10, 15)
	fmt.Println("Resultado:", res)
	// fmt.Println("Memoria:", calc.memoria) // ❌ Error: memoria no es accesible

	fmt.Println("Memoria:", calc.MostrarMemoria())
}

func litFunc(s string) {
	s = strings.ToLower(s)
}

func punteroFunc(s *string) {
	*s = strings.ToLower(*s)
}

func manejoErr(err error) {
	err = errors.New("seteado en la func manejoErr")
}

func manejoErrPuntero(err *error) {
	*err = errors.New("seteado en la func manejoErr")
}

// manejaString es una funcion que toma como parámetro un string y lo devuelve en mayus
//
// En caso de que la longitud del string sea mayor a 8, arroja un error.
func manejaString(s string) (string, error) {
	if len(s) > 8 {
		return "", errors.New("la longitud del string es mayor a 8")
	}

	return strings.ToUpper(s), nil
}

func presentacion(p Presentable) {
	fmt.Println(p.Presentar())
}

// En Go no usamos clases como en otros lenguajes orientados a objetos. En cambio, usamos structs,
// que son estructuras que nos permiten agrupar distintos tipos de datos bajo un mismo nombre.
//
// Una struct es como un molde o plantilla para representar una entidad con atributos. Por ejemplo,
// si queremos representar una persona, podemos tener una struct llamada Persona que tenga nombre,
// edad y correo electrónico.
//
// Esto nos da una forma clara y ordenada de trabajar con datos que están relacionados entre sí.
// Y lo mejor es que, sobre una struct, también podemos definir métodos, que son funciones
// asociadas a esa estructura.
//
// De esta manera, aunque Go no tiene clases ni herencia, podemos modelar el mundo real de forma
// estructurada, con datos y comportamientos, usando structs y métodos.
//
// Las structs son la base para construir estructuras más complejas, y nos ayudan a organizar mejor
// nuestro código y hacerlo más legible.
type persona struct {
	Nombre   string
	Apellido string
	Apodo    string
	Edad     int
}

func (p *persona) Presentar() string {
	return fmt.Sprintf("Hola, mi nombre es %s", p.NombreApellido())
}

func (p persona) NombreApellido() string {
	return fmt.Sprintf("%s - %s", p.Nombre, p.Apellido)
}

func (p *persona) ModificarEdad(edad int) {
	p.Edad = edad
}

// Hasta ahora estuvimos trabajando con estructuras en Go, que nos permiten agrupar datos relacionados.
// Pero… ¿qué pasa cuando queremos hablar de comportamientos más que de datos?
//
// Acá es donde entran las interfaces. Una interfaz en Go define un conjunto de métodos que una estructura
// debe implementar. Es como un contrato: si tu struct implementa todos los métodos que la interfaz declara,
// entonces satisface esa interfaz, sin necesidad de decirlo explícitamente.
//
// Por ejemplo, si yo tengo una interfaz Describible que tiene un método Describir() string, cualquier struct
// que tenga ese método va a cumplir con esa interfaz automáticamente. No tengo que hacer ningún implements
// ni herencia ni nada raro.
//
// Esto nos permite escribir código más flexible y desacoplado. Podemos escribir funciones que trabajen con
// cualquier cosa que sea 'Describible', sin importar si es una Persona, un Producto o un Animal
//
// En resumen, las interfaces nos permiten programar pensando en lo que hace un objeto, más que en qué
// tipo de objeto es. Eso es muy poderoso cuando diseñamos software que queremos que sea mantenible,
// escalable y reutilizable.
// "If it walks like a duck and it quacks like a duck, then it must be a duck
type Presentable interface {
	Presentar() string
}

type OtraStruct struct{}

func (o *OtraStruct) Presentar() string {
	return "Hola, mi nombre es anónimo"
}
