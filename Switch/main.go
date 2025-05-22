package main

import (
	"fmt"
	"time"
)

func main() {
	mensajeBase := "Escribe %d como %s"
	i := 3
	var palabra string

	switch i {
	case 1:
		palabra = "Uno"
	case 2:
		palabra = "Dos"
	case 3:
		palabra = "Tres"
	default:
		palabra = "Otro"
	}
	fmt.Printf(mensajeBase+"\n", i, palabra)

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("A descansar :3 n_n")
	default:
		fmt.Println("A trabajar :D")
	}

	time := time.Now()
	fmt.Println("La hora actual es: ", time.Hour())

	switch {
	case time.Hour() > 12:
		fmt.Println("Es tarde")
	default:
		fmt.Println("Es otra hora")
	}
}
