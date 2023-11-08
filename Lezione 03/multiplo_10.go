/*
	Scrivere un programma che legge da standard input un numero intero n e verifica se il numero è multiplo di 10.
	Suggerimento: per verificare se un numero sia multiplo di 10 potete utilizzare l'operatore % che calcola il resto della divisione tra interi.

*/

package main

import . "fmt"

func main() {
	var num int
	Print("Inserisci un numero:")
	Scan(&num)

	if num%10 == 0 {
		Println(num, "è un multiplo di 10")
	} else {
		Println(num, "non è un multiplo di 10")
	}
}