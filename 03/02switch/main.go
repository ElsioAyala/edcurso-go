package main

import "fmt"

func main() {

	/*var a uint8
	a = 1*/

	/*switch a {
	case 1:
		fmt.Println("El valor es 1")
	case 2:
		fmt.Println("El valor es 2")
	default:
		fmt.Println("El valor no es ni 1 ni 2")
	}*/

	/*switch a {
	case 1:
		fallthrough
	case 2:
		fallthrough
	case 3:
		fallthrough
	case 4:
		fallthrough
	case 5:
		fmt.Println("Estás entre semana")
	case 6:
		fallthrough
	case 7:
		fmt.Println("Estás en fin de semana :)")
	default:
		fmt.Println("No es un día valido")
	}*/

	/*switch a {
	case a == 1:
		fallthrough
	case 2:
		fallthrough
	case 3:
		fallthrough
	case 4:
		fallthrough
	case 5:
		fmt.Println("Estás entre semana")
	case 6:
		fallthrough
	case 7:
		fmt.Println("Estás en fin de semana :)")
	default:
		fmt.Println("No es un día valido")
	}*/

	switch a := 3; {
	case a >= 0 && a <= 5:
		fmt.Println("Están entre semana")
	case a == 6 || a == 7:
		fmt.Println("Estás en fin de semana :)")
	default:
		fmt.Println("No es un día valido")
	}

}
