package main

import (
	"fmt"
	"observer/display"
	"observer/station"
)

func main() {
	weatherData := station.NewWeatherData()

	_ = display.NewCurrentConditionDisplay(weatherData)
	_ = display.NewStatisticsDisplay(weatherData)
	_ = display.NewForecastDisplay(weatherData)
	_ = display.NewHeatIndexDisplay(weatherData)

	fmt.Println("=================== Notifying First Data ==================")
	weatherData.SetMeasurements(80, 65, 30.4)
	fmt.Println()

	fmt.Println("=================== Notifying Second Data =================")
	weatherData.SetMeasurements(82, 70, 29.2)
	fmt.Println()

	fmt.Println("=================== Notifying Third Data ==================")
	weatherData.SetMeasurements(78, 90, 29.2)
	fmt.Println()
}
