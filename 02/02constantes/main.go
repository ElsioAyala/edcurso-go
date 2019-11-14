package main

import "fmt"

func main() {
	// const nombre = "Pedro"
	// fmt.Println(nombre)

	var a int
	var b int8
	n := "Pedro"

	a = 121212
	b = 5

	c := a + int(b)
	fmt.Println(c)

	// Prioridad Aritmética
	// % () * / + -
	d := 6 + 6*(6-6)
	fmt.Println(d)

	fmt.Printf("Hola %s Cómo estás?\n", n)
	fmt.Printf("c es de tipo: %T", c)

	var nombre string
	var numero int
	var entiendes bool
	fmt.Println(nombre, numero, entiendes)
}
