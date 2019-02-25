package main

import "fmt"

func main() {
	fmt.Println(HolaMundo())
}

// Funciones publicas comienzan con mayusculas
func HolaMundo() string {
	return "Hola mundo desde go."
}
