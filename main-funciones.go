package main

import "fmt"

func mostrarAltura(altura int) string {
	// Condiciones
	var resultado string

	if altura >= 185 {
		resultado = ("Eres una persona alta")
	} else {
		resultado = ("Eres una persona bajita")
	}
	return resultado
}

func main() {
	var miAltura int
	fmt.Println("¿Cuanto mides?")
	fmt.Scan(&miAltura)

	fmt.Println(mostrarAltura(miAltura))
	fmt.Println(mostrarAltura(178))
	fmt.Println(mostrarAltura(145))

	fmt.Println("Mides", miAltura, "cm")

}

/*
Comentario multiple
*/
