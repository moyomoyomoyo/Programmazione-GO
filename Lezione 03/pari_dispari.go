/*Scrivere un programma che legge da standard input un intero n e stampa a video se il numero è pari o dispari.*/

package main

import . "fmt"

func main() {
	var num int
	Print("Inserire un numero: ")
	Scan(&num)

	if num%2 == 0 {
		Println(num, "è pari")
	} else {
		Println(num, "è dispari")

	}
}