/*
	Scrivere un programma che legga da standard input un numero intero n e che, come mostrato nell'Esempio di esecuzione, 
	stampi a video un quadrato di n righe costituite ciascuna da n simboli intervallati da spazi, alternando fra loro righe 
	costituite solo da simboli * (asterisco) intervallati da spazi, righe costituite solo da simboli + (più) intervallati da spazi 
	e righe costituite solo da simboli o (lettera o) intervallati da spazi.

	$ go run quadrato_righe_alterne_2.go
	Inserisci un numero: 5
	* * * * *
	+ + + + +
	o o o o o
	* * * * *
	+ + + + +

*/

package main

import . "fmt"

func main() {
	var n int
	Print("Inserisci un numero: ")
	Scan(&n)

	for r := 0; r < n; r++ {
		for c := 0; c < n; c++ {
			switch r % 3 {
			case 0:
				Print("* ")
			case 1:
				Print("+ ")
			case 2:
				Print("o ")
			}
		}
		Println()
	}
}