package main

import (
	"fmt"
	"strings"
	"slices"
)

func main() {
	separador := strings.Repeat("=", 15)
	var arregloCadenas []string
	fmt.Println("datos:", "contenido:",arregloCadenas,"condicion:",arregloCadenas == nil,"longitud:",len(arregloCadenas) == 0)
	fmt.Println(separador)

	arregloCadenas = make([]string, 3)
	arregloCadenas[0] = "Hola"
	arregloCadenas[1] = "Mundo"
	arregloCadenas[2] = "Golang"
	fmt.Println("datos:", "contenido:",arregloCadenas,"condicion:",arregloCadenas == nil,"longitud:",len(arregloCadenas) == 0)
	fmt.Println(separador)

	arregloCadenas = append(arregloCadenas, "!")
	fmt.Println("datos:", "contenido:",arregloCadenas,"condicion:",arregloCadenas == nil,"longitud:",len(arregloCadenas) == 0)
	fmt.Println(separador)

	segundoArreglo := []string{"Hola", "Mundo", "Golang"}
	fmt.Println("datos 2:", "contenido 2:",segundoArreglo,"condicion 2:",segundoArreglo == nil,"longitud 2:",len(segundoArreglo) == 0)
	tercerArreglo := []string{"Hola", "Mundo", "Golang"}
	if slices.Equal(segundoArreglo, tercerArreglo) {
		fmt.Println("Los arreglos son iguales")
	} else {
		fmt.Println("Los arreglos no son iguales")
	}
	fmt.Println(separador)
}
