package main

import "fmt"

func main() {
	var numero int = 1
	numero2 := 2
	fmt.Println(numero, numero2)

	var numeroEntero = 1
	fmt.Println(numeroEntero)
	var numeroFlotante = 1.2
	fmt.Println(numeroFlotante)
	resultado := numeroFlotante + float64(numeroEntero)
	fmt.Println(resultado)

	var nombre string = "Yamid"
	fmt.Println(nombre)
	apellido := "Rodriguez"
	fmt.Println(apellido)
	nombreCompleto := nombre + " " + apellido
	fmt.Println(nombreCompleto)

	nombre1 := "Yamid"
	nombre2 := "Horacio"
	apellido1 := "Rodriguez"
	edad := 22
	mensaje := fmt.Sprintf("Hola mi nombre es %s %s %s y tengo %d años", nombre1, nombre2, apellido1, edad)
	fmt.Println(mensaje)
	main_dos()
}


func main_dos() {
	nombre := "Ana"
	puntos := 120
	mensajeBase := "Jugador: %s, Puntos: %d" // Una cadena con verbos de formato

	fmt.Println("--- Usando fmt.Printf ---")
	// Printf INTERPRETA los verbos de formato y los reemplaza con los valores de las variables.
	// No añade un salto de línea por defecto, así que lo agregamos con \n.
	fmt.Printf(mensajeBase+"\n", nombre, puntos)
	fmt.Printf("Nombre: %s, Puntos: %d\n", nombre, puntos)

	fmt.Println("\n--- Usando fmt.Println ---")
	// Println NO INTERPRETA los verbos de formato. Los imprime literalmente.

	// Caso 1: Pasamos la cadena con verbos y luego las variables como argumentos separados.
	// Println imprimirá la primera cadena literalmente, luego un espacio, luego el valor de 'nombre',
	// luego un espacio, luego el valor de 'puntos', y finalmente un salto de línea.
	fmt.Println("Intentando usar verbos (incorrecto con Println):", mensajeBase, nombre, puntos)

	// Caso 2: Pasamos solo la cadena con verbos.
	// Println la imprimirá literalmente, seguida de un salto de línea.
	fmt.Println("Solo la cadena con verbos:", mensajeBase)

	// Caso 3: La forma "correcta" de usar Println para un resultado similar (sin formato directo).
	// Pasamos cada parte como un argumento separado. Println añade espacios entre ellos.
	fmt.Println("Forma correcta con Println:", "Jugador:", nombre, ",", "Puntos:", puntos)

	// Comparación directa con la misma cadena:
	fmt.Println("\n--- Comparación directa ---")
	cadena := "El nombre es %s y la edad es %d."
	valorStr := "Luis"
	valorInt := 33

	fmt.Print("Printf: ") // Usamos Print para no añadir salto de línea antes de Printf
	fmt.Printf(cadena+"\n", valorStr, valorInt)

	fmt.Print("Println: ") // Usamos Print para no añadir salto de línea antes de Println
	fmt.Println(cadena, valorStr, valorInt) // Imprime la cadena literal, luego las variables
}