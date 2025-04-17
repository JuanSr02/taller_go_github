package dipositivo

import (
	"errors"
	"github.com/fatih/color"
	"math/rand"
)

// dispositivo estructura principal
type dispositivo struct {
	nombre string
	estado bool
}

// New instancia un nuevo dispositivo
func New(n string) *dispositivo {

	// genera un entero aleatorio
	r := rand.Int()

	return &dispositivo{
		nombre: n,

		// si es par estará prendido y si es impar apagado
		estado: r%2 == 0,
	}
}

func (d *dispositivo) Encender() error {
	if d.estado {
		return errors.New("el dispositivo ya está encendido")
	}

	d.estado = true
	return nil
}

func (d *dispositivo) Apagar() error {
	if !d.estado {
		return errors.New("el dispositivo está apagado")
	}

	d.estado = false
	return nil
}

func (d *dispositivo) EstadoActual() string {
	if d.estado {
		color.Green("dispositivo encendido")
		return "dispositivo encendido"
	}

	color.Red("dispositivo apagado")
	return "dispositivo apagado"
}

// Controlable define el comportamiento de dispositivos.
type Controlable interface {
	// Encender enciende un device
	Encender() error

	// Apagar apaga un device
	Apagar() error

	// EstadoActual devuelve el estado actual e imprime un mensaje por consola
	EstadoActual() string
}
