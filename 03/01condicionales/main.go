package main

import "fmt"

func main() {
	/*isvalid := true*/

	if edad, nombre := 15, "alvarito"; edad < 14 {
		fmt.Println(nombre, "Es un niño")
	} else if edad < 18 {
		fmt.Println(nombre, "Es un adolescente")
	} else if edad < 30 {
		fmt.Println(nombre, "Es un Adulto")
	} else {
		fmt.Println(nombre, "Es un adulto grandecito")
	}

	/*if isvalid {
		fmt.Println("Esto está en el bloque verdadero")
		if 5 < 10 {
			fmt.Println("5 es menor a 10")
		} else {
			fmt.Println("5 no es menor a 10")
		}
	} else {
		fmt.Println("Esto está en el bloque falso")
	}
	*/
}
