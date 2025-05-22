package main

import "fmt"

func sumaNumeros(numeros ... int){
	fmt.Println(numeros, " ")
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	fmt.Println("La suma total es: ", total)
}

func main() {
	fmt.Println("Variadicas")
	var nombre string = "Yamid"
	fmt.Println("Hola", nombre)
	
	var edad int = 22
	fmt.Println("Hola", nombre, "tienes", edad, "años")

	sumaNumeros(1, 2, 3, 4, 5)
	sumaNumeros(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	sumaNumeros()
	sumaNumeros(1)
	sumaNumeros(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15)
	sumaNumeros(10, 20, 30)

	numeros := []int{1, 2, 3, 4, 5}
	sumaNumeros(numeros...)
}
