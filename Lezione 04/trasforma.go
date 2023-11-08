/*
	Scrivere un programma che legga da standard input una stringa e, considerando l’insieme delle lettere dell'alfabeto inglese,
	ristampi a video la stringa due volte: la prima volta in maiuscolo e la seconda volta in minuscolo.

	$ go run trasforma.go
	TestoDiProva!!!
	Testo maiuscolo: TESTODIPROVA!!!
	Testo minuscolo: testodiprova!!!

	$ go run trasforma.go
	Testo_Di_PrOvA....
	Testo maiuscolo: TESTO_DI_PROVA....
	Testo minuscolo: testo_di_prova....

*/

package main

import (
	. "fmt"
	"unicode"
)

func main() {
	var s string
	Print("Inserisci una stringa: ")
	Scan(&s)

	Println("Testo in maiuscolo")
	for _, r := range s {
		Print(string(unicode.ToUpper(r)))		
	}

	Println()

	Println("testo in minuscolo")
	for _, r := range s {
		Print(string(unicode.ToLower(r)))	
	}
	Println()

}