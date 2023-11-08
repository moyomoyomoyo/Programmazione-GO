/*
	Scrivere un programma che legga da standard input due numeri interi che chiameremo n e l e calcoli l’area di un poligono regolare con n lati di lunghezza l.
	Suggerimento: l'area di un poligono regolare può essere calcolata utilizzando le funzioni math.Pow (per il calcolo della potenza) e math.Tan (per il calcolo della tangente di un angolo) del package math nel seguente modo:

	var area float64 = (n * math.Pow(l, 2)) / (4 * math.Tan(math.Pi/n))

*/

package main

import . "fmt"
import math "math"

func main() {
	var n, l int
	Print("Inserisci il numero di lati del poligono:")
	Scan(&n)
	Print("Inserisci la lunghezza dei lati del poligono:")
	Scan(&l)

	Println("Area calcolata:", ((float64(n) * math.Pow(float64(l), 2)) / (4 * math.Tan(math.Pi/float64(n)))))
}