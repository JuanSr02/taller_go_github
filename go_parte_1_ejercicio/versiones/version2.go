package versiones

import (
	"fmt"
	"strings"
)

func Version2() {
	puntajes := map[int]int{
		1: 0,
		2: 0,
		3: 0,
		4: 0,
		5: 0,
	}
	var input string
	var numeroIngresado int
	valido := false
	for range 10 {
		// Pedir un numero al usuario
		for !valido {
			fmt.Print("Introduce un número entero entre 1 y 5: ")
			_, _ = fmt.Scanln(&input)

			// Limpiar la entrada de espacios
			input = strings.TrimSpace(input)

			// Intentar convertir a entero
			_, err := fmt.Sscanf(input, "%d", &numeroIngresado)

			if err != nil {
				fmt.Println("Error: Debes introducir un número entero válido.")
				continue
			}

			// Verificar rango
			if numeroIngresado >= 1 && numeroIngresado <= 5 {
				valido = true
			} else {
				fmt.Println("Error: El número debe estar entre 1 y 5.")
			}
		}
		puntajes[numeroIngresado] = puntajes[numeroIngresado] + 1
		valido = false
	}

	for key, value := range puntajes {
		if value == 1 {
			fmt.Printf("El valor: %d fue ingresado %d vez\n", key, value)
		} else {
			fmt.Printf("El valor: %d fue ingresado %d veces\n", key, value)
		}
	}

	sum1_2 := puntajes[1] + puntajes[2]
	sum4_5 := puntajes[4] + puntajes[5]

	if sum4_5 >= sum1_2 {
		fmt.Println("¡Buen resultado!")
	} else {
		fmt.Println("Resultado mejorable")
	}

}
