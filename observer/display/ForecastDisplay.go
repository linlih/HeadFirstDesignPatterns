package display

import (
	"fmt"
	"observer/station"
)

type ForecastDisplay struct {
	currentPressure float32
	lastPressure    float32
}

func NewForecastDisplay(weatherData *station.WeatherData) *ForecastDisplay {
	var forecastDisplay ForecastDisplay
	forecastDisplay.currentPressure = 29.92
	weatherData.RegisterObserver(&forecastDisplay)
	return &forecastDisplay
}

func (f *ForecastDisplay) Update(temp, humidity, pressure float32) {
	f.lastPressure = f.currentPressure
	f.currentPressure = pressure
	f.Display()
}

func (f *ForecastDisplay) Display() {
	fmt.Println("Forecast: ")
	if f.currentPressure > f.lastPressure {
		fmt.Println("Improving weather on the way!")
	} else if f.currentPressure == f.lastPressure {
		fmt.Println("More of the same")
	} else if f.currentPressure < f.lastPressure {
		fmt.Println("Watch out of cooler, rainy weather")
	}
}
