/*
	Scrivere un unico programma che:

	legga da standard input un valore intero che specifica il tipo di conversione da effettuare:
	1: secondi (inseriti dall’utente) in ore
	2: secondi inseriti dall’utente in minuti
	3: minuti inseriti dall’utente in ore
	4: minuti inseriti dall’utente in secondi
	5: ore inserite dall’utente in secondi
	6: ore inserite dall’utente in minuti
	7: minuti inseriti dall’utente in giorni e ore
	8: minuti inseriti dall’utente in anni e giorni

	gestendo l’insertimento di un valore di scelta non compreso tra 1 e 8;
	legga da standard input un valore reale da convertire;

	stampi a video il valore convertito.

	$ go run conversioni.go 
	Scegli la conversione:
	1) secondi -> ore
	2) secondi -> minuti
	3) minuti -> ore
	4) minuti -> secondi
	5) ore -> secondi
	6) ore -> minuti
	7) minuti -> giorni e ore
	8) minuti -> anni e giorni
	: 8
	Inserisci il valore da convertire: 7200
	7200 minuti corrispondono a 0 anni e 5 giorni

	$ go run conversioni.go
	Scegli la conversione:
	1) secondi -> ore
	2) secondi -> minuti
	3) minuti -> ore
	4) minuti -> secondi
	5) ore -> secondi
	6) ore -> minuti
	7) minuti -> giorni e ore
	8) minuti -> anni e giorni
	: 1
	Inserisci il valore da convertire: 3618
	3618 secondi corrispondono a 1.005 ore

	$ go run conversioni.go
	Scegli la conversione:
	1) secondi -> ore
	2) secondi -> minuti
	3) minuti -> ore
	4) minuti -> secondi
	5) ore -> secondi
	6) ore -> minuti
	7) minuti -> giorni e ore
	8) minuti -> anni e giorni
	: 9
	Scelta errata

*/


package main

import . "fmt"

func main() {
	Printf("Inserisci l'operazione da effettuare\n")

	Println("1: Secondi -> ore")
	Println("2: Secondi -> minuti")
	Println("3: Minuti -> ore")
	Println("4: Minuti -> secondi")
	Println("5: Ore -> secondi")
	Println("6: Ore -> minuti")
	Println("7: Minuti -> giorni, ore")
	Println("8: Minuti -> anni, giorni")


	var scelta, val int
	Print(" : ")
	Scan(&scelta)
	if scelta < 1 || scelta > 8 {
		Println("Scelta non valida")
	} else {
		Print("Inserire il valore da convertire: ")
		Scan(&val)

		if scelta == 1 {
			Println(val, "secondi equivalgono a", float64(val)/3600, "ore")
		} else if scelta == 2 {
			Println(val, "secondi equivalgono a", float64(val)/60, "minuti")
		} else if scelta == 3 {
			Println(val, "minuti equivalgono a", float64(val)/60, "ore")
		} else if scelta == 4 {
			Println(val, "minuti equivalgono a", float64(val)*60, "secondi")
		} else if scelta == 5 {
			Println(val, "ore equivalgono a", float64(val)*3600, "secondi")
		} else if scelta == 6 {
			Println(val, "ore equivalgono a", float64(val)*60, "minuti")
		} else if scelta == 7 {
			var p int
			p = (val / 60) / 24
			Println(val, "minuti equivalgono a", (val/60)/24, "giorni e ", ((float64(val)/60)/24-float64(p))*24, "ore")
		} else if scelta == 8 {
			Println(val, "minuti equivalgono a", ((val/60)/24)/365, "anni e", (val/60)/24, "giorni")
		}
	}

}