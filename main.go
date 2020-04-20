package main

import (
	"fmt"
	"math"
	"strconv"

	"github.com/inancgumus/screen"
)

//UserName ... name of app user
var UserName string

func main() {
	fmt.Printf("\n  Hi, ich bin Florians erstes Go Programm. Wer bist DU? ")
	fmt.Scanln(&UserName)
	menu()
}

//ReadUserNumber ... read user number, repeat if input is not valid
func ReadUserNumber() (input int) {
	var nuIn string
	fmt.Scanln(&nuIn)
	input, err := strconv.Atoi(nuIn)
	if err != nil || input < 0 {
		fmt.Printf("  Du hast keine Zahl eingegeben. Bitte erneut eingeben. (Genauigkeit muss >= 0 sein): ")
		input = ReadUserNumber()
	}
	return
}

//Calculation ... our calculation feature
func Calculation() {
	screen.Clear()
	screen.MoveTopLeft()
	fmt.Printf("\n  ####################################################\n")
	fmt.Printf("  ###           Mini Taschenrechner                ###\n")
	fmt.Printf("  ####################################################\n\n")
	fmt.Printf("  Nenne mir eine Zahl: ")
	var n1 int
	n1 = ReadUserNumber()
	fmt.Printf("  Nenne mir eine zweite Zahl: ")
	var n2 int
	n2 = ReadUserNumber()
	fmt.Printf("  Wieviele Stellen nach dem Komma: ")
	var n3 int
	n3 = ReadUserNumber()
	fmt.Printf("\n  " + rpad("", "-", 60) + "\n")
	fmt.Print("  " + rpad("| Die Summe von "+strconv.Itoa(n1)+" und "+strconv.Itoa(n2)+" ist: "+strconv.Itoa(n1+n2), " ", 59) + "|\n")
	fmt.Print("  " + rpad("| Die Differenz von "+strconv.Itoa(n1)+" und "+strconv.Itoa(n2)+" ist: "+strconv.Itoa(n1-n2), " ", 59) + "|\n")
	fmt.Print("  " + rpad("| Das Produkt von "+strconv.Itoa(n1)+" und "+strconv.Itoa(n2)+" ist: "+strconv.Itoa(n1*n2), " ", 59) + "|\n")
	if n2 == 0 {
		fmt.Print("  " + rpad("| Der Quotient konnte nicht berechnet werden da als zweite Zahl 0 angegeben wurde\n", " ", 59) + "|\n")
	} else {
		fmt.Print("  " + rpad("| Der Quotient von "+strconv.Itoa(n1)+" und "+strconv.Itoa(n2)+" ist: "+fmt.Sprintf("%g", math.Round(float64(n1)/float64(n2)*math.Pow(10, float64(n3)))/math.Pow(10, float64(n3))), " ", 59) + "|\n")
	}
	fmt.Printf("  " + rpad("", "-", 60) + "\n")
	fmt.Printf("\n  Bitte <enter> drücken um fortzufahren...")
	var nuIn string
	fmt.Scanln(&nuIn)
	menu()
}

func menu() {
	screen.Clear()
	screen.MoveTopLeft()
	fmt.Printf("\n  ####################################################\n")
	fmt.Printf("  ###                   Hauptmenü                  ###\n")
	fmt.Printf("  ####################################################\n\n")
	fmt.Printf("  Berechnungsprogramm starten: (1)\n")
	fmt.Printf("  Programm beenden: (2)\n")
	fmt.Printf("\n  " + UserName + ", bitte wählen: ")
	var n1 int
	n1 = ReadUserNumber()
	if n1 == 1 {
		Calculation()
	}
}

func rpad(s string, pad string, plength int) string {
	for i := len(s); i < plength; i++ {
		s = s + pad
	}
	return s
}
