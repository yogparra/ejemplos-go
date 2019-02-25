package main

import "fmt"

func main() {
	fmt.Println(HolaMundo("Guillermo"))
}

// Funciones publicas comienzan con mayusculas

const saludo = "Hola, "

func HolaMundo(nombre string) string {
	if nombre == "" {
		nombre = "Mundo"
	}
	return saludo + nombre
}
