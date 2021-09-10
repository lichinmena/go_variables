package main

import "fmt"

func main() {

	var nombre string
	nombre = "Luis Fernando"
	fmt.Println(nombre)

	//Imprime el tipo de dato
	var nombre2 = "Jose"
	fmt.Printf("El tipo de dato de Nombre2 es: %T", nombre2)
	fmt.Println()
	fmt.Printf("El valor de Nombre2 es: %s", nombre2)

	fmt.Println()

	nombre3 := "Daniel"
	fmt.Printf("El tipo de dato de Nombre3 es: %T", nombre3)
	fmt.Println()
	fmt.Printf("El tipo de dato es: %T y el valor es: %s", nombre3, nombre3)

	//Todas las variables que se declaran
	//deben ser usadas
	fmt.Println()
	var a1, a2 string
	fmt.Printf("a1 es %T y a2 es %T", a1, a2)
	fmt.Println()
	fmt.Printf("a1=%s\n", a1)
	fmt.Printf("a2=%s\n", a2)

	//Otra forma de declarar
	var (
		a3 string
		a4 string
		a5 int
	)
	fmt.Printf("a3=%s\n", a3)
	fmt.Printf("a4=%s\n", a4)
	fmt.Printf("a5=%d\n", a5)

	a6, a7, a8 := "Hola", "Mundo", 38
	fmt.Printf("a6=%s\n", a6)
	fmt.Printf("a7=%s\n", a7)
	fmt.Printf("a8=%d\n", a8)

}
