package main

import "fmt"

func main() {
	// Arrays / Listas
	var personas = [5]string{"Pepe", "Luis", "Juan", "Kike", "Javi"}
	fmt.Println(personas)

	fmt.Println(personas[0])

	// Asigno otro valor por indice
	personas[2] = "Juanito"
	fmt.Println(personas[2])

	fmt.Println(personas)

	// bucle for recorrer elementos de array
	for contador := 0; contador < len(personas); contador++ {

		// bloque codigo
		fmt.Println("-", personas[contador])
	}
}

/*
Comentario multiple
*/
