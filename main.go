package main

import "fmt"

func main() {
	// Misma linea
	fmt.Print("Hello World!")
	fmt.Println("Hello World2!")
	// Lineas separadas
	fmt.Println("Hello World!")
	fmt.Println("Hello World2!")

	// Variables
	var edad = 34
	var coche = "Mercedes"
	coche = "Toyota"

	// Salida de datos

	fmt.Println("Mi esdad es ", edad)
	// fmt.Println("Mi coche es ", coche)
	// %s solo 2parametros strings
	// %d el 2parametros enteros
	// fmt.Printf("%s %d", "Mi coche es ", coche)
	fmt.Printf("%s %s", "Mi coche es ", coche)

	// Entrada de datos
	var web string
	fmt.Println("Cual es tu sitio web?")
	fmt.Scan(&web)

	fmt.Println("Tu sitio web es:", web)

	// Condiciones
	var altura int
	fmt.Println("¿Cuanto mides?")
	fmt.Scan(&altura)

	if altura >= 185 {
		fmt.Println("Eres una persona alta")
		fmt.Println("Mides", altura, "cm")
	} else {
		fmt.Println("Eres una persona bajita")
		fmt.Println("Mides", altura, "cm")

	}
}

/*
Comentario multiple
*/
