package main

import (
	"play.ground/foo"
)

func main() {
	foo.Bar()
}
-- go.mod --
module play.ground
-- foo/foo.go --
package foo

import "fmt"

func Bar() {
	fmt.Println("This function lives in an another file!")
}
// El error "expected declaration, found '--'syntax" se produce cuando hay un error de sintaxis en el código Go. En este caso, parece que el archivo go.mod está vacío y no tiene ninguna declaración, lo cual es incorrecto. El archivo go.mod debe contener al menos una declaración de módulo, que se indica con la palabra clave module seguida del nombre del módulo.