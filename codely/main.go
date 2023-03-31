// El ejecutable de nustra aplicacion
package main

// 3 Formas de importar librerias
import (
	"fmt"
	. "fmt"
	_ "log"
)

// Para wardar varias constantes
const (
	language = "language"
)

// Declaramos variable
var global = true

// Funcion principal
func main() {
	var a float64 = 2
	// Declaramos variable forma corta
	b := 0
	// Salida por consola
	fmt.Println(a)
	// Salida por consola con libreira importada nos ahorramos fmt
	Println(b)

}

/*
Comentario multiple
*/
