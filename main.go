package main

import (
	"flag"
	"fmt"

	"github.com/jonhovd/funtemps/conv"
)

// Definerer flag-variablene i hoved-"scope"
var fahr float64
var cel float64
var kel float64
var out string
var funfacts string

// Bruker init (som anbefalt i dokumentasjonen) for å sikre at flagvariablene
// er initialisert.
func init() {

	/*
	   Her er eksempler på hvordan man implementerer parsing av flagg.
	   For eksempel, kommando
	       funtemps -F 0 -out C
	   skal returnere output: 0°F er -17.78°C
	*/

	// Definerer og initialiserer flagg-variablene
	flag.Float64Var(&fahr, "F", 0.0, "temperatur i grader fahrenheit")
	// Du må selv definere flag-variablene for "C" og "K"

	flag.StringVar(&out, "out", "C", "beregne temperatur i C - celsius, F - farhenheit, K- Kelvin")
	//flag.StringVar(&funfacts, "funfacts", "sun", "\"fun-facts\" om sun - Solen, luna - Månen og terra - Jorden")
	// Du må selv definere flag-variabelen for -t flagget, som bestemmer
	// hvilken temperaturskala skal brukes når funfacts skal vises

	//flag-variablen for C
	flag.Float64Var(&cel, "C", 0.0, "temperatur i grader celsius")
	//flag-variablen for K
	flag.Float64Var(&kel, "K", 0.0, "temperatur i grader Kelvin")

}

func main() {

	flag.Parse()

	/**
	    Her må logikken for flaggene og kall til funksjoner fra conv og funfacts
	    pakkene implementeres.

	    Det er anbefalt å sette opp en tabell med alle mulige kombinasjoner
	    av flagg. flag-pakken har funksjoner som man kan bruke for å teste
	    hvor mange flagg og argumenter er spesifisert på kommandolinje.

	        fmt.Println("len(flag.Args())", len(flag.Args()))
			    fmt.Println("flag.NFlag()", flag.NFlag())

	    Enkelte kombinasjoner skal ikke være gyldige og da må kontrollstrukturer
	    brukes for å utelukke ugyldige kombinasjoner:
	    -F, -C, -K kan ikke brukes samtidig
	    disse tre kan brukes med -out, men ikke med -funfacts
	    -funfacts kan brukes kun med -t
	    ...
	    Jobb deg gjennom alle tilfellene. Vær obs på at det er en del sjekk
	    implementert i flag-pakken og at den vil skrive ut "Usage" med
	    beskrivelsene av flagg-variablene, som angitt i parameter fire til
	    funksjonene Float64Var og StringVar
	*/

	if out == "C" && isFlagPassed("F") {
		celsius := conv.FarhenheitToCelsius(fahr)
		fmt.Printf("%.2f°F is %.2f°C\n", fahr, celsius)
	}

	if out == "F" && isFlagPassed("C") {
		fahrenheit := conv.CelsiusToFahrenheit(cel)
		fmt.Printf("%.2f°C is %.2f°F\n", cel, fahrenheit)
	}

	if out == "K" && isFlagPassed("C") {
		kelvin := conv.CelsiusToKelvin(cel)
		fmt.Printf("%.2f°C is %.2f°K\n", cel, kelvin)
	}

	if out == "C" && isFlagPassed("K") {
		celsius := conv.KelvinToCelsius(kel)
		fmt.Printf("%.2f°K is %.2f°C\n", kel, celsius)
	}

	if out == "F" && isFlagPassed("K") {
		fahrenheit := conv.KelvinToFarhenheit(kel)
		fmt.Printf("%.2f°K is %.2f°F\n", kel, fahrenheit)
	}

	if out == "K" && isFlagPassed("F") {
		kelvin := conv.FahrenheitToKelvin(fahr)
		fmt.Printf("%.2f°F is %.2f°K\n", fahr, kelvin)
	}

	// Her er noen eksempler du kan bruke i den manuelle testingen
	fmt.Println(fahr, out, funfacts)

	fmt.Println("len(flag.Args())", len(flag.Args()))
	fmt.Println("flag.NFlag()", flag.NFlag())

	fmt.Println(isFlagPassed("out"))

	// Eksempel på enkel logikk
	if out == "C" && isFlagPassed("F") {
		// Kalle opp funksjonen FahrenheitToCelsius(fahr), som da
		// skal returnere °C
		fmt.Println("0°F er -17.78°C")
	}
}

// Funksjonen sjekker om flagget er spesifisert på kommandolinje
// Du trenger ikke å bruke den, men den kan hjelpe med logikken
func isFlagPassed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}
