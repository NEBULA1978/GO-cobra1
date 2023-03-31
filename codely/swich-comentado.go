package main

import (
	"fmt"
)

func main() {
	// Creamos una slice de enteros con algunos números.
	numbers := []int{1, 2, 3, 4}

	// Recorremos la slice usando un for loop y la función range().
	// La variable i contendrá el índice actual y la variable n contendrá el valor en ese índice.
	for i, n := range numbers {
		// Usamos un switch statement para comprobar si el número actual es par o impar.
		switch n % 2 {
		case 0:
			// Si el número es par, imprimimos un mensaje diciendo que es par y luego usamos la palabra clave "fallthrough"
			// para continuar ejecutando el siguiente case (en este caso, el case por defecto).
			fmt.Printf("El número en el índice %d es par y ", i)
			fallthrough
		default:
			// Imprimimos el número actual y su índice.
			fmt.Printf("el número %d está en el índice %d\n", n, i)
		}
	}
}

// En este ejemplo, la salida será:

// el número 1 está en el índice 0
// El número en el índice 1 es par y el número 2 está en el índice 1
// el número 3 está en el índice 2
// El número en el índice 3 es par y el número 4 está en el índice 3
