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
func FarhenheitToCelsius(fahr float64) float64 {
	var celsius = (fahr - 32) * 5 / 9
	return celsius
}

// Konverterer Celsius til Farhenheit
func CelsiusToFahrenheit(cel float64) float64 {
	var farhenheit = cel*9/5 + 32
	return farhenheit
}

// Konverterer Kelvin til Farhenheit
func KelvinToFarhenheit(kel float64) float64 {
	var fahr = (kel-273.15)*9/5 + 32
	return fahr
}

// Konverterer Farhenheit til Kelvin
func FahrenheitToKelvin(fahr float64) float64 {
	var kelvin = (fahr-32)*5/9 + 273.15
	return kelvin
}

// Konverterer Celsius til Kelvin
func CelsiusToKelvin(cel float64) float64 {
	var kelvin = cel + 273.15
	return kelvin
}

// Konverterer Kelvin til Celsius
func KelvinToCelsius(kel float64) float64 {
	var celsius = kel - 273.15
	return celsius
}
