package main

// Scrivere un programma che legga da standard input le misure dell’altezza e della base di un rettangolo e ne calcoli il perimetro e l’area.

import . "fmt"

func main() {
	var b, h float64 

	Print("Inserire base:")
	Scan(&b)
	Print("Inserire altezza:")
	Scan(&h)

	Println("Perimetro rettangolo:", 2*b+2*h)
	Println("Area rettangolo:", b*h)
}
