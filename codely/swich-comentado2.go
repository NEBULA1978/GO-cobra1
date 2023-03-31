package main

import (
	"fmt"
)

func main() {
	// Ejemplo 1: switch con varios cases
	i := 2
	switch i {
	case 1:
		fmt.Println("i es igual a 1")
	case 2:
		fmt.Println("i es igual a 2")
	case 3:
		fmt.Println("i es igual a 3")
	default:
		fmt.Println("i no es igual a 1, 2 ni 3")
	}

	// Ejemplo 2: switch sin expresión
	j := 3
	switch {
	case j < 2:
		fmt.Println("j es menor que 2")
	case j < 5:
		fmt.Println("j es menor que 5")
	default:
		fmt.Println("j es mayor o igual a 5")
	}

	// Ejemplo 3: switch con una expresión de tipo string
	name := "Juan"
	switch name {
	case "Pedro":
		fmt.Println("¡Hola Pedro!")
	case "Maria":
		fmt.Println("¡Hola Maria!")
	default:
		fmt.Printf("¡Hola %s!\n", name)
	}
}

// La salida de este código será:

// i es igual a 2
// j es menor que 5
// ¡Hola Juan!
