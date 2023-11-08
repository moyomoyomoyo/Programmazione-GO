/*
	Scrivere un programma che legge da standard input un numero intero n (specificato senza segno se maggiore o uguale a 0) e stampi a video il numero con segno.

*/

package main

import . "fmt"

func main() {
	var num int
	Print("Inserisci un numero intero: ")
	Scan(&num)

	if num > 0 {
		Println("+", num)
	} else {
		Println(num)
	}
}