package main

import "fmt"

func sumaTodo(valores ...int) int {
	resultado := 0
	for _, valor := range valores {
		resultado += valor
	}
	return resultado
}

func main() {
	numeros := make([]int, 5)

	for i := range numeros {
		fmt.Printf("Número %d:", i+1)
		fmt.Scanf("%d", &numeros[i])
		fmt.Println("— Registro guardado —")
	}

	fmt.Println("Suma total:", sumaTodo(numeros...))
}
