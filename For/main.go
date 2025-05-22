package main

import (
	"fmt"
	"strings"
)

func main() {
	separador := strings.Repeat("=", 15)
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i++
	}
	fmt.Println(separador)

	for numero := 0; numero < 3; numero++ {
		fmt.Println(numero)
	}
	fmt.Println(separador)

	for rango := range 5 {
		fmt.Println("Rango: ", rango)
	}
	fmt.Println(separador)
	for {
		fmt.Println("Infinite loop")
		break
	}
	fmt.Println(separador)
	for valor := range 6 {
		if valor%2 == 0 {
			continue
		}
		fmt.Println(valor)
	}
}
