package main

import "fmt"

func main() {

	// hola mundo
	fmt.Println("hola mundo")

	// declaración de variables
	declarandoVariable := "hola"
	var otraVariable string

	fmt.Println(declarandoVariable, otraVariable)

	// for
	for i := 0; i < 10; i++ {
		fmt.Printf("imprimiendo i: %d \n", i)
	}

	// ifs
	var hola bool
	if !hola {
		fmt.Println("no me dijiste hola")
	} else {
		fmt.Println("chau")
	}

	// listas
	var array []string
	otroArray := []string{"ya tiene un elemento"}
	array = append(array, "agrego un elemento al array")

	fmt.Println(array, otroArray)

	// dict o maps
	var esteEsUnMapa map[string]string
	esteEsOtroMapa := map[string]string{}
	estaEsOtraForma := make(map[string]string)

	mapaConElemento := map[string]string{
		"hola": "chau",
	}

	fmt.Println("leo un elemento ", mapaConElemento["hola"])
	fmt.Println(esteEsOtroMapa, esteEsUnMapa, estaEsOtraForma, mapaConElemento)

	estaEsOtraForma["otro_elemento"] = "sarasa"
	fmt.Println(estaEsOtraForma)

	// ranges
	arrayConVariosElementos := []string{
		"hola",
		"chau",
		"chau",
		"chau",
	}

	mapaConVariosElementos := map[string]string{
		"hola":   "chau",
		"chau":   "chau",
		"random": "chau",
		"sarasa": "chau",
	}

	for i, v := range arrayConVariosElementos {
		fmt.Println("index: ", i, " elemento: ", v)
	}

	for k, v := range mapaConVariosElementos {
		fmt.Println("key: ", k, " value: ", v)
	}

	// tipos básicos
	var (
		anyObj           any
		float32Obj       float32
		float64Obj       float64
		stringObj        string
		intObj           int
		intObjEspecifico int64
		uintObj          uint
		boolObj          bool
		punteroObj       *string
	)

	fmt.Println(
		anyObj,
		float32Obj,
		float64Obj,
		stringObj,
		intObj,
		intObjEspecifico,
		uintObj,
		boolObj,
		punteroObj,
	)

	simpleFunc()
	funcWithParams("hola")
	funcWithVariadic("hola", "param1", "param2", "param3")
	fmt.Println(funcConRetornoSimple("hola"))
	fmt.Println(funcConDobleRetorno("hola"))

	// YAPA para el ejercicio:

	// agarrar el input del usuario (por consola)
	var x int

	// tengo que pasar el puntero de int para que escriba sobre la dirección de memoria y no el literal
	fmt.Println("acá podes ingresar un número por consola y capturarlo en la variable x")
	fmt.Scan(&x)

	fmt.Println("el numero ingresado desde la consola fue ", x)
}

func simpleFunc() {
	fmt.Println("funcion simple")
}

func funcWithParams(param1 string) {
	fmt.Println("funcion con parametro ", param1)
}

func funcWithVariadic(param1 string, variadic ...string) {
	fmt.Println("funcion con variadic ", param1, " variadic: ", variadic)
}

func funcConRetornoSimple(param1 string) string {
	fmt.Println("devuelve un string")
	return param1 + "_otro_string"
}

func funcConDobleRetorno(param1 string) (string, int) {
	return param1 + "_otro_string", 0
}
