package dipositivo

import (
	"errors"
	"github.com/fatih/color"
	"math/rand"
)

// Dispositivo estructura principal
type Dispositivo struct {
	Nombre string
	Estado bool
}

// New instancia un nuevo Dispositivo
func New(n string) *Dispositivo {

	// genera un entero aleatorio
	r := rand.Int()

	return &Dispositivo{
		Nombre: n,

		// si es par estará prendido y si es impar apagado
		Estado: r%2 == 0,
	}
}

func (d *Dispositivo) Encender() error {
	if d.Estado {
		return errors.New("el dispositivo ya está encendido")
	}

	d.Estado = true
	return nil
}

func (d *Dispositivo) Apagar() error {
	if !d.Estado {
		return errors.New("el dispositivo está apagado")
	}

	d.Estado = false
	return nil
}

func (d *Dispositivo) EstadoActual() string {
	if d.Estado {
		color.Green("Dispositivo encendido")
		return "Dispositivo encendido"
	}

	color.Red("Dispositivo apagado")
	return "Dispositivo apagado"
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
