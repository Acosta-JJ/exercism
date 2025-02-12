//Package weather aouaoueaoue.
package weather

//CurrentCondition defines the current codition.
var CurrentCondition string
//CurrentLocation defines the current Location.
var CurrentLocation string

//Forecast returns the CurrentLocation in CurrentCondition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
