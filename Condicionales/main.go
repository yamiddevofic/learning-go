package main

import "fmt"

func main() {
	nombre := "Yamid"
	edad := 22
	
	if edad >= 18 {
		fmt.Println("Hola ", nombre, " puedes votar")
	} 

	if nombre == "Yamid" {
		fmt.Println("Hola ", nombre)
	} else {
		fmt.Println("Hola desconocido")
	}

	if 8%2 == 0 {
		fmt.Println("El numero es par")
	} else {
		fmt.Println("El numero es impar")
	}

	if numero := 0; numero > 0 {
		fmt.Println("El numero es positivo")
	} else if numero < 0 {
		fmt.Println("El numero es negativo")
	} else {
		fmt.Println("El numero es cero")
	}
}
