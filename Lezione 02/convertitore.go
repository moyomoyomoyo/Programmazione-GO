/*
	Scrivere un programma che legga da standard input una distanza in Km ed effettui la conversione di tale distanza in miglia (1 Km = 0.62137 mi).
*/

package main

import . "fmt"

func main() {
	var km float64
	mi := 0.62137

	Print("Inserire Km:")
	Scan(&km)
	Println(km, "km sono", km*mi, "miglia")
}