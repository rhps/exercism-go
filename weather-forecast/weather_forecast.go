// Package weather to forecast the weather.
package weather

// CurrentCondition is a variable for current condition.
var CurrentCondition string
// CurrentLocation is a variable for current location.
var CurrentLocation string

// Forecast function to forecast weather.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
