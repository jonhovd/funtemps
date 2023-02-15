package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/jonhovd/funtemps/conv"
)

// Definerer flag-variablene i hoved-"scope"
var fahr float64
var cel float64
var kel float64
var out string

//var funfacts string

// Bruker init (som anbefalt i dokumentasjonen) for å sikre at flagvariablene
// er initialisert.
func init() {

	//flag-variablen for F
	flag.Float64Var(&fahr, "F", 0.0, "temperatur i grader fahrenheit")
	//flag-variablen for C
	flag.Float64Var(&cel, "C", 0.0, "temperatur i grader celsius")
	//flag-variablen for K
	flag.Float64Var(&kel, "K", 0.0, "temperatur i grader Kelvin")
	//flag-variablen for out
	flag.StringVar(&out, "out", "C", "beregne temperatur i C - celsius, F - farhenheit, K- Kelvin")

	//flag.StringVar(&funfacts, "funfacts", "sun", "\"fun-facts\" om sun - Solen, luna - Månen og terra - Jorden")

	// Du må selv definere flag-variabelen for -t flagget, som bestemmer
	// hvilken temperaturskala skal brukes når funfacts skal vises

}

func main() {

	flag.Parse()

	// Logikken for flaggene og kall til funksjonene fra conv:
	if out == "C" && isFlagPassed("F") {
		var celsius = conv.FarhenheitToCelsius(fahr)
		//fmt.Printf("%#v°F is %.2f°C\n", fahr, celsius)
		fmt.Printf("%.2f°F is %s°C\n", fahr, formatNumber(celsius))
	}

	if out == "F" && isFlagPassed("C") {
		var fahrenheit = conv.CelsiusToFahrenheit(cel)
		//fmt.Printf("%#v°C is %.2f°F\n", cel, fahrenheit)
		fmt.Printf("%.2f°C is %s°F\n", cel, formatNumber(fahrenheit))
	}

	if out == "K" && isFlagPassed("C") {
		var kelvin = conv.CelsiusToKelvin(cel)
		//fmt.Printf("%#v°C is %.2f°K\n", cel, kelvin)
		fmt.Printf("%.2f°C is %s°K\n", cel, formatNumber(kelvin))
	}

	if out == "C" && isFlagPassed("K") {
		var celsius = conv.KelvinToCelsius(kel)
		//fmt.Printf("%#v°K is %.2f°C\n", kel, celsius)
		fmt.Printf("%.2f°K is %s°C\n", kel, formatNumber(celsius))
	}

	if out == "F" && isFlagPassed("K") {
		var fahrenheit = conv.KelvinToFarhenheit(kel)
		//fmt.Printf("%#v°K is %.2f°F\n", kel, fahrenheit)
		fmt.Printf("%.2f°K is %s°F\n", kel, formatNumber(fahrenheit))
	}

	if out == "K" && isFlagPassed("F") {
		var kelvin = conv.FahrenheitToKelvin(fahr)
		//fmt.Printf("%#v°F is %.2f°K\n", fahr, kelvin)
		fmt.Printf("%.2f°F is %s°K\n", fahr, formatNumber(kelvin))
	}

}

func formatNumber(number float64) string {
	numStr := fmt.Sprintf("%.2f", number)
	numParts := strings.Split(numStr, ".")
	intPart := numParts[0]
	intPartLen := len(intPart)

	// Remove trailing zeros after decimal point
	decPart := strings.TrimRight(numParts[1], "0")
	if decPart == "" {
		return intPart
	}

	// Format integer part with thousands separators
	if intPartLen <= 3 {
		return intPart + "." + decPart
	}

	start := intPartLen % 3
	if start == 0 {
		start = 3
	}

	var result string
	for i, digit := range intPart {
		if i == start {
			result += " "
			start += 3
		}
		result += string(digit)
	}

	return result + "." + decPart
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
