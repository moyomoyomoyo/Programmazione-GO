/*
	Scrivere un programma che legga da standard input una stringa senza spazi e, considerando solamente l’insieme delle lettere dell'alfabeto inglese, stampi

	il numero di lettere maiuscole;
	il numero di lettere minuscole;
	il numero di consonanti;
	il numero di vocali.
	Suggerimento: Le lettere dell'alfabeto inglese sono considerate nello standard US-ASCII (integrato nello standard Unicode). Per i caratteri considerati nello standard US-ASCII, il codice Unicode varia tra 0 e 127. In particolare, per le lettere dell’alfabeto inglese il codice Unicode varia negli intervalli:

	[65, 90] per le lettere MAIUSCOLE (con ‘A’=65 e ‘Z’=90)
	[97, 122] per le lettere minuscole (con ‘a’=97 e ‘z’=122)

	$ go run analisi.go
	Ciaoà
	Maiuscole: 1
	Minuscole: 3
	Vocali: 3
	Consonanti: 1

	$ go run analisi.go
	Certo!Sto,bene!ìììììì
	Maiuscole: 2
	Minuscole: 10
	Vocali: 5
	Consonanti: 7

	$ go run analisi.go
	aaAA
	Maiuscole: 2
	Minuscole: 2
	Vocali: 4
	Consonanti: 0
*/

package main

import (
	. "fmt"
	"unicode"
)

func main() {
	var s string
	Print("Insert a string: ")
	Scan(&s)

	var minuscole, maiuscole, consonanti, vocali int
	for _, r := range s {

		if unicode.IsLetter(r) && len(string(r)) == 1{
			if unicode.IsUpper(r){
				maiuscole++
			} else if unicode.IsLower(r) {
				minuscole++
			}
			
			if isVocal(r){
				vocali++
			} else {
				consonanti++
			}

		}
	}

	Println("Maiuscole:\t", maiuscole)
	Println("Minuscole:\t", minuscole)
	Println("Vocali:	\t", vocali)
	Println("Consonanti:\t", consonanti)
}

func isVocal(r rune) bool {
	switch r {
		case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
			return true
		}

	return false
}
