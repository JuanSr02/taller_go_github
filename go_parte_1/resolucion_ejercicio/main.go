package main

import "fmt"

const (
	_votesQuantity = 10
	_voteMin       = 1
	_voteMax       = 5
)

func main() {
	votesQuantity := 0
	votes := map[int]int{}
	for {
		var vote int
		fmt.Println("votos realizados hasta ahora: ", votesQuantity)
		fmt.Println("Ingrese su voto: ")
		_, _ = fmt.Scanf("%d", &vote)

		if vote < _voteMin || vote > _voteMax {
			fmt.Println("...el voto debe ser un número entre el 1 y el 5...")
			continue
		}

		v, ok := votes[vote]
		if !ok {
			votes[vote] = 1
		} else {
			votes[vote] = v + 1
		}

		votesQuantity++
		if votesQuantity == _votesQuantity {
			break
		}
	}

	votosPositivos := 0
	votosNegativos := 0
	fmt.Println("Cantidad de votos por puntaje: ")
	for puntaje, cantidadVotos := range votes {
		fmt.Printf("puntaje:%d - cantidad_de_votos:%d\n", puntaje, cantidadVotos)

		if puntaje <= 2 {
			votosNegativos += cantidadVotos
		}

		if puntaje >= 4 {
			votosPositivos += cantidadVotos
		}
	}

	fmt.Printf("Cantidad de votos positivos: %d\n", votosPositivos)
	fmt.Printf("Cantidad de votos negativos: %d\n", votosNegativos)

	if votosPositivos > votosNegativos {
		fmt.Println("Buen Resultado!!")
	} else {
		fmt.Println("Resultado mejorable :(")
	}
}
