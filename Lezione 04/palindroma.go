/*
	Definizione: Una parola è palindroma se può essere letta normalmente, da sinistra verso destra, sia viceversa, cioè da destra verso sinistra.

	Scrivere un programma che:

	legga da standard input una stringa senza spazi;
	stampi a video il messaggio Palindroma nel caso in cui la stringa letta sia palindroma e Non palindroma altrimenti.

	$ go run palindroma.go
	anna
	Palindroma

	$ go run palindroma.go
	anni
	Non palindroma

	$ go run palindroma.go
	osso
	Palindroma

	$ go run palindroma.go
	èssè
	Palindroma

	$ go run palindroma.go
	èsse
	Non palindroma

*/

package main

import . "fmt"

func main () {
	var frase string
	Print("Inserisci una stringa: ")
	Scan(&frase)

	if isPalindrome(frase) {
		Println("è palindroma")
	} else {
		Println("non è palindroma")
	}

}

func isPalindrome(frase string) bool{
	var inversa string
	for _, s := range frase {
		inversa = string(s) + inversa
	}

	if frase == inversa {
		return true
	}

	return false
}