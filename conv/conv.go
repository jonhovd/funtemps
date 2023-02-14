package conv

//Konverteringsformlene for Celcius, Kelvin og Farhenheit:

// Konverterer Farhenheit til Celsius
// Celsius = (Farhrenheit - 32)*(5/9)
func FarhenheitToCelsius(fahr float64) float64 {
	var celsius = (fahr - 32) * 5 / 9
	return celsius
}

// Konverterer Celsius til Farhenheit
// Farhrenheit = Celsius*(9/5) + 32
func CelsiusToFahrenheit(cel float64) float64 {
	var farhenheit = cel*9/5 + 32
	return farhenheit
}

// Konverterer Kelvin til Farhenheit
// Farhrenheit = (Kelvin - 273.15)*(9/5) + 32
func KelvinToFarhenheit(kel float64) float64 {
	var farhenheit = (kel-273.15)*9/5 + 32
	return farhenheit
}

// Konverterer Farhenheit til Kelvin
// Kelvin = (Farhenheit - 32) * (5/9) + 273.15
func FahrenheitToKelvin(fahr float64) float64 {
	var kelvin = (fahr-32)*5/9 + 273.15
	return kelvin
}

// Konverterer Celsius til Kelvin
// Kelvin = Celsius + 273.15
func CelsiusToKelvin(cel float64) float64 {
	var kelvin = cel + 273.15
	return kelvin
}

// Konverterer Kelvin til Celsius
// Celsius = Kelvin - 273.15
func KelvinToCelsius(kel float64) float64 {
	var celsius = kel - 273.15
	return celsius
}
