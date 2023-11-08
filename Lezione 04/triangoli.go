/*
	Scrivere un programma che legga da standard input un intero n > 1 e stampi, utilizzando il carattere *, 
	il perimetro di due triangoli rettangoli con base e altezza di lunghezza n, posizionati come mostrato nell'Esempio d'esecuzione.

	$ go run triangoli.go
	2
	**
	 *
	 * 
	 **

	$ go run triangoli.go
	3
	***
	 **
	  *
	  *  
	  ** 
	  ***

	$ go run triangoli.go
	4
	****
	 * *
	  **
	   *
	   *   
	   **  
	   * * 
	   ****

	$ go run triangoli.go 6
	******
	 *   *
	  *  *
	   * *
	    **
	     *
	     *     
		 **    
		 * *   
		 *  *  
		 *   * 
		 ******
*/

package main

import . "fmt"

func main() {
	var n int
	Print("Inserisci un numero (maggiore di 1): ")
	Scan(&n)

	if n > 1 {

		for r := 0; r < n*2; r++ {

			for c := 0; c < n+(n-1); c++ {

				if c == n-1 {
					Print("* ")
				} else if c < n-1 && r < n-1 {
					if r == 0 || c == r {
						Print("* ")
					} else {
						Print("  ")
					}
				} else if c >= n && r > n {

					if r == 2*n-1 || c == r-1 {
						Print("* ")
					} else {
						Print("  ")
					}

				} else {
					Print("  ")
				}
			}
			Println("  ")
		}

	} else {
		Println("I told u, l'unica cosa che non dovevi fare l'hai fatta")
	}
}