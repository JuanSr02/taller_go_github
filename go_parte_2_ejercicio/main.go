package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/fatih/color"
)

// Struct dispositivo
type Dispositivo struct {
	Nombre string
	Estado bool
}

// Interfaz controlable.
type Controlable interface {
	Encender() error
	Apagar() error
	EstadoActual() string
}

// Encender activa el dispositivo
func (d *Dispositivo) Encender() error {
	if d.Estado {
		return errors.New("ya está encendido")
	}
	d.Estado = true
	return nil
}

// Apagar desactiva el dispositivo
func (d *Dispositivo) Apagar() error {
	if !d.Estado {
		return errors.New("ya está apagado")
	}
	d.Estado = false
	return nil
}

// EstadoActual devuelve una cadena que describe el estado actual del dispositivo
func (d *Dispositivo) EstadoActual() string {
	estado := "apagado"
	if d.Estado {
		estado = "encendido"
	}
	return fmt.Sprintf("El dispositivo %s está %s", d.Nombre, estado)
}

func mostrarEstadoDispositivos(dispositivos []Controlable) {
	// Mostrar todos los dispositivos con colores
	fmt.Println("\nEstado actual de todos los dispositivos:")
	if len(dispositivos) == 0 {
		color.Yellow("No hay dispositivos registrados")
		return
	}
	for i, disp := range dispositivos {
		// Verificar si está encendido o apagado
		if strings.Contains(disp.EstadoActual(), "encendido") {
			color.Green("[%d] %s", i+1, disp.EstadoActual())
		} else {
			color.Red("[%d] %s", i+1, disp.EstadoActual())
		}
	}
}

// CrearNuevoDispositivo solicita al usuario un nombre y crea un nuevo dispositivo
func CrearNuevoDispositivo(scanner *bufio.Scanner) (*Dispositivo, error) {
	fmt.Print("Ingrese el nombre del dispositivo (máximo 10 caracteres): ")
	scanner.Scan()
	nombre := strings.TrimSpace(scanner.Text())

	if len(nombre) == 0 {
		return nil, errors.New("el nombre no puede estar vacío")
	}

	if len(nombre) > 10 {
		return nil, errors.New("el nombre no puede tener más de 10 caracteres")
	}

	return &Dispositivo{
		Nombre: nombre,
		Estado: false, // Por defecto, el dispositivo está apagado
	}, nil
}

// ControlarDispositivo permite al usuario seleccionar un dispositivo y cambiarlo de estado
func ControlarDispositivo(dispositivos []Controlable, scanner *bufio.Scanner) error {
	if len(dispositivos) == 0 {
		return errors.New("no hay dispositivos registrados")
	}

	mostrarEstadoDispositivos(dispositivos)

	fmt.Print("\nSeleccione el número del dispositivo que desea controlar: ")
	scanner.Scan()
	numStr := strings.TrimSpace(scanner.Text())

	num, err := strconv.Atoi(numStr)
	if err != nil {
		return errors.New("ingrese un número válido")
	}

	if num < 1 || num > len(dispositivos) {
		return fmt.Errorf("seleccione un número entre 1 y %d", len(dispositivos))
	}

	// Obtener el dispositivo seleccionado (arrays empiezan en 0)
	dispositivo := dispositivos[num-1]

	// Mostrar opciones de acción
	fmt.Println("\n¿Qué desea hacer con este dispositivo?")
	fmt.Println("1. Encender")
	fmt.Println("2. Apagar")
	fmt.Print("Seleccione una opción: ")

	scanner.Scan()
	opcionStr := strings.TrimSpace(scanner.Text())
	opcion, err := strconv.Atoi(opcionStr)
	if err != nil {
		return errors.New("ingrese un número válido")
	}

	switch opcion {
	case 1:
		err := dispositivo.Encender()
		if err != nil {
			return err
		}
		color.Green("Dispositivo encendido exitosamente")
	case 2:
		err := dispositivo.Apagar()
		if err != nil {
			return err
		}
		color.Red("Dispositivo apagado exitosamente")
	default:
		return errors.New("opción no válida")
	}

	return nil
}

// MostrarMenu imprime el menú de opciones
func MostrarMenu() {
	fmt.Println("\n===== SISTEMA DE CONTROL DE DISPOSITIVOS IOT =====")
	fmt.Println("1. Crear un nuevo dispositivo")
	fmt.Println("2. Controlar un dispositivo (encender/apagar)")
	fmt.Println("3. Ver todos los dispositivos")
	fmt.Println("4. Salir")
	fmt.Print("Seleccione una opción: ")
}

func main() {
	dispositivos := []Controlable{}
	scanner := bufio.NewScanner(os.Stdin)

	// Predefinir algunos dispositivos para facilitar las pruebas
	dispositivos = append(
		dispositivos,
		&Dispositivo{Nombre: "Lampara", Estado: false},
		&Dispositivo{Nombre: "Termostato", Estado: false},
	)

	fmt.Println("Bienvenido al Sistema de Control de Dispositivos IoT")

	for {
		MostrarMenu()

		scanner.Scan()
		opcionStr := strings.TrimSpace(scanner.Text())
		opcion, err := strconv.Atoi(opcionStr)

		if err != nil {
			color.Yellow("Por favor, ingrese un número válido")
			continue
		}

		switch opcion {
		case 1:
			// Crear nuevo dispositivo
			nuevoDispositivo, err := CrearNuevoDispositivo(scanner)
			if err != nil {
				color.Yellow("Error: %s", err)
			} else {
				dispositivos = append(dispositivos, nuevoDispositivo)
				color.Green("Dispositivo '%s' creado exitosamente", nuevoDispositivo.Nombre)
			}

		case 2:
			// Controlar un dispositivo
			err := ControlarDispositivo(dispositivos, scanner)
			if err != nil {
				color.Yellow("Error: %s", err)
			}

		case 3:
			// Ver todos los dispositivos
			mostrarEstadoDispositivos(dispositivos)

		case 4:
			// Salir
			fmt.Println("Gracias por usar el Sistema de Control de Dispositivos IoT. ¡Hasta pronto!")
			return

		default:
			color.Yellow("Opción no válida. Por favor, seleccione una opción del menú")
		}
	}
}

/*func main() {
	// Inicializar el generador de números aleatorios
	rand.Seed(time.Now().UnixNano())

	// Crear una lista de dispositivos
	dispositivos := []Controlable{
		&Dispositivo{Nombre: "Lampara", Estado: false},
		&Dispositivo{Nombre: "Termostato", Estado: false},
		&Dispositivo{Nombre: "TV", Estado: false},
		&Dispositivo{Nombre: "Cerradura", Estado: false},
		&Dispositivo{Nombre: "Ventilador", Estado: false},
	}

	// Encender o apagar dispositivos aleatoriamente
	fmt.Println("Cambiando estados de dispositivos aleatoriamente...")
	for i := range dispositivos {
		// 50% de probabilidad de encender o apagar
		if rand.Intn(2) == 1 {
			err := dispositivos[i].Encender()
			if err != nil {
				color.Yellow("Error al encender: %s", err)
			}
		} else {
			err := dispositivos[i].Apagar()
			if err != nil {
				color.Yellow("Error al apagar: %s", err)
			}
		}
	}



}*/
