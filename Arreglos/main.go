package main

import (
	"fmt"
	"strings"
)

func main() {
	separador := strings.Repeat("=", 15)
	var arreglo [6]int
	fmt.Println("Arreglo creado: ", arreglo)
	arreglo[5] = 100
	arreglo[4] = 1
	fmt.Println("Arreglo modificado: ", arreglo)
	fmt.Printf("Tipo de arreglo: %T\n", arreglo)
	fmt.Printf("Longitud del arreglo: %d\n", len(arreglo))
	fmt.Println("Elemento 5: ", arreglo[5])
	fmt.Println(separador)

	arreglo2 := [6]int{1, 2, 3, 4, 5, 6}
	fmt.Println("Arreglo2: ", arreglo2)
	fmt.Println(separador)

	lista2 := [...] int{1, 2, 3, 4, 5, 6}
	fmt.Println("Lista inferida: ", lista2)
	fmt.Println("Longitud de lista2: ", len(lista2))
	fmt.Println(separador)
}
