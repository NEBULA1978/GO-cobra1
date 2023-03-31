package main

import "fmt"

func main() {
	var nombre string
	var cuentaTwitter string

	fmt.Print("Introduce tu nombre: ")
	fmt.Scanln(&nombre)

	fmt.Print("Introduce tu cuenta de Twitter: ")
	fmt.Scanln(&cuentaTwitter)

	fmt.Printf("¡Hola! Soy %s, un modelo de lenguaje entrenado por OpenAI.\n", nombre)
	fmt.Printf("Puedes seguirme en Twitter como %s para conocer más sobre mis habilidades.\n", cuentaTwitter)
}

// Al ejecutar este programa, verás lo siguiente en la consola:
// Introduce tu nombre: Juan
// Introduce tu cuenta de Twitter: juancito23
// ¡Hola! Soy Juan, un modelo de lenguaje entrenado por OpenAI.
// Puedes seguirme en Twitter como juancito23 para conocer más sobre mis habilidades.
