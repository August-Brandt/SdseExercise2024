package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Temperature struct {
	Value float64
	Unit  string // 'C' for celcius and 'F' for fahrenheit
}

func celsiusToFahrenheit(temp *Temperature) *Temperature {
	return &Temperature{temp.Value*1.8 + 32, "F"}
}

func fahrenheitToCelsius(temp *Temperature) *Temperature {
	return &Temperature{(temp.Value - 32) / 1.8, "C"}
}

func main() {
	// Check for correct number of inputs
	if len(os.Args) != 3 {
		fmt.Println("Usage:\n\ttemperature [temperature] [unit]")
		fmt.Println("Example: temperature 100 C")
		os.Exit(1)
	}

	// Parse temperature value
	value, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		fmt.Println("Invalid input: Please enter a numerical value for the temperature")
		os.Exit(1)
	}
	unit := strings.ToUpper(os.Args[2])
	
	// Perform the convertion
	startTemp := &Temperature{value, unit} 
	var convertedTemp *Temperature
	switch startTemp.Unit {
	case "F":
		convertedTemp = fahrenheitToCelsius(startTemp)
	case "C":
		convertedTemp = celsiusToFahrenheit(startTemp)
	default:
		fmt.Println("Invalid input: Please enter 'F' for fahrenheit or 'C' for celcius")
		os.Exit(1)
	}

	// the '%.0f %s' is the format for a float with no decimals followed by a string 
	fmt.Printf("%.0f %s\n", convertedTemp.Value, convertedTemp.Unit)   
}
