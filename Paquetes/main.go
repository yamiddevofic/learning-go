package main

import (
	"fmt"
	"os"
)

func main() {
	envVar := os.Getenv("HOME")
	if envVar == "" {
		fmt.Println("No se encontro la variable de entorno")
		os.Exit(1)
	}
	fmt.Println("La variable de entorno HOME es: " + envVar)
	file, err := os.Create("ejemplo.txt")
	if err != nil {
		fmt.Printf("No se pudo crear el archivo ejemplo.txt %v\n", err)
		return
	}
	defer file.Close()
	fmt.Println("Archivo creado exitosamente")
}
