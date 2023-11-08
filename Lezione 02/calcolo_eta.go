/*
	Scrivere un programma che legga da standard input le età di due persone (espresse in anni) e calcoli:

	la somma delle età inserite;
	la media delle età inserite;
	la media delle età inserite arrotondata per difetto all'intero inferiore;
	la media delle età inserite arrotondata per eccesso all'intero superiore;
	la somma e la media delle età che le due persone avranno fra 10 anni.
	Suggerimento: la media arrotondata per difetto può essere calcolata usando la funzione math.Floor del package math nel seguente modo:

	var mediaArrotondataDifetto float64 = math.Floor(media)
	Similarmente, la media arrotondata per eccesso può essere calcolata usando la funzione math.Ceil nel seguente modo:

	var mediaArrotondataEccesso float64 = math.Ceil(media)
	
*/

package main

import (
	. "fmt"
	"math"
)

func main() {
	var y1, y2 float64

	Print("Inserire l'età della prima persona:")
	Scan(&y1)
	Print("Inserire l'età della prima persona:")
	Scan(&y2)

	Println("Somma età:", y1+y2)
	Println("Media età:", (y1+y2)/2)
	Println("La media delle età arrontondata per difetto all'intero inferiore:", math.Floor((y1 + y2)/2))
	Println("La media delle età arrontondata per eccesso all'intero superiore:", math.Ceil((y1+y2)/2))
	Println("Somma età tra 10 anni:", y1+y2+20)
	Println("Media età tra 10 anni:", (y1+y2+20)/2)
}