package display

import (
	"fmt"
	"observer/station"
)

type StatisticsDisplay struct {
	maxTemp     float32
	minTemp     float32
	tempSum     float32
	numReadings float32
}

func NewStatisticsDisplay(weatherData *station.WeatherData) *StatisticsDisplay {
	var statisticsDisplay StatisticsDisplay
	statisticsDisplay.maxTemp = 0
	statisticsDisplay.minTemp = 200
	statisticsDisplay.tempSum = 0
	weatherData.RegisterObserver(&statisticsDisplay)
	return &statisticsDisplay
}

func (s *StatisticsDisplay) Update(temp, humidity, pressure float32) {
	s.tempSum += temp
	s.numReadings += 1
	if temp > s.maxTemp {
		s.maxTemp = temp
	}
	if temp < s.minTemp {
		s.minTemp = temp
	}
	s.Display()
}

func (s *StatisticsDisplay) Display() {
	fmt.Printf("Avg/Max/Min temperature: %.2f/%.2f/%.2f\n", s.tempSum/s.numReadings, s.maxTemp, s.minTemp)
}
