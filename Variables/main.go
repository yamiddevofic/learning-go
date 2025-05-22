package main

import "fmt"

func main() {
	// Variables para cada tipo de dato
	var numero int = 1
	var numero2 int = 2
	var nombre string = "Yamid"
	var boolean bool = true
	var flotante float64 = 1.2 // 64 bits = 8 bytes = 15 decimales = más precisión
	var flotante2 float32 = 1.2 // 32 bits = 4 bytes = 6 decimales = menos precisión
	fruta := "Manzana" // Inferencia de tipos

	fmt.Println("numero: ", numero)
	fmt.Println("numero2: ", numero2)
	fmt.Println("nombre: ", nombre)
	fmt.Println("boolean: ", boolean)
	fmt.Println("flotante: ", flotante)
	fmt.Println("flotante2: ", flotante2)
	fmt.Println("fruta: ", fruta)

	//tipos de datos
	fmt.Printf("tipo de numero: %T\n", numero)
	fmt.Printf("tipo de numero2: %T\n", numero2)
	fmt.Printf("tipo de nombre: %T\n", nombre)
	fmt.Printf("tipo de boolean: %T\n", boolean)
	fmt.Printf("tipo de flotante: %T\n", flotante)
	fmt.Printf("tipo de flotante2: %T\n", flotante2)
	fmt.Printf("tipo de fruta: %T\n", fruta)
}
	