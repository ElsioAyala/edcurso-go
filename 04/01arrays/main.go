package main

import "fmt"

func main() {
	/*var nombres [3] string

	nombres[0] = "Daniel"
	nombres[1] = "Alvaro"
	nombres[2] = "Alexis"

	fmt.Println(nombres)*/

	nombres := [3]string{"Daniel", "Alvaro", "Alexys"}

	tamaño := len(nombres)
	fmt.Println("El tamaño del arreglo es de:", tamaño)

	fmt.Println(nombres)

	//fmt.Println(nombres[inico:final(excluyente)])
	fmt.Println(nombres[0:2])
}
