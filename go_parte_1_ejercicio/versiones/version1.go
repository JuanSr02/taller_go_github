package versiones

import (
	"fmt"
	"math/rand"
	"time"
)

func Version1() {
	// Crear una nueva fuente de aleatoriedad
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	puntajes := map[int]int{
		1: 0,
		2: 0,
		3: 0,
		4: 0,
		5: 0,
	}

	for range 10 {
		// Generar un número aleatorio entre 1 y 5
		numeroAleatorio := r.Intn(5) + 1
		puntajes[numeroAleatorio] = puntajes[numeroAleatorio] + 1
	}

	for key, value := range puntajes {
		if value==1{
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
