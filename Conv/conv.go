package conv

// Konverterer Celsius til Fahrenheit
func CelsiusToFahrenheit(value float64) float64 {
	// legger inn formellen  Farhrenheit = Celsius*(9/5) + 32
	answer := value*1.8 + 32.00
	return answer
}

// Konverterer Celsius til Kelvin
func CelsiusToKelvin(value float64) float64 {
	// legger inn formellen kelvin = celsius + 273.15
	answer := value + 273.15
	return answer
}

// konverterer Kelvin til Fahrenheit
func KelvinToFahrenheit(value float64) float64 {
	// legger inn formellen (Kelvin - 273.15)*(9/5) + 32
	answer := (value-273.15)*1.8 + 32 //(K − 273.15) × 1.8 + 32
	return answer
}

// Konverterer Kelving til Celsius
func KelvinToCelsius(value float64) float64 {
	// legger inn formellen celsius = Kelvin - 273.15
	answer := value - 273.15
	return answer
}

// Konverterer Fahrenheit til Celsius
func FahrenheitToCelsius(value float64) float64 {
	// Legger inn formellen Celsius = (Farhrenheit - 32)/(5/9)
	answer := (value-32) / (1.8)
	return answer // (value - 32) * 5 / 9
}

// Konverterer Fahrenheit til kelvin
func FahrenheitToKelvin(value float64) float64 {
	// legger inn formellen kelvin = (Farhenheit - 32) * (5/9) + 273.15
	answer := (value-32)*1.8 + 273.15
	return answer
}



