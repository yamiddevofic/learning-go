package main

import (
	"fmt"
	"maps"
)

func main() {
	mapa := make(map[string]int)

	mapa["Amid"] = 5
	mapa["Horacio"] = 7

	fmt.Println(mapa)

	version1 := mapa["Amid"]
	version2 := mapa["Horacio"]
	fmt.Println(version1, version2)

	// longitud del mapa
	fmt.Println(len(mapa))

	dato, existe := mapa["Amid"]
	fmt.Println(dato, existe)

	_, existe = mapa["Amid"]
	fmt.Println(existe)

	dato, _ = mapa["Amid"]
	fmt.Println(dato)

	delete(mapa, "Amid")
	fmt.Println(mapa)

	clear(mapa)
	fmt.Println(mapa)

	nuevoMapa1 := map[string]int{"Juan": 1, "Pedro": 2}
	fmt.Println(nuevoMapa1)

	nuevoMapa2 := map[string]int{"Juan": 1, "Pedro": 2}
	fmt.Println(nuevoMapa2)

	if maps.Equal(nuevoMapa1, nuevoMapa2) {
		fmt.Println("Los mapas son iguales")
	} else {
		fmt.Println("Los mapas no son iguales")
	}
}