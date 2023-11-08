/*
	Scrivere un programma che legge da standard input un numero intero e stampa "Fizz" se il numero è multiplo di 3, 
	"Buzz" se il numero è multiplo di 5, "Fizz Buzz" se è multiplo sia di 3 sia di 5, niente altrimenti.
	
*/

package main

import . "fmt"

func main() {
	var num int
	Print("Inserire un numero: ")
	Scan(&num)

	if num%3 == 0 {
		Println("Fizz")
	}
	if num%5 == 0 {
		Println("Buzz")
	}
	Println()
}