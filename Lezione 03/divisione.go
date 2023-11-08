/*
	Scrivere un programma che legga da standard input due numeri interi a e b e calcoli il risultato della divisione a/b.
 	Se b è uguale a 0, il programma stampa Impossibile.

	$ go run divisione.go
	Inserisci due numeri:
	5 2
	Quoziente = 2.5

	$ go run divisione.go
	Inserisci due numeri:
	5 0
	Impossibile

*/

package main

import . "fmt"

func main() {
	var a, b int
	Print("Inserire un numero a: ")
	Scan(&a)

	Print("Inserire un numero b: ")
	Scan(&b)

	if b == 0 {
		Println("Impossibile")
	} else {
		Println("Quoziente:", float64(a)/float64(b))
	}
}