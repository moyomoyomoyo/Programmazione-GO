/*
	Scrivere un programma che legga da standard input il raggio di un cerchio e ne calcoli circonferenza e area.

	Suggerimento: l'area del cerchio si calcola facendo raggio x raggio x pi_greco, mentre la circonferenza facendo 2 x raggio x pi_greco. 
	Il valore numerico di pi_greco è memorizzato nella costante Pi del package math, a cui ci si può riferire con math.Pi.
	
*/

package main

import (
	. "fmt"
	"math"
)

const pi_greco = math.Pi

func main() {
	var r float64 //r = raggio

	Print("Inserire raggio del cerchio:")
	Scan(&r)

	Println("Perimetro cerchio:", 2*r*pi_greco)
	Println("Area cerchio:", r*r*pi_greco)
}