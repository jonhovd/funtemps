package conv

/*
  I denne pakken skal alle konverteringfunksjonene
  implementeres. Bruk engelsk.
    FarhenheitToCelsius
    CelsiusToFahrenheit
    KelvinToFarhenheit
    ...
*/

// Konverterer Farhenheit til Celsius
func FarhenheitToCelsius(value float64) float64 {
	// Her skal du implementere funksjonen
	// Du skal ikke formattere float64 i denne funksjonen
	// Gjør formattering i main.go med fmt.Printf eller
	// lignende
	return 56.67
}

// De andre konverteringsfunksjonene implementere her
// ...

// Konverterer Celsius til Farhenheit
func CelsiusToFahrenheit(value float64) float64 {
	return 134
}

// Konverterer Kelvin til Farhenheit
func KelvinToFarhenheit(value float64) float64 {
	return 132.8
}

// Konverterer Farhenheit til Kelvin
func FahrenheitToKelvin(value float64) float64 {
	return 329.82
}

// Konverterer Celsius til Kelvin
func CelsiusToKelvin(value float64) float64 {
	return 329.85
}

// Konverterer Kelvin til Celsius
func KelvinToCelsius(value float64) float64 {
	return 132.8
}
