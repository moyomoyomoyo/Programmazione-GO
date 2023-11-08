/*
	Scrivere un programma che prenda in ingresso un valore intero rappresentante una quantità di tempo in secondi 
	e restituisca il numero di ore, minuti e secondi corrspondenti

*/

package main

import . "fmt"

func main() {
	var h, m, s int
	Print("Inserire i secondi:")
	Scan(&s)

	m = s/60 

	if m%60 == 0 {
		h = m/60
	}

	Println(h, ":", m, ":", s)
}