/*
	Scrivere un programma che legga da standard input un numero intero n e che, come mostrato nell'Esempio di esecuzione, stampi a video un quadrato di n righe costituite ciascuna da n simboli intervallati da spazi, alternando fra loro righe costituite solo da simboli * (asterisco) intervallati da spazi e righe costituite solo da simboli + (più) intervallati da spazi.
	Suggerimento: potete utilizzare due cicli for annidati ed utilizzare l'operatore % per distinguere le righe pari da quelle dispari.

	$ go run quadrato_righe_alterne_1.go
	Inserisci un numero: 5
	* * * * *
	+ + + + +
	* * * * *
	+ + + + +
	* * * * *

*/


package main

import . "fmt"

func main()  {
	var n int
	Print("Inserisci un numero: ")
	Scan(&n)

	for r := 0; r < n; r++ {
		for c := 0; c < n; c++ {
			if r%2 == 0 {
				Print("* ")
			} else {
				Print("+ ")
			}
		}
		Println()
	}
}