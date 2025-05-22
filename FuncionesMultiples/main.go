package main

import "fmt"

func valoresMultiples(partes ...string) (string, string, string, string) {
	return partes[0], partes[1], partes[2], partes[3]
}

func valoresMultiples2(nombre1, nombre2, apellido string) (map[string]string, map[string]string) {
	mapa1 := make(map[string]string)
	mapa2 := make(map[string]string)
	mapa1[nombre1] = apellido
	mapa2[nombre2] = apellido

	return mapa1, mapa2
}

func main() {
	nombreCompleto := make([] string, 4)

	for i := range nombreCompleto {
		fmt.Printf("Nombre %d:", i+1)
		fmt.Scanf("%s", &nombreCompleto[i])
		fmt.Println("— Registro guardado —")
	}

	fmt.Println("Nombre completo:", nombreCompleto)
	fmt.Println(valoresMultiples(nombreCompleto...))

	//=============================
	nombre1 := "Yamid"
	nombre2 := "Horacio"
	apellido := "Rodriguez"

	mapaPrimerNombre, mapaSegundoNombre := valoresMultiples2(nombre1, nombre2, apellido)

	fmt.Println(mapaPrimerNombre)
	fmt.Println(mapaSegundoNombre)
	
	nombre, _ := valoresMultiples2(nombre1, nombre2, apellido)
	fmt.Println("Nombre:", nombre)
}


