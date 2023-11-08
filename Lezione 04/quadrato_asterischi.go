/*
	Scrivere un programma che legga da standard input un numero intero n e stampi a video un quadrato di n asterischi intervallati da spazi come mostrato nell'Esempio di esecuzione.
	Suggerimento: potete utilizzare due cicli for annidati.

	$ go run quadratoasterischi.go
	Inserisci un numero: 3
	* * *
	* * *
	* * *
*/

package main

import . "fmt"

func main()  {
	var n int
	Print("Inserisci un numero: ")
	Scan(&n)

	for r := 0; r < n; r++ {
		for c := 0; c < n; c++ {
			Print("* ")
		}
		Println()
	}
}