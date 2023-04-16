package funtemps

import (
	"flag"
	"fmt"
	"main/Conv" 
	"math"
	"os"
	"strings"
)

var fahr float64
var cels float64
var kelv float64
var out string
var funfacts string

func init() {
	flag.Float64Var(&fahr, "F", 0.00, "temperatur i grader fahrenheit")
	flag.Float64Var(&cels, "C", 0.00, "temperatur i grader Celsius")
	flag.Float64Var(&kelv, "K", 0.00, "temperatur i grader Kelvin")
	flag.StringVar(&out, "out", "C", "beregne temperatur i C - celsius, F - farhenheit, K- Kelvin")

}

func main() {

	flag.Parse()
  var desi int
  arguments := strings.Split(os.Args[2], ".")
  fmt.Println(arguments)
  if len(arguments) <= 1 {
  fmt.Println("") 
  } else {
    desi = (len(arguments[1]))
  }

	if flag.NFlag() < 2 {
		fmt.Println("expected two flags eks. funtemps -C 100 -out F")
		os.Exit(1)
	}
	if isFlagPassed("out") == false {
		fmt.Println("You must include an out")
	}

	// logikk
	if out == "K" && isFlagPassed("C") {
		// celsius til kelvin
    result := math.Pow10(desi)
		answer := conv.CelsiusToKelvin(cels)
    answer = (math.Round(answer*float64(result))/float64(result))
	  test := fmt.Sprintf("%[1]g°C er %[2]g°K\n", cels, answer)
    fmt.Println(test)
	} else if out == "K" && isFlagPassed("F") {
		// fahrenheit to kelvin
    result := math.Pow10(desi)
		answer := conv.FahrenheitToKelvin(fahr)
    answer = (math.Round(answer*float64(result))/float64(result))
		test := fmt.Sprintf("%[1]g°F er %[2]g°K\n", fahr, answer)
    fmt.Println(test)
	} else if out == "F" && isFlagPassed("C") {
		// celsius til fahrenheit
    result := math.Pow10(desi)
		answer := conv.CelsiusToFahrenheit(cels)
    answer = (math.Round(answer*float64(result))/float64(result))
		test := fmt.Sprintf("%[1]g°C er %[2]g°F\n", cels, answer)
    fmt.Println(test)
	} else if out == "F" && isFlagPassed("K") {
		// kelvin til fahrenheit
     result := math.Pow10(desi)
		answer := conv.KelvinToFahrenheit(kelv)
    answer = (math.Round(answer*float64(result))/float64(result))
		test := fmt.Sprintf("%[1]g°K er %[2]g°F\n", kelv, answer)
    fmt.Println(test)
	} else if out == "C" && isFlagPassed("F") {
		// Kalle opp funksjonen FahrenheitToCelsius(fahr), som da
		// skal returnere °C
    result := math.Pow10(desi)
		answer := conv.FahrenheitToCelsius(fahr)
    answer = (math.Round(answer*float64(result))/float64(result))
		test := fmt.Sprintf("%[1]g°F er %[2]g°C\n", fahr, answer)
    fmt.Println(test)
	} else if out == "C" && isFlagPassed("K") {
		//Kelvin til celsius
    result := math.Pow10(desi)
		answer := conv.KelvinToCelsius(kelv)
    answer = (math.Round(answer*float64(result))/float64(result))
		test := fmt.Sprintf("%[1]g°K er %[2]g°C\n", kelv, answer)
    fmt.Println(test)
	} else {
    fmt.Println("Usage: ")
  }
}

func isFlagPassed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}

